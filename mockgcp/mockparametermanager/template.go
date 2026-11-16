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

package mockparametermanager

import (
	"context"
	"strings"

	pb "cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// Creates a new [Template][google.cloud.parametermanager.v1.Template].
func (s *ParameterManagerV1) CreateTemplate(ctx context.Context, req *pb.CreateTemplateRequest) (*pb.Template, error) {
	templateID := req.TemplateId
	if templateID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "TemplateId is required")
	}

	reqTemplateName := req.Parent + "/templates/" + templateID

	name, err := s.parseTemplateName(reqTemplateName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.Template)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = obj.CreateTime

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Post response does not include labels
	response := proto.CloneOf(obj)
	response.Labels = nil
	return response, nil
}

// Gets metadata for a given [Template][google.cloud.parametermanager.v1.Template].
func (s *ParameterManagerV1) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.Template, error) {
	name, err := s.parseTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	var template pb.Template
	fqn := name.String()
	if err := s.storage.Get(ctx, fqn, &template); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &template, nil
}

// Update metadata for a given [Template][google.cloud.parametermanager.v1.Template].
func (s *ParameterManagerV1) UpdateTemplate(ctx context.Context, req *pb.UpdateTemplateRequest) (*pb.Template, error) {
	name, err := s.parseTemplateName(req.Template.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.Template{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.CloneOf(existing)
	updated.Name = name.String()
	updated.UpdateTime = timestamppb.Now()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}
	for _, path := range paths {
		switch path {
		case "labels":
			updated.Labels = req.Template.GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	// Patch response does not include labels
	response := proto.CloneOf(updated)
	response.Labels = nil
	return response, nil
}

// Deletes a [Template][google.cloud.parametermanager.v1.Template].
func (s *ParameterManagerV1) DeleteTemplate(ctx context.Context, req *pb.DeleteTemplateRequest) (*emptypb.Empty, error) {
	name, err := s.parseTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Template{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// Lists Templates in a given project and location.
func (s *ParameterManagerV1) ListTemplates(ctx context.Context, req *pb.ListTemplatesRequest) (*pb.ListTemplatesResponse, error) {
	prefix := req.Parent + "/templates/"
	response := &pb.ListTemplatesResponse{}
	templateKind := (&pb.Template{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, templateKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		template := obj.(*pb.Template)
		response.Templates = append(response.Templates, template)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

type templateName struct {
	Project      *projects.ProjectData
	Location     string
	TemplateName string
}

func (n *templateName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/templates/" + n.TemplateName
}

// parseTemplateName parses a string into a templateName.
// The expected form is projects/<projectID>/locations/<location>/templates/<templateName>
func (s *MockService) parseTemplateName(name string) (*templateName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "templates" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &templateName{
			Project:      project,
			Location:     tokens[3],
			TemplateName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"name %q is not valid. expected format is projects/<projectID>/locations/<location>/templates/<templateName>",
			name,
		)
	}
}
