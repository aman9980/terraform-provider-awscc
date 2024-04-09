// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ce

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ce_cost_category", costCategoryResource)
}

// costCategoryResource returns the Terraform awscc_ce_cost_category resource.
// This Terraform resource corresponds to the CloudFormation AWS::CE::CostCategory resource.
func costCategoryResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Cost category ARN",
		//	  "pattern": "^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Cost category ARN",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultValue
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default value for the cost category",
		//	  "maxLength": 50,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"default_value": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The default value for the cost category",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EffectiveStart
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ISO 8601 date time with offset format",
		//	  "maxLength": 25,
		//	  "minLength": 20,
		//	  "pattern": "^\\d{4}-\\d\\d-\\d\\dT\\d\\d:\\d\\d:\\d\\d(([+-]\\d\\d:\\d\\d)|Z)$",
		//	  "type": "string"
		//	}
		"effective_start": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ISO 8601 date time with offset format",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 50,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RuleVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "CostCategoryExpression.v1"
		//	  ],
		//	  "type": "string"
		//	}
		"rule_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CostCategoryExpression.v1",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Rules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "JSON array format of Expression in Billing and Cost Management API",
		//	  "type": "string"
		//	}
		"rules": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "JSON array format of Expression in Billing and Cost Management API",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: SplitChargeRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Json array format of CostCategorySplitChargeRule in Billing and Cost Management API",
		//	  "type": "string"
		//	}
		"split_charge_rules": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Json array format of CostCategorySplitChargeRule in Billing and Cost Management API",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Cost Category enables you to map your cost and usage into meaningful categories. You can use Cost Category to organize your costs using a rule-based engine.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::CostCategory").WithTerraformTypeName("awscc_ce_cost_category")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"default_value":      "DefaultValue",
		"effective_start":    "EffectiveStart",
		"name":               "Name",
		"rule_version":       "RuleVersion",
		"rules":              "Rules",
		"split_charge_rules": "SplitChargeRules",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
