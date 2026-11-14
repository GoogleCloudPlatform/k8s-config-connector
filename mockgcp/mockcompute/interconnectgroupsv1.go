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

type InterconnectGroupsV1 struct {
	*MockService
	pb.UnimplementedInterconnectGroupsServer
}

type globalInterconnectGroupName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalInterconnectGroupName) String() string {
	return "projects/" + n.Project.ID + "/global/interconnectGroups/" + n.Name
}

func (s *MockService) parseGlobalInterconnectGroupName(name string) (*globalInterconnectGroupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "interconnectGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &globalInterconnectGroupName{
			Project: project,
			Name:    tokens[4],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *InterconnectGroupsV1) populateConfigured(obj *pb.InterconnectGroup) {
	if obj.Intent == nil || obj.Intent.GetTopologyCapability() == "NO_SLA" {
		obj.Configured = &pb.InterconnectGroupConfigured{
			TopologyCapability: &pb.InterconnectGroupConfiguredTopologyCapability{
				SupportedSla: PtrTo("NO_SLA"),
			},
		}
	} else {
		// When intent is set to a capability requiring interconnects but none are present
		obj.Configured = &pb.InterconnectGroupConfigured{
			TopologyCapability: &pb.InterconnectGroupConfiguredTopologyCapability{
				SupportedSla: PtrTo("NO_SLA"),
				IntendedCapabilityBlockers: []*pb.InterconnectGroupConfiguredTopologyCapabilityIntendedCapabilityBlockers{
					{
						BlockerType:       PtrTo("NO_INTERCONNECTS"),
						DocumentationLink: PtrTo("https://cloud.google.com/network-connectivity/docs/interconnect/support/fix-sla#no_interconnects"),
						Explanation:       PtrTo("There are no Interconnects in the given group. Configure new Interconnects, or add existing Interconnects to the group."),
					},
				},
			},
		}
	}
}

func (s *InterconnectGroupsV1) Get(ctx context.Context, req *pb.GetInterconnectGroupRequest) (*pb.InterconnectGroup, error) {
	reqName := "projects/" + req.GetProject() + "/global/interconnectGroups/" + req.GetInterconnectGroup()
	name, err := s.parseGlobalInterconnectGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InterconnectGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *InterconnectGroupsV1) Insert(ctx context.Context, req *pb.InsertInterconnectGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/interconnectGroups/" + req.GetInterconnectGroupResource().GetName()
	name, err := s.parseGlobalInterconnectGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetInterconnectGroupResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#InterconnectGroup")
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.PhysicalStructure == nil {
		obj.PhysicalStructure = &pb.InterconnectGroupPhysicalStructure{}
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

func (s *InterconnectGroupsV1) Patch(ctx context.Context, req *pb.PatchInterconnectGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/interconnectGroups/" + req.GetInterconnectGroup()
	name, err := s.parseGlobalInterconnectGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InterconnectGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	reqObj := req.GetInterconnectGroupResource()
	if updateMask := req.GetUpdateMask(); updateMask != "" {
		for _, path := range strings.Split(updateMask, ",") {
			switch path {
			case "description":
				obj.Description = reqObj.Description
			case "intent":
				obj.Intent = reqObj.Intent
			case "interconnects":
				obj.Interconnects = reqObj.Interconnects
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
		if reqObj.Interconnects != nil {
			obj.Interconnects = reqObj.Interconnects
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

func (s *InterconnectGroupsV1) Delete(ctx context.Context, req *pb.DeleteInterconnectGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/interconnectGroups/" + req.GetInterconnectGroup()
	name, err := s.parseGlobalInterconnectGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.InterconnectGroup{}
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

func (s *InterconnectGroupsV1) List(ctx context.Context, req *pb.ListInterconnectGroupsRequest) (*pb.InterconnectGroupsListResponse, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	response := &pb.InterconnectGroupsListResponse{
		Id:       PtrTo("0123456789"),
		Kind:     PtrTo("compute#interconnectGroupList"),
		SelfLink: PtrTo(BuildComputeSelfLink(ctx, "projects/"+project.ID+"/global/interconnectGroups")),
	}

	findKind := (&pb.InterconnectGroup{}).ProtoReflect().Descriptor()
	prefix := "projects/" + project.ID + "/global/interconnectGroups/"
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		ig := obj.(*pb.InterconnectGroup)
		response.Items = append(response.Items, ig)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
