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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type MachineImagesV1 struct {
	*MockService
	pb.UnimplementedMachineImagesServer
}

func (s *MachineImagesV1) Get(ctx context.Context, req *pb.GetMachineImageRequest) (*pb.MachineImage, error) {
	name, err := s.parseMachineImageName(req.GetProject(), req.GetMachineImage())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.MachineImage{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

type MachineImageName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *MachineImageName) String() string {
	return "projects/" + n.Project.ID + "/global/machineImages/" + n.Name
}

func (s *MockService) parseMachineImageName(projectName, name string) (*MachineImageName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &MachineImageName{
		Project: project,
		Name:    name,
	}, nil
}

func (s *MachineImagesV1) Insert(ctx context.Context, req *pb.InsertMachineImageRequest) (*pb.Operation, error) {
	if req.GetMachineImageResource() == nil || req.GetMachineImageResource().GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "machine_image_resource.name is required")
	}

	name, err := s.parseMachineImageName(req.GetProject(), req.GetMachineImageResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetMachineImageResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#machineImage")
	obj.Status = PtrTo("READY")
	obj.StorageLocations = []string{"us"}
	if obj.TotalStorageBytes == nil {
		obj.TotalStorageBytes = PtrTo(int64(733080832))
	}
	if obj.LabelFingerprint == nil {
		obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.GetLabels()))
	}

	sourceInstance := obj.GetSourceInstance()
	if sourceInstance == "" && req.SourceInstance != nil {
		sourceInstance = *req.SourceInstance
	}

	if sourceInstance != "" {
		instanceFQN := toRelativePath(sourceInstance)
		if idx := strings.Index(instanceFQN, "projects/"); idx != -1 {
			instanceFQN = instanceFQN[idx:]
		}
		obj.SourceInstance = PtrTo(BuildComputeSelfLink(ctx, instanceFQN))

		instanceObj := &pb.Instance{}
		if err := s.storage.Get(ctx, instanceFQN, instanceObj); err == nil {
			s.populateFromInstance(ctx, obj, instanceFQN, instanceObj)
		} else {
			s.populateDefaultInstanceProperties(ctx, obj, instanceFQN)
		}
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
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *MachineImagesV1) populateFromInstance(ctx context.Context, obj *pb.MachineImage, instanceFQN string, instance *pb.Instance) {
	// Parse instanceFQN: projects/<project>/zones/<zone>/instances/<name>
	parts := strings.Split(instanceFQN, "/")
	projectID := ""
	zone := ""
	instanceName := ""
	for i := 0; i < len(parts)-1; i++ {
		switch parts[i] {
		case "projects":
			projectID = parts[i+1]
		case "zones":
			zone = parts[i+1]
		case "instances":
			instanceName = parts[i+1]
		}
	}

	diskSource := BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/disks/%s", projectID, zone, instanceName))
	savedDiskSource := fmt.Sprintf("projects/%s/zones/%s/disks/%s", projectID, zone, instanceName)

	machineType := instance.GetMachineType()
	if strings.Contains(machineType, "/") {
		machineType = lastComponent(machineType)
	}
	if machineType == "" {
		machineType = "e2-medium"
	}

	guestOsFeatures := []*pb.GuestOsFeature{
		{Type: PtrTo("UEFI_COMPATIBLE")},
		{Type: PtrTo("VIRTIO_SCSI_MULTIQUEUE")},
		{Type: PtrTo("GVNIC")},
		{Type: PtrTo("SEV_CAPABLE")},
		{Type: PtrTo("SEV_LIVE_MIGRATABLE_V2")},
		{Type: PtrTo("SUSPEND_SAFE_FPR")},
	}

	instanceLicenses := []string{
		"https://www.googleapis.com/compute/beta/projects/debian-cloud/global/licenses/debian-12-bookworm",
	}
	sourceInstanceLicenses := []string{
		fmt.Sprintf("projects/%s/global/licenses/debian-12-bookworm", projectID),
	}

	instancePropertiesDisks := []*pb.AttachedDisk{
		{
			Architecture:    PtrTo("X86_64"),
			AutoDelete:      PtrTo(true),
			Boot:            PtrTo(true),
			DeviceName:      PtrTo("persistent-disk-0"),
			DiskSizeGb:      PtrTo(int64(10)),
			GuestOsFeatures: guestOsFeatures,
			Index:           PtrTo(int32(0)),
			Interface:       PtrTo("SCSI"),
			Kind:            PtrTo("compute#attachedDisk"),
			Licenses:        instanceLicenses,
			Mode:            PtrTo("READ_WRITE"),
			SavedState:      PtrTo("DISK_SAVED_STATE_UNSPECIFIED"),
			Source:          PtrTo(diskSource),
			Type:            PtrTo("PERSISTENT"),
		},
	}

	sourceInstanceDisks := []*pb.SavedAttachedDisk{
		{
			AutoDelete:         PtrTo(true),
			Boot:               PtrTo(true),
			DeviceName:         PtrTo("persistent-disk-0"),
			DiskSizeGb:         PtrTo(int64(10)),
			DiskType:           PtrTo("pd-standard"),
			GuestOsFeatures:    guestOsFeatures,
			Index:              PtrTo(int32(0)),
			Interface:          PtrTo("SCSI"),
			Kind:               PtrTo("compute#savedAttachedDisk"),
			Licenses:           sourceInstanceLicenses,
			Mode:               PtrTo("READ_WRITE"),
			Source:             PtrTo(diskSource),
			StorageBytes:       PtrTo(int64(733080832)),
			StorageBytesStatus: PtrTo("UP_TO_DATE"),
			Type:               PtrTo("PERSISTENT"),
		},
	}

	var networkInterfaces []*pb.NetworkInterface
	for _, nic := range instance.GetNetworkInterfaces() {
		enrichedNic := proto.Clone(nic).(*pb.NetworkInterface)
		if enrichedNic.Fingerprint == nil {
			enrichedNic.Fingerprint = PtrTo("P4gE13HPT-E=")
		}
		if enrichedNic.Kind == nil {
			enrichedNic.Kind = PtrTo("compute#networkInterface")
		}
		if enrichedNic.Name == nil {
			enrichedNic.Name = PtrTo("nic0")
		}
		if enrichedNic.NetworkIP == nil {
			enrichedNic.NetworkIP = PtrTo("10.128.0.34")
		}
		if enrichedNic.StackType == nil {
			enrichedNic.StackType = PtrTo("IPV4_ONLY")
		}
		if enrichedNic.Subnetwork == nil {
			enrichedNic.Subnetwork = PtrTo(BuildComputeSelfLink(ctx, "projects/"+projectID+"/regions/us-central1/subnetworks/default"))
		}
		if enrichedNic.Network != nil && !strings.HasPrefix(enrichedNic.GetNetwork(), "https://") && !strings.HasPrefix(enrichedNic.GetNetwork(), "projects/") {
			enrichedNic.Network = PtrTo(BuildComputeSelfLink(ctx, "projects/"+projectID+"/global/networks/"+enrichedNic.GetNetwork()))
		}
		networkInterfaces = append(networkInterfaces, enrichedNic)
	}
	if len(networkInterfaces) == 0 {
		networkInterfaces = []*pb.NetworkInterface{
			{
				Fingerprint: PtrTo("P4gE13HPT-E="),
				Kind:        PtrTo("compute#networkInterface"),
				Name:        PtrTo("nic0"),
				Network:     PtrTo(BuildComputeSelfLink(ctx, "projects/"+projectID+"/global/networks/default")),
				NetworkIP:   PtrTo("10.128.0.34"),
				StackType:   PtrTo("IPV4_ONLY"),
				Subnetwork:  PtrTo(BuildComputeSelfLink(ctx, "projects/"+projectID+"/regions/us-central1/subnetworks/default")),
			},
		}
	}

	scheduling := instance.GetScheduling()
	if scheduling == nil {
		scheduling = &pb.Scheduling{}
	} else {
		scheduling = proto.Clone(scheduling).(*pb.Scheduling)
	}
	if scheduling.AutomaticRestart == nil {
		scheduling.AutomaticRestart = PtrTo(true)
	}
	if scheduling.OnHostMaintenance == nil {
		scheduling.OnHostMaintenance = PtrTo("MIGRATE")
	}
	if scheduling.Preemptible == nil {
		scheduling.Preemptible = PtrTo(false)
	}
	if scheduling.ProvisioningModel == nil {
		scheduling.ProvisioningModel = PtrTo("STANDARD")
	}

	tags := instance.GetTags()
	if tags == nil || tags.Fingerprint == nil {
		tags = &pb.Tags{
			Fingerprint: PtrTo("42WmSpB8rSM="),
		}
	}

	metadata := instance.GetMetadata()
	if metadata == nil {
		metadata = &pb.Metadata{
			Fingerprint: PtrTo("x4iOk7cA2GY="),
			Kind:        PtrTo("compute#metadata"),
		}
	} else {
		metadata = proto.Clone(metadata).(*pb.Metadata)
		if metadata.Fingerprint == nil {
			metadata.Fingerprint = PtrTo("x4iOk7cA2GY=")
		}
		if metadata.Kind == nil {
			metadata.Kind = PtrTo("compute#metadata")
		}
	}

	obj.InstanceProperties = &pb.InstanceProperties{
		CanIpForward:      instance.CanIpForward,
		Description:       PtrTo(""),
		Disks:             instancePropertiesDisks,
		MachineType:       PtrTo(machineType),
		Metadata:          metadata,
		MinCpuPlatform:    PtrTo(""),
		NetworkInterfaces: networkInterfaces,
		Scheduling:        scheduling,
		Tags:              tags,
	}

	obj.SourceInstanceProperties = &pb.SourceInstanceProperties{
		CanIpForward:       instance.CanIpForward,
		DeletionProtection: PtrTo(false),
		Disks:              sourceInstanceDisks,
		MachineType:        PtrTo(machineType),
		Metadata:           metadata,
		NetworkInterfaces:  networkInterfaces,
		Scheduling:         scheduling,
		Tags:               tags,
	}

	obj.SavedDisks = []*pb.SavedDisk{
		{
			Kind:               PtrTo("compute#savedDisk"),
			SourceDisk:         PtrTo(savedDiskSource),
			StorageBytes:       PtrTo(int64(733080832)),
			StorageBytesStatus: PtrTo("UP_TO_DATE"),
		},
	}
}

func (s *MachineImagesV1) populateDefaultInstanceProperties(ctx context.Context, obj *pb.MachineImage, instanceFQN string) {
	parts := strings.Split(instanceFQN, "/")
	projectID := ""
	zone := "us-central1-a"
	instanceName := ""
	for i := 0; i < len(parts)-1; i++ {
		switch parts[i] {
		case "projects":
			projectID = parts[i+1]
		case "zones":
			zone = parts[i+1]
		case "instances":
			instanceName = parts[i+1]
		}
	}

	fakeInstance := &pb.Instance{
		CanIpForward: PtrTo(false),
		MachineType:  PtrTo("e2-medium"),
	}
	s.populateFromInstance(ctx, obj, fmt.Sprintf("projects/%s/zones/%s/instances/%s", projectID, zone, instanceName), fakeInstance)
}

func (s *MachineImagesV1) Delete(ctx context.Context, req *pb.DeleteMachineImageRequest) (*pb.Operation, error) {
	name, err := s.parseMachineImageName(req.GetProject(), req.GetMachineImage())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.MachineImage{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *MachineImagesV1) List(ctx context.Context, req *pb.ListMachineImagesRequest) (*pb.MachineImageList, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	prefix := "projects/" + project.ID + "/global/machineImages/"
	var items []*pb.MachineImage
	kind := (&pb.MachineImage{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		items = append(items, obj.(*pb.MachineImage))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.MachineImageList{
		Id:       PtrTo(fmt.Sprintf("%d", s.generateID())),
		Items:    items,
		Kind:     PtrTo("compute#machineImageList"),
		SelfLink: PtrTo(BuildComputeSelfLink(ctx, "projects/"+project.ID+"/global/machineImages")),
	}, nil
}

func (s *MachineImagesV1) SetLabels(ctx context.Context, req *pb.SetLabelsMachineImageRequest) (*pb.Operation, error) {
	name, err := s.parseMachineImageName(req.GetProject(), req.GetResource())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.MachineImage{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if obj.GetLabelFingerprint() != req.GetGlobalSetLabelsRequestResource().GetLabelFingerprint() {
		return nil, status.Errorf(codes.FailedPrecondition, "LabelFingerprint mismatch")
	}

	obj.Labels = req.GetGlobalSetLabelsRequestResource().GetLabels()
	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.GetLabels()))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setLabels"),
		User:          PtrTo("user@example.com"),
		Status:        PtrTo(pb.Operation_DONE),
		Progress:      PtrTo(int32(100)),
		EndTime:       PtrTo(s.nowString()),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}
