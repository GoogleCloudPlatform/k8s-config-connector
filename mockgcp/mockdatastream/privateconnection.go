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
// proto.service: google.cloud.datastream.v1.Datastream
// proto.message: google.cloud.datastream.v1.PrivateConnection

package mockdatastream

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

func (s *DatastreamV1) GetPrivateConnection(ctx context.Context, req *pb.GetPrivateConnectionRequest) (*pb.PrivateConnection, error) {
	name, err := s.parsePrivateConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PrivateConnection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DatastreamV1) CreatePrivateConnection(ctx context.Context, req *pb.CreatePrivateConnectionRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/privateConnections/%s", req.GetParent(), req.GetPrivateConnectionId())
	name, err := s.parsePrivateConnectionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetPrivateConnection()).(*pb.PrivateConnection)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.PrivateConnection_CREATED

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

func (s *DatastreamV1) DeletePrivateConnection(ctx context.Context, req *pb.DeletePrivateConnectionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePrivateConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.PrivateConnection{}
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

type privateConnectionName struct {
	Project             *projects.ProjectData
	Location            string
	PrivateConnectionID string
}

func (n *privateConnectionName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/privateConnections/" + n.PrivateConnectionID
}

// parsePrivateConnectionName parses a string into a privateConnectionName.
// The expected form is `projects/*/locations/*/privateConnections/*`.
func (s *MockService) parsePrivateConnectionName(name string) (*privateConnectionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "privateConnections" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &privateConnectionName{
			Project:             project,
			Location:            tokens[3],
			PrivateConnectionID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
