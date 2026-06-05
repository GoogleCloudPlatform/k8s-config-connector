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
  --api-version "dns.cnrm.cloud.google.com/v1beta1" \
  --resource "DNSManagedZone:ManagedZone" \
  --ignore-field "*:kind" \
  --require-field "ManagedZoneCloudLoggingConfig:enableLogging,ManagedZoneForwardingConfig:targetNameServers,ManagedZonePeeringConfig:targetNetwork,ManagedZoneServiceDirectoryConfig:namespace" \
  --output-file "${REPO_ROOT}/apis/dns/v1beta1/types.generated.go"

# Generate OpenAPI mappers
go run "${REPO_ROOT}/dev/tools/openapi-to-krm/cmd/generate-mapper/main.go" \
  --mapper "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1beta1/DNSManagedZoneSpec::google.golang.org/api/dns/v1/ManagedZone" \
  --mapper "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1beta1/DNSManagedZoneStatus::google.golang.org/api/dns/v1/ManagedZone" \
  --output-file "${REPO_ROOT}/pkg/controller/direct/dns/zz_generated.mappers.go"

cd ${REPO_ROOT}
dev/tasks/generate-crds
