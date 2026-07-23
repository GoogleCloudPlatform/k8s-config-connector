// Copyright 2026 Google LLC
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

// CreateAuthzPolicy creates a new AuthzPolicy in the simulated networksecurity service.
// This is aligned with the direct KCC controller and verified against both mock and real GCP behavior.
func (s *NetworkSecurityV1Server) CreateAuthzPolicy(ctx context.Context, req *pb.CreateAuthzPolicyRequest) (*longrunning.Operation, error) {
	name := req.Parent + "/authzPolicies/" + req.AuthzPolicyId

	fqn := name

	obj := proto.CloneOf(req.AuthzPolicy)
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
		result := proto.CloneOf(obj)
		return result, nil
	})
}

// GetAuthzPolicy retrieves the simulated AuthzPolicy from storage.
func (s *NetworkSecurityV1Server) GetAuthzPolicy(ctx context.Context, req *pb.GetAuthzPolicyRequest) (*pb.AuthzPolicy, error) {
	name, err := s.parseAuthzPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AuthzPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *NetworkSecurityV1Server) UpdateAuthzPolicy(ctx context.Context, req *pb.UpdateAuthzPolicyRequest) (*longrunning.Operation, error) {
	name, err := s.parseAuthzPolicyName(req.GetAuthzPolicy().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pb.AuthzPolicy{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	updated := proto.CloneOf(obj)
	updated.UpdateTime = timestamppb.New(time.Now())

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// If update_mask is not provided, all fields should be overwritten.
		updated = proto.CloneOf(req.GetAuthzPolicy())
		updated.CreateTime = obj.CreateTime
		updated.UpdateTime = timestamppb.New(time.Now())
		updated.Name = obj.Name
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				updated.Description = req.GetAuthzPolicy().GetDescription()
			case "target":
				updated.Target = req.GetAuthzPolicy().GetTarget()
			case "http_rules":
				updated.HttpRules = req.GetAuthzPolicy().GetHttpRules()
			case "network_rules":
				updated.NetworkRules = req.GetAuthzPolicy().GetNetworkRules()
			case "action":
				updated.Action = req.GetAuthzPolicy().GetAction()
			case "custom_provider":
				updated.CustomProvider = req.GetAuthzPolicy().GetCustomProvider()
			case "labels":
				updated.Labels = req.GetAuthzPolicy().GetLabels()
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
		result := proto.CloneOf(updated)
		return result, nil
	})
}

func (s *NetworkSecurityV1Server) DeleteAuthzPolicy(ctx context.Context, req *pb.DeleteAuthzPolicyRequest) (*longrunning.Operation, error) {
	name, err := s.parseAuthzPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, name.String(), &pb.AuthzPolicy{}); err != nil {
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

type authzPolicyName struct {
	Project       *projects.ProjectData
	Location      string
	AuthzPolicyID string
}

func (n *authzPolicyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/authzPolicies/" + n.AuthzPolicyID
}

func (s *NetworkSecurityV1Server) parseAuthzPolicyName(name string) (*authzPolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "authzPolicies" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		name := &authzPolicyName{
			Project:       project,
			Location:      tokens[3],
			AuthzPolicyID: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
