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
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pbv2 "cloud.google.com/go/iam/apiv2/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pbhttp_v2 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/iam/v2"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/admin/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked IAM service.
type MockService struct {
	*common.MockEnvironment
	storage    storage.Storage
	operations *operations.Operations
}

type IAMServer struct {
	*MockService
	pb.UnimplementedIAMServer
}

// New creates a MockService
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"iam.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterIAMServer(grpcServer, &IAMServer{MockService: s})
	pbv2.RegisterPoliciesServer(grpcServer, &IAMV2PoliciesServer{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterIAMHandler,
		pbhttp_v2.RegisterPoliciesHandler)
	if err != nil {
		return nil, err
	}

	mux.RewriteHeaders = func(ctx context.Context, response http.ResponseWriter, payload proto.Message) {
		response.Header().Del("Cache-Control")
	}

	return mux, nil
}
