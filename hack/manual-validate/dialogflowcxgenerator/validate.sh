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

# Human-style validation of the DialogflowCXGenerator controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/dialogflowcxgenerator/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the dialogflowcxgenerator-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (listing generators)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/dialogflowcx/v1alpha1/dialogflowcxgenerator/dialogflowcxgenerator-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DialogflowCXGenerator
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

AGENT_K8S="dialogflowcxagent-${MV_UNIQUE_ID}"
GEN_K8S="dialogflowcxgenerator-${MV_UNIQUE_ID}"

GEN_URL=""
AGENT_URL=""
GENERATORS_LIST_URL=""
AGENT_LIST_URL="https://global-dialogflow.googleapis.com/v3/projects/${GCP_PROJECT_ID}/locations/global/agents?alt=json"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get dialogflowcxgenerator "${GEN_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup DialogflowCXGenerator "${GEN_K8S}" 300
    mv_kubectl get dialogflowcxagent "${AGENT_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup DialogflowCXAgent "${AGENT_K8S}" 300
  fi

  if [[ -n "${GEN_URL:-}" ]]; then
    local code
    code="$(mv_rest cleanup leakcheck-generator GET "${GEN_URL}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (generator gone)" PASS cloud/cleanup-leakcheck-generator.json "HTTP 404"
    elif [[ "${code}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (generator gone)" FAIL cloud/cleanup-leakcheck-generator.json \
        "could not verify (HTTP ${code}); check manually: ${GEN_URL}"
    else
      mv_record cleanup "no leaked GCP resources (generator gone)" FAIL cloud/cleanup-leakcheck-generator.json \
        "HTTP ${code}; clean up manually: DELETE ${GEN_URL}"
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
GENERATORS_LIST_URL="https://dialogflow.googleapis.com/v3/projects/${GCP_PROJECT_ID}/locations/global/agents/${AGENT_ID}/generators?%24alt=json%3Benum-encoding%3Dint"
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
mv_apply create "${CREATE}"
mv_wait_ready create DialogflowCXGenerator "${GEN_K8S}" 300
SNAP="$(mv_snapshot create DialogflowCXGenerator "${GEN_K8S}")"
STATUS_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${STATUS_REF:-<none>}"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_PROMPT="$(jq -r '.spec.promptText.text // empty' "${SNAP}")"

start=$(date +%s)
while :; do
  mv_rest create list-generators GET "${GENERATORS_LIST_URL}" >/dev/null
  CLOUD_NAME="$(jq -r --arg dn "${WANT_DISPLAY_NAME}" '[.generators[]? | select(.displayName==$dn)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-generators.json")"
  [[ -n "${CLOUD_NAME}" ]] && break
  (($(date +%s) - start > 60)) && break
  sleep 3
done

if [[ -n "${CLOUD_NAME}" ]]; then
  mv_record create "ListGenerators contains generator matching displayName" PASS cloud/create-list-generators.json "${CLOUD_NAME}"
  GEN_ID="${CLOUD_NAME##*/}"
  GEN_URL="https://dialogflow.googleapis.com/v3/${CLOUD_NAME}?%24alt=json%3Benum-encoding%3Dint"
else
  mv_record create "ListGenerators contains generator matching displayName" FAIL cloud/create-list-generators.json "not found in ${GENERATORS_LIST_URL}"
  if [[ -n "${STATUS_REF}" ]]; then
    CLOUD_NAME="${STATUS_REF}"
    GEN_ID="${CLOUD_NAME##*/}"
    GEN_URL="https://dialogflow.googleapis.com/v3/${CLOUD_NAME}?%24alt=json%3Benum-encoding%3Dint"
  fi
fi

if [[ -n "${STATUS_REF}" && -n "${CLOUD_NAME}" ]]; then
  mv_expect_eq create "status.externalRef matches cloud resource name" "${CLOUD_NAME}" "${STATUS_REF}" "${SNAP#"${MV_RUN_DIR}"/}"
fi

mv_wait_http create generator "${GEN_URL}" 200

# Compare spec fields against cloud response
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi

if [[ -n "${WANT_PROMPT}" ]]; then
  mv_expect_jq create "cloud promptText.text matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" '.promptText.text' "${WANT_PROMPT}"
fi

if jq -e '.spec.placeholders' "${SNAP}" >/dev/null 2>&1; then
  WANT_PLACEHOLDERS="$(jq -c '.spec.placeholders' "${SNAP}")"
  GOT_PLACEHOLDERS="$(jq -c '.placeholders // []' "${MV_RUN_DIR}/cloud/create-generator.json")"
  mv_expect_eq create "cloud placeholders match spec" "${WANT_PLACEHOLDERS}" "${GOT_PLACEHOLDERS}" "cloud/create-generator.json"
fi

if jq -e '.spec.modelParameter.temperature' "${SNAP}" >/dev/null 2>&1; then
  WANT_TEMP="$(jq -r '.spec.modelParameter.temperature | tostring' "${SNAP}")"
  mv_expect_jq create "cloud modelParameter.temperature matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" '(.modelParameter.temperature // empty) | tostring' "${WANT_TEMP}"
fi

if jq -e '.spec.modelParameter.maxDecodeSteps' "${SNAP}" >/dev/null 2>&1; then
  WANT_STEPS="$(jq -r '.spec.modelParameter.maxDecodeSteps | tostring' "${SNAP}")"
  mv_expect_jq create "cloud modelParameter.maxDecodeSteps matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" '(.modelParameter.maxDecodeSteps // empty) | tostring' "${WANT_STEPS}"
fi

if jq -e '.spec.modelParameter.topP' "${SNAP}" >/dev/null 2>&1; then
  WANT_TOP_P="$(jq -r '.spec.modelParameter.topP | tostring' "${SNAP}")"
  mv_expect_jq create "cloud modelParameter.topP matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" '(.modelParameter.topP // empty) | tostring' "${WANT_TOP_P}"
fi

if jq -e '.spec.modelParameter.topK' "${SNAP}" >/dev/null 2>&1; then
  WANT_TOP_K="$(jq -r '.spec.modelParameter.topK | tostring' "${SNAP}")"
  mv_expect_jq create "cloud modelParameter.topK matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" '(.modelParameter.topK // empty) | tostring' "${WANT_TOP_K}"
fi

mv_record create "gcloud CLI surface" INFO "cloud/create-generator.json" "no gcloud CLI command for dialogflow cx generators; REST GET and ListGenerators used out of band"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create verify-list-generators GET "${GENERATORS_LIST_URL}")"
  mv_expect_http create "REST GET verify-list-generators returns 200" 200 "${LIST_CODE}" "cloud/create-verify-list-generators.json"
  FOUND_NAME="$(jq -r --arg NAME "${CLOUD_NAME}" '[.generators[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-verify-list-generators.json")"
  mv_expect_eq create "ListGenerators returns created generator" "${CLOUD_NAME}" "${FOUND_NAME}" cloud/create-verify-list-generators.json
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DialogflowCXGenerator "${GEN_K8S}" 300
SNAP="$(mv_snapshot update DialogflowCXGenerator "${GEN_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_UPD_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_UPD_PROMPT="$(jq -r '.spec.promptText.text // empty' "${SNAP}")"

mv_wait_http update generator "${GEN_URL}" 200

# Compare updated spec fields against cloud response
if [[ -n "${WANT_UPD_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" '.displayName' "${WANT_UPD_DISPLAY_NAME}"
fi

if [[ -n "${WANT_UPD_PROMPT}" ]]; then
  mv_expect_jq update "cloud promptText.text updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" '.promptText.text' "${WANT_UPD_PROMPT}"
fi

if jq -e '.spec.placeholders' "${SNAP}" >/dev/null 2>&1; then
  WANT_UPD_PLACEHOLDERS="$(jq -c '.spec.placeholders' "${SNAP}")"
  GOT_UPD_PLACEHOLDERS="$(jq -c '.placeholders // []' "${MV_RUN_DIR}/cloud/update-generator.json")"
  mv_expect_eq update "cloud placeholders updated to new spec" "${WANT_UPD_PLACEHOLDERS}" "${GOT_UPD_PLACEHOLDERS}" "cloud/update-generator.json"
fi

if jq -e '.spec.modelParameter.temperature' "${SNAP}" >/dev/null 2>&1; then
  WANT_UPD_TEMP="$(jq -r '.spec.modelParameter.temperature | tostring' "${SNAP}")"
  mv_expect_jq update "cloud modelParameter.temperature updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" '(.modelParameter.temperature // empty) | tostring' "${WANT_UPD_TEMP}"
fi

if jq -e '.spec.modelParameter.maxDecodeSteps' "${SNAP}" >/dev/null 2>&1; then
  WANT_UPD_STEPS="$(jq -r '.spec.modelParameter.maxDecodeSteps | tostring' "${SNAP}")"
  mv_expect_jq update "cloud modelParameter.maxDecodeSteps updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" '(.modelParameter.maxDecodeSteps // empty) | tostring' "${WANT_UPD_STEPS}"
fi

if jq -e '.spec.modelParameter.topP' "${SNAP}" >/dev/null 2>&1; then
  WANT_UPD_TOP_P="$(jq -r '.spec.modelParameter.topP | tostring' "${SNAP}")"
  mv_expect_jq update "cloud modelParameter.topP updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" '(.modelParameter.topP // empty) | tostring' "${WANT_UPD_TOP_P}"
fi

if jq -e '.spec.modelParameter.topK' "${SNAP}" >/dev/null 2>&1; then
  WANT_UPD_TOP_K="$(jq -r '.spec.modelParameter.topK | tostring' "${SNAP}")"
  mv_expect_jq update "cloud modelParameter.topK updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" '(.modelParameter.topK // empty) | tostring' "${WANT_UPD_TOP_K}"
fi

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update verify-list-generators GET "${GENERATORS_LIST_URL}")"
  mv_expect_http update "REST GET verify-list-generators returns 200" 200 "${LIST_CODE}" "cloud/update-verify-list-generators.json"
  if [[ -n "${WANT_UPD_DISPLAY_NAME}" ]]; then
    UPD_NAME="$(jq -r --arg dn "${WANT_UPD_DISPLAY_NAME}" '[.generators[]? | select(.displayName==$dn)][0].name // empty' "${MV_RUN_DIR}/cloud/update-verify-list-generators.json")"
    mv_expect_eq update "ListGenerators reflects updated displayName" "${CLOUD_NAME}" "${UPD_NAME}" cloud/update-verify-list-generators.json
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DialogflowCXGenerator "${GEN_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DialogflowCXGenerator "${GEN_K8S}" 300
mv_wait_http delete generator-gone "${GEN_URL}" 404

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete verify-list-generators GET "${GENERATORS_LIST_URL}")"
  mv_expect_http delete "REST GET verify-list-generators returns 200" 200 "${LIST_CODE}" "cloud/delete-verify-list-generators.json"
  FOUND_DELETED="$(jq -r --arg NAME "${CLOUD_NAME}" '[.generators[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-verify-list-generators.json")"
  mv_expect_eq delete "ListGenerators no longer contains deleted generator" "" "${FOUND_DELETED}" cloud/delete-verify-list-generators.json
fi

# Dialogflow CX API does not emit Cloud Audit Log (Admin Activity) entries for
# Generator child resource mutations (only for top-level Agents).
# Verified delete out of band via REST GET 404 and ListGenerators omission above.
mv_record delete "Cloud Audit Logs surface" INFO "cloud/delete-generator-gone.json" \
  "Dialogflow CX API does not emit Admin Activity audit logs for Generator sub-resources; verified via REST GET 404 and ListGenerators"

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown DialogflowCXAgent "${AGENT_K8S}" 300
mv_wait_http teardown agent-gone "${AGENT_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
