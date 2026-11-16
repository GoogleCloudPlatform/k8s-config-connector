#!/usr/bin/env bash
# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# deploy.sh installs KRMSyncer into the destination GKE cluster and
# configures it to pull resources from the source GKE cluster:
#
#   1. Run read-only preflight checks, so misconfiguration fails before
#      anything is created: Workload Identity is enabled on the destination
#      cluster, the GSA (if it exists) belongs to it, the source cluster exists
#      and the selected control plane endpoint (Default, DNS or PrivateIP) is
#      available, the destination cluster is reachable and you can create
#      cluster-scoped objects on it, the manifests render, and (unless
#      --skip-build) Docker is running.
#   2. Configure Workload Identity: a Google Service Account (GSA) for the
#      controller, impersonated by the controller's Kubernetes ServiceAccount,
#      with roles/container.viewer on the source cluster's project.
#   3. Build and push the controller image (skipped with --skip-build). The
#      default image goes to an Artifact Registry repository in
#      --image-location (default: the destination cluster's region), created
#      if missing, and the destination cluster's node service accounts get
#      roles/artifactregistry.reader on it.
#   4. Deploy the KRMSyncer CRD, RBAC and controller to the destination cluster.
#   5. Deploy the KRMSyncer configuration (config/templates/krmsyncer.yaml),
#      which references the source cluster in spec.remote.gkeCluster.
#      --sync-mode=status (default) syncs only status, plus spec.resourceID for
#      the GVKs listed in
#      config/templates/service-generated-resource-id-gvks.yaml;
#      --sync-mode=full syncs spec and status of all KCC resources.
#
# Each destination cluster gets its own GSA, krmsyncer-<cluster>-<hash> (see
# gsa_name in lib.sh), so grants to one destination cluster don't leak to
# others.
#
# No credentials are stored in either cluster and nothing is created on the
# source cluster. The controller looks up the source cluster's endpoint and CA
# through the GKE API and authenticates as the GSA. roles/container.viewer
# grants read access to Kubernetes objects in the source project's clusters
# through IAM, so no RBAC objects are needed on the source cluster.
#
# The script deploys to the destination cluster through the kubectl context
# created by `gcloud container clusters get-credentials`
# (gke_<project>_<location>_<cluster>), or the one set with --dest-context.

set -o errexit
set -o nounset
set -o pipefail

# --- Constants ---------------------------------------------------------------

MODULE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEMPLATES_DIR="${MODULE_DIR}/config/templates"

# shellcheck source=lib.sh
source "${MODULE_DIR}/lib.sh"

IMAGE_PLACEHOLDER="controller:latest"

# Artifact Registry Docker repository holding the default controller image.
IMAGE_REPOSITORY="krmsyncer"

# KRMSyncer CR rendered from config/templates/krmsyncer.yaml.
SYNC_CONFIG_NAME="kcc-resource-syncer"

usage() {
  cat <<EOF
Usage: $(basename "$0") --source-cluster <NAME> --source-location <LOCATION> \\
         --dest-cluster <NAME> --dest-location <LOCATION> --project <PROJECT> [flags]

Deploys KRMSyncer to the destination GKE cluster and configures it to pull
resources from the source GKE cluster. The destination cluster needs a kubectl
context (by default the one created by \`gcloud container clusters get-credentials\`)
and Workload Identity enabled.

Flags:
      --source-cluster   Name of the source GKE cluster (required)
      --source-location  Location (region or zone) of the source cluster (required)
      --source-project   GCP project of the source cluster (default: --project)
      --source-endpoint  Control plane endpoint used to reach the source cluster:
                         Default, DNS or PrivateIP (default: Default)
      --dest-cluster     Name of the destination GKE cluster (required)
      --dest-location    Location (region or zone) of the destination cluster (required)
      --dest-context     kubectl context of the destination cluster
                         (default: gke_<project>_<dest-location>_<dest-cluster>)
      --project          GCP project of the destination cluster and the controller's
                         Google Service Account (required)
  -n, --namespace        Namespace for the KRMSyncer CR (default: ${CONTROLLER_NAMESPACE})
      --sync-mode        What the KRMSyncer CR replicates from KCC resources (default: status):
                         status: only status, plus spec.resourceID for the GVKs in
                                 config/templates/service-generated-resource-id-gvks.yaml
                         full:   spec and status
  -i, --image            Controller image (default:
                         <image-location>-docker.pkg.dev/<project>/${IMAGE_REPOSITORY}/controller:latest)
      --image-location   Artifact Registry location of the default image: a region
                         (e.g. us-central1) or multi-region (us, europe, asia)
                         (default: the region of --dest-location); ignored with --image
      --skip-build       Do not build/push the image; deploy --image as is
  -h, --help             Show this help

Example:
  $(basename "$0") --source-cluster src --source-location us-west1 \\
    --dest-cluster dst --dest-location us-central1 --project my-project \\
    --source-endpoint DNS
EOF
}

# --- Flags -------------------------------------------------------------------

SOURCE_CLUSTER=""
SOURCE_LOCATION=""
SOURCE_PROJECT=""
SOURCE_ENDPOINT="Default"
DEST_CLUSTER=""
DEST_LOCATION=""
DEST_CONTEXT=""
PROJECT=""
NAMESPACE="${CONTROLLER_NAMESPACE}"
SYNC_MODE="status"
IMAGE=""
IMAGE_LOCATION=""
SKIP_BUILD="false"

# Set by the preflight checks.
GSA_EXISTS="false"
MANIFESTS=""
SYNC_CONFIG=""
GO_VERSION=""

parse_flags() {
  while [[ $# -gt 0 ]]; do
    local flag="$1"
    case "${flag}" in
      -h|--help) usage; exit 0 ;;
      --skip-build) SKIP_BUILD="true"; shift; continue ;;
      --skip-build=*)
        SKIP_BUILD="$(bool_value "${flag#*=}")" ||
          die "--skip-build takes no value, true or false (got \"${flag#*=}\")"
        shift; continue
        ;;
      --*=*)
        # Support --flag=value by splitting it into --flag value.
        set -- "${flag%%=*}" "${flag#*=}" "${@:2}"
        continue
        ;;
    esac
    [[ $# -ge 2 ]] || die "flag ${flag} requires a value"
    case "${flag}" in
      --source-cluster)  SOURCE_CLUSTER="$2" ;;
      --source-location) SOURCE_LOCATION="$2" ;;
      --source-project)  SOURCE_PROJECT="$2" ;;
      --source-endpoint) SOURCE_ENDPOINT="$2" ;;
      --dest-cluster)    DEST_CLUSTER="$2" ;;
      --dest-location)   DEST_LOCATION="$2" ;;
      --dest-context)    DEST_CONTEXT="$2" ;;
      --project)         PROJECT="$2" ;;
      -n|--namespace)    NAMESPACE="$2" ;;
      --sync-mode)       SYNC_MODE="$2" ;;
      -i|--image)        IMAGE="$2" ;;
      --image-location)  IMAGE_LOCATION="$2" ;;
      *) usage >&2; die "unknown flag: ${flag}" ;;
    esac
    shift 2
  done

  local missing=()
  [[ -n "${SOURCE_CLUSTER}" ]]  || missing+=("--source-cluster")
  [[ -n "${SOURCE_LOCATION}" ]] || missing+=("--source-location")
  [[ -n "${DEST_CLUSTER}" ]]    || missing+=("--dest-cluster")
  [[ -n "${DEST_LOCATION}" ]]   || missing+=("--dest-location")
  [[ -n "${PROJECT}" ]]         || missing+=("--project")
  if [[ ${#missing[@]} -gt 0 ]]; then
    usage >&2
    die "required flag(s) not set: ${missing[*]}"
  fi

  case "${SOURCE_ENDPOINT}" in
    Default|DNS|PrivateIP) ;;
    *) die "--source-endpoint must be Default, DNS or PrivateIP (got \"${SOURCE_ENDPOINT}\")" ;;
  esac

  case "${SYNC_MODE}" in
    status|full) ;;
    *) die "--sync-mode must be status or full (got \"${SYNC_MODE}\")" ;;
  esac
}

check_deps() {
  local deps=(kubectl gcloud)
  [[ "${SKIP_BUILD}" == "true" ]] || deps+=(docker)
  local dep
  for dep in "${deps[@]}"; do
    command -v "${dep}" >/dev/null 2>&1 || die "dependency \"${dep}\" is required but not found"
  done
}

resolve_defaults() {
  [[ -f "${MODULE_DIR}/Dockerfile" && -d "${MODULE_DIR}/config/default" ]] ||
    die "${MODULE_DIR} is not the syncer directory (missing Dockerfile or config/default)"

  [[ -n "${SOURCE_PROJECT}" ]] || SOURCE_PROJECT="${PROJECT}"
  [[ "${SOURCE_PROJECT}/${SOURCE_LOCATION}/${SOURCE_CLUSTER}" != "${PROJECT}/${DEST_LOCATION}/${DEST_CLUSTER}" ]] ||
    die "source and destination are the same cluster (${PROJECT}/${DEST_LOCATION}/${DEST_CLUSTER})"

  resolve_dest_context

  # The default image lives in an Artifact Registry repository in
  # --image-location, by default the destination cluster's region.
  [[ -n "${IMAGE_LOCATION}" ]] || IMAGE_LOCATION="$(location_region "${DEST_LOCATION}")"
  DEFAULT_IMAGE_REPO="${IMAGE_LOCATION}-docker.pkg.dev/${PROJECT}/${IMAGE_REPOSITORY}"
  [[ -n "${IMAGE}" ]] || IMAGE="${DEFAULT_IMAGE_REPO}/controller:latest"

  DEST_ID="${PROJECT}/${DEST_LOCATION}/${DEST_CLUSTER}"
  GSA_NAME="$(gsa_name "${DEST_LOCATION}" "${DEST_CLUSTER}")"
  GSA_EMAIL="${GSA_NAME}@${PROJECT}.iam.gserviceaccount.com"
  WI_MEMBER="serviceAccount:${PROJECT}.svc.id.goog[${CONTROLLER_NAMESPACE}/${CONTROLLER_KSA}]"
}

print_summary() {
  echo "=================================================="
  echo "          KRMSyncer Deployment Setup              "
  echo "=================================================="
  log_info "Source Cluster         : projects/${SOURCE_PROJECT}/locations/${SOURCE_LOCATION}/clusters/${SOURCE_CLUSTER}"
  log_info "Source Endpoint        : ${SOURCE_ENDPOINT}"
  log_info "Dest Cluster Context   : ${DEST_CONTEXT}"
  log_info "GCP Project            : ${PROJECT}"
  log_info "Syncer CR Namespace    : ${NAMESPACE}"
  log_info "Sync Mode              : ${SYNC_MODE}"
  log_info "Controller Namespace   : ${CONTROLLER_NAMESPACE}"
  log_info "Controller Image       : ${IMAGE}"
  log_info "Google Service Account : ${GSA_EMAIL}"
  echo "=================================================="
}

# --- Helpers -----------------------------------------------------------------

# render_template prints a template with ${KEY} placeholders replaced.
# Usage: render_template <file> KEY=value...
render_template() {
  local content
  content="$(<"$1")"
  shift
  local kv
  for kv in "$@"; do
    content="${content//"\${${kv%%=*}}"/"${kv#*=}"}"
  done
  printf '%s\n' "${content}"
}

# print_resource_id_rule prints a KRMSyncer rule that syncs status and
# spec.resourceID for one GVK. Usage: print_resource_id_rule <file> <group> <version> <kind>
print_resource_id_rule() {
  [[ -n "$2" && -n "$3" && -n "$4" ]] ||
    die "$1: incomplete entry (group: \"$2\", version: \"$3\", kind: \"$4\")"
  printf '  - group: %s\n    version: %s\n    kind: %s\n    syncFields: [status, spec.resourceID]\n' "$2" "$3" "$4"
}

# resource_id_rules prints a [status, spec.resourceID] KRMSyncer rule for each
# GVK in the list at $1 (service-generated-resource-id-gvks.yaml). The file is
# a YAML list of {group, version, kind} maps with unquoted values, one key per
# line; anything else is rejected rather than silently skipped.
resource_id_rules() {
  local file="$1" line count=0 group="" version="" kind=""
  local re='^(- |  )(group|version|kind):[[:space:]]*([^[:space:]\"'\'']+)[[:space:]]*$'
  while IFS= read -r line || [[ -n "${line}" ]]; do
    line="${line%$'\r'}"
    line="${line%%#*}"
    [[ -n "${line//[[:space:]]/}" ]] || continue
    [[ "${line}" =~ ${re} ]] || die "${file}: unexpected line: ${line}"
    if [[ "${BASH_REMATCH[1]}" == "- " ]]; then
      [[ ${count} -eq 0 ]] || print_resource_id_rule "${file}" "${group}" "${version}" "${kind}"
      count=$((count + 1))
      group="" version="" kind=""
    elif [[ ${count} -eq 0 ]]; then
      die "${file}: unexpected line before the first entry: ${line}"
    fi
    case "${BASH_REMATCH[2]}" in
      group)   group="${BASH_REMATCH[3]}" ;;
      version) version="${BASH_REMATCH[3]}" ;;
      kind)    kind="${BASH_REMATCH[3]}" ;;
    esac
  done <"${file}"
  [[ ${count} -gt 0 ]] || die "${file}: no GVKs found"
  print_resource_id_rule "${file}" "${group}" "${version}" "${kind}"
}

# render_sync_config renders the KRMSyncer CR for --sync-mode into
# SYNC_CONFIG. In status mode, the resource ID rules are appended after the
# wildcard status-only rule, which must stay first (see
# config/templates/krmsyncer.yaml). It sets a global instead of printing, so
# errors in the command substitutions below abort the script (errexit is not
# inherited by nested command substitutions).
render_sync_config() {
  local template="${TEMPLATES_DIR}/krmsyncer.yaml"
  local gvks="${TEMPLATES_DIR}/service-generated-resource-id-gvks.yaml"
  local sync_fields="[spec, status]"
  [[ "${SYNC_MODE}" == "full" ]] || sync_fields="[status]"
  [[ -f "${template}" ]] || die "${template} not found"

  local config rules=""
  config="$(render_template "${template}" \
    "SYNC_CONFIG_NAME=${SYNC_CONFIG_NAME}" \
    "NAMESPACE=${NAMESPACE}" \
    "SOURCE_PROJECT=${SOURCE_PROJECT}" \
    "SOURCE_LOCATION=${SOURCE_LOCATION}" \
    "SOURCE_CLUSTER=${SOURCE_CLUSTER}" \
    "SOURCE_ENDPOINT=${SOURCE_ENDPOINT}" \
    "SYNC_FIELDS=${sync_fields}")"
  # The appended rules extend spec.rules, so it must end the template.
  [[ "${config}" == *$'\n'"    syncFields: ${sync_fields}" ]] ||
    die "${template} must end with the wildcard rule's syncFields (spec.rules must be the last field)"
  if [[ "${SYNC_MODE}" == "status" ]]; then
    [[ -f "${gvks}" ]] || die "${gvks} not found"
    rules="$(resource_id_rules "${gvks}")"
  fi
  SYNC_CONFIG="${config}"
  [[ -z "${rules}" ]] || SYNC_CONFIG+=$'\n'"${rules}"
}

# ensure_image_repository creates the Artifact Registry Docker repository for
# the default image if it doesn't exist. Custom --image registries are left to
# the user.
ensure_image_repository() {
  [[ "${IMAGE}" == "${DEFAULT_IMAGE_REPO}/"* ]] || return 0
  local rc=0
  gcloud_describe gcloud artifacts repositories describe "${IMAGE_REPOSITORY}" \
    --location "${IMAGE_LOCATION}" --project "${PROJECT}" >/dev/null || rc=$?
  case "${rc}" in
    0)
      log_info "Artifact Registry repository ${DEFAULT_IMAGE_REPO} already exists."
      return
      ;;
    1) ;;
    *) die "cannot describe Artifact Registry repository ${DEFAULT_IMAGE_REPO} (is the Artifact Registry API enabled? gcloud services enable artifactregistry.googleapis.com --project ${PROJECT})" ;;
  esac
  log_info "Creating Artifact Registry repository ${DEFAULT_IMAGE_REPO}..."
  gcloud artifacts repositories create "${IMAGE_REPOSITORY}" \
    --repository-format docker --location "${IMAGE_LOCATION}" --project "${PROJECT}" \
    --description "KRMSyncer controller images" ||
    die "failed to create Artifact Registry repository ${DEFAULT_IMAGE_REPO} (is the Artifact Registry API enabled? gcloud services enable artifactregistry.googleapis.com --project ${PROJECT})"
}

# node_service_accounts prints, one per line, the Google Service Accounts the
# destination cluster's nodes run as: those of its node pools and the node
# auto-provisioning (and Autopilot) default. "default" maps to the project's
# Compute Engine default service account.
node_service_accounts() {
  local accounts sa project_number=""
  accounts="$(
    gcloud container node-pools list --cluster "${DEST_CLUSTER}" \
      --location "${DEST_LOCATION}" --project "${PROJECT}" \
      --format='value(config.serviceAccount)'
    gcloud container clusters describe "${DEST_CLUSTER}" \
      --location "${DEST_LOCATION}" --project "${PROJECT}" \
      --format='value(autoscaling.autoprovisioningNodePoolDefaults.serviceAccount)'
  )" || return 1
  while IFS= read -r sa; do
    [[ -n "${sa}" ]] || continue
    if [[ "${sa}" == "default" ]]; then
      [[ -n "${project_number}" ]] ||
        project_number="$(gcloud projects describe "${PROJECT}" --format='value(projectNumber)')" || return 1
      sa="${project_number}-compute@developer.gserviceaccount.com"
    fi
    printf '%s\n' "${sa}"
  done <<<"${accounts}" | sort -u
}

# grant_image_pull_access grants the destination cluster's node service
# accounts roles/artifactregistry.reader on the default image's repository, so
# nodes can pull the controller image even when they run as a minimal custom
# service account. Custom --image registries are left to the user.
grant_image_pull_access() {
  [[ "${IMAGE}" == "${DEFAULT_IMAGE_REPO}/"* ]] || return 0
  if ! gcloud artifacts repositories describe "${IMAGE_REPOSITORY}" \
      --location "${IMAGE_LOCATION}" --project "${PROJECT}" >/dev/null 2>&1; then
    log_warn "Artifact Registry repository ${DEFAULT_IMAGE_REPO} not found; skipping image pull grants."
    return
  fi

  local accounts
  if ! accounts="$(node_service_accounts)" || [[ -z "${accounts}" ]]; then
    log_warn "Could not determine the node service accounts of ${DEST_CLUSTER}; make sure they have"
    log_warn "roles/artifactregistry.reader on ${DEFAULT_IMAGE_REPO}."
    return
  fi

  local sa
  while IFS= read -r sa; do
    log_info "Granting roles/artifactregistry.reader on ${DEFAULT_IMAGE_REPO} to ${sa}..."
    quiet gcloud artifacts repositories add-iam-policy-binding "${IMAGE_REPOSITORY}" \
      --location "${IMAGE_LOCATION}" --project "${PROJECT}" \
      --member "serviceAccount:${sa}" --role roles/artifactregistry.reader \
      --condition=None --quiet ||
      log_warn "failed to grant roles/artifactregistry.reader to ${sa}; nodes may fail to pull ${IMAGE}"
  done <<<"${accounts}"
  log_success "Node image pull access configured."
}

# quiet runs a command, printing its output only if it fails.
quiet() {
  local out
  if ! out="$("$@" 2>&1)"; then
    printf '%s\n' "${out}" >&2
    return 1
  fi
}

# gcloud_describe runs a gcloud describe command and prints its stdout. It
# returns 1 if the resource doesn't exist (NOT_FOUND), and prints the error and
# returns 2 on any other failure, so permission, auth or network errors aren't
# mistaken for a missing resource.
gcloud_describe() {
  local out err err_file rc=0
  err_file="$(mktemp)"
  out="$("$@" 2>"${err_file}")" || rc=$?
  err="$(<"${err_file}")"
  rm -f "${err_file}"
  if [[ ${rc} -eq 0 ]]; then
    printf '%s\n' "${out}"
    return 0
  fi
  [[ "${err}" != *NOT_FOUND* ]] || return 1
  printf '%s\n' "${err}" >&2
  return 2
}

# retry_gsa_not_found runs a command, retrying for up to a minute while it fails
# because a Google Service Account doesn't exist yet. GSA creation is eventually
# consistent, so IAM calls right after `gcloud iam service-accounts create` can
# fail with "does not exist". Output is printed only if the command fails.
retry_gsa_not_found() {
  local attempt out
  for attempt in 1 2 3 4 5 6 7; do
    out="$("$@" 2>&1)" && return 0
    case "${out}" in
      *"does not exist"*|*NOT_FOUND*) ;;
      *) break ;;
    esac
    [[ ${attempt} -lt 7 ]] || break
    log_info "  Service account not visible yet; retrying in 10s..."
    sleep 10
  done
  printf '%s\n' "${out}" >&2
  return 1
}

# --- Preflight checks --------------------------------------------------------
#
# Read-only checks that run before anything is created or changed, so a
# misconfiguration fails fast instead of leaving a partially provisioned setup.

check_workload_identity() {
  log_info "Checking Workload Identity on destination cluster ${DEST_CLUSTER}..."
  local pool
  pool="$(gcloud container clusters describe "${DEST_CLUSTER}" \
    --location "${DEST_LOCATION}" --project "${PROJECT}" \
    --format='value(workloadIdentityConfig.workloadPool)')"
  [[ -n "${pool}" ]] || die "Workload Identity is not enabled on ${DEST_CLUSTER}; enable it with:
  gcloud container clusters update ${DEST_CLUSTER} --location ${DEST_LOCATION} --project ${PROJECT} --workload-pool=${PROJECT}.svc.id.goog
(existing node pools also need --workload-metadata=GKE_METADATA)"
  log_success "Workload Identity enabled (pool: ${pool})."
}

# check_gsa sets GSA_EXISTS and fails if the GSA exists but belongs to another
# destination cluster.
check_gsa() {
  log_info "Checking Google Service Account ${GSA_EMAIL}..."
  local description rc=0
  description="$(gcloud_describe gcloud iam service-accounts describe "${GSA_EMAIL}" \
    --project "${PROJECT}" --format='value(description)')" || rc=$?
  case "${rc}" in
    0)
      [[ "${description}" == "$(gsa_description "${DEST_ID}")" ]] ||
        die "${GSA_EMAIL} exists but was not created by this script for ${DEST_ID} (description: \"${description}\")"
      GSA_EXISTS="true"
      ;;
    1) GSA_EXISTS="false" ;;
    *) die "cannot describe Google Service Account ${GSA_EMAIL}" ;;
  esac
}

# print_reachability_hint prints a command that tests, from inside the
# destination cluster, whether the source endpoint at $1 is reachable. IP
# endpoints depend on VPC routing, firewall rules and authorized networks,
# which the script can't verify from outside the cluster.
print_reachability_hint() {
  log_info "IP endpoints need a network path from the destination cluster (VPC routing, firewall,"
  log_info "authorized networks, and control plane global access if the clusters are in different regions)."
  log_info "To verify it, run:"
  log_info "  kubectl --context=${DEST_CONTEXT} run nettest --rm -i --restart=Never --image=busybox -- nc -zv -w 10 ${1} 443"
  log_info "\"open\" means there is a network path; a timeout means traffic is dropped (routing, firewall or authorized networks)."
}

# check_source_cluster verifies that the source cluster exists and exposes the
# selected endpoint, so misconfiguration fails here instead of in the controller.
check_source_cluster() {
  log_info "Looking up source cluster ${SOURCE_CLUSTER} in ${SOURCE_PROJECT}/${SOURCE_LOCATION}..."
  local described
  described="$(gcloud container clusters describe "${SOURCE_CLUSTER}" \
    --location "${SOURCE_LOCATION}" --project "${SOURCE_PROJECT}" \
    --format='csv[no-heading](endpoint,privateClusterConfig.privateEndpoint,controlPlaneEndpointsConfig.ipEndpointsConfig.privateEndpoint,controlPlaneEndpointsConfig.dnsEndpointConfig.endpoint,controlPlaneEndpointsConfig.dnsEndpointConfig.allowExternalTraffic)')" ||
    die "cannot describe source cluster ${SOURCE_CLUSTER} in ${SOURCE_PROJECT}/${SOURCE_LOCATION}"

  local endpoint private_legacy private_ip dns_endpoint dns_external
  # CSV keeps empty fields in place (whitespace-separated output would collapse them).
  IFS=',' read -r endpoint private_legacy private_ip dns_endpoint dns_external <<<"${described}"
  [[ -n "${private_ip}" ]] || private_ip="${private_legacy}"

  local update="gcloud container clusters update ${SOURCE_CLUSTER} --location ${SOURCE_LOCATION} --project ${SOURCE_PROJECT}"
  case "${SOURCE_ENDPOINT}" in
    Default)
      [[ -n "${endpoint}" ]] || die "source cluster has no IP endpoint (IP access may be disabled); use --source-endpoint DNS"
      log_info "Source endpoint (Default): ${endpoint}"
      print_reachability_hint "${endpoint}"
      ;;
    PrivateIP)
      [[ -n "${private_ip}" ]] || die "source cluster has no private endpoint; use --source-endpoint DNS or Default"
      log_info "Source endpoint (PrivateIP): ${private_ip}. The controller must be able to reach the source cluster's VPC."
      print_reachability_hint "${private_ip}"
      ;;
    DNS)
      [[ -n "${dns_endpoint}" && "${dns_external}" == [Tt]rue ]] ||
        die "source cluster's DNS endpoint is not enabled for external traffic; enable it with:
  ${update} --enable-dns-access
or use --source-endpoint Default or PrivateIP"
      log_info "Source endpoint (DNS): ${dns_endpoint}"
      ;;
  esac
  log_success "Source cluster found."
}

# check_dest_cluster checks that the destination cluster is reachable and that
# you can create the cluster-scoped objects of the controller.
check_dest_cluster() {
  log_info "Checking access to destination cluster ${DEST_CONTEXT}..."
  kube_dest get --raw /version >/dev/null ||
    die "cannot reach destination cluster with context \"${DEST_CONTEXT}\""
  local resource
  for resource in customresourcedefinitions.apiextensions.k8s.io \
      clusterroles.rbac.authorization.k8s.io clusterrolebindings.rbac.authorization.k8s.io; do
    [[ "$(kube_dest auth can-i create "${resource}" 2>/dev/null)" == "yes" ]] ||
      die "you cannot create ${resource} in the destination cluster (roles/container.admin or cluster-admin is needed)"
  done
}

# render_manifests renders the controller manifests into MANIFESTS and the
# KRMSyncer CR into SYNC_CONFIG.
render_manifests() {
  local config_dir="${MODULE_DIR}/config/default"
  log_info "Rendering manifests from ${config_dir}..."
  MANIFESTS="$(kubectl kustomize "${config_dir}")"
  [[ "${MANIFESTS}" == *"image: ${IMAGE_PLACEHOLDER}"* ]] ||
    die "image \"${IMAGE_PLACEHOLDER}\" not found in rendered manifests; check config/manager/manager.yaml"
  render_sync_config
}

# check_build sets GO_VERSION and checks that Docker can build the image.
check_build() {
  # Build with the toolchain directive, the Go release the syncer is built with. The go
  # directive is only the minimum version for importers, so it's the fallback.
  GO_VERSION="$(awk '$1 == "toolchain" && NF == 2 { sub(/^go/, "", $2); print $2 }' "${MODULE_DIR}/go.mod")"
  if [[ -z "${GO_VERSION}" ]]; then
    GO_VERSION="$(awk '$1 == "go" && NF == 2 { print $2 }' "${MODULE_DIR}/go.mod")"
  fi
  [[ -n "${GO_VERSION}" ]] || die "no toolchain or go directive found in ${MODULE_DIR}/go.mod"
  docker info >/dev/null 2>&1 || die "cannot connect to the Docker daemon; is it running?"
}

preflight_checks() {
  check_workload_identity
  check_gsa
  check_source_cluster
  check_dest_cluster
  render_manifests
  [[ "${SKIP_BUILD}" == "true" ]] || check_build
  log_success "Preflight checks passed."
}

# --- Steps -------------------------------------------------------------------

configure_workload_identity() {
  if [[ "${GSA_EXISTS}" == "true" ]]; then
    log_info "Google Service Account ${GSA_EMAIL} already exists."
  else
    log_info "Creating Google Service Account ${GSA_EMAIL}..."
    gcloud iam service-accounts create "${GSA_NAME}" --project "${PROJECT}" \
      --display-name "KRMSyncer controller (${DEST_CLUSTER})" \
      --description "$(gsa_description "${DEST_ID}")"
  fi

  # roles/container.viewer includes container.clusters.get (look up the source
  # cluster's endpoint and CA, and authenticate to it), container.clusters.connect
  # (needed for the DNS endpoint) and read access to Kubernetes objects in the
  # project's clusters, so no RBAC objects are needed on the source cluster.
  log_info "Granting roles/container.viewer on project ${SOURCE_PROJECT} to ${GSA_EMAIL}..."
  retry_gsa_not_found gcloud projects add-iam-policy-binding "${SOURCE_PROJECT}" \
    --member "serviceAccount:${GSA_EMAIL}" --role roles/container.viewer \
    --condition=None --quiet

  log_info "Allowing ${WI_MEMBER} to impersonate ${GSA_EMAIL}..."
  retry_gsa_not_found gcloud iam service-accounts add-iam-policy-binding "${GSA_EMAIL}" --project "${PROJECT}" \
    --member "${WI_MEMBER}" --role roles/iam.workloadIdentityUser \
    --condition=None --quiet
  log_success "Workload Identity configured."
}

build_and_push_image() {
  if [[ "${SKIP_BUILD}" == "true" ]]; then
    log_info "Skipping build; deploying ${IMAGE} as is."
    return
  fi

  log_info "Building controller image ${IMAGE} (Go ${GO_VERSION})..."
  docker build --platform linux/amd64 \
    --build-arg "GO_VERSION=${GO_VERSION}" -t "${IMAGE}" "${MODULE_DIR}"
  ensure_image_repository
  log_info "Pushing controller image ${IMAGE}..."
  docker push "${IMAGE}" ||
    die "failed to push ${IMAGE} (you may need to run: gcloud auth configure-docker ${IMAGE%%/*})"
  log_success "Controller image pushed."
}

deploy_controller() {
  log_info "Applying CRD, RBAC and controller Deployment..."
  printf '%s\n' "${MANIFESTS//"image: ${IMAGE_PLACEHOLDER}"/"image: ${IMAGE}"}" | kube_dest apply -f -

  log_info "Annotating ServiceAccount for Workload Identity..."
  kube_dest -n "${CONTROLLER_NAMESPACE}" annotate serviceaccount "${CONTROLLER_KSA}" \
    "iam.gke.io/gcp-service-account=${GSA_EMAIL}" --overwrite

  log_info "Waiting for KRMSyncer CRD to be established..."
  kube_dest wait --for condition=established --timeout=60s "crd/${KRMSYNCER_CRD}"

  # Restart so pods pick up the ServiceAccount annotation and a newly pushed image.
  log_info "Restarting controller..."
  kube_dest -n "${CONTROLLER_NAMESPACE}" rollout restart deployment "${CONTROLLER_DEPLOYMENT}"
  kube_dest -n "${CONTROLLER_NAMESPACE}" rollout status deployment "${CONTROLLER_DEPLOYMENT}" --timeout=180s
  log_success "KRMSyncer controller is running."
}

deploy_sync_config() {
  log_info "Applying KRMSyncer ${NAMESPACE}/${SYNC_CONFIG_NAME} (--sync-mode=${SYNC_MODE})..."
  kube_dest create namespace "${NAMESPACE}" --dry-run=client -o yaml | kube_dest apply -f - >/dev/null
  printf '%s\n' "${SYNC_CONFIG}" | kube_dest apply -f -
  log_success "KRMSyncer configuration applied."
}

# --- Main --------------------------------------------------------------------

main() {
  parse_flags "$@"
  check_deps
  resolve_defaults
  print_summary

  log_header "Step 1: Preflight Checks"
  preflight_checks
  # Workload Identity runs before the image build: new IAM bindings can take a
  # few minutes to propagate, and the build gives them time.
  log_header "Step 2: Configure Workload Identity"
  configure_workload_identity
  log_header "Step 3: Build and Push Controller Image"
  build_and_push_image
  grant_image_pull_access
  log_header "Step 4: Deploy KRMSyncer Controller"
  deploy_controller
  log_header "Step 5: Deploy KRMSyncer Configuration"
  deploy_sync_config

  echo
  log_info "View controller logs with:"
  log_info "  kubectl --context=${DEST_CONTEXT} -n ${CONTROLLER_NAMESPACE} logs deploy/${CONTROLLER_DEPLOYMENT} -f"
  log_info "New IAM bindings can take a few minutes to propagate. If the logs show auth or permission"
  log_info "errors right after setup, the controller retries with backoff; to retry immediately, run:"
  log_info "  kubectl --context=${DEST_CONTEXT} -n ${CONTROLLER_NAMESPACE} rollout restart deploy/${CONTROLLER_DEPLOYMENT}"
  log_info "If ${GSA_EMAIL} was granted access to a previous source project, revoke it with:"
  log_info "  ${MODULE_DIR}/cleanup.sh --dest-cluster ${DEST_CLUSTER} --dest-location ${DEST_LOCATION} --project ${PROJECT} --dest-context ${DEST_CONTEXT} --source-project <OLD_SOURCE_PROJECT>"
}

main "$@"
