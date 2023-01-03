// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudtrail_trail", trailDataSource)
}

// trailDataSource returns the Terraform awscc_cloudtrail_trail data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudTrail::Trail resource.
func trailDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CloudWatchLogsLogGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn.",
		//	  "type": "string"
		//	}
		"cloudwatch_logs_log_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CloudWatchLogsRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.",
		//	  "type": "string"
		//	}
		"cloudwatch_logs_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnableLogFileValidation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether log file validation is enabled. The default is false.",
		//	  "type": "boolean"
		//	}
		"enable_log_file_validation": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether log file validation is enabled. The default is false.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EventSelectors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The type of email sending events to publish to the event destination.",
		//	    "properties": {
		//	      "DataResources": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "CloudTrail supports data event logging for Amazon S3 objects and AWS Lambda functions. You can specify up to 250 resources for an individual event selector, but the total number of data resources cannot exceed 250 across all event selectors in a trail. This limit does not apply if you configure resource logging for all data events.",
		//	          "properties": {
		//	            "Type": {
		//	              "description": "The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.",
		//	              "type": "string"
		//	            },
		//	            "Values": {
		//	              "description": "An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            }
		//	          },
		//	          "required": [
		//	            "Type"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "ExcludeManagementEventSources": {
		//	        "description": "An optional list of service event sources from which you do not want management events to be logged on your trail. In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing \"kms.amazonaws.com\". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "IncludeManagementEvents": {
		//	        "description": "Specify if you want your event selector to include management events for your trail.",
		//	        "type": "boolean"
		//	      },
		//	      "ReadWriteType": {
		//	        "description": "Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.",
		//	        "enum": [
		//	          "All",
		//	          "ReadOnly",
		//	          "WriteOnly"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 5,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"event_selectors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DataResources
					"data_resources": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Type
								"type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Values
								"values": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ExcludeManagementEventSources
					"exclude_management_event_sources": schema.SetAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "An optional list of service event sources from which you do not want management events to be logged on your trail. In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing \"kms.amazonaws.com\". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: IncludeManagementEvents
					"include_management_events": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "Specify if you want your event selector to include management events for your trail.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ReadWriteType
					"read_write_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IncludeGlobalServiceEvents
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the trail is publishing events from global services such as IAM to the log files.",
		//	  "type": "boolean"
		//	}
		"include_global_service_events": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the trail is publishing events from global services such as IAM to the log files.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InsightSelectors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A string that contains insight types that are logged on a trail.",
		//	    "properties": {
		//	      "InsightType": {
		//	        "description": "The type of insight to log on a trail.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"insight_selectors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: InsightType
					"insight_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of insight to log on a trail.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IsLogging
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether the CloudTrail is currently logging AWS API calls.",
		//	  "type": "boolean"
		//	}
		"is_logging": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether the CloudTrail is currently logging AWS API calls.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IsMultiRegionTrail
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.",
		//	  "type": "boolean"
		//	}
		"is_multi_region_trail": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IsOrganizationTrail
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.",
		//	  "type": "boolean"
		//	}
		"is_organization_trail": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KMSKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: S3BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.",
		//	  "type": "string"
		//	}
		"s3_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: S3KeyPrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.",
		//	  "maxLength": 200,
		//	  "type": "string"
		//	}
		"s3_key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SnsTopicArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"sns_topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SnsTopicName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"sns_topic_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An arbitrary set of tags (key-value pairs) for this trail.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TrailName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 3,
		//	  "pattern": "(^[a-zA-Z0-9]$)|(^[a-zA-Z0-9]([a-zA-Z0-9\\._-])*[a-zA-Z0-9]$)",
		//	  "type": "string"
		//	}
		"trail_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudTrail::Trail",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudTrail::Trail").WithTerraformTypeName("awscc_cloudtrail_trail")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                              "Arn",
		"cloudwatch_logs_log_group_arn":    "CloudWatchLogsLogGroupArn",
		"cloudwatch_logs_role_arn":         "CloudWatchLogsRoleArn",
		"data_resources":                   "DataResources",
		"enable_log_file_validation":       "EnableLogFileValidation",
		"event_selectors":                  "EventSelectors",
		"exclude_management_event_sources": "ExcludeManagementEventSources",
		"include_global_service_events":    "IncludeGlobalServiceEvents",
		"include_management_events":        "IncludeManagementEvents",
		"insight_selectors":                "InsightSelectors",
		"insight_type":                     "InsightType",
		"is_logging":                       "IsLogging",
		"is_multi_region_trail":            "IsMultiRegionTrail",
		"is_organization_trail":            "IsOrganizationTrail",
		"key":                              "Key",
		"kms_key_id":                       "KMSKeyId",
		"read_write_type":                  "ReadWriteType",
		"s3_bucket_name":                   "S3BucketName",
		"s3_key_prefix":                    "S3KeyPrefix",
		"sns_topic_arn":                    "SnsTopicArn",
		"sns_topic_name":                   "SnsTopicName",
		"tags":                             "Tags",
		"trail_name":                       "TrailName",
		"type":                             "Type",
		"value":                            "Value",
		"values":                           "Values",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
