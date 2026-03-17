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

export REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "${REPO_ROOT}"

export KUBECONFIG="/tmp/kubeconfig-mcl"
rm -f "${KUBECONFIG}"

export LEASE_NAMESPACE="cnrm-system"
export WATCH_NAMESPACE="tenant-a"
export LEASE_NAME="kcc-leader-lease"
export CTX_A="kind-kcc-cluster-a"
export CTX_B="kind-kcc-cluster-b"

# 1. Cleanup and Setup Environment
echo ">>> Cleaning up old clusters..."
kind delete cluster --name kcc-cluster-a 2>/dev/null || true
kind delete cluster --name kcc-cluster-b 2>/dev/null || true
docker rm -f fake-gcs 2>/dev/null || true

echo ">>> Initializing test environment (Kind clusters)..."
chmod +x tests/e2e/mcl/setup-ha-test-clusters.sh
./tests/e2e/mcl/setup-ha-test-clusters.sh

echo ">>> Environment ready."
echo "Cluster A: ${CTX_A}"
echo "Cluster B: ${CTX_B}"

cleanup() {
  if [ $? -ne 0 ]; then
    echo ">>> Detected failure. Leaving clusters running for debugging: ${KUBECONFIG}"
    echo "--- Cluster A Manager Logs ---"
    kubectl --context "${CTX_A}" logs -n "${WATCH_NAMESPACE}" deployment/cnrm-controller-manager --tail=100 || true
    echo "--- Cluster A Syncer Logs ---"
    kubectl --context "${CTX_A}" logs -n "${LEASE_NAMESPACE}" deployment/syncer-controller-manager --tail=100 || true
    echo "--- Cluster A MCL Controller Logs ---"
    kubectl --context "${CTX_A}" logs -n "${LEASE_NAMESPACE}" deployment/multiclusterlease-controller-manager --tail=100 || true
    echo "--- Cluster B Manager Logs ---"
    kubectl --context "${CTX_B}" logs -n "${WATCH_NAMESPACE}" deployment/cnrm-controller-manager --tail=100 || true
    echo "--- Cluster B Syncer Logs ---"
    kubectl --context "${CTX_B}" logs -n "${LEASE_NAMESPACE}" deployment/syncer-controller-manager --tail=100 || true
    echo "--- Cluster B MCL Controller Logs ---"
    kubectl --context "${CTX_B}" logs -n "${LEASE_NAMESPACE}" deployment/multiclusterlease-controller-manager --tail=100 || true
    return
  fi

  echo ">>> Deleting kind clusters..."
  kind delete cluster --name kcc-cluster-a 2>/dev/null || true
  kind delete cluster --name kcc-cluster-b 2>/dev/null || true
  docker rm -f fake-gcs 2>/dev/null || true
  rm -f "${KUBECONFIG}"
}
trap cleanup EXIT

# 2. Wait for deployments
echo ">>> Waiting for deployments to be ready..."
kubectl --context "${CTX_A}" wait --for=condition=available deployment/cnrm-controller-manager -n "${WATCH_NAMESPACE}" --timeout=120s
kubectl --context "${CTX_B}" wait --for=condition=available deployment/cnrm-controller-manager -n "${WATCH_NAMESPACE}" --timeout=120s

# 3. Run the BATS tests
echo ">>> Running BATS tests..."
if ! bats tests/e2e/mcl/mcl_syncer_integration.bats; then
  echo ">>> ERROR: BATS tests failed."
  exit 1
fi
echo ">>> SUCCESS: HA/DR E2E tests passed."