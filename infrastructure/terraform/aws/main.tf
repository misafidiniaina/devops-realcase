module "network" {
  source = "./modules/network"

  name                    = local.name
  vpc_cidr                = var.vpc_cidr
  availability_zone_count = var.availability_zone_count
  nat_gateway_mode        = var.nat_gateway_mode
  enable_vpc_endpoints    = var.enable_vpc_endpoints
  tags                    = local.tags
}

module "security" {
  source = "./modules/security"

  name     = local.name
  vpc_id   = module.network.vpc_id
  app_port = var.app_port
  tags     = local.tags
}

resource "aws_cloudwatch_log_group" "application" {
  name              = "/ecs/${local.name}"
  retention_in_days = var.log_retention_in_days

  tags = merge(local.tags, {
    Name = "/ecs/${local.name}"
  })
}

module "database" {
  source = "./modules/database"

  identifier              = "${local.name}-postgres"
  db_name                 = var.db_name
  db_username             = var.db_username
  instance_class          = var.db_instance_class
  private_subnet_ids      = module.network.private_subnet_ids
  security_group_id       = module.security.database_security_group_id
  backup_retention_period = var.db_backup_retention_period
  multi_az                = var.db_multi_az
  skip_final_snapshot     = var.db_skip_final_snapshot
  deletion_protection     = var.db_deletion_protection
  tags                    = local.tags
}

module "ecr" {
  source = "./modules/ecr"

  name = local.name
  tags = local.tags
}

module "iam" {
  source = "./modules/iam"

  name                = local.name
  ecr_repository_arns = module.ecr.repository_arns
  log_group_arn       = aws_cloudwatch_log_group.application.arn
  database_secret_arn = module.database.master_user_secret_arn
  tags                = local.tags
}
