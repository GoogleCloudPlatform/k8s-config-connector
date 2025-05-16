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

package mockserviceusage

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb_v1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/serviceusage/v1"
	pb_v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/serviceusage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked serviceusage service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	serviceusagev1      *ServiceUsageV1
	serviceusagev1beta1 *ServiceUsageV1Beta1
}

// New creates a MockService
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.serviceusagev1 = &ServiceUsageV1{MockService: s}
	s.serviceusagev1beta1 = &ServiceUsageV1Beta1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"serviceusage.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb_v1.RegisterServiceUsageServer(grpcServer, s.serviceusagev1)
	pb_v1beta1.RegisterServiceUsageServer(grpcServer, s.serviceusagev1beta1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb_v1.RegisterServiceUsageHandler,
		pb_v1beta1.RegisterServiceUsageHandler,
		s.operations.RegisterOperationsPath("/v1beta1/operations/{name}"),
		s.operations.RegisterOperationsPath("/v1/operations/{name}"),
	)
	if err != nil {
		return nil, fmt.Errorf("creating http mux: %w", err)
	}

	return mux, nil
}
