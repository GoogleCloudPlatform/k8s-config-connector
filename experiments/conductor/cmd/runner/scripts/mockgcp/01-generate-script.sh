#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

cd $(dirname "$0")
SCRIPT_DIR=`pwd`

PROMPT=${SCRIPT_DIR}/01-generate-script.prompt

if [[ -z "${WORKDIR}" ]]; then
  echo "WORKDIR is required"
  exit 1
fi

if [[ -z "${BRANCH_NAME}" ]]; then
  echo "BRANCH_NAME is required"
  exit 1
fi

if [[ -z "${GCLOUD_COMMAND}" ]]; then
  echo "GCLOUD_COMMAND is required"
  exit 1
fi

if [[ -z "${LOG_DIR}" ]]; then
  echo "LOG_DIR is required"
  exit 1
fi

if [[ -z "${EXPECTED_PATH}" ]]; then
  echo "EXPECTED_PATH is required"
  exit 1
fi

mkdir -p ${LOG_DIR}
cat ${PROMPT} | \
    envsubst '$GCLOUD_COMMAND,$EXPECTED_PATH' > ${LOG_DIR}/prompt

cd ${WORKDIR}

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

git co master
git co ${BRANCH_NAME}

cd ${REPO_ROOT}/mockgcp

codebot --prompt=${LOG_DIR}/prompt | tee ${LOG_DIR}/codebot.log

git status
git add ${EXPECTED_PATH}/script.yaml
git add .
git commit -m "autogen: generate-script for ${GCLOUD_COMMAND}"

echo "Done"