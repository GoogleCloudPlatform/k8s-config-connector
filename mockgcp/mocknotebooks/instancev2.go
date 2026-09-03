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

package mocknotebooks

import (
	"context"
	"fmt"
	"sort"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"

	pb_v2 "cloud.google.com/go/notebooks/apiv2/notebookspb"
)

func (s *NotebookServiceV2) GetInstance(ctx context.Context, req *pb_v2.GetInstanceRequest) (*pb_v2.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb_v2.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *NotebookServiceV2) CreateInstance(ctx context.Context, req *pb_v2.CreateInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/instances/" + req.InstanceId
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Instance).(*pb_v2.Instance)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pb_v2.State_PROVISIONING
	obj.HealthState = pb_v2.HealthState_HEALTH_STATE_UNSPECIFIED
	obj.Creator = "someone@somewhere.com"
	obj.ProxyUri = fmt.Sprintf("https://%s-dot-us-central1.notebooks.googleusercontent.com", req.InstanceId)

	if obj.GetGceSetup() == nil {
		obj.Infrastructure = &pb_v2.Instance_GceSetup{
			GceSetup: &pb_v2.GceSetup{},
		}
	}
	gceSetup := obj.GetGceSetup()
	if gceSetup.MachineType == "" {
		gceSetup.MachineType = "n1-standard-1"
	}
	if gceSetup.ServiceAccounts == nil {
		gceSetup.ServiceAccounts = []*pb_v2.ServiceAccount{
			{
				Email:  fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number),
				Scopes: []string{"https://www.googleapis.com/auth/cloud-platform", "https://www.googleapis.com/auth/userinfo.email"},
			},
		}
	} else {
		for _, sa := range gceSetup.ServiceAccounts {
			if sa.Email == "" {
				sa.Email = fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number)
			}
			if sa.Scopes == nil {
				sa.Scopes = []string{"https://www.googleapis.com/auth/cloud-platform", "https://www.googleapis.com/auth/userinfo.email"}
			}
		}
	}
	if gceSetup.BootDisk == nil {
		gceSetup.BootDisk = &pb_v2.BootDisk{
			DiskSizeGb:     150,
			DiskType:       pb_v2.DiskType_PD_STANDARD,
			DiskEncryption: pb_v2.DiskEncryption_GMEK,
		}
	}
	if gceSetup.DataDisks == nil {
		gceSetup.DataDisks = []*pb_v2.DataDisk{
			{
				DiskSizeGb:     100,
				DiskType:       pb_v2.DiskType_PD_STANDARD,
				DiskEncryption: pb_v2.DiskEncryption_GMEK,
			},
		}
	}
	if gceSetup.ShieldedInstanceConfig == nil {
		gceSetup.ShieldedInstanceConfig = &pb_v2.ShieldedInstanceConfig{
			EnableIntegrityMonitoring: true,
			EnableVtpm:                true,
		}
	}
	if gceSetup.Tags == nil {
		gceSetup.Tags = []string{}
	}
	gceSetup.Tags = append(gceSetup.Tags, "deeplearning-vm", "notebook-instance")
	sort.Strings(gceSetup.Tags)

	if len(gceSetup.NetworkInterfaces) == 0 {
		gceSetup.NetworkInterfaces = []*pb_v2.NetworkInterface{
			{
				Network: fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/networks/default", name.Project.ID),
				Subnet:  fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/us-central1/subnetworks/default", name.Project.ID),
			},
		}
	} else {
		var networkInterfaces []*pb_v2.NetworkInterface
		for _, n := range gceSetup.NetworkInterfaces {
			n = &pb_v2.NetworkInterface{
				Network: fmt.Sprintf("https://www.googleapis.com/compute/v1/%s", n.Network),
				Subnet:  fmt.Sprintf("https://www.googleapis.com/compute/v1/%s", n.Subnet),
				NicType: pb_v2.NetworkInterface_GVNIC,
			}
			networkInterfaces = append(networkInterfaces, n)
		}
		gceSetup.NetworkInterfaces = networkInterfaces
	}
	if obj.Labels == nil {
		obj.Labels = make(map[string]string)
	}
	systemLabels := map[string]string{
		"consumer-project-id":     name.Project.ID,
		"consumer-project-number": fmt.Sprintf("%v", name.Project.Number),
		"notebooks-product":       "workbench-instances",
		"resource-name":           name.name,
	}
	for k, v := range systemLabels {
		obj.Labels[k] = v
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.region)
	metadata := &pb_v2.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "create",
		Endpoint:              "CreateInstance",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		obj.State = pb_v2.State_ACTIVE
		obj.HealthState = pb_v2.HealthState_HEALTHY
		return obj, nil
	})
}

func (s *NotebookServiceV2) UpdateInstance(ctx context.Context, req *pb_v2.UpdateInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.Instance.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb_v2.Instance{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb_v2.Instance)
	if req.UpdateMask != nil {
		allowedMutablePaths := map[string]bool{
			"labels":                                                         true,
			"disable_proxy_access":                                           true,
			"gce_setup.min_cpu_platform":                                     true,
			"gce_setup.metadata":                                             true,
			"gce_setup.machine_type":                                         true,
			"gce_setup.accelerator_configs":                                  true,
			"gce_setup.accelerator_configs.type":                             true,
			"gce_setup.accelerator_configs.core_count":                       true,
			"gce_setup.gpu_driver_config":                                    true,
			"gce_setup.gpu_driver_config.enable_gpu_driver":                  true,
			"gce_setup.gpu_driver_config.custom_gpu_driver_path":             true,
			"gce_setup.shielded_instance_config":                             true,
			"gce_setup.shielded_instance_config.enable_secure_boot":          true,
			"gce_setup.shielded_instance_config.enable_vtpm":                 true,
			"gce_setup.shielded_instance_config.enable_integrity_monitoring": true,
			"gce_setup.reservation_affinity":                                 true,
			"gce_setup.reservation_affinity.consume_reservation_type":        true,
			"gce_setup.reservation_affinity.key":                             true,
			"gce_setup.reservation_affinity.values":                          true,
			"gce_setup.tags":                                                 true,
			"gce_setup.container_image":                                      true,
			"gce_setup.container_image.repository":                           true,
			"gce_setup.container_image.tag":                                  true,
			"gce_setup.disable_public_ip":                                    true,
		}

		for _, path := range req.UpdateMask.Paths {
			if !allowedMutablePaths[path] {
				return nil, status.Errorf(codes.InvalidArgument, "field %q is immutable", path)
			}
		}

		requiresStopped := map[string]bool{
			"gce_setup.accelerator_configs":                                  true,
			"gce_setup.accelerator_configs.type":                             true,
			"gce_setup.accelerator_configs.core_count":                       true,
			"gce_setup.machine_type":                                         true,
			"gce_setup.shielded_instance_config":                             true,
			"gce_setup.shielded_instance_config.enable_secure_boot":          true,
			"gce_setup.shielded_instance_config.enable_vtpm":                 true,
			"gce_setup.shielded_instance_config.enable_integrity_monitoring": true,
			"gce_setup.reservation_affinity":                                 true,
			"gce_setup.reservation_affinity.consume_reservation_type":        true,
			"gce_setup.reservation_affinity.key":                             true,
			"gce_setup.reservation_affinity.values":                          true,
			"gce_setup.container_image":                                      true,
			"gce_setup.container_image.repository":                           true,
			"gce_setup.container_image.tag":                                  true,
			"gce_setup.disable_public_ip":                                    true,
			"disable_proxy_access":                                           true,
		}

		for _, path := range req.UpdateMask.Paths {
			if requiresStopped[path] && existing.State != pb_v2.State_STOPPED {
				return nil, status.Errorf(codes.FailedPrecondition, "instance in state %q must be stopped before updating one of the following: [accelerator_configs machine_type shielded_instance_config reservation_affinity container_image disable_public_ip disable_proxy_access]", existing.State)
			}
		}

		if err := fields.UpdateByFieldMask(updated, req.Instance, req.UpdateMask.Paths); err != nil {
			return nil, err
		}
	} else {
		proto.Merge(updated, req.Instance)
	}

	updated.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.region)
	metadata := &pb_v2.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
		Endpoint:              "UpdateInstance",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return updated, nil
	})
}

func (s *NotebookServiceV2) DeleteInstance(ctx context.Context, req *pb_v2.DeleteInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb_v2.Instance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.region)
	metadata := &pb_v2.OperationMetadata{
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
