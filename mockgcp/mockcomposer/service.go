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

package mockcomposer

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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orchestration/airflow/service/v1"
)

// MockService represents a mocked composer service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projects   *projects.ProjectStore
	operations *operations.Operations

	v1 *ComposerV1
	// v1beta1 *ComposerV1Beta1
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		kube:       env.GetKubeClient(),
		storage:    storage,
		projects:   env.GetProjects(),
		operations: operations.NewOperationsService(storage),
	}
	s.v1 = &ComposerV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHost() string {
	return "composer.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterEnvironmentsServer(grpcServer, s.v1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()

	// TODO: Is all of v1beta1 a direct mapping to v1?
	rewriteV1Beta1ToV1 := func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		u := r.URL
		u.Path = "/v1/" + strings.TrimPrefix(u.Path, "/v1beta1/")
		r.URL = u

		mux.ServeHTTP(w, r)
	}

	if err := mux.HandlePath("GET", "/v1beta1/{path=**}", rewriteV1Beta1ToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("POST", "/v1beta1/{path=**}", rewriteV1Beta1ToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("DELETE", "/v1beta1/{path=**}", rewriteV1Beta1ToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("PATCH", "/v1beta1/{path=**}", rewriteV1Beta1ToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("PUT", "/v1beta1/{path=**}", rewriteV1Beta1ToV1); err != nil {
		return nil, err
	}

	if err := pb.RegisterEnvironmentsHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	return mux, nil
}
