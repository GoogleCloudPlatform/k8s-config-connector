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

// +generated:mapper
// krm.group: privateca.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.security.privateca.v1

package privateca

import (
	pb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CAPool_IssuancePolicy_FromProto(mapCtx *direct.MapContext, in *pb.CaPool_IssuancePolicy) *krm.CAPool_IssuancePolicy {
	if in == nil {
		return nil
	}
	out := &krm.CAPool_IssuancePolicy{}
	out.AllowedKeyTypes = direct.Slice_FromProto(mapCtx, in.AllowedKeyTypes, CAPool_IssuancePolicy_AllowedKeyType_FromProto)
	// MISSING: BackdateDuration
	out.MaximumLifetime = direct.StringDuration_FromProto(mapCtx, in.GetMaximumLifetime())
	out.AllowedIssuanceModes = CAPool_IssuancePolicy_IssuanceModes_FromProto(mapCtx, in.GetAllowedIssuanceModes())
	out.BaselineValues = X509Parameters_FromProto(mapCtx, in.GetBaselineValues())
	out.IdentityConstraints = CertificateIdentityConstraints_FromProto(mapCtx, in.GetIdentityConstraints())
	out.PassthroughExtensions = CertificateExtensionConstraints_FromProto(mapCtx, in.GetPassthroughExtensions())
	return out
}
func CAPool_IssuancePolicy_ToProto(mapCtx *direct.MapContext, in *krm.CAPool_IssuancePolicy) *pb.CaPool_IssuancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.CaPool_IssuancePolicy{}
	out.AllowedKeyTypes = direct.Slice_ToProto(mapCtx, in.AllowedKeyTypes, CAPool_IssuancePolicy_AllowedKeyType_ToProto)
	// MISSING: BackdateDuration
	out.MaximumLifetime = direct.StringDuration_ToProto(mapCtx, in.MaximumLifetime)
	out.AllowedIssuanceModes = CAPool_IssuancePolicy_IssuanceModes_ToProto(mapCtx, in.AllowedIssuanceModes)
	out.BaselineValues = X509Parameters_ToProto(mapCtx, in.BaselineValues)
	out.IdentityConstraints = CertificateIdentityConstraints_ToProto(mapCtx, in.IdentityConstraints)
	out.PassthroughExtensions = CertificateExtensionConstraints_ToProto(mapCtx, in.PassthroughExtensions)
	return out
}
func CAPool_IssuancePolicy_AllowedKeyType_FromProto(mapCtx *direct.MapContext, in *pb.CaPool_IssuancePolicy_AllowedKeyType) *krm.CAPool_IssuancePolicy_AllowedKeyType {
	if in == nil {
		return nil
	}
	out := &krm.CAPool_IssuancePolicy_AllowedKeyType{}
	out.Rsa = CAPool_IssuancePolicy_AllowedKeyType_RsaKeyType_FromProto(mapCtx, in.GetRsa())
	out.EllipticCurve = CAPool_IssuancePolicy_AllowedKeyType_EcKeyType_FromProto(mapCtx, in.GetEllipticCurve())
	return out
}
func CAPool_IssuancePolicy_AllowedKeyType_ToProto(mapCtx *direct.MapContext, in *krm.CAPool_IssuancePolicy_AllowedKeyType) *pb.CaPool_IssuancePolicy_AllowedKeyType {
	if in == nil {
		return nil
	}
	out := &pb.CaPool_IssuancePolicy_AllowedKeyType{}
	if oneof := CAPool_IssuancePolicy_AllowedKeyType_RsaKeyType_ToProto(mapCtx, in.Rsa); oneof != nil {
		out.KeyType = &pb.CaPool_IssuancePolicy_AllowedKeyType_Rsa{Rsa: oneof}
	}
	if oneof := CAPool_IssuancePolicy_AllowedKeyType_EcKeyType_ToProto(mapCtx, in.EllipticCurve); oneof != nil {
		out.KeyType = &pb.CaPool_IssuancePolicy_AllowedKeyType_EllipticCurve{EllipticCurve: oneof}
	}
	return out
}
func CAPool_IssuancePolicy_AllowedKeyType_EcKeyType_FromProto(mapCtx *direct.MapContext, in *pb.CaPool_IssuancePolicy_AllowedKeyType_EcKeyType) *krm.CAPool_IssuancePolicy_AllowedKeyType_EcKeyType {
	if in == nil {
		return nil
	}
	out := &krm.CAPool_IssuancePolicy_AllowedKeyType_EcKeyType{}
	out.SignatureAlgorithm = direct.Enum_FromProto(mapCtx, in.GetSignatureAlgorithm())
	return out
}
func CAPool_IssuancePolicy_AllowedKeyType_EcKeyType_ToProto(mapCtx *direct.MapContext, in *krm.CAPool_IssuancePolicy_AllowedKeyType_EcKeyType) *pb.CaPool_IssuancePolicy_AllowedKeyType_EcKeyType {
	if in == nil {
		return nil
	}
	out := &pb.CaPool_IssuancePolicy_AllowedKeyType_EcKeyType{}
	out.SignatureAlgorithm = direct.Enum_ToProto[pb.CaPool_IssuancePolicy_AllowedKeyType_EcKeyType_EcSignatureAlgorithm](mapCtx, in.SignatureAlgorithm)
	return out
}
func CAPool_IssuancePolicy_AllowedKeyType_RsaKeyType_FromProto(mapCtx *direct.MapContext, in *pb.CaPool_IssuancePolicy_AllowedKeyType_RsaKeyType) *krm.CAPool_IssuancePolicy_AllowedKeyType_RsaKeyType {
	if in == nil {
		return nil
	}
	out := &krm.CAPool_IssuancePolicy_AllowedKeyType_RsaKeyType{}
	out.MinModulusSize = direct.LazyPtr(in.GetMinModulusSize())
	out.MaxModulusSize = direct.LazyPtr(in.GetMaxModulusSize())
	return out
}
func CAPool_IssuancePolicy_AllowedKeyType_RsaKeyType_ToProto(mapCtx *direct.MapContext, in *krm.CAPool_IssuancePolicy_AllowedKeyType_RsaKeyType) *pb.CaPool_IssuancePolicy_AllowedKeyType_RsaKeyType {
	if in == nil {
		return nil
	}
	out := &pb.CaPool_IssuancePolicy_AllowedKeyType_RsaKeyType{}
	out.MinModulusSize = direct.ValueOf(in.MinModulusSize)
	out.MaxModulusSize = direct.ValueOf(in.MaxModulusSize)
	return out
}
func CAPool_IssuancePolicy_IssuanceModes_FromProto(mapCtx *direct.MapContext, in *pb.CaPool_IssuancePolicy_IssuanceModes) *krm.CAPool_IssuancePolicy_IssuanceModes {
	if in == nil {
		return nil
	}
	out := &krm.CAPool_IssuancePolicy_IssuanceModes{}
	out.AllowCsrBasedIssuance = direct.LazyPtr(in.GetAllowCsrBasedIssuance())
	out.AllowConfigBasedIssuance = direct.LazyPtr(in.GetAllowConfigBasedIssuance())
	return out
}
func CAPool_IssuancePolicy_IssuanceModes_ToProto(mapCtx *direct.MapContext, in *krm.CAPool_IssuancePolicy_IssuanceModes) *pb.CaPool_IssuancePolicy_IssuanceModes {
	if in == nil {
		return nil
	}
	out := &pb.CaPool_IssuancePolicy_IssuanceModes{}
	out.AllowCsrBasedIssuance = direct.ValueOf(in.AllowCsrBasedIssuance)
	out.AllowConfigBasedIssuance = direct.ValueOf(in.AllowConfigBasedIssuance)
	return out
}
func CAPool_PublishingOptions_FromProto(mapCtx *direct.MapContext, in *pb.CaPool_PublishingOptions) *krm.CAPool_PublishingOptions {
	if in == nil {
		return nil
	}
	out := &krm.CAPool_PublishingOptions{}
	out.PublishCACert = direct.LazyPtr(in.GetPublishCaCert())
	out.PublishCrl = direct.LazyPtr(in.GetPublishCrl())
	// MISSING: EncodingFormat
	return out
}
func CAPool_PublishingOptions_ToProto(mapCtx *direct.MapContext, in *krm.CAPool_PublishingOptions) *pb.CaPool_PublishingOptions {
	if in == nil {
		return nil
	}
	out := &pb.CaPool_PublishingOptions{}
	out.PublishCaCert = direct.ValueOf(in.PublishCACert)
	out.PublishCrl = direct.ValueOf(in.PublishCrl)
	// MISSING: EncodingFormat
	return out
}
func CertificateExtensionConstraints_FromProto(mapCtx *direct.MapContext, in *pb.CertificateExtensionConstraints) *krm.CertificateExtensionConstraints {
	if in == nil {
		return nil
	}
	out := &krm.CertificateExtensionConstraints{}
	out.KnownExtensions = direct.EnumSlice_FromProto(mapCtx, in.KnownExtensions)
	out.AdditionalExtensions = direct.Slice_FromProto(mapCtx, in.AdditionalExtensions, ObjectID_FromProto)
	return out
}
func CertificateExtensionConstraints_ToProto(mapCtx *direct.MapContext, in *krm.CertificateExtensionConstraints) *pb.CertificateExtensionConstraints {
	if in == nil {
		return nil
	}
	out := &pb.CertificateExtensionConstraints{}
	out.KnownExtensions = direct.EnumSlice_ToProto[pb.CertificateExtensionConstraints_KnownCertificateExtension](mapCtx, in.KnownExtensions)
	out.AdditionalExtensions = direct.Slice_ToProto(mapCtx, in.AdditionalExtensions, ObjectID_ToProto)
	return out
}
func CertificateIdentityConstraints_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIdentityConstraints) *krm.CertificateIdentityConstraints {
	if in == nil {
		return nil
	}
	out := &krm.CertificateIdentityConstraints{}
	out.CelExpression = Expr_FromProto(mapCtx, in.GetCelExpression())
	out.AllowSubjectPassthrough = in.AllowSubjectPassthrough
	out.AllowSubjectAltNamesPassthrough = in.AllowSubjectAltNamesPassthrough
	return out
}
func CertificateIdentityConstraints_ToProto(mapCtx *direct.MapContext, in *krm.CertificateIdentityConstraints) *pb.CertificateIdentityConstraints {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIdentityConstraints{}
	out.CelExpression = Expr_ToProto(mapCtx, in.CelExpression)
	out.AllowSubjectPassthrough = in.AllowSubjectPassthrough
	out.AllowSubjectAltNamesPassthrough = in.AllowSubjectAltNamesPassthrough
	return out
}
func KeyUsage_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage) *krm.KeyUsage {
	if in == nil {
		return nil
	}
	out := &krm.KeyUsage{}
	out.BaseKeyUsage = KeyUsage_KeyUsageOptions_FromProto(mapCtx, in.GetBaseKeyUsage())
	out.ExtendedKeyUsage = KeyUsage_ExtendedKeyUsageOptions_FromProto(mapCtx, in.GetExtendedKeyUsage())
	out.UnknownExtendedKeyUsages = direct.Slice_FromProto(mapCtx, in.UnknownExtendedKeyUsages, ObjectID_FromProto)
	return out
}
func KeyUsage_ToProto(mapCtx *direct.MapContext, in *krm.KeyUsage) *pb.KeyUsage {
	if in == nil {
		return nil
	}
	out := &pb.KeyUsage{}
	out.BaseKeyUsage = KeyUsage_KeyUsageOptions_ToProto(mapCtx, in.BaseKeyUsage)
	out.ExtendedKeyUsage = KeyUsage_ExtendedKeyUsageOptions_ToProto(mapCtx, in.ExtendedKeyUsage)
	out.UnknownExtendedKeyUsages = direct.Slice_ToProto(mapCtx, in.UnknownExtendedKeyUsages, ObjectID_ToProto)
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
func ObjectID_FromProto(mapCtx *direct.MapContext, in *pb.ObjectId) *krm.ObjectID {
	if in == nil {
		return nil
	}
	out := &krm.ObjectID{}
	out.ObjectIDPath = in.ObjectIdPath
	return out
}
func ObjectID_ToProto(mapCtx *direct.MapContext, in *krm.ObjectID) *pb.ObjectId {
	if in == nil {
		return nil
	}
	out := &pb.ObjectId{}
	out.ObjectIdPath = in.ObjectIDPath
	return out
}
func PrivateCACAPoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CaPool) *krm.PrivateCACAPoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCACAPoolObservedState{}
	// MISSING: Name
	// MISSING: Labels
	return out
}
func PrivateCACAPoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCACAPoolObservedState) *pb.CaPool {
	if in == nil {
		return nil
	}
	out := &pb.CaPool{}
	// MISSING: Name
	// MISSING: Labels
	return out
}
func PrivateCACAPoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.CaPool) *krm.PrivateCACAPoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCACAPoolSpec{}
	// MISSING: Name
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	out.IssuancePolicy = CAPool_IssuancePolicy_FromProto(mapCtx, in.GetIssuancePolicy())
	out.PublishingOptions = CAPool_PublishingOptions_FromProto(mapCtx, in.GetPublishingOptions())
	// MISSING: Labels
	return out
}
func PrivateCACAPoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCACAPoolSpec) *pb.CaPool {
	if in == nil {
		return nil
	}
	out := &pb.CaPool{}
	// MISSING: Name
	out.Tier = direct.Enum_ToProto[pb.CaPool_Tier](mapCtx, in.Tier)
	out.IssuancePolicy = CAPool_IssuancePolicy_ToProto(mapCtx, in.IssuancePolicy)
	out.PublishingOptions = CAPool_PublishingOptions_ToProto(mapCtx, in.PublishingOptions)
	// MISSING: Labels
	return out
}
func X509Parameters_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters) *krm.X509Parameters {
	if in == nil {
		return nil
	}
	out := &krm.X509Parameters{}
	out.KeyUsage = KeyUsage_FromProto(mapCtx, in.GetKeyUsage())
	out.CAOptions = X509Parameters_CAOptions_FromProto(mapCtx, in.GetCaOptions())
	out.PolicyIds = direct.Slice_FromProto(mapCtx, in.PolicyIds, ObjectID_FromProto)
	out.AiaOcspServers = in.AiaOcspServers
	// MISSING: NameConstraints
	out.AdditionalExtensions = direct.Slice_FromProto(mapCtx, in.AdditionalExtensions, X509Extension_FromProto)
	return out
}
func X509Parameters_ToProto(mapCtx *direct.MapContext, in *krm.X509Parameters) *pb.X509Parameters {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters{}
	out.KeyUsage = KeyUsage_ToProto(mapCtx, in.KeyUsage)
	out.CaOptions = X509Parameters_CAOptions_ToProto(mapCtx, in.CAOptions)
	out.PolicyIds = direct.Slice_ToProto(mapCtx, in.PolicyIds, ObjectID_ToProto)
	out.AiaOcspServers = in.AiaOcspServers
	// MISSING: NameConstraints
	out.AdditionalExtensions = direct.Slice_ToProto(mapCtx, in.AdditionalExtensions, X509Extension_ToProto)
	return out
}
func X509Parameters_CAOptions_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters_CaOptions) *krm.X509Parameters_CAOptions {
	if in == nil {
		return nil
	}
	out := &krm.X509Parameters_CAOptions{}
	out.IsCA = in.IsCa
	out.MaxIssuerPathLength = in.MaxIssuerPathLength
	return out
}
func X509Parameters_CAOptions_ToProto(mapCtx *direct.MapContext, in *krm.X509Parameters_CAOptions) *pb.X509Parameters_CaOptions {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters_CaOptions{}
	out.IsCa = in.IsCA
	out.MaxIssuerPathLength = in.MaxIssuerPathLength
	return out
}
func X509Parameters_NameConstraints_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters_NameConstraints) *krm.X509Parameters_NameConstraints {
	if in == nil {
		return nil
	}
	out := &krm.X509Parameters_NameConstraints{}
	out.Critical = direct.LazyPtr(in.GetCritical())
	out.PermittedDNSNames = in.PermittedDnsNames
	out.ExcludedDNSNames = in.ExcludedDnsNames
	out.PermittedIPRanges = in.PermittedIpRanges
	out.ExcludedIPRanges = in.ExcludedIpRanges
	out.PermittedEmailAddresses = in.PermittedEmailAddresses
	out.ExcludedEmailAddresses = in.ExcludedEmailAddresses
	out.PermittedUris = in.PermittedUris
	out.ExcludedUris = in.ExcludedUris
	return out
}
func X509Parameters_NameConstraints_ToProto(mapCtx *direct.MapContext, in *krm.X509Parameters_NameConstraints) *pb.X509Parameters_NameConstraints {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters_NameConstraints{}
	out.Critical = direct.ValueOf(in.Critical)
	out.PermittedDnsNames = in.PermittedDNSNames
	out.ExcludedDnsNames = in.ExcludedDNSNames
	out.PermittedIpRanges = in.PermittedIPRanges
	out.ExcludedIpRanges = in.ExcludedIPRanges
	out.PermittedEmailAddresses = in.PermittedEmailAddresses
	out.ExcludedEmailAddresses = in.ExcludedEmailAddresses
	out.PermittedUris = in.PermittedUris
	out.ExcludedUris = in.ExcludedUris
	return out
}
