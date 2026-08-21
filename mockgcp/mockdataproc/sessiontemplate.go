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

// +tool:mockgcp-support
// proto.service: google.cloud.dataproc.v1.SessionTemplateController
// proto.message: google.cloud.dataproc.v1.SessionTemplate

package mockdataproc

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type sessionTemplateControllerServer struct {
	*MockService
	pb.UnimplementedSessionTemplateControllerServer
}

func (s *sessionTemplateControllerServer) GetSessionTemplate(ctx context.Context, req *pb.GetSessionTemplateRequest) (*pb.SessionTemplate, error) {
	name, err := s.parseSessionTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SessionTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *sessionTemplateControllerServer) CreateSessionTemplate(ctx context.Context, req *pb.CreateSessionTemplateRequest) (*pb.SessionTemplate, error) {
	name, err := s.parseSessionTemplateName(req.SessionTemplate.GetName())
	if err != nil {
		// Fallback to name parsed from parent + ID if template Name is not set,
		// but since KCC always sets name, we can also use req.Parent + "/sessionTemplates/" + ...
		// Let's parse req.SessionTemplate.GetName() or construct from parent
		if req.SessionTemplate.GetName() == "" {
			return nil, status.Errorf(codes.InvalidArgument, "template name is required")
		}
		name, err = s.parseSessionTemplateName(req.SessionTemplate.GetName())
		if err != nil {
			return nil, err
		}
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.SessionTemplate).(*pb.SessionTemplate)
	obj.Name = fqn
	obj.Uuid = "00000000-0000-0000-0000-000000000001"
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Creator = "test-user@google.com"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *sessionTemplateControllerServer) UpdateSessionTemplate(ctx context.Context, req *pb.UpdateSessionTemplateRequest) (*pb.SessionTemplate, error) {
	name, err := s.parseSessionTemplateName(req.SessionTemplate.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.SessionTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update fields
	obj.Description = req.SessionTemplate.GetDescription()
	obj.SessionConfig = req.SessionTemplate.GetSessionConfig()
	obj.Labels = req.SessionTemplate.GetLabels()
	obj.RuntimeConfig = req.SessionTemplate.GetRuntimeConfig()
	obj.EnvironmentConfig = req.SessionTemplate.GetEnvironmentConfig()

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *sessionTemplateControllerServer) DeleteSessionTemplate(ctx context.Context, req *pb.DeleteSessionTemplateRequest) (*emptypb.Empty, error) {
	name, err := s.parseSessionTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.SessionTemplate{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *sessionTemplateControllerServer) ListSessionTemplates(ctx context.Context, req *pb.ListSessionTemplatesRequest) (*pb.ListSessionTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSessionTemplates not implemented")
}

type sessionTemplateName struct {
	Project  string
	Location string
	Template string
}

func (n *sessionTemplateName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/sessionTemplates/%s", n.Project, n.Location, n.Template)
}

func (s *MockService) parseSessionTemplateName(name string) (*sessionTemplateName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "sessionTemplates" {
		return &sessionTemplateName{
			Project:  tokens[1],
			Location: tokens[3],
			Template: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
