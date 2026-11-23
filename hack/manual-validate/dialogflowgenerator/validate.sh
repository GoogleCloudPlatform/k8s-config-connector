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

# Human-style validation of the DialogflowGenerator controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/dialogflowgenerator/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the dialogflowgenerator-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (listing generators)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/dialogflow/v1alpha1/dialogflowgenerator/dialogflowgenerator-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DialogflowGenerator
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS=""
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
GEN_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
LOCATION="$(grep -m1 '^[[:space:]]*location:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
LOCATION="${LOCATION:-global}"

RESOURCE_ID="$(grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
GEN_ID="${RESOURCE_ID:-${GEN_K8S}}"

if [[ "${LOCATION}" == "global" ]]; then
  DF="https://dialogflow.googleapis.com/v2/projects/${GCP_PROJECT_ID}/locations/global"
else
  DF="https://${LOCATION}-dialogflow.googleapis.com/v2/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
fi
LIST_URL="${DF}/generators"
GEN_URL="${LIST_URL}/${GEN_ID}"
CLOUD_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/generators/${GEN_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get dialogflowgenerator "${GEN_K8S}" >/dev/null 2>&1 && mv_delete cleanup DialogflowGenerator "${GEN_K8S}" 300
  fi
  local code
  code="$(mv_rest cleanup leakcheck-generator GET "${GEN_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (generator gone)" PASS cloud/cleanup-leakcheck-generator.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (generator gone)" FAIL cloud/cleanup-leakcheck-generator.json \
      "could not verify (HTTP ${code}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${GEN_URL}"
  else
    mv_record cleanup "no leaked GCP resources (generator gone)" FAIL cloud/cleanup-leakcheck-generator.json \
      "HTTP ${code}; clean up manually: DELETE ${GEN_URL}"
  fi

  code="$(mv_rest cleanup leakcheck-list GET "${LIST_URL}")"
  if [[ "${code}" == "200" ]]; then
    local leaked
    leaked="$(jq -r --arg uid "${MV_UNIQUE_ID}" '[.generators[]? | select(.name | contains($uid))] | length' "${MV_RUN_DIR}/cloud/cleanup-leakcheck-list.json" 2>/dev/null || echo 0)"
    if [[ "${leaked}" == "0" ]]; then
      mv_record cleanup "no leaked GCP resources (no generators with uniqueId in name)" PASS cloud/cleanup-leakcheck-list.json "none found"
    else
      mv_record cleanup "no leaked GCP resources (no generators with uniqueId in name)" FAIL cloud/cleanup-leakcheck-list.json \
        "found ${leaked} leaked generators in ${LIST_URL}"
    fi
  else
    mv_record cleanup "no leaked GCP resources" FAIL cloud/cleanup-leakcheck-list.json "could not list generators: HTTP ${code}"
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
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
else
  mv_record deps "dependencies" INFO "" "none required for DialogflowGenerator"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create DialogflowGenerator "${GEN_K8S}" 300
SNAP="$(mv_snapshot create DialogflowGenerator "${GEN_K8S}")"
STATUS_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${STATUS_REF:-<none>}"

if [[ -n "${STATUS_REF}" ]]; then
  mv_expect_eq create "status.externalRef matches cloud resource name" "${CLOUD_NAME}" "${STATUS_REF}" "${SNAP#"${MV_RUN_DIR}"/}"
fi

mv_wait_http create generator "${GEN_URL}" 200

# Compare spec fields against cloud response
WANT_DESC="$(jq -r '.spec.description // empty' "${SNAP}")"
if [[ -n "${WANT_DESC}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" '.description' "${WANT_DESC}"
fi

WANT_MODEL="$(jq -r '.spec.publishedModel // empty' "${SNAP}")"
if [[ -n "${WANT_MODEL}" ]]; then
  mv_expect_jq create "cloud publishedModel matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" '.publishedModel' "${WANT_MODEL}"
fi

WANT_TRIGGER="$(jq -r '.spec.triggerEvent // empty' "${SNAP}")"
if [[ -n "${WANT_TRIGGER}" ]]; then
  mv_expect_jq create "cloud triggerEvent matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" '.triggerEvent' "${WANT_TRIGGER}"
fi

WANT_MAX_TOKENS="$(jq -r '.spec.inferenceParameter.maxOutputTokens // empty | tostring' "${SNAP}")"
if [[ -n "${WANT_MAX_TOKENS}" ]]; then
  mv_expect_jq create "cloud inferenceParameter.maxOutputTokens matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" \
    '.inferenceParameter.maxOutputTokens | tostring' "${WANT_MAX_TOKENS}"
fi

WANT_TEMP="$(jq -r '.spec.inferenceParameter.temperature // empty | tostring' "${SNAP}")"
if [[ -n "${WANT_TEMP}" ]]; then
  mv_expect_jq create "cloud inferenceParameter.temperature matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" \
    '.inferenceParameter.temperature | tostring' "${WANT_TEMP}"
fi

WANT_TOP_K="$(jq -r '.spec.inferenceParameter.topK // empty | tostring' "${SNAP}")"
if [[ -n "${WANT_TOP_K}" ]]; then
  mv_expect_jq create "cloud inferenceParameter.topK matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" \
    '.inferenceParameter.topK | tostring' "${WANT_TOP_K}"
fi

WANT_TOP_P="$(jq -r '.spec.inferenceParameter.topP // empty | tostring' "${SNAP}")"
if [[ -n "${WANT_TOP_P}" ]]; then
  mv_expect_jq create "cloud inferenceParameter.topP matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" \
    '.inferenceParameter.topP | tostring' "${WANT_TOP_P}"
fi

WANT_FF_TEXT="$(jq -r '.spec.freeFormContext.text // empty' "${SNAP}")"
if [[ -n "${WANT_FF_TEXT}" ]]; then
  mv_expect_jq create "cloud freeFormContext.text matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" \
    '.freeFormContext.text' "${WANT_FF_TEXT}"
fi

# Check summarizationContext if set (covers freeFormSuggestion and summarySuggestion)
WANT_FF_RESP="$(jq -r '[.spec.summarizationContext.fewShotExamples[]?.output.freeFormSuggestion.response // empty][0] // empty' "${SNAP}")"
if [[ -n "${WANT_FF_RESP}" ]]; then
  mv_expect_jq create "cloud fewShotExamples output.freeFormSuggestion.response matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" \
    '[.summarizationContext.fewShotExamples[]?.output.freeFormSuggestion.response // empty][0]' "${WANT_FF_RESP}"
fi

WANT_SUM_RESP="$(jq -r '[.spec.summarizationContext.fewShotExamples[]?.output.summarySuggestion.summarySections[]?.summary // empty][0] // empty' "${SNAP}")"
if [[ -n "${WANT_SUM_RESP}" ]]; then
  mv_expect_jq create "cloud fewShotExamples output.summarySuggestion.summarySections summary matches spec" "${MV_RUN_DIR}/cloud/create-generator.json" \
    '[.summarizationContext.fewShotExamples[]?.output.summarySuggestion.summarySections[]?.summary // empty][0]' "${WANT_SUM_RESP}"
fi

mv_record create "gcloud CLI surface" INFO "cloud/create-generator.json" "no gcloud CLI command for dialogflow generators; REST GET used out of band"

# Functional check: ListGenerators contains the created generator
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create verify-list-generators GET "${LIST_URL}")"
  mv_expect_http create "REST GET verify-list-generators returns 200" 200 "${LIST_CODE}" "cloud/create-verify-list-generators.json"
  FOUND_NAME="$(jq -r --arg NAME "${CLOUD_NAME}" '[.generators[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-verify-list-generators.json")"
  mv_expect_eq create "ListGenerators returns created generator" "${CLOUD_NAME}" "${FOUND_NAME}" cloud/create-verify-list-generators.json
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DialogflowGenerator "${GEN_K8S}" 300
SNAP="$(mv_snapshot update DialogflowGenerator "${GEN_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update generator "${GEN_URL}" 200

WANT_UPDATED_DESC="$(jq -r '.spec.description // empty' "${SNAP}")"
if [[ -n "${WANT_UPDATED_DESC}" ]]; then
  mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" '.description' "${WANT_UPDATED_DESC}"
fi

WANT_UPDATED_TRIGGER="$(jq -r '.spec.triggerEvent // empty' "${SNAP}")"
if [[ -n "${WANT_UPDATED_TRIGGER}" ]]; then
  mv_expect_jq update "cloud triggerEvent updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" '.triggerEvent' "${WANT_UPDATED_TRIGGER}"
fi

WANT_UPDATED_MAX_TOKENS="$(jq -r '.spec.inferenceParameter.maxOutputTokens // empty | tostring' "${SNAP}")"
if [[ -n "${WANT_UPDATED_MAX_TOKENS}" ]]; then
  mv_expect_jq update "cloud inferenceParameter.maxOutputTokens updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" \
    '.inferenceParameter.maxOutputTokens | tostring' "${WANT_UPDATED_MAX_TOKENS}"
fi

WANT_UPDATED_TEMP="$(jq -r '.spec.inferenceParameter.temperature // empty | tostring' "${SNAP}")"
if [[ -n "${WANT_UPDATED_TEMP}" ]]; then
  mv_expect_jq update "cloud inferenceParameter.temperature updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" \
    '.inferenceParameter.temperature | tostring' "${WANT_UPDATED_TEMP}"
fi

WANT_UPDATED_TOP_K="$(jq -r '.spec.inferenceParameter.topK // empty | tostring' "${SNAP}")"
if [[ -n "${WANT_UPDATED_TOP_K}" ]]; then
  mv_expect_jq update "cloud inferenceParameter.topK updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" \
    '.inferenceParameter.topK | tostring' "${WANT_UPDATED_TOP_K}"
fi

WANT_UPDATED_TOP_P="$(jq -r '.spec.inferenceParameter.topP // empty | tostring' "${SNAP}")"
if [[ -n "${WANT_UPDATED_TOP_P}" ]]; then
  mv_expect_jq update "cloud inferenceParameter.topP updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" \
    '.inferenceParameter.topP | tostring' "${WANT_UPDATED_TOP_P}"
fi

WANT_UPDATED_FF_TEXT="$(jq -r '.spec.freeFormContext.text // empty' "${SNAP}")"
if [[ -n "${WANT_UPDATED_FF_TEXT}" ]]; then
  mv_expect_jq update "cloud freeFormContext.text updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" \
    '.freeFormContext.text' "${WANT_UPDATED_FF_TEXT}"
fi

WANT_UPDATED_FF_RESP="$(jq -r '[.spec.summarizationContext.fewShotExamples[]?.output.freeFormSuggestion.response // empty][0] // empty' "${SNAP}")"
if [[ -n "${WANT_UPDATED_FF_RESP}" ]]; then
  mv_expect_jq update "cloud fewShotExamples output.freeFormSuggestion.response updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" \
    '[.summarizationContext.fewShotExamples[]?.output.freeFormSuggestion.response // empty][0]' "${WANT_UPDATED_FF_RESP}"
fi

WANT_UPDATED_SUM_RESP="$(jq -r '[.spec.summarizationContext.fewShotExamples[]?.output.summarySuggestion.summarySections[]?.summary // empty][0] // empty' "${SNAP}")"
if [[ -n "${WANT_UPDATED_SUM_RESP}" ]]; then
  mv_expect_jq update "cloud fewShotExamples output.summarySuggestion.summarySections summary updated to new spec" "${MV_RUN_DIR}/cloud/update-generator.json" \
    '[.summarizationContext.fewShotExamples[]?.output.summarySuggestion.summarySections[]?.summary // empty][0]' "${WANT_UPDATED_SUM_RESP}"
fi

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-generator.json")"
UPDATE_T="$(jq -r '.updateTime // empty' "${MV_RUN_DIR}/cloud/update-generator.json")"
mv_record update "createTime preserved, updateTime updated" INFO cloud/update-generator.json "createTime=${CREATE_T} updateTime=${UPDATE_T}"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update verify-list-generators GET "${LIST_URL}")"
  mv_expect_http update "REST GET verify-list-generators returns 200" 200 "${LIST_CODE}" "cloud/update-verify-list-generators.json"
  if [[ -n "${WANT_UPDATED_DESC}" ]]; then
    UPDATED_DESC="$(jq -r --arg NAME "${CLOUD_NAME}" '[.generators[]? | select(.name==$NAME)][0].description // empty' "${MV_RUN_DIR}/cloud/update-verify-list-generators.json")"
    mv_expect_eq update "ListGenerators reflects updated description" "${WANT_UPDATED_DESC}" "${UPDATED_DESC}" cloud/update-verify-list-generators.json
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DialogflowGenerator "${GEN_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DialogflowGenerator "${GEN_K8S}" 300
mv_wait_http delete generator-gone "${GEN_URL}" 404
mv_record delete "audit log write calls" INFO "cloud/delete-generator-gone.json" "Dialogflow Generators API does not emit Cloud Audit Logs; verified deletion via REST GET 404"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete verify-list-generators GET "${LIST_URL}")"
  mv_expect_http delete "REST GET verify-list-generators returns 200" 200 "${LIST_CODE}" "cloud/delete-verify-list-generators.json"
  FOUND_DELETED="$(jq -r --arg NAME "${CLOUD_NAME}" '[.generators[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-verify-list-generators.json")"
  mv_expect_eq delete "ListGenerators no longer contains deleted generator" "" "${FOUND_DELETED}" cloud/delete-verify-list-generators.json
fi

# --- Phase: dependency teardown ----------------------------------------------
if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
