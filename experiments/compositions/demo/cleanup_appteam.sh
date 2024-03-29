#!/bin/bash

scriptpath=$(realpath $0)
base=$(dirname $scriptpath)

name=$1
project=$2

kubectl delete appteams.alice.alice -n config-control ${name}
kubectl delete context.composition.google.com context -n ${project}

kubectl delete iamserviceaccount.iam.cnrm.cloud.google.com kcc-${project} -n config-control
kubectl delete iampartialpolicy.iam.cnrm.cloud.google.com ${project}-sa-workload-identity-binding -n config-control
kubectl delete iampartialpolicy.iam.cnrm.cloud.google.com kcc-owners-permissions-${project} -n config-control
kubectl delete storagebuckets.storage.cnrm.cloud.google.com test-bucket-${project} -n ${project}
kubectl delete project.resourcemanager.cnrm.cloud.google.com ${project} -n config-control

echo "waiting for project to be deleted ......"
sleep 300

kubectl delete configconnectorcontext.core.cnrm.cloud.google.com configconnectorcontext.core.cnrm.cloud.google.com -n ${project}
kubectl delete namespace ${project}