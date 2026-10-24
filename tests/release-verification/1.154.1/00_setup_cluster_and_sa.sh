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

# Script: 00_setup_cluster_and_sa.sh
# Purpose: Create GKE Cluster with Workload Identity and set up GCP Service Account for KCC.

PROJECT_ID="${PROJECT_ID:-$(gcloud config get-value project 2>/dev/null)}"
CLUSTER_NAME="${CLUSTER_NAME:-kcc-release-1154-verify}"
ZONE="${ZONE:-us-central1-a}"
GSA_NAME="${GSA_NAME:-cnrm-system}"

if [[ -z "${PROJECT_ID}" ]]; then
  echo "Error: PROJECT_ID is not set."
  echo "Please set PROJECT_ID environment variable (e.g. export PROJECT_ID=my-gcp-project)."
  exit 1
fi

echo "============================================================"
echo "Project ID:           ${PROJECT_ID}"
echo "Cluster Name:         ${CLUSTER_NAME}"
echo "Zone:                 ${ZONE}"
echo "GCP Service Account:  ${GSA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"
echo "============================================================"

# Step 0: Enable required GCP services
echo "Enabling required GCP APIs (container, iamcredentials, cloudresourcemanager)..."
gcloud services enable container.googleapis.com iamcredentials.googleapis.com cloudresourcemanager.googleapis.com --project="${PROJECT_ID}" || true

# Step 1: Create GKE Cluster if not already existing
if gcloud container clusters describe "${CLUSTER_NAME}" --zone="${ZONE}" --project="${PROJECT_ID}" >/dev/null 2>&1; then
  echo "GKE cluster '${CLUSTER_NAME}' already exists."
else
  echo "Creating GKE cluster '${CLUSTER_NAME}' with Workload Identity..."
  gcloud container clusters create "${CLUSTER_NAME}" \
    --zone="${ZONE}" \
    --project="${PROJECT_ID}" \
    --workload-pool="${PROJECT_ID}.svc.id.goog" \
    --num-nodes=2 \
    --machine-type=e2-standard-4
fi

# Step 2: Get cluster credentials for kubectl
echo "Getting kubectl credentials for cluster..."
gcloud container clusters get-credentials "${CLUSTER_NAME}" --zone="${ZONE}" --project="${PROJECT_ID}"

# Step 3: Create GCP Service Account for KCC if not existing
GSA_EMAIL="${GSA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"
if gcloud iam service-accounts describe "${GSA_EMAIL}" --project="${PROJECT_ID}" >/dev/null 2>&1; then
  echo "GCP Service Account '${GSA_EMAIL}' already exists."
else
  echo "Creating GCP Service Account '${GSA_NAME}'..."
  gcloud iam service-accounts create "${GSA_NAME}" \
    --project="${PROJECT_ID}" \
    --display-name="KCC Release Verification SA"
fi

# Step 5: Annotate default namespace with GCP Project ID
echo "Annotating default namespace with GCP Project ID (${PROJECT_ID})..."
kubectl annotate namespace default cnrm.cloud.google.com/project-id="${PROJECT_ID}" --overwrite

echo "Cluster and GCP Service Account setup completed successfully."
