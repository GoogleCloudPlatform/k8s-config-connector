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
	"hash/fnv"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func generateAnalysisRuleID(name string) string {
	h := fnv.New64a()
	h.Write([]byte(name))
	return strconv.FormatUint(h.Sum64(), 10)
}

func (s *ContactCenterInsightsServer) GetAnalysisRule(ctx context.Context, req *pb.GetAnalysisRuleRequest) (*pb.AnalysisRule, error) {
	name, err := s.parseAnalysisRuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AnalysisRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No AnalysisRule found for project: `%d` and AnalysisRule Id: `%s`.", name.Project.Number, name.AnalysisRule)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) CreateAnalysisRule(ctx context.Context, req *pb.CreateAnalysisRuleRequest) (*pb.AnalysisRule, error) {
	parentName, err := s.parseAnalysisRuleName(req.Parent + "/analysisRules/dummy")
	if err != nil {
		return nil, err
	}

	analysisRuleID := generateAnalysisRuleID(req.GetAnalysisRule().GetDisplayName())
	parentName.AnalysisRule = analysisRuleID

	fqn := parentName.String()

	obj := proto.Clone(req.AnalysisRule).(*pb.AnalysisRule)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) UpdateAnalysisRule(ctx context.Context, req *pb.UpdateAnalysisRuleRequest) (*pb.AnalysisRule, error) {
	reqName := req.GetAnalysisRule().GetName()
	name, err := s.parseAnalysisRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AnalysisRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No AnalysisRule found for project: `%d` and AnalysisRule Id: `%s`.", name.Project.Number, name.AnalysisRule)
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.AnalysisRule, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update CCInsightsAnalysisRule fields: %v", err)
		}
	} else {
		obj = proto.Clone(req.AnalysisRule).(*pb.AnalysisRule)
	}

	obj.Name = fqn

	now := timestamppb.Now()
	obj.UpdateTime = now

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) DeleteAnalysisRule(ctx context.Context, req *pb.DeleteAnalysisRuleRequest) (*emptypb.Empty, error) {
	name, err := s.parseAnalysisRuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AnalysisRule{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No AnalysisRule found for project: `%d` and AnalysisRule Id: `%s`.", name.Project.Number, name.AnalysisRule)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ContactCenterInsightsServer) ListAnalysisRules(ctx context.Context, req *pb.ListAnalysisRulesRequest) (*pb.ListAnalysisRulesResponse, error) {
	parentName, err := s.parseAnalysisRuleName(req.Parent + "/analysisRules/dummy")
	if err != nil {
		return nil, err
	}

	prefix := "projects/" + strconv.FormatInt(parentName.Project.Number, 10) + "/locations/" + parentName.Location + "/analysisRules/"
	var list []*pb.AnalysisRule
	findOptions := storage.ListOptions{
		Prefix: prefix,
	}
	if err := s.storage.List(ctx, (&pb.AnalysisRule{}).ProtoReflect().Descriptor(), findOptions, func(obj proto.Message) error {
		list = append(list, obj.(*pb.AnalysisRule))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListAnalysisRulesResponse{AnalysisRules: list}, nil
}
