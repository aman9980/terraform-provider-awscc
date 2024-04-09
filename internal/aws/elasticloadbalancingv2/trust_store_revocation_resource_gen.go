// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_elasticloadbalancingv2_trust_store_revocation", trustStoreRevocationResource)
}

// trustStoreRevocationResource returns the Terraform awscc_elasticloadbalancingv2_trust_store_revocation resource.
// This Terraform resource corresponds to the CloudFormation AWS::ElasticLoadBalancingV2::TrustStoreRevocation resource.
func trustStoreRevocationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: RevocationContents
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The attributes required to create a trust store revocation.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "RevocationType": {
		//	        "type": "string"
		//	      },
		//	      "S3Bucket": {
		//	        "type": "string"
		//	      },
		//	      "S3Key": {
		//	        "type": "string"
		//	      },
		//	      "S3ObjectVersion": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"revocation_contents": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: RevocationType
					"revocation_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: S3Bucket
					"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: S3Key
					"s3_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: S3ObjectVersion
					"s3_object_version": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The attributes required to create a trust store revocation.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
				setplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// RevocationContents is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: RevocationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID associated with the revocation.",
		//	  "format": "int64",
		//	  "type": "integer"
		//	}
		"revocation_id": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The ID associated with the revocation.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TrustStoreArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the trust store.",
		//	  "type": "string"
		//	}
		"trust_store_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the trust store.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TrustStoreRevocations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The data associated with a trust store revocation",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "NumberOfRevokedEntries": {
		//	        "format": "int64",
		//	        "type": "integer"
		//	      },
		//	      "RevocationId": {
		//	        "type": "string"
		//	      },
		//	      "RevocationType": {
		//	        "type": "string"
		//	      },
		//	      "TrustStoreArn": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"trust_store_revocations": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: NumberOfRevokedEntries
					"number_of_revoked_entries": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: RevocationId
					"revocation_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: RevocationType
					"revocation_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: TrustStoreArn
					"trust_store_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The data associated with a trust store revocation",
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::ElasticLoadBalancingV2::TrustStoreRevocation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticLoadBalancingV2::TrustStoreRevocation").WithTerraformTypeName("awscc_elasticloadbalancingv2_trust_store_revocation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"number_of_revoked_entries": "NumberOfRevokedEntries",
		"revocation_contents":       "RevocationContents",
		"revocation_id":             "RevocationId",
		"revocation_type":           "RevocationType",
		"s3_bucket":                 "S3Bucket",
		"s3_key":                    "S3Key",
		"s3_object_version":         "S3ObjectVersion",
		"trust_store_arn":           "TrustStoreArn",
		"trust_store_revocations":   "TrustStoreRevocations",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/RevocationContents",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
