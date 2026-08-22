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
// http.host: datalabeling.googleapis.com
// proto.service: google.cloud.datalabeling.v1beta1.DataLabelingService

package mockdatalabeling

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	longrunning "google.golang.org/genproto/googleapis/longrunning"

	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked datalabeling service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations
}

type DataLabelingV1beta1 struct {
	*MockService
	pb.UnimplementedDataLabelingServiceServer
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
	return []string{"datalabeling.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterDataLabelingServiceServer(grpcServer, &DataLabelingV1beta1{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	mux.AddService(pb.NewDataLabelingServiceClient(conn))

	return mux, nil
}

type InstructionName struct {
	Project       *projects.ProjectData
	InstructionID string
}

func (n *InstructionName) String() string {
	return fmt.Sprintf("projects/%s/instructions/%s", n.Project.ID, n.InstructionID)
}

func (s *MockService) parseInstructionName(name string) (*InstructionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "instructions" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &InstructionName{
			Project:       project,
			InstructionID: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *DataLabelingV1beta1) GetInstruction(ctx context.Context, req *pb.GetInstructionRequest) (*pb.Instruction, error) {
	name, err := s.parseInstructionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instruction{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DataLabelingV1beta1) CreateInstruction(ctx context.Context, req *pb.CreateInstructionRequest) (*longrunning.Operation, error) {
	// Generate unique instruction ID
	instructionID := "instruction-" + uuid.New().String()[:8]

	projectName, err := projects.ParseProjectName(req.Parent)
	if err != nil {
		return nil, err
	}
	project, err := s.Projects.GetProject(projectName)
	if err != nil {
		return nil, err
	}

	name := &InstructionName{
		Project:       project,
		InstructionID: instructionID,
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.Instruction).(*pb.Instruction)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.CreateInstructionMetadata{
		Instruction: fqn,
		CreateTime:  timestamppb.New(now),
	}

	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *DataLabelingV1beta1) DeleteInstruction(ctx context.Context, req *pb.DeleteInstructionRequest) (*emptypb.Empty, error) {
	name, err := s.parseInstructionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Instruction{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
