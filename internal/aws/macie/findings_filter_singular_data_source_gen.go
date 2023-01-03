// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package macie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_macie_findings_filter", findingsFilterDataSource)
}

// findingsFilterDataSource returns the Terraform awscc_macie_findings_filter data source.
// This Terraform data source corresponds to the CloudFormation AWS::Macie::FindingsFilter resource.
func findingsFilterDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Action
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Findings filter action.",
		//	  "enum": [
		//	    "ARCHIVE",
		//	    "NOOP"
		//	  ],
		//	  "type": "string"
		//	}
		"action": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Findings filter action.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Findings filter ARN.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Findings filter ARN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Findings filter description",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Findings filter description",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FindingCriteria
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Findings filter criteria.",
		//	  "properties": {
		//	    "Criterion": {
		//	      "description": "Map of filter criteria.",
		//	      "patternProperties": {
		//	        "": {
		//	          "properties": {
		//	            "eq": {
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "gt": {
		//	              "format": "int64",
		//	              "type": "integer"
		//	            },
		//	            "gte": {
		//	              "format": "int64",
		//	              "type": "integer"
		//	            },
		//	            "lt": {
		//	              "format": "int64",
		//	              "type": "integer"
		//	            },
		//	            "lte": {
		//	              "format": "int64",
		//	              "type": "integer"
		//	            },
		//	            "neq": {
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"finding_criteria": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Criterion
				"criterion":               // Pattern: ""
				schema.MapNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: eq
							"eq": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: gt
							"gt": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: gte
							"gte": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: lt
							"lt": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: lte
							"lte": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: neq
							"neq": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Map of filter criteria.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Findings filter criteria.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FindingsFilterListItems
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Findings filters list.",
		//	  "items": {
		//	    "description": "Returned by ListHandler representing filter name and ID.",
		//	    "properties": {
		//	      "Id": {
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"findings_filter_list_items": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Findings filters list.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Findings filter ID.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Findings filter ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Findings filter name",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Findings filter name",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Position
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Findings filter position.",
		//	  "type": "integer"
		//	}
		"position": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Findings filter position.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Macie::FindingsFilter",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::FindingsFilter").WithTerraformTypeName("awscc_macie_findings_filter")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                     "Action",
		"arn":                        "Arn",
		"criterion":                  "Criterion",
		"description":                "Description",
		"eq":                         "eq",
		"finding_criteria":           "FindingCriteria",
		"findings_filter_list_items": "FindingsFilterListItems",
		"gt":                         "gt",
		"gte":                        "gte",
		"id":                         "Id",
		"lt":                         "lt",
		"lte":                        "lte",
		"name":                       "Name",
		"neq":                        "neq",
		"position":                   "Position",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
