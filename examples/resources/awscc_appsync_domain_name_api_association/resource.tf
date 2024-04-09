resource "awscc_appsync_domain_name" "example" {
  domain_name     = "api.example.com"
  certificate_arn = "arn:aws:acm:<aws_region>:certificate/<certificate_id>"
}

resource "awscc_appsync_domain_name_api_association" "example2" {
  api_id      = "<appsync_app_id>"
  domain_name  = awscc_appsync_domain_name.example.domain_name
}
