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
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

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
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux(runtime.WithErrorHandler(customErrorHandler))

	if err := pb_v1.RegisterProjectsHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	if err := pb_v3.RegisterProjectsHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	return mux, nil
}

type wrappedStatus struct {
	Error *wrappedError `json:"error,omitempty"`
}

type wrappedError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}

// customErrorHandler wraps errors in an error blockk
func customErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	s := status.Convert(err)
	// pb := s.Proto()

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	httpStatusCode := runtime.HTTPStatusFromCode(s.Code())
	wrapped := &wrappedStatus{
		Error: &wrappedError{
			Code:    httpStatusCode,
			Message: s.Message(),
		},
	}

	switch s.Code() {
	case codes.PermissionDenied:
		wrapped.Error.Status = "PERMISSION_DENIED"
	case codes.AlreadyExists:
		wrapped.Error.Status = "ALREADY_EXISTS"
	}

	buf, merr := json.Marshal(wrapped)
	if merr != nil {
		klog.Warningf("Failed to marshal error message %q: %v", s, merr)
		runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
		return
	}

	w.WriteHeader(httpStatusCode)
	if _, err := w.Write(buf); err != nil {
		klog.Warningf("Failed to write response: %v", err)
	}
}
