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

package security

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/security/privateca/apiv1beta1/privatecapb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/security/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CertificateAuthority_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority) *krm.CertificateAuthority {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority{}
	// MISSING: Name
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	out.Config = CertificateConfig_FromProto(mapCtx, in.GetConfig())
	out.Lifetime = direct.StringDuration_FromProto(mapCtx, in.GetLifetime())
	out.KeySpec = CertificateAuthority_KeyVersionSpec_FromProto(mapCtx, in.GetKeySpec())
	out.CertificatePolicy = CertificateAuthority_CertificateAuthorityPolicy_FromProto(mapCtx, in.GetCertificatePolicy())
	out.IssuingOptions = CertificateAuthority_IssuingOptions_FromProto(mapCtx, in.GetIssuingOptions())
	out.SubordinateConfig = SubordinateConfig_FromProto(mapCtx, in.GetSubordinateConfig())
	// MISSING: State
	// MISSING: PemCaCertificates
	// MISSING: CaCertificateDescriptions
	out.GcsBucket = direct.LazyPtr(in.GetGcsBucket())
	// MISSING: AccessUrls
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	return out
}
func CertificateAuthority_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority) *pb.CertificateAuthority {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority{}
	// MISSING: Name
	out.Type = direct.Enum_ToProto[pb.CertificateAuthority_Type](mapCtx, in.Type)
	out.Tier = direct.Enum_ToProto[pb.CertificateAuthority_Tier](mapCtx, in.Tier)
	out.Config = CertificateConfig_ToProto(mapCtx, in.Config)
	out.Lifetime = direct.StringDuration_ToProto(mapCtx, in.Lifetime)
	out.KeySpec = CertificateAuthority_KeyVersionSpec_ToProto(mapCtx, in.KeySpec)
	out.CertificatePolicy = CertificateAuthority_CertificateAuthorityPolicy_ToProto(mapCtx, in.CertificatePolicy)
	out.IssuingOptions = CertificateAuthority_IssuingOptions_ToProto(mapCtx, in.IssuingOptions)
	out.SubordinateConfig = SubordinateConfig_ToProto(mapCtx, in.SubordinateConfig)
	// MISSING: State
	// MISSING: PemCaCertificates
	// MISSING: CaCertificateDescriptions
	out.GcsBucket = direct.ValueOf(in.GcsBucket)
	// MISSING: AccessUrls
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	return out
}
func CertificateAuthorityObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority) *krm.CertificateAuthorityObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthorityObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Type
	// MISSING: Tier
	// MISSING: Config
	// MISSING: Lifetime
	// MISSING: KeySpec
	// MISSING: CertificatePolicy
	// MISSING: IssuingOptions
	// MISSING: SubordinateConfig
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.PemCaCertificates = in.PemCaCertificates
	out.CaCertificateDescriptions = direct.Slice_FromProto(mapCtx, in.CaCertificateDescriptions, CertificateDescription_FromProto)
	// MISSING: GcsBucket
	out.AccessUrls = CertificateAuthority_AccessUrls_FromProto(mapCtx, in.GetAccessUrls())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	// MISSING: Labels
	return out
}
func CertificateAuthorityObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthorityObservedState) *pb.CertificateAuthority {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Type
	// MISSING: Tier
	// MISSING: Config
	// MISSING: Lifetime
	// MISSING: KeySpec
	// MISSING: CertificatePolicy
	// MISSING: IssuingOptions
	// MISSING: SubordinateConfig
	out.State = direct.Enum_ToProto[pb.CertificateAuthority_State](mapCtx, in.State)
	out.PemCaCertificates = in.PemCaCertificates
	out.CaCertificateDescriptions = direct.Slice_ToProto(mapCtx, in.CaCertificateDescriptions, CertificateDescription_ToProto)
	// MISSING: GcsBucket
	out.AccessUrls = CertificateAuthority_AccessUrls_ToProto(mapCtx, in.AccessUrls)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	// MISSING: Labels
	return out
}
func CertificateAuthority_AccessUrls_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_AccessUrls) *krm.CertificateAuthority_AccessUrls {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_AccessUrls{}
	out.CaCertificateAccessURL = direct.LazyPtr(in.GetCaCertificateAccessUrl())
	out.CrlAccessURL = direct.LazyPtr(in.GetCrlAccessUrl())
	return out
}
func CertificateAuthority_AccessUrls_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_AccessUrls) *pb.CertificateAuthority_AccessUrls {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_AccessUrls{}
	out.CaCertificateAccessUrl = direct.ValueOf(in.CaCertificateAccessURL)
	out.CrlAccessUrl = direct.ValueOf(in.CrlAccessURL)
	return out
}
func CertificateAuthority_CertificateAuthorityPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_CertificateAuthorityPolicy) *krm.CertificateAuthority_CertificateAuthorityPolicy {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_CertificateAuthorityPolicy{}
	out.AllowedConfigList = CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList_FromProto(mapCtx, in.GetAllowedConfigList())
	out.OverwriteConfigValues = ReusableConfigWrapper_FromProto(mapCtx, in.GetOverwriteConfigValues())
	out.AllowedLocationsAndOrganizations = direct.Slice_FromProto(mapCtx, in.AllowedLocationsAndOrganizations, Subject_FromProto)
	out.AllowedCommonNames = in.AllowedCommonNames
	out.AllowedSans = CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames_FromProto(mapCtx, in.GetAllowedSans())
	out.MaximumLifetime = direct.StringDuration_FromProto(mapCtx, in.GetMaximumLifetime())
	out.AllowedIssuanceModes = CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes_FromProto(mapCtx, in.GetAllowedIssuanceModes())
	return out
}
func CertificateAuthority_CertificateAuthorityPolicy_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_CertificateAuthorityPolicy) *pb.CertificateAuthority_CertificateAuthorityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_CertificateAuthorityPolicy{}
	if oneof := CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList_ToProto(mapCtx, in.AllowedConfigList); oneof != nil {
		out.ConfigPolicy = &pb.CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList_{AllowedConfigList: oneof}
	}
	if oneof := ReusableConfigWrapper_ToProto(mapCtx, in.OverwriteConfigValues); oneof != nil {
		out.ConfigPolicy = &pb.CertificateAuthority_CertificateAuthorityPolicy_OverwriteConfigValues{OverwriteConfigValues: oneof}
	}
	out.AllowedLocationsAndOrganizations = direct.Slice_ToProto(mapCtx, in.AllowedLocationsAndOrganizations, Subject_ToProto)
	out.AllowedCommonNames = in.AllowedCommonNames
	out.AllowedSans = CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames_ToProto(mapCtx, in.AllowedSans)
	out.MaximumLifetime = direct.StringDuration_ToProto(mapCtx, in.MaximumLifetime)
	out.AllowedIssuanceModes = CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes_ToProto(mapCtx, in.AllowedIssuanceModes)
	return out
}
func CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList) *krm.CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList{}
	out.AllowedConfigValues = direct.Slice_FromProto(mapCtx, in.AllowedConfigValues, ReusableConfigWrapper_FromProto)
	return out
}
func CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList) *pb.CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList{}
	out.AllowedConfigValues = direct.Slice_ToProto(mapCtx, in.AllowedConfigValues, ReusableConfigWrapper_ToProto)
	return out
}
func CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames) *krm.CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames{}
	out.AllowedDnsNames = in.AllowedDnsNames
	out.AllowedUris = in.AllowedUris
	out.AllowedEmailAddresses = in.AllowedEmailAddresses
	out.AllowedIps = in.AllowedIps
	out.AllowGlobbingDnsWildcards = direct.LazyPtr(in.GetAllowGlobbingDnsWildcards())
	out.AllowCustomSans = direct.LazyPtr(in.GetAllowCustomSans())
	return out
}
func CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames) *pb.CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames{}
	out.AllowedDnsNames = in.AllowedDnsNames
	out.AllowedUris = in.AllowedUris
	out.AllowedEmailAddresses = in.AllowedEmailAddresses
	out.AllowedIps = in.AllowedIps
	out.AllowGlobbingDnsWildcards = direct.ValueOf(in.AllowGlobbingDnsWildcards)
	out.AllowCustomSans = direct.ValueOf(in.AllowCustomSans)
	return out
}
func CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes) *krm.CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes{}
	out.AllowCsrBasedIssuance = direct.LazyPtr(in.GetAllowCsrBasedIssuance())
	out.AllowConfigBasedIssuance = direct.LazyPtr(in.GetAllowConfigBasedIssuance())
	return out
}
func CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes) *pb.CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes{}
	out.AllowCsrBasedIssuance = direct.ValueOf(in.AllowCsrBasedIssuance)
	out.AllowConfigBasedIssuance = direct.ValueOf(in.AllowConfigBasedIssuance)
	return out
}
func CertificateAuthority_IssuingOptions_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_IssuingOptions) *krm.CertificateAuthority_IssuingOptions {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_IssuingOptions{}
	out.IncludeCaCertURL = direct.LazyPtr(in.GetIncludeCaCertUrl())
	out.IncludeCrlAccessURL = direct.LazyPtr(in.GetIncludeCrlAccessUrl())
	return out
}
func CertificateAuthority_IssuingOptions_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_IssuingOptions) *pb.CertificateAuthority_IssuingOptions {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_IssuingOptions{}
	out.IncludeCaCertUrl = direct.ValueOf(in.IncludeCaCertURL)
	out.IncludeCrlAccessUrl = direct.ValueOf(in.IncludeCrlAccessURL)
	return out
}
func CertificateAuthority_KeyVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_KeyVersionSpec) *krm.CertificateAuthority_KeyVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_KeyVersionSpec{}
	out.CloudKMSKeyVersion = direct.LazyPtr(in.GetCloudKmsKeyVersion())
	out.Algorithm = direct.Enum_FromProto(mapCtx, in.GetAlgorithm())
	return out
}
func CertificateAuthority_KeyVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_KeyVersionSpec) *pb.CertificateAuthority_KeyVersionSpec {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_KeyVersionSpec{}
	if oneof := CertificateAuthority_KeyVersionSpec_CloudKmsKeyVersion_ToProto(mapCtx, in.CloudKMSKeyVersion); oneof != nil {
		out.KeyVersion = oneof
	}
	if oneof := CertificateAuthority_KeyVersionSpec_Algorithm_ToProto(mapCtx, in.Algorithm); oneof != nil {
		out.KeyVersion = oneof
	}
	return out
}
func CertificateConfig_FromProto(mapCtx *direct.MapContext, in *pb.CertificateConfig) *krm.CertificateConfig {
	if in == nil {
		return nil
	}
	out := &krm.CertificateConfig{}
	out.SubjectConfig = CertificateConfig_SubjectConfig_FromProto(mapCtx, in.GetSubjectConfig())
	out.ReusableConfig = ReusableConfigWrapper_FromProto(mapCtx, in.GetReusableConfig())
	out.PublicKey = PublicKey_FromProto(mapCtx, in.GetPublicKey())
	return out
}
func CertificateConfig_ToProto(mapCtx *direct.MapContext, in *krm.CertificateConfig) *pb.CertificateConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateConfig{}
	out.SubjectConfig = CertificateConfig_SubjectConfig_ToProto(mapCtx, in.SubjectConfig)
	out.ReusableConfig = ReusableConfigWrapper_ToProto(mapCtx, in.ReusableConfig)
	out.PublicKey = PublicKey_ToProto(mapCtx, in.PublicKey)
	return out
}
func CertificateConfig_SubjectConfig_FromProto(mapCtx *direct.MapContext, in *pb.CertificateConfig_SubjectConfig) *krm.CertificateConfig_SubjectConfig {
	if in == nil {
		return nil
	}
	out := &krm.CertificateConfig_SubjectConfig{}
	out.Subject = Subject_FromProto(mapCtx, in.GetSubject())
	out.CommonName = direct.LazyPtr(in.GetCommonName())
	out.SubjectAltName = SubjectAltNames_FromProto(mapCtx, in.GetSubjectAltName())
	return out
}
func CertificateConfig_SubjectConfig_ToProto(mapCtx *direct.MapContext, in *krm.CertificateConfig_SubjectConfig) *pb.CertificateConfig_SubjectConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateConfig_SubjectConfig{}
	out.Subject = Subject_ToProto(mapCtx, in.Subject)
	out.CommonName = direct.ValueOf(in.CommonName)
	out.SubjectAltName = SubjectAltNames_ToProto(mapCtx, in.SubjectAltName)
	return out
}
func CertificateDescription_FromProto(mapCtx *direct.MapContext, in *pb.CertificateDescription) *krm.CertificateDescription {
	if in == nil {
		return nil
	}
	out := &krm.CertificateDescription{}
	out.SubjectDescription = CertificateDescription_SubjectDescription_FromProto(mapCtx, in.GetSubjectDescription())
	out.ConfigValues = ReusableConfigValues_FromProto(mapCtx, in.GetConfigValues())
	out.PublicKey = PublicKey_FromProto(mapCtx, in.GetPublicKey())
	out.SubjectKeyID = CertificateDescription_KeyId_FromProto(mapCtx, in.GetSubjectKeyId())
	out.AuthorityKeyID = CertificateDescription_KeyId_FromProto(mapCtx, in.GetAuthorityKeyId())
	out.CrlDistributionPoints = in.CrlDistributionPoints
	out.AiaIssuingCertificateUrls = in.AiaIssuingCertificateUrls
	out.CertFingerprint = CertificateDescription_CertificateFingerprint_FromProto(mapCtx, in.GetCertFingerprint())
	return out
}
func CertificateDescription_ToProto(mapCtx *direct.MapContext, in *krm.CertificateDescription) *pb.CertificateDescription {
	if in == nil {
		return nil
	}
	out := &pb.CertificateDescription{}
	out.SubjectDescription = CertificateDescription_SubjectDescription_ToProto(mapCtx, in.SubjectDescription)
	out.ConfigValues = ReusableConfigValues_ToProto(mapCtx, in.ConfigValues)
	out.PublicKey = PublicKey_ToProto(mapCtx, in.PublicKey)
	out.SubjectKeyId = CertificateDescription_KeyId_ToProto(mapCtx, in.SubjectKeyID)
	out.AuthorityKeyId = CertificateDescription_KeyId_ToProto(mapCtx, in.AuthorityKeyID)
	out.CrlDistributionPoints = in.CrlDistributionPoints
	out.AiaIssuingCertificateUrls = in.AiaIssuingCertificateUrls
	out.CertFingerprint = CertificateDescription_CertificateFingerprint_ToProto(mapCtx, in.CertFingerprint)
	return out
}
func CertificateDescription_CertificateFingerprint_FromProto(mapCtx *direct.MapContext, in *pb.CertificateDescription_CertificateFingerprint) *krm.CertificateDescription_CertificateFingerprint {
	if in == nil {
		return nil
	}
	out := &krm.CertificateDescription_CertificateFingerprint{}
	out.Sha256Hash = direct.LazyPtr(in.GetSha256Hash())
	return out
}
func CertificateDescription_CertificateFingerprint_ToProto(mapCtx *direct.MapContext, in *krm.CertificateDescription_CertificateFingerprint) *pb.CertificateDescription_CertificateFingerprint {
	if in == nil {
		return nil
	}
	out := &pb.CertificateDescription_CertificateFingerprint{}
	out.Sha256Hash = direct.ValueOf(in.Sha256Hash)
	return out
}
func CertificateDescription_KeyId_FromProto(mapCtx *direct.MapContext, in *pb.CertificateDescription_KeyId) *krm.CertificateDescription_KeyId {
	if in == nil {
		return nil
	}
	out := &krm.CertificateDescription_KeyId{}
	out.KeyID = direct.LazyPtr(in.GetKeyId())
	return out
}
func CertificateDescription_KeyId_ToProto(mapCtx *direct.MapContext, in *krm.CertificateDescription_KeyId) *pb.CertificateDescription_KeyId {
	if in == nil {
		return nil
	}
	out := &pb.CertificateDescription_KeyId{}
	out.KeyId = direct.ValueOf(in.KeyID)
	return out
}
func CertificateDescription_SubjectDescription_FromProto(mapCtx *direct.MapContext, in *pb.CertificateDescription_SubjectDescription) *krm.CertificateDescription_SubjectDescription {
	if in == nil {
		return nil
	}
	out := &krm.CertificateDescription_SubjectDescription{}
	out.Subject = Subject_FromProto(mapCtx, in.GetSubject())
	out.CommonName = direct.LazyPtr(in.GetCommonName())
	out.SubjectAltName = SubjectAltNames_FromProto(mapCtx, in.GetSubjectAltName())
	out.HexSerialNumber = direct.LazyPtr(in.GetHexSerialNumber())
	out.Lifetime = direct.StringDuration_FromProto(mapCtx, in.GetLifetime())
	out.NotBeforeTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNotBeforeTime())
	out.NotAfterTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNotAfterTime())
	return out
}
func CertificateDescription_SubjectDescription_ToProto(mapCtx *direct.MapContext, in *krm.CertificateDescription_SubjectDescription) *pb.CertificateDescription_SubjectDescription {
	if in == nil {
		return nil
	}
	out := &pb.CertificateDescription_SubjectDescription{}
	out.Subject = Subject_ToProto(mapCtx, in.Subject)
	out.CommonName = direct.ValueOf(in.CommonName)
	out.SubjectAltName = SubjectAltNames_ToProto(mapCtx, in.SubjectAltName)
	out.HexSerialNumber = direct.ValueOf(in.HexSerialNumber)
	out.Lifetime = direct.StringDuration_ToProto(mapCtx, in.Lifetime)
	out.NotBeforeTime = direct.StringTimestamp_ToProto(mapCtx, in.NotBeforeTime)
	out.NotAfterTime = direct.StringTimestamp_ToProto(mapCtx, in.NotAfterTime)
	return out
}
func KeyUsage_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage) *krm.KeyUsage {
	if in == nil {
		return nil
	}
	out := &krm.KeyUsage{}
	out.BaseKeyUsage = KeyUsage_KeyUsageOptions_FromProto(mapCtx, in.GetBaseKeyUsage())
	out.ExtendedKeyUsage = KeyUsage_ExtendedKeyUsageOptions_FromProto(mapCtx, in.GetExtendedKeyUsage())
	out.UnknownExtendedKeyUsages = direct.Slice_FromProto(mapCtx, in.UnknownExtendedKeyUsages, ObjectId_FromProto)
	return out
}
func KeyUsage_ToProto(mapCtx *direct.MapContext, in *krm.KeyUsage) *pb.KeyUsage {
	if in == nil {
		return nil
	}
	out := &pb.KeyUsage{}
	out.BaseKeyUsage = KeyUsage_KeyUsageOptions_ToProto(mapCtx, in.BaseKeyUsage)
	out.ExtendedKeyUsage = KeyUsage_ExtendedKeyUsageOptions_ToProto(mapCtx, in.ExtendedKeyUsage)
	out.UnknownExtendedKeyUsages = direct.Slice_ToProto(mapCtx, in.UnknownExtendedKeyUsages, ObjectId_ToProto)
	return out
}
func KeyUsage_ExtendedKeyUsageOptions_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage_ExtendedKeyUsageOptions) *krm.KeyUsage_ExtendedKeyUsageOptions {
	if in == nil {
		return nil
	}
	out := &krm.KeyUsage_ExtendedKeyUsageOptions{}
	out.ServerAuth = direct.LazyPtr(in.GetServerAuth())
	out.ClientAuth = direct.LazyPtr(in.GetClientAuth())
	out.CodeSigning = direct.LazyPtr(in.GetCodeSigning())
	out.EmailProtection = direct.LazyPtr(in.GetEmailProtection())
	out.TimeStamping = direct.LazyPtr(in.GetTimeStamping())
	out.OcspSigning = direct.LazyPtr(in.GetOcspSigning())
	return out
}
func KeyUsage_ExtendedKeyUsageOptions_ToProto(mapCtx *direct.MapContext, in *krm.KeyUsage_ExtendedKeyUsageOptions) *pb.KeyUsage_ExtendedKeyUsageOptions {
	if in == nil {
		return nil
	}
	out := &pb.KeyUsage_ExtendedKeyUsageOptions{}
	out.ServerAuth = direct.ValueOf(in.ServerAuth)
	out.ClientAuth = direct.ValueOf(in.ClientAuth)
	out.CodeSigning = direct.ValueOf(in.CodeSigning)
	out.EmailProtection = direct.ValueOf(in.EmailProtection)
	out.TimeStamping = direct.ValueOf(in.TimeStamping)
	out.OcspSigning = direct.ValueOf(in.OcspSigning)
	return out
}
func KeyUsage_KeyUsageOptions_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage_KeyUsageOptions) *krm.KeyUsage_KeyUsageOptions {
	if in == nil {
		return nil
	}
	out := &krm.KeyUsage_KeyUsageOptions{}
	out.DigitalSignature = direct.LazyPtr(in.GetDigitalSignature())
	out.ContentCommitment = direct.LazyPtr(in.GetContentCommitment())
	out.KeyEncipherment = direct.LazyPtr(in.GetKeyEncipherment())
	out.DataEncipherment = direct.LazyPtr(in.GetDataEncipherment())
	out.KeyAgreement = direct.LazyPtr(in.GetKeyAgreement())
	out.CertSign = direct.LazyPtr(in.GetCertSign())
	out.CrlSign = direct.LazyPtr(in.GetCrlSign())
	out.EncipherOnly = direct.LazyPtr(in.GetEncipherOnly())
	out.DecipherOnly = direct.LazyPtr(in.GetDecipherOnly())
	return out
}
func KeyUsage_KeyUsageOptions_ToProto(mapCtx *direct.MapContext, in *krm.KeyUsage_KeyUsageOptions) *pb.KeyUsage_KeyUsageOptions {
	if in == nil {
		return nil
	}
	out := &pb.KeyUsage_KeyUsageOptions{}
	out.DigitalSignature = direct.ValueOf(in.DigitalSignature)
	out.ContentCommitment = direct.ValueOf(in.ContentCommitment)
	out.KeyEncipherment = direct.ValueOf(in.KeyEncipherment)
	out.DataEncipherment = direct.ValueOf(in.DataEncipherment)
	out.KeyAgreement = direct.ValueOf(in.KeyAgreement)
	out.CertSign = direct.ValueOf(in.CertSign)
	out.CrlSign = direct.ValueOf(in.CrlSign)
	out.EncipherOnly = direct.ValueOf(in.EncipherOnly)
	out.DecipherOnly = direct.ValueOf(in.DecipherOnly)
	return out
}
func ObjectId_FromProto(mapCtx *direct.MapContext, in *pb.ObjectId) *krm.ObjectId {
	if in == nil {
		return nil
	}
	out := &krm.ObjectId{}
	out.ObjectIDPath = in.ObjectIdPath
	return out
}
func ObjectId_ToProto(mapCtx *direct.MapContext, in *krm.ObjectId) *pb.ObjectId {
	if in == nil {
		return nil
	}
	out := &pb.ObjectId{}
	out.ObjectIdPath = in.ObjectIDPath
	return out
}
func PublicKey_FromProto(mapCtx *direct.MapContext, in *pb.PublicKey) *krm.PublicKey {
	if in == nil {
		return nil
	}
	out := &krm.PublicKey{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Key = in.GetKey()
	return out
}
func PublicKey_ToProto(mapCtx *direct.MapContext, in *krm.PublicKey) *pb.PublicKey {
	if in == nil {
		return nil
	}
	out := &pb.PublicKey{}
	out.Type = direct.Enum_ToProto[pb.PublicKey_KeyType](mapCtx, in.Type)
	out.Key = in.Key
	return out
}
func ReusableConfigValues_FromProto(mapCtx *direct.MapContext, in *pb.ReusableConfigValues) *krm.ReusableConfigValues {
	if in == nil {
		return nil
	}
	out := &krm.ReusableConfigValues{}
	out.KeyUsage = KeyUsage_FromProto(mapCtx, in.GetKeyUsage())
	out.CaOptions = ReusableConfigValues_CaOptions_FromProto(mapCtx, in.GetCaOptions())
	out.PolicyIds = direct.Slice_FromProto(mapCtx, in.PolicyIds, ObjectId_FromProto)
	out.AiaOcspServers = in.AiaOcspServers
	out.AdditionalExtensions = direct.Slice_FromProto(mapCtx, in.AdditionalExtensions, X509Extension_FromProto)
	return out
}
func ReusableConfigValues_ToProto(mapCtx *direct.MapContext, in *krm.ReusableConfigValues) *pb.ReusableConfigValues {
	if in == nil {
		return nil
	}
	out := &pb.ReusableConfigValues{}
	out.KeyUsage = KeyUsage_ToProto(mapCtx, in.KeyUsage)
	out.CaOptions = ReusableConfigValues_CaOptions_ToProto(mapCtx, in.CaOptions)
	out.PolicyIds = direct.Slice_ToProto(mapCtx, in.PolicyIds, ObjectId_ToProto)
	out.AiaOcspServers = in.AiaOcspServers
	out.AdditionalExtensions = direct.Slice_ToProto(mapCtx, in.AdditionalExtensions, X509Extension_ToProto)
	return out
}
func ReusableConfigValues_CaOptions_FromProto(mapCtx *direct.MapContext, in *pb.ReusableConfigValues_CaOptions) *krm.ReusableConfigValues_CaOptions {
	if in == nil {
		return nil
	}
	out := &krm.ReusableConfigValues_CaOptions{}
	out.IsCa = direct.BoolValue_FromProto(mapCtx, in.GetIsCa())
	out.MaxIssuerPathLength = Int32Value_FromProto(mapCtx, in.GetMaxIssuerPathLength())
	return out
}
func ReusableConfigValues_CaOptions_ToProto(mapCtx *direct.MapContext, in *krm.ReusableConfigValues_CaOptions) *pb.ReusableConfigValues_CaOptions {
	if in == nil {
		return nil
	}
	out := &pb.ReusableConfigValues_CaOptions{}
	out.IsCa = direct.BoolValue_ToProto(mapCtx, in.IsCa)
	out.MaxIssuerPathLength = Int32Value_ToProto(mapCtx, in.MaxIssuerPathLength)
	return out
}
func ReusableConfigWrapper_FromProto(mapCtx *direct.MapContext, in *pb.ReusableConfigWrapper) *krm.ReusableConfigWrapper {
	if in == nil {
		return nil
	}
	out := &krm.ReusableConfigWrapper{}
	out.ReusableConfig = direct.LazyPtr(in.GetReusableConfig())
	out.ReusableConfigValues = ReusableConfigValues_FromProto(mapCtx, in.GetReusableConfigValues())
	return out
}
func ReusableConfigWrapper_ToProto(mapCtx *direct.MapContext, in *krm.ReusableConfigWrapper) *pb.ReusableConfigWrapper {
	if in == nil {
		return nil
	}
	out := &pb.ReusableConfigWrapper{}
	if oneof := ReusableConfigWrapper_ReusableConfig_ToProto(mapCtx, in.ReusableConfig); oneof != nil {
		out.ConfigValues = oneof
	}
	if oneof := ReusableConfigValues_ToProto(mapCtx, in.ReusableConfigValues); oneof != nil {
		out.ConfigValues = &pb.ReusableConfigWrapper_ReusableConfigValues{ReusableConfigValues: oneof}
	}
	return out
}
func Subject_FromProto(mapCtx *direct.MapContext, in *pb.Subject) *krm.Subject {
	if in == nil {
		return nil
	}
	out := &krm.Subject{}
	out.CountryCode = direct.LazyPtr(in.GetCountryCode())
	out.Organization = direct.LazyPtr(in.GetOrganization())
	out.OrganizationalUnit = direct.LazyPtr(in.GetOrganizationalUnit())
	out.Locality = direct.LazyPtr(in.GetLocality())
	out.Province = direct.LazyPtr(in.GetProvince())
	out.StreetAddress = direct.LazyPtr(in.GetStreetAddress())
	out.PostalCode = direct.LazyPtr(in.GetPostalCode())
	return out
}
func Subject_ToProto(mapCtx *direct.MapContext, in *krm.Subject) *pb.Subject {
	if in == nil {
		return nil
	}
	out := &pb.Subject{}
	out.CountryCode = direct.ValueOf(in.CountryCode)
	out.Organization = direct.ValueOf(in.Organization)
	out.OrganizationalUnit = direct.ValueOf(in.OrganizationalUnit)
	out.Locality = direct.ValueOf(in.Locality)
	out.Province = direct.ValueOf(in.Province)
	out.StreetAddress = direct.ValueOf(in.StreetAddress)
	out.PostalCode = direct.ValueOf(in.PostalCode)
	return out
}
func SubjectAltNames_FromProto(mapCtx *direct.MapContext, in *pb.SubjectAltNames) *krm.SubjectAltNames {
	if in == nil {
		return nil
	}
	out := &krm.SubjectAltNames{}
	out.DnsNames = in.DnsNames
	out.Uris = in.Uris
	out.EmailAddresses = in.EmailAddresses
	out.IPAddresses = in.IpAddresses
	out.CustomSans = direct.Slice_FromProto(mapCtx, in.CustomSans, X509Extension_FromProto)
	return out
}
func SubjectAltNames_ToProto(mapCtx *direct.MapContext, in *krm.SubjectAltNames) *pb.SubjectAltNames {
	if in == nil {
		return nil
	}
	out := &pb.SubjectAltNames{}
	out.DnsNames = in.DnsNames
	out.Uris = in.Uris
	out.EmailAddresses = in.EmailAddresses
	out.IpAddresses = in.IPAddresses
	out.CustomSans = direct.Slice_ToProto(mapCtx, in.CustomSans, X509Extension_ToProto)
	return out
}
func SubordinateConfig_FromProto(mapCtx *direct.MapContext, in *pb.SubordinateConfig) *krm.SubordinateConfig {
	if in == nil {
		return nil
	}
	out := &krm.SubordinateConfig{}
	out.CertificateAuthority = direct.LazyPtr(in.GetCertificateAuthority())
	out.PemIssuerChain = SubordinateConfig_SubordinateConfigChain_FromProto(mapCtx, in.GetPemIssuerChain())
	return out
}
func SubordinateConfig_ToProto(mapCtx *direct.MapContext, in *krm.SubordinateConfig) *pb.SubordinateConfig {
	if in == nil {
		return nil
	}
	out := &pb.SubordinateConfig{}
	if oneof := SubordinateConfig_CertificateAuthority_ToProto(mapCtx, in.CertificateAuthority); oneof != nil {
		out.SubordinateConfig = oneof
	}
	if oneof := SubordinateConfig_SubordinateConfigChain_ToProto(mapCtx, in.PemIssuerChain); oneof != nil {
		out.SubordinateConfig = &pb.SubordinateConfig_PemIssuerChain{PemIssuerChain: oneof}
	}
	return out
}
func SubordinateConfig_SubordinateConfigChain_FromProto(mapCtx *direct.MapContext, in *pb.SubordinateConfig_SubordinateConfigChain) *krm.SubordinateConfig_SubordinateConfigChain {
	if in == nil {
		return nil
	}
	out := &krm.SubordinateConfig_SubordinateConfigChain{}
	out.PemCertificates = in.PemCertificates
	return out
}
func SubordinateConfig_SubordinateConfigChain_ToProto(mapCtx *direct.MapContext, in *krm.SubordinateConfig_SubordinateConfigChain) *pb.SubordinateConfig_SubordinateConfigChain {
	if in == nil {
		return nil
	}
	out := &pb.SubordinateConfig_SubordinateConfigChain{}
	out.PemCertificates = in.PemCertificates
	return out
}
func X509Extension_FromProto(mapCtx *direct.MapContext, in *pb.X509Extension) *krm.X509Extension {
	if in == nil {
		return nil
	}
	out := &krm.X509Extension{}
	out.ObjectID = ObjectId_FromProto(mapCtx, in.GetObjectId())
	out.Critical = direct.LazyPtr(in.GetCritical())
	out.Value = in.GetValue()
	return out
}
func X509Extension_ToProto(mapCtx *direct.MapContext, in *krm.X509Extension) *pb.X509Extension {
	if in == nil {
		return nil
	}
	out := &pb.X509Extension{}
	out.ObjectId = ObjectId_ToProto(mapCtx, in.ObjectID)
	out.Critical = direct.ValueOf(in.Critical)
	out.Value = in.Value
	return out
}
