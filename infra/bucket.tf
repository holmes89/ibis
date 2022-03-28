resource "aws_s3_bucket" "games" {
  bucket = "${module.this.name}-games"
}