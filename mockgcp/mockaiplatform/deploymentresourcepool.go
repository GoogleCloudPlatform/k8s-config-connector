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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1beta1"
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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The DeploymentResourcePool does not exist.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *deploymentResourcePoolService) CreateDeploymentResourcePool(ctx context.Context, req *pb.CreateDeploymentResourcePoolRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/deploymentResourcePools/" + req.DeploymentResourcePoolId
	name, err := s.parseDeploymentResourcePoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.DeploymentResourcePool).(*pb.DeploymentResourcePool)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.SatisfiesPzi = false
	obj.SatisfiesPzs = false

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateDeploymentResourcePoolOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.DeploymentResourcePool)
		return result, nil
	})
}

func (s *deploymentResourcePoolService) UpdateDeploymentResourcePool(ctx context.Context, req *pb.UpdateDeploymentResourcePoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseDeploymentResourcePoolName(req.GetDeploymentResourcePool().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.DeploymentResourcePool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "dedicated_resources.min_replica_count", "dedicatedResources.minReplicaCount":
			if obj.DedicatedResources == nil {
				obj.DedicatedResources = &pb.DedicatedResources{}
			}
			obj.DedicatedResources.MinReplicaCount = req.GetDeploymentResourcePool().GetDedicatedResources().GetMinReplicaCount()
		case "dedicated_resources.max_replica_count", "dedicatedResources.maxReplicaCount":
			if obj.DedicatedResources == nil {
				obj.DedicatedResources = &pb.DedicatedResources{}
			}
			obj.DedicatedResources.MaxReplicaCount = req.GetDeploymentResourcePool().GetDedicatedResources().GetMaxReplicaCount()
		case "dedicated_resources.autoscaling_metric_specs", "dedicatedResources.autoscalingMetricSpecs":
			if obj.DedicatedResources == nil {
				obj.DedicatedResources = &pb.DedicatedResources{}
			}
			obj.DedicatedResources.AutoscalingMetricSpecs = req.GetDeploymentResourcePool().GetDedicatedResources().GetAutoscalingMetricSpecs()
		case "disable_container_logging", "disableContainerLogging":
			obj.DisableContainerLogging = req.GetDeploymentResourcePool().GetDisableContainerLogging()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
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
		result := proto.Clone(obj).(*pb.DeploymentResourcePool)
		return result, nil
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

type deploymentResourcePoolName struct {
	Project                  *projects.ProjectData
	Location                 string
	DeploymentResourcePoolId string
}

func (n *deploymentResourcePoolName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/deploymentResourcePools/%s", n.Project.ID, n.Location, n.DeploymentResourcePoolId)
}

func (s *MockService) parseDeploymentResourcePoolName(name string) (*deploymentResourcePoolName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deploymentResourcePools" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &deploymentResourcePoolName{
			Project:                  project,
			Location:                 tokens[3],
			DeploymentResourcePoolId: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
