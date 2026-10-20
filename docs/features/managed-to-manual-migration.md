# Migrating Config Connector: Managed to Customer-Managed (Standalone) Installation

Enterprises operating Config Connector (KCC) via Google-managed installation pathways—**GKE Addon** or **Config Controller (ACP)**—often reach operational maturity where transitioning to **Customer-Managed Standalone Installation** (on GKE Standard or Autopilot) becomes critical. 

This guide outlines the architectural invariants, automated tooling (`kcc-migrate`), and native zero-binary runbooks required to perform zero-downtime migrations without triggering accidental cloud asset deletion, management conflict lockouts, or GitOps webhook rejections.

---

## 1. Migration Pathways Overview

Config Connector supports three primary installation models:
1. **Standalone Installation (Customer-Managed on GKE Standard or Autopilot):** Direct reconcilers, rapid release velocity, independent cluster lifecycle management, and multi-region disaster recovery topologies.
2. **GKE Addon (Cluster-Managed):** Turnkey activation, but version upgrades are tied to GKE control plane release trains.
3. **Config Controller (Google-Managed Service Wrapper):** Managed control plane wrapper with single-region constraints and multi-month release train version lags.

This guide details both migration pathways:

| Workflow | Source Topology | Target Topology | Key Operational Consideration |
| :--- | :--- | :--- | :--- |
| **Workflow 1: GKE Addon to Standalone** | GKE Cluster with Addon enabled | Same GKE Cluster with Standalone Operator | Non-destructive operator replacement; avoiding finalizer cascades during addon de-registration. |
| **Workflow 2: Config Controller to Standalone** | Config Controller Instance (ACP) | Customer-Managed GKE Cluster | Cross-cluster migration; lease handover; topological DAG dependency ordering; state adoption. |

---

## 2. Core Safety Invariants

A naive export-and-apply migration or uncoordinated teardown causes catastrophic production outages. Both migration pathways enforce five mandatory safety invariants:

| Safety Invariant | Operational Risk | Required Safety Directive | Enforcement Mechanism |
| :--- | :--- | :--- | :--- |
| **Deletion Policy** | Tearing down or disabling the source environment triggers KCC finalizers, issuing destructive HTTP `DELETE` calls to live GCP APIs. | `cnrm.cloud.google.com/deletion-policy: abandon` | Stamped on all resource manifests before any deletion or teardown operations. |
| **Conflict Policy** | Source cluster maintains `cnrm-lease-holder-id` and `cnrm-lease-expiration` labels on GCP resources, blocking target cluster reconcilers for hours. | `cnrm.cloud.google.com/management-conflict-prevention-policy: none` | Stamped on manifests applied to target cluster to allow immediate adoption. |
| **GitOps 3-Way Merge** | GitOps tools (ArgoCD, Config Sync) trigger KCC's `deny-immutable-field-updates` admission webhook if `spec.resourceID` or `state-into-spec` are omitted. | `cnrm.cloud.google.com/state-into-spec: absent` | Stamped during sanitization and synchronized into GitOps repositories. |
| **Teardown Sequence** | Deleting namespace `ConfigConnectorContext` terminates `cnrm-controller-manager` pods, stranding tenant CRs in `Terminating` deadlock. | Strict execution order | Workload CR finalizers are stripped *before* operator or context components are removed. |
| **Topological DAG Ordering** | Applying dependent resources out of order (e.g. `SQLDatabase` before `SQLInstance`, or IAM bindings before SA) causes reference resolution errors and reconcile loops. | Topological DAG Sorting (Tiers 0–5) | Managed automatically via `kcc-migrate adopt --sort=true`. |

---

## 3. Workflow 1: GKE Addon to Standalone Installation (Same-Cluster)

When transitioning from the GKE Addon to Standalone KCC on the same cluster, the goal is to decouple existing KCC custom resources from the cluster-managed addon controllers before disabling the addon via `gcloud`.

### Step 1: Backup All KCC Custom Resources

Export all live KCC resources across all namespaces to prevent data loss:

```bash
# Discover all active Config Connector CRDs
CRDS=$(kubectl get crds -o jsonpath='{.items[?(@.spec.group=~".*\\.cnrm\\.cloud\\.google\\.com")].metadata.name}')

mkdir -p ./kcc-addon-backup

for crd in ${CRDS}; do
  echo "Backing up ${crd}..."
  kubectl get "${crd}" --all-namespaces -o yaml > "./kcc-addon-backup/${crd}.yaml"
done
```

### Step 2: Inject `deletion-policy: abandon`

Prevent KCC from issuing GCP `DELETE` calls when the addon operator is removed:

```bash
for crd in ${CRDS}; do
  kubectl get "${crd}" --all-namespaces -o json | jq -r '.items[] | "\(.metadata.namespace) \(.metadata.name)"' | while read -r ns name; do
    if [ -n "${name}" ]; then
      if [ "${ns}" = "null" ] || [ -z "${ns}" ]; then
        kubectl annotate "${crd}" "${name}" "cnrm.cloud.google.com/deletion-policy=abandon" --overwrite
      else
        kubectl annotate "${crd}" "${name}" -n "${ns}" "cnrm.cloud.google.com/deletion-policy=abandon" --overwrite
      fi
    fi
  done
done
```

### Step 3: Strip Finalizers from Custom Resources

Remove `cnrm.cloud.google.com/finalizer` from all workload resources so that if CRDs are temporarily removed during addon disablement, Kubernetes will not deadlock:

```bash
for crd in ${CRDS}; do
  kubectl get "${crd}" --all-namespaces -o json | jq -r '.items[] | "\(.metadata.namespace) \(.metadata.name)"' | while read -r ns name; do
    if [ -n "${name}" ]; then
      if [ "${ns}" = "null" ] || [ -z "${ns}" ]; then
        kubectl patch "${crd}" "${name}" --type=merge -p '{"metadata":{"finalizers":[]}}'
      else
        kubectl patch "${crd}" "${name}" -n "${ns}" --type=merge -p '{"metadata":{"finalizers":[]}}'
      fi
    fi
  done
done
```

### Step 4: Disable the GKE Addon

Disable the addon on the GKE cluster:

```bash
gcloud container clusters update "${CLUSTER_NAME}" \
    --region="${REGION}" \
    --update-addons ConfigConnector=DISABLED
```

### Step 5: Install Standalone Config Connector Operator

Install the standalone Config Connector operator:

```bash
# Apply official standalone operator release bundle
kubectl apply -f "https://storage.googleapis.com/config-connector-operator-release/latest/release-bundle.tar.gz"
```

Configure your cluster or namespaced `ConfigConnector` custom resource:

```yaml
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
```

Once the standalone operator reconciles, re-apply the backed-up resources (sanitized of old status and runtime metadata). Standalone KCC will adopt the existing GCP resources cleanly.

---

## 4. Workflow 2: Config Controller (ACP) to Standalone KCC (Cross-Cluster)

Workflow 2 migrates resources across clusters: from a managed Config Controller instance to a target Standalone GKE cluster.

### Approach A: Automated Migration via `kcc-migrate` CLI

The `kcc-migrate` CLI tool (located under `dev/tools/kcc-migrate/`) provides an end-to-end, enterprise-grade automated pipeline.

#### CLI Architecture and Phases

1. **`export`**: Concurrently scans active KCC CRDs and exports live KRM resources using a 20-worker pool (~12s full scan).
2. **`sanitize`**: Strips runtime-managed metadata (`uid`, `resourceVersion`, `generation`, `managedFields`, `status`, source `finalizers`) and injects mandatory safety annotations (`deletion-policy: abandon`, `management-conflict-prevention-policy: none`, `state-into-spec: absent`).
3. **`lint`**: Statically audits sanitized manifests against schema and invariant rules, failing fast on invalid configurations.
4. **`adopt`**: Topologically orders manifests across 6 dependency tiers (DAG sorting) and executes server-side dry-run validation followed by live adoption into the target Standalone cluster.
5. **`detach`**: Audits and cleanly removes finalizers from the source Config Controller instance, guaranteeing 0 cloud asset deletions.
6. **`rollback`**: Provides two safety rollback modes:
   - *Pre-detach abort*: Purges partially applied manifests from target cluster while leaving source untouched.
   - *Post-detach repatriation*: Emergency repatriator re-attaching source Config Controller while abandoning target cluster.

#### Quick Start Execution

```bash
# Step 1: Export resources from source Config Controller
kcc-migrate export \
    --context=source-config-controller-ctx \
    --namespaces=tenant-prod,tenant-shared \
    --output-dir=./raw-manifests

# Step 2: Sanitize manifests and inject safety invariants
kcc-migrate sanitize \
    --input-dir=./raw-manifests \
    --output-dir=./sanitized-manifests

# Step 3: Run static validation and invariant checks
kcc-migrate lint \
    --input-dir=./sanitized-manifests \
    --strict=true

# Step 4: Dry-run and adopt into target Standalone KCC cluster with DAG sorting
kcc-migrate adopt \
    --context=target-standalone-kcc-ctx \
    --input-dir=./sanitized-manifests \
    --sort=true \
    --dry-run=true

kcc-migrate adopt \
    --context=target-standalone-kcc-ctx \
    --input-dir=./sanitized-manifests \
    --sort=true \
    --dry-run=false

# Step 5: Verify target adoption status
kubectl --context=target-standalone-kcc-ctx get gcp -A

# Step 6: Non-destructively detach from source Config Controller
kcc-migrate detach \
    --context=source-config-controller-ctx \
    --namespaces=tenant-prod,tenant-shared
```

#### Topological Dependency Tiers

When `--sort=true` is specified during `adopt`, resources are applied in deterministic order to ensure parent references resolve immediately:

- **Tier 0 (Root Foundations & Identities):** `KMSKeyRing`, `ComputeNetwork`, `IAMServiceAccount`, `StorageBucket`.
- **Tier 1 (Sub-foundations & Subnets):** `KMSCryptoKey`, `ComputeSubnetwork`, `SecretManagerSecret`, `PubSubTopic`.
- **Tier 2 (Core Stateful Engines & Firewalls):** `SQLInstance`, `SpannerInstance`, `RedisInstance`, `ComputeFirewall`.
- **Tier 3 (Stateful Children & Databases):** `SQLDatabase`, `SQLUser`, `DNSRecordSet`, `PubSubSubscription`.
- **Tier 4 (Routing & Endpoints):** `ComputeBackendService`, `ComputeForwardingRule`.
- **Tier 5 (IAM Policy Members & Bindings):** `IAMPolicyMember`, `IAMPolicy`, `IAMPartialPolicy`.

---

### Approach B: Native Tooling Fallback (Zero-Binary Runbook)

For air-gapped, high-compliance environments where third-party CLI binaries cannot be deployed to jumpboxes, execute the migration using standard `kubectl`, `jq`, and `yq`.

#### Phase 1: Freeze GitOps Synchronization
Suspend automated synchronization (e.g. RootSync / RepoSync / ArgoCD Application) on the source Config Controller to prevent reconciler fighting:

```bash
kubectl --context=source-ctx patch rootsync root-sync -n config-management-system --type=merge -p '{"spec":{"pause":true}}' 2>/dev/null || true
```

#### Phase 2: Parallel Resource Export
Export all KCC custom resources:

```bash
mkdir -p ./raw
CRD_LIST=$(kubectl --context=source-ctx get crds -o jsonpath='{.items[?(@.spec.group=~".*\\.cnrm\\.cloud\\.google\\.com")].metadata.name}')

for crd in ${CRD_LIST}; do
  kubectl --context=source-ctx get "${crd}" -A -o yaml > "./raw/${crd}.yaml"
done
```

#### Phase 3: Manifest Sanitization via `yq`
Process each exported file to strip runtime fields and stamp safety annotations:

```bash
mkdir -p ./sanitized

for f in ./raw/*.yaml; do
  [ -e "$f" ] || continue
  yq eval '
    del(.items[].metadata.uid) |
    del(.items[].metadata.resourceVersion) |
    del(.items[].metadata.generation) |
    del(.items[].metadata.creationTimestamp) |
    del(.items[].metadata.managedFields) |
    del(.items[].metadata.finalizers) |
    del(.items[].status) |
    .items[].metadata.annotations["cnrm.cloud.google.com/deletion-policy"] = "abandon" |
    .items[].metadata.annotations["cnrm.cloud.google.com/management-conflict-prevention-policy"] = "none" |
    .items[].metadata.annotations["cnrm.cloud.google.com/state-into-spec"] = "absent"
  ' "$f" > "./sanitized/$(basename "$f")"
done
```

#### Phase 4: Server-Side Dry-Run & Target Adoption
Apply the sanitized manifests to the target cluster using Server-Side Apply:

```bash
# Server-side dry run
kubectl --context=target-ctx apply --server-side --dry-run=server -R -f ./sanitized/

# Live adoption
kubectl --context=target-ctx apply --server-side -R -f ./sanitized/
```

#### Phase 5: Verification
Confirm that all resources have reconciled to `Ready: True`:

```bash
kubectl --context=target-ctx get gcp -A -o wide
```

#### Phase 6: Source Detachment
Safely strip finalizers on the source cluster:

```bash
for crd in ${CRD_LIST}; do
  kubectl --context=source-ctx get "${crd}" -A -o json | jq -r '.items[] | "\(.metadata.namespace) \(.metadata.name)"' | while read -r ns name; do
    if [ -n "${name}" ]; then
      kubectl --context=source-ctx annotate "${crd}" "${name}" -n "${ns}" "cnrm.cloud.google.com/deletion-policy=abandon" --overwrite
      kubectl --context=source-ctx patch "${crd}" "${name}" -n "${ns}" --type=merge -p '{"metadata":{"finalizers":[]}}'
    fi
  done
done
```

---

## 5. Multi-Region Disaster Recovery (Active-Passive Failover)

Following a successful migration to Standalone KCC, organizations requiring multi-region high availability can establish an active-passive dual-cluster topology:

- **Primary Cluster (Region A):** Active reconciler driving cloud infrastructure state.
- **Secondary Cluster (Region B):** Passive standby cluster ready to assume reconciliation within 30 seconds.
- **Distributed Leader Arbitration:** Utilizes a `MultiClusterLease` daemon backed by Google Cloud Storage generation locks, ensuring exactly one active reconciling cluster and eliminating split-brain risks.

Refer to the [Cloud SQL Enterprise DR Guide](cloudsql-enterprise-dr.md) for database-specific replication and automated failover patterns.
