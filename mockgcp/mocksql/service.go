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

package mocksql

import (
	"context"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked sql service.
type MockService struct {
	storage storage.Storage

	projects   projects.ProjectStore
	operations *operations

	users *sqlUsersService
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		storage:    storage,
		projects:   env.Projects,
		operations: &operations{storage: storage},
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"sqladmin.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	s.users = &sqlUsersService{MockService: s}
	pb.RegisterSqlDatabasesServiceServer(grpcServer, &sqlDatabaseServer{MockService: s})
	pb.RegisterSqlInstancesServiceServer(grpcServer, &sqlInstancesService{MockService: s})
	pb.RegisterSqlUsersServiceServer(grpcServer, s.users)
	pb.RegisterSqlOperationsServiceServer(grpcServer, &sqlOperationsService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterSqlDatabasesServiceHandler,
		pb.RegisterSqlInstancesServiceHandler,
		pb.RegisterSqlUsersServiceHandler,
		pb.RegisterSqlOperationsServiceHandler)
	if err != nil {
		return nil, err
	}

	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 {
			for errIdx := range error.Errors {
				if strings.HasPrefix(error.Errors[errIdx].Message, "databaseInstance") {
					error.Errors[errIdx].Message = "The Cloud SQL instance does not exist."
					error.Errors[errIdx].Reason = "instanceDoesNotExist"
				}
			}
			if strings.HasPrefix(error.Message, "databaseInstance") {
				error.Message = "The Cloud SQL instance does not exist."
				error.Status = ""
			}
		}
	}

	return mux, nil
}
