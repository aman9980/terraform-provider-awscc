// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_gateway_route_table_association", gatewayRouteTableAssociationResource)
}

// gatewayRouteTableAssociationResource returns the Terraform awscc_ec2_gateway_route_table_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::GatewayRouteTableAssociation resource.
func gatewayRouteTableAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The route table association ID.",
		//	  "type": "string"
		//	}
		"association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The route table association ID.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the gateway.",
		//	  "type": "string"
		//	}
		"gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the gateway.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RouteTableId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the route table.",
		//	  "type": "string"
		//	}
		"route_table_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the route table.",
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
		Description: "Associates a gateway with a route table. The gateway and route table must be in the same VPC. This association causes the incoming traffic to the gateway to be routed according to the routes in the route table.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::GatewayRouteTableAssociation").WithTerraformTypeName("awscc_ec2_gateway_route_table_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id": "AssociationId",
		"gateway_id":     "GatewayId",
		"route_table_id": "RouteTableId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
