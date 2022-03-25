provider "aws" {
  assume_role {
    role_arn = var.aws_assume_role_arn
  }
  region = var.region
}

data "archive_file" "dummy" {
  type        = "zip"
  output_path = "${path.module}/main.zip"

  source {
    content  = "hello"
    filename = "dummy.txt"
  }
}