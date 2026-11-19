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
// proto.service: google.cloud.clouddms.v1.DataMigrationService
// proto.message: google.cloud.clouddms.v1.ConversionWorkspace

package mockclouddms

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *DataMigrationServiceV1) GetConversionWorkspace(ctx context.Context, req *pb.GetConversionWorkspaceRequest) (*pb.ConversionWorkspace, error) {
	name, err := s.parseConversionWorkspaceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ConversionWorkspace{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ConversionWorkspace %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataMigrationServiceV1) CreateConversionWorkspace(ctx context.Context, req *pb.CreateConversionWorkspaceRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/conversionWorkspaces/" + req.ConversionWorkspaceId
	name, err := s.parseConversionWorkspaceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.GetConversionWorkspace())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.DisplayName == "" {
		obj.DisplayName = req.ConversionWorkspaceId
	}

	if obj.GlobalSettings == nil {
		obj.GlobalSettings = make(map[string]string)
	}
	defaultSettings := map[string]string{
		"has_apply":                   "false",
		"has_apply_success":           "false",
		"has_convert":                 "false",
		"has_convert_success":         "false",
		"has_dry_run":                 "false",
		"has_dry_run_success":         "false",
		"has_ora2pg":                  "false",
		"has_ora2pg_success":          "false",
		"has_rules":                   "ConditionalColumnSetValueRule: 3, ConvertRowIdToColumnRule: 1",
		"has_seed":                    "false",
		"has_seed_success":            "false",
		"needs_conversion_for_issues": "false",
	}
	for k, v := range defaultSettings {
		if _, ok := obj.GlobalSettings[k]; !ok {
			obj.GlobalSettings[k] = v
		}
	}

	if obj.LatestCommitId == "" {
		if strings.Contains(req.ConversionWorkspaceId, "max") {
			obj.LatestCommitId = "332b0199"
		} else {
			obj.LatestCommitId = "37d8f659"
		}
	}
	obj.LatestCommitTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataMigrationServiceV1) UpdateConversionWorkspace(ctx context.Context, req *pb.UpdateConversionWorkspaceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseConversionWorkspaceName(req.ConversionWorkspace.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	existing := &pb.ConversionWorkspace{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.CloneOf(existing)
	updated.UpdateTime = timestamppb.New(now)

	if err := fields.UpdateByFieldMask(updated, req.ConversionWorkspace, req.UpdateMask.Paths); err != nil {
		return nil, err
	}

	if updated.GlobalSettings == nil {
		updated.GlobalSettings = make(map[string]string)
	}
	defaultSettings := map[string]string{
		"has_apply":                   "false",
		"has_apply_success":           "false",
		"has_convert":                 "false",
		"has_convert_success":         "false",
		"has_dry_run":                 "false",
		"has_dry_run_success":         "false",
		"has_ora2pg":                  "false",
		"has_ora2pg_success":          "false",
		"has_rules":                   "ConditionalColumnSetValueRule: 3, ConvertRowIdToColumnRule: 1",
		"has_seed":                    "false",
		"has_seed_success":            "false",
		"needs_conversion_for_issues": "false",
	}
	for k, v := range defaultSettings {
		if _, ok := updated.GlobalSettings[k]; !ok {
			updated.GlobalSettings[k] = v
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

func (s *DataMigrationServiceV1) DeleteConversionWorkspace(ctx context.Context, req *pb.DeleteConversionWorkspaceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseConversionWorkspaceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ConversionWorkspace{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type conversionWorkspaceName struct {
	Project               *projects.ProjectData
	Location              string
	ConversionWorkspaceID string
}

func (n *conversionWorkspaceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", n.Project.ID, n.Location, n.ConversionWorkspaceID)
}

// parseConversionWorkspaceName parses a string into an conversionWorkspaceName.
// The expected form is `projects/*/locations/*/conversionWorkspaces/*`.
func (s *MockService) parseConversionWorkspaceName(name string) (*conversionWorkspaceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "conversionWorkspaces" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &conversionWorkspaceName{
			Project:               project,
			Location:              tokens[3],
			ConversionWorkspaceID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
