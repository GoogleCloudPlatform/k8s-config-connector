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

package mockmonitoring

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	monitoringpb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	dashboardpb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	metricsscopepb "cloud.google.com/go/monitoring/metricsscope/apiv1/metricsscopepb"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked apikeys service.
type MockService struct {
	*common.MockEnvironment
	storage    storage.Storage
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
	return []string{"monitoring.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	monitoringpb.RegisterAlertPolicyServiceServer(grpcServer, &AlertPolicyService{MockService: s})
	monitoringpb.RegisterGroupServiceServer(grpcServer, &GroupService{MockService: s})
	monitoringpb.RegisterMetricServiceServer(grpcServer, &metricService{MockService: s})
	monitoringpb.RegisterNotificationChannelServiceServer(grpcServer, &NotificationChannelService{MockService: s})
	monitoringpb.RegisterServiceMonitoringServiceServer(grpcServer, &serviceMonitoringService{MockService: s})
	monitoringpb.RegisterUptimeCheckServiceServer(grpcServer, &UptimeCheckService{MockService: s})

	dashboardpb.RegisterDashboardsServiceServer(grpcServer, &DashboardsService{MockService: s})

	metricsscopepb.RegisterMetricsScopesServer(grpcServer, &metricsScopeService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(monitoringpb.NewAlertPolicyServiceClient(conn))
	grpcMux.AddService(monitoringpb.NewGroupServiceClient(conn))
	grpcMux.AddService(monitoringpb.NewMetricServiceClient(conn))
	grpcMux.AddService(monitoringpb.NewNotificationChannelServiceClient(conn))
	grpcMux.AddService(monitoringpb.NewServiceMonitoringServiceClient(conn))
	grpcMux.AddService(monitoringpb.NewUptimeCheckServiceClient(conn))

	grpcMux.AddService(dashboardpb.NewDashboardsServiceClient(conn))

	grpcMux.AddService(metricsscopepb.NewMetricsScopesClient(conn))

	grpcMux.AddOperationsPath("/v1/{prefix=projects/*/databases/*}/operations/{name=**}", conn)

	return grpcMux, nil
}
