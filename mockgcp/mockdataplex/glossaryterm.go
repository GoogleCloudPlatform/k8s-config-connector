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
// proto.service: google.cloud.dataplex.v1.BusinessGlossaryService
// proto.message: google.cloud.dataplex.v1.GlossaryTerm

package mockdataplex

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

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *BusinessGlossaryService) GetGlossaryTerm(ctx context.Context, req *pb.GetGlossaryTermRequest) (*pb.GlossaryTerm, error) {
	name, err := s.parseGlossaryTermName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GlossaryTerm{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Term: %s not found.", name.FormattedString())
		}
		return nil, err
	}

	return obj, nil
}

func (s *BusinessGlossaryService) CreateGlossaryTerm(ctx context.Context, req *pb.CreateGlossaryTermRequest) (*pb.GlossaryTerm, error) {
	reqName := req.Parent + "/terms/" + req.TermId
	name, err := s.parseGlossaryTermName(reqName)
	if err != nil {
		return nil, err
	}

	if req.Term == nil {
		return nil, status.Errorf(codes.InvalidArgument, "term is required")
	}

	fqn := name.String()

	now := time.Now()
	obj := proto.Clone(req.Term).(*pb.GlossaryTerm)
	obj.Name = fqn
	obj.Parent = req.Parent
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = fmt.Sprintf("%d", now.UnixNano())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BusinessGlossaryService) UpdateGlossaryTerm(ctx context.Context, req *pb.UpdateGlossaryTermRequest) (*pb.GlossaryTerm, error) {
	if req.GetTerm() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "term is required")
	}

	name, err := s.parseGlossaryTermName(req.GetTerm().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.GlossaryTerm{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Term: %s not found.", name.FormattedString())
		}
		return nil, err
	}
	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	updateMask := req.GetUpdateMask()
	if updateMask == nil || len(updateMask.GetPaths()) == 0 {
		obj.Description = req.GetTerm().GetDescription()
		obj.DisplayName = req.GetTerm().GetDisplayName()
		obj.Labels = req.GetTerm().GetLabels()
	} else {
		for _, path := range updateMask.GetPaths() {
			switch path {
			case "description":
				obj.Description = req.GetTerm().GetDescription()
			case "display_name", "displayName":
				obj.DisplayName = req.GetTerm().GetDisplayName()
			case "labels":
				obj.Labels = req.GetTerm().GetLabels()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "mock does not implement update of %q", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BusinessGlossaryService) DeleteGlossaryTerm(ctx context.Context, req *pb.DeleteGlossaryTermRequest) (*emptypb.Empty, error) {
	name, err := s.parseGlossaryTermName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.GlossaryTerm{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Term: %s not found.", name.FormattedString())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type glossaryTermName struct {
	Project    *projects.ProjectData
	Location   string
	GlossaryID string
	TermID     string
}

func (n *glossaryTermName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/terms/%s", n.Project.ID, n.Location, n.GlossaryID, n.TermID)
}

func (n *glossaryTermName) FormattedString() string {
	return fmt.Sprintf("projects/%d/locations/%s/glossaries/%s/terms/%s", n.Project.Number, n.Location, n.GlossaryID, n.TermID)
}

func (s *MockService) parseGlossaryTermName(name string) (*glossaryTermName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "glossaries" && tokens[6] == "terms" {
		if tokens[7] == "" {
			return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
		}
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &glossaryTermName{
			Project:    project,
			Location:   tokens[3],
			GlossaryID: tokens[5],
			TermID:     tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
