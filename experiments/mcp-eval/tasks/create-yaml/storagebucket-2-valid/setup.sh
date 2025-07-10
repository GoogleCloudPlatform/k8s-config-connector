#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
mkdir -p ${REPO_ROOT}/.build/tasks
rm -rf ${REPO_ROOT}/.build/tasks/storagebucket-2-valid/*
