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

cd $(dirname "$0")
SCRIPT_DIR=`pwd`

if [[ -z "${WORKDIR}" ]]; then
  echo "WORKDIR is required"
  exit 1
fi

if [[ -z "${BRANCH_NAME}" ]]; then
  echo "BRANCH_NAME is required"
  exit 1
fi

cd ${WORKDIR}

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

git co master
git co ${BRANCH_NAME}

TESTDIR=pkg/test/resourcefixture/testdata/basic/${SERVICE}/v1alpha1/${CRD_KIND,,}/${CRD_KIND,,}-minimal
mkdir -p ${TESTDIR}

cat <<EOF > ${TESTDIR}/create.yaml
apiVersion: ${CRD_GROUP}/${CRD_VERSION}
kind: ${CRD_KIND}
metadata:
  name: ${CRD_KIND,,}-minimal-\${uniqueId}
spec:
  projectRef:
    external: \${projectId}
  locationID: us-west2
  description: "Initial description"
EOF

cat <<EOF > ${TESTDIR}/update.yaml
apiVersion: ${CRD_GROUP}/${CRD_VERSION}
kind: ${CRD_KIND}
metadata:
  name: ${CRD_KIND,,}-minimal-\${uniqueId}
spec:
  projectRef:
    external: \${projectId}
  locationID: us-west2
  description: "Updated description"
EOF

git status
git add .
git commit -m "${CRD_KIND}: Create minimal test"

codebot --prompt=/dev/stdin <<EOF
Please add a case statement for Group "${CRD_GROUP}" and Kind "${CRD_KIND}" to the switch statement in MaybeSkip,
in the file config/tests/samples/create/harness.go

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to case statement into the list of cases.
* Try to insert it in sorted order first by group, and then by kind.
* If the case already exists, do not make any changes.
EOF

git status
git add .
git commit -m "${CRD_KIND}: Support for testing with mockgcp"


echo "Done"