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

// +tool:mockgcp-service
// http.host: vectorsearch.googleapis.com
// proto.service: google.cloud.vectorsearch.v1.VectorSearchService

package mockvectorsearch

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/vectorsearch/apiv1/vectorsearchpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pbhttp "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/vectorsearch/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked vectorsearch service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	vectorSearchServer *VectorSearchServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.vectorSearchServer = &VectorSearchServer{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"vectorsearch.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterVectorSearchServiceServer(grpcServer, s.vectorSearchServer)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pbhttp.RegisterVectorSearchServiceHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"))
	if err != nil {
		return nil, err
	}

	mux.RewriteHeaders = func(ctx context.Context, response http.ResponseWriter, payload proto.Message) {
		response.Header().Del("Cache-Control")
	}

	return mux, nil
}
