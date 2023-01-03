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
	registry.AddDataSourceFactory("awscc_apigateway_documentation_version", documentationVersionDataSource)
}

// documentationVersionDataSource returns the Terraform awscc_apigateway_documentation_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::DocumentationVersion resource.
func documentationVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the API documentation snapshot.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the API documentation snapshot.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DocumentationVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version identifier of the API documentation snapshot.",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"documentation_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version identifier of the API documentation snapshot.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the API.",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the API.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::DocumentationVersion",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::DocumentationVersion").WithTerraformTypeName("awscc_apigateway_documentation_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":           "Description",
		"documentation_version": "DocumentationVersion",
		"rest_api_id":           "RestApiId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
