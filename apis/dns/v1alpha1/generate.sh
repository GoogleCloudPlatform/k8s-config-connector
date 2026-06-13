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

# Generate KRM Go types from the OpenAPI / Discovery API JSON specification, ignoring all 'kind' fields and setting required fields
go run "${REPO_ROOT}/dev/tools/openapi-to-krm/main.go" \
  --schema-file "${REPO_ROOT}/apis/dns/v1beta1/dns-api.json" \
  --api-version "dns.cnrm.cloud.google.com/v1alpha1" \
  --resource "DNSResponsePolicy:ResponsePolicy" \
  --resource "DNSResponsePolicyRule:ResponsePolicyRule" \
  --ignore-field "*:kind" \
  --output-file "${REPO_ROOT}/apis/dns/v1alpha1/types.generated.go"

# Generate OpenAPI mappers
go run "${REPO_ROOT}/dev/tools/openapi-to-krm/cmd/generate-mapper/main.go" \
  --mapper "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1alpha1/DNSResponsePolicySpec::google.golang.org/api/dns/v1/ResponsePolicy" \
  --mapper "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1alpha1/DNSResponsePolicyStatus::google.golang.org/api/dns/v1/ResponsePolicy" \
  --mapper "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1alpha1/DNSResponsePolicyRuleSpec::google.golang.org/api/dns/v1/ResponsePolicyRule" \
  --mapper "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1alpha1/DNSResponsePolicyRuleStatus::google.golang.org/api/dns/v1/ResponsePolicyRule" \
  --output-file "${REPO_ROOT}/pkg/controller/direct/dns/zz_generated.v1alpha1.mappers.go"

cd ${REPO_ROOT}
dev/tasks/generate-crds
