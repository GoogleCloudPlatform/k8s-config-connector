#!/bin/bash
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


set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
source "${REPO_ROOT}/dev/tools/goimports.sh"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

# We need the newer googleapis version for ReservationGroup, so we force-generate the proto pb file even if SKIP_GENERATE_PROTOS is set.
SKIP_GENERATE_PROTOS= ./generate-proto.sh 2b625c91510a2e8320a778bc88af8b65bc4a19a2

go run . generate-types \
    --service google.cloud.bigquery.reservation.v1 \
    --api-version "bigqueryreservation.cnrm.cloud.google.com/v1alpha1" \
    --include-skipped-output \
    --resource BigQueryReservationCapacityCommitment:CapacityCommitment \
    --resource BigQueryReservationReservationGroup:ReservationGroup

go run . generate-types \
    --service google.cloud.bigquery.reservation.v1 \
    --api-version "bigqueryreservation.cnrm.cloud.google.com/v1beta1" \
    --include-skipped-output \
    --resource BigQueryReservationReservation:Reservation \
    --resource BigQueryReservationAssignment:Assignment

go run . generate-mapper \
    --multiversion \
    --service google.cloud.bigquery.reservation.v1 \
    --api-version "bigqueryreservation.cnrm.cloud.google.com/v1beta1" \
    --include-skipped-output

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/bigqueryreservation/

# Restore the global googleapis.pb to the stable version defined in git.versions so subsequent API generations are not affected.
DEFAULT_GOOGLE_API_VERSION=$(grep https://github.com/googleapis/googleapis ${REPO_ROOT}/apis/git.versions | awk '{print $2}')
cp "${REPO_ROOT}/.build/googleapis-${DEFAULT_GOOGLE_API_VERSION}.pb" "${REPO_ROOT}/.build/googleapis.pb"
