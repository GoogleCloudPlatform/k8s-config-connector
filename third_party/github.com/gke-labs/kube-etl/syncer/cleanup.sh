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

# cleanup.sh revokes roles/container.viewer grants that deploy.sh gave a
# destination cluster's Google Service Account (GSA) on source projects the
# cluster no longer syncs from. deploy.sh only adds IAM bindings, so re-running
# it with a different --source-project leaves the old grant in place.
#
#   1. Find the destination cluster's GSA, named as in deploy.sh (gsa_name in
#      lib.sh).
#   2. Find the source projects still in use: spec.remote.gkeCluster.project of
#      every KRMSyncer CR in the destination cluster.
#   3. Revoke the GSA's roles/container.viewer binding on each --source-project
#      that is not in use and still has the binding (use --dry-run to preview).
#
# Only the GSA's own unconditional roles/container.viewer binding is removed;
# other members of the role keep their access. The script refuses to touch a
# GSA that deploy.sh did not create for this destination cluster, since
# another cluster might depend on its grants.

set -o errexit
set -o nounset
set -o pipefail

# --- Constants ---------------------------------------------------------------

MODULE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# shellcheck source=lib.sh
source "${MODULE_DIR}/lib.sh"

VIEWER_ROLE="roles/container.viewer"

usage() {
  cat <<EOF
Usage: $(basename "$0") --dest-cluster <NAME> --dest-location <LOCATION> --project <PROJECT> \\
         --source-project <PROJECT> [--source-project <PROJECT>...] [flags]

Revokes roles/container.viewer on source projects that the destination
cluster's KRMSyncer GSA was granted by deploy.sh but no longer syncs from.
Projects referenced by a KRMSyncer CR in the destination cluster are never
revoked. The destination cluster needs a kubectl context (by default the one
created by \`gcloud container clusters get-credentials\`).

Flags:
      --dest-cluster     Name of the destination GKE cluster (required)
      --dest-location    Location (region or zone) of the destination cluster (required)
      --dest-context     kubectl context of the destination cluster
                         (default: gke_<project>_<dest-location>_<dest-cluster>)
      --project          GCP project of the destination cluster and the GSA (required)
      --source-project   Previous source project to revoke (required; repeatable)
      --dry-run          Print what would be revoked, without changing anything
  -h, --help             Show this help

Example:
  $(basename "$0") --dest-cluster dst --dest-location us-central1 --project my-project \\
    --source-project old-source-project
EOF
}

# --- Flags -------------------------------------------------------------------

DEST_CLUSTER=""
DEST_LOCATION=""
DEST_CONTEXT=""
PROJECT=""
CANDIDATES=()
DRY_RUN="false"

parse_flags() {
  while [[ $# -gt 0 ]]; do
    local flag="$1"
    case "${flag}" in
      -h|--help) usage; exit 0 ;;
      --dry-run) DRY_RUN="true"; shift; continue ;;
      --dry-run=*)
        DRY_RUN="$(bool_value "${flag#*=}")" ||
          die "--dry-run takes no value, true or false (got \"${flag#*=}\")"
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
      --dest-cluster)   DEST_CLUSTER="$2" ;;
      --dest-location)  DEST_LOCATION="$2" ;;
      --dest-context)   DEST_CONTEXT="$2" ;;
      --project)        PROJECT="$2" ;;
      --source-project) CANDIDATES+=("$2") ;;
      *) usage >&2; die "unknown flag: ${flag}" ;;
    esac
    shift 2
  done

  local missing=()
  [[ -n "${DEST_CLUSTER}" ]]  || missing+=("--dest-cluster")
  [[ -n "${DEST_LOCATION}" ]] || missing+=("--dest-location")
  [[ -n "${PROJECT}" ]]       || missing+=("--project")
  [[ ${#CANDIDATES[@]} -gt 0 ]] || missing+=("--source-project")
  if [[ ${#missing[@]} -gt 0 ]]; then
    usage >&2
    die "required flag(s) not set: ${missing[*]}"
  fi
}

check_deps() {
  local dep
  for dep in kubectl gcloud; do
    command -v "${dep}" >/dev/null 2>&1 || die "dependency \"${dep}\" is required but not found"
  done
}

resolve_defaults() {
  DEST_ID="${PROJECT}/${DEST_LOCATION}/${DEST_CLUSTER}"
  resolve_dest_context
  kube_dest get --raw /version >/dev/null ||
    die "cannot reach destination cluster with context \"${DEST_CONTEXT}\""
}

# --- Helpers -----------------------------------------------------------------

# contains returns whether $1 is one of the remaining arguments.
contains() {
  local needle="$1" item
  shift
  for item in "$@"; do
    [[ "${item}" != "${needle}" ]] || return 0
  done
  return 1
}

# has_viewer_binding returns whether the GSA has an unconditional
# roles/container.viewer binding on project $1.
has_viewer_binding() {
  local members
  members="$(gcloud projects get-iam-policy "$1" \
    --flatten='bindings[].members' \
    --filter="bindings.role=${VIEWER_ROLE} AND -bindings.condition:*" \
    --format='value(bindings.members)')" || die "cannot read the IAM policy of project $1"
  grep -qxF "serviceAccount:${GSA_EMAIL}" <<<"${members}"
}

# --- Steps -------------------------------------------------------------------

resolve_gsa() {
  GSA_EMAIL="$(gsa_name "${DEST_LOCATION}" "${DEST_CLUSTER}")@${PROJECT}.iam.gserviceaccount.com"
  log_info "Google Service Account: ${GSA_EMAIL}"

  local description
  description="$(gcloud iam service-accounts describe "${GSA_EMAIL}" --project "${PROJECT}" \
    --format='value(description)')" || die "cannot describe ${GSA_EMAIL}"
  [[ "${description}" == "$(gsa_description "${DEST_ID}")" ]] ||
    die "${GSA_EMAIL} was not created by deploy.sh for ${DEST_ID} (description: \"${description}\"); other clusters may depend on its grants"
}

find_projects_in_use() {
  IN_USE=()
  if ! kube_dest get crd "${KRMSYNCER_CRD}" >/dev/null 2>&1; then
    log_warn "The KRMSyncer CRD is not installed in ${DEST_CONTEXT}; no source project is in use."
    return
  fi
  local projects
  projects="$(kube_dest get krmsyncers --all-namespaces \
    -o jsonpath='{range .items[*]}{.spec.remote.gkeCluster.project}{"\n"}{end}')" ||
    die "cannot list KRMSyncer CRs in ${DEST_CONTEXT}"
  local project
  while IFS= read -r project; do
    [[ -z "${project}" ]] || contains "${project}" ${IN_USE[@]+"${IN_USE[@]}"} || IN_USE+=("${project}")
  done <<<"${projects}"
  log_info "Source projects in use: ${IN_USE[*]:-(none)}"
}

find_candidates() {
  TO_REVOKE=()
  local project
  for project in "${CANDIDATES[@]}"; do
    if contains "${project}" ${IN_USE[@]+"${IN_USE[@]}"}; then
      log_info "Keeping ${project}: a KRMSyncer CR still syncs from it."
    elif has_viewer_binding "${project}"; then
      TO_REVOKE+=("${project}")
    else
      log_info "Nothing to do in ${project}: no ${VIEWER_ROLE} binding for ${GSA_EMAIL}."
    fi
  done
}

revoke() {
  if [[ ${#TO_REVOKE[@]} -eq 0 ]]; then
    log_success "No stale grants."
    return
  fi

  log_info "Will revoke ${VIEWER_ROLE} from ${GSA_EMAIL} on: ${TO_REVOKE[*]}"
  if [[ "${DRY_RUN}" == "true" ]]; then
    log_info "Dry run; nothing changed."
    return
  fi

  local project failed=0
  for project in "${TO_REVOKE[@]}"; do
    if gcloud projects remove-iam-policy-binding "${project}" \
        --member "serviceAccount:${GSA_EMAIL}" --role "${VIEWER_ROLE}" \
        --condition=None --quiet >/dev/null; then
      log_success "Revoked ${VIEWER_ROLE} on ${project}."
    else
      log_warn "Could not revoke ${VIEWER_ROLE} on ${project}."
      failed=1
    fi
  done
  [[ ${failed} -eq 0 ]] || die "some grants could not be revoked"
}

# --- Main --------------------------------------------------------------------

main() {
  parse_flags "$@"
  check_deps
  resolve_defaults
  resolve_gsa
  find_projects_in_use
  find_candidates
  revoke
}

main "$@"
