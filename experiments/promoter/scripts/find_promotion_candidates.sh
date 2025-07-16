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


# This script finds KCC Kind names that are ready to be promoted from v1alpha1 to v1beta1.
#
# A Kind is considered ready for promotion if it meets the following criteria:
# 1. It has a v1alpha1 types file (apis/<service>/v1alpha1/<resource>_types.go)
#    but no v1beta1 types file (apis/<service>/v1beta1/<resource>_types.go).
# 2. It has a direct controller implementation (pkg/controller/direct/<service>/<resource>_controller.go).
# 3. It has existing test fixtures (pkg/test/resourcefixture/testdata/basic/<service>/v1alpha1/<kind>*)

# Initialize an empty array to store the candidates
candidates=()

# Find all v1alpha1 resource type definitions
while IFS= read -r alpha_types_file; do
  # Extract the service and resource name from the path.
  service=$(echo "${alpha_types_file}" | cut -d'/' -f2)
  resource=$(basename "${alpha_types_file}" _types.go)

  # Extract the Kind name from the `GroupVersion.WithKind` line in the _types.go file.
  kind_name=$(grep 'GroupVersion.WithKind' "${alpha_types_file}" | sed -n 's/.*WithKind("\([^"]*\)").*/\1/p' | head -n 1)
  if [ -z "${kind_name}" ]; then
      # Log to stderr so it doesn't interfere with the JSON output
      echo "Could not determine Kind for ${resource} in ${alpha_types_file}" >&2
      continue
  fi
  
  # Condition 1: Check if a v1beta1 version of the _types.go file does NOT exist.
  beta_types_file="apis/${service}/v1beta1/${resource}_types.go"
  if [ -f "${beta_types_file}" ]; then
    continue
  fi

  # Condition 2: Check if a direct controller for the resource exists.
  resource_no_underscores=$(echo "${resource}" | tr -d '_')
  controller_file="pkg/controller/direct/${service}/${resource_no_underscores}_controller.go"
  if [ ! -f "${controller_file}" ]; then
    kind_lowercase=$(echo "${kind_name}" | tr '[:upper:]' '[:lower:]')
    controller_file="pkg/controller/direct/${service}/${kind_lowercase}_controller.go"
    if [ ! -f "${controller_file}" ]; then
      continue
    fi
  fi

  # Condition 3: Check if a test fixture for the v1alpha1 resource exists.
  test_fixture_dir_name=$(echo "${kind_name}" | tr '[:upper:]' '[:lower:]')
  test_fixture_path="pkg/test/resourcefixture/testdata/basic/${service}/v1alpha1/${test_fixture_dir_name}"
  if ! ls "${test_fixture_path}"* >/dev/null 2>&1; then
    test_fixture_path="pkg/test/resourcefixture/testdata/basic/${service}/v1alpha1/${resource_no_underscores}"
    if ! ls "${test_fixture_path}"* >/dev/null 2>&1; then
      continue
    fi
  fi

  # If all conditions are met, add the candidate to the array
  candidates+=("${kind_name}|${service}|${alpha_types_file}|${controller_file}|${test_fixture_path}")

done < <(find apis -type f -path '*/v1alpha1/*_types.go')

# Use jq to create the JSON output, sorting by kind
printf "%s\n" "${candidates[@]}" | \
  sort | \
  jq -R 'split("|") | {kind: .[0], service: .[1], apiPath: .[2], controllerPath: .[3], testFixturePath: .[4]}' | \
  jq -s '.'
