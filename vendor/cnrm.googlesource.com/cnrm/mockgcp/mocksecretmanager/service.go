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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	secretmanager_http "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/secretmanager/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	secretmanager "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ExpectedHost = "secretmanager.googleapis.com"

// MockService represents a mocked secret manager service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projects *projects.ProjectStore
}

// NewMockService creates a mockSecretManager
func NewMockService(kube client.Client, storage storage.Storage) *MockService {
	s := &MockService{
		kube:     kube,
		storage:  storage,
		projects: projects.NewProjectStore(),
	}
	return s
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	secretmanager.RegisterSecretManagerServiceServer(grpcServer, s)
	// longrunning.RegisterOperationsServer(grpcServer, s)
}

func (s *MockService) NewMux(ctx context.Context, conn *grpc.ClientConn) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()
	if err := secretmanager_http.RegisterSecretManagerServiceHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	return mux, nil
}
