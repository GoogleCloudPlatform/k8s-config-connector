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

// +tool:mockgcp-service
// http.host: livestream.googleapis.com
// proto.service: google.cloud.video.livestream.v1.LivestreamService

package mocklivestream

import (
	"context"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked livestream service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	return &MockService{
		MockEnvironment: env,
		storage:         storage,
	}
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"livestream.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	return nil, nil
}
