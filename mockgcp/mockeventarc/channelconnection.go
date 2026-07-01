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

package mockeventarc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *EventarcV1) GetChannelConnection(ctx context.Context, req *pb.GetChannelConnectionRequest) (*pb.ChannelConnection, error) {
	name, err := s.parseChannelConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ChannelConnection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *EventarcV1) CreateChannelConnection(ctx context.Context, req *pb.CreateChannelConnectionRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/channelConnections/%s", req.GetParent(), req.GetChannelConnectionId())
	name, err := s.parseChannelConnectionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetChannelConnection()).(*pb.ChannelConnection)
	obj.Name = fqn
	obj.Uid = name.ChannelConnection
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Return an LRO that does not finish immediately
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	lro, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, err
	}
	lro.Done = false
	lro.Metadata, err = anypb.New(lroMetadata)
	if err != nil {
		return nil, err
	}
	lro.Metadata.TypeUrl = "type.googleapis.com/google.cloud.eventarc.v1.OperationMetadata"
	lro.Name = fmt.Sprintf("projects/%s/locations/%s/operations/%s", name.Project.ID, name.Location, strings.Split(lro.Name, "/")[len(strings.Split(lro.Name, "/"))-1])

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		return obj, nil
	})
}

func (s *EventarcV1) DeleteChannelConnection(ctx context.Context, req *pb.DeleteChannelConnectionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseChannelConnectionName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.ChannelConnection{}

	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	deletedObj.UpdateTime = nil
	deletedObj.CreateTime = nil
	deletedObj.Uid = ""

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return deletedObj, nil
	})
}

type channelConnectionName struct {
	Project           *projects.ProjectData
	Location          string
	ChannelConnection string
}

func (n *channelConnectionName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/channelConnections/%s", n.Project.ID, n.Location, n.ChannelConnection)
}

func (s *MockService) parseChannelConnectionName(name string) (*channelConnectionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "channelConnections" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &channelConnectionName{
			Project:           project,
			Location:          tokens[3],
			ChannelConnection: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
