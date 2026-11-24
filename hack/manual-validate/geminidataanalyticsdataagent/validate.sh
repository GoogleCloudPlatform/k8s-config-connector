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

# Human-style validation of the GeminiDataAnalyticsDataAgent controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/geminidataanalyticsdataagent/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the geminidataanalyticsdataagent-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (getIamPolicy and list via REST)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/geminidataanalytics/v1alpha1/geminidataanalyticsdataagent/geminidataanalyticsdataagent-maximal}"
if [[ "${FIXTURE_DIR}" != /* ]]; then
  FIXTURE_DIR="${MV_REPO_ROOT}/${FIXTURE_DIR}"
fi
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init GeminiDataAnalyticsDataAgent
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS=""
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE=""
if [[ -f "${FIXTURE_DIR}/update.yaml" ]]; then
  UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
fi

DATA_AGENT_K8S="$( (grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}' | tr -d "\"'" ) || true )"
DATA_AGENT_ID="$( (grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}' | tr -d "\"'" ) || true )"
DATA_AGENT_ID="${DATA_AGENT_ID:-${DATA_AGENT_K8S}}"

LOCATION="$( (grep -A 5 'spec:' "${CREATE}" | grep -m1 '^[[:space:]]*location:[[:space:]]*' | awk '{print $2}' | tr -d "\"'" ) || true )"
LOCATION="${LOCATION:-global}"

DATA_AGENT_BASE="https://geminidataanalytics.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
DATA_AGENT_URL="${DATA_AGENT_BASE}/dataAgents/${DATA_AGENT_ID}"
EXPECTED_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/dataAgents/${DATA_AGENT_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get geminidataanalyticsdataagent "${DATA_AGENT_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup GeminiDataAnalyticsDataAgent "${DATA_AGENT_K8S}" 300
    if [[ -n "${DEPS}" ]]; then
      :
    fi
  fi

  if [[ -n "${DATA_AGENT_URL:-}" ]]; then
    local code
    code="$(mv_rest cleanup leakcheck-dataagent GET "${DATA_AGENT_URL}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (dataagent gone)" PASS cloud/cleanup-leakcheck-dataagent.json "HTTP 404"
    elif [[ "${code}" == "200" ]]; then
      local del_time
      del_time="$(jq -r '.deleteTime // empty' "${MV_RUN_DIR}/cloud/cleanup-leakcheck-dataagent.json" 2>/dev/null || true)"
      if [[ -n "${del_time}" ]]; then
        mv_record cleanup "no leaked GCP resources (dataagent soft-deleted)" PASS cloud/cleanup-leakcheck-dataagent.json "deleteTime: ${del_time}"
      else
        mv_record cleanup "no leaked GCP resources (dataagent soft-deleted)" FAIL cloud/cleanup-leakcheck-dataagent.json \
          "HTTP 200 but deleteTime not set; clean up manually: DELETE ${DATA_AGENT_URL}"
      fi
    else
      mv_record cleanup "no leaked GCP resources (dataagent soft-deleted)" FAIL cloud/cleanup-leakcheck-dataagent.json \
        "could not verify (HTTP ${code}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${DATA_AGENT_URL}"
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
mv_enable_apis setup geminidataanalytics.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS}" ]]; then
  mv_apply deps "${DEPS}"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create GeminiDataAnalyticsDataAgent "${DATA_AGENT_K8S}" 300

mv_wait_http create dataagent-get "${DATA_AGENT_URL}" 200

# gcloud CLI tool availability check
mv_record create "gcloud surface for GeminiDataAnalyticsDataAgent" INFO "" "no gcloud CLI for Gemini Data Analytics DataAgent (geminidataanalytics.googleapis.com); using raw REST"

SNAP="$(mv_snapshot create GeminiDataAnalyticsDataAgent "${DATA_AGENT_K8S}")"

# Compare spec fields against cloud response
mv_expect_jq create "cloud name matches expected" "${MV_RUN_DIR}/cloud/create-dataagent-get.json" ".name" "${EXPECTED_NAME}"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "displayName matches spec" "${MV_RUN_DIR}/cloud/create-dataagent-get.json" ".displayName" "${WANT_DISPLAY_NAME}"
fi

WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "description matches spec" "${MV_RUN_DIR}/cloud/create-dataagent-get.json" ".description" "${WANT_DESCRIPTION}"
fi

for k in $(jq -r '.spec.labels | keys[]? // empty' "${SNAP}"); do
  val="$(jq -r ".spec.labels[\"${k}\"]" "${SNAP}")"
  mv_expect_jq create "cloud label '${k}' matches spec" "${MV_RUN_DIR}/cloud/create-dataagent-get.json" ".labels[\"${k}\"]" "${val}"
done

# Functional check: getIamPolicy
if [[ "${FUNCTIONAL}" == "1" ]]; then
  IAM_CODE="$(mv_rest create get-iam-policy POST "${DATA_AGENT_URL}:getIamPolicy" "{}")"
  mv_expect_http create "REST POST getIamPolicy returns 200" 200 "${IAM_CODE}" "cloud/create-get-iam-policy.json"
fi

# --- Phase: update -----------------------------------------------------------
if [[ -n "${UPDATE}" ]]; then
  mv_apply update "${UPDATE}"
  mv_wait_ready update GeminiDataAnalyticsDataAgent "${DATA_AGENT_K8S}" 300

  mv_wait_http update dataagent-get "${DATA_AGENT_URL}" 200

  SNAP_UPDATE="$(mv_snapshot update GeminiDataAnalyticsDataAgent "${DATA_AGENT_K8S}")"

  mv_expect_jq update "k8s generation advanced to 2" "${SNAP_UPDATE}" '.metadata.generation // empty' "2"

  WANT_UPDATED_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP_UPDATE}")"
  if [[ -n "${WANT_UPDATED_DISPLAY_NAME}" ]]; then
    mv_expect_jq update "updated displayName matches spec" "${MV_RUN_DIR}/cloud/update-dataagent-get.json" ".displayName" "${WANT_UPDATED_DISPLAY_NAME}"
  fi

  WANT_UPDATED_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP_UPDATE}")"
  if [[ -n "${WANT_UPDATED_DESCRIPTION}" ]]; then
    mv_expect_jq update "updated description matches spec" "${MV_RUN_DIR}/cloud/update-dataagent-get.json" ".description" "${WANT_UPDATED_DESCRIPTION}"
  fi

  for k in $(jq -r '.spec.labels | keys[]? // empty' "${SNAP_UPDATE}"); do
    val="$(jq -r ".spec.labels[\"${k}\"]" "${SNAP_UPDATE}")"
    mv_expect_jq update "updated cloud label '${k}' matches spec" "${MV_RUN_DIR}/cloud/update-dataagent-get.json" ".labels[\"${k}\"]" "${val}"
  done

  if [[ "${FUNCTIONAL}" == "1" ]]; then
    IAM_CODE_UPDATE="$(mv_rest update get-iam-policy POST "${DATA_AGENT_URL}:getIamPolicy" "{}")"
    mv_expect_http update "REST POST getIamPolicy returns 200" 200 "${IAM_CODE_UPDATE}" "cloud/update-get-iam-policy.json"
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP_STEADY="$(mv_snapshot steady GeminiDataAnalyticsDataAgent "${DATA_AGENT_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP_STEADY}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete GeminiDataAnalyticsDataAgent "${DATA_AGENT_K8S}" 300

# GET on soft-deleted DataAgent returns 200 with deleteTime set
GET_DEL_CODE="$(mv_rest delete dataagent-get GET "${DATA_AGENT_URL}")"
mv_expect_http delete "REST GET dataagent returns 200" 200 "${GET_DEL_CODE}" "cloud/delete-dataagent-get.json"
DEL_TIME="$(jq -r '.deleteTime // empty' "${MV_RUN_DIR}/cloud/delete-dataagent-get.json")"
mv_expect_jq delete "dataAgent has deleteTime set" "${MV_RUN_DIR}/cloud/delete-dataagent-get.json" \
  'if (.deleteTime != null and .deleteTime != "") then "has_deleteTime" else "missing_deleteTime" end' "has_deleteTime"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-dataagents GET "${DATA_AGENT_BASE}/dataAgents")"
  mv_expect_http delete "REST GET list-dataagents returns 200" 200 "${LIST_CODE}" "cloud/delete-list-dataagents.json"
  FOUND_DELETED="$(jq -r --arg ID "${DATA_AGENT_ID}" '[.dataAgents[]? | select(.name | endswith("/dataAgents/" + $ID))][0].name // empty' "${MV_RUN_DIR}/cloud/delete-list-dataagents.json")"
  mv_expect_eq delete "List dataAgents no longer contains deleted dataAgent" "" "${FOUND_DELETED}" cloud/delete-list-dataagents.json
fi

# Audit log check
expect_dataagent_audit() {
  local phase="$1" label="$2" want="$3" timeout="${4:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.resourceName:\"${DATA_AGENT_ID}\" AND timestamp>=\"${MV_STARTED_UTC}\""
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    got="$(jq -r '[.[].protoPayload.methodName | split(".") | last] as $m | reduce range(0; $m | length) as $i ([]; if $i == 0 or $m[$i] != $m[$i - 1] then . + [$m[$i]] else . end) | join(",")' "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
    [[ "${got}" == "${want}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 15
  done
  echo "filter: ${filter}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.filter.txt"
  jq -r '.[] | "\(.timestamp)  \(.protoPayload.methodName)  \(.protoPayload.authenticationInfo.principalEmail // "-")"' \
    "${MV_RUN_DIR}/${out}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.summary.txt" 2>/dev/null || true
  mv_trace "gcloud logging read '${filter}' --project ${GCP_PROJECT_ID} --order=asc --format=json" \
    "$(jq 'length' "${MV_RUN_DIR}/${out}" 2>/dev/null || echo 0) entries: ${got:-none}" "${out}"
  mv_expect_eq "${phase}" "audit log write calls for ${DATA_AGENT_ID}" "${want}" "${got}" "${out}"
}

expect_dataagent_audit delete audit-dataagent "CreateDataAgent,UpdateDataAgent,DeleteDataAgent"

# --- Phase: teardown ---------------------------------------------------------
if [[ -n "${DEPS}" ]]; then
  :
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
