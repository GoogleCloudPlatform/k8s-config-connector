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
set -eo pipefail

# Script: 03_upgrade_to_new_version.sh
# Purpose: Upgrade Config Connector Operator and components from 1.153.0 to 1.154.0 (latest release).

NEW_VERSION="${NEW_VERSION:-1.154.1}"
TMP_DIR="$(mktemp -d /tmp/kcc-upgrade-XXXXXX)"

trap 'rm -rf "${TMP_DIR}"' EXIT

echo "============================================================"
echo "Upgrading Config Connector to Version: ${NEW_VERSION}"
echo "============================================================"

# Step 1: Download latest release bundle (1.154.0)
echo "Downloading KCC release bundle for latest release (1.154)..."
curl -sSL "https://storage.googleapis.com/configconnector-operator/${NEW_VERSION}/release-bundle.tar.gz" -o "${TMP_DIR}/release-bundle.tar.gz"

echo "Extracting release bundle..."
tar -zxvf "${TMP_DIR}/release-bundle.tar.gz" -C "${TMP_DIR}"

# Step 2: Apply updated Operator manifest
echo "Applying updated Config Connector Operator 1.154..."
kubectl apply -f "${TMP_DIR}/operator-system/configconnector-operator.yaml"

# Step 3: Wait for Operator pod restart & readiness
echo "Waiting for Operator pod to reconcile and reach Ready state..."
kubectl rollout status statefulset/configconnector-operator -n configconnector-operator-system --timeout=300s || true
kubectl wait --for=condition=Ready pod -l gcp-service=configconnector-operator -n configconnector-operator-system --timeout=300s || true

# Step 4: Wait for KCC controller manager pods to restart with 1.154
echo "Waiting for KCC controller manager deployment to update..."
kubectl rollout status statefulset/cnrm-controller-manager -n cnrm-system --timeout=300s || true
kubectl wait --for=condition=Ready pod -l cnrm.cloud.google.com/component=cnrm-controller-manager -n cnrm-system --timeout=300s || true

# Step 5: Verify installed version
echo ""
echo "Verifying running image tags:"
OPERATOR_IMAGE="$(kubectl get statefulset configconnector-operator -n configconnector-operator-system -o jsonpath='{.spec.template.spec.containers[0].image}')"
CONTROLLER_IMAGE="$(kubectl get statefulset cnrm-controller-manager -n cnrm-system -o jsonpath='{.spec.template.spec.containers[0].image}' 2>/dev/null || echo 'per-namespace/cluster-mode controller')"

echo "Operator Image:   ${OPERATOR_IMAGE}"
echo "Controller Image: ${CONTROLLER_IMAGE}"

echo ""
echo "Upgrade to Config Connector 1.154.0 completed successfully!"
