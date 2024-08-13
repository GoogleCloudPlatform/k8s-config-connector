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

if [[ $# -ne 1 ]];
then
  echo "usage: ./get_cloudsql.sh <namespace>"
  exit 1
fi

namespace=$1

echo "ServiceIdentity ----------------------------------------"
kubectl get  serviceidentity -n $namespace
echo
echo "SqlInstance --------------------------------------------"
kubectl get  sqlinstances.sql.cnrm.cloud.google.com -n $namespace
echo
echo "KMSKeyRings --------------------------------------------"
kubectl get  kmskeyring -n $namespace
echo
echo "KMSCryptoKeys ------------------------------------------"
kubectl get  kmscryptokey -n $namespace
echo
echo "IAMPolicyMember ----------------------------------------"
kubectl get  iampolicymember -n $namespace
echo
echo "ServiceUsage -------------------------------------------"
kubectl get services.serviceusage.cnrm.cloud.google.com -n $namespace
