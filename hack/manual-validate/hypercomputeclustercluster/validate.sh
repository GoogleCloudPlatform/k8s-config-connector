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

# Human-style validation of the HypercomputeClusterCluster controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/hypercomputeclustercluster/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: the hypercomputeclustercluster-maximal test fixture)
#   FUNCTIONAL=0     skip functional checks (gcloud cluster-director list and compute instance verification)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${MV_REPO_ROOT}/pkg/test/resourcefixture/testdata/basic/hypercomputecluster/v1alpha1/hypercomputeclustercluster/hypercomputeclustercluster-maximal}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init HypercomputeClusterCluster
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

# Names derived from the manifests so out-of-band checks never ask KCC.
CLUSTER_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}' | tr -d '"'\''')"
LOCATION="$(grep -m1 '^[[:space:]]*location:[[:space:]]*' "${CREATE}" | awk '{print $2}' | tr -d '"'\''')"
LOCATION="${LOCATION:-us-central1}"
CLUSTER_ID="$(grep -m1 '^[[:space:]]*resourceID:[[:space:]]*' "${CREATE}" | awk '{print $2}' | tr -d '"'\''')"
CLUSTER_ID="${CLUSTER_ID:-${CLUSTER_K8S}}"

EXPECTED_NAME="projects/${GCP_PROJECT_ID}/locations/${LOCATION}/clusters/${CLUSTER_ID}"
CLUSTER_URL="https://hypercomputecluster.googleapis.com/v1/${EXPECTED_NAME}"
LIST_URL="https://hypercomputecluster.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/${LOCATION}/clusters"
LOGIN_VM=""
LOGIN_ZONE=""

# Helper to convert mv_cli output to pure JSON for jq / mv_expect_jq
mv_strip_cli_json() {
  local src="$1" dst="$2"
  sed -e '1{/^\$ /d;}' -e '/^exit_code: [0-9]*$/d' "${src}" >"${dst}"
}

# Custom audit log check for HypercomputeClusterCluster:
# In Cloud Audit Logs:
# 1) HypercomputeCluster API methods are Long Running Operations (LROs) that emit
#    two audit logs per call: initiation (operation.first == true or operation.last != true)
#    and completion (operation.last == true). We filter to initiated calls (operation.last != true).
# 2) Under the hood, HypercomputeCluster creates Compute Engine VMs and disks which also
#    emit audit logs under compute.googleapis.com; we scope to service hypercomputecluster.googleapis.com.
expect_cluster_audit() {
  local phase="$1" label="$2" cluster_id="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.serviceName=\"hypercomputecluster.googleapis.com\" AND \"${cluster_id}\" AND timestamp>=\"${MV_STARTED_UTC}\""
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
  mv_expect_eq "${phase}" "audit log write calls on ${cluster_id}" "${want}" "${got}" "${out}"
}

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get hypercomputeclustercluster "${CLUSTER_K8S}" >/dev/null 2>&1 && mv_delete cleanup HypercomputeClusterCluster "${CLUSTER_K8S}" 900
    if [[ -n "${DEPS:-}" ]]; then
      :
    fi
  fi
  local code
  code="$(mv_rest cleanup leakcheck-cluster GET "${CLUSTER_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (cluster gone)" PASS cloud/cleanup-leakcheck-cluster.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (cluster gone)" FAIL cloud/cleanup-leakcheck-cluster.json \
      "could not verify (HTTP ${code}); check manually: gcloud cluster-director clusters describe ${CLUSTER_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP resources (cluster gone)" FAIL cloud/cleanup-leakcheck-cluster.json \
      "HTTP ${code}; clean up manually: gcloud cluster-director clusters delete ${CLUSTER_ID} --location=${LOCATION} --project=${GCP_PROJECT_ID} --quiet"
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
mv_enable_apis setup hypercomputecluster.googleapis.com compute.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
else
  mv_record deps "dependencies" INFO "" "none required for HypercomputeClusterCluster (uses default network)"
fi

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create HypercomputeClusterCluster "${CLUSTER_K8S}" 900
SNAP="$(mv_snapshot create HypercomputeClusterCluster "${CLUSTER_K8S}")"
STATUS_REF="$(jq -r '.status.externalRef // empty' "${SNAP}")"
mv_record create "status.externalRef" INFO "${SNAP#"${MV_RUN_DIR}"/}" "${STATUS_REF:-<none>}"

WANT_DESC="$(jq -r '.spec.description // empty' "${SNAP}")"
WANT_LABELS="$(jq -r '.spec.labels // empty' "${SNAP}")"
WANT_COMPUTE="$(jq -r '.spec.computeResources // empty' "${SNAP}")"
WANT_NETWORK="$(jq -r '.spec.networkResources // empty' "${SNAP}")"
WANT_ORCH="$(jq -r '.spec.orchestrator // empty' "${SNAP}")"

mv_wait_http create cluster "${CLUSTER_URL}" 200 300

if [[ -n "${STATUS_REF}" ]]; then
  mv_expect_eq create "status.externalRef matches cloud resource name" "${EXPECTED_NAME}" "${STATUS_REF}" "${SNAP#"${MV_RUN_DIR}"/}"
fi

mv_expect_jq create "cloud cluster name matches expected" "${MV_RUN_DIR}/cloud/create-cluster.json" '.name' "${EXPECTED_NAME}"

if [[ -n "${WANT_DESC}" ]]; then
  mv_expect_jq create "cloud description matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" '.description' "${WANT_DESC}"
fi

# Labels: check each label in spec
if [[ -n "${WANT_LABELS}" && "${WANT_LABELS}" != "null" ]]; then
  for k in $(jq -r '.spec.labels // {} | keys[]' "${SNAP}"); do
    val="$(jq -r --arg k "${k}" '.spec.labels[$k]' "${SNAP}")"
    mv_expect_jq create "cloud label ${k} matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" ".labels[\"${k}\"]" "${val}"
  done
fi

# Network resources
if [[ -n "${WANT_NETWORK}" && "${WANT_NETWORK}" != "null" ]]; then
  for k in $(jq -r '.spec.networkResources // {} | keys[]' "${SNAP}"); do
    net="$(jq -r --arg k "${k}" '.spec.networkResources[$k].config.existingNetwork.networkRef.external // empty' "${SNAP}")"
    subnet="$(jq -r --arg k "${k}" '.spec.networkResources[$k].config.existingNetwork.subnetworkRef.external // empty' "${SNAP}")"
    if [[ -n "${net}" ]]; then
      mv_expect_jq create "cloud networkResources[${k}] network matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        ".networkResources[\"${k}\"].config.existingNetwork.network" "${net}"
    fi
    if [[ -n "${subnet}" ]]; then
      mv_expect_jq create "cloud networkResources[${k}] subnetwork matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        ".networkResources[\"${k}\"].config.existingNetwork.subnetwork" "${subnet}"
    fi
  done
fi

# Compute resources
if [[ -n "${WANT_COMPUTE}" && "${WANT_COMPUTE}" != "null" ]]; then
  for k in $(jq -r '.spec.computeResources // {} | keys[]' "${SNAP}"); do
    mt="$(jq -r --arg k "${k}" '.spec.computeResources[$k].config.newOnDemandInstances.machineType // empty' "${SNAP}")"
    z="$(jq -r --arg k "${k}" '.spec.computeResources[$k].config.newOnDemandInstances.zone // empty' "${SNAP}")"
    if [[ -n "${mt}" ]]; then
      mv_expect_jq create "cloud computeResources[${k}] machineType matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        ".computeResources[\"${k}\"].config.newOnDemandInstances.machineType" "${mt}"
    fi
    if [[ -n "${z}" ]]; then
      mv_expect_jq create "cloud computeResources[${k}] zone matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        ".computeResources[\"${k}\"].config.newOnDemandInstances.zone" "${z}"
    fi
  done
fi

# Orchestrator Slurm
if [[ -n "${WANT_ORCH}" && "${WANT_ORCH}" != "null" ]]; then
  default_part="$(jq -r '.spec.orchestrator.slurm.defaultPartition // empty' "${SNAP}")"
  if [[ -n "${default_part}" ]]; then
    mv_expect_jq create "cloud orchestrator slurm defaultPartition matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
      '.orchestrator.slurm.defaultPartition' "${default_part}"
  fi
  login_mt="$(jq -r '.spec.orchestrator.slurm.loginNodes.machineType // empty' "${SNAP}")"
  if [[ -n "${login_mt}" ]]; then
    mv_expect_jq create "cloud orchestrator slurm loginNodes machineType matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
      '.orchestrator.slurm.loginNodes.machineType' "${login_mt}"
  fi
  LOGIN_ZONE="$(jq -r '.spec.orchestrator.slurm.loginNodes.zone // empty' "${SNAP}")"
  if [[ -n "${LOGIN_ZONE}" ]]; then
    mv_expect_jq create "cloud orchestrator slurm loginNodes zone matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
      '.orchestrator.slurm.loginNodes.zone' "${LOGIN_ZONE}"
  fi
  login_count="$(jq -r '.spec.orchestrator.slurm.loginNodes.count // empty' "${SNAP}")"
  if [[ -n "${login_count}" ]]; then
    mv_expect_jq create "cloud orchestrator slurm loginNodes count matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
      '((.orchestrator.slurm.loginNodes.count // empty) | tostring)' "${login_count}"
  fi
  login_os_login="$(jq -r '.spec.orchestrator.slurm.loginNodes.enableOSLogin // empty' "${SNAP}")"
  if [[ -n "${login_os_login}" ]]; then
    mv_expect_jq create "cloud orchestrator slurm loginNodes enableOsLogin matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
      '((.orchestrator.slurm.loginNodes.enableOsLogin // empty) | tostring)' "${login_os_login}"
  fi
  login_public_ips="$(jq -r '.spec.orchestrator.slurm.loginNodes.enablePublicIPs // empty' "${SNAP}")"
  if [[ -n "${login_public_ips}" ]]; then
    mv_expect_jq create "cloud orchestrator slurm loginNodes enablePublicIps matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
      '((.orchestrator.slurm.loginNodes.enablePublicIps // empty) | tostring)' "${login_public_ips}"
  fi

  # Slurm NodeSets
  nodeset_count="$(jq -r '.spec.orchestrator.slurm.nodeSets // [] | length' "${SNAP}")"
  for ((i=0; i<nodeset_count; i++)); do
    ns_id="$(jq -r --argjson i "${i}" '.spec.orchestrator.slurm.nodeSets[$i].id // empty' "${SNAP}")"
    ns_comp_id="$(jq -r --argjson i "${i}" '.spec.orchestrator.slurm.nodeSets[$i].computeID // empty' "${SNAP}")"
    ns_static_count="$(jq -r --argjson i "${i}" '.spec.orchestrator.slurm.nodeSets[$i].staticNodeCount // empty' "${SNAP}")"
    ns_boot_disk_size="$(jq -r --argjson i "${i}" '.spec.orchestrator.slurm.nodeSets[$i].computeInstance.bootDisk.sizeGB // empty' "${SNAP}")"
    ns_boot_disk_type="$(jq -r --argjson i "${i}" '.spec.orchestrator.slurm.nodeSets[$i].computeInstance.bootDisk.type // empty' "${SNAP}")"
    ns_startup_script="$(jq -r --argjson i "${i}" '.spec.orchestrator.slurm.nodeSets[$i].computeInstance.startupScript // empty' "${SNAP}")"

    if [[ -n "${ns_id}" ]]; then
      mv_expect_jq create "cloud nodeSets[${i}] id matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        "([.orchestrator.slurm.nodeSets[]? | select(.id==\"${ns_id}\")][0].id // empty)" "${ns_id}"
    fi
    if [[ -n "${ns_comp_id}" ]]; then
      mv_expect_jq create "cloud nodeSets[${i}] computeId matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        "([.orchestrator.slurm.nodeSets[]? | select(.id==\"${ns_id}\")][0].computeId // empty)" "${ns_comp_id}"
    fi
    if [[ -n "${ns_static_count}" ]]; then
      mv_expect_jq create "cloud nodeSets[${i}] staticNodeCount matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        "(([.orchestrator.slurm.nodeSets[]? | select(.id==\"${ns_id}\")][0].staticNodeCount // empty) | tostring)" "${ns_static_count}"
    fi
    if [[ -n "${ns_boot_disk_size}" ]]; then
      mv_expect_jq create "cloud nodeSets[${i}] bootDisk sizeGb matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        "(([.orchestrator.slurm.nodeSets[]? | select(.id==\"${ns_id}\")][0].computeInstance.bootDisk.sizeGb // empty) | tostring)" "${ns_boot_disk_size}"
    fi
    if [[ -n "${ns_boot_disk_type}" ]]; then
      mv_expect_jq create "cloud nodeSets[${i}] bootDisk type matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        "([.orchestrator.slurm.nodeSets[]? | select(.id==\"${ns_id}\")][0].computeInstance.bootDisk.type // empty)" "${ns_boot_disk_type}"
    fi
    if [[ -n "${ns_startup_script}" ]]; then
      mv_expect_jq create "cloud nodeSets[${i}] startupScript matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        "([.orchestrator.slurm.nodeSets[]? | select(.id==\"${ns_id}\")][0].computeInstance.startupScript // empty)" "${ns_startup_script}"
    fi
    for lk in $(jq -r --argjson i "${i}" '.spec.orchestrator.slurm.nodeSets[$i].computeInstance.labels // {} | keys[]' "${SNAP}"); do
      lv="$(jq -r --argjson i "${i}" --arg lk "${lk}" '.spec.orchestrator.slurm.nodeSets[$i].computeInstance.labels[$lk]' "${SNAP}")"
      mv_expect_jq create "cloud nodeSets[${i}] label ${lk} matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        "([.orchestrator.slurm.nodeSets[]? | select(.id==\"${ns_id}\")][0].computeInstance.labels[\"${lk}\"] // empty)" "${lv}"
    done
  done

  # Slurm Partitions
  part_count="$(jq -r '.spec.orchestrator.slurm.partitions // [] | length' "${SNAP}")"
  for ((i=0; i<part_count; i++)); do
    part_id="$(jq -r --argjson i "${i}" '.spec.orchestrator.slurm.partitions[$i].id // empty' "${SNAP}")"
    part_nodesets="$(jq -r --argjson i "${i}" '[.spec.orchestrator.slurm.partitions[$i].nodeSetIds[]?] | sort | join(",")' "${SNAP}")"
    if [[ -n "${part_id}" ]]; then
      mv_expect_jq create "cloud partitions[${i}] id matches spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        "([.orchestrator.slurm.partitions[]? | select(.id==\"${part_id}\")][0].id // empty)" "${part_id}"
    fi
    if [[ -n "${part_nodesets}" ]]; then
      mv_expect_jq create "cloud partitions[${i}] nodeSetIds match spec" "${MV_RUN_DIR}/cloud/create-cluster.json" \
        "([.orchestrator.slurm.partitions[]? | select(.id==\"${part_id}\")][0].nodeSetIds // [] | sort | join(\",\"))" "${part_nodesets}"
    fi
  done
fi

# Out-of-band describe via gcloud cluster-director
mv_cli create gcloud-describe -- gcloud cluster-director clusters describe "${CLUSTER_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json
mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-describe.txt" "${MV_RUN_DIR}/cloud/create-gcloud-describe.json"
mv_expect_jq create "gcloud describe name matches expected" "${MV_RUN_DIR}/cloud/create-gcloud-describe.json" '.name' "${EXPECTED_NAME}"

# Functional checks
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli create gcloud-list -- gcloud cluster-director clusters list --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json
  mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-list.txt" "${MV_RUN_DIR}/cloud/create-gcloud-list.json"
  FOUND_CLUSTER="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/create-gcloud-list.json")"
  mv_expect_eq create "gcloud cluster-director clusters list returns created cluster" "${EXPECTED_NAME}" "${FOUND_CLUSTER}" cloud/create-gcloud-list.json

  # Extract provisioned login node VM from cloud response and verify it is RUNNING in Compute Engine
  LOGIN_VM="$(jq -r '.orchestrator.slurm.loginNodes.instances[0].instance // empty' "${MV_RUN_DIR}/cloud/create-cluster.json")"
  if [[ -n "${LOGIN_VM}" ]]; then
    vm_zone="${LOGIN_ZONE:-us-central1-a}"
    mv_cli create gcloud-describe-instance -- gcloud compute instances describe "${LOGIN_VM##*/}" --zone="${vm_zone}" --project="${GCP_PROJECT_ID}" --format=json
    mv_strip_cli_json "${MV_RUN_DIR}/cloud/create-gcloud-describe-instance.txt" "${MV_RUN_DIR}/cloud/create-gcloud-describe-instance.json"
    mv_expect_jq create "login node VM is RUNNING in Compute Engine" "${MV_RUN_DIR}/cloud/create-gcloud-describe-instance.json" '.status' "RUNNING"
  fi
fi

# --- Phase: update -----------------------------------------------------------
if [[ -n "${UPDATE:-}" ]]; then
  mv_apply update "${UPDATE}"
  sleep 2
  mv_wait_ready update HypercomputeClusterCluster "${CLUSTER_K8S}" 900
  SNAP="$(mv_snapshot update HypercomputeClusterCluster "${CLUSTER_K8S}")"
  mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

  mv_wait_http update cluster "${CLUSTER_URL}" 200 300

  WANT_DESC="$(jq -r '.spec.description // empty' "${SNAP}")"
  if [[ -n "${WANT_DESC}" ]]; then
    mv_expect_jq update "cloud description updated to new spec" "${MV_RUN_DIR}/cloud/update-cluster.json" '.description' "${WANT_DESC}"
  fi

  for k in $(jq -r '.spec.labels // {} | keys[]' "${SNAP}"); do
    val="$(jq -r --arg k "${k}" '.spec.labels[$k]' "${SNAP}")"
    mv_expect_jq update "cloud label ${k} updated to new spec" "${MV_RUN_DIR}/cloud/update-cluster.json" ".labels[\"${k}\"]" "${val}"
  done

  # Verify updated nodeset label
  nodeset_count="$(jq -r '.spec.orchestrator.slurm.nodeSets // [] | length' "${SNAP}")"
  for ((i=0; i<nodeset_count; i++)); do
    ns_id="$(jq -r --argjson i "${i}" '.spec.orchestrator.slurm.nodeSets[$i].id // empty' "${SNAP}")"
    for lk in $(jq -r --argjson i "${i}" '.spec.orchestrator.slurm.nodeSets[$i].computeInstance.labels // {} | keys[]' "${SNAP}"); do
      lv="$(jq -r --argjson i "${i}" --arg lk "${lk}" '.spec.orchestrator.slurm.nodeSets[$i].computeInstance.labels[$lk]' "${SNAP}")"
      mv_expect_jq update "cloud nodeSets[${i}] label ${lk} updated to new spec" "${MV_RUN_DIR}/cloud/update-cluster.json" \
        "([.orchestrator.slurm.nodeSets[]? | select(.id==\"${ns_id}\")][0].computeInstance.labels[\"${lk}\"] // empty)" "${lv}"
    done
  done

  # Unchanged fields intact:
  default_part="$(jq -r '.spec.orchestrator.slurm.defaultPartition // empty' "${SNAP}")"
  if [[ -n "${default_part}" ]]; then
    mv_expect_jq update "cloud orchestrator slurm defaultPartition intact" "${MV_RUN_DIR}/cloud/update-cluster.json" \
      '.orchestrator.slurm.defaultPartition' "${default_part}"
  fi

  mv_cli update gcloud-describe -- gcloud cluster-director clusters describe "${CLUSTER_ID}" --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json
  mv_strip_cli_json "${MV_RUN_DIR}/cloud/update-gcloud-describe.txt" "${MV_RUN_DIR}/cloud/update-gcloud-describe.json"
  mv_expect_jq update "gcloud describe description updated" "${MV_RUN_DIR}/cloud/update-gcloud-describe.json" '.description' "${WANT_DESC}"
fi

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady HypercomputeClusterCluster "${CLUSTER_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete HypercomputeClusterCluster "${CLUSTER_K8S}" 900
mv_wait_http delete cluster-gone "${CLUSTER_URL}" 404 300

EXPECTED_AUDIT="CreateCluster,DeleteCluster"
if [[ -n "${UPDATE:-}" ]]; then
  EXPECTED_AUDIT="CreateCluster,UpdateCluster,DeleteCluster"
fi
expect_cluster_audit delete audit-cluster "${CLUSTER_ID}" "${EXPECTED_AUDIT}" 300

if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli delete gcloud-list -- gcloud cluster-director clusters list --location="${LOCATION}" --project="${GCP_PROJECT_ID}" --format=json
  mv_strip_cli_json "${MV_RUN_DIR}/cloud/delete-gcloud-list.txt" "${MV_RUN_DIR}/cloud/delete-gcloud-list.json"
  DELETED_CLUSTER="$(jq -r --arg NAME "${EXPECTED_NAME}" '[.[]? | select(.name==$NAME)][0].name // empty' "${MV_RUN_DIR}/cloud/delete-gcloud-list.json")"
  mv_expect_eq delete "gcloud cluster-director clusters list no longer contains cluster" "" "${DELETED_CLUSTER}" cloud/delete-gcloud-list.json

  if [[ -n "${LOGIN_VM:-}" ]]; then
    vm_zone="${LOGIN_ZONE:-us-central1-a}"
    if mv_cli delete gcloud-describe-instance -- gcloud compute instances describe "${LOGIN_VM##*/}" --zone="${vm_zone}" --project="${GCP_PROJECT_ID}" --format=json; then
      mv_record delete "login node VM removed from Compute Engine" FAIL cloud/delete-gcloud-describe-instance.txt "instance still exists"
    else
      mv_record delete "login node VM removed from Compute Engine" PASS cloud/delete-gcloud-describe-instance.txt "instance no longer exists"
    fi
  fi
fi

# --- Phase: dependency teardown ----------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  :
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
