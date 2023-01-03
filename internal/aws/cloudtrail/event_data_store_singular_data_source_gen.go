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
	registry.AddDataSourceFactory("awscc_cloudtrail_event_data_store", eventDataStoreDataSource)
}

// eventDataStoreDataSource returns the Terraform awscc_cloudtrail_event_data_store data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudTrail::EventDataStore resource.
func eventDataStoreDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdvancedEventSelectors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The advanced event selectors that were used to select events for the data store.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Advanced event selectors let you create fine-grained selectors for the following AWS CloudTrail event record ?elds. They help you control costs by logging only those events that are important to you.",
		//	    "properties": {
		//	      "FieldSelectors": {
		//	        "description": "Contains all selector statements in an advanced event selector.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "A single selector statement in an advanced event selector.",
		//	          "properties": {
		//	            "EndsWith": {
		//	              "description": "An operator that includes events that match the last few characters of the event record field specified as the value of Field.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "Equals": {
		//	              "description": "An operator that includes events that match the exact value of the event record field specified as the value of Field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "Field": {
		//	              "description": "A field in an event record on which to filter events to be logged. Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.",
		//	              "maxLength": 1000,
		//	              "minLength": 1,
		//	              "pattern": "([\\w|\\d|\\.|_]+)",
		//	              "type": "string"
		//	            },
		//	            "NotEndsWith": {
		//	              "description": "An operator that excludes events that match the last few characters of the event record field specified as the value of Field.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "NotEquals": {
		//	              "description": "An operator that excludes events that match the exact value of the event record field specified as the value of Field.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "NotStartsWith": {
		//	              "description": "An operator that excludes events that match the first few characters of the event record field specified as the value of Field.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "StartsWith": {
		//	              "description": "An operator that includes events that match the first few characters of the event record field specified as the value of Field.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            }
		//	          },
		//	          "required": [
		//	            "Field"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "minItems": 1,
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "Name": {
		//	        "description": "An optional, descriptive name for an advanced event selector, such as \"Log data events for only two S3 buckets\".",
		//	        "maxLength": 1000,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "FieldSelectors"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"advanced_event_selectors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: FieldSelectors
					"field_selectors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: EndsWith
								"ends_with": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that includes events that match the last few characters of the event record field specified as the value of Field.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Equals
								"equals": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that includes events that match the exact value of the event record field specified as the value of Field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Field
								"field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "A field in an event record on which to filter events to be logged. Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: NotEndsWith
								"not_ends_with": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that excludes events that match the last few characters of the event record field specified as the value of Field.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: NotEquals
								"not_equals": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that excludes events that match the exact value of the event record field specified as the value of Field.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: NotStartsWith
								"not_starts_with": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that excludes events that match the first few characters of the event record field specified as the value of Field.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: StartsWith
								"starts_with": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that includes events that match the first few characters of the event record field specified as the value of Field.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "Contains all selector statements in an advanced event selector.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "An optional, descriptive name for an advanced event selector, such as \"Log data events for only two S3 buckets\".",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The advanced event selectors that were used to select events for the data store.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedTimestamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of the event data store's creation.",
		//	  "type": "string"
		//	}
		"created_timestamp": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp of the event data store's creation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EventDataStoreArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the event data store.",
		//	  "type": "string"
		//	}
		"event_data_store_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the event data store.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MultiRegionEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the event data store includes events from all regions, or only from the region in which it was created.",
		//	  "type": "boolean"
		//	}
		"multi_region_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the event data store includes events from all regions, or only from the region in which it was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the event data store.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the event data store.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OrganizationEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates that an event data store is collecting logged events for an organization.",
		//	  "type": "boolean"
		//	}
		"organization_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates that an event data store is collecting logged events for an organization.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RetentionPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The retention period, in days.",
		//	  "type": "integer"
		//	}
		"retention_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The retention period, in days.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of an event data store. Values are ENABLED and PENDING_DELETION.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of an event data store. Values are ENABLED and PENDING_DELETION.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An arbitrary set of tags (key-value pairs) for this event data store.",
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
		// Property: TerminationProtectionEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the event data store is protected from termination.",
		//	  "type": "boolean"
		//	}
		"termination_protection_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the event data store is protected from termination.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedTimestamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.",
		//	  "type": "string"
		//	}
		"updated_timestamp": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudTrail::EventDataStore",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudTrail::EventDataStore").WithTerraformTypeName("awscc_cloudtrail_event_data_store")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"advanced_event_selectors":       "AdvancedEventSelectors",
		"created_timestamp":              "CreatedTimestamp",
		"ends_with":                      "EndsWith",
		"equals":                         "Equals",
		"event_data_store_arn":           "EventDataStoreArn",
		"field":                          "Field",
		"field_selectors":                "FieldSelectors",
		"key":                            "Key",
		"kms_key_id":                     "KmsKeyId",
		"multi_region_enabled":           "MultiRegionEnabled",
		"name":                           "Name",
		"not_ends_with":                  "NotEndsWith",
		"not_equals":                     "NotEquals",
		"not_starts_with":                "NotStartsWith",
		"organization_enabled":           "OrganizationEnabled",
		"retention_period":               "RetentionPeriod",
		"starts_with":                    "StartsWith",
		"status":                         "Status",
		"tags":                           "Tags",
		"termination_protection_enabled": "TerminationProtectionEnabled",
		"updated_timestamp":              "UpdatedTimestamp",
		"value":                          "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
