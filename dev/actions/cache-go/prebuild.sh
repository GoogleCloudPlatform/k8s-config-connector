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

  echo "Downloading modules..."
  GOWORK=off go mod download || true

  # Only prebuild the e2e tests for the root module to avoid compiling the entire codebase
  # and bloating the cache (since compiling everything results in >10GB cache).
  if [[ "${dir}" == "." ]]; then
    echo "  Prebuilding e2e tests..."
    GOWORK=off go test -c -o /dev/null ./tests/e2e 2>/dev/null || true
  fi

  popd > /dev/null
done

echo "Pre-downloading envtest assets and prebuilding setup-envtest..."
./dev/tasks/setup-envtest || true

echo "=== Cache Statistics ==="
echo "Total size of ~/.cache:"
du -h --max-depth=1 ~/.cache 2>/dev/null || true

echo "Any huge files in ~/.cache (larger than 100M):"
find ~/.cache -size +100M -exec ls -lh {} \; 2>/dev/null || true

echo "Total size of envtest assets:"
du -h --max-depth=2 ~/.local/share/kubebuilder-envtest 2>/dev/null || true

echo "Total size of Go module cache (GOMODCACHE):"
GOMOD_PATH=$(go env GOMODCACHE)
if [[ -d "${GOMOD_PATH}" ]]; then
  du -sh "${GOMOD_PATH}" 2>/dev/null || true
fi

echo "Dependency pre-building completed successfully!"
