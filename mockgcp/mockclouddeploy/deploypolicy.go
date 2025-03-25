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
// proto.service: google.cloud.deploy.v1.CloudDeploy
// proto.message: google.cloud.deploy.v1.DeployPolicy

package mockclouddeploy

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/deploy/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *CloudDeployV1) GetDeployPolicy(ctx context.Context, req *pb.GetDeployPolicyRequest) (*pb.DeployPolicy, error) {
	name, err := s.parseDeployPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DeployPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CloudDeployV1) CreateDeployPolicy(ctx context.Context, req *pb.CreateDeployPolicyRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/deployPolicies/%s", req.GetParent(), req.GetDeployPolicyId())
	name, err := s.parseDeployPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetDeployPolicy()).(*pb.DeployPolicy)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = "mock-uid"
	obj.Suspended = false

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:    "v1",
		CreateTime:    timestamppb.New(now),
		Target:        name.String(),
		Verb:          "create",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *CloudDeployV1) UpdateDeployPolicy(ctx context.Context, req *pb.UpdateDeployPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDeployPolicyName(req.GetDeployPolicy().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.DeployPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	now := time.Now()

	obj.UpdateTime = timestamppb.New(now)

	for _, path := range paths {
		switch path {
		case "suspended":
			obj.Suspended = req.GetDeployPolicy().GetSuspended()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *CloudDeployV1) DeleteDeployPolicy(ctx context.Context, req *pb.DeleteDeployPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDeployPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.DeployPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type deployPolicyName struct {
	Project    *projects.ProjectData
	Location   string
	DeployName string
}

func (n *deployPolicyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/deployPolicies/" + n.DeployName
}

// parseDeployPolicyName parses a string into a deployPolicyName.
// The expected form is `projects/*/locations/*/deployPolicies/*`.
func (s *MockService) parseDeployPolicyName(name string) (*deployPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deployPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &deployPolicyName{
			Project:    project,
			Location:   tokens[3],
			DeployName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

```
</out>


