
variable "region" {
  type        = string
  description = "AWS Region for S3 bucket"
  default     = "us-east-2"
}

variable "parent_zone_name" {
  type        = string
  description = "DNS zone name (E.g. staging.cloudposse.co)"
}

variable "root_zone_name" {
  type        = string
  description = "DNS root zone name (E.g. cloudposse.co)"
}

variable "aliases" {
  type        = list(string)
  description = "DNS aliases"
  default     = []
}

variable "cors_allowed_origins" {
  type        = list(string)
  description = "CORS allowed origins"
  default     = []
}

variable "website_enabled" {
  type        = bool
  description = "Website enabled"
  default     = false
}

variable "redirect_all_requests_to" {
  type        = string
  description = "A hostname to redirect all website requests for this distribution to. If this is set, it overrides other website settings"
  default     = ""
}

variable "origin_bucket" {
  type        = string
  description = "Origin bucket"
  default     = ""
}

variable "allow_headers" {
  description = "Allow headers"
  type        = list(string)

  default = [
    "Authorization",
    "Content-Type",
    "X-Amz-Date",
    "X-Amz-Security-Token",
    "X-Api-Key",
  ]
}

# var.allow_methods
variable "allow_methods" {
  description = "Allow methods"
  type        = list(string)

  default = [
    "OPTIONS",
    "HEAD",
    "GET",
    "POST",
    "PUT",
    "PATCH",
    "DELETE",
  ]
}

# var.allow_origin
variable "allow_origin" {
  description = "Allow origin"
  type        = string
  default     = "*"
}

# var.allow_max_age
variable "allow_max_age" {
  description = "Allow response caching time"
  type        = string
  default     = "7200"
}