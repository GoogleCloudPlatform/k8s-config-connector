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
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb_v1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v1"
	pb_v3 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked privateca service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	projectsInternal *ProjectsInternal
	projectsV1       *ProjectsV1
	projectsV3       *ProjectsV3

	tagKeys     *TagKeys
	tagValues   *TagValues
	tagBindings *TagBindingsServer
}

type TagKeys struct {
	*MockService
	pb_v3.UnimplementedTagKeysServer
}

type TagValues struct {
	*MockService
	pb_v3.UnimplementedTagValuesServer
}

type TagBindingsServer struct {
	*MockService
	pb_v3.UnimplementedTagBindingsServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.projectsInternal = &ProjectsInternal{MockService: s}
	s.projectsV1 = &ProjectsV1{MockService: s}
	s.projectsV3 = &ProjectsV3{MockService: s}
	s.tagKeys = &TagKeys{MockService: s}
	s.tagValues = &TagValues{MockService: s}
	s.tagBindings = &TagBindingsServer{MockService: s}
	return s
}

func (s *MockService) GetProjectStore() projects.ProjectStore {
	return s.projectsInternal
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"cloudresourcemanager.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb_v1.RegisterProjectsServer(grpcServer, s.projectsV1)
	pb_v3.RegisterProjectsServer(grpcServer, s.projectsV3)
	pb_v3.RegisterFoldersServer(grpcServer, &Folders{MockService: s})
	pb_v3.RegisterTagKeysServer(grpcServer, s.tagKeys)
	pb_v3.RegisterTagValuesServer(grpcServer, s.tagValues)
	pb_v3.RegisterTagBindingsServer(grpcServer, s.tagBindings)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb_v1.RegisterProjectsHandler,
		pb_v3.RegisterProjectsHandler,
		pb_v3.RegisterFoldersHandler,
		pb_v3.RegisterTagKeysHandler,
		pb_v3.RegisterTagValuesHandler,
		pb_v3.RegisterTagBindingsHandler,
		s.operations.RegisterOperationsPath("/v1/operations/{name}"),
		s.operations.RegisterOperationsPath("/v3/operations/{name}"))
	if err != nil {
		return nil, err
	}

	mux.RewriteHeaders = func(ctx context.Context, response http.ResponseWriter, payload proto.Message) {
		response.Header().Del("Cache-Control")
	}

	return mux, nil
}
