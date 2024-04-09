// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediatailor

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_mediatailor_channel", channelResource)
}

// channelResource returns the Terraform awscc_mediatailor_channel resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaTailor::Channel resource.
func channelResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe ARN of the channel.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The ARN of the channel.</p>",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ChannelName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"channel_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FillerSlate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "\u003cp\u003eSlate VOD source configuration.\u003c/p\u003e",
		//	  "properties": {
		//	    "SourceLocationName": {
		//	      "description": "\u003cp\u003eThe name of the source location where the slate VOD source is stored.\u003c/p\u003e",
		//	      "type": "string"
		//	    },
		//	    "VodSourceName": {
		//	      "description": "\u003cp\u003eThe slate VOD source name. The VOD source must already exist in a source location before it can be used for slate.\u003c/p\u003e",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"filler_slate": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SourceLocationName
				"source_location_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>The name of the source location where the slate VOD source is stored.</p>",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: VodSourceName
				"vod_source_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>The slate VOD source name. The VOD source must already exist in a source location before it can be used for slate.</p>",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "<p>Slate VOD source configuration.</p>",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LogConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "\u003cp\u003eThe log configuration for the channel.\u003c/p\u003e",
		//	  "properties": {
		//	    "LogTypes": {
		//	      "description": "\u003cp\u003eThe log types.\u003c/p\u003e",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "AS_RUN"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"log_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LogTypes
				"log_types": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "<p>The log types.</p>",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.OneOf(
								"AS_RUN",
							),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "<p>The log configuration for the channel.</p>",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Outputs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe channel's output properties.\u003c/p\u003e",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "\u003cp\u003eThe output configuration for this channel.\u003c/p\u003e",
		//	    "properties": {
		//	      "DashPlaylistSettings": {
		//	        "additionalProperties": false,
		//	        "description": "\u003cp\u003eDash manifest configuration parameters.\u003c/p\u003e",
		//	        "properties": {
		//	          "ManifestWindowSeconds": {
		//	            "description": "\u003cp\u003eThe total duration (in seconds) of each manifest. Minimum value: \u003ccode\u003e30\u003c/code\u003e seconds. Maximum value: \u003ccode\u003e3600\u003c/code\u003e seconds.\u003c/p\u003e",
		//	            "type": "number"
		//	          },
		//	          "MinBufferTimeSeconds": {
		//	            "description": "\u003cp\u003eMinimum amount of content (measured in seconds) that a player must keep available in the buffer. Minimum value: \u003ccode\u003e2\u003c/code\u003e seconds. Maximum value: \u003ccode\u003e60\u003c/code\u003e seconds.\u003c/p\u003e",
		//	            "type": "number"
		//	          },
		//	          "MinUpdatePeriodSeconds": {
		//	            "description": "\u003cp\u003eMinimum amount of time (in seconds) that the player should wait before requesting updates to the manifest. Minimum value: \u003ccode\u003e2\u003c/code\u003e seconds. Maximum value: \u003ccode\u003e60\u003c/code\u003e seconds.\u003c/p\u003e",
		//	            "type": "number"
		//	          },
		//	          "SuggestedPresentationDelaySeconds": {
		//	            "description": "\u003cp\u003eAmount of time (in seconds) that the player should be from the live point at the end of the manifest. Minimum value: \u003ccode\u003e2\u003c/code\u003e seconds. Maximum value: \u003ccode\u003e60\u003c/code\u003e seconds.\u003c/p\u003e",
		//	            "type": "number"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "HlsPlaylistSettings": {
		//	        "additionalProperties": false,
		//	        "description": "\u003cp\u003eHLS playlist configuration parameters.\u003c/p\u003e",
		//	        "properties": {
		//	          "AdMarkupType": {
		//	            "description": "\u003cp\u003eDetermines the type of SCTE 35 tags to use in ad markup. Specify \u003ccode\u003eDATERANGE\u003c/code\u003e to use \u003ccode\u003eDATERANGE\u003c/code\u003e tags (for live or VOD content). Specify \u003ccode\u003eSCTE35_ENHANCED\u003c/code\u003e to use \u003ccode\u003eEXT-X-CUE-OUT\u003c/code\u003e and \u003ccode\u003eEXT-X-CUE-IN\u003c/code\u003e tags (for VOD content only).\u003c/p\u003e",
		//	            "items": {
		//	              "enum": [
		//	                "DATERANGE",
		//	                "SCTE35_ENHANCED"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "ManifestWindowSeconds": {
		//	            "description": "\u003cp\u003eThe total duration (in seconds) of each manifest. Minimum value: \u003ccode\u003e30\u003c/code\u003e seconds. Maximum value: \u003ccode\u003e3600\u003c/code\u003e seconds.\u003c/p\u003e",
		//	            "type": "number"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ManifestName": {
		//	        "description": "\u003cp\u003eThe name of the manifest for the channel. The name appears in the \u003ccode\u003ePlaybackUrl\u003c/code\u003e.\u003c/p\u003e",
		//	        "type": "string"
		//	      },
		//	      "SourceGroup": {
		//	        "description": "\u003cp\u003eA string used to match which \u003ccode\u003eHttpPackageConfiguration\u003c/code\u003e is used for each \u003ccode\u003eVodSource\u003c/code\u003e.\u003c/p\u003e",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ManifestName",
		//	      "SourceGroup"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"outputs": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DashPlaylistSettings
					"dash_playlist_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ManifestWindowSeconds
							"manifest_window_seconds": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "<p>The total duration (in seconds) of each manifest. Minimum value: <code>30</code> seconds. Maximum value: <code>3600</code> seconds.</p>",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: MinBufferTimeSeconds
							"min_buffer_time_seconds": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "<p>Minimum amount of content (measured in seconds) that a player must keep available in the buffer. Minimum value: <code>2</code> seconds. Maximum value: <code>60</code> seconds.</p>",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: MinUpdatePeriodSeconds
							"min_update_period_seconds": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "<p>Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest. Minimum value: <code>2</code> seconds. Maximum value: <code>60</code> seconds.</p>",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: SuggestedPresentationDelaySeconds
							"suggested_presentation_delay_seconds": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "<p>Amount of time (in seconds) that the player should be from the live point at the end of the manifest. Minimum value: <code>2</code> seconds. Maximum value: <code>60</code> seconds.</p>",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "<p>Dash manifest configuration parameters.</p>",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: HlsPlaylistSettings
					"hls_playlist_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AdMarkupType
							"ad_markup_type": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "<p>Determines the type of SCTE 35 tags to use in ad markup. Specify <code>DATERANGE</code> to use <code>DATERANGE</code> tags (for live or VOD content). Specify <code>SCTE35_ENHANCED</code> to use <code>EXT-X-CUE-OUT</code> and <code>EXT-X-CUE-IN</code> tags (for VOD content only).</p>",
								Optional:    true,
								Computed:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.ValueStringsAre(
										stringvalidator.OneOf(
											"DATERANGE",
											"SCTE35_ENHANCED",
										),
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ManifestWindowSeconds
							"manifest_window_seconds": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "<p>The total duration (in seconds) of each manifest. Minimum value: <code>30</code> seconds. Maximum value: <code>3600</code> seconds.</p>",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "<p>HLS playlist configuration parameters.</p>",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ManifestName
					"manifest_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>The name of the manifest for the channel. The name appears in the <code>PlaybackUrl</code>.</p>",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: SourceGroup
					"source_group": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>A string used to match which <code>HttpPackageConfiguration</code> is used for each <code>VodSource</code>.</p>",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "<p>The channel's output properties.</p>",
			Required:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
			// Outputs is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: PlaybackMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "LOOP",
		//	    "LINEAR"
		//	  ],
		//	  "type": "string"
		//	}
		"playback_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"LOOP",
					"LINEAR",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to assign to the channel.",
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
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
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
			Description: "The tags to assign to the channel.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "BASIC",
		//	    "STANDARD"
		//	  ],
		//	  "type": "string"
		//	}
		"tier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"BASIC",
					"STANDARD",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TimeShiftConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "\u003cp\u003eThe configuration for time-shifted viewing.\u003c/p\u003e",
		//	  "properties": {
		//	    "MaxTimeDelaySeconds": {
		//	      "description": "\u003cp\u003eThe maximum time delay for time-shifted viewing. The minimum allowed maximum time delay is 0 seconds, and the maximum allowed maximum time delay is 21600 seconds (6 hours).\u003c/p\u003e",
		//	      "type": "number"
		//	    }
		//	  },
		//	  "required": [
		//	    "MaxTimeDelaySeconds"
		//	  ],
		//	  "type": "object"
		//	}
		"time_shift_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaxTimeDelaySeconds
				"max_time_delay_seconds": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "<p>The maximum time delay for time-shifted viewing. The minimum allowed maximum time delay is 0 seconds, and the maximum allowed maximum time delay is 21600 seconds (6 hours).</p>",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "<p>The configuration for time-shifted viewing.</p>",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
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
		Description: "Definition of AWS::MediaTailor::Channel Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaTailor::Channel").WithTerraformTypeName("awscc_mediatailor_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"ad_markup_type":                       "AdMarkupType",
		"arn":                                  "Arn",
		"channel_name":                         "ChannelName",
		"dash_playlist_settings":               "DashPlaylistSettings",
		"filler_slate":                         "FillerSlate",
		"hls_playlist_settings":                "HlsPlaylistSettings",
		"key":                                  "Key",
		"log_configuration":                    "LogConfiguration",
		"log_types":                            "LogTypes",
		"manifest_name":                        "ManifestName",
		"manifest_window_seconds":              "ManifestWindowSeconds",
		"max_time_delay_seconds":               "MaxTimeDelaySeconds",
		"min_buffer_time_seconds":              "MinBufferTimeSeconds",
		"min_update_period_seconds":            "MinUpdatePeriodSeconds",
		"outputs":                              "Outputs",
		"playback_mode":                        "PlaybackMode",
		"source_group":                         "SourceGroup",
		"source_location_name":                 "SourceLocationName",
		"suggested_presentation_delay_seconds": "SuggestedPresentationDelaySeconds",
		"tags":                                 "Tags",
		"tier":                                 "Tier",
		"time_shift_configuration":             "TimeShiftConfiguration",
		"value":                                "Value",
		"vod_source_name":                      "VodSourceName",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Outputs",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
