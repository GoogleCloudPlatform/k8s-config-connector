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

package mockeventarc

// +tool:mockgcp-service
// http.host: eventarc.googleapis.com
// proto.service: google.cloud.eventarc.v1.Eventarc

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked eventarc service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1 *EventarcV1
}

type EventarcV1 struct {
	*MockService
	pb.UnimplementedEventarcServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &EventarcV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"eventarc.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterEventarcServer(grpcServer, s.v1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	mux.AddService(pb.NewEventarcClient(conn))
	mux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	return mux, nil
}
