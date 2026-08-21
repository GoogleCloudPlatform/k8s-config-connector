// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"google.golang.org/grpc"
	sdkgrpc "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
)

// RegisterServers registers each resource with the gRPC server.
func RegisterServers(s *grpc.Server) {
	sdkgrpc.RegisterComputeFirewallPolicyServiceServer(s, &FirewallPolicyServer{})
	sdkgrpc.RegisterComputeFirewallPolicyAssociationServiceServer(s, &FirewallPolicyAssociationServer{})
	sdkgrpc.RegisterComputeFirewallPolicyRuleServiceServer(s, &FirewallPolicyRuleServer{})
	sdkgrpc.RegisterComputeForwardingRuleServiceServer(s, &ForwardingRuleServer{})
	sdkgrpc.RegisterComputeInstanceServiceServer(s, &InstanceServer{})
	sdkgrpc.RegisterComputeInstanceGroupManagerServiceServer(s, &InstanceGroupManagerServer{})
	sdkgrpc.RegisterComputeInterconnectAttachmentServiceServer(s, &InterconnectAttachmentServer{})
	sdkgrpc.RegisterComputeNetworkServiceServer(s, &NetworkServer{})
	sdkgrpc.RegisterComputeNetworkFirewallPolicyServiceServer(s, &NetworkFirewallPolicyServer{})
	sdkgrpc.RegisterComputeNetworkFirewallPolicyAssociationServiceServer(s, &NetworkFirewallPolicyAssociationServer{})
	sdkgrpc.RegisterComputeNetworkFirewallPolicyRuleServiceServer(s, &NetworkFirewallPolicyRuleServer{})
	sdkgrpc.RegisterComputePacketMirroringServiceServer(s, &PacketMirroringServer{})
	sdkgrpc.RegisterComputeServiceAttachmentServiceServer(s, &ServiceAttachmentServer{})
	sdkgrpc.RegisterComputeSubnetworkServiceServer(s, &SubnetworkServer{})
	sdkgrpc.RegisterComputeVpnTunnelServiceServer(s, &VpnTunnelServer{})
	sdkgrpc.RegisterComputeAddressServiceServer(s, &AddressServer{})
	sdkgrpc.RegisterComputeAutoscalerServiceServer(s, &AutoscalerServer{})
	sdkgrpc.RegisterComputeBackendBucketServiceServer(s, &BackendBucketServer{})
	sdkgrpc.RegisterComputeBackendServiceServiceServer(s, &BackendServiceServer{})
	sdkgrpc.RegisterComputeDiskServiceServer(s, &DiskServer{})
	sdkgrpc.RegisterComputeFirewallServiceServer(s, &FirewallServer{})
	sdkgrpc.RegisterComputeHealthCheckServiceServer(s, &HealthCheckServer{})
	sdkgrpc.RegisterComputeHttpHealthCheckServiceServer(s, &HttpHealthCheckServer{})
	sdkgrpc.RegisterComputeHttpsHealthCheckServiceServer(s, &HttpsHealthCheckServer{})
	sdkgrpc.RegisterComputeImageServiceServer(s, &ImageServer{})
	sdkgrpc.RegisterComputeInstanceTemplateServiceServer(s, &InstanceTemplateServer{})
	sdkgrpc.RegisterComputeInterconnectServiceServer(s, &InterconnectServer{})
	sdkgrpc.RegisterComputeManagedSslCertificateServiceServer(s, &ManagedSslCertificateServer{})
	sdkgrpc.RegisterComputeNetworkEndpointServiceServer(s, &NetworkEndpointServer{})
	sdkgrpc.RegisterComputeNetworkEndpointGroupServiceServer(s, &NetworkEndpointGroupServer{})
	sdkgrpc.RegisterComputeNetworkPeeringServiceServer(s, &NetworkPeeringServer{})
	sdkgrpc.RegisterComputeReservationServiceServer(s, &ReservationServer{})
	sdkgrpc.RegisterComputeRouterServiceServer(s, &RouterServer{})
	sdkgrpc.RegisterComputeRouterInterfaceServiceServer(s, &RouterInterfaceServer{})
	sdkgrpc.RegisterComputeRouterNatServiceServer(s, &RouterNatServer{})
	sdkgrpc.RegisterComputeRouterPeerServiceServer(s, &RouterPeerServer{})
	sdkgrpc.RegisterComputeSnapshotServiceServer(s, &SnapshotServer{})
	sdkgrpc.RegisterComputeSslCertificateServiceServer(s, &SslCertificateServer{})
	sdkgrpc.RegisterComputeSslPolicyServiceServer(s, &SslPolicyServer{})
	sdkgrpc.RegisterComputeTargetHttpProxyServiceServer(s, &TargetHttpProxyServer{})
	sdkgrpc.RegisterComputeTargetHttpsProxyServiceServer(s, &TargetHttpsProxyServer{})
	sdkgrpc.RegisterComputeTargetPoolServiceServer(s, &TargetPoolServer{})
	sdkgrpc.RegisterComputeTargetSslProxyServiceServer(s, &TargetSslProxyServer{})
	sdkgrpc.RegisterComputeTargetVpnGatewayServiceServer(s, &TargetVpnGatewayServer{})
	sdkgrpc.RegisterComputeUrlMapServiceServer(s, &UrlMapServer{})
	sdkgrpc.RegisterComputeVpnGatewayServiceServer(s, &VpnGatewayServer{})
}
