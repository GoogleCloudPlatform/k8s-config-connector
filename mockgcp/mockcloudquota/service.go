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

package mockcloudquota

// +tool:mockgcp-service
// http.host: cloudquotas.googleapis.com
// proto.service: google.api.cloudquotas.v1.CloudQuotas

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/cloudquotas/v1beta"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked cloudquotas service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1 *CloudQuotasV1
}

type CloudQuotasV1 struct {
	*MockService
	pb.UnimplementedCloudQuotasServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &CloudQuotasV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"cloudquotas.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterCloudQuotasServer(grpcServer, s.v1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterCloudQuotasHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"))

	if err != nil {
		return nil, err
	}

	return mux, nil
}

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
// http.host: cloudquotas.googleapis.com
// proto.service: google.api.cloudquotas.v1beta.QuotaAdjusterSettingsManager

package mockcloudquota

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/cloudquotas/v1beta"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations
}

type QuotaAdjusterSettingsManagerV1Beta struct {
	*MockService
	pb.UnimplementedQuotaAdjusterSettingsManagerServer
}

func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"cloudquotas.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterQuotaAdjusterSettingsManagerServer(grpcServer, &QuotaAdjusterSettingsManagerV1Beta{s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterQuotaAdjusterSettingsManagerHandler)
	if err != nil {
		return nil, err
	}
	return mux, nil
}
