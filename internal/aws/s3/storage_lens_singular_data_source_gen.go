// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3_storage_lens", storageLensDataSource)
}

// storageLensDataSource returns the Terraform awscc_s3_storage_lens data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3::StorageLens resource.
func storageLensDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: StorageLensConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies the details of Amazon S3 Storage Lens configuration.",
		//	  "properties": {
		//	    "AccountLevel": {
		//	      "additionalProperties": false,
		//	      "description": "Account-level metrics configurations.",
		//	      "properties": {
		//	        "ActivityMetrics": {
		//	          "additionalProperties": false,
		//	          "description": "Enables activity metrics.",
		//	          "properties": {
		//	            "IsEnabled": {
		//	              "description": "Specifies whether activity metrics are enabled or disabled.",
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "AdvancedCostOptimizationMetrics": {
		//	          "additionalProperties": false,
		//	          "description": "Enables advanced cost optimization metrics.",
		//	          "properties": {
		//	            "IsEnabled": {
		//	              "description": "Specifies whether advanced cost optimization metrics are enabled or disabled.",
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "AdvancedDataProtectionMetrics": {
		//	          "additionalProperties": false,
		//	          "description": "Enables advanced data protection metrics.",
		//	          "properties": {
		//	            "IsEnabled": {
		//	              "description": "Specifies whether advanced data protection metrics are enabled or disabled.",
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "BucketLevel": {
		//	          "additionalProperties": false,
		//	          "description": "Bucket-level metrics configurations.",
		//	          "properties": {
		//	            "ActivityMetrics": {
		//	              "additionalProperties": false,
		//	              "description": "Enables activity metrics.",
		//	              "properties": {
		//	                "IsEnabled": {
		//	                  "description": "Specifies whether activity metrics are enabled or disabled.",
		//	                  "type": "boolean"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "AdvancedCostOptimizationMetrics": {
		//	              "additionalProperties": false,
		//	              "description": "Enables advanced cost optimization metrics.",
		//	              "properties": {
		//	                "IsEnabled": {
		//	                  "description": "Specifies whether advanced cost optimization metrics are enabled or disabled.",
		//	                  "type": "boolean"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "AdvancedDataProtectionMetrics": {
		//	              "additionalProperties": false,
		//	              "description": "Enables advanced data protection metrics.",
		//	              "properties": {
		//	                "IsEnabled": {
		//	                  "description": "Specifies whether advanced data protection metrics are enabled or disabled.",
		//	                  "type": "boolean"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "DetailedStatusCodesMetrics": {
		//	              "additionalProperties": false,
		//	              "description": "Enables detailed status codes metrics.",
		//	              "properties": {
		//	                "IsEnabled": {
		//	                  "description": "Specifies whether detailed status codes metrics are enabled or disabled.",
		//	                  "type": "boolean"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "PrefixLevel": {
		//	              "additionalProperties": false,
		//	              "description": "Prefix-level metrics configurations.",
		//	              "properties": {
		//	                "StorageMetrics": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "IsEnabled": {
		//	                      "description": "Specifies whether prefix-level storage metrics are enabled or disabled.",
		//	                      "type": "boolean"
		//	                    },
		//	                    "SelectionCriteria": {
		//	                      "additionalProperties": false,
		//	                      "description": "Selection criteria for prefix-level metrics.",
		//	                      "properties": {
		//	                        "Delimiter": {
		//	                          "description": "Delimiter to divide S3 key into hierarchy of prefixes.",
		//	                          "type": "string"
		//	                        },
		//	                        "MaxDepth": {
		//	                          "description": "Max depth of prefixes of S3 key that Amazon S3 Storage Lens will analyze.",
		//	                          "type": "integer"
		//	                        },
		//	                        "MinStorageBytesPercentage": {
		//	                          "description": "The minimum storage bytes threshold for the prefixes to be included in the analysis.",
		//	                          "type": "number"
		//	                        }
		//	                      },
		//	                      "type": "object"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                }
		//	              },
		//	              "required": [
		//	                "StorageMetrics"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "DetailedStatusCodesMetrics": {
		//	          "additionalProperties": false,
		//	          "description": "Enables detailed status codes metrics.",
		//	          "properties": {
		//	            "IsEnabled": {
		//	              "description": "Specifies whether detailed status codes metrics are enabled or disabled.",
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "BucketLevel"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "AwsOrg": {
		//	      "additionalProperties": false,
		//	      "description": "The AWS Organizations ARN to use in the Amazon S3 Storage Lens configuration.",
		//	      "properties": {
		//	        "Arn": {
		//	          "description": "The Amazon Resource Name (ARN) of the specified resource.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Arn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "DataExport": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies how Amazon S3 Storage Lens metrics should be exported.",
		//	      "properties": {
		//	        "CloudWatchMetrics": {
		//	          "additionalProperties": false,
		//	          "description": "CloudWatch metrics settings for the Amazon S3 Storage Lens metrics export.",
		//	          "properties": {
		//	            "IsEnabled": {
		//	              "description": "Specifies whether CloudWatch metrics are enabled or disabled.",
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "IsEnabled"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "S3BucketDestination": {
		//	          "additionalProperties": false,
		//	          "description": "S3 bucket destination settings for the Amazon S3 Storage Lens metrics export.",
		//	          "properties": {
		//	            "AccountId": {
		//	              "description": "The AWS account ID that owns the destination S3 bucket.",
		//	              "type": "string"
		//	            },
		//	            "Arn": {
		//	              "description": "The ARN of the bucket to which Amazon S3 Storage Lens exports will be placed.",
		//	              "type": "string"
		//	            },
		//	            "Encryption": {
		//	              "description": "Configures the server-side encryption for Amazon S3 Storage Lens report files with either S3-managed keys (SSE-S3) or KMS-managed keys (SSE-KMS).",
		//	              "properties": {
		//	                "SSEKMS": {
		//	                  "additionalProperties": false,
		//	                  "description": "AWS KMS server-side encryption.",
		//	                  "properties": {
		//	                    "KeyId": {
		//	                      "description": "The ARN of the KMS key to use for encryption.",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "required": [
		//	                    "KeyId"
		//	                  ],
		//	                  "type": "object"
		//	                },
		//	                "SSES3": {
		//	                  "additionalProperties": false,
		//	                  "description": "S3 default server-side encryption.",
		//	                  "type": "object"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "Format": {
		//	              "description": "Specifies the file format to use when exporting Amazon S3 Storage Lens metrics export.",
		//	              "enum": [
		//	                "CSV",
		//	                "Parquet"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "OutputSchemaVersion": {
		//	              "description": "The version of the output schema to use when exporting Amazon S3 Storage Lens metrics.",
		//	              "enum": [
		//	                "V_1"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "Prefix": {
		//	              "description": "The prefix to use for Amazon S3 Storage Lens export.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "OutputSchemaVersion",
		//	            "Format",
		//	            "AccountId",
		//	            "Arn"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Exclude": {
		//	      "additionalProperties": false,
		//	      "description": "S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.",
		//	      "properties": {
		//	        "Buckets": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "The Amazon Resource Name (ARN) of the specified resource.",
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "Regions": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "An AWS Region.",
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Id": {
		//	      "description": "The ID that identifies the Amazon S3 Storage Lens configuration.",
		//	      "maxLength": 64,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9\\-_.]+$",
		//	      "type": "string"
		//	    },
		//	    "Include": {
		//	      "additionalProperties": false,
		//	      "description": "S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.",
		//	      "properties": {
		//	        "Buckets": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "The Amazon Resource Name (ARN) of the specified resource.",
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "Regions": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "An AWS Region.",
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "IsEnabled": {
		//	      "description": "Specifies whether the Amazon S3 Storage Lens configuration is enabled or disabled.",
		//	      "type": "boolean"
		//	    },
		//	    "StorageLensArn": {
		//	      "description": "The ARN for the Amazon S3 Storage Lens configuration.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Id",
		//	    "AccountLevel",
		//	    "IsEnabled"
		//	  ],
		//	  "type": "object"
		//	}
		"storage_lens_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AccountLevel
				"account_level": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ActivityMetrics
						"activity_metrics": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: IsEnabled
								"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies whether activity metrics are enabled or disabled.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Enables activity metrics.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: AdvancedCostOptimizationMetrics
						"advanced_cost_optimization_metrics": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: IsEnabled
								"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies whether advanced cost optimization metrics are enabled or disabled.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Enables advanced cost optimization metrics.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: AdvancedDataProtectionMetrics
						"advanced_data_protection_metrics": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: IsEnabled
								"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies whether advanced data protection metrics are enabled or disabled.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Enables advanced data protection metrics.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: BucketLevel
						"bucket_level": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ActivityMetrics
								"activity_metrics": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: IsEnabled
										"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Description: "Specifies whether activity metrics are enabled or disabled.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Enables activity metrics.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: AdvancedCostOptimizationMetrics
								"advanced_cost_optimization_metrics": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: IsEnabled
										"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Description: "Specifies whether advanced cost optimization metrics are enabled or disabled.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Enables advanced cost optimization metrics.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: AdvancedDataProtectionMetrics
								"advanced_data_protection_metrics": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: IsEnabled
										"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Description: "Specifies whether advanced data protection metrics are enabled or disabled.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Enables advanced data protection metrics.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: DetailedStatusCodesMetrics
								"detailed_status_codes_metrics": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: IsEnabled
										"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Description: "Specifies whether detailed status codes metrics are enabled or disabled.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Enables detailed status codes metrics.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: PrefixLevel
								"prefix_level": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: StorageMetrics
										"storage_metrics": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: IsEnabled
												"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
													Description: "Specifies whether prefix-level storage metrics are enabled or disabled.",
													Computed:    true,
												}, /*END ATTRIBUTE*/
												// Property: SelectionCriteria
												"selection_criteria": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
													Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
														// Property: Delimiter
														"delimiter": schema.StringAttribute{ /*START ATTRIBUTE*/
															Description: "Delimiter to divide S3 key into hierarchy of prefixes.",
															Computed:    true,
														}, /*END ATTRIBUTE*/
														// Property: MaxDepth
														"max_depth": schema.Int64Attribute{ /*START ATTRIBUTE*/
															Description: "Max depth of prefixes of S3 key that Amazon S3 Storage Lens will analyze.",
															Computed:    true,
														}, /*END ATTRIBUTE*/
														// Property: MinStorageBytesPercentage
														"min_storage_bytes_percentage": schema.Float64Attribute{ /*START ATTRIBUTE*/
															Description: "The minimum storage bytes threshold for the prefixes to be included in the analysis.",
															Computed:    true,
														}, /*END ATTRIBUTE*/
													}, /*END SCHEMA*/
													Description: "Selection criteria for prefix-level metrics.",
													Computed:    true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Prefix-level metrics configurations.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Bucket-level metrics configurations.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: DetailedStatusCodesMetrics
						"detailed_status_codes_metrics": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: IsEnabled
								"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies whether detailed status codes metrics are enabled or disabled.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Enables detailed status codes metrics.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Account-level metrics configurations.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AwsOrg
				"aws_org": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Arn
						"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The Amazon Resource Name (ARN) of the specified resource.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The AWS Organizations ARN to use in the Amazon S3 Storage Lens configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DataExport
				"data_export": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CloudWatchMetrics
						"cloudwatch_metrics": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: IsEnabled
								"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies whether CloudWatch metrics are enabled or disabled.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "CloudWatch metrics settings for the Amazon S3 Storage Lens metrics export.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: S3BucketDestination
						"s3_bucket_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AccountId
								"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The AWS account ID that owns the destination S3 bucket.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Arn
								"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The ARN of the bucket to which Amazon S3 Storage Lens exports will be placed.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Encryption
								"encryption": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: SSEKMS
										"ssekms": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: KeyId
												"key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
													Description: "The ARN of the KMS key to use for encryption.",
													Computed:    true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Description: "AWS KMS server-side encryption.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: SSES3
										"sses3": schema.MapAttribute{ /*START ATTRIBUTE*/
											ElementType: types.StringType,
											Description: "S3 default server-side encryption.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Configures the server-side encryption for Amazon S3 Storage Lens report files with either S3-managed keys (SSE-S3) or KMS-managed keys (SSE-KMS).",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Format
								"format": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies the file format to use when exporting Amazon S3 Storage Lens metrics export.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: OutputSchemaVersion
								"output_schema_version": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The version of the output schema to use when exporting Amazon S3 Storage Lens metrics.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Prefix
								"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The prefix to use for Amazon S3 Storage Lens export.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "S3 bucket destination settings for the Amazon S3 Storage Lens metrics export.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies how Amazon S3 Storage Lens metrics should be exported.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Exclude
				"exclude": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Buckets
						"buckets": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Regions
						"regions": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Id
				"id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID that identifies the Amazon S3 Storage Lens configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Include
				"include": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Buckets
						"buckets": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Regions
						"regions": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IsEnabled
				"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether the Amazon S3 Storage Lens configuration is enabled or disabled.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: StorageLensArn
				"storage_lens_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN for the Amazon S3 Storage Lens configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies the details of Amazon S3 Storage Lens configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
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
			Description: "A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::S3::StorageLens",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::StorageLens").WithTerraformTypeName("awscc_s3_storage_lens")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                         "AccountId",
		"account_level":                      "AccountLevel",
		"activity_metrics":                   "ActivityMetrics",
		"advanced_cost_optimization_metrics": "AdvancedCostOptimizationMetrics",
		"advanced_data_protection_metrics":   "AdvancedDataProtectionMetrics",
		"arn":                                "Arn",
		"aws_org":                            "AwsOrg",
		"bucket_level":                       "BucketLevel",
		"buckets":                            "Buckets",
		"cloudwatch_metrics":                 "CloudWatchMetrics",
		"data_export":                        "DataExport",
		"delimiter":                          "Delimiter",
		"detailed_status_codes_metrics":      "DetailedStatusCodesMetrics",
		"encryption":                         "Encryption",
		"exclude":                            "Exclude",
		"format":                             "Format",
		"id":                                 "Id",
		"include":                            "Include",
		"is_enabled":                         "IsEnabled",
		"key":                                "Key",
		"key_id":                             "KeyId",
		"max_depth":                          "MaxDepth",
		"min_storage_bytes_percentage":       "MinStorageBytesPercentage",
		"output_schema_version":              "OutputSchemaVersion",
		"prefix":                             "Prefix",
		"prefix_level":                       "PrefixLevel",
		"regions":                            "Regions",
		"s3_bucket_destination":              "S3BucketDestination",
		"selection_criteria":                 "SelectionCriteria",
		"ssekms":                             "SSEKMS",
		"sses3":                              "SSES3",
		"storage_lens_arn":                   "StorageLensArn",
		"storage_lens_configuration":         "StorageLensConfiguration",
		"storage_metrics":                    "StorageMetrics",
		"tags":                               "Tags",
		"value":                              "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
