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

package mockpubsub

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked pubsub service.
type MockService struct {
	*common.MockEnvironment
	operations *operations.Operations

	topics        storage.TypedStorage[*pb.Topic]
	snapshots     storage.TypedStorage[*pb.Snapshot]
	schemas       storage.TypedStorage[*pb.Schema]
	subscriptions storage.TypedStorage[*pb.Subscription]
}

// New creates a MockService.
func New(env *common.MockEnvironment, store storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		operations:      operations.NewOperationsService(store),
	}
	s.topics = storage.For[*pb.Topic](store)
	s.snapshots = storage.For[*pb.Snapshot](store)
	s.schemas = storage.For[*pb.Schema](store)
	s.subscriptions = storage.For[*pb.Subscription](store)
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"pubsub.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterPublisherServer(grpcServer, &publisherService{MockService: s})
	pb.RegisterSubscriberServer(grpcServer, &subscriberService{MockService: s})
	pb.RegisterSchemaServiceServer(grpcServer, &schemaService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	mux.AddService(pb.NewPublisherClient(conn))
	mux.AddService(pb.NewSubscriberClient(conn))
	mux.AddService(pb.NewSchemaServiceClient(conn))

	return mux, nil
}
