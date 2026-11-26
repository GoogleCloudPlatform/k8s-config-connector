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

# Human-style validation of the DesignCenterApplicationTemplate controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/designcenterapplicationtemplate/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the designcenterapplicationtemplate-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (listing application templates)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/designcenter/v1alpha1/designcenterapplicationtemplate/designcenterapplicationtemplate-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DesignCenterApplicationTemplate
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
SPACE_K8S="dc-space-${MV_UNIQUE_ID}"
TEMPLATE_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
SPACE_ID="${SPACE_K8S}"
TEMPLATE_ID="${TEMPLATE_K8S}"
LOCATION="us-central1"

DC="https://designcenter.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
SPACE_URL="${DC}/spaces/${SPACE_ID}"
TEMPLATE_URL="${SPACE_URL}/applicationTemplates/${TEMPLATE_ID}"
EXPECTED_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/spaces/${SPACE_ID}/applicationTemplates/${TEMPLATE_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get designcenterapplicationtemplate "${TEMPLATE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterApplicationTemplate "${TEMPLATE_K8S}" 300
    mv_kubectl get designcenterspace "${SPACE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterSpace "${SPACE_K8S}" 300
  fi
  local code
  code="$(mv_rest cleanup leakcheck-space GET "${SPACE_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (space gone)" PASS cloud/cleanup-leakcheck-space.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (space gone)" FAIL cloud/cleanup-leakcheck-space.json \
      "could not verify (HTTP ${code}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${SPACE_URL}"
  else
    mv_record cleanup "no leaked GCP resources (space gone)" FAIL cloud/cleanup-leakcheck-space.json \
      "HTTP ${code}; clean up manually: DELETE ${SPACE_URL}"
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
mv_enable_apis setup designcenter.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
mv_apply deps "${DEPS}"
mv_wait_ready deps DesignCenterSpace "${SPACE_K8S}" 300
mv_wait_http deps space "${SPACE_URL}" 200

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create DesignCenterApplicationTemplate "${TEMPLATE_K8S}" 300
SNAP="$(mv_snapshot create DesignCenterApplicationTemplate "${TEMPLATE_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_COMPOSITION_TYPE="$(jq -r '.spec.compositionType // empty' "${SNAP}")"
WANT_INPUTS_ROOT="$(jq -r 'if .spec.inputsAsRootModuleVariables != null then (.spec.inputsAsRootModuleVariables|tostring) else empty end' "${SNAP}")"
WANT_PARAM_KEY="$(jq -r '.spec.applicationParameters[0].key // empty' "${SNAP}")"
WANT_PARAM_VAL="$(jq -r '.spec.applicationParameters[0].value.stringValue // empty' "${SNAP}")"

mv_wait_http create template "${TEMPLATE_URL}" 200

# Compare spec fields against cloud response
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-template.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-template.json" '.description' "${WANT_DESCRIPTION}"
fi
if [[ -n "${WANT_COMPOSITION_TYPE}" ]]; then
  mv_expect_jq create "cloud compositionType matches spec" "${MV_RUN_DIR}/cloud/create-template.json" \
    'if .compositionType == 1 or .compositionType == "STANDARD" then "STANDARD" elif .compositionType == 2 or .compositionType == "COMPOSITE" then "COMPOSITE" else (.compositionType|tostring) end' \
    "${WANT_COMPOSITION_TYPE}"
fi
if [[ -n "${WANT_INPUTS_ROOT}" ]]; then
  mv_expect_jq create "cloud inputsAsRootModuleVariables matches spec" "${MV_RUN_DIR}/cloud/create-template.json" \
    '(.inputsAsRootModuleVariables // false) | tostring' "${WANT_INPUTS_ROOT}"
fi
if [[ -n "${WANT_PARAM_KEY}" ]]; then
  mv_expect_jq create "cloud applicationParameters key matches spec" "${MV_RUN_DIR}/cloud/create-template.json" \
    '(.applicationParameters[0].key // empty)' "${WANT_PARAM_KEY}"
  mv_expect_jq create "cloud applicationParameters value matches spec" "${MV_RUN_DIR}/cloud/create-template.json" \
    '(.applicationParameters[0].value // empty)' "${WANT_PARAM_VAL}"
fi

mv_expect_jq create "cloud resource name matches expected ID" "${MV_RUN_DIR}/cloud/create-template.json" '.name' "${EXPECTED_NAME}"
mv_record create "gcloud CLI surface" INFO "cloud/create-template.json" "no gcloud CLI command for designcenter; REST GET used out of band"

# Functional check: ListApplicationTemplates contains the newly created template
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create list-templates GET "${SPACE_URL}/applicationTemplates")"
  mv_expect_http create "REST GET list-templates returns 200" 200 "${LIST_CODE}" "cloud/create-list-templates.json"
  FOUND_NAME="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.applicationTemplates[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-templates.json")"
  mv_expect_eq create "ListApplicationTemplates returns created template" "${EXPECTED_NAME}" "${FOUND_NAME}" cloud/create-list-templates.json
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DesignCenterApplicationTemplate "${TEMPLATE_K8S}" 300
SNAP="$(mv_snapshot update DesignCenterApplicationTemplate "${TEMPLATE_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_COMPOSITION_TYPE="$(jq -r '.spec.compositionType // empty' "${SNAP}")"
WANT_INPUTS_ROOT="$(jq -r 'if .spec.inputsAsRootModuleVariables != null then (.spec.inputsAsRootModuleVariables|tostring) else empty end' "${SNAP}")"
WANT_PARAM_KEY="$(jq -r '.spec.applicationParameters[0].key // empty' "${SNAP}")"
WANT_PARAM_VAL="$(jq -r '.spec.applicationParameters[0].value.stringValue // empty' "${SNAP}")"

mv_wait_http update template "${TEMPLATE_URL}" 200

if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-template.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-template.json" '.description' "${WANT_DESCRIPTION}"
fi
if [[ -n "${WANT_COMPOSITION_TYPE}" ]]; then
  mv_expect_jq update "cloud compositionType matches spec" "${MV_RUN_DIR}/cloud/update-template.json" \
    'if .compositionType == 1 or .compositionType == "STANDARD" then "STANDARD" elif .compositionType == 2 or .compositionType == "COMPOSITE" then "COMPOSITE" else (.compositionType|tostring) end' \
    "${WANT_COMPOSITION_TYPE}"
fi
if [[ -n "${WANT_INPUTS_ROOT}" ]]; then
  mv_expect_jq update "cloud inputsAsRootModuleVariables matches spec" "${MV_RUN_DIR}/cloud/update-template.json" \
    '(.inputsAsRootModuleVariables // false) | tostring' "${WANT_INPUTS_ROOT}"
fi
if [[ -n "${WANT_PARAM_KEY}" ]]; then
  mv_expect_jq update "cloud applicationParameters key matches spec" "${MV_RUN_DIR}/cloud/update-template.json" \
    '(.applicationParameters[0].key // empty)' "${WANT_PARAM_KEY}"
  mv_expect_jq update "cloud applicationParameters value matches spec" "${MV_RUN_DIR}/cloud/update-template.json" \
    '(.applicationParameters[0].value // empty)' "${WANT_PARAM_VAL}"
fi

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-template.json")"
UPDATE_T="$(jq -r '.updateTime // empty' "${MV_RUN_DIR}/cloud/update-template.json")"
mv_record update "createTime preserved, updateTime updated" INFO cloud/update-template.json "createTime=${CREATE_T} updateTime=${UPDATE_T}"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update list-templates GET "${SPACE_URL}/applicationTemplates")"
  mv_expect_http update "REST GET list-templates returns 200" 200 "${LIST_CODE}" "cloud/update-list-templates.json"
  if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
    UPDATED_DISPLAY_NAME="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.applicationTemplates[]? | select(.name==$NAME)][0].displayName // empty' "${MV_RUN_DIR}/cloud/update-list-templates.json")"
    mv_expect_eq update "ListApplicationTemplates reflects updated displayName" "${WANT_DISPLAY_NAME}" "${UPDATED_DISPLAY_NAME}" cloud/update-list-templates.json
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DesignCenterApplicationTemplate "${TEMPLATE_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# Audit log check for DesignCenterApplicationTemplate:
# In Cloud Audit Logs:
# 1) CreateApplicationTemplate has protoPayload.resourceName as parent space and
#    request.applicationTemplateId as the template ID.
# 2) UpdateApplicationTemplate and DeleteApplicationTemplate have protoPayload.resourceName
#    as the full template resource name.
# Filtering on the template ID catches all three methods.
expect_template_audit() {
  local phase="$1" label="$2" template_id="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND \"${template_id}\" AND timestamp>=\"${MV_STARTED_UTC}\""
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    got="$(jq -r '[.[] | select(.operation.last != true) | .protoPayload.methodName | split(".") | last] | join(",")' "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
    [[ "${got}" == "${want}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 15
  done
  echo "filter: ${filter}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.filter.txt"
  mv_expect_eq "${phase}" "audit log write calls on ${template_id}" "${want}" "${got}" "${out}"
}

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DesignCenterApplicationTemplate "${TEMPLATE_K8S}" 300
mv_wait_http delete template-gone "${TEMPLATE_URL}" 404
expect_template_audit delete audit-template "${TEMPLATE_ID}" \
  "CreateApplicationTemplate,UpdateApplicationTemplate,DeleteApplicationTemplate"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-templates GET "${SPACE_URL}/applicationTemplates")"
  mv_expect_http delete "REST GET list-templates returns 200" 200 "${LIST_CODE}" "cloud/delete-list-templates.json"
  FOUND_DELETED="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.applicationTemplates[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-list-templates.json")"
  mv_expect_eq delete "ListApplicationTemplates no longer contains deleted template" "" "${FOUND_DELETED}" cloud/delete-list-templates.json
fi

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown DesignCenterSpace "${SPACE_K8S}" 300
mv_wait_http teardown space-gone "${SPACE_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
