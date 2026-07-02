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

package mockaistreams

// +tool:mockgcp-service
// http.host: aistreams.googleapis.com
// proto.service: google.partner.aistreams.v1alpha1.AIStreams

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/partner/aistreams/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked aistreams service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1 *AIStreamsV1
}

type AIStreamsV1 struct {
	*MockService
	pb.UnimplementedAIStreamsServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &AIStreamsV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"aistreams.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterAIStreamsServer(grpcServer, s.v1)

	// Register under canonical name as well for raw gRPC controllers
	canonicalDesc := pb.AIStreams_ServiceDesc
	canonicalDesc.ServiceName = "google.partner.aistreams.v1alpha1.AIStreams"
	grpcServer.RegisterService(&canonicalDesc, s.v1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterAIStreamsHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"))
	if err != nil {
		return nil, err
	}
	return mux, nil
}
