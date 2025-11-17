// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:mockgcp-support
// proto.service: google.cloud.networksecurity.v1beta1.NetworkSecurity
// proto.message: google.cloud.networksecurity.v1beta1.AuthorizationPolicy

package mocknetworksecurity

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *NetworkSecurityServer) CreateAuthorizationPolicy(ctx context.Context, req *pb.CreateAuthorizationPolicyRequest) (*longrunning.Operation, error) {
	name := req.Parent + "/authorizationPolicies/" + req.AuthorizationPolicyId

	fqn := name

	obj := proto.Clone(req.AuthorizationPolicy).(*pb.AuthorizationPolicy)
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
		ApiVersion:            "v1beta1",
	}
	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj).(*pb.AuthorizationPolicy)
		return result, nil
	})
}

func (s *NetworkSecurityServer) GetAuthorizationPolicy(ctx context.Context, req *pb.GetAuthorizationPolicyRequest) (*pb.AuthorizationPolicy, error) {
	name, err := s.parseAuthorizationPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AuthorizationPolicy{}
	obj.Name = fqn
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	return obj, nil
}

func (s *NetworkSecurityServer) UpdateAuthorizationPolicy(ctx context.Context, req *pb.UpdateAuthorizationPolicyRequest) (*longrunning.Operation, error) {
	name, err := s.parseAuthorizationPolicyName(req.GetAuthorizationPolicy().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pb.AuthorizationPolicy{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(req.GetAuthorizationPolicy()).(*pb.AuthorizationPolicy)
	updated.CreateTime = obj.CreateTime
	updated.UpdateTime = timestamppb.New(time.Now())
	switch req.GetUpdateMask().GetPaths()[0] {
	case "rules":
		updated.Rules = req.GetAuthorizationPolicy().GetRules()
	case "action":
		updated.Action = req.GetAuthorizationPolicy().GetAction()
	default:
		return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled	 in mock", req.GetUpdateMask().GetPaths()[0])
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
		ApiVersion:            "v1beta1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(updated).(*pb.AuthorizationPolicy)
		return result, nil
	})
}

func (s *NetworkSecurityServer) DeleteAuthorizationPolicy(ctx context.Context, req *pb.DeleteAuthorizationPolicyRequest) (*longrunning.Operation, error) {
	name, err := s.parseAuthorizationPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, name.String(), &pb.AuthorizationPolicy{}); err != nil {
		return nil, err
	}
	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1beta1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type authorizationPolicyName struct {
	Project               *projects.ProjectData
	Location              string
	AuthorizationPolicyID string
}

func (n *authorizationPolicyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/authorizationPolicies/" + n.AuthorizationPolicyID
}

func (s *NetworkSecurityServer) parseAuthorizationPolicyName(name string) (*authorizationPolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "authorizationPolicies" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		name := &authorizationPolicyName{
			Project:               project,
			Location:              tokens[3],
			AuthorizationPolicyID: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
