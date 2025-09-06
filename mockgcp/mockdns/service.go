// Copyright 2025 Google LLC
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

package mockdns

import (
	"context"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/dns/v1"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked dns service.
type MockService struct {
	storage storage.Storage

	projects projects.ProjectStore

	operations *dnsOperations
}

// New creates a dnsService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		storage:    storage,
		projects:   env.Projects,
		operations: newDNSOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"dns.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterManagedZonesServerServer(grpcServer, &managedZonesService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterManagedZonesServerHandler,
	)

	if err != nil {
		return nil, err
	}

	// Terraform uses the /v1beta2/ endpoints, but we prefer to implement v1.
	// Rewrite the the request (they seem to be compatible).
	rewriteBetaToV1 := func(w http.ResponseWriter, r *http.Request) {
		u := r.URL
		if strings.HasPrefix(u.Path, "/dns/v1beta2/") {
			u2 := *u
			u2.Path = "/dns/v1/" + strings.TrimPrefix(u.Path, "/dns/v1beta2/")
			r = httpmux.RewriteRequest(r, &u2)
		}

		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(rewriteBetaToV1), nil
}
