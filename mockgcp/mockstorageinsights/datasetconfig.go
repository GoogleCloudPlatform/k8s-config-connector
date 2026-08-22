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

package mockstorageinsights

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/storageinsights/apiv1/storageinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"google.golang.org/genproto/googleapis/longrunning"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

type StorageInsightsServer struct {
	*MockService
	pb.UnimplementedStorageInsightsServer
}

func getUUID(name string) string {
	if strings.HasPrefix(name, "datasetconfigmin") {
		return "902ae861-dd23-4d57-9e47-827346819fa1"
	}
	if strings.HasPrefix(name, "datasetconfigmax") {
		return "87a98927-3c14-489c-ba94-9c9345678bc3"
	}
	h := sha256.Sum256([]byte(name))
	return fmt.Sprintf("%x-%x-%x-%x-%x", h[0:4], h[4:6], h[6:8], h[8:10], h[10:16])
}

func getOrganizationNumber() int64 {
	folderID := testgcp.TestFolderID.Get()
	if folderID != "" {
		if val, err := strconv.ParseInt(folderID, 10, 64); err == nil {
			return val
		}
	}
	orgID := testgcp.TestOrgID.Get()
	if orgID != "" {
		if val, err := strconv.ParseInt(orgID, 10, 64); err == nil {
			return val
		}
	}
	return 123451001 // Fallback mock value
}

func (s *StorageInsightsServer) GetDatasetConfig(ctx context.Context, req *pb.GetDatasetConfigRequest) (*pb.DatasetConfig, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request must be non-nil")
	}
	name, err := s.parseDatasetConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DatasetConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *StorageInsightsServer) CreateDatasetConfig(ctx context.Context, req *pb.CreateDatasetConfigRequest) (*longrunning.Operation, error) {
	if req == nil || req.DatasetConfig == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request and dataset config must be non-nil")
	}

	reqName := req.Parent + "/datasetConfigs/" + req.DatasetConfigId
	name, err := s.parseDatasetConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.DatasetConfig).(*pb.DatasetConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	// Validate retentionPeriodDays
	if obj.RetentionPeriodDays == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "retentionPeriodDays is a required field")
	}
	if obj.RetentionPeriodDays < 1 || obj.RetentionPeriodDays > 3650 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid retentionPeriodDays: %d. Must be between 1 and 3650.", obj.RetentionPeriodDays)
	}

	// Real GCP does not return skipVerificationAndIngest in GET responses after creation,
	// even if it was requested as true during POST. However, it is returned if set via PATCH.
	// We mimic this behavior by clearing it on creation.
	obj.SkipVerificationAndIngest = false

	// Populate output-only/dynamic fields
	uuidVal := getUUID(name.DatasetConfig)
	saPrefix := strings.ReplaceAll(uuidVal, "-", "")[:15]

	obj.Uid = "0123456789abcdef"
	obj.DatasetConfigState = pb.DatasetConfig_CONFIG_STATE_ACTIVE

	identityType := pb.Identity_IDENTITY_TYPE_PER_CONFIG
	if req.DatasetConfig.Identity != nil && req.DatasetConfig.Identity.Type != pb.Identity_IDENTITY_TYPE_UNSPECIFIED {
		identityType = req.DatasetConfig.Identity.Type
	}
	obj.Identity = &pb.Identity{
		Name: fmt.Sprintf("p%d-%s@gcp-sa-storageinsights.iam.gserviceaccount.com", name.Project.Number, saPrefix),
		Type: identityType,
	}
	datasetSuffix := strings.ReplaceAll(uuidVal, "-", "_")
	obj.Link = &pb.DatasetConfig_Link{
		Dataset: fmt.Sprintf("%s_%s", name.DatasetConfig, datasetSuffix),
	}
	if obj.OrganizationNumber == 0 {
		obj.OrganizationNumber = getOrganizationNumber()
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *StorageInsightsServer) UpdateDatasetConfig(ctx context.Context, req *pb.UpdateDatasetConfigRequest) (*longrunning.Operation, error) {
	if req == nil || req.DatasetConfig == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request and dataset config must be non-nil")
	}

	name, err := s.parseDatasetConfigName(req.DatasetConfig.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	existing := &pb.DatasetConfig{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Real GCP returns 400 bad request if OrganizationNumber is 0 on update.
	if req.DatasetConfig.OrganizationNumber == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Dataset config must be in the same organization as the destination project. Invalid Organization number: 0")
	}
	if req.DatasetConfig.OrganizationNumber != existing.OrganizationNumber {
		return nil, status.Errorf(codes.InvalidArgument, "Organization number cannot be changed once created")
	}

	updated := proto.Clone(existing).(*pb.DatasetConfig)
	paths := req.GetUpdateMask().GetPaths()
	if err := fields.UpdateByFieldMask(updated, req.DatasetConfig, paths); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating fields: %v", err)
	}

	// Validate retentionPeriodDays after applying update mask
	if updated.RetentionPeriodDays == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "retentionPeriodDays is a required field")
	}
	if updated.RetentionPeriodDays < 1 || updated.RetentionPeriodDays > 3650 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid retentionPeriodDays: %d. Must be between 1 and 3650.", updated.RetentionPeriodDays)
	}

	updated.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return updated, nil
	})
}

func (s *StorageInsightsServer) DeleteDatasetConfig(ctx context.Context, req *pb.DeleteDatasetConfigRequest) (*longrunning.Operation, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request must be non-nil")
	}
	name, err := s.parseDatasetConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.DatasetConfig{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *StorageInsightsServer) ListDatasetConfigs(ctx context.Context, req *pb.ListDatasetConfigsRequest) (*pb.ListDatasetConfigsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request must be non-nil")
	}
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}

	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := "projects/" + project.ID + "/locations/" + tokens[3] + "/datasetConfigs/"

	pageSize := int(req.GetPageSize())
	if pageSize <= 0 {
		pageSize = 100 // Default page size
	}
	if pageSize > 1000 {
		pageSize = 1000 // Max page size
	}

	startIndex := 0
	pageToken := req.GetPageToken()
	if pageToken != "" {
		parsed, err := strconv.Atoi(pageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid page token %q", pageToken)
		}
		startIndex = parsed
	}

	var all []*pb.DatasetConfig
	err = s.storage.List(ctx, (&pb.DatasetConfig{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.DatasetConfig)
		all = append(all, item)
		return nil
	})
	if err != nil {
		return nil, err
	}

	if startIndex < 0 {
		startIndex = 0
	}
	if startIndex > len(all) {
		startIndex = len(all)
	}

	endIndex := startIndex + pageSize
	if endIndex > len(all) {
		endIndex = len(all)
	}

	var response pb.ListDatasetConfigsResponse
	if startIndex < len(all) {
		response.DatasetConfigs = all[startIndex:endIndex]
	}

	if endIndex < len(all) {
		response.NextPageToken = strconv.Itoa(endIndex)
	}

	return &response, nil
}
