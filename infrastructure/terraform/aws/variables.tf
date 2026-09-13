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

variable "db_password" {
  description = "Initial PostgreSQL administrator password. Pass this through a secret manager or tfvars excluded from Git."
  type        = string
  sensitive   = true
  nullable    = false
}

variable "db_instance_class" {
  description = "RDS instance class for the environment."
  type        = string
  default     = "db.t4g.micro"
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

variable "tags" {
  description = "Additional tags applied to all supported resources."
  type        = map(string)
  default     = {}
}
