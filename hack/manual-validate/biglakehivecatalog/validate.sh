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

# Human-style validation of the BigLakeHiveCatalog controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/biglakehivecatalog/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the biglakehivecatalog-maximal test fixture)
#   FUNCTIONAL=0     skip database creation/deletion functional check
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/bigquerybiglake/v1alpha1/biglakehivecatalog/biglakehivecatalog-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init BigLakeHiveCatalog
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Inspect the fixture's create.yaml to extract naming pattern
CATALOG_K8S="$(grep -E '^  name:' "${FIXTURE_DIR}/create.yaml" | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g")"
CATALOG_ID="$(grep -E 'resourceID:' "${FIXTURE_DIR}/create.yaml" | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g")"
BUCKET_NAME="$(grep -E '^  name:' "${FIXTURE_DIR}/dependencies.yaml" | head -1 | awk '{print $2}' | sed "s/\${projectId}/${GCP_PROJECT_ID}/g; s/\${uniqueId}/${MV_UNIQUE_ID}/g")"
DB_ID="db_${MV_UNIQUE_ID}"

EXPECTED_CATALOG_NAME="projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}"
CATALOG_URL="https://biglake.googleapis.com/hive/v1/projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}"
BUCKET_URL="https://storage.googleapis.com/storage/v1/b/${BUCKET_NAME}"

# Helper to convert mv_cli output to pure JSON for jq / mv_expect_jq
mv_strip_cli_json() {
  local src="$1" dst="$2"
  sed -e '1{/^\$ /d;}' -e '/^exit_code: [0-9]*$/d' "${src}" >"${dst}"
}

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"

  # Clean up any leftover functional check database if created
  gcloud biglake hive databases delete "${DB_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --quiet >/dev/null 2>&1 || true

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get biglakehivecatalog "${CATALOG_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup BigLakeHiveCatalog "${CATALOG_K8S}" 300

    # Storage bucket may contain objects created during functional test / Hive catalog metadata;
    # clean objects out so bucket deletion in KCC succeeds.
    gcloud storage rm -r "gs://${BUCKET_NAME}/**" >/dev/null 2>&1 || true
    mv_kubectl get storagebucket "${BUCKET_NAME}" >/dev/null 2>&1 && \
      mv_delete cleanup StorageBucket "${BUCKET_NAME}" 300
  fi

  local code
  code="$(mv_rest cleanup leakcheck-catalog GET "${CATALOG_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (catalog gone)" PASS cloud/cleanup-leakcheck-catalog.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
      "could not verify (HTTP ${code}); check manually: gcloud biglake hive catalogs describe ${CATALOG_ID} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP resources (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
      "HTTP ${code}; clean up manually: gcloud biglake hive catalogs delete ${CATALOG_ID} --project=${GCP_PROJECT_ID}"
  fi

  # Clean up bucket if still lingering
  gcloud storage rm -r "gs://${BUCKET_NAME}" >/dev/null 2>&1 || true

  local bcode
  bcode="$(mv_rest cleanup leakcheck-bucket GET "${BUCKET_URL}")"
  if [[ "${bcode}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (bucket gone)" PASS cloud/cleanup-leakcheck-bucket.json "HTTP 404"
  else
    mv_record cleanup "no leaked GCP resources (bucket gone)" FAIL cloud/cleanup-leakcheck-bucket.json \
      "HTTP ${bcode}; check bucket: ${BUCKET_URL}"
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
mv_enable_apis setup biglake.googleapis.com storage.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
mv_apply deps "${DEPS}"
mv_wait_ready deps StorageBucket "${BUCKET_NAME}" 300
mv_wait_http deps bucket "${BUCKET_URL}" 200

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create BigLakeHiveCatalog "${CATALOG_K8S}" 300
SNAP="$(mv_snapshot create BigLakeHiveCatalog "${CATALOG_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

WANT_LOCATION_URI="$(jq -r '.spec.locationURI' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_PRIMARY_LOCATION="$(jq -r '.spec.primaryLocation // "us-central1"' "${SNAP}")"

mv_wait_http create catalog "${CATALOG_URL}" 200
mv_expect_jq create "cloud locationUri matches spec" "${MV_RUN_DIR}/cloud/create-catalog.json" '.locationUri' "${WANT_LOCATION_URI}"
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-catalog.json" '.description' "${WANT_DESCRIPTION}"
fi
mv_expect_jq create "cloud name matches expected" "${MV_RUN_DIR}/cloud/create-catalog.json" '.name' "${EXPECTED_CATALOG_NAME}"
mv_expect_jq create "cloud replicas contains primary location" "${MV_RUN_DIR}/cloud/create-catalog.json" \
  '([.replicas[]? | select(.region=="'"${WANT_PRIMARY_LOCATION}"'")] | length > 0) | tostring' "true"

mv_cli create gcloud-describe-catalog -- gcloud biglake hive catalogs describe "${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-describe-catalog.txt" "${MV_RUN_DIR}/cloud/create-gcloud-describe-catalog.json"
mv_expect_jq create "gcloud describe catalog matches name" "${MV_RUN_DIR}/cloud/create-gcloud-describe-catalog.json" '.name' "${EXPECTED_CATALOG_NAME}"

mv_cli create gcloud-list-catalogs -- gcloud biglake hive catalogs list --project="${GCP_PROJECT_ID}" --format=json || true
mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-list-catalogs.txt" "${MV_RUN_DIR}/cloud/create-gcloud-list-catalogs.json"
FOUND_CATALOG="$(jq -r --arg NAME "${EXPECTED_CATALOG_NAME}" '[.[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-gcloud-list-catalogs.json")"
mv_expect_eq create "gcloud list catalogs includes the catalog" "${EXPECTED_CATALOG_NAME}" "${FOUND_CATALOG}" cloud/create-gcloud-list-catalogs.json

if [[ "${FUNCTIONAL}" == "1" ]]; then
  # Functional check: Create a BigLake Hive database inside this catalog using gcloud
  if mv_cli create functional-create-database -- gcloud biglake hive databases create "${DB_ID}" \
      --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" \
      --description="Validation functional test database"; then
    mv_record create "functional check: create database in catalog" PASS cloud/create-functional-create-database.txt ""
  else
    mv_record create "functional check: create database in catalog" FAIL cloud/create-functional-create-database.txt "failed to create database in catalog"
  fi

  # Verify database exists in catalog
  mv_cli create functional-describe-database -- gcloud biglake hive databases describe "${DB_ID}" \
      --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
  mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-functional-describe-database.txt" "${MV_RUN_DIR}/cloud/create-functional-describe-database.json"
  mv_expect_jq create "functional check: describe database matches name" "${MV_RUN_DIR}/cloud/create-functional-describe-database.json" \
    '.name' "${EXPECTED_CATALOG_NAME}/databases/${DB_ID}"

  # Delete the database so the catalog can be updated and eventually deleted
  if mv_cli create functional-delete-database -- gcloud biglake hive databases delete "${DB_ID}" \
      --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --quiet; then
    mv_record create "functional check: cleanup database in catalog" PASS cloud/create-functional-delete-database.txt ""
  else
    mv_record create "functional check: cleanup database in catalog" FAIL cloud/create-functional-delete-database.txt "failed to delete database"
  fi
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
# Give the apiserver a moment to bump generation before polling for Ready.
sleep 2
mv_wait_ready update BigLakeHiveCatalog "${CATALOG_K8S}" 300
SNAP="$(mv_snapshot update BigLakeHiveCatalog "${CATALOG_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_DESCRIPTION_UPDATE="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_LOCATION_URI_UPDATE="$(jq -r '.spec.locationURI' "${SNAP}")"

mv_wait_http update catalog "${CATALOG_URL}" 200
if [[ -n "${WANT_DESCRIPTION_UPDATE}" ]]; then
  mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-catalog.json" '.description' "${WANT_DESCRIPTION_UPDATE}"
fi
mv_expect_jq update "cloud locationUri retained" "${MV_RUN_DIR}/cloud/update-catalog.json" '.locationUri' "${WANT_LOCATION_URI_UPDATE}"
mv_expect_jq update "cloud replicas still contains primary location" "${MV_RUN_DIR}/cloud/update-catalog.json" \
  '([.replicas[]? | select(.region=="'"${WANT_PRIMARY_LOCATION}"'")] | length > 0) | tostring' "true"

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady BigLakeHiveCatalog "${CATALOG_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete BigLakeHiveCatalog "${CATALOG_K8S}" 300
mv_wait_http delete catalog-gone "${CATALOG_URL}" 404

# Cloud Audit Logs check: BigLake Hive Catalog operations logged under HiveMetastoreService
# Note: includes functional check database create/delete if FUNCTIONAL=1
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_expect_audit delete audit-catalog "catalogs/${CATALOG_ID}" \
    "CreateHiveCatalog,CreateHiveDatabase,DeleteHiveDatabase,UpdateHiveCatalog,DeleteHiveCatalog"
else
  mv_expect_audit delete audit-catalog "catalogs/${CATALOG_ID}" \
    "CreateHiveCatalog,UpdateHiveCatalog,DeleteHiveCatalog"
fi

# --- Phase: dependency teardown ----------------------------------------------
# Clean any objects created in storage bucket
gcloud storage rm -r "gs://${BUCKET_NAME}/**" >/dev/null 2>&1 || true
mv_delete teardown StorageBucket "${BUCKET_NAME}" 300
mv_wait_http teardown bucket-gone "${BUCKET_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
