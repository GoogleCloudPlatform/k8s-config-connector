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

# create KCC CRDs
if [[ $1 == '--kcc' ]]; then
   kubectl apply -f $base/../manifests/kcc-crds/iam_v1beta1_iampolicymember.yaml
   kubectl apply -f $base/../manifests/kcc-crds/kms_v1beta1_kmscryptokey.yaml
   kubectl apply -f $base/../manifests/kcc-crds/kms_v1beta1_kmskeyring.yaml
   kubectl apply -f $base/../manifests/kcc-crds/serviceusage_v1beta1_serviceidentity.yaml
   kubectl apply -f $base/../manifests/kcc-crds/sql_v1beta1_sqlinstance.yaml
fi

# create Composition CRDs
kubectl apply -f $base/../../composition/config/crd/bases/composition.google.com_compositions.yaml
kubectl apply -f $base/../../composition/config/crd/bases/composition.google.com_contexts.yaml
kubectl apply -f $base/../../composition/config/crd/bases/composition.google.com_plans.yaml
kubectl apply -f $base/../../facade/config/crd/bases/facade.facade_cloudsqls.yaml
# Create Alice Namespace
kubectl create namespace facade || true

kubectl apply -f $base/../manifests/test1/composition.yaml  # create composition CR
kubectl apply -f $base/../manifests/test1/context.yaml      # create context CR
kubectl apply -f $base/../manifests/test1/cloudsql1.yaml    # create facade CR