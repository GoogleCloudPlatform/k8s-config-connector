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

package mockcloudidentity

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/apps/cloudidentity/groups/v1beta1"
)

// MockService represents a mocked cloudidentity service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"cloudidentity.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterGroupsServerServer(grpcServer, &groupsServer{MockService: s})
	pb.RegisterGroupsMembershipsServerServer(grpcServer, &groupsMembershipsServer{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterGroupsServerHandler,
		pb.RegisterGroupsMembershipsServerHandler,
		s.operations.RegisterOperationsPath("/v1beta1/operations/{name}"),
	)
	if err != nil {
		return nil, err
	}

	// Returns slightly non-standard errors
	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 && strings.Contains(error.Message, "Membership") {
			error.Errors = nil
		}
	}

	// DCL sends a trailing slash, but that is not technically correct
	// and it trips up grpc-gateway (https://github.com/grpc-ecosystem/grpc-gateway/issues/472)
	// e.g. POST https://cloudidentity.googleapis.com/v1beta1/groups/1946c1f7c26/memberships/?alt=json
	removeTrailingSlash := func(w http.ResponseWriter, r *http.Request) {
		u := r.URL
		if strings.HasSuffix(u.Path, "/memberships/") {
			u2 := *u
			u2.Path = strings.TrimSuffix(u2.Path, "/")
			r = httpmux.RewriteRequest(r, &u2)
		}

		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(removeTrailingSlash), nil
}
