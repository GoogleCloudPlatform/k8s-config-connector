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

package mockbilling

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"
	"k8s.io/klog/v2"

	pb "cloud.google.com/go/billing/apiv1/billingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked billing service.
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
	return []string{"cloudbilling.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterCloudBillingServer(grpcServer, &BillingV1{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	// options := httpmux.Options{
	// 		// This API sends e.g. billingEnabled: false, billingAccountName: ""
	// 		EmitUnpopulated: true,
	// 	}

	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(pb.NewCloudBillingClient(conn))

	return grpcMux, nil
}

// Functionality for use by mocks
func (s *MockService) MockCreateBillingAccount(billingAccount *pb.BillingAccount) (*pb.BillingAccount, error) {
	name, err := s.parseBillingAccountName(billingAccount.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	if err := s.storage.Create(context.Background(), fqn, billingAccount); err != nil {
		return nil, err
	}
	klog.Infof("created billing account %s", fqn)
	return billingAccount, nil
}
