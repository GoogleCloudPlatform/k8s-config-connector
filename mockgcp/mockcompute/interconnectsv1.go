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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type InterconnectsV1 struct {
	*MockService
	pb.UnimplementedInterconnectsServer
}

func (s *InterconnectsV1) Get(ctx context.Context, req *pb.GetInterconnectRequest) (*pb.Interconnect, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/interconnects/" + req.GetInterconnect()
	name, err := s.parseInterconnectName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Interconnect{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *InterconnectsV1) Insert(ctx context.Context, req *pb.InsertInterconnectRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/interconnects/" + req.GetInterconnectResource().GetName()
	name, err := s.parseInterconnectName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetInterconnectResource()).(*pb.Interconnect)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#interconnect")
	if obj.LabelFingerprint == nil {
		obj.LabelFingerprint = PtrTo(computeFingerprint(obj))
	}
	if obj.OperationalStatus == nil {
		obj.OperationalStatus = PtrTo("OS_ACTIVE")
	}
	if obj.State == nil {
		obj.State = PtrTo("ACTIVE")
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InterconnectsV1) Patch(ctx context.Context, req *pb.PatchInterconnectRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/interconnects/" + req.GetInterconnect()
	name, err := s.parseInterconnectName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Interconnect{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if req.GetInterconnectResource().Labels != nil {
		obj.Labels = req.GetInterconnectResource().GetLabels()
	}
	if req.GetInterconnectResource().AdminEnabled != nil {
		obj.AdminEnabled = req.GetInterconnectResource().AdminEnabled
	}
	if req.GetInterconnectResource().CustomerName != nil {
		obj.CustomerName = req.GetInterconnectResource().CustomerName
	}
	if req.GetInterconnectResource().Description != nil {
		obj.Description = req.GetInterconnectResource().Description
	}

	obj.LabelFingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *InterconnectsV1) Delete(ctx context.Context, req *pb.DeleteInterconnectRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/interconnects/" + req.GetInterconnect()

	name, err := s.parseInterconnectName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Interconnect{}
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
		return deleted, nil
	})
}

type interconnectName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *interconnectName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/interconnects/" + n.Name
}

func (s *MockService) parseInterconnectName(fqn string) (*interconnectName, error) {
	tokens := strings.Split(fqn, "/")
	if len(tokens) != 5 || tokens[0] != "projects" || tokens[2] != "global" || tokens[3] != "interconnects" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid interconnect name %q", fqn)
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	return &interconnectName{
		Project: project,
		Name:    tokens[4],
	}, nil
}
