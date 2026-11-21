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

# Pin a googleapis SHA that contains the ReservationGroup definition
PROTO_SHA="2193a2bfcecb92b92aad7a4d81baa428cafd7dfd"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

if [[ -z "${SKIP_GENERATE_PROTOS:-}" ]]; then
  ./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}
fi

${CONTROLLERBUILDER} generate-types \
    --service google.cloud.bigquery.reservation.v1 \
    --api-version "bigqueryreservation.cnrm.cloud.google.com/v1alpha1" \
    --include-skipped-output \
    --resource BigQueryReservationCapacityCommitment:CapacityCommitment \
    --resource BigQueryReservationReservationGroup:ReservationGroup \
    --resource BigQueryReservationBiReservation:BiReservation \
    --proto-source-path ${PROTO_OUT}

${CONTROLLERBUILDER} generate-types \
    --service google.cloud.bigquery.reservation.v1 \
    --api-version "bigqueryreservation.cnrm.cloud.google.com/v1beta1" \
    --include-skipped-output \
    --resource BigQueryReservationReservation:Reservation \
    --resource BigQueryReservationAssignment:Assignment \
    --proto-source-path ${PROTO_OUT}

${CONTROLLERBUILDER} generate-mapper \
    --multiversion \
    --service google.cloud.bigquery.reservation.v1 \
    --api-version "bigqueryreservation.cnrm.cloud.google.com/v1beta1" \
    --include-skipped-output \
    --proto-source-path ${PROTO_OUT}

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/bigqueryreservation/
