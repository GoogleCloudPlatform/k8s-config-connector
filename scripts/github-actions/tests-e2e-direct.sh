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

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"

cd ${REPO_ROOT}/

echo "Downloading envtest assets..."
export KUBEBUILDER_ASSETS=$(go run sigs.k8s.io/controller-runtime/tools/setup-envtest@latest use -p path)

echo "Running e2e tests samples for LoggingLogMetric direct reconciliation..."

KCC_USE_DIRECT_RECONCILERS=LoggingLogMetric \
GOLDEN_OBJECT_CHECKS=1 \
GOLDEN_REQUEST_CHECKS=1 \
E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=mock \
  go test -test.count=1 -timeout 600s -v ./tests/e2e -run 'TestAllInSeries/samples/linear-log-metric|TestAllInSeries/samples/exponential-log-metric|TestAllInSeries/samples/int-log-metric|TestAllInSeries/samples/explicit-log-metric'

echo "Running e2e tests fixtures for LoggingLogMetric direct reconciliation..."

KCC_USE_DIRECT_RECONCILERS=LoggingLogMetric \
GOLDEN_OBJECT_CHECKS=1 \
GOLDEN_REQUEST_CHECKS=1 \
E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=mock \
  go test -test.count=1 -timeout 600s -v ./tests/e2e -run 'TestAllInSeries/fixtures/explicitlogmetric|TestAllInSeries/fixtures/exponentiallogmetric|TestAllInSeries/fixtures/linearlogmetric'

echo "Running scenarios tests for LoggingLogMetric direct reconciliation..."

KCC_USE_DIRECT_RECONCILERS=LoggingLogMetric \
GOLDEN_REQUEST_CHECKS=1 E2E_KUBE_TARGET=envtest E2E_GCP_TARGET=mock RUN_E2E=1 \
  go test -test.count=1 -timeout 360s -v ./tests/e2e -run TestE2EScript/scenarios/direct