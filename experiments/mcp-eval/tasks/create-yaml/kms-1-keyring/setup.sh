#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
mkdir -p ${REPO_ROOT}/.build/tasks/kms-1-keyring
rm -f ${REPO_ROOT}/.build/tasks/kms-1-keyring/*
