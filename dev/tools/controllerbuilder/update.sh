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
export GEMINI_API_KEY=${GEMINI_API_KEY}

# example usage
go run . update-types \
    --parent-message-full-name "google.monitoring.dashboard.v1.Dashboard" \
    --new-field "row_layout" \
    --api-dir ${REPO_ROOT}/apis/monitoring/v1beta1 \
    --ignored-fields "google.monitoring.dashboard.v1.PickTimeSeriesFilter.interval"
