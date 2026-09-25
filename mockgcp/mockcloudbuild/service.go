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
package mockcloudbuild

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	pbv2 "cloud.google.com/go/cloudbuild/apiv2/cloudbuildpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pbhttpv2 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/devtools/cloudbuild/v2"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/devtools/cloudbuild/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked cloudbuild service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1 *CloudBuildV1
	v2 *CloudBuildConnectionServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &CloudBuildV1{MockService: s}
	s.v2 = &CloudBuildConnectionServer{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"cloudbuild.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterCloudBuildServer(grpcServer, s.v1)
	pbv2.RegisterRepositoryManagerServer(grpcServer, s.v2)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterCloudBuildHandler,
		pbhttpv2.RegisterRepositoryManagerHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"),
		s.operations.RegisterOperationsPath("/v2/{prefix=**}/operations/{name}"))
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
