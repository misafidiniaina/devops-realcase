locals {
  name = "${var.project_name}-${var.environment}"

  tags = merge(
    {
      Environment = var.environment
      ManagedBy   = "Terraform"
      Project     = var.project_name
    },
    var.tags
  )
}
