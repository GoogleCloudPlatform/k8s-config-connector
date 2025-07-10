#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
mkdir -p ${REPO_ROOT}/.build/tasks/iam-1-serviceaccount
rm -f ${REPO_ROOT}/.build/tasks/iam-1-serviceaccount/*
