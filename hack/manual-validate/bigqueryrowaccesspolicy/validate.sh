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

# Human-style validation of the BigQueryRowAccessPolicy controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/bigqueryrowaccesspolicy/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the bigqueryrowaccesspolicy-maximal test fixture)
#   FUNCTIONAL=0     skip the data-level check (insert rows + query through the policy)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/bigquery/v1alpha1/bigqueryrowaccesspolicy/bigqueryrowaccesspolicy-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init BigQueryRowAccessPolicy
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Names derived from the fixture (resourceIDs) so out-of-band checks never ask KCC.
DS_K8S="bigquerydataset${MV_UNIQUE_ID}"
TBL_K8S="bigquerytable${MV_UNIQUE_ID}"
RAP_K8S="bigqueryrowaccesspolicy${MV_UNIQUE_ID}"
DS_ID="bigquerydataset_${MV_UNIQUE_ID}"
TBL_ID="bigquerytable_${MV_UNIQUE_ID}"
RAP_ID="bqrap_${MV_UNIQUE_ID}"
LOCATION="us-central1"

BQ="https://bigquery.googleapis.com/bigquery/v2/projects/${GCP_PROJECT_ID}"
DS_URL="${BQ}/datasets/${DS_ID}"
TBL_URL="${DS_URL}/tables/${TBL_ID}"
RAP_URL="${TBL_URL}/rowAccessPolicies/${RAP_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get bigqueryrowaccesspolicy "${RAP_K8S}" >/dev/null 2>&1 && mv_delete cleanup BigQueryRowAccessPolicy "${RAP_K8S}" 300
    mv_kubectl get bigquerytable "${TBL_K8S}" >/dev/null 2>&1 && mv_delete cleanup BigQueryTable "${TBL_K8S}" 300
    mv_kubectl get bigquerydataset "${DS_K8S}" >/dev/null 2>&1 && mv_delete cleanup BigQueryDataset "${DS_K8S}" 300
  fi
  local code
  code="$(mv_rest cleanup leakcheck-dataset GET "${DS_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (dataset gone)" PASS cloud/cleanup-leakcheck-dataset.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (dataset gone)" FAIL cloud/cleanup-leakcheck-dataset.json \
      "could not verify (HTTP ${code}); check manually: bq show ${GCP_PROJECT_ID}:${DS_ID}"
  else
    mv_record cleanup "no leaked GCP resources (dataset gone)" FAIL cloud/cleanup-leakcheck-dataset.json \
      "HTTP ${code}; clean up manually: bq rm -r -f -d ${GCP_PROJECT_ID}:${DS_ID}"
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

# Functional check: query the table as the operator. allAuthenticatedUsers is a
# grantee, so the operator sees exactly the rows the filter predicate allows.
query_regions() {
  local phase="$1" label="$2"
  mv_cli "${phase}" "${label}" -- bq --project_id="${GCP_PROJECT_ID}" --location="${LOCATION}" query \
    --use_legacy_sql=false --format=json \
    "SELECT region FROM \`${GCP_PROJECT_ID}.${DS_ID}.${TBL_ID}\` ORDER BY region" || true
  # bq --format=json prints the result set as a single-line JSON array.
  local rows
  rows="$(grep -m1 '^\[' "${MV_RUN_DIR}/cloud/${phase}-${label}.txt" || true)"
  echo "${rows:-[]}" | jq -r '[.[].region] | join(",")' 2>/dev/null || echo "<unparseable>"
}

grantee_members() {
  # filteredDataViewer members on the policy, via raw REST getIamPolicy.
  local phase="$1" code
  code="$(mv_rest "${phase}" rap-getiampolicy POST "${RAP_URL}:getIamPolicy" '{}')"
  mv_expect_http "${phase}" "REST getIamPolicy on row access policy" 200 "${code}" "cloud/${phase}-rap-getiampolicy.json"
  jq -r '[.bindings[]? | select(.role=="roles/bigquery.filteredDataViewer") | .members[]] | sort | join(",")' \
    "${MV_RUN_DIR}/cloud/${phase}-rap-getiampolicy.json"
}

# ---------------------------------------------------------------------------
mv_up

# --- Phase: dependencies -----------------------------------------------------
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
mv_apply deps "${DEPS}"
mv_wait_ready deps BigQueryDataset "${DS_K8S}" 300
mv_wait_ready deps BigQueryTable "${TBL_K8S}" 300
mv_wait_http deps dataset "${DS_URL}" 200
mv_wait_http deps table "${TBL_URL}" 200
mv_cli deps bq-show-table -- bq --project_id="${GCP_PROJECT_ID}" show --format=prettyjson "${GCP_PROJECT_ID}:${DS_ID}.${TBL_ID}" || true

if [[ "${FUNCTIONAL}" == "1" ]]; then
  # Seed data BEFORE the policy exists so both rows are visible as a baseline.
  mv_cli deps seed-rows -- bq --project_id="${GCP_PROJECT_ID}" --location="${LOCATION}" query --use_legacy_sql=false \
    "INSERT INTO \`${GCP_PROJECT_ID}.${DS_ID}.${TBL_ID}\` (region, amount) VALUES ('US', 1), ('EU', 2)" &&
    mv_record deps "seeded rows US,EU via bq DML" PASS cloud/deps-seed-rows.txt "" ||
    mv_record deps "seeded rows US,EU via bq DML" FAIL cloud/deps-seed-rows.txt "see file"
  mv_expect_eq deps "baseline query (no policy) sees all rows" "EU,US" "$(query_regions deps query-baseline)" cloud/deps-query-baseline.txt
fi

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create BigQueryRowAccessPolicy "${RAP_K8S}" 300
SNAP="$(mv_snapshot create BigQueryRowAccessPolicy "${RAP_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"
WANT_FILTER="$(jq -r '.spec.filterPredicate' "${SNAP}")"
WANT_GRANTEES="$(jq -r '[.spec.grantees[]?] | sort | join(",")' "${SNAP}")"

mv_wait_http create rap "${RAP_URL}" 200
mv_expect_jq create "cloud filterPredicate matches spec" "${MV_RUN_DIR}/cloud/create-rap.json" '.filterPredicate' "${WANT_FILTER}"
mv_expect_jq create "cloud policy id matches resourceID" "${MV_RUN_DIR}/cloud/create-rap.json" '.rowAccessPolicyReference.policyId' "${RAP_ID}"
mv_expect_eq create "cloud grantees (filteredDataViewer) match spec" "${WANT_GRANTEES}" "$(grantee_members create)" cloud/create-rap-getiampolicy.json
mv_cli create bq-ls-policies -- bq --project_id="${GCP_PROJECT_ID}" ls --row_access_policies --format=prettyjson "${GCP_PROJECT_ID}:${DS_ID}.${TBL_ID}" || true
if grep -q "\"${RAP_ID}\"" "${MV_RUN_DIR}/cloud/create-bq-ls-policies.txt"; then
  mv_record create "bq ls --row_access_policies lists the policy" PASS cloud/create-bq-ls-policies.txt ""
else
  mv_record create "bq ls --row_access_policies lists the policy" FAIL cloud/create-bq-ls-policies.txt "policy id not found"
fi
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_expect_eq create "query through policy returns only US rows" "US" "$(query_regions create query-filtered)" cloud/create-query-filtered.txt
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
mv_apply update "${UPDATE}"
# Give the apiserver a moment to bump generation before polling for Ready.
sleep 2
mv_wait_ready update BigQueryRowAccessPolicy "${RAP_K8S}" 300
SNAP="$(mv_snapshot update BigQueryRowAccessPolicy "${RAP_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"
WANT_FILTER="$(jq -r '.spec.filterPredicate' "${SNAP}")"
WANT_GRANTEES="$(jq -r '[.spec.grantees[]?] | sort | join(",")' "${SNAP}")"
mv_wait_http update rap "${RAP_URL}" 200
mv_expect_jq update "cloud filterPredicate updated to new spec" "${MV_RUN_DIR}/cloud/update-rap.json" '.filterPredicate' "${WANT_FILTER}"
mv_expect_eq update "cloud grantees still match spec" "${WANT_GRANTEES}" "$(grantee_members update)" cloud/update-rap-getiampolicy.json
CREATE_T="$(jq -r '.creationTime // empty' "${MV_RUN_DIR}/cloud/create-rap.json")"
UPDATE_T="$(jq -r '.creationTime // empty' "${MV_RUN_DIR}/cloud/update-rap.json")"
# BigQuery resets creationTime on rowAccessPolicies.update (verified via audit logs on
# 2026-11-20), so it can't prove in-place update. The audit-log check after delete does.
mv_record update "creationTime before/after update (API resets it on update)" INFO cloud/update-rap.json "${CREATE_T} -> ${UPDATE_T}"
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_expect_eq update "query through updated policy returns only EU rows" "EU" "$(query_regions update query-filtered)" cloud/update-query-filtered.txt
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady BigQueryRowAccessPolicy "${RAP_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete BigQueryRowAccessPolicy "${RAP_K8S}" 300
mv_wait_http delete rap-gone "${RAP_URL}" 404
# What GCP actually received for the policy over its whole life: exactly one create,
# one in-place update, one delete. No recreate, no spurious writes while steady.
mv_expect_audit delete audit-rap "rowAccessPolicies/${RAP_ID}" \
  "CreateRowAccessPolicy,UpdateRowAccessPolicy,DeleteRowAccessPolicy"
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_expect_eq delete "query after policy deletion sees all rows again" "EU,US" "$(query_regions delete query-unfiltered)" cloud/delete-query-unfiltered.txt
fi

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown BigQueryTable "${TBL_K8S}" 300
mv_wait_http teardown table-gone "${TBL_URL}" 404
mv_delete teardown BigQueryDataset "${DS_K8S}" 300
mv_wait_http teardown dataset-gone "${DS_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
