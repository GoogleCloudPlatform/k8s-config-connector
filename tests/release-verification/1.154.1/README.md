# KCC Release 1.154 Verification Test Suite

This directory contains automated scripts and Kubernetes manifests to verify the Config Connector (KCC) release `1.154.1` through an **upgrade scenario** (1.153.0 -> 1.154.1).

Per instructions, this suite ignores **New Alpha Resources** and **Reconciliation Improvements** and focuses exclusively on:
1. **Bug Fixes**
2. **New Fields**

---

## Release 1.154 Verification Matrix

### 1. New Fields Focus
| Resource | Target Field / Feature | Verification File |
| :--- | :--- | :--- |
| `StorageBucket` | `spec.ipFilter` | `manifests/post-upgrade/storagebucket-ipfilter.yaml` |
| `PubSubTopic` | `spec.messageStoragePolicy.enforceInTransit` | `manifests/post-upgrade/pubsubtopic-enforceintransit.yaml` |
| `ComputeSubnetwork` | `spec.reservedInternalRange` | `manifests/post-upgrade/computesubnetwork-reservedrange.yaml` |
| `ComputeSecurityPolicy` | `spec.region` | `manifests/post-upgrade/computesecuritypolicy-region.yaml` |
| `DNSRecordSet` | `spec.routingPolicy.geo[].healthCheckRef` & `rrdatasRefs` | `manifests/post-upgrade/dnsrecordset-routingpolicy.yaml` |
| `ComputeAddress` | `spec.ipCollection` | `manifests/post-upgrade/computeaddress-ipcollection.yaml` |
| `MonitoringAlertPolicy` | `spec.conditions[].conditionSql` | `manifests/post-upgrade/monitoringalertpolicy-conditionsql.yaml` |
| `ComputeRouterNAT` | `spec.type: PRIVATE` (Private NAT) | `manifests/post-upgrade/computerouternat-privatenat.yaml` |

### 2. Bug Fixes Focus
| Resource / Area | Fix Description | Verification File |
| :--- | :--- | :--- |
| `ComposerEnvironment` | Fix `storageConfig.bucketRef` mapping (PR #11001) | `manifests/pre-upgrade/composer-environment-bugfix.yaml` |
| `MemorystoreInstance` | Prevent infinite drift & false update loops (PR #11547, #11559) | `manifests/pre-upgrade/memorystore-instance-bugfix.yaml` |
| `BigQueryTable` | Fix perpetual diff on tables inheriting dataset encryption (PR #9623) | `manifests/pre-upgrade/bigquery-table-bugfix.yaml` |
| `NotebooksInstance` | Fix direct controller for NotebookInstance to resolve references (PR #9810) | `manifests/pre-upgrade/notebooks-instance-bugfix.yaml` |
| `KMSAutokeyConfig` | Autokey config identity & deletion resolution (PR #9658) | `manifests/pre-upgrade/kms-autokeyconfig-bugfix.yaml` |

---

## Step-by-Step Execution Guide

### Prerequisites
- `gcloud` CLI installed and authenticated with GCP credentials (`gcloud auth login`).
- `kubectl` CLI installed.
- Access to a GCP project with sufficient IAM permissions to create GKE clusters and GCP resources.

### Execution Order

#### Step 0: Set Environment Variables & Create GKE Cluster
```bash
export PROJECT_ID="your-gcp-project-id"
export CLUSTER_NAME="kcc-release-1154-verify"
export ZONE="us-central1-a"

./00_setup_cluster_and_sa.sh
```

#### Step 1: Install Previous KCC Version (1.153.0)
Installs Config Connector Operator version `1.153.0` and configures `ConfigConnector` custom resource in cluster mode.
```bash
./01_install_old_version.sh
```

#### Step 2: Deploy Pre-Upgrade Resources
Deploys base resources and resources targeting pre-existing bug fixes on version 1.153.0.
```bash
./02_apply_pre_upgrade_resources.sh
```

#### Step 3: Upgrade KCC to Version 1.154.0
Upgrades the Config Connector Operator and controller components to `1.154.0` (`latest`).
```bash
./03_upgrade_to_new_version.sh
```

#### Step 4: Apply Post-Upgrade Resources & Verify New Fields / Bug Fixes
Applies updated manifests with 1.154 new fields and checks resource conditions for readiness and stability.
```bash
./04_apply_post_upgrade_verification.sh
```

#### Step 4b: Check for Drift Loops & False Update Loops (Automated)
Monitors `resourceVersion` stability across a time interval and checks controller manager logs.
```bash
./check_drift_loop.sh
```

#### Step 5: Cleanup Resources
Deletes test resources and optionally tears down the cluster.
```bash
./05_cleanup.sh
```
