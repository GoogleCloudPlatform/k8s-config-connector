#!/usr/bin/env bats

# Copyright 2025 Google LLC
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
    export KIND_CLUSTER_A="cluster-a"
    export KIND_CLUSTER_B="cluster-b"
    export GCS_BUCKET="kcc-mcl-test"
    export LEASE_NAME="configconnector.core.cnrm.cloud.google.com"
    export LEASE_NAMESPACE="cnrm-system"
    # URL encoded leases/cnrm-system/configconnector.core.cnrm.cloud.google.com
    export GCS_LOCK_PATH="leases%2Fcnrm-system%2Fconfigconnector.core.cnrm.cloud.google.com"
}

@test "Initial lock acquisition by Cluster A" {
    # Wait for cluster-a to acquire the lock
    run wait_for_lock "cluster-a" 60
    [ "$status" -eq 0 ]
    
    # Verify lock content in GCS
    run get_lock_holder
    [ "$status" -eq 0 ]
    [[ "$output" == *"cluster-a"* ]]
}

@test "Cluster B takes over lock after Cluster A is scaled to 0" {
    # Scale Cluster A to 0
    kubectl --context kind-cluster-a -n cnrm-system scale statefulset cnrm-controller-manager --replicas=0
    
    # Wait for Cluster B to acquire the lock
    # It might take some time for the lease to expire (15s default) and Cluster B to take over
    run wait_for_lock "cluster-b" 120
    [ "$status" -eq 0 ]
    
    # Verify lock content in GCS
    run get_lock_holder
    [ "$status" -eq 0 ]
    [[ "$output" == *"cluster-b"* ]]
}

wait_for_lock() {
    local expected_holder=$1
    local timeout=$2
    local start_time=$(date +%s)
    
    while true; do
        local current_time=$(date +%s)
        if (( current_time - start_time > timeout )); then
            echo "Timeout waiting for lock acquisition by $expected_holder"
            return 1
        fi
        
        local holder=$(get_lock_holder)
        if [[ "$holder" == *"$expected_holder"* ]]; then
            return 0
        fi
        sleep 5
    done
}

get_lock_holder() {
    # Query the GCS emulator API
    local response=$(curl -s "http://localhost:4443/storage/v1/b/${GCS_BUCKET}/o/${GCS_LOCK_PATH}?alt=media")
    if [[ "$response" == *"Not Found"* || -z "$response" ]]; then
        echo "none"
    else
        echo "$response" | jq -r .holderIdentity 2>/dev/null || echo "error"
    fi
}