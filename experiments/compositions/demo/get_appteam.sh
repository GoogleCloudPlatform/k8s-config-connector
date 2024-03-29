#!/bin/bash

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