# KCC Multi-Cluster Leader Election & Syncer Integration: Final Deployment & Operations Guide

This document is the definitive guide for setting up, verifying, and operating the active-passive Config Connector (KCC) architecture using Multi-Cluster Leader Election (MCL) and KRMSyncer. 

This guide focuses on the **Namespaced Mode** installation, which is the recommended and most popular deployment model for KCC. In this architecture, each KCC Manager creates and manages a strictly isolated 1:1 `KRMSyncer` mapped to its designated tenant namespace, entirely eliminating API conflict "thundering herd" issues during failovers.

---

## ⚡ Quick Start: Token Replacement

This guide uses placeholders (e.g., `<CLUSTER_1_CONTEXT>`) to remain environment-agnostic. Before following the steps, run the following command to replace all tokens with your specific environment values:

```bash
# 1. Fill in your environment values here:
C1_CTX="my-cluster-1-context"
C2_CTX="my-cluster-2-context"
BUCKET="my-mcl-lease-bucket"
KCC_GSA="cnrm-system@my-project.iam.gserviceaccount.com"
MCL_GSA="mcl-controller@my-project.iam.gserviceaccount.com"
C1_ID="cluster-1"
C2_ID="cluster-2"
TENANT_NS="tenant-a"
LEASE_NAME="kcc-managers"

# 2. Run this sed command to update the guide in-place:
sed -i -e "s/<CLUSTER_1_CONTEXT>/$C1_CTX/g" \
       -e "s/<CLUSTER_2_CONTEXT>/$C2_CTX/g" \
       -e "s/<GCS_BUCKET_NAME>/$BUCKET/g" \
       -e "s/<KCC_GSA_EMAIL>/$KCC_GSA/g" \
       -e "s/<MCL_GSA_EMAIL>/$MCL_GSA/g" \
       -e "s/<CLUSTER_1_ID>/$C1_ID/g" \
       -e "s/<CLUSTER_2_ID>/$C2_ID/g" \
       -e "s/<TENANT_NS>/$TENANT_NS/g" \
       -e "s/<LEASE_NAME>/$LEASE_NAME/g" \
       guide_final.md
```

---

## 1. Deployment Guide

### Prerequisites
1. Two Kubernetes clusters accessed via kubectl contexts: `<CLUSTER_1_CONTEXT>` and `<CLUSTER_2_CONTEXT>`.
2. A Google Cloud Storage (GCS) bucket to act as the global lock state: `gs://<GCS_BUCKET_NAME>`.
3. A Google Service Account (GSA) for the MCL controller with `roles/storage.admin` access to the GCS bucket: `<MCL_GSA_EMAIL>`.
4. A GSA for Config Connector to manage GCP resources: `<KCC_GSA_EMAIL>`.
5. Workload Identity enabled on both clusters.

### Step 1: Install Core CRDs
Install the necessary Custom Resource Definitions on both clusters:
```bash
# KCC CRDs
kubectl --context=<CLUSTER_1_CONTEXT> apply -f config/crds/resources/
kubectl --context=<CLUSTER_2_CONTEXT> apply -f config/crds/resources/

# Multi-Cluster Leader Election (MCL) CRDs
kubectl --context=<CLUSTER_1_CONTEXT> apply -f ../multicluster-leader-election/config/crd/bases/
kubectl --context=<CLUSTER_2_CONTEXT> apply -f ../multicluster-leader-election/config/crd/bases/

# KRMSyncer CRDs
kubectl --context=<CLUSTER_1_CONTEXT> apply -f ../kube-etl/syncer/config/crd/bases/
kubectl --context=<CLUSTER_2_CONTEXT> apply -f ../kube-etl/syncer/config/crd/bases/
```

### Step 2: Deploy KCC Operator and Configure Manager
Deploy the Config Connector Operator to both clusters. We will configure KCC for **Namespaced Mode**, which requires a `ConfigConnector` object in `cnrm-system` and a `ConfigConnectorContext` in the tenant namespace.

**Cluster 1 (`<CLUSTER_1_CONTEXT>`):**
```yaml
cat <<EOF | kubectl --context=<CLUSTER_1_CONTEXT> apply -f - 
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
  experiments:
    multiClusterLease:
      leaseName: <LEASE_NAME>
      namespace: cnrm-system
      clusterCandidateIdentity: <CLUSTER_1_ID>
---
apiVersion: v1
kind: Namespace
metadata:
  name: <TENANT_NS>
---
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: <TENANT_NS>
spec:
  googleServiceAccount: <KCC_GSA_EMAIL>
EOF
```

**Cluster 2 (`<CLUSTER_2_CONTEXT>`):**
```yaml
cat <<EOF | kubectl --context=<CLUSTER_2_CONTEXT> apply -f - 
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
  experiments:
    multiClusterLease:
      leaseName: <LEASE_NAME>
      namespace: cnrm-system
      clusterCandidateIdentity: <CLUSTER_2_ID>
---
apiVersion: v1
kind: Namespace
metadata:
  name: <TENANT_NS>
---
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: <TENANT_NS>
spec:
  googleServiceAccount: <KCC_GSA_EMAIL>
EOF
```

### Step 3: Apply Additional RBAC for KCC Manager
In Namespaced Mode, the KCC Manager runs with a dedicated ServiceAccount (`cnrm-controller-manager-<TENANT_NS>`). It needs explicit permissions to interact with the global lease and create its 1:1 mapped Syncer object.

```yaml
cat <<EOF > kcc-mcl-syncer-rbac.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cnrm-syncer-mcl-role
rules:
- apiGroups: ["syncer.gkelabs.io"]
  resources: ["krmsyncers", "krmsyncers/status"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
- apiGroups: ["multicluster.gkelabs.io", "multicluster.core.cnrm.cloud.google.com"]
  resources: ["multiclusterleases"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cnrm-syncer-mcl-rolebinding-<TENANT_NS>
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cnrm-syncer-mcl-role
subjects:
- kind: ServiceAccount
  name: cnrm-controller-manager-<TENANT_NS>
  namespace: cnrm-system
EOF

kubectl --context=<CLUSTER_1_CONTEXT> apply -f kcc-mcl-syncer-rbac.yaml
kubectl --context=<CLUSTER_2_CONTEXT> apply -f kcc-mcl-syncer-rbac.yaml
```

### Step 4: Deploy MCL and Syncer Controllers
Deploy the supporting controllers to both clusters. We use Workload Identity and point the MCL controller to the specific bucket created for the project.

```bash
# 1. Bind the Workload Identity GSA to the MCL controller's KSA
kubectl --context=<CLUSTER_1_CONTEXT> annotate serviceaccount default -n multiclusterlease-system iam.gke.io/gcp-service-account=<MCL_GSA_EMAIL> --overwrite
kubectl --context=<CLUSTER_2_CONTEXT> annotate serviceaccount default -n multiclusterlease-system iam.gke.io/gcp-service-account=<MCL_GSA_EMAIL> --overwrite

# 2. Deploy the controllers using Kustomize 
kubectl --context=<CLUSTER_1_CONTEXT> apply -k ../multicluster-leader-election/config/default
kubectl --context=<CLUSTER_2_CONTEXT> apply -k ../multicluster-leader-election/config/default

kubectl --context=<CLUSTER_1_CONTEXT> apply -k ../kube-etl/syncer/config/default
kubectl --context=<CLUSTER_2_CONTEXT> apply -k ../kube-etl/syncer/config/default

# 3. Patch the deployments to use the correct bucket and remove explicit credential mounts
PATCH_STR='[{"op": "remove", "path": "/spec/template/spec/volumes/0"}, {"op": "remove", "path": "/spec/template/spec/containers/0/volumeMounts/0"}, {"op": "remove", "path": "/spec/template/spec/containers/0/env/0"}, {"op": "replace", "path": "/spec/template/spec/containers/0/args", "value": ["--gcs-bucket=<GCS_BUCKET_NAME>", "--namespace=multiclusterlease-system", "--health-probe-bind-address=:8081", "--metrics-bind-address=127.0.0.1:8080", "--leader-elect"]}]'

kubectl --context=<CLUSTER_1_CONTEXT> patch deployment multiclusterlease-controller-manager -n multiclusterlease-system --type='json' -p="$PATCH_STR"
kubectl --context=<CLUSTER_2_CONTEXT> patch deployment multiclusterlease-controller-manager -n multiclusterlease-system --type='json' -p="$PATCH_STR"

# Restart pods to apply patches
kubectl --context=<CLUSTER_1_CONTEXT> delete pod -n multiclusterlease-system -l control-plane=controller-manager
kubectl --context=<CLUSTER_2_CONTEXT> delete pod -n multiclusterlease-system -l control-plane=controller-manager
```

### Step 5: Setup Cross-Cluster Authentication
Create secrets containing the opposing cluster's `kubeconfig` so the `KRMSyncer` can securely pull data when operating as a follower.
```bash
# Generate raw kubeconfigs
kubectl config view --raw --context=<CLUSTER_1_CONTEXT> > cluster1-kubeconfig
kubectl config view --raw --context=<CLUSTER_2_CONTEXT> > cluster2-kubeconfig

# On Cluster 1 (to allow pulling from Cluster 2)
kubectl --context=<CLUSTER_1_CONTEXT> create secret generic <CLUSTER_2_ID> --namespace cnrm-system --from-file=kubeconfig=cluster2-kubeconfig

# On Cluster 2 (to allow pulling from Cluster 1)
kubectl --context=<CLUSTER_2_CONTEXT> create secret generic <CLUSTER_1_ID> --namespace cnrm-system --from-file=kubeconfig=cluster1-kubeconfig
```

---

## 2. Operational Verification & Validation

To guarantee the architecture is functioning correctly, an operator should run the following validation sequence. This simulates a hard crash of the active leader to verify the failover mechanisms and the automated creation of the replication syncers.

### Step 1: Observe Current Steady State
Check which cluster currently holds the global lease.
```bash
kubectl --context=<CLUSTER_1_CONTEXT> get multiclusterlease <LEASE_NAME> -n cnrm-system -o jsonpath='{.status.globalHolderIdentity}'
```
*(Assume `<CLUSTER_1_ID>` is currently the leader for the following steps)*

### Step 2: Verify Initial KRMSyncer Generation
In Namespaced mode, the KCC Manager dynamically generates a 1:1 `KRMSyncer` object named after the lease and the target namespace (e.g., `<LEASE_NAME>-<TENANT_NS>`).

**On the Leader (`<CLUSTER_1_CONTEXT>`):** The syncer should exist but be suspended, ready to activate if it ever loses leadership.
```bash
kubectl --context=<CLUSTER_1_CONTEXT> get krmsyncer <LEASE_NAME>-<TENANT_NS> -n cnrm-system -o jsonpath='{"Mode: "}{.spec.mode}{"\nSuspend: "}{.spec.suspend}{"\n"}'
# Expected Output:
# Mode: pull
# Suspend: true
```

**On the Follower (`<CLUSTER_2_CONTEXT>`):** The syncer should be actively pulling data from the leader.
```bash
kubectl --context=<CLUSTER_2_CONTEXT> get krmsyncer <LEASE_NAME>-<TENANT_NS> -n cnrm-system -o jsonpath='{"Mode: "}{.spec.mode}{"\nSuspend: "}{.spec.suspend}{"\nPulling From: "}{.spec.remote.clusterConfig.kubeConfigSecretRef.name}{"\n"}'
# Expected Output:
# Mode: pull
# Suspend: false
# Pulling From: <CLUSTER_1_ID>
```

### Step 3: Trigger a Hard Failover
Simulate a catastrophic failure by deleting the KCC manager pod on the active leader (`<CLUSTER_1_CONTEXT>`).
```bash
kubectl --context=<CLUSTER_1_CONTEXT> delete pod -n cnrm-system -l cnrm.cloud.google.com/component=cnrm-controller-manager
```

### Step 4: Validate the Failover Transition
Wait 30 seconds for the original lease to expire and for the follower cluster to claim the GCS lock. Then verify the state has completely flipped.

**1. Verify GCS Lock Transfer:**
```bash
gcloud storage cat gs://<GCS_BUCKET_NAME>/leases/cnrm-system/<LEASE_NAME>
# Expected: "holderIdentity":"<CLUSTER_2_ID>"
```

**2. Verify Local Lease Consensus:**
```bash
kubectl --context=<CLUSTER_1_CONTEXT> get multiclusterlease <LEASE_NAME> -n cnrm-system -o jsonpath='{.status.globalHolderIdentity}'
kubectl --context=<CLUSTER_2_CONTEXT> get multiclusterlease <LEASE_NAME> -n cnrm-system -o jsonpath='{.status.globalHolderIdentity}'
# Both clusters MUST return: <CLUSTER_2_ID>
```

**3. Verify Syncer Role Reversal (The Critical Step):**
```bash
# Cluster 2 (The New Leader) should now be suspended
kubectl --context=<CLUSTER_2_CONTEXT> get krmsyncer <LEASE_NAME>-<TENANT_NS> -n cnrm-system -o jsonpath='{"Suspend: "}{.spec.suspend}{"\n"}'
# Expected: Suspend: true

# Cluster 1 (The New Follower) should wake up its syncer and point it at Cluster 2
kubectl --context=<CLUSTER_1_CONTEXT> get krmsyncer <LEASE_NAME>-<TENANT_NS> -n cnrm-system -o jsonpath='{"Suspend: "}{.spec.suspend}{"\nPulling From: "}{.spec.remote.clusterConfig.kubeConfigSecretRef.name}{"\n"}'
# Expected: 
# Suspend: false
# Pulling From: <CLUSTER_2_ID>
```
If these validations pass, the active-passive namespaced architecture is fully functional and highly available.
