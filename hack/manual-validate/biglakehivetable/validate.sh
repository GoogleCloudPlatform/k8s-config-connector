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

# Human-style validation of the BigLakeHiveTable controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/biglakehivetable/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the biglakehivetable-maximal test fixture)
#   FUNCTIONAL=0     skip table IAM policy functional check
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/bigquerybiglake/v1alpha1/biglakehivetable/biglakehivetable-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init BigLakeHiveTable
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Inspect the fixture's create.yaml and dependencies.yaml to extract naming patterns
TABLE_K8S="$(grep -E '^  name:' "${FIXTURE_DIR}/create.yaml" | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g; s/\${projectId}/${GCP_PROJECT_ID}/g")"
TABLE_ID="$(grep -E 'resourceID:' "${FIXTURE_DIR}/create.yaml" | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g; s/\${projectId}/${GCP_PROJECT_ID}/g")"
if [[ -z "${TABLE_ID}" ]]; then
  TABLE_ID="${TABLE_K8S}"
fi

DB_K8S="$(grep -A 2 'parentRef:' "${FIXTURE_DIR}/create.yaml" | grep -E 'name:' | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g; s/\${projectId}/${GCP_PROJECT_ID}/g")"
DB_ID="$(grep -A 10 'kind: BigLakeHiveDatabase' "${FIXTURE_DIR}/dependencies.yaml" | grep -E 'resourceID:' | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g; s/\${projectId}/${GCP_PROJECT_ID}/g")"
if [[ -z "${DB_ID}" ]]; then
  DB_ID="${DB_K8S}"
fi

CATALOG_K8S="$(grep -A 10 'kind: BigLakeHiveDatabase' "${FIXTURE_DIR}/dependencies.yaml" | grep -A 2 'parentRef:' | grep -E 'name:' | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g; s/\${projectId}/${GCP_PROJECT_ID}/g")"
CATALOG_ID="$(grep -A 10 'kind: BigLakeHiveCatalog' "${FIXTURE_DIR}/dependencies.yaml" | grep -E 'resourceID:' | head -1 | awk '{print $2}' | sed "s/\${uniqueId}/${MV_UNIQUE_ID}/g; s/\${projectId}/${GCP_PROJECT_ID}/g")"
if [[ -z "${CATALOG_ID}" ]]; then
  CATALOG_ID="${CATALOG_K8S}"
fi

BUCKET_NAME="$(grep -A 10 'kind: StorageBucket' "${FIXTURE_DIR}/dependencies.yaml" | grep -E 'name:' | head -1 | awk '{print $2}' | sed "s/\${projectId}/${GCP_PROJECT_ID}/g; s/\${uniqueId}/${MV_UNIQUE_ID}/g")"

EXPECTED_TABLE_NAME="projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}/databases/${DB_ID}/tables/${TABLE_ID}"
TABLE_URL="https://biglake.googleapis.com/hive/v1/projects/${GCP_PROJECT_ID}/catalogs/${CATALOG_ID}/databases/${DB_ID}/tables/${TABLE_ID}"
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

  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get biglakehivetable "${TABLE_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup BigLakeHiveTable "${TABLE_K8S}" 300

    mv_kubectl get biglakehivedatabase "${DB_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup BigLakeHiveDatabase "${DB_K8S}" 300

    mv_kubectl get biglakehivecatalog "${CATALOG_K8S}" >/dev/null 2>&1 && \
      mv_delete cleanup BigLakeHiveCatalog "${CATALOG_K8S}" 300

    # Storage bucket may contain objects created during Hive table metadata;
    # clean objects out so bucket deletion in KCC succeeds.
    gcloud storage rm -r "gs://${BUCKET_NAME}/**" >/dev/null 2>&1 || true
    mv_kubectl get storagebucket "${BUCKET_NAME}" >/dev/null 2>&1 && \
      mv_delete cleanup StorageBucket "${BUCKET_NAME}" 300
  fi

  local code
  code="$(mv_rest cleanup leakcheck-table GET "${TABLE_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (table gone)" PASS cloud/cleanup-leakcheck-table.json "HTTP 404"
  elif [[ "${code}" == "400" ]] && grep -q "does not exist" "${MV_RUN_DIR}/cloud/cleanup-leakcheck-table.json" 2>/dev/null; then
    mv_record cleanup "no leaked GCP resources (table gone)" PASS cloud/cleanup-leakcheck-table.json "HTTP 400 (parent database or catalog does not exist)"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (table gone)" FAIL cloud/cleanup-leakcheck-table.json \
      "could not verify (HTTP ${code}); check manually: gcloud biglake hive tables describe ${TABLE_ID} --database=${DB_ID} --catalog=${CATALOG_ID} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP resources (table gone)" FAIL cloud/cleanup-leakcheck-table.json \
      "HTTP ${code}; clean up manually: gcloud biglake hive tables delete ${TABLE_ID} --database=${DB_ID} --catalog=${CATALOG_ID} --project=${GCP_PROJECT_ID}"
  fi

  local dcode
  dcode="$(mv_rest cleanup leakcheck-db GET "${DB_URL}")"
  if [[ "${dcode}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (database gone)" PASS cloud/cleanup-leakcheck-db.json "HTTP 404"
  elif [[ "${dcode}" == "400" ]] && grep -q "does not exist" "${MV_RUN_DIR}/cloud/cleanup-leakcheck-db.json" 2>/dev/null; then
    mv_record cleanup "no leaked GCP resources (database gone)" PASS cloud/cleanup-leakcheck-db.json "HTTP 400 (parent catalog does not exist)"
  elif [[ "${dcode}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (database gone)" FAIL cloud/cleanup-leakcheck-db.json \
      "could not verify (HTTP ${dcode}); check database: ${DB_URL}"
  else
    mv_record cleanup "no leaked GCP resources (database gone)" FAIL cloud/cleanup-leakcheck-db.json \
      "HTTP ${dcode}; clean up manually: gcloud biglake hive databases delete ${DB_ID} --catalog=${CATALOG_ID} --project=${GCP_PROJECT_ID}"
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
mv_wait_ready deps BigLakeHiveDatabase "${DB_K8S}" 300
mv_wait_http deps bucket "${BUCKET_URL}" 200
mv_wait_http deps catalog "${CATALOG_URL}" 200
mv_wait_http deps db "${DB_URL}" 200
mv_cli deps gcloud-describe-db -- gcloud biglake hive databases describe "${DB_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" || true

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create BigLakeHiveTable "${TABLE_K8S}" 300
SNAP="$(mv_snapshot create BigLakeHiveTable "${TABLE_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

mv_wait_http create table "${TABLE_URL}" 200
mv_expect_jq create "cloud name matches expected" "${MV_RUN_DIR}/cloud/create-table.json" '.name' "${EXPECTED_TABLE_NAME}"

# Compare every spec field in SNAP to cloud/create-table.json
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-table.json" '.description' "${WANT_DESCRIPTION}"
fi

# Parameters (Note: GCP sets table_type=HIVE by default; controller keeps user parameters)
if jq -e '.spec.parameters' "${SNAP}" >/dev/null 2>&1; then
  for k in $(jq -r '.spec.parameters | keys[]' "${SNAP}"); do
    val="$(jq -r --arg k "$k" '.spec.parameters[$k]' "${SNAP}")"
    mv_expect_jq create "cloud parameters.${k} matches spec" "${MV_RUN_DIR}/cloud/create-table.json" ".parameters[\"${k}\"]" "${val}"
  done
fi

# Partition keys
if jq -e '.spec.partitionKeys' "${SNAP}" >/dev/null 2>&1; then
  mv_expect_jq create "cloud partitionKeys count matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
    '.partitionKeys | length | tostring' "$(jq '.spec.partitionKeys | length' "${SNAP}")"
  pk_len=$(jq '.spec.partitionKeys | length' "${SNAP}")
  for ((i=0; i<pk_len; i++)); do
    pk_name="$(jq -r --argjson i "$i" '.spec.partitionKeys[$i].name' "${SNAP}")"
    pk_type="$(jq -r --argjson i "$i" '.spec.partitionKeys[$i].type' "${SNAP}")"
    pk_comment="$(jq -r --argjson i "$i" '.spec.partitionKeys[$i].comment // empty' "${SNAP}")"
    mv_expect_jq create "cloud partitionKeys[${i}].name matches spec" "${MV_RUN_DIR}/cloud/create-table.json" ".partitionKeys[${i}].name" "${pk_name}"
    mv_expect_jq create "cloud partitionKeys[${i}].type matches spec" "${MV_RUN_DIR}/cloud/create-table.json" ".partitionKeys[${i}].type" "${pk_type}"
    if [[ -n "${pk_comment}" ]]; then
      mv_expect_jq create "cloud partitionKeys[${i}].comment matches spec" "${MV_RUN_DIR}/cloud/create-table.json" ".partitionKeys[${i}].comment" "${pk_comment}"
    fi
  done
fi

# Storage descriptor fields
if jq -e '.spec.storageDescriptor' "${SNAP}" >/dev/null 2>&1; then
  loc_uri="$(jq -r '.spec.storageDescriptor.locationURI // empty' "${SNAP}")"
  if [[ -n "${loc_uri}" ]]; then
    mv_expect_jq create "cloud storageDescriptor.locationUri matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
      '.storageDescriptor.locationUri' "${loc_uri}"
  fi
  input_fmt="$(jq -r '.spec.storageDescriptor.inputFormat // empty' "${SNAP}")"
  if [[ -n "${input_fmt}" ]]; then
    mv_expect_jq create "cloud storageDescriptor.inputFormat matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
      '.storageDescriptor.inputFormat' "${input_fmt}"
  fi
  output_fmt="$(jq -r '.spec.storageDescriptor.outputFormat // empty' "${SNAP}")"
  if [[ -n "${output_fmt}" ]]; then
    mv_expect_jq create "cloud storageDescriptor.outputFormat matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
      '.storageDescriptor.outputFormat' "${output_fmt}"
  fi
  num_buckets="$(jq -r '.spec.storageDescriptor.numBuckets // empty' "${SNAP}")"
  if [[ -n "${num_buckets}" ]]; then
    mv_expect_jq create "cloud storageDescriptor.numBuckets matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
      '.storageDescriptor.numBuckets | tostring' "${num_buckets}"
  fi
  if jq -e '.spec.storageDescriptor.bucketCols' "${SNAP}" >/dev/null 2>&1; then
    bc_len=$(jq '.spec.storageDescriptor.bucketCols | length' "${SNAP}")
    mv_expect_jq create "cloud storageDescriptor.bucketCols length matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
      '.storageDescriptor.bucketCols | length | tostring' "${bc_len}"
    for ((i=0; i<bc_len; i++)); do
      bcol="$(jq -r --argjson i "$i" '.spec.storageDescriptor.bucketCols[$i]' "${SNAP}")"
      mv_expect_jq create "cloud storageDescriptor.bucketCols[${i}] matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
        ".storageDescriptor.bucketCols[${i}]" "${bcol}"
    done
  fi
  if jq -e '.spec.storageDescriptor.sortCols' "${SNAP}" >/dev/null 2>&1; then
    sort_len=$(jq '.spec.storageDescriptor.sortCols | length' "${SNAP}")
    for ((i=0; i<sort_len; i++)); do
      sc_col="$(jq -r --argjson i "$i" '.spec.storageDescriptor.sortCols[$i].col' "${SNAP}")"
      sc_order="$(jq -r --argjson i "$i" '.spec.storageDescriptor.sortCols[$i].order' "${SNAP}")"
      mv_expect_jq create "cloud storageDescriptor.sortCols[${i}].col matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
        ".storageDescriptor.sortCols[${i}].col" "${sc_col}"
      # In proto, order 0 may be omitted by protobuf JSON serialization if default enum/int
      if [[ "${sc_order}" != "0" ]]; then
        mv_expect_jq create "cloud storageDescriptor.sortCols[${i}].order matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
          ".storageDescriptor.sortCols[${i}].order | tostring" "${sc_order}"
      fi
    done
  fi
  if jq -e '.spec.storageDescriptor.parameters' "${SNAP}" >/dev/null 2>&1; then
    for k in $(jq -r '.spec.storageDescriptor.parameters | keys[]' "${SNAP}"); do
      val="$(jq -r --arg k "$k" '.spec.storageDescriptor.parameters[$k]' "${SNAP}")"
      mv_expect_jq create "cloud storageDescriptor.parameters.${k} matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
        ".storageDescriptor.parameters[\"${k}\"]" "${val}"
    done
  fi

  # SerdeInfo
  if jq -e '.spec.storageDescriptor.serdeInfo' "${SNAP}" >/dev/null 2>&1; then
    serde_name="$(jq -r '.spec.storageDescriptor.serdeInfo.name // empty' "${SNAP}")"
    if [[ -n "${serde_name}" ]]; then
      mv_expect_jq create "cloud storageDescriptor.serdeInfo.name matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
        '.storageDescriptor.serdeInfo.name' "${serde_name}"
    fi
    serde_lib="$(jq -r '.spec.storageDescriptor.serdeInfo.serializationLib // empty' "${SNAP}")"
    if [[ -n "${serde_lib}" ]]; then
      mv_expect_jq create "cloud storageDescriptor.serdeInfo.serializationLib matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
        '.storageDescriptor.serdeInfo.serializationLib' "${serde_lib}"
    fi
    serde_desc="$(jq -r '.spec.storageDescriptor.serdeInfo.description // empty' "${SNAP}")"
    if [[ -n "${serde_desc}" ]]; then
      mv_expect_jq create "cloud storageDescriptor.serdeInfo.description matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
        '.storageDescriptor.serdeInfo.description' "${serde_desc}"
    fi
    if jq -e '.spec.storageDescriptor.serdeInfo.parameters' "${SNAP}" >/dev/null 2>&1; then
      for k in $(jq -r '.spec.storageDescriptor.serdeInfo.parameters | keys[]' "${SNAP}"); do
        val="$(jq -r --arg k "$k" '.spec.storageDescriptor.serdeInfo.parameters[$k]' "${SNAP}")"
        mv_expect_jq create "cloud storageDescriptor.serdeInfo.parameters.${k} matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
          ".storageDescriptor.serdeInfo.parameters[\"${k}\"]" "${val}"
      done
    fi
  fi

  # Columns
  if jq -e '.spec.storageDescriptor.columns' "${SNAP}" >/dev/null 2>&1; then
    mv_expect_jq create "cloud storageDescriptor.columns count matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
      '.storageDescriptor.columns | length | tostring' "$(jq '.spec.storageDescriptor.columns | length' "${SNAP}")"
    col_len=$(jq '.spec.storageDescriptor.columns | length' "${SNAP}")
    for ((i=0; i<col_len; i++)); do
      c_name="$(jq -r --argjson i "$i" '.spec.storageDescriptor.columns[$i].name' "${SNAP}")"
      c_type="$(jq -r --argjson i "$i" '.spec.storageDescriptor.columns[$i].type' "${SNAP}")"
      c_comment="$(jq -r --argjson i "$i" '.spec.storageDescriptor.columns[$i].comment // empty' "${SNAP}")"
      mv_expect_jq create "cloud storageDescriptor.columns[${i}].name matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
        ".storageDescriptor.columns[${i}].name" "${c_name}"
      mv_expect_jq create "cloud storageDescriptor.columns[${i}].type matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
        ".storageDescriptor.columns[${i}].type" "${c_type}"
      if [[ -n "${c_comment}" ]]; then
        mv_expect_jq create "cloud storageDescriptor.columns[${i}].comment matches spec" "${MV_RUN_DIR}/cloud/create-table.json" \
          ".storageDescriptor.columns[${i}].comment" "${c_comment}"
      fi
    done
  fi
fi

# Out-of-band CLI describe and list
mv_cli create gcloud-describe-table -- gcloud biglake hive tables describe "${TABLE_ID}" \
  --database="${DB_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-describe-table.txt" "${MV_RUN_DIR}/cloud/create-gcloud-describe-table.json"
mv_expect_jq create "gcloud describe table matches name" "${MV_RUN_DIR}/cloud/create-gcloud-describe-table.json" '.name' "${EXPECTED_TABLE_NAME}"

mv_cli create gcloud-list-tables -- gcloud biglake hive tables list \
  --database="${DB_ID}" --catalog="${CATALOG_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-list-tables.txt" "${MV_RUN_DIR}/cloud/create-gcloud-list-tables.json"
FOUND_TABLE="$(jq -r --arg NAME "${EXPECTED_TABLE_NAME}" '[.[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-gcloud-list-tables.json")"
mv_expect_eq create "gcloud list tables includes the table" "${EXPECTED_TABLE_NAME}" "${FOUND_TABLE}" cloud/create-gcloud-list-tables.json

# Functional check: table IAM policy inspection
if [[ "${FUNCTIONAL}" == "1" ]]; then
  if mv_cli create functional-get-iam-policy -- gcloud biglake hive tables get-iam-policy "${TABLE_ID}" \
      --catalog="${CATALOG_ID}" --database="${DB_ID}" --project="${GCP_PROJECT_ID}" --format=json; then
    mv_record create "functional check: get table IAM policy" PASS cloud/create-functional-get-iam-policy.txt ""
  else
    mv_record create "functional check: get table IAM policy" FAIL cloud/create-functional-get-iam-policy.txt "failed to get table IAM policy"
  fi
  mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-functional-get-iam-policy.txt" "${MV_RUN_DIR}/cloud/create-functional-get-iam-policy.json"
  mv_expect_jq create "functional check: table IAM policy returns etag" "${MV_RUN_DIR}/cloud/create-functional-get-iam-policy.json" \
    '(.etag != null and .etag != "") | tostring' "true"
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
# Give the apiserver a moment to bump generation before polling for Ready.
sleep 2
mv_wait_ready update BigLakeHiveTable "${TABLE_K8S}" 300
SNAP="$(mv_snapshot update BigLakeHiveTable "${TABLE_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update table "${TABLE_URL}" 200

# Compare changed fields to update spec
WANT_DESCRIPTION_UPDATE="$(jq -r '.spec.description // empty' "${SNAP}")"
if [[ -n "${WANT_DESCRIPTION_UPDATE}" ]]; then
  mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-table.json" '.description' "${WANT_DESCRIPTION_UPDATE}"
fi

if jq -e '.spec.parameters' "${SNAP}" >/dev/null 2>&1; then
  for k in $(jq -r '.spec.parameters | keys[]' "${SNAP}"); do
    val="$(jq -r --arg k "$k" '.spec.parameters[$k]' "${SNAP}")"
    mv_expect_jq update "cloud parameters.${k} updated to new spec" "${MV_RUN_DIR}/cloud/update-table.json" ".parameters[\"${k}\"]" "${val}"
  done
fi

if jq -e '.spec.partitionKeys' "${SNAP}" >/dev/null 2>&1; then
  pk_len=$(jq '.spec.partitionKeys | length' "${SNAP}")
  for ((i=0; i<pk_len; i++)); do
    pk_comment="$(jq -r --argjson i "$i" '.spec.partitionKeys[$i].comment // empty' "${SNAP}")"
    if [[ -n "${pk_comment}" ]]; then
      mv_expect_jq update "cloud partitionKeys[${i}].comment updated to new spec" "${MV_RUN_DIR}/cloud/update-table.json" \
        ".partitionKeys[${i}].comment" "${pk_comment}"
    fi
  done
fi

if jq -e '.spec.storageDescriptor' "${SNAP}" >/dev/null 2>&1; then
  num_buckets="$(jq -r '.spec.storageDescriptor.numBuckets // empty' "${SNAP}")"
  if [[ -n "${num_buckets}" ]]; then
    mv_expect_jq update "cloud storageDescriptor.numBuckets updated to new spec" "${MV_RUN_DIR}/cloud/update-table.json" \
      '.storageDescriptor.numBuckets | tostring' "${num_buckets}"
  fi
  if jq -e '.spec.storageDescriptor.parameters' "${SNAP}" >/dev/null 2>&1; then
    for k in $(jq -r '.spec.storageDescriptor.parameters | keys[]' "${SNAP}"); do
      val="$(jq -r --arg k "$k" '.spec.storageDescriptor.parameters[$k]' "${SNAP}")"
      mv_expect_jq update "cloud storageDescriptor.parameters.${k} updated to new spec" "${MV_RUN_DIR}/cloud/update-table.json" \
        ".storageDescriptor.parameters[\"${k}\"]" "${val}"
    done
  fi
  if jq -e '.spec.storageDescriptor.serdeInfo.description' "${SNAP}" >/dev/null 2>&1; then
    serde_desc="$(jq -r '.spec.storageDescriptor.serdeInfo.description // empty' "${SNAP}")"
    if [[ -n "${serde_desc}" ]]; then
      mv_expect_jq update "cloud storageDescriptor.serdeInfo.description updated to new spec" "${MV_RUN_DIR}/cloud/update-table.json" \
        '.storageDescriptor.serdeInfo.description' "${serde_desc}"
    fi
  fi
  if jq -e '.spec.storageDescriptor.columns' "${SNAP}" >/dev/null 2>&1; then
    mv_expect_jq update "cloud storageDescriptor.columns count updated to new spec" "${MV_RUN_DIR}/cloud/update-table.json" \
      '.storageDescriptor.columns | length | tostring' "$(jq '.spec.storageDescriptor.columns | length' "${SNAP}")"
    col_len=$(jq '.spec.storageDescriptor.columns | length' "${SNAP}")
    for ((i=0; i<col_len; i++)); do
      c_name="$(jq -r --argjson i "$i" '.spec.storageDescriptor.columns[$i].name' "${SNAP}")"
      c_type="$(jq -r --argjson i "$i" '.spec.storageDescriptor.columns[$i].type' "${SNAP}")"
      c_comment="$(jq -r --argjson i "$i" '.spec.storageDescriptor.columns[$i].comment // empty' "${SNAP}")"
      mv_expect_jq update "cloud storageDescriptor.columns[${i}].name matches update spec" "${MV_RUN_DIR}/cloud/update-table.json" \
        ".storageDescriptor.columns[${i}].name" "${c_name}"
      mv_expect_jq update "cloud storageDescriptor.columns[${i}].type matches update spec" "${MV_RUN_DIR}/cloud/update-table.json" \
        ".storageDescriptor.columns[${i}].type" "${c_type}"
      if [[ -n "${c_comment}" ]]; then
        mv_expect_jq update "cloud storageDescriptor.columns[${i}].comment matches update spec" "${MV_RUN_DIR}/cloud/update-table.json" \
          ".storageDescriptor.columns[${i}].comment" "${c_comment}"
      fi
    done
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady BigLakeHiveTable "${TABLE_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete BigLakeHiveTable "${TABLE_K8S}" 300
mv_wait_http delete table-gone "${TABLE_URL}" 404

# Cloud Audit Logs check: BigLake Hive Table operations logged under HiveMetastoreService
# Note: Hive tables are also mirrored by BigQuery (logging under bigquery.googleapis.com).
# We filter to serviceName="biglake.googleapis.com" to verify the exact HiveMetastoreService admin calls.
expect_biglake_table_audit() {
  local phase="$1" label="$2" substr="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.serviceName=\"biglake.googleapis.com\" AND protoPayload.resourceName:\"${substr}\" AND timestamp>=\"${MV_STARTED_UTC}\""
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    got="$(mv_audit_calls "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
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

expect_biglake_table_audit delete audit-table "tables/${TABLE_ID}" \
  "CreateHiveTable,UpdateHiveTable,DeleteHiveTable"

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown BigLakeHiveDatabase "${DB_K8S}" 300
mv_wait_http teardown db-gone "${DB_URL}" 404

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
