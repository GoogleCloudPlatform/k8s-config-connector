#!/bin/bash
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

name=$1
project=$2

echo "AppTeam ----------------------------------------"
kubectl get appteam ${name} -n config-control
echo
echo "Composition Context --------------------------------"
kubectl get context.composition.google.com context -n ${project}
echo
echo "IAMServiceAccount ----------------------------------------"
kubectl get iamserviceaccount kcc-${project} -n config-control
echo
echo "IAMPartialPolicy --------------------------------------------"
kubectl get iampartialpolicy -n config-control ${project}-sa-workload-identity-binding
kubectl get iampartialpolicy -n config-control kcc-owners-permissions-${project}
echo
echo "StorageBuckets --------------------------------------------"
kubectl get storagebucket -n ${project} test-bucket-${project}
echo
echo "CCC --------------------------------------------"
kubectl get configconnectorcontext -n ${project}
echo
echo "Project ------------------------------------------"
kubectl get project ${project} -n config-control
echo