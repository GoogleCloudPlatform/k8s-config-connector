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
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/admin/v1"
	pbv1beta "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/v1beta"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked IAM service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projects   *projects.ProjectStore
	operations *operations.Operations

	serverV1              *ServerV1
	workloadIdentityPools *WorkloadIdentityPoolsV1Beta
}

// New creates a MockService
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		kube:       env.GetKubeClient(),
		storage:    storage,
		operations: operations.NewOperationsService(storage),
		projects:   env.GetProjects(),
	}
	s.serverV1 = &ServerV1{MockService: s}
	s.workloadIdentityPools = &WorkloadIdentityPoolsV1Beta{MockService: s}
	return s
}

func (s *MockService) ExpectedHost() string {
	return "iam.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterIAMServer(grpcServer, s.serverV1)
	pbv1beta.RegisterWorkloadIdentityPoolsServer(grpcServer, s.workloadIdentityPools)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()
	if err := pb.RegisterIAMHandler(ctx, mux, conn); err != nil {
		return nil, err
	}
	if err := pbv1beta.RegisterWorkloadIdentityPoolsHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	rewriteV1ToV1Beta := func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		u := r.URL
		u.Path = "/v1beta/" + strings.TrimPrefix(u.Path, "/v1/")
		r.URL = u

		mux.ServeHTTP(w, r)
	}

	if err := mux.HandlePath("GET", "/v1/projects/{project}/locations/{location}/workloadIdentityPools/{name}", rewriteV1ToV1Beta); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("DELETE", "/v1/projects/{project}/locations/{location}/workloadIdentityPools/{name}", rewriteV1ToV1Beta); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("POST", "/v1/projects/{project}/locations/{location}/workloadIdentityPools", rewriteV1ToV1Beta); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("PATCH", "/v1/projects/{project}/locations/{location}/workloadIdentityPools/{name}", rewriteV1ToV1Beta); err != nil {
		return nil, err
	}

	if err := mux.HandlePath("GET", "/v1/projects/{project}/locations/{location}/workloadIdentityPools/{name}/providers/{providerName}", rewriteV1ToV1Beta); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("DELETE", "/v1/projects/{project}/locations/{location}/workloadIdentityPools/{name}/providers/{providerName}", rewriteV1ToV1Beta); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("POST", "/v1/projects/{project}/locations/{location}/workloadIdentityPools/{name}/providers", rewriteV1ToV1Beta); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("PATCH", "/v1/projects/{project}/locations/{location}/workloadIdentityPools/{name}/providers/{providerName}", rewriteV1ToV1Beta); err != nil {
		return nil, err
	}

	return mux, nil
}
