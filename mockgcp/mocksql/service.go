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
	"fmt"
	"net/http"
	"strings"
	"sync"

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

	// mutex guards nextInstanceCreateFailure.
	mutex sync.Mutex
	// nextInstanceCreateFailure, if set, makes the next instances.insert fail asynchronously.
	// It is set through the RunTestCommand backdoor.
	nextInstanceCreateFailure *operationFailure
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

// RunTestCommand implements mockgcp.SupportsTestCommands, a backdoor for fault injection
// that script-driven tests trigger with a MockGCPBackdoor object.
func (s *MockService) RunTestCommand(ctx context.Context, serviceName string, command string) error {
	switch command {
	case "FailNextInstanceCreateWithStockout":
		// Real GCP can't be made to do this on demand: the create operation finishes with
		// status DONE and an ERROR_STOCKOUT error, and the instance is left in FAILED state.
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.nextInstanceCreateFailure = &operationFailure{
			Code:    "ERROR_STOCKOUT",
			Message: "The zone or region does not have sufficient resources to handle the request at the moment.",
		}
		return nil
	default:
		return fmt.Errorf("test-command %q not known", command)
	}
}

// takeNextInstanceCreateFailure returns and clears the pending instance-creation failure, if any.
func (s *MockService) takeNextInstanceCreateFailure() *operationFailure {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	failure := s.nextInstanceCreateFailure
	s.nextInstanceCreateFailure = nil
	return failure
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
	pb.RegisterSqlSslCertsServiceServer(grpcServer, &sqlSslCertsServer{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterSqlDatabasesServiceHandler,
		pb.RegisterSqlInstancesServiceHandler,
		pb.RegisterSqlUsersServiceHandler,
		pb.RegisterSqlOperationsServiceHandler,
		pb.RegisterSqlSslCertsServiceHandler)
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
