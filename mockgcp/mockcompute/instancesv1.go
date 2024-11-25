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

	obj := proto.Clone(req.GetInstanceResource()).(*pb.Instance)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instance")
	obj.Zone = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))
	obj.Status = PtrTo("RUNNING")
	if obj.LabelFingerprint == nil {
		obj.LabelFingerprint = PtrTo(computeFingerprint(obj))
	}
	// if obj.MachineType == nil {
	// 	machineType := "pd-standard"
	// 	obj.MachineType = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/machineTypes/%s", name.Project.ID, name.Zone, machineType)))
	// }

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
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

	return s.newLRO(ctx, name.Project.ID)
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

	attachedDisk := proto.Clone(req.GetAttachedDiskResource()).(*pb.AttachedDisk)
	attachedDisk.Kind = PtrTo("compute#attachedDisk")
	attachedDisk.Index = PtrTo(int32(len(obj.Disks)))

	obj.Disks = append(obj.Disks, attachedDisk)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
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

	return s.newLRO(ctx, name.Project.ID)
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

	return s.newLRO(ctx, name.Project.ID)
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

	return s.newLRO(ctx, name.Project.ID)
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

	return s.newLRO(ctx, name.Project.ID)
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

	return s.newLRO(ctx, name.Project.ID)
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
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
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

	return s.newLRO(ctx, name.Project.ID)
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
