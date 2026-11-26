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

# Human-style validation of the BigLakeHiveDatabase controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/biglakehivedatabase/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the biglakehivedatabase-maximal test fixture)
#   FUNCTIONAL=0     skip table creation/deletion functional check
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/bigquerybiglake/v1alpha1/biglakehivedatabase/biglakehivedatabase-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init BigLakeHiveDatabase
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Inspect the fixture's create.yaml and dependencies.yaml to extract naming patterns
DB_K8S="$(grep -E '^  name:' "${FIXTURE_DIR}/create.yaml" | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g; s/\${projectId}/${GCP_PROJECT_ID}/g")"
DB_ID="$(grep -E 'resourceID:' "${FIXTURE_DIR}/create.yaml" | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g; s/\${projectId}/${GCP_PROJECT_ID}/g")"
if [[ -z "${DB_ID}" ]]; then
  DB_ID="${DB_K8S}"
fi

CATALOG_K8S="$(grep -A 2 'parentRef:' "${FIXTURE_DIR}/create.yaml" | grep -E 'name:' | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g; s/\${projectId}/${GCP_PROJECT_ID}/g")"
CATALOG_ID="$(grep -A 10 'kind: BigLakeHiveCatalog' "${FIXTURE_DIR}/dependencies.yaml" | grep -E 'resourceID:' | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g; s/\${projectId}/${GCP_PROJECT_ID}/g")"
if [[ -z "${CATALOG_ID}" ]]; then
  CATALOG_ID="${CATALOG_K8S}"
fi

BUCKET_NAME="$(grep -A 10 'kind: StorageBucket' "${FIXTURE_DIR}/dependencies.yaml" | grep -E 'name:' | head -1 | awk '{print $2}' | sed "s/\${projectId}/${GCP_PROJECT_ID}/g; s/\${uniqueId}/${MV_UNIQUE_ID}/g")"
TABLE_ID="t_${MV_UNIQUE_ID}"

EXPECTED_DB_NAME="projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}/databases/${DB_ID}"
DB_URL="https://biglake.googleapis.com/hive/v1/projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}/databases/${DB_ID}"
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

  # Clean up any leftover functional check table if created
  gcloud biglake hive tables delete "${TABLE_ID}" --catalog="${CATALOG_ID}" --database="${DB_ID}" --project="${GCP_PROJECT_ID}" --quiet >/dev/null 2>&1 || true

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get biglakehivedatabase "${DB_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup BigLakeHiveDatabase "${DB_K8S}" 300

    mv_kubectl get biglakehivecatalog "${CATALOG_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup BigLakeHiveCatalog "${CATALOG_K8S}" 300

    # Storage bucket may contain objects created during functional test / Hive metadata;
    # clean objects out so bucket deletion in KCC succeeds.
    gcloud storage rm -r "gs://${BUCKET_NAME}/**" >/dev/null 2>&1 || true
    mv_kubectl get storagebucket "${BUCKET_NAME}" >/dev/null 2>&1 && \
      mv_delete cleanup StorageBucket "${BUCKET_NAME}" 300
  fi

  local code
  code="$(mv_rest cleanup leakcheck-db GET "${DB_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (database gone)" PASS cloud/cleanup-leakcheck-db.json "HTTP 404"
  elif [[ "${code}" == "400" ]] && grep -q "does not exist" "${MV_RUN_DIR}/cloud/cleanup-leakcheck-db.json" 2>/dev/null; then
    mv_record cleanup "no leaked GCP resources (database gone)" PASS cloud/cleanup-leakcheck-db.json "HTTP 400 (parent catalog does not exist)"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (database gone)" FAIL cloud/cleanup-leakcheck-db.json \
      "could not verify (HTTP ${code}); check manually: gcloud biglake hive databases describe ${DB_ID} --catalog=${CATALOG_ID} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP resources (database gone)" FAIL cloud/cleanup-leakcheck-db.json \
      "HTTP ${code}; clean up manually: gcloud biglake hive databases delete ${DB_ID} --catalog=${CATALOG_ID} --project=${GCP_PROJECT_ID}"
  fi

  local ccode
  ccode="$(mv_rest cleanup leakcheck-catalog GET "${CATALOG_URL}")"
  if [[ "${ccode}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (catalog gone)" PASS cloud/cleanup-leakcheck-catalog.json "HTTP 404"
  else
    mv_record cleanup "no leaked GCP resources (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
      "HTTP ${ccode}; check catalog: ${CATALOG_URL}"
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
mv_wait_ready deps BigLakeHiveCatalog "${CATALOG_K8S}" 300
mv_wait_http deps bucket "${BUCKET_URL}" 200
mv_wait_http deps catalog "${CATALOG_URL}" 200
mv_cli deps gcloud-describe-catalog -- gcloud biglake hive catalogs describe "${CATALOG_ID}" --project="${GCP_PROJECT_ID}" || true

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create BigLakeHiveDatabase "${DB_K8S}" 300
SNAP="$(mv_snapshot create BigLakeHiveDatabase "${DB_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

WANT_LOCATION_URI="$(jq -r '.spec.locationURI // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"

mv_wait_http create db "${DB_URL}" 200
mv_expect_jq create "cloud name matches expected" "${MV_RUN_DIR}/cloud/create-db.json" '.name' "${EXPECTED_DB_NAME}"
if [[ -n "${WANT_LOCATION_URI}" ]]; then
  mv_expect_jq create "cloud locationUri matches spec" "${MV_RUN_DIR}/cloud/create-db.json" '.locationUri' "${WANT_LOCATION_URI}"
fi
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-db.json" '.description' "${WANT_DESCRIPTION}"
fi

if jq -e '.spec.parameters' "${SNAP}" >/dev/null 2>&1; then
  for k in $(jq -r '.spec.parameters | keys[]' "${SNAP}"); do
    val="$(jq -r --arg k "$k" '.spec.parameters[$k]' "${SNAP}")"
    mv_expect_jq create "cloud parameters.${k} matches spec" "${MV_RUN_DIR}/cloud/create-db.json" ".parameters[\"${k}\"]" "${val}"
  done
fi

mv_cli create gcloud-describe-db -- gcloud biglake hive databases describe "${DB_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-describe-db.txt" "${MV_RUN_DIR}/cloud/create-gcloud-describe-db.json"
mv_expect_jq create "gcloud describe database matches name" "${MV_RUN_DIR}/cloud/create-gcloud-describe-db.json" '.name' "${EXPECTED_DB_NAME}"

mv_cli create gcloud-list-dbs -- gcloud biglake hive databases list --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-list-dbs.txt" "${MV_RUN_DIR}/cloud/create-gcloud-list-dbs.json"
FOUND_DB="$(jq -r --arg NAME "${EXPECTED_DB_NAME}" '[.[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-gcloud-list-dbs.json")"
mv_expect_eq create "gcloud list databases includes the database" "${EXPECTED_DB_NAME}" "${FOUND_DB}" cloud/create-gcloud-list-dbs.json

if [[ "${FUNCTIONAL}" == "1" ]]; then
  # Functional check: Create a BigLake Hive table inside this database using gcloud
  TABLE_JSON="${MV_RUN_DIR}/cloud/create-table-req.json"
  TABLE_LOCATION="gs://${BUCKET_NAME}/db/${TABLE_ID}"
  cat <<EOF >"${TABLE_JSON}"
{
  "name": "${TABLE_ID}",
  "storageDescriptor": {
    "locationUri": "${TABLE_LOCATION}",
    "inputFormat": "org.apache.hadoop.mapred.TextInputFormat",
    "outputFormat": "org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat",
    "serdeInfo": {
      "name": "${TABLE_ID}",
      "serializationLib": "org.apache.hadoop.hive.serde2.lazy.LazySimpleSerDe"
    },
    "columns": [
      {"name": "id", "type": "int"},
      {"name": "name", "type": "string"}
    ]
  },
  "parameters": {
    "EXTERNAL": "TRUE"
  }
}
EOF
  if mv_cli create functional-create-table -- gcloud biglake hive tables create "${TABLE_ID}" \
      --catalog="${CATALOG_ID}" --database="${DB_ID}" --project="${GCP_PROJECT_ID}" \
      --file="${TABLE_JSON}"; then
    mv_record create "functional check: create table in database" PASS cloud/create-functional-create-table.txt ""
  else
    mv_record create "functional check: create table in database" FAIL cloud/create-functional-create-table.txt "failed to create table in database"
  fi

  # Verify table exists in database
  mv_cli create functional-describe-table -- gcloud biglake hive tables describe "${TABLE_ID}" \
      --catalog="${CATALOG_ID}" --database="${DB_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
  mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-functional-describe-table.txt" "${MV_RUN_DIR}/cloud/create-functional-describe-table.json"
  mv_expect_jq create "functional check: describe table matches name" "${MV_RUN_DIR}/cloud/create-functional-describe-table.json" \
    '.name' "${EXPECTED_DB_NAME}/tables/${TABLE_ID}"

  # Delete the table so the database can be updated and eventually deleted
  if mv_cli create functional-delete-table -- gcloud biglake hive tables delete "${TABLE_ID}" \
      --catalog="${CATALOG_ID}" --database="${DB_ID}" --project="${GCP_PROJECT_ID}" --quiet; then
    mv_record create "functional check: cleanup table in database" PASS cloud/create-functional-delete-table.txt ""
  else
    mv_record create "functional check: cleanup table in database" FAIL cloud/create-functional-delete-table.txt "failed to delete table"
  fi
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
# Give the apiserver a moment to bump generation before polling for Ready.
sleep 2
mv_wait_ready update BigLakeHiveDatabase "${DB_K8S}" 300
SNAP="$(mv_snapshot update BigLakeHiveDatabase "${DB_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_DESCRIPTION_UPDATE="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_LOCATION_URI_UPDATE="$(jq -r '.spec.locationURI // empty' "${SNAP}")"

mv_wait_http update db "${DB_URL}" 200
if [[ -n "${WANT_DESCRIPTION_UPDATE}" ]]; then
  mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-db.json" '.description' "${WANT_DESCRIPTION_UPDATE}"
fi
if [[ -n "${WANT_LOCATION_URI_UPDATE}" ]]; then
  mv_expect_jq update "cloud locationUri retained" "${MV_RUN_DIR}/cloud/update-db.json" '.locationUri' "${WANT_LOCATION_URI_UPDATE}"
fi

if jq -e '.spec.parameters' "${SNAP}" >/dev/null 2>&1; then
  for k in $(jq -r '.spec.parameters | keys[]' "${SNAP}"); do
    val="$(jq -r --arg k "$k" '.spec.parameters[$k]' "${SNAP}")"
    mv_expect_jq update "cloud parameters.${k} updated to new spec" "${MV_RUN_DIR}/cloud/update-db.json" ".parameters[\"${k}\"]" "${val}"
  done
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady BigLakeHiveDatabase "${DB_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete BigLakeHiveDatabase "${DB_K8S}" 300
mv_wait_http delete db-gone "${DB_URL}" 404

# Cloud Audit Logs check: BigLake Hive Database operations logged under HiveMetastoreService
# Note: includes functional check table create/delete if FUNCTIONAL=1
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_expect_audit delete audit-db "databases/${DB_ID}" \
    "CreateHiveDatabase,CreateHiveTable,DeleteHiveTable,UpdateHiveDatabase,DeleteHiveDatabase"
else
  mv_expect_audit delete audit-db "databases/${DB_ID}" \
    "CreateHiveDatabase,UpdateHiveDatabase,DeleteHiveDatabase"
fi

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown BigLakeHiveCatalog "${CATALOG_K8S}" 300
mv_wait_http teardown catalog-gone "${CATALOG_URL}" 404

# Clean any objects created in storage bucket
gcloud storage rm -r "gs://${BUCKET_NAME}/**" >/dev/null 2>&1 || true
mv_delete teardown StorageBucket "${BUCKET_NAME}" 300
mv_wait_http teardown bucket-gone "${BUCKET_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
