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

# delete KCC CRDs
if [[ $1 == '--kcc' ]]; then
   kubectl delete -f $base/../manifests/kcc-crds/iam_v1beta1_iampolicymember.yaml || true
   kubectl delete -f $base/../manifests/kcc-crds/kms_v1beta1_kmscryptokey.yaml || true
   kubectl delete -f $base/../manifests/kcc-crds/kms_v1beta1_kmskeyring.yaml || true
   kubectl delete -f $base/../manifests/kcc-crds/serviceusage_v1beta1_serviceidentity.yaml || true
   kubectl delete -f $base/../manifests/kcc-crds/sql_v1beta1_sqlinstance.yaml || true
fi

# cleanup
kubectl delete namespace facade || true
kubectl delete -f $base/../manifests/test1/cloudsql1.yaml || true
kubectl delete -f $base/../manifests/test1/cloudsql1.yaml || true
kubectl delete -f $base/../manifests/test1/cloudsql2.yaml || true
kubectl delete -f $base/../manifests/test1/composition.yaml || true
kubectl delete -f $base/../manifests/test1/plan1.yaml || true
kubectl delete -f $base/../manifests/test1/plan2.yaml || true
kubectl delete -f $base/../manifests/test1/context.yaml || true
kubectl delete -f $base/../../composition/config/crd/bases/composition.google.com_compositions.yaml || true
kubectl delete -f $base/../../composition/config/crd/bases/composition.google.com_contexts.yaml || true
kubectl delete -f $base/../../composition/config/crd/bases/composition.google.com_plans.yaml || true
kubectl delete -f $base/../../facade/config/crd/bases/facade.facade_cloudsqls.yaml || true