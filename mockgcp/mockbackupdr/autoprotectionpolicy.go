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
// proto.service: google.cloud.backupdr.v1.BackupDR
// proto.message: google.cloud.backupdr.v1.AutoProtectionPolicy

package mockbackupdr

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *BackupDRV1) GetAutoProtectionPolicy(ctx context.Context, req *pb.GetAutoProtectionPolicyRequest) (*pb.AutoProtectionPolicy, error) {
	name, err := s.parseAutoProtectionPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AutoProtectionPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupDRV1) CreateAutoProtectionPolicy(ctx context.Context, req *pb.CreateAutoProtectionPolicyRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/autoProtectionPolicies/%s", req.GetParent(), req.GetAutoProtectionPolicyId())
	name, err := s.parseAutoProtectionPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.GetAutoProtectionPolicy())
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.AutoProtectionPolicy_CREATING
	s.normalizeBackupPlanDetails(obj)
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		obj.State = pb.AutoProtectionPolicy_ACTIVE
		lroMetadata.EndTime = timestamppb.New(time.Now())
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return obj, nil
	})
}

func (s *BackupDRV1) UpdateAutoProtectionPolicy(ctx context.Context, req *pb.UpdateAutoProtectionPolicyRequest) (*longrunningpb.Operation, error) {
	reqPolicy := req.GetAutoProtectionPolicy()
	if reqPolicy == nil {
		return nil, status.Errorf(codes.InvalidArgument, "auto_protection_policy must be provided")
	}

	name, err := s.parseAutoProtectionPolicyName(reqPolicy.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.AutoProtectionPolicy{}
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
			obj.Description = reqPolicy.GetDescription()
		case "criteria":
			obj.Criteria = reqPolicy.GetCriteria()
		case "backup_plan_details", "backupPlanDetails":
			obj.BackupPlanDetails = reqPolicy.GetBackupPlanDetails()
			s.normalizeBackupPlanDetails(obj)
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not mutable", path)
		}
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = fields.ComputeWeakEtag(obj)

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return obj, nil
	})
}

func (s *BackupDRV1) DeleteAutoProtectionPolicy(ctx context.Context, req *pb.DeleteAutoProtectionPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAutoProtectionPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.AutoProtectionPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *BackupDRV1) ListAutoProtectionPolicies(ctx context.Context, req *pb.ListAutoProtectionPoliciesRequest) (*pb.ListAutoProtectionPoliciesResponse, error) {
	name, err := s.parseAutoProtectionPolicyName(req.GetParent() + "/autoProtectionPolicies/dummy")
	if err != nil {
		return nil, err
	}

	response := &pb.ListAutoProtectionPoliciesResponse{}

	findPrefix := strings.TrimSuffix(name.String(), "dummy")

	listKind := (&pb.AutoProtectionPolicy{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		policy := obj.(*pb.AutoProtectionPolicy)
		if strings.HasPrefix(policy.GetName(), findPrefix) {
			response.AutoProtectionPolicies = append(response.AutoProtectionPolicies, policy)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *MockService) normalizeBackupPlanDetails(obj *pb.AutoProtectionPolicy) {
	for _, detail := range obj.GetBackupPlanDetails() {
		if detail == nil || detail.BackupPlan == "" {
			continue
		}
		tokens := strings.Split(detail.BackupPlan, "/")
		if len(tokens) >= 2 && tokens[0] == "projects" {
			project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
			if err == nil {
				tokens[1] = fmt.Sprintf("%d", project.Number)
				detail.BackupPlan = strings.Join(tokens, "/")
			}
		}
	}
}

type autoProtectionPolicyName struct {
	Project              *projects.ProjectData
	Location             string
	AutoProtectionPolicy string
}

func (n *autoProtectionPolicyName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/autoProtectionPolicies/%s", n.Project.ID, n.Location, n.AutoProtectionPolicy)
}

func (s *MockService) parseAutoProtectionPolicyName(name string) (*autoProtectionPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "autoProtectionPolicies" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &autoProtectionPolicyName{
			Project:              project,
			Location:             tokens[3],
			AutoProtectionPolicy: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
