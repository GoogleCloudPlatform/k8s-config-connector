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
// proto.message: google.cloud.dataplex.v1.Glossary

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunning "google.golang.org/genproto/googleapis/longrunning"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

type BusinessGlossaryService struct {
	*MockService
	pb.UnimplementedBusinessGlossaryServiceServer
}

func (s *BusinessGlossaryService) GetGlossary(ctx context.Context, req *pb.GetGlossaryRequest) (*pb.Glossary, error) {
	name, err := s.parseGlossaryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Glossary{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	s.populateDefaultsForGlossary(obj)

	return obj, nil
}

func (s *BusinessGlossaryService) CreateGlossary(ctx context.Context, req *pb.CreateGlossaryRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/glossaries/" + req.GlossaryId
	name, err := s.parseGlossaryName(reqName)
	if err != nil {
		return nil, err
	}

	if req.Glossary == nil {
		return nil, status.Errorf(codes.InvalidArgument, "glossary is required")
	}

	fqn := name.String()

	now := time.Now()
	obj := proto.Clone(req.Glossary).(*pb.Glossary)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = fmt.Sprintf("%d", now.UnixNano())
	obj.Etag = "abcdef0123A="

	s.populateDefaultsForGlossary(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		Target:     name.String(),
		Verb:       "create",
		CreateTime: timestamppb.New(now),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.RequestedCancellation = false
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *BusinessGlossaryService) UpdateGlossary(ctx context.Context, req *pb.UpdateGlossaryRequest) (*longrunning.Operation, error) {
	if req.GetGlossary() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "glossary is required")
	}

	name, err := s.parseGlossaryName(req.GetGlossary().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Glossary{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}
	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	updateMask := req.GetUpdateMask()
	if updateMask == nil || len(updateMask.GetPaths()) == 0 {
		obj.Description = req.GetGlossary().GetDescription()
		obj.DisplayName = req.GetGlossary().GetDisplayName()
		obj.Labels = req.GetGlossary().GetLabels()
	} else {
		for _, path := range updateMask.GetPaths() {
			switch path {
			case "description":
				obj.Description = req.GetGlossary().GetDescription()
			case "display_name", "displayName":
				obj.DisplayName = req.GetGlossary().GetDisplayName()
			case "labels":
				obj.Labels = req.GetGlossary().GetLabels()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "mock does not implement update of %q", path)
			}
		}
	}

	s.populateDefaultsForGlossary(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.RequestedCancellation = false
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *BusinessGlossaryService) DeleteGlossary(ctx context.Context, req *pb.DeleteGlossaryRequest) (*longrunning.Operation, error) {
	name, err := s.parseGlossaryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Glossary{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	now := time.Now()
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		Target:     name.String(),
		Verb:       "delete",
		CreateTime: timestamppb.New(now),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.RequestedCancellation = false
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *BusinessGlossaryService) populateDefaultsForGlossary(obj *pb.Glossary) {
	// TODO: TermCount and CategoryCount default to 0 and are omitted by proto3 JSON serialization.
}

type glossaryName struct {
	Project    *projects.ProjectData
	Location   string
	GlossaryID string
}

func (n *glossaryName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", n.Project.ID, n.Location, n.GlossaryID)
}

func (s *MockService) parseGlossaryName(name string) (*glossaryName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "glossaries" {
		if tokens[5] == "" {
			return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
		}
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &glossaryName{
			Project:    project,
			Location:   tokens[3],
			GlossaryID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
