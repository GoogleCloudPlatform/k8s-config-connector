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

# Human-style validation of the NetAppVolume controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/netappvolume/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the netappvolume-maximal test fixture)
#   FUNCTIONAL=0     skip gcloud list, storage pool volumeCount, and mount options functional checks
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/netapp/v1alpha1/netappvolume/netappvolume-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init NetAppVolume
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Names derived from the fixture (resourceIDs) so out-of-band checks never ask KCC.
POOL_K8S="netappvolume-dep-${MV_UNIQUE_ID}"
POOL_ID="netappvolume-dep-${MV_UNIQUE_ID}"
VOLUME_K8S="netappvolume-${MV_UNIQUE_ID}"
VOLUME_ID="netappvolume-${MV_UNIQUE_ID}"
LOCATION="us-central1"

NETAPP_API="https://netapp.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
POOL_URL="${NETAPP_API}/storagePools/${POOL_ID}"
VOLUME_URL="${NETAPP_API}/volumes/${VOLUME_ID}"

# Helper to strip mv_cli command and exit_code lines to get pure JSON
mv_strip_cli_json() {
  local src="$1" dst="$2"
  sed -e '1{/^\$ /d;}' -e '/^exit_code: [0-9]*$/d' "${src}" >"${dst}"
}

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get netappvolume "${VOLUME_K8S}" >/dev/null 2>&1 && mv_delete cleanup NetAppVolume "${VOLUME_K8S}" 600
    mv_kubectl get netappstoragepool "${POOL_K8S}" >/dev/null 2>&1 && mv_delete cleanup NetAppStoragePool "${POOL_K8S}" 600
  fi
  local code
  code="$(mv_rest cleanup leakcheck-volume GET "${VOLUME_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (volume gone)" PASS cloud/cleanup-leakcheck-volume.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (volume gone)" FAIL cloud/cleanup-leakcheck-volume.json \
      "could not verify (HTTP ${code}); check manually: gcloud netapp volumes describe ${VOLUME_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP resources (volume gone)" FAIL cloud/cleanup-leakcheck-volume.json \
      "HTTP ${code}; clean up manually: gcloud netapp volumes delete ${VOLUME_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID} --force --quiet"
  fi

  code="$(mv_rest cleanup leakcheck-pool GET "${POOL_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (storage pool gone)" PASS cloud/cleanup-leakcheck-pool.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (storage pool gone)" FAIL cloud/cleanup-leakcheck-pool.json \
      "could not verify (HTTP ${code}); check manually: gcloud netapp storage-pools describe ${POOL_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP resources (storage pool gone)" FAIL cloud/cleanup-leakcheck-pool.json \
      "HTTP ${code}; clean up manually: gcloud netapp storage-pools delete ${POOL_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID} --quiet"
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
mv_enable_apis setup netapp.googleapis.com compute.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
  mv_apply deps "${DEPS}"
  mv_wait_ready deps NetAppStoragePool "${POOL_K8S}" 600
  mv_wait_http deps pool "${POOL_URL}" 200
  mv_expect_jq deps "pool capacityGib is 2048" "${MV_RUN_DIR}/cloud/deps-pool.json" '.capacityGib' "2048"
  mv_expect_jq deps "pool serviceLevel is STANDARD" "${MV_RUN_DIR}/cloud/deps-pool.json" '.serviceLevel' "STANDARD"
  mv_expect_jq deps "pool state is READY" "${MV_RUN_DIR}/cloud/deps-pool.json" '.state' "READY"

  mv_cli deps gcloud-pool-describe -- gcloud netapp storage-pools describe "${POOL_ID}" \
    --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
else
  mv_record deps "no dependencies required" INFO "" "no dependencies.yaml in fixture"
fi

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create NetAppVolume "${VOLUME_K8S}" 600
SNAP="$(mv_snapshot create NetAppVolume "${VOLUME_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

WANT_SHARE_NAME="$(jq -r '.spec.shareName' "${SNAP}")"
WANT_CAPACITY="$(jq -r '.spec.capacityGib | tostring' "${SNAP}")"
WANT_DESC="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_PROTOCOL="$(jq -r '.spec.protocols[0]' "${SNAP}")"
WANT_UNIX_PERMS="$(jq -r '.spec.unixPermissions' "${SNAP}")"
WANT_SEC_STYLE="$(jq -r '.spec.securityStyle' "${SNAP}")"
WANT_SNAP_DIR="$(jq -r '.spec.snapshotDirectory | tostring' "${SNAP}")"
WANT_CLIENTS="$(jq -r '.spec.exportPolicy.rules[0].allowedClients' "${SNAP}")"
WANT_ACCESS="$(jq -r '.spec.exportPolicy.rules[0].accessType' "${SNAP}")"
WANT_NFSV3="$(jq -r '.spec.exportPolicy.rules[0].nfsv3 | tostring' "${SNAP}")"
WANT_NFSV4="$(jq -r '.spec.exportPolicy.rules[0].nfsv4 | tostring' "${SNAP}")"
WANT_SNAP_ENABLED="$(jq -r '.spec.snapshotPolicy.enabled | tostring' "${SNAP}")"
WANT_SNAP_HOUR="$(jq -r '.spec.snapshotPolicy.dailySchedule.hour | tostring' "${SNAP}")"
WANT_SNAP_MIN="$(jq -r '.spec.snapshotPolicy.dailySchedule.minute | tostring' "${SNAP}")"
WANT_SNAP_KEEP="$(jq -r '.spec.snapshotPolicy.dailySchedule.snapshotsToKeep | tostring' "${SNAP}")"
WANT_LABEL_ENV="$(jq -r '.spec.labels["env"] // empty' "${SNAP}")"
WANT_STORAGE_POOL="$(jq -r '.spec.storagePoolRef.name' "${SNAP}")"

mv_wait_http create volume "${VOLUME_URL}" 200
mv_expect_jq create "cloud shareName matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.shareName' "${WANT_SHARE_NAME}"
mv_expect_jq create "cloud capacityGib matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.capacityGib' "${WANT_CAPACITY}"
mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.description' "${WANT_DESC}"
mv_expect_jq create "cloud protocols[0] matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.protocols[0]' "${WANT_PROTOCOL}"
mv_expect_jq create "cloud unixPermissions matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.unixPermissions' "${WANT_UNIX_PERMS}"
mv_expect_jq create "cloud securityStyle matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.securityStyle' "${WANT_SEC_STYLE}"
mv_expect_jq create "cloud snapshotDirectory matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.snapshotDirectory | tostring' "${WANT_SNAP_DIR}"
mv_expect_jq create "cloud exportPolicy allowedClients matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.exportPolicy.rules[0].allowedClients' "${WANT_CLIENTS}"
mv_expect_jq create "cloud exportPolicy accessType matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.exportPolicy.rules[0].accessType' "${WANT_ACCESS}"
mv_expect_jq create "cloud exportPolicy nfsv3 matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.exportPolicy.rules[0].nfsv3 | tostring' "${WANT_NFSV3}"
mv_expect_jq create "cloud exportPolicy nfsv4 matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.exportPolicy.rules[0].nfsv4 | tostring' "${WANT_NFSV4}"
mv_expect_jq create "cloud snapshotPolicy enabled matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.snapshotPolicy.enabled | tostring' "${WANT_SNAP_ENABLED}"
mv_expect_jq create "cloud snapshotPolicy dailySchedule hour matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.snapshotPolicy.dailySchedule.hour | tostring' "${WANT_SNAP_HOUR}"
mv_expect_jq create "cloud snapshotPolicy dailySchedule minute matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.snapshotPolicy.dailySchedule.minute | tostring' "${WANT_SNAP_MIN}"
mv_expect_jq create "cloud snapshotPolicy dailySchedule snapshotsToKeep matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.snapshotPolicy.dailySchedule.snapshotsToKeep | tostring' "${WANT_SNAP_KEEP}"
mv_expect_jq create "cloud label env matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '.labels["env"]' "${WANT_LABEL_ENV}"
mv_expect_jq create "cloud storagePool matches spec" "${MV_RUN_DIR}/cloud/create-volume.json" '(.storagePool | split("/") | last)' "${WANT_STORAGE_POOL}"
mv_expect_jq create "cloud state is READY" "${MV_RUN_DIR}/cloud/create-volume.json" '.state' "READY"

# gcloud describe out-of-band verification
mv_cli create gcloud-describe -- gcloud netapp volumes describe "${VOLUME_ID}" \
  --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-describe.txt" "${MV_RUN_DIR}/cloud/create-gcloud-describe.json"
mv_expect_jq create "gcloud describe capacityGib matches spec" "${MV_RUN_DIR}/cloud/create-gcloud-describe.json" '.capacityGib' "${WANT_CAPACITY}"
mv_expect_jq create "gcloud describe state is READY" "${MV_RUN_DIR}/cloud/create-gcloud-describe.json" '.state' "READY"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  # Functional check 1: gcloud list confirms volume is discoverable under the location
  mv_cli create gcloud-list -- gcloud netapp volumes list \
    --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
  mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-list.txt" "${MV_RUN_DIR}/cloud/create-gcloud-list.json"
  if jq -e --arg id "${VOLUME_ID}" 'map(select((.name | split("/") | last) == $id)) | length > 0' \
    "${MV_RUN_DIR}/cloud/create-gcloud-list.json" >/dev/null 2>&1; then
    mv_record create "functional check: volume listed in location" PASS cloud/create-gcloud-list.json "found ${VOLUME_ID} in list"
  else
    mv_record create "functional check: volume listed in location" FAIL cloud/create-gcloud-list.json "volume not found in list"
  fi

  # Functional check 2: storage pool tracks the volume (volumeCount and volumeCapacityGib)
  mv_rest create check-pool GET "${POOL_URL}" >/dev/null
  mv_expect_jq create "functional check: storage pool volumeCount is 1" "${MV_RUN_DIR}/cloud/create-check-pool.json" '(.volumeCount // 0) | tostring' "1"
  mv_expect_jq create "functional check: storage pool volumeCapacityGib matches volume" "${MV_RUN_DIR}/cloud/create-check-pool.json" '(.volumeCapacityGib // 0) | tostring' "${WANT_CAPACITY}"

  # Functional check 3: record allocated mount options from GCP
  MOUNT_IP="$(jq -r '.mountOptions[0].ipAddress // "<none>"' "${MV_RUN_DIR}/cloud/create-volume.json")"
  MOUNT_EXPORT="$(jq -r '.mountOptions[0].exportFull // "<none>"' "${MV_RUN_DIR}/cloud/create-volume.json")"
  mv_record create "functional check: volume mount endpoint allocated" INFO cloud/create-volume.json "ipAddress=${MOUNT_IP} export=${MOUNT_EXPORT}"
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update NetAppVolume "${VOLUME_K8S}" 600
SNAP="$(mv_snapshot update NetAppVolume "${VOLUME_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_CAPACITY_UPDATED="$(jq -r '.spec.capacityGib | tostring' "${SNAP}")"
WANT_DESC_UPDATED="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_LABEL_ENV_UPDATED="$(jq -r '.spec.labels["env"] // empty' "${SNAP}")"
WANT_UNIX_PERMS_UPDATED="$(jq -r '.spec.unixPermissions' "${SNAP}")"
WANT_CLIENTS_UPDATED="$(jq -r '.spec.exportPolicy.rules[0].allowedClients' "${SNAP}")"
WANT_SNAP_HOUR_UPDATED="$(jq -r '.spec.snapshotPolicy.dailySchedule.hour | tostring' "${SNAP}")"

mv_wait_http update volume "${VOLUME_URL}" 200
mv_expect_jq update "cloud capacityGib updated to new spec" "${MV_RUN_DIR}/cloud/update-volume.json" '.capacityGib' "${WANT_CAPACITY_UPDATED}"
mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-volume.json" '.description' "${WANT_DESC_UPDATED}"
mv_expect_jq update "cloud label env updated to new spec" "${MV_RUN_DIR}/cloud/update-volume.json" '.labels["env"]' "${WANT_LABEL_ENV_UPDATED}"
mv_expect_jq update "cloud unixPermissions updated to new spec" "${MV_RUN_DIR}/cloud/update-volume.json" '.unixPermissions' "${WANT_UNIX_PERMS_UPDATED}"
mv_expect_jq update "cloud exportPolicy allowedClients updated to new spec" "${MV_RUN_DIR}/cloud/update-volume.json" '.exportPolicy.rules[0].allowedClients' "${WANT_CLIENTS_UPDATED}"
mv_expect_jq update "cloud snapshotPolicy dailySchedule hour updated to new spec" "${MV_RUN_DIR}/cloud/update-volume.json" '.snapshotPolicy.dailySchedule.hour | tostring' "${WANT_SNAP_HOUR_UPDATED}"

# Verify unchanged fields remain intact
mv_expect_jq update "cloud shareName preserved" "${MV_RUN_DIR}/cloud/update-volume.json" '.shareName' "${WANT_SHARE_NAME}"
mv_expect_jq update "cloud protocols preserved" "${MV_RUN_DIR}/cloud/update-volume.json" '.protocols[0]' "${WANT_PROTOCOL}"
mv_expect_jq update "cloud securityStyle preserved" "${MV_RUN_DIR}/cloud/update-volume.json" '.securityStyle' "${WANT_SEC_STYLE}"
mv_expect_jq update "cloud snapshotDirectory preserved" "${MV_RUN_DIR}/cloud/update-volume.json" '.snapshotDirectory | tostring' "${WANT_SNAP_DIR}"
mv_expect_jq update "cloud snapshotPolicy enabled preserved" "${MV_RUN_DIR}/cloud/update-volume.json" '.snapshotPolicy.enabled | tostring' "${WANT_SNAP_ENABLED}"
mv_expect_jq update "cloud snapshotPolicy dailySchedule minute preserved" "${MV_RUN_DIR}/cloud/update-volume.json" '.snapshotPolicy.dailySchedule.minute | tostring' "${WANT_SNAP_MIN}"
mv_expect_jq update "cloud snapshotPolicy dailySchedule snapshotsToKeep preserved" "${MV_RUN_DIR}/cloud/update-volume.json" '.snapshotPolicy.dailySchedule.snapshotsToKeep | tostring' "${WANT_SNAP_KEEP}"

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-volume.json")"
UPDATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/update-volume.json")"
mv_expect_eq update "createTime preserved across update (in-place update)" "${CREATE_T}" "${UPDATE_T}" cloud/update-volume.json

mv_cli update gcloud-describe -- gcloud netapp volumes describe "${VOLUME_ID}" \
  --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
mv_strip_cli_json "${MV_RUN_DIR}/cloud/update-gcloud-describe.txt" "${MV_RUN_DIR}/cloud/update-gcloud-describe.json"
mv_expect_jq update "gcloud describe capacityGib updated" "${MV_RUN_DIR}/cloud/update-gcloud-describe.json" '.capacityGib' "${WANT_CAPACITY_UPDATED}"
mv_expect_jq update "gcloud describe description updated" "${MV_RUN_DIR}/cloud/update-gcloud-describe.json" '.description' "${WANT_DESC_UPDATED}"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  # Functional check: storage pool volumeCapacityGib updated to 150
  mv_rest update check-pool GET "${POOL_URL}" >/dev/null
  mv_expect_jq update "functional check: storage pool volumeCapacityGib updated" "${MV_RUN_DIR}/cloud/update-check-pool.json" '(.volumeCapacityGib // 0) | tostring' "${WANT_CAPACITY_UPDATED}"
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady NetAppVolume "${VOLUME_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete NetAppVolume "${VOLUME_K8S}" 600
mv_wait_http delete volume-gone "${VOLUME_URL}" 404

# What GCP actually received for the volume over its whole life:
# NetApp Volumes logs CreateVolume with resourceName set to parent location (locations/${LOCATION}),
# and logs UpdateVolume with resourceName set to volumes/${VOLUME_ID}.
# Note: GCP NetApp Volumes API does not emit DeleteVolume in Cloud Audit Logs activity logs.
# No recreate occurred, and no spurious writes while steady.
mv_expect_audit delete audit-volume "volumes/${VOLUME_ID}" \
  "UpdateVolume" 120

# Verify CreateVolume in audit log for the NetApp service
CREATE_AUDIT_OUT="cloud/delete-audit-create.json"
FILTER_CREATE="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.serviceName=\"netapp.googleapis.com\" AND protoPayload.methodName=\"google.cloud.netapp.v1.NetApp.CreateVolume\" AND timestamp>=\"${MV_STARTED_UTC}\""
gcloud logging read "${FILTER_CREATE}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
  >"${MV_RUN_DIR}/${CREATE_AUDIT_OUT}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${CREATE_AUDIT_OUT}"
GOT_CREATE="$(jq -r '.[0].protoPayload.methodName | split(".") | last // empty' "${MV_RUN_DIR}/${CREATE_AUDIT_OUT}")"
mv_expect_eq delete "audit log CreateVolume called" "CreateVolume" "${GOT_CREATE}" "${CREATE_AUDIT_OUT}"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli delete gcloud-list -- gcloud netapp volumes list \
    --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
  mv_strip_cli_json "${MV_RUN_DIR}/cloud/delete-gcloud-list.txt" "${MV_RUN_DIR}/cloud/delete-gcloud-list.json"
  if jq -e --arg id "${VOLUME_ID}" 'map(select((.name | split("/") | last) == $id)) | length == 0' \
    "${MV_RUN_DIR}/cloud/delete-gcloud-list.json" >/dev/null 2>&1; then
    mv_record delete "functional check: volume no longer listed in location" PASS cloud/delete-gcloud-list.json "volume confirmed absent"
  else
    mv_record delete "functional check: volume no longer listed in location" FAIL cloud/delete-gcloud-list.json "volume still in list"
  fi

  # Pool volumeCount returns to 0
  mv_rest delete check-pool GET "${POOL_URL}" >/dev/null
  mv_expect_jq delete "functional check: storage pool volumeCount is 0" "${MV_RUN_DIR}/cloud/delete-check-pool.json" '(.volumeCount // 0) | tostring' "0"
fi

# --- Phase: dependency teardown ----------------------------------------------
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  mv_delete teardown NetAppStoragePool "${POOL_K8S}" 600
  mv_wait_http teardown pool-gone "${POOL_URL}" 404
fi

# --- Phase: teardown ---------------------------------------------------------
if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
