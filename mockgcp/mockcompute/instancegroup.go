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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/regions"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type InstanceGroupsV1 struct {
	*MockService
	pb.UnimplementedInstanceGroupsServer
}

type zonalIGName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalIGName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/instanceGroups/" + n.Name
}

// parseZonalIGName parses a string into a zonalIGName.
// The expected form is `projects/*/zones/*/instanceGroups/*`.
func (s *MockService) parseZonalIGName(name string) (*zonalIGName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "instanceGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zonalIGName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *InstanceGroupsV1) Get(ctx context.Context, req *pb.GetInstanceGroupRequest) (*pb.InstanceGroup, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalIGName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstanceGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *InstanceGroupsV1) Insert(ctx context.Context, req *pb.InsertInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroupResource().GetName()
	name, err := s.parseZonalIGName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetInstanceGroupResource()).(*pb.InstanceGroup)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instanceGroup")
	obj.Zone = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))
	if obj.Size == nil {
		obj.Size = PtrTo(int32(0))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	instancesObj := &pb.InstanceGroupsListInstances{
		Kind:     PtrTo("compute#instanceGroupsListInstances"),
		SelfLink: PtrTo(BuildComputeSelfLink(ctx, fqn+"/listInstances")),
	}
	if err := s.storage.Create(ctx, fqn+"/instances", instancesObj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.instanceGroups.insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstanceGroupsV1) Delete(ctx context.Context, req *pb.DeleteInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalIGName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstanceGroup{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	instancesObj := &pb.InstanceGroupsListInstances{}
	_ = s.storage.Delete(ctx, fqn+"/instances", instancesObj)

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.instanceGroups.delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstanceGroupsV1) List(ctx context.Context, req *pb.ListInstanceGroupsRequest) (*pb.InstanceGroupList, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/zones/%s/instanceGroups/", project.ID, req.GetZone())

	var items []*pb.InstanceGroup
	kind := (&pb.InstanceGroup{}).ProtoReflect().Descriptor()
	err = s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		items = append(items, obj.(*pb.InstanceGroup))
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.InstanceGroupList{
		Items: items,
		Kind:  PtrTo("compute#instanceGroupList"),
	}, nil
}

func (s *InstanceGroupsV1) AddInstances(ctx context.Context, req *pb.AddInstancesInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalIGName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	igObj := &pb.InstanceGroup{}
	if err := s.storage.Get(ctx, fqn, igObj); err != nil {
		return nil, err
	}

	instancesObj := &pb.InstanceGroupsListInstances{}
	if err := s.storage.Get(ctx, fqn+"/instances", instancesObj); err != nil {
		return nil, err
	}

	for _, ref := range req.GetInstanceGroupsAddInstancesRequestResource().GetInstances() {
		instanceURL := ref.GetInstance()
		if instanceURL == "" {
			continue
		}
		exists := false
		for _, item := range instancesObj.Items {
			if item.GetInstance() == instanceURL {
				exists = true
				break
			}
		}
		if !exists {
			instancesObj.Items = append(instancesObj.Items, &pb.InstanceWithNamedPorts{
				Instance: PtrTo(instanceURL),
				Status:   PtrTo("RUNNING"),
			})
		}
	}

	igObj.Size = PtrTo(int32(len(instancesObj.Items)))

	if err := s.storage.Update(ctx, fqn, igObj); err != nil {
		return nil, err
	}
	if err := s.storage.Update(ctx, fqn+"/instances", instancesObj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      igObj.Id,
		TargetLink:    igObj.SelfLink,
		OperationType: PtrTo("compute.instanceGroups.addInstances"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return igObj, nil
	})
}

func (s *InstanceGroupsV1) RemoveInstances(ctx context.Context, req *pb.RemoveInstancesInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalIGName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	igObj := &pb.InstanceGroup{}
	if err := s.storage.Get(ctx, fqn, igObj); err != nil {
		return nil, err
	}

	instancesObj := &pb.InstanceGroupsListInstances{}
	if err := s.storage.Get(ctx, fqn+"/instances", instancesObj); err != nil {
		return nil, err
	}

	var kept []*pb.InstanceWithNamedPorts
	for _, item := range instancesObj.Items {
		removed := false
		for _, ref := range req.GetInstanceGroupsRemoveInstancesRequestResource().GetInstances() {
			if item.GetInstance() == ref.GetInstance() {
				removed = true
				break
			}
		}
		if !removed {
			kept = append(kept, item)
		}
	}
	instancesObj.Items = kept

	igObj.Size = PtrTo(int32(len(instancesObj.Items)))

	if err := s.storage.Update(ctx, fqn, igObj); err != nil {
		return nil, err
	}
	if err := s.storage.Update(ctx, fqn+"/instances", instancesObj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      igObj.Id,
		TargetLink:    igObj.SelfLink,
		OperationType: PtrTo("compute.instanceGroups.removeInstances"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return igObj, nil
	})
}

func (s *InstanceGroupsV1) ListInstances(ctx context.Context, req *pb.ListInstancesInstanceGroupsRequest) (*pb.InstanceGroupsListInstances, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalIGName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	instancesObj := &pb.InstanceGroupsListInstances{}
	if err := s.storage.Get(ctx, fqn+"/instances", instancesObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return instancesObj, nil
}

func (s *InstanceGroupsV1) SetNamedPorts(ctx context.Context, req *pb.SetNamedPortsInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalIGName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	igObj := &pb.InstanceGroup{}
	if err := s.storage.Get(ctx, fqn, igObj); err != nil {
		return nil, err
	}

	igObj.NamedPorts = req.GetInstanceGroupsSetNamedPortsRequestResource().GetNamedPorts()
	igObj.Fingerprint = PtrTo(computeFingerprint(igObj))

	if err := s.storage.Update(ctx, fqn, igObj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      igObj.Id,
		TargetLink:    igObj.SelfLink,
		OperationType: PtrTo("compute.instanceGroups.setNamedPorts"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return igObj, nil
	})
}

func (s *InstanceGroupsV1) AggregatedList(ctx context.Context, req *pb.AggregatedListInstanceGroupsRequest) (*pb.InstanceGroupAggregatedList, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	response := &pb.InstanceGroupAggregatedList{}
	response.Id = PtrTo("0123456789")
	response.Kind = PtrTo("compute#instanceGroupAggregatedList")
	response.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/aggregated/instanceGroups", project.ID)))
	response.Items = make(map[string]*pb.InstanceGroupsScopedList)

	for _, region := range regions.GetAllRegions(ctx) {
		for _, zone := range region.Zones(ctx) {
			response.Items["zones/"+zone] = &pb.InstanceGroupsScopedList{}
		}
	}

	findPrefix := fmt.Sprintf("projects/%s/", project.ID)
	findKind := (&pb.InstanceGroup{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		ig := obj.(*pb.InstanceGroup)

		zoneName := lastComponent(ig.GetZone())
		if zoneName != "" {
			key := fmt.Sprintf("zones/%s", zoneName)
			scopedList := response.Items[key]
			if scopedList == nil {
				scopedList = &pb.InstanceGroupsScopedList{}
				response.Items[key] = scopedList
			}
			scopedList.InstanceGroups = append(scopedList.InstanceGroups, ig)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
