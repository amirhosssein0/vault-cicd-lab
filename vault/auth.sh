#!/bin/bash

# Enable Kubernetes Auth Method
vault auth enable kubernetes

# Configure Kubernetes Auth
vault write auth/kubernetes/config \
  kubernetes_host="https://kubernetes.default.svc:443"

# Write Policy
vault policy write app-policy /vault/policy.hcl

# Create Role
vault write auth/kubernetes/role/app-role \
  bound_service_account_names=app \
  bound_service_account_namespaces=develop,staging \
  policies=app-policy \
  ttl=24h