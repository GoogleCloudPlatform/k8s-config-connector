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
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/ftp/v1alpha1/ftpuser/ftpuser-maximal}"

mv_init FTPUser
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Names derived from the fixture (resourceIDs) so out-of-band checks never ask KCC.
BUCKET_K8S="ftpbucket-max-${MV_UNIQUE_ID}"
SA_K8S="ftpsa-max-${MV_UNIQUE_ID}"
SERVER_K8S="ftpserv-max-${MV_UNIQUE_ID}"
USER_K8S="ftpumax-${MV_UNIQUE_ID}"
BUCKET_ID="${BUCKET_K8S}"
SA_EMAIL="${SA_K8S}@${GCP_PROJECT_ID}.iam.gserviceaccount.com"
SERVER_ID="${SERVER_K8S}"
USER_ID="${USER_K8S}"
LOCATION="us-west1"

GCS_URL="https://storage.googleapis.com/storage/v1/b/${BUCKET_ID}"
IAM_URL="https://iam.googleapis.com/v1/projects/${GCP_PROJECT_ID}/serviceAccounts/${SA_EMAIL}"
FTP_SERVER_URL="https://ftp.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}/servers/${SERVER_ID}"
FTP_USER_URL="${FTP_SERVER_URL}/users/${USER_ID}?view=USER_VIEW_FULL"
FTP_USER_CLEANUP_URL="${FTP_SERVER_URL}/users/${USER_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get ftpuser "${USER_K8S}" >/dev/null 2>&1 && mv_delete cleanup FTPUser "${USER_K8S}" 300
    mv_kubectl get ftpserver "${SERVER_K8S}" >/dev/null 2>&1 && mv_delete cleanup FTPServer "${SERVER_K8S}" 600
    mv_kubectl get iamserviceaccount "${SA_K8S}" >/dev/null 2>&1 && mv_delete cleanup IAMServiceAccount "${SA_K8S}" 300
    mv_kubectl get storagebucket "${BUCKET_K8S}" >/dev/null 2>&1 && mv_delete cleanup StorageBucket "${BUCKET_K8S}" 300
  fi
  local code
  code="$(mv_rest cleanup leakcheck-user GET "${FTP_USER_CLEANUP_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (FTPUser gone)" PASS cloud/cleanup-leakcheck-user.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (FTPUser gone)" FAIL cloud/cleanup-leakcheck-user.json \
      "could not verify (HTTP ${code}); check manually via REST: ${FTP_USER_CLEANUP_URL}"
  else
    mv_record cleanup "no leaked GCP resources (FTPUser gone)" FAIL cloud/cleanup-leakcheck-user.json \
      "HTTP ${code}; clean up manually via REST DELETE: ${FTP_USER_CLEANUP_URL}"
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
mv_enable_apis setup ftp.googleapis.com storage.googleapis.com iam.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
mv_apply deps "${DEPS}"
mv_wait_ready deps StorageBucket "${BUCKET_K8S}" 300
mv_wait_ready deps IAMServiceAccount "${SA_K8S}" 300
mv_wait_ready deps FTPServer "${SERVER_K8S}" 600
mv_wait_http deps bucket "${GCS_URL}" 200
mv_wait_http deps serviceaccount "${IAM_URL}" 200
mv_wait_http deps ftpserver "${FTP_SERVER_URL}" 200

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create FTPUser "${USER_K8S}" 300
SNAP="$(mv_snapshot create FTPUser "${USER_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

WANT_SA="$(jq -r '.spec.customerServiceAccountRef.name // empty' "${SNAP}")"
WANT_LABEL="$(jq -r '.metadata.labels["test-label"] // empty' "${SNAP}")"
WANT_CRED_NAME="$(jq -r '.spec.userCredentials[0].credentialName' "${SNAP}")"
WANT_CRED_KEY="$(jq -r '.spec.userCredentials[0].sshPublicKeyBody' "${SNAP}")"
WANT_DIR_BUCKET="$(jq -r '.spec.storageDirectoryMappings[0].bucketRef.name' "${SNAP}")"
WANT_DIR_PREFIX="$(jq -r '.spec.storageDirectoryMappings[0].bucketPrefix' "${SNAP}")"
WANT_DIR_PATH="$(jq -r '.spec.storageDirectoryMappings[0].directory' "${SNAP}")"
WANT_DIR_PERM="$(jq -r '.spec.storageDirectoryMappings[0].permission' "${SNAP}")"

mv_wait_http create user "${FTP_USER_URL}" 200

# Out-of-band REST check on Cloud FTP User
# Map permission READ_WRITE -> 2 (or string "READ_WRITE" depending on proto representation)
# In proto int enum: READ_WRITE is 2, READ_ONLY is 1. If enum-encoding is not set, string "READ_WRITE" or int 2.
USER_JSON="${MV_RUN_DIR}/cloud/create-user.json"
mv_expect_jq create "cloud user name matches resourceID" "${USER_JSON}" '.username' "${USER_ID}"
mv_expect_jq create "cloud customerServiceAccount matches spec" "${USER_JSON}" '.customerServiceAccount' "${SA_EMAIL}"
mv_expect_jq create "cloud test-label matches spec" "${USER_JSON}" '.labels["test-label"]' "${WANT_LABEL}"
mv_expect_jq create "cloud credentialName matches spec" "${USER_JSON}" '.userCredentials[0].credentialName' "${WANT_CRED_NAME}"
# Cloud FTP strips trailing comment from ssh public key
mv_expect_jq create "cloud sshPublicKeyBody matches spec key body" "${USER_JSON}" \
  '.userCredentials[0].sshPublicKeyBody' "$(echo "${WANT_CRED_KEY}" | awk '{print $1" "$2}')"
mv_expect_jq create "cloud storageDirectory bucket matches spec" "${USER_JSON}" '.storageDirectoryMappings[0].bucket' "${WANT_DIR_BUCKET}"
mv_expect_jq create "cloud storageDirectory bucketPrefix matches spec" "${USER_JSON}" '.storageDirectoryMappings[0].bucketPrefix' "${WANT_DIR_PREFIX}"
mv_expect_jq create "cloud storageDirectory directory matches spec" "${USER_JSON}" '.storageDirectoryMappings[0].directory' "${WANT_DIR_PATH}"
mv_expect_jq create "cloud storageDirectory permission matches spec" "${USER_JSON}" \
  'if .storageDirectoryMappings[0].permission == 2 or .storageDirectoryMappings[0].permission == "READ_WRITE" then "READ_WRITE" else (.storageDirectoryMappings[0].permission | tostring) end' \
  "${WANT_DIR_PERM}"
mv_expect_jq create "cloud state is ACTIVE" "${USER_JSON}" \
  'if .state == 2 or .state == "ACTIVE" then "ACTIVE" else (.state | tostring) end' \
  "ACTIVE"

# Note: Cloud FTP has no gcloud surface (no gcloud ftp command exists).
mv_record create "gcloud ftp surface check" INFO "cloud/create-user.json" "Cloud FTP has no gcloud CLI surface; out-of-band verification via direct REST"

# Functional check: ListUsers on parent FTPServer via REST outside KCC
LIST_USERS_URL="${FTP_SERVER_URL}/users"
code="$(mv_rest create list-users GET "${LIST_USERS_URL}")"
mv_expect_http create "REST ListUsers on FTPServer" 200 "${code}" "cloud/create-list-users.json"
mv_expect_jq create "REST ListUsers includes created FTPUser" "${MV_RUN_DIR}/cloud/create-list-users.json" \
  '([.users[]? | select(.name | endswith("/users/'"${USER_ID}"'"))] | length > 0) | tostring' "true"

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
# Give the apiserver a moment to bump generation before polling for Ready.
sleep 2
mv_wait_ready update FTPUser "${USER_K8S}" 300
SNAP="$(mv_snapshot update FTPUser "${USER_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_LABEL_UPD="$(jq -r '.metadata.labels["test-label"] // empty' "${SNAP}")"
WANT_CRED_NAME_UPD="$(jq -r '.spec.userCredentials[0].credentialName' "${SNAP}")"
WANT_CRED_KEY_UPD="$(jq -r '.spec.userCredentials[0].sshPublicKeyBody' "${SNAP}")"
WANT_DIR_PREFIX_UPD="$(jq -r '.spec.storageDirectoryMappings[0].bucketPrefix' "${SNAP}")"
WANT_DIR_PATH_UPD="$(jq -r '.spec.storageDirectoryMappings[0].directory' "${SNAP}")"
WANT_DIR_PERM_UPD="$(jq -r '.spec.storageDirectoryMappings[0].permission' "${SNAP}")"

mv_wait_http update user "${FTP_USER_URL}" 200
USER_UPD_JSON="${MV_RUN_DIR}/cloud/update-user.json"

mv_expect_jq update "cloud test-label updated to new spec" "${USER_UPD_JSON}" '.labels["test-label"]' "${WANT_LABEL_UPD}"
mv_expect_jq update "cloud credentialName updated to new spec" "${USER_UPD_JSON}" '.userCredentials[0].credentialName' "${WANT_CRED_NAME_UPD}"
mv_expect_jq update "cloud sshPublicKeyBody updated to new spec" "${USER_UPD_JSON}" \
  '.userCredentials[0].sshPublicKeyBody' "$(echo "${WANT_CRED_KEY_UPD}" | awk '{print $1" "$2}')"
mv_expect_jq update "cloud storageDirectory bucketPrefix updated to new spec" "${USER_UPD_JSON}" '.storageDirectoryMappings[0].bucketPrefix' "${WANT_DIR_PREFIX_UPD}"
mv_expect_jq update "cloud storageDirectory directory updated to new spec" "${USER_UPD_JSON}" '.storageDirectoryMappings[0].directory' "${WANT_DIR_PATH_UPD}"
mv_expect_jq update "cloud storageDirectory permission updated to new spec" "${USER_UPD_JSON}" \
  'if .storageDirectoryMappings[0].permission == 1 or .storageDirectoryMappings[0].permission == "READ_ONLY" then "READ_ONLY" else (.storageDirectoryMappings[0].permission | tostring) end' \
  "${WANT_DIR_PERM_UPD}"

# Unchanged fields intact
mv_expect_jq update "cloud customerServiceAccount unchanged" "${USER_UPD_JSON}" '.customerServiceAccount' "${SA_EMAIL}"
mv_expect_jq update "cloud storageDirectory bucket unchanged" "${USER_UPD_JSON}" '.storageDirectoryMappings[0].bucket' "${WANT_DIR_BUCKET}"

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady FTPUser "${USER_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete FTPUser "${USER_K8S}" 300
mv_wait_http delete user-gone "${FTP_USER_CLEANUP_URL}" 404

# Functional check: ListUsers after delete no longer lists the user
code="$(mv_rest delete list-users GET "${LIST_USERS_URL}")"
mv_expect_http delete "REST ListUsers after delete returns 200" 200 "${code}" "cloud/delete-list-users.json"
mv_expect_jq delete "REST ListUsers no longer includes deleted FTPUser" "${MV_RUN_DIR}/cloud/delete-list-users.json" \
  '([.users[]? | select(.name | endswith("/users/'"${USER_ID}"'"))] | length == 0) | tostring' "true"

# What GCP actually received for the user over its whole life: exactly one create,
# one in-place update, one delete. No recreate, no spurious writes while steady.
mv_expect_audit delete audit-user "users/${USER_ID}" \
  "CreateUser,UpdateUser,DeleteUser"

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown FTPServer "${SERVER_K8S}" 600
mv_wait_http teardown server-gone "${FTP_SERVER_URL}" 404
mv_delete teardown IAMServiceAccount "${SA_K8S}" 300
mv_wait_http teardown sa-gone "${IAM_URL}" 404
mv_delete teardown StorageBucket "${BUCKET_K8S}" 300
mv_wait_http teardown bucket-gone "${GCS_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
