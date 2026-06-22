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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type InstanceTemplatesV1 struct {
	*MockService
	pb.UnimplementedInstanceTemplatesServer
}

func (s *InstanceTemplatesV1) Insert(ctx context.Context, req *pb.InsertInstanceTemplateRequest) (*pb.Operation, error) {
	name := req.GetInstanceTemplateResource().GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	fqn := "projects/" + req.GetProject() + "/global/instanceTemplates/" + name
	selfLink := fmt.Sprintf("https://compute.googleapis.com/compute/v1/projects/%s/global/instanceTemplates/%s", req.GetProject(), name)

	obj := proto.CloneOf(req.GetInstanceTemplateResource())
	obj.SelfLink = PtrTo(selfLink)
	obj.Kind = PtrTo("compute#instanceTemplate")
	obj.Id = PtrTo(s.generateID())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Name = PtrTo(name)

	if obj.Properties == nil {
		obj.Properties = &pb.InstanceProperties{}
	}
	for i, disk := range obj.Properties.Disks {
		if disk.DeviceName == nil {
			disk.DeviceName = PtrTo(fmt.Sprintf("persistent-disk-%d", i))
		}
		if disk.Index == nil {
			disk.Index = PtrTo(int32(i))
		}
		if i == 0 {
			disk.Boot = PtrTo(true)
		}
		disk.Kind = PtrTo("compute#attachedDisk")
	}
	for _, networkInterface := range obj.Properties.NetworkInterfaces {
		if networkInterface.Name == nil {
			networkInterface.Name = PtrTo("nic0")
		}
		networkInterface.Kind = PtrTo("compute#networkInterface")
		if networkInterface.Network != nil {
			networkInterface.Network = PtrTo(ExpandComputeLink(ctx, *networkInterface.Network))
		}
		if networkInterface.Subnetwork != nil {
			networkInterface.Subnetwork = PtrTo(ExpandComputeLink(ctx, *networkInterface.Subnetwork))
		}
		for _, accessConfig := range networkInterface.AccessConfigs {
			accessConfig.Kind = PtrTo("compute#accessConfig")
		}
	}
	if obj.Properties.Scheduling == nil {
		obj.Properties.Scheduling = &pb.Scheduling{}
	}
	if obj.Properties.Scheduling.OnHostMaintenance == nil {
		obj.Properties.Scheduling.OnHostMaintenance = PtrTo("MIGRATE")
	}
	if obj.Properties.Scheduling.AutomaticRestart == nil {
		obj.Properties.Scheduling.AutomaticRestart = PtrTo(true)
	}
	if obj.Properties.Scheduling.Preemptible == nil {
		obj.Properties.Scheduling.Preemptible = PtrTo(false)
	}
	if obj.Properties.Metadata == nil {
		obj.Properties.Metadata = &pb.Metadata{}
	}
	if obj.Properties.Metadata.Fingerprint == nil {
		obj.Properties.Metadata.Fingerprint = PtrTo(computeFingerprint(obj.Properties.Metadata))
	}
	if obj.Properties.Metadata.Kind == nil {
		obj.Properties.Metadata.Kind = PtrTo("compute#metadata")
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.instanceTemplates.insert"),
		TargetLink:    PtrTo(selfLink),
		TargetId:      obj.Id,
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, req.GetProject(), op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstanceTemplatesV1) Get(ctx context.Context, req *pb.GetInstanceTemplateRequest) (*pb.InstanceTemplate, error) {
	if req.GetInstanceTemplate() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "InstanceTemplate is required")
	}
	fqn := "projects/" + req.GetProject() + "/global/instanceTemplates/" + req.GetInstanceTemplate()
	obj := &pb.InstanceTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Try lookup by numeric ID
			foundByNumericID := false
			prefix := "projects/" + req.GetProject() + "/global/instanceTemplates/"
			kind := (&pb.InstanceTemplate{}).ProtoReflect().Descriptor()
			errList := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(item proto.Message) error {
				template := item.(*pb.InstanceTemplate)
				if fmt.Sprintf("%d", template.GetId()) == req.GetInstanceTemplate() {
					obj = template
					foundByNumericID = true
				}
				return nil
			})
			if errList == nil && foundByNumericID {
				return obj, nil
			}

			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *InstanceTemplatesV1) Delete(ctx context.Context, req *pb.DeleteInstanceTemplateRequest) (*pb.Operation, error) {
	if req.GetInstanceTemplate() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "InstanceTemplate is required")
	}
	fqn := "projects/" + req.GetProject() + "/global/instanceTemplates/" + req.GetInstanceTemplate()
	deleted := &pb.InstanceTemplate{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			// Try lookup by numeric ID first to find the actual name
			prefix := "projects/" + req.GetProject() + "/global/instanceTemplates/"
			kind := (&pb.InstanceTemplate{}).ProtoReflect().Descriptor()
			var actualFQN string
			_ = s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(item proto.Message) error {
				template := item.(*pb.InstanceTemplate)
				if fmt.Sprintf("%d", template.GetId()) == req.GetInstanceTemplate() {
					actualFQN = "projects/" + req.GetProject() + "/global/instanceTemplates/" + template.GetName()
				}
				return nil
			})
			if actualFQN != "" {
				if err := s.storage.Delete(ctx, actualFQN, deleted); err == nil {
					goto deletedOK
				}
			}
		}
		return nil, err
	}

deletedOK:
	op := &pb.Operation{
		// Different than the other CRUD operation, DELETE operation type does not have "compute.instanceTemplates."
		OperationType: PtrTo("delete"),
		TargetLink:    deleted.SelfLink,
		TargetId:      deleted.Id,
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, req.GetProject(), op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}
