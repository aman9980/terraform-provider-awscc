// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_kinesisvideo_stream", streamDataSource)
}

// streamDataSource returns the Terraform awscc_kinesisvideo_stream data source.
// This Terraform data source corresponds to the CloudFormation AWS::KinesisVideo::Stream resource.
func streamDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Kinesis Video stream.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Kinesis Video stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataRetentionInHours
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of hours till which Kinesis Video will retain the data in the stream",
		//	  "maximum": 87600,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"data_retention_in_hours": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of hours till which Kinesis Video will retain the data in the stream",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeviceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the device that is writing to the stream.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_.-]+",
		//	  "type": "string"
		//	}
		"device_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the device that is writing to the stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": ".+",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MediaType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The media type of the stream. Consumers of the stream can use this information when processing the stream.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[\\w\\-\\.\\+]+/[\\w\\-\\.\\+]+(,[\\w\\-\\.\\+]+/[\\w\\-\\.\\+]+)*",
		//	  "type": "string"
		//	}
		"media_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The media type of the stream. Consumers of the stream can use this information when processing the stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Kinesis Video stream.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_.-]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Kinesis Video stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs associated with the Kinesis Video Stream.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associated with the Kinesis Video Stream.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. Specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. Specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
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
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. Specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. Specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs associated with the Kinesis Video Stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::KinesisVideo::Stream",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::KinesisVideo::Stream").WithTerraformTypeName("awscc_kinesisvideo_stream")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"data_retention_in_hours": "DataRetentionInHours",
		"device_name":             "DeviceName",
		"key":                     "Key",
		"kms_key_id":              "KmsKeyId",
		"media_type":              "MediaType",
		"name":                    "Name",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
