#!/usr/bin/env bash

# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "${REPO_ROOT}"

export KUBECONFIG="${KUBECONFIG:-/tmp/kubeconfig-mcl}"

CLUSTER_A="kcc-cluster-a"
CLUSTER_B="kcc-cluster-b"
LEASE_NAMESPACE="cnrm-system"
WATCH_NAMESPACE="tenant-a"
LEASE_NAME="kcc-leader-lease"
TMP_DIR="/tmp/mcl-e2e-setup"

mkdir -p "${TMP_DIR}"

# 1. Build KCC Binary and Syncer Binary
echo ">>> Building KCC image..."
export CGO_ENABLED=0
go build -o "${TMP_DIR}/manager" ./cmd/manager

cat <<EOF > "${TMP_DIR}/Dockerfile.kcc"
FROM gcr.io/distroless/static:latest
COPY manager /manager
ENTRYPOINT ["/manager"]
EOF
docker build -t kcc-manager:latest -f "${TMP_DIR}/Dockerfile.kcc" "${TMP_DIR}"

echo ">>> Building Syncer image..."
go build -o "${TMP_DIR}/syncer" github.com/gke-labs/kube-etl/syncer

cat <<EOF > "${TMP_DIR}/Dockerfile.syncer"
FROM gcr.io/distroless/static:latest
COPY syncer /manager
ENTRYPOINT ["/manager"]
EOF
docker build -t real-syncer:latest -f "${TMP_DIR}/Dockerfile.syncer" "${TMP_DIR}"

echo ">>> Building MCL Controller image..."
go build -o "${TMP_DIR}/mcl-controller" github.com/gke-labs/multicluster-leader-election/cmd/manager

cat <<EOF > "${TMP_DIR}/Dockerfile.mcl"
FROM gcr.io/distroless/static:latest
COPY mcl-controller /manager
ENTRYPOINT ["/manager"]
EOF
docker build -t mcl-controller:latest -f "${TMP_DIR}/Dockerfile.mcl" "${TMP_DIR}"

setup_cluster() {
  local name=$1
  local context="kind-${name}"

  if ! kind get clusters | grep -q "^${name}$" ; then
    echo "Creating kind cluster: ${name}..."
    kind create cluster --name "${name}" --image kindest/node:v1.30.0
  fi

  # Start fake-gcs if not already running on the kind network
  if ! docker ps | grep -q "fake-gcs"; then
    echo "Starting fake-gcs-server on kind network..."
    docker rm -f fake-gcs 2>/dev/null || true
    mkdir -p "${TMP_DIR}/fake-gcs-data/mcl-lease-bucket"
    chmod -R 777 "${TMP_DIR}/fake-gcs-data"
    docker run -d --name fake-gcs --network kind \
      -v "${TMP_DIR}/fake-gcs-data":/data \
      fsouza/fake-gcs-server -scheme http -public-host fake-gcs:4443 -external-url http://fake-gcs:4443
    sleep 3
  fi

  echo "Loading images into ${name}..."
  kind load docker-image kcc-manager:latest --name "${name}"
  kind load docker-image real-syncer:latest --name "${name}"
  kind load docker-image mcl-controller:latest --name "${name}"

  echo "Installing essential CRDs in ${name}..."
  kubectl --context "${context}" apply -f operator/config/crd/bases/core.cnrm.cloud.google.com_configconnectors.yaml
  kubectl --context "${context}" apply -f operator/config/crd/bases/core.cnrm.cloud.google.com_configconnectorcontexts.yaml
  kubectl --context "${context}" apply -f tests/e2e/mcl/mcl_crd.yaml
  kubectl --context "${context}" apply -f tests/e2e/mcl/syncer_crd.yaml
  
  echo "Installing all KCC CRDs to prevent syncer dynamic puller crash loops..."
  kubectl --context "${context}" apply -f config/crds/resources/ >/dev/null 2>&1 || true

  echo "Creating namespaces in ${name}..."
  kubectl --context "${context}" create namespace "${LEASE_NAMESPACE}" || true
  kubectl --context "${context}" create namespace "${WATCH_NAMESPACE}" || true

  # RBAC
  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: kcc-admin-binding
subjects:
- kind: ServiceAccount
  name: default
  namespace: ${WATCH_NAMESPACE}
- kind: ServiceAccount
  name: default
  namespace: ${LEASE_NAMESPACE}
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
EOF

  # Deploy Real Syncer
  echo "Deploying Real KRMSyncer Controller in ${name}..."
  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: syncer-controller-manager
  namespace: ${LEASE_NAMESPACE}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: syncer-controller
  template:
    metadata:
      labels:
        app: syncer-controller
    spec:
      containers:
      - name: manager
        image: real-syncer:latest
        imagePullPolicy: IfNotPresent
EOF

  # Deploy Real MCL Controller
  echo "Deploying Real MCL Controller in ${name}..."
  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: multiclusterlease-controller-manager
  namespace: ${LEASE_NAMESPACE}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mcl-controller
  template:
    metadata:
      labels:
        app: mcl-controller
    spec:
      containers:
      - name: manager
        image: mcl-controller:latest
        imagePullPolicy: IfNotPresent
        args:
        - --gcs-bucket=mcl-lease-bucket
        env:
        - name: STORAGE_EMULATOR_HOST
          value: http://fake-gcs:4443
EOF

  # Create mock credentials
  kubectl --context "${context}" create secret generic mock-gcp-creds -n "${WATCH_NAMESPACE}" --from-literal=key.json='{"type":"service_account","project_id":"mock-project","private_key_id":"mock","private_key":"-----BEGIN PRIVATE KEY-----\nmock\n-----END PRIVATE KEY-----\n","client_email":"mock@mock-project.iam.gserviceaccount.com","client_id":"mock","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_x509_cert_url":"https://www.googleapis.com/robot/v1/metadata/x509/mock%40mock-project.iam.gserviceaccount.com"}' || true

  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: ${WATCH_NAMESPACE}
spec:
  googleServiceAccount: "mock@mock-project.iam.gserviceaccount.com"
EOF

  # Deploy KCC manager (Namespaced Mode)
  echo "Deploying KCC manager in ${name} for namespace ${WATCH_NAMESPACE}..."

  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cnrm-controller-manager
  namespace: ${WATCH_NAMESPACE}
  labels:
    control-plane: controller-manager
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: controller-manager
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      containers:
      - name: manager
        image: kcc-manager:latest
        imagePullPolicy: IfNotPresent
        args:
        - --leader-election-type=multicluster
        - --syncing-mode=pull
        - --scoped-namespace=${WATCH_NAMESPACE}
        env:
        - name: GCP_ENDPOINT
          value: http://mockgcp.local
        - name: GOOGLE_APPLICATION_CREDENTIALS
          value: /etc/gcp/key.json
        volumeMounts:
        - name: gcp-creds
          mountPath: /etc/gcp
      volumes:
      - name: gcp-creds
        secret:
          secretName: mock-gcp-creds
EOF

  # We also need a ConfigConnector object just so kccmanager.go can read the experiments config
  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
  experiments:
    multiClusterLease:
      leaseName: ${LEASE_NAME}
      namespace: ${LEASE_NAMESPACE}
      clusterCandidateIdentity: ${name}
EOF
}

# Setup Cluster A
setup_cluster "${CLUSTER_A}"

# Setup Cluster B
setup_cluster "${CLUSTER_B}"

echo ">>> Setting up Cross-Cluster Kubeconfigs for KRMSyncer..."
# Extract internal kubeconfigs with the correct docker network IPs
docker exec "${CLUSTER_A}-control-plane" cat /etc/kubernetes/admin.conf > "${TMP_DIR}/cluster-a.kubeconfig"
docker exec "${CLUSTER_B}-control-plane" cat /etc/kubernetes/admin.conf > "${TMP_DIR}/cluster-b.kubeconfig"

A_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' "${CLUSTER_A}-control-plane")
B_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' "${CLUSTER_B}-control-plane")

sed -i "s/.*server:.*/    server: https:\/\/${A_IP}:6443\n    insecure-skip-tls-verify: true/g" "${TMP_DIR}/cluster-a.kubeconfig"
sed -i '/certificate-authority-data:/d' "${TMP_DIR}/cluster-a.kubeconfig"

sed -i "s/.*server:.*/    server: https:\/\/${B_IP}:6443\n    insecure-skip-tls-verify: true/g" "${TMP_DIR}/cluster-b.kubeconfig"
sed -i '/certificate-authority-data:/d' "${TMP_DIR}/cluster-b.kubeconfig"

kubectl --context "kind-${CLUSTER_B}" create secret generic "${CLUSTER_A}" -n "${LEASE_NAMESPACE}" --from-file=kubeconfig="${TMP_DIR}/cluster-a.kubeconfig" || true
kubectl --context "kind-${CLUSTER_A}" create secret generic "${CLUSTER_B}" -n "${LEASE_NAMESPACE}" --from-file=kubeconfig="${TMP_DIR}/cluster-b.kubeconfig" || true

# Cleanup local temp files
rm -rf "${TMP_DIR}"
echo "Setup Complete!"
