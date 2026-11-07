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
  export CLUSTER_A="kcc-cluster-a"
  export CLUSTER_B="kcc-cluster-b"
  export CTX_A="kind-${CLUSTER_A}"
  export CTX_B="kind-${CLUSTER_B}"
  export NAMESPACE="cnrm-system"
  export LEASE_NAME="kcc-leader-lease"
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
  # 0. Wait for MCL Controllers to be ready
  echo "Waiting for MCL Controllers in both clusters..."
  kubectl --context="${CTX_A}" wait --for=condition=available deployment/mcl-controller -n "${NAMESPACE}" --timeout=60s
  kubectl --context="${CTX_B}" wait --for=condition=available deployment/mcl-controller -n "${NAMESPACE}" --timeout=60s

  # 1. Ensure initial state: Cluster A is leader
  # Both are running, but Cluster A typically wins if it starts first or has alphabetical priority
  echo "Waiting for Cluster A to hold the global lease..."
  retry 10 5 "kubectl --context='${CTX_A}' get multiclusterlease '${LEASE_NAME}' -n '${NAMESPACE}' -o jsonpath='{.status.globalHolderIdentity}' | grep '${CLUSTER_A}'"

  # 2. Confirm Cluster A logs show it leading
  echo "Verifying Cluster A logs for 'started leading'"
  retry 10 3 "kubectl --context='${CTX_A}' logs -n '${NAMESPACE}' -l control-plane=controller-manager | grep 'started leading'"

  # 3. TRIGGER FAILOVER: Kill Cluster A KCC
  echo "Scaling down Cluster A KCC (simulating failure)..."
  kubectl --context="${CTX_A}" scale deployment cnrm-controller-manager -n "${NAMESPACE}" --replicas=0

  # 4. Verify Cluster B acquires leadership
  # MCL controller in Cluster B will see the GCS lease is stale (after timeout) and allow B to take over.
  echo "Waiting for Cluster B to take over leadership (this may take ~15-20s for timeout)..."
  retry 10 5 "kubectl --context='${CTX_B}' get multiclusterlease '${LEASE_NAME}' -n '${NAMESPACE}' -o jsonpath='{.status.globalHolderIdentity}' | grep '${CLUSTER_B}'"

  # 5. Check KCC logs in Cluster B to confirm it started leading
  echo "Verifying Cluster B logs for 'started leading'"
  retry 10 3 "kubectl --context='${CTX_B}' logs -n '${NAMESPACE}' -l control-plane=controller-manager | grep 'started leading'"
}
