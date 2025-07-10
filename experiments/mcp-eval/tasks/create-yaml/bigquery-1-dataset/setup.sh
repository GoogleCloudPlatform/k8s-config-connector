#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
mkdir -p ${REPO_ROOT}/.build/tasks/bigquery-1-dataset
rm -f ${REPO_ROOT}/.build/tasks/bigquery-1-dataset/*
