// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ecr_repository", repositoryDataSource)
}

// repositoryDataSource returns the Terraform awscc_ecr_repository data source.
// This Terraform data source corresponds to the CloudFormation AWS::ECR::Repository resource.
func repositoryDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EncryptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.\n\nBy default, when no encryption configuration is set or the AES256 encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.\n\nFor more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html",
		//	  "properties": {
		//	    "EncryptionType": {
		//	      "description": "The encryption type to use.",
		//	      "enum": [
		//	        "AES256",
		//	        "KMS"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "KmsKey": {
		//	      "description": "If you use the KMS encryption type, specify the CMK to use for encryption. The alias, key ID, or full ARN of the CMK can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed CMK for Amazon ECR will be used.",
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "EncryptionType"
		//	  ],
		//	  "type": "object"
		//	}
		"encryption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EncryptionType
				"encryption_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The encryption type to use.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: KmsKey
				"kms_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "If you use the KMS encryption type, specify the CMK to use for encryption. The alias, key ID, or full ARN of the CMK can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed CMK for Amazon ECR will be used.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.\n\nBy default, when no encryption configuration is set or the AES256 encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.\n\nFor more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageScanningConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The image scanning configuration for the repository. This setting determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
		//	  "properties": {
		//	    "ScanOnPush": {
		//	      "description": "The setting that determines whether images are scanned after being pushed to a repository.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"image_scanning_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ScanOnPush
				"scan_on_push": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "The setting that determines whether images are scanned after being pushed to a repository.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The image scanning configuration for the repository. This setting determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageTagMutability
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The image tag mutability setting for the repository.",
		//	  "enum": [
		//	    "MUTABLE",
		//	    "IMMUTABLE"
		//	  ],
		//	  "type": "string"
		//	}
		"image_tag_mutability": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The image tag mutability setting for the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LifecyclePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The LifecyclePolicy property type specifies a lifecycle policy. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html",
		//	  "properties": {
		//	    "LifecyclePolicyText": {
		//	      "description": "The JSON repository policy text to apply to the repository.",
		//	      "maxLength": 30720,
		//	      "minLength": 100,
		//	      "type": "string"
		//	    },
		//	    "RegistryId": {
		//	      "description": "The AWS account ID associated with the registry that contains the repository. If you do not specify a registry, the default registry is assumed. ",
		//	      "maxLength": 12,
		//	      "minLength": 12,
		//	      "pattern": "^[0-9]{12}$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"lifecycle_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LifecyclePolicyText
				"lifecycle_policy_text": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The JSON repository policy text to apply to the repository.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RegistryId
				"registry_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The AWS account ID associated with the registry that contains the repository. If you do not specify a registry, the default registry is assumed. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The LifecyclePolicy property type specifies a lifecycle policy. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.",
		//	  "maxLength": 256,
		//	  "minLength": 2,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryPolicyText
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. ",
		//	  "type": "string"
		//	}
		"repository_policy_text": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"repository_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
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
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ECR::Repository",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::Repository").WithTerraformTypeName("awscc_ecr_repository")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                          "Arn",
		"encryption_configuration":     "EncryptionConfiguration",
		"encryption_type":              "EncryptionType",
		"image_scanning_configuration": "ImageScanningConfiguration",
		"image_tag_mutability":         "ImageTagMutability",
		"key":                          "Key",
		"kms_key":                      "KmsKey",
		"lifecycle_policy":             "LifecyclePolicy",
		"lifecycle_policy_text":        "LifecyclePolicyText",
		"registry_id":                  "RegistryId",
		"repository_name":              "RepositoryName",
		"repository_policy_text":       "RepositoryPolicyText",
		"repository_uri":               "RepositoryUri",
		"scan_on_push":                 "ScanOnPush",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
