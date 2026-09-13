variable "name" {
  description = "Base name used for security groups."
  type        = string
}

variable "vpc_id" {
  description = "VPC ID where security groups are created."
  type        = string
}

variable "app_port" {
  description = "Port exposed by the application service."
  type        = number
}

variable "tags" {
  description = "Tags applied to security groups."
  type        = map(string)
  default     = {}
}
