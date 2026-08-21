// Copyright 2025 Google LLC
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

// +tool:mockgcp-service
// http.host: iap.googleapis.com
// proto.service: google.cloud.iap.v1.IdentityAwareProxyAdminService
// proto.service: google.cloud.iap.v1.IdentityAwareProxyOAuthService

package mockiap

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	pb "cloud.google.com/go/iap/apiv1/iappb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked iap service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations
}

type IdentityAwareProxyAdminService struct {
	*MockService
	pb.UnimplementedIdentityAwareProxyAdminServiceServer
}

type IdentityAwareProxyOAuthService struct {
	*MockService
	pb.UnimplementedIdentityAwareProxyOAuthServiceServer
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
	return []string{"iap.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterIdentityAwareProxyAdminServiceServer(grpcServer, &IdentityAwareProxyAdminService{MockService: s})
	pb.RegisterIdentityAwareProxyOAuthServiceServer(grpcServer, &IdentityAwareProxyOAuthService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(pb.NewIdentityAwareProxyAdminServiceClient(conn))
	grpcMux.AddService(pb.NewIdentityAwareProxyOAuthServiceClient(conn))

	// grpcMux.AddOperationsPath("/v1/{prefix=**}/operations/{name=**}", conn)

	return grpcMux, nil
}
