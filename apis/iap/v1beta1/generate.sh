#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
source "${REPO_ROOT}/dev/tools/goimports.sh"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

./generate-proto.sh

go run . generate-types \
  --service google.cloud.iap.v1 \
  --api-version iap.cnrm.cloud.google.com/v1beta1 \
  --include-skipped-output \
  --resource IAPSettings:IapSettings

go run . generate-mapper \
  --service google.cloud.iap.v1 \
  --api-version iap.cnrm.cloud.google.com/v1beta1 \
  --include-skipped-output

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/iap/
