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
NAMESPACE="cnrm-system"
LEASE_NAME="kcc-leader-lease"
TMP_DIR="/tmp/mcl-e2e-setup"

mkdir -p "${TMP_DIR}"

# 1. Build KCC Binary and Image
echo ">>> Building KCC image..."
export CGO_ENABLED=0
go build -o "${TMP_DIR}/manager" ./cmd/manager

cat <<EOF > "${TMP_DIR}/Dockerfile.kcc"
FROM gcr.io/distroless/static:latest
COPY manager /manager
ENTRYPOINT ["/manager"]
EOF
docker build -t kcc-manager:latest -f "${TMP_DIR}/Dockerfile.kcc" "${TMP_DIR}"

setup_cluster() {
  local name=$1
  local context="kind-${name}"

  if ! kind get clusters | grep -q "^${name}$"; then
    echo "Creating kind cluster: ${name}..."
    kind create cluster --name "${name}"
  fi

  echo "Loading images into ${name}..."
  kind load docker-image kcc-manager:latest --name "${name}"

  # Get host gateway IP (usually 172.18.0.1 for kind)
  # We force IPv4 to avoid unreachable network errors
  HOST_GATEWAY=$(docker network inspect kind -f '{{(index .IPAM.Config 0).Gateway}}' | grep -v ":" || docker network inspect kind -f '{{(index .IPAM.Config 1).Gateway}}' | grep -v ":")

  echo "Installing essential CRDs in ${name}..."
  kubectl --context "${context}" apply -f operator/config/crd/bases/core.cnrm.cloud.google.com_configconnectors.yaml
  kubectl --context "${context}" apply -f tests/e2e/mcl/mcl_crd.yaml

  if [[ -n "${EXTRA_CRDS:-}" ]]; then
    IFS=',' read -ra ADDR <<< "${EXTRA_CRDS}"
    for crd in "${ADDR[@]}"; do
      path="${crd}"
      if [[ ! "$path" = /* ]]; then
        path="${REPO_ROOT}/${path}"
      fi
      echo "Installing extra CRD: ${path} in ${name}..."
      kubectl --context "${context}" apply -f "${path}"
    done
  fi

  echo "Creating namespace ${name}..."
  kubectl --context "${context}" create namespace "${NAMESPACE}" || true

  # RBAC
  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: kcc-admin-binding
subjects:
- kind: ServiceAccount
  name: default
  namespace: ${NAMESPACE}
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
EOF

  # Deploy KCC manager
  echo "Deploying KCC manager in ${name} (pointing to host MockGCP: ${HOST_GATEWAY})..."
  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cnrm-controller-manager
  namespace: ${NAMESPACE}
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
        - --scoped-namespace=${NAMESPACE}
        env:
        - name: HUB_KUBECONFIG
          value: /etc/hub/kubeconfig
        - name: GCP_ENDPOINT
          value: http://${HOST_GATEWAY}:8082
        volumeMounts:
        - name: hub-kubeconfig
          mountPath: /etc/hub
      volumes:
      - name: hub-kubeconfig
        secret:
          secretName: kcc-lease-hub-kubeconfig
          optional: true
EOF

  # ConfigConnector object
  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: cluster
  experiments:
    multiClusterLease:
      leaseName: ${LEASE_NAME}
      namespace: ${NAMESPACE}
      clusterCandidateIdentity: ${name}
EOF
}

# Setup Cluster A
setup_cluster "${CLUSTER_A}"

# Setup Cluster B
setup_cluster "${CLUSTER_B}"

# Cleanup local temp files
rm -rf "${TMP_DIR}"
echo "Setup Complete!"
