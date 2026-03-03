# Standardized Backup and Restore for Config Connector

Config Connector provides a set of CLI commands to manage the backup and restoration of your Google Cloud resources managed by Kubernetes.

## Overview

The backup and restore feature allows you to:
- **Configure** automated, scheduled backups of all Config Connector resources in your cluster.
- **Create** immediate, manual backups to a Google Cloud Storage (GCS) bucket.
- **Status** monitor the health of recent backup jobs and view available backup artifacts.
- **Restore** resources from a backup into a target cluster, with intelligent resource sorting and safety-first defaults.

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

### 4. Restore from a Backup

To restore resources from a specific backup into a target cluster:

```bash
config-connector backup restore \
    --source-bucket <GCS_BUCKET_NAME> \
    --backup-timestamp <TIMESTAMP> \
    --project <PROJECT_ID>
```

**Note:** It is highly recommended to use the `--dry-run` flag first to validate what will be restored.

## How it Works

### Backup Process
The `create` command discovers all resources in the cluster that belong to the Config Connector API groups (`*.cnrm.cloud.google.com`). For each resource:
1. It retrieves the current state from the Kubernetes API.
2. It sanitizes the resource by removing live fields (e.g., `uid`, `resourceVersion`, `managedFields`, `status`).
3. It saves the sanitized YAML to the specified GCS bucket.

### Restore Process
The `restore` command:
1. Loads all YAML files from the specified backup path in GCS.
2. Validates that the resource types (GVKs) are supported by the target cluster.
3. Sorts resources to handle basic dependencies (e.g., Projects before Networks, Networks before Subnetworks).
4. Sets the `cnrm.cloud.google.com/deletion-policy: abandon` annotation on all resources for safety.
5. Applies the resources using Server-Side Apply.

## Best Practices

### Safety First with Dry Run
Always perform a dry run before a real restore:
```bash
config-connector backup restore ... --dry-run
```

### Understanding the Deletion Policy
By default, restored resources are annotated with `cnrm.cloud.google.com/deletion-policy: abandon`. This ensures that if you accidentally delete the Kubernetes object, the underlying Google Cloud resource will NOT be deleted. Once you have verified the restored resources, you may choose to remove this annotation if you want Config Connector to manage the full lifecycle, including deletion.

### GCS Bucket Management
- **Security:** Ensure that only authorized users/service accounts have access to the backup bucket.
- **Lifecycle:** Use GCS Lifecycle Management to rotate or delete old backup artifacts to manage costs.
- **Location:** Place the bucket in the same region as your cluster for better performance and reduced egress costs.

### Resource Dependencies
While the restore command handles basic resource ordering (Containers -> Networking -> Others -> IAM), complex dependency chains may require manual intervention or multiple restore passes.

### Cross-Cluster Restore
When restoring to a different cluster, ensure that:
- The target cluster has Config Connector installed and configured with appropriate permissions.
- Any necessary Namespaces are already created in the target cluster, as the backup tool primarily focuses on KCC-managed resources.
