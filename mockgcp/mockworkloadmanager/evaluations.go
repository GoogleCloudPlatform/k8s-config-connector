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

package mockworkloadmanager

import (
	"context"
	"fmt"
	"strings"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/workloadmanager/apiv1/workloadmanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type WorkloadManagerV1 struct {
	*MockService

	pb.UnimplementedWorkloadManagerServer
}

type evaluationName struct {
	Project      *projects.ProjectData
	Location     string
	EvaluationID string
}

func (n *evaluationName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/evaluations/%s", n.Project.ID, n.Location, n.EvaluationID)
}

func (s *WorkloadManagerV1) parseEvaluationName(name string) (*evaluationName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "evaluations" {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is invalid", name)
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	return &evaluationName{
		Project:      project,
		Location:     tokens[3],
		EvaluationID: tokens[5],
	}, nil
}

// Creates a new Evaluation in a given project and location.
func (s *WorkloadManagerV1) CreateEvaluation(ctx context.Context, req *pb.CreateEvaluationRequest) (*longrunningpb.Operation, error) {
	evaluationID := req.EvaluationId
	if evaluationID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "EvaluationId is required")
	}

	reqName := req.Parent + "/evaluations/" + evaluationID
	name, err := s.parseEvaluationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Evaluation).(*pb.Evaluation)
	obj.Name = fqn
	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	obj.ResourceStatus = &pb.ResourceStatus{
		State: pb.ResourceStatus_ACTIVE,
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: now,
		Target:     fqn,
		Verb:       "create",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

// Gets details of a single Evaluation.
func (s *WorkloadManagerV1) GetEvaluation(ctx context.Context, req *pb.GetEvaluationRequest) (*pb.Evaluation, error) {
	name, err := s.parseEvaluationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Evaluation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Evaluation %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

// Updates the parameters of a single Evaluation.
func (s *WorkloadManagerV1) UpdateEvaluation(ctx context.Context, req *pb.UpdateEvaluationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEvaluationName(req.Evaluation.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.Evaluation{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Evaluation)
	if req.Evaluation.Description != "" {
		updated.Description = req.Evaluation.Description
	}
	if req.Evaluation.ResourceFilter != nil {
		updated.ResourceFilter = req.Evaluation.ResourceFilter
	}
	if req.Evaluation.RuleNames != nil {
		updated.RuleNames = req.Evaluation.RuleNames
	}
	if req.Evaluation.Labels != nil {
		updated.Labels = req.Evaluation.Labels
	}
	if req.Evaluation.Schedule != nil {
		updated.Schedule = req.Evaluation.Schedule
	}
	if req.Evaluation.CustomRulesBucket != "" {
		updated.CustomRulesBucket = req.Evaluation.CustomRulesBucket
	}
	if req.Evaluation.BigQueryDestination != nil {
		updated.BigQueryDestination = req.Evaluation.BigQueryDestination
	}

	now := timestamppb.Now()
	updated.UpdateTime = now

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: now,
		Target:     fqn,
		Verb:       "update",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

// Deletes a single Evaluation.
func (s *WorkloadManagerV1) DeleteEvaluation(ctx context.Context, req *pb.DeleteEvaluationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEvaluationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.Evaluation{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: now,
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}
