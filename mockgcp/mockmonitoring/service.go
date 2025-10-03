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
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	dashboardpb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/dashboard/v1"
	metricsscopepb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/metricsscope/v1"
	monitoringpb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/v3"
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
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		monitoringpb.RegisterAlertPolicyServiceHandler,
		monitoringpb.RegisterGroupServiceHandler,
		monitoringpb.RegisterMetricServiceHandler,
		monitoringpb.RegisterNotificationChannelServiceHandler,
		monitoringpb.RegisterServiceMonitoringServiceHandler,
		monitoringpb.RegisterUptimeCheckServiceHandler,
		dashboardpb.RegisterDashboardsServiceHandler,
		metricsscopepb.RegisterMetricsScopesHandler,
	)
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
