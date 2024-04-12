variable "domain_name" {
  description = "The domain name for the website"
  type        = string
}

variable "index_content" {
  description = "HTML content for the index page"
  type        = string
  default     = <<EOF
<html>
<head>
<title>Hello World</title>
</head>
<body>
<h1>Hello World!</h1>
</body>
</html>
EOF
}
