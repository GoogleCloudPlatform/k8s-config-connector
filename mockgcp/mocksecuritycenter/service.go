// Copyright 2026 Google LLC
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

package mocksecuritycenter

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked securitycenter service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"securitycenter.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterSecurityCenterServer(grpcServer, &securityCenterServer{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(pb.NewSecurityCenterClient(conn))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Rewrite /v1/organizations/12345/locations/global/bigQueryExports to /v1/organizations/12345/bigQueryExports
		// The proto doesn't have the locations/global bindings yet for bigQueryExports.
		path := r.URL.Path
		if strings.Contains(path, "/locations/global/bigQueryExports") {
			r.URL.Path = strings.Replace(path, "/locations/global/bigQueryExports", "/bigQueryExports", 1)
		}
		grpcMux.ServeHTTP(w, r)
	}), nil
}
