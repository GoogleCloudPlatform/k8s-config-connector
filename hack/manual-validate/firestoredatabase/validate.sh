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

# Human-style validation of the FirestoreDatabase controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/firestoredatabase/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the firestoredatabase-full test fixture)
#   FUNCTIONAL=0     skip data-level document/entity checks
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/firestore/v1beta1/firestoredatabase/firestoredatabase-full}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init FirestoreDatabase
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixture so out-of-band checks never ask KCC.
DB_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
DB_ID="$(grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
DB_ID="${DB_ID:-${DB_K8S}}"

API_BASE="https://firestore.googleapis.com/v1"
DB_URL="${API_BASE}/projects/${GCP_PROJECT_ID}/databases/${DB_ID}"
DOC_URL="${DB_URL}/documents/kccvalidation/doc1"

DELETE_HEARTBEAT_PID=""

cleanup() {
  local rc=$?
  set +e
  kill "${DELETE_HEARTBEAT_PID:-}" 2>/dev/null || true
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"

  # Delete test document if created (best effort)
  if [[ "${FUNCTIONAL}" == "1" ]]; then
    mv_rest cleanup rm-doc DELETE "${DOC_URL}" >/dev/null 2>&1 || true
  fi

  # In case delete protection was left enabled, disable it so deletion succeeds.
  mv_rest cleanup disable-delete-protection PATCH "${DB_URL}?updateMask=deleteProtectionState" \
    '{"name":"projects/'"${GCP_PROJECT_ID}"'/databases/'"${DB_ID}"'","deleteProtectionState":"DELETE_PROTECTION_DISABLED"}' >/dev/null 2>&1 || true

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get firestoredatabase "${DB_K8S}" >/dev/null 2>&1 && mv_delete cleanup FirestoreDatabase "${DB_K8S}" 300
    if [[ -n "${DEPS:-}" ]]; then
      # Delete dependencies in reverse order if any exist
      :
    fi
  fi

  local code
  code="$(mv_rest cleanup leakcheck-db GET "${DB_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (database gone)" PASS cloud/cleanup-leakcheck-db.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (database gone)" FAIL cloud/cleanup-leakcheck-db.json \
      "could not verify (HTTP ${code}); check manually: gcloud firestore databases describe --database=${DB_ID} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP resources (database gone)" FAIL cloud/cleanup-leakcheck-db.json \
      "HTTP ${code}; clean up manually: gcloud firestore databases delete --database=${DB_ID} --project=${GCP_PROJECT_ID} --quiet"
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
mv_enable_apis setup firestore.googleapis.com datastore.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create FirestoreDatabase "${DB_K8S}" 300
SNAP="$(mv_snapshot create FirestoreDatabase "${DB_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

mv_wait_http create db "${DB_URL}" 200

# Compare spec fields against cloud response
WANT_LOCATION="$(jq -r '.spec.locationID' "${SNAP}")"
mv_expect_jq create "cloud locationId matches spec" "${MV_RUN_DIR}/cloud/create-db.json" '.locationId' "${WANT_LOCATION}"

WANT_TYPE="$(jq -r '.spec.type // "FIRESTORE_NATIVE"' "${SNAP}")"
mv_expect_jq create "cloud database type matches spec" "${MV_RUN_DIR}/cloud/create-db.json" '.type' "${WANT_TYPE}"

WANT_CONCURRENCY="$(jq -r '.spec.concurrencyMode // empty' "${SNAP}")"
if [[ -n "${WANT_CONCURRENCY}" ]]; then
  mv_expect_jq create "cloud concurrencyMode matches spec" "${MV_RUN_DIR}/cloud/create-db.json" '.concurrencyMode' "${WANT_CONCURRENCY}"
fi

WANT_PITR="$(jq -r '.spec.pointInTimeRecoveryEnablement // empty' "${SNAP}")"
if [[ -n "${WANT_PITR}" ]]; then
  mv_expect_jq create "cloud pointInTimeRecoveryEnablement matches spec" "${MV_RUN_DIR}/cloud/create-db.json" '.pointInTimeRecoveryEnablement' "${WANT_PITR}"
fi

WANT_DELETE_PROTECTION="$(jq -r '.spec.deleteProtectionState // empty' "${SNAP}")"
if [[ -n "${WANT_DELETE_PROTECTION}" ]]; then
  mv_expect_jq create "cloud deleteProtectionState matches spec" "${MV_RUN_DIR}/cloud/create-db.json" '.deleteProtectionState' "${WANT_DELETE_PROTECTION}"
fi

# Describe via gcloud CLI out of band
mv_cli create gcloud-describe -- gcloud firestore databases describe --database="${DB_ID}" --project="${GCP_PROJECT_ID}" --format=json && \
  mv_record create "gcloud firestore databases describe succeeds" PASS cloud/create-gcloud-describe.txt "" || \
  mv_record create "gcloud firestore databases describe succeeds" FAIL cloud/create-gcloud-describe.txt "gcloud describe failed"

mv_cli create gcloud-list -- gcloud firestore databases list --project="${GCP_PROJECT_ID}" --format=json
if grep -q "databases/${DB_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-list.txt"; then
  mv_record create "gcloud list databases includes database" PASS cloud/create-gcloud-list.txt ""
else
  mv_record create "gcloud list databases includes database" FAIL cloud/create-gcloud-list.txt "database not in list output"
fi

# Functional check: prove the database can store data
if [[ "${FUNCTIONAL}" == "1" ]]; then
  if [[ "${WANT_TYPE}" == "FIRESTORE_NATIVE" ]]; then
    code="$(mv_rest create write-doc PATCH "${DOC_URL}" '{"fields":{"greeting":{"stringValue":"hello-kcc"}}}')"
    mv_expect_http create "functional check: write document in native mode" 200 "${code}" "cloud/create-write-doc.json"
    code="$(mv_rest create read-doc GET "${DOC_URL}")"
    mv_expect_http create "functional check: read document in native mode" 200 "${code}" "cloud/create-read-doc.json"
    mv_expect_jq create "functional check: document field matches" "${MV_RUN_DIR}/cloud/create-read-doc.json" '.fields.greeting.stringValue' "hello-kcc"
  elif [[ "${WANT_TYPE}" == "DATASTORE_MODE" ]]; then
    code="$(mv_rest create allocate-id POST "https://datastore.googleapis.com/v1/projects/${GCP_PROJECT_ID}:allocateIds" '{"keys":[{"partitionId":{"databaseId":"'"${DB_ID}"'"},"path":[{"kind":"KccValidation"}]}]}')"
    mv_expect_http create "functional check: allocate key in datastore mode" 200 "${code}" "cloud/create-allocate-id.json"
    mv_expect_jq create "functional check: allocated key has correct databaseId" "${MV_RUN_DIR}/cloud/create-allocate-id.json" '.keys[0].partitionId.databaseId' "${DB_ID}"
  fi
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
# Give the apiserver a moment to bump generation before polling for Ready.
sleep 2
mv_wait_ready update FirestoreDatabase "${DB_K8S}" 300
SNAP="$(mv_snapshot update FirestoreDatabase "${DB_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update db "${DB_URL}" 200

# Check mutable fields that were updated
NEW_CONCURRENCY="$(jq -r '.spec.concurrencyMode // empty' "${SNAP}")"
if [[ -n "${NEW_CONCURRENCY}" ]]; then
  mv_expect_jq update "cloud concurrencyMode updated to new spec" "${MV_RUN_DIR}/cloud/update-db.json" '.concurrencyMode' "${NEW_CONCURRENCY}"
fi

NEW_PITR="$(jq -r '.spec.pointInTimeRecoveryEnablement // empty' "${SNAP}")"
if [[ -n "${NEW_PITR}" ]]; then
  mv_expect_jq update "cloud pointInTimeRecoveryEnablement updated to new spec" "${MV_RUN_DIR}/cloud/update-db.json" '.pointInTimeRecoveryEnablement' "${NEW_PITR}"
fi

NEW_DELETE_PROTECTION="$(jq -r '.spec.deleteProtectionState // empty' "${SNAP}")"
if [[ -n "${NEW_DELETE_PROTECTION}" ]]; then
  mv_expect_jq update "cloud deleteProtectionState updated to new spec" "${MV_RUN_DIR}/cloud/update-db.json" '.deleteProtectionState' "${NEW_DELETE_PROTECTION}"
fi

# Check immutable fields: locationId and type must be preserved (even when type is omitted in update)
mv_expect_jq update "cloud locationId unchanged" "${MV_RUN_DIR}/cloud/update-db.json" '.locationId' "${WANT_LOCATION}"
UPDATE_TYPE="$(jq -r '.spec.type // empty' "${SNAP}")"
EXPECTED_TYPE="${UPDATE_TYPE:-${WANT_TYPE}}"
mv_expect_jq update "cloud database type preserved on update" "${MV_RUN_DIR}/cloud/update-db.json" '.type' "${EXPECTED_TYPE}"

# Verify update was in-place (createTime is unchanged)
CREATE_TIME="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-db.json")"
UPDATE_TIME="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/update-db.json")"
if [[ -n "${CREATE_TIME}" && "${CREATE_TIME}" == "${UPDATE_TIME}" ]]; then
  mv_record update "database updated in place (createTime unchanged)" PASS cloud/update-db.json "${CREATE_TIME}"
else
  mv_record update "database updated in place (createTime unchanged)" FAIL cloud/update-db.json "createTime was ${CREATE_TIME}, now ${UPDATE_TIME}"
fi

if [[ "${FUNCTIONAL}" == "1" ]]; then
  if [[ "${WANT_TYPE}" == "FIRESTORE_NATIVE" ]]; then
    code="$(mv_rest update write-doc PATCH "${DOC_URL}" '{"fields":{"greeting":{"stringValue":"hello-kcc-updated"}}}')"
    mv_expect_http update "functional check: update document in native mode" 200 "${code}" "cloud/update-write-doc.json"
    code="$(mv_rest update read-updated-doc GET "${DOC_URL}")"
    mv_expect_http update "functional check: read updated document in native mode" 200 "${code}" "cloud/update-read-updated-doc.json"
    mv_expect_jq update "functional check: updated document field matches" "${MV_RUN_DIR}/cloud/update-read-updated-doc.json" '.fields.greeting.stringValue' "hello-kcc-updated"
  elif [[ "${WANT_TYPE}" == "DATASTORE_MODE" ]]; then
    code="$(mv_rest update allocate-id-update POST "https://datastore.googleapis.com/v1/projects/${GCP_PROJECT_ID}:allocateIds" '{"keys":[{"partitionId":{"databaseId":"'"${DB_ID}"'"},"path":[{"kind":"KccValidation"}]}]}')"
    mv_expect_http update "functional check: allocate key in datastore mode after update" 200 "${code}" "cloud/update-allocate-id-update.json"
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady FirestoreDatabase "${DB_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
# Background heartbeat so interactive/CLI runners don't time out during long GCP LRO waits.
( while kill -0 $$ 2>/dev/null; do sleep 30; echo "[$(date -u +%Y-%m-%dT%H:%M:%SZ)] [INFO] delete: waiting for database deletion operation..." >&2; done ) &
DELETE_HEARTBEAT_PID=$!

mv_delete delete FirestoreDatabase "${DB_K8S}" 900
kill "${DELETE_HEARTBEAT_PID}" 2>/dev/null || true
wait "${DELETE_HEARTBEAT_PID}" 2>/dev/null || true

mv_wait_http delete db-gone "${DB_URL}" 404

# Cloud Audit Logs check: CreateDatabase, UpdateDatabase, DeleteDatabase
mv_expect_audit delete audit-db "databases/${DB_ID}" \
  "CreateDatabase,UpdateDatabase,DeleteDatabase"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  if [[ "${WANT_TYPE}" == "FIRESTORE_NATIVE" ]]; then
    code="$(mv_rest delete doc-after-delete GET "${DOC_URL}")"
    mv_expect_http delete "functional check: document GET on deleted database returns 404" 404 "${code}" "cloud/delete-doc-after-delete.json"
  elif [[ "${WANT_TYPE}" == "DATASTORE_MODE" ]]; then
    code="$(mv_rest delete allocate-after-delete POST "https://datastore.googleapis.com/v1/projects/${GCP_PROJECT_ID}:allocateIds" '{"keys":[{"partitionId":{"databaseId":"'"${DB_ID}"'"},"path":[{"kind":"KccValidation"}]}]}')"
    mv_expect_http delete "functional check: allocateIds on deleted database returns 404" 404 "${code}" "cloud/delete-allocate-after-delete.json"
  fi
fi

# --- Phase: dependency teardown ----------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  :
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
