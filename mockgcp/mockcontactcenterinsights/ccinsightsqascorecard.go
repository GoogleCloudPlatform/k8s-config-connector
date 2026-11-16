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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No QaScorecardRevision found for (project %d, qa_scorecard_id %s, revision_id latest).", name.Project.Number, name.QaScorecard)
		}
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

	// When a QaScorecard is created, GCP automatically creates an initial revision in state EDITABLE.
	initialRevID := "16702347239979751940"
	initialRevFQN := fqn + "/revisions/" + initialRevID
	initialRev := &pb.QaScorecardRevision{
		Name: initialRevFQN,
		Snapshot: &pb.QaScorecard{
			Name:        fqn,
			DisplayName: obj.DisplayName,
			CreateTime:  now,
			UpdateTime:  now,
		},
		CreateTime:   now,
		AlternateIds: []string{"latest"},
		State:        pb.QaScorecardRevision_EDITABLE,
	}
	if err := s.storage.Create(ctx, initialRevFQN, initialRev); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) UpdateQaScorecard(ctx context.Context, req *pb.UpdateQaScorecardRequest) (*pb.QaScorecard, error) {
	reqName := req.GetQaScorecard().GetName()
	name, err := s.parseQaScorecardName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.QaScorecard{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := timestamppb.Now()
	obj.UpdateTime = now
	obj.DisplayName = req.GetQaScorecard().GetDisplayName()
	obj.Description = req.GetQaScorecard().GetDescription()

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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Delete any revisions under this scorecard
	var revNames []string
	findOptions := storage.ListOptions{
		Prefix: fqn + "/revisions/",
	}
	_ = s.storage.List(ctx, (&pb.QaScorecardRevision{}).ProtoReflect().Descriptor(), findOptions, func(o proto.Message) error {
		rev := o.(*pb.QaScorecardRevision)
		revNames = append(revNames, rev.Name)
		return nil
	})
	for _, revName := range revNames {
		_ = s.storage.Delete(ctx, revName, &pb.QaScorecardRevision{})
	}

	return &emptypb.Empty{}, nil
}
