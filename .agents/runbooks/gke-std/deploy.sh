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
# limitations under the License

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/params.env"

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "${REPO_ROOT}"
RUN_DIR="${SCRIPT_DIR}"

echo "===> [Step 1] Creating Artifact Registry repository..."
if ! gcloud artifacts repositories describe "${ARTIFACT_REPO}" --location="${REGION}" --project="${PROJECT}" >/dev/null 2>&1; then
    gcloud artifacts repositories create "${ARTIFACT_REPO}" \
        --repository-format=docker \
        --location="${REGION}" \
        --description="KCC source build images for ${RESOURCE_PREFIX}" \
        --labels="repo-agent-instance=${RESOURCE_PREFIX}" \
        --project="${PROJECT}"
fi

echo "===> [Step 2] Creating Google Service Account and assigning project IAM permissions..."
if ! gcloud iam service-accounts describe "${GSA_EMAIL}" --project="${PROJECT}" >/dev/null 2>&1; then
    gcloud iam service-accounts create "${GSA_NAME}" \
        --display-name="${GSA_NAME}" \
        --description="GSA for Config Connector in ${CLUSTER_NAME}" \
        --project="${PROJECT}"
fi

gcloud projects add-iam-policy-binding "${PROJECT}" \
    --member="serviceAccount:${GSA_EMAIL}" \
    --role="roles/owner" \
    --quiet

PROJECT_NUMBER="$(gcloud projects describe "${PROJECT}" --format="value(projectNumber)")"
CB_SA="${PROJECT_NUMBER}-compute@developer.gserviceaccount.com"
gcloud projects add-iam-policy-binding "${PROJECT}" \
    --member="serviceAccount:${CB_SA}" \
    --role="roles/storage.admin" \
    --quiet
gcloud projects add-iam-policy-binding "${PROJECT}" \
    --member="serviceAccount:${CB_SA}" \
    --role="roles/artifactregistry.writer" \
    --quiet
gcloud projects add-iam-policy-binding "${PROJECT}" \
    --member="serviceAccount:${CB_SA}" \
    --role="roles/logging.logWriter" \
    --quiet

gcloud storage buckets add-iam-policy-binding "gs://${PROJECT}_cloudbuild" \
    --member="serviceAccount:${CB_SA}" \
    --role="roles/storage.admin" \
    --quiet 2>/dev/null || true

echo "===> [Step 3] Creating GKE Cluster with Workload Identity..."
if ! gcloud container clusters describe "${CLUSTER_NAME}" --zone="${ZONE}" --project="${PROJECT}" >/dev/null 2>&1; then
    gcloud container clusters create "${CLUSTER_NAME}" \
        --zone="${ZONE}" \
        --project="${PROJECT}" \
        --workload-pool="${PROJECT}.svc.id.goog" \
        --num-nodes=3 \
        --machine-type=e2-standard-4 \
        --labels="repo-agent-instance=${RESOURCE_PREFIX}"
fi

gcloud container clusters get-credentials "${CLUSTER_NAME}" \
    --zone="${ZONE}" \
    --project="${PROJECT}"

echo "===> [Step 4] Configuring Workload Identity IAM binding..."
gcloud iam service-accounts add-iam-policy-binding "${GSA_EMAIL}" \
    --member="serviceAccount:${PROJECT}.svc.id.goog[cnrm-system/cnrm-controller-manager]" \
    --role="roles/iam.workloadIdentityUser" \
    --project="${PROJECT}"

echo "===> [Step 5] Building container images from source via Cloud Build..."
if ! gcloud artifacts docker images list "${IMAGE_PREFIX%/}" --include-tags --filter="TAGS:${IMAGE_TAG}" --format="value(TAGS)" 2>/dev/null | grep -q "${IMAGE_TAG}"; then
    gcloud builds submit \
        --project="${PROJECT}" \
        --ignore-file="${RUN_DIR}/.gcloudignore" \
        --config="${RUN_DIR}/cloudbuild.yaml" \
        --substitutions="_IMAGE_PREFIX=${IMAGE_PREFIX},_IMAGE_TAG=${IMAGE_TAG}" \
        .
else
    echo "Images with tag ${IMAGE_TAG} already exist in Artifact Registry repository ${ARTIFACT_REPO}. Skipping build."
fi

echo "===> [Step 6] Generating Kustomize image patches..."
cp config/installbundle/components/manager/base/manager_image_patch_template.yaml config/installbundle/components/manager/base/manager_image_patch.yaml
sed -i -e "s@image: .*@image: ${IMAGE_PREFIX}controller:${IMAGE_TAG}@" config/installbundle/components/manager/base/manager_image_patch.yaml

cat <<EOF > config/installbundle/components/recorder/recorder_image_patch.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: resource-stats-recorder
spec:
  template:
    spec:
      containers:
      - image: ${IMAGE_PREFIX}recorder:${IMAGE_TAG}
        name: recorder
        resources:
          limits:
            memory: 512Mi
          requests:
            cpu: 100m
            memory: 512Mi
EOF

cat <<EOF > config/installbundle/components/webhook/webhook_image_patch.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webhook-manager
spec:
  template:
    spec:
      containers:
      - image: ${IMAGE_PREFIX}webhook:${IMAGE_TAG}
        name: webhook
        resources:
          limits:
            memory: 512Mi
          requests:
            cpu: 250m
            memory: 512Mi
        env:
        - name: GOMEMLIMIT
          value: 460MiB
EOF

cp config/installbundle/components/deletiondefender/deletiondefender_image_patch_template.yaml config/installbundle/components/deletiondefender/deletiondefender_image_patch.yaml
sed -i -e "s@image: .*@image: ${IMAGE_PREFIX}deletiondefender:${IMAGE_TAG}@" config/installbundle/components/deletiondefender/deletiondefender_image_patch.yaml

cp config/installbundle/components/unmanageddetector/unmanageddetector_image_patch_template.yaml config/installbundle/components/unmanageddetector/unmanageddetector_image_patch.yaml
sed -i -e "s@image: .*@image: ${IMAGE_PREFIX}unmanageddetector:${IMAGE_TAG}@" config/installbundle/components/unmanageddetector/unmanageddetector_image_patch.yaml

echo "===> [Step 7] Installing Config Connector CRDs..."
kubectl apply -f operator/config/crd/bases/
kubectl apply -f config/crds/resources/

echo "===> [Step 8] Deploying Config Connector controller and system manifests..."
kubectl kustomize config/installbundle/releases/scopes/cluster/withworkloadidentity | \
    sed -e "s/cnrm-system@\${PROJECT_ID?}\.iam\.gserviceaccount\.com/${GSA_EMAIL}/g" | \
    sed -e "s/\${PROJECT_ID?}/${PROJECT}/g" | \
    kubectl apply -f -

echo "===> [Step 9] Configuring managed test namespace..."
kubectl create namespace "${TEST_NAMESPACE}" --dry-run=client -o yaml | kubectl apply -f -
kubectl annotate namespace "${TEST_NAMESPACE}" "cnrm.cloud.google.com/project-id=${PROJECT}" --overwrite

echo "===> [Verify 1] Waiting for cnrm-system pods to become Ready..."
kubectl wait -n cnrm-system --for=condition=Ready pod --all --timeout=300s
kubectl get pods -n cnrm-system

echo "===> [Verify 2] Creating test PubSubTopic resource..."
cat <<EOF | kubectl apply -f -
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  name: "${TEST_TOPIC_NAME}"
  namespace: "${TEST_NAMESPACE}"
  labels:
    repo-agent-instance: "${RESOURCE_PREFIX}"
EOF

echo "===> [Verify 3] Waiting for PubSubTopic reconciliation (Ready: True)..."
kubectl wait -n "${TEST_NAMESPACE}" --for=condition=Ready "pubsubtopic/${TEST_TOPIC_NAME}" --timeout=180s
kubectl get pubsubtopic "${TEST_TOPIC_NAME}" -n "${TEST_NAMESPACE}" -o yaml

echo "===> [Verify 4] Verifying PubSubTopic in Google Cloud Platform..."
gcloud pubsub topics describe "${TEST_TOPIC_NAME}" --project="${PROJECT}"

echo "===> Deploy and verification completed successfully!"
