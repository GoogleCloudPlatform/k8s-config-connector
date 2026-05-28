// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockspanner

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	databasepb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	instancepb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked privateca service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	databaseV1 *SpannerDatabaseV1
	instanceV1 *SpannerInstanceV1
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.databaseV1 = &SpannerDatabaseV1{MockService: s}
	s.instanceV1 = &SpannerInstanceV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"spanner.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	databasepb.RegisterDatabaseAdminServer(grpcServer, s.databaseV1)
	instancepb.RegisterInstanceAdminServer(grpcServer, s.instanceV1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	mux.AddService(databasepb.NewDatabaseAdminClient(conn))
	mux.AddService(instancepb.NewInstanceAdminClient(conn))
	mux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	mux.OverrideHeaders(func(response http.ResponseWriter) {
		response.Header().Del("X-Content-Type-Options")
	})

	return mux, nil
}
