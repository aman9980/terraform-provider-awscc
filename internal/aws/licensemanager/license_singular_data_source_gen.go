// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_licensemanager_license", licenseDataSource)
}

// licenseDataSource returns the Terraform awscc_licensemanager_license data source.
// This Terraform data source corresponds to the CloudFormation AWS::LicenseManager::License resource.
func licenseDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Beneficiary
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Beneficiary of the license.",
		//	  "type": "string"
		//	}
		"beneficiary": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Beneficiary of the license.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConsumptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "BorrowConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AllowEarlyCheckIn": {
		//	          "type": "boolean"
		//	        },
		//	        "MaxTimeToLiveInMinutes": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "MaxTimeToLiveInMinutes",
		//	        "AllowEarlyCheckIn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ProvisionalConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "MaxTimeToLiveInMinutes": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "MaxTimeToLiveInMinutes"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "RenewType": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"consumption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BorrowConfiguration
				"borrow_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AllowEarlyCheckIn
						"allow_early_check_in": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: MaxTimeToLiveInMinutes
						"max_time_to_live_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ProvisionalConfiguration
				"provisional_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: MaxTimeToLiveInMinutes
						"max_time_to_live_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RenewType
				"renew_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Entitlements
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "AllowCheckIn": {
		//	        "type": "boolean"
		//	      },
		//	      "MaxCount": {
		//	        "type": "integer"
		//	      },
		//	      "Name": {
		//	        "type": "string"
		//	      },
		//	      "Overage": {
		//	        "type": "boolean"
		//	      },
		//	      "Unit": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "Unit"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"entitlements": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AllowCheckIn
					"allow_check_in": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: MaxCount
					"max_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Overage
					"overage": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Unit
					"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: HomeRegion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Home region for the created license.",
		//	  "type": "string"
		//	}
		"home_region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Home region for the created license.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Issuer
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Name": {
		//	      "type": "string"
		//	    },
		//	    "SignKey": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name"
		//	  ],
		//	  "type": "object"
		//	}
		"issuer": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SignKey
				"sign_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LicenseArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name is a unique name for each resource.",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"license_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name is a unique name for each resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LicenseMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Name": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"license_metadata": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LicenseName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name for the created license.",
		//	  "type": "string"
		//	}
		"license_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name for the created license.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProductName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Product name for the created license.",
		//	  "type": "string"
		//	}
		"product_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Product name for the created license.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProductSKU
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ProductSKU of the license.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"product_sku": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ProductSKU of the license.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Validity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Begin": {
		//	      "description": "Validity begin date for the license.",
		//	      "format": "date-time",
		//	      "type": "string"
		//	    },
		//	    "End": {
		//	      "description": "Validity begin date for the license.",
		//	      "format": "date-time",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Begin",
		//	    "End"
		//	  ],
		//	  "type": "object"
		//	}
		"validity": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Begin
				"begin": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Validity begin date for the license.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: End
				"end": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Validity begin date for the license.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of the license.",
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of the license.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::LicenseManager::License",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LicenseManager::License").WithTerraformTypeName("awscc_licensemanager_license")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_check_in":              "AllowCheckIn",
		"allow_early_check_in":        "AllowEarlyCheckIn",
		"begin":                       "Begin",
		"beneficiary":                 "Beneficiary",
		"borrow_configuration":        "BorrowConfiguration",
		"consumption_configuration":   "ConsumptionConfiguration",
		"end":                         "End",
		"entitlements":                "Entitlements",
		"home_region":                 "HomeRegion",
		"issuer":                      "Issuer",
		"license_arn":                 "LicenseArn",
		"license_metadata":            "LicenseMetadata",
		"license_name":                "LicenseName",
		"max_count":                   "MaxCount",
		"max_time_to_live_in_minutes": "MaxTimeToLiveInMinutes",
		"name":                        "Name",
		"overage":                     "Overage",
		"product_name":                "ProductName",
		"product_sku":                 "ProductSKU",
		"provisional_configuration":   "ProvisionalConfiguration",
		"renew_type":                  "RenewType",
		"sign_key":                    "SignKey",
		"status":                      "Status",
		"unit":                        "Unit",
		"validity":                    "Validity",
		"value":                       "Value",
		"version":                     "Version",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
