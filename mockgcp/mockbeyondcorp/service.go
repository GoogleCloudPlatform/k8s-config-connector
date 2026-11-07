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

// +tool:mockgcp-service
// http.host: beyondcorp.googleapis.com
// proto.service: google.cloud.beyondcorp.appgateways.v1.AppGatewaysService
// proto.service: google.cloud.beyondcorp.appconnectors.v1.AppConnectorsService

package mockbeyondcorp

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	appconnectionspb "cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb"
	appconnectorspb "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb"
	appgatewayspb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"
	clientconnectorservicespb "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1/clientconnectorservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked beyondcorp service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1                      *AppGatewaysV1
	appConnections          *appConnectionsServer
	appConnectors           *appConnectorServer
	clientConnectorServices *clientConnectorServicesServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &AppGatewaysV1{MockService: s}
	s.appConnections = &appConnectionsServer{MockService: s}
	s.appConnectors = &appConnectorServer{MockService: s}
	s.clientConnectorServices = &clientConnectorServicesServer{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"beyondcorp.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	appgatewayspb.RegisterAppGatewaysServiceServer(grpcServer, s.v1)
	appconnectorspb.RegisterAppConnectorsServiceServer(grpcServer, s.appConnectors)
	appconnectionspb.RegisterAppConnectionsServiceServer(grpcServer, s.appConnections)
	clientconnectorservicespb.RegisterClientConnectorServicesServiceServer(grpcServer, s.clientConnectorServices)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(appgatewayspb.NewAppGatewaysServiceClient(conn))
	grpcMux.AddService(appconnectorspb.NewAppConnectorsServiceClient(conn))
	grpcMux.AddService(appconnectionspb.NewAppConnectionsServiceClient(conn))
	grpcMux.AddService(clientconnectorservicespb.NewClientConnectorServicesServiceClient(conn))

	grpcMux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	return grpcMux, nil
}
