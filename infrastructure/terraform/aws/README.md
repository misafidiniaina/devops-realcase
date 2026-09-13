# AWS foundation

This Terraform root module defines the first AWS foundation for the Cloud Platform Lab. The application code is intentionally outside this module and is not changed by it.

## Architecture

- A tagged VPC with public and private subnets across two availability zones.
- Configurable NAT topology: `none` for a restricted development environment, `single` when general outbound access is needed, or `per_az` for high availability.
- Private interface endpoints for ECR, CloudWatch Logs, and Secrets Manager, plus an S3 gateway endpoint.
- Dedicated security groups for the future load balancer, application service, and PostgreSQL database.
- A private, encrypted PostgreSQL RDS instance with RDS-managed master credentials in Secrets Manager.
- Immutable, scan-on-push ECR repositories for backend and frontend images.
- A least-privilege ECS task execution role for ECR image pulls and CloudWatch logging.
- A CloudWatch log group with bounded retention for the future ECS service.

The module does not yet create ECS services, an application load balancer, DNS, or TLS certificates. Those should be added in the deployment phase after the foundation is reviewed.

## Prerequisites

- Terraform 1.10 or newer;
- an AWS account and a region with at least two availability zones;
- AWS credentials supplied through the AWS CLI, environment variables, or an approved identity provider;
- an S3 bucket for remote Terraform state, created outside this stack.

Do not put AWS access keys, database passwords, or other secrets in this repository. Terraform uses the standard AWS credential chain, and RDS generates and rotates the master password through Secrets Manager.

## Remote state setup

Create a dedicated state bucket with versioning and public access blocked. Then configure the bucket name locally:

```bash
cp backend.hcl.example backend.hcl
# Edit backend.hcl and set the real state bucket name.
terraform init -backend-config=backend.hcl
```

The S3 backend uses native state locking through `use_lockfile = true`. Never commit `backend.hcl`, local state, or Terraform variable files.

## Initialize and review

From this directory:

```bash
terraform fmt -check -recursive
terraform validate
terraform plan
```

`terraform plan` reads AWS account metadata and shows the proposed changes. It does not create resources.

Useful environment-specific overrides include:

```bash
terraform plan \
  -var='environment=dev' \
  -var='nat_gateway_mode=single' \
  -var='db_skip_final_snapshot=true'
```

For a shared or production environment, use `nat_gateway_mode=per_az`, enable `db_multi_az`, set `db_deletion_protection=true`, and set `db_skip_final_snapshot=false`.

## Apply and destroy

Only apply after reviewing the plan:

```bash
terraform apply
terraform destroy
```

The default values are intended for a disposable development environment. NAT gateways, interface endpoints, RDS, ECR, and CloudWatch resources may incur AWS charges. Destroy development resources when they are no longer needed.
