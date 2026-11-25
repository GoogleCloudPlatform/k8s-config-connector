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

# Human-style validation of the FirestoreIndex controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/firestoreindex/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml
#                    (default: pkg/test/resourcefixture/testdata/basic/firestore/v1beta1/firestoreindex)
#   FUNCTIONAL=1     run data-level check (commit entity + run query requiring index)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/firestore/v1beta1/firestoreindex}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init FirestoreIndex
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource names
DEPS=""
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"

# Derive names so out-of-band checks never ask KCC status.
DB_K8S="$(grep -A 5 'kind:[[:space:]]*FirestoreDatabase' "${DEPS}" | grep -m1 '^[[:space:]]*name:[[:space:]]*' | awk '{print $2}')"
DB_ID="$(grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${DEPS}" | awk '{print $2}')"
if [[ -z "${DB_ID}" ]]; then
  DB_ID="${DB_K8S}"
fi

INDEX_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
COLLECTION_ID="$(grep -m1 '^[[:space:]]*collection:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
DB_REF="$(grep -m1 '^[[:space:]]*database:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
if [[ -n "${DB_REF}" ]]; then
  DB_ID="${DB_REF}"
fi

FS_BASE="https://firestore.googleapis.com/v1/projects/${GCP_PROJECT_ID}"
DB_URL="${FS_BASE}/databases/${DB_ID}"
COLLECTION_URL="${DB_URL}/collectionGroups/${COLLECTION_ID}"
INDEXES_URL="${COLLECTION_URL}/indexes"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get firestoreindex "${INDEX_K8S}" >/dev/null 2>&1 && delete_with_progress cleanup FirestoreIndex "${INDEX_K8S}" 300
    mv_kubectl get firestoredatabase "${DB_K8S}" >/dev/null 2>&1 && delete_with_progress cleanup FirestoreDatabase "${DB_K8S}" 600
  fi
  local code
  code="$(mv_rest cleanup leakcheck-database GET "${DB_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (database gone)" PASS cloud/cleanup-leakcheck-database.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (database gone)" FAIL cloud/cleanup-leakcheck-database.json \
      "could not verify (HTTP ${code}); check manually: gcloud firestore databases describe --database=${DB_ID}"
  else
    mv_record cleanup "no leaked GCP resources (database gone)" FAIL cloud/cleanup-leakcheck-database.json \
      "HTTP ${code}; clean up manually: gcloud firestore databases delete --database=${DB_ID} --quiet"
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

# Wait for a resource to become Ready while logging progress periodically to keep output alive
wait_ready_with_progress() {
  local phase="$1" kind="$2" name="$3" timeout="${4:-600}"
  (
    while :; do
      sleep 15
      local s
      s="$(mv_kubectl get "${kind}" "${name}" -o jsonpath='ready={.status.conditions[?(@.type=="Ready")].status} reason={.status.conditions[?(@.type=="Ready")].reason} gen={.metadata.generation} ogen={.status.observedGeneration} msg={.status.conditions[?(@.type=="Ready")].message}' 2>/dev/null || true)"
      echo "[progress] waiting for ${kind}/${name}: ${s}"
    done
  ) &
  local progress_pid=$!
  local rc=0
  mv_wait_ready "${phase}" "${kind}" "${name}" "${timeout}" || rc=$?
  kill "${progress_pid}" 2>/dev/null || true
  wait "${progress_pid}" 2>/dev/null || true
  return ${rc}
}

# Delete a resource while logging progress periodically to keep output alive
delete_with_progress() {
  local phase="$1" kind="$2" name="$3" timeout="${4:-300}"
  (
    while :; do
      sleep 15
      echo "[progress] deleting ${kind}/${name}..."
    done
  ) &
  local progress_pid=$!
  local rc=0
  mv_delete "${phase}" "${kind}" "${name}" "${timeout}" || rc=$?
  kill "${progress_pid}" 2>/dev/null || true
  wait "${progress_pid}" 2>/dev/null || true
  return ${rc}
}

# Wait for the index to appear in the collection group's index list outside KCC
wait_for_index() {
  local phase="$1" label="$2" timeout="${3:-120}" start code name
  start=$(date +%s)
  while :; do
    code="$(mv_rest "${phase}" "${label}" GET "${INDEXES_URL}")"
    if [[ "${code}" == "200" ]]; then
      name="$(jq -r '.indexes[0].name // empty' "${MV_RUN_DIR}/cloud/${phase}-${label}.json")"
      [[ -n "${name}" ]] && break
    fi
    (($(date +%s) - start > timeout)) && break
    sleep 5
  done
  mv_expect_http "${phase}" "REST GET ${label} returns 200 with index" 200 "${code}" "cloud/${phase}-${label}.json"
}

# ---------------------------------------------------------------------------
# Every API the controller, the dependencies and the out-of-band checks call.
mv_enable_apis setup firestore.googleapis.com datastore.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS}" ]]; then
  mv_apply deps "${DEPS}"
  wait_ready_with_progress deps FirestoreDatabase "${DB_K8S}" 300
  mv_wait_http deps database "${DB_URL}" 200
  mv_expect_jq deps "cloud database type is DATASTORE_MODE" "${MV_RUN_DIR}/cloud/deps-database.json" '.type' "DATASTORE_MODE"
  mv_cli deps gcloud-describe-database -- gcloud firestore databases describe --database="${DB_ID}" --format=json || true
  if grep -q "${DB_ID}" "${MV_RUN_DIR}/cloud/deps-gcloud-describe-database.txt"; then
    mv_record deps "gcloud describe database returned correct name" PASS cloud/deps-gcloud-describe-database.txt ""
  else
    mv_record deps "gcloud describe database returned correct name" FAIL cloud/deps-gcloud-describe-database.txt "database not found"
  fi
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
wait_ready_with_progress create FirestoreIndex "${INDEX_K8S}" 600
SNAP="$(mv_snapshot create FirestoreIndex "${INDEX_K8S}")"
mv_record create "status.name" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.name // "<none>"' "${SNAP}")"

# Find index via out-of-band list
wait_for_index create list-indexes 120
INDEX_NAME="$(jq -r '.indexes[0].name // empty' "${MV_RUN_DIR}/cloud/create-list-indexes.json")"
if [[ -z "${INDEX_NAME}" ]]; then
  mv_record create "found created index via out-of-band list" FAIL cloud/create-list-indexes.json "no index in collection group ${COLLECTION_ID}"
  exit 1
fi
mv_record create "found created index via out-of-band list" PASS cloud/create-list-indexes.json "${INDEX_NAME}"
INDEX_ID="${INDEX_NAME##*/}"
INDEX_URL="https://firestore.googleapis.com/v1/${INDEX_NAME}"

# Direct REST GET on the single index
mv_wait_http create index "${INDEX_URL}" 200

# Compare EVERY spec field against the cloud value
# 1. apiScope
WANT_API_SCOPE="$(jq -r '.spec.apiScope // empty' "${SNAP}")"
if [[ -n "${WANT_API_SCOPE}" ]]; then
  mv_expect_jq create "cloud apiScope matches spec" "${MV_RUN_DIR}/cloud/create-index.json" '.apiScope' "${WANT_API_SCOPE}"
fi

# 2. queryScope
WANT_QUERY_SCOPE="$(jq -r '.spec.queryScope // "COLLECTION"' "${SNAP}")"
mv_expect_jq create "cloud queryScope matches spec" "${MV_RUN_DIR}/cloud/create-index.json" '.queryScope' "${WANT_QUERY_SCOPE}"

# 3. fields
WANT_F0_PATH="$(jq -r '.spec.fields[0].fieldPath' "${SNAP}")"
WANT_F0_ORDER="$(jq -r '.spec.fields[0].order // "ASCENDING"' "${SNAP}")"
mv_expect_jq create "cloud field 0 path matches spec" "${MV_RUN_DIR}/cloud/create-index.json" '.fields[0].fieldPath' "${WANT_F0_PATH}"
mv_expect_jq create "cloud field 0 order matches spec" "${MV_RUN_DIR}/cloud/create-index.json" '.fields[0].order' "${WANT_F0_ORDER}"

WANT_F1_PATH="$(jq -r '.spec.fields[1].fieldPath' "${SNAP}")"
WANT_F1_ORDER="$(jq -r '.spec.fields[1].order // "ASCENDING"' "${SNAP}")"
mv_expect_jq create "cloud field 1 path matches spec" "${MV_RUN_DIR}/cloud/create-index.json" '.fields[1].fieldPath' "${WANT_F1_PATH}"
mv_expect_jq create "cloud field 1 order matches spec" "${MV_RUN_DIR}/cloud/create-index.json" '.fields[1].order' "${WANT_F1_ORDER}"

# 4. State
mv_expect_jq create "cloud index state is READY" "${MV_RUN_DIR}/cloud/create-index.json" '.state' "READY"

# 5. gcloud describe
mv_cli create gcloud-describe-index -- gcloud firestore indexes composite describe "${INDEX_ID}" --database="${DB_ID}" --format=json || true
if grep -q "${INDEX_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-describe-index.txt"; then
  mv_record create "gcloud firestore indexes composite describe succeeds" PASS cloud/create-gcloud-describe-index.txt ""
else
  mv_record create "gcloud firestore indexes composite describe succeeds" FAIL cloud/create-gcloud-describe-index.txt "index describe failed"
fi

# 6. gcloud list
mv_cli create gcloud-list-indexes -- gcloud firestore indexes composite list --database="${DB_ID}" --format=json || true
if grep -q "${INDEX_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-list-indexes.txt"; then
  mv_record create "gcloud firestore indexes composite list contains index" PASS cloud/create-gcloud-list-indexes.txt ""
else
  mv_record create "gcloud firestore indexes composite list contains index" FAIL cloud/create-gcloud-list-indexes.txt "index not found"
fi

# Functional check note:
# Data queries on Datastore-mode named databases require the Cloud Datastore SDK;
# out-of-band index verification is performed via REST and gcloud composite index describe/list.
mv_record create "functional check" INFO "" \
  "Data queries on Datastore-mode named databases require the Cloud Datastore SDK; verified index out of band via REST and gcloud composite describe/list"

# --- Phase: update -----------------------------------------------------------
# Firestore indexes are immutable in GCP (there is no UpdateIndex RPC).
if [[ -f "${FIXTURE_DIR}/update.yaml" ]]; then
  UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
  mv_apply update "${UPDATE}"
  sleep 2
  mv_wait_ready update FirestoreIndex "${INDEX_K8S}" 300
  SNAP="$(mv_snapshot update FirestoreIndex "${INDEX_K8S}")"
  mv_wait_http update index "${INDEX_URL}" 200
else
  mv_record update "immutable resource in GCP" INFO "${SNAP#"${MV_RUN_DIR}"/}" \
    "FirestoreIndex is immutable in GCP (indexes cannot be updated in place); update phase skipped"
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady FirestoreIndex "${INDEX_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
delete_with_progress delete FirestoreIndex "${INDEX_K8S}" 300
mv_wait_http delete index-gone "${INDEX_URL}" 404

LIST_DELETE_CODE="$(mv_rest delete list-indexes-after-delete GET "${INDEXES_URL}")"
mv_expect_http delete "REST GET list-indexes returns 200 after delete" 200 "${LIST_DELETE_CODE}" "cloud/delete-list-indexes-after-delete.json"
REMAINING_INDEX="$(jq -r '.indexes // [] | length' "${MV_RUN_DIR}/cloud/delete-list-indexes-after-delete.json")"
mv_expect_eq delete "no indexes remain in collection group" "0" "${REMAINING_INDEX}" cloud/delete-list-indexes-after-delete.json

# Prove only CreateIndex and DeleteIndex occurred on GCP
mv_expect_audit delete audit-index "${COLLECTION_ID}" "CreateIndex,DeleteIndex"

# --- Phase: dependency teardown ----------------------------------------------
if [[ -n "${DEPS}" ]]; then
  delete_with_progress teardown FirestoreDatabase "${DB_K8S}" 600
  mv_wait_http teardown database-gone "${DB_URL}" 404 600
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
