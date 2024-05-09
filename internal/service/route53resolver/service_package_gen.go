// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	route53resolver_sdkv1 "github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceEndpoint,
			TypeName: "aws_route53_resolver_endpoint",
		},
		{
			Factory:  DataSourceFirewallConfig,
			TypeName: "aws_route53_resolver_firewall_config",
		},
		{
			Factory:  DataSourceFirewallDomainList,
			TypeName: "aws_route53_resolver_firewall_domain_list",
		},
		{
			Factory:  DataSourceFirewallRuleGroup,
			TypeName: "aws_route53_resolver_firewall_rule_group",
		},
		{
			Factory:  DataSourceFirewallRuleGroupAssociation,
			TypeName: "aws_route53_resolver_firewall_rule_group_association",
		},
		{
			Factory:  DataSourceResolverFirewallRules,
			TypeName: "aws_route53_resolver_firewall_rules",
		},
		{
			Factory:  DataSourceQueryLogConfig,
			TypeName: "aws_route53_resolver_query_log_config",
		},
		{
			Factory:  DataSourceRule,
			TypeName: "aws_route53_resolver_rule",
		},
		{
			Factory:  DataSourceRules,
			TypeName: "aws_route53_resolver_rules",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceConfig,
			TypeName: "aws_route53_resolver_config",
		},
		{
			Factory:  ResourceDNSSECConfig,
			TypeName: "aws_route53_resolver_dnssec_config",
		},
		{
			Factory:  ResourceEndpoint,
			TypeName: "aws_route53_resolver_endpoint",
			Name:     "Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceFirewallConfig,
			TypeName: "aws_route53_resolver_firewall_config",
		},
		{
			Factory:  ResourceFirewallDomainList,
			TypeName: "aws_route53_resolver_firewall_domain_list",
			Name:     "Firewall Domain List",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceFirewallRule,
			TypeName: "aws_route53_resolver_firewall_rule",
		},
		{
			Factory:  ResourceFirewallRuleGroup,
			TypeName: "aws_route53_resolver_firewall_rule_group",
			Name:     "Firewall Rule Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceFirewallRuleGroupAssociation,
			TypeName: "aws_route53_resolver_firewall_rule_group_association",
			Name:     "Rule Group Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceQueryLogConfig,
			TypeName: "aws_route53_resolver_query_log_config",
			Name:     "Query Log Config",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceQueryLogConfigAssociation,
			TypeName: "aws_route53_resolver_query_log_config_association",
		},
		{
			Factory:  ResourceRule,
			TypeName: "aws_route53_resolver_rule",
			Name:     "Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceRuleAssociation,
			TypeName: "aws_route53_resolver_rule_association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Route53Resolver
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*route53resolver_sdkv1.Route53Resolver, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return route53resolver_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config[names.AttrEndpoint].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
