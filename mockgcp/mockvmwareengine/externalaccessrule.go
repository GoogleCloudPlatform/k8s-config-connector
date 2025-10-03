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
// proto.message: google.cloud.vmwareengine.v1.ExternalAccessRule

package mockvmwareengine

import (
	"context"
	"fmt"
	"math/rand"
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

func (s *VMwareEngineV1) GetExternalAccessRule(ctx context.Context, req *pb.GetExternalAccessRuleRequest) (*pb.ExternalAccessRule, error) {
	name, err := s.parseExternalAccessRuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ExternalAccessRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *VMwareEngineV1) CreateExternalAccessRule(ctx context.Context, req *pb.CreateExternalAccessRuleRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/externalAccessRules/" + req.ExternalAccessRuleId
	name, err := s.parseExternalAccessRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Check if the parent Network Policy exists
	if _, err := s.GetNetworkPolicy(ctx, &pb.GetNetworkPolicyRequest{Name: name.NetworkPolicyName()}); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.FailedPrecondition, "parent resource '%s' not found", name.NetworkPolicyName())
		}
		return nil, err
	}

	// TODO: Validate destination_ip_ranges external_address references if provided

	now := time.Now()

	obj := proto.Clone(req.GetExternalAccessRule()).(*pb.ExternalAccessRule)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.ExternalAccessRule_ACTIVE // Or CREATING, then update in LRO
	obj.Uid = fmt.Sprintf("%x", rand.Int63())

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
		metadata.EndTime = timestamppb.Now()
		// If state was CREATING, update it to ACTIVE here
		// obj.State = pb.ExternalAccessRule_ACTIVE
		// obj.UpdateTime = timestamppb.Now()
		// if err := s.storage.Update(ctx, fqn, obj); err != nil {
		// 	return nil, err
		// }
		return obj, nil
	})
}

func (s *VMwareEngineV1) UpdateExternalAccessRule(ctx context.Context, req *pb.UpdateExternalAccessRuleRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseExternalAccessRuleName(req.GetExternalAccessRule().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.ExternalAccessRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Validate destination_ip_ranges external_address references if provided

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// For each field mentioned in update_mask, update the corresponding field in the object.
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetExternalAccessRule().GetDescription()
		case "priority":
			obj.Priority = req.GetExternalAccessRule().GetPriority()
		case "action":
			obj.Action = req.GetExternalAccessRule().GetAction()
		case "ip_protocol", "ipProtocol":
			obj.IpProtocol = req.GetExternalAccessRule().GetIpProtocol()
		case "source_ip_ranges", "sourceIpRanges":
			obj.SourceIpRanges = req.GetExternalAccessRule().GetSourceIpRanges()
		case "source_ports", "sourcePorts":
			obj.SourcePorts = req.GetExternalAccessRule().GetSourcePorts()
		case "destination_ip_ranges", "destinationIpRanges":
			obj.DestinationIpRanges = req.GetExternalAccessRule().GetDestinationIpRanges()
		case "destination_ports", "destinationPorts":
			obj.DestinationPorts = req.GetExternalAccessRule().GetDestinationPorts()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid for update", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *VMwareEngineV1) DeleteExternalAccessRule(ctx context.Context, req *pb.DeleteExternalAccessRuleRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseExternalAccessRuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ExternalAccessRule{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			// Deleting a non-existent resource should succeed according to API spec
		} else {
			return nil, err
		}
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

type externalAccessRuleName struct {
	Project              *projects.ProjectData
	Location             string
	NetworkPolicyID      string
	ExternalAccessRuleID string
}

func (n *externalAccessRuleName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s/externalAccessRules/%s", n.Project.ID, n.Location, n.NetworkPolicyID, n.ExternalAccessRuleID)
}

func (n *externalAccessRuleName) NetworkPolicyName() string {
	return fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", n.Project.ID, n.Location, n.NetworkPolicyID)
}

// parseExternalAccessRuleName parses a string into an externalAccessRuleName.
// The expected form is `projects/*/locations/*/networkPolicies/*/externalAccessRules/*`.
func (s *MockService) parseExternalAccessRuleName(name string) (*externalAccessRuleName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "networkPolicies" && tokens[6] == "externalAccessRules" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		ruleName := &externalAccessRuleName{
			Project:              project,
			Location:             tokens[3],
			NetworkPolicyID:      tokens[5],
			ExternalAccessRuleID: tokens[7],
		}

		return ruleName, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
