// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_volume", volumeDataSource)
}

// volumeDataSource returns the Terraform awscc_ec2_volume data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::Volume resource.
func volumeDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AutoEnableIO
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"auto_enable_io": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Encrypted
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Iops
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"iops": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MultiAttachEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"multi_attach_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: OutpostArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"outpost_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Size
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SnapshotId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"snapshot_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Throughput
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"throughput": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: VolumeType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"volume_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::Volume",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Volume").WithTerraformTypeName("awscc_ec2_volume")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_enable_io":       "AutoEnableIO",
		"availability_zone":    "AvailabilityZone",
		"encrypted":            "Encrypted",
		"id":                   "Id",
		"iops":                 "Iops",
		"key":                  "Key",
		"kms_key_id":           "KmsKeyId",
		"multi_attach_enabled": "MultiAttachEnabled",
		"outpost_arn":          "OutpostArn",
		"size":                 "Size",
		"snapshot_id":          "SnapshotId",
		"tags":                 "Tags",
		"throughput":           "Throughput",
		"value":                "Value",
		"volume_type":          "VolumeType",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
