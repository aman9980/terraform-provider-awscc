// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package appstream

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_appstream_directory_config", directoryConfigResource)
}

// directoryConfigResource returns the Terraform awscc_appstream_directory_config resource.
// This Terraform resource corresponds to the CloudFormation AWS::AppStream::DirectoryConfig resource.
func directoryConfigResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CertificateBasedAuthProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CertificateAuthorityArn": {
		//	      "type": "string"
		//	    },
		//	    "Status": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"certificate_based_auth_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CertificateAuthorityArn
				"certificate_authority_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DirectoryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"directory_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OrganizationalUnitDistinguishedNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"organizational_unit_distinguished_names": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceAccountCredentials
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AccountName": {
		//	      "type": "string"
		//	    },
		//	    "AccountPassword": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "AccountName",
		//	    "AccountPassword"
		//	  ],
		//	  "type": "object"
		//	}
		"service_account_credentials": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AccountName
				"account_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: AccountPassword
				"account_password": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					// AccountPassword is a write-only property.
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
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
		Description: "Resource Type definition for AWS::AppStream::DirectoryConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppStream::DirectoryConfig").WithTerraformTypeName("awscc_appstream_directory_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_name":                            "AccountName",
		"account_password":                        "AccountPassword",
		"certificate_authority_arn":               "CertificateAuthorityArn",
		"certificate_based_auth_properties":       "CertificateBasedAuthProperties",
		"directory_name":                          "DirectoryName",
		"organizational_unit_distinguished_names": "OrganizationalUnitDistinguishedNames",
		"service_account_credentials":             "ServiceAccountCredentials",
		"status":                                  "Status",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ServiceAccountCredentials/AccountPassword",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
