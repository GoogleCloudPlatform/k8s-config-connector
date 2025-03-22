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

package mockkms

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/kms/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked kms service.
type MockService struct {
	*common.MockEnvironment
	storage              storage.Storage
	operations           *operations.Operations
	v1AutokeyAdminServer *autokeyAdminServer
	v1AutokeyServer      *autokeyServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1AutokeyAdminServer = &autokeyAdminServer{MockService: s}
	s.v1AutokeyServer = &autokeyServer{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"cloudkms.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterKeyManagementServiceServer(grpcServer, &kmsServer{MockService: s})
	pb.RegisterAutokeyAdminServer(grpcServer, s.v1AutokeyAdminServer)
	pb.RegisterAutokeyServer(grpcServer, s.v1AutokeyServer)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterKeyManagementServiceHandler,
		pb.RegisterAutokeyAdminHandler,
		pb.RegisterAutokeyHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"),
	)
	if err != nil {
		return nil, err
	}

	// Returns slightly non-standard errors
	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 && (strings.Contains(error.Message, "KeyRing") || strings.Contains(error.Message, "CryptoKey") || strings.Contains(error.Message, "ImportJob")) {
			error.Errors = nil
		}
	}
	return mux, nil
}
