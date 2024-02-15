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

package mockiam

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/admin/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked IAM service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projects projects.ProjectStore

	serverV1 *ServerV1
}

type ServerV1 struct {
	*MockService
	pb.UnimplementedIAMServer
}

// New creates a MockService
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		kube:     env.GetKubeClient(),
		storage:  storage,
		projects: env.GetProjects(),
	}
	s.serverV1 = &ServerV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHost() string {
	return "iam.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterIAMServer(grpcServer, s.serverV1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	return httpmux.NewServeMux(ctx, conn, pb.RegisterIAMHandler)
}
