// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigatewayv2_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSApiGatewayV2AuthorizerDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApiGatewayV2::Authorizer", "awscc_apigatewayv2_authorizer", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSApiGatewayV2AuthorizerDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApiGatewayV2::Authorizer", "awscc_apigatewayv2_authorizer", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
