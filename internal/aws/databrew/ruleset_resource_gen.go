// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package databrew

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_databrew_ruleset", rulesetResource)
}

// rulesetResource returns the Terraform awscc_databrew_ruleset resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataBrew::Ruleset resource.
func rulesetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the Ruleset",
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the Ruleset",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the Ruleset",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the Ruleset",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Rules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of the data quality rules in the ruleset",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Data quality rule for a target resource (dataset)",
		//	    "properties": {
		//	      "CheckExpression": {
		//	        "description": "Expression with rule conditions",
		//	        "maxLength": 1024,
		//	        "minLength": 4,
		//	        "pattern": "^[\u003e\u003c0-9A-Za-z_.,:)(!= ]+$",
		//	        "type": "string"
		//	      },
		//	      "ColumnSelectors": {
		//	        "insertionOrder": true,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "Selector of a column from a dataset for profile job configuration. One selector includes either a column name or a regular expression",
		//	          "properties": {
		//	            "Name": {
		//	              "description": "The name of a column from a dataset",
		//	              "maxLength": 255,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Regex": {
		//	              "description": "A regular expression for selecting a column from a dataset",
		//	              "maxLength": 255,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "minItems": 1,
		//	        "type": "array"
		//	      },
		//	      "Disabled": {
		//	        "description": "Boolean value to disable/enable a rule",
		//	        "type": "boolean"
		//	      },
		//	      "Name": {
		//	        "description": "Name of the rule",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "SubstitutionMap": {
		//	        "insertionOrder": true,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "A key-value pair to associate expression's substitution variable names with their values",
		//	          "properties": {
		//	            "Value": {
		//	              "description": "Value or column name",
		//	              "maxLength": 1024,
		//	              "minLength": 0,
		//	              "type": "string"
		//	            },
		//	            "ValueReference": {
		//	              "description": "Variable name",
		//	              "maxLength": 128,
		//	              "minLength": 2,
		//	              "pattern": "^:[A-Za-z0-9_]+$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "ValueReference",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "Threshold": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Type": {
		//	            "description": "Threshold type for a rule",
		//	            "enum": [
		//	              "GREATER_THAN_OR_EQUAL",
		//	              "LESS_THAN_OR_EQUAL",
		//	              "GREATER_THAN",
		//	              "LESS_THAN"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Unit": {
		//	            "description": "Threshold unit for a rule",
		//	            "enum": [
		//	              "COUNT",
		//	              "PERCENTAGE"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "description": "Threshold value for a rule",
		//	            "type": "number"
		//	          }
		//	        },
		//	        "required": [
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "CheckExpression"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CheckExpression
					"check_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Expression with rule conditions",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(4, 1024),
							stringvalidator.RegexMatches(regexp.MustCompile("^[><0-9A-Za-z_.,:)(!= ]+$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: ColumnSelectors
					"column_selectors": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Name
								"name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of a column from a dataset",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 255),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Regex
								"regex": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "A regular expression for selecting a column from a dataset",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 255),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Optional: true,
						Computed: true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.SizeAtLeast(1),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Disabled
					"disabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "Boolean value to disable/enable a rule",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
							boolplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Name of the rule",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: SubstitutionMap
					"substitution_map": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Value
								"value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Value or column name",
									Required:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(0, 1024),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: ValueReference
								"value_reference": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Variable name",
									Required:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(2, 128),
										stringvalidator.RegexMatches(regexp.MustCompile("^:[A-Za-z0-9_]+$"), ""),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Threshold
					"threshold": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Threshold type for a rule",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"GREATER_THAN_OR_EQUAL",
										"LESS_THAN_OR_EQUAL",
										"GREATER_THAN",
										"LESS_THAN",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Unit
							"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Threshold unit for a rule",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"COUNT",
										"PERCENTAGE",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "Threshold value for a rule",
								Required:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of the data quality rules in the ruleset",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource",
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
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
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
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn of the target resource (dataset) to apply the ruleset to",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn of the target resource (dataset) to apply the ruleset to",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
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
		Description: "Resource schema for AWS::DataBrew::Ruleset.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataBrew::Ruleset").WithTerraformTypeName("awscc_databrew_ruleset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"check_expression": "CheckExpression",
		"column_selectors": "ColumnSelectors",
		"description":      "Description",
		"disabled":         "Disabled",
		"key":              "Key",
		"name":             "Name",
		"regex":            "Regex",
		"rules":            "Rules",
		"substitution_map": "SubstitutionMap",
		"tags":             "Tags",
		"target_arn":       "TargetArn",
		"threshold":        "Threshold",
		"type":             "Type",
		"unit":             "Unit",
		"value":            "Value",
		"value_reference":  "ValueReference",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
