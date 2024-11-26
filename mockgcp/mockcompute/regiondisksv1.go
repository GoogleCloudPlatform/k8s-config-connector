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

	obj := proto.Clone(req.GetDiskResource()).(*pb.Disk)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#disk")
	obj.Region = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))
	obj.Status = PtrTo("READY")
	if obj.Type == nil {
		diskType := "pd-standard"
		obj.Type = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s/regiondiskTypes/%s", name.Project.ID, name.Region, diskType)))
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
