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

package mockcompute

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/protobuf/proto"
)

type NodeTemplatesV1 struct {
	*MockService
	pb.UnimplementedNodeTemplatesServer
}

func (s *NodeTemplatesV1) Get(ctx context.Context, req *pb.GetNodeTemplateRequest) (*pb.NodeTemplate, error) {
	name, err := s.newNodeTemplateName(req.GetProject(), req.GetRegion(), req.GetNodeTemplate())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NodeTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *NodeTemplatesV1) Insert(ctx context.Context, req *pb.InsertNodeTemplateRequest) (*pb.Operation, error) {
	name, err := s.newNodeTemplateName(req.GetProject(), req.GetRegion(), req.GetNodeTemplateResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetNodeTemplateResource()).(*pb.NodeTemplate)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#nodetemplate")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *NodeTemplatesV1) Delete(ctx context.Context, req *pb.DeleteNodeTemplateRequest) (*pb.Operation, error) {
	name, err := s.newNodeTemplateName(req.GetProject(), req.GetRegion(), req.GetNodeTemplate())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NodeTemplate{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type nodeTemplateName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *nodeTemplateName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/nodeTemplates/" + n.Name
}

// newNodeTemplateName builds a normalized nodeTemplateName from the constituent parts.
// The expected form is `projects/{project}/regions/{region}/nodeTemplates/{nodeTemplate}`.
func (s *MockService) newNodeTemplateName(project string, region string, name string) (*nodeTemplateName, error) {
	projectObj, err := s.Projects.GetProjectByID(project)
	if err != nil {
		return nil, err
	}

	return &nodeTemplateName{
		Project: projectObj,
		Region:  region,
		Name:    name,
	}, nil
}
