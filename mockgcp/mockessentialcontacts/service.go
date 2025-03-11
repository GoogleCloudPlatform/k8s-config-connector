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

// +tool:mockgcp-service
// http.host: essentialcontacts.googleapis.com
// proto.service: google.cloud.essentialcontacts.v1.EssentialContactsService

package mockessentialcontacts

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/essentialcontacts/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked essentialcontacts service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1 *EssentialContactsV1
}

type EssentialContactsV1 struct {
	*MockService
	pb.UnimplementedEssentialContactsServiceServer
}

func (s *EssentialContactsV1) CreateContact(ctx context.Context, req *pb.CreateContactRequest) (*pb.Contact, error) {
	name, err := s.storage.GenerateName(req.Parent)
	if err != nil {
		return nil, err
	}

	contact := req.Contact
	contact.Name = name

	if err := s.storage.Create(ctx, contact); err != nil {
		return nil, err
	}

	return contact, nil
}


// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &EssentialContactsV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"essentialcontacts.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterEssentialContactsServiceServer(grpcServer, s.v1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{}, pb.RegisterEssentialContactsServiceHandler)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func (s *MockService) Do(req *http.Request) (*http.Response, error) {
	switch req.URL.Path {
	case "/v1/projects/mock-project/contacts":
		if req.Method == http.MethodPost {
			var createReq pb.CreateContactRequest
			if err := httpmux.ReadRequest(req, &createReq); err != nil {
				return nil, err
			}
			contact, err := s.v1.CreateContact(req.Context(), &createReq)
			if err != nil {
				return nil, err
			}
			return httpmux.RespondProto(req, contact)
		}
	}
	return nil, common.NotFoundError(req.URL.Path)
}

