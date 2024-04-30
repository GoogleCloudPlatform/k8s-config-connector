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

# Step 1. Apply the CRD for Alice team
kubectl apply -k crdBuilder/config/crd

# Step 2. Apply the compostion
kubectl apply -f composition-aks/01-composition.yaml

# Step 3. Create namespace for Alice team
kubectl apply -f composition-aks/02-context.yaml

# Step 4. Grant permission according to https://cloud.google.com/config-connector/docs/how-to/install-namespaced
gcloud iam service-accounts create alice-1 --project zicong-gke-multi-cloud-dev-2
gcloud projects add-iam-policy-binding zicong-gke-multi-cloud-dev-2 \
    --member="serviceAccount:alice-1@zicong-gke-multi-cloud-dev-2.iam.gserviceaccount.com" \
    --role="roles/owner"
gcloud iam service-accounts add-iam-policy-binding \
    alice-1@zicong-gke-multi-cloud-dev-2.iam.gserviceaccount.com \
    --member="serviceAccount:zicong-gke-multi-cloud-dev-2.svc.id.goog[cnrm-system/cnrm-controller-manager-alice-1]" \
    --role="roles/iam.workloadIdentityUser" \
    --project zicong-gke-multi-cloud-dev-2
gcloud projects add-iam-policy-binding zicong-gke-multi-cloud-dev-2 \
    --member="serviceAccount:alice-1@zicong-gke-multi-cloud-dev-2.iam.gserviceaccount.com" \
    --role="roles/monitoring.metricWriter"