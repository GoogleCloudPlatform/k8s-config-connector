#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if [[ -z "${GCLOUD_COMMAND}" ]]; then
  echo "GCLOUD_COMMAND is required"
  exit 1
fi

export WORKDIR=~/kccai/work1/
export BRANCH_NAME=gcloud_${SERVICE}_${RESOURCE}
export LOG_DIR=/tmp/conductor/${BRANCH_NAME}

export EXPECTED_PATH=mock${SERVICE}/testdata/${RESOURCE}/crud
./01-generate-script.sh

export RUN_TEST=${EXPECTED_PATH}
./02-run-script-real-gcp.sh 

./03a-add-to-makefile.sh

./03-implement-mocks.sh

./04-run-script-mockgcp.sh
