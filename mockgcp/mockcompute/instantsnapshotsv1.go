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

type InstantSnapshotsV1 struct {
	*MockService
	pb.UnimplementedInstantSnapshotsServer
}

type RegionInstantSnapshotsV1 struct {
	*MockService
	pb.UnimplementedRegionInstantSnapshotsServer
}

func (s *InstantSnapshotsV1) Get(ctx context.Context, req *pb.GetInstantSnapshotRequest) (*pb.InstantSnapshot, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instantSnapshots/" + req.GetInstantSnapshot()
	name, err := s.parseZonalInstantSnapshotName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstantSnapshot{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *InstantSnapshotsV1) Insert(ctx context.Context, req *pb.InsertInstantSnapshotRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instantSnapshots/" + req.GetInstantSnapshotResource().GetName()
	name, err := s.parseZonalInstantSnapshotName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetInstantSnapshotResource())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instantSnapshot")
	obj.Status = PtrTo("READY")
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.SelfLinkWithId = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.Zone = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))
	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.Labels))
	obj.SatisfiesPzi = PtrTo(true)
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.ResourceStatus == nil {
		obj.ResourceStatus = &pb.InstantSnapshotResourceStatus{
			StorageSizeBytes: PtrTo(int64(0)),
		}
	}

	s.populateSourceDiskInfo(ctx, obj)

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

func (s *InstantSnapshotsV1) Delete(ctx context.Context, req *pb.DeleteInstantSnapshotRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instantSnapshots/" + req.GetInstantSnapshot()
	name, err := s.parseZonalInstantSnapshotName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstantSnapshot{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	deleted := &pb.InstantSnapshot{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *InstantSnapshotsV1) SetLabels(ctx context.Context, req *pb.SetLabelsInstantSnapshotRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instantSnapshots/" + req.GetResource()
	name, err := s.parseZonalInstantSnapshotName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.InstantSnapshot{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	rb := req.GetZoneSetLabelsRequestResource()
	if rb.GetLabelFingerprint() != "" && rb.GetLabelFingerprint() != obj.GetLabelFingerprint() {
		return nil, status.Errorf(codes.FailedPrecondition, "Labels fingerprint does not match")
	}

	obj.Labels = rb.GetLabels()
	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.Labels))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setLabels"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionInstantSnapshotsV1) Get(ctx context.Context, req *pb.GetRegionInstantSnapshotRequest) (*pb.InstantSnapshot, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/instantSnapshots/" + req.GetInstantSnapshot()
	name, err := s.parseRegionalInstantSnapshotName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstantSnapshot{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegionInstantSnapshotsV1) Insert(ctx context.Context, req *pb.InsertRegionInstantSnapshotRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/instantSnapshots/" + req.GetInstantSnapshotResource().GetName()
	name, err := s.parseRegionalInstantSnapshotName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetInstantSnapshotResource())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instantSnapshot")
	obj.Status = PtrTo("READY")
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.SelfLinkWithId = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))
	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.Labels))
	obj.SatisfiesPzi = PtrTo(true)
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.ResourceStatus == nil {
		obj.ResourceStatus = &pb.InstantSnapshotResourceStatus{
			StorageSizeBytes: PtrTo(int64(0)),
		}
	}

	s.populateSourceDiskInfo(ctx, obj)

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

func (s *RegionInstantSnapshotsV1) Delete(ctx context.Context, req *pb.DeleteRegionInstantSnapshotRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/instantSnapshots/" + req.GetInstantSnapshot()
	name, err := s.parseRegionalInstantSnapshotName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstantSnapshot{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	deleted := &pb.InstantSnapshot{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *RegionInstantSnapshotsV1) SetLabels(ctx context.Context, req *pb.SetLabelsRegionInstantSnapshotRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/instantSnapshots/" + req.GetResource()
	name, err := s.parseRegionalInstantSnapshotName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.InstantSnapshot{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	rb := req.GetRegionSetLabelsRequestResource()
	if rb.GetLabelFingerprint() != "" && rb.GetLabelFingerprint() != obj.GetLabelFingerprint() {
		return nil, status.Errorf(codes.FailedPrecondition, "Labels fingerprint does not match")
	}

	obj.Labels = rb.GetLabels()
	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.Labels))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setLabels"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *MockService) populateSourceDiskInfo(ctx context.Context, obj *pb.InstantSnapshot) {
	sourceDisk := obj.GetSourceDisk()
	if sourceDisk == "" {
		return
	}

	u := sourceDisk
	if idx := strings.Index(u, "projects/"); idx != -1 {
		u = u[idx:]
	}

	if diskName, err := s.parseZonalDiskName(u); err == nil {
		diskFQN := diskName.String()
		diskObj := &pb.Disk{}
		if err := s.storage.Get(ctx, diskFQN, diskObj); err == nil {
			obj.SourceDiskId = PtrTo(fmt.Sprintf("%d", diskObj.GetId()))
			if diskObj.SizeGb != nil {
				obj.DiskSizeGb = diskObj.SizeGb
			}
		}
	} else if regDiskName, err := s.parseZonalRegionDiskName(u); err == nil {
		diskFQN := regDiskName.String()
		diskObj := &pb.Disk{}
		if err := s.storage.Get(ctx, diskFQN, diskObj); err == nil {
			obj.SourceDiskId = PtrTo(fmt.Sprintf("%d", diskObj.GetId()))
			if diskObj.SizeGb != nil {
				obj.DiskSizeGb = diskObj.SizeGb
			}
		}
	}

	if obj.DiskSizeGb == nil {
		obj.DiskSizeGb = PtrTo(int64(10))
	}
	if obj.SourceDiskId == nil {
		obj.SourceDiskId = PtrTo("1234567890123456789")
	}
}

type zonalInstantSnapshotName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalInstantSnapshotName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/instantSnapshots/" + n.Name
}

func (s *MockService) parseZonalInstantSnapshotName(name string) (*zonalInstantSnapshotName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "instantSnapshots" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zonalInstantSnapshotName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type regionalInstantSnapshotName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalInstantSnapshotName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/instantSnapshots/" + n.Name
}

func (s *MockService) parseRegionalInstantSnapshotName(name string) (*regionalInstantSnapshotName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "instantSnapshots" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalInstantSnapshotName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
