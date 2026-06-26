// Copyright 2023 Google LLC
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

package mockcertificatemanager

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

type cmStorage struct {
	storage.Storage
}

func (s *cmStorage) Get(ctx context.Context, fqn string, dest proto.Message) error {
	err := s.Storage.Get(ctx, fqn, dest)
	if err != nil && status.Code(err) == codes.NotFound {
		typeName := string(dest.ProtoReflect().Descriptor().Name())
		if typeName == "Certificate" {
			return status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		if typeName == "CertificateMap" || typeName == "CertificateMapEntry" {
			return status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
	}
	return err
}

func (s *cmStorage) Delete(ctx context.Context, fqn string, dest proto.Message) error {
	err := s.Storage.Delete(ctx, fqn, dest)
	if err != nil && status.Code(err) == codes.NotFound {
		typeName := string(dest.ProtoReflect().Descriptor().Name())
		if typeName == "Certificate" {
			return status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		if typeName == "CertificateMap" || typeName == "CertificateMapEntry" {
			return status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
	}
	return err
}

// MockService represents a mocked certificatemanager service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1 *CertificateManagerV1
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	wrappedStorage := &cmStorage{Storage: storage}
	s := &MockService{
		MockEnvironment: env,
		storage:         wrappedStorage,
		operations:      operations.NewOperationsService(wrappedStorage),
	}
	s.v1 = &CertificateManagerV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"certificatemanager.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterCertificateManagerServer(grpcServer, s.v1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	grpcMux.AddService(pb.NewCertificateManagerClient(conn))
	grpcMux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	return grpcMux, nil
}
