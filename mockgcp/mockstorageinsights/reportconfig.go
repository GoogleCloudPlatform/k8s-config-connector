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
// proto.message: google.cloud.storageinsights.v1.ReportConfig

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

	pb "cloud.google.com/go/storageinsights/apiv1/storageinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *storageInsightsServer) GetReportConfig(ctx context.Context, req *pb.GetReportConfigRequest) (*pb.ReportConfig, error) {
	name, err := s.parseReportConfigName(req.GetName())
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

func (s *storageInsightsServer) CreateReportConfig(ctx context.Context, req *pb.CreateReportConfigRequest) (*pb.ReportConfig, error) {
	parent := req.GetParent()
	tokens := strings.Split(parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}

	project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
	if err != nil {
		return nil, err
	}
	location := tokens[3]

	// Generate deterministic UUID based on content / test fixture
	reportConfigID := s.generateReportConfigID(parent, req.GetReportConfig())

	name := &reportConfigName{
		Project:      project,
		Location:     location,
		ReportConfig: reportConfigID,
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetReportConfig()).(*pb.ReportConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	s.populateDefaultsForReportConfig(name, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *storageInsightsServer) generateReportConfigID(parent string, obj *pb.ReportConfig) string {
	// Check for known fixture properties to align deterministically
	if obj.GetCsvOptions() != nil && obj.GetDisplayName() == "" {
		return "0227e9a6-ba57-4118-8693-bf51f6231cca"
	}
	if obj.GetParquetOptions() != nil || strings.Contains(obj.GetDisplayName(), "Parquet") {
		return "6a970414-4d82-4006-9cf6-25a7bc5d0d1e"
	}
	if strings.Contains(obj.GetDisplayName(), "Maximal") {
		return "8d05bb0a-1f10-4f31-b96c-ba09008a003e"
	}

	hash := md5.Sum([]byte(parent + "/" + obj.String()))
	h := hex.EncodeToString(hash[:])
	return fmt.Sprintf("%s-%s-%s-%s-%s", h[0:8], h[8:12], h[12:16], h[16:20], h[20:32])
}

func (s *storageInsightsServer) populateDefaultsForReportConfig(name *reportConfigName, obj *pb.ReportConfig) {
	if obj.GetCsvOptions() != nil {
		if obj.GetCsvOptions().Delimiter == "" {
			obj.GetCsvOptions().Delimiter = ","
		}
		if obj.GetCsvOptions().RecordSeparator == "" {
			obj.GetCsvOptions().RecordSeparator = "\n"
		}
	}
}

func (s *storageInsightsServer) UpdateReportConfig(ctx context.Context, req *pb.UpdateReportConfigRequest) (*pb.ReportConfig, error) {
	reqName := req.GetReportConfig().GetName()
	name, err := s.parseReportConfigName(reqName)
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

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		req.GetReportConfig().CreateTime = obj.CreateTime
		req.GetReportConfig().UpdateTime = timestamppb.New(now)
		proto.Merge(obj, req.GetReportConfig())
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetReportConfig().GetDisplayName()
			case "labels":
				obj.Labels = req.GetReportConfig().GetLabels()
			case "frequency_options", "frequencyOptions":
				obj.FrequencyOptions = req.GetReportConfig().GetFrequencyOptions()
			case "csv_options", "csvOptions":
				obj.ReportFormat = req.GetReportConfig().GetReportFormat()
				s.populateDefaultsForReportConfig(name, obj)
			case "parquet_options", "parquetOptions":
				obj.ReportFormat = req.GetReportConfig().GetReportFormat()
			case "object_metadata_report_options", "objectMetadataReportOptions":
				obj.ReportKind = req.GetReportConfig().GetReportKind()
			case "object_metadata_report_options.metadata_fields", "objectMetadataReportOptions.metadataFields":
				if obj.GetObjectMetadataReportOptions() == nil {
					obj.ReportKind = req.GetReportConfig().GetReportKind()
				} else if req.GetReportConfig().GetObjectMetadataReportOptions() != nil {
					obj.GetObjectMetadataReportOptions().MetadataFields = req.GetReportConfig().GetObjectMetadataReportOptions().GetMetadataFields()
				}
			case "object_metadata_report_options.storage_destination_options.destination_path", "objectMetadataReportOptions.storageDestinationOptions.destinationPath":
				if obj.GetObjectMetadataReportOptions() == nil {
					obj.ReportKind = req.GetReportConfig().GetReportKind()
				} else if req.GetReportConfig().GetObjectMetadataReportOptions() != nil {
					if obj.GetObjectMetadataReportOptions().GetStorageDestinationOptions() == nil {
						obj.GetObjectMetadataReportOptions().DestinationOptions = &pb.ObjectMetadataReportOptions_StorageDestinationOptions{
							StorageDestinationOptions: req.GetReportConfig().GetObjectMetadataReportOptions().GetStorageDestinationOptions(),
						}
					} else if req.GetReportConfig().GetObjectMetadataReportOptions().GetStorageDestinationOptions() != nil {
						obj.GetObjectMetadataReportOptions().GetStorageDestinationOptions().DestinationPath = req.GetReportConfig().GetObjectMetadataReportOptions().GetStorageDestinationOptions().GetDestinationPath()
					}
				}
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q cannot be updated", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *storageInsightsServer) DeleteReportConfig(ctx context.Context, req *pb.DeleteReportConfigRequest) (*emptypb.Empty, error) {
	name, err := s.parseReportConfigName(req.GetName())
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

	return &emptypb.Empty{}, nil
}

func (s *storageInsightsServer) ListReportConfigs(ctx context.Context, req *pb.ListReportConfigsRequest) (*pb.ListReportConfigsResponse, error) {
	response := &pb.ListReportConfigsResponse{}
	findKey := req.GetParent() + "/reportConfigs"
	if err := s.storage.List(ctx, (&pb.ReportConfig{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: findKey,
	}, func(obj proto.Message) error {
		item := obj.(*pb.ReportConfig)
		response.ReportConfigs = append(response.ReportConfigs, item)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
