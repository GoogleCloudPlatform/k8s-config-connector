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

func KMSCryptoKeyStatus_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKey) *krmkmsv1beta1.KMSCryptoKeyStatus {
	return KMSCryptoKeyStatus_v1beta1_FromProto(mapCtx, in)
}

func KMSCryptoKeyVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyVersion) *krmkmsv1alpha1.KMSCryptoKeyVersionSpec {
	return KMSCryptoKeyVersionSpec_v1alpha1_FromProto(mapCtx, in)
}

func KMSCryptoKeyStatus_ToProto(mapCtx *direct.MapContext, in *krmkmsv1beta1.KMSCryptoKeyStatus) *pb.CryptoKey {
	return KMSCryptoKeyStatus_v1beta1_ToProto(mapCtx, in)
}

func KMSCryptoKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKey) *krmkmsv1beta1.KMSCryptoKeySpec {
	return KMSCryptoKeySpec_v1beta1_FromProto(mapCtx, in)
}

func KMSCryptoKeySpec_ToProto(mapCtx *direct.MapContext, in *krmkmsv1beta1.KMSCryptoKeySpec) *pb.CryptoKey {
	return KMSCryptoKeySpec_v1beta1_ToProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func KMSImportJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krmkmsv1beta1.KMSImportJobObservedState {
	return KMSImportJobObservedState_v1beta1_FromProto(mapCtx, in)
}

func KMSImportJobObservedState_ToProto(mapCtx *direct.MapContext, in *krmkmsv1beta1.KMSImportJobObservedState) *pb.ImportJob {
	return KMSImportJobObservedState_v1beta1_ToProto(mapCtx, in)
}

func KMSCryptoKeyVersionSpec_ToProto(mapCtx *direct.MapContext, in *krmkmsv1alpha1.KMSCryptoKeyVersionSpec) *pb.CryptoKeyVersion {
	return KMSCryptoKeyVersionSpec_v1alpha1_ToProto(mapCtx, in)
}

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

func CryptokeyversionCertChainsStatus_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation_CertificateChains) *krmkmsv1alpha1.CryptokeyversionCertChainsStatus {
	return CryptokeyversionCertChainsStatus_v1alpha1_FromProto(mapCtx, in)
}

func CryptokeyversionExternalProtectionLevelOptionsStatus_ToProto(mapCtx *direct.MapContext, in *krmkmsv1alpha1.CryptokeyversionExternalProtectionLevelOptionsStatus) *pb.ExternalProtectionLevelOptions {
	return CryptokeyversionExternalProtectionLevelOptionsStatus_v1alpha1_ToProto(mapCtx, in)
}

func CryptokeyversionExternalProtectionLevelOptionsStatus_FromProto(mapCtx *direct.MapContext, in *pb.ExternalProtectionLevelOptions) *krmkmsv1alpha1.CryptokeyversionExternalProtectionLevelOptionsStatus {
	return CryptokeyversionExternalProtectionLevelOptionsStatus_v1alpha1_FromProto(mapCtx, in)
}

func CryptokeyversionCertChainsStatus_ToProto(mapCtx *direct.MapContext, in *krmkmsv1alpha1.CryptokeyversionCertChainsStatus) *pb.KeyOperationAttestation_CertificateChains {
	return CryptokeyversionCertChainsStatus_v1alpha1_ToProto(mapCtx, in)
}

func KeyringimportjobAttestationStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmkmsv1alpha1.KeyringimportjobAttestationStatus) *pb.KeyOperationAttestation {
	return KeyringimportjobAttestationStatus_ToProto(mapCtx, in)
}

func KeyringimportjobAttestationStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation) *krmkmsv1alpha1.KeyringimportjobAttestationStatus {
	return KeyringimportjobAttestationStatus_FromProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func KMSKeyHandleSpec_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krmkmsv1beta1.KMSKeyHandleSpec {
	return KMSKeyHandleSpec_v1beta1_FromProto(mapCtx, in)
}

func KMSAutokeyConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krmkmsv1beta1.KMSAutokeyConfigObservedState {
	return KMSAutokeyConfigObservedState_v1beta1_FromProto(mapCtx, in)
}

func KMSKeyHandleSpec_ToProto(mapCtx *direct.MapContext, in *krmkmsv1beta1.KMSKeyHandleSpec) *pb.KeyHandle {
	return KMSKeyHandleSpec_v1beta1_ToProto(mapCtx, in)
}

func KMSKeyHandleObservedState_ToProto(mapCtx *direct.MapContext, in *krmkmsv1beta1.KMSKeyHandleObservedState) *pb.KeyHandle {
	return KMSKeyHandleObservedState_v1beta1_ToProto(mapCtx, in)
}

func KMSKeyHandleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krmkmsv1beta1.KMSKeyHandleObservedState {
	return KMSKeyHandleObservedState_v1beta1_FromProto(mapCtx, in)
}

func KMSAutokeyConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmkmsv1beta1.KMSAutokeyConfigObservedState) *pb.AutokeyConfig {
	return KMSAutokeyConfigObservedState_v1beta1_ToProto(mapCtx, in)
}
