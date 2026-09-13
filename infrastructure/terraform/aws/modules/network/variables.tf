variable "name" {
  description = "Base name used for network resources."
  type        = string
}

variable "vpc_cidr" {
  description = "CIDR range for the VPC."
  type        = string
}

variable "availability_zone_count" {
  description = "Number of availability zones to use."
  type        = number
}

variable "nat_gateway_mode" {
  description = "NAT gateway topology: none, single, or per_az."
  type        = string

  validation {
    condition     = contains(["none", "single", "per_az"], var.nat_gateway_mode)
    error_message = "nat_gateway_mode must be one of: none, single, per_az."
  }
}

variable "enable_vpc_endpoints" {
  description = "Create private VPC endpoints for AWS services used by ECS tasks."
  type        = bool
}

variable "interface_endpoint_services" {
  description = "AWS services exposed through private interface endpoints."
  type        = set(string)
  default     = ["ecr.api", "ecr.dkr", "logs", "secretsmanager"]
}

variable "tags" {
  description = "Tags applied to network resources."
  type        = map(string)
  default     = {}
}
