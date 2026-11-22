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

# Human-style validation of the MigrationCenterImportDataFile controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/migrationcenterimportdatafile/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml (and optional update.yaml)
#                    (default: the migrationcenterimportdatafile-maximal test fixture)
#   FUNCTIONAL=1     run functional check (verify signedUri by checking URL prefix and list API)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/migrationcenter/v1alpha1/migrationcenterimportdatafile/migrationcenterimportdatafile-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init MigrationCenterImportDataFile
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource names
DEPS=""
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"

# Derive names from the rendered manifests so out-of-band checks never ask KCC.
IDF_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
SOURCE_K8S="$(grep -E '^[[:space:]]*name:[[:space:]]*migrationcentersource-' "${DEPS}" 2>/dev/null | head -1 | awk '{print $2}' || true)"
JOB_K8S="$(grep -E '^[[:space:]]*name:[[:space:]]*migrationcenterimportjob-' "${DEPS}" 2>/dev/null | head -1 | awk '{print $2}' || true)"
LOCATION="us-central1"

MC_BASE="https://migrationcenter.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
SOURCE_URL="${MC_BASE}/sources/${SOURCE_K8S}?%24alt=json%3Benum-encoding%3Dint"
JOB_URL="${MC_BASE}/importJobs/${JOB_K8S}?%24alt=json%3Benum-encoding%3Dint"
IDF_URL="${MC_BASE}/importJobs/${JOB_K8S}/importDataFiles/${IDF_K8S}?%24alt=json%3Benum-encoding%3Dint"
IDF_LIST_URL="${MC_BASE}/importJobs/${JOB_K8S}/importDataFiles?%24alt=json%3Benum-encoding%3Dint"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get migrationcenterimportdatafile "${IDF_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup MigrationCenterImportDataFile "${IDF_K8S}" 300
    if [[ -n "${JOB_K8S}" ]]; then
      mv_kubectl get migrationcenterimportjob "${JOB_K8S}" >/dev/null 2>&1 && \
        mv_delete cleanup MigrationCenterImportJob "${JOB_K8S}" 300
    fi
    if [[ -n "${SOURCE_K8S}" ]]; then
      mv_kubectl get migrationcentersource "${SOURCE_K8S}" >/dev/null 2>&1 && \
        mv_delete cleanup MigrationCenterSource "${SOURCE_K8S}" 300
    fi
  fi

  local code
  code="$(mv_rest cleanup leakcheck-idf GET "${IDF_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (import data file gone)" PASS cloud/cleanup-leakcheck-idf.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (import data file gone)" FAIL cloud/cleanup-leakcheck-idf.json \
      "could not verify (HTTP ${code}); check manually: ${IDF_URL}"
  else
    mv_record cleanup "no leaked GCP resources (import data file gone)" FAIL cloud/cleanup-leakcheck-idf.json \
      "HTTP ${code}; clean up manually: DELETE ${IDF_URL}"
  fi

  if [[ -n "${JOB_K8S}" ]]; then
    local jcode
    jcode="$(mv_rest cleanup leakcheck-job GET "${JOB_URL}")"
    if [[ "${jcode}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (import job gone)" PASS cloud/cleanup-leakcheck-job.json "HTTP 404"
    else
      mv_record cleanup "no leaked GCP resources (import job gone)" FAIL cloud/cleanup-leakcheck-job.json \
        "HTTP ${jcode}; check job: ${JOB_URL}"
    fi
  fi

  if [[ -n "${SOURCE_K8S}" ]]; then
    local scode
    scode="$(mv_rest cleanup leakcheck-source GET "${SOURCE_URL}")"
    if [[ "${scode}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (source gone)" PASS cloud/cleanup-leakcheck-source.json "HTTP 404"
    else
      mv_record cleanup "no leaked GCP resources (source gone)" FAIL cloud/cleanup-leakcheck-source.json \
        "HTTP ${scode}; check source: ${SOURCE_URL}"
    fi
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
# Every API the controller, the dependencies and the out-of-band checks call.
mv_enable_apis setup migrationcenter.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
  if [[ -n "${SOURCE_K8S}" ]]; then
    mv_wait_ready deps MigrationCenterSource "${SOURCE_K8S}" 300
    mv_wait_http deps source "${SOURCE_URL}" 200
  fi
  if [[ -n "${JOB_K8S}" ]]; then
    mv_wait_ready deps MigrationCenterImportJob "${JOB_K8S}" 300
    mv_wait_http deps job "${JOB_URL}" 200
  fi
else
  mv_record deps "dependencies" INFO "" "none required"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create MigrationCenterImportDataFile "${IDF_K8S}" 300
SNAP="$(mv_snapshot create MigrationCenterImportDataFile "${IDF_K8S}")"
STATUS_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${STATUS_REF:-<none>}"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_FORMAT="$(jq -r '.spec.format // empty' "${SNAP}")"

mv_wait_http create idf "${IDF_URL}" 200

# Format mapping from proto enum string to integer or enum string
case "${WANT_FORMAT}" in
  "IMPORT_JOB_FORMAT_RVTOOLS_XLSX")
    WANT_FORMAT_INT="1"
    ;;
  "IMPORT_JOB_FORMAT_RVTOOLS_CSV")
    WANT_FORMAT_INT="2"
    ;;
  "IMPORT_JOB_FORMAT_STRATOZONE_CSV")
    WANT_FORMAT_INT="3"
    ;;
  *)
    WANT_FORMAT_INT="${WANT_FORMAT}"
    ;;
esac

# Compare spec fields against cloud response
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-idf.json" '.displayName // empty' "${WANT_DISPLAY_NAME}"
fi

# Check format (handles both integer and string enum representations)
CLOUD_FORMAT="$(jq -r '.format // empty' "${MV_RUN_DIR}/cloud/create-idf.json")"
if [[ "${CLOUD_FORMAT}" == "${WANT_FORMAT}" || "${CLOUD_FORMAT}" == "${WANT_FORMAT_INT}" ]]; then
  mv_record create "cloud format matches spec" PASS cloud/create-idf.json "format=${CLOUD_FORMAT}"
else
  mv_record create "cloud format matches spec" FAIL cloud/create-idf.json "want '${WANT_FORMAT}' or '${WANT_FORMAT_INT}', got '${CLOUD_FORMAT}'"
fi

# Cloud state should be ACTIVE (2 in int enum)
CLOUD_STATE="$(jq -r '.state // empty' "${MV_RUN_DIR}/cloud/create-idf.json")"
if [[ "${CLOUD_STATE}" == "ACTIVE" || "${CLOUD_STATE}" == "2" ]]; then
  mv_record create "cloud state is ACTIVE" PASS cloud/create-idf.json "state=${CLOUD_STATE}"
else
  mv_record create "cloud state is ACTIVE" FAIL cloud/create-idf.json "got state=${CLOUD_STATE}"
fi

# Cloud name matches expected identity
EXPECTED_CLOUD_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/importJobs/${JOB_K8S}/importDataFiles/${IDF_K8S}"
mv_expect_jq create "cloud name matches expected identity" "${MV_RUN_DIR}/cloud/create-idf.json" '.name' "${EXPECTED_CLOUD_NAME}"

# Record gcloud CLI availability
mv_record create "gcloud CLI surface" INFO "cloud/create-idf.json" "no gcloud CLI command for migration-center import-data-files; REST API used out of band"

# Functional check: ListImportDataFiles contains the created import data file, and signedUri is present
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create verify-list-idfs GET "${IDF_LIST_URL}")"
  mv_expect_http create "REST GET verify-list-idfs returns 200" 200 "${LIST_CODE}" "cloud/create-verify-list-idfs.json"
  FOUND_NAME="$(jq -r --arg NAME "${EXPECTED_CLOUD_NAME}" '[.importDataFiles[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-verify-list-idfs.json")"
  mv_expect_eq create "ListImportDataFiles returns created import data file" "${EXPECTED_CLOUD_NAME}" "${FOUND_NAME}" cloud/create-verify-list-idfs.json

  SIGNED_URI="$(jq -r '.uploadFileInfo.signedUri // empty' "${MV_RUN_DIR}/cloud/create-idf.json")"
  if [[ -n "${SIGNED_URI}" && "${SIGNED_URI}" =~ ^https://storage\.googleapis\.com/ ]]; then
    mv_record create "uploadFileInfo.signedUri generated and valid" PASS cloud/create-idf.json "signedUri prefix valid"
  else
    mv_record create "uploadFileInfo.signedUri generated and valid" FAIL cloud/create-idf.json "signedUri empty or invalid: ${SIGNED_URI}"
  fi
fi

# --- Phase: update (MigrationCenterImportDataFile is immutable in GCP) -------
if [[ -f "${FIXTURE_DIR}/update.yaml" ]]; then
  UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
  mv_apply update "${UPDATE}"
  sleep 2
  mv_wait_ready update MigrationCenterImportDataFile "${IDF_K8S}" 300
  SNAP="$(mv_snapshot update MigrationCenterImportDataFile "${IDF_K8S}")"
  mv_wait_http update idf "${IDF_URL}" 200
else
  mv_record update "immutable resource" INFO "${SNAP#"${MV_RUN_DIR}"/}" "MigrationCenterImportDataFile has no Update RPC in GCP API; update phase skipped"
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady MigrationCenterImportDataFile "${IDF_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete MigrationCenterImportDataFile "${IDF_K8S}" 300
mv_wait_http delete idf-gone "${IDF_URL}" 404

# Out-of-band verification after delete:
# GCP Migration Center API does not emit Cloud Audit Log (Admin Activity) entries for
# ImportDataFile mutations (only for top-level Sources and ImportJobs).
# We verify delete out of band via:
# 1. Direct REST GET returning 404 (verified above via mv_wait_http).
# 2. ListImportDataFiles no longer contains the deleted import data file.
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete verify-list-idfs GET "${IDF_LIST_URL}")"
  mv_expect_http delete "REST GET verify-list-idfs returns 200" 200 "${LIST_CODE}" "cloud/delete-verify-list-idfs.json"
  FOUND_DELETED="$(jq -r --arg NAME "${EXPECTED_CLOUD_NAME}" '[.importDataFiles[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-verify-list-idfs.json")"
  mv_expect_eq delete "ListImportDataFiles no longer contains deleted import data file" "" "${FOUND_DELETED}" cloud/delete-verify-list-idfs.json
fi

# Audit logging surface status
mv_record delete "Cloud Audit Logs surface" INFO "cloud/delete-idf-gone.json" \
  "GCP MigrationCenter API does not emit Admin Activity audit logs for ImportDataFiles sub-resources; verified via REST GET 404 and ListImportDataFiles"

# --- Phase: dependency teardown ----------------------------------------------
if [[ -n "${JOB_K8S}" ]]; then
  mv_delete teardown MigrationCenterImportJob "${JOB_K8S}" 300
  mv_wait_http teardown job-gone "${JOB_URL}" 404
fi

if [[ -n "${SOURCE_K8S}" ]]; then
  mv_delete teardown MigrationCenterSource "${SOURCE_K8S}" 300
  mv_wait_http teardown source-gone "${SOURCE_URL}" 404
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
