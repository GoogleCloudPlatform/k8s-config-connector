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

@test "Resources in testdata/failover are reconciled in Cluster B after failover" {
  PROJECT_ID="mock-project"

  # Find all subdirectories
  for resource_dir in "${FAILOVER_DIR}"/*/; do
    [ -d "${resource_dir}" ] || continue
    RESOURCE_DIR_NAME=$(basename "${resource_dir}")


    test_file="${resource_dir}object.yaml"
    [ -f "${test_file}" ] || continue

    echo ">>> Testing resource type: ${RESOURCE_DIR_NAME} using file: ${test_file}"

    # Set LOG_DIR to the resource directory so everything is saved together
    LOG_DIR="${resource_dir}"
    
    # Generate unique names for each test to avoid collisions in GCP/K8s
    export RESOURCE_NAME="${RESOURCE_DIR_NAME}-mcl-${RANDOM}"
    export PROJECT_ID="${PROJECT_ID}"
    export NAMESPACE="${NAMESPACE}"

    # 1. Extract Kind and apply to Cluster A
    RESOURCE_YAML=$(envsubst < "${test_file}")
    RESOURCE_KIND=$(echo "${RESOURCE_YAML}" | grep "^kind:" | awk '{print $2}')
    RESOURCE_KIND_LOWER=$(echo "${RESOURCE_KIND}" | tr '[:upper:]' '[:lower:]')

    echo "Applying ${RESOURCE_KIND} ${RESOURCE_NAME} to Cluster A..."
    echo "${RESOURCE_YAML}" | kubectl --context="${CTX_A}" apply -f -

    # 2. Verify it's synced to Cluster B by the syncer
    echo "Waiting for ${RESOURCE_NAME} to be synced to Cluster B..."
    retry 10 5 "kubectl --context='${CTX_B}' get ${RESOURCE_KIND_LOWER} '${RESOURCE_NAME}' -n '${NAMESPACE}'"

    # Save YAML from Cluster A before failover
    echo "Saving Cluster A object YAML..."
    kubectl --context="${CTX_A}" get "${RESOURCE_KIND_LOWER}" "${RESOURCE_NAME}" -n "${NAMESPACE}" -o yaml > "${LOG_DIR}/clusterA_initial.yaml" || true

    # 3. Trigger failover: Kill Cluster A KCC
    echo "Capturing Cluster A logs before scale down..."
    kubectl --context "${CTX_A}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=-1 > "${LOG_DIR}/clusterA_manager.log" 2>/dev/null || true

    echo "Scaling down Cluster A KCC..."
    kubectl --context="${CTX_A}" scale deployment cnrm-controller-manager -n "${NAMESPACE}" --replicas=0

    # 4. Wait for Cluster B to take over leadership
    echo "Waiting for Cluster B to take over leadership..."
    retry 15 5 "kubectl --context='${CTX_B}' get multiclusterlease '${LEASE_NAME}' -n '${NAMESPACE}' -o jsonpath='{.status.globalHolderIdentity}' | grep '${CLUSTER_B}'"

    # 5. Verify Cluster B has the resource synced
    echo "Verifying Cluster B has the resource synced..."
    retry 10 5 "kubectl --context='${CTX_B}' get ${RESOURCE_KIND_LOWER} '${RESOURCE_NAME}' -n '${NAMESPACE}'"

    # Save YAML from Cluster B after sync/failover
    echo "Saving Cluster B object YAML..."
    kubectl --context="${CTX_B}" get "${RESOURCE_KIND_LOWER}" "${RESOURCE_NAME}" -n "${NAMESPACE}" -o yaml > "${LOG_DIR}/clusterB_reconcile.yaml" || true

    # 6. Capture Cluster B logs for this resource
    echo "Capturing Cluster B logs..."
    kubectl --context "${CTX_B}" logs -n "${NAMESPACE}" deployment/cnrm-controller-manager --tail=-1 > "${LOG_DIR}/clusterB_manager.log" 2>/dev/null || true

    # 7. Restore Cluster A for the next iteration (if any)
    echo "Scaling up Cluster A KCC for next test iteration..."
    kubectl --context="${CTX_A}" scale deployment cnrm-controller-manager -n "${NAMESPACE}" --replicas=1
    kubectl --context="${CTX_A}" wait --for=condition=available deployment/cnrm-controller-manager -n "${NAMESPACE}" --timeout=120s
  done
}
