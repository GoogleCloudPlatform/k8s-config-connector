#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/params.env"

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "${REPO_ROOT}"

echo "===> [Teardown 1] Deleting test KRM resources and test namespace..."
kubectl delete pubsubtopic "${TEST_TOPIC_NAME}" -n "${TEST_NAMESPACE}" --ignore-not-found=true --timeout=120s || true
kubectl delete namespace "${TEST_NAMESPACE}" --ignore-not-found=true || true

echo "===> [Teardown 2] Deleting Config Connector manifests and CRDs..."
kubectl kustomize config/installbundle/releases/scopes/cluster/withworkloadidentity | \
    sed -e "s/cnrm-system@\${PROJECT_ID?}\.iam\.gserviceaccount\.com/${GSA_EMAIL}/g" | \
    sed -e "s/\${PROJECT_ID?}/${PROJECT}/g" | \
    kubectl delete -f - --ignore-not-found=true || true

kubectl delete -f operator/config/crd/bases/ --ignore-not-found=true || true
kubectl delete -f config/crds/resources/ --ignore-not-found=true || true

echo "===> [Teardown 3] Deleting GKE Cluster..."
gcloud container clusters delete "${CLUSTER_NAME}" \
    --zone="${ZONE}" \
    --project="${PROJECT}" \
    --quiet || true

echo "===> [Teardown 4] Removing IAM policy bindings and deleting GSA..."
gcloud projects remove-iam-policy-binding "${PROJECT}" \
    --member="serviceAccount:${GSA_EMAIL}" \
    --role="roles/owner" \
    --quiet || true

gcloud iam service-accounts delete "${GSA_EMAIL}" \
    --project="${PROJECT}" \
    --quiet || true

echo "===> [Teardown 5] Deleting Artifact Registry repository..."
gcloud artifacts repositories delete "${ARTIFACT_REPO}" \
    --location="${REGION}" \
    --project="${PROJECT}" \
    --quiet || true

echo "===> [Teardown 6] Cleaning up local generated patch files..."
rm -f config/installbundle/components/manager/base/manager_image_patch.yaml
rm -f config/installbundle/components/recorder/recorder_image_patch.yaml
rm -f config/installbundle/components/webhook/webhook_image_patch.yaml
rm -f config/installbundle/components/deletiondefender/deletiondefender_image_patch.yaml
rm -f config/installbundle/components/unmanageddetector/unmanageddetector_image_patch.yaml

echo "===> Teardown completed successfully!"
