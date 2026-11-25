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

# Human-style validation of the DiscoveryEngineWidgetConfig controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/discoveryenginewidgetconfig/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the discoveryenginewidgetconfig-maximal test fixture)
#   FUNCTIONAL=0     skip the functional checks (reading via Engine endpoint)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/discoveryengine/v1alpha1/discoveryenginewidgetconfig/discoveryenginewidgetconfig-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init DiscoveryEngineWidgetConfig
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
DATASTORE_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*datastore-' "${DEPS}" | awk '{print $2}' || true)"
DATASTORE_K8S="${DATASTORE_K8S:-datastore-${MV_UNIQUE_ID}}"
DATASTORE_ID="${DATASTORE_K8S}"

ENGINE_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*engine-' "${DEPS}" | awk '{print $2}' || true)"
ENGINE_K8S="${ENGINE_K8S:-engine-${MV_UNIQUE_ID}}"
ENGINE_ID="${ENGINE_K8S}"

WIDGET_CONFIG_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
WIDGET_CONFIG_K8S="${WIDGET_CONFIG_K8S:-widgetconfig-${MV_UNIQUE_ID}}"

WIDGET_CONFIG_ID="$(grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
WIDGET_CONFIG_ID="${WIDGET_CONFIG_ID:-default_search_widget_config}"

LOCATION="$(grep -m1 '^[[:space:]]*location:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
if [[ -z "${LOCATION}" && -n "${DEPS:-}" ]]; then
  LOCATION="$(grep -m1 '^[[:space:]]*location:[[:space:]]*' "${DEPS}" | awk '{print $2}' || true)"
fi
LOCATION="${LOCATION:-global}"

COLLECTION="$(grep -m1 '^[[:space:]]*collection:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
if [[ -z "${COLLECTION}" && -n "${DEPS:-}" ]]; then
  COLLECTION="$(grep -m1 '^[[:space:]]*collection:[[:space:]]*' "${DEPS}" | awk '{print $2}' || true)"
fi
COLLECTION="${COLLECTION:-default_collection}"

API_BASE="https://discoveryengine.googleapis.com/v1"
COLLECTION_URL="${API_BASE}/projects/${GCP_PROJECT_ID}/locations/${LOCATION}/collections/${COLLECTION}"
DATASTORE_URL="${COLLECTION_URL}/dataStores/${DATASTORE_ID}"
ENGINE_URL="${COLLECTION_URL}/engines/${ENGINE_ID}"
WIDGET_CONFIG_URL="${DATASTORE_URL}/widgetConfigs/${WIDGET_CONFIG_ID}"
ENGINE_WIDGET_CONFIG_URL="${ENGINE_URL}/widgetConfigs/${WIDGET_CONFIG_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get discoveryenginewidgetconfig "${WIDGET_CONFIG_K8S}" >/dev/null 2>&1 && mv_delete cleanup DiscoveryEngineWidgetConfig "${WIDGET_CONFIG_K8S}" 300
    mv_kubectl get discoveryengineengine "${ENGINE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DiscoveryEngineEngine "${ENGINE_K8S}" 300
    mv_kubectl get discoveryenginedatastore "${DATASTORE_K8S}" >/dev/null 2>&1 && mv_delete cleanup DiscoveryEngineDataStore "${DATASTORE_K8S}" 300
  fi

  local dcode
  dcode="$(mv_rest cleanup leakcheck-datastore GET "${DATASTORE_URL}")"
  if [[ "${dcode}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (dataStore gone)" PASS cloud/cleanup-leakcheck-datastore.json "HTTP 404"
  elif [[ "${dcode}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (dataStore gone)" FAIL cloud/cleanup-leakcheck-datastore.json \
      "could not verify (HTTP ${dcode}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" -H \"x-goog-user-project: ${GCP_PROJECT_ID}\" ${DATASTORE_URL}"
  else
    mv_record cleanup "no leaked GCP resources (dataStore gone)" FAIL cloud/cleanup-leakcheck-datastore.json \
      "HTTP ${dcode}; clean up manually: curl -X DELETE -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" -H \"x-goog-user-project: ${GCP_PROJECT_ID}\" ${DATASTORE_URL}"
  fi

  local ecode
  ecode="$(mv_rest cleanup leakcheck-engine GET "${ENGINE_URL}")"
  if [[ "${ecode}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (engine gone)" PASS cloud/cleanup-leakcheck-engine.json "HTTP 404"
  elif [[ "${ecode}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (engine gone)" FAIL cloud/cleanup-leakcheck-engine.json \
      "could not verify (HTTP ${ecode}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" -H \"x-goog-user-project: ${GCP_PROJECT_ID}\" ${ENGINE_URL}"
  else
    mv_record cleanup "no leaked GCP resources (engine gone)" FAIL cloud/cleanup-leakcheck-engine.json \
      "HTTP ${ecode}; clean up manually: curl -X DELETE -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" -H \"x-goog-user-project: ${GCP_PROJECT_ID}\" ${ENGINE_URL}"
  fi

  local wcode
  wcode="$(mv_rest cleanup leakcheck-widgetconfig GET "${WIDGET_CONFIG_URL}")"
  if [[ "${wcode}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (widgetConfig gone)" PASS cloud/cleanup-leakcheck-widgetconfig.json "HTTP 404"
  else
    mv_record cleanup "no leaked GCP resources (widgetConfig gone)" FAIL cloud/cleanup-leakcheck-widgetconfig.json "HTTP ${wcode}"
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
mv_enable_apis setup discoveryengine.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
  mv_wait_ready deps DiscoveryEngineDataStore "${DATASTORE_K8S}" 300
  mv_wait_ready deps DiscoveryEngineEngine "${ENGINE_K8S}" 300
  mv_wait_http deps datastore "${DATASTORE_URL}" 200
  mv_wait_http deps engine "${ENGINE_URL}" 200
  mv_wait_http deps widgetconfig-initialized "${WIDGET_CONFIG_URL}" 200
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create DiscoveryEngineWidgetConfig "${WIDGET_CONFIG_K8S}" 300
SNAP="$(mv_snapshot create DiscoveryEngineWidgetConfig "${WIDGET_CONFIG_K8S}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "$(jq -r '.status.externalRef // "<none>"' "${SNAP}")"

mv_wait_http create widgetconfig "${WIDGET_CONFIG_URL}" 200

# Note on gcloud CLI support:
# gcloud does not support DiscoveryEngine WidgetConfig (REST API only).
mv_record create "gcloud CLI surface" INFO env.yaml "no gcloud CLI surface for discoveryengine widgetConfigs; verified via REST"

# Compare spec fields against cloud response
WANT_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_ENABLE_RESULT_SCORE="$(jq -r 'if .spec.enableResultScore != null then (.spec.enableResultScore | tostring) else empty end' "${SNAP}")"
WANT_ALLOW_PUBLIC_ACCESS="$(jq -r 'if .spec.accessSettings.allowPublicAccess != null then (.spec.accessSettings.allowPublicAccess | tostring) else empty end' "${SNAP}")"
WANT_ALLOWLISTED_DOMAINS="$(jq -r 'if .spec.accessSettings.allowlistedDomains != null then ([.spec.accessSettings.allowlistedDomains[]?] | sort | join(",")) else empty end' "${SNAP}")"
WANT_LOGO_URL="$(jq -r '.spec.uiBranding.logo.url // empty' "${SNAP}")"
WANT_ENABLE_AUTOCOMPLETE="$(jq -r 'if .spec.uiSettings.enableAutocomplete != null then (.spec.uiSettings.enableAutocomplete | tostring) else empty end' "${SNAP}")"
WANT_ENABLE_QUALITY_FEEDBACK="$(jq -r 'if .spec.uiSettings.enableQualityFeedback != null then (.spec.uiSettings.enableQualityFeedback | tostring) else empty end' "${SNAP}")"
WANT_ENABLE_SAFE_SEARCH="$(jq -r 'if .spec.uiSettings.enableSafeSearch != null then (.spec.uiSettings.enableSafeSearch | tostring) else empty end' "${SNAP}")"
WANT_ENABLE_SEARCH_AS_YOU_TYPE="$(jq -r 'if .spec.uiSettings.enableSearchAsYouType != null then (.spec.uiSettings.enableSearchAsYouType | tostring) else empty end' "${SNAP}")"
WANT_DISABLE_USER_EVENTS="$(jq -r 'if .spec.uiSettings.disableUserEventsCollection != null then (.spec.uiSettings.disableUserEventsCollection | tostring) else empty end' "${SNAP}")"
WANT_INTERACTION_TYPE="$(jq -r '.spec.uiSettings.interactionType // empty' "${SNAP}")"
WANT_RESULT_COUNT="$(jq -r 'if .spec.uiSettings.generativeAnswerConfig.resultCount != null then (.spec.uiSettings.generativeAnswerConfig.resultCount | tostring) else empty end' "${SNAP}")"

if [[ -n "${WANT_DISPLAY_NAME}" ]]; then
  mv_expect_jq create "cloud displayName matches spec" "${MV_RUN_DIR}/cloud/create-widgetconfig.json" '.displayName // empty' "${WANT_DISPLAY_NAME}"
fi
if [[ -n "${WANT_ENABLE_RESULT_SCORE}" ]]; then
  mv_expect_jq create "cloud enableResultScore matches spec" "${MV_RUN_DIR}/cloud/create-widgetconfig.json" '(.enableResultScore // false) | tostring' "${WANT_ENABLE_RESULT_SCORE}"
fi
if [[ -n "${WANT_ALLOW_PUBLIC_ACCESS}" ]]; then
  mv_expect_jq create "cloud accessSettings.allowPublicAccess matches spec" "${MV_RUN_DIR}/cloud/create-widgetconfig.json" '(.accessSettings.allowPublicAccess // false) | tostring' "${WANT_ALLOW_PUBLIC_ACCESS}"
fi
if [[ -n "${WANT_ALLOWLISTED_DOMAINS}" ]]; then
  GOT_DOMAINS="$(jq -r '[.accessSettings.allowlistedDomains[]?] | sort | join(",")' "${MV_RUN_DIR}/cloud/create-widgetconfig.json")"
  mv_expect_eq create "cloud accessSettings.allowlistedDomains match spec" "${WANT_ALLOWLISTED_DOMAINS}" "${GOT_DOMAINS}" cloud/create-widgetconfig.json
fi
if [[ -n "${WANT_LOGO_URL}" ]]; then
  mv_expect_jq create "cloud uiBranding.logo.url matches spec" "${MV_RUN_DIR}/cloud/create-widgetconfig.json" '.uiBranding.logo.url // empty' "${WANT_LOGO_URL}"
fi
if [[ -n "${WANT_ENABLE_AUTOCOMPLETE}" ]]; then
  mv_expect_jq create "cloud uiSettings.enableAutocomplete matches spec" "${MV_RUN_DIR}/cloud/create-widgetconfig.json" '(.uiSettings.enableAutocomplete // false) | tostring' "${WANT_ENABLE_AUTOCOMPLETE}"
fi
if [[ -n "${WANT_ENABLE_QUALITY_FEEDBACK}" ]]; then
  mv_expect_jq create "cloud uiSettings.enableQualityFeedback matches spec" "${MV_RUN_DIR}/cloud/create-widgetconfig.json" '(.uiSettings.enableQualityFeedback // false) | tostring' "${WANT_ENABLE_QUALITY_FEEDBACK}"
fi
if [[ -n "${WANT_ENABLE_SAFE_SEARCH}" ]]; then
  mv_expect_jq create "cloud uiSettings.enableSafeSearch matches spec" "${MV_RUN_DIR}/cloud/create-widgetconfig.json" '(.uiSettings.enableSafeSearch // false) | tostring' "${WANT_ENABLE_SAFE_SEARCH}"
fi
if [[ -n "${WANT_ENABLE_SEARCH_AS_YOU_TYPE}" ]]; then
  mv_expect_jq create "cloud uiSettings.enableSearchAsYouType matches spec" "${MV_RUN_DIR}/cloud/create-widgetconfig.json" '(.uiSettings.enableSearchAsYouType // false) | tostring' "${WANT_ENABLE_SEARCH_AS_YOU_TYPE}"
fi
if [[ -n "${WANT_DISABLE_USER_EVENTS}" ]]; then
  mv_expect_jq create "cloud uiSettings.disableUserEventsCollection matches spec" "${MV_RUN_DIR}/cloud/create-widgetconfig.json" '(.uiSettings.disableUserEventsCollection // false) | tostring' "${WANT_DISABLE_USER_EVENTS}"
fi
if [[ -n "${WANT_INTERACTION_TYPE}" ]]; then
  GOT_INTERACTION="$(jq -r '.uiSettings.interactionType // empty' "${MV_RUN_DIR}/cloud/create-widgetconfig.json")"
  if [[ "${GOT_INTERACTION}" == "${WANT_INTERACTION_TYPE}" || \
        ( "${WANT_INTERACTION_TYPE}" == "SEARCH_ONLY" && "${GOT_INTERACTION}" == "1" ) || \
        ( "${WANT_INTERACTION_TYPE}" == "SEARCH_WITH_ANSWER" && "${GOT_INTERACTION}" == "2" ) ]]; then
    mv_record create "cloud uiSettings.interactionType matches spec" PASS cloud/create-widgetconfig.json "${GOT_INTERACTION}"
  else
    mv_record create "cloud uiSettings.interactionType matches spec" FAIL cloud/create-widgetconfig.json "want ${WANT_INTERACTION_TYPE}, got ${GOT_INTERACTION}"
  fi
fi
if [[ -n "${WANT_RESULT_COUNT}" ]]; then
  mv_expect_jq create "cloud uiSettings.generativeAnswerConfig.resultCount matches spec" "${MV_RUN_DIR}/cloud/create-widgetconfig.json" '(.uiSettings.generativeAnswerConfig.resultCount // empty) | tostring' "${WANT_RESULT_COUNT}"
fi

CONFIG_ID="$(jq -r '.configId // empty' "${MV_RUN_DIR}/cloud/create-widgetconfig.json")"
if [[ -n "${CONFIG_ID}" ]]; then
  mv_record create "configId populated in cloud" PASS cloud/create-widgetconfig.json "${CONFIG_ID}"
else
  mv_record create "configId populated in cloud" FAIL cloud/create-widgetconfig.json "empty configId"
fi

if [[ "${FUNCTIONAL}" == "1" ]]; then
  ENGINE_WC_CODE="$(mv_rest create get-engine-widgetconfig GET "${ENGINE_WIDGET_CONFIG_URL}")"
  mv_expect_http create "REST GET widgetConfig via Engine endpoint returns 200" 200 "${ENGINE_WC_CODE}" "cloud/create-get-engine-widgetconfig.json"
  ENGINE_WC_CONFIG_ID="$(jq -r '.configId // empty' "${MV_RUN_DIR}/cloud/create-get-engine-widgetconfig.json")"
  mv_expect_eq create "widgetConfig via Engine endpoint has matching configId" "${CONFIG_ID}" "${ENGINE_WC_CONFIG_ID}" cloud/create-get-engine-widgetconfig.json
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update DiscoveryEngineWidgetConfig "${WIDGET_CONFIG_K8S}" 300
SNAP="$(mv_snapshot update DiscoveryEngineWidgetConfig "${WIDGET_CONFIG_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update widgetconfig "${WIDGET_CONFIG_URL}" 200

# Compare updated spec fields against cloud response
WANT_UPDATED_DISPLAY_NAME="$(jq -r '.spec.displayName // empty' "${SNAP}")"
WANT_UPDATED_ENABLE_RESULT_SCORE="$(jq -r 'if .spec.enableResultScore != null then (.spec.enableResultScore | tostring) else empty end' "${SNAP}")"
WANT_UPDATED_ALLOW_PUBLIC_ACCESS="$(jq -r 'if .spec.accessSettings.allowPublicAccess != null then (.spec.accessSettings.allowPublicAccess | tostring) else empty end' "${SNAP}")"
WANT_UPDATED_ALLOWLISTED_DOMAINS="$(jq -r 'if .spec.accessSettings.allowlistedDomains != null then ([.spec.accessSettings.allowlistedDomains[]?] | sort | join(",")) else empty end' "${SNAP}")"
WANT_UPDATED_LOGO_URL="$(jq -r '.spec.uiBranding.logo.url // empty' "${SNAP}")"
WANT_UPDATED_ENABLE_AUTOCOMPLETE="$(jq -r 'if .spec.uiSettings.enableAutocomplete != null then (.spec.uiSettings.enableAutocomplete | tostring) else empty end' "${SNAP}")"
WANT_UPDATED_ENABLE_QUALITY_FEEDBACK="$(jq -r 'if .spec.uiSettings.enableQualityFeedback != null then (.spec.uiSettings.enableQualityFeedback | tostring) else empty end' "${SNAP}")"
WANT_UPDATED_ENABLE_SAFE_SEARCH="$(jq -r 'if .spec.uiSettings.enableSafeSearch != null then (.spec.uiSettings.enableSafeSearch | tostring) else empty end' "${SNAP}")"
WANT_UPDATED_ENABLE_SEARCH_AS_YOU_TYPE="$(jq -r 'if .spec.uiSettings.enableSearchAsYouType != null then (.spec.uiSettings.enableSearchAsYouType | tostring) else empty end' "${SNAP}")"
WANT_UPDATED_DISABLE_USER_EVENTS="$(jq -r 'if .spec.uiSettings.disableUserEventsCollection != null then (.spec.uiSettings.disableUserEventsCollection | tostring) else empty end' "${SNAP}")"
WANT_UPDATED_INTERACTION_TYPE="$(jq -r '.spec.uiSettings.interactionType // empty' "${SNAP}")"
WANT_UPDATED_RESULT_COUNT="$(jq -r 'if .spec.uiSettings.generativeAnswerConfig.resultCount != null then (.spec.uiSettings.generativeAnswerConfig.resultCount | tostring) else empty end' "${SNAP}")"

if [[ -n "${WANT_UPDATED_DISPLAY_NAME}" ]]; then
  mv_expect_jq update "cloud displayName updated to new spec" "${MV_RUN_DIR}/cloud/update-widgetconfig.json" '.displayName // empty' "${WANT_UPDATED_DISPLAY_NAME}"
fi
if [[ -n "${WANT_UPDATED_ENABLE_RESULT_SCORE}" ]]; then
  mv_expect_jq update "cloud enableResultScore updated to new spec" "${MV_RUN_DIR}/cloud/update-widgetconfig.json" '(.enableResultScore // false) | tostring' "${WANT_UPDATED_ENABLE_RESULT_SCORE}"
fi
if [[ -n "${WANT_UPDATED_ALLOW_PUBLIC_ACCESS}" ]]; then
  mv_expect_jq update "cloud accessSettings.allowPublicAccess updated to new spec" "${MV_RUN_DIR}/cloud/update-widgetconfig.json" '(.accessSettings.allowPublicAccess // false) | tostring' "${WANT_UPDATED_ALLOW_PUBLIC_ACCESS}"
fi
if [[ -n "${WANT_UPDATED_ALLOWLISTED_DOMAINS}" ]]; then
  GOT_DOMAINS="$(jq -r '[.accessSettings.allowlistedDomains[]?] | sort | join(",")' "${MV_RUN_DIR}/cloud/update-widgetconfig.json")"
  mv_expect_eq update "cloud accessSettings.allowlistedDomains updated to new spec" "${WANT_UPDATED_ALLOWLISTED_DOMAINS}" "${GOT_DOMAINS}" cloud/update-widgetconfig.json
fi
if [[ -n "${WANT_UPDATED_LOGO_URL}" ]]; then
  mv_expect_jq update "cloud uiBranding.logo.url updated to new spec" "${MV_RUN_DIR}/cloud/update-widgetconfig.json" '.uiBranding.logo.url // empty' "${WANT_UPDATED_LOGO_URL}"
fi
if [[ -n "${WANT_UPDATED_ENABLE_AUTOCOMPLETE}" ]]; then
  mv_expect_jq update "cloud uiSettings.enableAutocomplete updated to new spec" "${MV_RUN_DIR}/cloud/update-widgetconfig.json" '(.uiSettings.enableAutocomplete // false) | tostring' "${WANT_UPDATED_ENABLE_AUTOCOMPLETE}"
fi
if [[ -n "${WANT_UPDATED_ENABLE_QUALITY_FEEDBACK}" ]]; then
  mv_expect_jq update "cloud uiSettings.enableQualityFeedback updated to new spec" "${MV_RUN_DIR}/cloud/update-widgetconfig.json" '(.uiSettings.enableQualityFeedback // false) | tostring' "${WANT_UPDATED_ENABLE_QUALITY_FEEDBACK}"
fi
if [[ -n "${WANT_UPDATED_ENABLE_SAFE_SEARCH}" ]]; then
  mv_expect_jq update "cloud uiSettings.enableSafeSearch updated to new spec" "${MV_RUN_DIR}/cloud/update-widgetconfig.json" '(.uiSettings.enableSafeSearch // false) | tostring' "${WANT_UPDATED_ENABLE_SAFE_SEARCH}"
fi
if [[ -n "${WANT_UPDATED_ENABLE_SEARCH_AS_YOU_TYPE}" ]]; then
  mv_expect_jq update "cloud uiSettings.enableSearchAsYouType updated to new spec" "${MV_RUN_DIR}/cloud/update-widgetconfig.json" '(.uiSettings.enableSearchAsYouType // false) | tostring' "${WANT_UPDATED_ENABLE_SEARCH_AS_YOU_TYPE}"
fi
if [[ -n "${WANT_UPDATED_DISABLE_USER_EVENTS}" ]]; then
  mv_expect_jq update "cloud uiSettings.disableUserEventsCollection updated to new spec" "${MV_RUN_DIR}/cloud/update-widgetconfig.json" '(.uiSettings.disableUserEventsCollection // false) | tostring' "${WANT_UPDATED_DISABLE_USER_EVENTS}"
fi
if [[ -n "${WANT_UPDATED_INTERACTION_TYPE}" ]]; then
  GOT_INTERACTION="$(jq -r '.uiSettings.interactionType // empty' "${MV_RUN_DIR}/cloud/update-widgetconfig.json")"
  if [[ "${GOT_INTERACTION}" == "${WANT_UPDATED_INTERACTION_TYPE}" || \
        ( "${WANT_UPDATED_INTERACTION_TYPE}" == "SEARCH_ONLY" && "${GOT_INTERACTION}" == "1" ) || \
        ( "${WANT_UPDATED_INTERACTION_TYPE}" == "SEARCH_WITH_ANSWER" && "${GOT_INTERACTION}" == "2" ) ]]; then
    mv_record update "cloud uiSettings.interactionType updated to new spec" PASS cloud/update-widgetconfig.json "${GOT_INTERACTION}"
  else
    mv_record update "cloud uiSettings.interactionType updated to new spec" FAIL cloud/update-widgetconfig.json "want ${WANT_UPDATED_INTERACTION_TYPE}, got ${GOT_INTERACTION}"
  fi
fi
if [[ -n "${WANT_UPDATED_RESULT_COUNT}" ]]; then
  mv_expect_jq update "cloud uiSettings.generativeAnswerConfig.resultCount updated to new spec" "${MV_RUN_DIR}/cloud/update-widgetconfig.json" '(.uiSettings.generativeAnswerConfig.resultCount // empty) | tostring' "${WANT_UPDATED_RESULT_COUNT}"
fi

UPDATE_CONFIG_ID="$(jq -r '.configId // empty' "${MV_RUN_DIR}/cloud/update-widgetconfig.json")"
mv_expect_eq update "configId preserved across update (in-place update)" "${CONFIG_ID}" "${UPDATE_CONFIG_ID}" cloud/update-widgetconfig.json

if [[ "${FUNCTIONAL}" == "1" ]]; then
  ENGINE_WC_CODE="$(mv_rest update get-engine-widgetconfig GET "${ENGINE_WIDGET_CONFIG_URL}")"
  mv_expect_http update "REST GET widgetConfig via Engine endpoint returns 200 after update" 200 "${ENGINE_WC_CODE}" "cloud/update-get-engine-widgetconfig.json"
  UPDATED_ENGINE_DISPLAY="$(jq -r '.displayName // empty' "${MV_RUN_DIR}/cloud/update-get-engine-widgetconfig.json")"
  mv_expect_eq update "Engine endpoint reflects updated displayName" "${WANT_UPDATED_DISPLAY_NAME}" "${UPDATED_ENGINE_DISPLAY}" cloud/update-get-engine-widgetconfig.json
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady DiscoveryEngineWidgetConfig "${WIDGET_CONFIG_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete DiscoveryEngineWidgetConfig "${WIDGET_CONFIG_K8S}" 300
mv_record delete "WidgetConfig deletion lifecycle in GCP" INFO "cloud/delete-widgetconfig-lifecycle.json" \
  "GCP Discovery Engine does not support deleting WidgetConfig (singleton resource). KCC removes finalizer without calling GCP Delete; resource is deleted when parent DataStore is deleted."

# Audit log check:
# Reads Cloud Audit Logs (admin activity) for write calls on the widgetConfig resource.
# WidgetConfig in GCP only supports UpdateWidgetConfig (called on initial configuration and update).
expect_widgetconfig_audit() {
  local phase="$1" label="$2" want="$3" timeout="${4:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.resourceName:\"${WIDGET_CONFIG_ID}\" AND timestamp>=\"${MV_STARTED_UTC}\""
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
  mv_expect_eq "${phase}" "audit log write calls for ${WIDGET_CONFIG_ID}" "${want}" "${got}" "${out}"
}

expect_widgetconfig_audit delete audit-widgetconfig "UpdateWidgetConfig,UpdateWidgetConfig"

# --- Phase: teardown ---------------------------------------------------------
mv_delete teardown DiscoveryEngineEngine "${ENGINE_K8S}" 300
mv_wait_http teardown engine-gone "${ENGINE_URL}" 404
mv_delete teardown DiscoveryEngineDataStore "${DATASTORE_K8S}" 300
mv_wait_http teardown datastore-gone "${DATASTORE_URL}" 404
mv_wait_http teardown widgetconfig-gone "${WIDGET_CONFIG_URL}" 404

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
