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
	pb "cloud.google.com/go/security/privateca/apiv1beta1/privatecapb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/security/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Certificate_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) *krm.Certificate {
	if in == nil {
		return nil
	}
	out := &krm.Certificate{}
	// MISSING: Name
	out.PemCsr = direct.LazyPtr(in.GetPemCsr())
	out.Config = CertificateConfig_FromProto(mapCtx, in.GetConfig())
	out.Lifetime = direct.StringDuration_FromProto(mapCtx, in.GetLifetime())
	// MISSING: RevocationDetails
	// MISSING: PemCertificate
	// MISSING: CertificateDescription
	// MISSING: PemCertificateChain
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	return out
}
func Certificate_ToProto(mapCtx *direct.MapContext, in *krm.Certificate) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	// MISSING: Name
	if oneof := Certificate_PemCsr_ToProto(mapCtx, in.PemCsr); oneof != nil {
		out.CertificateConfig = oneof
	}
	if oneof := CertificateConfig_ToProto(mapCtx, in.Config); oneof != nil {
		out.CertificateConfig = &pb.Certificate_Config{Config: oneof}
	}
	out.Lifetime = direct.StringDuration_ToProto(mapCtx, in.Lifetime)
	// MISSING: RevocationDetails
	// MISSING: PemCertificate
	// MISSING: CertificateDescription
	// MISSING: PemCertificateChain
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
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
func CertificateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) *krm.CertificateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: PemCsr
	// MISSING: Config
	// MISSING: Lifetime
	out.RevocationDetails = Certificate_RevocationDetails_FromProto(mapCtx, in.GetRevocationDetails())
	out.PemCertificate = direct.LazyPtr(in.GetPemCertificate())
	out.CertificateDescription = CertificateDescription_FromProto(mapCtx, in.GetCertificateDescription())
	out.PemCertificateChain = in.PemCertificateChain
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	return out
}
func CertificateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateObservedState) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: PemCsr
	// MISSING: Config
	// MISSING: Lifetime
	out.RevocationDetails = Certificate_RevocationDetails_ToProto(mapCtx, in.RevocationDetails)
	out.PemCertificate = direct.ValueOf(in.PemCertificate)
	out.CertificateDescription = CertificateDescription_ToProto(mapCtx, in.CertificateDescription)
	out.PemCertificateChain = in.PemCertificateChain
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	return out
}
func Certificate_RevocationDetails_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_RevocationDetails) *krm.Certificate_RevocationDetails {
	if in == nil {
		return nil
	}
	out := &krm.Certificate_RevocationDetails{}
	out.RevocationState = direct.Enum_FromProto(mapCtx, in.GetRevocationState())
	out.RevocationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevocationTime())
	return out
}
func Certificate_RevocationDetails_ToProto(mapCtx *direct.MapContext, in *krm.Certificate_RevocationDetails) *pb.Certificate_RevocationDetails {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_RevocationDetails{}
	out.RevocationState = direct.Enum_ToProto[pb.RevocationReason](mapCtx, in.RevocationState)
	out.RevocationTime = direct.StringTimestamp_ToProto(mapCtx, in.RevocationTime)
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
