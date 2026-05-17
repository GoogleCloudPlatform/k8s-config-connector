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

package mockrun

import (
	"context"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/run/apiv2/runpb"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked run service.
type MockService struct {
	*common.MockEnvironment
	storage     storage.Storage
	operations  *operations.Operations
	v2          *RunV2
	servicesV2  *ServicesV2
	workerPools *workerPools
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v2 = &RunV2{MockService: s}
	s.servicesV2 = &ServicesV2{MockService: s}
	s.workerPools = &workerPools{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"run.googleapis.com"}

}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterJobsServer(grpcServer, s.v2)
	pb.RegisterServicesServer(grpcServer, s.servicesV2)
	pb.RegisterWorkerPoolsServer(grpcServer, s.workerPools)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	mux.AddService(pb.NewJobsClient(conn))
	mux.AddService(pb.NewServicesClient(conn))
	mux.AddService(pb.NewWorkerPoolsClient(conn))
	mux.AddOperationsPath("/v2/{prefix=**}/operations/{name}", conn)

	return mux, nil
}
