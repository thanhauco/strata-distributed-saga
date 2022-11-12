# Terraform Deployment Guide

```bash
cd terraform
terraform init
terraform workspace new staging
terraform apply -var="environment=staging"
```
