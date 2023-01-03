// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_resource", resourceDataSource)
}

// resourceDataSource returns the Terraform awscc_apigateway_resource data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::Resource resource.
func resourceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ParentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The parent resource's identifier.",
		//	  "type": "string"
		//	}
		"parent_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The parent resource's identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PathPart
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The last path segment for this resource.",
		//	  "type": "string"
		//	}
		"path_part": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The last path segment for this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique primary identifier for a Resource",
		//	  "type": "string"
		//	}
		"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique primary identifier for a Resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the RestApi resource in which you want to create this resource..",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the RestApi resource in which you want to create this resource..",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::Resource",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Resource").WithTerraformTypeName("awscc_apigateway_resource")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"parent_id":   "ParentId",
		"path_part":   "PathPart",
		"resource_id": "ResourceId",
		"rest_api_id": "RestApiId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
