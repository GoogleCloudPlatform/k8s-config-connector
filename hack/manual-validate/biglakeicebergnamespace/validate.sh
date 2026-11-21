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
#   FUNCTIONAL=0     skip the functional check (listing tables in the namespace)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/bigquerybiglake/v1alpha1/biglakeicebergnamespace/biglakeicebergnamespace-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init BigLakeIcebergNamespace
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Names derived from the fixture (resourceIDs) so out-of-band checks never ask KCC.
if [[ "${FIXTURE_DIR}" == *"minimal"* ]]; then
  BUCKET_K8S="${GCP_PROJECT_ID}-ns-min-${MV_UNIQUE_ID}"
  CATALOG_K8S="biglakeicebergcatalog-ns-min-${MV_UNIQUE_ID}"
  NS_K8S="biglakeicebergnamespace-min-${MV_UNIQUE_ID}"
  BUCKET_ID="${BUCKET_K8S}"
  CATALOG_ID="${CATALOG_K8S}"
  NS_ID="nsmin${MV_UNIQUE_ID}"
else
  BUCKET_K8S="${GCP_PROJECT_ID}-ns-max-${MV_UNIQUE_ID}"
  CATALOG_K8S="biglakeicebergcatalog-ns-max-${MV_UNIQUE_ID}"
  NS_K8S="biglakeicebergnamespace-max-${MV_UNIQUE_ID}"
  BUCKET_ID="${BUCKET_K8S}"
  CATALOG_ID="${CATALOG_K8S}"
  NS_ID="nsmax${MV_UNIQUE_ID}"
fi

BUCKET_URL="https://storage.googleapis.com/storage/v1/b/${BUCKET_ID}"
CATALOG_URL="https://biglake.googleapis.com/iceberg/v1/restcatalog/extensions/projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}"
NS_URL="https://biglake.googleapis.com/iceberg/v1/restcatalog/v1/projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}/namespaces/${NS_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get biglakeicebergnamespace "${NS_K8S}" >/dev/null 2>&1 && mv_delete cleanup BigLakeIcebergNamespace "${NS_K8S}" 300
    mv_kubectl get biglakeicebergcatalog "${CATALOG_K8S}" >/dev/null 2>&1 && mv_delete cleanup BigLakeIcebergCatalog "${CATALOG_K8S}" 300
    mv_kubectl get storagebucket "${BUCKET_K8S}" >/dev/null 2>&1 && mv_delete cleanup StorageBucket "${BUCKET_K8S}" 300
  fi

  local code_ns code_cat code_bucket
  code_ns="$(mv_rest cleanup leakcheck-ns GET "${NS_URL}")"
  if [[ "${code_ns}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (namespace gone)" PASS cloud/cleanup-leakcheck-ns.json "HTTP 404"
  elif [[ "${code_ns}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (namespace gone)" FAIL cloud/cleanup-leakcheck-ns.json \
      "could not verify (HTTP ${code_ns}); check manually: gcloud biglake iceberg namespaces describe ${NS_ID} --catalog=${CATALOG_ID} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP resources (namespace gone)" FAIL cloud/cleanup-leakcheck-ns.json \
      "HTTP ${code_ns}; clean up manually: gcloud biglake iceberg namespaces delete ${NS_ID} --catalog=${CATALOG_ID} --project=${GCP_PROJECT_ID}"
  fi

  code_cat="$(mv_rest cleanup leakcheck-catalog GET "${CATALOG_URL}")"
  if [[ "${code_cat}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (catalog gone)" PASS cloud/cleanup-leakcheck-catalog.json "HTTP 404"
  elif [[ "${code_cat}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
      "could not verify (HTTP ${code_cat}); check manually: gcloud biglake iceberg catalogs describe ${CATALOG_ID} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP resources (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
      "HTTP ${code_cat}; clean up manually: gcloud biglake iceberg catalogs delete ${CATALOG_ID} --project=${GCP_PROJECT_ID}"
  fi

  code_bucket="$(mv_rest cleanup leakcheck-bucket GET "${BUCKET_URL}")"
  if [[ "${code_bucket}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (bucket gone)" PASS cloud/cleanup-leakcheck-bucket.json "HTTP 404"
  elif [[ "${code_bucket}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (bucket gone)" FAIL cloud/cleanup-leakcheck-bucket.json \
      "could not verify (HTTP ${code_bucket}); check manually: gcloud storage buckets describe gs://${BUCKET_ID}"
  else
    mv_record cleanup "no leaked GCP resources (bucket gone)" FAIL cloud/cleanup-leakcheck-bucket.json \
      "HTTP ${code_bucket}; clean up manually: gcloud storage rm --recursive gs://${BUCKET_ID}"
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
mv_wait_ready deps StorageBucket "${BUCKET_K8S}" 300
mv_wait_ready deps BigLakeIcebergCatalog "${CATALOG_K8S}" 300
mv_wait_http deps bucket "${BUCKET_URL}" 200
mv_wait_http deps catalog "${CATALOG_URL}" 200
mv_cli deps gcloud-describe-bucket -- gcloud storage buckets describe "gs://${BUCKET_ID}" --format=json || true
mv_cli deps gcloud-describe-catalog -- gcloud biglake iceberg catalogs describe "${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --format=json || true

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create BigLakeIcebergNamespace "${NS_K8S}" 300
SNAP="$(mv_snapshot create BigLakeIcebergNamespace "${NS_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"
WANT_NS="$(jq -r '[.spec.namespace[]?] | join(",")' "${SNAP}")"
WANT_PROPERTIES="$(jq -r '.spec.properties // {}' "${SNAP}")"

mv_wait_http create ns "${NS_URL}" 200
mv_expect_jq create "cloud namespace matches spec" "${MV_RUN_DIR}/cloud/create-ns.json" '[.namespace[]?] | join(",")' "${WANT_NS}"

for k in $(echo "${WANT_PROPERTIES}" | jq -r 'keys[]?'); do
  v="$(echo "${WANT_PROPERTIES}" | jq -r --arg k "${k}" '.[$k]')"
  mv_expect_jq create "cloud property '${k}' matches spec" "${MV_RUN_DIR}/cloud/create-ns.json" ".properties[\"${k}\"]" "${v}"
done

WANT_LOCATION="gs://${BUCKET_ID}/${NS_ID}"
mv_expect_jq create "cloud property 'location' is server-populated" "${MV_RUN_DIR}/cloud/create-ns.json" '.properties.location' "${WANT_LOCATION}"

mv_cli create gcloud-describe-ns -- gcloud biglake iceberg namespaces describe "${NS_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
if grep -q "\"${NS_ID}\"" "${MV_RUN_DIR}/cloud/create-gcloud-describe-ns.txt"; then
  mv_record create "gcloud biglake iceberg namespaces describe succeeds" PASS cloud/create-gcloud-describe-ns.txt ""
else
  mv_record create "gcloud biglake iceberg namespaces describe succeeds" FAIL cloud/create-gcloud-describe-ns.txt "namespace id not found in describe output"
fi

mv_cli create gcloud-list-ns -- gcloud biglake iceberg namespaces list --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
if grep -q "namespaces/${NS_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-list-ns.txt"; then
  mv_record create "gcloud biglake iceberg namespaces list contains the namespace" PASS cloud/create-gcloud-list-ns.txt ""
else
  mv_record create "gcloud biglake iceberg namespaces list contains the namespace" FAIL cloud/create-gcloud-list-ns.txt "namespace id not in list output"
fi

if [[ "${FUNCTIONAL}" == "1" ]]; then
  if mv_cli create gcloud-list-tables -- gcloud biglake iceberg tables list --catalog="${CATALOG_ID}" --namespace="${NS_ID}" --project="${GCP_PROJECT_ID}" --format=json; then
    mv_record create "functional: list tables in namespace succeeds (valid and queryable)" PASS cloud/create-gcloud-list-tables.txt ""
  else
    mv_record create "functional: list tables in namespace succeeds (valid and queryable)" FAIL cloud/create-gcloud-list-tables.txt "failed to list tables in namespace"
  fi
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
# Give the apiserver a moment to bump generation before polling for Ready.
sleep 2
mv_wait_ready update BigLakeIcebergNamespace "${NS_K8S}" 300
SNAP="$(mv_snapshot update BigLakeIcebergNamespace "${NS_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"
WANT_NS="$(jq -r '[.spec.namespace[]?] | join(",")' "${SNAP}")"
WANT_PROPERTIES="$(jq -r '.spec.properties // {}' "${SNAP}")"

mv_wait_http update ns "${NS_URL}" 200
mv_expect_jq update "cloud namespace still matches spec" "${MV_RUN_DIR}/cloud/update-ns.json" '[.namespace[]?] | join(",")' "${WANT_NS}"

for k in $(echo "${WANT_PROPERTIES}" | jq -r 'keys[]?'); do
  v="$(echo "${WANT_PROPERTIES}" | jq -r --arg k "${k}" '.[$k]')"
  mv_expect_jq update "cloud updated property '${k}' matches spec" "${MV_RUN_DIR}/cloud/update-ns.json" ".properties[\"${k}\"]" "${v}"
done

OLD_PURPOSE="$(jq -r '.spec.properties.purpose // empty' "${MV_RUN_DIR}/k8s/create-biglakeicebergnamespace-${NS_K8S}.json")"
if [[ -n "${OLD_PURPOSE}" ]]; then
  mv_expect_jq update "removed property 'purpose' is absent in cloud" "${MV_RUN_DIR}/cloud/update-ns.json" '.properties.purpose // "absent"' "absent"
fi

mv_expect_jq update "cloud property 'location' is preserved" "${MV_RUN_DIR}/cloud/update-ns.json" '.properties.location' "${WANT_LOCATION}"

mv_cli update gcloud-describe-ns -- gcloud biglake iceberg namespaces describe "${NS_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --format=json || true

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady BigLakeIcebergNamespace "${NS_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete BigLakeIcebergNamespace "${NS_K8S}" 300
mv_wait_http delete ns-gone "${NS_URL}" 404

if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli delete gcloud-list-tables -- gcloud biglake iceberg tables list --catalog="${CATALOG_ID}" --namespace="${NS_ID}" --project="${GCP_PROJECT_ID}" || true
  if grep -qi 'NoSuchNamespaceException\|NOT_FOUND\|does not exist' "${MV_RUN_DIR}/cloud/delete-gcloud-list-tables.txt"; then
    mv_record delete "functional: list tables fails after namespace delete" PASS cloud/delete-gcloud-list-tables.txt ""
  else
    mv_record delete "functional: list tables fails after namespace delete" FAIL cloud/delete-gcloud-list-tables.txt "did not return expected error"
  fi
fi

# Cloud Audit Logs check: exactly one Create, one Update, one Delete.
mv_expect_audit delete audit-ns "namespaces/${NS_ID}" \
  "CreateIcebergNamespace,UpdateIcebergNamespace,DeleteIcebergNamespace"

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown BigLakeIcebergCatalog "${CATALOG_K8S}" 300
mv_wait_http teardown catalog-gone "${CATALOG_URL}" 404
mv_delete teardown StorageBucket "${BUCKET_K8S}" 300
mv_wait_http teardown bucket-gone "${BUCKET_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
