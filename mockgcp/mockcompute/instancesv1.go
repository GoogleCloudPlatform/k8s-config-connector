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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type InstancesV1 struct {
	*MockService
	pb.UnimplementedInstancesServer
}

func (s *InstancesV1) Get(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *InstancesV1) Insert(ctx context.Context, req *pb.InsertInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstanceResource().GetName()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetInstanceResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instance")
	obj.Zone = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))
	obj.Status = PtrTo("RUNNING")
	if obj.LabelFingerprint == nil {
		obj.LabelFingerprint = PtrTo(computeFingerprint(obj))
	}
	for i, disk := range obj.Disks {
		if disk.Source == nil {
			disk.Source = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/disks/%s", name.Project.ID, name.Zone, obj.GetName())))
		}
		if i == 0 {
			disk.Boot = PtrTo(true)
		}
		// Auto-create the disk in storage if it doesn't already exist
		diskFQN := fmt.Sprintf("projects/%s/zones/%s/disks/%s", name.Project.ID, name.Zone, obj.GetName())
		diskID := s.generateID()

		var sourceImage *string
		var sourceImageID *string
		diskSize := int64(10)
		diskType := BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/diskTypes/pd-standard", name.Project.ID, name.Zone))
		var diskLabels map[string]string

		if disk.InitializeParams != nil {
			if disk.InitializeParams.SourceImage != nil {
				img := disk.InitializeParams.SourceImage
				tokens := strings.Split(*img, "/")
				if len(tokens) == 2 {
					sourceImage = PtrTo(BuildComputeSelfLink(ctx, "projects/debian-cloud/global/images/debian-11-bullseye-v20231010"))
					sourceImageID = PtrTo("2443108620951880213")
				} else if len(tokens) == 6 {
					sourceImage = PtrTo(BuildComputeSelfLink(ctx, "projects/debian-cloud/global/images/debian-11-bullseye-v20231010"))
					sourceImageID = PtrTo("2443108620951880213")
				} else {
					sourceImage = img
				}
			}
			if disk.InitializeParams.DiskSizeGb != nil {
				diskSize = *disk.InitializeParams.DiskSizeGb
			}
			if disk.InitializeParams.DiskType != nil {
				dt := disk.InitializeParams.DiskType
				if !strings.HasPrefix(*dt, "https://") && !strings.Contains(*dt, "/") {
					diskType = BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/diskTypes/%s", name.Project.ID, name.Zone, *dt))
				} else {
					diskType = *dt
				}
			}
			if disk.InitializeParams.Labels != nil {
				diskLabels = disk.InitializeParams.Labels
			}
		}

		mockDisk := &pb.Disk{
			Id:                     &diskID,
			Kind:                   PtrTo("compute#disk"),
			Name:                   PtrTo(obj.GetName()),
			SelfLink:               PtrTo(BuildComputeSelfLink(ctx, diskFQN)),
			CreationTimestamp:      PtrTo(s.nowString()),
			Zone:                   PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone))),
			Status:                 PtrTo("READY"),
			Type:                   PtrTo(diskType),
			PhysicalBlockSizeBytes: PtrTo(int64(4096)),
			SizeGb:                 PtrTo(diskSize),
			SourceImage:            sourceImage,
			SourceImageId:          sourceImageID,
			Labels:                 diskLabels,
		}
		_ = s.storage.Create(ctx, diskFQN, mockDisk)
		// InitializeParams is only used during insertion and is not preserved in the persistent GET response on real GCP
		disk.InitializeParams = nil
	}
	// if obj.MachineType == nil {
	// 	machineType := "pd-standard"
	// 	obj.MachineType = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/machineTypes/%s", name.Project.ID, name.Zone, machineType)))
	// }

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

// Updates a Instance resource in the specified project using the data included in the request.
func (s *InstancesV1) Update(ctx context.Context, req *pb.UpdateInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetInstanceResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstancesV1) AttachDisk(ctx context.Context, req *pb.AttachDiskInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	attachedDisk := proto.CloneOf(req.GetAttachedDiskResource())
	attachedDisk.Kind = PtrTo("compute#attachedDisk")
	attachedDisk.Index = PtrTo(int32(len(obj.Disks)))

	obj.Disks = append(obj.Disks, attachedDisk)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("attachDisk"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstancesV1) DetachDisk(ctx context.Context, req *pb.DetachDiskInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	var keepDisks []*pb.AttachedDisk
	for _, disk := range obj.Disks {
		if disk.GetDeviceName() != req.GetDeviceName() {
			keepDisks = append(keepDisks, disk)
		}
	}
	obj.Disks = keepDisks

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("detachDisk"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstancesV1) Stop(ctx context.Context, req *pb.StopInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Status = PtrTo("TERMINATED")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("stop"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstancesV1) Start(ctx context.Context, req *pb.StartInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Status = PtrTo("RUNNING")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("start"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstancesV1) SetServiceAccount(ctx context.Context, req *pb.SetServiceAccountInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.ServiceAccounts = []*pb.ServiceAccount{
		{
			Email:  req.GetInstancesSetServiceAccountRequestResource().Email,
			Scopes: req.GetInstancesSetServiceAccountRequestResource().Scopes,
		},
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating instance: %v", err)
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setServiceAccount"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstancesV1) SetMachineType(ctx context.Context, req *pb.SetMachineTypeInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.MachineType = req.GetInstancesSetMachineTypeRequestResource().MachineType

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setMachineType"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstancesV1) Delete(ctx context.Context, req *pb.DeleteInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Clean up auto-created disk
	diskFQN := fmt.Sprintf("projects/%s/zones/%s/disks/%s", name.Project.ID, name.Zone, name.Name)
	_ = s.storage.Delete(ctx, diskFQN, &pb.Disk{})

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

func (s *InstancesV1) SetLabels(ctx context.Context, req *pb.SetLabelsInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Labels = req.GetInstancesSetLabelsRequestResource().GetLabels()

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

func (s *InstancesV1) SetMetadata(ctx context.Context, req *pb.SetMetadataInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Metadata = req.GetMetadataResource()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setMetadata"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstancesV1) SetTags(ctx context.Context, req *pb.SetTagsInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Tags = req.GetTagsResource()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setTags"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstancesV1) SetScheduling(ctx context.Context, req *pb.SetSchedulingInstanceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instances/" + req.GetInstance()
	name, err := s.parseZonalInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Scheduling = req.GetSchedulingResource()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setScheduling"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type zonalInstanceName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalInstanceName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/instances/" + n.Name
}

// parseZonalInstanceName parses a string into a zonalInstanceName.
// The expected form is `projects/*/zones/*/instance/*`.
func (s *MockService) parseZonalInstanceName(name string) (*zonalInstanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zonalInstanceName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
