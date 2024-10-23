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

package mockvpcaccess

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/vpcaccess/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked vpc service.
type MockService struct {
	*common.MockEnvironment
	storage    storage.Storage
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
	return []string{"vpcaccess.googleapis.com"}
}

type vpcAccessService struct {
	*MockService
	pb.UnimplementedVpcAccessServiceServer
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterVpcAccessServiceServer(grpcServer, &vpcAccessService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterVpcAccessServiceHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"),
	)
	if err != nil {
		return nil, err
	}

	// Returns slightly non-standard errors
	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 {
			error.Errors = nil
		}
	}

	// Terraform uses the /v1beta1/ endpoints, but we have protos only for v1.
	// Also, we probably want to be implementing the newer versions
	// as that makes it easier to move KCC to newer API versions.
	// So far, it seems we can get away with a simple mapping.
	rewriteBetaToV1 := func(w http.ResponseWriter, r *http.Request) {
		u := r.URL
		if strings.HasPrefix(u.Path, "/v1beta1/") {
			u.Path = "/v1/" + strings.TrimPrefix(u.Path, "/v1beta1/")
			r.URL = u
		}

		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(rewriteBetaToV1), nil
}
