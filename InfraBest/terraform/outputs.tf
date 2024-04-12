output "website_url" {
  value = aws_s3_bucket.website.website_endpoint
}

output "cloudfront_distribution_id" {
  value = aws_cloudfront_distribution.website_cdn.id
}
