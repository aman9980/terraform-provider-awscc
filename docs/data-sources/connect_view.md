---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_view Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Connect::View
---

# awscc_connect_view (Data Source)

Data Source schema for AWS::Connect::View



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `actions` (List of String) The actions of the view in an array.
- `description` (String) The description of the view.
- `instance_arn` (String) The Amazon Resource Name (ARN) of the instance.
- `name` (String) The name of the view.
- `tags` (Attributes Set) One or more tags. (see [below for nested schema](#nestedatt--tags))
- `template` (String) The template of the view as JSON.
- `view_arn` (String) The Amazon Resource Name (ARN) of the view.
- `view_content_sha_256` (String) The view content hash.
- `view_id` (String) The view id of the view.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters
- `value` (String) The value for the tag. . You can specify a value that is maximum of 256 Unicode characters