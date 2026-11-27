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

# Human-style validation of the DialogflowGeneratorEvaluation controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/dialogflowgeneratorevaluation/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml
#                    (default: the dialogflowgeneratorevaluation-maximal test fixture)
#   FUNCTIONAL=0     skip functional checks (listing evaluations, verify done=true)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/dialogflow/v1alpha1/dialogflowgeneratorevaluation/dialogflowgeneratorevaluation-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DialogflowGeneratorEvaluation
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"

# Names derived from the fixture so out-of-band checks never ask KCC.
BUCKET_NAME="${GCP_PROJECT_ID}-df-eval-max-${MV_UNIQUE_ID}"
GENERATOR_K8S="dialogflowgenerator-${MV_UNIQUE_ID}"
GENERATOR_ID="${GENERATOR_K8S}"
EVAL_K8S="dfgeneratoreval-max-${MV_UNIQUE_ID}"
LOCATION="global"

DF="https://dialogflow.googleapis.com/v2/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
GENERATOR_URL="${DF}/generators/${GENERATOR_ID}"
EVALUATIONS_URL="${GENERATOR_URL}/evaluations"
BUCKET_URL="https://storage.googleapis.com/storage/v1/b/${BUCKET_NAME}"
EVAL_URL=""
EVAL_NAME=""

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get dialogflowgeneratorevaluation "${EVAL_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup DialogflowGeneratorEvaluation "${EVAL_K8S}" 300
    mv_kubectl get dialogflowgenerator "${GENERATOR_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup DialogflowGenerator "${GENERATOR_K8S}" 300
    gcloud storage rm -r "gs://${BUCKET_NAME}/**" >/dev/null 2>&1 || true
    mv_kubectl get storagebucket "${BUCKET_NAME}" >/dev/null 2>&1 && \
      mv_delete cleanup StorageBucket "${BUCKET_NAME}" 300
  fi

  # Leak check generator evaluation: if EVAL_URL is unset, discover it via list
  if [[ -z "${EVAL_URL}" && -n "${EVALUATIONS_URL:-}" ]]; then
    mv_rest cleanup leakcheck-list-evals GET "${EVALUATIONS_URL}" >/dev/null 2>&1 || true
    local disc_name
    disc_name="$(jq -r '.generatorEvaluations[0].name // empty' "${MV_RUN_DIR}/cloud/cleanup-leakcheck-list-evals.json" 2>/dev/null || true)"
    if [[ -n "${disc_name}" ]]; then
      EVAL_NAME="${disc_name}"
      EVAL_URL="https://dialogflow.googleapis.com/v2/${disc_name}"
    fi
  fi

  if [[ -n "${EVAL_URL}" ]]; then
    local ecode
    ecode="$(mv_rest cleanup leakcheck-eval GET "${EVAL_URL}")"
    if [[ "${ecode}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (evaluation gone)" PASS cloud/cleanup-leakcheck-eval.json "HTTP 404"
    elif [[ "${ecode}" == "200" ]]; then
      # Delete leaked evaluation directly via out-of-band REST
      mv_rest cleanup leakcheck-delete-eval DELETE "${EVAL_URL}" >/dev/null 2>&1 || true
      mv_record cleanup "no leaked GCP resources (evaluation gone)" FAIL cloud/cleanup-leakcheck-eval.json \
        "HTTP ${ecode}; evaluation was leaked; deleted out-of-band: DELETE ${EVAL_URL}"
    else
      mv_record cleanup "no leaked GCP resources (evaluation gone)" FAIL cloud/cleanup-leakcheck-eval.json \
        "could not verify (HTTP ${ecode}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${EVAL_URL}"
    fi
  elif [[ -n "${EVALUATIONS_URL:-}" ]]; then
    local lecode
    lecode="$(mv_rest cleanup leakcheck-list-evals GET "${EVALUATIONS_URL}")"
    if [[ "${lecode}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (evaluation collection gone)" PASS cloud/cleanup-leakcheck-list-evals.json "HTTP 404"
    fi
  fi

  # Leak check generator via REST
  local gcode
  gcode="$(mv_rest cleanup leakcheck-generator GET "${GENERATOR_URL}")"
  if [[ "${gcode}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (generator gone)" PASS cloud/cleanup-leakcheck-generator.json "HTTP 404"
  elif [[ "${gcode}" == "200" ]]; then
    mv_rest cleanup leakcheck-delete-gen DELETE "${GENERATOR_URL}" >/dev/null 2>&1 || true
    mv_record cleanup "no leaked GCP resources (generator gone)" FAIL cloud/cleanup-leakcheck-generator.json \
      "HTTP ${gcode}; generator was leaked; deleted out-of-band: DELETE ${GENERATOR_URL}"
  else
    mv_record cleanup "no leaked GCP resources (generator gone)" FAIL cloud/cleanup-leakcheck-generator.json \
      "could not verify (HTTP ${gcode}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${GENERATOR_URL}"
  fi

  # Clean objects and check bucket leak
  gcloud storage rm -r "gs://${BUCKET_NAME}/**" >/dev/null 2>&1 || true
  gcloud storage rm -r "gs://${BUCKET_NAME}" >/dev/null 2>&1 || true

  local bcode
  bcode="$(mv_rest cleanup leakcheck-bucket GET "${BUCKET_URL}")"
  if [[ "${bcode}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (bucket gone)" PASS cloud/cleanup-leakcheck-bucket.json "HTTP 404"
  else
    mv_record cleanup "no leaked GCP resources (bucket gone)" FAIL cloud/cleanup-leakcheck-bucket.json \
      "HTTP ${bcode}; check bucket: ${BUCKET_URL}"
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

# Cloud Audit Logs check for Dialogflow GeneratorEvaluation:
# Note: Dialogflow ES/v2 generators and generator evaluations do not emit entries to
# Cloud Audit Logs activity logs; if none are found within timeout, record INFO.
expect_eval_audit() {
  local phase="$1" label="$2" generator_id="$3" want="$4" timeout="${5:-60}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.methodName:\"GeneratorEvaluation\" AND \"${generator_id}\" AND timestamp>=\"${MV_STARTED_UTC}\""
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    got="$(jq -r '[.[] | select(.operation.last != true) | .protoPayload.methodName | split(".") | last] | join(",")' "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
    [[ -n "${got}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 10
  done
  echo "filter: ${filter}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.filter.txt"
  if [[ -z "${got}" ]]; then
    mv_record "${phase}" "audit logs for generator evaluation" INFO "${out}" \
      "Dialogflow does not emit Cloud Audit Logs for GeneratorEvaluation operations"
  else
    mv_expect_eq "${phase}" "audit log write calls for generator evaluation (${generator_id})" "${want}" "${got}" "${out}"
  fi
}

# ---------------------------------------------------------------------------
# Every API the controller, the dependencies and the out-of-band checks call.
mv_enable_apis setup dialogflow.googleapis.com storage.googleapis.com logging.googleapis.com
mv_record setup "gcloud CLI support" INFO "" "gcloud has no command group for dialogflow generator evaluations; verified out-of-band via raw REST (curl)"
mv_up

# --- Phase: dependencies -----------------------------------------------------
mv_apply deps "${DEPS}"
mv_wait_ready deps StorageBucket "${BUCKET_NAME}" 300
mv_wait_ready deps DialogflowGenerator "${GENERATOR_K8S}" 300
mv_wait_http deps bucket "${BUCKET_URL}" 200
mv_wait_http deps generator "${GENERATOR_URL}" 200

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
# Evaluation creation initiates an LRO on Dialogflow; poll until Ready.
mv_wait_ready create DialogflowGeneratorEvaluation "${EVAL_K8S}" 120
SNAP="$(mv_snapshot create DialogflowGeneratorEvaluation "${EVAL_K8S}")"
EXTERNAL_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${EXTERNAL_REF:-<none>}"

# Out-of-band discovery via ListGeneratorEvaluations
start=$(date +%s)
while :; do
  LIST_CODE="$(mv_rest create list-evaluations GET "${EVALUATIONS_URL}")"
  EVAL_NAME="$(jq -r '.generatorEvaluations[0].name // empty' "${MV_RUN_DIR}/cloud/create-list-evaluations.json")"
  [[ -n "${EVAL_NAME}" ]] && break
  (($(date +%s) - start > 120)) && break
  sleep 5
done
if [[ -z "${EVAL_NAME}" && -n "${EXTERNAL_REF}" ]]; then
  EVAL_NAME="${EXTERNAL_REF}"
fi
mv_expect_http create "REST GET list-evaluations returns 200" 200 "${LIST_CODE}" "cloud/create-list-evaluations.json"
mv_expect_eq create "evaluation found in list" "true" "$([[ -n "${EVAL_NAME}" ]] && echo true || echo false)" "cloud/create-list-evaluations.json"

EVAL_URL="https://dialogflow.googleapis.com/v2/${EVAL_NAME}"
mv_wait_http create eval "${EVAL_URL}" 200

# Compare every spec field in applied object with cloud response
WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi

WANT_GCS_PATH="$(jq -r '.spec.generatorEvaluationConfig.outputGCSBucketPath // empty' "${SNAP}")"
if [[ -n "${WANT_GCS_PATH}" ]]; then
  mv_expect_jq create "cloud outputGcsBucketPath matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '.generatorEvaluationConfig.outputGcsBucketPath' "${WANT_GCS_PATH}"
fi

WANT_EVAL_VERSION="$(jq -r '.spec.generatorEvaluationConfig.summarizationConfig.evaluatorVersion // empty' "${SNAP}")"
if [[ -n "${WANT_EVAL_VERSION}" ]]; then
  mv_expect_jq create "cloud evaluatorVersion matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '.generatorEvaluationConfig.summarizationConfig.evaluatorVersion' "${WANT_EVAL_VERSION}"
fi

WANT_SAMPLE_SIZE="$(jq -r '.spec.generatorEvaluationConfig.inputDataConfig.sampleSize // empty' "${SNAP}")"
if [[ -n "${WANT_SAMPLE_SIZE}" ]]; then
  mv_expect_jq create "cloud sampleSize matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '(.generatorEvaluationConfig.inputDataConfig.sampleSize // empty | tostring)' "${WANT_SAMPLE_SIZE}"
fi

WANT_SUMMARY_GEN_OPT="$(jq -r '.spec.generatorEvaluationConfig.inputDataConfig.summaryGenerationOption // empty' "${SNAP}")"
if [[ -n "${WANT_SUMMARY_GEN_OPT}" ]]; then
  mv_expect_jq create "cloud summaryGenerationOption matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" \
    'if .generatorEvaluationConfig.inputDataConfig.summaryGenerationOption == 1 or .generatorEvaluationConfig.inputDataConfig.summaryGenerationOption == "ALWAYS_GENERATE" then "ALWAYS_GENERATE" elif .generatorEvaluationConfig.inputDataConfig.summaryGenerationOption == 2 or .generatorEvaluationConfig.inputDataConfig.summaryGenerationOption == "NEVER_GENERATE" then "NEVER_GENERATE" else (.generatorEvaluationConfig.inputDataConfig.summaryGenerationOption // empty | tostring) end' \
    "${WANT_SUMMARY_GEN_OPT}"
fi

WANT_START_TIME="$(jq -r '.spec.generatorEvaluationConfig.inputDataConfig.agentAssistInputDataConfig.startTime // empty' "${SNAP}")"
if [[ -n "${WANT_START_TIME}" ]]; then
  mv_expect_jq create "cloud agentAssistInputDataConfig startTime matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '.generatorEvaluationConfig.inputDataConfig.agentAssistInputDataConfig.startTime' "${WANT_START_TIME}"
fi

WANT_END_TIME="$(jq -r '.spec.generatorEvaluationConfig.inputDataConfig.agentAssistInputDataConfig.endTime // empty' "${SNAP}")"
if [[ -n "${WANT_END_TIME}" ]]; then
  mv_expect_jq create "cloud agentAssistInputDataConfig endTime matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '.generatorEvaluationConfig.inputDataConfig.agentAssistInputDataConfig.endTime' "${WANT_END_TIME}"
fi

WANT_INIT_DESC="$(jq -r '.spec.initialGenerator.description // empty' "${SNAP}")"
if [[ -n "${WANT_INIT_DESC}" ]]; then
  mv_expect_jq create "cloud initialGenerator description matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '.initialGenerator.description' "${WANT_INIT_DESC}"
fi

WANT_INIT_MODEL="$(jq -r '.spec.initialGenerator.publishedModel // empty' "${SNAP}")"
if [[ -n "${WANT_INIT_MODEL}" ]]; then
  mv_expect_jq create "cloud initialGenerator publishedModel matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '.initialGenerator.publishedModel' "${WANT_INIT_MODEL}"
fi

WANT_INIT_TRIGGER="$(jq -r '.spec.initialGenerator.triggerEvent // empty' "${SNAP}")"
if [[ -n "${WANT_INIT_TRIGGER}" ]]; then
  mv_expect_jq create "cloud initialGenerator triggerEvent matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" \
    'if .initialGenerator.triggerEvent == 1 or .initialGenerator.triggerEvent == "END_OF_UTTERANCE" then "END_OF_UTTERANCE" elif .initialGenerator.triggerEvent == 2 or .initialGenerator.triggerEvent == "MANUAL" then "MANUAL" else (.initialGenerator.triggerEvent // empty | tostring) end' \
    "${WANT_INIT_TRIGGER}"
fi

WANT_MAX_TOKENS="$(jq -r '.spec.initialGenerator.inferenceParameter.maxOutputTokens // empty' "${SNAP}")"
if [[ -n "${WANT_MAX_TOKENS}" ]]; then
  mv_expect_jq create "cloud inferenceParameter maxOutputTokens matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '(.initialGenerator.inferenceParameter.maxOutputTokens // empty | tostring)' "${WANT_MAX_TOKENS}"
fi

WANT_TEMP="$(jq -r '.spec.initialGenerator.inferenceParameter.temperature // empty' "${SNAP}")"
if [[ -n "${WANT_TEMP}" ]]; then
  mv_expect_jq create "cloud inferenceParameter temperature matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '(.initialGenerator.inferenceParameter.temperature // empty | tostring)' "${WANT_TEMP}"
fi

WANT_TOP_K="$(jq -r '.spec.initialGenerator.inferenceParameter.topK // empty' "${SNAP}")"
if [[ -n "${WANT_TOP_K}" ]]; then
  mv_expect_jq create "cloud inferenceParameter topK matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '(.initialGenerator.inferenceParameter.topK // empty | tostring)' "${WANT_TOP_K}"
fi

WANT_TOP_P="$(jq -r '.spec.initialGenerator.inferenceParameter.topP // empty' "${SNAP}")"
if [[ -n "${WANT_TOP_P}" ]]; then
  mv_expect_jq create "cloud inferenceParameter topP matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" '(.initialGenerator.inferenceParameter.topP // empty | tostring)' "${WANT_TOP_P}"
fi

WANT_SECTION_TYPE="$(jq -r '.spec.initialGenerator.summarizationContext.summarizationSections[0].type // empty' "${SNAP}")"
if [[ -n "${WANT_SECTION_TYPE}" ]]; then
  mv_expect_jq create "cloud summarizationSection type matches spec" "${MV_RUN_DIR}/cloud/create-eval.json" \
    'if .initialGenerator.summarizationContext.summarizationSections[0].type == 1 or .initialGenerator.summarizationContext.summarizationSections[0].type == "SITUATION" then "SITUATION" elif .initialGenerator.summarizationContext.summarizationSections[0].type == 2 or .initialGenerator.summarizationContext.summarizationSections[0].type == "ACTION" then "ACTION" elif .initialGenerator.summarizationContext.summarizationSections[0].type == 3 or .initialGenerator.summarizationContext.summarizationSections[0].type == "RESOLUTION" then "RESOLUTION" else (.initialGenerator.summarizationContext.summarizationSections[0].type // empty | tostring) end' \
    "${WANT_SECTION_TYPE}"
fi

if [[ "${FUNCTIONAL}" == "1" ]]; then
  # Functional check 1: Dialogflow completed evaluation pipeline (done=true).
  mv_expect_jq create "functional check: evaluation execution status is done" \
    "${MV_RUN_DIR}/cloud/create-eval.json" '(.evaluationStatus.done // false | tostring)' "true"

  # Functional check 2: ListGeneratorEvaluations contains the evaluation.
  LIST_COUNT="$(jq -r '[.generatorEvaluations[]? | select(.name=="'"${EVAL_NAME}"'")] | length' "${MV_RUN_DIR}/cloud/create-list-evaluations.json")"
  mv_expect_eq create "functional check: ListGeneratorEvaluations contains the evaluation" "1" "${LIST_COUNT}" "cloud/create-list-evaluations.json"
fi

# --- Phase: update -----------------------------------------------------------
if [[ -f "${FIXTURE_DIR}/update.yaml" ]]; then
  UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
  mv_apply update "${UPDATE}"
  sleep 2
  mv_wait_ready update DialogflowGeneratorEvaluation "${EVAL_K8S}" 300
  SNAP="$(mv_snapshot update DialogflowGeneratorEvaluation "${EVAL_K8S}")"
  mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"
  mv_wait_http update eval "${EVAL_URL}" 200
else
  # DialogflowGeneratorEvaluation is immutable: Dialogflow API has no UpdateGeneratorEvaluation RPC,
  # and the KCC controller rejects in-place updates.
  mv_record update "update phase: DialogflowGeneratorEvaluation is immutable" INFO "" \
    "Dialogflow API does not provide an update method; evaluation records are immutable"
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DialogflowGeneratorEvaluation "${EVAL_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DialogflowGeneratorEvaluation "${EVAL_K8S}" 300
mv_wait_http delete eval-gone "${EVAL_URL}" 404

# Cloud Audit Logs check: verify exact write methods invoked on Dialogflow GeneratorEvaluation
expect_eval_audit delete audit-eval "${GENERATOR_ID}" \
  "CreateGeneratorEvaluation,DeleteGeneratorEvaluation"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-evaluations GET "${EVALUATIONS_URL}")"
  mv_expect_http delete "REST GET list-evaluations after delete returns 200" 200 "${LIST_CODE}" "cloud/delete-list-evaluations.json"
  LIST_DELETED_COUNT="$(jq -r '[.generatorEvaluations[]? | select(.name=="'"${EVAL_NAME}"'")] | length' "${MV_RUN_DIR}/cloud/delete-list-evaluations.json")"
  mv_expect_eq delete "ListGeneratorEvaluations no longer contains deleted evaluation" "0" "${LIST_DELETED_COUNT}" "cloud/delete-list-evaluations.json"
fi

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown DialogflowGenerator "${GENERATOR_K8S}" 300
mv_wait_http teardown generator-gone "${GENERATOR_URL}" 404

# Clean any output objects in storage bucket before deleting bucket
gcloud storage rm -r "gs://${BUCKET_NAME}/**" >/dev/null 2>&1 || true
mv_delete teardown StorageBucket "${BUCKET_NAME}" 300
mv_wait_http teardown bucket-gone "${BUCKET_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
