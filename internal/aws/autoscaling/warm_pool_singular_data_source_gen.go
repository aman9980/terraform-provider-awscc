// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_autoscaling_warm_pool", warmPoolDataSource)
}

// warmPoolDataSource returns the Terraform awscc_autoscaling_warm_pool data source.
// This Terraform data source corresponds to the CloudFormation AWS::AutoScaling::WarmPool resource.
func warmPoolDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AutoScalingGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"auto_scaling_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceReusePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ReuseOnScaleIn": {
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"instance_reuse_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ReuseOnScaleIn
				"reuse_on_scale_in": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MaxGroupPreparedCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"max_group_prepared_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MinSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"min_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PoolState
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"pool_state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AutoScaling::WarmPool",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AutoScaling::WarmPool").WithTerraformTypeName("awscc_autoscaling_warm_pool")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_scaling_group_name":     "AutoScalingGroupName",
		"instance_reuse_policy":       "InstanceReusePolicy",
		"max_group_prepared_capacity": "MaxGroupPreparedCapacity",
		"min_size":                    "MinSize",
		"pool_state":                  "PoolState",
		"reuse_on_scale_in":           "ReuseOnScaleIn",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
