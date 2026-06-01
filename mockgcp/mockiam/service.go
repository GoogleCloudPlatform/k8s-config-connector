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

// +tool:mockgcp-service
// http.host: iam.googleapis.com
// proto.service: google.iam.admin.v1.IAM

package mockiam

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	"cloud.google.com/go/iam/admin/apiv1/adminpb"
	pbv2 "cloud.google.com/go/iam/apiv2/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked IAM service.
type MockService struct {
	*common.MockEnvironment
	storage    storage.Storage
	operations *operations.Operations
}

type IAMServer struct {
	*MockService
	adminpb.UnimplementedIAMServer
}

// New creates a MockService
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"iam.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	adminpb.RegisterIAMServer(grpcServer, &IAMServer{MockService: s})
	pbv2.RegisterPoliciesServer(grpcServer, &IAMV2PoliciesServer{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(adminpb.NewIAMClient(conn))
	grpcMux.AddService(pbv2.NewPoliciesClient(conn))

	return grpcMux, nil
}
