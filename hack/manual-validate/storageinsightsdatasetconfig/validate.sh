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

# Human-style validation of the StorageInsightsDatasetConfig controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/storageinsightsdatasetconfig/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the storageinsightsdatasetconfig-maximal test fixture)
#   FUNCTIONAL=0     skip functional checks (listing dataset configs via gcloud CLI)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/storageinsights/v1alpha1/storageinsightsdatasetconfig/storageinsightsdatasetconfig-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init StorageInsightsDatasetConfig
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
CONFIG_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
CONFIG_ID="$(grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
CONFIG_ID="${CONFIG_ID:-${CONFIG_K8S}}"
LOCATION="$(grep -m1 '^[[:space:]]*location:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
LOCATION="${LOCATION:-us-central1}"

if [[ -n "${DEPS:-}" ]]; then
  BUCKET_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${DEPS}" | awk '{print $2}' || true)"
  BUCKET_NAME="${BUCKET_K8S}"
fi

API_BASE="https://storageinsights.googleapis.com/v1"
PARENT="projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
CONFIG_URL="${API_BASE}/${PARENT}/datasetConfigs/${CONFIG_ID}"
LIST_URL="${API_BASE}/${PARENT}/datasetConfigs"
EXPECTED_NAME="${PARENT}/datasetConfigs/${CONFIG_ID}"

if [[ -n "${BUCKET_NAME:-}" ]]; then
  BUCKET_URL="https://storage.googleapis.com/storage/v1/b/${BUCKET_NAME}"
fi

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    if [[ -n "${CONFIG_K8S:-}" ]]; then
      mv_kubectl get storageinsightsdatasetconfig "${CONFIG_K8S}" >/dev/null 2>&1 && mv_delete cleanup StorageInsightsDatasetConfig "${CONFIG_K8S}" 300
    fi
    if [[ -n "${BUCKET_K8S:-}" ]]; then
      mv_kubectl get storagebucket "${BUCKET_K8S}" >/dev/null 2>&1 && mv_delete cleanup StorageBucket "${BUCKET_K8S}" 300
    fi
  fi

  if [[ -n "${CONFIG_URL:-}" ]]; then
    local code
    code="$(mv_rest cleanup leakcheck-datasetconfig GET "${CONFIG_URL}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (datasetConfig gone)" PASS cloud/cleanup-leakcheck-datasetconfig.json "HTTP 404"
    elif [[ "${code}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (datasetConfig gone)" FAIL cloud/cleanup-leakcheck-datasetconfig.json \
        "could not verify (HTTP ${code}); check manually: gcloud storage insights dataset-configs describe ${CONFIG_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID}"
    else
      mv_record cleanup "no leaked GCP resources (datasetConfig gone)" FAIL cloud/cleanup-leakcheck-datasetconfig.json \
        "HTTP ${code}; clean up manually: gcloud storage insights dataset-configs delete ${CONFIG_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID} --quiet"
    fi
  fi

  if [[ -n "${BUCKET_URL:-}" ]]; then
    local bcode
    bcode="$(mv_rest cleanup leakcheck-bucket GET "${BUCKET_URL}")"
    if [[ "${bcode}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (bucket gone)" PASS cloud/cleanup-leakcheck-bucket.json "HTTP 404"
    elif [[ "${bcode}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (bucket gone)" FAIL cloud/cleanup-leakcheck-bucket.json \
        "could not verify (HTTP ${bcode}); check manually: gcloud storage buckets describe gs://${BUCKET_NAME} --project=${GCP_PROJECT_ID}"
    else
      mv_record cleanup "no leaked GCP resources (bucket gone)" FAIL cloud/cleanup-leakcheck-bucket.json \
        "HTTP ${bcode}; clean up manually: gcloud storage rm --recursive gs://${BUCKET_NAME}"
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

# ---------------------------------------------------------------------------
# Every API the controller, dependencies and out-of-band checks call.
mv_enable_apis setup storage.googleapis.com storageinsights.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
  mv_wait_ready deps StorageBucket "${BUCKET_K8S}" 300
  mv_wait_http deps bucket "${BUCKET_URL}" 200
  mv_cli deps gcloud-storage-buckets-describe -- gcloud storage buckets describe "gs://${BUCKET_NAME}" --project="${GCP_PROJECT_ID}" --format=json || true
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create StorageInsightsDatasetConfig "${CONFIG_K8S}" 300
SNAP="$(mv_snapshot create StorageInsightsDatasetConfig "${CONFIG_K8S}")"
STATUS_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${STATUS_REF:-<none>}"

WANT_DESC="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_LABEL_FOO="$(jq -r '.spec.labels.foo // empty' "${SNAP}")"
WANT_RETENTION="$(jq -r '(.spec.retentionPeriodDays // empty) | tostring' "${SNAP}")"
WANT_NEW_BUCKETS="$(jq -r '(.spec.includeNewlyCreatedBuckets // false) | tostring' "${SNAP}")"
WANT_ID_TYPE="$(jq -r '.spec.identity.type // empty' "${SNAP}")"

mv_wait_http create datasetconfig "${CONFIG_URL}" 200

# Verify cloud name matches expected resource path
mv_expect_jq create "cloud name matches expected identity" "${MV_RUN_DIR}/cloud/create-datasetconfig.json" '.name' "${EXPECTED_NAME}"

# Verify spec fields in cloud response
if [[ -n "${WANT_DESC}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-datasetconfig.json" '.description // empty' "${WANT_DESC}"
fi

if [[ -n "${WANT_LABEL_FOO}" ]]; then
  mv_expect_jq create "cloud label foo matches spec" "${MV_RUN_DIR}/cloud/create-datasetconfig.json" '.labels.foo // empty' "${WANT_LABEL_FOO}"
fi
mv_expect_jq create "cloud label managed-by-cnrm is set by controller" "${MV_RUN_DIR}/cloud/create-datasetconfig.json" '.labels["managed-by-cnrm"] // empty' "true"

mv_expect_jq create "cloud sourceProjects projectNumbers match spec" "${MV_RUN_DIR}/cloud/create-datasetconfig.json" '.sourceProjects.projectNumbers[0] // empty' "${GCP_PROJECT_NUMBER}"
mv_expect_jq create "cloud includeCloudStorageLocations locations match spec" "${MV_RUN_DIR}/cloud/create-datasetconfig.json" '.includeCloudStorageLocations.locations[0] // empty' "${LOCATION}"
if [[ -n "${BUCKET_NAME:-}" ]]; then
  mv_expect_jq create "cloud includeCloudStorageBuckets bucketName matches spec" "${MV_RUN_DIR}/cloud/create-datasetconfig.json" '.includeCloudStorageBuckets.cloudStorageBuckets[0].bucketName // empty' "${BUCKET_NAME}"
fi
mv_expect_jq create "cloud includeNewlyCreatedBuckets matches spec" "${MV_RUN_DIR}/cloud/create-datasetconfig.json" '(.includeNewlyCreatedBuckets // false) | tostring' "${WANT_NEW_BUCKETS}"
mv_expect_jq create "cloud retentionPeriodDays matches spec" "${MV_RUN_DIR}/cloud/create-datasetconfig.json" '(.retentionPeriodDays // empty) | tostring' "${WANT_RETENTION}"

# Cloud identity type and generated service account
CLOUD_ID_TYPE="$(jq -r '.identity.type // empty' "${MV_RUN_DIR}/cloud/create-datasetconfig.json")"
if [[ "${CLOUD_ID_TYPE}" == "IDENTITY_TYPE_PER_CONFIG" || "${CLOUD_ID_TYPE}" == "1" ]]; then
  mv_record create "cloud identity.type matches spec" PASS cloud/create-datasetconfig.json "type=${CLOUD_ID_TYPE}"
else
  mv_record create "cloud identity.type matches spec" FAIL cloud/create-datasetconfig.json "want IDENTITY_TYPE_PER_CONFIG or 1, got ${CLOUD_ID_TYPE}"
fi

IDENTITY_NAME="$(jq -r '.identity.name // empty' "${MV_RUN_DIR}/cloud/create-datasetconfig.json")"
if [[ -n "${IDENTITY_NAME}" && "${IDENTITY_NAME}" =~ @gcp-sa-storageinsights\.iam\.gserviceaccount\.com$ ]]; then
  mv_record create "identity service account created and valid" PASS cloud/create-datasetconfig.json "${IDENTITY_NAME}"
else
  mv_record create "identity service account created and valid" FAIL cloud/create-datasetconfig.json "invalid identity.name: ${IDENTITY_NAME}"
fi

# Linked dataset name
LINKED_DATASET="$(jq -r '.link.dataset // empty' "${MV_RUN_DIR}/cloud/create-datasetconfig.json")"
if [[ -n "${LINKED_DATASET}" && "${LINKED_DATASET}" =~ ^${CONFIG_ID}_ ]]; then
  mv_record create "link.dataset created and matches config ID" PASS cloud/create-datasetconfig.json "${LINKED_DATASET}"
else
  mv_record create "link.dataset created and matches config ID" FAIL cloud/create-datasetconfig.json "invalid link.dataset: ${LINKED_DATASET}"
fi

# Cloud state is ACTIVE
CLOUD_STATE="$(jq -r '.datasetConfigState // empty' "${MV_RUN_DIR}/cloud/create-datasetconfig.json")"
if [[ "${CLOUD_STATE}" == "CONFIG_STATE_ACTIVE" || "${CLOUD_STATE}" == "1" ]]; then
  mv_record create "cloud datasetConfigState is ACTIVE" PASS cloud/create-datasetconfig.json "state=${CLOUD_STATE}"
else
  mv_record create "cloud datasetConfigState is ACTIVE" FAIL cloud/create-datasetconfig.json "got state=${CLOUD_STATE}"
fi

# CLI describe
mv_cli create gcloud-describe -- gcloud storage insights dataset-configs describe "${CONFIG_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
if grep -q "/datasetConfigs/${CONFIG_ID}\"" "${MV_RUN_DIR}/cloud/create-gcloud-describe.txt"; then
  mv_record create "gcloud storage insights dataset-configs describe succeeds" PASS cloud/create-gcloud-describe.txt ""
else
  mv_record create "gcloud storage insights dataset-configs describe succeeds" FAIL cloud/create-gcloud-describe.txt "could not describe dataset config via gcloud"
fi

# Functional check: List dataset configs includes the new config
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli create gcloud-list -- gcloud storage insights dataset-configs list --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
  if grep -q "/datasetConfigs/${CONFIG_ID}\"" "${MV_RUN_DIR}/cloud/create-gcloud-list.txt"; then
    mv_record create "gcloud storage insights dataset-configs list contains config" PASS cloud/create-gcloud-list.txt ""
  else
    mv_record create "gcloud storage insights dataset-configs list contains config" FAIL cloud/create-gcloud-list.txt "config id not found in list"
  fi
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update StorageInsightsDatasetConfig "${CONFIG_K8S}" 300
SNAP="$(mv_snapshot update StorageInsightsDatasetConfig "${CONFIG_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_UPDATE_DESC="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_UPDATE_RETENTION="$(jq -r '(.spec.retentionPeriodDays // empty) | tostring' "${SNAP}")"
WANT_UPDATE_NEW_BUCKETS="$(jq -r '(.spec.includeNewlyCreatedBuckets // false) | tostring' "${SNAP}")"

mv_wait_http update datasetconfig "${CONFIG_URL}" 200

# Check mutable fields were updated
mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-datasetconfig.json" '.description // empty' "${WANT_UPDATE_DESC}"
mv_expect_jq update "cloud retentionPeriodDays updated to new spec" "${MV_RUN_DIR}/cloud/update-datasetconfig.json" '(.retentionPeriodDays // empty) | tostring' "${WANT_UPDATE_RETENTION}"
mv_expect_jq update "cloud includeNewlyCreatedBuckets updated to new spec" "${MV_RUN_DIR}/cloud/update-datasetconfig.json" '(.includeNewlyCreatedBuckets // false) | tostring' "${WANT_UPDATE_NEW_BUCKETS}"

# Check immutable / unchanged fields remain intact
mv_expect_jq update "cloud label foo remains intact" "${MV_RUN_DIR}/cloud/update-datasetconfig.json" '.labels.foo // empty' "${WANT_LABEL_FOO}"
mv_expect_jq update "cloud label managed-by-cnrm remains intact" "${MV_RUN_DIR}/cloud/update-datasetconfig.json" '.labels["managed-by-cnrm"] // empty' "true"
if [[ -n "${BUCKET_NAME:-}" ]]; then
  mv_expect_jq update "cloud includeCloudStorageBuckets bucketName remains intact" "${MV_RUN_DIR}/cloud/update-datasetconfig.json" '.includeCloudStorageBuckets.cloudStorageBuckets[0].bucketName // empty' "${BUCKET_NAME}"
fi
mv_expect_jq update "cloud sourceProjects projectNumbers remains intact" "${MV_RUN_DIR}/cloud/update-datasetconfig.json" '.sourceProjects.projectNumbers[0] // empty' "${GCP_PROJECT_NUMBER}"

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady StorageInsightsDatasetConfig "${CONFIG_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete StorageInsightsDatasetConfig "${CONFIG_K8S}" 300
mv_wait_http delete datasetconfig-gone "${CONFIG_URL}" 404
mv_expect_audit delete audit-datasetconfig "datasetConfigs/${CONFIG_ID}" \
  "CreateDatasetConfig,UpdateDatasetConfig,DeleteDatasetConfig"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli delete gcloud-list-after-delete -- gcloud storage insights dataset-configs list --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
  if grep -q "/datasetConfigs/${CONFIG_ID}\"" "${MV_RUN_DIR}/cloud/delete-gcloud-list-after-delete.txt"; then
    mv_record delete "dataset config no longer in gcloud list" FAIL cloud/delete-gcloud-list-after-delete.txt "config id still found"
  else
    mv_record delete "dataset config no longer in gcloud list" PASS cloud/delete-gcloud-list-after-delete.txt ""
  fi
fi

# --- Phase: dependency teardown ----------------------------------------------
if [[ -n "${BUCKET_K8S:-}" ]]; then
  mv_delete teardown StorageBucket "${BUCKET_K8S}" 300
  mv_wait_http teardown bucket-gone "${BUCKET_URL}" 404
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
