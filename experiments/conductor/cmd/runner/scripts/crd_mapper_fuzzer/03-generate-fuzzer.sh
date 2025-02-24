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

controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > pkg/controller/direct/${SERVICE}/${RESOURCE}_fuzzer.go
// +tool:fuzz-gen
// proto.message: ${PROTO_MESSAGE}
EOF

# TODO: We could easily do this "classically"
codebot --prompt=/dev/stdin <<EOF
Add an unnamed (_) go import for  
"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/${SERVICE}"
to the imports in pkg/controller/direct/register/register.go
EOF

# TODO: go mod tidy?

git status
git add .
git commit -m "Generate fuzzer ${SERVICE} ${RESOURCE}"

echo "Done"