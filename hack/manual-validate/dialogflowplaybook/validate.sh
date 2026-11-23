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

# Human-style validation of the DialogflowPlaybook controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/dialogflowplaybook/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the dialogflowplaybook-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (listing playbooks)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/dialogflow/v1alpha1/dialogflowplaybook/dialogflowplaybook-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DialogflowPlaybook
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

AGENT_K8S="dialogflowcxagent-${MV_UNIQUE_ID}"
PLAYBOOK_K8S="dialogflowplaybook-${MV_UNIQUE_ID}"

PLAYBOOK_URL=""
AGENT_URL=""

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get dialogflowplaybook "${PLAYBOOK_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup DialogflowPlaybook "${PLAYBOOK_K8S}" 300
    mv_kubectl get dialogflowcxagent "${AGENT_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup DialogflowCXAgent "${AGENT_K8S}" 300
  fi

  if [[ -n "${PLAYBOOK_URL:-}" ]]; then
    local code
    code="$(mv_rest cleanup leakcheck-playbook GET "${PLAYBOOK_URL}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (playbook gone)" PASS cloud/cleanup-leakcheck-playbook.json "HTTP 404"
    elif [[ "${code}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (playbook gone)" FAIL cloud/cleanup-leakcheck-playbook.json \
        "could not verify (HTTP ${code}); check manually: ${PLAYBOOK_URL}"
    else
      mv_record cleanup "no leaked GCP resources (playbook gone)" FAIL cloud/cleanup-leakcheck-playbook.json \
        "HTTP ${code}; clean up manually: DELETE ${PLAYBOOK_URL}"
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
mv_apply deps "${DEPS}"
mv_wait_ready deps DialogflowCXAgent "${AGENT_K8S}" 300
SNAP_AGENT="$(mv_snapshot deps DialogflowCXAgent "${AGENT_K8S}")"
AGENT_NAME="$(jq -r '.status.name // .status.externalRef // empty' "${SNAP_AGENT}")"
AGENT_ID="${AGENT_NAME##*/}"
AGENT_URL="https://global-dialogflow.googleapis.com/v3/projects/${GCP_PROJECT_ID}/locations/global/agents/${AGENT_ID}?alt=json"
PLAYBOOK_LIST_URL="https://dialogflow.googleapis.com/v3/projects/${GCP_PROJECT_ID}/locations/global/agents/${AGENT_ID}/playbooks?%24alt=json%3Benum-encoding%3Dint"
mv_wait_http deps agent "${AGENT_URL}" 200

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create DialogflowPlaybook "${PLAYBOOK_K8S}" 300
SNAP="$(mv_snapshot create DialogflowPlaybook "${PLAYBOOK_K8S}")"
STATUS_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${STATUS_REF:-<none>}"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_GOAL="$(jq -r '.spec.goal // empty' "${SNAP}")"
WANT_GUIDELINES="$(jq -r '.spec.instruction.guidelines // empty' "${SNAP}")"

start=$(date +%s)
while :; do
  mv_rest create list-playbooks GET "${PLAYBOOK_LIST_URL}" >/dev/null
  CLOUD_NAME="$(jq -r --arg dn "${WANT_DISPLAY_NAME}" '[.playbooks[]? | select(.displayName==$dn)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-playbooks.json")"
  [[ -n "${CLOUD_NAME}" ]] && break
  (($(date +%s) - start > 60)) && break
  sleep 3
done

if [[ -n "${CLOUD_NAME}" ]]; then
  mv_record create "ListPlaybooks contains playbook matching displayName" PASS cloud/create-list-playbooks.json "${CLOUD_NAME}"
  PLAYBOOK_ID="${CLOUD_NAME##*/}"
  PLAYBOOK_URL="https://dialogflow.googleapis.com/v3/${CLOUD_NAME}?%24alt=json%3Benum-encoding%3Dint"
else
  mv_record create "ListPlaybooks contains playbook matching displayName" FAIL cloud/create-list-playbooks.json "not found in ${PLAYBOOK_LIST_URL}"
  if [[ -n "${STATUS_REF}" ]]; then
    CLOUD_NAME="${STATUS_REF}"
    PLAYBOOK_ID="${CLOUD_NAME##*/}"
    PLAYBOOK_URL="https://dialogflow.googleapis.com/v3/${CLOUD_NAME}?%24alt=json%3Benum-encoding%3Dint"
  fi
fi

if [[ -n "${STATUS_REF}" && -n "${CLOUD_NAME}" ]]; then
  mv_expect_eq create "status.externalRef matches cloud resource name" "${CLOUD_NAME}" "${STATUS_REF}" "${SNAP#"${MV_RUN_DIR}"/}"
fi

mv_wait_http create playbook "${PLAYBOOK_URL}" 200

# Compare spec fields against cloud response
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-playbook.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_GOAL}" ]]; then
  mv_expect_jq create "cloud goal matches spec" "${MV_RUN_DIR}/cloud/create-playbook.json" '.goal' "${WANT_GOAL}"
fi
if [[ -n "${WANT_GUIDELINES}" ]]; then
  mv_expect_jq create "cloud instruction guidelines match spec" "${MV_RUN_DIR}/cloud/create-playbook.json" '.instruction.guidelines' "${WANT_GUIDELINES}"
fi
if jq -e '.spec.instruction.steps // empty' "${SNAP}" >/dev/null 2>&1; then
  WANT_STEPS="$(jq -c '[.spec.instruction.steps[]?.text]' "${SNAP}")"
  GOT_STEPS="$(jq -c '[.instruction.steps[]?.text]' "${MV_RUN_DIR}/cloud/create-playbook.json")"
  mv_expect_eq create "cloud instruction steps match spec" "${WANT_STEPS}" "${GOT_STEPS}" "cloud/create-playbook.json"
fi
if jq -e '.spec.inputParameterDefinitions // empty' "${SNAP}" >/dev/null 2>&1; then
  WANT_IN_PARAMS="$(jq -c '[.spec.inputParameterDefinitions[]? | {name: .name, description: .description}]' "${SNAP}")"
  GOT_IN_PARAMS="$(jq -c '[.inputParameterDefinitions[]? | {name: .name, description: .description}]' "${MV_RUN_DIR}/cloud/create-playbook.json")"
  mv_expect_eq create "cloud inputParameterDefinitions match spec" "${WANT_IN_PARAMS}" "${GOT_IN_PARAMS}" "cloud/create-playbook.json"
fi
if jq -e '.spec.outputParameterDefinitions // empty' "${SNAP}" >/dev/null 2>&1; then
  WANT_OUT_PARAMS="$(jq -c '[.spec.outputParameterDefinitions[]? | {name: .name, description: .description}]' "${SNAP}")"
  GOT_OUT_PARAMS="$(jq -c '[.outputParameterDefinitions[]? | {name: .name, description: .description}]' "${MV_RUN_DIR}/cloud/create-playbook.json")"
  mv_expect_eq create "cloud outputParameterDefinitions match spec" "${WANT_OUT_PARAMS}" "${GOT_OUT_PARAMS}" "cloud/create-playbook.json"
fi

# ToolCall fields check: action, inputParameters, tool
WANT_TOOL_ACTION="$(jq -r '.. | .toolCall?.action // empty' "${SNAP}" | head -1)"
if [[ -n "${WANT_TOOL_ACTION}" ]]; then
  GOT_TOOL_ACTION="$(jq -r '.. | .toolCall?.action // empty' "${MV_RUN_DIR}/cloud/create-playbook.json" | head -1)"
  mv_expect_eq create "cloud toolCall action matches spec" "${WANT_TOOL_ACTION}" "${GOT_TOOL_ACTION}" "cloud/create-playbook.json"
fi
WANT_TOOL="$(jq -r '.. | .toolCall?.tool // empty' "${SNAP}" | head -1)"
if [[ -n "${WANT_TOOL}" ]]; then
  GOT_TOOL="$(jq -r '.. | .toolCall?.tool // empty' "${MV_RUN_DIR}/cloud/create-playbook.json" | head -1)"
  mv_expect_eq create "cloud toolCall tool matches spec" "${WANT_TOOL}" "${GOT_TOOL}" "cloud/create-playbook.json"
fi
if jq -e '.. | .toolCall?.inputParameters // empty' "${SNAP}" >/dev/null 2>&1; then
  WANT_INPUT_PARAMS="$(jq -c '.. | .toolCall?.inputParameters // empty' "${SNAP}" | head -1)"
  GOT_INPUT_PARAMS="$(jq -c '.. | .toolCall?.inputParameters // empty' "${MV_RUN_DIR}/cloud/create-playbook.json" | head -1)"
  mv_expect_eq create "cloud toolCall inputParameters match spec" "${WANT_INPUT_PARAMS}" "${GOT_INPUT_PARAMS}" "cloud/create-playbook.json"
fi

mv_record create "gcloud CLI surface" INFO "cloud/create-playbook.json" "no gcloud CLI command for dialogflow cx playbooks; REST GET used out of band"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create verify-list-playbooks GET "${PLAYBOOK_LIST_URL}")"
  mv_expect_http create "REST GET verify-list-playbooks returns 200" 200 "${LIST_CODE}" "cloud/create-verify-list-playbooks.json"
  FOUND_NAME="$(jq -r --arg NAME "${CLOUD_NAME}" '[.playbooks[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-verify-list-playbooks.json")"
  mv_expect_eq create "ListPlaybooks returns created playbook" "${CLOUD_NAME}" "${FOUND_NAME}" cloud/create-verify-list-playbooks.json
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DialogflowPlaybook "${PLAYBOOK_K8S}" 300
SNAP="$(mv_snapshot update DialogflowPlaybook "${PLAYBOOK_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_UPD_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_UPD_GOAL="$(jq -r '.spec.goal // empty' "${SNAP}")"
WANT_UPD_GUIDELINES="$(jq -r '.spec.instruction.guidelines // empty' "${SNAP}")"

mv_wait_http update playbook "${PLAYBOOK_URL}" 200

if [[ -n "${WANT_UPD_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated" "${MV_RUN_DIR}/cloud/update-playbook.json" '.displayName' "${WANT_UPD_DISPLAY_NAME}"
fi
if [[ -n "${WANT_UPD_GOAL}" ]]; then
  mv_expect_jq update "cloud goal updated" "${MV_RUN_DIR}/cloud/update-playbook.json" '.goal' "${WANT_UPD_GOAL}"
fi
if [[ -n "${WANT_UPD_GUIDELINES}" ]]; then
  mv_expect_jq update "cloud instruction guidelines updated" "${MV_RUN_DIR}/cloud/update-playbook.json" '.instruction.guidelines' "${WANT_UPD_GUIDELINES}"
fi

# ToolCall fields check on update: action, inputParameters, tool
WANT_UPD_TOOL_ACTION="$(jq -r '.. | .toolCall?.action // empty' "${SNAP}" | head -1)"
if [[ -n "${WANT_UPD_TOOL_ACTION}" ]]; then
  GOT_UPD_TOOL_ACTION="$(jq -r '.. | .toolCall?.action // empty' "${MV_RUN_DIR}/cloud/update-playbook.json" | head -1)"
  mv_expect_eq update "cloud toolCall action updated" "${WANT_UPD_TOOL_ACTION}" "${GOT_UPD_TOOL_ACTION}" "cloud/update-playbook.json"
fi
WANT_UPD_TOOL="$(jq -r '.. | .toolCall?.tool // empty' "${SNAP}" | head -1)"
if [[ -n "${WANT_UPD_TOOL}" ]]; then
  GOT_UPD_TOOL="$(jq -r '.. | .toolCall?.tool // empty' "${MV_RUN_DIR}/cloud/update-playbook.json" | head -1)"
  mv_expect_eq update "cloud toolCall tool updated" "${WANT_UPD_TOOL}" "${GOT_UPD_TOOL}" "cloud/update-playbook.json"
fi
if jq -e '.. | .toolCall?.inputParameters // empty' "${SNAP}" >/dev/null 2>&1; then
  WANT_UPD_INPUT_PARAMS="$(jq -c '.. | .toolCall?.inputParameters // empty' "${SNAP}" | head -1)"
  GOT_UPD_INPUT_PARAMS="$(jq -c '.. | .toolCall?.inputParameters // empty' "${MV_RUN_DIR}/cloud/update-playbook.json" | head -1)"
  mv_expect_eq update "cloud toolCall inputParameters updated" "${WANT_UPD_INPUT_PARAMS}" "${GOT_UPD_INPUT_PARAMS}" "cloud/update-playbook.json"
fi

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update verify-list-playbooks GET "${PLAYBOOK_LIST_URL}")"
  mv_expect_http update "REST GET verify-list-playbooks returns 200" 200 "${LIST_CODE}" "cloud/update-verify-list-playbooks.json"
  if [[ -n "${WANT_UPD_DISPLAY_NAME}" ]]; then
    UPD_NAME="$(jq -r --arg dn "${WANT_UPD_DISPLAY_NAME}" '[.playbooks[]? | select(.displayName==$dn)][0].name // empty' "${MV_RUN_DIR}/cloud/update-verify-list-playbooks.json")"
    mv_expect_eq update "ListPlaybooks reflects updated displayName" "${CLOUD_NAME}" "${UPD_NAME}" cloud/update-verify-list-playbooks.json
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DialogflowPlaybook "${PLAYBOOK_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DialogflowPlaybook "${PLAYBOOK_K8S}" 300
mv_wait_http delete playbook-gone "${PLAYBOOK_URL}" 404

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete verify-list-playbooks GET "${PLAYBOOK_LIST_URL}")"
  mv_expect_http delete "REST GET verify-list-playbooks returns 200" 200 "${LIST_CODE}" "cloud/delete-verify-list-playbooks.json"
  FOUND_DELETED="$(jq -r --arg NAME "${CLOUD_NAME}" '[.playbooks[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-verify-list-playbooks.json")"
  mv_expect_eq delete "ListPlaybooks no longer contains deleted playbook" "" "${FOUND_DELETED}" cloud/delete-verify-list-playbooks.json
fi

# Dialogflow CX API does not emit Cloud Audit Log (Admin Activity) entries for
# Playbook child resource mutations (only for top-level Agents).
# Verified delete out of band via REST GET 404 and ListPlaybooks omission above.
mv_record delete "Cloud Audit Logs surface" INFO "cloud/delete-playbook-gone.json" \
  "Dialogflow CX API does not emit Admin Activity audit logs for Playbooks sub-resources; verified via REST GET 404 and ListPlaybooks"

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown DialogflowCXAgent "${AGENT_K8S}" 300
mv_wait_http teardown agent-gone "${AGENT_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
