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

package mockapigee

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked privateca service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	ops := operations.NewOperationsService(storage)

	// gcloud is very particular about the operation format
	ops.SetOperationFormat("{uuid}")

	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      ops,
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"apigee.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterOrganizationsEndpointAttachmentsServerServer(grpcServer, &endpointAttachmentsServer{MockService: s})
	pb.RegisterOrganizationsEnvironmentsServerServer(grpcServer, &environmentsServer{MockService: s})
	pb.RegisterOrganizationsEnvgroupsServerServer(grpcServer, &EnvgroupV1{MockService: s})
	pb.RegisterOrganizationsEnvgroupsAttachmentsServerServer(grpcServer, &envgroupsAttachmentsServer{MockService: s})
	pb.RegisterOrganizationsInstancesServerServer(grpcServer, &instancesServer{MockService: s})
	pb.RegisterOrganizationsInstancesAttachmentsServerServer(grpcServer, &instancesAttachmentsServer{MockService: s})
	pb.RegisterOrganizationsServerServer(grpcServer, &organizationsServer{MockService: s})
	pb.RegisterProjectsServerServer(grpcServer, &projectsServer{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterOrganizationsEndpointAttachmentsServerHandler,
		pb.RegisterOrganizationsEnvironmentsServerHandler,
		pb.RegisterOrganizationsEnvgroupsServerHandler,
		pb.RegisterOrganizationsEnvgroupsAttachmentsServerHandler,
		pb.RegisterOrganizationsInstancesServerHandler,
		pb.RegisterOrganizationsInstancesAttachmentsServerHandler,
		pb.RegisterOrganizationsServerHandler,
		pb.RegisterProjectsServerHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"))
	if err != nil {
		return nil, err
	}

	// Returns slightly non-standard errors
	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 {
			error.Errors = nil
		}
	}

	return mux, nil
}
