#!/bin/bash -x
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


scriptpath=$(realpath $0)
base=$(dirname $scriptpath)

name=clearing
project=$1

kubectl delete appteams.facade.facade -n config-control ${name}
kubectl delete context.composition.google.com context -n ${project}

# These are redundant
kubectl delete iamserviceaccount.iam.cnrm.cloud.google.com kcc-${project} -n config-control
kubectl delete iampartialpolicy.iam.cnrm.cloud.google.com ${project}-sa-workload-identity-binding -n config-control
kubectl delete iampartialpolicy.iam.cnrm.cloud.google.com kcc-owners-permissions-${project} -n config-control
kubectl delete storagebuckets.storage.cnrm.cloud.google.com test-bucket-${project} -n ${project}
kubectl delete project.resourcemanager.cnrm.cloud.google.com ${project} -n config-control

echo "waiting for project to be deleted ......"
sleep 30

kubectl delete configconnectorcontext.core.cnrm.cloud.google.com configconnectorcontext.core.cnrm.cloud.google.com -n ${project}
kubectl delete namespace ${project}