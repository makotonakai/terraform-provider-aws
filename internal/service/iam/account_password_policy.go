package iam

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// @SDKResource("aws_iam_account_password_policy")
func ResourceAccountPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceAccountPasswordPolicyUpdate,
		ReadWithoutTimeout:   resourceAccountPasswordPolicyRead,
		UpdateWithoutTimeout: resourceAccountPasswordPolicyUpdate,
		DeleteWithoutTimeout: resourceAccountPasswordPolicyDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"allow_users_to_change_password": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"expire_passwords": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hard_expiry": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"max_password_age": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minimum_password_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"password_reuse_prevention": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"require_lowercase_characters": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"require_numbers": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"require_symbols": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"require_uppercase_characters": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAccountPasswordPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).IAMConn()

	input := &iam.UpdateAccountPasswordPolicyInput{}

	if v, ok := d.GetOk("allow_users_to_change_password"); ok {
		input.AllowUsersToChangePassword = aws.Bool(v.(bool))
	}
	if v, ok := d.GetOk("hard_expiry"); ok {
		input.HardExpiry = aws.Bool(v.(bool))
	}
	if v, ok := d.GetOk("max_password_age"); ok {
		input.MaxPasswordAge = aws.Int64(int64(v.(int)))
	}
	if v, ok := d.GetOk("minimum_password_length"); ok {
		input.MinimumPasswordLength = aws.Int64(int64(v.(int)))
	}
	if v, ok := d.GetOk("password_reuse_prevention"); ok {
		input.PasswordReusePrevention = aws.Int64(int64(v.(int)))
	}
	if v, ok := d.GetOk("require_lowercase_characters"); ok {
		input.RequireLowercaseCharacters = aws.Bool(v.(bool))
	}
	if v, ok := d.GetOk("require_numbers"); ok {
		input.RequireNumbers = aws.Bool(v.(bool))
	}
	if v, ok := d.GetOk("require_symbols"); ok {
		input.RequireSymbols = aws.Bool(v.(bool))
	}
	if v, ok := d.GetOk("require_uppercase_characters"); ok {
		input.RequireUppercaseCharacters = aws.Bool(v.(bool))
	}

	_, err := conn.UpdateAccountPasswordPolicyWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "updating IAM Account Password Policy: %s", err)
	}

	if d.IsNewResource() {
		d.SetId("iam-account-password-policy")
	}

	return append(diags, resourceAccountPasswordPolicyRead(ctx, d, meta)...)
}

func resourceAccountPasswordPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).IAMConn()

	policy, err := FindAccountPasswordPolicy(ctx, conn)

	if !d.IsNewResource() && tfawserr.ErrCodeEquals(err, iam.ErrCodeNoSuchEntityException) {
		log.Printf("[WARN] IAM Account Password Policy (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading IAM Account Password Policy (%s): %s", d.Id(), err)
	}

	d.Set("allow_users_to_change_password", policy.AllowUsersToChangePassword)
	d.Set("expire_passwords", policy.ExpirePasswords)
	d.Set("hard_expiry", policy.HardExpiry)
	d.Set("max_password_age", policy.MaxPasswordAge)
	d.Set("minimum_password_length", policy.MinimumPasswordLength)
	d.Set("password_reuse_prevention", policy.PasswordReusePrevention)
	d.Set("require_lowercase_characters", policy.RequireLowercaseCharacters)
	d.Set("require_numbers", policy.RequireNumbers)
	d.Set("require_symbols", policy.RequireSymbols)
	d.Set("require_uppercase_characters", policy.RequireUppercaseCharacters)

	return diags
}

func resourceAccountPasswordPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).IAMConn()

	log.Printf("[DEBUG] Deleting IAM Account Password Policy: %s", d.Id())
	_, err := conn.DeleteAccountPasswordPolicyWithContext(ctx, &iam.DeleteAccountPasswordPolicyInput{})

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting IAM Account Password Policy (%s): %s", d.Id(), err)
	}

	return diags
}

func FindAccountPasswordPolicy(ctx context.Context, conn *iam.IAM) (*iam.PasswordPolicy, error) {
	input := &iam.GetAccountPasswordPolicyInput{}

	output, err := conn.GetAccountPasswordPolicyWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, iam.ErrCodeNoSuchEntityException) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.PasswordPolicy == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output.PasswordPolicy, nil
}
