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

# Step 1. Apply CR to create composition of ASK and its attachment
kubectl apply -f 03-attached-1.yaml

# Step 2. login to the cluster to apply the manifest
AKS_NAME=$(cat 03-attached-1.yaml |yq .metadata.name)
ATTACHED_REGION=$(cat 03-attached-1.yaml |yq .spec.gcpRegion)
ATTACHED_PLATFORM_VERSION=$(cat 03-attached-1.yaml |yq .spec.attachedPlatformVersion)

gcloud container attached clusters generate-install-manifest \
  $AKS_NAME \
  --location=${ATTACHED_REGION} \
  --platform-version ${ATTACHED_PLATFORM_VERSION} \
  --output-file=/tmp/install-agent-${AKS_NAME}.yaml

az aks get-credentials --name ${AKS_NAME}-aks \
  --resource-group ${AKS_NAME}-rg 

kubectl apply -f /tmp/install-agent-${AKS_NAME}.yaml


## Commands to check progress

kubectl get AttachedAKS -n alice-2

kubectl get ResourceGroup.resources.azure.com \
  -n alice-2

kubectl get managedcluster.containerservice.azure.com \
  -n alice-2

kubectl get cm -n alice-2

kubectl get containerattachedcluster \
  -n alice-2
