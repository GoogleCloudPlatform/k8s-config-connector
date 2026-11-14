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

# We must compile Compute using a newer googleapis pin than the global default in apis/git.versions to include RolloutPlan.
PROTO_SHA="7e7f1e9533723bcf3821e681e564b27f92fd46f5"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT} "$@"

# --- v1alpha1 ---
${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource ComputeCrossSiteNetwork:CrossSiteNetwork \
  --resource ComputeInterconnectAttachmentGroup:InterconnectAttachmentGroup \
  --resource ComputeInterconnectGroup:InterconnectGroup \
  --include-skipped-output

${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --resource ComputeNetworkEdgeSecurityService:NetworkEdgeSecurityService \
  --resource ComputeNetworkAttachment:NetworkAttachment \
  --resource ComputeInterconnect:Interconnect \
  --resource ComputeFutureReservation:google.cloud.compute.v1beta.FutureReservation \
  --resource ComputeRegionPerInstanceConfig:PerInstanceConfig \
  --resource ComputeAutoscaler:Autoscaler \
  --resource ComputeBackendServiceSignedURLKey:SignedUrlKey \
  --resource ComputeRegionAutoscaler:Autoscaler \
  --resource ComputeOrganizationSecurityPolicy:SecurityPolicy \
  --resource ComputeOrganizationSecurityPolicyAssociation:SecurityPolicyAssociation \
  --resource ComputeOrganizationSecurityPolicyRule:FirewallPolicyRule \
  --resource ComputeNetworkEndpoint:NetworkEndpoint \
  --resource ComputeMachineImage:MachineImage \
  --resource ComputeRegionSSLPolicy:SslPolicy \
  --resource ComputeInstantSnapshot:InstantSnapshot \
  --resource ComputeHealthCheckService:HealthCheckService \
  --resource ComputeVMExtensionPolicy:VmExtensionPolicy \
  --resource ComputeNotificationEndpoint:NotificationEndpoint \
  --resource ComputeHealthAggregationPolicy:HealthAggregationPolicy \
  --resource ComputeGlobalVMExtensionPolicy:GlobalVmExtensionPolicy \
  --include-skipped-output

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --resource ComputeHealthCheckService:HealthCheckService \
  --resource ComputeHealthAggregationPolicy:HealthAggregationPolicy \
  --resource ComputeInterconnectAttachmentGroup:InterconnectAttachmentGroup \
  --resource ComputeInterconnectGroup:InterconnectGroup

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --resource ComputeCrossSiteNetwork:CrossSiteNetwork \
  --pattern ComputeCrossSiteNetwork="projects/{project}/global/crossSiteNetworks/{crossSiteNetwork}"

rm -f ${REPO_ROOT}/apis/compute/v1alpha1/computeregionautoscaler_types.go

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --resource ComputeVMExtensionPolicy:VmExtensionPolicy \
  --pattern ComputeVMExtensionPolicy=projects/{project}/zones/{zone}/vmExtensionPolicies/{vmextensionpolicy}

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --resource ComputeNotificationEndpoint:NotificationEndpoint

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --resource ComputeGlobalVMExtensionPolicy:GlobalVmExtensionPolicy \
  --pattern ComputeGlobalVMExtensionPolicy=projects/{project}/global/vmExtensionPolicies/{globalvmextensionpolicy}

${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --resource ComputeRolloutPlan:RolloutPlan \
  --include-skipped-output \
  --prepopulate-resource

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --resource ComputeRolloutPlan:RolloutPlan \
  --pattern ComputeRolloutPlan="projects/{project}/global/rolloutPlans/{rolloutPlan}"

${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --prepopulate-resource \
  --resource ComputeStoragePool:StoragePool

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1alpha1 \
  --resource ComputeStoragePool:StoragePool \
  --pattern ComputeStoragePool=projects/{project}/zones/{zone}/storagePools/{storagePool}
# --- v1beta1 ---
${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1beta1 \
  --resource ComputeBackendService:BackendService \
  --resource ComputeBackendBucket:BackendBucket \
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
  --resource ComputeProjectMetadata:Metadata \
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
  --resource ComputeManagedSSLCertificate:SslCertificate \
  --resource ComputeSubnetwork:Subnetwork \
  --resource ComputeTargetGRPCProxy:TargetGrpcProxy \
  --resource ComputeTargetHTTPSProxy:TargetHttpsProxy \
  --resource ComputeTargetHTTPProxy:TargetHttpProxy \
  --resource ComputeTargetInstance:TargetInstance \
  --resource ComputeTargetPool:TargetPool \
  --resource ComputeTargetSSLProxy:TargetSslProxy \
  --resource ComputeTargetTCPProxy:TargetTcpProxy \
  --resource ComputeTargetVPNGateway:TargetVpnGateway \
  --resource ComputeURLMap:UrlMap \
  --resource ComputeVPNGateway:VpnGateway \
  --resource ComputeVPNTunnel:VpnTunnel \
  --resource ComputeInterconnectAttachment:InterconnectAttachment \
  --resource ComputePacketMirroring:PacketMirroring \
  --resource ComputeServiceAttachment:ServiceAttachment \
  --include-skipped-output

${CONTROLLERBUILDER} generate-mapper \
  --proto-source-path ${PROTO_OUT} \
  --multiversion \
  --service google.cloud.compute.v1 \
  --api-version compute.cnrm.cloud.google.com/v1beta1 \
  --include-skipped-output \
  --prepopulate-resource \
  --skip-proto GlobalVmExtensionPolicy \
  --skip-proto GlobalVmExtensionPolicyExtensionPolicy \
  --skip-proto GlobalVmExtensionPolicyInstanceSelector \
  --skip-proto GlobalVmExtensionPolicyLabelSelector \
  --skip-proto GlobalVmExtensionPolicyList \
  --skip-proto GlobalVmExtensionPolicyRolloutOperation \
  --skip-proto GlobalVmExtensionPolicyRolloutOperationRolloutInput \
  --skip-proto GlobalVmExtensionPolicyRolloutOperationRolloutStatus \
  --skip-proto GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata \
  --skip-proto GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus

cd ${REPO_ROOT}
if [[ -z "${SKIP_GENERATE_CRDS:-}" ]]; then
  dev/tasks/generate-crds
fi

if [ -d "${REPO_ROOT}/pkg/controller/direct/compute" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/compute/
fi
