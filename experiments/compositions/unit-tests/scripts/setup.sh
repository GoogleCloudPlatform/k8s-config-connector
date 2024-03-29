#!/bin/bash

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
kubectl apply -f $base/../../alice/config/crd/bases/alice.alice_cloudsqls.yaml
# Create Alice Namespace
kubectl create namespace alice || true

kubectl apply -f $base/../manifests/test1/composition.yaml  # create composition CR
kubectl apply -f $base/../manifests/test1/context.yaml      # create context CR
kubectl apply -f $base/../manifests/test1/cloudsql1.yaml    # create alice CR