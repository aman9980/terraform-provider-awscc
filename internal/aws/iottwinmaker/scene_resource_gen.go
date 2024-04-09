// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
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
	registry.AddResourceFactory("awscc_iottwinmaker_scene", sceneResource)
}

// sceneResource returns the Terraform awscc_iottwinmaker_scene resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTTwinMaker::Scene resource.
func sceneResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the scene.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "arn:((aws)|(aws-cn)|(aws-us-gov)):iottwinmaker:[a-z0-9-]+:[0-9]{12}:[\\/a-zA-Z0-9_\\-\\.:]+",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the scene.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Capabilities
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of capabilities that the scene uses to render.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 256,
		//	    "minLength": 0,
		//	    "pattern": ".*",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"capabilities": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of capabilities that the scene uses to render.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 50),
				setvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(0, 256),
					stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContentLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The relative path that specifies the location of the content definition file.",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "pattern": "[sS]3://[A-Za-z0-9._/-]+",
		//	  "type": "string"
		//	}
		"content_location": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The relative path that specifies the location of the content definition file.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("[sS]3://[A-Za-z0-9._/-]+"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationDateTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time when the scene was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_date_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The date and time when the scene was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the scene.",
		//	  "maxLength": 512,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the scene.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 512),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GeneratedSceneMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A key-value pair of generated scene metadata for the scene.",
		//	  "maxLength": 50,
		//	  "minLength": 0,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 2048,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"generated_scene_metadata": // Pattern: ""
		schema.MapAttribute{        /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value pair of generated scene metadata for the scene.",
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SceneId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the scene.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z_0-9][a-zA-Z_\\-0-9]*[a-zA-Z0-9]+",
		//	  "type": "string"
		//	}
		"scene_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the scene.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z_0-9][a-zA-Z_\\-0-9]*[a-zA-Z0-9]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SceneMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A key-value pair of scene metadata for the scene.",
		//	  "maxLength": 50,
		//	  "minLength": 0,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 2048,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"scene_metadata":    // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value pair of scene metadata for the scene.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A key-value pair to associate with a resource.",
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value pair to associate with a resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdateDateTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time of the current update.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"update_date_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The date and time of the current update.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WorkspaceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the scene.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z_0-9][a-zA-Z_\\-0-9]*[a-zA-Z0-9]+",
		//	  "type": "string"
		//	}
		"workspace_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the scene.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z_0-9][a-zA-Z_\\-0-9]*[a-zA-Z0-9]+"), ""),
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
		Description: "Resource schema for AWS::IoTTwinMaker::Scene",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTTwinMaker::Scene").WithTerraformTypeName("awscc_iottwinmaker_scene")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                      "Arn",
		"capabilities":             "Capabilities",
		"content_location":         "ContentLocation",
		"creation_date_time":       "CreationDateTime",
		"description":              "Description",
		"generated_scene_metadata": "GeneratedSceneMetadata",
		"scene_id":                 "SceneId",
		"scene_metadata":           "SceneMetadata",
		"tags":                     "Tags",
		"update_date_time":         "UpdateDateTime",
		"workspace_id":             "WorkspaceId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
