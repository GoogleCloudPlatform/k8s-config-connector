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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ImageGuestOsFeatures_FromProto maps GuestOsFeature from pb to KRM.
func ImageGuestOsFeatures_FromProto(mapCtx *direct.MapContext, in *pb.GuestOsFeature) *krm.ImageGuestOsFeatures {
	if in == nil {
		return nil
	}
	out := &krm.ImageGuestOsFeatures{}
	out.Type = in.GetType()
	return out
}

// ImageGuestOsFeatures_ToProto maps GuestOsFeature from KRM to pb.
func ImageGuestOsFeatures_ToProto(mapCtx *direct.MapContext, in *krm.ImageGuestOsFeatures) *pb.GuestOsFeature {
	if in == nil {
		return nil
	}
	out := &pb.GuestOsFeature{}
	out.Type = direct.LazyPtr(in.Type)
	return out
}

// ImageImageEncryptionKey_FromProto maps CustomerEncryptionKey from pb to KRM.
// We handcode this because KMSCryptoKeyRef and IAMServiceAccountRef require custom structure translation.
func ImageImageEncryptionKey_FromProto(mapCtx *direct.MapContext, in *pb.CustomerEncryptionKey) *krm.ImageImageEncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.ImageImageEncryptionKey{}
	if in.GetKmsKeyName() != "" {
		out.KmsKeySelfLinkRef = &kmsv1beta1.KMSCryptoKeyRef{
			External: in.GetKmsKeyName(),
		}
	}
	if in.GetKmsKeyServiceAccount() != "" {
		out.KmsKeyServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{
			External: in.GetKmsKeyServiceAccount(),
		}
	}
	return out
}

// ImageImageEncryptionKey_ToProto maps CustomerEncryptionKey from KRM to pb.
// We handcode this because KMSCryptoKeyRef and IAMServiceAccountRef require custom structure translation.
func ImageImageEncryptionKey_ToProto(mapCtx *direct.MapContext, in *krm.ImageImageEncryptionKey) *pb.CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.CustomerEncryptionKey{}
	if in.KmsKeySelfLinkRef != nil {
		if in.KmsKeySelfLinkRef.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", in.KmsKeySelfLinkRef.Name)
		}
		out.KmsKeyName = direct.LazyPtr(in.KmsKeySelfLinkRef.External)
	}
	if in.KmsKeyServiceAccountRef != nil {
		if in.KmsKeyServiceAccountRef.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", in.KmsKeyServiceAccountRef.Name)
		}
		out.KmsKeyServiceAccount = direct.LazyPtr(in.KmsKeyServiceAccountRef.External)
	}
	return out
}

// ImageRawDisk_FromProto maps RawDisk from pb to KRM.
func ImageRawDisk_FromProto(mapCtx *direct.MapContext, in *pb.RawDisk) *krm.ImageRawDisk {
	if in == nil {
		return nil
	}
	out := &krm.ImageRawDisk{}
	out.ContainerType = in.ContainerType
	out.Sha1 = in.Sha1Checksum
	out.Source = in.GetSource()
	return out
}

// ImageRawDisk_ToProto maps RawDisk from KRM to pb.
func ImageRawDisk_ToProto(mapCtx *direct.MapContext, in *krm.ImageRawDisk) *pb.RawDisk {
	if in == nil {
		return nil
	}
	out := &pb.RawDisk{}
	out.ContainerType = in.ContainerType
	out.Sha1Checksum = in.Sha1
	out.Source = direct.LazyPtr(in.Source)
	return out
}

// ComputeDiskRef_FromProto maps a disk string reference from pb to KRM ComputeDiskRef.
func ComputeDiskRef_FromProto(mapCtx *direct.MapContext, in string) *krm.ComputeDiskRef {
	if in == "" {
		return nil
	}
	return &krm.ComputeDiskRef{
		External: in,
	}
}

// ComputeDiskRef_ToProto maps a KRM ComputeDiskRef to pb string reference.
func ComputeDiskRef_ToProto(mapCtx *direct.MapContext, in *krm.ComputeDiskRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}

// ComputeSnapshotRef_FromProto maps a snapshot string reference from pb to KRM ComputeSnapshotRef.
func ComputeSnapshotRef_FromProto(mapCtx *direct.MapContext, in string) *krm.ComputeSnapshotRef {
	if in == "" {
		return nil
	}
	return &krm.ComputeSnapshotRef{
		External: in,
	}
}

// ComputeSnapshotRef_ToProto maps a KRM ComputeSnapshotRef to pb string reference.
func ComputeSnapshotRef_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSnapshotRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}

// ComputeImageRef_FromProto maps an image string reference from pb to KRM ComputeImageRef.
func ComputeImageRef_FromProto(mapCtx *direct.MapContext, in string) *krm.ComputeImageRef {
	if in == "" {
		return nil
	}
	return &krm.ComputeImageRef{
		External: in,
	}
}

// ComputeImageRef_ToProto maps a KRM ComputeImageRef to pb string reference.
func ComputeImageRef_ToProto(mapCtx *direct.MapContext, in *krm.ComputeImageRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}

// ComputeImageSpec_v1beta1_FromProto maps a pb.Image to a krm.ComputeImageSpec.
// We hand-code this function because KRM has multiple resource reference fields (like diskRef, sourceImageRef, sourceSnapshotRef)
// that map to/from flat string fields in the proto definition, requiring custom reference parsing.
func ComputeImageSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Image) *krm.ComputeImageSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeImageSpec{}
	out.Description = in.Description
	out.DiskRef = ComputeDiskRef_FromProto(mapCtx, in.GetSourceDisk())
	out.DiskSizeGb = in.DiskSizeGb
	out.Family = in.Family
	out.GuestOsFeatures = direct.Slice_FromProto(mapCtx, in.GuestOsFeatures, ImageGuestOsFeatures_FromProto)
	out.ImageEncryptionKey = ImageImageEncryptionKey_FromProto(mapCtx, in.GetImageEncryptionKey())
	out.Licenses = in.Licenses
	out.RawDisk = ImageRawDisk_FromProto(mapCtx, in.GetRawDisk())
	out.SourceImageRef = ComputeImageRef_FromProto(mapCtx, in.GetSourceImage())
	out.SourceSnapshotRef = ComputeSnapshotRef_FromProto(mapCtx, in.GetSourceSnapshot())
	out.StorageLocations = in.StorageLocations
	return out
}

// ComputeImageSpec_v1beta1_ToProto maps a krm.ComputeImageSpec to a pb.Image.
// We hand-code this function because KRM has multiple resource reference fields (like diskRef, sourceImageRef, sourceSnapshotRef)
// that map to/from flat string fields in the proto definition, requiring custom reference parsing.
func ComputeImageSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeImageSpec) *pb.Image {
	if in == nil {
		return nil
	}
	out := &pb.Image{}
	out.Description = in.Description
	if in.DiskRef != nil {
		out.SourceDisk = ComputeDiskRef_ToProto(mapCtx, in.DiskRef)
	}
	out.DiskSizeGb = in.DiskSizeGb
	out.Family = in.Family
	out.GuestOsFeatures = direct.Slice_ToProto(mapCtx, in.GuestOsFeatures, ImageGuestOsFeatures_ToProto)
	out.ImageEncryptionKey = ImageImageEncryptionKey_ToProto(mapCtx, in.ImageEncryptionKey)
	out.Licenses = in.Licenses
	out.RawDisk = ImageRawDisk_ToProto(mapCtx, in.RawDisk)
	if in.SourceImageRef != nil {
		out.SourceImage = ComputeImageRef_ToProto(mapCtx, in.SourceImageRef)
	}
	if in.SourceSnapshotRef != nil {
		out.SourceSnapshot = ComputeSnapshotRef_ToProto(mapCtx, in.SourceSnapshotRef)
	}
	out.StorageLocations = in.StorageLocations
	return out
}

// ComputeImageStatus_v1beta1_FromProto maps a pb.Image to a krm.ComputeImageStatus.
func ComputeImageStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Image) *krm.ComputeImageStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeImageStatus{}
	out.ArchiveSizeBytes = in.ArchiveSizeBytes
	out.CreationTimestamp = in.CreationTimestamp
	out.LabelFingerprint = in.LabelFingerprint
	out.SelfLink = in.SelfLink
	return out
}

// ComputeImageStatus_v1beta1_ToProto maps a krm.ComputeImageStatus to a pb.Image.
func ComputeImageStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeImageStatus) *pb.Image {
	if in == nil {
		return nil
	}
	out := &pb.Image{}
	out.ArchiveSizeBytes = in.ArchiveSizeBytes
	out.CreationTimestamp = in.CreationTimestamp
	out.LabelFingerprint = in.LabelFingerprint
	out.SelfLink = in.SelfLink
	return out
}
