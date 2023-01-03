// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_servicecatalogappregistry_application", applicationDataSource)
}

// applicationDataSource returns the Terraform awscc_servicecatalogappregistry_application data source.
// This Terraform data source corresponds to the CloudFormation AWS::ServiceCatalogAppRegistry::Application resource.
func applicationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "arn:aws[-a-z]*:servicecatalog:[a-z]{2}(-gov)?-[a-z]+-\\d:\\d{12}:/applications/[a-z0-9]+",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the application. ",
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the application. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "[a-z0-9]{26}",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the application. ",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "\\w+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the application. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ServiceCatalogAppRegistry::Application",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalogAppRegistry::Application").WithTerraformTypeName("awscc_servicecatalogappregistry_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":         "Arn",
		"description": "Description",
		"id":          "Id",
		"name":        "Name",
		"tags":        "Tags",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
