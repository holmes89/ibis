output "url" {
  value = aws_api_gateway_deployment.rest_api.invoke_url
}

output "ci_id" {
  value = aws_iam_access_key.gh.id
}

output "ci_secret" {
  value = aws_iam_access_key.gh.secret
}

output "api_function_name" {
  value = join("", aws_lambda_function.api.*.function_name)
}

output "cf_dist_id" {
  value = module.cdn.cf_id
}

output "ui_bucket" {
  value = module.cdn.s3_bucket
}