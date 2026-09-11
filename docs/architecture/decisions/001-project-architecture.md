# ADR-001: Project Architecture

## Status

Accepted

## Context

The project must demonstrate practical Cloud, DevOps and DevSecOps
engineering capabilities rather than simply deploying a basic application.

## Decision

The platform will use:

- React for the frontend
- Go for the backend
- PostgreSQL for persistence
- Docker for containerization
- Terraform for infrastructure provisioning
- Kubernetes for workload orchestration
- GitHub Actions for CI
- Argo CD for GitOps
- Prometheus and Grafana for metrics
- Loki for logs
- Alertmanager for alerting

AWS EKS and GCP GKE will be used to demonstrate transferable
cloud engineering skills.

## Consequences

The project will require more infrastructure than a simple application,
but will provide stronger evidence of production-oriented engineering
skills.

The application itself will remain intentionally simple so that the
focus stays on the platform.
