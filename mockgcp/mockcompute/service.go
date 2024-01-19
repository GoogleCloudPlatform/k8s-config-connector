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

package mockcompute

import (
	"context"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

// MockService represents a mocked compute service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projects   projects.ProjectStore
	operations *operations.Operations

	networksv1 *NetworksV1
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		kube:       env.GetKubeClient(),
		storage:    storage,
		projects:   env.GetProjects(),
		operations: operations.NewOperationsService(storage),
	}
	s.networksv1 = &NetworksV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHost() string {
	return "compute.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterNetworksServer(grpcServer, s.networksv1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()

	// Terraform uses the /beta/ endpoints, but we have protos only for v1.
	// Also, we probably want to be implementing the newer versions
	// as that makes it easier to move KCC to newer API versions.
	// So far, it seems that all of beta is a direct mapping to v1 - though
	// I'm sure eventually we'll find something that needs special handling.
	rewriteBetaToV1 := func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		u := r.URL
		u.Path = "/compute/v1/" + strings.TrimPrefix(u.Path, "/compute/beta/")
		r.URL = u

		mux.ServeHTTP(w, r)
	}

	if err := mux.HandlePath("GET", "/compute/beta/{path=**}", rewriteBetaToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("POST", "/compute/beta/{path=**}", rewriteBetaToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("DELETE", "/compute/beta/{path=**}", rewriteBetaToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("PATCH", "/compute/beta/{path=**}", rewriteBetaToV1); err != nil {
		return nil, err
	}

	if err := pb.RegisterNetworksHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	return mux, nil
}
