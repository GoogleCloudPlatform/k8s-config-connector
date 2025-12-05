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

// +tool:mockgcp-support
// proto.service: google.cloud.clouddms.v1.DataMigrationService
// proto.message: google.cloud.clouddms.v1.MigrationJob

package mockclouddms

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *DataMigrationServiceV1) GetMigrationJob(ctx context.Context, req *pb.GetMigrationJobRequest) (*pb.MigrationJob, error) {
	name, err := s.parseMigrationJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MigrationJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "MigrationJob %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataMigrationServiceV1) CreateMigrationJob(ctx context.Context, req *pb.CreateMigrationJobRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/migrationJobs/" + req.MigrationJobId
	name, err := s.parseMigrationJobName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetMigrationJob()).(*pb.MigrationJob)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	if obj.State == pb.MigrationJob_STATE_UNSPECIFIED {
		obj.State = pb.MigrationJob_NOT_STARTED
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "create",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataMigrationServiceV1) DeleteMigrationJob(ctx context.Context, req *pb.DeleteMigrationJobRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseMigrationJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.MigrationJob{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type migrationJobName struct {
	Project          *projects.ProjectData
	Location         string
	MigrationJobName string
}

func (n *migrationJobName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", n.Project.ID, n.Location, n.MigrationJobName)
}

// parseMigrationJobName parses a string into an migrationJobName.
// The expected form is `projects/*/locations/*/migrationJobs/*`.
func (s *MockService) parseMigrationJobName(name string) (*migrationJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "migrationJobs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &migrationJobName{
			Project:          project,
			Location:         tokens[3],
			MigrationJobName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
