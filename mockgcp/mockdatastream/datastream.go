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
// proto.message: google.cloud.datastream.v1.Stream

package mockdatastream

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datastream/v1"
)

func (s *DatastreamV1) GetStream(ctx context.Context, req *pb.GetStreamRequest) (*pb.Stream, error) {
	name, err := s.parseStreamName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Stream{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "stream %q not found", req.Name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DatastreamV1) CreateStream(ctx context.Context, req *pb.CreateStreamRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/streams/%s", req.GetParent(), req.GetStreamId())
	name, err := s.parseStreamName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetStream()).(*pb.Stream)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.State == pb.Stream_STATE_UNSPECIFIED {
		obj.State = pb.Stream_NOT_STARTED
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	meta := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
	}
	return s.operations.StartLRO(ctx, prefix, meta, func() (proto.Message, error) {
		meta.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DatastreamV1) UpdateStream(ctx context.Context, req *pb.UpdateStreamRequest) (*longrunning.Operation, error) {
	name, err := s.parseStreamName(req.GetStream().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Stream{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	proto.Merge(obj, req.GetStream())
	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	meta := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
	}

	return s.operations.StartLRO(ctx, prefix, meta, func() (proto.Message, error) {
		meta.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DatastreamV1) DeleteStream(ctx context.Context, req *pb.DeleteStreamRequest) (*longrunning.Operation, error) {
	name, err := s.parseStreamName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	deleted := &pb.Stream{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	meta := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
	}
	return s.operations.StartLRO(ctx, prefix, meta, func() (proto.Message, error) {
		meta.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type streamName struct {
	Project  *projects.ProjectData
	Location string
	StreamID string
}

func (n *streamName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/streams/%s", n.Project.Number, n.Location, n.StreamID)
}

func (s *DatastreamV1) parseStreamName(name string) (*streamName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "streams" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		streamName := &streamName{
			Project:  project,
			Location: tokens[3],
			StreamID: tokens[5],
		}

		return streamName, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
