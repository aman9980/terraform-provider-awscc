// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3objectlambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3objectlambda_access_point_policy", accessPointPolicyDataSource)
}

// accessPointPolicyDataSource returns the Terraform awscc_s3objectlambda_access_point_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3ObjectLambda::AccessPointPolicy resource.
func accessPointPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ObjectLambdaAccessPoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.",
		//	  "maxLength": 45,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$",
		//	  "type": "string"
		//	}
		"object_lambda_access_point": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. ",
		//	  "type": "object"
		//	}
		"policy_document": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::S3ObjectLambda::AccessPointPolicy",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3ObjectLambda::AccessPointPolicy").WithTerraformTypeName("awscc_s3objectlambda_access_point_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"object_lambda_access_point": "ObjectLambdaAccessPoint",
		"policy_document":            "PolicyDocument",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
