#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

cd $(dirname "$0")
SCRIPT_DIR=`pwd`

PROMPT_DIR=${SCRIPT_DIR}

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


if [[ -z "${PROTO_PACKAGE}" ]]; then
  echo "PROTO_PACKAGE is required"
  exit 1
fi

mkdir -p ${LOG_DIR}

cd ${WORKDIR}

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

git co master
git co ${BRANCH_NAME}

cd ${REPO_ROOT}/mockgcp

cat ${PROMPT_DIR}/03-add-to-makefile.prompt | \
    envsubst '$PROTO_PACKAGE' > ${LOG_DIR}/03-add-to-makefile.prompt

codebot --prompt=${LOG_DIR}/03-add-to-makefile.prompt


git status
git add Makefile
git add .
git commit -m "autogen: update Makefile for ${PROTO_PACKAGE}"

make generate-grpc-for-google-protos

git add .
git commit -m "autogen: updated GRPC/proto dependencies for ${PROTO_PACKAGE}"

echo "Done"





