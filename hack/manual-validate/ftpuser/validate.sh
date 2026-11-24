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

# Human-style validation of the FTPUser controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/ftpuser/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the ftpuser-maximal test fixture)
#   FUNCTIONAL=0     skip functional checks (listing users under the FTPServer via REST)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/ftp/v1alpha1/ftpuser/ftpuser-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init FTPUser
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixture so out-of-band checks never ask KCC.
BUCKET_K8S="ftpbucket-max-${MV_UNIQUE_ID}"
SA_K8S="ftpsa-max-${MV_UNIQUE_ID}"
SERVER_K8S="ftpserv-max-${MV_UNIQUE_ID}"
USER_K8S="ftpumax-${MV_UNIQUE_ID}"
LOCATION="us-west1"

API_BASE="https://ftp.googleapis.com/v1"
PARENT="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/servers/${SERVER_K8S}"
SERVER_URL="${API_BASE}/${PARENT}"
USER_URL="${SERVER_URL}/users/${USER_K8S}?view=USER_VIEW_FULL"
USER_URL_SIMPLE="${SERVER_URL}/users/${USER_K8S}"
LIST_USERS_URL="${SERVER_URL}/users?view=USER_VIEW_FULL"
EXPECTED_USER_NAME="${PARENT}/users/${USER_K8S}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get ftpuser "${USER_K8S}" >/dev/null 2>&1 && mv_delete cleanup FTPUser "${USER_K8S}" 300
    mv_kubectl get ftpserver "${SERVER_K8S}" >/dev/null 2>&1 && mv_delete cleanup FTPServer "${SERVER_K8S}" 300
    mv_kubectl get iamserviceaccount "${SA_K8S}" >/dev/null 2>&1 && mv_delete cleanup IAMServiceAccount "${SA_K8S}" 300
    mv_kubectl get storagebucket "${BUCKET_K8S}" >/dev/null 2>&1 && mv_delete cleanup StorageBucket "${BUCKET_K8S}" 300
  fi

  local code
  code="$(mv_rest cleanup leakcheck-user GET "${USER_URL_SIMPLE}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (FTPUser gone)" PASS cloud/cleanup-leakcheck-user.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (FTPUser gone)" FAIL cloud/cleanup-leakcheck-user.json \
      "could not verify (HTTP ${code}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${USER_URL_SIMPLE}"
  else
    mv_record cleanup "no leaked GCP resources (FTPUser gone)" FAIL cloud/cleanup-leakcheck-user.json \
      "HTTP ${code}; clean up manually: curl -X DELETE -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${USER_URL_SIMPLE}"
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
mv_enable_apis setup ftp.googleapis.com iam.googleapis.com storage.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
mv_apply deps "${DEPS}"
mv_wait_ready deps StorageBucket "${BUCKET_K8S}" 300
mv_wait_ready deps IAMServiceAccount "${SA_K8S}" 300
mv_wait_ready deps FTPServer "${SERVER_K8S}" 300
mv_wait_http deps ftpserver "${SERVER_URL}" 200

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create FTPUser "${USER_K8S}" 300
SNAP="$(mv_snapshot create FTPUser "${USER_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

mv_wait_http create ftpuser "${USER_URL}" 200

# Compare spec fields against cloud response
WANT_CSA="$(jq -r '.spec.customerServiceAccountRef.name // empty' "${SNAP}")@${GCP_PROJECT_ID}.iam.gserviceaccount.com"
WANT_CRED_NAME="$(jq -r '.spec.userCredentials[0].credentialName // empty' "${SNAP}")"
WANT_CRED_KEY="$(jq -r '.spec.userCredentials[0].sshPublicKeyBody // empty' "${SNAP}" | awk '{print $1 " " $2}')"
WANT_DIR_BUCKET="$(jq -r '.spec.storageDirectoryMappings[0].bucketRef.name // empty' "${SNAP}")"
WANT_DIR_PREFIX="$(jq -r '.spec.storageDirectoryMappings[0].bucketPrefix // empty' "${SNAP}")"
WANT_DIR_PATH="$(jq -r '.spec.storageDirectoryMappings[0].directory // empty' "${SNAP}")"
WANT_LABEL_VAL="$(jq -r '.metadata.labels["test-label"] // empty' "${SNAP}")"

# Verify cloud resource name
mv_expect_jq create "cloud resource name matches expected name" "${MV_RUN_DIR}/cloud/create-ftpuser.json" '.name' "${EXPECTED_USER_NAME}"

# Verify customer service account
mv_expect_jq create "cloud customerServiceAccount matches spec" "${MV_RUN_DIR}/cloud/create-ftpuser.json" '.customerServiceAccount' "${WANT_CSA}"

# Verify credentials
mv_expect_jq create "cloud credentialName matches spec" "${MV_RUN_DIR}/cloud/create-ftpuser.json" '.userCredentials[0].credentialName' "${WANT_CRED_NAME}"
GOT_KEY="$(jq -r '.userCredentials[0].sshPublicKeyBody // empty' "${MV_RUN_DIR}/cloud/create-ftpuser.json" | awk '{print $1 " " $2}')"
mv_expect_eq create "cloud sshPublicKeyBody matches spec" "${WANT_CRED_KEY}" "${GOT_KEY}" cloud/create-ftpuser.json

# Verify directory mappings
mv_expect_jq create "cloud mapping bucket matches spec" "${MV_RUN_DIR}/cloud/create-ftpuser.json" '.storageDirectoryMappings[0].bucket' "${WANT_DIR_BUCKET}"
mv_expect_jq create "cloud mapping bucketPrefix matches spec" "${MV_RUN_DIR}/cloud/create-ftpuser.json" '.storageDirectoryMappings[0].bucketPrefix' "${WANT_DIR_PREFIX}"
mv_expect_jq create "cloud mapping directory matches spec" "${MV_RUN_DIR}/cloud/create-ftpuser.json" '.storageDirectoryMappings[0].directory' "${WANT_DIR_PATH}"

# Verify permission (READ_WRITE = 2)
GOT_PERM="$(jq -r '.storageDirectoryMappings[0].permission // empty' "${MV_RUN_DIR}/cloud/create-ftpuser.json")"
if [[ "${GOT_PERM}" == "READ_WRITE" || "${GOT_PERM}" == "2" ]]; then
  mv_record create "cloud mapping permission matches READ_WRITE" PASS cloud/create-ftpuser.json "${GOT_PERM}"
else
  mv_record create "cloud mapping permission matches READ_WRITE" FAIL cloud/create-ftpuser.json "want 'READ_WRITE' or '2', got '${GOT_PERM}'"
fi

# Verify labels
if [[ -n "${WANT_LABEL_VAL}" ]]; then
  mv_expect_jq create "cloud label 'test-label' matches metadata" "${MV_RUN_DIR}/cloud/create-ftpuser.json" '.labels["test-label"]' "${WANT_LABEL_VAL}"
fi
mv_expect_jq create "cloud label 'managed-by-cnrm' present" "${MV_RUN_DIR}/cloud/create-ftpuser.json" '.labels["managed-by-cnrm"]' "true"

# Note on gcloud CLI support:
# gcloud does not support Cloud FTP (REST API only).
mv_record create "gcloud CLI surface" INFO env.yaml "no gcloud CLI surface for Cloud FTP; verified via REST"

# Functional check: ListUsers under server out of band
if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest create list-users GET "${LIST_USERS_URL}")"
  mv_expect_http create "REST GET list-users returns 200" 200 "${LIST_CODE}" "cloud/create-list-users.json"
  FOUND_USER="$(jq -r --arg NAME "${EXPECTED_USER_NAME}" '[.users[]? | select(.name == $NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-list-users.json")"
  if [[ -n "${FOUND_USER}" ]]; then
    mv_record create "List users contains created FTPUser" PASS cloud/create-list-users.json "${FOUND_USER}"
  else
    mv_record create "List users contains created FTPUser" FAIL cloud/create-list-users.json "user ${EXPECTED_USER_NAME} not found"
  fi
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update FTPUser "${USER_K8S}" 300
SNAP="$(mv_snapshot update FTPUser "${USER_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update ftpuser "${USER_URL}" 200

# Verify changed fields
WANT_UPD_CRED_NAME="$(jq -r '.spec.userCredentials[0].credentialName // empty' "${SNAP}")"
WANT_UPD_CRED_KEY="$(jq -r '.spec.userCredentials[0].sshPublicKeyBody // empty' "${SNAP}" | awk '{print $1 " " $2}')"
WANT_UPD_DIR_PREFIX="$(jq -r '.spec.storageDirectoryMappings[0].bucketPrefix // empty' "${SNAP}")"
WANT_UPD_DIR_PATH="$(jq -r '.spec.storageDirectoryMappings[0].directory // empty' "${SNAP}")"
WANT_UPD_LABEL_VAL="$(jq -r '.metadata.labels["test-label"] // empty' "${SNAP}")"

mv_expect_jq update "cloud credentialName updated to new spec" "${MV_RUN_DIR}/cloud/update-ftpuser.json" '.userCredentials[0].credentialName' "${WANT_UPD_CRED_NAME}"
GOT_UPD_KEY="$(jq -r '.userCredentials[0].sshPublicKeyBody // empty' "${MV_RUN_DIR}/cloud/update-ftpuser.json" | awk '{print $1 " " $2}')"
mv_expect_eq update "cloud sshPublicKeyBody updated to new spec" "${WANT_UPD_CRED_KEY}" "${GOT_UPD_KEY}" cloud/update-ftpuser.json

mv_expect_jq update "cloud mapping bucketPrefix updated to new spec" "${MV_RUN_DIR}/cloud/update-ftpuser.json" '.storageDirectoryMappings[0].bucketPrefix' "${WANT_UPD_DIR_PREFIX}"
mv_expect_jq update "cloud mapping directory updated to new spec" "${MV_RUN_DIR}/cloud/update-ftpuser.json" '.storageDirectoryMappings[0].directory' "${WANT_UPD_DIR_PATH}"

GOT_UPD_PERM="$(jq -r '.storageDirectoryMappings[0].permission // empty' "${MV_RUN_DIR}/cloud/update-ftpuser.json")"
if [[ "${GOT_UPD_PERM}" == "READ_ONLY" || "${GOT_UPD_PERM}" == "1" ]]; then
  mv_record update "cloud mapping permission updated to READ_ONLY" PASS cloud/update-ftpuser.json "${GOT_UPD_PERM}"
else
  mv_record update "cloud mapping permission updated to READ_ONLY" FAIL cloud/update-ftpuser.json "want 'READ_ONLY' or '1', got '${GOT_UPD_PERM}'"
fi

if [[ -n "${WANT_UPD_LABEL_VAL}" ]]; then
  mv_expect_jq update "cloud label 'test-label' updated to new metadata" "${MV_RUN_DIR}/cloud/update-ftpuser.json" '.labels["test-label"]' "${WANT_UPD_LABEL_VAL}"
fi

# Verify unchanged fields
mv_expect_jq update "cloud mapping bucket unchanged" "${MV_RUN_DIR}/cloud/update-ftpuser.json" '.storageDirectoryMappings[0].bucket' "${WANT_DIR_BUCKET}"
mv_expect_jq update "cloud customerServiceAccount unchanged" "${MV_RUN_DIR}/cloud/update-ftpuser.json" '.customerServiceAccount' "${WANT_CSA}"

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-ftpuser.json")"
UPDATE_CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/update-ftpuser.json")"
mv_expect_eq update "createTime preserved across update (in-place update)" "${CREATE_T}" "${UPDATE_CREATE_T}" cloud/update-ftpuser.json

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest update list-users GET "${LIST_USERS_URL}")"
  mv_expect_http update "REST GET list-users returns 200 after update" 200 "${LIST_CODE}" "cloud/update-list-users.json"
  UPDATED_DIR="$(jq -r --arg NAME "${EXPECTED_USER_NAME}" '[.users[]? | select(.name == $NAME)][0].storageDirectoryMappings[0].directory // empty' "${MV_RUN_DIR}/cloud/update-list-users.json")"
  mv_expect_eq update "List users reflects updated directory mapping" "${WANT_UPD_DIR_PATH}" "${UPDATED_DIR}" cloud/update-list-users.json
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady FTPUser "${USER_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete FTPUser "${USER_K8S}" 300
mv_wait_http delete user-gone "${USER_URL_SIMPLE}" 404

if [[ "${FUNCTIONAL}" == "1" ]]; then
  LIST_CODE="$(mv_rest delete list-users GET "${LIST_USERS_URL}")"
  mv_expect_http delete "REST GET list-users returns 200 after delete" 200 "${LIST_CODE}" "cloud/delete-list-users.json"
  FOUND_DELETED="$(jq -r --arg NAME "${EXPECTED_USER_NAME}" '[.users[]? | select(.name == $NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-list-users.json")"
  mv_expect_eq delete "List users no longer contains deleted FTPUser" "" "${FOUND_DELETED}" cloud/delete-list-users.json
fi

# Cloud Audit Logs check:
# Verifies the lifecycle calls received by GCP for this FTPUser:
# Exactly CreateUser, UpdateUser, and DeleteUser (no spurious writes, no recreations).
# Cloud FTP produces pairs of start/finish audit log entries for asynchronous LRO calls.
expect_ftpuser_audit() {
  local phase="$1" label="$2" substr="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.resourceName:\"${substr}\" AND timestamp>=\"${MV_STARTED_UTC}\""
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    # Extract distinct consecutive short method names, or filter by operation.last==true
    # In Cloud FTP, each LRO operation emits start and completion audit logs.
    got="$(jq -r 'reduce ([.[].protoPayload.methodName | split(".") | last][]? // empty) as $item ([]; if .[-1] != $item then . + [$item] else . end) | join(",")' "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
    [[ "${got}" == "${want}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 15
  done
  echo "filter: ${filter}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.filter.txt"
  jq -r '.[] | "\(.timestamp)  \(.protoPayload.methodName)  \(.protoPayload.authenticationInfo.principalEmail // "-")"' \
    "${MV_RUN_DIR}/${out}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.summary.txt" 2>/dev/null || true
  mv_trace "gcloud logging read '${filter}' --project ${GCP_PROJECT_ID} --order=asc --format=json" \
    "$(jq 'length' "${MV_RUN_DIR}/${out}" 2>/dev/null || echo 0) entries: ${got:-none}" "cloud/${phase}-${label}.summary.txt"
  mv_expect_eq "${phase}" "audit log write calls on ${substr##*/}" "${want}" "${got}" "${out}"
}

expect_ftpuser_audit delete audit-user "users/${USER_K8S}" "CreateUser,UpdateUser,DeleteUser"

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown FTPServer "${SERVER_K8S}" 300
mv_wait_http teardown server-gone "${SERVER_URL}" 404
mv_delete teardown IAMServiceAccount "${SA_K8S}" 300
mv_delete teardown StorageBucket "${BUCKET_K8S}" 300

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
