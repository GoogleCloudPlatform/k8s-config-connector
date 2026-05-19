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
// proto.service: google.cloud.modelarmor.v1.ModelArmor
// proto.message: google.cloud.modelarmor.v1.Template

package mockmodelarmor

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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/modelarmor/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type ModelArmorV1 struct {
	*MockService
	pb.UnimplementedModelArmorServer
}

func (s *ModelArmorV1) ListTemplates(ctx context.Context, req *pb.ListTemplatesRequest) (*pb.ListTemplatesResponse, error) {
	// parent is like projects/*/locations/*
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}

	res := &pb.ListTemplatesResponse{}
	kind := (&pb.Template{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: req.Parent}, func(obj proto.Message) error {
		res.Templates = append(res.Templates, obj.(*pb.Template))
		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ModelArmorV1) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.Template, error) {
	name, err := s.parseTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Template{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ModelArmorV1) CreateTemplate(ctx context.Context, req *pb.CreateTemplateRequest) (*pb.Template, error) {
	reqName := req.Parent + "/templates/" + req.TemplateId
	name, err := s.parseTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.Template)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	if obj.TemplateMetadata == nil {
		obj.TemplateMetadata = &pb.Template_TemplateMetadata{}
	}
	s.populateDefaultsForTemplate(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ModelArmorV1) populateDefaultsForTemplate(obj *pb.Template) {

}

func (s *ModelArmorV1) UpdateTemplate(ctx context.Context, req *pb.UpdateTemplateRequest) (*pb.Template, error) {
	name, err := s.parseTemplateName(req.GetTemplate().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Template{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Use update_mask
	proto.Merge(obj, req.Template)
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ModelArmorV1) DeleteTemplate(ctx context.Context, req *pb.DeleteTemplateRequest) (*emptypb.Empty, error) {
	name, err := s.parseTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Template{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ModelArmorV1) GetFloorSetting(ctx context.Context, req *pb.GetFloorSettingRequest) (*pb.FloorSetting, error) {
	name, err := s.parseFloorSettingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FloorSetting{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Return a default floor setting if not found
			now := time.Now()
			obj = &pb.FloorSetting{
				Name:       fqn,
				CreateTime: timestamppb.New(now),
				UpdateTime: timestamppb.New(now),
			}
			// Should we create it in storage too?
			// Some GCP services do, some don't. Let's create it for consistency.
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
			return obj, nil
		}
		return nil, err
	}

	return obj, nil
}

func (s *ModelArmorV1) UpdateFloorSetting(ctx context.Context, req *pb.UpdateFloorSettingRequest) (*pb.FloorSetting, error) {
	name, err := s.parseFloorSettingName(req.GetFloorSetting().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.FloorSetting{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Create it if it doesn't exist? (Standard for some settings)
			now := time.Now()
			obj = proto.Clone(req.FloorSetting).(*pb.FloorSetting)
			obj.Name = fqn
			obj.CreateTime = timestamppb.New(now)
			obj.UpdateTime = timestamppb.New(now)
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
			return obj, nil
		}
		return nil, err
	}

	// TODO: Use update_mask
	proto.Merge(obj, req.FloorSetting)
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ModelArmorV1) SanitizeUserPrompt(ctx context.Context, req *pb.SanitizeUserPromptRequest) (*pb.SanitizeUserPromptResponse, error) {
	return &pb.SanitizeUserPromptResponse{
		SanitizationResult: &pb.SanitizationResult{
			FilterMatchState: pb.FilterMatchState_NO_MATCH_FOUND,
			InvocationResult: pb.InvocationResult_SUCCESS,
		},
	}, nil
}

func (s *ModelArmorV1) SanitizeModelResponse(ctx context.Context, req *pb.SanitizeModelResponseRequest) (*pb.SanitizeModelResponseResponse, error) {
	return &pb.SanitizeModelResponseResponse{
		SanitizationResult: &pb.SanitizationResult{
			FilterMatchState: pb.FilterMatchState_NO_MATCH_FOUND,
			InvocationResult: pb.InvocationResult_SUCCESS,
		},
	}, nil
}

type templateName struct {
	Project      string
	Location     string
	TemplateName string
}

func (n *templateName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/templates/%s", n.Project, n.Location, n.TemplateName)
}

// parseTemplateName parses a string into a templateName.
// The expected form is `projects/*/locations/*/templates/*`.
func (s *MockService) parseTemplateName(name string) (*templateName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "templates" {
		name := &templateName{
			Project:      tokens[1],
			Location:     tokens[3],
			TemplateName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type floorSettingName struct {
	ParentType string
	ParentID   string
	Location   string
}

func (n *floorSettingName) String() string {
	return fmt.Sprintf("%s/%s/locations/%s/floorSetting", n.ParentType, n.ParentID, n.Location)
}

func (s *MockService) parseFloorSettingName(name string) (*floorSettingName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 5 && (tokens[0] == "projects" || tokens[0] == "folders" || tokens[0] == "organizations") && tokens[2] == "locations" && tokens[4] == "floorSetting" {
		return &floorSettingName{
			ParentType: tokens[0],
			ParentID:   tokens[1],
			Location:   tokens[3],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
