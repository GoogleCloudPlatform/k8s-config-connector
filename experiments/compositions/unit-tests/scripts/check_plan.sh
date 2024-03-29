#!/bin/bash

scriptpath=$(realpath $0)
base=$(dirname $scriptpath)

cr=$1
namespace=$2
expected=$3

tmpfile=$(mktemp expanded-output.XXXXXXXX)
kubectl get plans $cr -n $namespace -o json | jq ".spec.stages"   > $tmpfile
echo ".spec.stages:"
cat $tmpfile
diff $tmpfile $expected || echo " ***************** FAILED ********************"
rm $tmpfile