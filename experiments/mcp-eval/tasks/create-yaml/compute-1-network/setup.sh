#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
mkdir -p ${REPO_ROOT}/.build/tasks/compute-1-network
rm -f ${REPO_ROOT}/.build/tasks/compute-1-network/*
