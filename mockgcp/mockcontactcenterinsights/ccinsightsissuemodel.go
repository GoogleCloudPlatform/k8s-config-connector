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

package mockcontactcenterinsights

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *ContactCenterInsightsServer) GetIssueModel(ctx context.Context, req *pb.GetIssueModelRequest) (*pb.IssueModel, error) {
	name, err := s.parseIssueModelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.IssueModel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No IssueModel found for ID: %s", name.IssueModel)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) CreateIssueModel(ctx context.Context, req *pb.CreateIssueModelRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetIssueModel().GetName()
	name, err := s.parseIssueModelName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.IssueModel).(*pb.IssueModel)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.IssueModel_UNDEPLOYED

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, nil, func() (proto.Message, error) {
		obj.State = pb.IssueModel_DEPLOYED
		obj.UpdateTime = timestamppb.New(time.Now())
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *ContactCenterInsightsServer) UpdateIssueModel(ctx context.Context, req *pb.UpdateIssueModelRequest) (*pb.IssueModel, error) {
	reqName := req.GetIssueModel().GetName()
	name, err := s.parseIssueModelName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.IssueModel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No IssueModel found for ID: %s", name.IssueModel)
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.IssueModel, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update CCInsightsIssueModel fields: %v", err)
		}
	} else {
		obj = proto.Clone(req.IssueModel).(*pb.IssueModel)
	}

	obj.Name = fqn
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) DeleteIssueModel(ctx context.Context, req *pb.DeleteIssueModelRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseIssueModelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.IssueModel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No IssueModel found for ID: %s", name.IssueModel)
		}
		return nil, err
	}

	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, nil, func() (proto.Message, error) {
		if err := s.storage.Delete(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}
