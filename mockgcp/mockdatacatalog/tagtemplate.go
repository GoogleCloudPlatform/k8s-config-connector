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
// proto.service: google.cloud.datacatalog.v1.DataCatalog
// proto.message: google.cloud.datacatalog.v1.TagTemplate

package mockdatacatalog

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datacatalog/v1"
)

func (s *DataCatalogV1) GetTagTemplate(ctx context.Context, req *pb.GetTagTemplateRequest) (*pb.TagTemplate, error) {
	name, err := s.parseTagTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TagTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "TagTemplate %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataCatalogV1) CreateTagTemplate(ctx context.Context, req *pb.CreateTagTemplateRequest) (*pb.TagTemplate, error) {
	reqName := req.Parent + "/tagTemplates/" + req.TagTemplateId
	name, err := s.parseTagTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.TagTemplate).(*pb.TagTemplate)
	obj.Name = fqn
	s.populateDefaultsForTagTemplate(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DataCatalogV1) populateDefaultsForTagTemplate(obj *pb.TagTemplate) {

}

func (s *DataCatalogV1) UpdateTagTemplate(ctx context.Context, req *pb.UpdateTagTemplateRequest) (*pb.TagTemplate, error) {
	reqName := req.GetTagTemplate().GetName()
	name, err := s.parseTagTemplateName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.TagTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	s.populateDefaultsForTagTemplate(obj)

	// TODO: Can we use a fieldmask here?
	obj.DisplayName = req.GetTagTemplate().GetDisplayName()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *DataCatalogV1) DeleteTagTemplate(ctx context.Context, req *pb.DeleteTagTemplateRequest) (*emptypb.Empty, error) {
	name, err := s.parseTagTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TagTemplate{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type tagTemplateName struct {
	Project   *projects.ProjectData
	Location  string
	Id        string
	FieldName string
}

func (n *tagTemplateName) String() string {
	if n.FieldName == "" {
		return fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", n.Project.ID, n.Location, n.Id)
	} else {
		return fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s/fields/%s", n.Project.ID, n.Location, n.Id, n.FieldName)
	}
}

// parseTagTemplateName parses a string into a tagTemplateName.
// The expected form is `projects/*/locations/*/tagTemplates/*`.
func (s *MockService) parseTagTemplateName(name string) (*tagTemplateName, error) {
	tokens := strings.Split(name, "/")
	projectTokenIdx := -1
	locationTokenIdx := -1
	idTokenIdx := -1
	//fieldNameTokenIdx := -1

	if len(tokens) >= 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "tagTemplates" {
		projectTokenIdx = 1
		locationTokenIdx = 3
		idTokenIdx = 5
	}

	if len(tokens) >= 8 && tokens[6] == "fields" {
		//fieldNameTokenIdx = 7
		return nil, status.Errorf(codes.Unimplemented, "mock for projects.locations.tagTemplates.fields not implemented yet")
	}

	if projectTokenIdx == -1 {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}

	project, err := s.Projects.GetProjectByID(tokens[projectTokenIdx])
	if err != nil {
		return nil, err
	}

	name := &tagTemplateName{
		Project:  project,
		Location: tokens[locationTokenIdx],
		Id:       tokens[idTokenIdx],
	}

	return name, nil
}

