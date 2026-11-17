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

package mockces

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/ces/apiv1/cespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *AgentServiceV1) GetGuardrail(ctx context.Context, req *pb.GetGuardrailRequest) (*pb.Guardrail, error) {
	name, err := s.parseGuardrailName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Guardrail{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Guardrail %s does not exist in the app.", name.GuardrailID)
		}
		return nil, err
	}
	result := proto.Clone(obj).(*pb.Guardrail)
	result.Etag = "abcdef0123A="
	return result, nil
}

func (s *AgentServiceV1) CreateGuardrail(ctx context.Context, req *pb.CreateGuardrailRequest) (*pb.Guardrail, error) {
	parent, err := s.parseAppName(req.GetParent())
	if err != nil {
		return nil, err
	}

	parentApp := &pb.App{}
	if err := s.storage.Get(ctx, parent.String(), parentApp); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Parent resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	guardrailID := req.GetGuardrailId()
	if guardrailID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "guardrailId is required")
	}

	name := &guardrailName{
		Project:     parent.Project,
		Location:    parent.Location,
		AppID:       parent.AppID,
		GuardrailID: guardrailID,
	}

	fqn := name.String()

	existing := &pb.Guardrail{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Resource '%s' already exists", fqn)
	}

	obj := proto.Clone(req.GetGuardrail()).(*pb.Guardrail)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "abcdef0123A="

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Real GCP CreateGuardrail returns the guardrail without etag
	result := proto.Clone(obj).(*pb.Guardrail)
	result.Etag = ""
	return result, nil
}

func (s *AgentServiceV1) UpdateGuardrail(ctx context.Context, req *pb.UpdateGuardrailRequest) (*pb.Guardrail, error) {
	reqGuardrail := req.GetGuardrail()
	if reqGuardrail == nil {
		return nil, status.Errorf(codes.InvalidArgument, "guardrail is required")
	}

	name, err := s.parseGuardrailName(reqGuardrail.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Guardrail{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Guardrail %s does not exist in the app.", name.GuardrailID)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Merge(obj, reqGuardrail)
	} else {
		if err := fields.UpdateByFieldMask(obj, reqGuardrail, paths); err != nil {
			return nil, err
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Real GCP UpdateGuardrail returns etag
	result := proto.Clone(obj).(*pb.Guardrail)
	result.Etag = "abcdef0123A="
	return result, nil
}

func (s *AgentServiceV1) DeleteGuardrail(ctx context.Context, req *pb.DeleteGuardrailRequest) (*emptypb.Empty, error) {
	name, err := s.parseGuardrailName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Guardrail{}
	if err := s.storage.Get(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Guardrail %s does not exist in the app.", name.GuardrailID)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, &pb.Guardrail{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AgentServiceV1) ListGuardrails(ctx context.Context, req *pb.ListGuardrailsRequest) (*pb.ListGuardrailsResponse, error) {
	parent, err := s.parseAppName(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &pb.ListGuardrailsResponse{}
	findPrefix := fmt.Sprintf("projects/%s/locations/%s/apps/%s/guardrails/", parent.Project.ID, parent.Location, parent.AppID)

	guardrailKind := (&pb.Guardrail{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, guardrailKind, storage.ListOptions{}, func(obj proto.Message) error {
		item := obj.(*pb.Guardrail)
		if strings.HasPrefix(item.GetName(), findPrefix) {
			itemCopy := proto.Clone(item).(*pb.Guardrail)
			response.Guardrails = append(response.Guardrails, itemCopy)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
