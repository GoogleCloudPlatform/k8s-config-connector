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


# W0318 16:49:25.190246 1183864 results.go:65] error from apply on serviceusage.cnrm.cloud.google.com/v1beta1, Kind=ServiceIdentity config-control/sqladmin.googleapis.com: error from apply: namespaces "config-control" not found
# W0318 16:49:25.374357 1183864 results.go:65] error from apply on kms.cnrm.cloud.google.com/v1beta1, Kind=KMSKeyRing config-control/kmscryptokeyring-us-east1: error from apply: namespaces "config-control" not found
# W0318 16:49:25.564951 1183864 results.go:65] error from apply on kms.cnrm.cloud.google.com/v1beta1, Kind=KMSCryptoKey config-control/kmscryptokey-enc-us-east1: error from apply: namespaces "config-control" not found
# W0318 16:49:25.751457 1183864 results.go:65] error from apply on iam.cnrm.cloud.google.com/v1beta1, Kind=IAMPolicyMember config-control/sql-kms-east-policybinding: error from apply: namespaces "config-control" not found
# W0318 16:49:25.934380 1183864 results.go:65] error from apply on sql.cnrm.cloud.google.com/v1beta1, Kind=SQLInstance config-control/facadeDB-main: error from apply: namespaces "config-control" not found
# W0318 16:49:26.116118 1183864 results.go:65] error from apply on kms.cnrm.cloud.google.com/v1beta1, Kind=KMSKeyRing config-control/kmscryptokeyring-us-central1: error from apply: namespaces "config-control" not found
# W0318 16:49:26.298863 1183864 results.go:65] error from apply on kms.cnrm.cloud.google.com/v1beta1, Kind=KMSCryptoKey config-control/kmscryptokey-enc-us-central1: error from apply: namespaces "config-control" not found
# W0318 16:49:26.462008 1183864 results.go:65] error from apply on iam.cnrm.cloud.google.com/v1beta1, Kind=IAMPolicyMember config-control/sql-kms-east-policybinding: error from apply: namespaces "config-control" not found
# W0318 16:49:26.642842 1183864 results.go:65] error from apply on sql.cnrm.cloud.google.com/v1beta1, Kind=SQLInstance config-control/facadeDB-replica: error from apply: namespaces "config-control" not found

failed=false
kubectl get  sqlinstances.sql.cnrm.cloud.google.com -n facade facadedb-main || falied=true
kubectl get  sqlinstances.sql.cnrm.cloud.google.com -n facade facadedb-replica || falied=true
kubectl get  kmskeyring -n facade  kmscryptokeyring-us-central1 || failed=true
kubectl get  kmskeyring -n facade  kmscryptokeyring-us-east1 || failed=true
kubectl get  kmscryptokey -n facade kmscryptokey-enc-us-central1 || failed=true
kubectl get  kmscryptokey -n facade kmscryptokey-enc-us-east1 || failed=true
kubectl get  iampolicymember -n facade sql-kms-us-east1-policybinding || failed=true
kubectl get  iampolicymember -n facade sql-kms-us-central1-policybinding || failed=true
kubectl get  serviceidentity -n facade sqladmin.googleapis.com || failed=true

if [[ $failed == true ]]; then
    echo "------ FAILED ---------------------------------"
    echo "One or more resources are missing"
fi

