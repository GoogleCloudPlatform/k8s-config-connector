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

// +tool:mockgcp-service
// http.host: notebooks.googleapis.com
// proto.service: google.cloud.notebooks.v1.NotebookService

package mocknotebooks

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	pb_v1beta1 "cloud.google.com/go/notebooks/apiv1beta1/notebookspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked notebooks service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations
}

type NotebookServiceV1 struct {
	*MockService
	pb.UnimplementedNotebookServiceServer
}

type NotebookServiceV1beta1 struct {
	*MockService
	pb_v1beta1.UnimplementedNotebookServiceServer
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
	return []string{"notebooks.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterNotebookServiceServer(grpcServer, &NotebookServiceV1{MockService: s})
	pb_v1beta1.RegisterNotebookServiceServer(grpcServer, &NotebookServiceV1beta1{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error creating mux: %w", err)
	}

	mux.AddService(pb.NewNotebookServiceClient(conn))
	mux.AddService(pb_v1beta1.NewNotebookServiceClient(conn))
	mux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)
	mux.AddOperationsPath("/v1beta1/{prefix=**}/operations/{name}", conn)

	return mux, nil
}
