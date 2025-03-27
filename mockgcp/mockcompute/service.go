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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
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
	pb.RegisterGlobalOperationsServer(grpcServer, &GlobalOperationsV1{MockService: s})
	pb.RegisterGlobalOrganizationOperationsServer(grpcServer, &GlobalOrganizationOperationsV1{MockService: s})

	pb.RegisterNodeGroupsServer(grpcServer, &NodeGroupsV1{MockService: s})
	pb.RegisterNodeTemplatesServer(grpcServer, &NodeTemplatesV1{MockService: s})

	pb.RegisterAddressesServer(grpcServer, &RegionalAddressesV1{MockService: s})
	pb.RegisterGlobalAddressesServer(grpcServer, &GlobalAddressesV1{MockService: s})
	pb.RegisterSslCertificatesServer(grpcServer, &GlobalSSLCertificatesV1{MockService: s})
	pb.RegisterRegionSslCertificatesServer(grpcServer, &RegionalSSLCertificatesV1{MockService: s})
	pb.RegisterTargetSslProxiesServer(grpcServer, &TargetSslProxyV1{MockService: s})
	pb.RegisterTargetTcpProxiesServer(grpcServer, &GlobalTargetTcpProxyV1{MockService: s})
	pb.RegisterRegionTargetTcpProxiesServer(grpcServer, &RegionalTargetTcpProxyV1{MockService: s})

	pb.RegisterServiceAttachmentsServer(grpcServer, &RegionalServiceAttachmentV1{MockService: s})

	pb.RegisterFirewallPoliciesServer(grpcServer, &FirewallPoliciesV1{MockService: s})

	pb.RegisterGlobalForwardingRulesServer(grpcServer, &GlobalForwardingRulesV1{MockService: s})
	pb.RegisterForwardingRulesServer(grpcServer, &RegionalForwardingRulesV1{MockService: s})

	pb.RegisterImagesServer(grpcServer, &ImagesV1{MockService: s})

	pb.RegisterInstancesServer(grpcServer, &InstancesV1{MockService: s})

	pb.RegisterZonesServer(grpcServer, &ZonesV1{MockService: s})

	pb.RegisterNetworkAttachmentsServer(grpcServer, &networkAttachmentsV1{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{})
	if err != nil {
		return nil, err
	}

	if err := pb.RegisterBackendBucketsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterBackendServicesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionBackendServicesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterExternalVpnGatewaysHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterNetworkEdgeSecurityServicesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterNetworksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterNetworkAttachmentsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterSubnetworksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterVpnGatewaysHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterTargetVpnGatewaysHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterTargetGrpcProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterTargetHttpProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionTargetHttpProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterTargetHttpsProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionTargetHttpsProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterTargetSslProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterTargetTcpProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionTargetTcpProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterUrlMapsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionUrlMapsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterNodeGroupsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterNodeTemplatesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterDisksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionDisksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterFirewallPoliciesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterForwardingRulesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterGlobalForwardingRulesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterRegionOperationsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterGlobalOperationsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterGlobalOrganizationOperationsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterAddressesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterGlobalAddressesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterRegionHealthChecksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterHealthChecksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	// for ssl certs and the managedsslcerts
	if err := pb.RegisterSslCertificatesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionSslCertificatesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterServiceAttachmentsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterImagesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterInstancesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterZonesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterNetworkAttachmentsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	// Returns slightly non-standard errors
	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		// Does not return status (at least for 404)
		error.Status = ""
	}

	// Does not return Cache-Control header
	mux.RewriteHeaders = func(ctx context.Context, response http.ResponseWriter, payload proto.Message) {
		response.Header().Del("Cache-Control")
	}

	// Terraform uses the /beta/ endpoints, but we have protos only for v1.
	// Also, we probably want to be implementing the newer versions
	// as that makes it easier to move KCC to newer API versions.
	// So far, it seems that all of beta is a direct mapping to v1 - though
	// I'm sure eventually we'll find something that needs special handling.
	rewriteBetaToV1 := func(w http.ResponseWriter, r *http.Request) {
		u := r.URL
		if strings.HasPrefix(u.Path, "/compute/beta/") {
			u2 := *u
			u2.Path = "/compute/v1/" + strings.TrimPrefix(u.Path, "/compute/beta/")
			r = httpmux.RewriteRequest(r, &u2)
		}

		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(rewriteBetaToV1), nil
}
