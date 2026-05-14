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
  --service google.cloud.networkmanagement.v1 \
  --include-skipped-output \
  --api-version networkmanagement.cnrm.cloud.google.com/v1alpha1  \
  --resource NetworkManagementConnectivityTest:ConnectivityTest \
  --proto-override google.cloud.networkmanagement.v1.AbortInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.AppEngineVersionInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.CloudFunctionInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.CloudRunRevisionInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.CloudSQLInstanceInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.DeliverInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.DirectVpcEgressConnectionInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.DropInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.EdgeLocation:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.Endpoint:forwarding_rule_target=output-only \
  --proto-override google.cloud.networkmanagement.v1.Endpoint:load_balancer_id=output-only \
  --proto-override google.cloud.networkmanagement.v1.Endpoint:load_balancer_type=output-only \
  --proto-override google.cloud.networkmanagement.v1.EndpointInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.FirewallInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.ForwardInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.ForwardingRuleInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.GKEMasterInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.GkePodInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.GoogleServiceInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.InstanceInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.IpMasqueradingSkippedInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.LatencyDistribution:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.LatencyPercentile:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.LoadBalancerBackend:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.LoadBalancerBackendInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.LoadBalancerInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.NatInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.NetworkInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.ProbingAbortCause:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.ProbingDetails:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.ProbingDetails.EdgeLocation:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.ProxyConnectionInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.ReachabilityDetails:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.RedisClusterInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.RedisInstanceInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.RouteInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.ServerlessExternalConnectionInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.ServerlessNegInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.SingleEdgeResponse:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.Step:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.StorageBucketInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.Trace:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.VpcConnectorInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.VpnGatewayInfo:*=output-only \
  --proto-override google.cloud.networkmanagement.v1.VpnTunnelInfo:*=output-only \

go run . generate-mapper \
  --service google.cloud.networkmanagement.v1 \
  --api-version networkmanagement.cnrm.cloud.google.com/v1alpha1

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/networkmanagement/
