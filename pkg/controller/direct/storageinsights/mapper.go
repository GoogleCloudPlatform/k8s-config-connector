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

package storageinsights

import (
	"fmt"
	"strings"

	pb "cloud.google.com/go/storageinsights/apiv1/storageinsightspb"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storageinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func parseBucketName(external string) string {
	if strings.Contains(external, "/buckets/") {
		parts := strings.Split(external, "/buckets/")
		return parts[len(parts)-1]
	}
	if strings.HasPrefix(external, "gs://") {
		return strings.TrimPrefix(external, "gs://")
	}
	return external
}

func DatasetConfig_CloudStorageBuckets_CloudStorageBucket_FromProto(mapCtx *direct.MapContext, in *pb.DatasetConfig_CloudStorageBuckets_CloudStorageBucket) *krm.DatasetConfig_CloudStorageBuckets_CloudStorageBucket {
	if in == nil {
		return nil
	}
	out := &krm.DatasetConfig_CloudStorageBuckets_CloudStorageBucket{}
	if in.CloudStorageBucket != nil {
		switch opt := in.CloudStorageBucket.(type) {
		case *pb.DatasetConfig_CloudStorageBuckets_CloudStorageBucket_BucketName:
			out.BucketRef = &storagev1beta1.StorageBucketRef{External: opt.BucketName}
		case *pb.DatasetConfig_CloudStorageBuckets_CloudStorageBucket_BucketPrefixRegex:
			out.BucketPrefixRegex = &opt.BucketPrefixRegex
		}
	}
	return out
}

func DatasetConfig_CloudStorageBuckets_CloudStorageBucket_ToProto(mapCtx *direct.MapContext, in *krm.DatasetConfig_CloudStorageBuckets_CloudStorageBucket) *pb.DatasetConfig_CloudStorageBuckets_CloudStorageBucket {
	if in == nil {
		return nil
	}
	out := &pb.DatasetConfig_CloudStorageBuckets_CloudStorageBucket{}
	if in.BucketRef != nil && in.BucketRef.External != "" {
		out.CloudStorageBucket = &pb.DatasetConfig_CloudStorageBuckets_CloudStorageBucket_BucketName{
			BucketName: parseBucketName(in.BucketRef.External),
		}
	} else if in.BucketPrefixRegex != nil {
		out.CloudStorageBucket = &pb.DatasetConfig_CloudStorageBuckets_CloudStorageBucket_BucketPrefixRegex{
			BucketPrefixRegex: *in.BucketPrefixRegex,
		}
	}
	return out
}

func StorageInsightsDatasetConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatasetConfig) *krm.StorageInsightsDatasetConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageInsightsDatasetConfigSpec{}
	out.Labels = in.Labels
	out.OrganizationNumber = direct.LazyPtr(in.GetOrganizationNumber())
	out.SourceProjects = DatasetConfig_SourceProjects_FromProto(mapCtx, in.GetSourceProjects())
	out.SourceFolders = DatasetConfig_SourceFolders_FromProto(mapCtx, in.GetSourceFolders())

	if scope, ok := in.SourceOptions.(*pb.DatasetConfig_OrganizationScope); ok {
		out.OrganizationScope = direct.PtrTo(scope.OrganizationScope)
	}

	if in.GetCloudStorageObjectPath() != "" {
		path := in.GetCloudStorageObjectPath()
		var bucket, object string
		if strings.HasPrefix(path, "gs://") {
			path = strings.TrimPrefix(path, "gs://")
			parts := strings.SplitN(path, "/", 2)
			if len(parts) > 0 {
				bucket = parts[0]
			}
			if len(parts) > 1 {
				object = parts[1]
			}
		} else {
			bucket = path
		}
		out.CloudStorageObjectRef = &krm.CloudStorageObjectRef{}
		if bucket != "" {
			out.CloudStorageObjectRef.BucketRef = &storagev1beta1.StorageBucketRef{External: bucket}
		}
		if object != "" {
			out.CloudStorageObjectRef.Object = direct.LazyPtr(object)
		}
	}

	out.IncludeCloudStorageLocations = DatasetConfig_CloudStorageLocations_FromProto(mapCtx, in.GetIncludeCloudStorageLocations())
	out.ExcludeCloudStorageLocations = DatasetConfig_CloudStorageLocations_FromProto(mapCtx, in.GetExcludeCloudStorageLocations())
	out.IncludeCloudStorageBuckets = DatasetConfig_CloudStorageBuckets_FromProto(mapCtx, in.GetIncludeCloudStorageBuckets())
	out.ExcludeCloudStorageBuckets = DatasetConfig_CloudStorageBuckets_FromProto(mapCtx, in.GetExcludeCloudStorageBuckets())
	out.IncludeNewlyCreatedBuckets = direct.PtrTo(in.GetIncludeNewlyCreatedBuckets())
	out.SkipVerificationAndIngest = direct.PtrTo(in.GetSkipVerificationAndIngest())
	out.RetentionPeriodDays = direct.LazyPtr(in.GetRetentionPeriodDays())
	out.Identity = Identity_FromProto(mapCtx, in.GetIdentity())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

func StorageInsightsDatasetConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageInsightsDatasetConfigSpec) *pb.DatasetConfig {
	if in == nil {
		return nil
	}
	out := &pb.DatasetConfig{}
	out.Labels = in.Labels
	out.OrganizationNumber = direct.ValueOf(in.OrganizationNumber)
	if oneof := DatasetConfig_SourceProjects_ToProto(mapCtx, in.SourceProjects); oneof != nil {
		out.SourceOptions = &pb.DatasetConfig_SourceProjects_{SourceProjects: oneof}
	}
	if oneof := DatasetConfig_SourceFolders_ToProto(mapCtx, in.SourceFolders); oneof != nil {
		out.SourceOptions = &pb.DatasetConfig_SourceFolders_{SourceFolders: oneof}
	}
	if in.OrganizationScope != nil {
		out.SourceOptions = &pb.DatasetConfig_OrganizationScope{OrganizationScope: *in.OrganizationScope}
	}

	if in.CloudStorageObjectRef != nil {
		bucket := ""
		if in.CloudStorageObjectRef.BucketRef != nil {
			bucket = parseBucketName(in.CloudStorageObjectRef.BucketRef.External)
			bucket = strings.TrimPrefix(bucket, "gs://")
		}
		object := ""
		if in.CloudStorageObjectRef.Object != nil {
			object = *in.CloudStorageObjectRef.Object
		}
		if bucket == "" {
			mapCtx.Errorf("cloudStorageObjectRef.bucketRef cannot be empty when cloudStorageObjectRef is specified")
		} else {
			var path string
			if object != "" {
				path = fmt.Sprintf("gs://%s/%s", bucket, object)
			} else {
				path = fmt.Sprintf("gs://%s", bucket)
			}
			out.SourceOptions = &pb.DatasetConfig_CloudStorageObjectPath{
				CloudStorageObjectPath: path,
			}
		}
	}

	if oneof := DatasetConfig_CloudStorageLocations_ToProto(mapCtx, in.IncludeCloudStorageLocations); oneof != nil {
		out.CloudStorageLocations = &pb.DatasetConfig_IncludeCloudStorageLocations{IncludeCloudStorageLocations: oneof}
	}
	if oneof := DatasetConfig_CloudStorageLocations_ToProto(mapCtx, in.ExcludeCloudStorageLocations); oneof != nil {
		out.CloudStorageLocations = &pb.DatasetConfig_ExcludeCloudStorageLocations{ExcludeCloudStorageLocations: oneof}
	}
	if oneof := DatasetConfig_CloudStorageBuckets_ToProto(mapCtx, in.IncludeCloudStorageBuckets); oneof != nil {
		out.CloudStorageBuckets = &pb.DatasetConfig_IncludeCloudStorageBuckets{IncludeCloudStorageBuckets: oneof}
	}
	if oneof := DatasetConfig_CloudStorageBuckets_ToProto(mapCtx, in.ExcludeCloudStorageBuckets); oneof != nil {
		out.CloudStorageBuckets = &pb.DatasetConfig_ExcludeCloudStorageBuckets{ExcludeCloudStorageBuckets: oneof}
	}
	out.IncludeNewlyCreatedBuckets = direct.ValueOf(in.IncludeNewlyCreatedBuckets)
	out.SkipVerificationAndIngest = direct.ValueOf(in.SkipVerificationAndIngest)
	out.RetentionPeriodDays = direct.ValueOf(in.RetentionPeriodDays)
	out.Identity = Identity_ToProto(mapCtx, in.Identity)
	out.Description = direct.ValueOf(in.Description)
	return out
}
