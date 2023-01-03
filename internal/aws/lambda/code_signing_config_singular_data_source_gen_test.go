// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lambda_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLambdaCodeSigningConfigDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Lambda::CodeSigningConfig", "awscc_lambda_code_signing_config", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSLambdaCodeSigningConfigDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Lambda::CodeSigningConfig", "awscc_lambda_code_signing_config", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
