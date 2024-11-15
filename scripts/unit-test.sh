#!/bin/bash
# Copyright 2022 Google LLC
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
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "${REPO_ROOT}"
make -C operator test
UNIT_TEST_PACKAGES=$(go list ./pkg/... ./cmd/... ./config/tests/...  ./scripts/resource-autogen/... ./tests/... ./mockgcp/... | grep -v tests/e2e)
if [ -z ${GITHUB_ACTION+x} ]; then
    go run gotest.tools/gotestsum@latest --format testname -- ${UNIT_TEST_PACKAGES} -coverprofile cover.out -count=1
else
    go run gotest.tools/gotestsum@latest --jsonfile unittest_result.json --format pkgname --format-hide-empty-pkg --format-hivis -- ${UNIT_TEST_PACKAGES} -coverprofile cover.out -count=1
fi
