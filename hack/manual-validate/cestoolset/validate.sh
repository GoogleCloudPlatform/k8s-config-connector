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

# Human-style validation of the CESToolset controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/cestoolset/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the cestoolset-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (listing toolsets via REST)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/ces/v1alpha1/cestoolset/cestoolset-maximal}"
if [[ "${FIXTURE_DIR}" != /* ]]; then
  FIXTURE_DIR="${MV_REPO_ROOT}/${FIXTURE_DIR}"
fi
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init CESToolset
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE=""
if [[ -f "${FIXTURE_DIR}/update.yaml" ]]; then
  UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
fi

# Resource identifiers derived from fixtures so out-of-band checks never ask KCC.
APP_K8S="$( (grep -A 5 'kind:[[:space:]]*CESApp' "${DEPS}" | grep -m1 '^[[:space:]]*name:[[:space:]]*' | awk '{print $2}' | tr -d "\"'" ) || true )"
APP_ID="$( (grep -A 10 'kind:[[:space:]]*CESApp' "${DEPS}" | grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' | awk '{print $2}' | tr -d "\"'" ) || true )"
APP_ID="${APP_ID:-${APP_K8S}}"
LOCATION="$( (grep -A 10 'kind:[[:space:]]*CESApp' "${DEPS}" | grep -m1 '^[[:space:]]*location:[[:space:]]*' | awk '{print $2}' | tr -d "\"'" ) || true )"
LOCATION="${LOCATION:-us}"

TOOLSET_K8S="$( (grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}' | tr -d "\"'" ) || true )"
TOOLSET_ID="$( (grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}' | tr -d "\"'" ) || true )"
TOOLSET_ID="${TOOLSET_ID:-${TOOLSET_K8S}}"

CES_BASE="https://ces.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
APP_URL="${CES_BASE}/apps/${APP_ID}"
TOOLSET_URL="${APP_URL}/toolsets/${TOOLSET_ID}"
TOOLSET_LIST_URL="${APP_URL}/toolsets"
EXPECTED_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/apps/${APP_ID}/toolsets/${TOOLSET_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get cestoolset "${TOOLSET_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup CESToolset "${TOOLSET_K8S}" 300
    mv_kubectl get cesapp "${APP_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup CESApp "${APP_K8S}" 300
  fi

  if [[ -n "${TOOLSET_URL:-}" ]]; then
    local code
    code="$(mv_rest cleanup leakcheck-toolset GET "${TOOLSET_URL}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (toolset gone)" PASS cloud/cleanup-leakcheck-toolset.json "HTTP 404"
    elif [[ "${code}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (toolset gone)" FAIL cloud/cleanup-leakcheck-toolset.json \
        "could not verify (HTTP ${code}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${TOOLSET_URL}"
    else
      mv_record cleanup "no leaked GCP resources (toolset gone)" FAIL cloud/cleanup-leakcheck-toolset.json \
        "HTTP ${code}; clean up manually: DELETE ${TOOLSET_URL}"
    fi
  fi

  if [[ -n "${APP_URL:-}" ]]; then
    local acode
    acode="$(mv_rest cleanup leakcheck-app GET "${APP_URL}")"
    if [[ "${acode}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (app gone)" PASS cloud/cleanup-leakcheck-app.json "HTTP 404"
    elif [[ "${acode}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (app gone)" FAIL cloud/cleanup-leakcheck-app.json \
        "could not verify (HTTP ${acode}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${APP_URL}"
    else
      mv_record cleanup "no leaked GCP resources (app gone)" FAIL cloud/cleanup-leakcheck-app.json \
        "HTTP ${acode}; clean up manually: DELETE ${APP_URL}"
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
mv_enable_apis setup ces.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
mv_apply deps "${DEPS}"
mv_wait_ready deps CESApp "${APP_K8S}" 300
mv_wait_http deps app "${APP_URL}" 200

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create CESToolset "${TOOLSET_K8S}" 300
SNAP="$(mv_snapshot create CESToolset "${TOOLSET_K8S}")"
STATUS_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${STATUS_REF:-<none>}"

if [[ -n "${STATUS_REF}" ]]; then
  mv_expect_eq create "status.externalRef matches cloud resource name" "${EXPECTED_NAME}" "${STATUS_REF}" "${SNAP#"${MV_RUN_DIR}"/}"
fi

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_EXECUTION_TYPE="$(jq -r '.spec.executionType // empty' "${SNAP}")"
WANT_ENABLE_FAKE_MODE="$(jq -r 'if .spec.toolFakeConfig?.enableFakeMode != null then .spec.toolFakeConfig.enableFakeMode | tostring else empty end' "${SNAP}")"
WANT_PYTHON_CODE="$(jq -r '.spec.toolFakeConfig?.codeBlock?.pythonCode // empty' "${SNAP}")"
WANT_BEARER_TOKEN="$(jq -r '.spec.openAPIToolset?.apiAuthentication?.bearerTokenConfig?.token // empty' "${SNAP}")"
WANT_IGNORE_UNKNOWN_FIELDS="$(jq -r 'if .spec.openAPIToolset?.ignoreUnknownFields != null then .spec.openAPIToolset.ignoreUnknownFields | tostring else empty end' "${SNAP}")"
WANT_OPENAPI_SCHEMA="$(jq -r '.spec.openAPIToolset?.openAPISchema // empty' "${SNAP}")"

mv_wait_http create toolset "${TOOLSET_URL}" 200

# Out-of-band tool availability
mv_record create "gcloud surface for CES" INFO "" "no gcloud CLI for Customer Engagement Suite (ces.googleapis.com); using raw REST"

# Compare spec fields against cloud response
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" '.description' "${WANT_DESCRIPTION}"
fi
if [[ -n "${WANT_EXECUTION_TYPE}" ]]; then
  mv_expect_jq create "cloud executionType matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    'if .executionType == 1 or .executionType == "SYNCHRONOUS" then "SYNCHRONOUS" elif .executionType == 2 or .executionType == "ASYNCHRONOUS" then "ASYNCHRONOUS" else (.executionType // "" | tostring) end' \
    "${WANT_EXECUTION_TYPE}"
fi
if [[ -n "${WANT_ENABLE_FAKE_MODE}" ]]; then
  mv_expect_jq create "cloud toolFakeConfig.enableFakeMode matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    'if .toolFakeConfig?.enableFakeMode == true then "true" else "false" end' \
    "${WANT_ENABLE_FAKE_MODE}"
fi
if [[ -n "${WANT_PYTHON_CODE}" ]]; then
  mv_expect_jq create "cloud toolFakeConfig pythonCode matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    '.toolFakeConfig?.codeBlock?.pythonCode // ""' \
    "${WANT_PYTHON_CODE}"
fi
if [[ -n "${WANT_BEARER_TOKEN}" ]]; then
  mv_expect_jq create "cloud openApiToolset bearerToken matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    '.openApiToolset?.apiAuthentication?.bearerTokenConfig?.token // ""' \
    "${WANT_BEARER_TOKEN}"
fi
if [[ -n "${WANT_IGNORE_UNKNOWN_FIELDS}" ]]; then
  mv_expect_jq create "cloud openApiToolset.ignoreUnknownFields matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    'if .openApiToolset?.ignoreUnknownFields == true then "true" else "false" end' \
    "${WANT_IGNORE_UNKNOWN_FIELDS}"
fi
if [[ -n "${WANT_OPENAPI_SCHEMA}" ]]; then
  mv_expect_jq create "cloud openApiToolset openApiSchema matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    '.openApiToolset?.openApiSchema // ""' \
    "${WANT_OPENAPI_SCHEMA}"
fi

# Functional check: listing toolsets in the parent CESApp proves toolset is queryable
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create list-toolsets GET "${TOOLSET_LIST_URL}")"
  mv_expect_http create "REST GET list-toolsets returns 200" 200 "${LIST_CODE}" "cloud/create-list-toolsets.json"
  FOUND_NAME="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.toolsets[]? | select(.name == $NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-toolsets.json")"
  mv_expect_eq create "ListToolsets contains created toolset" "${EXPECTED_NAME}" "${FOUND_NAME}" cloud/create-list-toolsets.json
fi

# --- Phase: update -----------------------------------------------------------
if [[ -n "${UPDATE}" ]]; then
  mv_apply update "${UPDATE}"
  sleep 2
  mv_wait_ready update CESToolset "${TOOLSET_K8S}" 300
  SNAP_UP="$(mv_snapshot update CESToolset "${TOOLSET_K8S}")"
  mv_expect_jq update "k8s generation advanced to 2" "${SNAP_UP}" '.metadata.generation' "2"

  WANT_UP_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP_UP}")"
  WANT_UP_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP_UP}")"
  WANT_UP_ENABLE_FAKE_MODE="$(jq -r 'if .spec.toolFakeConfig?.enableFakeMode != null then .spec.toolFakeConfig.enableFakeMode | tostring else empty end' "${SNAP_UP}")"
  WANT_UP_BEARER_TOKEN="$(jq -r '.spec.openAPIToolset?.apiAuthentication?.bearerTokenConfig?.token // empty' "${SNAP_UP}")"
  WANT_UP_IGNORE_UNKNOWN_FIELDS="$(jq -r 'if .spec.openAPIToolset?.ignoreUnknownFields != null then .spec.openAPIToolset.ignoreUnknownFields | tostring else empty end' "${SNAP_UP}")"
  WANT_UP_OPENAPI_SCHEMA="$(jq -r '.spec.openAPIToolset?.openAPISchema // empty' "${SNAP_UP}")"

  mv_wait_http update toolset "${TOOLSET_URL}" 200

  # Verify mutated fields in cloud response
  if [[ -n "${WANT_UP_DISPLAY_NAME}" ]]; then
    mv_expect_jq update "cloud displayName updated" "${MV_RUN_DIR}/cloud/update-toolset.json" '.displayName' "${WANT_UP_DISPLAY_NAME}"
  fi
  if [[ -n "${WANT_UP_DESCRIPTION}" ]]; then
    mv_expect_jq update "cloud description updated" "${MV_RUN_DIR}/cloud/update-toolset.json" '.description' "${WANT_UP_DESCRIPTION}"
  fi
  if [[ -n "${WANT_UP_ENABLE_FAKE_MODE}" ]]; then
    mv_expect_jq update "cloud toolFakeConfig.enableFakeMode updated" "${MV_RUN_DIR}/cloud/update-toolset.json" \
      'if .toolFakeConfig?.enableFakeMode == true then "true" else "false" end' \
      "${WANT_UP_ENABLE_FAKE_MODE}"
  fi
  if [[ -n "${WANT_UP_BEARER_TOKEN}" ]]; then
    mv_expect_jq update "cloud openApiToolset bearerToken updated" "${MV_RUN_DIR}/cloud/update-toolset.json" \
      '.openApiToolset?.apiAuthentication?.bearerTokenConfig?.token // ""' \
      "${WANT_UP_BEARER_TOKEN}"
  fi
  if [[ -n "${WANT_UP_IGNORE_UNKNOWN_FIELDS}" ]]; then
    mv_expect_jq update "cloud openApiToolset.ignoreUnknownFields updated" "${MV_RUN_DIR}/cloud/update-toolset.json" \
      'if .openApiToolset?.ignoreUnknownFields == true then "true" else "false" end' \
      "${WANT_UP_IGNORE_UNKNOWN_FIELDS}"
  fi
  if [[ -n "${WANT_UP_OPENAPI_SCHEMA}" ]]; then
    mv_expect_jq update "cloud openApiToolset openApiSchema updated" "${MV_RUN_DIR}/cloud/update-toolset.json" \
      '.openApiToolset?.openApiSchema // ""' \
      "${WANT_UP_OPENAPI_SCHEMA}"
  fi

  # Unchanged fields intact
  if [[ -n "${WANT_EXECUTION_TYPE}" ]]; then
    mv_expect_jq update "cloud executionType intact" "${MV_RUN_DIR}/cloud/update-toolset.json" \
      'if .executionType == 1 or .executionType == "SYNCHRONOUS" then "SYNCHRONOUS" elif .executionType == 2 or .executionType == "ASYNCHRONOUS" then "ASYNCHRONOUS" else (.executionType // "" | tostring) end' \
      "${WANT_EXECUTION_TYPE}"
  fi
  if [[ -n "${WANT_PYTHON_CODE}" ]]; then
    mv_expect_jq update "cloud toolFakeConfig pythonCode intact" "${MV_RUN_DIR}/cloud/update-toolset.json" \
      '.toolFakeConfig?.codeBlock?.pythonCode // ""' \
      "${WANT_PYTHON_CODE}"
  fi

  CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-toolset.json")"
  UPDATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/update-toolset.json")"
  mv_expect_eq update "createTime preserved across update (in-place update)" "${CREATE_T}" "${UPDATE_T}" cloud/update-toolset.json

  CREATE_UT="$(jq -r '.updateTime // empty' "${MV_RUN_DIR}/cloud/create-toolset.json")"
  UPDATE_UT="$(jq -r '.updateTime // empty' "${MV_RUN_DIR}/cloud/update-toolset.json")"
  mv_record update "updateTime before/after update" INFO cloud/update-toolset.json "${CREATE_UT} -> ${UPDATE_UT}"

  # Functional check: list toolsets reflects updated displayName
  if [[ "${FUNCTIONAL}" == "1" ]]; then
    LIST_CODE="$(mv_rest update list-toolsets GET "${TOOLSET_LIST_URL}")"
    mv_expect_http update "REST GET list-toolsets returns 200 after update" 200 "${LIST_CODE}" "cloud/update-list-toolsets.json"
    if [[ -n "${WANT_UP_DISPLAY_NAME}" ]]; then
      UPDATED_DISPLAY="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.toolsets[]? | select(.name == $NAME)][0].displayName // empty' "${MV_RUN_DIR}/cloud/update-list-toolsets.json")"
      mv_expect_eq update "List toolsets reflects updated displayName" "${WANT_UP_DISPLAY_NAME}" "${UPDATED_DISPLAY}" cloud/update-list-toolsets.json
    fi
  fi
else
  mv_record update "update manifest" INFO "" "no update.yaml in fixture; update phase skipped"
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP_STEADY="$(mv_snapshot steady CESToolset "${TOOLSET_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP_STEADY}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete CESToolset "${TOOLSET_K8S}" 300
mv_wait_http delete toolset-gone "${TOOLSET_URL}" 404

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-toolsets GET "${TOOLSET_LIST_URL}")"
  mv_expect_http delete "REST GET list-toolsets returns 200" 200 "${LIST_CODE}" "cloud/delete-list-toolsets.json"
  FOUND_DELETED="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.toolsets[]? | select(.name == $NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-list-toolsets.json")"
  mv_expect_eq delete "List toolsets no longer contains deleted toolset" "" "${FOUND_DELETED}" cloud/delete-list-toolsets.json
fi

# Out-of-band verification after delete:
# GCP CES Toolsets API does not emit Cloud Audit Log (Admin Activity) entries for
# Toolset mutations (only for top-level Apps).
# Deletion is proven out of band via:
# 1. Direct REST GET returning 404 (verified above).
# 2. ListToolsets no longer containing the deleted toolset (verified above).
mv_record delete "audit log write calls" INFO "cloud/delete-toolset-gone.json" \
  "CES Toolsets API does not emit Cloud Audit Logs; verified deletion via REST GET 404 and ListToolsets"

# --- Phase: teardown ---------------------------------------------------------
mv_delete teardown CESApp "${APP_K8S}" 300
mv_wait_http teardown app-gone "${APP_URL}" 404

if grep -q 'panic:' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log has no panics" FAIL manager.log "panic detected in controller manager"
else
  mv_record teardown "manager log has no panics" PASS manager.log "no panics detected"
fi
