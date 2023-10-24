#!/bin/bash
# Copyright 2023 Google LLC
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

echo "Running setup-envtest"
export KUBEBUILDER_ASSETS=$(go run sigs.k8s.io/controller-runtime/tools/setup-envtest@master use -p path)

echo "Running (legacy) unit tests with mockgcp..."
export E2E_GCP_TARGET=mock
# We don't (yet) have great coverage, so we can only run a few tests
go test -tags integration ./pkg/controller/dynamic/... -timeout 1800s -test.count=1 -v -run 'TestCreateNoChangeUpdateDelete/basic-service$'
