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

# Human-style validation of the NetAppStoragePool controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/netappstoragepool/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the netappstoragepool-maximal test fixture)
#   FUNCTIONAL=0     skip gcloud list and capacity functional checks
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/netapp/v1alpha1/netappstoragepool/netappstoragepool-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init NetAppStoragePool
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Names derived from the fixture (resourceIDs) so out-of-band checks never ask KCC.
POOL_K8S="netappstoragepool-${MV_UNIQUE_ID}"
POOL_ID="netappstoragepool-${MV_UNIQUE_ID}"
LOCATION="us-central1"

NETAPP_API="https://netapp.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
POOL_URL="${NETAPP_API}/storagePools/${POOL_ID}"

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
    mv_kubectl get netappstoragepool "${POOL_K8S}" >/dev/null 2>&1 && mv_delete cleanup NetAppStoragePool "${POOL_K8S}" 600
  fi
  local code
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
else
  mv_record deps "no dependencies required (references default VPC network)" INFO "" "default network pre-exists in project"
fi

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create NetAppStoragePool "${POOL_K8S}" 600
SNAP="$(mv_snapshot create NetAppStoragePool "${POOL_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

WANT_SERVICE_LEVEL="$(jq -r '.spec.serviceLevel' "${SNAP}")"
WANT_CAPACITY="$(jq -r '.spec.capacityGib' "${SNAP}")"
WANT_DESC="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_LABEL_ONE="$(jq -r '.metadata.labels["label-one"] // empty' "${SNAP}")"
WANT_ZONE="$(jq -r '.spec.zone // empty' "${SNAP}")"
WANT_REPLICA_ZONE="$(jq -r '.spec.replicaZone // empty' "${SNAP}")"

mv_wait_http create pool "${POOL_URL}" 200
mv_expect_jq create "cloud capacityGib matches spec" "${MV_RUN_DIR}/cloud/create-pool.json" '.capacityGib' "${WANT_CAPACITY}"
mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-pool.json" '.description' "${WANT_DESC}"
mv_expect_jq create "cloud serviceLevel matches spec" "${MV_RUN_DIR}/cloud/create-pool.json" '.serviceLevel' "${WANT_SERVICE_LEVEL}"
mv_expect_jq create "cloud label-one matches spec" "${MV_RUN_DIR}/cloud/create-pool.json" '.labels["label-one"]' "${WANT_LABEL_ONE}"
mv_expect_jq create "cloud state is READY" "${MV_RUN_DIR}/cloud/create-pool.json" '.state' "READY"
mv_expect_jq create "cloud network ends with default" "${MV_RUN_DIR}/cloud/create-pool.json" '(.network | split("/") | last)' "default"

# Note on zone and replicaZone: NetApp Volumes API omits write-only zone and replicaZone from GET/LIST responses.
# Verify whether they appear in the REST response, or document the API behavior.
if jq -e '.zone' "${MV_RUN_DIR}/cloud/create-pool.json" >/dev/null 2>&1; then
  mv_expect_jq create "cloud zone matches spec" "${MV_RUN_DIR}/cloud/create-pool.json" '.zone' "${WANT_ZONE}"
else
  mv_record create "cloud zone write-only field omitted on GET" INFO cloud/create-pool.json "spec requested zone=${WANT_ZONE}, API omits from GET"
fi
if jq -e '.replicaZone' "${MV_RUN_DIR}/cloud/create-pool.json" >/dev/null 2>&1; then
  mv_expect_jq create "cloud replicaZone matches spec" "${MV_RUN_DIR}/cloud/create-pool.json" '.replicaZone' "${WANT_REPLICA_ZONE}"
else
  mv_record create "cloud replicaZone write-only field omitted on GET" INFO cloud/create-pool.json "spec requested replicaZone=${WANT_REPLICA_ZONE}, API omits from GET"
fi

# gcloud describe out-of-band verification
mv_cli create gcloud-describe -- gcloud netapp storage-pools describe "${POOL_ID}" \
  --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-describe.txt" "${MV_RUN_DIR}/cloud/create-gcloud-describe.json"
mv_expect_jq create "gcloud describe capacityGib matches spec" "${MV_RUN_DIR}/cloud/create-gcloud-describe.json" '.capacityGib' "${WANT_CAPACITY}"
mv_expect_jq create "gcloud describe state is READY" "${MV_RUN_DIR}/cloud/create-gcloud-describe.json" '.state' "READY"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  # Functional check: gcloud list confirms pool is discoverable under the location
  mv_cli create gcloud-list -- gcloud netapp storage-pools list \
    --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
  mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-list.txt" "${MV_RUN_DIR}/cloud/create-gcloud-list.json"
  if jq -e --arg id "${POOL_ID}" 'map(select((.name | split("/") | last) == $id)) | length > 0' \
    "${MV_RUN_DIR}/cloud/create-gcloud-list.json" >/dev/null 2>&1; then
    mv_record create "functional check: storage pool listed in location" PASS cloud/create-gcloud-list.json "found ${POOL_ID} in list"
  else
    mv_record create "functional check: storage pool listed in location" FAIL cloud/create-gcloud-list.json "pool not found in list"
  fi

  # Record available throughput and volume count from GCP
  AVAILABLE_THROUGHPUT="$(jq -r '.availableThroughputMibps // "<none>"' "${MV_RUN_DIR}/cloud/create-pool.json")"
  mv_record create "functional check: available throughput reported" INFO cloud/create-pool.json "${AVAILABLE_THROUGHPUT} MiBps"
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update NetAppStoragePool "${POOL_K8S}" 600
SNAP="$(mv_snapshot update NetAppStoragePool "${POOL_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_CAPACITY_UPDATED="$(jq -r '.spec.capacityGib' "${SNAP}")"
WANT_DESC_UPDATED="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_LABEL_ONE_UPDATED="$(jq -r '.metadata.labels["label-one"] // empty' "${SNAP}")"
WANT_LABEL_NEW="$(jq -r '.metadata.labels["label-new"] // empty' "${SNAP}")"

mv_wait_http update pool "${POOL_URL}" 200
mv_expect_jq update "cloud capacityGib updated to new spec" "${MV_RUN_DIR}/cloud/update-pool.json" '.capacityGib' "${WANT_CAPACITY_UPDATED}"
mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-pool.json" '.description' "${WANT_DESC_UPDATED}"
mv_expect_jq update "cloud label-one updated to new spec" "${MV_RUN_DIR}/cloud/update-pool.json" '.labels["label-one"]' "${WANT_LABEL_ONE_UPDATED}"
mv_expect_jq update "cloud label-new added to new spec" "${MV_RUN_DIR}/cloud/update-pool.json" '.labels["label-new"]' "${WANT_LABEL_NEW}"

CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-pool.json")"
UPDATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/update-pool.json")"
mv_expect_eq update "createTime preserved across update (in-place update)" "${CREATE_T}" "${UPDATE_T}" cloud/update-pool.json

mv_cli update gcloud-describe -- gcloud netapp storage-pools describe "${POOL_ID}" \
  --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
mv_strip_cli_json "${MV_RUN_DIR}/cloud/update-gcloud-describe.txt" "${MV_RUN_DIR}/cloud/update-gcloud-describe.json"
mv_expect_jq update "gcloud describe capacityGib updated" "${MV_RUN_DIR}/cloud/update-gcloud-describe.json" '.capacityGib' "${WANT_CAPACITY_UPDATED}"

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady NetAppStoragePool "${POOL_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete NetAppStoragePool "${POOL_K8S}" 600
mv_wait_http delete pool-gone "${POOL_URL}" 404

# What GCP actually received for the pool over its whole life:
# NetApp Volumes logs CreateStoragePool with resourceName set to parent location (locations/${LOCATION}),
# and logs UpdateStoragePool and DeleteStoragePool with resourceName set to storagePools/${POOL_ID}.
# No recreate occurred, and no spurious writes while steady.
mv_expect_audit delete audit-pool "storagePools/${POOL_ID}" \
  "UpdateStoragePool,DeleteStoragePool" 120

# Verify CreateStoragePool in audit log for the NetApp service
CREATE_AUDIT_OUT="cloud/delete-audit-create.json"
FILTER_CREATE="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.serviceName=\"netapp.googleapis.com\" AND protoPayload.methodName=\"google.cloud.netapp.v1.NetApp.CreateStoragePool\" AND timestamp>=\"${MV_STARTED_UTC}\""
gcloud logging read "${FILTER_CREATE}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
  >"${MV_RUN_DIR}/${CREATE_AUDIT_OUT}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${CREATE_AUDIT_OUT}"
GOT_CREATE="$(jq -r '.[0].protoPayload.methodName | split(".") | last // empty' "${MV_RUN_DIR}/${CREATE_AUDIT_OUT}")"
mv_expect_eq delete "audit log CreateStoragePool called" "CreateStoragePool" "${GOT_CREATE}" "${CREATE_AUDIT_OUT}"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli delete gcloud-list -- gcloud netapp storage-pools list \
    --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
  mv_strip_cli_json "${MV_RUN_DIR}/cloud/delete-gcloud-list.txt" "${MV_RUN_DIR}/cloud/delete-gcloud-list.json"
  if jq -e --arg id "${POOL_ID}" 'map(select((.name | split("/") | last) == $id)) | length == 0' \
    "${MV_RUN_DIR}/cloud/delete-gcloud-list.json" >/dev/null 2>&1; then
    mv_record delete "functional check: pool no longer listed in location" PASS cloud/delete-gcloud-list.json "pool confirmed absent"
  else
    mv_record delete "functional check: pool no longer listed in location" FAIL cloud/delete-gcloud-list.json "pool still in list"
  fi
fi

# --- Phase: teardown ---------------------------------------------------------
if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
