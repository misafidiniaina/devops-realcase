# Platform Architecture

## Objective

Build a production-style platform capable of securely delivering,
deploying, observing and recovering a containerized application.

## Application

The application is an Infrastructure Inventory Platform.

It consists of:

- React frontend
- Go REST API
- PostgreSQL database

## Platform

The application will progressively be deployed using:

- Docker
- Kubernetes
- AWS EKS
- GCP GKE
- Terraform
- Argo CD

## Delivery

The delivery pipeline will follow:

Developer
→ GitHub
→ CI
→ Security Checks
→ Container Registry
→ GitOps
→ Argo CD
→ Kubernetes

## Observability

The platform will use:

- Prometheus
- Grafana
- Loki
- Alertmanager

## Security

Security will be integrated throughout the lifecycle:

- SAST
- Dependency scanning
- Container scanning
- Secret detection
- IaC scanning
- Kubernetes hardening
- IAM least privilege
- Network policies

## Reliability

The platform will demonstrate:

- Health checks
- Automatic pod recovery
- Horizontal scaling
- Rollbacks
- Monitoring and alerting
- Backup and restoration
- Disaster recovery
