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

# Human-style validation of the DesignCenterPolicy controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/designcenterpolicy/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the designcenterpolicy-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (listing policies)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/designcenter/v1alpha1/designcenterpolicy/designcenterpolicy-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DesignCenterPolicy
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
SPACE_K8S="$(awk '/kind: DesignCenterSpace/{f=1} f && /^[[:space:]]*name:/{print $2; exit}' "${DEPS}")"
TEMPLATE_K8S="$(awk '/kind: DesignCenterApplicationTemplate/{f=1} f && /^[[:space:]]*name:/{print $2; exit}' "${DEPS}")"
POLICY_K8S="$(awk '/kind: DesignCenterPolicy/{f=1} f && /^[[:space:]]*name:/{print $2; exit}' "${CREATE}")"

SPACE_ID="${SPACE_K8S}"
TEMPLATE_ID="${TEMPLATE_K8S}"
POLICY_ID="${POLICY_K8S}"

LOCATION="$(awk '/kind: DesignCenterSpace/{f=1} f && /^[[:space:]]*location:/{print $2; exit}' "${DEPS}")"
LOCATION="${LOCATION:-us-central1}"

DC="https://designcenter.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
SPACE_URL="${DC}/spaces/${SPACE_ID}"
TEMPLATE_URL="${SPACE_URL}/applicationTemplates/${TEMPLATE_ID}"
POLICY_URL="${TEMPLATE_URL}/policies/${POLICY_ID}"
EXPECTED_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/spaces/${SPACE_ID}/applicationTemplates/${TEMPLATE_ID}/policies/${POLICY_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get designcenterpolicy "${POLICY_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterPolicy "${POLICY_K8S}" 300
    mv_kubectl get designcenterapplicationtemplate "${TEMPLATE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterApplicationTemplate "${TEMPLATE_K8S}" 300
    mv_kubectl get designcenterspace "${SPACE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterSpace "${SPACE_K8S}" 300
  fi
  local code_pol code_tmpl code_space
  code_pol="$(mv_rest cleanup leakcheck-policy GET "${POLICY_URL}")"
  if [[ "${code_pol}" == "404" ]]; then
    mv_record cleanup "no leaked GCP policy" PASS cloud/cleanup-leakcheck-policy.json "HTTP 404"
  elif [[ "${code_pol}" != "200" ]]; then
    mv_record cleanup "no leaked GCP policy" FAIL cloud/cleanup-leakcheck-policy.json \
      "could not verify (HTTP ${code_pol}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${POLICY_URL}"
  else
    mv_record cleanup "no leaked GCP policy" FAIL cloud/cleanup-leakcheck-policy.json \
      "HTTP ${code_pol}; clean up manually: DELETE ${POLICY_URL}"
  fi

  code_tmpl="$(mv_rest cleanup leakcheck-template GET "${TEMPLATE_URL}")"
  if [[ "${code_tmpl}" == "404" ]]; then
    mv_record cleanup "no leaked GCP template" PASS cloud/cleanup-leakcheck-template.json "HTTP 404"
  elif [[ "${code_tmpl}" != "200" ]]; then
    mv_record cleanup "no leaked GCP template" FAIL cloud/cleanup-leakcheck-template.json \
      "could not verify (HTTP ${code_tmpl}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${TEMPLATE_URL}"
  else
    mv_record cleanup "no leaked GCP template" FAIL cloud/cleanup-leakcheck-template.json \
      "HTTP ${code_tmpl}; clean up manually: DELETE ${TEMPLATE_URL}"
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
mv_wait_ready create DesignCenterPolicy "${POLICY_K8S}" 300
SNAP="$(mv_snapshot create DesignCenterPolicy "${POLICY_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_POLICY_TYPE="$(jq -r '.spec.policyType // empty' "${SNAP}")"
WANT_POLICY_URI="$(jq -r '.spec.policyURI // empty' "${SNAP}")"
WANT_POLICY_REVISION_URI="$(jq -r '.spec.policyRevisionURI // empty' "${SNAP}")"
WANT_CONDITIONS="$(jq -c '(.spec.applyConditions // []) | sort_by(.conditionKey)' "${SNAP}")"

mv_wait_http create policy "${POLICY_URL}" 200

# Compare spec fields against cloud response
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-policy.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-policy.json" '.description' "${WANT_DESCRIPTION}"
fi
if [[ -n "${WANT_POLICY_TYPE}" ]]; then
  mv_expect_jq create "cloud policyType matches spec" "${MV_RUN_DIR}/cloud/create-policy.json" \
    'if .policyType == 1 or .policyType == "COMPLIANCE_FRAMEWORK" then "COMPLIANCE_FRAMEWORK" else (.policyType|tostring) end' \
    "${WANT_POLICY_TYPE}"
fi
if [[ -n "${WANT_POLICY_URI}" ]]; then
  mv_expect_jq create "cloud policyURI matches spec" "${MV_RUN_DIR}/cloud/create-policy.json" \
    '(.policyUri // .policyURI // empty)' "${WANT_POLICY_URI}"
fi
if [[ -n "${WANT_POLICY_REVISION_URI}" ]]; then
  mv_expect_jq create "cloud policyRevisionURI matches spec" "${MV_RUN_DIR}/cloud/create-policy.json" \
    '(.policyRevisionUri // .policyRevisionURI // empty)' "${WANT_POLICY_REVISION_URI}"
fi
if [[ "${WANT_CONDITIONS}" != "[]" ]]; then
  GOT_CONDITIONS="$(jq -c '(.applyConditions // []) | sort_by(.conditionKey)' "${MV_RUN_DIR}/cloud/create-policy.json")"
  mv_expect_eq create "cloud applyConditions matches spec" "${WANT_CONDITIONS}" "${GOT_CONDITIONS}" "cloud/create-policy.json"
fi

mv_expect_jq create "cloud resource name matches expected ID" "${MV_RUN_DIR}/cloud/create-policy.json" '.name' "${EXPECTED_NAME}"

# Out of band gcloud describe
if mv_cli create gcloud-describe -- gcloud alpha design-center spaces application-templates policies describe "${POLICY_ID}" --application-template="${TEMPLATE_ID}" --space="${SPACE_ID}" --project="${GCP_PROJECT_ID}" --location="${LOCATION}"; then
  if grep -q "${EXPECTED_NAME}" "${MV_RUN_DIR}/cloud/create-gcloud-describe.txt"; then
    mv_record create "gcloud alpha describe returns created policy" PASS cloud/create-gcloud-describe.txt ""
  else
    mv_record create "gcloud alpha describe returns created policy" FAIL cloud/create-gcloud-describe.txt "policy name not found in gcloud output"
  fi
else
  mv_record create "gcloud alpha describe command executed" INFO cloud/create-gcloud-describe.txt "gcloud alpha command returned non-zero"
fi

# Functional check: ListPolicies contains the newly created policy
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create list-policies GET "${TEMPLATE_URL}/policies")"
  mv_expect_http create "REST GET list-policies returns 200" 200 "${LIST_CODE}" "cloud/create-list-policies.json"
  FOUND_NAME="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.policies[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-policies.json")"
  mv_expect_eq create "ListPolicies returns created policy" "${EXPECTED_NAME}" "${FOUND_NAME}" cloud/create-list-policies.json
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DesignCenterPolicy "${POLICY_K8S}" 300
SNAP="$(mv_snapshot update DesignCenterPolicy "${POLICY_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_POLICY_TYPE="$(jq -r '.spec.policyType // empty' "${SNAP}")"
WANT_POLICY_URI="$(jq -r '.spec.policyURI // empty' "${SNAP}")"
WANT_POLICY_REVISION_URI="$(jq -r '.spec.policyRevisionURI // empty' "${SNAP}")"
WANT_CONDITIONS="$(jq -c '(.spec.applyConditions // []) | sort_by(.conditionKey)' "${SNAP}")"

mv_wait_http update policy "${POLICY_URL}" 200

if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-policy.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-policy.json" '.description' "${WANT_DESCRIPTION}"
fi
if [[ -n "${WANT_POLICY_TYPE}" ]]; then
  mv_expect_jq update "cloud policyType matches spec" "${MV_RUN_DIR}/cloud/update-policy.json" \
    'if .policyType == 1 or .policyType == "COMPLIANCE_FRAMEWORK" then "COMPLIANCE_FRAMEWORK" else (.policyType|tostring) end' \
    "${WANT_POLICY_TYPE}"
fi
if [[ -n "${WANT_POLICY_URI}" ]]; then
  mv_expect_jq update "cloud policyURI matches spec" "${MV_RUN_DIR}/cloud/update-policy.json" \
    '(.policyUri // .policyURI // empty)' "${WANT_POLICY_URI}"
fi
if [[ -n "${WANT_POLICY_REVISION_URI}" ]]; then
  mv_expect_jq update "cloud policyRevisionURI matches spec" "${MV_RUN_DIR}/cloud/update-policy.json" \
    '(.policyRevisionUri // .policyRevisionURI // empty)' "${WANT_POLICY_REVISION_URI}"
fi
if [[ "${WANT_CONDITIONS}" != "[]" ]]; then
  GOT_CONDITIONS="$(jq -c '(.applyConditions // []) | sort_by(.conditionKey)' "${MV_RUN_DIR}/cloud/update-policy.json")"
  mv_expect_eq update "cloud applyConditions updated to new spec" "${WANT_CONDITIONS}" "${GOT_CONDITIONS}" "cloud/update-policy.json"
fi

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-policy.json")"
UPDATE_T="$(jq -r '.updateTime // empty' "${MV_RUN_DIR}/cloud/update-policy.json")"
mv_record update "createTime preserved, updateTime updated" INFO cloud/update-policy.json "createTime=${CREATE_T} updateTime=${UPDATE_T}"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update list-policies GET "${TEMPLATE_URL}/policies")"
  mv_expect_http update "REST GET list-policies returns 200" 200 "${LIST_CODE}" "cloud/update-list-policies.json"
  if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
    UPDATED_DISPLAY_NAME="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.policies[]? | select(.name==$NAME)][0].displayName // empty' "${MV_RUN_DIR}/cloud/update-list-policies.json")"
    mv_expect_eq update "ListPolicies reflects updated displayName" "${WANT_DISPLAY_NAME}" "${UPDATED_DISPLAY_NAME}" cloud/update-list-policies.json
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DesignCenterPolicy "${POLICY_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# Audit log check for DesignCenterPolicy:
# In Cloud Audit Logs:
# 1) CreatePolicy has protoPayload.resourceName as parent applicationTemplate and
#    request.policyId as the policy ID.
# 2) UpdatePolicy and DeletePolicy have protoPayload.resourceName as the full policy resource name.
# Filtering on the policy ID and designcenter service catches all three write methods.
expect_policy_audit() {
  local phase="$1" label="$2" policy_id="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.serviceName=\"designcenter.googleapis.com\" AND \"${policy_id}\" AND timestamp>=\"${MV_STARTED_UTC}\""
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
  mv_expect_eq "${phase}" "audit log write calls on ${policy_id}" "${want}" "${got}" "${out}"
}

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DesignCenterPolicy "${POLICY_K8S}" 300
mv_wait_http delete policy-gone "${POLICY_URL}" 404
expect_policy_audit delete audit-policy "${POLICY_ID}" \
  "CreatePolicy,UpdatePolicy,DeletePolicy"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-policies GET "${TEMPLATE_URL}/policies")"
  mv_expect_http delete "REST GET list-policies returns 200" 200 "${LIST_CODE}" "cloud/delete-list-policies.json"
  FOUND_DELETED="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.policies[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-list-policies.json")"
  mv_expect_eq delete "ListPolicies no longer contains deleted policy" "" "${FOUND_DELETED}" cloud/delete-list-policies.json
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
