// Copyright 2026 Google LLC
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

package mockvmmigration

// +tool:mockgcp-service
// http.host: vmmigration.googleapis.com
// proto.service: google.cloud.vmmigration.v1.VmMigration

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked vmmigration service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1 *VmMigrationV1
}

type VmMigrationV1 struct {
	*MockService
	pb.UnimplementedVmMigrationServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &VmMigrationV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"vmmigration.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterVmMigrationServer(grpcServer, s.v1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("creating http mux: %w", err)
	}

	mux.AddService(pb.NewVmMigrationClient(conn))
	mux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	mux.OverrideHeaders(func(response http.ResponseWriter) {
		response.Header().Del("Cache-Control")
	})

	return mux, nil
}
