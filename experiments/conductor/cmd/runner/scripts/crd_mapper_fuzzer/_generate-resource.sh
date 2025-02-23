#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail


# TODO: Can we infer / default RESOURCE from PROTO_MESSAGE

export WORKDIR=~/kccai/work1/
export BRANCH_NAME=types_${SERVICE}_${CRD_KIND}
export LOG_DIR=/tmp/conductor/${BRANCH_NAME}

#./01-write-generator-script.sh

#./02-generate-spec-and-status.sh

./03-generate-fuzzer.sh


cat <<EOF

Workflow is now to iterate:

apis/dataproc/v1beta1/generate.sh  && dev/tasks/generate-crds 
go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w  pkg/controller/direct/dataproc/
go test -v ./pkg/fuzztesting/fuzztests/ -fuzz=FuzzAllMappers -fuzztime 600s
<Make changes so fuzzer passes>

If the CRD already exists we need to make sure that there are only description changes in 

git diff origin/master -- config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_dataprocclusters.dataproc.cnrm.cloud.google.com.yaml
EOF


diff -u3 \
<(git show origin/master:config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_dataprocclusters.dataproc.cnrm.cloud.google.com.yaml | crd-tools remove-descriptions) \
<(cat config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_dataprocclusters.dataproc.cnrm.cloud.google.com.yaml | crd-tools remove-descriptions) | less
