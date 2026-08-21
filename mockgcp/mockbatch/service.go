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

package mockbatch

// +tool:mockgcp-service
// http.host: batch.googleapis.com
// proto.service: google.cloud.batch.v1.BatchService

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	pb "cloud.google.com/go/batch/apiv1/batchpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	pb_v1alpha "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/batch/resourceallowance/pb"
)

// MockService represents a mocked batch service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1      *BatchV1
	v1alpha *BatchV1Alpha
}

type BatchV1 struct {
	*MockService
	pb.UnimplementedBatchServiceServer
}

type BatchV1Alpha struct {
	*MockService
	pb_v1alpha.UnimplementedBatchServiceServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &BatchV1{MockService: s}
	s.v1alpha = &BatchV1Alpha{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"batch.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterBatchServiceServer(grpcServer, s.v1)
	pb_v1alpha.RegisterBatchServiceServer(grpcServer, s.v1alpha)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building batch grpc mux: %w", err)
	}

	mux.AddService(pb.NewBatchServiceClient(conn))
	mux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	return mux, nil
}
