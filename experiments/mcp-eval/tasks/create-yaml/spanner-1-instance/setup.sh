#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
mkdir -p ${REPO_ROOT}/.build/tasks/spanner-1-instance
rm -f ${REPO_ROOT}/.build/tasks/spanner-1-instance/*
