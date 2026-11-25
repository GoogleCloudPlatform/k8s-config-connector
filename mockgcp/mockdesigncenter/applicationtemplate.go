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

package mockdesigncenter

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *DesignCenterServer) GetApplicationTemplate(ctx context.Context, req *pb.GetApplicationTemplateRequest) (*pb.ApplicationTemplate, error) {
	name, err := s.parseApplicationTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ApplicationTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return s.populateApplicationTemplate(obj), nil
}

func (s *DesignCenterServer) CreateApplicationTemplate(ctx context.Context, req *pb.CreateApplicationTemplateRequest) (*pb.ApplicationTemplate, error) {
	parent, err := s.parseSpaceName(req.Parent)
	if err != nil {
		return nil, err
	}

	// Reject if parent space does not exist
	parentSpace := &pb.Space{}
	if err := s.storage.Get(ctx, parent.String(), parentSpace); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	if req.ApplicationTemplateId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "applicationTemplateId is required")
	}

	if req.ApplicationTemplate == nil {
		return nil, status.Errorf(codes.InvalidArgument, "applicationTemplate is required")
	}

	applicationTemplateID := req.ApplicationTemplateId
	fqn := fmt.Sprintf("%s/applicationTemplates/%s", parent.String(), applicationTemplateID)

	name, err := s.parseApplicationTemplateName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.ApplicationTemplate).(*pb.ApplicationTemplate)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj = s.populateApplicationTemplate(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DesignCenterServer) UpdateApplicationTemplate(ctx context.Context, req *pb.UpdateApplicationTemplateRequest) (*pb.ApplicationTemplate, error) {
	if req.ApplicationTemplate == nil {
		return nil, status.Errorf(codes.InvalidArgument, "applicationTemplate is required")
	}

	name, err := s.parseApplicationTemplateName(req.ApplicationTemplate.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.ApplicationTemplate{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.ApplicationTemplate)
	now := time.Now()
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	if err := fields.UpdateByFieldMask(updated, req.ApplicationTemplate, paths); err != nil {
		return nil, err
	}

	updated = s.populateApplicationTemplate(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *DesignCenterServer) DeleteApplicationTemplate(ctx context.Context, req *pb.DeleteApplicationTemplateRequest) (*emptypb.Empty, error) {
	name, err := s.parseApplicationTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ApplicationTemplate{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DesignCenterServer) ListApplicationTemplates(ctx context.Context, req *pb.ListApplicationTemplatesRequest) (*pb.ListApplicationTemplatesResponse, error) {
	parent, err := s.parseSpaceName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/applicationTemplates/"
	response := &pb.ListApplicationTemplatesResponse{}

	findOpts := storage.ListOptions{
		Prefix: prefix,
	}
	kind := (&pb.ApplicationTemplate{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, findOpts, func(item proto.Message) error {
		obj := item.(*pb.ApplicationTemplate)
		response.ApplicationTemplates = append(response.ApplicationTemplates, s.populateApplicationTemplate(obj))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *DesignCenterServer) populateApplicationTemplate(obj *pb.ApplicationTemplate) *pb.ApplicationTemplate {
	if obj.CompositionType == pb.ApplicationCompositionType_APPLICATION_COMPOSITION_TYPE_UNSPECIFIED {
		obj.CompositionType = pb.ApplicationCompositionType_STANDARD
	}
	if obj.IacFormat == pb.ApplicationTemplate_IAC_FORMAT_UNSPECIFIED {
		obj.IacFormat = pb.ApplicationTemplate_TERRAFORM
	}

	sat := &pb.SerializedApplicationTemplate{
		Uri:                         obj.Name,
		CompositionType:             obj.CompositionType,
		IacFormat:                   obj.IacFormat,
		DisplayName:                 obj.DisplayName,
		Description:                 obj.Description,
		ApplicationParameters:       obj.ApplicationParameters,
		InputsAsRootModuleVariables: obj.InputsAsRootModuleVariables,
	}
	obj.SerializedApplicationTemplate = sat

	return obj
}
