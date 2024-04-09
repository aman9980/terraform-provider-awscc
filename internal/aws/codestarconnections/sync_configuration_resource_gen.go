// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package codestarconnections

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_codestarconnections_sync_configuration", syncConfigurationResource)
}

// syncConfigurationResource returns the Terraform awscc_codestarconnections_sync_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::CodeStarConnections::SyncConfiguration resource.
func syncConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Branch
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the branch of the repository from which resources are to be synchronized,",
		//	  "type": "string"
		//	}
		"branch": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the branch of the repository from which resources are to be synchronized,",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigFile
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The source provider repository path of the sync configuration file of the respective SyncType.",
		//	  "type": "string"
		//	}
		"config_file": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The source provider repository path of the sync configuration file of the respective SyncType.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: OwnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "the ID of the entity that owns the repository.",
		//	  "pattern": "[a-za-z0-9_\\.-]+",
		//	  "type": "string"
		//	}
		"owner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "the ID of the entity that owns the repository.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProviderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the external provider where your third-party code repository is configured.",
		//	  "pattern": "^(GitHub|Bitbucket|GitHubEnterprise|GitLab)$",
		//	  "type": "string"
		//	}
		"provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the external provider where your third-party code repository is configured.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RepositoryLinkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.",
		//	  "pattern": "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}",
		//	  "type": "string"
		//	}
		"repository_link_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: RepositoryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the repository that is being synced to.",
		//	  "pattern": "[a-za-z0-9_\\.-]+",
		//	  "type": "string"
		//	}
		"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the repository that is being synced to.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the resource that is being synchronized to the repository.",
		//	  "pattern": "[a-za-z0-9_\\.-]+",
		//	  "type": "string"
		//	}
		"resource_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the resource that is being synchronized to the repository.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[a-za-z0-9_\\.-]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: SyncType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.",
		//	  "type": "string"
		//	}
		"sync_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Schema for AWS::CodeStarConnections::SyncConfiguration resource which is used to enables an AWS resource to be synchronized from a source-provider.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeStarConnections::SyncConfiguration").WithTerraformTypeName("awscc_codestarconnections_sync_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"branch":             "Branch",
		"config_file":        "ConfigFile",
		"owner_id":           "OwnerId",
		"provider_type":      "ProviderType",
		"repository_link_id": "RepositoryLinkId",
		"repository_name":    "RepositoryName",
		"resource_name":      "ResourceName",
		"role_arn":           "RoleArn",
		"sync_type":          "SyncType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
