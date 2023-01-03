// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_rekognition_project", projectDataSource)
}

// projectDataSource returns the Terraform awscc_rekognition_project data source.
// This Terraform data source corresponds to the CloudFormation AWS::Rekognition::Project resource.
func projectDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "pattern": "(^arn:[a-z\\d-]+:rekognition:[a-z\\d-]+:\\d{12}:project/[a-zA-Z0-9_.\\-]{1,255}/[0-9]+$)",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the project",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9][a-zA-Z0-9_\\-]*",
		//	  "type": "string"
		//	}
		"project_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the project",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Rekognition::Project",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Rekognition::Project").WithTerraformTypeName("awscc_rekognition_project")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":          "Arn",
		"project_name": "ProjectName",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
