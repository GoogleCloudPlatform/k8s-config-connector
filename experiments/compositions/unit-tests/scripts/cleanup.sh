#!/bin/bash

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
kubectl delete namespace alice || true
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
kubectl delete -f $base/../../alice/config/crd/bases/alice.alice_cloudsqls.yaml || true