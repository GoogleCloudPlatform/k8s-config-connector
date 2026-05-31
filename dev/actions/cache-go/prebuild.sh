#!/usr/bin/env bash
# Copyright 2026 Google LLC
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

set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "${REPO_ROOT}"

echo "Starting Go dependency pre-building..."

# Find all directories with a go.mod file, excluding vendor
MODULE_DIRS=$(find . -name "go.mod" -not -path "*/vendor/*" -exec dirname {} \; | sort)

for dir in ${MODULE_DIRS}; do
  echo "=== Processing module in: ${dir} ==="
  pushd "${dir}" > /dev/null
  
  # Determine the module path of the current directory to avoid prebuilding itself
  # We use GOWORK=off to run strictly in module context
  MODULE_PATH=$(GOWORK=off go list -m)
  echo "Module path: ${MODULE_PATH}"

  echo "Downloading modules..."
  GOWORK=off go mod download || true

  echo "Identifying imported packages..."
  # List all packages that are dependencies of the current module's packages
  # Filter out: standard library, main module, empty lines, and any local packages starting with github.com/GoogleCloudPlatform/k8s-config-connector
  PACKAGES=$(GOWORK=off go list -f '{{if not .Standard}}{{if .Module}}{{if not .Module.Main}}{{.ImportPath}}{{end}}{{end}}{{end}}' -deps ./... 2>/dev/null | sort -u | grep -v -E "^(github.com/GoogleCloudPlatform/k8s-config-connector|$)" || true)

  if [[ -n "${PACKAGES}" ]]; then
    NUM_PKGS=$(echo "${PACKAGES}" | wc -l)
    echo "  Prebuilding ${NUM_PKGS} imported packages..."
    echo "${PACKAGES}" | xargs GOWORK=off go build -o /dev/null 2>/dev/null || true
  else
    echo "  No external imported packages found."
  fi

  popd > /dev/null
done

echo "Dependency pre-building completed successfully!"
