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

// +tool:mockgcp-support
// proto.service: google.cloud.aiplatform.v1beta1.DeploymentResourcePoolService
// proto.message: google.cloud.aiplatform.v1beta1.DeploymentResourcePool

package mockaiplatform

// deploymentResourcePoolService implements MockGCP and alignment for VertexAIDeploymentResourcePool.

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type deploymentResourcePoolService struct {
	*MockService
	pb.UnimplementedDeploymentResourcePoolServiceServer
}

func (s *deploymentResourcePoolService) GetDeploymentResourcePool(ctx context.Context, req *pb.GetDeploymentResourcePoolRequest) (*pb.DeploymentResourcePool, error) {
	name, err := s.parseDeploymentResourcePoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DeploymentResourcePool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *deploymentResourcePoolService) CreateDeploymentResourcePool(ctx context.Context, req *pb.CreateDeploymentResourcePoolRequest) (*longrunning.Operation, error) {
	parent := req.GetParent()
	id := req.GetDeploymentResourcePoolId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "deployment_resource_pool_id must be specified")
	}

	fqn := fmt.Sprintf("%s/deploymentResourcePools/%s", parent, id)
	_, err := s.parseDeploymentResourcePoolName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	obj := proto.Clone(req.DeploymentResourcePool).(*pb.DeploymentResourcePool)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	// Some fields are set by the API
	obj.SatisfiesPzi = true
	obj.SatisfiesPzs = true

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// We return a mock operation metadata
	op := &pb.CreateDeploymentResourcePoolOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *deploymentResourcePoolService) UpdateDeploymentResourcePool(ctx context.Context, req *pb.UpdateDeploymentResourcePoolRequest) (*longrunning.Operation, error) {
	reqName := req.DeploymentResourcePool.GetName()
	name, err := s.parseDeploymentResourcePoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.DeploymentResourcePool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DedicatedResources = req.DeploymentResourcePool.DedicatedResources
		obj.DisableContainerLogging = req.DeploymentResourcePool.DisableContainerLogging
	} else {
		for _, path := range paths {
			switch path {
			case "dedicated_resources", "dedicatedResources":
				obj.DedicatedResources = req.DeploymentResourcePool.DedicatedResources
			case "disable_container_logging", "disableContainerLogging":
				obj.DisableContainerLogging = req.DeploymentResourcePool.DisableContainerLogging
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UpdateDeploymentResourcePoolOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *deploymentResourcePoolService) DeleteDeploymentResourcePool(ctx context.Context, req *pb.DeleteDeploymentResourcePoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseDeploymentResourcePoolName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.DeploymentResourcePool{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type DeploymentResourcePoolName struct {
	Project                  *projects.ProjectData
	Location                 string
	DeploymentResourcePoolId string
}

func (s *deploymentResourcePoolService) parseDeploymentResourcePoolName(name string) (*DeploymentResourcePoolName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deploymentResourcePools" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &DeploymentResourcePoolName{
			Project:                  project,
			Location:                 tokens[3],
			DeploymentResourcePoolId: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}

func (n *DeploymentResourcePoolName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/deploymentResourcePools/%s", n.Project.ID, n.Location, n.DeploymentResourcePoolId)
}
