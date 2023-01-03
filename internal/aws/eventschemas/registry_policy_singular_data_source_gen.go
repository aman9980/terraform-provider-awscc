// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package eventschemas

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_eventschemas_registry_policy", registryPolicyDataSource)
}

// registryPolicyDataSource returns the Terraform awscc_eventschemas_registry_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::EventSchemas::RegistryPolicy resource.
func registryPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Policy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "object"
		//	}
		"policy": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RegistryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"registry_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RevisionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"revision_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EventSchemas::RegistryPolicy",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EventSchemas::RegistryPolicy").WithTerraformTypeName("awscc_eventschemas_registry_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":            "Id",
		"policy":        "Policy",
		"registry_name": "RegistryName",
		"revision_id":   "RevisionId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
