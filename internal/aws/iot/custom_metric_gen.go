// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iot_custom_metric", customMetricResourceType)
}

// customMetricResourceType returns the Terraform awscc_iot_custom_metric resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::CustomMetric resource type.
func customMetricResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"display_name": {
			// Property: DisplayName
			// CloudFormation resource type schema:
			// {
			//   "description": "Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.",
			//   "maxLength": 128,
			//   "type": "string"
			// }
			Description: "Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.",
			Type:        types.StringType,
			Optional:    true,
		},
		"metric_arn": {
			// Property: MetricArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Number (ARN) of the custom metric.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Number (ARN) of the custom metric.",
			Type:        types.StringType,
			Computed:    true,
		},
		"metric_name": {
			// Property: MetricName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// MetricName is a force-new attribute.
		},
		"metric_type": {
			// Property: MetricType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.",
			//   "enum": [
			//     "string-list",
			//     "ip-address-list",
			//     "number-list",
			//     "number"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.",
			Type:        types.StringType,
			Required:    true,
			// MetricType is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The tag's key.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The tag's value.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Description: "The tag's key.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The tag's value.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "A custom metric published by your devices to Device Defender.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::CustomMetric").WithTerraformTypeName("awscc_iot_custom_metric").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iot_custom_metric", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}