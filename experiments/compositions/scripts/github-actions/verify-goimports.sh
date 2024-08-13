#!/bin/bash

# Copyright 2022 The Kubernetes Authors.
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


# CI script to verify go imports

set -o errexit
set -o nounset
set -o pipefail

# cd to the repo root
files=$(go run golang.org/x/tools/cmd/goimports -format-only -l experiments/compositions/.)
# Sadly goimports doesn't use exit codes
if [[ -n "${files}" ]]; then
    echo "::error ::goimports should be run on these files:"
    echo "${files}"
    exit 1
fi