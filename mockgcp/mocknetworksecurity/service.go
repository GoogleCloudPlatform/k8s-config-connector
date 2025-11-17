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

// +tool:mockgcp-service
// http.host: networksecurity.googleapis.com
// proto.service: google.cloud.networksecurity.v1beta1.NetworkSecurity

package mocknetworksecurity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

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
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(pb.NewNetworkSecurityClient(conn))

	grpcMux.AddOperationsPath("/v1beta1/{prefix=**}/operations/{name}", conn)

	// mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
	// 	if error.Code == 404 {
	// 		error.Errors = nil
	// 	}
	// }

	return grpcMux, nil
}

type NetworkSecurityServer struct {
	*MockService
	pb.UnimplementedNetworkSecurityServer
}
