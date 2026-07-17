// Copyright 2026 Google LLC
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

package mockgeminidataanalytics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/geminidataanalytics/apiv1beta/geminidataanalyticspb"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked geminidataanalytics service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	v1beta *GeminiDataAnalyticsV1beta
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
	}
	s.v1beta = &GeminiDataAnalyticsV1beta{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"geminidataanalytics.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterDataChatServiceServer(grpcServer, s.v1beta)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(pb.NewDataChatServiceClient(conn))

	return grpcMux, nil
}
