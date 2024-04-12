#!/bin/bash
echo "Deploying the infrastructure..."
terraform init
terraform apply -auto-approve
