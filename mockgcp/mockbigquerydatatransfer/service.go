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

package mockbigquerydatatransfer

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked bigquerydatatransfer service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"bigquerydatatransfer.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	datatransferpb.RegisterDataTransferServiceServer(grpcServer, &dataTransferService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	mux.AddService(datatransferpb.NewDataTransferServiceClient(conn))

	return mux, nil
}
