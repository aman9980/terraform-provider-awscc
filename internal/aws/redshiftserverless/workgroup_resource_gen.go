// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_redshiftserverless_workgroup", workgroupResource)
}

// workgroupResource returns the Terraform awscc_redshiftserverless_workgroup resource.
// This Terraform resource corresponds to the CloudFormation AWS::RedshiftServerless::Workgroup resource.
func workgroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BaseCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The base compute capacity of the workgroup in Redshift Processing Units (RPUs).",
		//	  "type": "integer"
		//	}
		"base_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The base compute capacity of the workgroup in Redshift Processing Units (RPUs).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// BaseCapacity is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ConfigParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of parameters to set for finer control over a database. Available options are datestyle, enable_user_activity_logging, query_group, search_path, max_query_execution_time, and require_ssl.",
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
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 255),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ParameterValue
					"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 15000),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of parameters to set for finer control over a database. Available options are datestyle, enable_user_activity_logging, query_group, search_path, max_query_execution_time, and require_ssl.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// ConfigParameters is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: EnhancedVpcRouting
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.",
		//	  "type": "boolean"
		//	}
		"enhanced_vpc_routing": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				generic.BoolDefaultValue(false),
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The max compute capacity of the workgroup in Redshift Processing Units (RPUs).",
		//	  "type": "integer"
		//	}
		"max_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The max compute capacity of the workgroup in Redshift Processing Units (RPUs).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// MaxCapacity is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: NamespaceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The namespace the workgroup is associated with.",
		//	  "maxLength": 64,
		//	  "minLength": 3,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"namespace_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The namespace the workgroup is associated with.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Port
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The custom port to use when connecting to a workgroup. Valid port ranges are 5431-5455 and 8191-8215. The default is 5439.",
		//	  "type": "integer"
		//	}
		"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The custom port to use when connecting to a workgroup. Valid port ranges are 5431-5455 and 8191-8215. The default is 5439.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PubliclyAccessible
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "A value that specifies whether the workgroup can be accessible from a public network.",
		//	  "type": "boolean"
		//	}
		"publicly_accessible": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A value that specifies whether the workgroup can be accessible from a public network.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				generic.BoolDefaultValue(false),
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of security group IDs to associate with the workgroup.",
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
			Description: "A list of security group IDs to associate with the workgroup.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 32),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(0, 255),
					stringvalidator.RegexMatches(regexp.MustCompile("^sg-[0-9a-fA-F]{8,}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// SecurityGroupIds is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of subnet IDs the workgroup is associated with.",
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
			Description: "A list of subnet IDs the workgroup is associated with.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 32),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(0, 255),
					stringvalidator.RegexMatches(regexp.MustCompile("^subnet-[0-9a-fA-F]{8,}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// SubnetIds is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The map of the key-value pairs used to tag the workgroup.",
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
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The map of the key-value pairs used to tag the workgroup.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Tags is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Workgroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Definition for workgroup resource",
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
		//	    "MaxCapacity": {
		//	      "type": "integer"
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
										PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
											generic.Multiset(),
										}, /*END PLAN MODIFIERS*/
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
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EnhancedVpcRouting
				"enhanced_vpc_routing": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MaxCapacity
				"max_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
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
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SubnetIds
				"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
					}, /*END PLAN MODIFIERS*/
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
			Description: "Definition for workgroup resource",
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WorkgroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the workgroup.",
		//	  "maxLength": 64,
		//	  "minLength": 3,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"workgroup_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the workgroup.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 64),
			}, /*END VALIDATORS*/
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
		Description: "Definition of AWS::RedshiftServerless::Workgroup Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RedshiftServerless::Workgroup").WithTerraformTypeName("awscc_redshiftserverless_workgroup")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":              "Address",
		"availability_zone":    "AvailabilityZone",
		"base_capacity":        "BaseCapacity",
		"config_parameters":    "ConfigParameters",
		"creation_date":        "CreationDate",
		"endpoint":             "Endpoint",
		"enhanced_vpc_routing": "EnhancedVpcRouting",
		"key":                  "Key",
		"max_capacity":         "MaxCapacity",
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

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/BaseCapacity",
		"/properties/MaxCapacity",
		"/properties/ConfigParameters",
		"/properties/SecurityGroupIds",
		"/properties/SubnetIds",
		"/properties/Tags",
		"/properties/Tags/*/Key",
		"/properties/Tags/*/Value",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
