variable "identifier" {
  description = "Unique RDS instance identifier."
  type        = string
}

variable "db_name" {
  description = "Initial PostgreSQL database name."
  type        = string
}

variable "db_username" {
  description = "Initial PostgreSQL administrator username."
  type        = string
}

variable "instance_class" {
  description = "RDS instance class."
  type        = string
}

variable "private_subnet_ids" {
  description = "Private subnet IDs for the DB subnet group."
  type        = list(string)
}

variable "security_group_id" {
  description = "Security group allowed to connect to PostgreSQL."
  type        = string
}

variable "backup_retention_period" {
  description = "Number of days to retain automated backups."
  type        = number
}

variable "multi_az" {
  description = "Deploy a standby DB instance in another availability zone."
  type        = bool
}

variable "skip_final_snapshot" {
  description = "Skip the final snapshot on destroy. Use only for disposable environments."
  type        = bool
}

variable "deletion_protection" {
  description = "Prevent accidental RDS deletion."
  type        = bool
}

variable "tags" {
  description = "Tags applied to database resources."
  type        = map(string)
  default     = {}
}
