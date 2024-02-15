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

package mockresourcemanager

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb_v1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v1"
	pb_v3 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked privateca service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projectsInternal *ProjectsInternal
	operations       *operations.Operations

	projectsV1 *ProjectsV1
	projectsV3 *ProjectsV3
}

// New creates a MockService.
func New(kubeClient client.Client, storage storage.Storage) *MockService {
	s := &MockService{
		kube:       kubeClient,
		storage:    storage,
		operations: operations.NewOperationsService(storage),
	}
	s.projectsInternal = &ProjectsInternal{MockService: s}
	s.projectsV1 = &ProjectsV1{MockService: s}
	s.projectsV3 = &ProjectsV3{MockService: s}
	return s
}

func (s *MockService) GetInternalService() projects.ProjectStore {
	return s.projectsInternal
}

func (s *MockService) ExpectedHost() string {
	return "cloudresourcemanager.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb_v1.RegisterProjectsServer(grpcServer, s.projectsV1)
	pb_v3.RegisterProjectsServer(grpcServer, s.projectsV3)
	pb_v3.RegisterTagKeysServer(grpcServer, &TagKeys{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn,
		pb_v1.RegisterProjectsHandler,
		pb_v3.RegisterProjectsHandler,
		pb_v3.RegisterTagKeysHandler,
		s.operations.RegisterOperationsHandler("v3"))
	if err != nil {
		return nil, err
	}

	return mux, nil
}
