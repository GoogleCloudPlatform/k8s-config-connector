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
#   FUNCTIONAL=0     skip the functional catalog listing check
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/designcenter/v1alpha1/designcentercatalogtemplate/designcentercatalogtemplate-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DesignCenterCatalogTemplate
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

DEPS_SRC="${FIXTURE_DIR}/dependencies.yaml"
CREATE_SRC="${FIXTURE_DIR}/create.yaml"
UPDATE_SRC="${FIXTURE_DIR}/update.yaml"

# Derive names from the manifests
TEMPLATE_K8S="$(sed -e "s|\${uniqueId}|${MV_UNIQUE_ID}|g" "${CREATE_SRC}" | grep -E '^\s*name:' | head -1 | awk '{print $2}')"
TEMPLATE_ID="$(sed -e "s|\${uniqueId}|${MV_UNIQUE_ID}|g" "${CREATE_SRC}" | { grep -E '^\s*resourceID:' || true; } | head -1 | awk '{print $2}')"
if [[ -z "${TEMPLATE_ID}" ]]; then
  TEMPLATE_ID="${TEMPLATE_K8S}"
fi

SPACE_K8S="dc-space-${MV_UNIQUE_ID}"
SPACE_ID="${SPACE_K8S}"
CATALOG_K8S="dc-catalog-${MV_UNIQUE_ID}"
CATALOG_ID="${CATALOG_K8S}"
LOCATION="us-central1"

DC_API="https://designcenter.googleapis.com/v1"
SPACE_URL="${DC_API}/projects/${GCP_PROJECT_ID}/locations/${LOCATION}/spaces/${SPACE_ID}"
CATALOG_URL="${SPACE_URL}/catalogs/${CATALOG_ID}"
TEMPLATES_URL="${CATALOG_URL}/templates"
TEMPLATE_URL="${TEMPLATES_URL}/${TEMPLATE_ID}"
TEMPLATE_GCP_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/spaces/${SPACE_ID}/catalogs/${CATALOG_ID}/templates/${TEMPLATE_ID}"

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
    mv_record cleanup "no leaked template (template gone)" PASS cloud/cleanup-leakcheck-template.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked template (template gone)" FAIL cloud/cleanup-leakcheck-template.json \
      "could not verify (HTTP ${code}); check manually: curl ${TEMPLATE_URL}"
  else
    mv_record cleanup "no leaked template (template gone)" FAIL cloud/cleanup-leakcheck-template.json \
      "HTTP ${code}; clean up manually: curl -X DELETE ${TEMPLATE_URL}"
  fi
  code="$(mv_rest cleanup leakcheck-catalog GET "${CATALOG_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked catalog (catalog gone)" PASS cloud/cleanup-leakcheck-catalog.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked catalog (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
      "could not verify (HTTP ${code}); check manually: curl ${CATALOG_URL}"
  else
    mv_record cleanup "no leaked catalog (catalog gone)" FAIL cloud/cleanup-leakcheck-catalog.json \
      "HTTP ${code}; clean up manually: curl -X DELETE ${CATALOG_URL}"
  fi
  code="$(mv_rest cleanup leakcheck-space GET "${SPACE_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked space (space gone)" PASS cloud/cleanup-leakcheck-space.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked space (space gone)" FAIL cloud/cleanup-leakcheck-space.json \
      "could not verify (HTTP ${code}); check manually: gcloud design-center spaces describe ${SPACE_ID} --location=${LOCATION}"
  else
    mv_record cleanup "no leaked space (space gone)" FAIL cloud/cleanup-leakcheck-space.json \
      "HTTP ${code}; clean up manually: gcloud design-center spaces delete ${SPACE_ID} --location=${LOCATION}"
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

# expect_audit_template polls Cloud Audit Logs for write calls affecting the template.
# CreateCatalogTemplate logs against the parent catalog with request.catalogTemplateId.
# UpdateCatalogTemplate and DeleteCatalogTemplate log against the template resourceName.
# For LRO methods like DeleteCatalogTemplate, filter out the operation completion log (operation.last: true)
# to ensure exactly one log entry per action.
expect_audit_template() {
  local phase="$1" label="$2" want="CreateCatalogTemplate,UpdateCatalogTemplate,DeleteCatalogTemplate" timeout=240
  local out="cloud/${phase}-${label}.json"
  local filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.serviceName=\"designcenter.googleapis.com\" AND (protoPayload.resourceName:\"templates/${TEMPLATE_ID}\" OR protoPayload.request.catalogTemplateId=\"${TEMPLATE_ID}\") AND timestamp>=\"${MV_STARTED_UTC}\""
  local start got
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    got="$(jq -r '[.[] | select(.operation.last != true) | .protoPayload.methodName | split(".") | last] | join(",")' "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
    [[ "${got}" == "${want}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 15
  done
  echo "filter: ${filter}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.filter.txt"
  mv_expect_eq "${phase}" "audit log write calls on ${TEMPLATE_ID}" "${want}" "${got}" "${out}"
}

# ---------------------------------------------------------------------------
# Every API the controller, the dependencies and the out-of-band checks call.
mv_enable_apis setup designcenter.googleapis.com apphub.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
DEPS="$(mv_render "${DEPS_SRC}" dependencies.yaml)"
mv_apply deps "${DEPS}"
mv_wait_ready deps DesignCenterSpace "${SPACE_K8S}" 300
mv_wait_ready deps DesignCenterCatalog "${CATALOG_K8S}" 300
mv_wait_http deps space "${SPACE_URL}" 200
mv_wait_http deps catalog "${CATALOG_URL}" 200
mv_cli deps gcloud-describe-space -- gcloud design-center spaces describe "${SPACE_ID}" \
  --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
if grep -q "\"projects/${GCP_PROJECT_ID}/locations/${LOCATION}/spaces/${SPACE_ID}\"" "${MV_RUN_DIR}/cloud/deps-gcloud-describe-space.txt"; then
  mv_record deps "gcloud space describe matches resourceID" PASS cloud/deps-gcloud-describe-space.txt ""
else
  mv_record deps "gcloud space describe matches resourceID" FAIL cloud/deps-gcloud-describe-space.txt "space not found in gcloud output"
fi

# --- Phase: create -----------------------------------------------------------
CREATE="$(mv_render "${CREATE_SRC}" create.yaml)"
mv_apply create "${CREATE}"
mv_wait_ready create DesignCenterCatalogTemplate "${TEMPLATE_K8S}" 300
SNAP="$(mv_snapshot create DesignCenterCatalogTemplate "${TEMPLATE_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"
WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_DESCRIPTION="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_CATEGORY="$(jq -r '.spec.templateCategory // empty' "${SNAP}")"
WANT_TYPE="$(jq -r '.spec.type // empty' "${SNAP}")"

mv_wait_http create template "${TEMPLATE_URL}" 200
mv_expect_jq create "cloud template name matches expected" "${MV_RUN_DIR}/cloud/create-template.json" '.name' "${TEMPLATE_GCP_NAME}"
if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-template.json" '.displayName' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_DESCRIPTION}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-template.json" '.description' "${WANT_DESCRIPTION}"
fi
if [[ -n "${WANT_CATEGORY}" ]]; then
  mv_expect_jq create "cloud templateCategory matches spec" "${MV_RUN_DIR}/cloud/create-template.json" '.templateCategory' "${WANT_CATEGORY}"
fi
if [[ -n "${WANT_TYPE}" ]]; then
  mv_expect_jq create "cloud type matches spec" "${MV_RUN_DIR}/cloud/create-template.json" '.type' "${WANT_TYPE}"
fi

if [[ "${FUNCTIONAL}" == "1" ]]; then
  code="$(mv_rest create list-templates GET "${TEMPLATES_URL}")"
  mv_expect_http create "REST list templates returns 200" 200 "${code}" "cloud/create-list-templates.json"
  if jq -e --arg name "${TEMPLATE_GCP_NAME}" '.catalogTemplates[]? | select(.name == $name)' "${MV_RUN_DIR}/cloud/create-list-templates.json" >/dev/null; then
    mv_record create "template is listed in catalog templates" PASS cloud/create-list-templates.json ""
  else
    mv_record create "template is listed in catalog templates" FAIL cloud/create-list-templates.json "template not found in list"
  fi
fi

# --- Phase: update -----------------------------------------------------------
UPDATE="$(mv_render "${UPDATE_SRC}" update.yaml)"
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
if [[ -n "${WANT_TYPE}" ]]; then
  mv_expect_jq update "cloud type updated to new spec" "${MV_RUN_DIR}/cloud/update-template.json" '.type' "${WANT_TYPE}"
fi
if [[ -n "${WANT_CATEGORY}" ]]; then
  mv_expect_jq update "cloud templateCategory unchanged" "${MV_RUN_DIR}/cloud/update-template.json" '.templateCategory' "${WANT_CATEGORY}"
fi

CREATE_UUID="$(jq -r '.uuid // empty' "${MV_RUN_DIR}/cloud/create-template.json")"
UPDATE_UUID="$(jq -r '.uuid // empty' "${MV_RUN_DIR}/cloud/update-template.json")"
mv_expect_eq update "template uuid is preserved across update (not recreated)" "${CREATE_UUID}" "${UPDATE_UUID}" cloud/update-template.json

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DesignCenterCatalogTemplate "${TEMPLATE_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DesignCenterCatalogTemplate "${TEMPLATE_K8S}" 300
mv_wait_http delete template-gone "${TEMPLATE_URL}" 404
expect_audit_template delete audit-template

if [[ "${FUNCTIONAL}" == "1" ]]; then
  code="$(mv_rest delete list-templates GET "${TEMPLATES_URL}")"
  mv_expect_http delete "REST list templates returns 200" 200 "${code}" "cloud/delete-list-templates.json"
  if jq -e --arg name "${TEMPLATE_GCP_NAME}" '.catalogTemplates[]? | select(.name == $name)' "${MV_RUN_DIR}/cloud/delete-list-templates.json" >/dev/null; then
    mv_record delete "template is removed from catalog templates" FAIL cloud/delete-list-templates.json "template still found in list"
  else
    mv_record delete "template is removed from catalog templates" PASS cloud/delete-list-templates.json ""
  fi
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
