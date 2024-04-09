// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package controltower

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
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_controltower_enabled_control", enabledControlResource)
}

// enabledControlResource returns the Terraform awscc_controltower_enabled_control resource.
// This Terraform resource corresponds to the CloudFormation AWS::ControlTower::EnabledControl resource.
func enabledControlResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ControlIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn of the control.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:aws[0-9a-zA-Z_\\-:\\/]+$",
		//	  "type": "string"
		//	}
		"control_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn of the control.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[0-9a-zA-Z_\\-:\\/]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn for Organizational unit to which the control needs to be applied",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:aws[0-9a-zA-Z_\\-:\\/]+$",
		//	  "type": "string"
		//	}
		"target_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn for Organizational unit to which the control needs to be applied",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[0-9a-zA-Z_\\-:\\/]+$"), ""),
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
		Description: "Enables a control on a specified target.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ControlTower::EnabledControl").WithTerraformTypeName("awscc_controltower_enabled_control")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"control_identifier": "ControlIdentifier",
		"target_identifier":  "TargetIdentifier",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
