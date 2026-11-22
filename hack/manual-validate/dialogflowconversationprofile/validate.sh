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

# Human-style validation of the DialogflowConversationProfile controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/dialogflowconversationprofile/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the dialogflowconversationprofile-maximal test fixture)
#   FUNCTIONAL=0     skip the functional check (listing conversation profiles)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/dialogflow/v1alpha1/dialogflowconversationprofile/dialogflowconversationprofile-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DialogflowConversationProfile
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS=""
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
PROFILE_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
LOCATION="$(grep -m1 '^[[:space:]]*location:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
LOCATION="${LOCATION:-global}"

if [[ "${LOCATION}" == "global" ]]; then
  DF="https://dialogflow.googleapis.com/v2/projects/${GCP_PROJECT_ID}/locations/global"
else
  DF="https://${LOCATION}-dialogflow.googleapis.com/v2/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
fi
LIST_URL="${DF}/conversationProfiles"
PROFILE_URL=""
PROFILE_ID=""
CLOUD_NAME=""

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get dialogflowconversationprofile "${PROFILE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DialogflowConversationProfile "${PROFILE_K8S}" 300
  fi
  local code
  if [[ -n "${PROFILE_URL:-}" ]]; then
    code="$(mv_rest cleanup leakcheck-profile GET "${PROFILE_URL}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (profile gone)" PASS cloud/cleanup-leakcheck-profile.json "HTTP 404"
    elif [[ "${code}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (profile gone)" FAIL cloud/cleanup-leakcheck-profile.json \
        "could not verify (HTTP ${code}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${PROFILE_URL}"
    else
      mv_record cleanup "no leaked GCP resources (profile gone)" FAIL cloud/cleanup-leakcheck-profile.json \
        "HTTP ${code}; clean up manually: DELETE ${PROFILE_URL}"
    fi
  else
    code="$(mv_rest cleanup leakcheck-list GET "${LIST_URL}")"
    if [[ "${code}" == "200" ]]; then
      local leaked
      leaked="$(jq -r --arg uid "${MV_UNIQUE_ID}" '[.conversationProfiles[]? | select(.displayName | contains($uid))] | length' "${MV_RUN_DIR}/cloud/cleanup-leakcheck-list.json" 2>/dev/null || echo 0)"
      if [[ "${leaked}" == "0" ]]; then
        mv_record cleanup "no leaked GCP resources (no profiles with uniqueId in displayName)" PASS cloud/cleanup-leakcheck-list.json "none found"
      else
        mv_record cleanup "no leaked GCP resources (no profiles with uniqueId in displayName)" FAIL cloud/cleanup-leakcheck-list.json \
          "found ${leaked} leaked profiles in ${LIST_URL}"
      fi
    else
      mv_record cleanup "no leaked GCP resources" FAIL cloud/cleanup-leakcheck-list.json "could not list profiles: HTTP ${code}"
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

# Custom audit log check for DialogflowConversationProfile:
# Cloud Audit Logs record CreateConversationProfile, UpdateConversationProfile, and DeleteConversationProfile.
expect_profile_audit() {
  local phase="$1" label="$2" profile_id="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND \"${profile_id}\" AND timestamp>=\"${MV_STARTED_UTC}\""
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    got="$(jq -r '[.[].protoPayload.methodName | split(".") | last] | join(",")' "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
    [[ "${got}" == "${want}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 15
  done
  echo "filter: ${filter}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.filter.txt"
  mv_expect_eq "${phase}" "audit log write calls on ${profile_id}" "${want}" "${got}" "${out}"
}

# ---------------------------------------------------------------------------
# Every API the controller, the dependencies and the out-of-band checks call.
mv_enable_apis setup dialogflow.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
else
  mv_record deps "dependencies" INFO "" "none required for DialogflowConversationProfile"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create DialogflowConversationProfile "${PROFILE_K8S}" 300
SNAP="$(mv_snapshot create DialogflowConversationProfile "${PROFILE_K8S}")"
STATUS_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${STATUS_REF:-<none>}"

WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_LANGUAGE_CODE="$(jq -r '.spec.languageCode // empty' "${SNAP}")"
WANT_LOGGING="$(jq -r '.spec.loggingConfig.enableStackdriverLogging // false | tostring' "${SNAP}")"
WANT_TIME_ZONE="$(jq -r '.spec.timeZone // empty' "${SNAP}")"

# Find the server-generated profile out-of-band via ListConversationProfiles
start=$(date +%s)
while :; do
  mv_rest create list-profiles GET "${LIST_URL}" >/dev/null
  CLOUD_NAME="$(jq -r --arg dn "${WANT_DISPLAY_NAME}" '[.conversationProfiles[]? | select(.displayName==$dn)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-profiles.json")"
  [[ -n "${CLOUD_NAME}" ]] && break
  (($(date +%s) - start > 60)) && break
  sleep 3
done

if [[ -n "${CLOUD_NAME}" ]]; then
  mv_record create "ListConversationProfiles contains profile matching displayName" PASS cloud/create-list-profiles.json "${CLOUD_NAME}"
  PROFILE_ID="${CLOUD_NAME##*/}"
  PROFILE_URL="https://dialogflow.googleapis.com/v2/${CLOUD_NAME}"
else
  mv_record create "ListConversationProfiles contains profile matching displayName" FAIL cloud/create-list-profiles.json "not found in ${LIST_URL}"
  # Fall back to status.externalRef if listing didn't find it yet
  if [[ -n "${STATUS_REF}" ]]; then
    CLOUD_NAME="${STATUS_REF}"
    PROFILE_ID="${CLOUD_NAME##*/}"
    PROFILE_URL="https://dialogflow.googleapis.com/v2/${CLOUD_NAME}"
  fi
fi

if [[ -n "${STATUS_REF}" && -n "${CLOUD_NAME}" ]]; then
  mv_expect_eq create "status.externalRef matches cloud resource name" "${CLOUD_NAME}" "${STATUS_REF}" "${SNAP#"${MV_RUN_DIR}"/}"
fi

mv_wait_http create profile "${PROFILE_URL}" 200

# Compare spec fields against cloud response
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-profile.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_LANGUAGE_CODE}" ]]; then
  mv_expect_jq create "cloud languageCode matches spec" "${MV_RUN_DIR}/cloud/create-profile.json" '.languageCode' "${WANT_LANGUAGE_CODE}"
fi
if [[ -n "${WANT_LOGGING}" ]]; then
  mv_expect_jq create "cloud loggingConfig.enableStackdriverLogging matches spec" "${MV_RUN_DIR}/cloud/create-profile.json" \
    '.loggingConfig.enableStackdriverLogging // false | tostring' "${WANT_LOGGING}"
fi
if [[ -n "${WANT_TIME_ZONE}" ]]; then
  mv_expect_jq create "cloud timeZone matches spec" "${MV_RUN_DIR}/cloud/create-profile.json" '.timeZone' "${WANT_TIME_ZONE}"
fi

mv_record create "gcloud CLI surface" INFO "cloud/create-profile.json" "no gcloud CLI command for dialogflow conversation-profiles; REST GET used out of band"

# Functional check: ListConversationProfiles contains the created profile
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create verify-list-profiles GET "${LIST_URL}")"
  mv_expect_http create "REST GET verify-list-profiles returns 200" 200 "${LIST_CODE}" "cloud/create-verify-list-profiles.json"
  FOUND_NAME="$(jq -r --arg NAME "${CLOUD_NAME}" '[.conversationProfiles[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-verify-list-profiles.json")"
  mv_expect_eq create "ListConversationProfiles returns created profile" "${CLOUD_NAME}" "${FOUND_NAME}" cloud/create-verify-list-profiles.json
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DialogflowConversationProfile "${PROFILE_K8S}" 300
SNAP="$(mv_snapshot update DialogflowConversationProfile "${PROFILE_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_UPDATED_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_UPDATED_LANGUAGE_CODE="$(jq -r '.spec.languageCode // empty' "${SNAP}")"
WANT_UPDATED_LOGGING="$(jq -r '.spec.loggingConfig.enableStackdriverLogging // false | tostring' "${SNAP}")"

mv_wait_http update profile "${PROFILE_URL}" 200

if [[ -n "${WANT_UPDATED_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-profile.json" '.displayName' "${WANT_UPDATED_DISPLAY_NAME}"
fi
if [[ -n "${WANT_UPDATED_LANGUAGE_CODE}" ]]; then
  mv_expect_jq update "cloud languageCode matches spec" "${MV_RUN_DIR}/cloud/update-profile.json" '.languageCode' "${WANT_UPDATED_LANGUAGE_CODE}"
fi
if [[ -n "${WANT_UPDATED_LOGGING}" ]]; then
  mv_expect_jq update "cloud loggingConfig.enableStackdriverLogging updated to new spec" "${MV_RUN_DIR}/cloud/update-profile.json" \
    '.loggingConfig.enableStackdriverLogging // false | tostring' "${WANT_UPDATED_LOGGING}"
fi

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-profile.json")"
UPDATE_T="$(jq -r '.updateTime // empty' "${MV_RUN_DIR}/cloud/update-profile.json")"
mv_record update "createTime preserved, updateTime updated" INFO cloud/update-profile.json "createTime=${CREATE_T} updateTime=${UPDATE_T}"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update verify-list-profiles GET "${LIST_URL}")"
  mv_expect_http update "REST GET verify-list-profiles returns 200" 200 "${LIST_CODE}" "cloud/update-verify-list-profiles.json"
  if [[ -n "${WANT_UPDATED_DISPLAY_NAME}" ]]; then
    UPDATED_DISPLAY_NAME="$(jq -r --arg NAME "${CLOUD_NAME}" '[.conversationProfiles[]? | select(.name==$NAME)][0].displayName // empty' "${MV_RUN_DIR}/cloud/update-verify-list-profiles.json")"
    mv_expect_eq update "ListConversationProfiles reflects updated displayName" "${WANT_UPDATED_DISPLAY_NAME}" "${UPDATED_DISPLAY_NAME}" cloud/update-verify-list-profiles.json
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DialogflowConversationProfile "${PROFILE_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DialogflowConversationProfile "${PROFILE_K8S}" 300
mv_wait_http delete profile-gone "${PROFILE_URL}" 404
expect_profile_audit delete audit-profile "${PROFILE_ID}" \
  "CreateConversationProfile,UpdateConversationProfile,DeleteConversationProfile"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete verify-list-profiles GET "${LIST_URL}")"
  mv_expect_http delete "REST GET verify-list-profiles returns 200" 200 "${LIST_CODE}" "cloud/delete-verify-list-profiles.json"
  FOUND_DELETED="$(jq -r --arg NAME "${CLOUD_NAME}" '[.conversationProfiles[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-verify-list-profiles.json")"
  mv_expect_eq delete "ListConversationProfiles no longer contains deleted profile" "" "${FOUND_DELETED}" cloud/delete-verify-list-profiles.json
fi

# --- Phase: dependency teardown ----------------------------------------------
if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
