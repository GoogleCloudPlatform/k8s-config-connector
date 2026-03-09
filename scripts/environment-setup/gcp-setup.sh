#!/bin/bash
# Copyright 2022 Google LLC
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

# Script to set up GCP
set -o errexit

# Configure gcloud with your login credentials.
gcloud auth login
gcloud auth application-default login

# Set PROJECT_ID to your current project
export PROJECT_ID=$(gcloud config get-value project)

# Enable the container registry service for storing images created by the make
# docker-push command.
gcloud services enable containerregistry.googleapis.com

# Configure gcloud to allow docker to authorize and recognize the gcr.io
# registry.
gcloud auth configure-docker

# Enable GKE for your project.
gcloud services enable container.googleapis.com
# When creating GKE clusters, you must either provide a zone or set the default
# zone for gcloud. Set the default zone for gcloud to us-west1-a.
gcloud config set compute/zone us-west1-a
# Define the name of your GKE cluster as cnrm-dev.
export CLUSTER_NAME="cnrm-dev"

if [[ ! $(gcloud beta container clusters list | grep ${CLUSTER_NAME}) ]]; then
    # Create a GKE cluster with Workload Identity enabled.
    gcloud beta container clusters create ${CLUSTER_NAME} \
        --workload-pool=${PROJECT_ID}.svc.id.goog
fi

echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" \
    | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list

curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key \
    --keyring /usr/share/keyrings/cloud.google.gpg add -

sudo apt-get update && sudo apt-get install google-cloud-cli

sudo apt-get install google-cloud-sdk-gke-gcloud-auth-plugin

# Configure kubectl to communicate with the cluster.
gcloud container clusters get-credentials ${CLUSTER_NAME}
# Add an annotation to your default K8s namespace to bind it to your GCP project.
kubectl annotate namespace default \
    "cnrm.cloud.google.com/project-id=${PROJECT_ID}" \
    --overwrite

if [[ ! $(gcloud iam service-accounts list | grep "cnrm-system") ]]; then
    # Create a GCP Service Account.
    gcloud iam service-accounts create cnrm-system
fi

# Give the GCP Service Account elevated permissions on your project.
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member="serviceAccount:cnrm-system@${PROJECT_ID}.iam.gserviceaccount.com" \
    --role="roles/owner"

# Create a GCP IAM Policy Binding between the GCP Service Account and the
# Kubernetes Service Account that will later be created and used by the CNRM
# Controller Manager.
gcloud iam service-accounts add-iam-policy-binding cnrm-system@${PROJECT_ID}.iam.gserviceaccount.com \
    --member="serviceAccount:${PROJECT_ID}.svc.id.goog[cnrm-system/cnrm-controller-manager]" \
    --role="roles/iam.workloadIdentityUser"

# To ensure the logs are ingested, enable the stackdriver service.
gcloud services enable stackdriver.googleapis.com

GREEN='\033[0;32m'
NC='\033[0m'
echo -e "${GREEN}GCP SETUP SUCCESSFUL${NC}"
