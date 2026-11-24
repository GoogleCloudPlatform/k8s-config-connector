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

# Human-style validation of the DiscoveryEngineDataStore controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/discoveryenginedatastore/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the discoveryenginedatastoreminimal test fixture)
#   FUNCTIONAL=0     skip the functional checks (listing dataStores and schemas)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/discoveryengine/v1alpha1/discoveryenginedatastore/discoveryenginedatastoreminimal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DiscoveryEngineDataStore
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
DATASTORE_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
DATASTORE_ID="$(grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
DATASTORE_ID="${DATASTORE_ID:-${DATASTORE_K8S}}"
LOCATION="$(grep -m1 '^[[:space:]]*location:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
LOCATION="${LOCATION:-global}"
COLLECTION="$(grep -m1 '^[[:space:]]*collection:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
COLLECTION="${COLLECTION:-default_collection}"

API_BASE="https://discoveryengine.googleapis.com/v1"
COLLECTION_URL="${API_BASE}/projects/${GCP_PROJECT_ID}/locations/${LOCATION}/collections/${COLLECTION}"
DATASTORE_URL="${COLLECTION_URL}/dataStores/${DATASTORE_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get discoveryenginedatastore "${DATASTORE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DiscoveryEngineDataStore "${DATASTORE_K8S}" 300
    if [[ -n "${DEPS:-}" ]]; then
      # Delete dependencies in reverse order if any exist
      :
    fi
  fi

  local code
  code="$(mv_rest cleanup leakcheck-datastore GET "${DATASTORE_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (dataStore gone)" PASS cloud/cleanup-leakcheck-datastore.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (dataStore gone)" FAIL cloud/cleanup-leakcheck-datastore.json \
      "could not verify (HTTP ${code}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" -H \"x-goog-user-project: ${GCP_PROJECT_ID}\" ${DATASTORE_URL}"
  else
    mv_record cleanup "no leaked GCP resources (dataStore gone)" FAIL cloud/cleanup-leakcheck-datastore.json \
      "HTTP ${code}; clean up manually: curl -X DELETE -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" -H \"x-goog-user-project: ${GCP_PROJECT_ID}\" ${DATASTORE_URL}"
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
mv_enable_apis setup discoveryengine.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create DiscoveryEngineDataStore "${DATASTORE_K8S}" 300
SNAP="$(mv_snapshot create DiscoveryEngineDataStore "${DATASTORE_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

mv_wait_http create datastore "${DATASTORE_URL}" 200

# Compare spec fields against cloud response
WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_INDUSTRY_VERTICAL="$(jq -r '.spec.industryVertical // empty' "${SNAP}")"
WANT_CONTENT_CONFIG="$(jq -r '.spec.contentConfig // empty' "${SNAP}")"
WANT_SOLUTION_TYPES="$(jq -r '[.spec.solutionTypes[]?] | sort | join(",")' "${SNAP}")"

if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-datastore.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_INDUSTRY_VERTICAL}" ]]; then
  mv_expect_jq create "cloud industryVertical matches spec" "${MV_RUN_DIR}/cloud/create-datastore.json" '.industryVertical' "${WANT_INDUSTRY_VERTICAL}"
fi
if [[ -n "${WANT_CONTENT_CONFIG}" ]]; then
  mv_expect_jq create "cloud contentConfig matches spec" "${MV_RUN_DIR}/cloud/create-datastore.json" '.contentConfig' "${WANT_CONTENT_CONFIG}"
fi
if [[ -n "${WANT_SOLUTION_TYPES}" ]]; then
  GOT_SOLUTION_TYPES="$(jq -r '[.solutionTypes[]?] | sort | join(",")' "${MV_RUN_DIR}/cloud/create-datastore.json")"
  mv_expect_eq create "cloud solutionTypes match spec" "${WANT_SOLUTION_TYPES}" "${GOT_SOLUTION_TYPES}" cloud/create-datastore.json
fi

# Resource name in cloud ends with dataStore ID
GOT_NAME="$(jq -r '.name // empty' "${MV_RUN_DIR}/cloud/create-datastore.json")"
EXPECTED_SUFFIX="/collections/${COLLECTION}/dataStores/${DATASTORE_ID}"
if [[ "${GOT_NAME}" == *"${EXPECTED_SUFFIX}" ]]; then
  mv_record create "cloud resource name ends with expected dataStore ID" PASS cloud/create-datastore.json "${GOT_NAME}"
else
  mv_record create "cloud resource name ends with expected dataStore ID" FAIL cloud/create-datastore.json "want *${EXPECTED_SUFFIX}, got ${GOT_NAME}"
fi

# Check default schema ID was created in cloud
DEFAULT_SCHEMA="$(jq -r '.defaultSchemaId // empty' "${MV_RUN_DIR}/cloud/create-datastore.json")"
if [[ -n "${DEFAULT_SCHEMA}" ]]; then
  mv_record create "defaultSchemaId populated in cloud" PASS cloud/create-datastore.json "${DEFAULT_SCHEMA}"
else
  mv_record create "defaultSchemaId populated in cloud" FAIL cloud/create-datastore.json "empty defaultSchemaId"
fi

# Note on gcloud CLI support:
# gcloud does not support DiscoveryEngine DataStores (REST API only).
mv_record create "gcloud CLI surface" INFO env.yaml "no gcloud CLI surface for discoveryengine; verified via REST"

# Functional checks
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create list-datastores GET "${COLLECTION_URL}/dataStores")"
  mv_expect_http create "REST GET list-datastores returns 200" 200 "${LIST_CODE}" "cloud/create-list-datastores.json"
  FOUND_DS="$(jq -r --arg ID "${DATASTORE_ID}" '[.dataStores[]? | select(.name | endswith("/dataStores/" + $ID))][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-datastores.json")"
  if [[ -n "${FOUND_DS}" ]]; then
    mv_record create "List dataStores contains created dataStore" PASS cloud/create-list-datastores.json "${FOUND_DS}"
  else
    mv_record create "List dataStores contains created dataStore" FAIL cloud/create-list-datastores.json "dataStore ${DATASTORE_ID} not found in collection"
  fi

  SCHEMA_CODE="$(mv_rest create list-schemas GET "${DATASTORE_URL}/schemas")"
  mv_expect_http create "REST GET list-schemas returns 200" 200 "${SCHEMA_CODE}" "cloud/create-list-schemas.json"
  FOUND_SCHEMA="$(jq -r '[.schemas[]? | select(.name | endswith("/schemas/default_schema"))][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-schemas.json")"
  if [[ -n "${FOUND_SCHEMA}" ]]; then
    mv_record create "dataStore schemas contains default_schema" PASS cloud/create-list-schemas.json "${FOUND_SCHEMA}"
  else
    mv_record create "dataStore schemas contains default_schema" FAIL cloud/create-list-schemas.json "default_schema not found"
  fi
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DiscoveryEngineDataStore "${DATASTORE_K8S}" 300
SNAP="$(mv_snapshot update DiscoveryEngineDataStore "${DATASTORE_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update datastore "${DATASTORE_URL}" 200

# Verify changed fields
WANT_UPDATED_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
if [[ -n "${WANT_UPDATED_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-datastore.json" '.displayName' "${WANT_UPDATED_DISPLAY_NAME}"
fi

# Verify unchanged fields
if [[ -n "${WANT_INDUSTRY_VERTICAL}" ]]; then
  mv_expect_jq update "cloud industryVertical unchanged" "${MV_RUN_DIR}/cloud/update-datastore.json" '.industryVertical' "${WANT_INDUSTRY_VERTICAL}"
fi

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-datastore.json")"
UPDATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/update-datastore.json")"
mv_expect_eq update "createTime preserved across update (in-place update)" "${CREATE_T}" "${UPDATE_T}" cloud/update-datastore.json

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update list-datastores GET "${COLLECTION_URL}/dataStores")"
  mv_expect_http update "REST GET list-datastores returns 200 after update" 200 "${LIST_CODE}" "cloud/update-list-datastores.json"
  UPDATED_DISPLAY="$(jq -r --arg ID "${DATASTORE_ID}" '[.dataStores[]? | select(.name | endswith("/dataStores/" + $ID))][0].displayName // empty' "${MV_RUN_DIR}/cloud/update-list-datastores.json")"
  mv_expect_eq update "List dataStores reflects updated displayName" "${WANT_UPDATED_DISPLAY_NAME}" "${UPDATED_DISPLAY}" cloud/update-list-datastores.json
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DiscoveryEngineDataStore "${DATASTORE_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DiscoveryEngineDataStore "${DATASTORE_K8S}" 300
mv_wait_http delete datastore-gone "${DATASTORE_URL}" 404

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-datastores GET "${COLLECTION_URL}/dataStores")"
  mv_expect_http delete "REST GET list-datastores returns 200" 200 "${LIST_CODE}" "cloud/delete-list-datastores.json"
  FOUND_DELETED="$(jq -r --arg ID "${DATASTORE_ID}" '[.dataStores[]? | select(.name | endswith("/dataStores/" + $ID))][0].name // empty' "${MV_RUN_DIR}/cloud/delete-list-datastores.json")"
  mv_expect_eq delete "List dataStores no longer contains deleted dataStore" "" "${FOUND_DELETED}" cloud/delete-list-datastores.json
fi

# Audit log check:
# Reads Cloud Audit Logs (admin activity) for write calls on the dataStore resource.
# Note: In Discovery Engine Cloud Audit Logs, CreateDataStore is audited on the parent collection
# (resourceName: projects/.../collections/default_collection), while UpdateDataStore and
# DeleteDataStore are audited on the dataStore itself (resourceName: .../dataStores/${DATASTORE_ID}).
# This check verifies that over the resource lifecycle, exactly one in-place UpdateDataStore
# and one DeleteDataStore occurred on the dataStore, with no recreation or spurious writes.
expect_datastore_audit() {
  local phase="$1" label="$2" want="$3" timeout="${4:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.resourceName:\"${DATASTORE_ID}\" AND timestamp>=\"${MV_STARTED_UTC}\""
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    got="$(jq -r '[.[].protoPayload.methodName | split(".") | last] | join(",")' "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
    [[ "${got}" == "${want}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 15
  done
  echo "filter: ${filter}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.filter.txt"
  jq -r '.[] | "\(.timestamp)  \(.protoPayload.methodName)  \(.protoPayload.authenticationInfo.principalEmail // "-")"' \
    "${MV_RUN_DIR}/${out}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.summary.txt" 2>/dev/null || true
  mv_trace "gcloud logging read '${filter}' --project ${GCP_PROJECT_ID} --order=asc --format=json" \
    "$(jq 'length' "${MV_RUN_DIR}/${out}" 2>/dev/null || echo 0) entries: ${got:-none}" "${out}"
  mv_expect_eq "${phase}" "audit log write calls for ${DATASTORE_ID}" "${want}" "${got}" "${out}"
}

expect_datastore_audit delete audit-datastore "UpdateDataStore,DeleteDataStore"

# --- Phase: teardown ---------------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  # Teardown dependencies in reverse order if any exist
  :
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
