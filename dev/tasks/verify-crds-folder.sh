#!/bin/bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# you may obtain a copy of the License at
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

REPO_ROOT="$(git rev-parse --show-toplevel)"

if [[ -z "${VERSION:-}" ]]; then
  VERSION=$(ls operator/channels/packages/configconnector/ | sort -V | tail -n 1)
fi

TARGET_CRD_DIR=${REPO_ROOT}/crds
TEMP_DIR=$(mktemp -td verify-crds-folder.XXXXXXXX)
CRDS_FILE=${REPO_ROOT}/operator/channels/packages/configconnector/${VERSION}/crds.yaml


if [ ! -f "${CRDS_FILE}" ]; then
    echo "Error: CRDs file not found at ${CRDS_FILE}"
    exit 1

fi

# Parse CRDs into individual files
echo "Splitting manifest ${CRDS_FILE}"
cd ${REPO_ROOT} && go run -mod=readonly ${REPO_ROOT}/scripts/parse-crds/parse-crds.go \
    -file ${CRDS_FILE} \
    -output-dir ${TEMP_DIR}

echo "Adding license headers to split files"
for filepath in ${TEMP_DIR}/*.yaml; do
    filename=$(basename -- ${filepath})
    TMP_FILE=$(mktemp)

    metadata_name=$(awk '/^metadata:/ { found=1 } found && /name:/ { print $2; exit }' "${filepath}")
    if [[ -z "${metadata_name}" ]]; then
      echo "Error: Could not parse metadata.name from '${filepath}'. Aborting."
      exit 1
    fi

    plural_from_file=$(echo "${metadata_name}" | cut -d'.' -f1)
    group_name=$(echo "${metadata_name}" | cut -d'.' -f2)

    if [[ -z "${plural_from_file}" || -z "${group_name}" ]]; then
      echo "Error: Failed to extract plural or group from metadata.name '${metadata_name}' in file '${filepath}'. Aborting."
      exit 1
    fi

    long_filename="apiextensions.k8s.io_v1_customresourcedefinition_${plural_from_file}.${group_name}.cnrm.cloud.google.com.yaml"
    original_filepath=${REPO_ROOT}/config/crds/resources/${long_filename}
    
    if [ -f "${original_filepath}" ]; then
      commit_year=$(git log --reverse -n 1 --format=%ad --date=format:%Y -- "${original_filepath}")
      if [[ -z "${commit_year}" ]]; then
        echo "Error: Could not determine commit year for '${original_filepath}'. Aborting."
        exit 1
      fi
    else
      commit_year=$(date +%Y)
    fi

    LICENSE_HEADER=$(cat << EOF
# Copyright ${commit_year} Google LLC
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

    echo "${LICENSE_HEADER}" > ${TMP_FILE}
    echo "" >> ${TMP_FILE} # Add a empty line after the license header
    cat ${filepath} >> ${TMP_FILE}
    mv ${TMP_FILE} ${filepath}

done

echo "Verifying ${TARGET_CRD_DIR} is up-to-date"
if ! diff -r -N ${TARGET_CRD_DIR} ${TEMP_DIR}; then
    echo "Error: The ${TARGET_CRD_DIR} directory is out of date."
    echo "Please run 'VERSION=${VERSION} dev/tasks/sync-crds-folder.sh' and commit the changes."
    exit 1

fi

rm -rf ${TEMP_DIR}