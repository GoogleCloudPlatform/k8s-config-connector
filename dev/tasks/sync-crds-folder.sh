# Copyright 2024 Google LLC
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

set -o errexit
set -o nounset
set -o pipefail

VERSION=${1:-1.127.0}  # Default to 1.127.0 if no version provided
REPO_ROOT="$(git rev-parse --show-toplevel)"
TARGET_CRD_DIR=${REPO_ROOT}/crds
TEMP_DIR=$(mktemp -td sync-crds-folder.XXXXXXXX)
CRDS_FILE=${REPO_ROOT}/operator/channels/packages/configconnector/${VERSION}/crds.yaml

LICENSE_HEADER=$(cat << 'EOF'
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
EOF
)

if [ ! -f "${CRDS_FILE}" ]; then
    echo "Error: CRDs file not found at ${CRDS_FILE}"
    exit 1
fi

# Parse CRDs into individual files
cd ${REPO_ROOT} && go run -mod=readonly ${REPO_ROOT}/scripts/parse-crds/parse-crds.go \
    -file ${CRDS_FILE} \
    -output-dir ${TEMP_DIR}


rm -rf ${TARGET_CRD_DIR}
mkdir ${TARGET_CRD_DIR}
for filepath in ${TEMP_DIR}/*.yaml; do
    filename=$(basename -- ${filepath})
    echo "${LICENSE_HEADER}" > ${TARGET_CRD_DIR}/${filename}
    echo "" >> ${TARGET_CRD_DIR}/${filename} # Add a empty line after the license header
    cat ${filepath} >> ${TARGET_CRD_DIR}/${filename}
done

rm -rf ${TEMP_DIR}