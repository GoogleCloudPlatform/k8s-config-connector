#!/bin/bash

kubectl delete -f composition-hasql.yaml || true
kubectl delete -f composition-appteam.yaml || true
kubectl delete -f context.yaml || true