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

type InstanceTemplatesV1 struct {
	*MockService
	pb.UnimplementedInstanceTemplatesServer
}

func (s *InstanceTemplatesV1) Get(ctx context.Context, req *pb.GetInstanceTemplateRequest) (*pb.InstanceTemplate, error) {
	name, err := s.parseInstanceTemplateName(req.GetProject(), req.GetInstanceTemplate())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstanceTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *InstanceTemplatesV1) Insert(ctx context.Context, req *pb.InsertInstanceTemplateRequest) (*pb.Operation, error) {
	name, err := s.parseInstanceTemplateName(req.GetProject(), req.GetInstanceTemplateResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetInstanceTemplateResource()).(*pb.InstanceTemplate)
	obj.SelfLink = PtrTo("https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Kind = PtrTo("compute#instanceTemplate")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *InstanceTemplatesV1) Delete(ctx context.Context, req *pb.DeleteInstanceTemplateRequest) (*pb.Operation, error) {
	name, err := s.parseInstanceTemplateName(req.GetProject(), req.GetInstanceTemplate())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.InstanceTemplate{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type instanceTemplateName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *instanceTemplateName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/instanceTemplates/" + n.Name
}

func (s *MockService) parseInstanceTemplateName(projectName, name string) (*instanceTemplateName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &instanceTemplateName{
		Project: project,
		Name:    name,
	}, nil
}
