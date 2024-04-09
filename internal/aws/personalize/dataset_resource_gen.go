// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package personalize

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_personalize_dataset", datasetResource)
}

// datasetResource returns the Terraform awscc_personalize_dataset resource.
// This Terraform resource corresponds to the CloudFormation AWS::Personalize::Dataset resource.
func datasetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DatasetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the dataset",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"dataset_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the dataset",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatasetGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the dataset group to add the dataset to",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"dataset_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the dataset group to add the dataset to",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatasetImportJob
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Initial DatasetImportJob for the created dataset",
		//	  "properties": {
		//	    "DataSource": {
		//	      "additionalProperties": false,
		//	      "description": "The Amazon S3 bucket that contains the training data to import.",
		//	      "properties": {
		//	        "DataLocation": {
		//	          "description": "The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.",
		//	          "maxLength": 256,
		//	          "pattern": "(s3|http|https)://.+",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "DatasetArn": {
		//	      "description": "The ARN of the dataset that receives the imported data",
		//	      "maxLength": 256,
		//	      "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	      "type": "string"
		//	    },
		//	    "DatasetImportJobArn": {
		//	      "description": "The ARN of the dataset import job",
		//	      "maxLength": 256,
		//	      "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	      "type": "string"
		//	    },
		//	    "JobName": {
		//	      "description": "The name for the dataset import job.",
		//	      "maxLength": 63,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
		//	      "type": "string"
		//	    },
		//	    "RoleArn": {
		//	      "description": "The ARN of the IAM role that has permissions to read from the Amazon S3 data source.",
		//	      "maxLength": 256,
		//	      "pattern": "arn:([a-z\\d-]+):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"dataset_import_job": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataSource
				"data_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DataLocation
						"data_location": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(256),
								stringvalidator.RegexMatches(regexp.MustCompile("(s3|http|https)://.+"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The Amazon S3 bucket that contains the training data to import.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DatasetArn
				"dataset_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the dataset that receives the imported data",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(256),
						stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DatasetImportJobArn
				"dataset_import_job_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the dataset import job",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(256),
						stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: JobName
				"job_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name for the dataset import job.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 63),
						stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9\\-_]*"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the IAM role that has permissions to read from the Amazon S3 data source.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(256),
						stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Initial DatasetImportJob for the created dataset",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatasetType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of dataset",
		//	  "enum": [
		//	    "Interactions",
		//	    "Items",
		//	    "Users"
		//	  ],
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"dataset_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of dataset",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
				stringvalidator.OneOf(
					"Interactions",
					"Items",
					"Users",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the dataset",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the dataset",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9\\-_]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SchemaArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the schema to associate with the dataset. The schema defines the dataset fields.",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"schema_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the schema to associate with the dataset. The schema defines the dataset fields.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::Personalize::Dataset.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Personalize::Dataset").WithTerraformTypeName("awscc_personalize_dataset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"data_location":          "DataLocation",
		"data_source":            "DataSource",
		"dataset_arn":            "DatasetArn",
		"dataset_group_arn":      "DatasetGroupArn",
		"dataset_import_job":     "DatasetImportJob",
		"dataset_import_job_arn": "DatasetImportJobArn",
		"dataset_type":           "DatasetType",
		"job_name":               "JobName",
		"name":                   "Name",
		"role_arn":               "RoleArn",
		"schema_arn":             "SchemaArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(2160).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
