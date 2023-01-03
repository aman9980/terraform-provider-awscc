// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_imagebuilder_distribution_configuration", distributionConfigurationDataSource)
}

// distributionConfigurationDataSource returns the Terraform awscc_imagebuilder_distribution_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::ImageBuilder::DistributionConfiguration resource.
func distributionConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the distribution configuration.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the distribution configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the distribution configuration.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the distribution configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Distributions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The distributions of the distribution configuration.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The distributions of the distribution configuration.",
		//	    "properties": {
		//	      "AmiDistributionConfiguration": {
		//	        "additionalProperties": false,
		//	        "description": "The specific AMI settings (for example, launch permissions, AMI tags).",
		//	        "properties": {
		//	          "AmiTags": {
		//	            "additionalProperties": false,
		//	            "description": "The tags to apply to AMIs distributed to this Region.",
		//	            "patternProperties": {
		//	              "": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Description": {
		//	            "description": "The description of the AMI distribution configuration.",
		//	            "type": "string"
		//	          },
		//	          "KmsKeyId": {
		//	            "description": "The KMS key identifier used to encrypt the distributed image.",
		//	            "type": "string"
		//	          },
		//	          "LaunchPermissionConfiguration": {
		//	            "additionalProperties": false,
		//	            "description": "Launch permissions can be used to configure which AWS accounts can use the AMI to launch instances.",
		//	            "properties": {
		//	              "OrganizationArns": {
		//	                "description": "The ARN for an Amazon Web Services Organization that you want to share your AMI with.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "OrganizationalUnitArns": {
		//	                "description": "The ARN for an Organizations organizational unit (OU) that you want to share your AMI with.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "UserGroups": {
		//	                "description": "The name of the group.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "UserIds": {
		//	                "description": "The AWS account ID.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Name": {
		//	            "description": "The name of the AMI distribution configuration.",
		//	            "type": "string"
		//	          },
		//	          "TargetAccountIds": {
		//	            "description": "The ID of accounts to which you want to distribute an image.",
		//	            "insertionOrder": true,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ContainerDistributionConfiguration": {
		//	        "additionalProperties": false,
		//	        "description": "Container distribution settings for encryption, licensing, and sharing in a specific Region.",
		//	        "properties": {
		//	          "ContainerTags": {
		//	            "description": "Tags that are attached to the container distribution configuration.",
		//	            "insertionOrder": true,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "Description": {
		//	            "description": "The description of the container distribution configuration.",
		//	            "type": "string"
		//	          },
		//	          "TargetRepository": {
		//	            "additionalProperties": false,
		//	            "description": "The destination repository for the container distribution configuration.",
		//	            "properties": {
		//	              "RepositoryName": {
		//	                "description": "The repository name of target container repository.",
		//	                "type": "string"
		//	              },
		//	              "Service": {
		//	                "description": "The service of target container repository.",
		//	                "enum": [
		//	                  "ECR"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "FastLaunchConfigurations": {
		//	        "description": "The Windows faster-launching configurations to use for AMI distribution.",
		//	        "insertionOrder": true,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "The Windows faster-launching configuration to use for AMI distribution.",
		//	          "properties": {
		//	            "AccountId": {
		//	              "description": "The owner account ID for the fast-launch enabled Windows AMI.",
		//	              "type": "string"
		//	            },
		//	            "Enabled": {
		//	              "description": "A Boolean that represents the current state of faster launching for the Windows AMI. Set to true to start using Windows faster launching, or false to stop using it.",
		//	              "type": "boolean"
		//	            },
		//	            "LaunchTemplate": {
		//	              "additionalProperties": false,
		//	              "description": "The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.",
		//	              "properties": {
		//	                "LaunchTemplateId": {
		//	                  "description": "The ID of the launch template to use for faster launching for a Windows AMI.",
		//	                  "type": "string"
		//	                },
		//	                "LaunchTemplateName": {
		//	                  "description": "The name of the launch template to use for faster launching for a Windows AMI.",
		//	                  "type": "string"
		//	                },
		//	                "LaunchTemplateVersion": {
		//	                  "description": "The version of the launch template to use for faster launching for a Windows AMI.",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "MaxParallelLaunches": {
		//	              "description": "The maximum number of parallel instances that are launched for creating resources.",
		//	              "type": "integer"
		//	            },
		//	            "SnapshotConfiguration": {
		//	              "additionalProperties": false,
		//	              "description": "Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.",
		//	              "properties": {
		//	                "TargetResourceCount": {
		//	                  "description": "The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.",
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "LaunchTemplateConfigurations": {
		//	        "description": "A group of launchTemplateConfiguration settings that apply to image distribution.",
		//	        "insertionOrder": true,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "launchTemplateConfiguration settings that apply to image distribution.",
		//	          "properties": {
		//	            "AccountId": {
		//	              "description": "The account ID that this configuration applies to.",
		//	              "type": "string"
		//	            },
		//	            "LaunchTemplateId": {
		//	              "description": "Identifies the EC2 launch template to use.",
		//	              "type": "string"
		//	            },
		//	            "SetDefaultVersion": {
		//	              "description": "Set the specified EC2 launch template as the default launch template for the specified account.",
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "LicenseConfigurationArns": {
		//	        "description": "The License Manager Configuration to associate with the AMI in the specified Region.",
		//	        "insertionOrder": true,
		//	        "items": {
		//	          "description": "The Amazon Resource Name (ARN) of the License Manager configuration.",
		//	          "type": "string"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "Region": {
		//	        "description": "region",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Region"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"distributions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AmiDistributionConfiguration
					"ami_distribution_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AmiTags
							"ami_tags":          // Pattern: ""
							schema.MapAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "The tags to apply to AMIs distributed to this Region.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The description of the AMI distribution configuration.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: KmsKeyId
							"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The KMS key identifier used to encrypt the distributed image.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: LaunchPermissionConfiguration
							"launch_permission_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: OrganizationArns
									"organization_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "The ARN for an Amazon Web Services Organization that you want to share your AMI with.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: OrganizationalUnitArns
									"organizational_unit_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "The ARN for an Organizations organizational unit (OU) that you want to share your AMI with.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: UserGroups
									"user_groups": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "The name of the group.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: UserIds
									"user_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "The AWS account ID.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Launch permissions can be used to configure which AWS accounts can use the AMI to launch instances.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the AMI distribution configuration.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: TargetAccountIds
							"target_account_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "The ID of accounts to which you want to distribute an image.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The specific AMI settings (for example, launch permissions, AMI tags).",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ContainerDistributionConfiguration
					"container_distribution_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ContainerTags
							"container_tags": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "Tags that are attached to the container distribution configuration.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The description of the container distribution configuration.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: TargetRepository
							"target_repository": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: RepositoryName
									"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The repository name of target container repository.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Service
									"service": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The service of target container repository.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The destination repository for the container distribution configuration.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Container distribution settings for encryption, licensing, and sharing in a specific Region.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: FastLaunchConfigurations
					"fast_launch_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AccountId
								"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The owner account ID for the fast-launch enabled Windows AMI.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "A Boolean that represents the current state of faster launching for the Windows AMI. Set to true to start using Windows faster launching, or false to stop using it.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: LaunchTemplate
								"launch_template": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: LaunchTemplateId
										"launch_template_id": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The ID of the launch template to use for faster launching for a Windows AMI.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: LaunchTemplateName
										"launch_template_name": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The name of the launch template to use for faster launching for a Windows AMI.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: LaunchTemplateVersion
										"launch_template_version": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The version of the launch template to use for faster launching for a Windows AMI.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: MaxParallelLaunches
								"max_parallel_launches": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "The maximum number of parallel instances that are launched for creating resources.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: SnapshotConfiguration
								"snapshot_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: TargetResourceCount
										"target_resource_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Description: "The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "The Windows faster-launching configurations to use for AMI distribution.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: LaunchTemplateConfigurations
					"launch_template_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AccountId
								"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The account ID that this configuration applies to.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: LaunchTemplateId
								"launch_template_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Identifies the EC2 launch template to use.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: SetDefaultVersion
								"set_default_version": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Set the specified EC2 launch template as the default launch template for the specified account.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "A group of launchTemplateConfiguration settings that apply to image distribution.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: LicenseConfigurationArns
					"license_configuration_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "The License Manager Configuration to associate with the AMI in the specified Region.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Region
					"region": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "region",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The distributions of the distribution configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the distribution configuration.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the distribution configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The tags associated with the component.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The tags associated with the component.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ImageBuilder::DistributionConfiguration",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::DistributionConfiguration").WithTerraformTypeName("awscc_imagebuilder_distribution_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                           "AccountId",
		"ami_distribution_configuration":       "AmiDistributionConfiguration",
		"ami_tags":                             "AmiTags",
		"arn":                                  "Arn",
		"container_distribution_configuration": "ContainerDistributionConfiguration",
		"container_tags":                       "ContainerTags",
		"description":                          "Description",
		"distributions":                        "Distributions",
		"enabled":                              "Enabled",
		"fast_launch_configurations":           "FastLaunchConfigurations",
		"kms_key_id":                           "KmsKeyId",
		"launch_permission_configuration":      "LaunchPermissionConfiguration",
		"launch_template":                      "LaunchTemplate",
		"launch_template_configurations":       "LaunchTemplateConfigurations",
		"launch_template_id":                   "LaunchTemplateId",
		"launch_template_name":                 "LaunchTemplateName",
		"launch_template_version":              "LaunchTemplateVersion",
		"license_configuration_arns":           "LicenseConfigurationArns",
		"max_parallel_launches":                "MaxParallelLaunches",
		"name":                                 "Name",
		"organization_arns":                    "OrganizationArns",
		"organizational_unit_arns":             "OrganizationalUnitArns",
		"region":                               "Region",
		"repository_name":                      "RepositoryName",
		"service":                              "Service",
		"set_default_version":                  "SetDefaultVersion",
		"snapshot_configuration":               "SnapshotConfiguration",
		"tags":                                 "Tags",
		"target_account_ids":                   "TargetAccountIds",
		"target_repository":                    "TargetRepository",
		"target_resource_count":                "TargetResourceCount",
		"user_groups":                          "UserGroups",
		"user_ids":                             "UserIds",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
