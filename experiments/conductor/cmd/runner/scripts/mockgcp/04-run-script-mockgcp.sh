#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

set -x

cd $(dirname "$0")
SCRIPT_DIR=`pwd`

if [[ -z "${WORKDIR}" ]]; then
  echo "WORKDIR is required"
  exit 1
fi

if [[ -z "${BRANCH_NAME}" ]]; then
  echo "BRANCH_NAME is required"
  exit 1
fi

if [[ -z "${LOG_DIR}" ]]; then
  echo "LOG_DIR is required"
  exit 1
fi

if [[ -z "${RUN_TEST}" ]]; then
  echo "RUN_TEST is required"
  exit 1
fi

mkdir -p ${LOG_DIR}

cd ${WORKDIR}

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

git co master
git co ${BRANCH_NAME}

cd ${REPO_ROOT}/mockgcp

export WRITE_GOLDEN_OUTPUT=1
export E2E_GCP_TARGET=mock

echo "Running test"
# We ignore test failures, because we expect the golden output to be volatile at this stage
(go test ./mockgcptests -v -run TestScripts/${RUN_TEST} || true) | tee ${LOG_DIR}/test-mockgcp.log

echo "Done"