// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rolesanywhere

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_rolesanywhere_profile", profileDataSource)
}

// profileDataSource returns the Terraform awscc_rolesanywhere_profile data source.
// This Terraform data source corresponds to the CloudFormation AWS::RolesAnywhere::Profile resource.
func profileDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DurationSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 43200,
		//	  "minimum": 900,
		//	  "type": "number"
		//	}
		"duration_seconds": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Enabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ManagedPolicyArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"managed_policy_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ProfileArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"profile_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ProfileId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "[a-f0-9]{8}-([a-z0-9]{4}-){3}[a-z0-9]{12}",
		//	  "type": "string"
		//	}
		"profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RequireInstanceProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"require_instance_properties": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "maxLength": 1011,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"role_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SessionPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"session_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
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
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
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
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::RolesAnywhere::Profile",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RolesAnywhere::Profile").WithTerraformTypeName("awscc_rolesanywhere_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"duration_seconds":            "DurationSeconds",
		"enabled":                     "Enabled",
		"key":                         "Key",
		"managed_policy_arns":         "ManagedPolicyArns",
		"name":                        "Name",
		"profile_arn":                 "ProfileArn",
		"profile_id":                  "ProfileId",
		"require_instance_properties": "RequireInstanceProperties",
		"role_arns":                   "RoleArns",
		"session_policy":              "SessionPolicy",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
