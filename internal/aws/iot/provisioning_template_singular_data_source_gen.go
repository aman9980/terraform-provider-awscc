// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_provisioning_template", provisioningTemplateDataSource)
}

// provisioningTemplateDataSource returns the Terraform awscc_iot_provisioning_template data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::ProvisioningTemplate resource.
func provisioningTemplateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 500,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Enabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PreProvisioningHook
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "PayloadVersion": {
		//	      "type": "string"
		//	    },
		//	    "TargetArn": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"pre_provisioning_hook": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PayloadVersion
				"payload_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TargetArn
				"target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ProvisioningRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"provisioning_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
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
		// Property: TemplateArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"template_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TemplateBody
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"template_body": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TemplateName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9A-Za-z_-]+$",
		//	  "type": "string"
		//	}
		"template_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TemplateType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "FLEET_PROVISIONING",
		//	    "JITP"
		//	  ],
		//	  "type": "string"
		//	}
		"template_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoT::ProvisioningTemplate",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::ProvisioningTemplate").WithTerraformTypeName("awscc_iot_provisioning_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":           "Description",
		"enabled":               "Enabled",
		"key":                   "Key",
		"payload_version":       "PayloadVersion",
		"pre_provisioning_hook": "PreProvisioningHook",
		"provisioning_role_arn": "ProvisioningRoleArn",
		"tags":                  "Tags",
		"target_arn":            "TargetArn",
		"template_arn":          "TemplateArn",
		"template_body":         "TemplateBody",
		"template_name":         "TemplateName",
		"template_type":         "TemplateType",
		"value":                 "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
