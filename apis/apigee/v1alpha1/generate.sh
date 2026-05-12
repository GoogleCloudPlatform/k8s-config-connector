#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

# Use the pinned SHA from apis/git.versions
PROTO_SHA="731d7f2ab6e4e2ea15030c95039e2cb66174d4fb" 
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

go run . generate-types \
    --service mockgcp.cloud.apigee.v1 \
    --api-version apigee.cnrm.cloud.google.com/v1alpha1 \
    --resource ApigeeApiProduct:GoogleCloudApigeeV1ApiProduct \
    --proto-source-path ${PROTO_OUT} \
    --skip-scaffold-files
