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

# Human-style validation of the NetworkManagementVPCFlowLogsConfig controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/networkmanagementvpcflowlogsconfig/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the networkmanagementvpcflowlogsconfig-maximal test fixture)
#   FUNCTIONAL=0     skip the functional checks (listing configs)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/networkmanagement/v1alpha1/networkmanagementvpcflowlogsconfig/networkmanagementvpcflowlogsconfig-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init NetworkManagementVPCFlowLogsConfig
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource IDs
DEPS=""
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE=""
if [[ -f "${FIXTURE_DIR}/update.yaml" ]]; then
  UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"
fi

# Resource identifiers derived from fixture
CONFIG_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
LOCATION="$(grep -m1 '^[[:space:]]*location:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
LOCATION="${LOCATION:-global}"

RESOURCE_ID="$(grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}' || true)"
CONFIG_ID="${RESOURCE_ID:-${CONFIG_K8S}}"

BASE_URL="https://networkmanagement.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}"
LIST_URL="${BASE_URL}/vpcFlowLogsConfigs"
CONFIG_URL="${LIST_URL}/${CONFIG_ID}"
CLOUD_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/vpcFlowLogsConfigs/${CONFIG_ID}"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get networkmanagementvpcflowlogsconfig "${CONFIG_K8S}" >/dev/null 2>&1 && mv_delete cleanup NetworkManagementVPCFlowLogsConfig "${CONFIG_K8S}" 300
  fi
  local code
  code="$(mv_rest cleanup leakcheck-config GET "${CONFIG_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (config gone)" PASS cloud/cleanup-leakcheck-config.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (config gone)" FAIL cloud/cleanup-leakcheck-config.json \
      "could not verify (HTTP ${code}); check manually: curl -H \"Authorization: Bearer \$(${MV_REST_TOKEN_CMD})\" ${CONFIG_URL}"
  else
    mv_record cleanup "no leaked GCP resources (config gone)" FAIL cloud/cleanup-leakcheck-config.json \
      "HTTP ${code}; clean up manually: DELETE ${CONFIG_URL}"
  fi

  code="$(mv_rest cleanup leakcheck-list GET "${LIST_URL}")"
  if [[ "${code}" == "200" ]]; then
    local leaked
    leaked="$(jq -r --arg uid "${MV_UNIQUE_ID}" '[.vpcFlowLogsConfigs[]? | select(.name | contains($uid))] | length' "${MV_RUN_DIR}/cloud/cleanup-leakcheck-list.json" 2>/dev/null || echo 0)"
    if [[ "${leaked}" == "0" ]]; then
      mv_record cleanup "no leaked GCP resources (no configs with uniqueId in name)" PASS cloud/cleanup-leakcheck-list.json "none found"
    else
      mv_record cleanup "no leaked GCP resources (no configs with uniqueId in name)" FAIL cloud/cleanup-leakcheck-list.json \
        "found ${leaked} leaked configs in ${LIST_URL}"
    fi
  else
    mv_record cleanup "no leaked GCP resources" FAIL cloud/cleanup-leakcheck-list.json "could not list configs: HTTP ${code}"
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

# Custom audit log check for NetworkManagementVPCFlowLogsConfig:
# NetworkManagement VPC Flow Logs Config uses Long Running Operations (LRO).
# Cloud Audit Logs record two entries per operation: initiation and completion (operation.last == true).
# We filter out completion entries (operation.last == true) to reflect the API calls initiated by KCC.
expect_flowlogs_audit() {
  local phase="$1" label="$2" config_id="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND \"${config_id}\" AND timestamp>=\"${MV_STARTED_UTC}\""
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
  jq -r '.[] | "\(.timestamp)  \(.protoPayload.methodName)  \(.protoPayload.authenticationInfo.principalEmail // "-")"' \
    "${MV_RUN_DIR}/${out}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.summary.txt" 2>/dev/null || true
  mv_trace "gcloud logging read '${filter}' --project ${GCP_PROJECT_ID} --order=asc --format=json" \
    "$(jq 'length' "${MV_RUN_DIR}/${out}" 2>/dev/null || echo 0) entries: ${got:-none}" "cloud/${phase}-${label}.summary.txt"
  mv_expect_eq "${phase}" "audit log write calls on ${config_id}" "${want}" "${got}" "${out}"
}

# ---------------------------------------------------------------------------
# Every API the controller, the dependencies and the out-of-band checks call.
mv_enable_apis setup networkmanagement.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
else
  mv_record deps "dependencies" INFO "" "none required for NetworkManagementVPCFlowLogsConfig"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create NetworkManagementVPCFlowLogsConfig "${CONFIG_K8S}" 300
SNAP="$(mv_snapshot create NetworkManagementVPCFlowLogsConfig "${CONFIG_K8S}")"
STATUS_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${STATUS_REF:-<none>}"

if [[ -n "${STATUS_REF}" ]]; then
  mv_expect_eq create "status.externalRef matches cloud resource name" "${CLOUD_NAME}" "${STATUS_REF}" "${SNAP#"${MV_RUN_DIR}"/}"
fi

mv_wait_http create config "${CONFIG_URL}" 200

# Compare spec fields against cloud response
WANT_DESC="$(jq -r '.spec.description // empty' "${SNAP}")"
if [[ -n "${WANT_DESC}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-config.json" '.description' "${WANT_DESC}"
fi

WANT_STATE="$(jq -r '.spec.state // empty' "${SNAP}")"
if [[ -n "${WANT_STATE}" ]]; then
  mv_expect_jq create "cloud state matches spec" "${MV_RUN_DIR}/cloud/create-config.json" '.state' "${WANT_STATE}"
fi

WANT_INTERVAL="$(jq -r '.spec.aggregationInterval // empty' "${SNAP}")"
if [[ -n "${WANT_INTERVAL}" ]]; then
  mv_expect_jq create "cloud aggregationInterval matches spec" "${MV_RUN_DIR}/cloud/create-config.json" '.aggregationInterval' "${WANT_INTERVAL}"
fi

WANT_SAMPLING="$(jq -r '.spec.flowSampling // empty | tostring' "${SNAP}")"
if [[ -n "${WANT_SAMPLING}" ]]; then
  ACTUAL_SAMPLING="$(jq -r '.flowSampling // empty | tostring' "${MV_RUN_DIR}/cloud/create-config.json")"
  mv_expect_eq create "cloud flowSampling matches spec" "${WANT_SAMPLING}" "${ACTUAL_SAMPLING}" cloud/create-config.json
fi

WANT_METADATA="$(jq -r '.spec.metadata // empty' "${SNAP}")"
if [[ -n "${WANT_METADATA}" ]]; then
  mv_expect_jq create "cloud metadata matches spec" "${MV_RUN_DIR}/cloud/create-config.json" '.metadata' "${WANT_METADATA}"
fi

WANT_METADATA_FIELDS="$(jq -r '[.spec.metadataFields[]?] | sort | join(",")' "${SNAP}")"
if [[ -n "${WANT_METADATA_FIELDS}" ]]; then
  ACTUAL_METADATA_FIELDS="$(jq -r '[.metadataFields[]?] | sort | join(",")' "${MV_RUN_DIR}/cloud/create-config.json")"
  mv_expect_eq create "cloud metadataFields match spec" "${WANT_METADATA_FIELDS}" "${ACTUAL_METADATA_FIELDS}" cloud/create-config.json
fi

WANT_FILTER="$(jq -r '.spec.filterExpr // empty' "${SNAP}")"
if [[ -n "${WANT_FILTER}" ]]; then
  mv_expect_jq create "cloud filterExpr matches spec" "${MV_RUN_DIR}/cloud/create-config.json" '.filterExpr' "${WANT_FILTER}"
fi

WANT_LABELS="$(jq -r '.spec.labels // {} | to_entries | sort_by(.key) | from_entries' "${SNAP}")"
ACTUAL_LABELS="$(jq -r '.labels // {} | to_entries | sort_by(.key) | from_entries' "${MV_RUN_DIR}/cloud/create-config.json")"
mv_expect_eq create "cloud labels match spec" "${WANT_LABELS}" "${ACTUAL_LABELS}" cloud/create-config.json

# Target resource (vpnTunnel or interconnectAttachment)
WANT_VPN="$(jq -r '.spec.vpnTunnelRef.external // empty' "${SNAP}")"
if [[ -n "${WANT_VPN}" ]]; then
  VPN_SUFFIX="${WANT_VPN#*regions/}"
  ACTUAL_VPN="$(jq -r '.vpnTunnel // empty' "${MV_RUN_DIR}/cloud/create-config.json")"
  if [[ "${ACTUAL_VPN}" == *"regions/${VPN_SUFFIX}"* ]]; then
    mv_record create "cloud vpnTunnel matches spec target" PASS cloud/create-config.json "${ACTUAL_VPN}"
  else
    mv_record create "cloud vpnTunnel matches spec target" FAIL cloud/create-config.json "expected suffix regions/${VPN_SUFFIX}, got ${ACTUAL_VPN}"
  fi
fi

WANT_INTERCONNECT="$(jq -r '.spec.interconnectAttachmentRef.external // empty' "${SNAP}")"
if [[ -n "${WANT_INTERCONNECT}" ]]; then
  IC_SUFFIX="${WANT_INTERCONNECT#*regions/}"
  ACTUAL_IC="$(jq -r '.interconnectAttachment // empty' "${MV_RUN_DIR}/cloud/create-config.json")"
  if [[ "${ACTUAL_IC}" == *"regions/${IC_SUFFIX}"* ]]; then
    mv_record create "cloud interconnectAttachment matches spec target" PASS cloud/create-config.json "${ACTUAL_IC}"
  else
    mv_record create "cloud interconnectAttachment matches spec target" FAIL cloud/create-config.json "expected suffix regions/${IC_SUFFIX}, got ${ACTUAL_IC}"
  fi
fi

# Out-of-band check via gcloud CLI describe
mv_cli create gcloud-describe -- gcloud network-management vpc-flow-logs-configs describe "${CONFIG_ID}" \
  --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
if grep -q "${CONFIG_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-describe.txt"; then
  mv_record create "gcloud describe succeeds and contains config id" PASS cloud/create-gcloud-describe.txt ""
else
  mv_record create "gcloud describe succeeds and contains config id" FAIL cloud/create-gcloud-describe.txt "config id not found in describe output"
fi

# Functional check: List contains the created config
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli create gcloud-list -- gcloud network-management vpc-flow-logs-configs list \
    --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
  if grep -q "${CONFIG_ID}" "${MV_RUN_DIR}/cloud/create-gcloud-list.txt"; then
    mv_record create "gcloud list contains created config" PASS cloud/create-gcloud-list.txt "${CLOUD_NAME}"
  else
    mv_record create "gcloud list contains created config" FAIL cloud/create-gcloud-list.txt "not found in list output"
  fi
fi

# --- Phase: update -----------------------------------------------------------
if [[ -n "${UPDATE:-}" ]]; then
  mv_apply update "${UPDATE}"
  sleep 2
  mv_wait_ready update NetworkManagementVPCFlowLogsConfig "${CONFIG_K8S}" 300
  SNAP="$(mv_snapshot update NetworkManagementVPCFlowLogsConfig "${CONFIG_K8S}")"
  mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

  mv_wait_http update config "${CONFIG_URL}" 200

  # Check updated fields
  WANT_UPD_DESC="$(jq -r '.spec.description // empty' "${SNAP}")"
  if [[ -n "${WANT_UPD_DESC}" ]]; then
    mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-config.json" '.description' "${WANT_UPD_DESC}"
  fi

  WANT_UPD_STATE="$(jq -r '.spec.state // empty' "${SNAP}")"
  if [[ -n "${WANT_UPD_STATE}" ]]; then
    mv_expect_jq update "cloud state updated to new spec" "${MV_RUN_DIR}/cloud/update-config.json" '.state' "${WANT_UPD_STATE}"
  fi

  WANT_UPD_SAMPLING="$(jq -r '.spec.flowSampling // empty | tostring' "${SNAP}")"
  if [[ -n "${WANT_UPD_SAMPLING}" ]]; then
    ACTUAL_UPD_SAMPLING="$(jq -r '.flowSampling // empty | tostring' "${MV_RUN_DIR}/cloud/update-config.json")"
    mv_expect_eq update "cloud flowSampling updated to new spec" "${WANT_UPD_SAMPLING}" "${ACTUAL_UPD_SAMPLING}" cloud/update-config.json
  fi

  WANT_UPD_LABELS="$(jq -r '.spec.labels // {} | to_entries | sort_by(.key) | from_entries' "${SNAP}")"
  ACTUAL_UPD_LABELS="$(jq -r '.labels // {} | to_entries | sort_by(.key) | from_entries' "${MV_RUN_DIR}/cloud/update-config.json")"
  mv_expect_eq update "cloud labels updated to new spec" "${WANT_UPD_LABELS}" "${ACTUAL_UPD_LABELS}" cloud/update-config.json

  # Verify unchanged fields remain intact
  if [[ -n "${WANT_INTERVAL}" ]]; then
    mv_expect_jq update "cloud aggregationInterval unchanged" "${MV_RUN_DIR}/cloud/update-config.json" '.aggregationInterval' "${WANT_INTERVAL}"
  fi
  if [[ -n "${WANT_METADATA}" ]]; then
    mv_expect_jq update "cloud metadata unchanged" "${MV_RUN_DIR}/cloud/update-config.json" '.metadata' "${WANT_METADATA}"
  fi
  if [[ -n "${WANT_METADATA_FIELDS}" ]]; then
    ACTUAL_METADATA_FIELDS="$(jq -r '[.metadataFields[]?] | sort | join(",")' "${MV_RUN_DIR}/cloud/update-config.json")"
    mv_expect_eq update "cloud metadataFields unchanged" "${WANT_METADATA_FIELDS}" "${ACTUAL_METADATA_FIELDS}" cloud/update-config.json
  fi
  if [[ -n "${WANT_FILTER}" ]]; then
    mv_expect_jq update "cloud filterExpr unchanged" "${MV_RUN_DIR}/cloud/update-config.json" '.filterExpr' "${WANT_FILTER}"
  fi

  mv_cli update gcloud-describe -- gcloud network-management vpc-flow-logs-configs describe "${CONFIG_ID}" \
    --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
  if grep -q "${CONFIG_ID}" "${MV_RUN_DIR}/cloud/update-gcloud-describe.txt"; then
    mv_record update "gcloud describe succeeds after update" PASS cloud/update-gcloud-describe.txt ""
  else
    mv_record update "gcloud describe succeeds after update" FAIL cloud/update-gcloud-describe.txt "describe failed after update"
  fi
else
  mv_record update "update manifest" INFO "" "no update.yaml in fixture; update phase skipped"
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady NetworkManagementVPCFlowLogsConfig "${CONFIG_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete NetworkManagementVPCFlowLogsConfig "${CONFIG_K8S}" 300
mv_wait_http delete config-gone "${CONFIG_URL}" 404
expect_flowlogs_audit delete audit-flowlogs "${CONFIG_ID}" \
  "CreateVpcFlowLogsConfig,UpdateVpcFlowLogsConfig,DeleteVpcFlowLogsConfig" 240

if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli delete gcloud-list -- gcloud network-management vpc-flow-logs-configs list \
    --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json || true
  if grep -q "${CONFIG_ID}" "${MV_RUN_DIR}/cloud/delete-gcloud-list.txt"; then
    mv_record delete "gcloud list no longer contains deleted config" FAIL cloud/delete-gcloud-list.txt "still found in list output"
  else
    mv_record delete "gcloud list no longer contains deleted config" PASS cloud/delete-gcloud-list.txt ""
  fi
fi

# --- Phase: dependency teardown ----------------------------------------------
if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
