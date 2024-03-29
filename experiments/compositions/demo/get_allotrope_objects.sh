#!/bin/bash

scriptpath=$(realpath $0)
base=$(dirname $scriptpath)

namespace=config-control
echo "Composition ----------------------------------------"
kubectl get  composition
echo
echo "InputAPI --------------------------------------------"
kubectl get  cloudsqls -n $namespace
kubectl get  appteams -n $namespace
echo
echo "Plan ---------------------------------------------------"
kubectl get  plans -n $namespace
echo
echo "Context ------------------------------------------------"
kubectl get  contexts -n $namespace
echo