---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_appsync_domain_name_api_association Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::AppSync::DomainNameApiAssociation
---

# awscc_appsync_domain_name_api_association (Resource)

Resource Type definition for AWS::AppSync::DomainNameApiAssociation

## Example Usage

```terraform

# This will associate custom domain with generated ACM certificates:

resource "awscc_appsync_domain_name" "example" {
  domain_name     = "api.example.com"
  certificate_arn = "arn:aws:acm:<aws_region>:certificate/<certificate_id>"
}

##Please note, when creating ACM certificate add domain in following format `*.example.com`

# This will associate appsync API with custom domain:

resource "awscc_appsync_domain_name_api_association" "example2" {
  api_id      = "<appsync_app_id>"
  domain_name  = awscc_appsync_domain_name.example.domain_name
}

```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_id` (String)
- `domain_name` (String)

### Read-Only

- `api_association_identifier` (String)
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_appsync_domain_name_api_association.example <resource ID>
```
