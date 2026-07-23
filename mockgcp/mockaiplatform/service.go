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

package mockaiplatform

import (
	"context"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked aiplatform service.
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
	return []string{"{region}-aiplatform.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterTensorboardServiceServer(grpcServer, &tensorboardService{MockService: s})
	pb.RegisterDatasetServiceServer(grpcServer, &datasetService{MockService: s})
	pb.RegisterEndpointServiceServer(grpcServer, &endpointService{MockService: s})
	pb.RegisterMetadataServiceServer(grpcServer, &metadataStoreService{MockService: s})
	pb.RegisterFeaturestoreServiceServer(grpcServer, &featurestoreService{MockService: s})
	pb.RegisterModelServiceServer(grpcServer, &modelService{MockService: s})
	pb.RegisterNotebookServiceServer(grpcServer, &notebookService{MockService: s})
	pb.RegisterScheduleServiceServer(grpcServer, &scheduleService{MockService: s})
	pb.RegisterExampleStoreServiceServer(grpcServer, &exampleStoreService{MockService: s})
	pb.RegisterFeatureOnlineStoreAdminServiceServer(grpcServer, &featureOnlineStoreAdminService{MockService: s})
	pb.RegisterDeploymentResourcePoolServiceServer(grpcServer, &deploymentResourcePoolService{MockService: s})
	pb.RegisterPipelineServiceServer(grpcServer, &pipelineService{MockService: s})
	pb.RegisterJobServiceServer(grpcServer, &jobService{MockService: s})

	// Also register under v1 name so that v1 gRPC clients can call it
	desc := pb.FeatureOnlineStoreAdminService_ServiceDesc
	desc.ServiceName = "google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService"
	grpcServer.RegisterService(&desc, &featureOnlineStoreAdminService{MockService: s})

	descPipeline := pb.PipelineService_ServiceDesc
	descPipeline.ServiceName = "google.cloud.aiplatform.v1.PipelineService"
	grpcServer.RegisterService(&descPipeline, &pipelineService{MockService: s})

	descJob := pb.JobService_ServiceDesc
	descJob.ServiceName = "google.cloud.aiplatform.v1.JobService"
	grpcServer.RegisterService(&descJob, &jobService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	mux.AddService(pb.NewTensorboardServiceClient(conn))
	mux.AddService(pb.NewDatasetServiceClient(conn))
	mux.AddService(pb.NewEndpointServiceClient(conn))
	mux.AddService(pb.NewMetadataServiceClient(conn))
	mux.AddService(pb.NewFeaturestoreServiceClient(conn))
	mux.AddService(pb.NewModelServiceClient(conn))
	mux.AddService(pb.NewNotebookServiceClient(conn))
	mux.AddService(pb.NewScheduleServiceClient(conn))
	mux.AddService(pb.NewExampleStoreServiceClient(conn))
	mux.AddService(pb.NewFeatureOnlineStoreAdminServiceClient(conn))
	mux.AddService(pb.NewDeploymentResourcePoolServiceClient(conn))
	mux.AddService(pb.NewPipelineServiceClient(conn))
	mux.AddService(pb.NewJobServiceClient(conn))

	mux.AddOperationsPath("/v1beta1/{prefix=**}/operations/{name}", conn)
	mux.AddOperationsPath("/ui/{prefix=**}/operations/{name}", conn)

	return mux, nil
}
