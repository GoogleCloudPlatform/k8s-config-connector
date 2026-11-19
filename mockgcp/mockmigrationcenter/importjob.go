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

type importJobName struct {
	Project     *projects.ProjectData
	Location    string
	ImportJobID string
}

func (n *importJobName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/importJobs/%s", n.Project.ID, n.Location, n.ImportJobID)
}

func (s *MockService) parseImportJobName(name string) (*importJobName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "importJobs" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &importJobName{
			Project:     project,
			Location:    tokens[3],
			ImportJobID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not in expected format", name)
}

func (s *MockService) populateDefaultsForImportJob(obj *pb.ImportJob) {
	if obj.State == pb.ImportJob_IMPORT_JOB_STATE_UNSPECIFIED {
		obj.State = pb.ImportJob_IMPORT_JOB_STATE_PENDING
	}
}

func (s *MockService) formatProjectNumberInAssetSource(assetSource string) string {
	name, err := s.parseSourceName(assetSource)
	if err == nil {
		if name.Project.Number != 0 {
			return fmt.Sprintf("projects/%d/locations/%s/sources/%s", name.Project.Number, name.Location, name.SourceName)
		}
		return fmt.Sprintf("projects/%s/locations/%s/sources/%s", name.Project.ID, name.Location, name.SourceName)
	}
	return assetSource
}

func (s *MigrationCenterV1) GetImportJob(ctx context.Context, req *pb.GetImportJobRequest) (*pb.ImportJob, error) {
	name, err := s.parseImportJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ImportJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	s.populateDefaultsForImportJob(obj)
	return obj, nil
}

func (s *MigrationCenterV1) CreateImportJob(ctx context.Context, req *pb.CreateImportJobRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/importJobs/%s", req.GetParent(), req.GetImportJobId())
	name, err := s.parseImportJobName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetImportJob()).(*pb.ImportJob)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	s.populateDefaultsForImportJob(obj)

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
		lroObj := proto.Clone(obj).(*pb.ImportJob)
		if lroObj.AssetSource != "" {
			lroObj.AssetSource = s.formatProjectNumberInAssetSource(lroObj.AssetSource)
		}
		return lroObj, nil
	})
}

func (s *MigrationCenterV1) UpdateImportJob(ctx context.Context, req *pb.UpdateImportJobRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseImportJobName(req.GetImportJob().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ImportJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := time.Now()

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = req.GetImportJob().GetDisplayName()
		obj.Labels = req.GetImportJob().GetLabels()
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetImportJob().GetDisplayName()
			case "labels":
				obj.Labels = req.GetImportJob().GetLabels()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported in mock", path)
			}
		}
	}
	obj.UpdateTime = timestamppb.New(now)
	s.populateDefaultsForImportJob(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		lroObj := proto.Clone(obj).(*pb.ImportJob)
		if lroObj.AssetSource != "" {
			lroObj.AssetSource = s.formatProjectNumberInAssetSource(lroObj.AssetSource)
		}
		return lroObj, nil
	})
}

func (s *MigrationCenterV1) DeleteImportJob(ctx context.Context, req *pb.DeleteImportJobRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseImportJobName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ImportJob{}
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

func (s *MigrationCenterV1) ListImportJobs(ctx context.Context, req *pb.ListImportJobsRequest) (*pb.ListImportJobsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is invalid", req.GetParent())
	}

	var importJobs []*pb.ImportJob
	listKind := (&pb.ImportJob{}).ProtoReflect().Descriptor()
	err := s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		importJob := obj.(*pb.ImportJob)
		name, err := s.parseImportJobName(importJob.GetName())
		if err == nil && name.Project.ID == tokens[1] && name.Location == tokens[3] {
			s.populateDefaultsForImportJob(importJob)
			importJobs = append(importJobs, importJob)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListImportJobsResponse{
		ImportJobs: importJobs,
	}, nil
}
