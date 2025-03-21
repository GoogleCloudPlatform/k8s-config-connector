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
// proto.service: google.cloud.netapp.v1.NetApp
// proto.message: google.cloud.netapp.v1.BackupPolicy

package mocknetapp

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/netapp/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NetAppV1) GetBackupPolicy(ctx context.Context, req *pb.GetBackupPolicyRequest) (*pb.BackupPolicy, error) {
	name, err := s.parseBackupPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *NetAppV1) CreateBackupPolicy(ctx context.Context, req *pb.CreateBackupPolicyRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/backupPolicies/%s", req.GetParent(), req.GetBackupPolicyId())
	name, err := s.parseBackupPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetBackupPolicy()).(*pb.BackupPolicy)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.BackupPolicy_READY

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
		lroMetadata.EndTime = timestamppb.Now()
		if err := s.storage.Get(ctx, fqn, obj); err != nil {
			return nil, err
		}
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *NetAppV1) UpdateBackupPolicy(ctx context.Context, req *pb.UpdateBackupPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPolicyName(req.GetBackupPolicy().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackupPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetBackupPolicy().Description
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.Etag = ComputeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *NetAppV1) DeleteBackupPolicy(ctx context.Context, req *pb.DeleteBackupPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.BackupPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type backupPolicyName struct {
	Project        *projects.ProjectData
	Location       string
	BackupPolicyID string
}

func (n *backupPolicyName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupPolicies/%s", n.Project.ID, n.Location, n.BackupPolicyID)
}

// parseBackupPolicyName parses a string into a backupPolicyName.
// The expected form is `projects/*/locations/*/backupPolicies/*`.
func (s *MockService) parseBackupPolicyName(name string) (*backupPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupPolicyName{
			Project:        project,
			Location:       tokens[3],
			BackupPolicyID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

```
</out>


