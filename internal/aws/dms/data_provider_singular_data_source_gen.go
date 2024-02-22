// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package dms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_dms_data_provider", dataProviderDataSource)
}

// dataProviderDataSource returns the Terraform awscc_dms_data_provider data source.
// This Terraform data source corresponds to the CloudFormation AWS::DMS::DataProvider resource.
func dataProviderDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DataProviderArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The data provider ARN.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"data_provider_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The data provider ARN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataProviderCreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The data provider creation time.",
		//	  "maxLength": 40,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"data_provider_creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The data provider creation time.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataProviderIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes an identifier for the data provider. It is used for describing/deleting/modifying can be name/arn",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"data_provider_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes an identifier for the data provider. It is used for describing/deleting/modifying can be name/arn",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataProviderName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes a name to identify the data provider.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"data_provider_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes a name to identify the data provider.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The optional description of the data provider.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The optional description of the data provider.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Engine
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes a data engine for the data provider.",
		//	  "enum": [
		//	    "postgres",
		//	    "mysql",
		//	    "oracle",
		//	    "sqlserver",
		//	    "aurora",
		//	    "aurora_postgresql"
		//	  ],
		//	  "type": "string"
		//	}
		"engine": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes a data engine for the data provider.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ExactSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "The property describes the exact settings which can be modified",
		//	  "type": "boolean"
		//	}
		"exact_settings": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes the exact settings which can be modified",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Settings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property identifies the exact type of settings for the data provider.",
		//	  "properties": {
		//	    "MicrosoftSqlServerSettings": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CertificateArn": {
		//	          "type": "string"
		//	        },
		//	        "DatabaseName": {
		//	          "type": "string"
		//	        },
		//	        "Port": {
		//	          "type": "integer"
		//	        },
		//	        "ServerName": {
		//	          "type": "string"
		//	        },
		//	        "SslMode": {
		//	          "enum": [
		//	            "none",
		//	            "require",
		//	            "verify_ca",
		//	            "verify_full"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "MySqlSettings": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CertificateArn": {
		//	          "type": "string"
		//	        },
		//	        "Port": {
		//	          "type": "integer"
		//	        },
		//	        "ServerName": {
		//	          "type": "string"
		//	        },
		//	        "SslMode": {
		//	          "enum": [
		//	            "none",
		//	            "require",
		//	            "verify_ca",
		//	            "verify_full"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "OracleSettings": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AsmServer": {
		//	          "type": "string"
		//	        },
		//	        "CertificateArn": {
		//	          "type": "string"
		//	        },
		//	        "DatabaseName": {
		//	          "type": "string"
		//	        },
		//	        "Port": {
		//	          "type": "integer"
		//	        },
		//	        "SecretsManagerOracleAsmAccessRoleArn": {
		//	          "type": "string"
		//	        },
		//	        "SecretsManagerOracleAsmSecretId": {
		//	          "type": "string"
		//	        },
		//	        "SecretsManagerSecurityDbEncryptionAccessRoleArn": {
		//	          "type": "string"
		//	        },
		//	        "SecretsManagerSecurityDbEncryptionSecretId": {
		//	          "type": "string"
		//	        },
		//	        "ServerName": {
		//	          "type": "string"
		//	        },
		//	        "SslMode": {
		//	          "enum": [
		//	            "none",
		//	            "require",
		//	            "verify_ca",
		//	            "verify_full"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "PostgreSqlSettings": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CertificateArn": {
		//	          "type": "string"
		//	        },
		//	        "DatabaseName": {
		//	          "type": "string"
		//	        },
		//	        "Port": {
		//	          "type": "integer"
		//	        },
		//	        "ServerName": {
		//	          "type": "string"
		//	        },
		//	        "SslMode": {
		//	          "enum": [
		//	            "none",
		//	            "require",
		//	            "verify_ca",
		//	            "verify_full"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MicrosoftSqlServerSettings
				"microsoft_sql_server_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CertificateArn
						"certificate_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Port
						"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ServerName
						"server_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SslMode
						"ssl_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MySqlSettings
				"my_sql_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CertificateArn
						"certificate_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Port
						"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ServerName
						"server_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SslMode
						"ssl_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: OracleSettings
				"oracle_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AsmServer
						"asm_server": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: CertificateArn
						"certificate_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Port
						"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SecretsManagerOracleAsmAccessRoleArn
						"secrets_manager_oracle_asm_access_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SecretsManagerOracleAsmSecretId
						"secrets_manager_oracle_asm_secret_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SecretsManagerSecurityDbEncryptionAccessRoleArn
						"secrets_manager_security_db_encryption_access_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SecretsManagerSecurityDbEncryptionSecretId
						"secrets_manager_security_db_encryption_secret_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ServerName
						"server_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SslMode
						"ssl_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PostgreSqlSettings
				"postgre_sql_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CertificateArn
						"certificate_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Port
						"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ServerName
						"server_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SslMode
						"ssl_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The property identifies the exact type of settings for the data provider.",
			Computed:    true,
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
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		Description: "Data Source schema for AWS::DMS::DataProvider",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DMS::DataProvider").WithTerraformTypeName("awscc_dms_data_provider")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"asm_server":                    "AsmServer",
		"certificate_arn":               "CertificateArn",
		"data_provider_arn":             "DataProviderArn",
		"data_provider_creation_time":   "DataProviderCreationTime",
		"data_provider_identifier":      "DataProviderIdentifier",
		"data_provider_name":            "DataProviderName",
		"database_name":                 "DatabaseName",
		"description":                   "Description",
		"engine":                        "Engine",
		"exact_settings":                "ExactSettings",
		"key":                           "Key",
		"microsoft_sql_server_settings": "MicrosoftSqlServerSettings",
		"my_sql_settings":               "MySqlSettings",
		"oracle_settings":               "OracleSettings",
		"port":                          "Port",
		"postgre_sql_settings":          "PostgreSqlSettings",
		"secrets_manager_oracle_asm_access_role_arn":             "SecretsManagerOracleAsmAccessRoleArn",
		"secrets_manager_oracle_asm_secret_id":                   "SecretsManagerOracleAsmSecretId",
		"secrets_manager_security_db_encryption_access_role_arn": "SecretsManagerSecurityDbEncryptionAccessRoleArn",
		"secrets_manager_security_db_encryption_secret_id":       "SecretsManagerSecurityDbEncryptionSecretId",
		"server_name": "ServerName",
		"settings":    "Settings",
		"ssl_mode":    "SslMode",
		"tags":        "Tags",
		"value":       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
