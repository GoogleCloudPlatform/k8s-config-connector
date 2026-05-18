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

package mockmetastore

// +tool:mockgcp-service
// http.host: metastore.googleapis.com
// proto.service: google.cloud.metastore.v1.DataprocMetastore

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked Dataproc Metastore service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1           *DataprocMetastoreV1
	federationV1 *DataprocMetastoreFederationV1
}

type DataprocMetastoreV1 struct {
	*MockService
	pb.UnimplementedDataprocMetastoreServer
}

type DataprocMetastoreFederationV1 struct {
	*MockService
	pb.UnimplementedDataprocMetastoreFederationServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &DataprocMetastoreV1{MockService: s}
	s.federationV1 = &DataprocMetastoreFederationV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"metastore.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterDataprocMetastoreServer(grpcServer, s.v1)
	pb.RegisterDataprocMetastoreFederationServer(grpcServer, s.federationV1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	mux.AddService(pb.NewDataprocMetastoreClient(conn))
	mux.AddService(pb.NewDataprocMetastoreFederationClient(conn))
	mux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	return mux, nil
}
