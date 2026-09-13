variable "name" {
  description = "Repository name prefix."
  type        = string
}

variable "repository_names" {
  description = "Container repositories to create."
  type        = set(string)
  default     = ["backend", "frontend"]
}

variable "tags" {
  description = "Tags applied to ECR resources."
  type        = map(string)
  default     = {}
}
