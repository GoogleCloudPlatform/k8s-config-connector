#!/bin/bash
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


set -e
set -x

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

# example usage of inserting a field
go run . update-types insert \
    --parent "google.monitoring.dashboard.v1.Dashboard" \
    --field "row_layout" \
    --api-dir ${REPO_ROOT}/apis/monitoring/v1beta1

# example usage of syncing a message with all of its dependencies from proto package
go run . update-types sync \
    --service google.cloud.bigquery.datatransfer.v1 \
    --api-version bigquerydatatransfer.cnrm.cloud.google.com/v1beta1 \
    --legacy-mode  # this is a flag to indicate that the resource is previously generated with KRM fields without proto annotations

# Fix up formatting
${REPO_ROOT}/dev/tasks/fix-gofmt