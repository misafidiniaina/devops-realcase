output "vpc_id" {
  description = "Application VPC ID."
  value       = aws_vpc.main.id
}

output "public_subnet_ids" {
  description = "Public subnet IDs for load balancers."
  value       = aws_subnet.public[*].id
}

output "private_subnet_ids" {
  description = "Private subnet IDs for application and database workloads."
  value       = aws_subnet.private[*].id
}

output "load_balancer_security_group_id" {
  description = "Security group ID for the public load balancer."
  value       = aws_security_group.load_balancer.id
}

output "application_security_group_id" {
  description = "Security group ID for application workloads."
  value       = aws_security_group.application.id
}

output "database_security_group_id" {
  description = "Security group ID for PostgreSQL."
  value       = aws_security_group.database.id
}

output "database_endpoint" {
  description = "Private PostgreSQL endpoint."
  value       = aws_db_instance.postgres.address
}

output "database_port" {
  description = "PostgreSQL port."
  value       = aws_db_instance.postgres.port
}

output "ecr_repository_urls" {
  description = "ECR repository URLs keyed by application component."
  value       = { for component, repository in aws_ecr_repository.application : component => repository.repository_url }
}

output "ecs_task_execution_role_arn" {
  description = "IAM role ARN for future ECS task definitions."
  value       = aws_iam_role.ecs_task_execution.arn
}
