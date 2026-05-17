// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockcompute

import (
	"context"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/compute/apiv1/computepb"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked compute service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	*computeOperations
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment:   env,
		storage:           storage,
		computeOperations: newComputeOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	// service attachment has host "www.googleapis.com"
	return []string{"compute.googleapis.com", "www.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterFutureReservationsServer(grpcServer, &FutureReservationsV1{MockService: s})

	pb.RegisterBackendBucketsServer(grpcServer, &backendBuckets{MockService: s})

	pb.RegisterExternalVpnGatewaysServer(grpcServer, &externalVPNGateways{MockService: s})

	pb.RegisterNetworkEdgeSecurityServicesServer(grpcServer, &networkEdgeSecurityServicesV1{MockService: s})

	pb.RegisterNetworksServer(grpcServer, &NetworksV1{MockService: s})
	pb.RegisterNetworkAttachmentsServer(grpcServer, &networkAttachmentsV1{MockService: s})
	pb.RegisterSubnetworksServer(grpcServer, &SubnetsV1{MockService: s})
	pb.RegisterVpnGatewaysServer(grpcServer, &VPNGatewaysV1{MockService: s})
	pb.RegisterTargetVpnGatewaysServer(grpcServer, &TargetVpnGatewaysV1{MockService: s})
	pb.RegisterTargetGrpcProxiesServer(grpcServer, &TargetGrpcProxyV1{MockService: s})

	pb.RegisterTargetHttpProxiesServer(grpcServer, &GlobalTargetHTTPProxiesV1{MockService: s})
	pb.RegisterRegionTargetHttpProxiesServer(grpcServer, &RegionalTargetHTTPProxiesV1{MockService: s})

	pb.RegisterTargetHttpsProxiesServer(grpcServer, &GlobalTargetHTTPSProxiesV1{MockService: s})
	pb.RegisterRegionTargetHttpsProxiesServer(grpcServer, &RegionalTargetHTTPSProxiesV1{MockService: s})

	pb.RegisterUrlMapsServer(grpcServer, &GlobalURLMapsV1{MockService: s})
	pb.RegisterRegionUrlMapsServer(grpcServer, &RegionalURLMapsV1{MockService: s})

	pb.RegisterRegionHealthChecksServer(grpcServer, &RegionalHealthCheckV1{MockService: s})
	pb.RegisterHealthChecksServer(grpcServer, &GlobalHealthCheckV1{MockService: s})

	pb.RegisterBackendServicesServer(grpcServer, &GlobalBackendServicesV1{MockService: s})
	pb.RegisterRegionBackendServicesServer(grpcServer, &RegionalBackendServicesV1{MockService: s})

	pb.RegisterDisksServer(grpcServer, &DisksV1{MockService: s})
	pb.RegisterRegionDisksServer(grpcServer, &RegionalDisksV1{MockService: s})

	pb.RegisterRegionOperationsServer(grpcServer, &RegionalOperationsV1{MockService: s})
	pb.RegisterZoneOperationsServer(grpcServer, &ZonalOperationsV1{MockService: s})
	pb.RegisterGlobalOperationsServer(grpcServer, &GlobalOperationsV1{MockService: s})
	pb.RegisterGlobalOrganizationOperationsServer(grpcServer, &GlobalOrganizationOperationsV1{MockService: s})

	pb.RegisterNodeGroupsServer(grpcServer, &NodeGroupsV1{MockService: s})
	pb.RegisterNodeTemplatesServer(grpcServer, &NodeTemplatesV1{MockService: s})

	pb.RegisterAddressesServer(grpcServer, &RegionalAddressesV1{MockService: s})
	pb.RegisterGlobalAddressesServer(grpcServer, &GlobalAddressesV1{MockService: s})
	pb.RegisterSslCertificatesServer(grpcServer, &GlobalSSLCertificatesV1{MockService: s})
	pb.RegisterRegionSslCertificatesServer(grpcServer, &RegionalSSLCertificatesV1{MockService: s})
	pb.RegisterSslPoliciesServer(grpcServer, &GlobalSslPolicyV1{MockService: s})
	pb.RegisterTargetSslProxiesServer(grpcServer, &TargetSslProxyV1{MockService: s})
	pb.RegisterTargetTcpProxiesServer(grpcServer, &GlobalTargetTcpProxyV1{MockService: s})
	pb.RegisterRegionTargetTcpProxiesServer(grpcServer, &RegionalTargetTcpProxyV1{MockService: s})

	pb.RegisterRegionNetworkEndpointGroupsServer(grpcServer, &RegionNetworkEndpointGroupV1{MockService: s})

	pb.RegisterRoutesServer(grpcServer, &RoutesV1{MockService: s})

	pb.RegisterServiceAttachmentsServer(grpcServer, &RegionalServiceAttachmentV1{MockService: s})

	pb.RegisterFirewallPoliciesServer(grpcServer, &FirewallPoliciesV1{MockService: s})

	pb.RegisterGlobalForwardingRulesServer(grpcServer, &GlobalForwardingRulesV1{MockService: s})
	pb.RegisterForwardingRulesServer(grpcServer, &RegionalForwardingRulesV1{MockService: s})

	pb.RegisterImagesServer(grpcServer, &ImagesV1{MockService: s})

	pb.RegisterInstancesServer(grpcServer, &InstancesV1{MockService: s})
	pb.RegisterInstanceTemplatesServer(grpcServer, &InstanceTemplatesV1{MockService: s})

	pb.RegisterInstanceGroupManagersServer(grpcServer, &instanceGroupManagers{MockService: s})

	pb.RegisterZonesServer(grpcServer, &ZonesV1{MockService: s})
	pb.RegisterReservationsServer(grpcServer, &ReservationsV1{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	grpcMux.AddService(pb.NewFutureReservationsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.FutureReservations"))
	grpcMux.AddService(pb.NewBackendBucketsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.BackendBuckets"))
	grpcMux.AddService(pb.NewExternalVpnGatewaysClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.ExternalVpnGateways"))
	grpcMux.AddService(pb.NewNetworkEdgeSecurityServicesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.NetworkEdgeSecurityServices"))
	grpcMux.AddService(pb.NewNetworksClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.Networks"))
	grpcMux.AddService(pb.NewNetworkAttachmentsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.NetworkAttachments"))
	grpcMux.AddService(pb.NewSubnetworksClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.Subnetworks"))
	grpcMux.AddService(pb.NewVpnGatewaysClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.VpnGateways"))
	grpcMux.AddService(pb.NewTargetVpnGatewaysClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.TargetVpnGateways"))
	grpcMux.AddService(pb.NewTargetGrpcProxiesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.TargetGrpcProxies"))
	grpcMux.AddService(pb.NewTargetHttpProxiesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.TargetHttpProxies"))
	grpcMux.AddService(pb.NewRegionTargetHttpProxiesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.RegionTargetHttpProxies"))
	grpcMux.AddService(pb.NewTargetHttpsProxiesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.TargetHttpsProxies"))
	grpcMux.AddService(pb.NewRegionTargetHttpsProxiesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.RegionTargetHttpsProxies"))
	grpcMux.AddService(pb.NewSslPoliciesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.SslPolicies"))
	grpcMux.AddService(pb.NewTargetSslProxiesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.TargetSslProxies"))
	grpcMux.AddService(pb.NewTargetTcpProxiesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.TargetTcpProxies"))
	grpcMux.AddService(pb.NewRegionTargetTcpProxiesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.RegionTargetTcpProxies"))
	grpcMux.AddService(pb.NewUrlMapsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.UrlMaps"))
	grpcMux.AddService(pb.NewRegionUrlMapsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.RegionUrlMaps"))
	grpcMux.AddService(pb.NewNodeGroupsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.NodeGroups"))
	grpcMux.AddService(pb.NewNodeTemplatesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.NodeTemplates"))
	grpcMux.AddService(pb.NewDisksClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.Disks"))
	grpcMux.AddService(pb.NewRegionDisksClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.RegionDisks"))
	grpcMux.AddService(pb.NewFirewallPoliciesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.FirewallPolicies"))
	grpcMux.AddService(pb.NewForwardingRulesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.ForwardingRules"))
	grpcMux.AddService(pb.NewGlobalForwardingRulesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.GlobalForwardingRules"))
	grpcMux.AddService(pb.NewRegionOperationsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.RegionOperations"))
	grpcMux.AddService(pb.NewZoneOperationsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.ZoneOperations"))
	grpcMux.AddService(pb.NewGlobalOperationsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.GlobalOperations"))
	grpcMux.AddService(pb.NewGlobalOrganizationOperationsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.GlobalOrganizationOperations"))
	grpcMux.AddService(pb.NewAddressesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.Addresses"))
	grpcMux.AddService(pb.NewGlobalAddressesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.GlobalAddresses"))
	grpcMux.AddService(pb.NewRegionHealthChecksClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.RegionHealthChecks"))
	grpcMux.AddService(pb.NewHealthChecksClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.HealthChecks"))
	grpcMux.AddService(pb.NewSslCertificatesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.SslCertificates"))
	grpcMux.AddService(pb.NewRegionSslCertificatesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.RegionSslCertificates"))
	grpcMux.AddService(pb.NewRegionNetworkEndpointGroupsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.RegionNetworkEndpointGroups"))
	grpcMux.AddService(pb.NewServiceAttachmentsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.ServiceAttachments"))
	grpcMux.AddService(pb.NewImagesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.Images"))
	grpcMux.AddService(pb.NewInstancesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.Instances"))
	grpcMux.AddService(pb.NewInstanceTemplatesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.InstanceTemplates"))
	grpcMux.AddService(pb.NewZonesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.Zones"))
	grpcMux.AddService(pb.NewInstanceGroupManagersClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.InstanceGroupManagers"))
	grpcMux.AddService(pb.NewReservationsClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.Reservations"))
	grpcMux.AddService(pb.NewBackendServicesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.BackendServices"))
	grpcMux.AddService(pb.NewRegionBackendServicesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.RegionBackendServices"))
	grpcMux.AddService(pb.NewRoutesClient(conn), httptogrpc.WithServiceName("google.cloud.compute.v1.Routes"))

	grpcMux.OverrideHeaders(func(w http.ResponseWriter) {
		w.Header().Del("Cache-Control")
	})

	// Terraform uses the /beta/ endpoints, but we have protos only for v1.
	// Also, we probably want to be implementing the newer versions
	// as that makes it easier to move KCC to newer API versions.
	// So far, it seems that all of beta is a direct mapping to v1 - though
	// I'm sure eventually we'll find something that needs special handling.
	rewriteBetaToV1 := func(w http.ResponseWriter, r *http.Request) {
		u := r.URL
		if strings.HasPrefix(u.Path, "/compute/beta/") {
			u.Path = "/compute/v1/" + strings.TrimPrefix(u.Path, "/compute/beta/")
		}

		grpcMux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(rewriteBetaToV1), nil
}
