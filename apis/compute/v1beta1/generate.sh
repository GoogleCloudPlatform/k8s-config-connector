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
source "${REPO_ROOT}/dev/tools/goimports.sh"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

./generate-proto.sh
go run . generate-types \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1beta1 \
  --resource ComputeExternalVPNGateway:ExternalVpnGateway \
  --resource ComputeFirewall:Firewall \
  --resource ComputeFirewallPolicy:FirewallPolicy \
  --resource ComputeFirewallPolicyAssociation:FirewallPolicyAssociation \
  --resource ComputeFirewallPolicyRule:FirewallPolicyRule \
  --resource ComputeForwardingRule:ForwardingRule \
  --resource ComputeHTTPHealthCheck:HTTPHealthCheck \
  --resource ComputeHTTPSHealthCheck:HTTPSHealthCheck \
  --resource ComputeImage:Image \
  --resource ComputeHealthCheck:HealthCheck \
  --resource ComputeInstance:Instance \
  --resource ComputeInstanceGroup:InstanceGroup \
  --resource ComputeInstanceGroupManager:InstanceGroupManager \
  --resource ComputeNetwork:Network \
  --resource ComputeNetworkEndpointGroup:NetworkEndpointGroup \
  --resource ComputeNetworkFirewallPolicy:FirewallPolicy \
  --resource ComputeNetworkPeering:NetworkPeering \
  --resource ComputeNodeGroup:NodeGroup \
  --resource ComputeNodeTemplate:NodeTemplate \
  --resource ComputeReservation:Reservation \
  --resource ComputeRoute:Route \
  --resource ComputeRouter:Router \
  --resource ComputeRouterInterface:RouterInterface \
  --resource ComputeRouterNAT:RouterNat \
  --resource ComputeResourcePolicy:ResourcePolicy \
  --resource ComputeSecurityPolicy:SecurityPolicy \
  --resource ComputeSnapshot:Snapshot \
  --resource ComputeSSLPolicy:SslPolicy \
  --resource ComputeSSLCertificate:SslCertificate \
  --resource ComputeSubnetwork:Subnetwork \
  --resource ComputeTargetHTTPSProxy:TargetHttpsProxy \
  --resource ComputeTargetPool:TargetPool \
  --resource ComputeTargetSSLProxy:TargetSslProxy \
  --resource ComputeTargetTcpProxy:TargetTcpProxy \
  --resource ComputeURLMap:UrlMap \
  --resource ComputeInterconnectAttachment:InterconnectAttachment \
  --include-skipped-output

go run . generate-mapper \
  --multiversion \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1beta1 \
  --include-skipped-output

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/compute/
