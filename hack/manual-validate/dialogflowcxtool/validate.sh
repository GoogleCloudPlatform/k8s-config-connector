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

# Human-style validation of the DialogflowCXTool controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/dialogflowcxtool/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the dialogflowcxtool-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (listing tools)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/dialogflowcx/v1alpha1/dialogflowcxtool/dialogflowcxtool-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DialogflowCXTool
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

AGENT_K8S="dialogflowcxagent-${MV_UNIQUE_ID}"
TOOL_K8S="dialogflowcxtool-maximal-${MV_UNIQUE_ID}"

TOOL_URL=""
AGENT_URL=""
TOOLS_LIST_URL=""
AGENT_LIST_URL="https://global-dialogflow.googleapis.com/v3/projects/${GCP_PROJECT_ID}/locations/global/agents?alt=json"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get dialogflowcxtool "${TOOL_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup DialogflowCXTool "${TOOL_K8S}" 300
    mv_kubectl get dialogflowcxagent "${AGENT_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup DialogflowCXAgent "${AGENT_K8S}" 300
  fi

  if [[ -n "${TOOL_URL:-}" ]]; then
    local code
    code="$(mv_rest cleanup leakcheck-tool GET "${TOOL_URL}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (tool gone)" PASS cloud/cleanup-leakcheck-tool.json "HTTP 404"
    elif [[ "${code}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (tool gone)" FAIL cloud/cleanup-leakcheck-tool.json \
        "could not verify (HTTP ${code}); check manually: ${TOOL_URL}"
    else
      mv_record cleanup "no leaked GCP resources (tool gone)" FAIL cloud/cleanup-leakcheck-tool.json \
        "HTTP ${code}; clean up manually: DELETE ${TOOL_URL}"
    fi
  fi

  if [[ -n "${AGENT_URL:-}" ]]; then
    local acode
    acode="$(mv_rest cleanup leakcheck-agent GET "${AGENT_URL}")"
    if [[ "${acode}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (agent gone)" PASS cloud/cleanup-leakcheck-agent.json "HTTP 404"
    elif [[ "${acode}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (agent gone)" FAIL cloud/cleanup-leakcheck-agent.json \
        "could not verify (HTTP ${acode}); check manually: ${AGENT_URL}"
    else
      mv_record cleanup "no leaked GCP resources (agent gone)" FAIL cloud/cleanup-leakcheck-agent.json \
        "HTTP ${acode}; clean up manually: DELETE ${AGENT_URL}"
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
mv_enable_apis setup dialogflow.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
AGENT_K8S="$(python3 -c "import yaml; print(yaml.safe_load(open('${DEPS}'))['metadata']['name'])")"
mv_apply deps "${DEPS}"
mv_wait_ready deps DialogflowCXAgent "${AGENT_K8S}" 300
SNAP_AGENT="$(mv_snapshot deps DialogflowCXAgent "${AGENT_K8S}")"
AGENT_NAME="$(jq -r '.status.name // .status.externalRef // empty' "${SNAP_AGENT}")"
AGENT_ID="${AGENT_NAME##*/}"
AGENT_URL="https://global-dialogflow.googleapis.com/v3/projects/${GCP_PROJECT_ID}/locations/global/agents/${AGENT_ID}?alt=json"
TOOLS_LIST_URL="https://global-dialogflow.googleapis.com/v3/projects/${GCP_PROJECT_ID}/locations/global/agents/${AGENT_ID}/tools?%24alt=json%3Benum-encoding%3Dint"
mv_wait_http deps agent "${AGENT_URL}" 200

# Verify out of band that ListAgents returns the created agent
start=$(date +%s)
while :; do
  mv_rest deps list-agents GET "${AGENT_LIST_URL}" >/dev/null
  FOUND_AGENT="$(jq -r --arg dn "agent-${MV_UNIQUE_ID}" '[.agents[]? | select(.displayName==$dn)][0].name // empty' "${MV_RUN_DIR}/cloud/deps-list-agents.json")"
  [[ -n "${FOUND_AGENT}" ]] && break
  (($(date +%s) - start > 60)) && break
  sleep 3
done

if [[ -n "${FOUND_AGENT}" ]]; then
  mv_record deps "ListAgents contains agent matching displayName" PASS cloud/deps-list-agents.json "${FOUND_AGENT}"
else
  mv_record deps "ListAgents contains agent matching displayName" FAIL cloud/deps-list-agents.json "not found in ${AGENT_LIST_URL}"
fi

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
TOOL_K8S="$(python3 -c "import yaml; print(yaml.safe_load(open('${CREATE}'))['metadata']['name'])")"
mv_apply create "${CREATE}"
mv_wait_ready create DialogflowCXTool "${TOOL_K8S}" 300
SNAP="$(mv_snapshot create DialogflowCXTool "${TOOL_K8S}")"
STATUS_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${STATUS_REF:-<none>}"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"

start=$(date +%s)
while :; do
  mv_rest create list-tools GET "${TOOLS_LIST_URL}" >/dev/null
  CLOUD_NAME="$(jq -r --arg dn "${WANT_DISPLAY_NAME}" '[.tools[]? | select(.displayName==$dn)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-tools.json")"
  [[ -n "${CLOUD_NAME}" ]] && break
  (($(date +%s) - start > 60)) && break
  sleep 3
done

if [[ -n "${CLOUD_NAME}" ]]; then
  mv_record create "ListTools contains tool matching displayName" PASS cloud/create-list-tools.json "${CLOUD_NAME}"
  TOOL_ID="${CLOUD_NAME##*/}"
  TOOL_URL="https://global-dialogflow.googleapis.com/v3/${CLOUD_NAME}?%24alt=json%3Benum-encoding%3Dint"
else
  mv_record create "ListTools contains tool matching displayName" FAIL cloud/create-list-tools.json "not found in ${TOOLS_LIST_URL}"
  if [[ -n "${STATUS_REF}" ]]; then
    CLOUD_NAME="${STATUS_REF}"
    TOOL_ID="${CLOUD_NAME##*/}"
    TOOL_URL="https://global-dialogflow.googleapis.com/v3/${CLOUD_NAME}?%24alt=json%3Benum-encoding%3Dint"
  fi
fi

if [[ -n "${STATUS_REF}" && -n "${CLOUD_NAME}" ]]; then
  mv_expect_eq create "status.externalRef matches cloud resource name" "${CLOUD_NAME}" "${STATUS_REF}" "${SNAP#"${MV_RUN_DIR}"/}"
fi

mv_wait_http create tool "${TOOL_URL}" 200

# Compare spec fields against cloud response
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-tool.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi

if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-tool.json" '.description' "${WANT_DESCRIPTION}"
fi

if jq -e '.spec.openAPISpec.textSchema' "${SNAP}" >/dev/null 2>&1; then
  WANT_OPENAPI_TEXT="$(jq -r '.spec.openAPISpec.textSchema' "${SNAP}")"
  mv_expect_jq create "cloud openApiSpec.textSchema matches spec" "${MV_RUN_DIR}/cloud/create-tool.json" '.openApiSpec.textSchema' "${WANT_OPENAPI_TEXT}"
fi

if jq -e '.spec.functionSpec.inputSchema' "${SNAP}" >/dev/null 2>&1; then
  WANT_INPUT_SCHEMA="$(jq -c '.spec.functionSpec.inputSchema' "${SNAP}")"
  GOT_INPUT_SCHEMA="$(jq -c '.functionSpec.inputSchema // empty' "${MV_RUN_DIR}/cloud/create-tool.json")"
  mv_expect_eq create "cloud functionSpec.inputSchema matches spec" "${WANT_INPUT_SCHEMA}" "${GOT_INPUT_SCHEMA}" "cloud/create-tool.json"
fi

if jq -e '.spec.functionSpec.outputSchema' "${SNAP}" >/dev/null 2>&1; then
  WANT_OUTPUT_SCHEMA="$(jq -c '.spec.functionSpec.outputSchema' "${SNAP}")"
  GOT_OUTPUT_SCHEMA="$(jq -c '.functionSpec.outputSchema // empty' "${MV_RUN_DIR}/cloud/create-tool.json")"
  mv_expect_eq create "cloud functionSpec.outputSchema matches spec" "${WANT_OUTPUT_SCHEMA}" "${GOT_OUTPUT_SCHEMA}" "cloud/create-tool.json"
fi

mv_record create "gcloud CLI surface" INFO "cloud/create-tool.json" "no gcloud CLI command for dialogflow cx tools; REST GET and ListTools used out of band"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create verify-list-tools GET "${TOOLS_LIST_URL}")"
  mv_expect_http create "REST GET verify-list-tools returns 200" 200 "${LIST_CODE}" "cloud/create-verify-list-tools.json"
  FOUND_NAME="$(jq -r --arg NAME "${CLOUD_NAME}" '[.tools[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-verify-list-tools.json")"
  mv_expect_eq create "ListTools returns created tool" "${CLOUD_NAME}" "${FOUND_NAME}" cloud/create-verify-list-tools.json
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DialogflowCXTool "${TOOL_K8S}" 300
SNAP="$(mv_snapshot update DialogflowCXTool "${TOOL_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_UPD_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_UPD_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"

mv_wait_http update tool "${TOOL_URL}" 200

# Compare updated spec fields against cloud response
if [[ -n "${WANT_UPD_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-tool.json" '.displayName' "${WANT_UPD_DISPLAY_NAME}"
fi

if [[ -n "${WANT_UPD_DESCRIPTION}" ]]; then
  mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-tool.json" '.description' "${WANT_UPD_DESCRIPTION}"
fi

if jq -e '.spec.openAPISpec.textSchema' "${SNAP}" >/dev/null 2>&1; then
  WANT_UPD_OPENAPI_TEXT="$(jq -r '.spec.openAPISpec.textSchema' "${SNAP}")"
  mv_expect_jq update "cloud openApiSpec.textSchema updated to new spec" "${MV_RUN_DIR}/cloud/update-tool.json" '.openApiSpec.textSchema' "${WANT_UPD_OPENAPI_TEXT}"
fi

if jq -e '.spec.functionSpec.inputSchema' "${SNAP}" >/dev/null 2>&1; then
  WANT_UPD_INPUT_SCHEMA="$(jq -c '.spec.functionSpec.inputSchema' "${SNAP}")"
  GOT_UPD_INPUT_SCHEMA="$(jq -c '.functionSpec.inputSchema // empty' "${MV_RUN_DIR}/cloud/update-tool.json")"
  mv_expect_eq update "cloud functionSpec.inputSchema updated to new spec" "${WANT_UPD_INPUT_SCHEMA}" "${GOT_UPD_INPUT_SCHEMA}" "cloud/update-tool.json"
fi

if jq -e '.spec.functionSpec.outputSchema' "${SNAP}" >/dev/null 2>&1; then
  WANT_UPD_OUTPUT_SCHEMA="$(jq -c '.spec.functionSpec.outputSchema' "${SNAP}")"
  GOT_UPD_OUTPUT_SCHEMA="$(jq -c '.functionSpec.outputSchema // empty' "${MV_RUN_DIR}/cloud/update-tool.json")"
  mv_expect_eq update "cloud functionSpec.outputSchema updated to new spec" "${WANT_UPD_OUTPUT_SCHEMA}" "${GOT_UPD_OUTPUT_SCHEMA}" "cloud/update-tool.json"
fi

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update verify-list-tools GET "${TOOLS_LIST_URL}")"
  mv_expect_http update "REST GET verify-list-tools returns 200" 200 "${LIST_CODE}" "cloud/update-verify-list-tools.json"
  if [[ -n "${WANT_UPD_DISPLAY_NAME}" ]]; then
    UPD_NAME="$(jq -r --arg dn "${WANT_UPD_DISPLAY_NAME}" '[.tools[]? | select(.displayName==$dn)][0].name // empty' "${MV_RUN_DIR}/cloud/update-verify-list-tools.json")"
    mv_expect_eq update "ListTools reflects updated displayName" "${CLOUD_NAME}" "${UPD_NAME}" cloud/update-verify-list-tools.json
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DialogflowCXTool "${TOOL_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DialogflowCXTool "${TOOL_K8S}" 300
mv_wait_http delete tool-gone "${TOOL_URL}" 404

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete verify-list-tools GET "${TOOLS_LIST_URL}")"
  mv_expect_http delete "REST GET verify-list-tools returns 200" 200 "${LIST_CODE}" "cloud/delete-verify-list-tools.json"
  FOUND_DELETED="$(jq -r --arg NAME "${CLOUD_NAME}" '[.tools[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-verify-list-tools.json")"
  mv_expect_eq delete "ListTools no longer contains deleted tool" "" "${FOUND_DELETED}" cloud/delete-verify-list-tools.json
fi

# Dialogflow CX API does not emit Cloud Audit Log (Admin Activity) entries for
# Tool child resource mutations (only for top-level Agents).
# Verified delete out of band via REST GET 404 and ListTools omission above.
mv_record delete "Cloud Audit Logs surface" INFO "cloud/delete-tool-gone.json" \
  "Dialogflow CX API does not emit Admin Activity audit logs for Tool sub-resources; verified via REST GET 404 and ListTools"

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown DialogflowCXAgent "${AGENT_K8S}" 300
mv_wait_http teardown agent-gone "${AGENT_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
