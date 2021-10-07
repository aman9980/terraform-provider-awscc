// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_apigateway_authorizer", authorizerDataSourceType)
}

// authorizerDataSourceType returns the Terraform awscc_apigateway_authorizer data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ApiGateway::Authorizer resource type.
func authorizerDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auth_type": {
			// Property: AuthType
			// CloudFormation resource type schema:
			// {
			//   "description": "Optional customer-defined field, used in OpenAPI imports and exports without functional impact.",
			//   "type": "string"
			// }
			Description: "Optional customer-defined field, used in OpenAPI imports and exports without functional impact.",
			Type:        types.StringType,
			Computed:    true,
		},
		"authorizer_credentials": {
			// Property: AuthorizerCredentials
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer.",
			//   "type": "string"
			// }
			Description: "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer.",
			Type:        types.StringType,
			Computed:    true,
		},
		"authorizer_id": {
			// Property: AuthorizerId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"authorizer_result_ttl_in_seconds": {
			// Property: AuthorizerResultTtlInSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "The TTL in seconds of cached authorizer results.",
			//   "type": "integer"
			// }
			Description: "The TTL in seconds of cached authorizer results.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"authorizer_uri": {
			// Property: AuthorizerUri
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the authorizer's Uniform Resource Identifier (URI).",
			//   "type": "string"
			// }
			Description: "Specifies the authorizer's Uniform Resource Identifier (URI).",
			Type:        types.StringType,
			Computed:    true,
		},
		"identity_source": {
			// Property: IdentitySource
			// CloudFormation resource type schema:
			// {
			//   "description": "The identity source for which authorization is requested.",
			//   "type": "string"
			// }
			Description: "The identity source for which authorization is requested.",
			Type:        types.StringType,
			Computed:    true,
		},
		"identity_validation_expression": {
			// Property: IdentityValidationExpression
			// CloudFormation resource type schema:
			// {
			//   "description": "A validation expression for the incoming identity token.",
			//   "type": "string"
			// }
			Description: "A validation expression for the incoming identity token.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the authorizer.",
			//   "type": "string"
			// }
			Description: "The name of the authorizer.",
			Type:        types.StringType,
			Computed:    true,
		},
		"provider_ar_ns": {
			// Property: ProviderARNs
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"rest_api_id": {
			// Property: RestApiId
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the API.",
			//   "type": "string"
			// }
			Description: "The identifier of the API.",
			Type:        types.StringType,
			Computed:    true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The authorizer type.",
			//   "type": "string"
			// }
			Description: "The authorizer type.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ApiGateway::Authorizer",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

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

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
