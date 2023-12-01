// Copyright 2022 Google LLC
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
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type ZonalInstanceGroupsV1 struct {
	*MockService
	pb.UnimplementedInstanceGroupsServer
}

func (s *ZonalInstanceGroupsV1) Get(ctx context.Context, req *pb.GetInstanceGroupRequest) (*pb.InstanceGroup, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalInstanceGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstanceGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "instanceGroup %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading instanceGroup: %v", err)
		}
	}

	return obj, nil
}

func (s *ZonalInstanceGroupsV1) Insert(ctx context.Context, req *pb.InsertInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroupResource().GetName()
	name, err := s.parseZonalInstanceGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetInstanceGroupResource()).(*pb.InstanceGroup)
	obj.SelfLink = PtrTo("https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instanceGroup")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating instanceGroup: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ZonalInstanceGroupsV1) Delete(ctx context.Context, req *pb.DeleteInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalInstanceGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.InstanceGroup{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "instanceGroup %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting instanceGroup: %v", err)
		}
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ZonalInstanceGroupsV1) ListInstances(ctx context.Context, req *pb.ListInstancesInstanceGroupsRequest) (*pb.InstanceGroupsListInstances, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalInstanceGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String() + "/listInstances"

	obj := &pb.InstanceGroupsListInstances{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			// ignore - assume no instances have been added yet
		} else {
			return nil, status.Errorf(codes.Internal, "error reading instanceGroup: %v", err)
		}
	}

	return obj, nil
}

func (s *ZonalInstanceGroupsV1) AddInstances(ctx context.Context, req *pb.AddInstancesInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalInstanceGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String() + "/listInstances"

	create := false
	obj := &pb.InstanceGroupsListInstances{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			// ignore - assume no instances
			create = true
		} else {
			return nil, status.Errorf(codes.Internal, "error reading instanceGroup instances: %v", err)
		}
	}

	for _, instanceRef := range req.GetInstanceGroupsAddInstancesRequestResource().GetInstances() {
		instance := instanceRef.GetInstance()
		item := &pb.InstanceWithNamedPorts{
			Instance: &instance,
		}
		obj.Items = append(obj.Items, item)
	}

	if create {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, status.Errorf(codes.Internal, "error creating instanceGroup instances: %v", err)
		}

	} else {
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, status.Errorf(codes.Internal, "error updating instanceGroup instances: %v", err)
		}
	}

	return s.newLRO(ctx, name.Project.ID)
}

type zonalInstanceGroupName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalInstanceGroupName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/instanceGroups/" + n.Name
}

// parseZonalInstanceGroupName parses a string into a instancegroupName.
// The expected form is `projects/*/regions/*/instancegroup/*`.
func (s *MockService) parseZonalInstanceGroupName(name string) (*zonalInstanceGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "instanceGroups" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zonalInstanceGroupName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
