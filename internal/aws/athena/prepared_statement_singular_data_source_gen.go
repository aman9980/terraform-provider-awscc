// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package athena

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_athena_prepared_statement", preparedStatementDataSource)
}

// preparedStatementDataSource returns the Terraform awscc_athena_prepared_statement data source.
// This Terraform data source corresponds to the CloudFormation AWS::Athena::PreparedStatement resource.
func preparedStatementDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the prepared statement.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the prepared statement.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: QueryStatement
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The query string for the prepared statement.",
		//	  "maxLength": 262144,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"query_statement": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The query string for the prepared statement.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StatementName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the prepared statement.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"statement_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the prepared statement.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WorkGroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the workgroup to which the prepared statement belongs.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"work_group": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the workgroup to which the prepared statement belongs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Athena::PreparedStatement",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Athena::PreparedStatement").WithTerraformTypeName("awscc_athena_prepared_statement")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":     "Description",
		"query_statement": "QueryStatement",
		"statement_name":  "StatementName",
		"work_group":      "WorkGroup",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
