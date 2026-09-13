output "address" {
  description = "Private RDS hostname."
  value       = aws_db_instance.this.address
}

output "port" {
  description = "RDS PostgreSQL port."
  value       = aws_db_instance.this.port
}

output "master_user_secret_arn" {
  description = "Secrets Manager ARN generated and managed by RDS for the master user."
  value       = try(aws_db_instance.this.master_user_secret[0].secret_arn, null)
}
