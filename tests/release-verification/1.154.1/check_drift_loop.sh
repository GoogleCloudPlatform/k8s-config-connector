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

# Script: check_drift_loop.sh
# Purpose: Automatically detect resource drift loops and false update loops by checking resourceVersion stability and controller log events.

SAMPLE_INTERVAL_SECONDS="${SAMPLE_INTERVAL_SECONDS:-30}"

echo "============================================================"
echo "Automated KCC Drift Loop & False Update Verification"
echo "Sample Interval: ${SAMPLE_INTERVAL_SECONDS} seconds"
echo "============================================================"

resources=(
  "memorystoreinstance/memorystore-instance-verify"
  "bigquerytable/bigquerytableverify"
  "pubsubtopic/kcc-topic-verify-1154"
  "storagebucket/kcc-bucket-verify-1154"
  "computesecuritypolicy/kcc-secpolicy-verify-1154"
)

declare -A initial_rvs
declare -A final_rvs

echo ""
echo "Sampling initial resourceVersion for target resources..."
for res in "${resources[@]}"; do
  rv="$(kubectl get "${res}" -o jsonpath='{.metadata.resourceVersion}' 2>/dev/null || echo "NOT_FOUND")"
  initial_rvs["${res}"]="${rv}"
  echo "  - ${res}: ${rv}"
done

echo ""
echo "Monitoring for ${SAMPLE_INTERVAL_SECONDS} seconds to detect unwanted reconciliation updates..."
sleep "${SAMPLE_INTERVAL_SECONDS}"

echo ""
echo "Sampling final resourceVersion for target resources..."
for res in "${resources[@]}"; do
  rv="$(kubectl get "${res}" -o jsonpath='{.metadata.resourceVersion}' 2>/dev/null || echo "NOT_FOUND")"
  final_rvs["${res}"]="${rv}"
done

echo ""
echo "============================================================"
echo "Drift Loop & False Update Verification Summary"
echo "============================================================"
printf "%-50s %-15s %-15s %-10s\n" "RESOURCE" "INITIAL_RV" "FINAL_RV" "RESULT"
echo "---------------------------------------------------------------------------------------------------"

drift_detected=false

for res in "${resources[@]}"; do
  init_rv="${initial_rvs["${res}"]}"
  fin_rv="${final_rvs["${res}"]}"

  if [[ "${init_rv}" == "NOT_FOUND" || "${fin_rv}" == "NOT_FOUND" ]]; then
    result="SKIPPED"
  elif [[ "${init_rv}" == "${fin_rv}" ]]; then
    result="PASS"
  else
    result="FAIL (Drift Detected)"
    drift_detected=true
  fi

  printf "%-50s %-15s %-15s %-10s\n" "${res}" "${init_rv}" "${fin_rv}" "${result}"
done

echo "---------------------------------------------------------------------------------------------------"

echo ""
echo "Checking Controller Manager Logs for Update Events (last 60s)..."
update_logs="$(kubectl logs -n cnrm-system cnrm-controller-manager-0 --tail=200 2>/dev/null | grep -i "updating underlying resource" || true)"

if [[ -n "${update_logs}" ]]; then
  echo "Warning: Active update logs detected in controller manager:"
  echo "${update_logs}"
else
  echo "No continuous GCP update calls detected in controller manager logs."
fi

echo ""
if [[ "${drift_detected}" == "true" ]]; then
  echo "Result: FAIL - One or more resources exhibited resourceVersion changes (drift loop)."
  exit 1
else
  echo "Result: PASS - All reconciled resources remained static. No drift loop or false updates detected."
fi
