#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
mkdir -p ${REPO_ROOT}/.build/tasks/dns-1-zone
rm -f ${REPO_ROOT}/.build/tasks/dns-1-zone/*
