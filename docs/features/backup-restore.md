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
- The Config Connector controller service account must have permissions to manage `StorageBucket` and `IAMServiceAccount` resources in the target GCP project (typically `roles/owner` or `roles/editor` at the project level, or specific roles like `roles/storage.admin` and `roles/iam.serviceAccountAdmin`).

### 1. Configure Automated Backups

To set up a daily backup of all Config Connector resources:

```bash
config-connector backup configure \
    --project <PROJECT_ID> \
    --bucket <GCS_BUCKET_NAME> \
    --location <GCP_REGION>
```

This command applies the following Config Connector resources to your cluster:
- **StorageBucket**: Creates a GCS bucket for storing backups.
- **IAMServiceAccount**: Creates a dedicated GCP Service Account (`cnrm-backup`) for the backup manager.
- **IAMPolicyMember**: Sets up Workload Identity and grants `roles/storage.objectAdmin` on the bucket.
- **ServiceAccount**: Creates a Kubernetes Service Account for the backup manager.
- **CronJob**: A daily job that runs `config-connector backup create`.

**Note:** The actual creation of the GCP bucket and service account is performed by the Config Connector controller.

### 2. Manually Trigger a Backup

To create an immediate backup:

```bash
config-connector backup create --bucket <GCS_BUCKET_NAME> --project <PROJECT_ID>
```

**Note:** This command requires that the target GCS bucket already exists. If you haven't run `backup configure`, you must create the bucket manually.

Backups are stored in GCS under a path prefixed by the cluster name and timestamp: `gs://<BUCKET_NAME>/<CLUSTER_NAME>/YYYY-MM-DD-HH-MM-SS/`.

### 3. Check Backup Status

To see the status of recent backup jobs and available backup artifacts in GCS:

```bash
config-connector backup status --bucket <GCS_BUCKET_NAME> --project <PROJECT_ID>
```

The output includes:
- **Kubernetes Jobs**: Status of recent `CronJob` executions (Succeeded/Failed).
- **GCS Artifacts**: A list of the 10 most recent backups found in the GCS bucket for the cluster, including the total number of resources backed up.

## How it Works

### Backup Process
The `create` command discovers all resources in the cluster that belong to the Config Connector API groups (`*.cnrm.cloud.google.com`). For each resource:
1. It retrieves the current state from the Kubernetes API.
2. It sanitizes the resource by removing internal Kubernetes fields (e.g., `uid`, `resourceVersion`, `managedFields`, `creationTimestamp`).
3. It preserves the `spec` and `status` fields to ensure a complete snapshot of the resource's configuration and current observed state.
4. It saves the sanitized YAML to the specified GCS bucket.
5. It writes a `summary.json` file containing the count of resources backed up per Kind.

## Best Practices

### GCS Bucket Management
- **Security:** Ensure that only authorized users/service accounts have access to the backup bucket.
- **Lifecycle:** Use GCS Lifecycle Management to rotate or delete old backup artifacts to manage costs.
- **Location:** Place the bucket in the same region as your cluster for better performance and reduced egress costs.
