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

# Human-style validation of the MemorystoreACLPolicy controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/memorystoreaclpolicy/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the memorystoreaclpolicy-maximal test fixture)
#   FUNCTIONAL=0     skip functional checks (listing policies and revisions)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/memorystore/v1alpha1/memorystoreaclpolicy/memorystoreaclpolicy-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init MemorystoreACLPolicy
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
POLICY_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
POLICY_ID="$(grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
POLICY_ID="${POLICY_ID:-${POLICY_K8S}}"
LOCATION="$(grep -m1 '^[[:space:]]*location:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
LOCATION="${LOCATION:-us-central1}"

API_BASE="https://memorystore.googleapis.com/v1"
PARENT="projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
POLICY_URL="${API_BASE}/${PARENT}/aclPolicies/${POLICY_ID}"
EXPECTED_NAME="${PARENT}/aclPolicies/${POLICY_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get memorystoreaclpolicy "${POLICY_K8S}" >/dev/null 2>&1 && mv_delete cleanup MemorystoreACLPolicy "${POLICY_K8S}" 300
    if [[ -n "${DEPS:-}" ]]; then
      # Delete dependencies in reverse order if any exist
      :
    fi
  fi

  local code
  code="$(mv_rest cleanup leakcheck-policy GET "${POLICY_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (aclPolicy gone)" PASS cloud/cleanup-leakcheck-policy.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (aclPolicy gone)" FAIL cloud/cleanup-leakcheck-policy.json \
      "could not verify (HTTP ${code}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${POLICY_URL}"
  else
    mv_record cleanup "no leaked GCP resources (aclPolicy gone)" FAIL cloud/cleanup-leakcheck-policy.json \
      "HTTP ${code}; clean up manually: gcloud memorystore acl-policies delete ${POLICY_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID} --quiet"
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
# Every API the controller, dependencies and out-of-band checks call.
mv_enable_apis setup memorystore.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create MemorystoreACLPolicy "${POLICY_K8S}" 300
SNAP="$(mv_snapshot create MemorystoreACLPolicy "${POLICY_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

mv_wait_http create policy "${POLICY_URL}" 200

# Compare spec fields against cloud response
mv_expect_jq create "cloud resource name matches expected ID" "${MV_RUN_DIR}/cloud/create-policy.json" '.name' "${EXPECTED_NAME}"

WANT_RULES="$(jq -S -c '[.spec.rules[]? | {username: .username, rule: .rule}] | sort_by(.username)' "${SNAP}")"
GOT_RULES="$(jq -S -c '[.rules[]? | {username: .username, rule: .rule}] | sort_by(.username)' "${MV_RUN_DIR}/cloud/create-policy.json")"
mv_expect_eq create "cloud rules match spec" "${WANT_RULES}" "${GOT_RULES}" cloud/create-policy.json

# Describe via gcloud CLI out of band
mv_cli create gcloud-describe -- gcloud memorystore acl-policies describe "${POLICY_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json && \
  mv_record create "gcloud memorystore acl-policies describe succeeds" PASS cloud/create-gcloud-describe.txt "" || \
  mv_record create "gcloud memorystore acl-policies describe succeeds" FAIL cloud/create-gcloud-describe.txt "gcloud describe failed"

# Functional check: ListAclPolicies and list revisions out of band
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create list-policies GET "${API_BASE}/${PARENT}/aclPolicies")"
  mv_expect_http create "REST GET list-policies returns 200" 200 "${LIST_CODE}" "cloud/create-list-policies.json"
  FOUND_NAME="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.aclPolicies[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-policies.json")"
  mv_expect_eq create "ListAclPolicies returns created policy" "${EXPECTED_NAME}" "${FOUND_NAME}" cloud/create-list-policies.json

  REV_CODE="$(mv_rest create list-revisions GET "${POLICY_URL}/revisions")"
  mv_expect_http create "REST GET list-revisions returns 200" 200 "${REV_CODE}" "cloud/create-list-revisions.json"
  CREATE_REV_COUNT="$(jq '.aclPolicyRevisions // [] | length' "${MV_RUN_DIR}/cloud/create-list-revisions.json")"
  mv_record create "initial revision created on create" INFO cloud/create-list-revisions.json "count=${CREATE_REV_COUNT}"
  mv_expect_eq create "at least 1 revision exists after create" "1" "$(( CREATE_REV_COUNT >= 1 ? 1 : 0 ))" cloud/create-list-revisions.json

  mv_cli create gcloud-revisions -- gcloud memorystore acl-policies revisions list --acl-policy="${POLICY_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}"
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update MemorystoreACLPolicy "${POLICY_K8S}" 300
SNAP="$(mv_snapshot update MemorystoreACLPolicy "${POLICY_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update policy "${POLICY_URL}" 200

# Compare updated spec fields against cloud response
WANT_RULES_UPDATE="$(jq -S -c '[.spec.rules[]? | {username: .username, rule: .rule}] | sort_by(.username)' "${SNAP}")"
GOT_RULES_UPDATE="$(jq -S -c '[.rules[]? | {username: .username, rule: .rule}] | sort_by(.username)' "${MV_RUN_DIR}/cloud/update-policy.json")"
mv_expect_eq update "cloud rules updated to new spec" "${WANT_RULES_UPDATE}" "${GOT_RULES_UPDATE}" cloud/update-policy.json

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-policy.json")"
UPDATE_T="$(jq -r '.updateTime // empty' "${MV_RUN_DIR}/cloud/update-policy.json")"
mv_record update "createTime preserved, updateTime updated" INFO cloud/update-policy.json "createTime=${CREATE_T} updateTime=${UPDATE_T}"

mv_cli update gcloud-describe -- gcloud memorystore acl-policies describe "${POLICY_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json && \
  mv_record update "gcloud memorystore acl-policies describe succeeds after update" PASS cloud/update-gcloud-describe.txt "" || \
  mv_record update "gcloud memorystore acl-policies describe succeeds after update" FAIL cloud/update-gcloud-describe.txt "gcloud describe failed"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  REV_CODE="$(mv_rest update list-revisions GET "${POLICY_URL}/revisions")"
  mv_expect_http update "REST GET list-revisions returns 200 after update" 200 "${REV_CODE}" "cloud/update-list-revisions.json"
  UPDATE_REV_COUNT="$(jq '.aclPolicyRevisions // [] | length' "${MV_RUN_DIR}/cloud/update-list-revisions.json")"
  mv_record update "new revision created on update" INFO cloud/update-list-revisions.json "count=${UPDATE_REV_COUNT} (was ${CREATE_REV_COUNT})"
  mv_expect_eq update "revision count increased after update" "1" "$(( UPDATE_REV_COUNT > CREATE_REV_COUNT ? 1 : 0 ))" cloud/update-list-revisions.json

  mv_cli update gcloud-revisions -- gcloud memorystore acl-policies revisions list --acl-policy="${POLICY_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}"
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady MemorystoreACLPolicy "${POLICY_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# Audit log check:
# Reads Cloud Audit Logs (admin activity) for write calls on the policy resource.
# In Cloud Audit Logs, UpdateAclPolicy and DeleteAclPolicy are LROs and write both start
# and completion entries (operation.last == true). We filter out operation.last to match
# the API calls initiated by KCC.
expect_aclpolicy_audit() {
  local phase="$1" label="$2" policy_id="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND \"${policy_id}\" AND timestamp>=\"${MV_STARTED_UTC}\""
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
  mv_expect_eq "${phase}" "audit log write calls on ${policy_id}" "${want}" "${got}" "${out}"
}

# --- Phase: delete -----------------------------------------------------------
mv_delete delete MemorystoreACLPolicy "${POLICY_K8S}" 300
mv_wait_http delete policy-gone "${POLICY_URL}" 404
expect_aclpolicy_audit delete audit-policy "${POLICY_ID}" \
  "CreateAclPolicy,UpdateAclPolicy,DeleteAclPolicy"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-policies GET "${API_BASE}/${PARENT}/aclPolicies")"
  mv_expect_http delete "REST GET list-policies returns 200" 200 "${LIST_CODE}" "cloud/delete-list-policies.json"
  FOUND_DELETED="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.aclPolicies[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-list-policies.json")"
  mv_expect_eq delete "ListAclPolicies no longer contains deleted policy" "" "${FOUND_DELETED}" cloud/delete-list-policies.json
fi

# --- Phase: teardown ---------------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  # Teardown dependencies in reverse order if any exist
  :
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
