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

PROTO_SHA="cdc919ff596e263f2cc55a9780d2f74633da1ced"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

# --- v1alpha1 ---
${CONTROLLERBUILDER} generate-types \
  --service google.cloud.networksecurity.v1 \
  --api-version networksecurity.cnrm.cloud.google.com/v1alpha1 \
  --resource NetworkSecurityBackendAuthenticationConfig:BackendAuthenticationConfig \
  --resource NetworkSecurityInterceptDeployment:InterceptDeployment \
  --resource NetworkSecurityAddressGroup:AddressGroup \
  --resource NetworkSecurityInterceptEndpointGroup:InterceptEndpointGroup \
  --resource NetworkSecurityMirroringDeployment:MirroringDeployment \
  --resource NetworkSecurityMirroringDeploymentGroup:MirroringDeploymentGroup \
  --resource NetworkSecurityMirroringEndpointGroup:MirroringEndpointGroup \
  --resource NetworkSecurityMirroringEndpointGroupAssociation:MirroringEndpointGroupAssociation \
  --resource NetworkSecuritySACRealm:SACRealm \
  --resource NetworkSecuritySecurityProfile:SecurityProfile \
  --resource NetworkSecuritySecurityProfileGroup:SecurityProfileGroup \
  --resource NetworkSecurityFirewallEndpointAssociation:FirewallEndpointAssociation \
  --resource NetworkSecurityGatewaySecurityPolicy:GatewaySecurityPolicy \
  --resource NetworkSecurityGatewaySecurityPolicyRule:GatewaySecurityPolicyRule \
  --resource NetworkSecurityTLSInspectionPolicy:TlsInspectionPolicy \
  --resource NetworkSecurityAuthzPolicy:AuthzPolicy \
  --resource NetworkSecurityFirewallEndpoint:FirewallEndpoint \
  --resource NetworkSecurityDNSThreatDetector:DnsThreatDetector \
  --resource NetworkSecurityURLList:UrlList \
  --proto-source-path ${PROTO_OUT} \
  --prepopulate-spec \
  --emit-required-from-proto \
  --emit-plural-acronyms \
  --emit-message-maps \
  --place-server-set-fields \
  --detect-output-only-in-comments \
  --emit-parent-refs \
  --emit-sibling-refs \
  --emit-reference-hints \
  --emit-source-links

${CONTROLLERBUILDER} generate-types \
  --service google.cloud.networksecurity.v1alpha1 \
  --api-version networksecurity.cnrm.cloud.google.com/v1alpha1 \
  --resource NetworkSecurityPartnerSSERealm:PartnerSSERealm \
  --resource NetworkSecurityPartnerSSEGateway:PartnerSSEGateway \
  --resource NetworkSecurityTLSInspectionPolicy:TlsInspectionPolicy \
  --prune-unused-types=false \
  --proto-source-path ${PROTO_OUT} \
  --prepopulate-spec \
  --emit-required-from-proto \
  --emit-plural-acronyms \
  --emit-message-maps \
  --place-server-set-fields \
  --detect-output-only-in-comments \
  --emit-parent-refs \
  --emit-sibling-refs \
  --emit-reference-hints \
  --emit-source-links

${CONTROLLERBUILDER} generate-identity \
  --service google.cloud.networksecurity.v1 \
  --api-version networksecurity.cnrm.cloud.google.com/v1alpha1 \
  --resource NetworkSecurityGatewaySecurityPolicyRule:GatewaySecurityPolicyRule \
  --proto-source-path ${PROTO_OUT}

# --- v1beta1 ---
${CONTROLLERBUILDER} generate-types \
  --service google.cloud.networksecurity.v1beta1 \
  --api-version networksecurity.cnrm.cloud.google.com/v1beta1 \
  --resource NetworkSecurityAuthorizationPolicy:AuthorizationPolicy \
  --resource NetworkSecurityClientTLSPolicy:ClientTlsPolicy \
  --proto-source-path ${PROTO_OUT}

${CONTROLLERBUILDER} generate-mapper \
  --service google.cloud.networksecurity.v1,google.cloud.networksecurity.v1beta1 \
  --api-version "networksecurity.cnrm.cloud.google.com/v1beta1" \
  --proto-source-path ${PROTO_OUT} \
  --multiversion \
  --emit-plural-acronyms \
  --emit-message-maps

cd ${REPO_ROOT}
dev/tasks/generate-crds

if [ -d "${REPO_ROOT}/pkg/controller/direct/networksecurity" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/networksecurity/
fi
