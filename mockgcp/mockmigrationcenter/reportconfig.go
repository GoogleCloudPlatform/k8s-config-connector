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

package mockmigrationcenter

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/migrationcenter/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type reportConfigName struct {
	Project        *projects.ProjectData
	Location       string
	ReportConfigID string
}

func (n *reportConfigName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/reportConfigs/%s", n.Project.ID, n.Location, n.ReportConfigID)
}

func (s *MockService) parseReportConfigName(name string) (*reportConfigName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "reportConfigs" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &reportConfigName{
			Project:        project,
			Location:       tokens[3],
			ReportConfigID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not in expected format", name)
}

func (s *MigrationCenterV1) GetReportConfig(ctx context.Context, req *pb.GetReportConfigRequest) (*pb.ReportConfig, error) {
	name, err := s.parseReportConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ReportConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *MigrationCenterV1) CreateReportConfig(ctx context.Context, req *pb.CreateReportConfigRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/reportConfigs/%s", req.GetParent(), req.GetReportConfigId())
	name, err := s.parseReportConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetReportConfig()).(*pb.ReportConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

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
		result := proto.Clone(obj).(*pb.ReportConfig)
		// Real GCP converts group and preferenceSet project IDs to project numbers in the LRO response
		for _, assignment := range result.GetGroupPreferencesetAssignments() {
			if strings.HasPrefix(assignment.GetGroup(), fmt.Sprintf("projects/%s/", name.Project.ID)) {
				assignment.Group = fmt.Sprintf("projects/%d/", name.Project.Number) + strings.TrimPrefix(assignment.GetGroup(), fmt.Sprintf("projects/%s/", name.Project.ID))
			}
			if strings.HasPrefix(assignment.GetPreferenceSet(), fmt.Sprintf("projects/%s/", name.Project.ID)) {
				assignment.PreferenceSet = fmt.Sprintf("projects/%d/", name.Project.Number) + strings.TrimPrefix(assignment.GetPreferenceSet(), fmt.Sprintf("projects/%s/", name.Project.ID))
			}
		}
		return result, nil
	})
}

func (s *MigrationCenterV1) DeleteReportConfig(ctx context.Context, req *pb.DeleteReportConfigRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseReportConfigName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ReportConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *MigrationCenterV1) ListReportConfigs(ctx context.Context, req *pb.ListReportConfigsRequest) (*pb.ListReportConfigsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is invalid", req.GetParent())
	}

	var reportConfigs []*pb.ReportConfig
	listKind := (&pb.ReportConfig{}).ProtoReflect().Descriptor()
	err := s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		rc := obj.(*pb.ReportConfig)
		name, err := s.parseReportConfigName(rc.GetName())
		if err == nil && name.Project.ID == tokens[1] && name.Location == tokens[3] {
			reportConfigs = append(reportConfigs, rc)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListReportConfigsResponse{
		ReportConfigs: reportConfigs,
	}, nil
}
