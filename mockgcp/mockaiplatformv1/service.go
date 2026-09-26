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

package mockaiplatformv1

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked aiplatform v1 service.
type MockService struct {
	*common.MockEnvironment

	storage storage.Storage

	operations *operations.Operations
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"{region}-aiplatform.googleapis.com", "{region}-aiplatform.googleapis.com:443"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterPipelineServiceServer(grpcServer, &pipelineService{MockService: s})
	pb.RegisterFeatureOnlineStoreAdminServiceServer(grpcServer, &featureOnlineStoreAdminService{MockService: s})
	pb.RegisterModelServiceServer(grpcServer, &modelService{MockService: s})
	pb.RegisterPersistentResourceServiceServer(grpcServer, &persistentResourceService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	mux.AddService(pb.NewPipelineServiceClient(conn))
	mux.AddService(pb.NewFeatureOnlineStoreAdminServiceClient(conn))
	mux.AddService(pb.NewModelServiceClient(conn))
	mux.AddService(pb.NewPersistentResourceServiceClient(conn))

	mux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	return mux, nil
}

func computeEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}
