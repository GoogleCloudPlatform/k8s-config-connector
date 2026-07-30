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

# Script: 02_apply_pre_upgrade_resources.sh
# Purpose: Apply pre-upgrade resources on KCC 1.153.0 and wait/verify their reconciliation status.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PRE_UPGRADE_MANIFEST="${SCRIPT_DIR}/manifests/pre-upgrade/01-bugfixes-and-base.yaml"

echo "============================================================"
echo "Applying Pre-Upgrade Test Resources (KCC 1.153.0)"
echo "============================================================"

if [[ ! -f "${PRE_UPGRADE_MANIFEST}" ]]; then
  echo "Error: Manifest ${PRE_UPGRADE_MANIFEST} not found."
  exit 1
fi

# Step 1: Safely apply / re-apply pre-upgrade manifests (Idempotent)
kubectl apply -f "${PRE_UPGRADE_MANIFEST}"

echo ""
echo "Waiting for initial reconciliation by KCC 1.153.0..."

# Step 2: Poll for up to 60 seconds for resources to reconcile
resources=(
  "storagebucket/kcc-bucket-verify-1154"
  "pubsubtopic/kcc-topic-verify-1154"
  "memorystoreinstance/memorystore-instance-verify"
  "bigquerydataset/bigquerydatasetverify"
  "bigquerytable/bigquerytableverify"
  "computesecuritypolicy/kcc-secpolicy-verify-1154"
)

for attempt in {1..12}; do
  all_ready=true
  for res in "${resources[@]}"; do
    status="$(kubectl get "${res}" -o jsonpath='{.status.conditions[?(@.type=="Ready")].status}' 2>/dev/null || echo "False")"
    if [[ "${status}" != "True" ]]; then
      all_ready=false
      break
    fi
  done

  if [[ "${all_ready}" == "true" ]]; then
    echo "All pre-upgrade resources reconciled to Ready=True!"
    break
  fi

  echo "Waiting for resources to reconcile... (attempt ${attempt}/12)"
  sleep 5
done

echo ""
echo "============================================================"
echo "Pre-Upgrade Resource Status Summary (KCC 1.153.0)"
echo "============================================================"

for res in "${resources[@]}"; do
  echo "Resource: ${res}"
  kubectl get "${res}" -o jsonpath='{"  Ready Status: "}{.status.conditions[?(@.type=="Ready")].status}{"\n  Reason:       "}{.status.conditions[?(@.type=="Ready")].reason}{"\n  Message:      "}{.status.conditions[?(@.type=="Ready")].message}{"\n"}' 2>/dev/null || echo "  (pending status)"
  echo "---"
done

echo "Pre-upgrade manifest application complete. (Safe to re-run anytime)"
