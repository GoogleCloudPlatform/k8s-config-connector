// Copyright 2024 Google LLC
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

type RegionalDisksV1 struct {
	*MockService
	pb.UnimplementedRegionDisksServer
}

func (s *RegionalDisksV1) Get(ctx context.Context, req *pb.GetRegionDiskRequest) (*pb.Disk, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/disks/" + req.GetDisk()
	name, err := s.parseZonalRegionDiskName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Disk{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *RegionalDisksV1) Insert(ctx context.Context, req *pb.InsertRegionDiskRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/disks/" + req.GetDiskResource().GetName()
	name, err := s.parseZonalRegionDiskName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetDiskResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#disk")
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))
	obj.Status = PtrTo("READY")
	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.Labels))
	if obj.Type == nil {
		diskType := "pd-standard"
		obj.Type = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s/regiondiskTypes/%s", name.Project.ID, name.Region, diskType)))
	}
	if obj.PhysicalBlockSizeBytes == nil {
		obj.PhysicalBlockSizeBytes = PtrTo(int64(4096))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

// Updates a disk resource in the specified project using the data included in the request.
func (s *RegionalDisksV1) Update(ctx context.Context, req *pb.UpdateRegionDiskRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/disks/" + req.GetDisk()
	name, err := s.parseZonalRegionDiskName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Disk{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetDiskResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalDisksV1) Delete(ctx context.Context, req *pb.DeleteRegionDiskRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/disks/" + req.GetDisk()
	name, err := s.parseZonalRegionDiskName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Disk{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalDisksV1) Resize(ctx context.Context, req *pb.ResizeRegionDiskRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/disks/" + req.GetDisk()
	name, err := s.parseZonalRegionDiskName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Disk{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	sizeGb := req.GetRegionDisksResizeRequestResource().GetSizeGb()
	obj.SizeGb = &sizeGb

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalDisksV1) SetLabels(ctx context.Context, req *pb.SetLabelsRegionDiskRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/disks/" + req.GetResource()
	name, err := s.parseZonalRegionDiskName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Disk{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Labels = req.GetRegionSetLabelsRequestResource().GetLabels()
	obj.LabelFingerprint = req.GetRegionSetLabelsRequestResource().LabelFingerprint

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalDisksV1) AddResourcePolicies(ctx context.Context, req *pb.AddResourcePoliciesRegionDiskRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/disks/" + req.GetDisk()
	name, err := s.parseZonalRegionDiskName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Disk{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	policiesToAdd := req.GetRegionDisksAddResourcePoliciesRequestResource().GetResourcePolicies()
	for _, policy := range policiesToAdd {
		found := false
		for _, existing := range obj.ResourcePolicies {
			if existing == policy {
				found = true
				break
			}
		}
		if !found {
			obj.ResourcePolicies = append(obj.ResourcePolicies, policy)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("addResourcePolicies"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionalDisksV1) RemoveResourcePolicies(ctx context.Context, req *pb.RemoveResourcePoliciesRegionDiskRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/disks/" + req.GetDisk()
	name, err := s.parseZonalRegionDiskName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Disk{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	policiesToRemove := req.GetRegionDisksRemoveResourcePoliciesRequestResource().GetResourcePolicies()
	var newPolicies []string
	for _, existing := range obj.ResourcePolicies {
		toRemove := false
		for _, policy := range policiesToRemove {
			if existing == policy {
				toRemove = true
				break
			}
		}
		if !toRemove {
			newPolicies = append(newPolicies, existing)
		}
	}
	obj.ResourcePolicies = newPolicies

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("removeResourcePolicies"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type regionalDiskName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalDiskName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/disks/" + n.Name
}

// parseZonalRegionDiskName parses a string into a regionalDiskName.
// The expected form is `projects/*/regions/*/disks/*`.
func (s *MockService) parseZonalRegionDiskName(name string) (*regionalDiskName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "disks" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalDiskName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
