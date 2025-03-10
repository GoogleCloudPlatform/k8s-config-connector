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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orgpolicy/v2"
	"github.com/golang/protobuf/ptypes/empty"
)

type orgPolicyV2 struct {
	*MockService
	pb.UnimplementedOrgPolicyServer
}

func (s *orgPolicyV2) GetCustomConstraint(ctx context.Context, req *pb.GetCustomConstraintRequest) (*pb.CustomConstraint, error) {
	name, err := s.parseCustomConstraintName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CustomConstraint{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *orgPolicyV2) CreateCustomConstraint(ctx context.Context, req *pb.CreateCustomConstraintRequest) (*pb.CustomConstraint, error) {
	reqName := req.CustomConstraint.Name
	name, err := s.parseCustomConstraintName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.CustomConstraint).(*pb.CustomConstraint)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *orgPolicyV2) UpdateCustomConstraint(ctx context.Context, req *pb.UpdateCustomConstraintRequest) (*pb.CustomConstraint, error) {
	reqName := req.GetCustomConstraint().GetName()

	name, err := s.parseCustomConstraintName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.CustomConstraint{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *orgPolicyV2) DeleteCustomConstraint(ctx context.Context, req *pb.DeleteCustomConstraintRequest) (*empty.Empty, error) {
	name, err := s.parseCustomConstraintName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.CustomConstraint{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type CustomConstraintName struct {
	Organization         string
	CustomConstraintName string
}

func (n *CustomConstraintName) String() string {
	return "organizations/" + n.Organization + "/CustomConstraints/" + n.CustomConstraintName
}

// parseCustomConstraintName parses a string into a CustomConstraintName.
// The expected form is organizations/<organizationID>/CustomConstraints/<CustomConstraintName>
func (s *MockService) parseCustomConstraintName(name string) (*CustomConstraintName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "organizations" && tokens[2] == "CustomConstraints" {

		name := &CustomConstraintName{
			Organization:         tokens[1],
			CustomConstraintName: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
