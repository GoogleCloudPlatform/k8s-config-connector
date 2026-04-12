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
// http.host: dataplex.googleapis.com
// proto.service: google.cloud.dataplex.v1.CatalogService

package mockdataplex

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataplex/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)
// MockService represents a mocked dataplex service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	// Store the underlying GRPC servers
	dataplexService *DataplexService
	catalogService  *CatalogService
}

type DataplexService struct {
	*MockService
	pb.UnimplementedDataplexServiceServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.dataplexService = &DataplexService{MockService: s}
	s.catalogService = &CatalogService{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"dataplex.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterDataplexServiceServer(grpcServer, s.dataplexService)
	pb.RegisterCatalogServiceServer(grpcServer, s.catalogService)

	// Register under the original service names so that the KCC GRPC client works.
	dataplexDesc := pb.DataplexService_ServiceDesc
	dataplexDesc.ServiceName = "google.cloud.dataplex.v1.DataplexService"
	grpcServer.RegisterService(&dataplexDesc, s.dataplexService)

	catalogDesc := pb.CatalogService_ServiceDesc
	catalogDesc.ServiceName = "google.cloud.dataplex.v1.CatalogService"
	grpcServer.RegisterService(&catalogDesc, s.catalogService)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	grpcMux.AddService(pb.NewDataplexServiceClient(conn))
	grpcMux.AddService(pb.NewCatalogServiceClient(conn))
	grpcMux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	return grpcMux, nil
}
