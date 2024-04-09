// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_apigateway_deployment", deploymentResource)
}

// deploymentResource returns the Terraform awscc_apigateway_deployment resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGateway::Deployment resource.
func deploymentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DeploymentCanarySettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The input configuration for a canary deployment.",
		//	  "properties": {
		//	    "PercentTraffic": {
		//	      "description": "The percentage (0.0-100.0) of traffic routed to the canary deployment.",
		//	      "type": "number"
		//	    },
		//	    "StageVariableOverrides": {
		//	      "additionalProperties": false,
		//	      "description": "A stage variable overrides used for the canary release deployment. They can override existing stage variables or add new stage variables for the canary release deployment. These stage variables are represented as a string-to-string map between stage variable names and their values.",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "UseStageCache": {
		//	      "description": "A Boolean flag to indicate whether the canary release deployment uses the stage cache or not.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"deployment_canary_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PercentTraffic
				"percent_traffic": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The percentage (0.0-100.0) of traffic routed to the canary deployment.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StageVariableOverrides
				"stage_variable_overrides": // Pattern: ""
				schema.MapAttribute{        /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A stage variable overrides used for the canary release deployment. They can override existing stage variables or add new stage variables for the canary release deployment. These stage variables are represented as a string-to-string map between stage variable names and their values.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
						mapplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: UseStageCache
				"use_stage_cache": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "A Boolean flag to indicate whether the canary release deployment uses the stage cache or not.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The input configuration for a canary deployment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// DeploymentCanarySettings is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: DeploymentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"deployment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description for the Deployment resource to create.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description for the Deployment resource to create.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The string identifier of the associated RestApi.",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The string identifier of the associated RestApi.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StageDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The description of the Stage resource for the Deployment resource to create. To specify a stage description, you must also provide a stage name.",
		//	  "properties": {
		//	    "AccessLogSetting": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies settings for logging access in this stage.",
		//	      "properties": {
		//	        "DestinationArn": {
		//	          "description": "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ``amazon-apigateway-``.",
		//	          "type": "string"
		//	        },
		//	        "Format": {
		//	          "description": "A single line format of the access logs of data, as specified by selected $context variables. The format must include at least ``$context.requestId``.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "CacheClusterEnabled": {
		//	      "description": "Specifies whether a cache cluster is enabled for the stage.",
		//	      "type": "boolean"
		//	    },
		//	    "CacheClusterSize": {
		//	      "description": "The size of the stage's cache cluster. For more information, see [cacheClusterSize](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateStage.html#apigw-CreateStage-request-cacheClusterSize) in the *API Gateway API Reference*.",
		//	      "type": "string"
		//	    },
		//	    "CacheDataEncrypted": {
		//	      "description": "Indicates whether the cached responses are encrypted.",
		//	      "type": "boolean"
		//	    },
		//	    "CacheTtlInSeconds": {
		//	      "description": "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.",
		//	      "type": "integer"
		//	    },
		//	    "CachingEnabled": {
		//	      "description": "Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses. For more information, see [Enable API Gateway Caching in a Stage to Enhance API Performance](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html) in the *API Gateway Developer Guide*.",
		//	      "type": "boolean"
		//	    },
		//	    "CanarySetting": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies settings for the canary deployment in this stage.",
		//	      "properties": {
		//	        "PercentTraffic": {
		//	          "description": "The percent (0-100) of traffic diverted to a canary deployment.",
		//	          "type": "number"
		//	        },
		//	        "StageVariableOverrides": {
		//	          "additionalProperties": false,
		//	          "description": "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.",
		//	          "patternProperties": {
		//	            "": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "UseStageCache": {
		//	          "description": "A Boolean flag to indicate whether the canary deployment uses the stage cache or not.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ClientCertificateId": {
		//	      "description": "The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.",
		//	      "type": "string"
		//	    },
		//	    "DataTraceEnabled": {
		//	      "description": "Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs.",
		//	      "type": "boolean"
		//	    },
		//	    "Description": {
		//	      "description": "A description of the purpose of the stage.",
		//	      "type": "string"
		//	    },
		//	    "DocumentationVersion": {
		//	      "description": "The version identifier of the API documentation snapshot.",
		//	      "type": "string"
		//	    },
		//	    "LoggingLevel": {
		//	      "description": "The logging level for this method. For valid values, see the ``loggingLevel`` property of the [MethodSetting](https://docs.aws.amazon.com/apigateway/latest/api/API_MethodSetting.html) resource in the *Amazon API Gateway API Reference*.",
		//	      "type": "string"
		//	    },
		//	    "MethodSettings": {
		//	      "description": "Configures settings for all of the stage's methods.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "The ``MethodSetting`` property type configures settings for all methods in a stage.\n The ``MethodSettings`` property of the [Amazon API Gateway Deployment StageDescription](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html) property type contains a list of ``MethodSetting`` property types.",
		//	        "properties": {
		//	          "CacheDataEncrypted": {
		//	            "description": "Specifies whether the cached responses are encrypted.",
		//	            "type": "boolean"
		//	          },
		//	          "CacheTtlInSeconds": {
		//	            "description": "Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.",
		//	            "type": "integer"
		//	          },
		//	          "CachingEnabled": {
		//	            "description": "Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.",
		//	            "type": "boolean"
		//	          },
		//	          "DataTraceEnabled": {
		//	            "description": "Specifies whether data trace logging is enabled for this method, which affects the log entries pushed to Amazon CloudWatch Logs. This can be useful to troubleshoot APIs, but can result in logging sensitive data. We recommend that you don't enable this option for production APIs.",
		//	            "type": "boolean"
		//	          },
		//	          "HttpMethod": {
		//	            "description": "The HTTP method.",
		//	            "type": "string"
		//	          },
		//	          "LoggingLevel": {
		//	            "description": "Specifies the logging level for this method, which affects the log entries pushed to Amazon CloudWatch Logs. Valid values are ``OFF``, ``ERROR``, and ``INFO``. Choose ``ERROR`` to write only error-level entries to CloudWatch Logs, or choose ``INFO`` to include all ``ERROR`` events as well as extra informational events.",
		//	            "type": "string"
		//	          },
		//	          "MetricsEnabled": {
		//	            "description": "Specifies whether Amazon CloudWatch metrics are enabled for this method.",
		//	            "type": "boolean"
		//	          },
		//	          "ResourcePath": {
		//	            "description": "The resource path for this method. Forward slashes (``/``) are encoded as ``~1`` and the initial slash must include a forward slash. For example, the path value ``/resource/subresource`` must be encoded as ``/~1resource~1subresource``. To specify the root path, use only a slash (``/``).",
		//	            "type": "string"
		//	          },
		//	          "ThrottlingBurstLimit": {
		//	            "description": "Specifies the throttling burst limit.",
		//	            "type": "integer"
		//	          },
		//	          "ThrottlingRateLimit": {
		//	            "description": "Specifies the throttling rate limit.",
		//	            "type": "number"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "MetricsEnabled": {
		//	      "description": "Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.",
		//	      "type": "boolean"
		//	    },
		//	    "Tags": {
		//	      "description": "An array of arbitrary tags (key-value pairs) to associate with the stage.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "",
		//	        "properties": {
		//	          "Key": {
		//	            "description": "The key name of the tag",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "description": "The value for the tag",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Value",
		//	          "Key"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "ThrottlingBurstLimit": {
		//	      "description": "The target request burst rate limit. This allows more requests through for a period of time than the target rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide*.",
		//	      "type": "integer"
		//	    },
		//	    "ThrottlingRateLimit": {
		//	      "description": "The target request steady-state rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide*.",
		//	      "type": "number"
		//	    },
		//	    "TracingEnabled": {
		//	      "description": "Specifies whether active tracing with X-ray is enabled for this stage.\n For more information, see [Trace API Gateway API Execution with X-Ray](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html) in the *API Gateway Developer Guide*.",
		//	      "type": "boolean"
		//	    },
		//	    "Variables": {
		//	      "additionalProperties": false,
		//	      "description": "A map that defines the stage variables. Variable names must consist of alphanumeric characters, and the values must match the following regular expression: ``[A-Za-z0-9-._~:/?#\u0026=,]+``.",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"stage_description": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AccessLogSetting
				"access_log_setting": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DestinationArn
						"destination_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ``amazon-apigateway-``.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Format
						"format": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "A single line format of the access logs of data, as specified by selected $context variables. The format must include at least ``$context.requestId``.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies settings for logging access in this stage.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CacheClusterEnabled
				"cache_cluster_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether a cache cluster is enabled for the stage.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CacheClusterSize
				"cache_cluster_size": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The size of the stage's cache cluster. For more information, see [cacheClusterSize](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateStage.html#apigw-CreateStage-request-cacheClusterSize) in the *API Gateway API Reference*.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CacheDataEncrypted
				"cache_data_encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether the cached responses are encrypted.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CacheTtlInSeconds
				"cache_ttl_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CachingEnabled
				"caching_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses. For more information, see [Enable API Gateway Caching in a Stage to Enhance API Performance](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html) in the *API Gateway Developer Guide*.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CanarySetting
				"canary_setting": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: PercentTraffic
						"percent_traffic": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Description: "The percent (0-100) of traffic diverted to a canary deployment.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
								float64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: StageVariableOverrides
						"stage_variable_overrides": // Pattern: ""
						schema.MapAttribute{        /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
								mapplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: UseStageCache
						"use_stage_cache": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "A Boolean flag to indicate whether the canary deployment uses the stage cache or not.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies settings for the canary deployment in this stage.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ClientCertificateId
				"client_certificate_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DataTraceEnabled
				"data_trace_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Description
				"description": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A description of the purpose of the stage.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DocumentationVersion
				"documentation_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version identifier of the API documentation snapshot.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LoggingLevel
				"logging_level": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The logging level for this method. For valid values, see the ``loggingLevel`` property of the [MethodSetting](https://docs.aws.amazon.com/apigateway/latest/api/API_MethodSetting.html) resource in the *Amazon API Gateway API Reference*.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MethodSettings
				"method_settings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CacheDataEncrypted
							"cache_data_encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Specifies whether the cached responses are encrypted.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: CacheTtlInSeconds
							"cache_ttl_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: CachingEnabled
							"caching_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: DataTraceEnabled
							"data_trace_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Specifies whether data trace logging is enabled for this method, which affects the log entries pushed to Amazon CloudWatch Logs. This can be useful to troubleshoot APIs, but can result in logging sensitive data. We recommend that you don't enable this option for production APIs.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: HttpMethod
							"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The HTTP method.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: LoggingLevel
							"logging_level": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Specifies the logging level for this method, which affects the log entries pushed to Amazon CloudWatch Logs. Valid values are ``OFF``, ``ERROR``, and ``INFO``. Choose ``ERROR`` to write only error-level entries to CloudWatch Logs, or choose ``INFO`` to include all ``ERROR`` events as well as extra informational events.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: MetricsEnabled
							"metrics_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Specifies whether Amazon CloudWatch metrics are enabled for this method.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ResourcePath
							"resource_path": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The resource path for this method. Forward slashes (``/``) are encoded as ``~1`` and the initial slash must include a forward slash. For example, the path value ``/resource/subresource`` must be encoded as ``/~1resource~1subresource``. To specify the root path, use only a slash (``/``).",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ThrottlingBurstLimit
							"throttling_burst_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "Specifies the throttling burst limit.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ThrottlingRateLimit
							"throttling_rate_limit": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "Specifies the throttling rate limit.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Configures settings for all of the stage's methods.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MetricsEnabled
				"metrics_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Tags
				"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The key name of the tag",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The value for the tag",
								Required:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "An array of arbitrary tags (key-value pairs) to associate with the stage.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ThrottlingBurstLimit
				"throttling_burst_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The target request burst rate limit. This allows more requests through for a period of time than the target rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide*.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ThrottlingRateLimit
				"throttling_rate_limit": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The target request steady-state rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide*.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TracingEnabled
				"tracing_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether active tracing with X-ray is enabled for this stage.\n For more information, see [Trace API Gateway API Execution with X-Ray](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html) in the *API Gateway Developer Guide*.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Variables
				"variables":         // Pattern: ""
				schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A map that defines the stage variables. Variable names must consist of alphanumeric characters, and the values must match the following regular expression: ``[A-Za-z0-9-._~:/?#&=,]+``.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
						mapplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The description of the Stage resource for the Deployment resource to create. To specify a stage description, you must also provide a stage name.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// StageDescription is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: StageName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Stage resource for the Deployment resource to create.",
		//	  "type": "string"
		//	}
		"stage_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Stage resource for the Deployment resource to create.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// StageName is a write-only property.
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
		Description: "The ``AWS::ApiGateway::Deployment`` resource deploys an API Gateway ``RestApi`` resource to a stage so that clients can call the API over the internet. The stage acts as an environment.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Deployment").WithTerraformTypeName("awscc_apigateway_deployment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_log_setting":         "AccessLogSetting",
		"cache_cluster_enabled":      "CacheClusterEnabled",
		"cache_cluster_size":         "CacheClusterSize",
		"cache_data_encrypted":       "CacheDataEncrypted",
		"cache_ttl_in_seconds":       "CacheTtlInSeconds",
		"caching_enabled":            "CachingEnabled",
		"canary_setting":             "CanarySetting",
		"client_certificate_id":      "ClientCertificateId",
		"data_trace_enabled":         "DataTraceEnabled",
		"deployment_canary_settings": "DeploymentCanarySettings",
		"deployment_id":              "DeploymentId",
		"description":                "Description",
		"destination_arn":            "DestinationArn",
		"documentation_version":      "DocumentationVersion",
		"format":                     "Format",
		"http_method":                "HttpMethod",
		"key":                        "Key",
		"logging_level":              "LoggingLevel",
		"method_settings":            "MethodSettings",
		"metrics_enabled":            "MetricsEnabled",
		"percent_traffic":            "PercentTraffic",
		"resource_path":              "ResourcePath",
		"rest_api_id":                "RestApiId",
		"stage_description":          "StageDescription",
		"stage_name":                 "StageName",
		"stage_variable_overrides":   "StageVariableOverrides",
		"tags":                       "Tags",
		"throttling_burst_limit":     "ThrottlingBurstLimit",
		"throttling_rate_limit":      "ThrottlingRateLimit",
		"tracing_enabled":            "TracingEnabled",
		"use_stage_cache":            "UseStageCache",
		"value":                      "Value",
		"variables":                  "Variables",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/StageName",
		"/properties/StageDescription",
		"/properties/DeploymentCanarySettings",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
