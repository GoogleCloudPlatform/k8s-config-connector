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

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

go run . generate-types \
  --service google.firestore.admin.v1 \
  --api-version firestore.cnrm.cloud.google.com/v1alpha1  \
  --resource FirestoreDocument:google.firestore.v1.Document \
  --resource FirestoreField:Field \
  --resource FirestoreBackupSchedule:BackupSchedule

go run . generate-mapper \
  --multiversion \
  --service google.firestore.admin.v1 \
  --service google.firestore.v1 \
  --api-version firestore.cnrm.cloud.google.com/v1alpha1

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w pkg/controller/direct/firestore/
