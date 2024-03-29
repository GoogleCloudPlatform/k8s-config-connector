#!/bin/bash

scriptpath=$(realpath $0)
base=$(dirname $scriptpath)


name=$1
namespace=$2

kubectl delete -f cloudsqls.alice.alice -n $namespace ${name}

kubectl delete  sqlinstances.sql.cnrm.cloud.google.com -n $namespace ${name}-db-main
kubectl delete  sqlinstances.sql.cnrm.cloud.google.com -n $namespace ${name}-db-replica-us-central1
kubectl delete  kmskeyring -n $namespace  kmscryptokeyring-us-central1 
kubectl delete  kmskeyring -n $namespace  kmscryptokeyring-us-east1 
kubectl delete  kmscryptokey -n $namespace kmscryptokey-enc-us-central1 
kubectl delete  kmscryptokey -n $namespace kmscryptokey-enc-us-east1 
kubectl delete  iampolicymember -n $namespace sql-kms-us-east1-policybinding 
kubectl delete  iampolicymember -n $namespace sql-kms-us-central1-policybinding 
kubectl delete  serviceidentity -n $namespace sqladmin.googleapis.com 
kubectl delete  services.serviceusage.cnrm.cloud.google.com -n $namespace cloudkms.googleapis.com 
kubectl delete  services.serviceusage.cnrm.cloud.google.com -n $namespace iam.googleapis.com 
kubectl delete  services.serviceusage.cnrm.cloud.google.com -n $namespace serviceusage.googleapis.com
kubectl delete  services.serviceusage.cnrm.cloud.google.com -n $namespace sqladmin.googleapis.com 