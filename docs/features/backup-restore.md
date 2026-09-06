# Standardized Backup and Disaster Recovery (DR) for Config Connector

Config Connector provides a native, declarative Backup and Disaster Recovery (DR) suite within the `config-connector` CLI. This feature suite addresses critical business continuity requirements (such as Recovery Time Objective RTO < 4h and Recovery Point Objective RPO) for enterprise workloads managing Google Cloud infrastructure declaratively.

---

## Overview

The native Backup and DR toolset enables:
- **Automated Continuous Backups**: Declarative scheduling via Kubernetes `CronJob` with dedicated GCP Service Accounts and Workload Identity.
- **On-Demand Snapshots**: Immediate snapshots targeting either a Google Cloud Storage (GCS) bucket (`--bucket`) or a local directory (`--output-dir`).
- **Resilient Re-Acquisition (In-Place DR)**: Restoring K8s manifests without recreating or modifying underlying GCP cloud resources, eliminating data loss and downtime.
- **Cross-Region DR Failover**: Automated cross-region translation (`--target-region`, `--region-mapping`), namespace auto-provisioning (`--auto-create-namespaces`), and target project re-targeting (`--target-project`).
- **12-Tier Topological DAG Sorting**: Hierarchical ordering ensuring foundational resources (KMS, VPCs, IAM Service Accounts) are applied before dependent downstream infrastructure (Cloud SQL, Cloud Storage, Pub/Sub, IAM Policy Members).
- **Heterogeneous Cluster Support**: `--skip-missing-crds` gracefully bypasses resources whose CRDs are not installed in a target secondary cluster.
- **Audit & Monitoring**: `config-connector backup status` inspects automated backup job execution and lists snapshot inventories.

---

## Architecture & How It Works

### 1. Backup Pipeline (`config-connector backup create`)
The backup engine discovers all resources managed by Config Connector (`*.cnrm.cloud.google.com`):
1. **Discovery**: Queries the Kubernetes API server for all Config Connector Custom Resource Definitions (CRDs), excluding internal operator CRs (`core.cnrm.cloud.google.com`).
2. **Sanitization**: Strips internal Kubernetes metadata to ensure clean portability across clusters and namespaces:
   - `metadata.uid`
   - `metadata.resourceVersion`
   - `metadata.generation`
   - `metadata.managedFields`
   - `metadata.creationTimestamp`
   - `metadata.annotations["kubectl.kubernetes.io/last-applied-configuration"]`
3. **Artifact Structure**: Organizes resources cleanly by cluster, timestamp, namespace, and kind:
   ```text
   gs://<BUCKET_NAME>/<CLUSTER_NAME>/<TIMESTAMP>/
     summary.json
     <NAMESPACE>/
       storagebucket/
         <NAME>.yaml
       secretmanagersecret/
         <NAME>.yaml
       ...
   ```
4. **Summary Manifest**: Emits `summary.json` containing total counts per resource kind for fast integrity verification.

### 2. Restore Pipeline (`config-connector restore` / `config-connector backup restore`)
The restore engine rehydrates resources into the target cluster safely:
1. **Source Loading**: Reads manifests from either a GCS bucket (`--source-bucket`) or local directory (`--from-dir`), resolving explicit timestamps or `--backup-timestamp=latest`.
2. **CRD Discovery & Filtering**: Checks the target cluster's installed CRDs. If `--skip-missing-crds` is specified, unknown CRDs are logged and skipped rather than aborting.
3. **Selective Scope**: Supports restoring a single tenant or workload via `--filter-namespace=<namespace>`.
4. **Dynamic Namespace Creation**: When `--auto-create-namespaces` is specified, the CLI checks if target namespaces exist in the cluster and automatically provisions them before applying resources.
5. **DR Spec Transformations**:
   - **Regional Remapping**: Automatically rewrites `spec.location` and `spec.region` using `--target-region=<region>` or explicit `--region-mapping="us-central1=us-east1"`.
   - **Project Override**: Updates `spec.projectRef.external` and `metadata.annotations["cnrm.cloud.google.com/project-id"]` when `--target-project` is provided.
6. **Safe Cloud Re-Acquisition**:
   - Strips `status` blocks so Server-Side Apply succeeds without schema conflicts.
   - Injects `cnrm.cloud.google.com/deletion-policy: abandon` to ensure deletion of K8s objects never deletes GCP cloud infrastructure.
   - Injects `cnrm.cloud.google.com/management-conflict-prevention-policy: none` to enable instantaneous adoption by the controller manager.
7. **12-Tier Topological DAG Ordering**:
   Resources are topologically sorted and applied in strict dependency order:
   - **Tier 0**: Namespaces
   - **Tier 1**: Organizations, Folders
   - **Tier 2**: Projects
   - **Tier 3**: KMS KeyRings, CryptoKeys
   - **Tier 4**: Compute Networks, Subnetworks, Routers, Firewalls
   - **Tier 5**: IAM Service Accounts
   - **Tier 6**: IAM Custom Roles
   - **Tier 7**: Artifact Registry Repositories, Container Registry
   - **Tier 8**: Storage Buckets, Cloud SQL Instances, BigQuery Datasets, Spanner Instances
   - **Tier 9**: Pub/Sub Topics, Secret Manager Secrets
   - **Tier 10**: Pub/Sub Subscriptions, SQL Databases/Users, Secret Versions
   - **Tier 11**: Workloads, Deployments, Services
   - **Tier 12**: IAM Policy Members, IAM Policies, IAM Audit Configs
8. **Server-Side Apply**: Applies resources via Kubernetes Server-Side Apply using the field manager `kcc-backup-restore`.

---

## CLI Reference

### 1. `config-connector backup configure`

Provisions declarative infrastructure for scheduled automated backups:

```bash
config-connector backup configure \
    --project <PROJECT_ID> \
    --bucket <GCS_BUCKET_NAME> \
    --location <GCP_REGION> \
    --schedule "0 2 * * *"
```

**Flags**:
- `--project` *(required)*: GCP project ID where backup resources reside.
- `--bucket` *(required)*: GCS bucket name to store backups.
- `--location`: GCP region for the storage bucket (default: `us-central1`).
- `--schedule`: Cron schedule expression (default: `@daily`).

---

### 2. `config-connector backup create`

Executes an on-demand backup to GCS or local disk:

```bash
# Backup to Google Cloud Storage
config-connector backup create \
    --bucket <GCS_BUCKET_NAME> \
    --project <PROJECT_ID> \
    --cluster <CLUSTER_NAME>

# Backup to Local Filesystem
config-connector backup create \
    --output-dir /var/backups/kcc \
    --cluster <CLUSTER_NAME>
```

**Flags**:
- `--bucket`: Target GCS bucket for backup storage.
- `--output-dir`: Target local directory for backup storage (mutually exclusive with `--bucket`).
- `--cluster`: Cluster identifier (defaults to current kubeconfig context name).
- `--project`: Target GCP project ID for cluster identification.

---

### 3. `config-connector backup status`

Inspects recent backup job runs and available GCS snapshots:

```bash
config-connector backup status \
    --bucket <GCS_BUCKET_NAME> \
    --project <PROJECT_ID> \
    --cluster <CLUSTER_NAME>
```

**Flags**:
- `--bucket` *(required)*: Target GCS bucket to inspect.
- `--cluster`: Cluster identifier (defaults to current kubeconfig context).
- `--project`: Target GCP project ID.

---

### 4. `config-connector restore` / `config-connector backup restore`

Restores Config Connector resources from a backup snapshot:

```bash
# Dry-run validation from GCS backup
config-connector restore \
    --source-bucket <GCS_BUCKET_NAME> \
    --backup-timestamp latest \
    --dry-run

# In-Place Restore of a specific namespace
config-connector restore \
    --source-bucket <GCS_BUCKET_NAME> \
    --backup-timestamp latest \
    --filter-namespace payment-services

# Cross-Region DR Failover with auto-namespaces and regional remapping
config-connector restore \
    --source-bucket <GCS_BUCKET_NAME> \
    --backup-timestamp latest \
    --filter-namespace payment-services \
    --auto-create-namespaces \
    --skip-missing-crds \
    --target-region us-east1 \
    --region-mapping "us-central1=us-east1"
```

**Flags**:
- `--source-bucket`: GCS bucket containing the backup snapshot.
- `--from-dir`: Local directory containing backup snapshots.
- `--backup-timestamp`: Specific timestamp folder (e.g. `2026_09_06_15_25_48`) or `latest` (default: `latest`).
- `--cluster`: Source cluster name used when storing the backup (defaults to current kubeconfig context).
- `--dry-run`: Previews the restore execution plan and DAG topological sort without applying changes.
- `--filter-namespace`: Limits the restore to resources belonging to a specific namespace.
- `--auto-create-namespaces`: Automatically creates namespaces in the target cluster if they do not exist.
- `--skip-missing-crds`: Skips resources whose CRDs are not installed in the target cluster.
- `--target-region`: Dynamically overrides regional fields (`spec.location`, `spec.region`) on all regional resources.
- `--region-mapping`: Explicit source-to-target regional mappings in `source=target` format (e.g., `us-central1=us-east1`).
- `--target-project`: Overrides the target GCP project across all restored resources.

---

## Disaster Recovery Runbooks

### Runbook A: In-Place Disaster Recovery (Cluster Re-Acquisition)

**Scenario**: The primary GKE cluster experienced control plane corruption, accidental namespace deletion, or unrecoverable etcd loss. Underlying Google Cloud resources (databases, buckets, secrets, IAM) are still running in GCP.

**Objective**: Rehydrate the Kubernetes control plane and re-acquire management of cloud infrastructure without causing cloud resource recreation, data loss, or service interruption.

1. **Verify Target Cluster & Workload Identity**:
   ```bash
   kubectl cluster-info
   kubectl get pods -n cnrm-system
   ```
2. **Execute Dry-Run Preview**:
   ```bash
   config-connector restore \
       --source-bucket kcc-backup-vault \
       --backup-timestamp latest \
       --filter-namespace prod-workloads \
       --dry-run
   ```
3. **Execute Restore**:
   ```bash
   config-connector restore \
       --source-bucket kcc-backup-vault \
       --backup-timestamp latest \
       --filter-namespace prod-workloads \
       --auto-create-namespaces
   ```
4. **Validation**:
   - Restored resources will be annotated with `deletion-policy: abandon` and `management-conflict-prevention-policy: none`.
   - The Config Connector controller will adopt the existing GCP resources and transition them to `Ready: True (UpToDate)` within seconds.

---

### Runbook B: Cross-Region Multi-Cluster Failover (RTO < 4h)

**Scenario**: Complete regional outage affecting the primary region (e.g., `us-central1`). Business continuity requires standing up or activating standby infrastructure in a secondary region (e.g., `us-east1-c`).

**Objective**: Provision all infrastructure definitions in the secondary region with updated regional parameters, satisfying RTO < 4h requirements.

1. **Switch Kubeconfig to DR Cluster**:
   ```bash
   kubectl config use-context gke_prod-project_us-east1-c_prod-dr-cluster
   ```
2. **Validate Target Region & CRDs**:
   Ensure Config Connector is active on the DR cluster.
3. **Execute Cross-Region Restore**:
   ```bash
   config-connector restore \
       --source-bucket kcc-backup-vault \
       --backup-timestamp latest \
       --auto-create-namespaces \
       --skip-missing-crds \
       --target-region us-east1 \
       --region-mapping "us-central1=us-east1"
   ```
4. **Automated Transformations Handled**:
   - Namespaces missing on the DR cluster are automatically created.
   - Any regional resource configured in `us-central1` (e.g. StorageBucket locations, Cloud SQL regions) is translated to `us-east1`.
   - Missing CRDs on non-identical clusters are safely skipped without failing the pipeline.
   - Resources are applied strictly in topological order (IAM & KMS -> Networking -> Databases -> IAM Bindings).

---

## Live Environment Verification Matrix

The native Backup and DR suite was rigorously validated across 10 real-world scenarios in live Google Cloud environments (`gca-gke-2025` and `gca-gke-test`):

| Test # | Scenario | Services Tested | Result | Verification Notes |
|:---|:---|:---|:---|:---|
| **1** | Multi-Service Live Provisioning | Storage, SecretManager, Pub/Sub, IAM, Cloud SQL | **PASS** | 8 distinct KCC resources achieved `Ready: True` (`UpToDate`) on primary cluster `kcc-management-cluster`. |
| **2** | Local Filesystem Backup | All 8 resources + cluster resources | **PASS** | Backed up 36 resources to local disk directory with sanitized manifests and `summary.json`. |
| **3** | GCS Remote Backup & Status | GCS Bucket storage | **PASS** | Uploaded 36 resources to GCS; verified snapshot status and resource kind breakdown via `backup status`. |
| **4** | Sanitization & Safety Audit | Manifest parser | **PASS** | Verified stripped metadata (`uid`, `resourceVersion`, `managedFields`, `last-applied`) while preserving `spec.resourceID` and project IDs. |
| **5** | Dry-Run Simulation | Local & remote manifests | **PASS** | Simulated full restore with zero mutations, validating DAG dependency order. |
| **6** | In-Place Disaster Simulation | Namespace CR deletion | **PASS** | Deleted all 8 KCC CRs with `deletion-policy: abandon`. Confirmed K8s namespace was empty while underlying GCP resources remained 100% active and healthy. |
| **7** | In-Place Cloud Re-Acquisition | Storage, SecretManager, Pub/Sub, IAM, Cloud SQL | **PASS** | Re-applied backup manifests; Config Connector re-acquired all 8 existing cloud resources to `Ready: True` in <10 seconds without recreation or data loss. |
| **8** | Regional Outage Simulation | Context switch to secondary cluster | **PASS** | Primary region simulated offline; switched context to secondary cluster `prod-api-router-07` in `us-east1-c`. |
| **9** | Cross-Region DR Failover | Multi-service stack | **PASS** | Auto-created target namespace, remapped `us-central1` to `us-east1`, applied 8 resources in topological order, and handled CRD discrepancies gracefully. |
| **10** | Idempotency & Conflict Prevention | Re-run restore pipeline | **PASS** | Re-executed restore on DR cluster; 8/8 resources Server-Side Applied cleanly as no-ops with zero conflict errors. |
