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

# Script: 04_apply_post_upgrade_verification.sh
# Purpose: Apply 1.154 post-upgrade manifests (new fields) and verify bug fixes and new field reconciliation.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
POST_UPGRADE_MANIFEST="${SCRIPT_DIR}/manifests/post-upgrade/02-newfields-and-updates.yaml"

echo "============================================================"
echo "Applying Post-Upgrade Test Resources (KCC 1.154.0)"
echo "============================================================"

if [[ ! -f "${POST_UPGRADE_MANIFEST}" ]]; then
  echo "Error: Manifest ${POST_UPGRADE_MANIFEST} not found."
  exit 1
fi

kubectl apply -f "${POST_UPGRADE_MANIFEST}"

echo ""
echo "Waiting 45 seconds for post-upgrade reconciliation..."
sleep 45

echo "============================================================"
echo "Post-Upgrade Verification Results (Release 1.154.1)"
echo "============================================================"

echo ""
echo "[New Fields Verification]"
echo "1. StorageBucket (spec.ipFilter):"
kubectl get storagebucket kcc-bucket-verify-1154 -o jsonpath='{.spec.ipFilter}' && echo ""

echo "2. PubSubTopic (spec.messageStoragePolicy.enforceInTransit):"
kubectl get pubsubtopic kcc-topic-verify-1154 -o jsonpath='{.spec.messageStoragePolicy.enforceInTransit}' && echo ""

echo "3. ComputeSecurityPolicy (spec.region):"
kubectl get computesecuritypolicy kcc-secpolicy-verify-1154 -o jsonpath='{.spec.region}' && echo ""

echo "4. MonitoringAlertPolicy (spec.conditions[].conditionSql):"
kubectl get monitoringalertpolicy monitoringalertpolicy-sql-verify -o jsonpath='{.spec.conditions[*].conditionSql}' && echo ""

echo "5. ComputeRouterNAT (type: PRIVATE):"
kubectl get computerouternat kcc-routernat-privatenat-verify -o jsonpath='{.spec.type}' && echo ""

echo ""
echo "[Bug Fixes & Resource Status Check]"
echo "Checking status conditions across all test resources:"
echo ""

resources=(
  "storagebucket/kcc-bucket-verify-1154"
  "pubsubtopic/kcc-topic-verify-1154"
  "memorystoreinstance/memorystore-instance-verify"
  "bigquerydataset/bigquerydatasetverify"
  "bigquerytable/bigquerytableverify"
  "computesecuritypolicy/kcc-secpolicy-verify-1154"
  "monitoringalertpolicy/monitoringalertpolicy-sql-verify"
  "computerouternat/kcc-routernat-privatenat-verify"
)

for res in "${resources[@]}"; do
  echo "Resource: ${res}"
  kubectl get "${res}" -o jsonpath='{"  Status: "}{.status.conditions[*].status}{"\n  Reason: "}{.status.conditions[*].reason}{"\n  Message: "}{.status.conditions[*].message}{"\n"}' 2>/dev/null || echo "  (not ready or pending)"
  echo "---"
done

echo ""
echo "Post-upgrade verification completed."
