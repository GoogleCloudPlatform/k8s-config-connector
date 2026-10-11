#!/usr/bin/env bats

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

setup() {
  export REPO_ROOT="${REPO_ROOT:-$(git rev-parse --show-toplevel)}"
  export FAILOVER_DIR="${REPO_ROOT}/tests/e2e/testdata/failover"
  export LOG_DIR="${FAILOVER_DIR}/failover_logs"
  export CLUSTER_A="kcc-cluster-a"
  export CLUSTER_B="kcc-cluster-b"
  export CTX_A="kind-${CLUSTER_A}"
  export CTX_B="kind-${CLUSTER_B}"
  export NAMESPACE="cnrm-system"
  export LEASE_NAME="kcc-leader-lease"
  mkdir -p "${LOG_DIR}"
}

# Helper to retry a command until success or timeout
retry() {
  local count=$1
  local delay=$2
  shift 2
  local cmd="$@"
  for i in $(seq 1 $count); do
    if eval "$cmd"; then
      return 0
    fi
    echo "Attempt $i failed, retrying in ${delay}s..."
    sleep $delay
  done
  return 1
}

@test "Cluster B acquires leadership when Cluster A is scaled down" {
  # 0. Wait for CRDs to be ready
  echo "Waiting for MultiClusterLease CRDs in both clusters..."
  retry 5 2 "kubectl --context='${CTX_A}' get crd multiclusterleases.multicluster.core.cnrm.cloud.google.com"
  retry 5 2 "kubectl --context='${CTX_B}' get crd multiclusterleases.multicluster.core.cnrm.cloud.google.com"

  # 1. Ensure initial state: Cluster A is leader
  echo "Waiting for Cluster A to hold the global lease..."
  retry 10 5 "kubectl --context='${CTX_A}' get multiclusterlease '${LEASE_NAME}' -n '${NAMESPACE}' -o jsonpath='{.status.globalHolderIdentity}' | grep '${CLUSTER_A}'"

  # 2. Confirm Cluster A logs show it leading
  echo "Verifying Cluster A logs for 'started leading'"
  retry 10 3 "kubectl --context='${CTX_A}' logs -n '${NAMESPACE}' deployment/cnrm-controller-manager | grep 'started leading'"

  # 3. TRIGGER FAILOVER: Kill Cluster A KCC
  echo "Capturing Cluster A logs before scale down..."
  kubectl --context "${CTX_A}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=-1 > "${LOG_DIR}/clusterA_manager.log" 2>/dev/null || true

  echo "Scaling down Cluster A KCC (simulating failure)..."
  kubectl --context="${CTX_A}" scale deployment cnrm-controller-manager -n "${NAMESPACE}" --replicas=0

  # 4. Verify Cluster B acquires leadership
  echo "Waiting for Cluster B to take over leadership (this may take ~15-20s for timeout)..."
  retry 10 5 "kubectl --context='${CTX_B}' get multiclusterlease '${LEASE_NAME}' -n '${NAMESPACE}' -o jsonpath='{.status.globalHolderIdentity}' | grep '${CLUSTER_B}'"

  # 5. Check KCC logs in Cluster B to confirm it started leading
  echo "Verifying Cluster B logs for 'started leading'"
  retry 10 3 "kubectl --context='${CTX_B}' logs -n '${NAMESPACE}' deployment/cnrm-controller-manager | grep 'started leading'"
  
  # 6. Capture Cluster B logs
  echo "Capturing Cluster B logs..."
  kubectl --context "${CTX_B}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=-1 > "${LOG_DIR}/clusterB_manager.log" 2>/dev/null || true

  # 7. Restore Cluster A
  echo "Scaling up Cluster A KCC..."
  kubectl --context="${CTX_A}" scale deployment cnrm-controller-manager -n "${NAMESPACE}" --replicas=1
  kubectl --context="${CTX_A}" wait --for=condition=available deployment/cnrm-controller-manager -n "${NAMESPACE}" --timeout=120s
}
