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
// proto.service: google.cloud.notebooks.v1.NotebookService
// proto.message: google.cloud.notebooks.v1.Instance

package mocknotebooks

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *NotebookServiceV1) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// createBootDisk creates the boot disk configuration
func createBootDisk(projectID string, instanceName string) *pb.Instance_Disk {
	return &pb.Instance_Disk{
		AutoDelete: true,
		Boot:       true,
		DeviceName: "boot",
		DiskSizeGb: 150,
		GuestOsFeatures: []*pb.Instance_Disk_GuestOsFeature{
			{Type: "VIRTIO_SCSI_MULTIQUEUE"},
			{Type: "UEFI_COMPATIBLE"},
			{Type: "GVNIC"},
		},
		Interface: "SCSI",
		Kind:      "compute#attachedDisk",
		Licenses: []string{
			"https://www.googleapis.com/compute/v1/projects/click-to-deploy-images/global/licenses/c2d-tensorflow",
			"https://www.googleapis.com/compute/v1/projects/click-to-deploy-images/global/licenses/c2d-dl-platform-gvnic",
			"https://www.googleapis.com/compute/v1/projects/click-to-deploy-images/global/licenses/c2d-dl-platform-cpu-common",
			"https://www.googleapis.com/compute/v1/projects/click-to-deploy-images/global/licenses/c2d-dl-platform-debian-11",
			"https://www.googleapis.com/compute/v1/projects/click-to-deploy-images/global/licenses/c2d-dl-platform-dlvm",
		},
		Mode:   "READ_WRITE",
		Source: fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/us-central1-a/disks/%s-boot", projectID, instanceName),
		Type:   "PERSISTENT",
	}
}

// createDataDisk creates the data disk configuration
func createDataDisk(projectID string, instanceName string) *pb.Instance_Disk {
	return &pb.Instance_Disk{
		AutoDelete: true,
		DeviceName: "data",
		DiskSizeGb: 100,
		Index:      1,
		Interface:  "SCSI",
		Kind:       "compute#attachedDisk",
		Mode:       "READ_WRITE",
		Source:     fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/us-central1-a/disks/%s-data", projectID, instanceName),
		Type:       "PERSISTENT",
	}
}

func (s *NotebookServiceV1) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/instances/" + req.InstanceId
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Instance).(*pb.Instance)
	obj.Name = fqn
	obj.Environment = req.Instance.Environment
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pb.Instance_PROVISIONING
	s.setDefaultServiceAccount(obj, name)

	// Set additional required fields
	obj.Creator = "someone@somewhere.com"
	obj.DiskEncryption = pb.Instance_GMEK
	obj.AcceleratorConfig = nil
	obj.Disks = []*pb.Instance_Disk{
		createBootDisk(name.Project.ID, req.InstanceId),
		createDataDisk(name.Project.ID, req.InstanceId),
	}
	// Cant find the field InstanceMigrationEligibility in the proto file
	// for     "instanceMigrationEligibility": {},
	//obj.InstanceMigrationEligibility = &pb.Instance_InstanceMigrationEligibility{}
	obj.Labels = map[string]string{
		"goog-caip-notebook": "",
	}
	obj.MachineType = fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/us-central1-a/machineTypes/n1-standard-1", name.Project.ID)
	obj.Metadata = map[string]string{
		"container":                  "gcr.io/deeplearning-platform-release/base-cpu",
		"disable-swap-binaries":      "true",
		"enable-guest-attributes":    "TRUE",
		"notebooks-api":              "PROD",
		"notebooks-api-version":      "v1",
		"proxy-mode":                 "service_account",
		"serial-port-logging-enable": "true",
		"shutdown-script":            "/opt/deeplearning/bin/shutdown_script.sh",
		"warmup-libraries":           "matplotlib.pyplot",
	}
	//  cant find the field Migrated in the proto file
	// for     "migrated": false,
	//obj.Migrated = false
	obj.Network = fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/networks/default", name.Project.ID)
	obj.ServiceAccountScopes = []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/userinfo.email",
	}
	obj.ShieldedInstanceConfig = &pb.Instance_ShieldedInstanceConfig{
		EnableIntegrityMonitoring: true,
		EnableVtpm:                true,
	}
	obj.Subnet = fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/us-central1/subnetworks/default", name.Project.ID)
	obj.Tags = []string{
		"deeplearning-vm",
		"notebook-instance",
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.region)
	metadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "create",
		Endpoint:              "CreateInstance",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NotebookServiceV1) UpdateInstanceMetadataItems(ctx context.Context, req *pb.UpdateInstanceMetadataItemsRequest) (*pb.UpdateInstanceMetadataItemsResponse, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	updated := &pb.Instance{}
	updated.Metadata = req.GetItems()
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	obj := &pb.UpdateInstanceMetadataItemsResponse{}
	obj.Items = req.Items
	return obj, nil
}

func (s *NotebookServiceV1) UpdateShieldedInstanceConfig(ctx context.Context, req *pb.UpdateShieldedInstanceConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	updated := &pb.Instance{}
	updated.ShieldedInstanceConfig = req.GetShieldedInstanceConfig()
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.region)
	metadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
		Endpoint:              "UpdateShieldedInstanceConfig",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return updated, nil
	})
}

func (s *NotebookServiceV1) setDefaultServiceAccount(obj *pb.Instance, name *instanceName) {
	if obj.ServiceAccount == "" {
		obj.ServiceAccount = fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number)
	}
}

func (s *NotebookServiceV1) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.region)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		Endpoint:              "DeleteInstance",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type instanceName struct {
	Project *projects.ProjectData
	region  string
	name    string
}

func (n *instanceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", n.Project.ID, n.region, n.name)
}

// parseInstanceName parses a string into an instanceName.
// The expected form is `projects/*/locations/*/instances/*`.
func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &instanceName{
			Project: project,
			region:  tokens[3],
			name:    tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
