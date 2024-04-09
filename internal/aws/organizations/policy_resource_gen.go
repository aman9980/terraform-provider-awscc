// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package organizations

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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
	registry.AddResourceFactory("awscc_organizations_policy", policyResource)
}

// policyResource returns the Terraform awscc_organizations_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::Organizations::Policy resource.
func policyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of the Policy",
		//	  "pattern": "^arn:aws.*:organizations::[0-9]{12}:policy/o-[a-z0-9]{10}/(service_control|tag|backup|aiservices_opt_out)_policy/p-[a-z0-9]{8}",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of the Policy",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AwsManaged
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A boolean value that indicates whether the specified policy is an AWS managed policy. If true, then you can attach the policy to roots, OUs, or accounts, but you cannot edit it.",
		//	  "type": "boolean"
		//	}
		"aws_managed": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A boolean value that indicates whether the specified policy is an AWS managed policy. If true, then you can attach the policy to roots, OUs, or accounts, but you cannot edit it.",
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Content
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Policy text content. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.",
		//	  "maxLength": 1000000,
		//	  "minLength": 1,
		//	  "pattern": "[\\s\\S]*",
		//	  "type": "string"
		//	}
		"content": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Policy text content. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1000000),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\s\\S]*"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Human readable description of the policy",
		//	  "maxLength": 512,
		//	  "pattern": "[\\s\\S]*",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Human readable description of the policy",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(512),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\s\\S]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the Policy",
		//	  "maxLength": 130,
		//	  "pattern": "^p-[0-9a-zA-Z_]{8,128}$",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the Policy",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the Policy",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[\\s\\S]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the Policy",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\s\\S]*"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags that you want to attach to the newly created policy. For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A custom key-value pair associated with a resource within your organization.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key identifier, or name, of the tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "[\\s\\S]*",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The string value that's associated with the key of the tag. You can set the value of a tag to an empty string, but you can't set the value of a tag to null.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "[\\s\\S]*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
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
						Description: "The key identifier, or name, of the tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("[\\s\\S]*"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The string value that's associated with the key of the tag. You can set the value of a tag to an empty string, but you can't set the value of a tag to null.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("[\\s\\S]*"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags that you want to attach to the newly created policy. For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^(r-[0-9a-z]{4,32})|(\\d{12})|(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$",
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"target_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.ValueStringsAre(
					stringvalidator.RegexMatches(regexp.MustCompile("^(r-[0-9a-z]{4,32})|(\\d{12})|(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of policy to create. You can specify one of the following values: AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, TAG_POLICY",
		//	  "enum": [
		//	    "SERVICE_CONTROL_POLICY",
		//	    "AISERVICES_OPT_OUT_POLICY",
		//	    "BACKUP_POLICY",
		//	    "TAG_POLICY"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of policy to create. You can specify one of the following values: AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, TAG_POLICY",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SERVICE_CONTROL_POLICY",
					"AISERVICES_OPT_OUT_POLICY",
					"BACKUP_POLICY",
					"TAG_POLICY",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Policies in AWS Organizations enable you to manage different features of the AWS accounts in your organization.  You can use policies when all features are enabled in your organization.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Organizations::Policy").WithTerraformTypeName("awscc_organizations_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":         "Arn",
		"aws_managed": "AwsManaged",
		"content":     "Content",
		"description": "Description",
		"id":          "Id",
		"key":         "Key",
		"name":        "Name",
		"tags":        "Tags",
		"target_ids":  "TargetIds",
		"type":        "Type",
		"value":       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
