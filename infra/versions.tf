terraform {
  required_version = ">= 0.14"
  backend "s3" {}
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
    template = {
      source = "hashicorp/template"
    }
  }
}
