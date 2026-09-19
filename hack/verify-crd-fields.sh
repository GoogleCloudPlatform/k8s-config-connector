#!/bin/bash

# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# This script helps to compare the fields defined in a CRD YAML with the
# fields in a Go source file. It extracts the field paths from the CRD
# and normalizes them, making it easier to compare with the Go struct definitions.

set -e

CRD_FILE="${1}"
GO_FILE="${2}"

if [ -z "${CRD_FILE}" ] || [ -z "${GO_FILE}" ]; then
  echo "Usage: $0 <path_to_crd_yaml> <path_to_go_types_file>"
  exit 1
fi

CRD_FIELDS_TMP=$(mktemp)
CRD_FIELDS_FINAL="crd_fields.txt"

echo "Extracting fields from CRD: ${CRD_FILE}"

# Use yq to extract all nested paths from the spec properties.
# The output is a stream of dot-separated paths.
# This requires yq to be installed (https://github.com/mikefarah/yq)
yq e '.spec.versions[].schema.openAPIV3Schema.properties.spec.properties | .. | path | . as $p | select(length > 0) | $p | join(".")' "${CRD_FILE}" > "${CRD_FIELDS_TMP}"

echo "Normalizing CRD fields..."

# Normalize the paths to make them easier to compare with Go fields.
# This removes the verbose parts of the path from yq.
# We are left with a cleaner path like "template.template.containers.name".
sed -i -e 's/\.properties//g' \
       -e 's/\.items//g' \
       -e '/additionalProperties/d' \
       -e '/oneOf/d' \
       -e '/not/d' \
       -e '/anyOf/d' \
       "${CRD_FIELDS_TMP}"

# Sort the fields for consistent output and easier diffing.
sort "${CRD_FIELDS_TMP}" > "${CRD_FIELDS_FINAL}"

rm "${CRD_FIELDS_TMP}"

echo "CRD fields extracted and saved to ${CRD_FIELDS_FINAL}"
echo "Please compare this file with the Go structs in ${GO_FILE}"
echo "---"
echo "Example from ${CRD_FIELDS_FINAL}:"
head -n 5 "${CRD_FIELDS_FINAL}"
echo "..."

# At this point, I (the AI) would read the crd_fields.txt and the go file,
# and perform the comparison in my thought process.
# The script has served its purpose of providing a structured, flattened view of the CRD.

# For the Go file, a similar extraction is much harder with shell tools.
# A simple grep for json tags can be a starting point, but it doesn't handle nesting reliably.
# The most reliable way would be to manually compare the crd_fields.txt with the Go file.
