#!/bin/bash

scriptpath=$(realpath $0)
base=$(dirname $scriptpath)

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