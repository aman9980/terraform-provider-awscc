---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ce_cost_category Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CE::CostCategory
---

# awscc_ce_cost_category (Data Source)

Data Source schema for AWS::CE::CostCategory



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) Cost category ARN
- **default_value** (String) The default value for the cost category
- **effective_start** (String) ISO 8601 date time with offset format
- **name** (String)
- **rule_version** (String)
- **rules** (String) JSON array format of Expression in Billing and Cost Management API
- **split_charge_rules** (String) Json array format of CostCategorySplitChargeRule in Billing and Cost Management API

