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
// proto.service: google.cloud.dataproc.v1.AutoscalingPolicyService
// proto.message: google.cloud.dataproc.v1.AutoscalingPolicy

package mockdataproc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type autoscalingPolicyServiceServer struct {
	*MockService
	pb.UnimplementedAutoscalingPolicyServiceServer
}

func (s *autoscalingPolicyServiceServer) CreateAutoscalingPolicy(ctx context.Context, req *pb.CreateAutoscalingPolicyRequest) (*pb.AutoscalingPolicy, error) {
	reqName := fmt.Sprintf("%s/autoscalingPolicies/%s", req.GetParent(), "test-autoscaling-policy")
	name, err := s.parseAutoscalingPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetPolicy()).(*pb.AutoscalingPolicy)
	obj.Name = fqn
	s.populateDefaultsForAutoscalingPolicy(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *autoscalingPolicyServiceServer) GetAutoscalingPolicy(ctx context.Context, req *pb.GetAutoscalingPolicyRequest) (*pb.AutoscalingPolicy, error) {
	name, err := s.parseAutoscalingPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AutoscalingPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *autoscalingPolicyServiceServer) DeleteAutoscalingPolicy(ctx context.Context, req *pb.DeleteAutoscalingPolicyRequest) (*emptypb.Empty, error) {
	name, err := s.parseAutoscalingPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AutoscalingPolicy{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *autoscalingPolicyServiceServer) UpdateAutoscalingPolicy(ctx context.Context, req *pb.UpdateAutoscalingPolicyRequest) (*pb.AutoscalingPolicy, error) {
	name, err := s.parseAutoscalingPolicyName(req.Policy.Id)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := req.Policy
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *autoscalingPolicyServiceServer) ListAutoscalingPolicy(ctx context.Context, req *pb.ListAutoscalingPoliciesRequest) (*pb.ListAutoscalingPoliciesResponse, error) {
	name, err := s.parseAutoscalingPolicyName(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.ListAutoscalingPoliciesResponse{}

	AutoscalingPolicyKind := (&pb.AutoscalingPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, AutoscalingPolicyKind, storage.ListOptions{}, func(obj proto.Message) error {
		autoScalingPolicy := obj.(*pb.AutoscalingPolicy)
		if strings.HasPrefix(autoScalingPolicy.GetName(), name.String()) {
			response.Policies = append(response.Policies, autoScalingPolicy)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *autoscalingPolicyServiceServer) populateDefaultsForAutoscalingPolicy(obj *pb.AutoscalingPolicy) {
	if obj.GetBasicAlgorithm() == nil {
		obj.Algorithm = &pb.AutoscalingPolicy_BasicAlgorithm{}
	}
	if obj.GetBasicAlgorithm().CooldownPeriod == nil {
		obj.GetBasicAlgorithm().CooldownPeriod = durationpb.New(2 * time.Minute)
	}
	if obj.GetBasicAlgorithm().GetYarnConfig() == nil {
		obj.GetBasicAlgorithm().Config = &pb.BasicAutoscalingAlgorithm_YarnConfig{}
	}
	if obj.GetBasicAlgorithm().GetYarnConfig().GracefulDecommissionTimeout == nil {
		obj.GetBasicAlgorithm().GetYarnConfig().GracefulDecommissionTimeout = durationpb.New(24 * time.Hour)
	}
	if obj.GetBasicAlgorithm().GetYarnConfig().ScaleDownFactor == 0 {
		obj.GetBasicAlgorithm().GetYarnConfig().ScaleDownFactor = 1
	}
	if obj.GetBasicAlgorithm().GetYarnConfig().ScaleDownMinWorkerFraction == 0 {
		obj.GetBasicAlgorithm().GetYarnConfig().ScaleDownMinWorkerFraction = 1
	}
	if obj.GetBasicAlgorithm().GetYarnConfig().ScaleUpFactor == 0 {
		obj.GetBasicAlgorithm().GetYarnConfig().ScaleUpFactor = 0.5
	}
	if obj.GetBasicAlgorithm().GetYarnConfig().ScaleUpMinWorkerFraction == 0 {
		obj.GetBasicAlgorithm().GetYarnConfig().ScaleUpMinWorkerFraction = 0.5
	}

	if obj.WorkerConfig == nil {
		obj.WorkerConfig = &pb.InstanceGroupAutoscalingPolicyConfig{}
	}
	if obj.WorkerConfig.MaxInstances == 0 {
		obj.WorkerConfig.MaxInstances = 5
	}

	if obj.SecondaryWorkerConfig == nil {
		obj.SecondaryWorkerConfig = &pb.InstanceGroupAutoscalingPolicyConfig{}
	}
	if obj.SecondaryWorkerConfig.MaxInstances == 0 {
		obj.SecondaryWorkerConfig.MaxInstances = 1
	}
}

type autoscalingPolicyName struct {
	Project            *projects.ProjectData
	Region             string
	AutoscalingPolicy  string
	AutoscalingPolicy2 string
}

func (n *autoscalingPolicyName) String() string {
	return fmt.Sprintf("projects/%s/regions/%s/autoscalingPolicies/%s", n.Project.ID, n.Region, n.AutoscalingPolicy)
}

// parseAutoscalingPolicyName parses a string into an AutoscalingPolicyName.
// The expected form is `projects/*/regions/*/autoscalingPolicies/*`.
func (s *MockService) parseAutoscalingPolicyName(name string) (*autoscalingPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "autoscalingPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &autoscalingPolicyName{
			Project:           project,
			Region:            tokens[3],
			AutoscalingPolicy: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

// buildAutoscalingPolicyName builds a AutoscalingPolicyName from the components.
func (s *MockService) buildAutoscalingPolicyName(projectName, region, cluster string) (*autoscalingPolicyName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &autoscalingPolicyName{
		Project:           project,
		Region:            region,
		AutoscalingPolicy: cluster,
	}, nil
}
