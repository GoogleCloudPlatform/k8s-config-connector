#!/bin/bash
# KCC Multi-Cluster Leader Election & Syncer Integration Demo Script
# This script serves as a map of commands for an operator to deploy and verify the setup.
# Note: It is meant to be run interactively or line-by-line rather than all at once.

# -----------------------------------------------------------------------------
# 0. Set your environment variables here
# -----------------------------------------------------------------------------
C1_CTX="my-cluster-1-context"
C2_CTX="my-cluster-2-context"
BUCKET="my-mcl-lease-bucket"
KCC_GSA="cnrm-system@my-project.iam.gserviceaccount.com"
MCL_GSA="mcl-controller@my-project.iam.gserviceaccount.com"
C1_ID="cluster-1"
C2_ID="cluster-2"
TENANT_NS="tenant-a"
LEASE_NAME="kcc-managers"
SYNCER_NAME="${LEASE_NAME}-${TENANT_NS}"

# -----------------------------------------------------------------------------
# 1. Install Core CRDs
# -----------------------------------------------------------------------------
echo "Installing KCC CRDs..."
kubectl --context=$C1_CTX apply -f config/crds/resources/ 
kubectl --context=$C2_CTX apply -f config/crds/resources/ 

echo "Installing Multi-Cluster Leader Election (MCL) CRDs..."
kubectl --context=$C1_CTX apply -f ../multicluster-leader-election/config/crd/bases/ 
kubectl --context=$C2_CTX apply -f ../multicluster-leader-election/config/crd/bases/ 

echo "Installing KRMSyncer CRDs..."
kubectl --context=$C1_CTX apply -f ../kube-etl/syncer/config/crd/bases/ 
kubectl --context=$C2_CTX apply -f ../kube-etl/syncer/config/crd/bases/ 

# -----------------------------------------------------------------------------
# 2. Deploy KCC Operator and Configure Manager (Namespaced Mode)
# -----------------------------------------------------------------------------
echo "Configuring ConfigConnector (Namespaced Mode) on Cluster 1..."
cat <<EOF | kubectl --context=$C1_CTX apply -f - 
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
  experiments:
    multiClusterLease:
      leaseName: $LEASE_NAME
      namespace: cnrm-system
      clusterCandidateIdentity: $C1_ID
EOF

echo "Configuring ConfigConnector (Namespaced Mode) on Cluster 2..."
cat <<EOF | kubectl --context=$C2_CTX apply -f - 
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
  experiments:
    multiClusterLease:
      leaseName: $LEASE_NAME
      namespace: cnrm-system
      clusterCandidateIdentity: $C2_ID
EOF

echo "Creating tenant namespace and ConfigConnectorContext on both clusters..."
for CTX in $C1_CTX $C2_CTX; do
  kubectl --context=$CTX create namespace $TENANT_NS || true
  cat <<EOF | kubectl --context=$CTX apply -f - 
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: $TENANT_NS
spec:
  googleServiceAccount: $KCC_GSA
EOF
done

# -----------------------------------------------------------------------------
# 3. Apply Additional RBAC for KCC Manager
# -----------------------------------------------------------------------------
echo "Applying Syncer/MCL RBAC for the Namespaced KCC Manager..."
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
  name: cnrm-syncer-mcl-rolebinding-${TENANT_NS}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cnrm-syncer-mcl-role
subjects:
- kind: ServiceAccount
  name: cnrm-controller-manager-${TENANT_NS}
  namespace: cnrm-system
EOF

kubectl --context=$C1_CTX apply -f kcc-mcl-syncer-rbac.yaml
kubectl --context=$C2_CTX apply -f kcc-mcl-syncer-rbac.yaml

# -----------------------------------------------------------------------------
# 4. Deploy MCL and Syncer Controllers
# -----------------------------------------------------------------------------
echo "Binding Workload Identity GSA to MCL controller's KSA..."
kubectl --context=$C1_CTX annotate serviceaccount default -n multiclusterlease-system iam.gke.io/gcp-service-account=$MCL_GSA --overwrite
kubectl --context=$C2_CTX annotate serviceaccount default -n multiclusterlease-system iam.gke.io/gcp-service-account=$MCL_GSA --overwrite

echo "Deploying MCL controllers via Kustomize..."
kubectl --context=$C1_CTX apply -k ../multicluster-leader-election/config/default
kubectl --context=$C2_CTX apply -k ../multicluster-leader-election/config/default

echo "Deploying Syncer controllers..."
kubectl --context=$C1_CTX apply -k ../kube-etl/syncer/config/default
kubectl --context=$C2_CTX apply -k ../kube-etl/syncer/config/default

echo "Patching MCL deployments to use the correct bucket and remove explicit credential mounts..."
PATCH_STR='[{"op": "remove", "path": "/spec/template/spec/volumes/0"}, {"op": "remove", "path": "/spec/template/spec/containers/0/volumeMounts/0"}, {"op": "remove", "path": "/spec/template/spec/containers/0/env/0"}, {"op": "replace", "path": "/spec/template/spec/containers/0/args", "value": ["--gcs-bucket='$BUCKET'", "--namespace=multiclusterlease-system", "--health-probe-bind-address=:8081", "--metrics-bind-address=127.0.0.1:8080", "--leader-elect"]}]'

kubectl --context=$C1_CTX patch deployment multiclusterlease-controller-manager -n multiclusterlease-system --type='json' -p="$PATCH_STR"
kubectl --context=$C2_CTX patch deployment multiclusterlease-controller-manager -n multiclusterlease-system --type='json' -p="$PATCH_STR"

echo "Restarting MCL controller pods to apply patches..."
kubectl --context=$C1_CTX delete pod -n multiclusterlease-system -l control-plane=controller-manager
kubectl --context=$C2_CTX delete pod -n multiclusterlease-system -l control-plane=controller-manager

# -----------------------------------------------------------------------------
# 5. Setup Cross-Cluster Authentication
# -----------------------------------------------------------------------------
echo "Generating raw kubeconfigs..."
kubectl config view --raw --context=$C1_CTX > cluster1-kubeconfig
kubectl config view --raw --context=$C2_CTX > cluster2-kubeconfig

echo "Creating cross-cluster secrets for KRMSyncer..."
# On Cluster 1 (to allow pulling from Cluster 2)
kubectl --context=$C1_CTX create secret generic $C2_ID --namespace cnrm-system --from-file=kubeconfig=cluster2-kubeconfig

# On Cluster 2 (to allow pulling from Cluster 1)
kubectl --context=$C2_CTX create secret generic $C1_ID --namespace cnrm-system --from-file=kubeconfig=cluster1-kubeconfig

# -----------------------------------------------------------------------------
# 6. Operational Verification & Validation
# -----------------------------------------------------------------------------
echo "--- Step 1: Observe Current Steady State ---"
kubectl --context=$C1_CTX get multiclusterlease $LEASE_NAME -n cnrm-system -o jsonpath='{.status.globalHolderIdentity}'
echo ""

echo "--- Step 2: Verify Initial KRMSyncer Generation (Namespaced 1:1 Mapping) ---"
# Check on Cluster 1 (Assuming it is Leader, expected Suspend: true)
kubectl --context=$C1_CTX get krmsyncer $SYNCER_NAME -n cnrm-system -o jsonpath='{"Mode: "}{.spec.mode}{"\nSuspend: "}{.spec.suspend}{"\n"}'

# Check on Cluster 2 (Assuming it is Follower, expected Suspend: false)
kubectl --context=$C2_CTX get krmsyncer $SYNCER_NAME -n cnrm-system -o jsonpath='{"Mode: "}{.spec.mode}{"\nSuspend: "}{.spec.suspend}{"\nPulling From: "}{.spec.remote.clusterConfig.kubeConfigSecretRef.name}{"\n"}'

echo "--- Step 3: Trigger a Hard Failover ---"
# Simulate a catastrophic failure by deleting the KCC manager pod on the active leader (Cluster 1)
kubectl --context=$C1_CTX delete pod -n cnrm-system -l cnrm.cloud.google.com/component=cnrm-controller-manager

echo "Waiting 30s for the failover transition..."
sleep 30

echo "--- Step 4: Validate the Failover Transition ---"
# 1. Verify GCS Lock Transfer (Expected: "holderIdentity":"cluster-2")
gcloud storage cat gs://$BUCKET/leases/cnrm-system/$LEASE_NAME
echo ""

# 2. Verify Local Lease Consensus (Both should output cluster-2)
kubectl --context=$C1_CTX get multiclusterlease $LEASE_NAME -n cnrm-system -o jsonpath='{.status.globalHolderIdentity}'
echo ""
kubectl --context=$C2_CTX get multiclusterlease $LEASE_NAME -n cnrm-system -o jsonpath='{.status.globalHolderIdentity}'
echo ""

# 3. Verify Syncer Role Reversal
# Cluster 2 (The New Leader) should now be suspended (Expected: Suspend: true)
kubectl --context=$C2_CTX get krmsyncer $SYNCER_NAME -n cnrm-system -o jsonpath='{"Suspend: "}{.spec.suspend}{"\n"}'

# Cluster 1 (The New Follower) should wake up its syncer and point it at Cluster 2
kubectl --context=$C1_CTX get krmsyncer $SYNCER_NAME -n cnrm-system -o jsonpath='{"Suspend: "}{.spec.suspend}{"\nPulling From: "}{.spec.remote.clusterConfig.kubeConfigSecretRef.name}{"\n"}'