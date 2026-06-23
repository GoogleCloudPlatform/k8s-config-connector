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
	"strings"

	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// KMSAutokeyConfigSpec_FromProto converts the protobuf AutokeyConfig to the KRM type.
// This is handcoded because the GCP proto for AutokeyConfig from the googleapis submodule does not contain KeyProjectResolutionMode,
// but the Go SDK's kmspb does, so we override it to include mapping of KeyProjectResolutionMode.
func KMSAutokeyConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.KMSAutokeyConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSAutokeyConfigSpec{}
	if in.GetKeyProject() != "" {
		out.KeyProjectRef = &refsv1beta1.ProjectRef{External: in.GetKeyProject()}
	}
	out.KeyProjectResolutionMode = direct.Enum_FromProto(mapCtx, in.GetKeyProjectResolutionMode())
	return out
}

// KMSAutokeyConfigSpec_ToProto converts the KRM type KMSAutokeyConfigSpec to the protobuf representation.
// This is handcoded because the GCP proto for AutokeyConfig from the googleapis submodule does not contain KeyProjectResolutionMode,
// but the Go SDK's kmspb does, so we override it to include mapping of KeyProjectResolutionMode.
func KMSAutokeyConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSAutokeyConfigSpec) *pb.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutokeyConfig{}
	if in.KeyProjectRef != nil {
		out.KeyProject = in.KeyProjectRef.External
	}
	if in.KeyProjectResolutionMode != nil {
		out.KeyProjectResolutionMode = direct.Enum_ToProto[pb.AutokeyConfig_KeyProjectResolutionMode](mapCtx, in.KeyProjectResolutionMode)
	}
	return out
}

// KMSCryptoKeyVersionStatus_FromProto converts the protobuf CryptoKeyVersion status fields to the KRM status.
// This is handcoded because the baseline KRM schema nests ExternalProtectionLevelOptions inside Attestation,
// whereas in the GCP direct proto schema it is a direct field on CryptoKeyVersion.
func KMSCryptoKeyVersionStatus_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyVersion) *krmv1alpha1.KMSCryptoKeyVersionStatus {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.KMSCryptoKeyVersionStatus{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	out.Algorithm = direct.Enum_FromProto(mapCtx, in.GetAlgorithm())
	out.GenerateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetGenerateTime())

	// Map attestation if present in proto
	var attestation *krmv1alpha1.CryptokeyversionAttestationStatus
	if v := in.GetAttestation(); v != nil {
		attestation = CryptokeyversionAttestationStatus_FromProto(mapCtx, v)
	}

	// Map external_protection_level_options (nested inside attestation in KRM)
	if in.GetExternalProtectionLevelOptions() != nil {
		if attestation == nil {
			attestation = &krmv1alpha1.CryptokeyversionAttestationStatus{}
		}
		attestation.ExternalProtectionLevelOptions = CryptokeyversionExternalProtectionLevelOptionsStatus_FromProto(mapCtx, in.GetExternalProtectionLevelOptions())
	}

	if attestation != nil {
		out.Attestation = []krmv1alpha1.CryptokeyversionAttestationStatus{*attestation}
	}

	return out
}

// KMSCryptoKeyVersionStatus_ToProto converts the KRM status KMSCryptoKeyVersionStatus to protobuf.
// This is handcoded because the baseline KRM schema nests ExternalProtectionLevelOptions inside Attestation,
// whereas in the GCP direct proto schema it is a direct field on CryptoKeyVersion.
func KMSCryptoKeyVersionStatus_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.KMSCryptoKeyVersionStatus) *pb.CryptoKeyVersion {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKeyVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	out.Algorithm = direct.Enum_ToProto[pb.CryptoKeyVersion_CryptoKeyVersionAlgorithm](mapCtx, in.Algorithm)
	out.GenerateTime = direct.StringTimestamp_ToProto(mapCtx, in.GenerateTime)

	if len(in.Attestation) > 0 {
		att := &in.Attestation[0]
		out.Attestation = CryptokeyversionAttestationStatus_ToProto(mapCtx, att)
		if att.ExternalProtectionLevelOptions != nil {
			out.ExternalProtectionLevelOptions = CryptokeyversionExternalProtectionLevelOptionsStatus_ToProto(mapCtx, att.ExternalProtectionLevelOptions)
		}
	}

	return out
}

// CryptoKeyVersionTemplate_FromProto converts the protobuf CryptoKeyVersionTemplate to the KRM type.
// This is handcoded because the KRM's Algorithm field is a required non-pointer string,
// but the generated Enum_FromProto helper returns a pointer string.
func CryptoKeyVersionTemplate_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyVersionTemplate) *krm.CryptoKeyVersionTemplate {
	if in == nil {
		return nil
	}
	out := &krm.CryptoKeyVersionTemplate{}
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	if algoPtr := direct.Enum_FromProto(mapCtx, in.GetAlgorithm()); algoPtr != nil {
		out.Algorithm = *algoPtr
	}
	return out
}

// CryptoKeyVersionTemplate_ToProto converts the KRM type CryptoKeyVersionTemplate to the protobuf representation.
// This is handcoded because the KRM's Algorithm field is a non-pointer string,
// but the generated Enum_ToProto helper expects a pointer string.
func CryptoKeyVersionTemplate_ToProto(mapCtx *direct.MapContext, in *krm.CryptoKeyVersionTemplate) *pb.CryptoKeyVersionTemplate {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKeyVersionTemplate{}
	out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	out.Algorithm = direct.Enum_ToProto[pb.CryptoKeyVersion_CryptoKeyVersionAlgorithm](mapCtx, &in.Algorithm)
	return out
}

// KMSImportJobSpec_FromProto converts the protobuf ImportJob to the KRM type KMSImportJobSpec.
// This is handcoded because it handles specific casing transformations or fields that are not automatically mapped.
func KMSImportJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.KMSImportJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSImportJobSpec{}
	// MISSING: Name
	out.ImportMethod = direct.Enum_FromProto(mapCtx, in.GetImportMethod())
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	return out
}

// KMSImportJobSpec_ToProto converts the KRM type KMSImportJobSpec to the protobuf ImportJob representation.
// This is handcoded to transform enum strings to uppercase to handle mixed-case user input.
func KMSImportJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSImportJobSpec) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	// MISSING: Name

	if in.ImportMethod != nil {
		importMethodUpper := strings.ToUpper(*in.ImportMethod)
		out.ImportMethod = direct.Enum_ToProto[pb.ImportJob_ImportMethod](mapCtx, &importMethodUpper)
	} else {
		out.ImportMethod = direct.Enum_ToProto[pb.ImportJob_ImportMethod](mapCtx, in.ImportMethod)
	}
	if in.ProtectionLevel != nil {
		protectionLevelUpper := strings.ToUpper(*in.ProtectionLevel)
		out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, &protectionLevelUpper)
	} else {
		out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	}
	return out
}

// CryptokeyversionAttestationStatus_FromProto converts KeyOperationAttestation to KRM representation.
// This is handcoded because the 'content' field is represented as []byte in proto, but is a pointer-string in KRM.
func CryptokeyversionAttestationStatus_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation) *krmv1alpha1.CryptokeyversionAttestationStatus {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.CryptokeyversionAttestationStatus{}
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	if in.GetContent() != nil {
		out.Content = direct.LazyPtr(string(in.GetContent()))
	}
	out.CertChains = CryptokeyversionCertChainsStatus_FromProto(mapCtx, in.GetCertChains())
	return out
}

// CryptokeyversionAttestationStatus_ToProto converts KRM CryptokeyversionAttestationStatus to KeyOperationAttestation.
// This is handcoded because the 'content' field is represented as []byte in proto, but is a pointer-string in KRM.
func CryptokeyversionAttestationStatus_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.CryptokeyversionAttestationStatus) *pb.KeyOperationAttestation {
	if in == nil {
		return nil
	}
	out := &pb.KeyOperationAttestation{}
	out.Format = direct.Enum_ToProto[pb.KeyOperationAttestation_AttestationFormat](mapCtx, in.Format)
	if in.Content != nil {
		out.Content = []byte(*in.Content)
	}
	out.CertChains = CryptokeyversionCertChainsStatus_ToProto(mapCtx, in.CertChains)
	return out
}

// KMSSecretCiphertextAPI represents the client-side API representation of KMSSecretCiphertext
// for the KRM fuzzer to round-trip against.
type KMSSecretCiphertextAPI struct {
	CryptoKey                   string
	Plaintext                   string
	AdditionalAuthenticatedData *string
	ResourceID                  *string
	Ciphertext                  *string
}

// KMSSecretCiphertextSpec_FromAPI maps KMSSecretCiphertextAPI Spec fields to KRM Spec representation.
func KMSSecretCiphertextSpec_FromAPI(mapCtx *direct.MapContext, in *KMSSecretCiphertextAPI) *krmv1alpha1.KMSSecretCiphertextSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.KMSSecretCiphertextSpec{
		CryptoKey:  in.CryptoKey,
		ResourceID: in.ResourceID,
	}
	out.Plaintext = secret.Legacy{
		Value: direct.LazyPtr(in.Plaintext),
	}
	if in.AdditionalAuthenticatedData != nil {
		out.AdditionalAuthenticatedData = &secret.Legacy{
			Value: in.AdditionalAuthenticatedData,
		}
	}
	return out
}

// KMSSecretCiphertextSpec_ToAPI maps KRM KMSSecretCiphertextSpec representation to KMSSecretCiphertextAPI.
func KMSSecretCiphertextSpec_ToAPI(mapCtx *direct.MapContext, in *krmv1alpha1.KMSSecretCiphertextSpec) *KMSSecretCiphertextAPI {
	if in == nil {
		return nil
	}
	out := &KMSSecretCiphertextAPI{
		CryptoKey:  in.CryptoKey,
		ResourceID: in.ResourceID,
	}
	if in.Plaintext.Value != nil {
		out.Plaintext = *in.Plaintext.Value
	}
	if in.AdditionalAuthenticatedData != nil && in.AdditionalAuthenticatedData.Value != nil {
		out.AdditionalAuthenticatedData = in.AdditionalAuthenticatedData.Value
	}
	return out
}

// KMSSecretCiphertextStatus_FromAPI maps KMSSecretCiphertextAPI Status fields to KRM Status representation.
func KMSSecretCiphertextStatus_FromAPI(mapCtx *direct.MapContext, in *KMSSecretCiphertextAPI) *krmv1alpha1.KMSSecretCiphertextStatus {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.KMSSecretCiphertextStatus{
		Ciphertext: in.Ciphertext,
	}
	return out
}

// KMSSecretCiphertextStatus_ToAPI maps KRM KMSSecretCiphertextStatus representation to KMSSecretCiphertextAPI.
func KMSSecretCiphertextStatus_ToAPI(mapCtx *direct.MapContext, in *krmv1alpha1.KMSSecretCiphertextStatus) *KMSSecretCiphertextAPI {
	if in == nil {
		return nil
	}
	out := &KMSSecretCiphertextAPI{
		Ciphertext: in.Ciphertext,
	}
	return out
}
