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
  export LEASE_NAMESPACE="cnrm-system"
  export WATCH_NAMESPACE="tenant-a"
  export LEASE_NAME="kcc-leader-lease"
  export SYNCER_NAME="${LEASE_NAME}-${WATCH_NAMESPACE}"
}

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

@test "KRMSyncer configures correctly and isolates resources during failover" {
  echo "Waiting for KCC to create local leases..."
  retry 15 5 "kubectl --context='${CTX_A}' get multiclusterlease '${LEASE_NAME}' -n '${LEASE_NAMESPACE}'"
  retry 15 5 "kubectl --context='${CTX_B}' get multiclusterlease '${LEASE_NAME}' -n '${LEASE_NAMESPACE}'"

  echo "Waiting for REAL KRMSyncer to be ready in both clusters..."
  retry 15 5 "kubectl --context='${CTX_A}' wait --for=condition=available deployment/syncer-controller-manager -n '${LEASE_NAMESPACE}' --timeout=10s"
  retry 15 5 "kubectl --context='${CTX_B}' wait --for=condition=available deployment/syncer-controller-manager -n '${LEASE_NAMESPACE}' --timeout=10s"

  echo "Waiting for REAL MCL Controller to be ready in both clusters..."
  retry 15 5 "kubectl --context='${CTX_A}' wait --for=condition=available deployment/multiclusterlease-controller-manager -n '${LEASE_NAMESPACE}' --timeout=10s"
  retry 15 5 "kubectl --context='${CTX_B}' wait --for=condition=available deployment/multiclusterlease-controller-manager -n '${LEASE_NAMESPACE}' --timeout=10s"

  echo "Waiting for REAL MCL Controller to elect a leader..."
  retry 30 5 "kubectl --context='${CTX_A}' get multiclusterlease '${LEASE_NAME}' -n '${LEASE_NAMESPACE}' -o json | jq -e '.status.globalHolderIdentity != null and .status.globalHolderIdentity != \"\"'"

  local leader=$(kubectl --context="${CTX_A}" get multiclusterlease "${LEASE_NAME}" -n "${LEASE_NAMESPACE}" -o json | jq -r '.status.globalHolderIdentity')
  local follower
  local leader_ctx
  local follower_ctx
  if [ "$leader" = "${CLUSTER_A}" ]; then
     follower="${CLUSTER_B}"
     leader_ctx="${CTX_A}"
     follower_ctx="${CTX_B}"
  else
     follower="${CLUSTER_A}"
     leader_ctx="${CTX_B}"
     follower_ctx="${CTX_A}"
  fi
  
  echo "Elected Leader: $leader"

  echo "Verifying Follower ($follower) configured its Namespaced KRMSyncer to PULL from Leader..."
  retry 15 3 "kubectl --context='${follower_ctx}' get krmsyncer '${SYNCER_NAME}' -n '${LEASE_NAMESPACE}' -o json | jq -e '.spec.suspend == false'"

  echo "Verifying Leader ($leader) suspended its KRMSyncer..."
  retry 15 3 "kubectl --context='${leader_ctx}' get krmsyncer '${SYNCER_NAME}' -n '${LEASE_NAMESPACE}' -o json | jq -e '.spec.suspend == true'"

  # --------------------------------------------------------------------------------
  # DATA PLANE VERIFICATION: CREATE & SYNC
  # --------------------------------------------------------------------------------
  
  echo "Applying StorageBucket to both clusters (Simulating GitOps)..."
  kubectl --context="${CTX_A}" apply -f tests/e2e/mcl/storagebucket.yaml
  kubectl --context="${CTX_B}" apply -f tests/e2e/mcl/storagebucket.yaml

  echo "Waiting for Leader to reconcile the resource..."
  if ! retry 15 5 "kubectl --context='${leader_ctx}' get storagebucket test-ha-bucket -n '${WATCH_NAMESPACE}' -o json | grep -E 'UpdateFailed|UpToDate'"; then
    echo "Dumping StorageBucket state for debugging:"
    kubectl --context="${leader_ctx}" get storagebucket test-ha-bucket -n "${WATCH_NAMESPACE}" -o json
    exit 1
  fi

  echo "Waiting for REAL KRMSyncer to copy the status from Leader to Follower..."
  retry 30 3 "kubectl --context='${follower_ctx}' get storagebucket test-ha-bucket -n '${WATCH_NAMESPACE}' -o json | grep -E 'UpdateFailed|UpToDate'"

  # --------------------------------------------------------------------------------
  # FAILOVER
  # --------------------------------------------------------------------------------

  echo "Scaling down Leader MCL controller and KCC manager to simulate failure..."
  kubectl --context="${leader_ctx}" scale deployment cnrm-controller-manager -n "${WATCH_NAMESPACE}" --replicas=0
  kubectl --context="${leader_ctx}" scale deployment multiclusterlease-controller-manager -n "${LEASE_NAMESPACE}" --replicas=0

  echo "Waiting for Follower ($follower) to acquire the lease..."
  retry 30 5 "kubectl --context='${follower_ctx}' get multiclusterlease '${LEASE_NAME}' -n '${LEASE_NAMESPACE}' -o json | jq -e '.status.globalHolderIdentity == \"${follower}\"'"

  echo "Verifying New Leader ($follower) suspended its KRMSyncer..."
  retry 20 2 "kubectl --context='${follower_ctx}' get krmsyncer '${SYNCER_NAME}' -n '${LEASE_NAMESPACE}' -o json | jq -e '.spec.suspend == true'"

  # --------------------------------------------------------------------------------
  # MODIFY IN NEW LEADER
  # --------------------------------------------------------------------------------

  echo "Modifying the resource in New Leader and old leader (Adding a spec change to bump generation)..."
  kubectl --context="${follower_ctx}" patch storagebucket test-ha-bucket -n "${WATCH_NAMESPACE}" --type='merge' -p '{"spec":{"storageClass":"STANDARD"}}'
  kubectl --context="${leader_ctx}" patch storagebucket test-ha-bucket -n "${WATCH_NAMESPACE}" --type='merge' -p '{"spec":{"storageClass":"STANDARD"}}'
  
  echo "Waiting for New Leader to reconcile the modification (generation 2)..."
  retry 20 2 "kubectl --context='${follower_ctx}' get storagebucket test-ha-bucket -n '${WATCH_NAMESPACE}' -o json | jq -e '.status.observedGeneration == 2'"

  # --------------------------------------------------------------------------------
  # FAILBACK
  # --------------------------------------------------------------------------------

  echo "Scaling up Old Leader KCC (Restoring as Follower)..."
  kubectl --context="${leader_ctx}" scale deployment cnrm-controller-manager -n "${WATCH_NAMESPACE}" --replicas=1

  echo "Waiting for Old Leader (Now Follower) to configure its Syncer to pull from New Leader..."
  retry 20 2 "kubectl --context='${leader_ctx}' get krmsyncer '${SYNCER_NAME}' -n '${LEASE_NAMESPACE}' -o json | jq -e '.spec.suspend == false'"

  # Simulate GitOps replicating the spec change to Old Leader
  kubectl --context="${leader_ctx}" patch storagebucket test-ha-bucket -n "${WATCH_NAMESPACE}" --type='merge' -p '{"spec":{"storageClass":"STANDARD"}}'

  echo "Waiting for REAL KRMSyncer to copy the updated status from New Leader to Old Leader while Old Leader is a follower..."
  retry 30 3 "kubectl --context='${leader_ctx}' get storagebucket test-ha-bucket -n '${WATCH_NAMESPACE}' -o json | jq -e '.status.observedGeneration == 2'"

  echo "Scaling up Old Leader MCL controller to simulate failback..."
  kubectl --context="${leader_ctx}" scale deployment multiclusterlease-controller-manager -n "${LEASE_NAMESPACE}" --replicas=1

  # To force the failback, we can scale down the current leader's MCL controller
  echo "Scaling down Current Leader ($follower) MCL controller to force failback..."
  kubectl --context="${follower_ctx}" scale deployment multiclusterlease-controller-manager -n "${LEASE_NAMESPACE}" --replicas=0

  echo "Waiting for Original Leader ($leader) to re-acquire the lease..."
  retry 30 5 "kubectl --context='${leader_ctx}' get multiclusterlease '${LEASE_NAME}' -n '${LEASE_NAMESPACE}' -o json | jq -e '.status.globalHolderIdentity == \"${leader}\"'"

  echo "Verifying Original Leader suspended its KRMSyncer..."
  retry 20 2 "kubectl --context='${leader_ctx}' get krmsyncer '${SYNCER_NAME}' -n '${LEASE_NAMESPACE}' -o json | jq -e '.spec.suspend == true'"

  echo "Verifying Original Follower configured its Syncer to pull from Original Leader..."
  retry 20 2 "kubectl --context='${follower_ctx}' get krmsyncer '${SYNCER_NAME}' -n '${LEASE_NAMESPACE}' -o json | jq -e '.spec.suspend == false'"

  echo "Modifying the resource in both clusters (Adding another spec change)..."
  kubectl --context="${leader_ctx}" patch storagebucket test-ha-bucket -n "${WATCH_NAMESPACE}" --type='merge' -p '{"spec":{"storageClass":"NEARLINE"}}'
  kubectl --context="${follower_ctx}" patch storagebucket test-ha-bucket -n "${WATCH_NAMESPACE}" --type='merge' -p '{"spec":{"storageClass":"NEARLINE"}}'

  echo "Waiting for Original Leader to reconcile the failback modification (generation 3)..."
  retry 20 2 "kubectl --context='${leader_ctx}' get storagebucket test-ha-bucket -n '${WATCH_NAMESPACE}' -o json | jq -e '.status.observedGeneration == 3'"

  echo "Waiting for REAL KRMSyncer to copy the updated status from Original Leader back to Original Follower..."
  retry 30 3 "kubectl --context='${follower_ctx}' get storagebucket test-ha-bucket -n '${WATCH_NAMESPACE}' -o json | jq -e '.status.observedGeneration == 3'"

  # Cleanup fake-gcs is handled in run-ha-e2e-tests.sh
}