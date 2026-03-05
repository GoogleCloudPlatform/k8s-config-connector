#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

cd "$(dirname "$0")"

go run ../../../dev/tools/controllerbuilder/main.go generate-types \
  --service google.devtools.cloudbuild.v1 \
  --api-version cloudbuild.cnrm.cloud.google.com/v1beta1 \
  --resource CloudBuildTrigger:BuildTrigger \
  --skip-scaffold-files

go run ../../../dev/tools/controllerbuilder/main.go generate-mapper \
  --service google.devtools.cloudbuild.v1 \
  --api-version cloudbuild.cnrm.cloud.google.com/v1beta1
