#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
mkdir -p ${REPO_ROOT}/.build/tasks
rm -rf ${REPO_ROOT}/.build/tasks/storagebucket-3-apply/*
kubectl delete namespace storagebucket-apply  --ignore-not-found
kubectl create namespace storagebucket-apply 