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

package mockdatastream

// +tool:mockgcp-service
// http.host: datastream.googleapis.com
// proto.service: google.cloud.datastream.v1.Datastream

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datastream/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *DatastreamV1) GetConnectionProfile(ctx context.Context, req *pb.GetConnectionProfileRequest) (*pb.ConnectionProfile, error) {
	name, err := s.parseConnectionProfileName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ConnectionProfile{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DatastreamV1) CreateConnectionProfile(ctx context.Context, req *pb.CreateConnectionProfileRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/connectionProfiles/%s", req.GetParent(), req.GetConnectionProfileId())
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
	if obj.Connectivity == nil {
		obj.Connectivity = &pb.ConnectionProfile_StaticServiceIpConnectivity{}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DatastreamV1) UpdateConnectionProfile(ctx context.Context, req *pb.UpdateConnectionProfileRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetConnectionProfile().GetName()
	name, err := s.parseConnectionProfileName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.ConnectionProfile{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.UpdateTime = timestamppb.New(time.Now())

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetConnectionProfile().GetDisplayName()
		}
		// TODO: handle other fields
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DatastreamV1) DeleteConnectionProfile(ctx context.Context, req *pb.DeleteConnectionProfileRequest) (*longrunningpb.Operation, error) {
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
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type connectionProfileName struct {
	Project             *projects.ProjectData
	Location            string
	ConnectionProfileID string
}

func (n *connectionProfileName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/connectionProfiles/" + n.ConnectionProfileID
}

// parseConnectionProfileName parses a string into a connectionProfileName.
// The expected form is `projects/*/locations/*/connectionProfiles/*`.
func (s *MockService) parseConnectionProfileName(name string) (*connectionProfileName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "connectionProfiles" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &connectionProfileName{
			Project:             project,
			Location:            tokens[3],
			ConnectionProfileID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
