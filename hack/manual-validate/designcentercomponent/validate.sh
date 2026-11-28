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

# Human-style validation of the DesignCenterComponent controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/designcentercomponent/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the designcentercomponent-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (listing components)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/designcenter/v1alpha1/designcentercomponent/designcentercomponent-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DesignCenterComponent
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
SPACE_K8S="$(python3 -c "import yaml; docs = list(yaml.safe_load_all(open('${DEPS}'))); print([d['metadata']['name'] for d in docs if d.get('kind')=='DesignCenterSpace'][0])")"
TEMPLATE_K8S="$(python3 -c "import yaml; docs = list(yaml.safe_load_all(open('${DEPS}'))); print([d['metadata']['name'] for d in docs if d.get('kind')=='DesignCenterApplicationTemplate'][0])")"
COMPONENT_K8S="$(python3 -c "import yaml; doc = yaml.safe_load(open('${CREATE}')); print(doc['metadata']['name'])")"

SPACE_ID="${SPACE_K8S}"
TEMPLATE_ID="${TEMPLATE_K8S}"
COMPONENT_ID="${COMPONENT_K8S}"
LOCATION="$(python3 -c "import yaml; docs = list(yaml.safe_load_all(open('${DEPS}'))); locs = [d['spec'].get('location') for d in docs if d.get('kind')=='DesignCenterSpace']; print(locs[0] if locs and locs[0] else 'us-central1')")"
LOCATION="${LOCATION:-us-central1}"

DC="https://designcenter.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
SPACE_URL="${DC}/spaces/${SPACE_ID}"
TEMPLATE_URL="${SPACE_URL}/applicationTemplates/${TEMPLATE_ID}"
COMPONENT_URL="${TEMPLATE_URL}/components/${COMPONENT_ID}"
EXPECTED_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/spaces/${SPACE_ID}/applicationTemplates/${TEMPLATE_ID}/components/${COMPONENT_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get designcentercomponent "${COMPONENT_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterComponent "${COMPONENT_K8S}" 300
    mv_kubectl get designcenterapplicationtemplate "${TEMPLATE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterApplicationTemplate "${TEMPLATE_K8S}" 300
    mv_kubectl get designcenterspace "${SPACE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterSpace "${SPACE_K8S}" 300
  fi
  local code_comp code_tpl code_space
  code_comp="$(mv_rest cleanup leakcheck-component GET "${COMPONENT_URL}")"
  if [[ "${code_comp}" == "404" ]]; then
    mv_record cleanup "no leaked GCP component" PASS cloud/cleanup-leakcheck-component.json "HTTP 404"
  elif [[ "${code_comp}" != "200" ]]; then
    mv_record cleanup "no leaked GCP component" FAIL cloud/cleanup-leakcheck-component.json \
      "could not verify (HTTP ${code_comp}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${COMPONENT_URL}"
  else
    mv_record cleanup "no leaked GCP component" FAIL cloud/cleanup-leakcheck-component.json \
      "HTTP ${code_comp}; clean up manually: DELETE ${COMPONENT_URL}"
  fi

  code_tpl="$(mv_rest cleanup leakcheck-template GET "${TEMPLATE_URL}")"
  if [[ "${code_tpl}" == "404" ]]; then
    mv_record cleanup "no leaked GCP application template" PASS cloud/cleanup-leakcheck-template.json "HTTP 404"
  elif [[ "${code_tpl}" != "200" ]]; then
    mv_record cleanup "no leaked GCP application template" FAIL cloud/cleanup-leakcheck-template.json \
      "could not verify (HTTP ${code_tpl}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${TEMPLATE_URL}"
  else
    mv_record cleanup "no leaked GCP application template" FAIL cloud/cleanup-leakcheck-template.json \
      "HTTP ${code_tpl}; clean up manually: DELETE ${TEMPLATE_URL}"
  fi

  code_space="$(mv_rest cleanup leakcheck-space GET "${SPACE_URL}")"
  if [[ "${code_space}" == "404" ]]; then
    mv_record cleanup "no leaked GCP space" PASS cloud/cleanup-leakcheck-space.json "HTTP 404"
  elif [[ "${code_space}" != "200" ]]; then
    mv_record cleanup "no leaked GCP space" FAIL cloud/cleanup-leakcheck-space.json \
      "could not verify (HTTP ${code_space}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${SPACE_URL}"
  else
    mv_record cleanup "no leaked GCP space" FAIL cloud/cleanup-leakcheck-space.json \
      "HTTP ${code_space}; clean up manually: DELETE ${SPACE_URL}"
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
mv_wait_ready deps DesignCenterApplicationTemplate "${TEMPLATE_K8S}" 300
mv_wait_http deps template "${TEMPLATE_URL}" 200

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create DesignCenterComponent "${COMPONENT_K8S}" 300
SNAP="$(mv_snapshot create DesignCenterComponent "${COMPONENT_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

WANT_SHARED_TEMPLATE="$(jq -r '.spec.sharedTemplateRevisionURI // empty' "${SNAP}")"
WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_USE_AS_ROOT="$(jq -r 'if .spec.useAsRootModule != null then (.spec.useAsRootModule|tostring) else empty end' "${SNAP}")"
WANT_PARAM_DISPLAY_NAME="$(jq -r '(.spec.parameters[]? | select(.key=="agent-runtime-1_display_name") | .value.stringValue) // empty' "${SNAP}")"
WANT_PARAM_PROJECT="$(jq -r '(.spec.parameters[]? | select(.key=="agent-runtime-1_project_id") | .value.stringValue) // empty' "${SNAP}")"
WANT_PARAM_REGION="$(jq -r '(.spec.parameters[]? | select(.key=="agent-runtime-1_region") | .value.stringValue) // empty' "${SNAP}")"
WANT_ALIAS_VARIABLE="$(jq -r '(.spec.inputVariableAliases[0].variable) // empty' "${SNAP}")"
WANT_ALIAS_NAME="$(jq -r '(.spec.inputVariableAliases[0].alias) // empty' "${SNAP}")"

mv_wait_http create component "${COMPONENT_URL}" 200

# Compare spec fields against cloud response
if [[ -n "${WANT_SHARED_TEMPLATE}" ]]; then
  mv_expect_jq create "cloud sharedTemplateRevisionUri matches spec" "${MV_RUN_DIR}/cloud/create-component.json" \
    '.sharedTemplateRevisionUri // empty' "${WANT_SHARED_TEMPLATE}"
fi
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-component.json" \
    '.displayName // empty' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_USE_AS_ROOT}" ]]; then
  mv_expect_jq create "cloud useAsRootModule matches spec" "${MV_RUN_DIR}/cloud/create-component.json" \
    '(.useAsRootModule // false) | tostring' "${WANT_USE_AS_ROOT}"
fi
if [[ -n "${WANT_PARAM_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud parameter agent-runtime-1_display_name matches spec" "${MV_RUN_DIR}/cloud/create-component.json" \
    '(.parameters[]? | select(.key=="agent-runtime-1_display_name") | .value) // empty' "${WANT_PARAM_DISPLAY_NAME}"
fi
if [[ -n "${WANT_PARAM_PROJECT}" ]]; then
  mv_expect_jq create "cloud parameter agent-runtime-1_project_id matches spec" "${MV_RUN_DIR}/cloud/create-component.json" \
    '(.parameters[]? | select(.key=="agent-runtime-1_project_id") | .value) // empty' "${WANT_PARAM_PROJECT}"
fi
if [[ -n "${WANT_PARAM_REGION}" ]]; then
  mv_expect_jq create "cloud parameter agent-runtime-1_region matches spec" "${MV_RUN_DIR}/cloud/create-component.json" \
    '(.parameters[]? | select(.key=="agent-runtime-1_region") | .value) // empty' "${WANT_PARAM_REGION}"
fi
if [[ -n "${WANT_ALIAS_VARIABLE}" ]]; then
  mv_expect_jq create "cloud inputVariableAliases alias matches spec" "${MV_RUN_DIR}/cloud/create-component.json" \
    "(.inputVariableAliases[]? | select(.variable==\"${WANT_ALIAS_VARIABLE}\") | .alias) // empty" "${WANT_ALIAS_NAME}"
fi

mv_expect_jq create "cloud resource name matches expected ID" "${MV_RUN_DIR}/cloud/create-component.json" '.name' "${EXPECTED_NAME}"
mv_record create "gcloud CLI surface" INFO "cloud/create-component.json" "no gcloud CLI command for designcenter; REST GET used out of band"

# Functional check: ListComponents contains the newly created component
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create list-components GET "${TEMPLATE_URL}/components")"
  mv_expect_http create "REST GET list-components returns 200" 200 "${LIST_CODE}" "cloud/create-list-components.json"
  FOUND_NAME="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.components[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-components.json")"
  mv_expect_eq create "ListComponents returns created component" "${EXPECTED_NAME}" "${FOUND_NAME}" cloud/create-list-components.json
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DesignCenterComponent "${COMPONENT_K8S}" 300
SNAP="$(mv_snapshot update DesignCenterComponent "${COMPONENT_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_USE_AS_ROOT="$(jq -r 'if .spec.useAsRootModule != null then (.spec.useAsRootModule|tostring) else empty end' "${SNAP}")"
WANT_ALIAS_VARIABLE="$(jq -r '(.spec.inputVariableAliases[0].variable) // empty' "${SNAP}")"
WANT_ALIAS_NAME="$(jq -r '(.spec.inputVariableAliases[0].alias) // empty' "${SNAP}")"

mv_wait_http update component "${COMPONENT_URL}" 200

if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-component.json" \
    '.displayName // empty' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_USE_AS_ROOT}" ]]; then
  mv_expect_jq update "cloud useAsRootModule updated to new spec" "${MV_RUN_DIR}/cloud/update-component.json" \
    '(.useAsRootModule // false) | tostring' "${WANT_USE_AS_ROOT}"
fi
if [[ -n "${WANT_ALIAS_VARIABLE}" ]]; then
  mv_expect_jq update "cloud inputVariableAliases alias updated to new spec" "${MV_RUN_DIR}/cloud/update-component.json" \
    "(.inputVariableAliases[]? | select(.variable==\"${WANT_ALIAS_VARIABLE}\") | .alias) // empty" "${WANT_ALIAS_NAME}"
fi
if [[ -n "${WANT_SHARED_TEMPLATE}" ]]; then
  mv_expect_jq update "cloud sharedTemplateRevisionUri preserved" "${MV_RUN_DIR}/cloud/update-component.json" \
    '.sharedTemplateRevisionUri // empty' "${WANT_SHARED_TEMPLATE}"
fi
if [[ -n "${WANT_PARAM_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud parameter agent-runtime-1_display_name preserved" "${MV_RUN_DIR}/cloud/update-component.json" \
    '(.parameters[]? | select(.key=="agent-runtime-1_display_name") | .value) // empty' "${WANT_PARAM_DISPLAY_NAME}"
fi
if [[ -n "${WANT_PARAM_PROJECT}" ]]; then
  mv_expect_jq update "cloud parameter agent-runtime-1_project_id preserved" "${MV_RUN_DIR}/cloud/update-component.json" \
    '(.parameters[]? | select(.key=="agent-runtime-1_project_id") | .value) // empty' "${WANT_PARAM_PROJECT}"
fi
if [[ -n "${WANT_PARAM_REGION}" ]]; then
  mv_expect_jq update "cloud parameter agent-runtime-1_region preserved" "${MV_RUN_DIR}/cloud/update-component.json" \
    '(.parameters[]? | select(.key=="agent-runtime-1_region") | .value) // empty' "${WANT_PARAM_REGION}"
fi

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-component.json")"
UPDATE_T="$(jq -r '.updateTime // empty' "${MV_RUN_DIR}/cloud/update-component.json")"
mv_record update "createTime preserved, updateTime updated" INFO cloud/update-component.json "createTime=${CREATE_T} updateTime=${UPDATE_T}"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update list-components GET "${TEMPLATE_URL}/components")"
  mv_expect_http update "REST GET list-components returns 200" 200 "${LIST_CODE}" "cloud/update-list-components.json"
  if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
    UPDATED_DISPLAY_NAME="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.components[]? | select(.name==$NAME)][0].displayName // empty' "${MV_RUN_DIR}/cloud/update-list-components.json")"
    mv_expect_eq update "ListComponents reflects updated displayName" "${WANT_DISPLAY_NAME}" "${UPDATED_DISPLAY_NAME}" cloud/update-list-components.json
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DesignCenterComponent "${COMPONENT_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# Audit log check for DesignCenterComponent:
# In Cloud Audit Logs:
# 1) CreateComponent has protoPayload.resourceName as parent applicationTemplate and
#    request.componentId as the component ID.
# 2) UpdateComponent and DeleteComponent have protoPayload.resourceName
#    as the full component resource name.
# Filtering on the component ID catches all three methods.
expect_component_audit() {
  local phase="$1" label="$2" component_id="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.serviceName=\"designcenter.googleapis.com\" AND \"${component_id}\" AND timestamp>=\"${MV_STARTED_UTC}\""
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    got="$(mv_audit_calls "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
    [[ "${got}" == "${want}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 15
  done
  echo "filter: ${filter}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.filter.txt"
  mv_expect_eq "${phase}" "audit log write calls on ${component_id}" "${want}" "${got}" "${out}"
}

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DesignCenterComponent "${COMPONENT_K8S}" 300
mv_wait_http delete component-gone "${COMPONENT_URL}" 404
expect_component_audit delete audit-component "${COMPONENT_ID}" \
  "CreateComponent,UpdateComponent,DeleteComponent"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-components GET "${TEMPLATE_URL}/components")"
  mv_expect_http delete "REST GET list-components returns 200" 200 "${LIST_CODE}" "cloud/delete-list-components.json"
  FOUND_DELETED="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.components[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-list-components.json")"
  mv_expect_eq delete "ListComponents no longer contains deleted component" "" "${FOUND_DELETED}" cloud/delete-list-components.json
fi

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown DesignCenterApplicationTemplate "${TEMPLATE_K8S}" 300
mv_wait_http teardown template-gone "${TEMPLATE_URL}" 404

mv_delete teardown DesignCenterSpace "${SPACE_K8S}" 300
mv_wait_http teardown space-gone "${SPACE_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
