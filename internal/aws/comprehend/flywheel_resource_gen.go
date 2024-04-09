// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package comprehend

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_comprehend_flywheel", flywheelResource)
}

// flywheelResource returns the Terraform awscc_comprehend_flywheel resource.
// This Terraform resource corresponds to the CloudFormation AWS::Comprehend::Flywheel resource.
func flywheelResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ActiveModelArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "pattern": "arn:aws(-[^:]+)?:comprehend:[a-zA-Z0-9-]*:[0-9]{12}:(document-classifier|entity-recognizer)/[a-zA-Z0-9](-*[a-zA-Z0-9])*(/version/[a-zA-Z0-9](-*[a-zA-Z0-9])*)?",
		//	  "type": "string"
		//	}
		"active_model_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:aws(-[^:]+)?:comprehend:[a-zA-Z0-9-]*:[0-9]{12}:(document-classifier|entity-recognizer)/[a-zA-Z0-9](-*[a-zA-Z0-9])*(/version/[a-zA-Z0-9](-*[a-zA-Z0-9])*)?"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "arn:aws(-[^:]+)?:comprehend:[a-zA-Z0-9-]*:[0-9]{12}:flywheel/[a-zA-Z0-9](-*[a-zA-Z0-9])*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataAccessRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "arn:aws(-[^:]+)?:iam::[0-9]{12}:role/.+",
		//	  "type": "string"
		//	}
		"data_access_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:aws(-[^:]+)?:iam::[0-9]{12}:role/.+"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: DataLakeS3Uri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 512,
		//	  "pattern": "s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?",
		//	  "type": "string"
		//	}
		"data_lake_s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(512),
				stringvalidator.RegexMatches(regexp.MustCompile("s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataSecurityConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DataLakeKmsKeyId": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "ModelKmsKeyId": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "VolumeKmsKeyId": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "VpcConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "SecurityGroupIds": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 32,
		//	            "minLength": 1,
		//	            "pattern": "[-0-9a-zA-Z]+",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 5,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "Subnets": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 32,
		//	            "minLength": 1,
		//	            "pattern": "[-0-9a-zA-Z]+",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 16,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "required": [
		//	        "SecurityGroupIds",
		//	        "Subnets"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"data_security_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataLakeKmsKeyId
				"data_lake_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 2048),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ModelKmsKeyId
				"model_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 2048),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: VolumeKmsKeyId
				"volume_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 2048),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: VpcConfig
				"vpc_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: SecurityGroupIds
						"security_group_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Required:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeBetween(1, 5),
								setvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 32),
									stringvalidator.RegexMatches(regexp.MustCompile("[-0-9a-zA-Z]+"), ""),
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Subnets
						"subnets": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Required:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeBetween(1, 16),
								setvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 32),
									stringvalidator.RegexMatches(regexp.MustCompile("[-0-9a-zA-Z]+"), ""),
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FlywheelName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"flywheel_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ModelType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "DOCUMENT_CLASSIFIER",
		//	    "ENTITY_RECOGNIZER"
		//	  ],
		//	  "type": "string"
		//	}
		"model_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"DOCUMENT_CLASSIFIER",
					"ENTITY_RECOGNIZER",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TaskConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DocumentClassificationConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Labels": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 5000,
		//	            "type": "string"
		//	          },
		//	          "maxItems": 1000,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "Mode": {
		//	          "enum": [
		//	            "MULTI_CLASS",
		//	            "MULTI_LABEL"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Mode"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "EntityRecognitionConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "EntityTypes": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Type": {
		//	                "maxLength": 64,
		//	                "minLength": 1,
		//	                "pattern": "",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Type"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "maxItems": 25,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "LanguageCode": {
		//	      "enum": [
		//	        "en",
		//	        "es",
		//	        "fr",
		//	        "it",
		//	        "de",
		//	        "pt"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "LanguageCode"
		//	  ],
		//	  "type": "object"
		//	}
		"task_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DocumentClassificationConfig
				"document_classification_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Labels
						"labels": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeAtMost(1000),
								setvalidator.ValueStringsAre(
									stringvalidator.LengthAtMost(5000),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
								setplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Mode
						"mode": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"MULTI_CLASS",
									"MULTI_LABEL",
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EntityRecognitionConfig
				"entity_recognition_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EntityTypes
						"entity_types": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Type
									"type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Required: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 64),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Optional: true,
							Computed: true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeBetween(1, 25),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
								setplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LanguageCode
				"language_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"en",
							"es",
							"fr",
							"it",
							"de",
							"pt",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
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
		Description: "The AWS::Comprehend::Flywheel resource creates an Amazon Comprehend Flywheel that enables customer to train their model.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Comprehend::Flywheel").WithTerraformTypeName("awscc_comprehend_flywheel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"active_model_arn":               "ActiveModelArn",
		"arn":                            "Arn",
		"data_access_role_arn":           "DataAccessRoleArn",
		"data_lake_kms_key_id":           "DataLakeKmsKeyId",
		"data_lake_s3_uri":               "DataLakeS3Uri",
		"data_security_config":           "DataSecurityConfig",
		"document_classification_config": "DocumentClassificationConfig",
		"entity_recognition_config":      "EntityRecognitionConfig",
		"entity_types":                   "EntityTypes",
		"flywheel_name":                  "FlywheelName",
		"key":                            "Key",
		"labels":                         "Labels",
		"language_code":                  "LanguageCode",
		"mode":                           "Mode",
		"model_kms_key_id":               "ModelKmsKeyId",
		"model_type":                     "ModelType",
		"security_group_ids":             "SecurityGroupIds",
		"subnets":                        "Subnets",
		"tags":                           "Tags",
		"task_config":                    "TaskConfig",
		"type":                           "Type",
		"value":                          "Value",
		"volume_kms_key_id":              "VolumeKmsKeyId",
		"vpc_config":                     "VpcConfig",
	})

	opts = opts.WithCreateTimeoutInMinutes(240).WithDeleteTimeoutInMinutes(120)

	opts = opts.WithUpdateTimeoutInMinutes(10)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
