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

export KUBECONFIG="${REPO_ROOT}/.kubeconfig-mcl"
rm -f "${KUBECONFIG}"

export NAMESPACE="cnrm-system"
export LEASE_NAME="kcc-leader-lease"
export CTX_A="kind-kcc-cluster-a"
export CTX_B="kind-kcc-cluster-b"

# 1. Cleanup and Setup Environment
echo ">>> Cleaning up old clusters..."
kind delete cluster --name kcc-cluster-a 2>/dev/null || true
kind delete cluster --name kcc-cluster-b 2>/dev/null || true

echo ">>> Initializing test environment (Kind clusters)..."
chmod +x tests/e2e/mcl/setup-ha-test-clusters.sh
./tests/e2e/mcl/setup-ha-test-clusters.sh

echo ">>> Environment ready."
echo "Cluster A: ${CTX_A}"
echo "Cluster B: ${CTX_B}"

# 2. Wait for deployments and CR
echo ">>> Waiting for deployments and lease object to be ready..."
kubectl --context "${CTX_A}" wait --for=condition=available deployment/cnrm-controller-manager -n "${NAMESPACE}" --timeout=120s
kubectl --context "${CTX_B}" wait --for=condition=available deployment/cnrm-controller-manager -n "${NAMESPACE}" --timeout=120s

# Wait for KCC to create the lease object (if it hasn't already)
echo ">>> Waiting for MultiClusterLease object..."
for i in {1..30}; do
  if kubectl --context "${CTX_A}" get multiclusterlease "${LEASE_NAME}" -n "${NAMESPACE}" &>/dev/null; then
    echo "Found lease object."
    break
  fi
  echo "Still waiting for lease object..."
  sleep 5
done

echo ">>> Waiting for MultiClusterLease object B..."
for i in {1..30}; do
  if kubectl --context "${CTX_B}" get multiclusterlease "${LEASE_NAME}" -n "${NAMESPACE}" &>/dev/null; then
    echo "Found lease object."
    break
  fi
  echo "Still waiting for lease object..."
  sleep 5
done

# 3. Run the BATS tests
echo ">>> Running BATS tests..."
if ! bats tests/e2e/mcl/mcl_failover.bats; then
  echo ">>> ERROR: BATS tests failed. Dumping logs for debugging..."
  echo "--- Cluster A Manager Logs ---"
  kubectl --context "${CTX_A}" logs -n "${NAMESPACE}" -l control-plane=controller-manager --tail=100 || true
  echo "--- Cluster B Manager Logs ---"
  kubectl --context "${CTX_B}" logs -n "${NAMESPACE}" -l control-plane=controller-manager --tail=100 || true
  exit 1
fi
echo ">>> SUCCESS: HA/DR E2E tests passed."
