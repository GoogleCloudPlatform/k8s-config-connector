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
// proto.service: google.cloud.batch.v1alpha.BatchService
// proto.message: google.cloud.batch.v1alpha.ResourceAllowance

package mockbatch

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	pb_v1alpha "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/batch/resourceallowance/pb"
)

type resourceAllowanceName struct {
	Project           *projects.ProjectData
	Location          string
	ResourceAllowance string
}

func (n *resourceAllowanceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/resourceAllowances/%s", n.Project.ID, n.Location, n.ResourceAllowance)
}

func (s *MockService) parseResourceAllowanceName(name string) (*resourceAllowanceName, error) {
	r := regexp.MustCompile("^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/resourceAllowances/(?P<resource_allowance>[^/]+)$")
	match := r.FindStringSubmatch(name)
	if len(match) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
	m := make(map[string]string)
	for i, n := range r.SubexpNames() {
		if len(n) > 0 {
			m[n] = match[i]
		}
	}

	project, err := s.Projects.GetProjectByID(m["project"])
	if err != nil {
		return nil, err
	}

	return &resourceAllowanceName{
		Project:           project,
		Location:          m["location"],
		ResourceAllowance: m["resource_allowance"],
	}, nil
}

func (s *BatchV1Alpha) CreateResourceAllowance(ctx context.Context, req *pb_v1alpha.CreateResourceAllowanceRequest) (*pb_v1alpha.ResourceAllowance, error) {
	reqName := req.Parent + "/resourceAllowances/" + req.ResourceAllowanceId
	name, err := s.parseResourceAllowanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ResourceAllowance).(*pb_v1alpha.ResourceAllowance)
	obj.Name = fqn
	obj.Uid = "b9a676df-c595-4c81-9963-f44b8e44e50c"
	obj.CreateTime = timestamppb.Now()

	if obj.GetUsageResourceAllowance() != nil {
		if obj.GetUsageResourceAllowance().Status == nil {
			obj.GetUsageResourceAllowance().Status = &pb_v1alpha.UsageResourceAllowanceStatus{}
		}
		obj.GetUsageResourceAllowance().Status.State = pb_v1alpha.ResourceAllowanceState_RESOURCE_ALLOWANCE_ACTIVE

		if obj.GetUsageResourceAllowance().Spec != nil && obj.GetUsageResourceAllowance().Spec.Limit != nil {
			limitVal := obj.GetUsageResourceAllowance().Spec.Limit.Limit
			obj.GetUsageResourceAllowance().Status.LimitStatus = &pb_v1alpha.UsageResourceAllowanceStatus_LimitStatus{
				Limit:    limitVal,
				Consumed: float64Ptr(0.0),
			}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BatchV1Alpha) GetResourceAllowance(ctx context.Context, req *pb_v1alpha.GetResourceAllowanceRequest) (*pb_v1alpha.ResourceAllowance, error) {
	name, err := s.parseResourceAllowanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb_v1alpha.ResourceAllowance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BatchV1Alpha) UpdateResourceAllowance(ctx context.Context, req *pb_v1alpha.UpdateResourceAllowanceRequest) (*pb_v1alpha.ResourceAllowance, error) {
	name, err := s.parseResourceAllowanceName(req.ResourceAllowance.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb_v1alpha.ResourceAllowance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(req.ResourceAllowance).(*pb_v1alpha.ResourceAllowance)
	updated.Uid = obj.Uid
	updated.CreateTime = obj.CreateTime

	if updated.GetUsageResourceAllowance() != nil {
		if updated.GetUsageResourceAllowance().Status == nil {
			updated.GetUsageResourceAllowance().Status = &pb_v1alpha.UsageResourceAllowanceStatus{}
		}
		updated.GetUsageResourceAllowance().Status.State = pb_v1alpha.ResourceAllowanceState_RESOURCE_ALLOWANCE_ACTIVE

		if updated.GetUsageResourceAllowance().Spec != nil && updated.GetUsageResourceAllowance().Spec.Limit != nil {
			limitVal := updated.GetUsageResourceAllowance().Spec.Limit.Limit
			updated.GetUsageResourceAllowance().Status.LimitStatus = &pb_v1alpha.UsageResourceAllowanceStatus_LimitStatus{
				Limit:    limitVal,
				Consumed: float64Ptr(0.0),
			}
			if obj.GetUsageResourceAllowance() != nil && obj.GetUsageResourceAllowance().Status != nil && obj.GetUsageResourceAllowance().Status.LimitStatus != nil {
				updated.GetUsageResourceAllowance().Status.LimitStatus.Consumed = obj.GetUsageResourceAllowance().Status.LimitStatus.Consumed
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *BatchV1Alpha) DeleteResourceAllowance(ctx context.Context, req *pb_v1alpha.DeleteResourceAllowanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseResourceAllowanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb_v1alpha.ResourceAllowance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	operationMetadata := &pb_v1alpha.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		EndTime:               timestamppb.New(now),
		ApiVersion:            "v1alpha",
		RequestedCancellation: false,
		Verb:                  "delete",
		Target:                fqn,
	}
	return s.operations.DoneLRO(ctx, fqn, operationMetadata, nil)
}

func (s *BatchV1Alpha) ListResourceAllowances(ctx context.Context, req *pb_v1alpha.ListResourceAllowancesRequest) (*pb_v1alpha.ListResourceAllowancesResponse, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be specified")
	}

	var resourceAllowances []*pb_v1alpha.ResourceAllowance
	if err := s.storage.List(ctx, (&pb_v1alpha.ResourceAllowance{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: parent + "/resourceAllowances/",
	}, func(obj proto.Message) error {
		allowance := obj.(*pb_v1alpha.ResourceAllowance)
		resourceAllowances = append(resourceAllowances, allowance)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb_v1alpha.ListResourceAllowancesResponse{
		ResourceAllowances: resourceAllowances,
	}, nil
}

func float64Ptr(v float64) *float64 {
	return &v
}
