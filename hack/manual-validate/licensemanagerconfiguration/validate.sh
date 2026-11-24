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

# Human-style validation of the LicenseManagerConfiguration controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/licensemanagerconfiguration/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the licensemanagerconfiguration-maximal test fixture)
#   FUNCTIONAL=0     skip functional checks (listing configurations via REST)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/licensemanager/v1alpha1/licensemanagerconfiguration/licensemanagerconfiguration-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init LicenseManagerConfiguration
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
CONFIG_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
CONFIG_ID="$(grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
CONFIG_ID="${CONFIG_ID:-${CONFIG_K8S}}"
LOCATION="$(grep -m1 '^[[:space:]]*location:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
LOCATION="${LOCATION:-us-central1}"

API_BASE="https://licensemanager.googleapis.com/v1"
PARENT="projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
CONFIG_URL="${API_BASE}/${PARENT}/configurations/${CONFIG_ID}"
LIST_URL="${API_BASE}/${PARENT}/configurations"
EXPECTED_NAME="${PARENT}/configurations/${CONFIG_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get licensemanagerconfiguration "${CONFIG_K8S}" >/dev/null 2>&1 && mv_delete cleanup LicenseManagerConfiguration "${CONFIG_K8S}" 300
    if [[ -n "${DEPS:-}" ]]; then
      # Delete dependencies in reverse order if any exist
      :
    fi
  fi

  local code
  code="$(mv_rest cleanup leakcheck-config GET "${CONFIG_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (configuration gone)" PASS cloud/cleanup-leakcheck-config.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (configuration gone)" FAIL cloud/cleanup-leakcheck-config.json \
      "could not verify (HTTP ${code}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${CONFIG_URL}"
  else
    mv_record cleanup "no leaked GCP resources (configuration gone)" FAIL cloud/cleanup-leakcheck-config.json \
      "HTTP ${code}; clean up manually: curl -X DELETE -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${CONFIG_URL}"
  fi

  if [[ "${KEEP_ENV:-0}" == "1" ]]; then
    mv_log "KEEP_ENV=1: envtest + manager left running. export KUBECONFIG=${KUBECONFIG}; pids:${MV_PIDS}"
  else
    mv_down
  fi
  mv_report || rc=1
  exit ${rc}
}
trap cleanup EXIT

# ---------------------------------------------------------------------------
# Every API the controller, dependencies and out-of-band checks call.
mv_enable_apis setup licensemanager.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create LicenseManagerConfiguration "${CONFIG_K8S}" 300
SNAP="$(mv_snapshot create LicenseManagerConfiguration "${CONFIG_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

mv_wait_http create configuration "${CONFIG_URL}" 200

# Compare spec fields against cloud response
WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_PRODUCT="$(jq -r '.spec.product // empty' "${SNAP}")"
WANT_LICENSE_TYPE="$(jq -r '.spec.licenseType // empty' "${SNAP}")"
WANT_CURRENT_COUNT="$(jq -r '.spec.currentBillingInfo.userCountBilling.userCount // empty' "${SNAP}")"
WANT_NEXT_COUNT="$(jq -r '.spec.nextBillingInfo.userCountBilling.userCount // empty' "${SNAP}")"

# Verify cloud resource name
mv_expect_jq create "cloud resource name matches expected ID" "${MV_RUN_DIR}/cloud/create-configuration.json" '.name' "${EXPECTED_NAME}"

if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-configuration.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_PRODUCT}" ]]; then
  mv_expect_jq create "cloud product matches spec" "${MV_RUN_DIR}/cloud/create-configuration.json" '.product' "${WANT_PRODUCT}"
fi
if [[ -n "${WANT_CURRENT_COUNT}" ]]; then
  mv_expect_jq create "cloud currentBillingInfo userCount matches spec" "${MV_RUN_DIR}/cloud/create-configuration.json" '(.currentBillingInfo.userCountBilling.userCount | tostring)' "${WANT_CURRENT_COUNT}"
fi
if [[ -n "${WANT_NEXT_COUNT}" ]]; then
  mv_expect_jq create "cloud nextBillingInfo userCount matches spec" "${MV_RUN_DIR}/cloud/create-configuration.json" '(.nextBillingInfo.userCountBilling.userCount | tostring)' "${WANT_NEXT_COUNT}"
fi
if [[ -n "${WANT_LICENSE_TYPE}" ]]; then
  GOT_LICENSE_TYPE="$(jq -r '.licenseType // empty' "${MV_RUN_DIR}/cloud/create-configuration.json")"
  if [[ "${GOT_LICENSE_TYPE}" == "${WANT_LICENSE_TYPE}" || "${GOT_LICENSE_TYPE}" == "1" ]]; then
    mv_record create "cloud licenseType matches spec" PASS cloud/create-configuration.json "${GOT_LICENSE_TYPE}"
  else
    mv_record create "cloud licenseType matches spec" FAIL cloud/create-configuration.json "want '${WANT_LICENSE_TYPE}', got '${GOT_LICENSE_TYPE}'"
  fi
fi

# Verify every label in spec.labels is present with the correct value in GCP response
while IFS="=" read -r key val; do
  [[ -z "${key}" ]] && continue
  mv_expect_jq create "cloud label '${key}' matches spec" "${MV_RUN_DIR}/cloud/create-configuration.json" ".labels[\"${key}\"]" "${val}"
done < <(jq -r '.spec.labels // {} | to_entries[] | "\(.key)=\(.value)"' "${SNAP}")
mv_expect_jq create "cloud label 'managed-by-cnrm' present" "${MV_RUN_DIR}/cloud/create-configuration.json" '.labels["managed-by-cnrm"]' "true"

# Note on gcloud CLI support:
# gcloud does not support License Manager (REST API only).
mv_record create "gcloud CLI surface" INFO env.yaml "no gcloud CLI surface for licensemanager; verified via REST"

# Functional check: ListConfigurations out of band
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create list-configurations GET "${LIST_URL}")"
  mv_expect_http create "REST GET list-configurations returns 200" 200 "${LIST_CODE}" "cloud/create-list-configurations.json"
  FOUND_CONFIG="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.configurations[]? | select(.name == $NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-configurations.json")"
  if [[ -n "${FOUND_CONFIG}" ]]; then
    mv_record create "List configurations contains created configuration" PASS cloud/create-list-configurations.json "${FOUND_CONFIG}"
  else
    mv_record create "List configurations contains created configuration" FAIL cloud/create-list-configurations.json "configuration ${EXPECTED_NAME} not found"
  fi
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update LicenseManagerConfiguration "${CONFIG_K8S}" 300
SNAP="$(mv_snapshot update LicenseManagerConfiguration "${CONFIG_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update configuration "${CONFIG_URL}" 200

# Verify changed fields
WANT_UPDATED_DISPLAY="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_UPDATED_CURRENT="$(jq -r '.spec.currentBillingInfo.userCountBilling.userCount // empty' "${SNAP}")"
WANT_UPDATED_NEXT="$(jq -r '.spec.nextBillingInfo.userCountBilling.userCount // empty' "${SNAP}")"

if [[ -n "${WANT_UPDATED_DISPLAY}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-configuration.json" '.displayName' "${WANT_UPDATED_DISPLAY}"
fi
if [[ -n "${WANT_UPDATED_CURRENT}" ]]; then
  mv_expect_jq update "cloud currentBillingInfo userCount updated to new spec" "${MV_RUN_DIR}/cloud/update-configuration.json" '(.currentBillingInfo.userCountBilling.userCount | tostring)' "${WANT_UPDATED_CURRENT}"
fi
if [[ -n "${WANT_UPDATED_NEXT}" ]]; then
  mv_expect_jq update "cloud nextBillingInfo userCount updated to new spec" "${MV_RUN_DIR}/cloud/update-configuration.json" '(.nextBillingInfo.userCountBilling.userCount | tostring)' "${WANT_UPDATED_NEXT}"
fi

# Verify unchanged fields
if [[ -n "${WANT_PRODUCT}" ]]; then
  mv_expect_jq update "cloud product unchanged" "${MV_RUN_DIR}/cloud/update-configuration.json" '.product' "${WANT_PRODUCT}"
fi
CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-configuration.json")"
UPDATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/update-configuration.json")"
mv_expect_eq update "createTime preserved across update (in-place update)" "${CREATE_T}" "${UPDATE_T}" cloud/update-configuration.json

# Verify labels preserved across update
while IFS="=" read -r key val; do
  [[ -z "${key}" ]] && continue
  mv_expect_jq update "cloud label '${key}' matches spec after update" "${MV_RUN_DIR}/cloud/update-configuration.json" ".labels[\"${key}\"]" "${val}"
done < <(jq -r '.spec.labels // {} | to_entries[] | "\(.key)=\(.value)"' "${SNAP}")

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update list-configurations GET "${LIST_URL}")"
  mv_expect_http update "REST GET list-configurations returns 200 after update" 200 "${LIST_CODE}" "cloud/update-list-configurations.json"
  UPDATED_DISPLAY="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.configurations[]? | select(.name == $NAME)][0].displayName // empty' "${MV_RUN_DIR}/cloud/update-list-configurations.json")"
  mv_expect_eq update "List configurations reflects updated displayName" "${WANT_UPDATED_DISPLAY}" "${UPDATED_DISPLAY}" cloud/update-list-configurations.json
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady LicenseManagerConfiguration "${CONFIG_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete LicenseManagerConfiguration "${CONFIG_K8S}" 300
mv_wait_http delete config-gone "${CONFIG_URL}" 404

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-configurations GET "${LIST_URL}")"
  mv_expect_http delete "REST GET list-configurations returns 200" 200 "${LIST_CODE}" "cloud/delete-list-configurations.json"
  FOUND_DELETED="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.configurations[]? | select(.name == $NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-list-configurations.json")"
  mv_expect_eq delete "List configurations no longer contains deleted configuration" "" "${FOUND_DELETED}" cloud/delete-list-configurations.json
fi

# Audit log check:
# Reads Cloud Logging activity logs for write calls on the configuration resource.
# License Manager logs lifecycle events to licensemanager.googleapis.com%2Fconfiguration_state
# and licensemanager.googleapis.com%2Fconfiguration_current_billing_info with messages
# ending in 'created', 'updated', and 'deleted'.
expect_config_audit() {
  local phase="$1" label="$2" want="$3" timeout="${4:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="(logName=\"projects/${GCP_PROJECT_ID}/logs/licensemanager.googleapis.com%2Fconfiguration_state\" OR logName=\"projects/${GCP_PROJECT_ID}/logs/licensemanager.googleapis.com%2Fconfiguration_current_billing_info\") AND resource.labels.configuration_id=\"${CONFIG_ID}\" AND timestamp>=\"${MV_STARTED_UTC}\""
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    got="$(jq -r '[.[].jsonPayload.message | split(" ") | last] | join(",")' "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
    [[ "${got}" == "${want}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 15
  done
  echo "filter: ${filter}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.filter.txt"
  jq -r '.[] | "\(.timestamp)  \(.jsonPayload.message // "-")"' \
    "${MV_RUN_DIR}/${out}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.summary.txt" 2>/dev/null || true
  mv_trace "gcloud logging read '${filter}' --project ${GCP_PROJECT_ID} --order=asc --format=json" \
    "$(jq 'length' "${MV_RUN_DIR}/${out}" 2>/dev/null || echo 0) entries: ${got:-none}" "cloud/${phase}-${label}.summary.txt"
  mv_expect_eq "${phase}" "audit log lifecycle calls on ${CONFIG_ID}" "${want}" "${got}" "${out}"
}

expect_config_audit delete audit-config "created,updated,deleted"

# --- Phase: dependency teardown ----------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  # Teardown dependencies in reverse order if any exist
  :
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
