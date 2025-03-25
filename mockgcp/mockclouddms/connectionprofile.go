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
// proto.message: google.cloud.clouddms.v1.ConnectionProfile

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/clouddms/v1"
)

func (s *DataMigrationServiceV1) GetConnectionProfile(ctx context.Context, req *pb.GetConnectionProfileRequest) (*pb.ConnectionProfile, error) {
	name, err := s.parseConnectionProfileName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ConnectionProfile{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ConnectionProfile %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataMigrationServiceV1) CreateConnectionProfile(ctx context.Context, req *pb.CreateConnectionProfileRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/connectionProfiles/" + req.ConnectionProfileId
	name, err := s.parseConnectionProfileName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetConnectionProfile()).(*pb.ConnectionProfile)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.ConnectionProfile_READY
	if obj.GetMysql().GetPassword() != "" {
		obj.GetMysql().Password = ""
		obj.GetMysql().PasswordSet = true
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

func (s *DataMigrationServiceV1) UpdateConnectionProfile(ctx context.Context, req *pb.UpdateConnectionProfileRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseConnectionProfileName(req.GetConnectionProfile().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.ConnectionProfile{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	now := time.Now()
	updated := proto.Clone(existing).(*pb.ConnectionProfile)

	updated.UpdateTime = timestamppb.New(now)
	updated.DisplayName = req.GetConnectionProfile().GetDisplayName()

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

func (s *DataMigrationServiceV1) DeleteConnectionProfile(ctx context.Context, req *pb.DeleteConnectionProfileRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseConnectionProfileName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ConnectionProfile{}
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

type connectionProfileName struct {
	Project               *projects.ProjectData
	Location              string
	ConnectionProfileName string
}

func (n *connectionProfileName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", n.Project.ID, n.Location, n.ConnectionProfileName)
}

// parseConnectionProfileName parses a string into an connectionProfileName.
// The expected form is `projects/*/locations/*/connectionProfiles/*`.
func (s *MockService) parseConnectionProfileName(name string) (*connectionProfileName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "connectionProfiles" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &connectionProfileName{
			Project:               project,
			Location:              tokens[3],
			ConnectionProfileName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
