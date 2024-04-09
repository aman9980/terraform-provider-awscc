// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecr

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_ecr_replication_configuration", replicationConfigurationResource)
}

// replicationConfigurationResource returns the Terraform awscc_ecr_replication_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::ECR::ReplicationConfiguration resource.
func replicationConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: RegistryId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The RegistryId associated with the aws account.",
		//	  "type": "string"
		//	}
		"registry_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The RegistryId associated with the aws account.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReplicationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object representing the replication configuration for a registry.",
		//	  "properties": {
		//	    "Rules": {
		//	      "description": "An array of objects representing the replication rules for a replication configuration. A replication configuration may contain a maximum of 10 rules.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "An array of objects representing the details of a replication destination.",
		//	        "properties": {
		//	          "Destinations": {
		//	            "description": "An array of objects representing the details of a replication destination.",
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "description": "An array of objects representing the details of a replication destination.",
		//	              "properties": {
		//	                "Region": {
		//	                  "description": "A Region to replicate to.",
		//	                  "pattern": "[0-9a-z-]{2,25}",
		//	                  "type": "string"
		//	                },
		//	                "RegistryId": {
		//	                  "description": "The account ID of the destination registry to replicate to.",
		//	                  "pattern": "^[0-9]{12}$",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Region",
		//	                "RegistryId"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 25,
		//	            "minItems": 1,
		//	            "type": "array"
		//	          },
		//	          "RepositoryFilters": {
		//	            "description": "An array of objects representing the details of a repository filter.",
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "description": "An array of objects representing the details of a repository filter.",
		//	              "properties": {
		//	                "Filter": {
		//	                  "description": "The repository filter to be applied for replication.",
		//	                  "pattern": "^(?:[a-z0-9]+(?:[._-][a-z0-9]*)*/)*[a-z0-9]*(?:[._-][a-z0-9]*)*$",
		//	                  "type": "string"
		//	                },
		//	                "FilterType": {
		//	                  "description": "Type of repository filter",
		//	                  "enum": [
		//	                    "PREFIX_MATCH"
		//	                  ],
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Filter",
		//	                "FilterType"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 100,
		//	            "minItems": 0,
		//	            "type": "array"
		//	          }
		//	        },
		//	        "required": [
		//	          "Destinations"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 10,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "Rules"
		//	  ],
		//	  "type": "object"
		//	}
		"replication_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Rules
				"rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Destinations
							"destinations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Region
										"region": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "A Region to replicate to.",
											Required:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.RegexMatches(regexp.MustCompile("[0-9a-z-]{2,25}"), ""),
											}, /*END VALIDATORS*/
										}, /*END ATTRIBUTE*/
										// Property: RegistryId
										"registry_id": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The account ID of the destination registry to replicate to.",
											Required:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]{12}$"), ""),
											}, /*END VALIDATORS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "An array of objects representing the details of a replication destination.",
								Required:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.SizeBetween(1, 25),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: RepositoryFilters
							"repository_filters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Filter
										"filter": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The repository filter to be applied for replication.",
											Required:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.RegexMatches(regexp.MustCompile("^(?:[a-z0-9]+(?:[._-][a-z0-9]*)*/)*[a-z0-9]*(?:[._-][a-z0-9]*)*$"), ""),
											}, /*END VALIDATORS*/
										}, /*END ATTRIBUTE*/
										// Property: FilterType
										"filter_type": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "Type of repository filter",
											Required:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.OneOf(
													"PREFIX_MATCH",
												),
											}, /*END VALIDATORS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "An array of objects representing the details of a repository filter.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.SizeBetween(0, 100),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "An array of objects representing the replication rules for a replication configuration. A replication configuration may contain a maximum of 10 rules.",
					Required:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(0, 10),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object representing the replication configuration for a registry.",
			Required:    true,
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
		Description: "The AWS::ECR::ReplicationConfiguration resource configures the replication destinations for an Amazon Elastic Container Registry (Amazon Private ECR). For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/replication.html",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::ReplicationConfiguration").WithTerraformTypeName("awscc_ecr_replication_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"destinations":              "Destinations",
		"filter":                    "Filter",
		"filter_type":               "FilterType",
		"region":                    "Region",
		"registry_id":               "RegistryId",
		"replication_configuration": "ReplicationConfiguration",
		"repository_filters":        "RepositoryFilters",
		"rules":                     "Rules",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
