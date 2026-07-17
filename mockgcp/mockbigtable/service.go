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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

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
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	mux.AddService(pb.NewBigtableInstanceAdminClient(conn))
	mux.AddService(pb.NewBigtableTableAdminClient(conn))
	mux.AddOperationsPath("/v2/{prefix=**}/operations/{name}", conn)

	// Check if the 'view' field has been added upstream yet.
	// If it has, fail so we know to remove this workaround.
	{
		getDesc := (&pb.GetMaterializedViewRequest{}).ProtoReflect().Descriptor()
		listDesc := (&pb.ListMaterializedViewsRequest{}).ProtoReflect().Descriptor()
		if getDesc.Fields().ByName("view") != nil {
			return nil, fmt.Errorf("view field now exists in GetMaterializedViewRequest; please revert the workaround in mockgcp/mockbigtable/service.go")
		}
		if listDesc.Fields().ByName("view") != nil {
			return nil, fmt.Errorf("view field now exists in ListMaterializedViewsRequest; please revert the workaround in mockgcp/mockbigtable/service.go")
		}
	}

	// Workaround for https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8299
	// gcloud sends view=SCHEMA_VIEW for materialized views, but our generated protobuf
	// GetMaterializedViewRequest and ListMaterializedViewsRequest do not have a view field.
	// We strip it here so that dynamicgrpcgateway doesn't fail with ErrFieldNotFound.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "materializedViews") {
			q := r.URL.Query()
			if q.Has("view") {
				q.Del("view")
				r.URL.RawQuery = q.Encode()
			}
		}
		mux.ServeHTTP(w, r)
	}), nil
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
