// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkmanager_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSNetworkManagerTransitGatewayRegistrationDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::NetworkManager::TransitGatewayRegistration", "awscc_networkmanager_transit_gateway_registration", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSNetworkManagerTransitGatewayRegistrationDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::NetworkManager::TransitGatewayRegistration", "awscc_networkmanager_transit_gateway_registration", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
