#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
rm -f ${REPO_ROOT}/.build/tasks/sql-1-instance/*
