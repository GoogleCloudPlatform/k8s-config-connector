// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package mocknetworksecurity

import (
	"context"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	pb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked networksecurity service.
type MockService struct {
	*common.MockEnvironment
	storage    storage.Storage
	operations *operations.Operations

	v1 *NetworkSecurityServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"networksecurity.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterNetworkSecurityServer(grpcServer, &NetworkSecurityServer{MockService: s})
	pbv1.RegisterNetworkSecurityServer(grpcServer, &NetworkSecurityV1Server{MockService: s})
	pbv1.RegisterMirroringServer(grpcServer, &MirroringServer{MockService: s})
	pbv1.RegisterInterceptServer(grpcServer, &InterceptServer{MockService: s})
	pbv1.RegisterSSERealmServiceServer(grpcServer, &SSERealmServer{MockService: s})
	pbv1.RegisterFirewallActivationServer(grpcServer, &FirewallActivationServer{MockService: s})
	pbv1.RegisterDnsThreatDetectorServiceServer(grpcServer, &DnsThreatDetectorServer{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	mux.AddService(pb.NewNetworkSecurityClient(conn))
	mux.AddService(pbv1.NewNetworkSecurityClient(conn))
	mux.AddService(pbv1.NewMirroringClient(conn))
	mux.AddService(pbv1.NewInterceptClient(conn))
	mux.AddService(pbv1.NewSSERealmServiceClient(conn))
	mux.AddService(pbv1.NewFirewallActivationClient(conn))
	mux.AddService(pbv1.NewDnsThreatDetectorServiceClient(conn))
	mux.AddOperationsPath("/v1beta1/{prefix=**}/operations/{name}", conn)
	mux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	return mux, nil
}
