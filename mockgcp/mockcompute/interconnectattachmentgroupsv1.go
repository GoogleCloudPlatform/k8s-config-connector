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

package mockcompute

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type InterconnectAttachmentGroupsV1 struct {
	*MockService
	pb.UnimplementedInterconnectAttachmentGroupsServer
}

type globalInterconnectAttachmentGroupName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalInterconnectAttachmentGroupName) String() string {
	return "projects/" + n.Project.ID + "/global/interconnectAttachmentGroups/" + n.Name
}

func (s *MockService) parseGlobalInterconnectAttachmentGroupName(name string) (*globalInterconnectAttachmentGroupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "interconnectAttachmentGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &globalInterconnectAttachmentGroupName{
			Project: project,
			Name:    tokens[4],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *InterconnectAttachmentGroupsV1) populateConfigured(obj *pb.InterconnectAttachmentGroup) {
	if obj.Intent == nil || obj.Intent.GetAvailabilitySla() == "NO_SLA" {
		obj.Configured = &pb.InterconnectAttachmentGroupConfigured{
			AvailabilitySla: &pb.InterconnectAttachmentGroupConfiguredAvailabilitySLA{
				EffectiveSla: PtrTo("NO_SLA"),
			},
		}
	} else {
		// When intent is set to an SLA requiring attachments but none are present
		obj.Configured = &pb.InterconnectAttachmentGroupConfigured{
			AvailabilitySla: &pb.InterconnectAttachmentGroupConfiguredAvailabilitySLA{
				EffectiveSla: PtrTo("NO_SLA"),
				IntendedSlaBlockers: []*pb.InterconnectAttachmentGroupConfiguredAvailabilitySLAIntendedSlaBlockers{
					{
						BlockerType:       PtrTo("NO_ATTACHMENTS"),
						DocumentationLink: PtrTo("https://cloud.google.com/network-connectivity/docs/interconnect/support/fix-sla#no_attachments"),
						Explanation:       PtrTo("There are no attachments in the given group. Configure new attachments, or add existing attachments to the group."),
					},
				},
			},
		}
	}
}

func (s *InterconnectAttachmentGroupsV1) Get(ctx context.Context, req *pb.GetInterconnectAttachmentGroupRequest) (*pb.InterconnectAttachmentGroup, error) {
	reqName := "projects/" + req.GetProject() + "/global/interconnectAttachmentGroups/" + req.GetInterconnectAttachmentGroup()
	name, err := s.parseGlobalInterconnectAttachmentGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InterconnectAttachmentGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *InterconnectAttachmentGroupsV1) Insert(ctx context.Context, req *pb.InsertInterconnectAttachmentGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/interconnectAttachmentGroups/" + req.GetInterconnectAttachmentGroupResource().GetName()
	name, err := s.parseGlobalInterconnectAttachmentGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetInterconnectAttachmentGroupResource()).(*pb.InterconnectAttachmentGroup)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#interconnectAttachmentGroup")
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.LogicalStructure == nil {
		obj.LogicalStructure = &pb.InterconnectAttachmentGroupLogicalStructure{}
	}
	s.populateConfigured(obj)
	obj.Etag = PtrTo(computeFingerprint(obj))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      &id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InterconnectAttachmentGroupsV1) Patch(ctx context.Context, req *pb.PatchInterconnectAttachmentGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/interconnectAttachmentGroups/" + req.GetInterconnectAttachmentGroup()
	name, err := s.parseGlobalInterconnectAttachmentGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InterconnectAttachmentGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	reqObj := req.GetInterconnectAttachmentGroupResource()
	if updateMask := req.GetUpdateMask(); updateMask != "" {
		for _, path := range strings.Split(updateMask, ",") {
			switch path {
			case "description":
				obj.Description = reqObj.Description
			case "intent":
				obj.Intent = reqObj.Intent
			case "attachments":
				obj.Attachments = reqObj.Attachments
			case "interconnectGroup", "interconnect_group":
				obj.InterconnectGroup = reqObj.InterconnectGroup
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported", path)
			}
		}
	} else {
		if reqObj.Description != nil {
			obj.Description = reqObj.Description
		}
		if reqObj.Intent != nil {
			obj.Intent = reqObj.Intent
		}
		if reqObj.Attachments != nil {
			obj.Attachments = reqObj.Attachments
		}
		if reqObj.InterconnectGroup != nil {
			obj.InterconnectGroup = reqObj.InterconnectGroup
		}
	}
	s.populateConfigured(obj)
	obj.Etag = PtrTo(computeFingerprint(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InterconnectAttachmentGroupsV1) Delete(ctx context.Context, req *pb.DeleteInterconnectAttachmentGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/interconnectAttachmentGroups/" + req.GetInterconnectAttachmentGroup()
	name, err := s.parseGlobalInterconnectAttachmentGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.InterconnectAttachmentGroup{}
	if err := s.storage.Get(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *InterconnectAttachmentGroupsV1) List(ctx context.Context, req *pb.ListInterconnectAttachmentGroupsRequest) (*pb.InterconnectAttachmentGroupsListResponse, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	response := &pb.InterconnectAttachmentGroupsListResponse{
		Id:       PtrTo("0123456789"),
		Kind:     PtrTo("compute#interconnectAttachmentGroupList"),
		SelfLink: PtrTo(BuildComputeSelfLink(ctx, "projects/"+project.ID+"/global/interconnectAttachmentGroups")),
	}

	findKind := (&pb.InterconnectAttachmentGroup{}).ProtoReflect().Descriptor()
	prefix := "projects/" + project.ID + "/global/interconnectAttachmentGroups/"
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		iag := obj.(*pb.InterconnectAttachmentGroup)
		response.Items = append(response.Items, iag)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
