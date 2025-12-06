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

package mockclouddms

// +tool:mockgcp-service
// http.host: datamigration.googleapis.com
// proto.service: google.cloud.clouddms.v1.DataMigrationService

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	grpcpb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/clouddms/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
)

// MockService represents a mocked datamigration service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations
}

type DataMigrationServiceV1 struct {
	*MockService
	pb.UnimplementedDataMigrationServiceServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"datamigration.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterDataMigrationServiceServer(grpcServer, &DataMigrationServiceV1{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		grpcpb.RegisterDataMigrationServiceHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"))

	if err != nil {
		return nil, err
	}

	return mux, nil
}
