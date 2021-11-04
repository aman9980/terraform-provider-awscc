---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_s3outposts_bucket Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::S3Outposts::Bucket
---

# awscc_s3outposts_bucket (Data Source)

Data Source schema for AWS::S3Outposts::Bucket



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the specified bucket.
- **bucket_name** (String) A name for the bucket.
- **lifecycle_configuration** (Attributes) Rules that define how Amazon S3Outposts manages objects during their lifetime. (see [below for nested schema](#nestedatt--lifecycle_configuration))
- **outpost_id** (String) The id of the customer outpost on which the bucket resides.
- **tags** (Attributes List) An arbitrary set of tags (key-value pairs) for this S3Outposts bucket. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--lifecycle_configuration"></a>
### Nested Schema for `lifecycle_configuration`

Read-Only:

- **rules** (Attributes List) A list of lifecycle rules for individual objects in an Amazon S3Outposts bucket. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules))

<a id="nestedatt--lifecycle_configuration--rules"></a>
### Nested Schema for `lifecycle_configuration.rules`

Read-Only:

- **abort_incomplete_multipart_upload** (Attributes) Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3Outposts bucket. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--abort_incomplete_multipart_upload))
- **expiration_date** (String) Indicates when objects are deleted from Amazon S3Outposts. The date value must be in ISO 8601 format. The time is always midnight UTC.
- **expiration_in_days** (Number) Indicates the number of days after creation when objects are deleted from Amazon S3Outposts.
- **filter** (Attributes) The container for the filter of the lifecycle rule. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--filter))
- **id** (String) Unique identifier for the lifecycle rule. The value can't be longer than 255 characters.
- **status** (String)

<a id="nestedatt--lifecycle_configuration--rules--abort_incomplete_multipart_upload"></a>
### Nested Schema for `lifecycle_configuration.rules.abort_incomplete_multipart_upload`

Read-Only:

- **days_after_initiation** (Number) Specifies the number of days after which Amazon S3Outposts aborts an incomplete multipart upload.


<a id="nestedatt--lifecycle_configuration--rules--filter"></a>
### Nested Schema for `lifecycle_configuration.rules.filter`

Read-Only:

- **and_operator** (Attributes) The container for the AND condition for the lifecycle rule. A combination of Prefix and 1 or more Tags OR a minimum of 2 or more tags. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--filter--and_operator))
- **prefix** (String) Object key prefix that identifies one or more objects to which this rule applies.
- **tag** (Attributes) Specifies a tag used to identify a subset of objects for an Amazon S3Outposts bucket. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--filter--tag))

<a id="nestedatt--lifecycle_configuration--rules--filter--and_operator"></a>
### Nested Schema for `lifecycle_configuration.rules.filter.tag`

Read-Only:

- **prefix** (String) Prefix identifies one or more objects to which the rule applies.
- **tags** (Attributes List) All of these tags must exist in the object's tag set in order for the rule to apply. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--filter--tag--tags))

<a id="nestedatt--lifecycle_configuration--rules--filter--tag--tags"></a>
### Nested Schema for `lifecycle_configuration.rules.filter.tag.tags`

Read-Only:

- **key** (String)
- **value** (String)



<a id="nestedatt--lifecycle_configuration--rules--filter--tag"></a>
### Nested Schema for `lifecycle_configuration.rules.filter.tag`

Read-Only:

- **key** (String)
- **value** (String)





<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)

