// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_authorizer", authorizerDataSource)
}

// authorizerDataSource returns the Terraform awscc_apigateway_authorizer data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::Authorizer resource.
func authorizerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AuthType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Optional customer-defined field, used in OpenAPI imports and exports without functional impact.",
		//	  "type": "string"
		//	}
		"auth_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Optional customer-defined field, used in OpenAPI imports and exports without functional impact.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerCredentials
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer.",
		//	  "type": "string"
		//	}
		"authorizer_credentials": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"authorizer_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerResultTtlInSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The TTL in seconds of cached authorizer results.",
		//	  "type": "integer"
		//	}
		"authorizer_result_ttl_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The TTL in seconds of cached authorizer results.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the authorizer's Uniform Resource Identifier (URI).",
		//	  "type": "string"
		//	}
		"authorizer_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the authorizer's Uniform Resource Identifier (URI).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdentitySource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identity source for which authorization is requested.",
		//	  "type": "string"
		//	}
		"identity_source": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identity source for which authorization is requested.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdentityValidationExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A validation expression for the incoming identity token.",
		//	  "type": "string"
		//	}
		"identity_validation_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A validation expression for the incoming identity token.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the authorizer.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the authorizer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProviderARNs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"provider_ar_ns": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the API.",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the API.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The authorizer type.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The authorizer type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::Authorizer",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Authorizer").WithTerraformTypeName("awscc_apigateway_authorizer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auth_type":                        "AuthType",
		"authorizer_credentials":           "AuthorizerCredentials",
		"authorizer_id":                    "AuthorizerId",
		"authorizer_result_ttl_in_seconds": "AuthorizerResultTtlInSeconds",
		"authorizer_uri":                   "AuthorizerUri",
		"identity_source":                  "IdentitySource",
		"identity_validation_expression":   "IdentityValidationExpression",
		"name":                             "Name",
		"provider_ar_ns":                   "ProviderARNs",
		"rest_api_id":                      "RestApiId",
		"type":                             "Type",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
