// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_amplify_app", appDataSource)
}

// appDataSource returns the Terraform awscc_amplify_app data source.
// This Terraform data source corresponds to the CloudFormation AWS::Amplify::App resource.
func appDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessToken
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"access_token": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AppId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 20,
		//	  "minLength": 1,
		//	  "pattern": "d[a-z0-9]+",
		//	  "type": "string"
		//	}
		"app_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AppName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "(?s).+",
		//	  "type": "string"
		//	}
		"app_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AutoBranchCreationConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AutoBranchCreationPatterns": {
		//	      "items": {
		//	        "maxLength": 2048,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "BasicAuthConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "EnableBasicAuth": {
		//	          "type": "boolean"
		//	        },
		//	        "Password": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Username": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "BuildSpec": {
		//	      "maxLength": 25000,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "EnableAutoBranchCreation": {
		//	      "type": "boolean"
		//	    },
		//	    "EnableAutoBuild": {
		//	      "type": "boolean"
		//	    },
		//	    "EnablePerformanceMode": {
		//	      "type": "boolean"
		//	    },
		//	    "EnablePullRequestPreview": {
		//	      "type": "boolean"
		//	    },
		//	    "EnvironmentVariables": {
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Name": {
		//	            "maxLength": 255,
		//	            "pattern": "(?s).*",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "maxLength": 5500,
		//	            "pattern": "(?s).*",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "Framework": {
		//	      "maxLength": 255,
		//	      "pattern": "(?s).*",
		//	      "type": "string"
		//	    },
		//	    "PullRequestEnvironmentName": {
		//	      "maxLength": 20,
		//	      "pattern": "(?s).*",
		//	      "type": "string"
		//	    },
		//	    "Stage": {
		//	      "enum": [
		//	        "EXPERIMENTAL",
		//	        "BETA",
		//	        "PULL_REQUEST",
		//	        "PRODUCTION",
		//	        "DEVELOPMENT"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"auto_branch_creation_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AutoBranchCreationPatterns
				"auto_branch_creation_patterns": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: BasicAuthConfig
				"basic_auth_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EnableBasicAuth
						"enable_basic_auth": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Password
						"password": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Username
						"username": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: BuildSpec
				"build_spec": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EnableAutoBranchCreation
				"enable_auto_branch_creation": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EnableAutoBuild
				"enable_auto_build": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EnablePerformanceMode
				"enable_performance_mode": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EnablePullRequestPreview
				"enable_pull_request_preview": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EnvironmentVariables
				"environment_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Framework
				"framework": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PullRequestEnvironmentName
				"pull_request_environment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Stage
				"stage": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BasicAuthConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "EnableBasicAuth": {
		//	      "type": "boolean"
		//	    },
		//	    "Password": {
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "Username": {
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"basic_auth_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnableBasicAuth
				"enable_basic_auth": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Password
				"password": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Username
				"username": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BuildSpec
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 25000,
		//	  "minLength": 1,
		//	  "pattern": "(?s).+",
		//	  "type": "string"
		//	}
		"build_spec": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CustomHeaders
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 25000,
		//	  "minLength": 0,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"custom_headers": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CustomRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Condition": {
		//	        "maxLength": 2048,
		//	        "minLength": 0,
		//	        "pattern": "(?s).*",
		//	        "type": "string"
		//	      },
		//	      "Source": {
		//	        "maxLength": 2048,
		//	        "minLength": 1,
		//	        "pattern": "(?s).+",
		//	        "type": "string"
		//	      },
		//	      "Status": {
		//	        "maxLength": 7,
		//	        "minLength": 3,
		//	        "pattern": ".{3,7}",
		//	        "type": "string"
		//	      },
		//	      "Target": {
		//	        "maxLength": 2048,
		//	        "minLength": 1,
		//	        "pattern": "(?s).+",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Target",
		//	      "Source"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"custom_rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Condition
					"condition": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Source
					"source": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Status
					"status": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Target
					"target": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultDomain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"default_domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnableBranchAutoDeletion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_branch_auto_deletion": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentVariables
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Name": {
		//	        "maxLength": 255,
		//	        "pattern": "(?s).*",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 5500,
		//	        "pattern": "(?s).*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"environment_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IAMServiceRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "minLength": 1,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"iam_service_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "(?s).+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: OauthToken
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"oauth_token": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Platform
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "WEB",
		//	    "WEB_DYNAMIC",
		//	    "WEB_COMPUTE"
		//	  ],
		//	  "type": "string"
		//	}
		"platform": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Repository
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"repository": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "insertionOrder": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Amplify::App",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::App").WithTerraformTypeName("awscc_amplify_app")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_token":                  "AccessToken",
		"app_id":                        "AppId",
		"app_name":                      "AppName",
		"arn":                           "Arn",
		"auto_branch_creation_config":   "AutoBranchCreationConfig",
		"auto_branch_creation_patterns": "AutoBranchCreationPatterns",
		"basic_auth_config":             "BasicAuthConfig",
		"build_spec":                    "BuildSpec",
		"condition":                     "Condition",
		"custom_headers":                "CustomHeaders",
		"custom_rules":                  "CustomRules",
		"default_domain":                "DefaultDomain",
		"description":                   "Description",
		"enable_auto_branch_creation":   "EnableAutoBranchCreation",
		"enable_auto_build":             "EnableAutoBuild",
		"enable_basic_auth":             "EnableBasicAuth",
		"enable_branch_auto_deletion":   "EnableBranchAutoDeletion",
		"enable_performance_mode":       "EnablePerformanceMode",
		"enable_pull_request_preview":   "EnablePullRequestPreview",
		"environment_variables":         "EnvironmentVariables",
		"framework":                     "Framework",
		"iam_service_role":              "IAMServiceRole",
		"key":                           "Key",
		"name":                          "Name",
		"oauth_token":                   "OauthToken",
		"password":                      "Password",
		"platform":                      "Platform",
		"pull_request_environment_name": "PullRequestEnvironmentName",
		"repository":                    "Repository",
		"source":                        "Source",
		"stage":                         "Stage",
		"status":                        "Status",
		"tags":                          "Tags",
		"target":                        "Target",
		"username":                      "Username",
		"value":                         "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
