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
// apiVersion: logging.cnrm.cloud.google.com/v1beta1
// kind: LoggingLogSink
// service: google.logging.v2.ConfigServiceV2
// resource: LogSink

package mocklogging

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/logging/v2"
)

func (s *configService) GetSink(ctx context.Context, req *pb.GetSinkRequest) (*pb.LogSink, error) {
	name, err := s.parseLogSinkName(req.SinkName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.LogSink{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Sink %s does not exist", name.SinkName)
		}
		return nil, err
	}
	return obj, nil
}

func (s *configService) CreateSink(ctx context.Context, req *pb.CreateSinkRequest) (*pb.LogSink, error) {
	reqName := fmt.Sprintf("%s/sinks/%s", req.Parent, req.GetSink().GetName())
	name, err := s.parseLogSinkName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetSink()).(*pb.LogSink)
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	obj.WriterIdentity = fmt.Sprintf("serviceAccount:service-%d@gcp-sa-logging.iam.gserviceaccount.com", name.Project.Number)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *configService) UpdateSink(ctx context.Context, req *pb.UpdateSinkRequest) (*pb.LogSink, error) {
	reqName := req.SinkName
	name, err := s.parseLogSinkName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.LogSink{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.LogSink)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// Optional, but we require it in our mock.
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required by mock")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			updated.Description = req.GetSink().GetDescription()
		case "filter":
			updated.Filter = req.GetSink().GetFilter()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	updated.UpdateTime = timestamppb.New(time.Now())
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *configService) DeleteSink(ctx context.Context, req *pb.DeleteSinkRequest) (*empty.Empty, error) {
	name, err := s.parseLogSinkName(req.SinkName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.LogSink{}

	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

type logSinkName struct {
	Project  *projects.ProjectData
	SinkName string
}

func (n *logSinkName) String() string {
	return fmt.Sprintf("projects/%s/sinks/%s", n.Project.ID, n.SinkName)

}

// parseLogSinkName parses a string into a logSinkName.
// The expected form is `projects/*/sinks/*`
func (s *configService) parseLogSinkName(name string) (*logSinkName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "sinks" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &logSinkName{
			Project:  project,
			SinkName: tokens[3],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
