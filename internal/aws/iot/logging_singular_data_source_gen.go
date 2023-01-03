// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_logging", loggingDataSource)
}

// loggingDataSource returns the Terraform awscc_iot_logging data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::Logging resource.
func loggingDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "pattern": "^[0-9]{12}$",
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultLogLevel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
		//	  "enum": [
		//	    "ERROR",
		//	    "WARN",
		//	    "INFO",
		//	    "DEBUG",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"default_log_level": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the role that allows IoT to write to Cloudwatch logs.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the role that allows IoT to write to Cloudwatch logs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoT::Logging",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::Logging").WithTerraformTypeName("awscc_iot_logging")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":        "AccountId",
		"default_log_level": "DefaultLogLevel",
		"role_arn":          "RoleArn",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
