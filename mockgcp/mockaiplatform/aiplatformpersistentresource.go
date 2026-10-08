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

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type persistentResourceService struct {
	*MockService
	pb.UnimplementedPersistentResourceServiceServer
}

func (s *persistentResourceService) GetPersistentResource(ctx context.Context, req *pb.GetPersistentResourceRequest) (*pb.PersistentResource, error) {
	name, err := s.parsePersistentResourceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PersistentResource{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The PersistentResource does not exist.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *persistentResourceService) CreatePersistentResource(ctx context.Context, req *pb.CreatePersistentResourceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/persistentResources/" + req.PersistentResourceId
	name, err := s.parsePersistentResourceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.PersistentResource).(*pb.PersistentResource)
	obj.Name = fqn

	if obj.DisplayName == "" {
		obj.DisplayName = req.PersistentResourceId
	}

	obj.CreateTime = timestamppb.New(now)
	obj.StartTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.PersistentResource_RUNNING
	obj.Error = &rpcstatus.Status{}

	for _, pool := range obj.ResourcePools {
		if pool.Id == "" {
			if pool.MachineSpec != nil {
				pool.Id = pool.MachineSpec.MachineType
			}
		}
		if pool.DiskSpec == nil {
			pool.DiskSpec = &pb.DiskSpec{
				BootDiskSizeGb: 100,
				BootDiskType:   "pd-ssd",
			}
		} else {
			if pool.DiskSpec.BootDiskSizeGb == 0 {
				pool.DiskSpec.BootDiskSizeGb = 100
			}
			if pool.DiskSpec.BootDiskType == "" {
				pool.DiskSpec.BootDiskType = "pd-ssd"
			}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.CreatePersistentResourceOperationMetadata{
		GenericMetadata: &pb.GenericOperationMetadata{
			CreateTime: timestamppb.New(now),
			UpdateTime: timestamppb.New(now),
		},
		ProgressMessage: "Create PersistentResource request received. Checking project configuration and quota...",
	}

	opPrefix := name.String()
	return s.operations.StartLRO(ctx, opPrefix, metadata, func() (proto.Message, error) {
		metadata.ProgressMessage = "Your persistent resource is ready."
		// For create, return the basic details (just name)
		return &pb.PersistentResource{
			Name: name.String(),
		}, nil
	})
}

func (s *persistentResourceService) DeletePersistentResource(ctx context.Context, req *pb.DeletePersistentResourceRequest) (*longrunning.Operation, error) {
	name, err := s.parsePersistentResourceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.PersistentResource{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	metadata := &pb.DeleteOperationMetadata{
		GenericMetadata: &pb.GenericOperationMetadata{
			CreateTime: timestamppb.New(now),
			UpdateTime: timestamppb.New(now),
		},
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, metadata, &emptypb.Empty{})
}

type PersistentResourceName struct {
	Project              *projects.ProjectData
	Location             string
	PersistentResourceID string
}

func (n *PersistentResourceName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/persistentResources/%s", n.Project.Number, n.Location, n.PersistentResourceID)
}

func (s *MockService) parsePersistentResourceName(name string) (*PersistentResourceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "persistentResources" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &PersistentResourceName{
			Project:              project,
			Location:             tokens[3],
			PersistentResourceID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
