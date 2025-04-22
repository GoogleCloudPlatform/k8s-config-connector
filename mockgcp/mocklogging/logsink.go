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
// proto.service: google.logging.v2.ConfigServiceV2
// proto.message: google.logging.v2.LogSink

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

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *configServiceV2) GetSink(ctx context.Context, req *pb.GetSinkRequest) (*pb.LogSink, error) {
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

func (s *configServiceV2) CreateSink(ctx context.Context, req *pb.CreateSinkRequest) (*pb.LogSink, error) {
	reqName := fmt.Sprintf("%s/sinks/%s", req.Parent, req.GetSink().GetName())
	name, err := s.parseLogSinkName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetSink()).(*pb.LogSink)
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if name.Parent.Project != nil {
		obj.WriterIdentity = fmt.Sprintf("serviceAccount:service-%d@gcp-sa-logging.iam.gserviceaccount.com", name.Parent.Project.Number)
	}

	if name.Parent.Folder != "" {
		obj.WriterIdentity = fmt.Sprintf("serviceAccount:service-folder-%s@gcp-sa-logging.iam.gserviceaccount.com", name.Parent.Folder)
	}

	if name.Parent.Organization != "" {
		obj.WriterIdentity = fmt.Sprintf("serviceAccount:service-org-%s@gcp-sa-logging.iam.gserviceaccount.com", name.Parent.Organization)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *configServiceV2) UpdateSink(ctx context.Context, req *pb.UpdateSinkRequest) (*pb.LogSink, error) {
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

func (s *configServiceV2) DeleteSink(ctx context.Context, req *pb.DeleteSinkRequest) (*empty.Empty, error) {
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

type FolderOrgOrProject struct {
	Folder       string
	Organization string
	Project      *projects.ProjectData
}

func (s *configServiceV2) PopFolderOrgOrProject(tokens []string) (*FolderOrgOrProject, []string, error) {
	if len(tokens) >= 2 && tokens[0] == "projects" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, nil, err
		}

		name := &FolderOrgOrProject{
			Project: project,
		}

		return name, tokens[2:], nil
	}

	if len(tokens) >= 2 && tokens[0] == "folders" {
		name := &FolderOrgOrProject{
			Folder: tokens[1],
		}

		return name, tokens[2:], nil
	}

	if len(tokens) >= 2 && tokens[0] == "organizations" {
		name := &FolderOrgOrProject{
			Organization: tokens[1],
		}

		return name, tokens[2:], nil
	}

	return nil, nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", strings.Join(tokens, "/"))
}

type logSinkName struct {
	Parent   FolderOrgOrProject
	SinkName string
}

func (n *FolderOrgOrProject) String() string {
	if n.Organization != "" {
		return fmt.Sprintf("organizations/%s", n.Organization)
	}
	if n.Folder != "" {
		return fmt.Sprintf("folders/%s", n.Organization)
	}
	return fmt.Sprintf("projects/%s", n.Project.ID)
}

func (n *logSinkName) String() string {
	return fmt.Sprintf("%s/sinks/%s", n.Parent.String(), n.SinkName)
}

// parseLogSinkName parses a string into a logSinkName.
// The expected form is `projects/*/sinks/*`
func (s *configServiceV2) parseLogSinkName(name string) (*logSinkName, error) {
	tokens := strings.Split(name, "/")

	parent, remainder, err := s.PopFolderOrgOrProject(tokens)
	if err != nil {
		return nil, err
	}
	if len(remainder) == 2 && remainder[0] == "sinks" {
		name := &logSinkName{
			Parent:   *parent,
			SinkName: remainder[1],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
