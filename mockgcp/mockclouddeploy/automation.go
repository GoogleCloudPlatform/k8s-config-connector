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
// proto.service: google.cloud.deploy.v1.CloudDeploy
// proto.message: google.cloud.deploy.v1.Automation

package mockclouddeploy

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/deploy/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
)

func (s *cloudDeploy) ListAutomations(ctx context.Context, req *pb.ListAutomationsRequest) (*pb.ListAutomationsResponse, error) {
	parent, err := s.parseDeliveryPipelineName(req.Parent)
	if err != nil {
		return nil, err
	}

	var automations []*pb.Automation
	automationKind := (&pb.Automation{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, automationKind, storage.ListOptions{}, func(obj proto.Message) error {
		automation := obj.(*pb.Automation)
		name, err := s.parseAutomationName(automation.Name)
		if err != nil {
			return nil // Should not happen
		}

		if name.Project.ID == parent.Project.ID && name.Location == parent.Location && name.DeliveryPipeline == parent.DeliveryPipeline {
			automations = append(automations, automation)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListAutomationsResponse{
		Automations: automations,
	}, nil
}

func (s *cloudDeploy) GetAutomation(ctx context.Context, req *pb.GetAutomationRequest) (*pb.Automation, error) {
	name, err := s.parseAutomationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Automation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *cloudDeploy) CreateAutomation(ctx context.Context, req *pb.CreateAutomationRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/automations/%s", req.Parent, req.AutomationId)
	name, err := s.parseAutomationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.Automation).(*pb.Automation)
	obj.Name = fqn

	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = uuid.NewString()

	s.populateAutomationDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *cloudDeploy) UpdateAutomation(ctx context.Context, req *pb.UpdateAutomationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAutomationName(req.Automation.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Automation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound && req.AllowMissing {
			obj = proto.Clone(req.Automation).(*pb.Automation)
			obj.Name = fqn
			obj.Uid = uuid.NewString()
			obj.CreateTime = timestamppb.New(time.Now())
			obj.UpdateTime = timestamppb.New(time.Now())
			obj.Etag = uuid.NewString()
			s.populateAutomationDefaults(obj)

			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
		} else {
			if status.Code(err) == codes.NotFound {
				return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
			}
			return nil, err
		}
	} else {
		req.Automation.Uid = obj.GetUid()
		req.Automation.CreateTime = obj.GetCreateTime()
		req.Automation.UpdateTime = timestamppb.New(time.Now())
		req.Automation.Etag = obj.GetEtag()

		// Apply the update mask to the object.
		paths := req.GetUpdateMask().GetPaths()
		if len(paths) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
		}

		// Handle * in update mask
		hasWildcard := false
		for _, path := range paths {
			if path == "*" {
				hasWildcard = true
				break
			}
		}

		if hasWildcard {
			obj = proto.Clone(req.Automation).(*pb.Automation)
			obj.Name = fqn
			obj.Uid = req.Automation.Uid
			obj.CreateTime = req.Automation.CreateTime
			s.populateAutomationDefaults(obj)
		} else {
			if err := fields.UpdateByFieldMask(obj, req.Automation, req.UpdateMask.Paths); err != nil {
				return nil, fmt.Errorf("update field_mask.paths: %w", err)
			}
		}

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *cloudDeploy) DeleteAutomation(ctx context.Context, req *pb.DeleteAutomationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAutomationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Automation{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound && req.AllowMissing {
			// Success
		} else {
			if status.Code(err) == codes.NotFound {
				return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
			}
			return nil, err
		}
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *cloudDeploy) populateAutomationDefaults(obj *pb.Automation) {
	if obj.Annotations == nil {
		obj.Annotations = make(map[string]string)
	}
	if obj.Labels == nil {
		obj.Labels = make(map[string]string)
	}
	if obj.Selector != nil {
		for _, target := range obj.Selector.Targets {
			if target.Labels == nil {
				target.Labels = make(map[string]string)
			}
		}
	}
	for _, rule := range obj.Rules {
		if rule.GetPromoteReleaseRule() != nil {
			if rule.GetPromoteReleaseRule().Condition == nil {
				rule.GetPromoteReleaseRule().Condition = &pb.AutomationRuleCondition{}
			}
			if rule.GetPromoteReleaseRule().Condition.TargetsPresentCondition == nil {
				rule.GetPromoteReleaseRule().Condition.TargetsPresentCondition = &pb.TargetsPresentCondition{}
			}
		}
	}
}

type automationName struct {
	Project          *projects.ProjectData
	Location         string
	DeliveryPipeline string
	Automation       string
}

func (n *automationName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automations/%s", n.Project.ID, n.Location, n.DeliveryPipeline, n.Automation)
}

// parseAutomationName parses a string into a automationName.
// The expected form is `projects/*/locations/*/deliveryPipelines/*/automations/*`.
func (s *MockService) parseAutomationName(name string) (*automationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deliveryPipelines" && tokens[6] == "automations" {
		for i := 1; i < len(tokens); i += 2 {
			if tokens[i] == "" {
				return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
			}
		}

		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &automationName{
			Project:          project,
			Location:         tokens[3],
			DeliveryPipeline: tokens[5],
			Automation:       tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
