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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// SnapshotSnapshotEncryptionKey_v1beta1_FromProto converts pb.CustomerEncryptionKey to krm.SnapshotSnapshotEncryptionKey.
// This is handcoded because rawKey is represented as a sensitive secret.Legacy struct in KRM, and ref fields require custom mapping.
func SnapshotSnapshotEncryptionKey_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CustomerEncryptionKey) *krm.SnapshotSnapshotEncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.SnapshotSnapshotEncryptionKey{}
	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	if in.GetKmsKeyServiceAccount() != "" {
		out.KmsKeyServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetKmsKeyServiceAccount()}
	}
	if in.GetRawKey() != "" {
		out.RawKey = &secret.Legacy{
			Value: in.RawKey,
		}
	}
	if in.GetSha256() != "" {
		out.Sha256 = in.Sha256
	}
	return out
}

// SnapshotSnapshotEncryptionKey_v1beta1_ToProto converts krm.SnapshotSnapshotEncryptionKey to pb.CustomerEncryptionKey.
// This is handcoded because rawKey is represented as a sensitive secret.Legacy struct in KRM, and ref fields require custom mapping.
func SnapshotSnapshotEncryptionKey_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SnapshotSnapshotEncryptionKey) *pb.CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.CustomerEncryptionKey{}
	if in.KmsKeyRef != nil {
		out.KmsKeyName = &in.KmsKeyRef.External
	}
	if in.KmsKeyServiceAccountRef != nil {
		out.KmsKeyServiceAccount = &in.KmsKeyServiceAccountRef.External
	}
	if in.RawKey != nil && in.RawKey.Value != nil {
		out.RawKey = in.RawKey.Value
	}
	if in.Sha256 != nil {
		out.Sha256 = in.Sha256
	}
	return out
}

// SnapshotSourceDiskEncryptionKey_v1beta1_FromProto converts pb.CustomerEncryptionKey to krm.SnapshotSourceDiskEncryptionKey.
// This is handcoded because rawKey is represented as a sensitive secret.Legacy struct in KRM, and ref fields require custom mapping.
func SnapshotSourceDiskEncryptionKey_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CustomerEncryptionKey) *krm.SnapshotSourceDiskEncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.SnapshotSourceDiskEncryptionKey{}
	if in.GetKmsKeyServiceAccount() != "" {
		out.KmsKeyServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetKmsKeyServiceAccount()}
	}
	if in.GetRawKey() != "" {
		out.RawKey = &secret.Legacy{
			Value: in.RawKey,
		}
	}
	return out
}

// SnapshotSourceDiskEncryptionKey_v1beta1_ToProto converts krm.SnapshotSourceDiskEncryptionKey to pb.CustomerEncryptionKey.
// This is handcoded because rawKey is represented as a sensitive secret.Legacy struct in KRM, and ref fields require custom mapping.
func SnapshotSourceDiskEncryptionKey_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SnapshotSourceDiskEncryptionKey) *pb.CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.CustomerEncryptionKey{}
	if in.KmsKeyServiceAccountRef != nil {
		out.KmsKeyServiceAccount = &in.KmsKeyServiceAccountRef.External
	}
	if in.RawKey != nil && in.RawKey.Value != nil {
		out.RawKey = in.RawKey.Value
	}
	return out
}
