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

export NAMESPACE="cnrm-system"
export LEASE_NAME="kcc-leader-lease"
export CTX_A="kind-kcc-cluster-a"
export CTX_B="kind-kcc-cluster-b"

echo ">>> Initializing log directory..."
export LOG_DIR="${REPO_ROOT}/tests/e2e/testdata/failover/failover_logs"
mkdir -p "${LOG_DIR}"

# 0. Determine Test Filter and Extra CRDs
TEST_FILTER="${1:-}"
export EXTRA_CRDS=""

if [[ -z "${TEST_FILTER}" ]]; then
  echo ">>> No filter specified. Running basic failover tests..."
  TEST_FILES=("tests/e2e/mcl/mcl_failover.bats")
else
  echo ">>> Resource filter specified: ${TEST_FILTER}"
  export RESOURCE_FILTER="${TEST_FILTER}"
  TEST_FILES=("tests/e2e/mcl/mcl_resource_reconciliation.bats")
  
  # Try to find the CRD for the filtered resource
  CRD_FILE=$(ls config/crds/resources/*${TEST_FILTER}*.yaml 2>/dev/null | head -n 1)
  if [[ -f "${CRD_FILE}" ]]; then
    export EXTRA_CRDS="${CRD_FILE}"
    echo ">>> Found extra CRD to install: ${EXTRA_CRDS}"
  else
    echo ">>> WARNING: Could not find CRD for ${TEST_FILTER}"
  fi
fi

# 1. Cleanup and Setup Environment
echo ">>> Cleaning up old clusters..."
kind delete cluster --name kcc-cluster-a 2>/dev/null || true
kind delete cluster --name kcc-cluster-b 2>/dev/null || true

echo ">>> Initializing test environment (Kind clusters)..."
chmod +x tests/e2e/mcl/setup-ha-test-clusters.sh
./tests/e2e/mcl/setup-ha-test-clusters.sh

echo ">>> Building Fake Syncer and Fake MCL..."
go build -o /tmp/fake-syncer tests/e2e/mcl/fake-syncer-controller.go
go build -o /tmp/fake-mcl tests/e2e/mcl/fake-mcl-controller.go

echo ">>> Building MockGCP Server..."
CGO_ENABLED=0 go build -o /tmp/mockgcp-server tests/e2e/mcl/mockgcp-server.go

echo ">>> Cleaning up previous background processes..."
pkill -f /tmp/mockgcp-server || true
pkill -f /tmp/fake-syncer || true
pkill -f /tmp/fake-mcl || true

echo ">>> Starting Shared MockGCP Server on host..."
/tmp/mockgcp-server > /tmp/mockgcp-host.log 2>&1 &
MOCKGCP_PID=$!
timeout 30s bash -c "until curl -s http://127.0.0.1:8082 > /dev/null; do sleep 1; done"

echo ">>> Starting Fake Syncer (Cluster A -> Cluster B)..."
/tmp/fake-syncer --source-context="${CTX_A}" --target-context="${CTX_B}" --kubeconfig="${KUBECONFIG}" > /tmp/fake-syncer-a-b.log 2>&1 &
SYNCER_PID=$!

echo ">>> Starting Fake MCL Controllers..."
/tmp/fake-mcl --kubecontext="${CTX_A}" --kubeconfig="${KUBECONFIG}" > /tmp/fake-mcl-a.log 2>&1 &
MCL_A_PID=$!
/tmp/fake-mcl --kubecontext="${CTX_B}" --kubeconfig="${KUBECONFIG}" > /tmp/fake-mcl-b.log 2>&1 &
MCL_B_PID=$!

cleanup() {
  # If the last command failed, don't delete clusters/binaries so we can debug
  if [ $? -ne 0 ]; then
    echo ">>> Detected failure. Skipping cleanup of clusters and binaries for debugging."
    return
  fi

  if [[ -n "${RESOURCE_FILTER:-}" ]]; then
    # Move host mockgcp logs to resource dir if we were filtering
    RESOURCE_LOG_DIR="${REPO_ROOT}/tests/e2e/testdata/failover/${RESOURCE_FILTER}"
    if [[ -d "${RESOURCE_LOG_DIR}" ]]; then
      echo ">>> Moving host MockGCP logs to ${RESOURCE_LOG_DIR}..."
      cp /tmp/mockgcp-host.log "${RESOURCE_LOG_DIR}/mockgcp_shared.log" || true
      kubectl --context "${CTX_A}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=-1 > "${RESOURCE_LOG_DIR}/clusterA_manager.log" 2>/dev/null || true
      kubectl --context "${CTX_B}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=-1 > "${RESOURCE_LOG_DIR}/clusterB_manager.log" 2>/dev/null || true
    fi
  fi

  if [[ -z "${RESOURCE_FILTER:-}" ]]; then
    cp /tmp/mockgcp-host.log "${LOG_DIR}/mockgcp_shared.log" || true
    kubectl --context "${CTX_A}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=-1 >> "${LOG_DIR}/clusterA_manager.log" 2>/dev/null || true
    kubectl --context "${CTX_B}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=-1 >> "${LOG_DIR}/clusterB_manager.log" 2>/dev/null || true
  fi
  
  echo ">>> Cleaning up background processes..."
  kill $SYNCER_PID $MOCKGCP_PID $MCL_A_PID $MCL_B_PID 2>/dev/null || true

  if [[ "${SKIP_CLEANUP:-}" != "true" ]]; then
    echo ">>> Deleting kind clusters..."
    kind delete cluster --name kcc-cluster-a 2>/dev/null || true
    kind delete cluster --name kcc-cluster-b 2>/dev/null || true

    echo ">>> Deleting temporary binaries and logs..."
    rm -f /tmp/fake-syncer /tmp/fake-mcl /tmp/mockgcp-server /tmp/kubeconfig-mcl 2>/dev/null || true
    rm -f /tmp/fake-syncer-a-b.log /tmp/mockgcp-host.log /tmp/fake-mcl-a.log /tmp/fake-mcl-b.log 2>/dev/null || true
  else
    echo ">>> Skipping cluster cleanup (SKIP_CLEANUP=true)."
  fi
}
trap cleanup EXIT

echo ">>> Environment ready."
echo "Cluster A: ${CTX_A}"
echo "Cluster B: ${CTX_B}"

# 2. Wait for deployments and CR
echo ">>> Waiting for deployments and lease object to be ready..."
kubectl --context "${CTX_A}" wait --for=condition=available deployment/cnrm-controller-manager -n "${NAMESPACE}" --timeout=120s
kubectl --context "${CTX_B}" wait --for=condition=available deployment/cnrm-controller-manager -n "${NAMESPACE}" --timeout=120s

echo ">>> Waiting for MultiClusterLease objects..."
if ! timeout 60s bash -c "until kubectl --context ${CTX_A} get multiclusterlease ${LEASE_NAME} -n ${NAMESPACE} >/dev/null 2>&1; do sleep 2; done"; then
  echo ">>> ERROR: MultiClusterLease ${LEASE_NAME} not found in Cluster A (${CTX_A}) after 60s."
  kubectl --context "${CTX_A}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=50 || true
  exit 1
fi

if ! timeout 60s bash -c "until kubectl --context ${CTX_B} get multiclusterlease ${LEASE_NAME} -n ${NAMESPACE} >/dev/null 2>&1; do sleep 2; done"; then
  echo ">>> ERROR: MultiClusterLease ${LEASE_NAME} not found in Cluster B (${CTX_B}) after 60s."
  kubectl --context "${CTX_B}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=50 || true
  exit 1
fi

# 3. Run the BATS tests
echo ">>> Running BATS tests..."
for bats_file in "${TEST_FILES[@]}"; do
  echo ">>> Executing ${bats_file}..."
  if ! bats "${bats_file}"; then
    echo ">>> ERROR: BATS tests in ${bats_file} failed. Dumping logs for debugging..."
    echo "--- Cluster A Manager Logs ---"
    kubectl --context "${CTX_A}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=100 || true
    echo "--- Cluster B Manager Logs ---"
    kubectl --context "${CTX_B}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=100 || true
    exit 1
  fi
done

echo ">>> SUCCESS: HA/DR E2E tests passed."
