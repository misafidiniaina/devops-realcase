variable "name" {
  description = "Base name used for IAM resources."
  type        = string
}

variable "ecr_repository_arns" {
  description = "ECR repositories the ECS task execution role may pull from."
  type        = map(string)
}

variable "log_group_arn" {
  description = "CloudWatch log group ARN used by ECS tasks."
  type        = string
}

variable "database_secret_arn" {
  description = "RDS-managed Secrets Manager ARN read by ECS tasks."
  type        = string
}

variable "tags" {
  description = "Tags applied to IAM resources."
  type        = map(string)
  default     = {}
}
