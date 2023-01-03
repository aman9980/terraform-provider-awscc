// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotsitewise_access_policy", accessPolicyDataSource)
}

// accessPolicyDataSource returns the Terraform awscc_iotsitewise_access_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTSiteWise::AccessPolicy resource.
func accessPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessPolicyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the access policy.",
		//	  "type": "string"
		//	}
		"access_policy_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the access policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AccessPolicyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the access policy.",
		//	  "type": "string"
		//	}
		"access_policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the access policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AccessPolicyIdentity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The identity for this access policy. Choose either a user or a group but not both.",
		//	  "properties": {
		//	    "IamRole": {
		//	      "additionalProperties": false,
		//	      "description": "Contains information for an IAM role identity in an access policy.",
		//	      "properties": {
		//	        "arn": {
		//	          "description": "The ARN of the IAM role.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "IamUser": {
		//	      "additionalProperties": false,
		//	      "description": "Contains information for an IAM user identity in an access policy.",
		//	      "properties": {
		//	        "arn": {
		//	          "description": "The ARN of the IAM user.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "User": {
		//	      "additionalProperties": false,
		//	      "description": "Contains information for a user identity in an access policy.",
		//	      "properties": {
		//	        "id": {
		//	          "description": "The AWS SSO ID of the user.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"access_policy_identity": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IamRole
				"iam_role": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: arn
						"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the IAM role.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains information for an IAM role identity in an access policy.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IamUser
				"iam_user": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: arn
						"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the IAM user.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains information for an IAM user identity in an access policy.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: User
				"user": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: id
						"id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The AWS SSO ID of the user.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains information for a user identity in an access policy.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The identity for this access policy. Choose either a user or a group but not both.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AccessPolicyPermission
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The permission level for this access policy. Valid values are ADMINISTRATOR or VIEWER.",
		//	  "type": "string"
		//	}
		"access_policy_permission": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The permission level for this access policy. Valid values are ADMINISTRATOR or VIEWER.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AccessPolicyResource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The AWS IoT SiteWise Monitor resource for this access policy. Choose either portal or project but not both.",
		//	  "properties": {
		//	    "Portal": {
		//	      "additionalProperties": false,
		//	      "description": "A portal resource.",
		//	      "properties": {
		//	        "id": {
		//	          "description": "The ID of the portal.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Project": {
		//	      "additionalProperties": false,
		//	      "description": "A project resource.",
		//	      "properties": {
		//	        "id": {
		//	          "description": "The ID of the project.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"access_policy_resource": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Portal
				"portal": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: id
						"id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ID of the portal.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "A portal resource.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Project
				"project": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: id
						"id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ID of the project.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "A project resource.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The AWS IoT SiteWise Monitor resource for this access policy. Choose either portal or project but not both.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTSiteWise::AccessPolicy",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::AccessPolicy").WithTerraformTypeName("awscc_iotsitewise_access_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_policy_arn":        "AccessPolicyArn",
		"access_policy_id":         "AccessPolicyId",
		"access_policy_identity":   "AccessPolicyIdentity",
		"access_policy_permission": "AccessPolicyPermission",
		"access_policy_resource":   "AccessPolicyResource",
		"arn":                      "arn",
		"iam_role":                 "IamRole",
		"iam_user":                 "IamUser",
		"id":                       "id",
		"portal":                   "Portal",
		"project":                  "Project",
		"user":                     "User",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
