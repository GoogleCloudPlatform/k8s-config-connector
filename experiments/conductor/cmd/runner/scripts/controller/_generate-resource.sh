#!/bin/bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -o errexit
set -o nounset
set -o pipefail


export WORKDIR=~/kccai/work1/
export BRANCH_NAME=controller_${SERVICE}_${CRD_KIND}
export OLD_BRANCH_NAME=types_${SERVICE}_${CRD_KIND}
export LOG_DIR=/tmp/conductor/${BRANCH_NAME}



./01-create-controller-client.sh

./02-create-controller.sh

./03-create-identity.sh

./04-create-test.sh

./05-capture-golden-test-output.sh

cat <<EOF

Workflow is now to iterate:

apis/dataproc/v1beta1/generate.sh  && dev/tasks/generate-crds 
go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w  pkg/controller/direct/dataproc/
go test -v ./pkg/fuzztesting/fuzztests/ -fuzz=FuzzAllMappers -fuzztime 600s
<Make changes so fuzzer passes>

If the CRD already exists we need to make sure that there are only description changes in 

git diff origin/master -- config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_dataprocclusters.dataproc.cnrm.cloud.google.com.yaml

diff -u3 \
<(git show origin/master:config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_dataprocclusters.dataproc.cnrm.cloud.google.com.yaml | crd-tools remove-descriptions) \
<(cat config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_dataprocclusters.dataproc.cnrm.cloud.google.com.yaml | crd-tools remove-descriptions) | less
EOF

