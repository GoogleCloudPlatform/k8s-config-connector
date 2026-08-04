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

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krmkmsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	krmkmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func CryptokeyversionAttestationStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation) *krmkmsv1alpha1.CryptokeyversionAttestationStatus {
	return CryptokeyversionAttestationStatus_FromProto(mapCtx, in)
}

func CryptoKeyVersionTemplate_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyVersionTemplate) *krmkmsv1beta1.CryptoKeyVersionTemplate {
	return CryptoKeyVersionTemplate_FromProto(mapCtx, in)
}

func KMSCryptoKeyVersionStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyVersion) *krmkmsv1alpha1.KMSCryptoKeyVersionStatus {
	return KMSCryptoKeyVersionStatus_FromProto(mapCtx, in)
}

func CryptokeyversionAttestationStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmkmsv1alpha1.CryptokeyversionAttestationStatus) *pb.KeyOperationAttestation {
	return CryptokeyversionAttestationStatus_ToProto(mapCtx, in)
}

func CryptoKeyVersionTemplate_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmkmsv1beta1.CryptoKeyVersionTemplate) *pb.CryptoKeyVersionTemplate {
	return CryptoKeyVersionTemplate_ToProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func KMSKeyRingImportJobStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krmkmsv1alpha1.KMSKeyRingImportJobStatus {
	return KMSKeyRingImportJobStatus_FromProto(mapCtx, in)
}

func KMSCryptoKeyVersionStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmkmsv1alpha1.KMSCryptoKeyVersionStatus) *pb.CryptoKeyVersion {
	return KMSCryptoKeyVersionStatus_ToProto(mapCtx, in)
}

func KMSKeyRingImportJobStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmkmsv1alpha1.KMSKeyRingImportJobStatus) *pb.ImportJob {
	return KMSKeyRingImportJobStatus_ToProto(mapCtx, in)
}

func KMSKeyRingImportJobSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krmkmsv1alpha1.KMSKeyRingImportJobSpec {
	return KMSKeyRingImportJobSpec_FromProto(mapCtx, in)
}

func KMSKeyRingImportJobSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmkmsv1alpha1.KMSKeyRingImportJobSpec) *pb.ImportJob {
	return KMSKeyRingImportJobSpec_ToProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func KeyringimportjobAttestationStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmkmsv1alpha1.KeyringimportjobAttestationStatus) *pb.KeyOperationAttestation {
	return KeyringimportjobAttestationStatus_ToProto(mapCtx, in)
}

func KeyringimportjobAttestationStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation) *krmkmsv1alpha1.KeyringimportjobAttestationStatus {
	return KeyringimportjobAttestationStatus_FromProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---
