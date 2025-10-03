// Copyright 2023 Google LLC
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

package mockcertificatemanager

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/certificatemanager/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked certificatemanager service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1 *CertificateManagerV1
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &CertificateManagerV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"certificatemanager.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterCertificateManagerServer(grpcServer, s.v1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterCertificateManagerHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"),
	)
	if err != nil {
		return nil, err
	}
	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 {
			switch resource := strings.Split(error.Message, " ")[0]; resource {
			case "dnsAuthorization", "certificate", "certificateMap", "certificateMapEntry":
				error.Message = strings.Replace(error.Message, resource, "Resource", 1)
			}
			error.Message = strings.Replace(error.Message, `"`, `'`, 2)
			error.Message = strings.Replace(error.Message, "not found", "was not found", 1)
			error.Errors = nil
		}
	}
	return mux, nil
}
