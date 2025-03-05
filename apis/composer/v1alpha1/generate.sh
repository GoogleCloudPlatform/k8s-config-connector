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


set -o errexit
set -o nounset
set -o pipefail
set -x

cd ${REPO_ROOT}/dev/tools/controllerbuilder

go run . generate-types     --service google.cloud.orchestration.airflow.service.v1     --api-version composer.cnrm.cloud.google.com/v1alpha1     --resource ComposerEnvironment:Environment

go run . generate-mapper     --service google.cloud.orchestration.airflow.service.v1     --api-version composer.cnrm.cloud.google.com/v1alpha1

go run . generate-controller     --service google.cloud.orchestration.airflow.service.v1     --api-version composer.cnrm.cloud.google.com/v1alpha1 ---resource ComposerEnvironment:Environment