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
// proto.message: google.logging.v2.LogView

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

func (s *configServiceV2) GetView(ctx context.Context, req *pb.GetViewRequest) (*pb.LogView, error) {
	name, err := s.parseLogViewName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.LogView{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "View %s does not exist", name.ViewName)
		}
		return nil, err
	}
	return obj, nil
}

func (s *configServiceV2) CreateView(ctx context.Context, req *pb.CreateViewRequest) (*pb.LogView, error) {
	reqName := fmt.Sprintf("%s/views/%s", req.GetParent(), req.GetViewId())
	name, err := s.parseLogViewName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetView()).(*pb.LogView)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *configServiceV2) UpdateView(ctx context.Context, req *pb.UpdateViewRequest) (*pb.LogView, error) {
	reqName := req.Name
	name, err := s.parseLogViewName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.LogView{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	now := time.Now()
	updated := proto.Clone(existing).(*pb.LogView)

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required by mock")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			updated.Description = req.GetView().GetDescription()
		case "filter":
			updated.Filter = req.GetView().GetFilter()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	updated.UpdateTime = timestamppb.New(now)
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *configServiceV2) DeleteView(ctx context.Context, req *pb.DeleteViewRequest) (*empty.Empty, error) {
	name, err := s.parseLogViewName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.LogView{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

type logViewName struct {
	project      *projects.ProjectData
	organization string
	location     string
	bucketName   string
	ViewName     string
}

func (n *logViewName) String() string {
	if n.organization != "" {
		return fmt.Sprintf("organizations/%s/locations/%s/buckets/%s/views/%s", n.organization, n.location, n.bucketName, n.ViewName)

	}
	return fmt.Sprintf("projects/%s/locations/%s/buckets/%s/views/%s", n.project.ID, n.location, n.bucketName, n.ViewName)
}

// parseLogViewName parses a string into a logViewName.
func (s *MockService) parseLogViewName(name string) (*logViewName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "views" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &logViewName{
			project:    project,
			location:   tokens[3],
			bucketName: tokens[5],
			ViewName:   tokens[7],
		}
		return name, nil
	}

	if len(tokens) == 8 && tokens[0] == "organizations" && tokens[2] == "locations" && tokens[3] == "global" && tokens[4] == "buckets" && tokens[6] == "views" {
		name := &logViewName{
			organization: tokens[1],
			location:     tokens[3],
			bucketName:   tokens[5],
			ViewName:     tokens[7],
		}
		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
