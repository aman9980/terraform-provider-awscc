// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_iotsitewise_asset", assetResource)
}

// assetResource returns the Terraform awscc_iotsitewise_asset resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTSiteWise::Asset resource.
func assetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the asset",
		//	  "type": "string"
		//	}
		"asset_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the asset",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssetDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the asset",
		//	  "type": "string"
		//	}
		"asset_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the asset",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssetExternalId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The External ID of the asset",
		//	  "maxLength": 128,
		//	  "minLength": 2,
		//	  "pattern": "[a-zA-Z0-9_][a-zA-Z_\\-0-9.:]*[a-zA-Z0-9_]+",
		//	  "type": "string"
		//	}
		"asset_external_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The External ID of the asset",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(2, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9_][a-zA-Z_\\-0-9.:]*[a-zA-Z0-9_]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssetHierarchies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A hierarchy specifies allowed parent/child asset relationships.",
		//	    "properties": {
		//	      "ChildAssetId": {
		//	        "description": "The ID of the child asset to be associated.",
		//	        "type": "string"
		//	      },
		//	      "ExternalId": {
		//	        "description": "String-friendly customer provided external ID",
		//	        "maxLength": 128,
		//	        "minLength": 2,
		//	        "pattern": "[a-zA-Z0-9_][a-zA-Z_\\-0-9.:]*[a-zA-Z0-9_]+",
		//	        "type": "string"
		//	      },
		//	      "Id": {
		//	        "description": "Customer provided actual UUID for property",
		//	        "maxLength": 36,
		//	        "minLength": 36,
		//	        "pattern": "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$",
		//	        "type": "string"
		//	      },
		//	      "LogicalId": {
		//	        "description": "The LogicalID of a hierarchy in the parent asset's model.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ChildAssetId"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"asset_hierarchies": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ChildAssetId
					"child_asset_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the child asset to be associated.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: ExternalId
					"external_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "String-friendly customer provided external ID",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(2, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9_][a-zA-Z_\\-0-9.:]*[a-zA-Z0-9_]+"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Customer provided actual UUID for property",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(36, 36),
							stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LogicalId
					"logical_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The LogicalID of a hierarchy in the parent asset's model.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the asset",
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$",
		//	  "type": "string"
		//	}
		"asset_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the asset",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssetModelId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the asset model from which to create the asset.",
		//	  "type": "string"
		//	}
		"asset_model_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the asset model from which to create the asset.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique, friendly name for the asset.",
		//	  "type": "string"
		//	}
		"asset_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique, friendly name for the asset.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssetProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The asset property's definition, alias, unit, and notification state.",
		//	    "properties": {
		//	      "Alias": {
		//	        "description": "The property alias that identifies the property.",
		//	        "type": "string"
		//	      },
		//	      "ExternalId": {
		//	        "description": "String-friendly customer provided external ID",
		//	        "maxLength": 128,
		//	        "minLength": 2,
		//	        "pattern": "[a-zA-Z0-9_][a-zA-Z_\\-0-9.:]*[a-zA-Z0-9_]+",
		//	        "type": "string"
		//	      },
		//	      "Id": {
		//	        "description": "Customer provided actual UUID for property",
		//	        "maxLength": 36,
		//	        "minLength": 36,
		//	        "pattern": "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$",
		//	        "type": "string"
		//	      },
		//	      "LogicalId": {
		//	        "description": "Customer provided ID for property.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "NotificationState": {
		//	        "description": "The MQTT notification state (ENABLED or DISABLED) for this asset property.",
		//	        "enum": [
		//	          "ENABLED",
		//	          "DISABLED"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Unit": {
		//	        "description": "The unit of measure (such as Newtons or RPM) of the asset property. If you don't specify a value for this parameter, the service uses the value of the assetModelProperty in the asset model.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"asset_properties": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Alias
					"alias": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The property alias that identifies the property.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ExternalId
					"external_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "String-friendly customer provided external ID",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(2, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9_][a-zA-Z_\\-0-9.:]*[a-zA-Z0-9_]+"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Customer provided actual UUID for property",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(36, 36),
							stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LogicalId
					"logical_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Customer provided ID for property.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: NotificationState
					"notification_state": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The MQTT notification state (ENABLED or DISABLED) for this asset property.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"ENABLED",
								"DISABLED",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Unit
					"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The unit of measure (such as Newtons or RPM) of the asset property. If you don't specify a value for this parameter, the service uses the value of the assetModelProperty in the asset model.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the asset.",
		//	  "insertionOrder": false,
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
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the asset.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
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
		Description: "Resource schema for AWS::IoTSiteWise::Asset",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Asset").WithTerraformTypeName("awscc_iotsitewise_asset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":              "Alias",
		"asset_arn":          "AssetArn",
		"asset_description":  "AssetDescription",
		"asset_external_id":  "AssetExternalId",
		"asset_hierarchies":  "AssetHierarchies",
		"asset_id":           "AssetId",
		"asset_model_id":     "AssetModelId",
		"asset_name":         "AssetName",
		"asset_properties":   "AssetProperties",
		"child_asset_id":     "ChildAssetId",
		"external_id":        "ExternalId",
		"id":                 "Id",
		"key":                "Key",
		"logical_id":         "LogicalId",
		"notification_state": "NotificationState",
		"tags":               "Tags",
		"unit":               "Unit",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
