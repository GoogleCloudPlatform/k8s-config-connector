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
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

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
	obj.Name = fqn
	obj.Creator = "someone@somewhere.com"
	obj.HealthState = pb_v2.HealthState_HEALTHY
	obj.State = pb_v2.State_ACTIVE
	obj.ProxyUri = fmt.Sprintf("https://%s-dot-us-central1.notebooks.googleusercontent.com", name.name)

	if obj.GetGceSetup() == nil {
		obj.Infrastructure = &pb_v2.Instance_GceSetup{
			GceSetup: &pb_v2.GceSetup{},
		}
	}
	gceSetup := obj.GetGceSetup()
	if gceSetup.MachineType == "" {
		gceSetup.MachineType = "n1-standard-1"
	}
	if len(gceSetup.ServiceAccounts) == 0 {
		gceSetup.ServiceAccounts = []*pb_v2.ServiceAccount{
			{
				Email: fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number),
			},
		}
	}
	if gceSetup.BootDisk == nil {
		gceSetup.BootDisk = &pb_v2.BootDisk{
			DiskSizeGb:     150,
			DiskType:       pb_v2.DiskType_PD_STANDARD,
			DiskEncryption: pb_v2.DiskEncryption_GMEK,
		}
	}
	if len(gceSetup.NetworkInterfaces) == 0 {
		gceSetup.NetworkInterfaces = []*pb_v2.NetworkInterface{
			{
				Network: fmt.Sprintf("projects/%s/global/networks/default", name.Project.ID),
				Subnet:  fmt.Sprintf("projects/%s/regions/us-central1/subnetworks/default", name.Project.ID),
			},
		}
	}

	if obj.CreateTime == nil {
		obj.CreateTime = timestamppb.New(time.Now())
	}
	if obj.UpdateTime == nil {
		obj.UpdateTime = timestamppb.New(time.Now())
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
		if req.Instance.Labels != nil {
			updated.Labels = req.Instance.Labels
		}
		if req.Instance.GetGceSetup() != nil {
			if req.Instance.GetGceSetup().MachineType != "" {
				if updated.GetGceSetup() == nil {
					updated.Infrastructure = &pb_v2.Instance_GceSetup{GceSetup: &pb_v2.GceSetup{}}
				}
				updated.GetGceSetup().MachineType = req.Instance.GetGceSetup().MachineType
			}
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
