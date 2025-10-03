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

#!/bin/bash
set -x
set -e
set -o nounset

echo "KIND: $KIND, the config connector kind" # e.g. CloudbuildWorkerPool
echo "PROTO: $PROTO, the gcp proto name" # e.g. WorkerPool
echo "PACKAGE: $PACKAGE, the gcp proto package or service name" # e.g. google.devtools.cloudbuild.v1
echo "SERVICE: $SERVICE, the config connector api group name" # e.g. cloudbuild
# echo "SUFFIX: $SUFFIX, logs suffix id" # e.g. cloudbuild


REPO_ROOT="$(git rev-parse --show-toplevel)"
DEBUG_LOG="${REPO_ROOT}/resource-${KIND}-fail.log" 
CODEBOT_LOG="${REPO_ROOT}/resource-${KIND}-codebot.log" 
INS_PATH="${REPO_ROOT}/experiments/conductor/cmd/runner/scripts/controller"

function git_commit() {    
    git add apis/"${SERVICE}"/v1alpha1/
    git add pkg/controller/direct/register/register.go 
    git add pkg/controller/direct/"${SERVICE}"/
    git add pkg/gvks/supportedgvks/gvks_generated.go
    git add config/crds/resources/
    git add go.mod
    git add go.sum
    git add go.work

    result=$(git diff --staged)
    if [[ -n ${result} ]]; then
        git commit -m "$1"
    fi
    return 0
}

function make_ready_pr() {
    make lint

    make manifests
    git add config/installbundle/

    make ensure
    make fmt

    WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/...
    git add tests/apichecks/testdata/exceptions/

    git_commit "make read-pr"
}

function run_worker() {
    attempts=1
    while [[ attempts -le 4 ]]; do

        if go build "${REPO_ROOT}"/apis/"${SERVICE}"/v1alpha1/...
        then    
            return 0
        fi

        ins="${INS_PATH}"/llm-gcp-imports.txt
        msg=$(cat $ins)
        msg=$(eval "echo \"$msg\"")
        codebot --ui-type "bash" --base-dir "${REPO_ROOT}" --proto-dir "${REPO_ROOT}"/.build/third_party/googleapis/  <<EOF >> ${CODEBOT_LOG} 2>&1 
"${msg}"
EOF
        if [[ $? -eq 0 ]]; then
            git_commit "[llm] fix gcp imports for ${KIND}"
        fi 

        ins="${INS_PATH}"/llm-fix-gobuild.txt
        msg=$(cat $ins)
        msg=$(eval "echo \"$msg\"")
        codebot --ui-type "bash" --base-dir "${REPO_ROOT}" --proto-dir "${REPO_ROOT}"/.build/third_party/googleapis/  <<EOF >> ${CODEBOT_LOG} 2>&1 
"${msg}"
EOF
        if [[ $? -eq 0 ]]; then
            git_commit "[llm] fix gobuilds for ${KIND}"
        fi 

       ((attempts++))
    done 
    return 
}



if controllerbuilder generate-direct-reconciler \
    --service "${PACKAGE}" \
    --api-version "${SERVICE}".cnrm.cloud.google.com/v1alpha1 \
    --resource "${KIND}":"${PROTO}"
then 
    echo "controller ${KIND} added"
else 
    exit 1
fi

if "${REPO_ROOT}"/dev/tasks/generate-crds
then 
    echo "crd updated"
else 
    exit 1
fi

if run_worker
then
    if make_ready_pr
    then
        echo "Passed"
    fi 
else 
    echo "Unfortunately, this run failed. Please fix manually" 
fi