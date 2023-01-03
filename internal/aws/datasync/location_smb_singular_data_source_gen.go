// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_datasync_location_smb", locationSMBDataSource)
}

// locationSMBDataSource returns the Terraform awscc_datasync_location_smb data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataSync::LocationSMB resource.
func locationSMBDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AgentArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 128,
		//	    "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 4,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"agent_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Windows domain that the SMB server belongs to.",
		//	  "maxLength": 253,
		//	  "pattern": "^([A-Za-z0-9]+[A-Za-z0-9-.]*)*[A-Za-z0-9-]*[A-Za-z0-9]$",
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Windows domain that the SMB server belongs to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the SMB location that is created.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"location_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the SMB location that is created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocationUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the SMB location that was described.",
		//	  "maxLength": 4356,
		//	  "pattern": "^(efs|nfs|s3|smb|fsxw)://[a-zA-Z0-9./\\-]+$",
		//	  "type": "string"
		//	}
		"location_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the SMB location that was described.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MountOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "default": {
		//	    "Version": "AUTOMATIC"
		//	  },
		//	  "description": "The mount options used by DataSync to access the SMB server.",
		//	  "properties": {
		//	    "Version": {
		//	      "description": "The specific SMB version that you want DataSync to use to mount your SMB share.",
		//	      "enum": [
		//	        "AUTOMATIC",
		//	        "SMB2",
		//	        "SMB3"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"mount_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The specific SMB version that you want DataSync to use to mount your SMB share.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The mount options used by DataSync to access the SMB server.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Password
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
		//	  "maxLength": 104,
		//	  "pattern": "^.{0,104}$",
		//	  "type": "string"
		//	}
		"password": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServerHostname
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.",
		//	  "maxLength": 255,
		//	  "pattern": "^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$",
		//	  "type": "string"
		//	}
		"server_hostname": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Subdirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination",
		//	  "maxLength": 4096,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
		//	  "type": "string"
		//	}
		"subdirectory": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
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
						Description: "The key for an AWS resource tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for an AWS resource tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: User
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
		//	  "maxLength": 104,
		//	  "pattern": "^[^\\x5B\\x5D\\\\/:;|=,+*?]{1,104}$",
		//	  "type": "string"
		//	}
		"user": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DataSync::LocationSMB",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationSMB").WithTerraformTypeName("awscc_datasync_location_smb")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"agent_arns":      "AgentArns",
		"domain":          "Domain",
		"key":             "Key",
		"location_arn":    "LocationArn",
		"location_uri":    "LocationUri",
		"mount_options":   "MountOptions",
		"password":        "Password",
		"server_hostname": "ServerHostname",
		"subdirectory":    "Subdirectory",
		"tags":            "Tags",
		"user":            "User",
		"value":           "Value",
		"version":         "Version",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
