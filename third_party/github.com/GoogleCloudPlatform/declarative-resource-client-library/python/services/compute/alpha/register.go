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
	sdkgrpc "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/alpha/compute_alpha_go_proto"
)

// RegisterServers registers each resource with the gRPC server.
func RegisterServers(s *grpc.Server) {
	sdkgrpc.RegisterComputeAlphaFirewallPolicyServiceServer(s, &FirewallPolicyServer{})
	sdkgrpc.RegisterComputeAlphaFirewallPolicyAssociationServiceServer(s, &FirewallPolicyAssociationServer{})
	sdkgrpc.RegisterComputeAlphaFirewallPolicyRuleServiceServer(s, &FirewallPolicyRuleServer{})
	sdkgrpc.RegisterComputeAlphaForwardingRuleServiceServer(s, &ForwardingRuleServer{})
	sdkgrpc.RegisterComputeAlphaInstanceServiceServer(s, &InstanceServer{})
	sdkgrpc.RegisterComputeAlphaInstanceGroupManagerServiceServer(s, &InstanceGroupManagerServer{})
	sdkgrpc.RegisterComputeAlphaInterconnectAttachmentServiceServer(s, &InterconnectAttachmentServer{})
	sdkgrpc.RegisterComputeAlphaNetworkServiceServer(s, &NetworkServer{})
	sdkgrpc.RegisterComputeAlphaNetworkFirewallPolicyServiceServer(s, &NetworkFirewallPolicyServer{})
	sdkgrpc.RegisterComputeAlphaNetworkFirewallPolicyAssociationServiceServer(s, &NetworkFirewallPolicyAssociationServer{})
	sdkgrpc.RegisterComputeAlphaNetworkFirewallPolicyRuleServiceServer(s, &NetworkFirewallPolicyRuleServer{})
	sdkgrpc.RegisterComputeAlphaPacketMirroringServiceServer(s, &PacketMirroringServer{})
	sdkgrpc.RegisterComputeAlphaServiceAttachmentServiceServer(s, &ServiceAttachmentServer{})
	sdkgrpc.RegisterComputeAlphaSubnetworkServiceServer(s, &SubnetworkServer{})
	sdkgrpc.RegisterComputeAlphaVpnTunnelServiceServer(s, &VpnTunnelServer{})
	sdkgrpc.RegisterComputeAlphaAddressServiceServer(s, &AddressServer{})
	sdkgrpc.RegisterComputeAlphaAutoscalerServiceServer(s, &AutoscalerServer{})
	sdkgrpc.RegisterComputeAlphaBackendBucketServiceServer(s, &BackendBucketServer{})
	sdkgrpc.RegisterComputeAlphaBackendServiceServiceServer(s, &BackendServiceServer{})
	sdkgrpc.RegisterComputeAlphaDiskServiceServer(s, &DiskServer{})
	sdkgrpc.RegisterComputeAlphaFirewallServiceServer(s, &FirewallServer{})
	sdkgrpc.RegisterComputeAlphaHealthCheckServiceServer(s, &HealthCheckServer{})
	sdkgrpc.RegisterComputeAlphaHttpHealthCheckServiceServer(s, &HttpHealthCheckServer{})
	sdkgrpc.RegisterComputeAlphaHttpsHealthCheckServiceServer(s, &HttpsHealthCheckServer{})
	sdkgrpc.RegisterComputeAlphaImageServiceServer(s, &ImageServer{})
	sdkgrpc.RegisterComputeAlphaInstanceTemplateServiceServer(s, &InstanceTemplateServer{})
	sdkgrpc.RegisterComputeAlphaInterconnectServiceServer(s, &InterconnectServer{})
	sdkgrpc.RegisterComputeAlphaManagedSslCertificateServiceServer(s, &ManagedSslCertificateServer{})
	sdkgrpc.RegisterComputeAlphaNetworkEndpointServiceServer(s, &NetworkEndpointServer{})
	sdkgrpc.RegisterComputeAlphaNetworkEndpointGroupServiceServer(s, &NetworkEndpointGroupServer{})
	sdkgrpc.RegisterComputeAlphaNetworkPeeringServiceServer(s, &NetworkPeeringServer{})
	sdkgrpc.RegisterComputeAlphaReservationServiceServer(s, &ReservationServer{})
	sdkgrpc.RegisterComputeAlphaRouterServiceServer(s, &RouterServer{})
	sdkgrpc.RegisterComputeAlphaRouterInterfaceServiceServer(s, &RouterInterfaceServer{})
	sdkgrpc.RegisterComputeAlphaRouterNatServiceServer(s, &RouterNatServer{})
	sdkgrpc.RegisterComputeAlphaRouterPeerServiceServer(s, &RouterPeerServer{})
	sdkgrpc.RegisterComputeAlphaSnapshotServiceServer(s, &SnapshotServer{})
	sdkgrpc.RegisterComputeAlphaSslCertificateServiceServer(s, &SslCertificateServer{})
	sdkgrpc.RegisterComputeAlphaSslPolicyServiceServer(s, &SslPolicyServer{})
	sdkgrpc.RegisterComputeAlphaTargetHttpProxyServiceServer(s, &TargetHttpProxyServer{})
	sdkgrpc.RegisterComputeAlphaTargetHttpsProxyServiceServer(s, &TargetHttpsProxyServer{})
	sdkgrpc.RegisterComputeAlphaTargetPoolServiceServer(s, &TargetPoolServer{})
	sdkgrpc.RegisterComputeAlphaTargetSslProxyServiceServer(s, &TargetSslProxyServer{})
	sdkgrpc.RegisterComputeAlphaTargetVpnGatewayServiceServer(s, &TargetVpnGatewayServer{})
	sdkgrpc.RegisterComputeAlphaUrlMapServiceServer(s, &UrlMapServer{})
	sdkgrpc.RegisterComputeAlphaVpnGatewayServiceServer(s, &VpnGatewayServer{})
}
