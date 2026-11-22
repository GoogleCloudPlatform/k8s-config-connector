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
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/ces/v1alpha1/cestoolset/cestoolset-maximal}"

mv_init CESToolset
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Render manifests first so names and parameters can be derived for out-of-band checks.
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Helper to read values from rendered YAML manifests.
get_yaml_field() {
  local file="$1" key="$2" default="$3"
  python3 -c "
import yaml, sys
try:
    with open('${file}') as f:
        doc = yaml.safe_load(f)
    val = doc
    for k in '${key}'.split('.'):
        val = val[k]
    print(val)
except Exception:
    print('${default}')
"
}

# Names and IDs derived from the fixture manifests so out-of-band checks never ask KCC.
APP_K8S="$(get_yaml_field "${DEPS}" "metadata.name" "cesapp-${MV_UNIQUE_ID}")"
APP_ID="$(get_yaml_field "${DEPS}" "spec.resourceID" "${APP_K8S}")"
LOCATION="$(get_yaml_field "${DEPS}" "spec.location" "us")"

TOOLSET_K8S="$(get_yaml_field "${CREATE}" "metadata.name" "cestoolset-maximal-${MV_UNIQUE_ID}")"
TOOLSET_ID="$(get_yaml_field "${CREATE}" "spec.resourceID" "${TOOLSET_K8S}")"

CES_BASE="https://ces.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
APP_URL="${CES_BASE}/apps/${APP_ID}"
TOOLSET_URL="${APP_URL}/toolsets/${TOOLSET_ID}"
TOOLSETS_LIST_URL="${APP_URL}/toolsets"
RETRIEVE_TOOLS_URL="${TOOLSET_URL}:retrieveTools"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get cestoolset "${TOOLSET_K8S}" >/dev/null 2>&1 && mv_delete cleanup CESToolset "${TOOLSET_K8S}" 300
    mv_kubectl get cesapp "${APP_K8S}" >/dev/null 2>&1 && mv_delete cleanup CESApp "${APP_K8S}" 300
  fi
  local code
  code="$(mv_rest cleanup leakcheck-toolset GET "${TOOLSET_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (toolset gone)" PASS cloud/cleanup-leakcheck-toolset.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (toolset gone)" FAIL cloud/cleanup-leakcheck-toolset.json \
      "could not verify (HTTP ${code}); check manually: curl ${TOOLSET_URL}"
  else
    mv_record cleanup "no leaked GCP resources (toolset gone)" FAIL cloud/cleanup-leakcheck-toolset.json \
      "HTTP ${code}; clean up manually: curl -X DELETE ${TOOLSET_URL}"
  fi
  local app_code
  app_code="$(mv_rest cleanup leakcheck-app GET "${APP_URL}")"
  if [[ "${app_code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (app gone)" PASS cloud/cleanup-leakcheck-app.json "HTTP 404"
  elif [[ "${app_code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (app gone)" FAIL cloud/cleanup-leakcheck-app.json \
      "could not verify (HTTP ${app_code}); check manually: curl ${APP_URL}"
  else
    mv_record cleanup "no leaked GCP resources (app gone)" FAIL cloud/cleanup-leakcheck-app.json \
      "HTTP ${app_code}; clean up manually: curl -X DELETE ${APP_URL}"
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

# Note on CLI tooling:
# Google Cloud does not provide a 'gcloud ces' command surface.
# Out-of-band verification is performed via direct REST (curl with ADC credentials) and Cloud Audit Logs.
mv_record setup "CLI tooling surface" INFO "" "No gcloud surface for ces.googleapis.com; using REST and Cloud Audit Logs"

# --- Phase: dependencies -----------------------------------------------------
mv_apply deps "${DEPS}"
mv_wait_ready deps CESApp "${APP_K8S}" 300
mv_wait_http deps app "${APP_URL}" 200

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create CESToolset "${TOOLSET_K8S}" 300
SNAP="$(mv_snapshot create CESToolset "${TOOLSET_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

mv_wait_http create toolset "${TOOLSET_URL}" 200

# Verify resource name and spec fields against cloud response
mv_expect_jq create "cloud toolset name ends with resourceID" "${MV_RUN_DIR}/cloud/create-toolset.json" \
  '(.name | endswith("/toolsets/'"${TOOLSET_ID}"'")) | tostring' "true"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    '.displayName' "${WANT_DISPLAY_NAME}"
fi

WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    '.description' "${WANT_DESCRIPTION}"
fi

WANT_EXECUTION_TYPE="$(jq -r '.spec.executionType // empty' "${SNAP}")"
if [[ -n "${WANT_EXECUTION_TYPE}" ]]; then
  # Google Cloud REST API returns enum string or integer (1 for SYNCHRONOUS)
  mv_expect_jq create "cloud executionType matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    '(.executionType | if . == 1 then "SYNCHRONOUS" else . end) | tostring' "${WANT_EXECUTION_TYPE}"
fi

WANT_FAKE_MODE="$(jq -r '(.spec.toolFakeConfig.enableFakeMode // false) | tostring' "${SNAP}")"
mv_expect_jq create "cloud toolFakeConfig.enableFakeMode matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
  '(.toolFakeConfig.enableFakeMode // false) | tostring' "${WANT_FAKE_MODE}"

WANT_PYTHON_CODE="$(jq -r '.spec.toolFakeConfig.codeBlock.pythonCode // empty' "${SNAP}")"
if [[ -n "${WANT_PYTHON_CODE}" ]]; then
  mv_expect_jq create "cloud toolFakeConfig.codeBlock.pythonCode matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    '.toolFakeConfig.codeBlock.pythonCode' "${WANT_PYTHON_CODE}"
fi

WANT_IGNORE_UNKNOWN="$(jq -r '(.spec.openAPIToolset.ignoreUnknownFields // false) | tostring' "${SNAP}")"
mv_expect_jq create "cloud openApiToolset.ignoreUnknownFields matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
  '(.openApiToolset.ignoreUnknownFields // false) | tostring' "${WANT_IGNORE_UNKNOWN}"

WANT_TOKEN="$(jq -r '.spec.openAPIToolset.apiAuthentication.bearerTokenConfig.token // empty' "${SNAP}")"
if [[ -n "${WANT_TOKEN}" ]]; then
  mv_expect_jq create "cloud openApiToolset bearerTokenConfig token matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    '.openApiToolset.apiAuthentication.bearerTokenConfig.token' "${WANT_TOKEN}"
fi

WANT_SCHEMA="$(jq -r '.spec.openAPIToolset.openAPISchema // empty' "${SNAP}")"
if [[ -n "${WANT_SCHEMA}" ]]; then
  mv_expect_jq create "cloud openApiToolset openApiSchema matches spec" "${MV_RUN_DIR}/cloud/create-toolset.json" \
    '.openApiToolset.openApiSchema' "${WANT_SCHEMA}"
fi

# Out-of-band collection check: list toolsets on app
code="$(mv_rest create list-toolsets GET "${TOOLSETS_LIST_URL}")"
mv_expect_http create "REST GET list toolsets returns 200" 200 "${code}" "cloud/create-list-toolsets.json"
mv_expect_jq create "cloud list toolsets includes the created toolset" "${MV_RUN_DIR}/cloud/create-list-toolsets.json" \
  '([.toolsets[]? | .name | endswith("/toolsets/'"${TOOLSET_ID}"'")] | any) | tostring' "true"

# Functional check: retrieve tools defined in the toolset
code="$(mv_rest create retrieve-tools POST "${RETRIEVE_TOOLS_URL}" '{}')"
mv_expect_http create "functional: retrieveTools returns 200" 200 "${code}" "cloud/create-retrieve-tools.json"
mv_expect_jq create "functional: retrieveTools generates tools from OpenAPI schema" "${MV_RUN_DIR}/cloud/create-retrieve-tools.json" \
  '([.tools[]? | select(.openApiTool.name=="listItems" and .openApiTool.description=="List items")] | length > 0) | tostring' "true"

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
# Give the apiserver a moment to bump generation before polling for Ready.
sleep 2
mv_wait_ready update CESToolset "${TOOLSET_K8S}" 300
SNAP="$(mv_snapshot update CESToolset "${TOOLSET_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update toolset "${TOOLSET_URL}" 200

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-toolset.json" \
    '.displayName' "${WANT_DISPLAY_NAME}"
fi

WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-toolset.json" \
    '.description' "${WANT_DESCRIPTION}"
fi

WANT_FAKE_MODE="$(jq -r '(.spec.toolFakeConfig.enableFakeMode // false) | tostring' "${SNAP}")"
mv_expect_jq update "cloud toolFakeConfig.enableFakeMode updated to false" "${MV_RUN_DIR}/cloud/update-toolset.json" \
  '(.toolFakeConfig.enableFakeMode // false) | tostring' "${WANT_FAKE_MODE}"

WANT_IGNORE_UNKNOWN="$(jq -r '(.spec.openAPIToolset.ignoreUnknownFields // false) | tostring' "${SNAP}")"
mv_expect_jq update "cloud openApiToolset.ignoreUnknownFields updated to false" "${MV_RUN_DIR}/cloud/update-toolset.json" \
  '(.openApiToolset.ignoreUnknownFields // false) | tostring' "${WANT_IGNORE_UNKNOWN}"

WANT_TOKEN="$(jq -r '.spec.openAPIToolset.apiAuthentication.bearerTokenConfig.token // empty' "${SNAP}")"
if [[ -n "${WANT_TOKEN}" ]]; then
  mv_expect_jq update "cloud openApiToolset bearerTokenConfig token updated" "${MV_RUN_DIR}/cloud/update-toolset.json" \
    '.openApiToolset.apiAuthentication.bearerTokenConfig.token' "${WANT_TOKEN}"
fi

WANT_SCHEMA="$(jq -r '.spec.openAPIToolset.openAPISchema // empty' "${SNAP}")"
if [[ -n "${WANT_SCHEMA}" ]]; then
  mv_expect_jq update "cloud openApiToolset openApiSchema updated" "${MV_RUN_DIR}/cloud/update-toolset.json" \
    '.openApiToolset.openApiSchema' "${WANT_SCHEMA}"
fi

# Functional check: retrieve tools reflects updated schema
code="$(mv_rest update retrieve-tools POST "${RETRIEVE_TOOLS_URL}" '{}')"
mv_expect_http update "functional: retrieveTools after update returns 200" 200 "${code}" "cloud/update-retrieve-tools.json"
mv_expect_jq update "functional: retrieveTools generates updated tools from OpenAPI schema" "${MV_RUN_DIR}/cloud/update-retrieve-tools.json" \
  '([.tools[]? | select(.openApiTool.name=="listItems" and .openApiTool.description=="List items updated")] | length > 0) | tostring' "true"

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-toolset.json")"
UPDATE_T="$(jq -r '.updateTime // empty' "${MV_RUN_DIR}/cloud/update-toolset.json")"
mv_record update "createTime and updateTime in cloud" INFO cloud/update-toolset.json "created: ${CREATE_T}, updated: ${UPDATE_T}"

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady CESToolset "${TOOLSET_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete CESToolset "${TOOLSET_K8S}" 300
mv_wait_http delete toolset-gone "${TOOLSET_URL}" 404

# Functional check on deleted toolset
code="$(mv_rest delete retrieve-tools POST "${RETRIEVE_TOOLS_URL}" '{}')"
mv_expect_http delete "functional: retrieveTools on deleted toolset returns 404" 404 "${code}" "cloud/delete-retrieve-tools.json"

# List check: toolset no longer listed
code="$(mv_rest delete list-toolsets GET "${TOOLSETS_LIST_URL}")"
mv_expect_http delete "REST GET list toolsets returns 200" 200 "${code}" "cloud/delete-list-toolsets.json"
mv_expect_jq delete "cloud list toolsets does not include deleted toolset" "${MV_RUN_DIR}/cloud/delete-list-toolsets.json" \
  '([.toolsets[]? | .name | endswith("/toolsets/'"${TOOLSET_ID}"'")] | any) | tostring' "false"

# Verify Cloud Audit Logs: exactly CreateToolset, UpdateToolset, DeleteToolset
mv_expect_audit delete audit-toolset "toolsets/${TOOLSET_ID}" \
  "CreateToolset,UpdateToolset,DeleteToolset"

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown CESApp "${APP_K8S}" 300
mv_wait_http teardown app-gone "${APP_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
