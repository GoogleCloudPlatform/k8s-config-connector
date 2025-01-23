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
// http.host: gkehub.googleapis.com
// proto.service: google.cloud.gkehub.v1beta.GkeHub
// proto.service: google.cloud.gkehub.v1beta1.GkeHubMembershipService

package mockgkehub

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	v1betapb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta"
	v1beta1pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked gkehubfeature service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1beta  *GKEHubFeature
	v1beta1 *GKEHubMembership
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1beta = &GKEHubFeature{MockService: s}
	s.v1beta1 = &GKEHubMembership{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"gkehub.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	v1betapb.RegisterGkeHubServer(grpcServer, s.v1beta)
	v1beta1pb.RegisterGkeHubMembershipServiceServer(grpcServer, s.v1beta1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{}, v1betapb.RegisterGkeHubHandler, v1beta1pb.RegisterGkeHubMembershipServiceHandler, s.operations.RegisterOperationsPath("/v1beta/{prefix=**}/operations/{name}"), s.operations.RegisterOperationsPath("/v1beta1/{prefix=**}/operations/{name}"))
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
