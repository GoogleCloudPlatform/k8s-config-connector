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

// +tool:mockgcp-service
// http.host: iam.googleapis.com
// proto.service: google.iam.admin.v1.IAM

package mockiam

import (
	"context"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	v1_pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/admin/v1"
	v1beta_pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/v1beta"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked IAM service.
type MockService struct {
	*common.MockEnvironment
	storage      storage.Storage
	operations   *operations.Operations
	serverV1     *ServerV1
	serverV1Beta *ServerV1Beta
}

type ServerV1 struct {
	*MockService
	v1_pb.UnimplementedIAMServer
}

type ServerV1Beta struct {
	*MockService
	v1beta_pb.UnimplementedWorkloadIdentityPoolsServer
}

// New creates a MockService
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.serverV1 = &ServerV1{MockService: s}
	s.serverV1Beta = &ServerV1Beta{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"iam.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	v1_pb.RegisterIAMServer(grpcServer, s.serverV1)
	v1beta_pb.RegisterWorkloadIdentityPoolsServer(grpcServer, s.serverV1Beta)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		v1_pb.RegisterIAMHandler,
		v1beta_pb.RegisterWorkloadIdentityPoolsHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"))

	if err != nil {
		return nil, err
	}

	return mux, nil
}
