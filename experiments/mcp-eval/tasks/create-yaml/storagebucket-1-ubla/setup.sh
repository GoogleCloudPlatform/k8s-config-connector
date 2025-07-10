#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
mkdir -p ${REPO_ROOT}/.build/tasks
rm -f ${REPO_ROOT}/.build/tasks/storagebucket-1-ubla/*.yaml