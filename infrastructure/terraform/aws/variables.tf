variable "aws_region" {
  description = "AWS region where the foundation will be created."
  type        = string
  default     = "eu-west-1"
}

variable "project_name" {
  description = "Short project name used in resource names and tags."
  type        = string
  default     = "cloud-platform-lab"
}

variable "environment" {
  description = "Deployment environment name."
  type        = string
  default     = "dev"
}

variable "vpc_cidr" {
  description = "CIDR range for the application VPC."
  type        = string
  default     = "10.20.0.0/16"
}

variable "availability_zone_count" {
  description = "Number of availability zones used for the foundation."
  type        = number
  default     = 2

  validation {
    condition     = var.availability_zone_count >= 2
    error_message = "At least two availability zones are required."
  }
}

variable "nat_gateway_mode" {
  description = "NAT gateway topology: none, single, or per_az."
  type        = string
  default     = "none"

  validation {
    condition     = contains(["none", "single", "per_az"], var.nat_gateway_mode)
    error_message = "nat_gateway_mode must be one of: none, single, per_az."
  }
}

variable "enable_vpc_endpoints" {
  description = "Create private endpoints for AWS services used by ECS tasks."
  type        = bool
  default     = true
}

variable "app_port" {
  description = "Port used by the containerized backend service."
  type        = number
  default     = 8080
}

variable "db_name" {
  description = "Initial PostgreSQL database name."
  type        = string
  default     = "cloud_platform"
}

variable "db_username" {
  description = "Initial PostgreSQL administrator username."
  type        = string
  default     = "cloud_platform"
}

variable "db_instance_class" {
  description = "RDS instance class for the environment."
  type        = string
  default     = "db.t4g.micro"
}

variable "db_backup_retention_period" {
  description = "Number of days to retain automated RDS backups."
  type        = number
  default     = 7

  validation {
    condition     = var.db_backup_retention_period >= 0 && var.db_backup_retention_period <= 35
    error_message = "db_backup_retention_period must be between 0 and 35 days."
  }
}

variable "db_multi_az" {
  description = "Deploy a standby RDS instance in another availability zone."
  type        = bool
  default     = false
}

variable "db_skip_final_snapshot" {
  description = "Skip the final RDS snapshot when destroying the environment. Keep true only for disposable environments."
  type        = bool
  default     = true
}

variable "db_deletion_protection" {
  description = "Prevent accidental deletion of the RDS instance."
  type        = bool
  default     = false
}

variable "log_retention_in_days" {
  description = "CloudWatch application log retention period."
  type        = number
  default     = 30
}

variable "tags" {
  description = "Additional tags applied to all supported resources."
  type        = map(string)
  default     = {}
}
