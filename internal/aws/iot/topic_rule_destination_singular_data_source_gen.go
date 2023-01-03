// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_topic_rule_destination", topicRuleDestinationDataSource)
}

// topicRuleDestinationDataSource returns the Terraform awscc_iot_topic_rule_destination data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::TopicRuleDestination resource.
func topicRuleDestinationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN).",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HttpUrlProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "HTTP URL destination properties.",
		//	  "properties": {
		//	    "ConfirmationUrl": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"http_url_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ConfirmationUrl
				"confirmation_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "HTTP URL destination properties.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the TopicRuleDestination.",
		//	  "enum": [
		//	    "ENABLED",
		//	    "IN_PROGRESS",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the TopicRuleDestination.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StatusReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The reasoning for the current status of the TopicRuleDestination.",
		//	  "type": "string"
		//	}
		"status_reason": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The reasoning for the current status of the TopicRuleDestination.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "VPC destination properties.",
		//	  "properties": {
		//	    "RoleArn": {
		//	      "type": "string"
		//	    },
		//	    "SecurityGroups": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "SubnetIds": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "VpcId": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"vpc_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SecurityGroups
				"security_groups": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SubnetIds
				"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VpcId
				"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "VPC destination properties.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoT::TopicRuleDestination",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::TopicRuleDestination").WithTerraformTypeName("awscc_iot_topic_rule_destination")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"confirmation_url":    "ConfirmationUrl",
		"http_url_properties": "HttpUrlProperties",
		"role_arn":            "RoleArn",
		"security_groups":     "SecurityGroups",
		"status":              "Status",
		"status_reason":       "StatusReason",
		"subnet_ids":          "SubnetIds",
		"vpc_id":              "VpcId",
		"vpc_properties":      "VpcProperties",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
