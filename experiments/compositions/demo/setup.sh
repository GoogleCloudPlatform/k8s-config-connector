#!/bin/bash

scriptpath=$(realpath $0)
base=$(dirname $scriptpath)

# we are demoing on CC in config-control namespace
# create Composition CRDs
kubectl apply -f $base/../composition/config/crd/bases/composition.google.com_compositions.yaml
kubectl apply -f $base/../composition/config/crd/bases/composition.google.com_contexts.yaml
kubectl apply -f $base/../composition/config/crd/bases/composition.google.com_plans.yaml
kubectl apply -f $base/../alice/config/crd/bases/alice.alice_cloudsqls.yaml
kubectl apply -f $base/../alice/config/crd/bases/alice.alice_appteams.yaml

kubectl apply -f composition-hasql.yaml  # create composition CR
kubectl apply -f composition-appteam.yaml  # create composition CR
kubectl apply -f context.yaml      # create context CR