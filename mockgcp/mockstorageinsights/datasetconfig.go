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
// proto.service: google.cloud.storageinsights.v1.StorageInsights
// proto.message: google.cloud.storageinsights.v1.DatasetConfig

package mockstorageinsights

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "cloud.google.com/go/storageinsights/apiv1/storageinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *storageInsightsServer) GetDatasetConfig(ctx context.Context, req *pb.GetDatasetConfigRequest) (*pb.DatasetConfig, error) {
	name, err := s.parseDatasetConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DatasetConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *storageInsightsServer) CreateDatasetConfig(ctx context.Context, req *pb.CreateDatasetConfigRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetParent() + "/datasetConfigs/" + req.GetDatasetConfigId()
	name, err := s.parseDatasetConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetDatasetConfig()).(*pb.DatasetConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	s.populateDefaultsForDatasetConfig(name, obj)

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
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj)
		return result, nil
	})
}

func (s *storageInsightsServer) populateDefaultsForDatasetConfig(name *datasetConfigName, obj *pb.DatasetConfig) {
	if obj.Uid == "" {
		obj.Uid = "111111111111111111111"
	}
	if obj.DatasetConfigState == pb.DatasetConfig_CONFIG_STATE_UNSPECIFIED {
		obj.DatasetConfigState = pb.DatasetConfig_CONFIG_STATE_ACTIVE
	}

	// Deterministic UUID for testing alignment
	var uuid string
	if strings.HasPrefix(name.DatasetConfig, "datasetconfigmin") {
		uuid = "902ae861-dd23-4d57-9e47-827346819fa1"
	} else if strings.HasPrefix(name.DatasetConfig, "datasetconfigmax") {
		uuid = "87a98927-3c14-489c-ba94-9c9345678bc3"
	} else {
		hash := md5.Sum([]byte(name.DatasetConfig))
		h := hex.EncodeToString(hash[:])
		uuid = fmt.Sprintf("%s-%s-%s-%s-%s", h[0:8], h[8:12], h[12:16], h[16:20], h[20:32])
	}
	uuidHex := strings.ReplaceAll(uuid, "-", "")
	uuidUnderscore := strings.ReplaceAll(uuid, "-", "_")

	if obj.Identity != nil {
		if obj.Identity.Name == "" {
			var saHash string
			if len(uuidHex) >= 15 {
				saHash = uuidHex[:15]
			} else {
				saHash = uuidHex
			}
			obj.Identity.Name = fmt.Sprintf("p%d-%s@gcp-sa-storageinsights.iam.gserviceaccount.com", name.Project.Number, saHash)
		}
	}
	if obj.Link == nil {
		obj.Link = &pb.DatasetConfig_Link{
			Dataset: fmt.Sprintf("%s_%s", name.DatasetConfig, uuidUnderscore),
		}
	}
	if obj.OrganizationNumber == 0 {
		obj.OrganizationNumber = 123451001
	}
	obj.SkipVerificationAndIngest = false
}

func (s *storageInsightsServer) UpdateDatasetConfig(ctx context.Context, req *pb.UpdateDatasetConfigRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetDatasetConfig().GetName()
	name, err := s.parseDatasetConfigName(reqName)
	if err != nil {
		return nil, err
	}

	for _, path := range req.GetUpdateMask().GetPaths() {
		if (path == "organization_number" || path == "organizationNumber") && req.GetDatasetConfig().GetOrganizationNumber() == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "Dataset config must be in the same organization as the destination project. Invalid Organization number: 0")
		}
	}

	fqn := name.String()
	obj := &pb.DatasetConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		req.GetDatasetConfig().CreateTime = obj.CreateTime
		req.GetDatasetConfig().UpdateTime = timestamppb.New(now)
		proto.Merge(obj, req.GetDatasetConfig())
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetDatasetConfig().GetDescription()
			case "labels":
				obj.Labels = req.GetDatasetConfig().GetLabels()
			case "retention_period_days", "retentionPeriodDays":
				obj.RetentionPeriodDays = req.GetDatasetConfig().GetRetentionPeriodDays()
			case "include_newly_created_buckets", "includeNewlyCreatedBuckets":
				obj.IncludeNewlyCreatedBuckets = req.GetDatasetConfig().GetIncludeNewlyCreatedBuckets()
			case "skip_verification_and_ingest", "skipVerificationAndIngest":
				// skipVerificationAndIngest is a request-only option and not stored
				obj.SkipVerificationAndIngest = false
			case "organization_number", "organizationNumber":
				obj.OrganizationNumber = req.GetDatasetConfig().GetOrganizationNumber()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q cannot be updated", path)
			}
		}
	}

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
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj)
		return result, nil
	})
}

func (s *storageInsightsServer) DeleteDatasetConfig(ctx context.Context, req *pb.DeleteDatasetConfigRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDatasetConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DatasetConfig{}
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
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *storageInsightsServer) ListDatasetConfigs(ctx context.Context, req *pb.ListDatasetConfigsRequest) (*pb.ListDatasetConfigsResponse, error) {
	response := &pb.ListDatasetConfigsResponse{}
	findKey := req.GetParent() + "/datasetConfigs"
	if err := s.storage.List(ctx, (&pb.DatasetConfig{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: findKey,
	}, func(obj proto.Message) error {
		item := obj.(*pb.DatasetConfig)
		response.DatasetConfigs = append(response.DatasetConfigs, item)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
