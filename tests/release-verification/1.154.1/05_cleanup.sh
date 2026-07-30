#!/usr/bin/env bash
set -eo pipefail

# Script: 05_cleanup.sh
# Purpose: Delete test resources and optionally tear down GKE cluster.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PRE_UPGRADE_MANIFEST="${SCRIPT_DIR}/manifests/pre-upgrade/01-bugfixes-and-base.yaml"
POST_UPGRADE_MANIFEST="${SCRIPT_DIR}/manifests/post-upgrade/02-newfields-and-updates.yaml"

PROJECT_ID="${PROJECT_ID:-$(gcloud config get-value project 2>/dev/null)}"
CLUSTER_NAME="${CLUSTER_NAME:-kcc-release-1154-verify}"
ZONE="${ZONE:-us-central1-a}"

echo "============================================================"
echo "Cleaning up KCC Release 1.154 Verification Resources"
echo "============================================================"

if [[ -f "${POST_UPGRADE_MANIFEST}" ]]; then
  echo "Deleting post-upgrade test resources..."
  kubectl delete -f "${POST_UPGRADE_MANIFEST}" --ignore-not-found=true || true
fi

if [[ -f "${PRE_UPGRADE_MANIFEST}" ]]; then
  echo "Deleting pre-upgrade test resources..."
  kubectl delete -f "${PRE_UPGRADE_MANIFEST}" --ignore-not-found=true || true
fi

echo ""
echo "Test resources deleted from Kubernetes."

if [[ "${DELETE_CLUSTER}" == "true" ]]; then
  echo "Deleting GKE cluster '${CLUSTER_NAME}'..."
  gcloud container clusters delete "${CLUSTER_NAME}" --zone="${ZONE}" --project="${PROJECT_ID}" --quiet
  echo "Cluster '${CLUSTER_NAME}' deleted."
else
  echo ""
  echo "Note: GKE cluster '${CLUSTER_NAME}' was preserved."
  echo "To delete the cluster, run:"
  echo "  export DELETE_CLUSTER=true"
  echo "  ./05_cleanup.sh"
fi
