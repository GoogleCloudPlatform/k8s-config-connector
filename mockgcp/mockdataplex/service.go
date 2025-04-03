// Copyright 2024 Google LLC
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

package mockdataplex

// +tool:mockgcp-service
// http.host: dataplex.googleapis.com
// proto.service: google.cloud.dataplex.v1.ContentService

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataplex/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

// MockService represents a mocked dataplex service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	// Embed the previously defined DataplexV1 if merging
	// dataplexV1 *DataplexV1

	contentV1 *ContentServiceV1
}

type ContentServiceV1 struct {
	*MockService
	pb.UnimplementedContentServiceServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	// If merging, initialize other services like:
	// s.dataplexV1 = &DataplexV1{MockService: s}
	s.contentV1 = &ContentServiceV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"dataplex.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	// If merging, include registrations for other services like:
	// pb.RegisterDataplexServiceServer(grpcServer, s.dataplexV1)
	pb.RegisterContentServiceServer(grpcServer, s.contentV1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		// If merging, include handlers for other services like:
		// pb.RegisterDataplexServiceHandler,
		pb.RegisterContentServiceHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"))
	if err != nil {
		return nil, err
	}

	// Potentially shared mux configuration like RewriteHeaders or RewriteError
	mux.RewriteHeaders = func(ctx context.Context, response http.ResponseWriter, payload proto.Message) {
		response.Header().Del("Cache-Control")
	}
	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 {
			error.Errors = nil
		}
	}

	return mux, nil
}

// Implement IAM methods potentially shared across services or specific to ContentService
func (s *ContentServiceV1) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	// Simplified IAM Get - return empty policy or error
	// Replace with actual IAM logic if needed
	return &iampb.Policy{}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}

func (s *ContentServiceV1) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	// Simplified IAM Set - return request policy or error
	// Replace with actual IAM logic if needed
	return req.Policy, nil
	// return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}

func (s *ContentServiceV1) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	// Simplified IAM Test - return all permissions or error
	// Replace with actual IAM logic if needed
	return &iampb.TestIamPermissionsResponse{Permissions: req.Permissions}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}
