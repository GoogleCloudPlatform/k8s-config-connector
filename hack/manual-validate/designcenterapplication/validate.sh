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

# Human-style validation of the DesignCenterApplication controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/designcenterapplication/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the designcenterapplication-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (listing applications)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/designcenter/v1alpha1/designcenterapplication/designcenterapplication-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DesignCenterApplication
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
SPACE_K8S="dc-space-${MV_UNIQUE_ID}"
APPLICATION_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
SPACE_ID="${SPACE_K8S}"
APPLICATION_ID="${APPLICATION_K8S}"
LOCATION="us-central1"

DC="https://designcenter.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
SPACE_URL="${DC}/spaces/${SPACE_ID}"
APPLICATION_URL="${SPACE_URL}/applications/${APPLICATION_ID}"
EXPECTED_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/spaces/${SPACE_ID}/applications/${APPLICATION_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get designcenterapplication "${APPLICATION_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterApplication "${APPLICATION_K8S}" 300
    mv_kubectl get designcenterspace "${SPACE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterSpace "${SPACE_K8S}" 300
  fi
  local code_app code_space
  code_app="$(mv_rest cleanup leakcheck-application GET "${APPLICATION_URL}")"
  if [[ "${code_app}" == "404" ]]; then
    mv_record cleanup "no leaked GCP application" PASS cloud/cleanup-leakcheck-application.json "HTTP 404"
  elif [[ "${code_app}" != "200" ]]; then
    mv_record cleanup "no leaked GCP application" FAIL cloud/cleanup-leakcheck-application.json \
      "could not verify (HTTP ${code_app}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${APPLICATION_URL}"
  else
    mv_record cleanup "no leaked GCP application" FAIL cloud/cleanup-leakcheck-application.json \
      "HTTP ${code_app}; clean up manually: DELETE ${APPLICATION_URL}"
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
mv_enable_apis setup designcenter.googleapis.com apphub.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
mv_apply deps "${DEPS}"
mv_wait_ready deps DesignCenterSpace "${SPACE_K8S}" 300
mv_wait_http deps space "${SPACE_URL}" 200

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create DesignCenterApplication "${APPLICATION_K8S}" 300
SNAP="$(mv_snapshot create DesignCenterApplication "${APPLICATION_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_DEPLOYMENT_REGION="$(jq -r '.spec.deploymentRegion // empty' "${SNAP}")"
WANT_SCOPE_TYPE="$(jq -r '.spec.scope.type // empty' "${SNAP}")"
WANT_SOURCE_URI="$(jq -r '.spec.source.sharedTemplateRevisionURI // empty' "${SNAP}")"

mv_wait_http create application "${APPLICATION_URL}" 200

# Compare spec fields against cloud response
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-application.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-application.json" '.description' "${WANT_DESCRIPTION}"
fi
if [[ -n "${WANT_DEPLOYMENT_REGION}" ]]; then
  mv_expect_jq create "cloud deploymentRegion matches spec" "${MV_RUN_DIR}/cloud/create-application.json" '.deploymentRegion' "${WANT_DEPLOYMENT_REGION}"
fi
if [[ -n "${WANT_SCOPE_TYPE}" ]]; then
  mv_expect_jq create "cloud scope.type matches spec" "${MV_RUN_DIR}/cloud/create-application.json" \
    'if .scope.type == 1 or .scope.type == "REGIONAL" then "REGIONAL" else (.scope.type|tostring) end' \
    "${WANT_SCOPE_TYPE}"
fi
if [[ -n "${WANT_SOURCE_URI}" ]]; then
  mv_expect_jq create "cloud source sharedTemplateRevisionUri matches spec" "${MV_RUN_DIR}/cloud/create-application.json" \
    '(.source.sharedTemplateRevisionUri // .source.sharedTemplateRevisionURI // empty)' "${WANT_SOURCE_URI}"
fi

mv_expect_jq create "cloud resource name matches expected ID" "${MV_RUN_DIR}/cloud/create-application.json" '.name' "${EXPECTED_NAME}"

# Out of band gcloud describe
if mv_cli create gcloud-describe -- gcloud alpha design-center spaces applications describe "${APPLICATION_ID}" --space="${SPACE_ID}" --project="${GCP_PROJECT_ID}" --location="${LOCATION}"; then
  if grep -q "${EXPECTED_NAME}" "${MV_RUN_DIR}/cloud/create-gcloud-describe.txt"; then
    mv_record create "gcloud alpha describe returns created application" PASS cloud/create-gcloud-describe.txt ""
  else
    mv_record create "gcloud alpha describe returns created application" FAIL cloud/create-gcloud-describe.txt "application name not found in gcloud output"
  fi
else
  mv_record create "gcloud alpha describe command executed" INFO cloud/create-gcloud-describe.txt "gcloud alpha command returned non-zero"
fi

# Functional check: ListApplications contains the newly created application
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create list-applications GET "${SPACE_URL}/applications")"
  mv_expect_http create "REST GET list-applications returns 200" 200 "${LIST_CODE}" "cloud/create-list-applications.json"
  FOUND_NAME="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.applications[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-applications.json")"
  mv_expect_eq create "ListApplications returns created application" "${EXPECTED_NAME}" "${FOUND_NAME}" cloud/create-list-applications.json
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DesignCenterApplication "${APPLICATION_K8S}" 300
SNAP="$(mv_snapshot update DesignCenterApplication "${APPLICATION_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_DEPLOYMENT_REGION="$(jq -r '.spec.deploymentRegion // empty' "${SNAP}")"
WANT_SCOPE_TYPE="$(jq -r '.spec.scope.type // empty' "${SNAP}")"
WANT_SOURCE_URI="$(jq -r '.spec.source.sharedTemplateRevisionURI // empty' "${SNAP}")"

mv_wait_http update application "${APPLICATION_URL}" 200

if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-application.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-application.json" '.description' "${WANT_DESCRIPTION}"
fi
if [[ -n "${WANT_DEPLOYMENT_REGION}" ]]; then
  mv_expect_jq update "cloud deploymentRegion still matches spec" "${MV_RUN_DIR}/cloud/update-application.json" '.deploymentRegion' "${WANT_DEPLOYMENT_REGION}"
fi
if [[ -n "${WANT_SCOPE_TYPE}" ]]; then
  mv_expect_jq update "cloud scope.type still matches spec" "${MV_RUN_DIR}/cloud/update-application.json" \
    'if .scope.type == 1 or .scope.type == "REGIONAL" then "REGIONAL" else (.scope.type|tostring) end' \
    "${WANT_SCOPE_TYPE}"
fi
if [[ -n "${WANT_SOURCE_URI}" ]]; then
  mv_expect_jq update "cloud source sharedTemplateRevisionUri still matches spec" "${MV_RUN_DIR}/cloud/update-application.json" \
    '(.source.sharedTemplateRevisionUri // .source.sharedTemplateRevisionURI // empty)' "${WANT_SOURCE_URI}"
fi

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-application.json")"
UPDATE_T="$(jq -r '.updateTime // empty' "${MV_RUN_DIR}/cloud/update-application.json")"
mv_record update "createTime preserved, updateTime updated" INFO cloud/update-application.json "createTime=${CREATE_T} updateTime=${UPDATE_T}"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update list-applications GET "${SPACE_URL}/applications")"
  mv_expect_http update "REST GET list-applications returns 200" 200 "${LIST_CODE}" "cloud/update-list-applications.json"
  if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
    UPDATED_DISPLAY_NAME="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.applications[]? | select(.name==$NAME)][0].displayName // empty' "${MV_RUN_DIR}/cloud/update-list-applications.json")"
    mv_expect_eq update "ListApplications reflects updated displayName" "${WANT_DISPLAY_NAME}" "${UPDATED_DISPLAY_NAME}" cloud/update-list-applications.json
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DesignCenterApplication "${APPLICATION_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# Audit log check for DesignCenterApplication:
# In Cloud Audit Logs:
# 1) CreateApplication has protoPayload.resourceName as parent space and
#    request.applicationId as the application ID.
# 2) UpdateApplication and DeleteApplication have protoPayload.resourceName
#    as the full application resource name.
# Filtering on the application ID catches all three methods.
expect_application_audit() {
  local phase="$1" label="$2" app_id="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.serviceName=\"designcenter.googleapis.com\" AND \"${app_id}\" AND timestamp>=\"${MV_STARTED_UTC}\""
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
  mv_expect_eq "${phase}" "audit log write calls on ${app_id}" "${want}" "${got}" "${out}"
}

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DesignCenterApplication "${APPLICATION_K8S}" 300
mv_wait_http delete application-gone "${APPLICATION_URL}" 404
expect_application_audit delete audit-application "${APPLICATION_ID}" \
  "CreateApplication,UpdateApplication,DeleteApplication"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-applications GET "${SPACE_URL}/applications")"
  mv_expect_http delete "REST GET list-applications returns 200" 200 "${LIST_CODE}" "cloud/delete-list-applications.json"
  FOUND_DELETED="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.applications[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-list-applications.json")"
  mv_expect_eq delete "ListApplications no longer contains deleted application" "" "${FOUND_DELETED}" cloud/delete-list-applications.json
fi

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown DesignCenterSpace "${SPACE_K8S}" 300
mv_wait_http teardown space-gone "${SPACE_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
