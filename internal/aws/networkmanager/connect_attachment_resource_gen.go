// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_networkmanager_connect_attachment", connectAttachmentResourceType)
}

// connectAttachmentResourceType returns the Terraform awscc_networkmanager_connect_attachment resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkManager::ConnectAttachment resource type.
func connectAttachmentResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"attachment_id": {
			// Property: AttachmentId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the attachment.",
			//   "type": "string"
			// }
			Description: "The ID of the attachment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"attachment_policy_rule_number": {
			// Property: AttachmentPolicyRuleNumber
			// CloudFormation resource type schema:
			// {
			//   "description": "The policy rule number associated with the attachment.",
			//   "type": "integer"
			// }
			Description: "The policy rule number associated with the attachment.",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"attachment_type": {
			// Property: AttachmentType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of attachment.",
			//   "type": "string"
			// }
			Description: "The type of attachment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"core_network_arn": {
			// Property: CoreNetworkArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of a core network for the VPC attachment.",
			//   "type": "string"
			// }
			Description: "The ARN of a core network for the VPC attachment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"core_network_id": {
			// Property: CoreNetworkId
			// CloudFormation resource type schema:
			// {
			//   "description": "ID of the CoreNetwork that the attachment will be attached to.",
			//   "type": "string"
			// }
			Description: "ID of the CoreNetwork that the attachment will be attached to.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "Creation time of the attachment.",
			//   "type": "string"
			// }
			Description: "Creation time of the attachment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"edge_location": {
			// Property: EdgeLocation
			// CloudFormation resource type schema:
			// {
			//   "description": "Edge location of the attachment.",
			//   "type": "string"
			// }
			Description: "Edge location of the attachment.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"options": {
			// Property: Options
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Protocol options for connect attachment",
			//   "properties": {
			//     "Protocol": {
			//       "description": "Tunnel protocol for connect attachment",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Protocol options for connect attachment",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"protocol": {
						// Property: Protocol
						Description: "Tunnel protocol for connect attachment",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"owner_account_id": {
			// Property: OwnerAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the attachment account owner.",
			//   "type": "string"
			// }
			Description: "The ID of the attachment account owner.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"proposed_segment_change": {
			// Property: ProposedSegmentChange
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The attachment to move from one segment to another.",
			//   "properties": {
			//     "AttachmentPolicyRuleNumber": {
			//       "description": "New policy rule number of the attachment",
			//       "type": "integer"
			//     },
			//     "SegmentName": {
			//       "description": "Proposed segment name",
			//       "type": "string"
			//     },
			//     "Tags": {
			//       "description": "Proposed tags for the Segment.",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "A key-value pair to associate with a resource.",
			//         "properties": {
			//           "Key": {
			//             "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//             "type": "string"
			//           },
			//           "Value": {
			//             "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Key",
			//           "Value"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The attachment to move from one segment to another.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"attachment_policy_rule_number": {
						// Property: AttachmentPolicyRuleNumber
						Description: "New policy rule number of the attachment",
						Type:        types.Int64Type,
						Optional:    true,
					},
					"segment_name": {
						// Property: SegmentName
						Description: "Proposed segment name",
						Type:        types.StringType,
						Optional:    true,
					},
					"tags": {
						// Property: Tags
						Description: "Proposed tags for the Segment.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
									Type:        types.StringType,
									Required:    true,
								},
								"value": {
									// Property: Value
									Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Optional: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
				},
			),
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"resource_arn": {
			// Property: ResourceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The attachment resource ARN.",
			//   "type": "string"
			// }
			Description: "The attachment resource ARN.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"segment_name": {
			// Property: SegmentName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the segment attachment.",
			//   "type": "string"
			// }
			Description: "The name of the segment attachment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "State of the attachment.",
			//   "type": "string"
			// }
			Description: "State of the attachment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags for the attachment.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Tags for the attachment.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"transport_attachment_id": {
			// Property: TransportAttachmentId
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of transport attachment",
			//   "type": "string"
			// }
			Description: "Id of transport attachment",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"updated_at": {
			// Property: UpdatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "Last update time of the attachment.",
			//   "type": "string"
			// }
			Description: "Last update time of the attachment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "AWS::NetworkManager::ConnectAttachment Resource Type Definition",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::ConnectAttachment").WithTerraformTypeName("awscc_networkmanager_connect_attachment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attachment_id":                 "AttachmentId",
		"attachment_policy_rule_number": "AttachmentPolicyRuleNumber",
		"attachment_type":               "AttachmentType",
		"core_network_arn":              "CoreNetworkArn",
		"core_network_id":               "CoreNetworkId",
		"created_at":                    "CreatedAt",
		"edge_location":                 "EdgeLocation",
		"key":                           "Key",
		"options":                       "Options",
		"owner_account_id":              "OwnerAccountId",
		"proposed_segment_change":       "ProposedSegmentChange",
		"protocol":                      "Protocol",
		"resource_arn":                  "ResourceArn",
		"segment_name":                  "SegmentName",
		"state":                         "State",
		"tags":                          "Tags",
		"transport_attachment_id":       "TransportAttachmentId",
		"updated_at":                    "UpdatedAt",
		"value":                         "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}