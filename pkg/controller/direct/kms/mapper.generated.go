// Copyright 2025 Google LLC
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

// +generated:mapper
// krm.group: kms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.kms.v1

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AutokeyConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &krm.AutokeyConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.KeyProject = direct.LazyPtr(in.GetKeyProject())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func AutokeyConfig_ToProto(mapCtx *direct.MapContext, in *krm.AutokeyConfig) *pb.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutokeyConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.KeyProject = AutokeyConfig_KeyProject_ToProto(mapCtx, in.KeyProject)
	out.State = direct.Enum_ToProto[pb.AutokeyConfig_State](mapCtx, in.State)
	return out
}
func Certificate_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) *krm.Certificate {
	if in == nil {
		return nil
	}
	out := &krm.Certificate{}
	out.RawDer = in.GetRawDer()
	out.Parsed = direct.LazyPtr(in.GetParsed())
	out.Issuer = direct.LazyPtr(in.GetIssuer())
	out.Subject = direct.LazyPtr(in.GetSubject())
	// MISSING: SubjectAlternativeDNSNames
	// (near miss): "SubjectAlternativeDNSNames" vs "SubjectAlternativeDnsNames"
	out.NotBeforeTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNotBeforeTime())
	out.NotAfterTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNotAfterTime())
	out.SerialNumber = direct.LazyPtr(in.GetSerialNumber())
	out.Sha256Fingerprint = direct.LazyPtr(in.GetSha256Fingerprint())
	return out
}
func Certificate_ToProto(mapCtx *direct.MapContext, in *krm.Certificate) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	out.RawDer = in.RawDer
	out.Parsed = direct.ValueOf(in.Parsed)
	out.Issuer = direct.ValueOf(in.Issuer)
	out.Subject = direct.ValueOf(in.Subject)
	// MISSING: SubjectAlternativeDNSNames
	// (near miss): "SubjectAlternativeDNSNames" vs "SubjectAlternativeDnsNames"
	out.NotBeforeTime = direct.StringTimestamp_ToProto(mapCtx, in.NotBeforeTime)
	out.NotAfterTime = direct.StringTimestamp_ToProto(mapCtx, in.NotAfterTime)
	out.SerialNumber = direct.ValueOf(in.SerialNumber)
	out.Sha256Fingerprint = direct.ValueOf(in.Sha256Fingerprint)
	return out
}
func CryptoKey_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKey) *krm.CryptoKey {
	if in == nil {
		return nil
	}
	out := &krm.CryptoKey{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Primary = CryptoKeyVersion_FromProto(mapCtx, in.GetPrimary())
	out.Purpose = direct.Enum_FromProto(mapCtx, in.GetPurpose())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.NextRotationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextRotationTime())
	out.RotationPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRotationPeriod())
	out.VersionTemplate = CryptoKeyVersionTemplate_FromProto(mapCtx, in.GetVersionTemplate())
	out.Labels = in.Labels
	out.ImportOnly = direct.LazyPtr(in.GetImportOnly())
	out.DestroyScheduledDuration = direct.StringDuration_FromProto(mapCtx, in.GetDestroyScheduledDuration())
	out.CryptoKeyBackend = direct.LazyPtr(in.GetCryptoKeyBackend())
	out.KeyAccessJustificationsPolicy = KeyAccessJustificationsPolicy_FromProto(mapCtx, in.GetKeyAccessJustificationsPolicy())
	return out
}
func CryptoKey_ToProto(mapCtx *direct.MapContext, in *krm.CryptoKey) *pb.CryptoKey {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKey{}
	out.Name = direct.ValueOf(in.Name)
	out.Primary = CryptoKeyVersion_ToProto(mapCtx, in.Primary)
	out.Purpose = direct.Enum_ToProto[pb.CryptoKey_CryptoKeyPurpose](mapCtx, in.Purpose)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.NextRotationTime = direct.StringTimestamp_ToProto(mapCtx, in.NextRotationTime)
	if oneof := direct.StringDuration_ToProto(mapCtx, in.RotationPeriod); oneof != nil {
		out.RotationSchedule = &pb.CryptoKey_RotationPeriod{RotationPeriod: oneof}
	}
	out.VersionTemplate = CryptoKeyVersionTemplate_ToProto(mapCtx, in.VersionTemplate)
	out.Labels = in.Labels
	out.ImportOnly = direct.ValueOf(in.ImportOnly)
	out.DestroyScheduledDuration = direct.StringDuration_ToProto(mapCtx, in.DestroyScheduledDuration)
	out.CryptoKeyBackend = direct.ValueOf(in.CryptoKeyBackend)
	out.KeyAccessJustificationsPolicy = KeyAccessJustificationsPolicy_ToProto(mapCtx, in.KeyAccessJustificationsPolicy)
	return out
}
func CryptoKeyVersion_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyVersion) *krm.CryptoKeyVersion {
	if in == nil {
		return nil
	}
	out := &krm.CryptoKeyVersion{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	out.Algorithm = direct.Enum_FromProto(mapCtx, in.GetAlgorithm())
	out.Attestation = KeyOperationAttestation_FromProto(mapCtx, in.GetAttestation())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.GenerateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetGenerateTime())
	out.DestroyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDestroyTime())
	out.DestroyEventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDestroyEventTime())
	out.ImportJob = direct.LazyPtr(in.GetImportJob())
	out.ImportTime = direct.StringTimestamp_FromProto(mapCtx, in.GetImportTime())
	out.ImportFailureReason = direct.LazyPtr(in.GetImportFailureReason())
	out.GenerationFailureReason = direct.LazyPtr(in.GetGenerationFailureReason())
	out.ExternalDestructionFailureReason = direct.LazyPtr(in.GetExternalDestructionFailureReason())
	out.ExternalProtectionLevelOptions = ExternalProtectionLevelOptions_FromProto(mapCtx, in.GetExternalProtectionLevelOptions())
	out.ReimportEligible = direct.LazyPtr(in.GetReimportEligible())
	return out
}
func CryptoKeyVersion_ToProto(mapCtx *direct.MapContext, in *krm.CryptoKeyVersion) *pb.CryptoKeyVersion {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKeyVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.CryptoKeyVersion_CryptoKeyVersionState](mapCtx, in.State)
	out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	out.Algorithm = direct.Enum_ToProto[pb.CryptoKeyVersion_CryptoKeyVersionAlgorithm](mapCtx, in.Algorithm)
	out.Attestation = KeyOperationAttestation_ToProto(mapCtx, in.Attestation)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.GenerateTime = direct.StringTimestamp_ToProto(mapCtx, in.GenerateTime)
	out.DestroyTime = direct.StringTimestamp_ToProto(mapCtx, in.DestroyTime)
	out.DestroyEventTime = direct.StringTimestamp_ToProto(mapCtx, in.DestroyEventTime)
	out.ImportJob = direct.ValueOf(in.ImportJob)
	out.ImportTime = direct.StringTimestamp_ToProto(mapCtx, in.ImportTime)
	out.ImportFailureReason = direct.ValueOf(in.ImportFailureReason)
	out.GenerationFailureReason = direct.ValueOf(in.GenerationFailureReason)
	out.ExternalDestructionFailureReason = direct.ValueOf(in.ExternalDestructionFailureReason)
	out.ExternalProtectionLevelOptions = ExternalProtectionLevelOptions_ToProto(mapCtx, in.ExternalProtectionLevelOptions)
	out.ReimportEligible = direct.ValueOf(in.ReimportEligible)
	return out
}
func CryptoKeyVersionTemplate_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyVersionTemplate) *krm.CryptoKeyVersionTemplate {
	if in == nil {
		return nil
	}
	out := &krm.CryptoKeyVersionTemplate{}
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	out.Algorithm = direct.Enum_FromProto(mapCtx, in.GetAlgorithm())
	return out
}
func CryptoKeyVersionTemplate_ToProto(mapCtx *direct.MapContext, in *krm.CryptoKeyVersionTemplate) *pb.CryptoKeyVersionTemplate {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKeyVersionTemplate{}
	out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	out.Algorithm = direct.Enum_ToProto[pb.CryptoKeyVersion_CryptoKeyVersionAlgorithm](mapCtx, in.Algorithm)
	return out
}
func Digest_FromProto(mapCtx *direct.MapContext, in *pb.Digest) *krm.Digest {
	if in == nil {
		return nil
	}
	out := &krm.Digest{}
	out.Sha256 = in.GetSha256()
	out.Sha384 = in.GetSha384()
	out.Sha512 = in.GetSha512()
	return out
}
func Digest_ToProto(mapCtx *direct.MapContext, in *krm.Digest) *pb.Digest {
	if in == nil {
		return nil
	}
	out := &pb.Digest{}
	if oneof := Digest_Sha256_ToProto(mapCtx, in.Sha256); oneof != nil {
		out.Digest = oneof
	}
	if oneof := Digest_Sha384_ToProto(mapCtx, in.Sha384); oneof != nil {
		out.Digest = oneof
	}
	if oneof := Digest_Sha512_ToProto(mapCtx, in.Sha512); oneof != nil {
		out.Digest = oneof
	}
	return out
}
func EkmConfig_FromProto(mapCtx *direct.MapContext, in *pb.EkmConfig) *krm.EkmConfig {
	if in == nil {
		return nil
	}
	out := &krm.EkmConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DefaultEkmConnection = direct.LazyPtr(in.GetDefaultEkmConnection())
	return out
}
func EkmConfig_ToProto(mapCtx *direct.MapContext, in *krm.EkmConfig) *pb.EkmConfig {
	if in == nil {
		return nil
	}
	out := &pb.EkmConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.DefaultEkmConnection = direct.ValueOf(in.DefaultEkmConnection)
	return out
}
func EkmConnection_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection) *krm.EkmConnection {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnection{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ServiceResolvers = direct.Slice_FromProto(mapCtx, in.ServiceResolvers, EkmConnection_ServiceResolver_FromProto)
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.KeyManagementMode = direct.Enum_FromProto(mapCtx, in.GetKeyManagementMode())
	out.CryptoSpacePath = direct.LazyPtr(in.GetCryptoSpacePath())
	return out
}
func EkmConnection_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnection) *pb.EkmConnection {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.ServiceResolvers = direct.Slice_ToProto(mapCtx, in.ServiceResolvers, EkmConnection_ServiceResolver_ToProto)
	out.Etag = direct.ValueOf(in.Etag)
	out.KeyManagementMode = direct.Enum_ToProto[pb.EkmConnection_KeyManagementMode](mapCtx, in.KeyManagementMode)
	out.CryptoSpacePath = direct.ValueOf(in.CryptoSpacePath)
	return out
}
func EkmConnection_ServiceResolver_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection_ServiceResolver) *krm.EkmConnection_ServiceResolver {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnection_ServiceResolver{}
	out.ServiceDirectoryService = direct.LazyPtr(in.GetServiceDirectoryService())
	out.EndpointFilter = direct.LazyPtr(in.GetEndpointFilter())
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.ServerCertificates = direct.Slice_FromProto(mapCtx, in.ServerCertificates, Certificate_FromProto)
	return out
}
func EkmConnection_ServiceResolver_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnection_ServiceResolver) *pb.EkmConnection_ServiceResolver {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection_ServiceResolver{}
	out.ServiceDirectoryService = direct.ValueOf(in.ServiceDirectoryService)
	out.EndpointFilter = direct.ValueOf(in.EndpointFilter)
	out.Hostname = direct.ValueOf(in.Hostname)
	out.ServerCertificates = direct.Slice_ToProto(mapCtx, in.ServerCertificates, Certificate_ToProto)
	return out
}
func ExternalProtectionLevelOptions_FromProto(mapCtx *direct.MapContext, in *pb.ExternalProtectionLevelOptions) *krm.ExternalProtectionLevelOptions {
	if in == nil {
		return nil
	}
	out := &krm.ExternalProtectionLevelOptions{}
	// MISSING: ExternalKeyURI
	// (near miss): "ExternalKeyURI" vs "ExternalKeyUri"
	out.EkmConnectionKeyPath = direct.LazyPtr(in.GetEkmConnectionKeyPath())
	return out
}
func ExternalProtectionLevelOptions_ToProto(mapCtx *direct.MapContext, in *krm.ExternalProtectionLevelOptions) *pb.ExternalProtectionLevelOptions {
	if in == nil {
		return nil
	}
	out := &pb.ExternalProtectionLevelOptions{}
	// MISSING: ExternalKeyURI
	// (near miss): "ExternalKeyURI" vs "ExternalKeyUri"
	out.EkmConnectionKeyPath = direct.ValueOf(in.EkmConnectionKeyPath)
	return out
}
func ImportJob_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.ImportJob {
	if in == nil {
		return nil
	}
	out := &krm.ImportJob{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ImportMethod = direct.Enum_FromProto(mapCtx, in.GetImportMethod())
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.GenerateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetGenerateTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.ExpireEventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireEventTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.PublicKey = ImportJob_WrappingPublicKey_FromProto(mapCtx, in.GetPublicKey())
	out.Attestation = KeyOperationAttestation_FromProto(mapCtx, in.GetAttestation())
	return out
}
func ImportJob_ToProto(mapCtx *direct.MapContext, in *krm.ImportJob) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	out.Name = direct.ValueOf(in.Name)
	out.ImportMethod = direct.Enum_ToProto[pb.ImportJob_ImportMethod](mapCtx, in.ImportMethod)
	out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.GenerateTime = direct.StringTimestamp_ToProto(mapCtx, in.GenerateTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.ExpireEventTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireEventTime)
	out.State = direct.Enum_ToProto[pb.ImportJob_ImportJobState](mapCtx, in.State)
	out.PublicKey = ImportJob_WrappingPublicKey_ToProto(mapCtx, in.PublicKey)
	out.Attestation = KeyOperationAttestation_ToProto(mapCtx, in.Attestation)
	return out
}
func ImportJob_WrappingPublicKey_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob_WrappingPublicKey) *krm.ImportJob_WrappingPublicKey {
	if in == nil {
		return nil
	}
	out := &krm.ImportJob_WrappingPublicKey{}
	out.Pem = direct.LazyPtr(in.GetPem())
	return out
}
func ImportJob_WrappingPublicKey_ToProto(mapCtx *direct.MapContext, in *krm.ImportJob_WrappingPublicKey) *pb.ImportJob_WrappingPublicKey {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob_WrappingPublicKey{}
	out.Pem = direct.ValueOf(in.Pem)
	return out
}
func ImportJob_WrappingPublicKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob_WrappingPublicKey) *krm.ImportJob_WrappingPublicKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ImportJob_WrappingPublicKeyObservedState{}
	out.Pem = direct.LazyPtr(in.GetPem())
	return out
}
func ImportJob_WrappingPublicKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ImportJob_WrappingPublicKeyObservedState) *pb.ImportJob_WrappingPublicKey {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob_WrappingPublicKey{}
	out.Pem = direct.ValueOf(in.Pem)
	return out
}
func KMSAutokeyConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.KMSAutokeyConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KMSAutokeyConfigObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func KMSAutokeyConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KMSAutokeyConfigObservedState) *pb.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutokeyConfig{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.AutokeyConfig_State](mapCtx, in.State)
	return out
}
func KMSAutokeyConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.KMSAutokeyConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSAutokeyConfigSpec{}
	// MISSING: Name
	if in.GetKeyProject() != "" {
		out.KeyProjectRef = &refs.ProjectRef{External: in.GetKeyProject()}
	}
	return out
}
func KMSAutokeyConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSAutokeyConfigSpec) *pb.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutokeyConfig{}
	// MISSING: Name
	if in.KeyProjectRef != nil {
		out.KeyProject = in.KeyProjectRef.External
	}
	return out
}
func KMSImportJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.KMSImportJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KMSImportJobObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.GenerateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetGenerateTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.ExpireEventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireEventTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.PublicKey = ImportJob_WrappingPublicKeyObservedState_FromProto(mapCtx, in.GetPublicKey())
	out.Attestation = KeyOperationAttestationObservedState_FromProto(mapCtx, in.GetAttestation())
	return out
}
func KMSImportJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KMSImportJobObservedState) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.GenerateTime = direct.StringTimestamp_ToProto(mapCtx, in.GenerateTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.ExpireEventTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireEventTime)
	out.State = direct.Enum_ToProto[pb.ImportJob_ImportJobState](mapCtx, in.State)
	out.PublicKey = ImportJob_WrappingPublicKeyObservedState_ToProto(mapCtx, in.PublicKey)
	out.Attestation = KeyOperationAttestationObservedState_ToProto(mapCtx, in.Attestation)
	return out
}
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
func KMSImportJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSImportJobSpec) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	// MISSING: Name
	out.ImportMethod = direct.Enum_ToProto[pb.ImportJob_ImportMethod](mapCtx, in.ImportMethod)
	out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	return out
}
func KMSKeyHandleSpec_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krm.KMSKeyHandleSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSKeyHandleSpec{}
	// MISSING: Name
	// MISSING: KMSKey
	out.ResourceTypeSelector = direct.LazyPtr(in.GetResourceTypeSelector())
	return out
}
func KMSKeyHandleSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSKeyHandleSpec) *pb.KeyHandle {
	if in == nil {
		return nil
	}
	out := &pb.KeyHandle{}
	// MISSING: Name
	// MISSING: KMSKey
	out.ResourceTypeSelector = direct.ValueOf(in.ResourceTypeSelector)
	return out
}
func KMSKeyRingSpec_FromProto(mapCtx *direct.MapContext, in *pb.KeyRing) *krm.KMSKeyRingSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSKeyRingSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func KMSKeyRingSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSKeyRingSpec) *pb.KeyRing {
	if in == nil {
		return nil
	}
	out := &pb.KeyRing{}
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func KMSKeyRingStatus_FromProto(mapCtx *direct.MapContext, in *pb.KeyRing) *krm.KMSKeyRingStatus {
	if in == nil {
		return nil
	}
	out := &krm.KMSKeyRingStatus{}
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func KMSKeyRingStatus_ToProto(mapCtx *direct.MapContext, in *krm.KMSKeyRingStatus) *pb.KeyRing {
	if in == nil {
		return nil
	}
	out := &pb.KeyRing{}
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func KeyAccessJustificationsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.KeyAccessJustificationsPolicy) *krm.KeyAccessJustificationsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.KeyAccessJustificationsPolicy{}
	out.AllowedAccessReasons = direct.EnumSlice_FromProto(mapCtx, in.AllowedAccessReasons)
	return out
}
func KeyAccessJustificationsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.KeyAccessJustificationsPolicy) *pb.KeyAccessJustificationsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.KeyAccessJustificationsPolicy{}
	out.AllowedAccessReasons = direct.EnumSlice_ToProto[pb.AccessReason](mapCtx, in.AllowedAccessReasons)
	return out
}
func KeyHandle_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krm.KeyHandle {
	if in == nil {
		return nil
	}
	out := &krm.KeyHandle{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: KMSKey
	// (near miss): "KMSKey" vs "KmsKey"
	out.ResourceTypeSelector = direct.LazyPtr(in.GetResourceTypeSelector())
	return out
}
func KeyHandle_ToProto(mapCtx *direct.MapContext, in *krm.KeyHandle) *pb.KeyHandle {
	if in == nil {
		return nil
	}
	out := &pb.KeyHandle{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: KMSKey
	// (near miss): "KMSKey" vs "KmsKey"
	out.ResourceTypeSelector = direct.ValueOf(in.ResourceTypeSelector)
	return out
}
func KeyOperationAttestation_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation) *krm.KeyOperationAttestation {
	if in == nil {
		return nil
	}
	out := &krm.KeyOperationAttestation{}
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	out.Content = in.GetContent()
	out.CertChains = KeyOperationAttestation_CertificateChains_FromProto(mapCtx, in.GetCertChains())
	return out
}
func KeyOperationAttestation_ToProto(mapCtx *direct.MapContext, in *krm.KeyOperationAttestation) *pb.KeyOperationAttestation {
	if in == nil {
		return nil
	}
	out := &pb.KeyOperationAttestation{}
	out.Format = direct.Enum_ToProto[pb.KeyOperationAttestation_AttestationFormat](mapCtx, in.Format)
	out.Content = in.Content
	out.CertChains = KeyOperationAttestation_CertificateChains_ToProto(mapCtx, in.CertChains)
	return out
}
func KeyOperationAttestationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation) *krm.KeyOperationAttestationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KeyOperationAttestationObservedState{}
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	out.Content = in.GetContent()
	out.CertChains = KeyOperationAttestation_CertificateChains_FromProto(mapCtx, in.GetCertChains())
	return out
}
func KeyOperationAttestationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KeyOperationAttestationObservedState) *pb.KeyOperationAttestation {
	if in == nil {
		return nil
	}
	out := &pb.KeyOperationAttestation{}
	out.Format = direct.Enum_ToProto[pb.KeyOperationAttestation_AttestationFormat](mapCtx, in.Format)
	out.Content = in.Content
	out.CertChains = KeyOperationAttestation_CertificateChains_ToProto(mapCtx, in.CertChains)
	return out
}
func KeyOperationAttestation_CertificateChains_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation_CertificateChains) *krm.KeyOperationAttestation_CertificateChains {
	if in == nil {
		return nil
	}
	out := &krm.KeyOperationAttestation_CertificateChains{}
	out.CaviumCerts = in.CaviumCerts
	out.GoogleCardCerts = in.GoogleCardCerts
	out.GooglePartitionCerts = in.GooglePartitionCerts
	return out
}
func KeyOperationAttestation_CertificateChains_ToProto(mapCtx *direct.MapContext, in *krm.KeyOperationAttestation_CertificateChains) *pb.KeyOperationAttestation_CertificateChains {
	if in == nil {
		return nil
	}
	out := &pb.KeyOperationAttestation_CertificateChains{}
	out.CaviumCerts = in.CaviumCerts
	out.GoogleCardCerts = in.GoogleCardCerts
	out.GooglePartitionCerts = in.GooglePartitionCerts
	return out
}
func PublicKey_FromProto(mapCtx *direct.MapContext, in *pb.PublicKey) *krm.PublicKey {
	if in == nil {
		return nil
	}
	out := &krm.PublicKey{}
	out.Pem = direct.LazyPtr(in.GetPem())
	out.Algorithm = direct.Enum_FromProto(mapCtx, in.GetAlgorithm())
	out.PemCrc32c = direct.Int64Value_FromProto(mapCtx, in.GetPemCrc32c())
	out.Name = direct.LazyPtr(in.GetName())
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	return out
}
func PublicKey_ToProto(mapCtx *direct.MapContext, in *krm.PublicKey) *pb.PublicKey {
	if in == nil {
		return nil
	}
	out := &pb.PublicKey{}
	out.Pem = direct.ValueOf(in.Pem)
	out.Algorithm = direct.Enum_ToProto[pb.CryptoKeyVersion_CryptoKeyVersionAlgorithm](mapCtx, in.Algorithm)
	out.PemCrc32c = direct.Int64Value_ToProto(mapCtx, in.PemCrc32c)
	out.Name = direct.ValueOf(in.Name)
	out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	return out
}
