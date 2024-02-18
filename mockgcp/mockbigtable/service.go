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

package mockbigtable

import (
	"context"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
	// pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/bigtable/admin/v2"
)

// MockService represents a mocked bigtable service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projects   projects.ProjectStore
	operations *operations.Operations
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		kube:       env.GetKubeClient(),
		storage:    storage,
		projects:   env.GetProjects(),
		operations: operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHost() string {
	return "bigtable.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterBigtableInstanceAdminServer(grpcServer, &instanceAdminServer{MockService: s})
	s.operations.RegisterGRPCServices(grpcServer)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux := runtime.NewServeMux()

	// We don't support http, because we are supporting the upstream grpc service id (i.e. google... rather than mockgcp...)
	// if err := pb.RegisterApiKeysHandler(ctx, mux, conn); err != nil {
	// 	return nil, err
	// }

	return mux, nil
}
