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

# if [[ -z "${GCLOUD_COMMAND}" ]]; then
#   echo "GCLOUD_COMMAND is required"
#   exit 1
# fi

if [[ -z "${LOG_DIR}" ]]; then
  echo "LOG_DIR is required"
  exit 1
fi

# if [[ -z "${EXPECTED_PATH}" ]]; then
#   echo "EXPECTED_PATH is required"
#   exit 1
# fi

mkdir -p ${LOG_DIR}
# cat ${PROMPT} | \
#     envsubst '$GCLOUD_COMMAND,$EXPECTED_PATH' > ${LOG_DIR}/prompt

cd ${WORKDIR}

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

git co master
git co ${BRANCH_NAME}

cd apis/${SERVICE}/${CRD_VERSION}

# TODO: We should be able to in-place replace here
# TODO: Generate spec and status separately
# TODO: lower case (but it's nice right now because it always fails compilation)
# TODO: Auto-exclude any examples with the same proto??
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF >> ${RESOURCE}_types.go
// +kcc:proto=${PROTO_SERVICE}.${PROTO_RESOURCE}
EOF

# TODO: go mod tidy?

git status
git add .
git commit -m "Generate spec and status for ${SERVICE} ${RESOURCE}"

echo "Done"