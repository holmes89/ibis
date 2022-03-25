data "aws_iam_policy_document" "ci" {
  statement {
    actions = [
      "lambda:CreateFunction",
      "lambda:UpdateFunctionConfiguration",
      "lambda:UpdateFunctionCode",
    ]
    resources = [
      join("", aws_lambda_function.api.*.arn),
    ]
  }
  statement {
    actions = [
      "s3:DeleteObject",
      "s3:GetBucketLocation",
      "s3:GetObject",
      "s3:ListBucket",
      "s3:PutObject",
      "s3:PutObjectAcl",
      "s3:GetObjectAcl"
    ]
    resources = [
      module.cdn.s3_bucket_arn,
    ]
  }
}

resource "aws_iam_user_policy" "gh_ro" {
  name = "${module.this.name}-ci"
  user = aws_iam_user.gh.name

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = data.aws_iam_policy_document.ci.json
}

resource "aws_iam_access_key" "gh" {
  user    = aws_iam_user.gh.name
}

resource "aws_iam_user" "gh" {
  name = "ibis-ci"
  path = "/ci/"
}
