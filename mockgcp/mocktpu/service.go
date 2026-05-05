// Copyright 2025 Google LLC
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
// http.host: tpu.googleapis.com
// proto.service: google.cloud.tpu.v2.Tpu

package mocktpu

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pbv1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/tpu/v1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/tpu/v2"
	pbv2alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/tpu/v2alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked tpu service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	tpuServer         *TpuServer
	tpuV1Server       *TpuV1Server
	tpuV2Alpha1Server *TpuV2Alpha1Server
}

type TpuServer struct {
	*MockService
	pb.UnimplementedTpuServer
}

type TpuV1Server struct {
	*MockService
	pbv1.UnimplementedTpuServer
}

type TpuV2Alpha1Server struct {
	*MockService
	pbv2alpha1.UnimplementedTpuServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.tpuServer = &TpuServer{MockService: s}
	s.tpuV1Server = &TpuV1Server{MockService: s}
	s.tpuV2Alpha1Server = &TpuV2Alpha1Server{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"tpu.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pbv1.RegisterTpuServer(grpcServer, s.tpuV1Server)
	pb.RegisterTpuServer(grpcServer, s.tpuServer)
	pbv2alpha1.RegisterTpuServer(grpcServer, s.tpuV2Alpha1Server)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pbv1.RegisterTpuHandler,
		pb.RegisterTpuHandler,
		pbv2alpha1.RegisterTpuHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"),
		s.operations.RegisterOperationsPath("/v2/{prefix=**}/operations/{name}"),
		s.operations.RegisterOperationsPath("/v2alpha1/{prefix=**}/operations/{name}"))
	if err != nil {
		return nil, err
	}

	mux.RewriteHeaders = func(ctx context.Context, response http.ResponseWriter, payload proto.Message) {
		response.Header().Del("Cache-Control")
	}

	return mux, nil
}
