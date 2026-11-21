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
#   FUNCTIONAL=1     run functional check (list tables on namespace)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/bigquerybiglake/v1alpha1/biglakeicebergnamespace/biglakeicebergnamespace-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init BigLakeIcebergNamespace
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Render manifests and derive resource names from fixture so out-of-band checks never ask KCC.
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

BUCKET_K8S="$(python3 -c "import yaml; docs = list(yaml.safe_load_all(open('${DEPS}'))); buckets = [d['metadata']['name'] for d in docs if d.get('kind') == 'StorageBucket']; print(buckets[0] if buckets else '')")"
BUCKET_NAME="${BUCKET_K8S}"
CATALOG_K8S="$(python3 -c "import yaml; docs = list(yaml.safe_load_all(open('${DEPS}'))); cats = [d['metadata']['name'] for d in docs if d.get('kind') == 'BigLakeIcebergCatalog']; print(cats[0] if cats else '')")"
CATALOG_ID="${CATALOG_K8S}"

NS_K8S="$(python3 -c "import yaml; d = yaml.safe_load(open('${CREATE}')); print(d['metadata']['name'])")"
NS_ID="$(python3 -c "import yaml; d = yaml.safe_load(open('${CREATE}')); print(d.get('spec', {}).get('resourceID') or d['spec']['namespace'][0])")"

BUCKET_URL="https://storage.googleapis.com/storage/v1/b/${BUCKET_NAME}"
CATALOG_URL="https://biglake.googleapis.com/iceberg/v1/restcatalog/extensions/projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}"
NS_URL="https://biglake.googleapis.com/iceberg/v1/restcatalog/v1/projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}/namespaces/${NS_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    [[ -n "${NS_K8S:-}" ]] && mv_kubectl get biglakeicebergnamespace "${NS_K8S}" >/dev/null 2>&1 && mv_delete cleanup BigLakeIcebergNamespace "${NS_K8S}" 300
    [[ -n "${CATALOG_K8S:-}" ]] && mv_kubectl get biglakeicebergcatalog "${CATALOG_K8S}" >/dev/null 2>&1 && mv_delete cleanup BigLakeIcebergCatalog "${CATALOG_K8S}" 300
    [[ -n "${BUCKET_K8S:-}" ]] && mv_kubectl get storagebucket "${BUCKET_K8S}" >/dev/null 2>&1 && mv_delete cleanup StorageBucket "${BUCKET_K8S}" 300
  fi
  local code
  if [[ -n "${NS_URL:-}" ]]; then
    code="$(mv_rest cleanup leakcheck-namespace GET "${NS_URL}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (namespace gone)" PASS cloud/cleanup-leakcheck-namespace.json "HTTP 404"
    elif [[ "${code}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (namespace gone)" FAIL cloud/cleanup-leakcheck-namespace.json \
        "could not verify (HTTP ${code}); check manually: gcloud biglake iceberg namespaces describe ${NS_ID} --catalog=${CATALOG_ID} --project=${GCP_PROJECT_ID}"
    else
      mv_record cleanup "no leaked GCP resources (namespace gone)" FAIL cloud/cleanup-leakcheck-namespace.json \
        "HTTP ${code}; clean up manually: gcloud biglake iceberg namespaces delete ${NS_ID} --catalog=${CATALOG_ID} --project=${GCP_PROJECT_ID} --quiet"
    fi
  fi
  if [[ -n "${CATALOG_URL:-}" ]]; then
    code="$(mv_rest cleanup leakcheck-catalog GET "${CATALOG_URL}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (catalog gone)" PASS cloud/cleanup-leakcheck-catalog.json "HTTP 404"
    elif [[ "${code}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
        "could not verify (HTTP ${code}); check manually: gcloud biglake iceberg catalogs describe ${CATALOG_ID} --project=${GCP_PROJECT_ID}"
    else
      mv_record cleanup "no leaked GCP resources (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
        "HTTP ${code}; clean up manually: gcloud biglake iceberg catalogs delete ${CATALOG_ID} --project=${GCP_PROJECT_ID} --quiet"
    fi
  fi
  if [[ -n "${BUCKET_URL:-}" ]]; then
    code="$(mv_rest cleanup leakcheck-bucket GET "${BUCKET_URL}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (bucket gone)" PASS cloud/cleanup-leakcheck-bucket.json "HTTP 404"
    elif [[ "${code}" != "200" ]]; then
      mv_record cleanup "no leaked GCP resources (bucket gone)" FAIL cloud/cleanup-leakcheck-bucket.json \
        "could not verify (HTTP ${code}); check manually: gcloud storage buckets describe gs://${BUCKET_NAME}"
    else
      mv_record cleanup "no leaked GCP resources (bucket gone)" FAIL cloud/cleanup-leakcheck-bucket.json \
        "HTTP ${code}; clean up manually: gcloud storage rm -r gs://${BUCKET_NAME}"
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
# Every API the controller, the dependencies and the out-of-band checks call.
mv_enable_apis setup biglake.googleapis.com storage.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
mv_apply deps "${DEPS}"
if [[ -n "${BUCKET_K8S:-}" ]]; then
  mv_wait_ready deps StorageBucket "${BUCKET_K8S}" 300
  mv_wait_http deps bucket "${BUCKET_URL}" 200
fi
if [[ -n "${CATALOG_K8S:-}" ]]; then
  mv_wait_ready deps BigLakeIcebergCatalog "${CATALOG_K8S}" 300
  mv_wait_http deps catalog "${CATALOG_URL}" 200
  mv_cli deps gcloud-describe-catalog -- gcloud biglake iceberg catalogs describe "${CATALOG_ID}" --project="${GCP_PROJECT_ID}"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create BigLakeIcebergNamespace "${NS_K8S}" 300
SNAP="$(mv_snapshot create BigLakeIcebergNamespace "${NS_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

WANT_NAMESPACE="$(jq -r '[.spec.namespace[]] | join(",")' "${SNAP}")"

mv_wait_http create namespace "${NS_URL}" 200
mv_expect_jq create "cloud namespace matches spec" "${MV_RUN_DIR}/cloud/create-namespace.json" '.namespace | join(",")' "${WANT_NAMESPACE}"

# Verify each spec property matches in cloud
for key in $(jq -r '(.spec.properties // {}) | keys[]' "${SNAP}"); do
  want_val="$(jq -r ".spec.properties[\"${key}\"]" "${SNAP}")"
  mv_expect_jq create "cloud property ${key} matches spec" "${MV_RUN_DIR}/cloud/create-namespace.json" ".properties[\"${key}\"]" "${want_val}"
done

# Server-populated location property verification
mv_expect_jq create "cloud property location is populated by server" "${MV_RUN_DIR}/cloud/create-namespace.json" '(.properties.location | startswith("gs://")) | tostring' "true"

# Out-of-band gcloud describe
mv_cli create gcloud-describe -- gcloud biglake iceberg namespaces describe "${NS_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}"
if grep -q "${NS_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-describe.txt"; then
  mv_record create "gcloud describe namespace succeeds" PASS cloud/create-gcloud-describe.txt ""
else
  mv_record create "gcloud describe namespace succeeds" FAIL cloud/create-gcloud-describe.txt "namespace not found"
fi

# Out-of-band gcloud list namespaces
mv_cli create gcloud-list-namespaces -- gcloud biglake iceberg namespaces list --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}"
if grep -q "${NS_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-list-namespaces.txt"; then
  mv_record create "gcloud list namespaces contains namespace" PASS cloud/create-gcloud-list-namespaces.txt ""
else
  mv_record create "gcloud list namespaces contains namespace" FAIL cloud/create-gcloud-list-namespaces.txt "namespace not found in list"
fi

# Functional check: query tables under the namespace
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli create gcloud-list-tables -- gcloud biglake iceberg tables list --namespace="${NS_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}"
  if grep -q 'exit_code: 0' "${MV_RUN_DIR}/cloud/create-gcloud-list-tables.txt"; then
    mv_record create "functional: gcloud tables list succeeds on namespace" PASS cloud/create-gcloud-list-tables.txt ""
  else
    mv_record create "functional: gcloud tables list succeeds on namespace" FAIL cloud/create-gcloud-list-tables.txt "exit_code != 0"
  fi
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update BigLakeIcebergNamespace "${NS_K8S}" 300
SNAP="$(mv_snapshot update BigLakeIcebergNamespace "${NS_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update namespace "${NS_URL}" 200
mv_expect_jq update "cloud namespace still matches spec" "${MV_RUN_DIR}/cloud/update-namespace.json" '.namespace | join(",")' "${WANT_NAMESPACE}"

# Verify updated and added properties
for key in $(jq -r '(.spec.properties // {}) | keys[]' "${SNAP}"); do
  want_val="$(jq -r ".spec.properties[\"${key}\"]" "${SNAP}")"
  mv_expect_jq update "cloud property ${key} matches updated spec" "${MV_RUN_DIR}/cloud/update-namespace.json" ".properties[\"${key}\"]" "${want_val}"
done

# Verify removed properties
CREATE_SNAP="${MV_RUN_DIR}/k8s/create-biglakeicebergnamespace-${NS_K8S}.json"
for key in $(jq -r '(.spec.properties // {}) | keys[]' "${CREATE_SNAP}"); do
  if [[ "$(jq -r ".spec.properties[\"${key}\"] // empty" "${SNAP}")" == "" ]]; then
    mv_expect_jq update "removed property ${key} is absent from cloud" "${MV_RUN_DIR}/cloud/update-namespace.json" ".properties[\"${key}\"] // \"absent\"" "absent"
  fi
done

# Location property preserved
mv_expect_jq update "cloud property location is preserved" "${MV_RUN_DIR}/cloud/update-namespace.json" '(.properties.location | startswith("gs://")) | tostring' "true"

# Out-of-band gcloud describe
mv_cli update gcloud-describe -- gcloud biglake iceberg namespaces describe "${NS_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}"
if grep -q "${NS_ID}" "${MV_RUN_DIR}/cloud/update-gcloud-describe.txt"; then
  mv_record update "gcloud describe namespace succeeds after update" PASS cloud/update-gcloud-describe.txt ""
else
  mv_record update "gcloud describe namespace succeeds after update" FAIL cloud/update-gcloud-describe.txt "namespace not found"
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady BigLakeIcebergNamespace "${NS_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete BigLakeIcebergNamespace "${NS_K8S}" 300
mv_wait_http delete namespace-gone "${NS_URL}" 404

# Audit log check: Create, Update, Delete on Iceberg namespace
mv_expect_audit delete audit-namespace "namespaces/${NS_ID}" \
  "CreateIcebergNamespace,UpdateIcebergNamespace,DeleteIcebergNamespace"

if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli delete gcloud-list-tables -- gcloud biglake iceberg tables list --namespace="${NS_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" || true
  if grep -qi 'not found\|NoSuchNamespaceException\|404' "${MV_RUN_DIR}/cloud/delete-gcloud-list-tables.txt"; then
    mv_record delete "functional: tables list fails after namespace deletion" PASS cloud/delete-gcloud-list-tables.txt ""
  else
    mv_record delete "functional: tables list fails after namespace deletion" FAIL cloud/delete-gcloud-list-tables.txt "did not fail with NoSuchNamespace"
  fi
fi

mv_cli delete gcloud-list-namespaces -- gcloud biglake iceberg namespaces list --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}"
if grep -q "${NS_ID}" "${MV_RUN_DIR}/cloud/delete-gcloud-list-namespaces.txt"; then
  mv_record delete "gcloud list namespaces does not contain namespace" FAIL cloud/delete-gcloud-list-namespaces.txt "namespace still present in list"
else
  mv_record delete "gcloud list namespaces does not contain namespace" PASS cloud/delete-gcloud-list-namespaces.txt ""
fi

# --- Phase: dependency teardown ----------------------------------------------
if [[ -n "${CATALOG_K8S:-}" ]]; then
  mv_delete teardown BigLakeIcebergCatalog "${CATALOG_K8S}" 300
  mv_wait_http teardown catalog-gone "${CATALOG_URL}" 404
fi
if [[ -n "${BUCKET_K8S:-}" ]]; then
  mv_delete teardown StorageBucket "${BUCKET_K8S}" 300
  mv_wait_http teardown bucket-gone "${BUCKET_URL}" 404
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
