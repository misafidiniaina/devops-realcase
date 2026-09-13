output "vpc_id" {
  description = "Application VPC ID."
  value       = module.network.vpc_id
}

output "public_subnet_ids" {
  description = "Public subnet IDs for load balancers."
  value       = module.network.public_subnet_ids
}

output "private_subnet_ids" {
  description = "Private subnet IDs for application and database workloads."
  value       = module.network.private_subnet_ids
}

output "load_balancer_security_group_id" {
  description = "Security group ID for the public load balancer."
  value       = module.security.load_balancer_security_group_id
}

output "application_security_group_id" {
  description = "Security group ID for application workloads."
  value       = module.security.application_security_group_id
}

output "database_security_group_id" {
  description = "Security group ID for PostgreSQL."
  value       = module.security.database_security_group_id
}

output "database_endpoint" {
  description = "Private PostgreSQL endpoint."
  value       = module.database.address
}

output "database_port" {
  description = "PostgreSQL port."
  value       = module.database.port
}

output "database_master_user_secret_arn" {
  description = "Secrets Manager ARN for the RDS-managed master credentials."
  value       = module.database.master_user_secret_arn
}

output "ecr_repository_urls" {
  description = "ECR repository URLs keyed by application component."
  value       = module.ecr.repository_urls
}

output "ecs_task_execution_role_arn" {
  description = "IAM role ARN for future ECS task definitions."
  value       = module.iam.ecs_task_execution_role_arn
}
