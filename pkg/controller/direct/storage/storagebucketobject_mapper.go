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

func StorageBucketObjectSpec_FromProto(mapCtx *direct.MapContext, in *pb.Object) *krm.StorageBucketObjectSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketObjectSpec{}
	out.ContentEncoding = direct.LazyPtr(in.GetContentEncoding())
	out.ContentDisposition = direct.LazyPtr(in.GetContentDisposition())
	out.CacheControl = direct.LazyPtr(in.GetCacheControl())
	out.Acl = direct.Slice_FromProto(mapCtx, in.Acl, ObjectAccessControl_FromProto)
	out.ContentLanguage = direct.LazyPtr(in.GetContentLanguage())
	out.ContentType = direct.LazyPtr(in.GetContentType())
	out.StorageClass = direct.LazyPtr(in.GetStorageClass())
	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	out.TemporaryHold = direct.LazyPtr(in.GetTemporaryHold())
	out.Metadata = in.Metadata
	out.EventBasedHold = direct.BoolValue_FromProto(mapCtx, in.GetEventBasedHold())
	if in.GetBucket() != "" {
		out.BucketRef = &krm.StorageBucketRef{External: in.GetBucket()}
	}
	out.CustomTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCustomTime())
	if in.GetName() != "" {
		out.ResourceID = direct.LazyPtr(in.GetName())
	}
	return out
}

func StorageBucketObjectSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketObjectSpec) *pb.Object {
	if in == nil {
		return nil
	}
	out := &pb.Object{}
	out.ContentEncoding = direct.ValueOf(in.ContentEncoding)
	out.ContentDisposition = direct.ValueOf(in.ContentDisposition)
	out.CacheControl = direct.ValueOf(in.CacheControl)
	out.Acl = direct.Slice_ToProto(mapCtx, in.Acl, ObjectAccessControl_ToProto)
	out.ContentLanguage = direct.ValueOf(in.ContentLanguage)
	out.ContentType = direct.ValueOf(in.ContentType)
	out.StorageClass = direct.ValueOf(in.StorageClass)
	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	out.TemporaryHold = direct.ValueOf(in.TemporaryHold)
	out.Metadata = in.Metadata
	out.EventBasedHold = direct.BoolValue_ToProto(mapCtx, in.EventBasedHold)
	if in.BucketRef != nil {
		out.Bucket = in.BucketRef.External
	}
	out.CustomTime = direct.StringTimestamp_ToProto(mapCtx, in.CustomTime)
	out.Name = direct.ValueOf(in.ResourceID)
	return out
}

func StorageBucketObjectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Object) *krm.StorageBucketObjectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketObjectObservedState{}
	out.Metageneration = direct.LazyPtr(in.GetMetageneration())
	out.TimeDeleted = direct.StringTimestamp_FromProto(mapCtx, in.GetTimeDeleted())
	out.Size = direct.LazyPtr(in.GetSize())
	out.TimeCreated = direct.StringTimestamp_FromProto(mapCtx, in.GetTimeCreated())
	out.Crc32C = direct.UInt32Value_FromProto(mapCtx, in.GetCrc32C())
	out.ComponentCount = direct.LazyPtr(in.GetComponentCount())
	out.Md5Hash = direct.LazyPtr(in.GetMd5Hash())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Updated = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdated())
	out.TimeStorageClassUpdated = direct.StringTimestamp_FromProto(mapCtx, in.GetTimeStorageClassUpdated())
	out.RetentionExpirationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRetentionExpirationTime())
	out.Generation = direct.LazyPtr(in.GetGeneration())
	out.Owner = Owner_FromProto(mapCtx, in.GetOwner())
	out.CustomerEncryption = Object_CustomerEncryption_FromProto(mapCtx, in.GetCustomerEncryption())
	return out
}

func StorageBucketObjectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketObjectObservedState) *pb.Object {
	if in == nil {
		return nil
	}
	out := &pb.Object{}
	out.Metageneration = direct.ValueOf(in.Metageneration)
	out.TimeDeleted = direct.StringTimestamp_ToProto(mapCtx, in.TimeDeleted)
	out.Size = direct.ValueOf(in.Size)
	out.TimeCreated = direct.StringTimestamp_ToProto(mapCtx, in.TimeCreated)
	out.Crc32C = direct.UInt32Value_ToProto(mapCtx, in.Crc32C)
	out.ComponentCount = direct.ValueOf(in.ComponentCount)
	out.Md5Hash = direct.ValueOf(in.Md5Hash)
	out.Etag = direct.ValueOf(in.Etag)
	out.Updated = direct.StringTimestamp_ToProto(mapCtx, in.Updated)
	out.TimeStorageClassUpdated = direct.StringTimestamp_ToProto(mapCtx, in.TimeStorageClassUpdated)
	out.RetentionExpirationTime = direct.StringTimestamp_ToProto(mapCtx, in.RetentionExpirationTime)
	out.Generation = direct.ValueOf(in.Generation)
	out.Owner = Owner_ToProto(mapCtx, in.Owner)
	out.CustomerEncryption = Object_CustomerEncryption_ToProto(mapCtx, in.CustomerEncryption)
	return out
}
