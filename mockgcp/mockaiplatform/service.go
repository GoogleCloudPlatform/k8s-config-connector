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

package mockaiplatform

import (
	"context"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1beta1"
)

// MockService represents a mocked aiplatform service.
type MockService struct {
	*common.MockEnvironment

	storage storage.Storage

	operations *operations.Operations
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
	return []string{"{region}-aiplatform.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterTensorboardServiceServer(grpcServer, &tensorboardService{MockService: s})
	pb.RegisterDatasetServiceServer(grpcServer, &datasetService{MockService: s})
	pb.RegisterEndpointServiceServer(grpcServer, &endpointService{MockService: s})
	pb.RegisterMetadataServiceServer(grpcServer, &metadataStoreService{MockService: s})
	pb.RegisterFeaturestoreServiceServer(grpcServer, &featurestoreService{MockService: s})
	pb.RegisterModelServiceServer(grpcServer, &modelService{MockService: s})
	pb.RegisterNotebookServiceServer(grpcServer, &notebookService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterTensorboardServiceHandler,
		pb.RegisterDatasetServiceHandler,
		pb.RegisterEndpointServiceHandler,
		pb.RegisterMetadataServiceHandler,
		pb.RegisterFeaturestoreServiceHandler,
		pb.RegisterModelServiceHandler,
		pb.RegisterNotebookServiceHandler,
		s.operations.RegisterOperationsPath("/v1beta1/{prefix=**}/operations/{name}"),
		s.operations.RegisterOperationsPath("/ui/{prefix=**}/operations/{name}"))
	if err != nil {
		return nil, err
	}

	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 {
			error.Errors = nil
		}
	}

	return mux, nil
}
