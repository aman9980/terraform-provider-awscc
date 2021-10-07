// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package amplify_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSAmplifyAppDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Amplify::App", "awscc_amplify_app", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSAmplifyAppDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Amplify::App", "awscc_amplify_app", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}