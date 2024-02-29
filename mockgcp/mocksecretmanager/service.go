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

package mocksecretmanager

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/secretmanager/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked secret manager service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	v1 *SecretsV1
}

// New creates a mockSecretManager
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
	}
	s.v1 = &SecretsV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"secretmanager.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterSecretManagerServiceServer(grpcServer, s.v1)
	// longrunning.RegisterOperationsServer(grpcServer, s)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterSecretManagerServiceHandler,
	)
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
