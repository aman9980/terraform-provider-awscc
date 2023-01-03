// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sagemaker_model_package_group", modelPackageGroupDataSource)
}

// modelPackageGroupDataSource returns the Terraform awscc_sagemaker_model_package_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::SageMaker::ModelPackageGroup resource.
func modelPackageGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time at which the model package group was created.",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time at which the model package group was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModelPackageGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the model package group.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "arn:.*",
		//	  "type": "string"
		//	}
		"model_package_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the model package group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModelPackageGroupDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the model package group.",
		//	  "maxLength": 1024,
		//	  "pattern": "[\\p{L}\\p{M}\\p{Z}\\p{S}\\p{N}\\p{P}]*",
		//	  "type": "string"
		//	}
		"model_package_group_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the model package group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModelPackageGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the model package group.",
		//	  "maxLength": 63,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"model_package_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the model package group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModelPackageGroupPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"model_package_group_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ModelPackageGroupStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of a modelpackage group job.",
		//	  "enum": [
		//	    "Pending",
		//	    "InProgress",
		//	    "Completed",
		//	    "Failed",
		//	    "Deleting",
		//	    "DeleteFailed"
		//	  ],
		//	  "type": "string"
		//	}
		"model_package_group_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of a modelpackage group job.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SageMaker::ModelPackageGroup",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::ModelPackageGroup").WithTerraformTypeName("awscc_sagemaker_model_package_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_time":                   "CreationTime",
		"key":                             "Key",
		"model_package_group_arn":         "ModelPackageGroupArn",
		"model_package_group_description": "ModelPackageGroupDescription",
		"model_package_group_name":        "ModelPackageGroupName",
		"model_package_group_policy":      "ModelPackageGroupPolicy",
		"model_package_group_status":      "ModelPackageGroupStatus",
		"tags":                            "Tags",
		"value":                           "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
