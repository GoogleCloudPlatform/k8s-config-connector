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

package mockworkflowexecution

// +tool:mockgcp-service
// http.host: workflowexecutions.googleapis.com
// proto.service: google.cloud.workflows.executions.v1.Executions

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	// Note we use "real" protos (not mockgcp) ones as it's GRPC API.
	executionspb "cloud.google.com/go/workflows/executions/apiv1/executionspb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked workflows service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations
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
	return []string{"workflowexecutions.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	executionspb.RegisterExecutionsServer(grpcServer, &workflowExecutionService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("creating http mux: %w", err)
	}

	mux.AddService(executionspb.NewExecutionsClient(conn))
	mux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	mux.OverrideHeaders(func(response http.ResponseWriter) {
		response.Header().Del("Cache-Control")
	})

	return mux, nil
}
