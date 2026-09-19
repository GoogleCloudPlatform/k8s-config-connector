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
set -eo pipefail

# Script: 01_install_old_version.sh
# Purpose: Download and install Config Connector version 1.153.0 (old release).

OLD_VERSION="${OLD_VERSION:-1.153.0}"
PROJECT_ID="${PROJECT_ID:-$(gcloud config get-value project 2>/dev/null)}"
GSA_NAME="${GSA_NAME:-cnrm-system}"
GSA_EMAIL="${GSA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"
TMP_DIR="$(mktemp -d /tmp/kcc-install-XXXXXX)"

trap 'rm -rf "${TMP_DIR}"' EXIT

echo "============================================================"
echo "Installing KCC Version: ${OLD_VERSION}"
echo "Target Project:         ${PROJECT_ID}"
echo "GCP Service Account:    ${GSA_EMAIL}"
echo "============================================================"

# Step 1: Download and extract release bundle for 1.153.0
echo "Downloading KCC release bundle for version ${OLD_VERSION}..."
curl -sSL "https://storage.googleapis.com/configconnector-operator/${OLD_VERSION}/release-bundle.tar.gz" -o "${TMP_DIR}/release-bundle.tar.gz"

echo "Extracting release bundle..."
tar -zxvf "${TMP_DIR}/release-bundle.tar.gz" -C "${TMP_DIR}"

# Step 2: Apply Config Connector Operator
echo "Applying Config Connector Operator 1.153.0..."
kubectl apply -f "${TMP_DIR}/operator-system/configconnector-operator.yaml"

echo "Waiting for Config Connector Operator pod to be Ready..."
kubectl wait --for=condition=Ready pod -l gcp-service=configconnector-operator -n configconnector-operator-system --timeout=300s || true

# Step 3: Configure Workload Identity binding for cluster mode
echo "Configuring Workload Identity policy binding..."
gcloud iam service-accounts add-iam-policy-binding "${GSA_EMAIL}" \
  --project="${PROJECT_ID}" \
  --role="roles/iam.workloadIdentityUser" \
  --member="serviceAccount:${PROJECT_ID}.svc.id.goog[cnrm-system/cnrm-controller-manager]" \
  --condition=None >/dev/null

# Step 4: Apply ConfigConnector Custom Resource
echo "Configuring ConfigConnector CR in cluster mode..."
cat <<EOF | kubectl apply -f -
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: cluster
  googleServiceAccount: "${GSA_EMAIL}"
EOF

# Step 5: Wait for CNRM System namespace and controller manager to start
echo "Waiting for KCC controller manager components to be ready..."
kubectl wait --for=condition=Ready pod -l cnrm.cloud.google.com/component=cnrm-controller-manager -n cnrm-system --timeout=300s || {
  echo "Current pod status in cnrm-system:"
  kubectl get pods -n cnrm-system
}

echo "Config Connector version ${OLD_VERSION} installed successfully!"
