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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datastream/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *DatastreamV1) ListStreams(ctx context.Context, req *pb.ListStreamsRequest) (*pb.ListStreamsResponse, error) {
	parent := req.GetParent() // projects/*/locations/*

	res := &pb.ListStreamsResponse{}
	err := s.storage.List(ctx, (&pb.Stream{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: parent + "/"}, func(obj proto.Message) error {
		res.Streams = append(res.Streams, obj.(*pb.Stream))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DatastreamV1) GetStream(ctx context.Context, req *pb.GetStreamRequest) (*pb.Stream, error) {
	name, err := s.parseStreamName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Stream{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DatastreamV1) CreateStream(ctx context.Context, req *pb.CreateStreamRequest) (*longrunningpb.Operation, error) {
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
	obj.State = pb.Stream_RUNNING // Default to running? Or CREATED?

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

func (s *DatastreamV1) UpdateStream(ctx context.Context, req *pb.UpdateStreamRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetStream().GetName()
	name, err := s.parseStreamName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Stream{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.UpdateTime = timestamppb.New(time.Now())

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, req.GetStream(), paths); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating stream: %v", err)
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

func (s *DatastreamV1) RunStream(ctx context.Context, req *pb.RunStreamRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseStreamName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Stream{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.Stream_RUNNING
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "run",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DatastreamV1) DeleteStream(ctx context.Context, req *pb.DeleteStreamRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseStreamName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Stream{}
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

func (s *DatastreamV1) ListStreamObjects(ctx context.Context, req *pb.ListStreamObjectsRequest) (*pb.ListStreamObjectsResponse, error) {
	parent := req.GetParent() // projects/*/locations/*/streams/*

	res := &pb.ListStreamObjectsResponse{}
	err := s.storage.List(ctx, (&pb.StreamObject{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: parent + "/"}, func(obj proto.Message) error {
		res.StreamObjects = append(res.StreamObjects, obj.(*pb.StreamObject))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DatastreamV1) GetStreamObject(ctx context.Context, req *pb.GetStreamObjectRequest) (*pb.StreamObject, error) {
	name, err := s.parseStreamObjectName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.StreamObject{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

type streamName struct {
	Project  *projects.ProjectData
	Location string
	StreamID string
}

func (n *streamName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/streams/" + n.StreamID
}

func (s *MockService) parseStreamName(name string) (*streamName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "streams" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &streamName{
			Project:  project,
			Location: tokens[3],
			StreamID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type streamObjectName struct {
	Project  *projects.ProjectData
	Location string
	StreamID string
	ObjectID string
}

func (n *streamObjectName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/streams/" + n.StreamID + "/objects/" + n.ObjectID
}

// parseStreamObjectName parses a string into a streamObjectName.
// The expected form is `projects/*/locations/*/streams/*/objects/*`.
func (s *MockService) parseStreamObjectName(name string) (*streamObjectName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "streams" && tokens[6] == "objects" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &streamObjectName{
			Project:  project,
			Location: tokens[3],
			StreamID: tokens[5],
			ObjectID: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
