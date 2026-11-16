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
	"strconv"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *ContactCenterInsightsServer) GetQaScorecardRevision(ctx context.Context, req *pb.GetQaScorecardRevisionRequest) (*pb.QaScorecardRevision, error) {
	name, err := s.parseQaScorecardRevisionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.QaScorecardRevision{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No QaScorecardRevision found for (project %d, qa_scorecard_id %s, revision_id %s).", name.Project.Number, name.QaScorecard, name.Revision)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) CreateQaScorecardRevision(ctx context.Context, req *pb.CreateQaScorecardRevisionRequest) (*pb.QaScorecardRevision, error) {
	parentName, err := s.parseQaScorecardName(req.Parent)
	if err != nil {
		return nil, err
	}

	parentFQN := parentName.String()
	parentObj := &pb.QaScorecard{}
	if err := s.storage.Get(ctx, parentFQN, parentObj); err != nil {
		return nil, err
	}

	revisionID := req.QaScorecardRevisionId
	if revisionID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "qa_scorecard_revision_id must be set")
	}

	reqName := req.Parent + "/revisions/" + revisionID
	name, err := s.parseQaScorecardRevisionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.QaScorecardRevision).(*pb.QaScorecardRevision)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.AlternateIds = []string{"latest"}
	obj.State = pb.QaScorecardRevision_EDITABLE

	obj.Snapshot = &pb.QaScorecard{
		Name:        parentFQN,
		DisplayName: parentObj.DisplayName,
		Description: parentObj.Description,
		CreateTime:  parentObj.CreateTime,
		UpdateTime:  parentObj.UpdateTime,
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) TuneQaScorecardRevision(ctx context.Context, req *pb.TuneQaScorecardRevisionRequest) (*longrunning.Operation, error) {
	name, err := s.parseQaScorecardRevisionName(req.Parent)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	rev := &pb.QaScorecardRevision{}
	if err := s.storage.Get(ctx, fqn, rev); err != nil {
		return nil, err
	}

	prefix := "projects/" + strconv.FormatInt(name.Project.Number, 10) + "/locations/" + name.Location
	now := timestamppb.Now()
	metadata := &pb.TuneQaScorecardRevisionMetadata{
		CreateTime: now,
		Request:    req,
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		metadata.TuningCompletionRatio = 0.2

		rev.State = pb.QaScorecardRevision_READY
		if err := s.storage.Update(ctx, fqn, rev); err != nil {
			return nil, err
		}
		return nil, nil
	})
}

func (s *ContactCenterInsightsServer) DeployQaScorecardRevision(ctx context.Context, req *pb.DeployQaScorecardRevisionRequest) (*pb.QaScorecardRevision, error) {
	name, err := s.parseQaScorecardRevisionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.QaScorecardRevision{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.QaScorecardRevision_READY
	found := false
	for _, id := range obj.AlternateIds {
		if id == "latest" {
			found = true
			break
		}
	}
	if !found {
		obj.AlternateIds = append(obj.AlternateIds, "latest")
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) UndeployQaScorecardRevision(ctx context.Context, req *pb.UndeployQaScorecardRevisionRequest) (*pb.QaScorecardRevision, error) {
	name, err := s.parseQaScorecardRevisionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.QaScorecardRevision{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	var newAlternateIds []string
	for _, id := range obj.AlternateIds {
		if id != "latest" {
			newAlternateIds = append(newAlternateIds, id)
		}
	}
	obj.AlternateIds = newAlternateIds

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) DeleteQaScorecardRevision(ctx context.Context, req *pb.DeleteQaScorecardRevisionRequest) (*emptypb.Empty, error) {
	name, err := s.parseQaScorecardRevisionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.QaScorecardRevision{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ContactCenterInsightsServer) ListQaScorecardRevisions(ctx context.Context, req *pb.ListQaScorecardRevisionsRequest) (*pb.ListQaScorecardRevisionsResponse, error) {
	parentName, err := s.parseQaScorecardName(req.Parent)
	if err != nil {
		return nil, err
	}

	parentFQN := parentName.String()
	parentObj := &pb.QaScorecard{}
	if err := s.storage.Get(ctx, parentFQN, parentObj); err != nil {
		return nil, err
	}

	prefix := parentFQN + "/revisions/"
	var list []*pb.QaScorecardRevision
	findOptions := storage.ListOptions{
		Prefix: prefix,
	}
	if err := s.storage.List(ctx, (&pb.QaScorecardRevision{}).ProtoReflect().Descriptor(), findOptions, func(obj proto.Message) error {
		list = append(list, obj.(*pb.QaScorecardRevision))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListQaScorecardRevisionsResponse{QaScorecardRevisions: list}, nil
}
