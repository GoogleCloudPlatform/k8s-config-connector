# Standardized Backup and Disaster Recovery (DR) for Config Connector

Config Connector provides a native, declarative Backup and Disaster Recovery (DR) suite within the `config-connector` CLI. This feature suite addresses critical business continuity requirements (such as Recovery Time Objective RTO < 4h and Recovery Point Objective RPO) for enterprise workloads managing Google Cloud infrastructure declaratively.

---

## Overview

The native Backup and DR toolset enables:
- **Automated Continuous Backups**: Declarative scheduling via Kubernetes `CronJob` with dedicated GCP Service Accounts and Workload Identity.
- **Cross-Platform GKE Autopilot & GKE Standard Support**: Native compliance with Autopilot Pod Security Standards (`runAsNonRoot: true`, `seccompProfile: RuntimeDefault`, `capabilities.drop: [ALL]`, explicit CPU/memory requests) as well as GKE Standard clusters.
- **Hybrid Geo-Replication & Turbo Replication**: Dual-region GCS bucket pairs (`--dual-region`) with Turbo Replication (`--turbo-replication`, `rpo: ASYNC_TURBO`, <15m RPO SLA) and parallel dual-writing to an offsite replica bucket (`--replica-bucket`).
- **WORM Immutability & Data Protection**: Object Versioning enabled by default and WORM retention policy locking (`--retention-days`, `--lock-retention`) protecting backups against tampering, ransomware, or premature deletion.
- **Cryptographic SHA-256 Integrity Verification**: Generates per-manifest SHA-256 digests in `summary.json` and verifies integrity on restore (`--verify-integrity`), failing closed if any manifest is modified.
- **Automated Outage Failover**: Seamless failover to `--fallback-bucket` when the primary GCS bucket is unreachable during a regional cloud disaster.
- **Unified 360° Backup (Backup for GKE Integration)**: Optional holistic integration (`--include-cluster-backup`) unifying in-cluster workloads and PersistentVolume CSI snapshots via Backup for GKE (`gkebackup.googleapis.com`) with declarative cloud infrastructure.
- **On-Demand Snapshots**: Immediate snapshots targeting primary and replica GCS buckets or local directories.
- **Resilient Re-Acquisition (In-Place DR)**: Restoring K8s manifests without recreating or modifying underlying GCP cloud resources, eliminating data loss and downtime.
- **Cross-Region DR Failover**: Deep regional parameter remapping across StorageBuckets, Cloud SQL, Pub/Sub allowed persistence regions, and Secret Manager replicas.
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

Provisions declarative infrastructure for scheduled automated backups with least-privilege IAM and WORM compliance:

```bash
# Standard GKE configuration
config-connector backup configure \
    --project <PROJECT_ID> \
    --bucket <GCS_BUCKET_NAME> \
    --location <GCP_REGION> \
    --schedule "0 2 * * *" \
    --retention-days 30 \
    --lock-retention

# GKE Autopilot configuration with Dual-Region Turbo Replication & Secondary Replica
config-connector backup configure \
    --project <PROJECT_ID> \
    --cluster-project <CLUSTER_PROJECT_ID> \
    --cluster <CLUSTER_NAME> \
    --bucket <PRIMARY_BUCKET_NAME> \
    --dual-region nam4 \
    --turbo-replication \
    --retention-days 30 \
    --replica-bucket <REPLICA_BUCKET_NAME> \
    --replica-location us-east1 \
    --autopilot \
    --include-cluster-backup \
    --dry-run
```

**Flags**:
- `--project` *(required)*: GCP project ID where backup resources and storage buckets reside.
- `--cluster-project`: GCP project ID where the GKE cluster resides (for Workload Identity). Defaults to `--project`.
- `--bucket` *(required)*: Primary GCS bucket name to store backups.
- `--bucket-location`: GCP region for the storage bucket (default: `us-central1`).
- `--schedule`: Cron schedule expression or alias (`daily`, `weekly`, `hourly`, default: `daily`).
- `--autopilot`: Generates Autopilot-compliant Pod Security Standard contexts (`runAsNonRoot: true`, `seccompProfile: RuntimeDefault`, `capabilities.drop: [ALL]`, explicit CPU/memory requests).
- `--dual-region`: Dual-region bucket pair (e.g. `nam4`, `eur4`, `asia1`).
- `--turbo-replication`: Enables GCS Turbo Replication (`rpo: ASYNC_TURBO`, <15m RPO SLA) for dual-region pairs.
- `--versioning`: Enables GCS Object Versioning (default: `true`).
- `--retention-days`: Retention period in days for WORM compliance.
- `--lock-retention`: Permanently locks bucket retention policy (WORM compliance).
- `--replica-bucket`: Secondary offsite GCS bucket for cross-region disaster recovery dual-writing.
- `--replica-location`: Location for secondary replica bucket (default: `us-east1`).
- `--include-cluster-backup`: Configures Backup for GKE (`gkebackup.googleapis.com`) to back up cluster workloads and persistent volumes.
- `--gke-backup-plan`: Name of GKE BackupPlan to configure (defaults to `<cluster>-backup-plan`).
- `--dry-run`: Outputs generated Kubernetes manifests without applying them to the cluster.

---

### 2. `config-connector backup create`

Executes an on-demand, cryptographically verified backup with concurrent dual-writing:

```bash
# Backup to Primary and Replica GCS Buckets with Cluster Backup
config-connector backup create \
    --bucket <PRIMARY_GCS_BUCKET> \
    --replica-bucket <REPLICA_GCS_BUCKET> \
    --project <PROJECT_ID> \
    --cluster-project <CLUSTER_PROJECT_ID> \
    --cluster <CLUSTER_NAME> \
    --include-cluster-backup

# Backup to Local Filesystem
config-connector backup create \
    --output-dir /var/backups/kcc \
    --cluster <CLUSTER_NAME>
```

**Flags**:
- `--bucket`: Primary GCS bucket name for backup storage.
- `--replica-bucket`: Secondary offsite GCS bucket for concurrent cross-region dual-writing.
- `--replica-project`: GCP project ID for replica bucket (defaults to `--project`).
- `--output-dir`: Target local directory for backup storage. Path traversal sequences (`..`) are strictly rejected.
- `--cluster`: Cluster identifier (defaults to current kubeconfig context name).
- `--project`: Target GCP project ID for cluster identification.
- `--cluster-project`: GCP project ID where the cluster resides.
- `--include-cluster-backup`: Triggers Backup for GKE (`gkebackup.googleapis.com`) to snapshot cluster workloads and CSI persistent volumes. Gracefully logs warnings if agent is missing without failing cloud infrastructure backups.
- `--gke-backup-plan`: GKE BackupPlan identifier (defaults to `<cluster>-backup-plan`).

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

Restores Config Connector resources from a backup snapshot with cryptographic integrity verification, automatic failover, and deep regional translation:

```bash
# Dry-run validation from GCS backup with automated failover
config-connector restore \
    --source-bucket <PRIMARY_BUCKET_NAME> \
    --fallback-bucket <REPLICA_BUCKET_NAME> \
    --backup-timestamp latest \
    --dry-run

# Cross-Platform / Cross-Region DR Failover (Autopilot -> Standard)
config-connector restore \
    --source-bucket <REPLICA_BUCKET_NAME> \
    --backup-timestamp latest \
    --filter-namespace payment-services \
    --auto-create-namespaces \
    --skip-missing-crds \
    --target-region us-east1 \
    --region-mapping "us-central1=us-east1" \
    --verify-integrity \
    --include-cluster-backup
```

**Flags**:
- `--source-bucket`: GCS bucket containing the backup snapshot.
- `--fallback-bucket`: Automatic fallback GCS replica bucket if the source bucket is unreachable during a regional cloud disaster.
- `--from-dir`: Local directory containing backup snapshots. Path traversal sequences (`..`) are strictly rejected.
- `--backup-timestamp`: Specific timestamp folder (e.g. `2026_09_06_15_25_48`) or `latest` (default: `latest`).
- `--cluster`: Source cluster name used when storing the backup (defaults to current kubeconfig context).
- `--dry-run`: Previews the restore execution plan and DAG topological sort without applying changes.
- `--filter-namespace`: Limits the restore to resources belonging to a specific namespace.
- `--auto-create-namespaces`: Automatically creates namespaces in the target cluster if they do not exist (default: `true`).
- `--skip-missing-crds`: Gracefully skips resources whose CRDs are not installed in the target cluster.
- `--verify-integrity`: Enforces cryptographic SHA-256 digest validation against `summary.json`, failing closed if any manifest has been altered or corrupted (default: `true`).
- `--target-region`: Dynamically overrides regional fields (`spec.location`, `spec.region`, Pub/Sub `allowedPersistenceRegions`, Secret Manager `replicas[].location`).
- `--region-mapping`: Explicit source-to-target regional mappings in `source=target` format (e.g., `us-central1=us-east1`).
- `--target-project`: Overrides the target GCP project across all restored resources.
- `--deletion-policy`: Sets resource deletion policy annotation (`abandon` or `delete`, default: `abandon`).
- `--management-conflict-policy`: Sets acquisition conflict policy (`none`, `resource`, or `priority`, default: `none`).
- `--include-cluster-backup`: Correlates with and presents restoration instructions for in-cluster workloads and PersistentVolumes via Backup for GKE.

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

The native Backup and DR suite was rigorously validated across 12 real-world scenarios in live Google Cloud environments (`gca-gke-2025` and `gca-gke-test`):

| Test # | Scenario | Services Tested | Clusters Tested | Result | Verification Notes |
|:---|:---|:---|:---|:---|:---|
| **1** | GKE Autopilot Admission & Security Compliance | StorageBucket, IAMPolicyMember, CronJob | GKE Autopilot (`kcc-management-cluster`) | **PASS** | Evaluated server dry-run admission on live Autopilot cluster. Verified strict compliance with `runAsNonRoot: true`, `seccompProfile: RuntimeDefault`, `capabilities.drop: [ALL]`, and explicit CPU/memory requests. |
| **2** | GKE Standard Cluster Configuration | StorageBucket, IAMPolicyMember, CronJob | GKE Standard (`prod-api-router-07`) | **PASS** | Validated server-side apply on standard GKE cluster with least-privilege IAM (`objectCreator` + `objectViewer`) and automated CronJob scheduling. |
| **3** | Live Multi-Service Provisioning | Storage, Pub/Sub, SecretManager, IAM, Cloud SQL | GKE Autopilot (`kcc-management-cluster`) | **PASS** | Deployed multi-service workload; all 6 distinct KCC resources achieved `Ready: True` (`UpToDate`) on live Autopilot cluster. |
| **4** | Dual-Region & Concurrent Dual-Write Backup | GCS Primary & Replica Buckets | GKE Autopilot | **PASS** | Executed live backup writing concurrently to primary GCS (`us-central1`) and replica GCS (`us-east1`) with 100% manifest consistency and summary metadata. |
| **5** | Cryptographic SHA-256 Digest Verification | Manifest integrity hasher | Local & GCS | **PASS** | Verified that `summary.json` contains SHA-256 hashes for all 34 objects. Tested untampered restore dry-run which validated all 34 hashes successfully. |
| **6** | Security Defense: Path Traversal Attack | Filesystem validator | CLI | **PASS** | Executed `backup create --output-dir=/tmp/../../etc` and `restore --from-dir=/tmp/../../etc`; verified strict rejection with relative traversal defense errors. |
| **7** | Security Defense: Cryptographic Tamper Detection | SHA-256 verification engine | Local filesystem | **PASS** | Injected 1 altered byte into a backup manifest. Verified restore failed closed with `cryptographic checksum mismatch ... (tamper alert)`. |
| **8** | Outage Simulation: Automatic Bucket Failover | GCS Storage client | Live GCP | **PASS** | Pointed `--source-bucket` to an offline/non-existent bucket with `--fallback-bucket` replica. CLI detected failure, automatically failed over to replica, and completed restore cleanly. |
| **9** | Cross-Platform Regional DR: Autopilot -> Standard | Storage, Pub/Sub, SecretManager, IAM, Cloud SQL | GKE Autopilot -> GKE Standard | **PASS** | Restored backup taken from GKE Autopilot (`us-central1`) onto GKE Standard (`prod-api-router-07` in `us-east1-c`). Auto-created missing namespace and remapped all regional parameters. |
| **10** | Deep Regional Translation Verification | Pub/Sub, SecretManager, Storage | GKE Standard | **PASS** | Verified deep translation: remapped Pub/Sub `allowedPersistenceRegions` to `us-east1` and Secret Manager `replicas[].location` to `us-east1`. |
| **11** | Backup for GKE Unified Integration | `gkebackup.googleapis.com` | GKE Autopilot | **PASS** | Tested `--include-cluster-backup`; verified graceful degradation when GKE BackupPlan was missing (logged warning, recorded in `summary.json`, completed cloud backup). |
| **12** | Database & Asset Invariant Safety Verification | Cloud SQL (`baseline-payment-db`) | Live GCP (`gca-gke-2025`) | **PASS** | Confirmed pre-existing database instance `baseline-payment-db` remained 100% untouched, healthy, and `RUNNABLE` throughout all tests. |
