provider "aws" {
  alias  = "useast1"
  region = "us-east-1"
}


locals {
  ui_domain_name = "${module.this.name}.${var.parent_zone_name}"
}

module "ui_acm_request_certificate" {
  providers = {
    aws = aws.useast1
  }
  source                      = "cloudposse/acm-request-certificate/aws"
  version                     = "0.16.0"
  context                     = module.this.context
  zone_name                   = var.parent_zone_name
  domain_name                 = local.ui_domain_name
  ttl                         = 300
  wait_for_certificate_issued = true
  enabled                     = true
  subject_alternative_names   = []
}

module "cdn" {
  source                   = "cloudposse/cloudfront-s3-cdn/aws"
  version                  = "0.80.0"
  context                  = module.this.context
  attributes               = ["ui"]
  acm_certificate_arn      = module.ui_acm_request_certificate.arn
  aliases                  = var.aliases
  parent_zone_name         = var.parent_zone_name
  parent_zone_id           = null
  website_enabled          = var.website_enabled
  redirect_all_requests_to = var.redirect_all_requests_to
  dns_alias_enabled        = true
  origin_force_destroy     = true
  cors_allowed_headers     = ["*"]
  cors_allowed_methods     = ["GET", "HEAD"]
  cors_allowed_origins     = var.cors_allowed_origins
  cors_expose_headers      = ["ETag"]
  error_document           = "index.html"

  allow_ssl_requests_only = false

  depends_on = [module.ui_acm_request_certificate.validation_id]
}