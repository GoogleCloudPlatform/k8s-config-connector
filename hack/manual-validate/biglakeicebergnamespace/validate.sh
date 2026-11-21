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

# Human-style validation of the BigLakeIcebergNamespace controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/biglakeicebergnamespace/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the biglakeicebergnamespace-maximal test fixture)
#   FUNCTIONAL=0     skip table creation/deletion functional check
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/bigquerybiglake/v1alpha1/biglakeicebergnamespace/biglakeicebergnamespace-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init BigLakeIcebergNamespace
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Names derived from the fixture (resourceIDs) so out-of-band checks never ask KCC.
BUCKET_NAME="${GCP_PROJECT_ID}-ns-max-${MV_UNIQUE_ID}"
CATALOG_ID="biglakeicebergcatalog-ns-max-${MV_UNIQUE_ID}"
NS_ID="nsmax${MV_UNIQUE_ID}"
TABLE_ID="t_${MV_UNIQUE_ID}"

CATALOG_URL="https://biglake.googleapis.com/iceberg/v1/restcatalog/extensions/projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}?alt=json"
REST_CATALOG="https://biglake.googleapis.com/iceberg/v1/restcatalog/v1/projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}"
NS_URL="${REST_CATALOG}/namespaces/${NS_ID}?%24alt=json%3Benum-encoding%3Dint"
BUCKET_URL="https://storage.googleapis.com/storage/v1/b/${BUCKET_NAME}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  # Clean up any leftover functional check table if created
  gcloud biglake iceberg tables delete "${TABLE_ID}" --catalog="${CATALOG_ID}" --namespace="${NS_ID}" --project="${GCP_PROJECT_ID}" --quiet >/dev/null 2>&1 || true

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get biglakeicebergnamespace "biglakeicebergnamespace-max-${MV_UNIQUE_ID}" >/dev/null 2>&1 && \
      mv_delete cleanup BigLakeIcebergNamespace "biglakeicebergnamespace-max-${MV_UNIQUE_ID}" 300
    mv_kubectl get biglakeicebergcatalog "${CATALOG_ID}" >/dev/null 2>&1 && \
      mv_delete cleanup BigLakeIcebergCatalog "${CATALOG_ID}" 300
    # Storage bucket may contain objects created during functional test / Iceberg catalog metadata;
    # clean objects out so bucket deletion in KCC succeeds.
    gcloud storage rm -r "gs://${BUCKET_NAME}/**" >/dev/null 2>&1 || true
    mv_kubectl get storagebucket "${BUCKET_NAME}" >/dev/null 2>&1 && \
      mv_delete cleanup StorageBucket "${BUCKET_NAME}" 300
  fi

  local code
  code="$(mv_rest cleanup leakcheck-ns GET "${NS_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (namespace gone)" PASS cloud/cleanup-leakcheck-ns.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (namespace gone)" FAIL cloud/cleanup-leakcheck-ns.json \
      "could not verify (HTTP ${code}); check manually: gcloud biglake iceberg namespaces describe ${NS_ID} --catalog=${CATALOG_ID}"
  else
    mv_record cleanup "no leaked GCP resources (namespace gone)" FAIL cloud/cleanup-leakcheck-ns.json \
      "HTTP ${code}; clean up manually: gcloud biglake iceberg namespaces delete ${NS_ID} --catalog=${CATALOG_ID}"
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
mv_wait_ready deps BigLakeIcebergCatalog "${CATALOG_ID}" 300
mv_wait_http deps bucket "${BUCKET_URL}" 200
mv_wait_http deps catalog "${CATALOG_URL}" 200
mv_cli deps gcloud-describe-catalog -- gcloud biglake iceberg catalogs describe "${CATALOG_ID}" --project="${GCP_PROJECT_ID}" || true

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create BigLakeIcebergNamespace "biglakeicebergnamespace-max-${MV_UNIQUE_ID}" 300
SNAP="$(mv_snapshot create BigLakeIcebergNamespace "biglakeicebergnamespace-max-${MV_UNIQUE_ID}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"
WANT_NS="$(jq -r '[.spec.namespace[]?] | join(",")' "${SNAP}")"
WANT_OWNER="$(jq -r '.spec.properties.owner' "${SNAP}")"
WANT_PURPOSE="$(jq -r '.spec.properties.purpose' "${SNAP}")"

mv_wait_http create ns "${NS_URL}" 200
mv_expect_jq create "cloud namespace matches spec" "${MV_RUN_DIR}/cloud/create-ns.json" '[.namespace[]?] | join(",")' "${WANT_NS}"
mv_expect_jq create "cloud property owner matches spec" "${MV_RUN_DIR}/cloud/create-ns.json" '.properties.owner' "${WANT_OWNER}"
mv_expect_jq create "cloud property purpose matches spec" "${MV_RUN_DIR}/cloud/create-ns.json" '.properties.purpose' "${WANT_PURPOSE}"
# Real GCP Iceberg Catalog automatically populates the location property
EXPECTED_LOCATION="gs://${BUCKET_NAME}/${NS_ID}"
mv_expect_jq create "cloud property location matches bucket/ns" "${MV_RUN_DIR}/cloud/create-ns.json" '.properties.location' "${EXPECTED_LOCATION}"

mv_cli create gcloud-describe-ns -- gcloud biglake iceberg namespaces describe "${NS_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" || true
if grep -q "${NS_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-describe-ns.txt"; then
  mv_record create "gcloud describe namespace succeeds" PASS cloud/create-gcloud-describe-ns.txt ""
else
  mv_record create "gcloud describe namespace succeeds" FAIL cloud/create-gcloud-describe-ns.txt "namespace id not found in describe output"
fi

mv_cli create gcloud-list-ns -- gcloud biglake iceberg namespaces list --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" || true
if grep -q "${NS_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-list-ns.txt"; then
  mv_record create "gcloud list namespaces includes the namespace" PASS cloud/create-gcloud-list-ns.txt ""
else
  mv_record create "gcloud list namespaces includes the namespace" FAIL cloud/create-gcloud-list-ns.txt "namespace id not in list output"
fi

if [[ "${FUNCTIONAL}" == "1" ]]; then
  # Create an Iceberg table inside this namespace using gcloud to verify the namespace works
  TABLE_JSON="${MV_RUN_DIR}/cloud/create-table-req.json"
  cat <<EOF >"${TABLE_JSON}"
{
  "name": "${TABLE_ID}",
  "location": "gs://${BUCKET_NAME}/${NS_ID}/${TABLE_ID}",
  "schema": {
    "type": "struct",
    "schema-id": 0,
    "fields": [
      {
        "id": 1,
        "name": "id",
        "required": true,
        "type": "long"
      }
    ]
  }
}
EOF
  if mv_cli create functional-create-table -- gcloud biglake iceberg tables create \
      --catalog="${CATALOG_ID}" --namespace="${NS_ID}" --project="${GCP_PROJECT_ID}" \
      --create-from-file="${TABLE_JSON}"; then
    mv_record create "functional check: create table in namespace" PASS cloud/create-functional-create-table.txt ""
  else
    mv_record create "functional check: create table in namespace" FAIL cloud/create-functional-create-table.txt "failed to create table in namespace"
  fi

  # Delete the table so the namespace can be updated and eventually deleted
  if mv_cli create functional-delete-table -- gcloud biglake iceberg tables delete "${TABLE_ID}" \
      --catalog="${CATALOG_ID}" --namespace="${NS_ID}" --project="${GCP_PROJECT_ID}" --quiet; then
    mv_record create "functional check: cleanup table in namespace" PASS cloud/create-functional-delete-table.txt ""
  else
    mv_record create "functional check: cleanup table in namespace" FAIL cloud/create-functional-delete-table.txt "failed to delete table"
  fi
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
# Give the apiserver a moment to bump generation before polling for Ready.
sleep 2
mv_wait_ready update BigLakeIcebergNamespace "biglakeicebergnamespace-max-${MV_UNIQUE_ID}" 300
SNAP="$(mv_snapshot update BigLakeIcebergNamespace "biglakeicebergnamespace-max-${MV_UNIQUE_ID}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

WANT_OWNER_UPDATE="$(jq -r '.spec.properties.owner' "${SNAP}")"
WANT_ENV_UPDATE="$(jq -r '.spec.properties.env' "${SNAP}")"

mv_wait_http update ns "${NS_URL}" 200
mv_expect_jq update "cloud property owner updated to new spec" "${MV_RUN_DIR}/cloud/update-ns.json" '.properties.owner' "${WANT_OWNER_UPDATE}"
mv_expect_jq update "cloud property env added from new spec" "${MV_RUN_DIR}/cloud/update-ns.json" '.properties.env' "${WANT_ENV_UPDATE}"
mv_expect_jq update "cloud property purpose removed as expected" "${MV_RUN_DIR}/cloud/update-ns.json" '.properties.purpose // "null"' "null"
mv_expect_jq update "cloud property location retained" "${MV_RUN_DIR}/cloud/update-ns.json" '.properties.location' "${EXPECTED_LOCATION}"

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady BigLakeIcebergNamespace "biglakeicebergnamespace-max-${MV_UNIQUE_ID}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete BigLakeIcebergNamespace "biglakeicebergnamespace-max-${MV_UNIQUE_ID}" 300
mv_wait_http delete ns-gone "${NS_URL}" 404

# Cloud Audit Logs check: Iceberg REST catalog uses IcebergCatalogService
# Note: includes functional check table create/delete if FUNCTIONAL=1
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_expect_audit delete audit-ns "namespaces/${NS_ID}" \
    "CreateIcebergNamespace,CreateIcebergTable,DeleteIcebergTable,UpdateIcebergNamespace,DeleteIcebergNamespace"
else
  mv_expect_audit delete audit-ns "namespaces/${NS_ID}" \
    "CreateIcebergNamespace,UpdateIcebergNamespace,DeleteIcebergNamespace"
fi

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown BigLakeIcebergCatalog "${CATALOG_ID}" 300
mv_wait_http teardown catalog-gone "${CATALOG_URL}" 404

# Clean any objects created by functional table check in storage bucket
gcloud storage rm -r "gs://${BUCKET_NAME}/**" >/dev/null 2>&1 || true
mv_delete teardown StorageBucket "${BUCKET_NAME}" 300
mv_wait_http teardown bucket-gone "${BUCKET_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
