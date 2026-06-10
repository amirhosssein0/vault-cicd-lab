#!/bin/bash

# Create DB Password Secret
vault kv put secret/db-password value="super-secret-123"