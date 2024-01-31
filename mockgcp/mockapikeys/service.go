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

package mockapikeys

import (
	"context"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/apikeys/v2"
)

// MockService represents a mocked apikeys service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projects   projects.ProjectStore
	operations *operations.Operations

	v2 *APIKeysV2
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		kube:       env.GetKubeClient(),
		storage:    storage,
		projects:   env.GetProjects(),
		operations: operations.NewOperationsService(storage),
	}
	s.v2 = &APIKeysV2{MockService: s}
	return s
}

func (s *MockService) ExpectedHost() string {
	return "apikeys.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterApiKeysServer(grpcServer, s.v2)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn,
		pb.RegisterApiKeysHandler,
		s.operations.RegisterOperationsHandler("v2"),
	)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
