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

CONTROLLERBUILDER="${CONTROLLERBUILDER:-}"
if [[ -z "${CONTROLLERBUILDER}" ]]; then
  if [[ -x "${REPO_ROOT}/bin/controllerbuilder" ]]; then
    CONTROLLERBUILDER="${REPO_ROOT}/bin/controllerbuilder"
  else
    CONTROLLERBUILDER="go run ${REPO_ROOT}/dev/tools/controllerbuilder"
  fi
fi
source "${REPO_ROOT}/dev/tools/goimports.sh"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

PROTO_SHA="9b719a5c153f3ec7a14d5d45d4da720410c7b943"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

# Unset SKIP_GENERATE_PROTOS so this specific script fetches the newer proto
OLD_SKIP_GENERATE_PROTOS="${SKIP_GENERATE_PROTOS:-}"
unset SKIP_GENERATE_PROTOS

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

# Restore SKIP_GENERATE_PROTOS
if [[ -n "${OLD_SKIP_GENERATE_PROTOS}" ]]; then
  export SKIP_GENERATE_PROTOS="${OLD_SKIP_GENERATE_PROTOS}"
fi

${CONTROLLERBUILDER} generate-types \
    --proto-source-path ${PROTO_OUT} \
    --service google.cloud.bigquery.reservation.v1 \
    --api-version "bigqueryreservation.cnrm.cloud.google.com/v1alpha1" \
    --include-skipped-output \
    --resource BigQueryReservationCapacityCommitment:CapacityCommitment \
    --resource BigQueryReservationReservationGroup:ReservationGroup \
    --resource BigQueryReservationBiReservation:BiReservation

${CONTROLLERBUILDER} generate-types \
    --proto-source-path ${PROTO_OUT} \
    --service google.cloud.bigquery.reservation.v1 \
    --api-version "bigqueryreservation.cnrm.cloud.google.com/v1beta1" \
    --include-skipped-output \
    --resource BigQueryReservationReservation:Reservation \
    --resource BigQueryReservationAssignment:Assignment

${CONTROLLERBUILDER} generate-mapper \
    --proto-source-path ${PROTO_OUT} \
    --multiversion \
    --service google.cloud.bigquery.reservation.v1 \
    --api-version "bigqueryreservation.cnrm.cloud.google.com/v1beta1" \
    --include-skipped-output

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/bigqueryreservation/
