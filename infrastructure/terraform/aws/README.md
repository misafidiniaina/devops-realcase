# AWS foundation

This Terraform root module creates the first cloud foundation for the Cloud Platform Lab:

- a tagged VPC with public and private subnets across two availability zones;
- security groups for a future load balancer, application service, and PostgreSQL;
- a private, encrypted PostgreSQL RDS instance;
- immutable ECR repositories for the backend and frontend images;
- an ECS task execution role for the next deployment phase.

The application code is not managed or changed by this module. ECS services, the load balancer, NAT/VPC endpoints, DNS, and TLS will be added in later infrastructure phases.

## Prerequisites

- Terraform 1.7 or newer;
- AWS credentials configured through the AWS CLI, environment variables, or an approved identity provider;
- an AWS account and a region with at least two availability zones.

Terraform uses the standard AWS credential chain.  this repository.

## Initialize and review

From this directory:

```bash
cp terraform.tfvars.example terraform.tfvars
# Replace db_password in terraform.tfvars with a temporary development secret.
terraform init
terraform fmt -check
terraform validate
terraform plan
```

`terraform plan` requires AWS credentials to read the account and availability zones. It does not create resources.

## Apply and destroy

Only apply after reviewing the plan:

```bash
terraform apply
terraform destroy
```

The default settings are intended for a disposable development environment. In a shared or production environment, use a remote state backend, disable `db_skip_final_snapshot`, enable deletion protection, and source the database password from a secret manager.

## State and cost notes

Terraform state and local variable files must neveDo not put access keys or database passwords inr be committed. The RDS instance, ECR repositories, and other AWS resources may incur charges even when the application is not receiving traffic. Destroy the development environment when it is no longer needed.
