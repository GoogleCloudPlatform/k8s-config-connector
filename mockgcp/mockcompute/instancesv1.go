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
	"net"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
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
	// if obj.MachineType == nil {
	// 	machineType := "pd-standard"
	// 	obj.MachineType = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/machineTypes/%s", name.Project.ID, name.Zone, machineType)))
	// }

	// Resolve everything before writing any resources, so that a failed insert leaves nothing behind.
	region := regionFromZone(name.Zone)
	for i, nic := range obj.GetNetworkInterfaces() {
		if nic.GetNetworkIP() == "" || net.ParseIP(nic.GetNetworkIP()) != nil {
			continue
		}
		networkIP, err := s.resolveNetworkIP(ctx, name.Project.ID, region, i, nic.GetNetworkIP())
		if err != nil {
			return nil, err
		}
		nic.NetworkIP = PtrTo(networkIP)
	}

	var bootDisks []storedBootDisk
	for i, attached := range obj.GetDisks() {
		if !attached.GetBoot() || attached.GetInitializeParams() == nil || attached.GetSource() != "" || attached.GetType() == "SCRATCH" {
			continue
		}
		disk, diskFQN, err := s.buildBootDisk(ctx, name, obj.GetSelfLink(), attached, i)
		if err != nil {
			return nil, err
		}
		bootDisks = append(bootDisks, storedBootDisk{fqn: diskFQN, disk: disk})
	}

	for i, bootDisk := range bootDisks {
		if err := s.storage.Create(ctx, bootDisk.fqn, bootDisk.disk); err != nil {
			s.deleteBootDisks(ctx, bootDisks[:i])
			return nil, err
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		s.deleteBootDisks(ctx, bootDisks)
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

	s.releaseInstanceDisks(ctx, deleted)

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

// storedBootDisk is a disk created from the initializeParams of a boot disk.
type storedBootDisk struct {
	fqn  string
	disk *pb.Disk
}

// buildBootDisk builds the persistent disk described by the initializeParams of a boot disk, and rewrites the
// attached disk to reference it, as GCE does. It returns the disk and its storage key; nothing is stored.
func (s *InstancesV1) buildBootDisk(ctx context.Context, instance *zonalInstanceName, instanceSelfLink string, attached *pb.AttachedDisk, index int) (*pb.Disk, string, error) {
	params := attached.GetInitializeParams()
	projectID := instance.Project.ID

	diskName := params.GetDiskName()
	if diskName == "" {
		diskName = instance.Name
	}
	diskFQN := "projects/" + projectID + "/zones/" + instance.Zone + "/disks/" + diskName
	existing := &pb.Disk{}
	if err := s.storage.Get(ctx, diskFQN, existing); err == nil {
		return nil, "", status.Errorf(codes.AlreadyExists, "The resource '%s' already exists", diskFQN)
	} else if status.Code(err) != codes.NotFound {
		return nil, "", err
	}

	var image *pb.Image
	if params.GetSourceImage() != "" {
		var err error
		image, err = s.resolveSourceImage(ctx, projectID, params.GetSourceImage())
		if err != nil {
			return nil, "", err
		}
	}

	sizeGb := params.GetDiskSizeGb()
	if sizeGb == 0 {
		sizeGb = image.GetDiskSizeGb()
	}
	if sizeGb == 0 {
		sizeGb = 10
	}

	disk := &pb.Disk{
		Id:                     PtrTo(s.generateID()),
		Name:                   PtrTo(diskName),
		Kind:                   PtrTo("compute#disk"),
		SelfLink:               PtrTo(BuildComputeSelfLink(ctx, diskFQN)),
		CreationTimestamp:      PtrTo(s.nowString()),
		Zone:                   PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", projectID, instance.Zone))),
		Status:                 PtrTo("READY"),
		Type:                   PtrTo(expandDiskType(ctx, projectID, instance.Zone, params.GetDiskType())),
		SizeGb:                 PtrTo(sizeGb),
		Labels:                 params.GetLabels(),
		LabelFingerprint:       PtrTo(labelsFingerprint(params.GetLabels())),
		PhysicalBlockSizeBytes: PtrTo(int64(4096)),
		Users:                  []string{instanceSelfLink},
		DiskEncryptionKey:      attached.GetDiskEncryptionKey(),
	}
	if image != nil {
		disk.SourceImage = PtrTo(image.GetSelfLink())
		disk.SourceImageId = PtrTo(fmt.Sprintf("%d", image.GetId()))
	}

	attached.Source = PtrTo(BuildComputeSelfLink(ctx, diskFQN))
	attached.InitializeParams = nil
	attached.Kind = PtrTo("compute#attachedDisk")
	attached.Index = PtrTo(int32(index))
	if attached.Type == nil {
		attached.Type = PtrTo("PERSISTENT")
	}
	if attached.Mode == nil {
		attached.Mode = PtrTo("READ_WRITE")
	}
	if attached.Interface == nil {
		attached.Interface = PtrTo("SCSI")
	}
	if attached.DeviceName == nil {
		attached.DeviceName = PtrTo(fmt.Sprintf("persistent-disk-%d", index))
	}
	if attached.DiskSizeGb == nil {
		attached.DiskSizeGb = PtrTo(sizeGb)
	}

	return disk, diskFQN, nil
}

// deleteBootDisks removes disks created for an insert that failed.
func (s *InstancesV1) deleteBootDisks(ctx context.Context, bootDisks []storedBootDisk) {
	for _, bootDisk := range bootDisks {
		// Best effort: the disk is only orphaned if this fails.
		_ = s.storage.Delete(ctx, bootDisk.fqn, &pb.Disk{})
	}
}

// expandDiskType returns the full URL of a disk type. Bare names are resolved in the instance's zone.
func expandDiskType(ctx context.Context, projectID, zone, diskType string) string {
	if diskType == "" {
		diskType = "pd-standard"
	}
	if strings.HasPrefix(diskType, "http") {
		return diskType
	}
	path := strings.TrimPrefix(diskType, "/")
	switch {
	case strings.HasPrefix(path, "projects/"):
		return BuildComputeSelfLink(ctx, path)
	case strings.HasPrefix(path, "zones/"):
		return BuildComputeSelfLink(ctx, "projects/"+projectID+"/"+path)
	}
	return BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/diskTypes/%s", projectID, zone, path))
}

// resolveNetworkIP resolves a networkIP that refers to a regional address, such as
// projects/P/regions/R/addresses/A, to the IP address of that address. The address must be in the
// instance's project and region.
// The order of the checks below (path parsing, project lookup, address lookup, empty IP, cross-project,
// cross-region) and their error codes/messages are modeled on GCE but have not been verified against real
// GCP, which may report these cases in a different order or with different messages.
func (s *InstancesV1) resolveNetworkIP(ctx context.Context, instanceProject, instanceRegion string, index int, value string) (string, error) {
	field := fmt.Sprintf("resource.networkInterfaces[%d].networkIP", index)
	invalid := status.Errorf(codes.InvalidArgument, "Invalid value for field '%s': '%s'", field, value)

	tokens := strings.Split(strings.TrimPrefix(trimComputeURLPrefix(value), "/"), "/")
	var projectRef, region, addressName string
	switch {
	case len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "addresses":
		projectRef, region, addressName = tokens[1], tokens[3], tokens[5]
	case len(tokens) == 4 && tokens[0] == "regions" && tokens[2] == "addresses":
		projectRef, region, addressName = instanceProject, tokens[1], tokens[3]
	default:
		return "", invalid
	}

	project, err := s.Projects.GetProjectByIDOrNumber(projectRef)
	if err != nil {
		return "", err
	}

	fqn := "projects/" + project.ID + "/regions/" + region + "/addresses/" + addressName
	address := &pb.Address{}
	if err := s.storage.Get(ctx, fqn, address); err != nil {
		if status.Code(err) == codes.NotFound {
			return "", status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return "", err
	}

	if address.GetAddress() == "" {
		return "", invalid
	}

	if project.ID != instanceProject {
		return "", status.Errorf(codes.InvalidArgument, "Invalid value for field '%s': '%s'. IP address '%s' (%s) is reserved by another project.", field, value, fqn, address.GetAddress())
	}
	if region != instanceRegion {
		return "", status.Errorf(codes.InvalidArgument, "Invalid value for field '%s': '%s'. IP address '%s' (%s) is in region %s, but the instance is in region %s.", field, value, fqn, address.GetAddress(), region, instanceRegion)
	}
	return address.GetAddress(), nil
}

// releaseInstanceDisks is called once an instance is deleted. Disks the instance auto-deletes are removed;
// other disks no longer list the instance as a user. The instance is already gone, so failures are logged
// rather than returned: a retried delete would fail on the missing instance and never clean up.
func (s *InstancesV1) releaseInstanceDisks(ctx context.Context, instance *pb.Instance) {
	log := klog.FromContext(ctx)
	for _, attached := range instance.GetDisks() {
		diskFQN, ok := diskSourceFQN(attached.GetSource())
		if !ok {
			continue
		}
		disk := &pb.Disk{}
		if err := s.storage.Get(ctx, diskFQN, disk); err != nil {
			if status.Code(err) != codes.NotFound {
				log.Error(err, "failed to read disk of deleted instance", "disk", diskFQN)
			}
			continue
		}

		if attached.GetAutoDelete() {
			if err := s.storage.Delete(ctx, diskFQN, &pb.Disk{}); err != nil && status.Code(err) != codes.NotFound {
				log.Error(err, "failed to delete auto-delete disk of deleted instance", "disk", diskFQN)
			}
			continue
		}

		var users []string
		for _, user := range disk.GetUsers() {
			if user != instance.GetSelfLink() {
				users = append(users, user)
			}
		}
		if len(users) == len(disk.GetUsers()) {
			continue
		}
		disk.Users = users
		if err := s.storage.Update(ctx, diskFQN, disk); err != nil {
			log.Error(err, "failed to remove deleted instance from disk users", "disk", diskFQN)
		}
	}
}

// diskSourceFQN returns the storage key of the zonal disk referenced by a source URL or path.
func diskSourceFQN(source string) (string, bool) {
	if source == "" {
		return "", false
	}
	tokens := strings.Split(strings.TrimPrefix(trimComputeURLPrefix(source), "/"), "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "disks" {
		return "projects/" + tokens[1] + "/zones/" + tokens[3] + "/disks/" + tokens[5], true
	}
	return "", false
}

// regionFromZone returns the region containing a zone, e.g. us-central1 for us-central1-a.
func regionFromZone(zone string) string {
	if i := strings.LastIndex(zone, "-"); i != -1 {
		return zone[:i]
	}
	return zone
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
