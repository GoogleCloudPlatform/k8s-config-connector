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
	sdkgrpc "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
)

// RegisterServers registers each resource with the gRPC server.
func RegisterServers(s *grpc.Server) {
	sdkgrpc.RegisterComputeBetaFirewallPolicyServiceServer(s, &FirewallPolicyServer{})
	sdkgrpc.RegisterComputeBetaFirewallPolicyAssociationServiceServer(s, &FirewallPolicyAssociationServer{})
	sdkgrpc.RegisterComputeBetaFirewallPolicyRuleServiceServer(s, &FirewallPolicyRuleServer{})
	sdkgrpc.RegisterComputeBetaForwardingRuleServiceServer(s, &ForwardingRuleServer{})
	sdkgrpc.RegisterComputeBetaInstanceServiceServer(s, &InstanceServer{})
	sdkgrpc.RegisterComputeBetaInstanceGroupManagerServiceServer(s, &InstanceGroupManagerServer{})
	sdkgrpc.RegisterComputeBetaInterconnectAttachmentServiceServer(s, &InterconnectAttachmentServer{})
	sdkgrpc.RegisterComputeBetaNetworkServiceServer(s, &NetworkServer{})
	sdkgrpc.RegisterComputeBetaNetworkFirewallPolicyServiceServer(s, &NetworkFirewallPolicyServer{})
	sdkgrpc.RegisterComputeBetaNetworkFirewallPolicyAssociationServiceServer(s, &NetworkFirewallPolicyAssociationServer{})
	sdkgrpc.RegisterComputeBetaNetworkFirewallPolicyRuleServiceServer(s, &NetworkFirewallPolicyRuleServer{})
	sdkgrpc.RegisterComputeBetaPacketMirroringServiceServer(s, &PacketMirroringServer{})
	sdkgrpc.RegisterComputeBetaServiceAttachmentServiceServer(s, &ServiceAttachmentServer{})
	sdkgrpc.RegisterComputeBetaSubnetworkServiceServer(s, &SubnetworkServer{})
	sdkgrpc.RegisterComputeBetaVpnTunnelServiceServer(s, &VpnTunnelServer{})
	sdkgrpc.RegisterComputeBetaAddressServiceServer(s, &AddressServer{})
	sdkgrpc.RegisterComputeBetaAutoscalerServiceServer(s, &AutoscalerServer{})
	sdkgrpc.RegisterComputeBetaBackendBucketServiceServer(s, &BackendBucketServer{})
	sdkgrpc.RegisterComputeBetaBackendServiceServiceServer(s, &BackendServiceServer{})
	sdkgrpc.RegisterComputeBetaDiskServiceServer(s, &DiskServer{})
	sdkgrpc.RegisterComputeBetaFirewallServiceServer(s, &FirewallServer{})
	sdkgrpc.RegisterComputeBetaHealthCheckServiceServer(s, &HealthCheckServer{})
	sdkgrpc.RegisterComputeBetaHttpHealthCheckServiceServer(s, &HttpHealthCheckServer{})
	sdkgrpc.RegisterComputeBetaHttpsHealthCheckServiceServer(s, &HttpsHealthCheckServer{})
	sdkgrpc.RegisterComputeBetaImageServiceServer(s, &ImageServer{})
	sdkgrpc.RegisterComputeBetaInstanceTemplateServiceServer(s, &InstanceTemplateServer{})
	sdkgrpc.RegisterComputeBetaInterconnectServiceServer(s, &InterconnectServer{})
	sdkgrpc.RegisterComputeBetaManagedSslCertificateServiceServer(s, &ManagedSslCertificateServer{})
	sdkgrpc.RegisterComputeBetaNetworkEndpointServiceServer(s, &NetworkEndpointServer{})
	sdkgrpc.RegisterComputeBetaNetworkEndpointGroupServiceServer(s, &NetworkEndpointGroupServer{})
	sdkgrpc.RegisterComputeBetaNetworkPeeringServiceServer(s, &NetworkPeeringServer{})
	sdkgrpc.RegisterComputeBetaReservationServiceServer(s, &ReservationServer{})
	sdkgrpc.RegisterComputeBetaRouterServiceServer(s, &RouterServer{})
	sdkgrpc.RegisterComputeBetaRouterInterfaceServiceServer(s, &RouterInterfaceServer{})
	sdkgrpc.RegisterComputeBetaRouterNatServiceServer(s, &RouterNatServer{})
	sdkgrpc.RegisterComputeBetaRouterPeerServiceServer(s, &RouterPeerServer{})
	sdkgrpc.RegisterComputeBetaSnapshotServiceServer(s, &SnapshotServer{})
	sdkgrpc.RegisterComputeBetaSslCertificateServiceServer(s, &SslCertificateServer{})
	sdkgrpc.RegisterComputeBetaSslPolicyServiceServer(s, &SslPolicyServer{})
	sdkgrpc.RegisterComputeBetaTargetHttpProxyServiceServer(s, &TargetHttpProxyServer{})
	sdkgrpc.RegisterComputeBetaTargetHttpsProxyServiceServer(s, &TargetHttpsProxyServer{})
	sdkgrpc.RegisterComputeBetaTargetPoolServiceServer(s, &TargetPoolServer{})
	sdkgrpc.RegisterComputeBetaTargetSslProxyServiceServer(s, &TargetSslProxyServer{})
	sdkgrpc.RegisterComputeBetaTargetVpnGatewayServiceServer(s, &TargetVpnGatewayServer{})
	sdkgrpc.RegisterComputeBetaUrlMapServiceServer(s, &UrlMapServer{})
	sdkgrpc.RegisterComputeBetaVpnGatewayServiceServer(s, &VpnGatewayServer{})
}
