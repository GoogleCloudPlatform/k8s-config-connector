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

# lib.sh holds the names and helpers shared by deploy.sh and cleanup.sh.
# It is sourced, not executed.

CONTROLLER_NAMESPACE="krmsyncer-system"
CONTROLLER_DEPLOYMENT="krmsyncer-controller-manager"
KRMSYNCER_CRD="krmsyncers.syncer.gkelabs.io"

# Kubernetes ServiceAccount of the controller (config/rbac/service_account.yaml).
CONTROLLER_KSA="krmsyncer-controller-manager"

# --- Logging -----------------------------------------------------------------

log_info()    { echo -e "\033[0;34m[INFO]\033[0m $*"; }
log_warn()    { echo -e "\033[0;33m[WARN]\033[0m $*" >&2; }
log_success() { echo -e "\033[0;32m[OK]\033[0m $*"; }
log_header()  { echo -e "\n\033[1;36m=== $* ===\033[0m"; }
die()         { echo -e "\033[0;31m[ERROR]\033[0m $*" >&2; exit 1; }

# --- Helpers -----------------------------------------------------------------

# kube_dest runs kubectl against the destination cluster (DEST_CONTEXT).
kube_dest() { kubectl --context="${DEST_CONTEXT}" "$@"; }

# bool_value prints the value of a boolean flag given as --flag=$1: "true" or
# "false". It returns 1 for any other value.
bool_value() {
  case "$1" in
    true|false) printf '%s\n' "$1" ;;
    *) return 1 ;;
  esac
}

sha256_hex() {
  if command -v sha256sum >/dev/null 2>&1; then
    sha256sum | awk '{ print $1 }'
  else
    shasum -a 256 | awk '{ print $1 }'
  fi
}

# location_region prints the region of GKE location $1: zones
# (e.g. us-central1-a) map to their region (us-central1), regions as is.
# deploy.sh and the Makefile use it for the default Artifact Registry location
# of the controller image (--image-location / IMAGE_LOCATION).
location_region() {
  local location="$1"
  if [[ "${location}" =~ ^([a-z]+-[a-z]+[0-9]+)-[a-z]$ ]]; then
    printf '%s\n' "${BASH_REMATCH[1]}"
  else
    printf '%s\n' "${location}"
  fi
}

# gsa_name prints the name of the controller's Google Service Account for
# destination cluster $2 in location $1: krmsyncer-<cluster prefix>-<hash>, at
# most 29 characters (GSA IDs are limited to 30). The hash covers the location
# and name, so clusters with the same name in different locations, or names
# sharing a truncated prefix, get different GSAs.
gsa_name() {
  local location="$1" cluster="$2" hash prefix
  hash="$(printf '%s' "${location}/${cluster}" | sha256_hex)"
  hash="${hash:0:6}"
  prefix="$(printf '%s' "${cluster}" | tr 'A-Z' 'a-z')"
  prefix="${prefix//[^a-z0-9-]/-}"
  prefix="${prefix:0:12}"
  while [[ "${prefix}" == *- ]]; do prefix="${prefix%-}"; done
  if [[ -n "${prefix}" ]]; then
    printf 'krmsyncer-%s-%s\n' "${prefix}" "${hash}"
  else
    printf 'krmsyncer-%s\n' "${hash}"
  fi
}

# gsa_description prints the description deploy.sh sets on the GSA it creates
# for destination cluster $1 (<project>/<location>/<cluster>). Both scripts
# use it to check that a GSA belongs to that cluster.
gsa_description() {
  printf 'KRMSyncer controller for destination cluster %s\n' "$1"
}

# resolve_dest_context sets DEST_CONTEXT to the kubectl context of the
# destination cluster and checks that it exists. If DEST_CONTEXT is empty, it
# defaults to the name `gcloud container clusters get-credentials` creates,
# gke_<project>_<location>_<cluster>. A context set with --dest-context is
# also checked, on a best-effort basis, against the cluster's endpoints, and a
# mismatch is reported as a warning: Connect Gateway or proxied contexts point
# elsewhere legitimately.
resolve_dest_context() {
  local get_credentials="gcloud container clusters get-credentials ${DEST_CLUSTER} --location ${DEST_LOCATION} --project ${PROJECT}"
  if [[ -z "${DEST_CONTEXT}" ]]; then
    DEST_CONTEXT="gke_${PROJECT}_${DEST_LOCATION}_${DEST_CLUSTER}"
    kubectl config get-contexts "${DEST_CONTEXT}" >/dev/null 2>&1 ||
      die "kubectl context \"${DEST_CONTEXT}\" not found; run: ${get_credentials}
(or pass --dest-context if you renamed the context)"
    return
  fi

  kubectl config get-contexts "${DEST_CONTEXT}" >/dev/null 2>&1 ||
    die "kubectl context \"${DEST_CONTEXT}\" (--dest-context) not found in your kubeconfig"

  local server host endpoints
  server="$(kubectl config view --minify --context "${DEST_CONTEXT}" \
    -o jsonpath='{.clusters[0].cluster.server}' 2>/dev/null)" || server=""
  host="${server#*://}"
  host="${host%%/*}"
  host="${host%:*}"
  endpoints="$(gcloud container clusters describe "${DEST_CLUSTER}" \
    --location "${DEST_LOCATION}" --project "${PROJECT}" \
    --format='csv[no-heading](endpoint,privateClusterConfig.privateEndpoint,controlPlaneEndpointsConfig.ipEndpointsConfig.privateEndpoint,controlPlaneEndpointsConfig.dnsEndpointConfig.endpoint)' \
    2>/dev/null)" || endpoints=""
  if [[ -z "${host}" || ",${endpoints}," != *",${host},"* ]]; then
    log_warn "kubectl context \"${DEST_CONTEXT}\" points at ${server:-an unknown server}, which is not an endpoint of"
    log_warn "${PROJECT}/${DEST_LOCATION}/${DEST_CLUSTER}. Make sure it is the destination cluster."
  fi
}
