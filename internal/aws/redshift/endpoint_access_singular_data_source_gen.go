// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_redshift_endpoint_access", endpointAccessDataSource)
}

// endpointAccessDataSource returns the Terraform awscc_redshift_endpoint_access data source.
// This Terraform data source corresponds to the CloudFormation AWS::Redshift::EndpointAccess resource.
func endpointAccessDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Address
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The DNS address of the endpoint.",
		//	  "type": "string"
		//	}
		"address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The DNS address of the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClusterIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account",
		//	  "type": "string"
		//	}
		"cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EndpointCreateTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time (UTC) that the endpoint was created.",
		//	  "type": "string"
		//	}
		"endpoint_create_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time (UTC) that the endpoint was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EndpointName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the endpoint.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"endpoint_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EndpointStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the endpoint.",
		//	  "type": "string"
		//	}
		"endpoint_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Port
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The port number on which the cluster accepts incoming connections.",
		//	  "type": "integer"
		//	}
		"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The port number on which the cluster accepts incoming connections.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceOwner
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS account ID of the owner of the cluster.",
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"resource_owner": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS account ID of the owner of the cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subnet group name where Amazon Redshift chooses to deploy the endpoint.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"subnet_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The subnet group name where Amazon Redshift chooses to deploy the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.",
		//	  "properties": {
		//	    "NetworkInterfaces": {
		//	      "description": "One or more network interfaces of the endpoint. Also known as an interface endpoint.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Describes a network interface.",
		//	        "properties": {
		//	          "AvailabilityZone": {
		//	            "description": "The Availability Zone.",
		//	            "type": "string"
		//	          },
		//	          "NetworkInterfaceId": {
		//	            "description": "The network interface identifier.",
		//	            "type": "string"
		//	          },
		//	          "PrivateIpAddress": {
		//	            "description": "The IPv4 address of the network interface within the subnet.",
		//	            "type": "string"
		//	          },
		//	          "SubnetId": {
		//	            "description": "The subnet identifier.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "VpcEndpointId": {
		//	      "description": "The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.",
		//	      "type": "string"
		//	    },
		//	    "VpcId": {
		//	      "description": "The VPC identifier that the endpoint is associated.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"vpc_endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: NetworkInterfaces
				"network_interfaces": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AvailabilityZone
							"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The Availability Zone.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: NetworkInterfaceId
							"network_interface_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The network interface identifier.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: PrivateIpAddress
							"private_ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The IPv4 address of the network interface within the subnet.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: SubnetId
							"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The subnet identifier.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "One or more network interfaces of the endpoint. Also known as an interface endpoint.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VpcEndpointId
				"vpc_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VpcId
				"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The VPC identifier that the endpoint is associated.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcSecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of vpc security group ids to apply to the created endpoint access.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"vpc_security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of vpc security group ids to apply to the created endpoint access.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcSecurityGroups
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Describes the members of a VPC security group.",
		//	    "properties": {
		//	      "Status": {
		//	        "description": "The status of the VPC security group.",
		//	        "type": "string"
		//	      },
		//	      "VpcSecurityGroupId": {
		//	        "description": "The identifier of the VPC security group.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"vpc_security_groups": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Status
					"status": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The status of the VPC security group.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: VpcSecurityGroupId
					"vpc_security_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The identifier of the VPC security group.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Redshift::EndpointAccess",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::EndpointAccess").WithTerraformTypeName("awscc_redshift_endpoint_access")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":                "Address",
		"availability_zone":      "AvailabilityZone",
		"cluster_identifier":     "ClusterIdentifier",
		"endpoint_create_time":   "EndpointCreateTime",
		"endpoint_name":          "EndpointName",
		"endpoint_status":        "EndpointStatus",
		"network_interface_id":   "NetworkInterfaceId",
		"network_interfaces":     "NetworkInterfaces",
		"port":                   "Port",
		"private_ip_address":     "PrivateIpAddress",
		"resource_owner":         "ResourceOwner",
		"status":                 "Status",
		"subnet_group_name":      "SubnetGroupName",
		"subnet_id":              "SubnetId",
		"vpc_endpoint":           "VpcEndpoint",
		"vpc_endpoint_id":        "VpcEndpointId",
		"vpc_id":                 "VpcId",
		"vpc_security_group_id":  "VpcSecurityGroupId",
		"vpc_security_group_ids": "VpcSecurityGroupIds",
		"vpc_security_groups":    "VpcSecurityGroups",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
