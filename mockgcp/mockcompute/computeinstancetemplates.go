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
// proto.service: google.cloud.compute.v1.InstanceTemplates
// proto.message: google.cloud.compute.v1.InstanceTemplate

package mockcompute

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type instanceTemplates struct {
	*MockService
	pb.UnimplementedInstanceTemplatesServer
}

func (s *instanceTemplates) Get(ctx context.Context, req *pb.GetInstanceTemplateRequest) (*pb.InstanceTemplate, error) {
	reqName := fmt.Sprintf("projects/%s/global/instanceTemplates/%s", req.GetProject(), req.GetInstanceTemplate())
	name, err := s.parseInstanceTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstanceTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "InstanceTemplate %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *instanceTemplates) Insert(ctx context.Context, req *pb.InsertInstanceTemplateRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/instanceTemplates/%s", req.GetProject(), req.GetInstanceTemplateResource().GetName())
	name, err := s.parseInstanceTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now().Format(time.RFC3339)
	obj := proto.Clone(req.GetInstanceTemplateResource()).(*pb.InstanceTemplate)
	obj.Id = proto.Uint64(s.generateID())
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.Kind = PtrTo("compute#instanceTemplate")
	obj.CreationTimestamp = &now

	s.populateDefaultsForInstanceTemplate(ctx, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.instanceTemplates.insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *instanceTemplates) populateDefaultsForInstanceTemplate(ctx context.Context, obj *pb.InstanceTemplate) {
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.Properties == nil {
		obj.Properties = &pb.InstanceProperties{}
	}
	properties := obj.Properties

	if properties.Metadata == nil {
		properties.Metadata = &pb.Metadata{}
	}
	properties.Metadata.Kind = PtrTo("compute#metadata")

	for i, networkInterface := range properties.NetworkInterfaces {
		networkInterface.Kind = PtrTo("compute#networkInterface")
		if networkInterface.GetName() == "" {
			networkInterface.Name = PtrTo(fmt.Sprintf("nic%d", i))
		}
		networkInterface.Network = PtrTo(makeFullyQualifiedNetwork(ctx, networkInterface.GetNetwork()))

		for _, accessConfig := range networkInterface.GetAccessConfigs() {
			accessConfig.Kind = PtrTo("compute#accessConfig")
			if accessConfig.GetNetworkTier() == "" {
				accessConfig.NetworkTier = PtrTo("PREMIUM")
			}
		}
	}

	for i, disk := range properties.Disks {
		disk.Kind = PtrTo("compute#attachedDisk")
		if disk.GetDeviceName() == "" {
			disk.DeviceName = PtrTo(fmt.Sprintf("persistent-disk-%d", i))
		}
		disk.Index = PtrTo(int32(i))
	}

	if properties.Scheduling == nil {
		properties.Scheduling = &pb.Scheduling{}
	}
	if properties.Scheduling.AutomaticRestart == nil {
		properties.Scheduling.AutomaticRestart = PtrTo(true)
	}
	if properties.Scheduling.OnHostMaintenance == nil {
		properties.Scheduling.OnHostMaintenance = PtrTo("MIGRATE")
	}
	if properties.Scheduling.Preemptible == nil {
		properties.Scheduling.Preemptible = PtrTo(false)
	}
	if properties.Scheduling.ProvisioningModel == nil {
		properties.Scheduling.ProvisioningModel = PtrTo("STANDARD")
	}

	properties.Metadata.Fingerprint = nil
	properties.Metadata.Fingerprint = PtrTo(computeFingerprint(properties.Metadata))
}

func (s *instanceTemplates) Delete(ctx context.Context, req *pb.DeleteInstanceTemplateRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/instanceTemplates/%s", req.GetProject(), req.GetInstanceTemplate())
	name, err := s.parseInstanceTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.InstanceTemplate{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type instanceTemplateName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *instanceTemplateName) String() string {
	return "projects/" + n.Project.ID + "/global/instanceTemplates/" + n.Name
}

// parseInstanceTemplateName parses a string into a instanceTemplateName.
// The expected form is `projects/*/global/instanceTemplates/*`.
func (s *MockService) parseInstanceTemplateName(name string) (*instanceTemplateName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "instanceTemplates" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &instanceTemplateName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
