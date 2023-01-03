// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package robomaker_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoboMakerFleetDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RoboMaker::Fleet", "awscc_robomaker_fleet", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.DataSourceWithEmptyResourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrPair(fmt.Sprintf("data.%s", td.ResourceName), "id", td.ResourceName, "id"),
				resource.TestCheckResourceAttrPair(fmt.Sprintf("data.%s", td.ResourceName), "arn", td.ResourceName, "arn"),
			),
		},
	})
}

func TestAccAWSRoboMakerFleetDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RoboMaker::Fleet", "awscc_robomaker_fleet", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
