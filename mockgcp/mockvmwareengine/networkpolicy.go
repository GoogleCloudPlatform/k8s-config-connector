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
// proto.service: google.cloud.vmwareengine.v1.VmwareEngine
// proto.message: google.cloud.vmwareengine.v1.NetworkPolicy

package mockvmwareengine

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/vmwareengine/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *VMwareEngineV1) GetNetworkPolicy(ctx context.Context, req *pb.GetNetworkPolicyRequest) (*pb.NetworkPolicy, error) {
	name, err := s.parseNetworkPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NetworkPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *VMwareEngineV1) CreateNetworkPolicy(ctx context.Context, req *pb.CreateNetworkPolicyRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/networkPolicies/" + req.NetworkPolicyId
	name, err := s.parseNetworkPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Check that the referenced VmwareEngineNetwork exists.
	if _, err := s.GetVmwareEngineNetwork(ctx, &pb.GetVmwareEngineNetworkRequest{Name: req.NetworkPolicy.VmwareEngineNetwork}); err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "vmwareEngineNetwork '%s' not found", req.NetworkPolicy.VmwareEngineNetwork)
	}

	now := time.Now()

	obj := proto.Clone(req.GetNetworkPolicy()).(*pb.NetworkPolicy)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	// change project ID to project Number in VmwareEngineNetworkCanonical
	obj.VmwareEngineNetworkCanonical = req.NetworkPolicy.VmwareEngineNetwork
	parts := strings.Split(obj.VmwareEngineNetworkCanonical, "/")
	if len(parts) >= 6 && parts[0] == "projects" && parts[2] == "locations" && parts[4] == "vmwareEngineNetworks" {
		obj.VmwareEngineNetworkCanonical = fmt.Sprintf("projects/%d/locations/%s/vmwareEngineNetworks/%s",
			name.Project.Number, parts[3], parts[5])
	}

	setDefaultNetworkPolicyFields(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(now)
		obj.UpdateTime = timestamppb.New(now)
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *VMwareEngineV1) UpdateNetworkPolicy(ctx context.Context, req *pb.UpdateNetworkPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNetworkPolicyName(req.GetNetworkPolicy().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NetworkPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, path := range req.UpdateMask.Paths {
		switch path {
		case "description":
			obj.Description = req.NetworkPolicy.Description
		default:
			// TODO: add support for other fields
			return nil, status.Errorf(codes.Unimplemented, "update mask field %q is not implemented in mockgcp", path)
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *VMwareEngineV1) DeleteNetworkPolicy(ctx context.Context, req *pb.DeleteNetworkPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNetworkPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.NetworkPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		ApiVersion: "v1",
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type networkPolicyName struct {
	Project         *projects.ProjectData
	Location        string
	NetworkPolicyID string
}

func (n *networkPolicyName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", n.Project.ID, n.Location, n.NetworkPolicyID)
}

// parseNetworkPolicyName parses a string into a networkPolicyName.
// The expected form is `projects/*/locations/*/networkPolicies/*`.
func (s *MockService) parseNetworkPolicyName(name string) (*networkPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "networkPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &networkPolicyName{
			Project:         project,
			Location:        tokens[3],
			NetworkPolicyID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func setDefaultNetworkPolicyFields(obj *pb.NetworkPolicy) {
	obj.Uid = "111111111111111111111"
	if obj.ExternalIp != nil && obj.ExternalIp.Enabled {
		obj.ExternalIp = &pb.NetworkPolicy_NetworkService{
			Enabled: true,
			State:   pb.NetworkPolicy_NetworkService_ACTIVE,
		}
	} else {
		obj.ExternalIp = &pb.NetworkPolicy_NetworkService{
			Enabled: false,
			State:   pb.NetworkPolicy_NetworkService_UNPROVISIONED,
		}
	}
	if obj.InternetAccess != nil && obj.InternetAccess.Enabled {
		obj.InternetAccess = &pb.NetworkPolicy_NetworkService{
			Enabled: true,
			State:   pb.NetworkPolicy_NetworkService_ACTIVE,
		}
	} else {
		obj.InternetAccess = &pb.NetworkPolicy_NetworkService{
			Enabled: false,
			State:   pb.NetworkPolicy_NetworkService_UNPROVISIONED,
		}
	}
}
