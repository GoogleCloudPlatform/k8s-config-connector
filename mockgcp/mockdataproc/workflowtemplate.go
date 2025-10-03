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
// proto.service: google.cloud.dataproc.v1.WorkflowTemplateService
// proto.message: google.cloud.dataproc.v1.WorkflowTemplate

package mockdataproc

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

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *workflowTemplateServer) GetWorkflowTemplate(ctx context.Context, req *pb.GetWorkflowTemplateRequest) (*pb.WorkflowTemplate, error) {
	name, err := s.parseWorkflowTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkflowTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "WorkflowTemplate %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *workflowTemplateServer) CreateWorkflowTemplate(ctx context.Context, req *pb.CreateWorkflowTemplateRequest) (*pb.WorkflowTemplate, error) {
	reqName := req.Parent + "/workflowTemplates/" + req.Template.Id
	name, err := s.parseWorkflowTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetTemplate()).(*pb.WorkflowTemplate)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	s.populateDefaultsForWorkflowTemplate(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *workflowTemplateServer) populateDefaultsForWorkflowTemplate(obj *pb.WorkflowTemplate) {
	if obj.Placement == nil {
		obj.Placement = &pb.WorkflowTemplatePlacement{}
	}
	if obj.Placement.GetClusterSelector() == nil && obj.Placement.GetManagedCluster() == nil {
		obj.Placement = &pb.WorkflowTemplatePlacement{
			Placement: &pb.WorkflowTemplatePlacement_ManagedCluster{
				ManagedCluster: &pb.ManagedCluster{},
			},
		}
	}
}

func (s *workflowTemplateServer) UpdateWorkflowTemplate(ctx context.Context, req *pb.UpdateWorkflowTemplateRequest) (*pb.WorkflowTemplate, error) {
	reqName := req.GetTemplate().Name
	name, err := s.parseWorkflowTemplateName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.WorkflowTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	s.populateDefaultsForWorkflowTemplate(obj)

	// TODO: Can we use a fieldmask here?

	obj.Jobs = req.GetTemplate().Jobs
	obj.Parameters = req.GetTemplate().Parameters
	obj.DagTimeout = req.GetTemplate().DagTimeout
	obj.Version = req.GetTemplate().Version + 1
	obj.UpdateTime = timestamppb.New(time.Now())
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *workflowTemplateServer) DeleteWorkflowTemplate(ctx context.Context, req *pb.DeleteWorkflowTemplateRequest) (*emptypb.Empty, error) {
	name, err := s.parseWorkflowTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.WorkflowTemplate{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type workflowTemplateName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *workflowTemplateName) String() string {
	return fmt.Sprintf("projects/%s/regions/%s/workflowTemplates/%s", n.Project.ID, n.Region, n.Name)
}

// parseWorkflowTemplateName parses a string into an WorkflowTemplateName.
// The expected form is `projects/*/regions/*/workflowTemplates/*`.
func (s *MockService) parseWorkflowTemplateName(name string) (*workflowTemplateName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "workflowTemplates" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &workflowTemplateName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
