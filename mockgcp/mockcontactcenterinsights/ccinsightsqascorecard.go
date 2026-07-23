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

	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ContactCenterInsightsServer struct {
	*MockService
	pb.UnimplementedContactCenterInsightsServer
}

func (s *ContactCenterInsightsServer) GetQaScorecard(ctx context.Context, req *pb.GetQaScorecardRequest) (*pb.QaScorecard, error) {
	name, err := s.parseQaScorecardName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.QaScorecard{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) CreateQaScorecard(ctx context.Context, req *pb.CreateQaScorecardRequest) (*pb.QaScorecard, error) {
	reqName := req.Parent + "/qaScorecards/" + req.QaScorecardId
	name, err := s.parseQaScorecardName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.QaScorecard).(*pb.QaScorecard)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) UpdateQaScorecard(ctx context.Context, req *pb.UpdateQaScorecardRequest) (*pb.QaScorecard, error) {
	qaScorecardName := req.GetQaScorecard().GetName()

	name, err := s.parseQaScorecardName(qaScorecardName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.QaScorecard{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	paths := updateMask.GetPaths()
	if len(paths) == 0 {
		// Treat as all paths if empty
		paths = []string{
			"display_name", "description",
		}
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetQaScorecard().GetDisplayName()
		case "description":
			obj.Description = req.GetQaScorecard().GetDescription()
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) DeleteQaScorecard(ctx context.Context, req *pb.DeleteQaScorecardRequest) (*emptypb.Empty, error) {
	name, err := s.parseQaScorecardName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.QaScorecard{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
