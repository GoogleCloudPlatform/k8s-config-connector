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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DiskAsyncPrimaryDisk_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DiskAsyncReplication) *krm.DiskAsyncPrimaryDisk {
	if in == nil {
		return nil
	}
	out := &krm.DiskAsyncPrimaryDisk{}
	if in.GetDisk() != "" {
		out.DiskRef = krm.ComputeDiskRef{External: in.GetDisk()}
	} else {
		return nil
	}
	return out
}

func DiskAsyncPrimaryDisk_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DiskAsyncPrimaryDisk) *pb.DiskAsyncReplication {
	if in == nil {
		return nil
	}
	out := &pb.DiskAsyncReplication{}
	if in.DiskRef.External != "" {
		out.Disk = direct.PtrTo(in.DiskRef.External)
	}
	return out
}

func DiskDiskEncryptionKey_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CustomerEncryptionKey) *krm.DiskDiskEncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.DiskDiskEncryptionKey{}
	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	if in.GetKmsKeyServiceAccount() != "" {
		out.KmsKeyServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetKmsKeyServiceAccount()}
	}
	if in.RawKey != nil {
		out.RawKey = &secret.Legacy{Value: in.RawKey}
	}
	if in.RsaEncryptedKey != nil {
		out.RsaEncryptedKey = &secret.Legacy{Value: in.RsaEncryptedKey}
	}
	out.Sha256 = in.Sha256

	if out.KmsKeyRef == nil && out.KmsKeyServiceAccountRef == nil && out.RawKey == nil && out.RsaEncryptedKey == nil && out.Sha256 == nil {
		return nil
	}
	return out
}

func DiskDiskEncryptionKey_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DiskDiskEncryptionKey) *pb.CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.CustomerEncryptionKey{}
	if in.KmsKeyRef != nil && in.KmsKeyRef.External != "" {
		out.KmsKeyName = direct.PtrTo(in.KmsKeyRef.External)
	}
	if in.KmsKeyServiceAccountRef != nil && in.KmsKeyServiceAccountRef.External != "" {
		out.KmsKeyServiceAccount = direct.PtrTo(in.KmsKeyServiceAccountRef.External)
	}
	if in.RawKey != nil && in.RawKey.Value != nil {
		out.RawKey = in.RawKey.Value
	}
	if in.RsaEncryptedKey != nil && in.RsaEncryptedKey.Value != nil {
		out.RsaEncryptedKey = in.RsaEncryptedKey.Value
	}
	out.Sha256 = in.Sha256
	return out
}

func DiskGuestOsFeatures_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GuestOsFeature) *krm.DiskGuestOsFeatures {
	if in == nil {
		return nil
	}
	if in.GetType() == "" {
		return nil
	}
	return &krm.DiskGuestOsFeatures{
		Type: in.GetType(),
	}
}

func DiskGuestOsFeatures_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DiskGuestOsFeatures) *pb.GuestOsFeature {
	if in == nil {
		return nil
	}
	return &pb.GuestOsFeature{
		Type: direct.PtrTo(in.Type),
	}
}

func ComputeDiskSpec_ResourcePolicies_FromProto(mapCtx *direct.MapContext, in []string) []krm.ComputeResourcePolicyRef {
	if in == nil {
		return nil
	}
	var out []krm.ComputeResourcePolicyRef
	for _, i := range in {
		out = append(out, krm.ComputeResourcePolicyRef{
			External: i,
		})
	}
	return out
}

func ComputeDiskSpec_ResourcePolicies_ToProto(mapCtx *direct.MapContext, in []krm.ComputeResourcePolicyRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}

func DiskSourceImageEncryptionKey_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CustomerEncryptionKey) *krm.DiskSourceImageEncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.DiskSourceImageEncryptionKey{}
	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	if in.GetKmsKeyServiceAccount() != "" {
		out.KmsKeyServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetKmsKeyServiceAccount()}
	}
	out.RawKey = in.RawKey
	out.Sha256 = in.Sha256

	if out.KmsKeyRef == nil && out.KmsKeyServiceAccountRef == nil && out.RawKey == nil && out.Sha256 == nil {
		return nil
	}
	return out
}

func DiskSourceImageEncryptionKey_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DiskSourceImageEncryptionKey) *pb.CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.CustomerEncryptionKey{}
	if in.KmsKeyRef != nil && in.KmsKeyRef.External != "" {
		out.KmsKeyName = direct.PtrTo(in.KmsKeyRef.External)
	}
	if in.KmsKeyServiceAccountRef != nil && in.KmsKeyServiceAccountRef.External != "" {
		out.KmsKeyServiceAccount = direct.PtrTo(in.KmsKeyServiceAccountRef.External)
	}
	out.RawKey = in.RawKey
	out.Sha256 = in.Sha256
	return out
}

func DiskSourceSnapshotEncryptionKey_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CustomerEncryptionKey) *krm.DiskSourceSnapshotEncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.DiskSourceSnapshotEncryptionKey{}
	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	if in.GetKmsKeyServiceAccount() != "" {
		out.KmsKeyServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetKmsKeyServiceAccount()}
	}
	out.RawKey = in.RawKey
	out.Sha256 = in.Sha256

	if out.KmsKeyRef == nil && out.KmsKeyServiceAccountRef == nil && out.RawKey == nil && out.Sha256 == nil {
		return nil
	}
	return out
}

func DiskSourceSnapshotEncryptionKey_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DiskSourceSnapshotEncryptionKey) *pb.CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.CustomerEncryptionKey{}
	if in.KmsKeyRef != nil && in.KmsKeyRef.External != "" {
		out.KmsKeyName = direct.PtrTo(in.KmsKeyRef.External)
	}
	if in.KmsKeyServiceAccountRef != nil && in.KmsKeyServiceAccountRef.External != "" {
		out.KmsKeyServiceAccount = direct.PtrTo(in.KmsKeyServiceAccountRef.External)
	}
	out.RawKey = in.RawKey
	out.Sha256 = in.Sha256
	return out
}
