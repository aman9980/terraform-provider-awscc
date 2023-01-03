// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ssmincidents

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ssmincidents_response_plan", responsePlanDataSource)
}

// responsePlanDataSource returns the Terraform awscc_ssmincidents_response_plan data source.
// This Terraform data source corresponds to the CloudFormation AWS::SSMIncidents::ResponsePlan resource.
func responsePlanDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Actions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": [],
		//	  "description": "The list of actions.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The automation configuration to launch.",
		//	    "properties": {
		//	      "SsmAutomation": {
		//	        "additionalProperties": false,
		//	        "description": "The configuration to use when starting the SSM automation document.",
		//	        "properties": {
		//	          "DocumentName": {
		//	            "description": "The document name to use when starting the SSM automation document.",
		//	            "maxLength": 128,
		//	            "type": "string"
		//	          },
		//	          "DocumentVersion": {
		//	            "description": "The version of the document to use when starting the SSM automation document.",
		//	            "maxLength": 128,
		//	            "type": "string"
		//	          },
		//	          "DynamicParameters": {
		//	            "default": [],
		//	            "description": "The parameters with dynamic values to set when starting the SSM automation document.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "description": "A parameter with a dynamic value to set when starting the SSM automation document.",
		//	              "properties": {
		//	                "Key": {
		//	                  "maxLength": 50,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "additionalProperties": false,
		//	                  "description": "Value of the dynamic parameter to set when starting the SSM automation document.",
		//	                  "properties": {
		//	                    "Variable": {
		//	                      "description": "The variable types used as dynamic parameter value when starting the SSM automation document.",
		//	                      "enum": [
		//	                        "INCIDENT_RECORD_ARN",
		//	                        "INVOLVED_RESOURCES"
		//	                      ],
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                }
		//	              },
		//	              "required": [
		//	                "Value",
		//	                "Key"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 200,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "Parameters": {
		//	            "default": [],
		//	            "description": "The parameters to set when starting the SSM automation document.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "description": "A parameter to set when starting the SSM automation document.",
		//	              "properties": {
		//	                "Key": {
		//	                  "maxLength": 50,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "Values": {
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "description": "A value of the parameter to set when starting the SSM automation document.",
		//	                    "maxLength": 10000,
		//	                    "type": "string"
		//	                  },
		//	                  "maxItems": 10,
		//	                  "type": "array",
		//	                  "uniqueItems": true
		//	                }
		//	              },
		//	              "required": [
		//	                "Values",
		//	                "Key"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 200,
		//	            "minItems": 1,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "RoleArn": {
		//	            "description": "The role ARN to use when starting the SSM automation document.",
		//	            "maxLength": 1000,
		//	            "pattern": "^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	            "type": "string"
		//	          },
		//	          "TargetAccount": {
		//	            "description": "The account type to use when starting the SSM automation document.",
		//	            "enum": [
		//	              "IMPACTED_ACCOUNT",
		//	              "RESPONSE_PLAN_OWNER_ACCOUNT"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "RoleArn",
		//	          "DocumentName"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"actions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: SsmAutomation
					"ssm_automation": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: DocumentName
							"document_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The document name to use when starting the SSM automation document.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DocumentVersion
							"document_version": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The version of the document to use when starting the SSM automation document.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DynamicParameters
							"dynamic_parameters": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Key
										"key": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Variable
												"variable": schema.StringAttribute{ /*START ATTRIBUTE*/
													Description: "The variable types used as dynamic parameter value when starting the SSM automation document.",
													Computed:    true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Description: "Value of the dynamic parameter to set when starting the SSM automation document.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "The parameters with dynamic values to set when starting the SSM automation document.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Parameters
							"parameters": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Key
										"key": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Values
										"values": schema.ListAttribute{ /*START ATTRIBUTE*/
											ElementType: types.StringType,
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "The parameters to set when starting the SSM automation document.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: RoleArn
							"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The role ARN to use when starting the SSM automation document.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: TargetAccount
							"target_account": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The account type to use when starting the SSM automation document.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The configuration to use when starting the SSM automation document.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of actions.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the response plan.",
		//	  "maxLength": 1000,
		//	  "pattern": "^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the response plan.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ChatChannel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The chat channel configuration.",
		//	  "properties": {
		//	    "ChatbotSns": {
		//	      "default": [],
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "The ARN of the Chatbot SNS topic.",
		//	        "maxLength": 1000,
		//	        "pattern": "^arn:aws(-(cn|us-gov))?:sns:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"chat_channel": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ChatbotSns
				"chatbot_sns": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The chat channel configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The display name of the response plan.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The display name of the response plan.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Engagements
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": [],
		//	  "description": "The list of engagements to use.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "The ARN of the contact.",
		//	    "maxLength": 1000,
		//	    "pattern": "^arn:aws(-(cn|us-gov))?:ssm-contacts:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"engagements": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The list of engagements to use.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IncidentTemplate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The incident template configuration.",
		//	  "properties": {
		//	    "DedupeString": {
		//	      "description": "The deduplication string.",
		//	      "maxLength": 1000,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "Impact": {
		//	      "description": "The impact value.",
		//	      "maximum": 5,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "IncidentTags": {
		//	      "default": [],
		//	      "description": "Tags that get applied to incidents created by the StartIncident API action.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "A key-value pair to tag a resource.",
		//	        "properties": {
		//	          "Key": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Value",
		//	          "Key"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 50,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "NotificationTargets": {
		//	      "default": [],
		//	      "description": "The list of notification targets.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "A notification target.",
		//	        "properties": {
		//	          "SnsTopicArn": {
		//	            "description": "The ARN of the Chatbot SNS topic.",
		//	            "maxLength": 1000,
		//	            "pattern": "^arn:aws(-(cn|us-gov))?:sns:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "maxItems": 10,
		//	      "type": "array"
		//	    },
		//	    "Summary": {
		//	      "description": "The summary string.",
		//	      "maxLength": 4000,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "Title": {
		//	      "description": "The title string.",
		//	      "maxLength": 200,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Title",
		//	    "Impact"
		//	  ],
		//	  "type": "object"
		//	}
		"incident_template": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DedupeString
				"dedupe_string": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The deduplication string.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Impact
				"impact": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The impact value.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IncidentTags
				"incident_tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Tags that get applied to incidents created by the StartIncident API action.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: NotificationTargets
				"notification_targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: SnsTopicArn
							"sns_topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The ARN of the Chatbot SNS topic.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The list of notification targets.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Summary
				"summary": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The summary string.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Title
				"title": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The title string.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The incident template configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Integrations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": [],
		//	  "description": "The list of integrations.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "oneOf": [
		//	      {
		//	        "required": [
		//	          "PagerDutyConfiguration"
		//	        ]
		//	      }
		//	    ],
		//	    "properties": {
		//	      "PagerDutyConfiguration": {
		//	        "additionalProperties": false,
		//	        "description": "The pagerDuty configuration to use when starting the incident.",
		//	        "properties": {
		//	          "Name": {
		//	            "description": "The name of the pagerDuty configuration.",
		//	            "maxLength": 200,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "PagerDutyIncidentConfiguration": {
		//	            "additionalProperties": false,
		//	            "description": "The pagerDuty incident configuration.",
		//	            "properties": {
		//	              "ServiceId": {
		//	                "description": "The pagerDuty serviceId.",
		//	                "maxLength": 200,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "ServiceId"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "SecretId": {
		//	            "description": "The AWS secrets manager secretId storing the pagerDuty token.",
		//	            "maxLength": 512,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name",
		//	          "SecretId",
		//	          "PagerDutyIncidentConfiguration"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"integrations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: PagerDutyConfiguration
					"pager_duty_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the pagerDuty configuration.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: PagerDutyIncidentConfiguration
							"pager_duty_incident_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ServiceId
									"service_id": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The pagerDuty serviceId.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The pagerDuty incident configuration.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: SecretId
							"secret_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The AWS secrets manager secretId storing the pagerDuty token.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The pagerDuty configuration to use when starting the incident.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of integrations.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the response plan.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the response plan.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": [],
		//	  "description": "The tags to apply to the response plan.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to tag a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to apply to the response plan.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SSMIncidents::ResponsePlan",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMIncidents::ResponsePlan").WithTerraformTypeName("awscc_ssmincidents_response_plan")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":                           "Actions",
		"arn":                               "Arn",
		"chat_channel":                      "ChatChannel",
		"chatbot_sns":                       "ChatbotSns",
		"dedupe_string":                     "DedupeString",
		"display_name":                      "DisplayName",
		"document_name":                     "DocumentName",
		"document_version":                  "DocumentVersion",
		"dynamic_parameters":                "DynamicParameters",
		"engagements":                       "Engagements",
		"impact":                            "Impact",
		"incident_tags":                     "IncidentTags",
		"incident_template":                 "IncidentTemplate",
		"integrations":                      "Integrations",
		"key":                               "Key",
		"name":                              "Name",
		"notification_targets":              "NotificationTargets",
		"pager_duty_configuration":          "PagerDutyConfiguration",
		"pager_duty_incident_configuration": "PagerDutyIncidentConfiguration",
		"parameters":                        "Parameters",
		"role_arn":                          "RoleArn",
		"secret_id":                         "SecretId",
		"service_id":                        "ServiceId",
		"sns_topic_arn":                     "SnsTopicArn",
		"ssm_automation":                    "SsmAutomation",
		"summary":                           "Summary",
		"tags":                              "Tags",
		"target_account":                    "TargetAccount",
		"title":                             "Title",
		"value":                             "Value",
		"values":                            "Values",
		"variable":                          "Variable",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
