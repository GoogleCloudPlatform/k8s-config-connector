// Copyright 2024 Google LLC
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

// +tool:mockgcp-service
// http.host: compute.googleapis.com
// proto.service: google.cloud.compute.v1.NetworkEdgeSecurityServices

import (
	"context"
	"net/http"

	compute "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked compute service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.ComputeOperations

	backendServices                 *backendServicesServer
	regionBackendServices           *regionBackendServicesServer
	firewalls                       *firewallsServer
	forwardingRules                 *forwardingRulesServer
	globalAddresses                 *globalAddressesServer
	globalForwardingRules           *globalForwardingRulesServer
	globalNetworkEndpointGroups     *globalNetworkEndpointGroupsServer
	globalOperations                *globalOperationsServer
	healthChecks                    *healthChecksServer
	regionHealthChecks              *regionHealthChecksServer
	regionHealthCheckServices       *regionHealthCheckServicesServer
	instances                       *instancesServer
	instanceTemplates               *instanceTemplatesServer
	instanceGroupManagers           *instanceGroupManagersServer
	regionInstanceGroupManagers     *regionInstanceGroupManagersServer
	networkEndpointGroups           *networkEndpointGroupsServer
	regionNetworkEndpointGroups     *regionNetworkEndpointGroupsServer
	networks                        *networksServer
	projects                        *projectsServer
	regions                         *regionsServer
	regionOperations                *regionOperationsServer
	regionAutoscalers               *regionAutoscalersServer
	routers                         *routersServer
	routes                          *routesServer
	sslCertificates                 *sslCertificatesServer
	sslPolicies                     *sslPoliciesServer
	subnetworks                     *subnetworksServer
	targetGrpcProxies               *targetGrpcProxiesServer
	targetHttpProxies               *targetHttpProxiesServer
	targetHttpsProxies              *targetHttpsProxiesServer
	targetInstances                 *targetInstancesServer
	targetPools                     *targetPoolsServer
	targetSslProxies                *targetSslProxiesServer
	targetTcpProxies                *targetTcpProxiesServer
	urlMaps                         *urlMapsServer
	regionUrlMaps                   *regionUrlMapsServer
	vpnTunnels                      *vpnTunnelsServer
	zoneOperations                  *zoneOperationsServer
	networkAttachments              *networkAttachmentsServer
	serviceAttachments              *serviceAttachmentsServer
	securityPolicies                *securityPoliciesServer
	regionSecurityPolicies          *regionSecurityPoliciesServer
	regionNotificationEndpoints     *regionNotificationEndpointsServer
	regionNetworkFirewallPolicies   *regionNetworkFirewallPoliciesServer
	regionDisk                      *regionDisksServer
	regionBackendServicesServer     *regionBackendServicesServer
	networkEdgeSecurityServices     *networkEdgeSecurityServicesServer
	regionNetworkEdgeSecurityPolicy *regionNetworkEdgeSecurityPolicyServer
}

type networkEdgeSecurityServicesServer struct {
	*MockService
	compute.UnimplementedNetworkEdgeSecurityServicesServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewComputeOperations(storage),
	}
	s.backendServices = &backendServicesServer{MockService: s}
	s.regionBackendServices = &regionBackendServicesServer{MockService: s}
	s.firewalls = &firewallsServer{MockService: s}
	s.forwardingRules = &forwardingRulesServer{MockService: s}
	s.globalAddresses = &globalAddressesServer{MockService: s}
	s.globalForwardingRules = &globalForwardingRulesServer{MockService: s}
	s.globalNetworkEndpointGroups = &globalNetworkEndpointGroupsServer{MockService: s}
	s.globalOperations = &globalOperationsServer{MockService: s}
	s.healthChecks = &healthChecksServer{MockService: s}
	s.regionHealthChecks = &regionHealthChecksServer{MockService: s}
	s.regionHealthCheckServices = &regionHealthCheckServicesServer{MockService: s}
	s.instances = &instancesServer{MockService: s}
	s.instanceTemplates = &instanceTemplatesServer{MockService: s}
	s.instanceGroupManagers = &instanceGroupManagersServer{MockService: s}
	s.regionInstanceGroupManagers = &regionInstanceGroupManagersServer{MockService: s}
	s.networkEndpointGroups = &networkEndpointGroupsServer{MockService: s}
	s.regionNetworkEndpointGroups = &regionNetworkEndpointGroupsServer{MockService: s}
	s.networks = &networksServer{MockService: s}
	s.projects = &projectsServer{MockService: s}
	s.regions = &regionsServer{MockService: s}
	s.regionOperations = &regionOperationsServer{MockService: s}
	s.regionAutoscalers = &regionAutoscalersServer{MockService: s}
	s.routers = &routersServer{MockService: s}
	s.routes = &routesServer{MockService: s}
	s.sslCertificates = &sslCertificatesServer{MockService: s}
	s.sslPolicies = &sslPoliciesServer{MockService: s}
	s.subnetworks = &subnetworksServer{MockService: s}
	s.targetGrpcProxies = &targetGrpcProxiesServer{MockService: s}
	s.targetHttpProxies = &targetHttpProxiesServer{MockService: s}
	s.targetHttpsProxies = &targetHttpsProxiesServer{MockService: s}
	s.targetInstances = &targetInstancesServer{MockService: s}
	s.targetPools = &targetPoolsServer{MockService: s}
	s.targetSslProxies = &targetSslProxiesServer{MockService: s}
	s.targetTcpProxies = &targetTcpProxiesServer{MockService: s}
	s.urlMaps = &urlMapsServer{MockService: s}
	s.regionUrlMaps = &regionUrlMapsServer{MockService: s}
	s.vpnTunnels = &vpnTunnelsServer{MockService: s}
	s.zoneOperations = &zoneOperationsServer{MockService: s}
	s.networkAttachments = &networkAttachmentsServer{MockService: s}
	s.serviceAttachments = &serviceAttachmentsServer{MockService: s}
	s.securityPolicies = &securityPoliciesServer{MockService: s}
	s.regionSecurityPolicies = &regionSecurityPoliciesServer{MockService: s}
	s.regionNotificationEndpoints = &regionNotificationEndpointsServer{MockService: s}
	s.regionNetworkFirewallPolicies = &regionNetworkFirewallPoliciesServer{MockService: s}
	s.regionDisk = &regionDisksServer{MockService: s}
	s.regionBackendServicesServer = &regionBackendServicesServer{MockService: s}
	s.networkEdgeSecurityServices = &networkEdgeSecurityServicesServer{MockService: s}
	s.regionNetworkEdgeSecurityPolicy = &regionNetworkEdgeSecurityPolicyServer{MockService: s}

	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"compute.googleapis.com", "www.googleapis.com"}
}

// Register registers the handlers for this service.
func (s *MockService) Register(grpcServer *grpc.Server) {
	compute.RegisterBackendServicesServer(grpcServer, s.backendServices)
	compute.RegisterRegionBackendServicesServer(grpcServer, s.regionBackendServices)
	compute.RegisterFirewallsServer(grpcServer, s.firewalls)
	compute.RegisterForwardingRulesServer(grpcServer, s.forwardingRules)
	compute.RegisterGlobalAddressesServer(grpcServer, s.globalAddresses)
	compute.RegisterGlobalForwardingRulesServer(grpcServer, s.globalForwardingRules)
	compute.RegisterGlobalNetworkEndpointGroupsServer(grpcServer, s.globalNetworkEndpointGroups)
	compute.RegisterGlobalOperationsServer(grpcServer, s.globalOperations)
	compute.RegisterHealthChecksServer(grpcServer, s.healthChecks)
	compute.RegisterRegionHealthChecksServer(grpcServer, s.regionHealthChecks)
	compute.RegisterRegionHealthCheckServicesServer(grpcServer, s.regionHealthCheckServices)
	compute.RegisterInstancesServer(grpcServer, s.instances)
	compute.RegisterInstanceTemplatesServer(grpcServer, s.instanceTemplates)
	compute.RegisterInstanceGroupManagersServer(grpcServer, s.instanceGroupManagers)
	compute.RegisterRegionInstanceGroupManagersServer(grpcServer, s.regionInstanceGroupManagers)
	compute.RegisterNetworkEndpointGroupsServer(grpcServer, s.networkEndpointGroups)
	compute.RegisterRegionNetworkEndpointGroupsServer(grpcServer, s.regionNetworkEndpointGroups)
	compute.RegisterNetworksServer(grpcServer, s.networks)
	compute.RegisterProjectsServer(grpcServer, s.projects)
	compute.RegisterRegionsServer(grpcServer, s.regions)
	compute.RegisterRegionOperationsServer(grpcServer, s.regionOperations)
	compute.RegisterRegionAutoscalersServer(grpcServer, s.regionAutoscalers)
	compute.RegisterRoutersServer(grpcServer, s.routers)
	compute.RegisterRoutesServer(grpcServer, s.routes)
	compute.RegisterSslCertificatesServer(grpcServer, s.sslCertificates)
	compute.RegisterSslPoliciesServer(grpcServer, s.sslPolicies)
	compute.RegisterSubnetworksServer(grpcServer, s.subnetworks)
	compute.RegisterTargetGrpcProxiesServer(grpcServer, s.targetGrpcProxies)
	compute.RegisterTargetHttpProxiesServer(grpcServer, s.targetHttpProxies)
	compute.RegisterTargetHttpsProxiesServer(grpcServer, s.targetHttpsProxies)
	compute.RegisterTargetInstancesServer(grpcServer, s.targetInstances)
	compute.RegisterTargetPoolsServer(grpcServer, s.targetPools)
	compute.RegisterTargetSslProxiesServer(grpcServer, s.targetSslProxies)
	compute.RegisterTargetTcpProxiesServer(grpcServer, s.targetTcpProxies)
	compute.RegisterUrlMapsServer(grpcServer, s.urlMaps)
	compute.RegisterRegionUrlMapsServer(grpcServer, s.regionUrlMaps)
	compute.RegisterVpnTunnelsServer(grpcServer, s.vpnTunnels)
	compute.RegisterZoneOperationsServer(grpcServer, s.zoneOperations)
	compute.RegisterNetworkAttachmentsServer(grpcServer, s.networkAttachments)
	compute.RegisterServiceAttachmentsServer(grpcServer, s.serviceAttachments)
	compute.RegisterSecurityPoliciesServer(grpcServer, s.securityPolicies)
	compute.RegisterRegionSecurityPoliciesServer(grpcServer, s.regionSecurityPolicies)
	compute.RegisterRegionNotificationEndpointsServer(grpcServer, s.regionNotificationEndpoints)
	compute.RegisterRegionNetworkFirewallPoliciesServer(grpcServer, s.regionNetworkFirewallPolicies)
	compute.RegisterRegionDisksServer(grpcServer, s.regionDisk)
	compute.RegisterRegionBackendServicesServer(grpcServer, s.regionBackendServicesServer)
	compute.RegisterNetworkEdgeSecurityServicesServer(grpcServer, s.networkEdgeSecurityServices)
	compute.RegisterRegionNetworkEdgeSecurityPoliciesServer(grpcServer, s.regionNetworkEdgeSecurityPolicy)

}

// NewHTTPMux creates an HTTP handler that routes requests to the mock backend.
func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterBackendServicesHandler,
		pb.RegisterRegionBackendServicesHandler,
		pb.RegisterFirewallsHandler,
		pb.RegisterForwardingRulesHandler,
		pb.RegisterGlobalAddressesHandler,
		pb.RegisterGlobalForwardingRulesHandler,
		pb.RegisterGlobalNetworkEndpointGroupsHandler,
		pb.RegisterGlobalOperationsHandler,
		pb.RegisterHealthChecksHandler,
		pb.RegisterRegionHealthChecksHandler,
		pb.RegisterRegionHealthCheckServicesHandler,
		pb.RegisterInstancesHandler,
		pb.RegisterInstanceTemplatesHandler,
		pb.RegisterInstanceGroupManagersHandler,
		pb.RegisterRegionInstanceGroupManagersHandler,
		pb.RegisterNetworkEndpointGroupsHandler,
		pb.RegisterRegionNetworkEndpointGroupsHandler,
		pb.RegisterNetworksHandler,
		pb.RegisterProjectsHandler,
		pb.RegisterRegionsHandler,
		pb.RegisterRegionOperationsHandler,
		pb.RegisterRegionAutoscalersHandler,
		pb.RegisterRoutersHandler,
		pb.RegisterRoutesHandler,
		pb.RegisterSslCertificatesHandler,
		pb.RegisterSslPoliciesHandler,
		pb.RegisterSubnetworksHandler,
		pb.RegisterTargetGrpcProxiesHandler,
		pb.RegisterTargetHttpProxiesHandler,
		pb.RegisterTargetHttpsProxiesHandler,
		pb.RegisterTargetInstancesHandler,
		pb.RegisterTargetPoolsHandler,
		pb.RegisterTargetSslProxiesHandler,
		pb.RegisterTargetTcpProxiesHandler,
		pb.RegisterUrlMapsHandler,
		pb.RegisterRegionUrlMapsHandler,
		pb.RegisterVpnTunnelsHandler,
		pb.RegisterZoneOperationsHandler,
		pb.RegisterNetworkAttachmentsHandler,
		pb.RegisterServiceAttachmentsHandler,
		pb.RegisterSecurityPoliciesHandler,
		pb.RegisterRegionSecurityPoliciesHandler,
		pb.RegisterRegionNotificationEndpointsHandler,
		pb.RegisterRegionNetworkFirewallPoliciesHandler,
		pb.RegisterRegionDisksHandler,
		pb.RegisterRegionBackendServicesHandler,
		pb.RegisterNetworkEdgeSecurityServicesHandler,
		pb.RegisterRegionNetworkEdgeSecurityPoliciesHandler,

		s.operations.RegisterOperationsPath("/compute/v1/projects/{project}/global/operations/{name}"),
		s.operations.RegisterOperationsPath("/compute/v1/projects/{project}/regions/{region}/operations/{name}"),
		s.operations.RegisterOperationsPath("/compute/v1/projects/{project}/zones/{zone}/operations/{name}"),
	)
	if err != nil {
		return nil, err
	}

	mux.RewriteResponse = func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
		// Force response to http.StatusOK
		// See https://cloud.google.com/compute/docs/api/how-tos/api-requests-responses#handling_api_responses
		// See https://cloud.google.com/compute/docs/reference/rest/v1/forwardingRules/insert
		// See https://cloud.google.com/compute/docs/reference/rest/v1/regionOperations/get
		// See https://cloud.google.com/compute/docs/reference/rest/v1/globalOperations/get

		op, ok := resp.(*compute.Operation)
		if ok {
			if op.Status != nil && *op.Status == compute.Operation_DONE {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				return nil
			}
		}

		// Let the default logic handle things otherwise
		return httpmux.DefaultRewriteResponse(ctx, w, resp)
	}

	// TODO: We probably need this for mock compliance testing
	// mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
	// 	log.Printf("rewriting error %v", error)
	// 	if error.Code == 400 {
	// 		error.Errors = nil
	// 		error.Message = "Bad Request"
	// 		return
	// 	}
	// 	if error.Code == 404 {
	// 		// TODO: Need to check this is the expected format
	// 		error.Errors = nil
	// 		error.Message = fmt.Sprintf("The resource '%s' was not found", error.RequestURI)
	// 		return
	// 	}

	// }

	return mux, nil
}
