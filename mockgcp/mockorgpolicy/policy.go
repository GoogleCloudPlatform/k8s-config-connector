// Copyright 2022 Google LLC
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

package mockorgpolicy

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orgpolicy/v2"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *orgPolicyV2) GetPolicy(ctx context.Context, req *pb.GetPolicyRequest) (*pb.Policy, error) {
	name, err := s.parsePolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *orgPolicyV2) CreatePolicy(ctx context.Context, req *pb.CreatePolicyRequest) (*pb.Policy, error) {
	reqName := req.Policy.Name
	name, err := s.parsePolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Policy).(*pb.Policy)
	obj.Name = fqn
	if obj.Spec != nil {
		obj.Spec.UpdateTime = timestamppb.New(time.Now())
		obj.Spec.Etag = fields.ComputeWeakEtag(obj.Spec)
	}
	if obj.DryRunSpec != nil {
		obj.DryRunSpec.UpdateTime = timestamppb.New(time.Now())
		obj.DryRunSpec.Etag = fields.ComputeWeakEtag(obj.DryRunSpec)
	}
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *orgPolicyV2) UpdatePolicy(ctx context.Context, req *pb.UpdatePolicyRequest) (*pb.Policy, error) {
	reqName := req.GetPolicy().GetName()

	name, err := s.parsePolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// use the new object from update request
	obj = proto.Clone(req.GetPolicy()).(*pb.Policy)
	if obj.Spec != nil {
		obj.Spec.UpdateTime = timestamppb.New(time.Now())
		obj.Spec.Etag = fields.ComputeWeakEtag(obj.Spec)
	}
	if obj.DryRunSpec != nil {
		obj.DryRunSpec.UpdateTime = timestamppb.New(time.Now())
		obj.DryRunSpec.Etag = fields.ComputeWeakEtag(obj.DryRunSpec)
	}
	obj.Etag = fields.ComputeWeakEtag(obj)
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *orgPolicyV2) DeletePolicy(ctx context.Context, req *pb.DeletePolicyRequest) (*empty.Empty, error) {
	name, err := s.parsePolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Policy{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type PolicyName struct {
	Organization string
	Folder       string
	Project      string
	PolicyName   string
}

func (n *PolicyName) String() string {
	var parent string
	if n.Folder != "" {
		parent = "folders/" + n.Folder
	} else if n.Project != "" {
		parent = "projects/" + n.Project
	} else if n.Organization != "" {
		parent = "organizations/" + n.Organization
	} else {
		return ""
	}
	return parent + "/policies/" + n.PolicyName
}

// parsePolicyName parses a string into a PolicyName.
func (s *MockService) parsePolicyName(name string) (*PolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[2] == "policies" {
		var name *PolicyName
		name = &PolicyName{}
		switch tokens[0] {
		case "projects":
			name.Project = tokens[1]
		case "folders":
			name.Folder = tokens[1]
		case "organizations":
			name.Organization = tokens[1]
		default:
			return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", tokens[0])
		}
		name.PolicyName = tokens[3]

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
