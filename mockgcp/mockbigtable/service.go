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

package mockbigtable

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	grpcpb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/bigtable/admin/v2"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked bigtable service.
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
	return []string{"bigtableadmin.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterBigtableInstanceAdminServer(grpcServer, &instanceAdminServer{MockService: s})
	pb.RegisterBigtableTableAdminServer(grpcServer, &tableAdminServer{MockService: s})
	s.operations.RegisterGRPCServices(grpcServer)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		grpcpb.RegisterBigtableInstanceAdminHandler,
		grpcpb.RegisterBigtableTableAdminHandler,
		s.operations.RegisterOperationsPath("/v2/{prefix=**}/operations/{name}"))
	if err != nil {
		return nil, err
	}

	// Returns slightly non-standard errors
	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 {
			error.Errors = nil
		}
	}

	return mux, nil
}

func (s *MockService) RunTestCommand(ctx context.Context, serviceName string, command string) error {
	switch command {
	case "ScaleUp":
		return s.runScaleUpCommand(ctx)
	default:
		return fmt.Errorf("test-command %q not known", command)
	}
}

func (s *MockService) runScaleUpCommand(ctx context.Context) error {
	var clusters []*pb.Cluster

	findKind := (&pb.Cluster{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: "",
	}, func(obj proto.Message) error {
		cluster := obj.(*pb.Cluster)
		clusters = append(clusters, cluster)

		return nil
	}); err != nil {
		return err
	}

	for _, cluster := range clusters {
		maxServeNodes := cluster.GetClusterConfig().GetClusterAutoscalingConfig().GetAutoscalingLimits().GetMaxServeNodes()
		if cluster.GetServeNodes() < maxServeNodes {
			cluster.ServeNodes++
			fqn := cluster.Name
			if err := s.storage.Update(ctx, fqn, cluster); err != nil {
				return err
			}
		}
	}

	return nil
}
