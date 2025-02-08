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
	pb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/security/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CertificateExtensionConstraints_FromProto(mapCtx *direct.MapContext, in *pb.CertificateExtensionConstraints) *krm.CertificateExtensionConstraints {
	if in == nil {
		return nil
	}
	out := &krm.CertificateExtensionConstraints{}
	out.KnownExtensions = direct.EnumSlice_FromProto(mapCtx, in.KnownExtensions)
	out.AdditionalExtensions = direct.Slice_FromProto(mapCtx, in.AdditionalExtensions, ObjectId_FromProto)
	return out
}
func CertificateExtensionConstraints_ToProto(mapCtx *direct.MapContext, in *krm.CertificateExtensionConstraints) *pb.CertificateExtensionConstraints {
	if in == nil {
		return nil
	}
	out := &pb.CertificateExtensionConstraints{}
	out.KnownExtensions = direct.EnumSlice_ToProto[pb.CertificateExtensionConstraints_KnownCertificateExtension](mapCtx, in.KnownExtensions)
	out.AdditionalExtensions = direct.Slice_ToProto(mapCtx, in.AdditionalExtensions, ObjectId_ToProto)
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
func CertificateTemplate_FromProto(mapCtx *direct.MapContext, in *pb.CertificateTemplate) *krm.CertificateTemplate {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate{}
	// MISSING: Name
	out.MaximumLifetime = direct.StringDuration_FromProto(mapCtx, in.GetMaximumLifetime())
	out.PredefinedValues = X509Parameters_FromProto(mapCtx, in.GetPredefinedValues())
	out.IdentityConstraints = CertificateIdentityConstraints_FromProto(mapCtx, in.GetIdentityConstraints())
	out.PassthroughExtensions = CertificateExtensionConstraints_FromProto(mapCtx, in.GetPassthroughExtensions())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	return out
}
func CertificateTemplate_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate) *pb.CertificateTemplate {
	if in == nil {
		return nil
	}
	out := &pb.CertificateTemplate{}
	// MISSING: Name
	out.MaximumLifetime = direct.StringDuration_ToProto(mapCtx, in.MaximumLifetime)
	out.PredefinedValues = X509Parameters_ToProto(mapCtx, in.PredefinedValues)
	out.IdentityConstraints = CertificateIdentityConstraints_ToProto(mapCtx, in.IdentityConstraints)
	out.PassthroughExtensions = CertificateExtensionConstraints_ToProto(mapCtx, in.PassthroughExtensions)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	return out
}
func CertificateTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateTemplate) *krm.CertificateTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplateObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: MaximumLifetime
	// MISSING: PredefinedValues
	// MISSING: IdentityConstraints
	// MISSING: PassthroughExtensions
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	return out
}
func CertificateTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplateObservedState) *pb.CertificateTemplate {
	if in == nil {
		return nil
	}
	out := &pb.CertificateTemplate{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: MaximumLifetime
	// MISSING: PredefinedValues
	// MISSING: IdentityConstraints
	// MISSING: PassthroughExtensions
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
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
func SecurityCertificateTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateTemplate) *krm.SecurityCertificateTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecurityCertificateTemplateObservedState{}
	// MISSING: Name
	// MISSING: MaximumLifetime
	// MISSING: PredefinedValues
	// MISSING: IdentityConstraints
	// MISSING: PassthroughExtensions
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	return out
}
func SecurityCertificateTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecurityCertificateTemplateObservedState) *pb.CertificateTemplate {
	if in == nil {
		return nil
	}
	out := &pb.CertificateTemplate{}
	// MISSING: Name
	// MISSING: MaximumLifetime
	// MISSING: PredefinedValues
	// MISSING: IdentityConstraints
	// MISSING: PassthroughExtensions
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	return out
}
func SecurityCertificateTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.CertificateTemplate) *krm.SecurityCertificateTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecurityCertificateTemplateSpec{}
	// MISSING: Name
	// MISSING: MaximumLifetime
	// MISSING: PredefinedValues
	// MISSING: IdentityConstraints
	// MISSING: PassthroughExtensions
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	return out
}
func SecurityCertificateTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecurityCertificateTemplateSpec) *pb.CertificateTemplate {
	if in == nil {
		return nil
	}
	out := &pb.CertificateTemplate{}
	// MISSING: Name
	// MISSING: MaximumLifetime
	// MISSING: PredefinedValues
	// MISSING: IdentityConstraints
	// MISSING: PassthroughExtensions
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
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
func X509Parameters_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters) *krm.X509Parameters {
	if in == nil {
		return nil
	}
	out := &krm.X509Parameters{}
	out.KeyUsage = KeyUsage_FromProto(mapCtx, in.GetKeyUsage())
	out.CaOptions = X509Parameters_CaOptions_FromProto(mapCtx, in.GetCaOptions())
	out.PolicyIds = direct.Slice_FromProto(mapCtx, in.PolicyIds, ObjectId_FromProto)
	out.AiaOcspServers = in.AiaOcspServers
	out.NameConstraints = X509Parameters_NameConstraints_FromProto(mapCtx, in.GetNameConstraints())
	out.AdditionalExtensions = direct.Slice_FromProto(mapCtx, in.AdditionalExtensions, X509Extension_FromProto)
	return out
}
func X509Parameters_ToProto(mapCtx *direct.MapContext, in *krm.X509Parameters) *pb.X509Parameters {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters{}
	out.KeyUsage = KeyUsage_ToProto(mapCtx, in.KeyUsage)
	out.CaOptions = X509Parameters_CaOptions_ToProto(mapCtx, in.CaOptions)
	out.PolicyIds = direct.Slice_ToProto(mapCtx, in.PolicyIds, ObjectId_ToProto)
	out.AiaOcspServers = in.AiaOcspServers
	out.NameConstraints = X509Parameters_NameConstraints_ToProto(mapCtx, in.NameConstraints)
	out.AdditionalExtensions = direct.Slice_ToProto(mapCtx, in.AdditionalExtensions, X509Extension_ToProto)
	return out
}
func X509Parameters_CaOptions_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters_CaOptions) *krm.X509Parameters_CaOptions {
	if in == nil {
		return nil
	}
	out := &krm.X509Parameters_CaOptions{}
	out.IsCa = in.IsCa
	out.MaxIssuerPathLength = in.MaxIssuerPathLength
	return out
}
func X509Parameters_CaOptions_ToProto(mapCtx *direct.MapContext, in *krm.X509Parameters_CaOptions) *pb.X509Parameters_CaOptions {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters_CaOptions{}
	out.IsCa = in.IsCa
	out.MaxIssuerPathLength = in.MaxIssuerPathLength
	return out
}
func X509Parameters_NameConstraints_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters_NameConstraints) *krm.X509Parameters_NameConstraints {
	if in == nil {
		return nil
	}
	out := &krm.X509Parameters_NameConstraints{}
	out.Critical = direct.LazyPtr(in.GetCritical())
	out.PermittedDnsNames = in.PermittedDnsNames
	out.ExcludedDnsNames = in.ExcludedDnsNames
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
	out.PermittedDnsNames = in.PermittedDnsNames
	out.ExcludedDnsNames = in.ExcludedDnsNames
	out.PermittedIpRanges = in.PermittedIPRanges
	out.ExcludedIpRanges = in.ExcludedIPRanges
	out.PermittedEmailAddresses = in.PermittedEmailAddresses
	out.ExcludedEmailAddresses = in.ExcludedEmailAddresses
	out.PermittedUris = in.PermittedUris
	out.ExcludedUris = in.ExcludedUris
	return out
}
