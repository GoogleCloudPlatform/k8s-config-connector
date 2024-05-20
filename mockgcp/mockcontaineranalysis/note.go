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

package mockcontaineranalysis

import (
	"context"
	"strings"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgrafeas/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NoteName struct {
	Project  *projects.ProjectData
	NoteName string
}

type ContainerAnalysisV1 struct {
	*MockService
	pb.UnimplementedGrafeasServer
}

func (s *ContainerAnalysisV1) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.Note, error) {
	name, err := s.parseNoteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Note{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContainerAnalysisV1) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.Note, error) {
	reqName := req.Parent + "/notes/" + req.NoteId
	name, err := s.parseNoteName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Note).(*pb.Note)
	obj.Name = fqn
	obj.Kind = pb.NoteKind_ATTESTATION
	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContainerAnalysisV1) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*pb.Note, error) {
	reqName := req.GetName()

	name, err := s.parseNoteName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Note{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "short_description":
			obj.ShortDescription = req.Note.GetShortDescription()
		case "long_description":
			obj.LongDescription = req.Note.GetLongDescription()
		case "related_url":
			obj.RelatedUrl = req.Note.GetRelatedUrl()
		case "attestation.hint.human_readable_name":
			obj.Type = req.Note.GetType()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContainerAnalysisV1) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*empty.Empty, error) {
	name, err := s.parseNoteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Note{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (n *NoteName) String() string {
	return "projects/" + n.Project.ID + "/notes/" + n.NoteName
}

// parseNoteName parses a string into a NoteName.
// The expected form is projects/<projectID>/notes/<NoteName>
func (s *MockService) parseNoteName(name string) (*NoteName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "notes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &NoteName{
			Project:  project,
			NoteName: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
