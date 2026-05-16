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

// krm.group: storage.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.storage.v1

package storage

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/storage/v1"
)

func StorageBucketCors_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Cors) *krm.StorageBucketCors {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketCors{}
	out.MaxAgeSeconds = direct.PtrTo(int64(in.MaxAgeSeconds))
	out.Method = in.Method
	out.Origin = in.Origin
	out.ResponseHeader = in.ResponseHeader
	return out
}
func StorageBucketCors_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketCors) *pb.Bucket_Cors {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Cors{}
	if in.MaxAgeSeconds != nil {
		out.MaxAgeSeconds = int32(*in.MaxAgeSeconds)
	}
	out.Method = in.Method
	out.Origin = in.Origin
	out.ResponseHeader = in.ResponseHeader
	return out
}

func StorageBucketWebsite_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Website) *krm.StorageBucketWebsite {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketWebsite{}
	if in.MainPageSuffix != "" {
		out.MainPageSuffix = direct.PtrTo(in.MainPageSuffix)
	}
	if in.NotFoundPage != "" {
		out.NotFoundPage = direct.PtrTo(in.NotFoundPage)
	}
	return out
}
func StorageBucketWebsite_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketWebsite) *pb.Bucket_Website {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Website{}
	if in.MainPageSuffix != nil {
		out.MainPageSuffix = *in.MainPageSuffix
	}
	if in.NotFoundPage != nil {
		out.NotFoundPage = *in.NotFoundPage
	}
	return out
}

func StorageBucketVersioning_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Versioning) *krm.StorageBucketVersioning {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketVersioning{}
	out.Enabled = in.Enabled
	return out
}
func StorageBucketVersioning_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketVersioning) *pb.Bucket_Versioning {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Versioning{}
	out.Enabled = in.Enabled
	return out
}

func StorageBucketLogging_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Logging) *krm.StorageBucketLogging {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketLogging{}
	out.LogBucket = in.LogBucket
	if in.LogObjectPrefix != "" {
		out.LogObjectPrefix = direct.PtrTo(in.LogObjectPrefix)
	}
	return out
}
func StorageBucketLogging_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketLogging) *pb.Bucket_Logging {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Logging{}
	out.LogBucket = in.LogBucket
	if in.LogObjectPrefix != nil {
		out.LogObjectPrefix = *in.LogObjectPrefix
	}
	return out
}

func StorageBucketEncryption_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Encryption) *krm.StorageBucketEncryption {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketEncryption{}
	if in.DefaultKmsKeyName != "" {
		out.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.DefaultKmsKeyName}
	}
	return out
}
func StorageBucketEncryption_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketEncryption) *pb.Bucket_Encryption {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Encryption{}
	if in.KmsKeyRef != nil {
		out.DefaultKmsKeyName = in.KmsKeyRef.External
	}
	return out
}

func StorageBucketRetentionPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_RetentionPolicy) *krm.StorageBucketRetentionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketRetentionPolicy{}
	out.IsLocked = direct.PtrTo(in.IsLocked)
	out.RetentionPeriod = in.RetentionPeriod
	return out
}
func StorageBucketRetentionPolicy_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketRetentionPolicy) *pb.Bucket_RetentionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_RetentionPolicy{}
	if in.IsLocked != nil {
		out.IsLocked = *in.IsLocked
	}
	out.RetentionPeriod = in.RetentionPeriod
	return out
}

func StorageBucketAutoclass_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Autoclass) *krm.StorageBucketAutoclass {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketAutoclass{}
	out.Enabled = direct.PtrTo(in.Enabled)
	return out
}
func StorageBucketAutoclass_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketAutoclass) *pb.Bucket_Autoclass {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Autoclass{}
	if in.Enabled != nil {
		out.Enabled = *in.Enabled
	}
	return out
}
