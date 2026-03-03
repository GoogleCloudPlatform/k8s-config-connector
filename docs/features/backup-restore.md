# Standardized Backup and Restore for Config Connector

Config Connector provides a set of CLI commands to manage the backup and restoration of your Google Cloud resources managed by Kubernetes.

## Overview

The backup and restore feature allows you to:
- **Configure** automated, scheduled backups of all Config Connector resources in your cluster.
- **Create** immediate, manual backups to a Google Cloud Storage (GCS) bucket.
- **Status** monitor the health of recent backup jobs and view available backup artifacts.

## Getting Started

### Prerequisites

- A GKE cluster with Config Connector installed.
- Workload Identity enabled on the cluster.
- `config-connector` CLI installed and configured to access your cluster.

### 1. Configure Automated Backups

To set up a daily backup of all Config Connector resources:

```bash
config-connector backup configure \
    --project <PROJECT_ID> \
    --bucket <GCS_BUCKET_NAME> \
    --location <GCP_REGION>
```

This command performs the following actions:
- Creates a GCS bucket for storing backups (if it doesn't exist).
- Creates a GCP Service Account (`cnrm-backup`) with necessary permissions.
- Sets up Workload Identity for the backup manager.
- Creates a Kubernetes `CronJob` in the `cnrm-system` namespace that runs daily.

### 2. Manually Trigger a Backup

To create an immediate backup:

```bash
config-connector backup create --bucket <GCS_BUCKET_NAME> --project <PROJECT_ID>
```

Backups are stored in GCS under a path prefixed by the timestamp: `gs://<BUCKET_NAME>/YYYY-MM-DD-HH-MM-SS/`.

### 3. Check Backup Status

To see the status of recent backup jobs and available backup artifacts in GCS:

```bash
config-connector backup status --bucket <GCS_BUCKET_NAME> --project <PROJECT_ID>
```

## How it Works

### Backup Process
The `create` command discovers all resources in the cluster that belong to the Config Connector API groups (`*.cnrm.cloud.google.com`). For each resource:
1. It retrieves the current state from the Kubernetes API.
2. It sanitizes the resource by removing live fields (e.g., `uid`, `resourceVersion`, `managedFields`, `status`).
3. It saves the sanitized YAML to the specified GCS bucket.

## Best Practices

### GCS Bucket Management
- **Security:** Ensure that only authorized users/service accounts have access to the backup bucket.
- **Lifecycle:** Use GCS Lifecycle Management to rotate or delete old backup artifacts to manage costs.
- **Location:** Place the bucket in the same region as your cluster for better performance and reduced egress costs.
