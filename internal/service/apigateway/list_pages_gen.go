// Code generated by "internal/generate/listpages/main.go -ListOps=GetAuthorizers -Paginator=Position -AWSSDKVersion=2"; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

func getAuthorizersPages(ctx context.Context, conn *apigateway.Client, input *apigateway.GetAuthorizersInput, fn func(*apigateway.GetAuthorizersOutput, bool) bool) error {
	for {
		output, err := conn.GetAuthorizers(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.Position) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.Position = output.Position
	}
	return nil
}
