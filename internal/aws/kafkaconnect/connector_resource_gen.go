// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package kafkaconnect

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
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
	registry.AddResourceFactory("awscc_kafkaconnect_connector", connectorResource)
}

// connectorResource returns the Terraform awscc_kafkaconnect_connector resource.
// This Terraform resource corresponds to the CloudFormation AWS::KafkaConnect::Connector resource.
func connectorResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Capacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Information about the capacity allocated to the connector.",
		//	  "oneOf": [
		//	    {
		//	      "required": [
		//	        "AutoScaling"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "ProvisionedCapacity"
		//	      ]
		//	    }
		//	  ],
		//	  "properties": {
		//	    "AutoScaling": {
		//	      "additionalProperties": false,
		//	      "description": "Details about auto scaling of a connector.",
		//	      "properties": {
		//	        "MaxWorkerCount": {
		//	          "description": "The maximum number of workers for a connector.",
		//	          "type": "integer"
		//	        },
		//	        "McuCount": {
		//	          "description": "Specifies how many MSK Connect Units (MCU) as the minimum scaling unit.",
		//	          "enum": [
		//	            1,
		//	            2,
		//	            4,
		//	            8
		//	          ],
		//	          "type": "integer"
		//	        },
		//	        "MinWorkerCount": {
		//	          "description": "The minimum number of workers for a connector.",
		//	          "type": "integer"
		//	        },
		//	        "ScaleInPolicy": {
		//	          "additionalProperties": false,
		//	          "description": "Information about the scale in policy of the connector.",
		//	          "properties": {
		//	            "CpuUtilizationPercentage": {
		//	              "description": "Specifies the CPU utilization percentage threshold at which connector scale in should trigger.",
		//	              "maximum": 100,
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "required": [
		//	            "CpuUtilizationPercentage"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "ScaleOutPolicy": {
		//	          "additionalProperties": false,
		//	          "description": "Information about the scale out policy of the connector.",
		//	          "properties": {
		//	            "CpuUtilizationPercentage": {
		//	              "description": "Specifies the CPU utilization percentage threshold at which connector scale out should trigger.",
		//	              "maximum": 100,
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "required": [
		//	            "CpuUtilizationPercentage"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "MaxWorkerCount",
		//	        "MinWorkerCount",
		//	        "ScaleInPolicy",
		//	        "ScaleOutPolicy",
		//	        "McuCount"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ProvisionedCapacity": {
		//	      "additionalProperties": false,
		//	      "description": "Details about a fixed capacity allocated to a connector.",
		//	      "properties": {
		//	        "McuCount": {
		//	          "description": "Specifies how many MSK Connect Units (MCU) are allocated to the connector.",
		//	          "enum": [
		//	            1,
		//	            2,
		//	            4,
		//	            8
		//	          ],
		//	          "type": "integer"
		//	        },
		//	        "WorkerCount": {
		//	          "description": "Number of workers for a connector.",
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "WorkerCount"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"capacity": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AutoScaling
				"auto_scaling": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: MaxWorkerCount
						"max_worker_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The maximum number of workers for a connector.",
							Required:    true,
						}, /*END ATTRIBUTE*/
						// Property: McuCount
						"mcu_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "Specifies how many MSK Connect Units (MCU) as the minimum scaling unit.",
							Required:    true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.OneOf(
									1,
									2,
									4,
									8,
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: MinWorkerCount
						"min_worker_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The minimum number of workers for a connector.",
							Required:    true,
						}, /*END ATTRIBUTE*/
						// Property: ScaleInPolicy
						"scale_in_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CpuUtilizationPercentage
								"cpu_utilization_percentage": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Specifies the CPU utilization percentage threshold at which connector scale in should trigger.",
									Required:    true,
									Validators: []validator.Int64{ /*START VALIDATORS*/
										int64validator.Between(1, 100),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Information about the scale in policy of the connector.",
							Required:    true,
						}, /*END ATTRIBUTE*/
						// Property: ScaleOutPolicy
						"scale_out_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CpuUtilizationPercentage
								"cpu_utilization_percentage": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Specifies the CPU utilization percentage threshold at which connector scale out should trigger.",
									Required:    true,
									Validators: []validator.Int64{ /*START VALIDATORS*/
										int64validator.Between(1, 100),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Information about the scale out policy of the connector.",
							Required:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Details about auto scaling of a connector.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ProvisionedCapacity
				"provisioned_capacity": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: McuCount
						"mcu_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "Specifies how many MSK Connect Units (MCU) are allocated to the connector.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.OneOf(
									1,
									2,
									4,
									8,
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: WorkerCount
						"worker_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "Number of workers for a connector.",
							Required:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Details about a fixed capacity allocated to a connector.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Information about the capacity allocated to the connector.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConnectorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name for the created Connector.",
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*",
		//	  "type": "string"
		//	}
		"connector_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name for the created Connector.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConnectorConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration for the connector.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"connector_configuration": // Pattern: ""
		schema.MapAttribute{       /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The configuration for the connector.",
			Required:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConnectorDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A summary description of the connector.",
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"connector_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A summary description of the connector.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConnectorName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the connector.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"connector_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the connector.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KafkaCluster
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Details of how to connect to the Kafka cluster.",
		//	  "properties": {
		//	    "ApacheKafkaCluster": {
		//	      "additionalProperties": false,
		//	      "description": "Details of how to connect to an Apache Kafka cluster.",
		//	      "properties": {
		//	        "BootstrapServers": {
		//	          "description": "The bootstrap servers string of the Apache Kafka cluster.",
		//	          "type": "string"
		//	        },
		//	        "Vpc": {
		//	          "additionalProperties": false,
		//	          "description": "Information about a VPC used with the connector.",
		//	          "properties": {
		//	            "SecurityGroups": {
		//	              "description": "The AWS security groups to associate with the elastic network interfaces in order to specify what the connector has access to.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "Subnets": {
		//	              "description": "The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            }
		//	          },
		//	          "required": [
		//	            "SecurityGroups",
		//	            "Subnets"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "BootstrapServers",
		//	        "Vpc"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "ApacheKafkaCluster"
		//	  ],
		//	  "type": "object"
		//	}
		"kafka_cluster": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ApacheKafkaCluster
				"apache_kafka_cluster": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BootstrapServers
						"bootstrap_servers": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The bootstrap servers string of the Apache Kafka cluster.",
							Required:    true,
						}, /*END ATTRIBUTE*/
						// Property: Vpc
						"vpc": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: SecurityGroups
								"security_groups": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "The AWS security groups to associate with the elastic network interfaces in order to specify what the connector has access to.",
									Required:    true,
								}, /*END ATTRIBUTE*/
								// Property: Subnets
								"subnets": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets.",
									Required:    true,
									Validators: []validator.Set{ /*START VALIDATORS*/
										setvalidator.SizeAtLeast(1),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Information about a VPC used with the connector.",
							Required:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Details of how to connect to an Apache Kafka cluster.",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Details of how to connect to the Kafka cluster.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KafkaClusterClientAuthentication
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Details of the client authentication used by the Kafka cluster.",
		//	  "properties": {
		//	    "AuthenticationType": {
		//	      "description": "The type of client authentication used to connect to the Kafka cluster. Value NONE means that no client authentication is used.",
		//	      "enum": [
		//	        "NONE",
		//	        "IAM"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "AuthenticationType"
		//	  ],
		//	  "type": "object"
		//	}
		"kafka_cluster_client_authentication": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AuthenticationType
				"authentication_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of client authentication used to connect to the Kafka cluster. Value NONE means that no client authentication is used.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"NONE",
							"IAM",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Details of the client authentication used by the Kafka cluster.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KafkaClusterEncryptionInTransit
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Details of encryption in transit to the Kafka cluster.",
		//	  "properties": {
		//	    "EncryptionType": {
		//	      "description": "The type of encryption in transit to the Kafka cluster.",
		//	      "enum": [
		//	        "PLAINTEXT",
		//	        "TLS"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "EncryptionType"
		//	  ],
		//	  "type": "object"
		//	}
		"kafka_cluster_encryption_in_transit": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EncryptionType
				"encryption_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of encryption in transit to the Kafka cluster.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"PLAINTEXT",
							"TLS",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Details of encryption in transit to the Kafka cluster.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KafkaConnectVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.",
		//	  "type": "string"
		//	}
		"kafka_connect_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LogDelivery
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Details of what logs are delivered and where they are delivered.",
		//	  "properties": {
		//	    "WorkerLogDelivery": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies where worker logs are delivered.",
		//	      "properties": {
		//	        "CloudWatchLogs": {
		//	          "additionalProperties": false,
		//	          "description": "Details about delivering logs to Amazon CloudWatch Logs.",
		//	          "properties": {
		//	            "Enabled": {
		//	              "description": "Specifies whether the logs get sent to the specified CloudWatch Logs destination.",
		//	              "type": "boolean"
		//	            },
		//	            "LogGroup": {
		//	              "description": "The CloudWatch log group that is the destination for log delivery.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Enabled"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "Firehose": {
		//	          "additionalProperties": false,
		//	          "description": "Details about delivering logs to Amazon Kinesis Data Firehose.",
		//	          "properties": {
		//	            "DeliveryStream": {
		//	              "description": "The Kinesis Data Firehose delivery stream that is the destination for log delivery.",
		//	              "type": "string"
		//	            },
		//	            "Enabled": {
		//	              "description": "Specifies whether the logs get sent to the specified Kinesis Data Firehose delivery stream.",
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "Enabled"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "S3": {
		//	          "additionalProperties": false,
		//	          "description": "Details about delivering logs to Amazon S3.",
		//	          "properties": {
		//	            "Bucket": {
		//	              "description": "The name of the S3 bucket that is the destination for log delivery.",
		//	              "type": "string"
		//	            },
		//	            "Enabled": {
		//	              "description": "Specifies whether the logs get sent to the specified Amazon S3 destination.",
		//	              "type": "boolean"
		//	            },
		//	            "Prefix": {
		//	              "description": "The S3 prefix that is the destination for log delivery.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Enabled"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "WorkerLogDelivery"
		//	  ],
		//	  "type": "object"
		//	}
		"log_delivery": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: WorkerLogDelivery
				"worker_log_delivery": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CloudWatchLogs
						"cloudwatch_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies whether the logs get sent to the specified CloudWatch Logs destination.",
									Required:    true,
								}, /*END ATTRIBUTE*/
								// Property: LogGroup
								"log_group": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The CloudWatch log group that is the destination for log delivery.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Details about delivering logs to Amazon CloudWatch Logs.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Firehose
						"firehose": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: DeliveryStream
								"delivery_stream": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The Kinesis Data Firehose delivery stream that is the destination for log delivery.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies whether the logs get sent to the specified Kinesis Data Firehose delivery stream.",
									Required:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Details about delivering logs to Amazon Kinesis Data Firehose.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: S3
						"s3": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Bucket
								"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the S3 bucket that is the destination for log delivery.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies whether the logs get sent to the specified Amazon S3 destination.",
									Required:    true,
								}, /*END ATTRIBUTE*/
								// Property: Prefix
								"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The S3 prefix that is the destination for log delivery.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Details about delivering logs to Amazon S3.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies where worker logs are delivered.",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Details of what logs are delivered and where they are delivered.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Plugins
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of plugins to use with the connector.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Details about a Kafka Connect plugin which will be used with the connector.",
		//	    "properties": {
		//	      "CustomPlugin": {
		//	        "additionalProperties": false,
		//	        "description": "Details about a custom plugin.",
		//	        "properties": {
		//	          "CustomPluginArn": {
		//	            "description": "The Amazon Resource Name (ARN) of the custom plugin to use.",
		//	            "pattern": "arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*",
		//	            "type": "string"
		//	          },
		//	          "Revision": {
		//	            "description": "The revision of the custom plugin to use.",
		//	            "format": "int64",
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "required": [
		//	          "CustomPluginArn",
		//	          "Revision"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "CustomPlugin"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"plugins": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CustomPlugin
					"custom_plugin": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CustomPluginArn
							"custom_plugin_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The Amazon Resource Name (ARN) of the custom plugin to use.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*"), ""),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Revision
							"revision": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The revision of the custom plugin to use.",
								Required:    true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.AtLeast(1),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Details about a custom plugin.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of plugins to use with the connector.",
			Required:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceExecutionRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.",
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn):iam:.*",
		//	  "type": "string"
		//	}
		"service_execution_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("arn:(aws|aws-us-gov|aws-cn):iam:.*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource",
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
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A collection of tags associated with a resource",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WorkerConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies the worker configuration to use with the connector.",
		//	  "properties": {
		//	    "Revision": {
		//	      "description": "The revision of the worker configuration to use.",
		//	      "format": "int64",
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "WorkerConfigurationArn": {
		//	      "description": "The Amazon Resource Name (ARN) of the worker configuration to use.",
		//	      "pattern": "arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Revision",
		//	    "WorkerConfigurationArn"
		//	  ],
		//	  "type": "object"
		//	}
		"worker_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Revision
				"revision": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The revision of the worker configuration to use.",
					Required:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(1),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: WorkerConfigurationArn
				"worker_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon Resource Name (ARN) of the worker configuration to use.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies the worker configuration to use with the connector.",
			Optional:    true,
			Computed:    true,
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
		Description: "Resource Type definition for AWS::KafkaConnect::Connector",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::KafkaConnect::Connector").WithTerraformTypeName("awscc_kafkaconnect_connector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"apache_kafka_cluster":                "ApacheKafkaCluster",
		"authentication_type":                 "AuthenticationType",
		"auto_scaling":                        "AutoScaling",
		"bootstrap_servers":                   "BootstrapServers",
		"bucket":                              "Bucket",
		"capacity":                            "Capacity",
		"cloudwatch_logs":                     "CloudWatchLogs",
		"connector_arn":                       "ConnectorArn",
		"connector_configuration":             "ConnectorConfiguration",
		"connector_description":               "ConnectorDescription",
		"connector_name":                      "ConnectorName",
		"cpu_utilization_percentage":          "CpuUtilizationPercentage",
		"custom_plugin":                       "CustomPlugin",
		"custom_plugin_arn":                   "CustomPluginArn",
		"delivery_stream":                     "DeliveryStream",
		"enabled":                             "Enabled",
		"encryption_type":                     "EncryptionType",
		"firehose":                            "Firehose",
		"kafka_cluster":                       "KafkaCluster",
		"kafka_cluster_client_authentication": "KafkaClusterClientAuthentication",
		"kafka_cluster_encryption_in_transit": "KafkaClusterEncryptionInTransit",
		"kafka_connect_version":               "KafkaConnectVersion",
		"key":                                 "Key",
		"log_delivery":                        "LogDelivery",
		"log_group":                           "LogGroup",
		"max_worker_count":                    "MaxWorkerCount",
		"mcu_count":                           "McuCount",
		"min_worker_count":                    "MinWorkerCount",
		"plugins":                             "Plugins",
		"prefix":                              "Prefix",
		"provisioned_capacity":                "ProvisionedCapacity",
		"revision":                            "Revision",
		"s3":                                  "S3",
		"scale_in_policy":                     "ScaleInPolicy",
		"scale_out_policy":                    "ScaleOutPolicy",
		"security_groups":                     "SecurityGroups",
		"service_execution_role_arn":          "ServiceExecutionRoleArn",
		"subnets":                             "Subnets",
		"tags":                                "Tags",
		"value":                               "Value",
		"vpc":                                 "Vpc",
		"worker_configuration":                "WorkerConfiguration",
		"worker_configuration_arn":            "WorkerConfigurationArn",
		"worker_count":                        "WorkerCount",
		"worker_log_delivery":                 "WorkerLogDelivery",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
