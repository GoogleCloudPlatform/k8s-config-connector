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

package mockbeyondcorp

import (
	"context"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/beyondcorp/clientconnectorservices/v1"
)

// MockService represents a mocked beyondcorp service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1 *ClientConnectorServicesV1
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &ClientConnectorServicesV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"beyondcorp.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterClientConnectorServicesServiceServer(grpcServer, s.v1)

	// Register with the google.cloud... ServiceName as well, so that the direct controller (which uses the standard GCP client) can call us!
	desc := pb.ClientConnectorServicesService_ServiceDesc
	desc.ServiceName = "google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorServicesService"
	grpcServer.RegisterService(&desc, s.v1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterClientConnectorServicesServiceHandler,
		s.operations.RegisterOperationsPath("/v1/projects/{project}/locations/{location}/operations/{name}"),
	)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
