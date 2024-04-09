// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package config

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_config_conformance_pack", conformancePackResource)
}

// conformancePackResource returns the Terraform awscc_config_conformance_pack resource.
// This Terraform resource corresponds to the CloudFormation AWS::Config::ConformancePack resource.
func conformancePackResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConformancePackInputParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of ConformancePackInputParameter objects.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Input parameters in the form of key-value pairs for the conformance pack.",
		//	    "properties": {
		//	      "ParameterName": {
		//	        "description": "Key part of key-value pair with value being parameter value",
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      },
		//	      "ParameterValue": {
		//	        "description": "Value part of key-value pair with key being parameter Name",
		//	        "maxLength": 4096,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ParameterName",
		//	      "ParameterValue"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 60,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"conformance_pack_input_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ParameterName
					"parameter_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Key part of key-value pair with value being parameter value",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 255),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: ParameterValue
					"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Value part of key-value pair with key being parameter Name",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 4096),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of ConformancePackInputParameter objects.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 60),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConformancePackName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the conformance pack which will be assigned as the unique identifier.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z][-a-zA-Z0-9]*",
		//	  "type": "string"
		//	}
		"conformance_pack_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the conformance pack which will be assigned as the unique identifier.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z][-a-zA-Z0-9]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeliveryS3Bucket
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AWS Config stores intermediate files while processing conformance pack template.",
		//	  "maxLength": 63,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"delivery_s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AWS Config stores intermediate files while processing conformance pack template.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 63),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeliveryS3KeyPrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The prefix for delivery S3 bucket.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"delivery_s3_key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The prefix for delivery S3 bucket.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TemplateBody
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.",
		//	  "maxLength": 51200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"template_body": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 51200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// TemplateBody is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: TemplateS3Uri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": "s3://.*",
		//	  "type": "string"
		//	}
		"template_s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1024),
				stringvalidator.RegexMatches(regexp.MustCompile("s3://.*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// TemplateS3Uri is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: TemplateSSMDocumentDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.",
		//	  "properties": {
		//	    "DocumentName": {
		//	      "maxLength": 128,
		//	      "minLength": 3,
		//	      "type": "string"
		//	    },
		//	    "DocumentVersion": {
		//	      "maxLength": 128,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"template_ssm_document_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DocumentName
				"document_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(3, 128),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DocumentVersion
				"document_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 128),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// TemplateSSMDocumentDetails is a write-only property.
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
		Description: "A conformance pack is a collection of AWS Config rules and remediation actions that can be easily deployed as a single entity in an account and a region or across an entire AWS Organization.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::ConformancePack").WithTerraformTypeName("awscc_config_conformance_pack")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"conformance_pack_input_parameters": "ConformancePackInputParameters",
		"conformance_pack_name":             "ConformancePackName",
		"delivery_s3_bucket":                "DeliveryS3Bucket",
		"delivery_s3_key_prefix":            "DeliveryS3KeyPrefix",
		"document_name":                     "DocumentName",
		"document_version":                  "DocumentVersion",
		"parameter_name":                    "ParameterName",
		"parameter_value":                   "ParameterValue",
		"template_body":                     "TemplateBody",
		"template_s3_uri":                   "TemplateS3Uri",
		"template_ssm_document_details":     "TemplateSSMDocumentDetails",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/TemplateBody",
		"/properties/TemplateS3Uri",
		"/properties/TemplateSSMDocumentDetails",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
