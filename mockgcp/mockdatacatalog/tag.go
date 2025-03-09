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
// proto.message: google.cloud.datacatalog.v1.Tag

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

func (s *DataCatalogV1) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.Tag, error) {
	name, err := s.parseTagName(req.Parent)
	if err != nil {
		return nil, err
	}
	reqName := name.String() + "/"

	fqn := reqName + fmt.Sprintf("generated%d", s.IDGenerator.Next())

	obj := proto.Clone(req.Tag).(*pb.Tag)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DataCatalogV1) UpdateTag(ctx context.Context, req *pb.UpdateTagRequest) (*pb.Tag, error) {
	reqName := req.GetTag().GetName()
	name, err := s.parseTagName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Tag{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Nested FieldMasks.
	if req.GetUpdateMask() != nil && len(req.GetUpdateMask().Paths) > 0 {
		return nil, status.Errorf(codes.Unimplemented, "UpdateTag with updateMask is not yet implemented")
	}

	// TODO: the merge operation of 'req' onto 'obj' is incorrect
	obj = proto.Clone(req.GetTag()).(*pb.Tag)
	obj.Name = fqn

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *DataCatalogV1) DeleteTag(ctx context.Context, req *pb.DeleteTagRequest) (*emptypb.Empty, error) {
	name, err := s.parseTagName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deleted := &pb.Tag{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type entryName struct {
	Project     *projects.ProjectData
	Location    string
	EntryGroup  string
	Entry       string
	TagTemplate string
	Tag         string
}

// projects/{project}/locations/{location}/entryGroups/{entry_group}/entries/{entry}
// projects/{project}/locations/{location}/tagTemplates/{tag_template}
func (n *entryName) String() string {
	var b strings.Builder
	b.WriteString("projects/" + n.Project.ID + "/locations/" + n.Location + "/entryGroups/" + n.EntryGroup)
	if n.Entry != "" {
		b.WriteString("/entries/")
		b.WriteString(n.Entry)
	}
	if n.TagTemplate != "" {
		b.WriteString("/tagTemplates/")
		b.WriteString(n.TagTemplate)
	}
	if n.Tag != "" {
		b.WriteString("/tags/")
		b.WriteString(n.Tag)
	}
	return b.String()
}

// parseTagName parses a string into an entryName.
func (s *DataCatalogV1) parseTagName(name string) (*entryName, error) {
	tokens := strings.Split(name, "/")
	valid := false
	if len(tokens) == 10 {
		if tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "entryGroups" && tokens[6] == "entries" && tokens[8] == "tags" {
			valid = true
		}
	}
	if len(tokens) == 8 {
		if tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "entryGroups" && tokens[6] == "tags" {
			valid = true
		}
	}
	if !valid {
		return nil, fmt.Errorf("unable to parse %q as a valid entryName", name)
	}

	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	namePart := &entryName{
		Project:    project,
		Location:   tokens[3],
		EntryGroup: tokens[5],
	}
	if len(tokens) > 6 {
		namePart.Entry = tokens[7]
		if len(tokens) > 8 {
			namePart.Tag = tokens[9]
		}
	}
	return namePart, nil
}

