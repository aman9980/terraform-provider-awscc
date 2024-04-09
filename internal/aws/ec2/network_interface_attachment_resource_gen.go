// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_network_interface_attachment", networkInterfaceAttachmentResource)
}

// networkInterfaceAttachmentResource returns the Terraform awscc_ec2_network_interface_attachment resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::NetworkInterfaceAttachment resource.
func networkInterfaceAttachmentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the network interface attachment.",
		//	  "type": "string"
		//	}
		"attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the network interface attachment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeleteOnTermination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": true,
		//	  "description": "Whether to delete the network interface when the instance terminates. By default, this value is set to true.",
		//	  "type": "boolean"
		//	}
		"delete_on_termination": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether to delete the network interface when the instance terminates. By default, this value is set to true.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				generic.BoolDefaultValue(true),
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeviceIndex
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The network interface's position in the attachment order. For example, the first attached network interface has a DeviceIndex of 0.",
		//	  "type": "string"
		//	}
		"device_index": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The network interface's position in the attachment order. For example, the first attached network interface has a DeviceIndex of 0.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnaSrdSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "EnaSrdEnabled": {
		//	      "type": "boolean"
		//	    },
		//	    "EnaSrdUdpSpecification": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "EnaSrdUdpEnabled": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"ena_srd_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnaSrdEnabled
				"ena_srd_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EnaSrdUdpSpecification
				"ena_srd_udp_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EnaSrdUdpEnabled
						"ena_srd_udp_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the instance to which you will attach the ENI.",
		//	  "type": "string"
		//	}
		"instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the instance to which you will attach the ENI.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkInterfaceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the ENI that you want to attach.",
		//	  "type": "string"
		//	}
		"network_interface_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the ENI that you want to attach.",
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
		Description: "Resource Type definition for AWS::EC2::NetworkInterfaceAttachment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::NetworkInterfaceAttachment").WithTerraformTypeName("awscc_ec2_network_interface_attachment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attachment_id":             "AttachmentId",
		"delete_on_termination":     "DeleteOnTermination",
		"device_index":              "DeviceIndex",
		"ena_srd_enabled":           "EnaSrdEnabled",
		"ena_srd_specification":     "EnaSrdSpecification",
		"ena_srd_udp_enabled":       "EnaSrdUdpEnabled",
		"ena_srd_udp_specification": "EnaSrdUdpSpecification",
		"instance_id":               "InstanceId",
		"network_interface_id":      "NetworkInterfaceId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
