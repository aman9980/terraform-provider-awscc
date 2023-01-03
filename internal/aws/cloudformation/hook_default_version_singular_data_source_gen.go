// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudformation_hook_default_version", hookDefaultVersionDataSource)
}

// hookDefaultVersionDataSource returns the Terraform awscc_cloudformation_hook_default_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFormation::HookDefaultVersion resource.
func hookDefaultVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion",
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/.+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TypeName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
		//	  "pattern": "^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$",
		//	  "type": "string"
		//	}
		"type_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TypeVersionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the type version.",
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/.+$",
		//	  "type": "string"
		//	}
		"type_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the type version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of an existing version of the hook to set as the default.",
		//	  "pattern": "^[A-Za-z0-9-]{1,128}$",
		//	  "type": "string"
		//	}
		"version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of an existing version of the hook to set as the default.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFormation::HookDefaultVersion",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::HookDefaultVersion").WithTerraformTypeName("awscc_cloudformation_hook_default_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":              "Arn",
		"type_name":        "TypeName",
		"type_version_arn": "TypeVersionArn",
		"version_id":       "VersionId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
