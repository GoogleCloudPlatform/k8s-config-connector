# Copyright 2024 Google LLC
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

# Apply the compostion
kubectl apply -f 01-composition.yaml

# Create a GCP service account for this team and
#   grant KCC permission according to https://cloud.google.com/config-connector/docs/how-to/install-namespaced
export NAMESPACE=team-eks
export GCP_SA_NAME="${NAMESPACE}"
export PROJECT_ID=$(gcloud config get-value project)
export GSA_EMAIL="${GCP_SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"

gcloud iam service-accounts create ${GCP_SA_NAME} --project ${PROJECT_ID}
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member="serviceAccount:${GSA_EMAIL}" \
    --role="roles/owner"
gcloud iam service-accounts add-iam-policy-binding \
    ${GSA_EMAIL} \
    --member="serviceAccount:${PROJECT_ID}.svc.id.goog[cnrm-system/cnrm-controller-manager-${NAMESPACE}]" \
    --role="roles/iam.workloadIdentityUser" \
    --project ${PROJECT_ID}
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member="serviceAccount:${GSA_EMAIL}" \
    --role="roles/monitoring.metricWriter"

# Create namespace for Alice team
cat 02-context.yaml | envsubst | \
  kubectl apply  -f -

