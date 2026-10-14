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
// proto.service: google.cloud.networkservices.v1.NetworkServices
// proto.message: google.cloud.networkservices.v1.ServiceLbPolicy

package mocknetworkservices

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NetworkServicesServer) GetServiceLbPolicy(ctx context.Context, req *pb.GetServiceLbPolicyRequest) (*pb.ServiceLbPolicy, error) {
	name, err := s.parseServiceLbPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ServiceLbPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkServicesServer) ListServiceLbPolicies(ctx context.Context, req *pb.ListServiceLbPoliciesRequest) (*pb.ListServiceLbPoliciesResponse, error) {
	response := &pb.ListServiceLbPoliciesResponse{}

	parent, err := s.parseServiceLbPolicyParent(req.Parent)
	if err != nil {
		return nil, err
	}
	prefix := parent.String() + "/serviceLbPolicies/"

	findKind := (&pb.ServiceLbPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		item := obj.(*pb.ServiceLbPolicy)
		response.ServiceLbPolicies = append(response.ServiceLbPolicies, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *NetworkServicesServer) CreateServiceLbPolicy(ctx context.Context, req *pb.CreateServiceLbPolicyRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/serviceLbPolicies/" + req.ServiceLbPolicyId
	name, err := s.parseServiceLbPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.ServiceLbPolicy)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NetworkServicesServer) UpdateServiceLbPolicy(ctx context.Context, req *pb.UpdateServiceLbPolicyRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetServiceLbPolicy().GetName()

	name, err := s.parseServiceLbPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.ServiceLbPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		req.ServiceLbPolicy.CreateTime = obj.CreateTime
		req.ServiceLbPolicy.UpdateTime = timestamppb.New(now)
		req.ServiceLbPolicy.Name = obj.Name
		obj = req.ServiceLbPolicy
	} else {
		for _, path := range paths {
			switch path {
			case "labels":
				obj.Labels = req.GetServiceLbPolicy().GetLabels()
			case "description":
				obj.Description = req.GetServiceLbPolicy().GetDescription()
			case "loadBalancingAlgorithm", "load_balancing_algorithm":
				obj.LoadBalancingAlgorithm = req.GetServiceLbPolicy().GetLoadBalancingAlgorithm()
			case "autoCapacityDrain", "auto_capacity_drain":
				obj.AutoCapacityDrain = req.GetServiceLbPolicy().GetAutoCapacityDrain()
			case "failoverConfig", "failover_config":
				obj.FailoverConfig = req.GetServiceLbPolicy().GetFailoverConfig()
			case "isolationConfig", "isolation_config":
				obj.IsolationConfig = req.GetServiceLbPolicy().GetIsolationConfig()
			case "name":
				if req.GetServiceLbPolicy().GetName() != obj.GetName() {
					return nil, status.Errorf(codes.InvalidArgument, "field name is immutable")
				}
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
		obj.UpdateTime = timestamppb.New(now)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NetworkServicesServer) DeleteServiceLbPolicy(ctx context.Context, req *pb.DeleteServiceLbPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseServiceLbPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ServiceLbPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type serviceLbPolicyParent struct {
	Project  *projects.ProjectData
	Location string
}

func (p *serviceLbPolicyParent) String() string {
	return "projects/" + p.Project.ID + "/locations/" + p.Location
}

func (s *NetworkServicesServer) parseServiceLbPolicyParent(parent string) (*serviceLbPolicyParent, error) {
	tokens := strings.Split(parent, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &serviceLbPolicyParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

type serviceLbPolicyName struct {
	Project             *projects.ProjectData
	Location            string
	ServiceLbPolicyName string
}

func (n *serviceLbPolicyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/serviceLbPolicies/" + n.ServiceLbPolicyName
}

func (s *NetworkServicesServer) parseServiceLbPolicyName(name string) (*serviceLbPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "serviceLbPolicies" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &serviceLbPolicyName{
			Project:             project,
			Location:            tokens[3],
			ServiceLbPolicyName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
