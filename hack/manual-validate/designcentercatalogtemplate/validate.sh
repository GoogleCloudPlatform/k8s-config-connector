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

# Human-style validation of the DesignCenterCatalogTemplate controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/designcentercatalogtemplate/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the designcentercatalogtemplate-maximal test fixture)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/designcenter/v1alpha1/designcentercatalogtemplate/designcentercatalogtemplate-maximal}"

mv_init DesignCenterCatalogTemplate
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Names derived from the fixture so out-of-band checks never ask KCC.
SPACE_K8S="dc-space-${MV_UNIQUE_ID}"
SPACE_ID="${SPACE_K8S}"
CATALOG_K8S="dc-catalog-${MV_UNIQUE_ID}"
CATALOG_ID="${CATALOG_K8S}"

TEMPLATE_NAME_RAW="$(grep -E '^[[:space:]]*name:[[:space:]]' "${FIXTURE_DIR}/create.yaml" | head -1 | awk '{print $2}')"
TEMPLATE_K8S="$(sed -e "s|\${uniqueId}|${MV_UNIQUE_ID}|g" <<< "${TEMPLATE_NAME_RAW}")"
TEMPLATE_ID="${TEMPLATE_K8S}"
LOCATION="us-central1"

DC_API="https://designcenter.googleapis.com/v1"
SPACE_URL="${DC_API}/projects/${GCP_PROJECT_ID}/locations/${LOCATION}/spaces/${SPACE_ID}"
CATALOG_URL="${SPACE_URL}/catalogs/${CATALOG_ID}"
TEMPLATE_URL="${CATALOG_URL}/templates/${TEMPLATE_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get designcentercatalogtemplate "${TEMPLATE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterCatalogTemplate "${TEMPLATE_K8S}" 300
    mv_kubectl get designcentercatalog "${CATALOG_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterCatalog "${CATALOG_K8S}" 300
    mv_kubectl get designcenterspace "${SPACE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DesignCenterSpace "${SPACE_K8S}" 300
  fi

  local code
  code="$(mv_rest cleanup leakcheck-template GET "${TEMPLATE_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP template (template gone)" PASS cloud/cleanup-leakcheck-template.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP template (template gone)" FAIL cloud/cleanup-leakcheck-template.json \
      "could not verify (HTTP ${code}); check manually: gcloud design-center spaces catalogs templates describe ${TEMPLATE_ID} --catalog=${CATALOG_ID} --space=${SPACE_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP template (template gone)" FAIL cloud/cleanup-leakcheck-template.json \
      "HTTP ${code}; clean up manually: gcloud design-center spaces catalogs templates delete ${TEMPLATE_ID} --catalog=${CATALOG_ID} --space=${SPACE_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID} --quiet"
  fi

  code="$(mv_rest cleanup leakcheck-catalog GET "${CATALOG_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP catalog (catalog gone)" PASS cloud/cleanup-leakcheck-catalog.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP catalog (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
      "could not verify (HTTP ${code}); check manually: gcloud design-center spaces catalogs describe ${CATALOG_ID} --space=${SPACE_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP catalog (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
      "HTTP ${code}; clean up manually: gcloud design-center spaces catalogs delete ${CATALOG_ID} --space=${SPACE_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID} --quiet"
  fi

  code="$(mv_rest cleanup leakcheck-space GET "${SPACE_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP space (space gone)" PASS cloud/cleanup-leakcheck-space.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP space (space gone)" FAIL cloud/cleanup-leakcheck-space.json \
      "could not verify (HTTP ${code}); check manually: gcloud design-center spaces describe ${SPACE_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP space (space gone)" FAIL cloud/cleanup-leakcheck-space.json \
      "HTTP ${code}; clean up manually: gcloud design-center spaces delete ${SPACE_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID} --quiet"
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
mv_enable_apis setup designcenter.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
mv_apply deps "${DEPS}"
mv_wait_ready deps DesignCenterSpace "${SPACE_K8S}" 300
mv_wait_ready deps DesignCenterCatalog "${CATALOG_K8S}" 300
mv_wait_http deps space "${SPACE_URL}" 200
mv_wait_http deps catalog "${CATALOG_URL}" 200
mv_cli deps gcloud-describe-space -- gcloud design-center spaces describe "${SPACE_ID}" \
  --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
mv_cli deps gcloud-describe-catalog -- gcloud design-center spaces catalogs describe "${CATALOG_ID}" \
  --space="${SPACE_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create DesignCenterCatalogTemplate "${TEMPLATE_K8S}" 300
SNAP="$(mv_snapshot create DesignCenterCatalogTemplate "${TEMPLATE_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"
WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_CATEGORY="$(jq -r '.spec.templateCategory // empty' "${SNAP}")"
WANT_TYPE="$(jq -r '.spec.type // empty' "${SNAP}")"

mv_wait_http create template "${TEMPLATE_URL}" 200

# Spec field checks against cloud REST response
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-template.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-template.json" '.description' "${WANT_DESCRIPTION}"
fi
if [[ -n "${WANT_CATEGORY}" ]]; then
  mv_expect_jq create "cloud templateCategory matches spec" "${MV_RUN_DIR}/cloud/create-template.json" \
    'if .templateCategory == 1 or .templateCategory == "COMPONENT_TEMPLATE" then "COMPONENT_TEMPLATE" else (.templateCategory | tostring) end' \
    "${WANT_CATEGORY}"
fi
if [[ -n "${WANT_TYPE}" ]]; then
  mv_expect_jq create "cloud type matches spec" "${MV_RUN_DIR}/cloud/create-template.json" \
    'if .type == 1 or .type == "SERVICE" then "SERVICE" elif .type == 2 or .type == "WORKLOAD" then "WORKLOAD" else (.type | tostring) end' \
    "${WANT_TYPE}"
fi
mv_expect_jq create "cloud resource name matches expected" "${MV_RUN_DIR}/cloud/create-template.json" \
  '.name' "projects/${GCP_PROJECT_ID}/locations/${LOCATION}/spaces/${SPACE_ID}/catalogs/${CATALOG_ID}/templates/${TEMPLATE_ID}"

# gcloud describe check
mv_cli create gcloud-describe-template -- gcloud design-center spaces catalogs templates describe "${TEMPLATE_ID}" \
  --catalog="${CATALOG_ID}" --space="${SPACE_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
if grep -q "templates/${TEMPLATE_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-describe-template.txt"; then
  mv_record create "gcloud describe confirms template exists" PASS cloud/create-gcloud-describe-template.txt ""
else
  mv_record create "gcloud describe confirms template exists" FAIL cloud/create-gcloud-describe-template.txt "template id not found in describe output"
fi

# Functional check: list templates in the parent catalog
mv_cli create gcloud-list-templates -- gcloud design-center spaces catalogs templates list \
  --catalog="${CATALOG_ID}" --space="${SPACE_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
if grep -q "templates/${TEMPLATE_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-list-templates.txt"; then
  mv_record create "functional: template listed in parent catalog" PASS cloud/create-gcloud-list-templates.txt ""
else
  mv_record create "functional: template listed in parent catalog" FAIL cloud/create-gcloud-list-templates.txt "template id not listed"
fi

# Functional check: query revisions sub-resource
if mv_cli create gcloud-list-revisions -- gcloud design-center spaces catalogs templates revisions list \
  --template="${TEMPLATE_ID}" --catalog="${CATALOG_ID}" --space="${SPACE_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json; then
  mv_record create "functional: template revisions query operates on template" PASS cloud/create-gcloud-list-revisions.txt ""
else
  mv_record create "functional: template revisions query operates on template" FAIL cloud/create-gcloud-list-revisions.txt "gcloud revisions list failed"
fi

# --- Phase: update -----------------------------------------------------------
if [[ -f "${FIXTURE_DIR}/update.yaml" ]]; then
  UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
  mv_apply update "${UPDATE}"
  # Give the apiserver a moment to bump generation before polling for Ready.
  sleep 2
  mv_wait_ready update DesignCenterCatalogTemplate "${TEMPLATE_K8S}" 300
  SNAP="$(mv_snapshot update DesignCenterCatalogTemplate "${TEMPLATE_K8S}")"
  mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"
  WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
  WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
  WANT_CATEGORY="$(jq -r '.spec.templateCategory // empty' "${SNAP}")"
  WANT_TYPE="$(jq -r '.spec.type // empty' "${SNAP}")"

  mv_wait_http update template "${TEMPLATE_URL}" 200

  if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
    mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-template.json" '.displayName' "${WANT_DISPLAY_NAME}"
  fi
  if [[ -n "${WANT_DESCRIPTION}" ]]; then
    mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-template.json" '.description' "${WANT_DESCRIPTION}"
  fi
  if [[ -n "${WANT_CATEGORY}" ]]; then
    mv_expect_jq update "cloud templateCategory remains intact" "${MV_RUN_DIR}/cloud/update-template.json" \
      'if .templateCategory == 1 or .templateCategory == "COMPONENT_TEMPLATE" then "COMPONENT_TEMPLATE" else (.templateCategory | tostring) end' \
      "${WANT_CATEGORY}"
  fi
  if [[ -n "${WANT_TYPE}" ]]; then
    mv_expect_jq update "cloud type updated to new spec" "${MV_RUN_DIR}/cloud/update-template.json" \
      'if .type == 1 or .type == "SERVICE" then "SERVICE" elif .type == 2 or .type == "WORKLOAD" then "WORKLOAD" else (.type | tostring) end' \
      "${WANT_TYPE}"
  fi

  CREATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/create-template.json")"
  UPDATE_T="$(jq -r '.createTime // empty' "${MV_RUN_DIR}/cloud/update-template.json")"
  if [[ -n "${CREATE_T}" && -n "${UPDATE_T}" ]]; then
    mv_expect_eq update "createTime unchanged (in-place update, not recreated)" "${CREATE_T}" "${UPDATE_T}" cloud/update-template.json
  fi

  # Functional check: updated template still listed in parent catalog
  mv_cli update gcloud-list-templates -- gcloud design-center spaces catalogs templates list \
    --catalog="${CATALOG_ID}" --space="${SPACE_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
  if grep -q "templates/${TEMPLATE_ID}" "${MV_RUN_DIR}/cloud/update-gcloud-list-templates.txt"; then
    mv_record update "functional: updated template listed in parent catalog" PASS cloud/update-gcloud-list-templates.txt ""
  else
    mv_record update "functional: updated template listed in parent catalog" FAIL cloud/update-gcloud-list-templates.txt "template id not listed"
  fi
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DesignCenterCatalogTemplate "${TEMPLATE_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DesignCenterCatalogTemplate "${TEMPLATE_K8S}" 300
mv_wait_http delete template-gone "${TEMPLATE_URL}" 404

# Audit log check: verify admin-activity writes on template via parent catalog
EXPECTED_AUDIT="CreateCatalogTemplate,UpdateCatalogTemplate,DeleteCatalogTemplate,DeleteCatalogTemplate"
if [[ ! -f "${FIXTURE_DIR}/update.yaml" ]]; then
  EXPECTED_AUDIT="CreateCatalogTemplate,DeleteCatalogTemplate,DeleteCatalogTemplate"
fi
mv_expect_audit delete audit-template "catalogs/${CATALOG_ID}" "${EXPECTED_AUDIT}"

# Functional check: template no longer listed in parent catalog
mv_cli delete gcloud-list-templates -- gcloud design-center spaces catalogs templates list \
  --catalog="${CATALOG_ID}" --space="${SPACE_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
if grep -q "templates/${TEMPLATE_ID}" "${MV_RUN_DIR}/cloud/delete-gcloud-list-templates.txt"; then
  mv_record delete "functional: template no longer listed in parent catalog" FAIL cloud/delete-gcloud-list-templates.txt "template still listed after deletion"
else
  mv_record delete "functional: template no longer listed in parent catalog" PASS cloud/delete-gcloud-list-templates.txt ""
fi

# --- Phase: dependency teardown ----------------------------------------------
mv_delete teardown DesignCenterCatalog "${CATALOG_K8S}" 300
mv_wait_http teardown catalog-gone "${CATALOG_URL}" 404
mv_delete teardown DesignCenterSpace "${SPACE_K8S}" 300
mv_wait_http teardown space-gone "${SPACE_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
