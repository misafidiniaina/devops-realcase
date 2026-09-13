output "repository_urls" {
  description = "ECR repository URLs keyed by component."
  value       = { for component, repository in aws_ecr_repository.this : component => repository.repository_url }
}

output "repository_arns" {
  description = "ECR repository ARNs keyed by component."
  value       = { for component, repository in aws_ecr_repository.this : component => repository.arn }
}
