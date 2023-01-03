// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_task_template", taskTemplateDataSource)
}

// taskTemplateDataSource returns the Terraform awscc_connect_task_template data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::TaskTemplate resource.
func taskTemplateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier (arn) of the task template.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/task-template/[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier (arn) of the task template.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClientToken
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "the client token string in uuid format",
		//	  "pattern": "^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$",
		//	  "type": "string"
		//	}
		"client_token": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "the client token string in uuid format",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Constraints
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The constraints for the task template",
		//	  "properties": {
		//	    "InvisibleFields": {
		//	      "description": "The list of the task template's invisible fields",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Invisible field info",
		//	        "properties": {
		//	          "Id": {
		//	            "additionalProperties": false,
		//	            "description": "the identifier (name) for the task template field",
		//	            "properties": {
		//	              "Name": {
		//	                "description": "The name of the task template field",
		//	                "maxLength": 100,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Name"
		//	            ],
		//	            "type": "object"
		//	          }
		//	        },
		//	        "required": [
		//	          "Id"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 50,
		//	      "type": "array"
		//	    },
		//	    "ReadOnlyFields": {
		//	      "description": "The list of the task template's read only fields",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "ReadOnly field info",
		//	        "properties": {
		//	          "Id": {
		//	            "additionalProperties": false,
		//	            "description": "the identifier (name) for the task template field",
		//	            "properties": {
		//	              "Name": {
		//	                "description": "The name of the task template field",
		//	                "maxLength": 100,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Name"
		//	            ],
		//	            "type": "object"
		//	          }
		//	        },
		//	        "required": [
		//	          "Id"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 50,
		//	      "type": "array"
		//	    },
		//	    "RequiredFields": {
		//	      "description": "The list of the task template's required fields",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Required field info",
		//	        "properties": {
		//	          "Id": {
		//	            "additionalProperties": false,
		//	            "description": "the identifier (name) for the task template field",
		//	            "properties": {
		//	              "Name": {
		//	                "description": "The name of the task template field",
		//	                "maxLength": 100,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Name"
		//	            ],
		//	            "type": "object"
		//	          }
		//	        },
		//	        "required": [
		//	          "Id"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 50,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"constraints": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: InvisibleFields
				"invisible_fields": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Id
							"id": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Name
									"name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The name of the task template field",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "the identifier (name) for the task template field",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The list of the task template's invisible fields",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ReadOnlyFields
				"read_only_fields": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Id
							"id": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Name
									"name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The name of the task template field",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "the identifier (name) for the task template field",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The list of the task template's read only fields",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RequiredFields
				"required_fields": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Id
							"id": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Name
									"name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The name of the task template field",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "the identifier (name) for the task template field",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The list of the task template's required fields",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The constraints for the task template",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ContactFlowArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the contact flow.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"contact_flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the contact flow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Defaults
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "the default value for the task template's field",
		//	    "properties": {
		//	      "DefaultValue": {
		//	        "description": "the default value for the task template's field",
		//	        "maxLength": 4096,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Id": {
		//	        "additionalProperties": false,
		//	        "description": "the identifier (name) for the task template field",
		//	        "properties": {
		//	          "Name": {
		//	            "description": "The name of the task template field",
		//	            "maxLength": 100,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "Id",
		//	      "DefaultValue"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array"
		//	}
		"defaults": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DefaultValue
					"default_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "the default value for the task template's field",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the task template field",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "the identifier (name) for the task template field",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the task template.",
		//	  "maxLength": 255,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the task template.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Fields
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of task template's fields",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A task template field object.",
		//	    "properties": {
		//	      "Description": {
		//	        "description": "The description of the task template's field",
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      },
		//	      "Id": {
		//	        "additionalProperties": false,
		//	        "description": "the identifier (name) for the task template field",
		//	        "properties": {
		//	          "Name": {
		//	            "description": "The name of the task template field",
		//	            "maxLength": 100,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "SingleSelectOptions": {
		//	        "description": "list of field options to be used with single select",
		//	        "items": {
		//	          "description": "Single select field identifier",
		//	          "maxLength": 100,
		//	          "minLength": 1,
		//	          "pattern": "^[A-Za-z0-9](?:[A-Za-z0-9_.,\\s-]*[A-Za-z0-9_.,-])?$",
		//	          "type": "string"
		//	        },
		//	        "maxItems": 50,
		//	        "type": "array"
		//	      },
		//	      "Type": {
		//	        "description": "The type of the task template's field",
		//	        "enum": [
		//	          "NAME",
		//	          "DESCRIPTION",
		//	          "SCHEDULED_TIME",
		//	          "QUICK_CONNECT",
		//	          "URL",
		//	          "NUMBER",
		//	          "TEXT",
		//	          "TEXT_AREA",
		//	          "DATE_TIME",
		//	          "BOOLEAN",
		//	          "SINGLE_SELECT",
		//	          "EMAIL"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Id",
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array"
		//	}
		"fields": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The description of the task template's field",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the task template field",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "the identifier (name) for the task template field",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: SingleSelectOptions
					"single_select_options": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "list of field options to be used with single select",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of the task template's field",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of task template's fields",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier (arn) of the instance.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier (arn) of the instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the task template.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the task template.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the task template",
		//	  "enum": [
		//	    "ACTIVE",
		//	    "INACTIVE"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the task template",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Connect::TaskTemplate",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::TaskTemplate").WithTerraformTypeName("awscc_connect_task_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"client_token":          "ClientToken",
		"constraints":           "Constraints",
		"contact_flow_arn":      "ContactFlowArn",
		"default_value":         "DefaultValue",
		"defaults":              "Defaults",
		"description":           "Description",
		"fields":                "Fields",
		"id":                    "Id",
		"instance_arn":          "InstanceArn",
		"invisible_fields":      "InvisibleFields",
		"key":                   "Key",
		"name":                  "Name",
		"read_only_fields":      "ReadOnlyFields",
		"required_fields":       "RequiredFields",
		"single_select_options": "SingleSelectOptions",
		"status":                "Status",
		"tags":                  "Tags",
		"type":                  "Type",
		"value":                 "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
