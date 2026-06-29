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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

// Zonal InstanceGroupManagers
type instanceGroupManagers struct {
	*MockService
	pb.UnimplementedInstanceGroupManagersServer
}

func (s *instanceGroupManagers) Get(ctx context.Context, req *pb.GetInstanceGroupManagerRequest) (*pb.InstanceGroupManager, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroupManagers/" + req.GetInstanceGroupManager()
	name, err := s.parseZonalIGMName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstanceGroupManager{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *instanceGroupManagers) Insert(ctx context.Context, req *pb.InsertInstanceGroupManagerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroupManagers/" + req.GetInstanceGroupManagerResource().GetName()
	name, err := s.parseZonalIGMName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetInstanceGroupManagerResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instanceGroupManager")
	obj.Zone = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))

	// Underlying InstanceGroup
	obj.InstanceGroup = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/instanceGroups/%s", name.Project.ID, name.Zone, *obj.Name)))

	// Fingerprint
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	// Status fields
	obj.Status = &pb.InstanceGroupManagerStatus{
		IsStable: PtrTo(true),
		VersionTarget: &pb.InstanceGroupManagerStatusVersionTarget{
			IsReached: PtrTo(true),
		},
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
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *instanceGroupManagers) Patch(ctx context.Context, req *pb.PatchInstanceGroupManagerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroupManagers/" + req.GetInstanceGroupManager()
	name, err := s.parseZonalIGMName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.InstanceGroupManager{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// We handle slices separately to avoid appending
	autoHealingPolicies := req.GetInstanceGroupManagerResource().GetAutoHealingPolicies()
	versions := req.GetInstanceGroupManagerResource().GetVersions()

	// Clone the update resource and clear slices to avoid append-merge
	patchSrc := proto.CloneOf(req.GetInstanceGroupManagerResource())
	patchSrc.AutoHealingPolicies = nil
	patchSrc.Versions = nil

	proto.Merge(obj, patchSrc)

	if autoHealingPolicies != nil {
		obj.AutoHealingPolicies = autoHealingPolicies
	}
	if versions != nil {
		obj.Versions = versions
	}

	// Update fingerprint on change
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *instanceGroupManagers) Delete(ctx context.Context, req *pb.DeleteInstanceGroupManagerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroupManagers/" + req.GetInstanceGroupManager()
	name, err := s.parseZonalIGMName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.InstanceGroupManager{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

// Regional InstanceGroupManagers
type regionInstanceGroupManagers struct {
	*MockService
	pb.UnimplementedRegionInstanceGroupManagersServer
}

func (s *regionInstanceGroupManagers) Get(ctx context.Context, req *pb.GetRegionInstanceGroupManagerRequest) (*pb.InstanceGroupManager, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/instanceGroupManagers/" + req.GetInstanceGroupManager()
	name, err := s.parseRegionalIGMName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstanceGroupManager{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *regionInstanceGroupManagers) Insert(ctx context.Context, req *pb.InsertRegionInstanceGroupManagerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/instanceGroupManagers/" + req.GetInstanceGroupManagerResource().GetName()
	name, err := s.parseRegionalIGMName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetInstanceGroupManagerResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instanceGroupManager")
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))

	// Underlying InstanceGroup
	obj.InstanceGroup = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s/instanceGroups/%s", name.Project.ID, name.Region, *obj.Name)))

	// Fingerprint
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	// Status fields
	obj.Status = &pb.InstanceGroupManagerStatus{
		IsStable: PtrTo(true),
		VersionTarget: &pb.InstanceGroupManagerStatusVersionTarget{
			IsReached: PtrTo(true),
		},
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
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *regionInstanceGroupManagers) Patch(ctx context.Context, req *pb.PatchRegionInstanceGroupManagerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/instanceGroupManagers/" + req.GetInstanceGroupManager()
	name, err := s.parseRegionalIGMName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.InstanceGroupManager{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// We handle slices separately to avoid appending
	autoHealingPolicies := req.GetInstanceGroupManagerResource().GetAutoHealingPolicies()
	versions := req.GetInstanceGroupManagerResource().GetVersions()

	// Clone the update resource and clear slices to avoid append-merge
	patchSrc := proto.CloneOf(req.GetInstanceGroupManagerResource())
	patchSrc.AutoHealingPolicies = nil
	patchSrc.Versions = nil

	proto.Merge(obj, patchSrc)

	if autoHealingPolicies != nil {
		obj.AutoHealingPolicies = autoHealingPolicies
	}
	if versions != nil {
		obj.Versions = versions
	}

	// Update fingerprint on change
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *regionInstanceGroupManagers) Delete(ctx context.Context, req *pb.DeleteRegionInstanceGroupManagerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/instanceGroupManagers/" + req.GetInstanceGroupManager()
	name, err := s.parseRegionalIGMName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.InstanceGroupManager{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type zonalIGMName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalIGMName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/instanceGroupManagers/" + n.Name
}

// parseZonalIGMName parses a string into a zonalIGMName.
// The expected form is `projects/*/zones/*/instanceGroupManagers/*`.
func (s *MockService) parseZonalIGMName(name string) (*zonalIGMName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "instanceGroupManagers" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zonalIGMName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type regionalIGMName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalIGMName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/instanceGroupManagers/" + n.Name
}

// parseRegionalIGMName parses a string into a regionalIGMName.
// The expected form is `projects/*/regions/*/instanceGroupManagers/*`.
func (s *MockService) parseRegionalIGMName(name string) (*regionalIGMName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "instanceGroupManagers" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalIGMName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
