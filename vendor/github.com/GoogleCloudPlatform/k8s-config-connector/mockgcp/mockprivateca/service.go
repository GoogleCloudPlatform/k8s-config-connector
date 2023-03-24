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

package mockprivateca

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb_http "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/security/privateca/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ExpectedHost = "privateca.googleapis.com"

// MockService represents a mocked privateca service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projects   *projects.ProjectStore
	operations *operations.Operations

	v1 *PrivateCAV1
}

// New creates a MockService.
func New(kube client.Client, storage storage.Storage) *MockService {
	s := &MockService{
		kube:       kube,
		storage:    storage,
		projects:   projects.NewProjectStore(),
		operations: operations.NewOperationsService(storage),
	}
	s.v1 = &PrivateCAV1{MockService: s}
	return s
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterCertificateAuthorityServiceServer(grpcServer, s.v1)
}

func (s *MockService) NewMux(ctx context.Context, conn *grpc.ClientConn) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()

	if err := pb_http.RegisterCertificateAuthorityServiceHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	return mux, nil
}
