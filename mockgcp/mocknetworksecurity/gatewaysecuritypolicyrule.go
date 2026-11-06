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

package mocknetworksecurity

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *NetworkSecurityV1Server) CreateGatewaySecurityPolicyRule(ctx context.Context, req *pb.CreateGatewaySecurityPolicyRuleRequest) (*longrunning.Operation, error) {
	name := req.Parent + "/rules/" + req.GatewaySecurityPolicyRuleId

	fqn := name

	obj := proto.Clone(req.GatewaySecurityPolicyRule).(*pb.GatewaySecurityPolicyRule)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name,
		Verb:                  "create",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj)
		return result, nil
	})
}

func (s *NetworkSecurityV1Server) GetGatewaySecurityPolicyRule(ctx context.Context, req *pb.GetGatewaySecurityPolicyRuleRequest) (*pb.GatewaySecurityPolicyRule, error) {
	name, err := s.parseGatewaySecurityPolicyRuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GatewaySecurityPolicyRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *NetworkSecurityV1Server) UpdateGatewaySecurityPolicyRule(ctx context.Context, req *pb.UpdateGatewaySecurityPolicyRuleRequest) (*longrunning.Operation, error) {
	name, err := s.parseGatewaySecurityPolicyRuleName(req.GetGatewaySecurityPolicyRule().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pb.GatewaySecurityPolicyRule{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(obj).(*pb.GatewaySecurityPolicyRule)
	updated.UpdateTime = timestamppb.New(time.Now())

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// If update_mask is not provided, all fields should be overwritten.
		updated = proto.Clone(req.GetGatewaySecurityPolicyRule()).(*pb.GatewaySecurityPolicyRule)
		updated.CreateTime = obj.CreateTime
		updated.UpdateTime = timestamppb.New(time.Now())
		updated.Name = obj.Name
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				updated.Description = req.GetGatewaySecurityPolicyRule().GetDescription()
			case "priority":
				updated.Priority = req.GetGatewaySecurityPolicyRule().GetPriority()
			case "enabled":
				updated.Enabled = req.GetGatewaySecurityPolicyRule().GetEnabled()
			case "sessionMatcher":
				updated.SessionMatcher = req.GetGatewaySecurityPolicyRule().GetSessionMatcher()
			case "applicationMatcher":
				updated.ApplicationMatcher = req.GetGatewaySecurityPolicyRule().GetApplicationMatcher()
			case "basicProfile":
				updated.Profile = req.GetGatewaySecurityPolicyRule().GetProfile()
			case "tlsInspectionEnabled":
				updated.TlsInspectionEnabled = req.GetGatewaySecurityPolicyRule().GetTlsInspectionEnabled()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
			}
		}
	}

	if err := s.storage.Update(ctx, name.String(), updated); err != nil {
		return nil, err
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(updated)
		return result, nil
	})
}

func (s *NetworkSecurityV1Server) DeleteGatewaySecurityPolicyRule(ctx context.Context, req *pb.DeleteGatewaySecurityPolicyRuleRequest) (*longrunning.Operation, error) {
	name, err := s.parseGatewaySecurityPolicyRuleName(req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, name.String(), &pb.GatewaySecurityPolicyRule{}); err != nil {
		return nil, err
	}
	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type gatewaySecurityPolicyRuleName struct {
	Project                   *projects.ProjectData
	Location                  string
	GatewaySecurityPolicy     string
	GatewaySecurityPolicyRule string
}

func (n *gatewaySecurityPolicyRuleName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/gatewaySecurityPolicies/" + n.GatewaySecurityPolicy + "/rules/" + n.GatewaySecurityPolicyRule
}

func (s *NetworkSecurityV1Server) parseGatewaySecurityPolicyRuleName(name string) (*gatewaySecurityPolicyRuleName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "gatewaySecurityPolicies" && tokens[6] == "rules" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		ruleName := &gatewaySecurityPolicyRuleName{
			Project:                   project,
			Location:                  tokens[3],
			GatewaySecurityPolicy:     tokens[5],
			GatewaySecurityPolicyRule: tokens[7],
		}
		return ruleName, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
