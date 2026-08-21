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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type InstanceGroups struct {
	*MockService
	pb.UnimplementedInstanceGroupsServer
}

type zonalInstanceGroupName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalInstanceGroupName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/instanceGroups/" + n.Name
}

// parseZonalInstanceGroupName parses a string into a zonalInstanceGroupName.
// The expected form is `projects/*/zones/*/instanceGroups/*`.
func (s *MockService) parseZonalInstanceGroupName(name string) (*zonalInstanceGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "instanceGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
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

func (s *InstanceGroups) Get(ctx context.Context, req *pb.GetInstanceGroupRequest) (*pb.InstanceGroup, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalInstanceGroupName(reqName)
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

func (s *InstanceGroups) Insert(ctx context.Context, req *pb.InsertInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroupResource().GetName()
	name, err := s.parseZonalInstanceGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetInstanceGroupResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instanceGroup")
	obj.Zone = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))
	obj.Size = PtrTo(int32(0))
	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(PlaceholderFingerprint)
	}

	if obj.Network != nil && *obj.Network != "" {
		networkPath := *obj.Network
		if !strings.HasPrefix(networkPath, "http") {
			if !strings.Contains(networkPath, "/") {
				networkPath = "global/networks/" + networkPath
			}
			if !strings.HasPrefix(networkPath, "projects/") {
				networkPath = "projects/" + name.Project.ID + "/" + networkPath
			}
			obj.Network = PtrTo(BuildComputeSelfLink(ctx, networkPath))
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Create empty instances list
	instancesFQN := fqn + "/instances"
	instancesObj := &pb.InstanceGroupsListInstances{
		Kind: PtrTo("compute#instanceGroupsListInstances"),
		Id:   PtrTo(fmt.Sprintf("%d", id)),
	}
	if err := s.storage.Create(ctx, instancesFQN, instancesObj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstanceGroups) Delete(ctx context.Context, req *pb.DeleteInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalInstanceGroupName(reqName)
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

	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	instancesFQN := fqn + "/instances"
	instancesObj := &pb.InstanceGroupsListInstances{}
	if err := s.storage.Get(ctx, instancesFQN, instancesObj); err == nil {
		_ = s.storage.Delete(ctx, instancesFQN, instancesObj)
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstanceGroups) List(ctx context.Context, req *pb.ListInstanceGroupsRequest) (*pb.InstanceGroupList, error) {
	projectID := req.GetProject()
	zone := req.GetZone()
	project, err := s.Projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	response := &pb.InstanceGroupList{}
	response.Id = PtrTo("0123456789")
	response.Kind = PtrTo("compute#instanceGroupList")
	response.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/instanceGroups", project.ID, zone)))

	findPrefix := fmt.Sprintf("projects/%s/zones/%s/instanceGroups/", project.ID, zone)
	findKind := (&pb.InstanceGroup{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		ig := obj.(*pb.InstanceGroup)
		response.Items = append(response.Items, ig)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *InstanceGroups) AggregatedList(ctx context.Context, req *pb.AggregatedListInstanceGroupsRequest) (*pb.InstanceGroupAggregatedList, error) {
	projectID := req.GetProject()
	project, err := s.Projects.GetProjectByID(projectID)
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

	findPrefix := fmt.Sprintf("projects/%s/zones/", project.ID)
	findKind := (&pb.InstanceGroup{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		ig := obj.(*pb.InstanceGroup)

		key := fmt.Sprintf("zones/%s", lastComponent(ig.GetZone()))
		scopedList := response.Items[key]
		if scopedList == nil {
			scopedList = &pb.InstanceGroupsScopedList{}
			response.Items[key] = scopedList
		}
		scopedList.InstanceGroups = append(scopedList.InstanceGroups, ig)

		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *InstanceGroups) ListInstances(ctx context.Context, req *pb.ListInstancesInstanceGroupsRequest) (*pb.InstanceGroupsListInstances, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalInstanceGroupName(reqName)
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

	instancesFQN := fqn + "/instances"
	instancesObj := &pb.InstanceGroupsListInstances{}
	if err := s.storage.Get(ctx, instancesFQN, instancesObj); err != nil {
		if status.Code(err) == codes.NotFound {
			instancesObj.Kind = PtrTo("compute#instanceGroupsListInstances")
			return instancesObj, nil
		}
		return nil, err
	}

	return instancesObj, nil
}

func (s *InstanceGroups) AddInstances(ctx context.Context, req *pb.AddInstancesInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalInstanceGroupName(reqName)
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

	instancesFQN := fqn + "/instances"
	instancesObj := &pb.InstanceGroupsListInstances{}
	if err := s.storage.Get(ctx, instancesFQN, instancesObj); err != nil {
		if status.Code(err) == codes.NotFound {
			instancesObj = &pb.InstanceGroupsListInstances{
				Kind: PtrTo("compute#instanceGroupsListInstances"),
				Id:   PtrTo(fmt.Sprintf("%d", *obj.Id)),
			}
		} else {
			return nil, err
		}
	}

	existingInstances := make(map[string]bool)
	for _, item := range instancesObj.Items {
		existingInstances[item.GetInstance()] = true
	}

	for _, ref := range req.GetInstanceGroupsAddInstancesRequestResource().GetInstances() {
		instanceURL := ref.GetInstance()
		if !strings.HasPrefix(instanceURL, "http") {
			instanceURL = BuildComputeSelfLink(ctx, strings.TrimPrefix(instanceURL, "/"))
		}
		if existingInstances[instanceURL] {
			continue
		}
		instancesObj.Items = append(instancesObj.Items, &pb.InstanceWithNamedPorts{
			Instance:   PtrTo(instanceURL),
			Status:     PtrTo("RUNNING"),
			NamedPorts: obj.NamedPorts,
		})
	}

	if err := s.storage.Update(ctx, instancesFQN, instancesObj); err != nil {
		return nil, err
	}

	obj.Size = PtrTo(int32(len(instancesObj.Items)))
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("addInstances"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstanceGroups) RemoveInstances(ctx context.Context, req *pb.RemoveInstancesInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalInstanceGroupName(reqName)
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

	instancesFQN := fqn + "/instances"
	instancesObj := &pb.InstanceGroupsListInstances{}
	if err := s.storage.Get(ctx, instancesFQN, instancesObj); err != nil {
		if status.Code(err) == codes.NotFound {
			instancesObj = &pb.InstanceGroupsListInstances{
				Kind: PtrTo("compute#instanceGroupsListInstances"),
				Id:   PtrTo(fmt.Sprintf("%d", *obj.Id)),
			}
		} else {
			return nil, err
		}
	}

	toRemove := make(map[string]bool)
	for _, ref := range req.GetInstanceGroupsRemoveInstancesRequestResource().GetInstances() {
		instanceURL := ref.GetInstance()
		if !strings.HasPrefix(instanceURL, "http") {
			instanceURL = BuildComputeSelfLink(ctx, strings.TrimPrefix(instanceURL, "/"))
		}
		toRemove[instanceURL] = true
	}

	var keptItems []*pb.InstanceWithNamedPorts
	for _, item := range instancesObj.Items {
		if !toRemove[item.GetInstance()] {
			keptItems = append(keptItems, item)
		}
	}
	instancesObj.Items = keptItems

	if err := s.storage.Update(ctx, instancesFQN, instancesObj); err != nil {
		return nil, err
	}

	obj.Size = PtrTo(int32(len(instancesObj.Items)))
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("removeInstances"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstanceGroups) SetNamedPorts(ctx context.Context, req *pb.SetNamedPortsInstanceGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroups/" + req.GetInstanceGroup()
	name, err := s.parseZonalInstanceGroupName(reqName)
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

	namedPorts := req.GetInstanceGroupsSetNamedPortsRequestResource().GetNamedPorts()
	obj.NamedPorts = namedPorts

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	instancesFQN := fqn + "/instances"
	instancesObj := &pb.InstanceGroupsListInstances{}
	if err := s.storage.Get(ctx, instancesFQN, instancesObj); err == nil {
		for _, item := range instancesObj.Items {
			item.NamedPorts = namedPorts
		}
		_ = s.storage.Update(ctx, instancesFQN, instancesObj)
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setNamedPorts"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}
