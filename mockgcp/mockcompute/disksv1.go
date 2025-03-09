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

// +tool:mockgcp-support
// proto.service: google.cloud.compute.v1.Disks
// proto.message: google.cloud.compute.v1.Disk

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type DisksV1 struct {
	*MockService
	pb.UnimplementedDisksServer
}

func (s *DisksV1) Get(ctx context.Context, req *pb.GetDiskRequest) (*pb.Disk, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/disks/" + req.GetDisk()
	name, err := s.parseZonalDiskName(reqName)
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

func (s *DisksV1) Insert(ctx context.Context, req *pb.InsertDiskRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/disks/" + req.GetDiskResource().GetName()
	name, err := s.parseZonalDiskName(reqName)
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
	obj.Zone = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))
	obj.Status = PtrTo("READY")
	if obj.Type == nil {
		diskType := "pd-standard"
		obj.Type = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/diskTypes/%s", name.Project.ID, name.Zone, diskType)))
	}
	if obj.PhysicalBlockSizeBytes == nil {
		obj.PhysicalBlockSizeBytes = PtrTo(int64(4096))
	}
	if obj.EnableConfidentialCompute == nil {
		obj.EnableConfidentialCompute = PtrTo(false)
	}
	if obj.SourceImage != nil {
		tokens := strings.Split(*obj.SourceImage, "/")
		if len(tokens) == 2 {
			// debian-cloud/debian-11
			obj.SourceImage = PtrTo(buildComputeSelfLink(ctx, "projects/debian-cloud/global/images/debian-11-bullseye-v20231010"))
			obj.SourceImageId = PtrTo("2443108620951880213")
		}
		if len(tokens) == 6 {
			// projects/debian-cloud/global/images/family/debian-11
			obj.SourceImage = PtrTo(buildComputeSelfLink(ctx, "projects/debian-cloud/global/images/debian-11-bullseye-v20231010"))
			obj.SourceImageId = PtrTo("2443108620951880213")
		}
	}
	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.GetLabels()))
	obj.SatisfiesPzi = PtrTo(true)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

// Updates a Disk resource in the specified project using the data included in the request.
func (s *DisksV1) Update(ctx context.Context, req *pb.UpdateDiskRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/disks/" + req.GetDisk()
	name, err := s.parseZonalDiskName(reqName)
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

	op := &pb.Operation{
		OperationType: PtrTo("update"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *DisksV1) Delete(ctx context.Context, req *pb.DeleteDiskRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/disks/" + req.GetDisk()
	name, err := s.parseZonalDiskName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := (&pb.Disk{})
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type zonalDiskName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalDiskName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/disks/" + n.Name
}

// parseZonalDiskName parses a string into a zonalDiskName.
// The expected form is `projects/*/zones/*/disk/*`.
func (s *MockService) parseZonalDiskName(name string) (*zonalDiskName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "disks" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zonalDiskName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
