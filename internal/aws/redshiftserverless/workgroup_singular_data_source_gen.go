// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package redshiftserverless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_redshiftserverless_workgroup", workgroupDataSource)
}

// workgroupDataSource returns the Terraform awscc_redshiftserverless_workgroup data source.
// This Terraform data source corresponds to the CloudFormation AWS::RedshiftServerless::Workgroup resource.
func workgroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BaseCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"base_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ParameterKey": {
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      },
		//	      "ParameterValue": {
		//	        "maxLength": 15000,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"config_parameters": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ParameterKey
					"parameter_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ParameterValue
					"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnhancedVpcRouting
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enhanced_vpc_routing": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NamespaceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 3,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"namespace_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PubliclyAccessible
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"publicly_accessible": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 255,
		//	    "minLength": 0,
		//	    "pattern": "^sg-[0-9a-fA-F]{8,}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 32,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 255,
		//	    "minLength": 0,
		//	    "pattern": "^subnet-[0-9a-fA-F]{8,}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 32,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Workgroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "BaseCapacity": {
		//	      "type": "integer"
		//	    },
		//	    "ConfigParameters": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "ParameterKey": {
		//	            "maxLength": 255,
		//	            "minLength": 0,
		//	            "type": "string"
		//	          },
		//	          "ParameterValue": {
		//	            "maxLength": 15000,
		//	            "minLength": 0,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "CreationDate": {
		//	      "type": "string"
		//	    },
		//	    "Endpoint": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Address": {
		//	          "type": "string"
		//	        },
		//	        "Port": {
		//	          "type": "integer"
		//	        },
		//	        "VpcEndpoints": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "NetworkInterfaces": {
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "AvailabilityZone": {
		//	                      "type": "string"
		//	                    },
		//	                    "NetworkInterfaceId": {
		//	                      "type": "string"
		//	                    },
		//	                    "PrivateIpAddress": {
		//	                      "type": "string"
		//	                    },
		//	                    "SubnetId": {
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "VpcEndpointId": {
		//	                "type": "string"
		//	              },
		//	              "VpcId": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "EnhancedVpcRouting": {
		//	      "type": "boolean"
		//	    },
		//	    "NamespaceName": {
		//	      "maxLength": 64,
		//	      "minLength": 3,
		//	      "pattern": "^[a-z0-9-]+$",
		//	      "type": "string"
		//	    },
		//	    "PubliclyAccessible": {
		//	      "type": "boolean"
		//	    },
		//	    "SecurityGroupIds": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "pattern": "^sg-[0-9a-fA-F]{8,}$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "Status": {
		//	      "enum": [
		//	        "CREATING",
		//	        "AVAILABLE",
		//	        "MODIFYING",
		//	        "DELETING"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "SubnetIds": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "pattern": "^subnet-[0-9a-fA-F]{8,}$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "WorkgroupArn": {
		//	      "type": "string"
		//	    },
		//	    "WorkgroupId": {
		//	      "type": "string"
		//	    },
		//	    "WorkgroupName": {
		//	      "maxLength": 64,
		//	      "minLength": 3,
		//	      "pattern": "^[a-z0-9-]*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"workgroup": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BaseCapacity
				"base_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ConfigParameters
				"config_parameters": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ParameterKey
							"parameter_key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: ParameterValue
							"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: CreationDate
				"creation_date": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Endpoint
				"endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Address
						"address": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Port
						"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: VpcEndpoints
						"vpc_endpoints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: NetworkInterfaces
									"network_interfaces": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
										NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: AvailabilityZone
												"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: NetworkInterfaceId
												"network_interface_id": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: PrivateIpAddress
												"private_ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: SubnetId
												"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
										}, /*END NESTED OBJECT*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: VpcEndpointId
									"vpc_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: VpcId
									"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EnhancedVpcRouting
				"enhanced_vpc_routing": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: NamespaceName
				"namespace_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PubliclyAccessible
				"publicly_accessible": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SubnetIds
				"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: WorkgroupArn
				"workgroup_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: WorkgroupId
				"workgroup_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: WorkgroupName
				"workgroup_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: WorkgroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 3,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"workgroup_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::RedshiftServerless::Workgroup",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RedshiftServerless::Workgroup").WithTerraformTypeName("awscc_redshiftserverless_workgroup")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":              "Address",
		"availability_zone":    "AvailabilityZone",
		"base_capacity":        "BaseCapacity",
		"config_parameters":    "ConfigParameters",
		"creation_date":        "CreationDate",
		"endpoint":             "Endpoint",
		"enhanced_vpc_routing": "EnhancedVpcRouting",
		"key":                  "Key",
		"namespace_name":       "NamespaceName",
		"network_interface_id": "NetworkInterfaceId",
		"network_interfaces":   "NetworkInterfaces",
		"parameter_key":        "ParameterKey",
		"parameter_value":      "ParameterValue",
		"port":                 "Port",
		"private_ip_address":   "PrivateIpAddress",
		"publicly_accessible":  "PubliclyAccessible",
		"security_group_ids":   "SecurityGroupIds",
		"status":               "Status",
		"subnet_id":            "SubnetId",
		"subnet_ids":           "SubnetIds",
		"tags":                 "Tags",
		"value":                "Value",
		"vpc_endpoint_id":      "VpcEndpointId",
		"vpc_endpoints":        "VpcEndpoints",
		"vpc_id":               "VpcId",
		"workgroup":            "Workgroup",
		"workgroup_arn":        "WorkgroupArn",
		"workgroup_id":         "WorkgroupId",
		"workgroup_name":       "WorkgroupName",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
