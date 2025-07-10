#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
mkdir -p ${REPO_ROOT}/.build/tasks/pubsub-1-topic
rm -f ${REPO_ROOT}/.build/tasks/pubsub-1-topic/*
