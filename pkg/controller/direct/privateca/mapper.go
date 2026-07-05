// Copyright 2024 Google LLC
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

package privateca

import (
	pb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/kmsrefs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	exprpb "google.golang.org/genproto/googleapis/type/expr"
)

func Expr_FromProto(mapCtx *direct.MapContext, in *exprpb.Expr) *krm.Expr {
	if in == nil {
		return nil
	}
	out := &krm.Expr{}
	out.Expression = direct.LazyPtr(in.GetExpression())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}

func Expr_ToProto(mapCtx *direct.MapContext, in *krm.Expr) *exprpb.Expr {
	if in == nil {
		return nil
	}
	out := &exprpb.Expr{}
	out.Expression = direct.ValueOf(in.Expression)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Location = direct.ValueOf(in.Location)
	return out
}

func X509Extension_FromProto(mapCtx *direct.MapContext, in *pb.X509Extension) *krm.X509Extension {
	if in == nil {
		return nil
	}
	out := &krm.X509Extension{}
	out.ObjectID = ObjectID_FromProto(mapCtx, in.GetObjectId())
	out.Critical = direct.LazyPtr(in.GetCritical())
	out.Value = in.GetValue()
	return out
}

func X509Extension_ToProto(mapCtx *direct.MapContext, in *krm.X509Extension) *pb.X509Extension {
	if in == nil {
		return nil
	}
	out := &pb.X509Extension{}
	out.ObjectId = ObjectID_ToProto(mapCtx, in.ObjectID)
	out.Critical = direct.ValueOf(in.Critical)
	out.Value = in.Value
	return out
}

func PrivateCACertificateTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.CertificateTemplate) *krm.PrivateCACertificateTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCACertificateTemplateSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.IdentityConstraints = CertificateTemplate_IdentityConstraints_FromProto(mapCtx, in.GetIdentityConstraints())
	out.PassthroughExtensions = CertificateTemplate_PassthroughExtensions_FromProto(mapCtx, in.GetPassthroughExtensions())
	out.PredefinedValues = CertificateTemplate_X509Parameters_FromProto(mapCtx, in.GetPredefinedValues())
	return out
}

func PrivateCACertificateTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCACertificateTemplateSpec) *pb.CertificateTemplate {
	if in == nil {
		return nil
	}
	out := &pb.CertificateTemplate{}
	out.Description = direct.ValueOf(in.Description)
	out.IdentityConstraints = CertificateTemplate_IdentityConstraints_ToProto(mapCtx, in.IdentityConstraints)
	out.PassthroughExtensions = CertificateTemplate_PassthroughExtensions_ToProto(mapCtx, in.PassthroughExtensions)
	out.PredefinedValues = CertificateTemplate_X509Parameters_ToProto(mapCtx, in.PredefinedValues)
	return out
}

func CertificateTemplate_IdentityConstraints_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIdentityConstraints) *krm.CertificateTemplate_IdentityConstraints {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate_IdentityConstraints{}
	out.AllowSubjectAltNamesPassthrough = in.AllowSubjectAltNamesPassthrough
	out.AllowSubjectPassthrough = in.AllowSubjectPassthrough
	out.CelExpression = CertificateTemplate_Expr_FromProto(mapCtx, in.GetCelExpression())
	return out
}

func CertificateTemplate_IdentityConstraints_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate_IdentityConstraints) *pb.CertificateIdentityConstraints {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIdentityConstraints{}
	out.AllowSubjectAltNamesPassthrough = in.AllowSubjectAltNamesPassthrough
	out.AllowSubjectPassthrough = in.AllowSubjectPassthrough
	out.CelExpression = CertificateTemplate_Expr_ToProto(mapCtx, in.CelExpression)
	return out
}

func CertificateTemplate_Expr_FromProto(mapCtx *direct.MapContext, in *exprpb.Expr) *krm.CertificateTemplate_Expr {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate_Expr{}
	out.Expression = direct.LazyPtr(in.GetExpression())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}

func CertificateTemplate_Expr_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate_Expr) *exprpb.Expr {
	if in == nil {
		return nil
	}
	out := &exprpb.Expr{}
	out.Expression = direct.ValueOf(in.Expression)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Location = direct.ValueOf(in.Location)
	return out
}

func CertificateTemplate_PassthroughExtensions_FromProto(mapCtx *direct.MapContext, in *pb.CertificateExtensionConstraints) *krm.CertificateTemplate_PassthroughExtensions {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate_PassthroughExtensions{}
	out.KnownExtensions = direct.EnumSlice_FromProto(mapCtx, in.KnownExtensions)
	out.AdditionalExtensions = direct.Slice_FromProto[pb.ObjectId, krm.CertificateTemplate_ObjectID](mapCtx, in.AdditionalExtensions, CertificateTemplate_ObjectID_FromProto)
	return out
}

func CertificateTemplate_PassthroughExtensions_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate_PassthroughExtensions) *pb.CertificateExtensionConstraints {
	if in == nil {
		return nil
	}
	out := &pb.CertificateExtensionConstraints{}
	out.KnownExtensions = direct.EnumSlice_ToProto[pb.CertificateExtensionConstraints_KnownCertificateExtension](mapCtx, in.KnownExtensions)
	out.AdditionalExtensions = direct.Slice_ToProto[krm.CertificateTemplate_ObjectID, pb.ObjectId](mapCtx, in.AdditionalExtensions, CertificateTemplate_ObjectID_ToProto)
	return out
}

func CertificateTemplate_ObjectID_FromProto(mapCtx *direct.MapContext, in *pb.ObjectId) *krm.CertificateTemplate_ObjectID {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate_ObjectID{}
	if in.GetObjectIdPath() != nil {
		out.ObjectIDPath = make([]int64, len(in.GetObjectIdPath()))
		for i, v := range in.GetObjectIdPath() {
			out.ObjectIDPath[i] = int64(v)
		}
	}
	return out
}

func CertificateTemplate_ObjectID_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate_ObjectID) *pb.ObjectId {
	if in == nil {
		return nil
	}
	out := &pb.ObjectId{}
	if in.ObjectIDPath != nil {
		out.ObjectIdPath = make([]int32, len(in.ObjectIDPath))
		for i, v := range in.ObjectIDPath {
			out.ObjectIdPath[i] = int32(v)
		}
	}
	return out
}

func CertificateTemplate_X509Parameters_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters) *krm.CertificateTemplate_X509Parameters {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate_X509Parameters{}
	out.KeyUsage = CertificateTemplate_KeyUsage_FromProto(mapCtx, in.GetKeyUsage())
	out.CAOptions = CertificateTemplate_CAOptions_FromProto(mapCtx, in.GetCaOptions())
	out.PolicyIds = direct.Slice_FromProto[pb.ObjectId, krm.CertificateTemplate_ObjectID](mapCtx, in.GetPolicyIds(), CertificateTemplate_ObjectID_FromProto)
	out.AiaOcspServers = in.GetAiaOcspServers()
	out.AdditionalExtensions = direct.Slice_FromProto[pb.X509Extension, krm.CertificateTemplate_X509Extension](mapCtx, in.GetAdditionalExtensions(), CertificateTemplate_X509Extension_FromProto)
	return out
}

func CertificateTemplate_X509Parameters_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate_X509Parameters) *pb.X509Parameters {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters{}
	out.KeyUsage = CertificateTemplate_KeyUsage_ToProto(mapCtx, in.KeyUsage)
	out.CaOptions = CertificateTemplate_CAOptions_ToProto(mapCtx, in.CAOptions)
	out.PolicyIds = direct.Slice_ToProto[krm.CertificateTemplate_ObjectID, pb.ObjectId](mapCtx, in.PolicyIds, CertificateTemplate_ObjectID_ToProto)
	out.AiaOcspServers = in.AiaOcspServers
	out.AdditionalExtensions = direct.Slice_ToProto[krm.CertificateTemplate_X509Extension, pb.X509Extension](mapCtx, in.AdditionalExtensions, CertificateTemplate_X509Extension_ToProto)
	return out
}

func CertificateTemplate_CAOptions_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters_CaOptions) *krm.CertificateTemplate_CAOptions {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate_CAOptions{}
	out.IsCA = in.IsCa
	if in.MaxIssuerPathLength != nil {
		out.MaxIssuerPathLength = direct.LazyPtr(int64(*in.MaxIssuerPathLength))
	}
	return out
}

func CertificateTemplate_CAOptions_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate_CAOptions) *pb.X509Parameters_CaOptions {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters_CaOptions{}
	out.IsCa = in.IsCA
	if in.MaxIssuerPathLength != nil {
		out.MaxIssuerPathLength = direct.LazyPtr(int32(*in.MaxIssuerPathLength))
	}
	return out
}

func CertificateTemplate_X509Extension_FromProto(mapCtx *direct.MapContext, in *pb.X509Extension) *krm.CertificateTemplate_X509Extension {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate_X509Extension{}
	out.ObjectID = CertificateTemplate_ObjectID_FromProto(mapCtx, in.GetObjectId())
	out.Critical = direct.LazyPtr(in.GetCritical())
	if in.Value != nil {
		out.Value = direct.LazyPtr(string(in.Value))
	}
	return out
}

func CertificateTemplate_X509Extension_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate_X509Extension) *pb.X509Extension {
	if in == nil {
		return nil
	}
	out := &pb.X509Extension{}
	out.ObjectId = CertificateTemplate_ObjectID_ToProto(mapCtx, in.ObjectID)
	out.Critical = direct.ValueOf(in.Critical)
	if in.Value != nil {
		out.Value = []byte(*in.Value)
	}
	return out
}

func CertificateTemplate_KeyUsage_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage) *krm.CertificateTemplate_KeyUsage {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate_KeyUsage{}
	out.BaseKeyUsage = CertificateTemplate_KeyUsage_KeyUsageOptions_FromProto(mapCtx, in.GetBaseKeyUsage())
	out.ExtendedKeyUsage = CertificateTemplate_KeyUsage_ExtendedKeyUsageOptions_FromProto(mapCtx, in.GetExtendedKeyUsage())
	out.UnknownExtendedKeyUsages = direct.Slice_FromProto[pb.ObjectId, krm.CertificateTemplate_ObjectID](mapCtx, in.GetUnknownExtendedKeyUsages(), CertificateTemplate_ObjectID_FromProto)
	return out
}

func CertificateTemplate_KeyUsage_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate_KeyUsage) *pb.KeyUsage {
	if in == nil {
		return nil
	}
	out := &pb.KeyUsage{}
	out.BaseKeyUsage = CertificateTemplate_KeyUsage_KeyUsageOptions_ToProto(mapCtx, in.BaseKeyUsage)
	out.ExtendedKeyUsage = CertificateTemplate_KeyUsage_ExtendedKeyUsageOptions_ToProto(mapCtx, in.ExtendedKeyUsage)
	out.UnknownExtendedKeyUsages = direct.Slice_ToProto[krm.CertificateTemplate_ObjectID, pb.ObjectId](mapCtx, in.UnknownExtendedKeyUsages, CertificateTemplate_ObjectID_ToProto)
	return out
}

func CertificateTemplate_KeyUsage_KeyUsageOptions_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage_KeyUsageOptions) *krm.CertificateTemplate_KeyUsage_KeyUsageOptions {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate_KeyUsage_KeyUsageOptions{}
	out.CertSign = direct.LazyPtr(in.GetCertSign())
	out.ContentCommitment = direct.LazyPtr(in.GetContentCommitment())
	out.CrlSign = direct.LazyPtr(in.GetCrlSign())
	out.DataEncipherment = direct.LazyPtr(in.GetDataEncipherment())
	out.DecipherOnly = direct.LazyPtr(in.GetDecipherOnly())
	out.DigitalSignature = direct.LazyPtr(in.GetDigitalSignature())
	out.EncipherOnly = direct.LazyPtr(in.GetEncipherOnly())
	out.KeyAgreement = direct.LazyPtr(in.GetKeyAgreement())
	out.KeyEncipherment = direct.LazyPtr(in.GetKeyEncipherment())
	return out
}

func CertificateTemplate_KeyUsage_KeyUsageOptions_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate_KeyUsage_KeyUsageOptions) *pb.KeyUsage_KeyUsageOptions {
	if in == nil {
		return nil
	}
	out := &pb.KeyUsage_KeyUsageOptions{}
	out.CertSign = direct.ValueOf(in.CertSign)
	out.ContentCommitment = direct.ValueOf(in.ContentCommitment)
	out.CrlSign = direct.ValueOf(in.CrlSign)
	out.DataEncipherment = direct.ValueOf(in.DataEncipherment)
	out.DecipherOnly = direct.ValueOf(in.DecipherOnly)
	out.DigitalSignature = direct.ValueOf(in.DigitalSignature)
	out.EncipherOnly = direct.ValueOf(in.EncipherOnly)
	out.KeyAgreement = direct.ValueOf(in.KeyAgreement)
	out.KeyEncipherment = direct.ValueOf(in.KeyEncipherment)
	return out
}

func CertificateTemplate_KeyUsage_ExtendedKeyUsageOptions_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage_ExtendedKeyUsageOptions) *krm.CertificateTemplate_KeyUsage_ExtendedKeyUsageOptions {
	if in == nil {
		return nil
	}
	out := &krm.CertificateTemplate_KeyUsage_ExtendedKeyUsageOptions{}
	out.ClientAuth = direct.LazyPtr(in.GetClientAuth())
	out.CodeSigning = direct.LazyPtr(in.GetCodeSigning())
	out.EmailProtection = direct.LazyPtr(in.GetEmailProtection())
	out.OcspSigning = direct.LazyPtr(in.GetOcspSigning())
	out.ServerAuth = direct.LazyPtr(in.GetServerAuth())
	out.TimeStamping = direct.LazyPtr(in.GetTimeStamping())
	return out
}

func CertificateTemplate_KeyUsage_ExtendedKeyUsageOptions_ToProto(mapCtx *direct.MapContext, in *krm.CertificateTemplate_KeyUsage_ExtendedKeyUsageOptions) *pb.KeyUsage_ExtendedKeyUsageOptions {
	if in == nil {
		return nil
	}
	out := &pb.KeyUsage_ExtendedKeyUsageOptions{}
	out.ClientAuth = direct.ValueOf(in.ClientAuth)
	out.CodeSigning = direct.ValueOf(in.CodeSigning)
	out.EmailProtection = direct.ValueOf(in.EmailProtection)
	out.OcspSigning = direct.ValueOf(in.OcspSigning)
	out.ServerAuth = direct.ValueOf(in.ServerAuth)
	out.TimeStamping = direct.ValueOf(in.TimeStamping)
	return out
}

func PrivateCACertificateTemplateStatus_FromProto(mapCtx *direct.MapContext, in *pb.CertificateTemplate) *krm.PrivateCACertificateTemplateStatus {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCACertificateTemplateStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func PrivateCACertificateTemplateStatus_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCACertificateTemplateStatus) *pb.CertificateTemplate {
	if in == nil {
		return nil
	}
	out := &pb.CertificateTemplate{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

// CertificateConfig mapping functions because KRM uses value fields instead of pointer fields
// for SubjectConfig and X509Config to support required-field OpenAPI validation.
func CertificateConfig_FromProto(mapCtx *direct.MapContext, in *pb.CertificateConfig) *krm.CertificateConfig {
	if in == nil {
		return nil
	}
	out := &krm.CertificateConfig{}
	if sc := CertificateConfig_SubjectConfig_FromProto(mapCtx, in.GetSubjectConfig()); sc != nil {
		out.SubjectConfig = *sc
	}
	if xc := CertificateConfig_X509Config_FromProto(mapCtx, in.GetX509Config()); xc != nil {
		out.X509Config = *xc
	}
	return out
}

func CertificateConfig_ToProto(mapCtx *direct.MapContext, in *krm.CertificateConfig) *pb.CertificateConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateConfig{}
	out.SubjectConfig = CertificateConfig_SubjectConfig_ToProto(mapCtx, &in.SubjectConfig)
	out.X509Config = CertificateConfig_X509Config_ToProto(mapCtx, &in.X509Config)
	return out
}

// CertificateConfig_SubjectConfig mapping functions because KRM uses value field instead of pointer field
// for Subject.
func CertificateConfig_SubjectConfig_FromProto(mapCtx *direct.MapContext, in *pb.CertificateConfig_SubjectConfig) *krm.CertificateConfig_SubjectConfig {
	if in == nil {
		return nil
	}
	out := &krm.CertificateConfig_SubjectConfig{}
	if s := Subject_FromProto(mapCtx, in.GetSubject()); s != nil {
		out.Subject = *s
	}
	out.SubjectAltName = SubjectAltNames_FromProto(mapCtx, in.GetSubjectAltName())
	return out
}

func CertificateConfig_SubjectConfig_ToProto(mapCtx *direct.MapContext, in *krm.CertificateConfig_SubjectConfig) *pb.CertificateConfig_SubjectConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateConfig_SubjectConfig{}
	out.Subject = Subject_ToProto(mapCtx, &in.Subject)
	out.SubjectAltName = SubjectAltNames_ToProto(mapCtx, in.SubjectAltName)
	return out
}

func CertificateAuthority_ObjectID_FromProto(mapCtx *direct.MapContext, in *pb.ObjectId) *krm.CertificateAuthority_ObjectID {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_ObjectID{}
	if path := in.GetObjectIdPath(); path != nil {
		out.ObjectIDPath = make([]int64, len(path))
		for i, v := range path {
			out.ObjectIDPath[i] = int64(v)
		}
	}
	return out
}

func CertificateAuthority_ObjectID_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_ObjectID) *pb.ObjectId {
	if in == nil {
		return nil
	}
	out := &pb.ObjectId{}
	if path := in.ObjectIDPath; path != nil {
		out.ObjectIdPath = make([]int32, len(path))
		for i, v := range path {
			out.ObjectIdPath[i] = int32(v)
		}
	}
	return out
}

func CertificateAuthority_ObjectIDStatus_FromProto(mapCtx *direct.MapContext, in *pb.ObjectId) *krm.CertificateAuthority_ObjectIDStatus {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_ObjectIDStatus{}
	if path := in.GetObjectIdPath(); path != nil {
		out.ObjectIDPath = make([]int64, len(path))
		for i, v := range path {
			out.ObjectIDPath[i] = int64(v)
		}
	}
	return out
}

func CertificateAuthority_ObjectIDStatus_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_ObjectIDStatus) *pb.ObjectId {
	if in == nil {
		return nil
	}
	out := &pb.ObjectId{}
	if path := in.ObjectIDPath; path != nil {
		out.ObjectIdPath = make([]int32, len(path))
		for i, v := range path {
			out.ObjectIdPath[i] = int32(v)
		}
	}
	return out
}

func CertificateAuthority_X509Extension_FromProto(mapCtx *direct.MapContext, in *pb.X509Extension) *krm.CertificateAuthority_X509Extension {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_X509Extension{}
	if oid := CertificateAuthority_ObjectID_FromProto(mapCtx, in.GetObjectId()); oid != nil {
		out.ObjectID = *oid
	}
	out.Critical = direct.LazyPtr(in.GetCritical())
	out.Value = string(in.GetValue())
	return out
}

func CertificateAuthority_X509Extension_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_X509Extension) *pb.X509Extension {
	if in == nil {
		return nil
	}
	out := &pb.X509Extension{}
	out.ObjectId = CertificateAuthority_ObjectID_ToProto(mapCtx, &in.ObjectID)
	out.Critical = direct.ValueOf(in.Critical)
	out.Value = []byte(in.Value)
	return out
}

func CertificateAuthority_X509ExtensionStatus_FromProto(mapCtx *direct.MapContext, in *pb.X509Extension) *krm.CertificateAuthority_X509ExtensionStatus {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_X509ExtensionStatus{}
	out.ObjectID = CertificateAuthority_ObjectIDStatus_FromProto(mapCtx, in.GetObjectId())
	out.Critical = direct.LazyPtr(in.GetCritical())
	if val := in.GetValue(); val != nil {
		strVal := string(val)
		out.Value = &strVal
	}
	return out
}

func CertificateAuthority_X509ExtensionStatus_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_X509ExtensionStatus) *pb.X509Extension {
	if in == nil {
		return nil
	}
	out := &pb.X509Extension{}
	out.ObjectId = CertificateAuthority_ObjectIDStatus_ToProto(mapCtx, in.ObjectID)
	out.Critical = direct.ValueOf(in.Critical)
	if in.Value != nil {
		out.Value = []byte(*in.Value)
	}
	return out
}

func CertificateAuthority_KeyUsage_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage) *krm.CertificateAuthority_KeyUsage {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_KeyUsage{}
	out.BaseKeyUsage = KeyUsage_KeyUsageOptions_FromProto(mapCtx, in.GetBaseKeyUsage())
	out.ExtendedKeyUsage = KeyUsage_ExtendedKeyUsageOptions_FromProto(mapCtx, in.GetExtendedKeyUsage())
	out.UnknownExtendedKeyUsages = direct.Slice_FromProto(mapCtx, in.GetUnknownExtendedKeyUsages(), CertificateAuthority_ObjectID_FromProto)
	return out
}

func CertificateAuthority_KeyUsage_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_KeyUsage) *pb.KeyUsage {
	if in == nil {
		return nil
	}
	out := &pb.KeyUsage{}
	out.BaseKeyUsage = KeyUsage_KeyUsageOptions_ToProto(mapCtx, in.BaseKeyUsage)
	out.ExtendedKeyUsage = KeyUsage_ExtendedKeyUsageOptions_ToProto(mapCtx, in.ExtendedKeyUsage)
	out.UnknownExtendedKeyUsages = direct.Slice_ToProto(mapCtx, in.UnknownExtendedKeyUsages, CertificateAuthority_ObjectID_ToProto)
	return out
}

func CertificateAuthority_KeyUsageStatus_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage) *krm.CertificateAuthority_KeyUsageStatus {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_KeyUsageStatus{}
	out.BaseKeyUsage = KeyUsage_KeyUsageOptions_FromProto(mapCtx, in.GetBaseKeyUsage())
	out.ExtendedKeyUsage = KeyUsage_ExtendedKeyUsageOptions_FromProto(mapCtx, in.GetExtendedKeyUsage())
	out.UnknownExtendedKeyUsages = direct.Slice_FromProto(mapCtx, in.GetUnknownExtendedKeyUsages(), CertificateAuthority_ObjectIDStatus_FromProto)
	return out
}

func CertificateAuthority_KeyUsageStatus_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_KeyUsageStatus) *pb.KeyUsage {
	if in == nil {
		return nil
	}
	out := &pb.KeyUsage{}
	out.BaseKeyUsage = KeyUsage_KeyUsageOptions_ToProto(mapCtx, in.BaseKeyUsage)
	out.ExtendedKeyUsage = KeyUsage_ExtendedKeyUsageOptions_ToProto(mapCtx, in.ExtendedKeyUsage)
	out.UnknownExtendedKeyUsages = direct.Slice_ToProto(mapCtx, in.UnknownExtendedKeyUsages, CertificateAuthority_ObjectIDStatus_ToProto)
	return out
}

func CertificateConfig_X509Config_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters) *krm.CertificateConfig_X509Config {
	if in == nil {
		return nil
	}
	out := &krm.CertificateConfig_X509Config{}
	out.AdditionalExtensions = direct.Slice_FromProto(mapCtx, in.GetAdditionalExtensions(), CertificateAuthority_X509Extension_FromProto)
	out.CaOptions = X509Parameters_CAOptions_FromProto(mapCtx, in.GetCaOptions())
	out.KeyUsage = CertificateAuthority_KeyUsage_FromProto(mapCtx, in.GetKeyUsage())
	out.PolicyIds = direct.Slice_FromProto(mapCtx, in.GetPolicyIds(), CertificateAuthority_ObjectID_FromProto)
	return out
}

func CertificateConfig_X509Config_ToProto(mapCtx *direct.MapContext, in *krm.CertificateConfig_X509Config) *pb.X509Parameters {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters{}
	out.AdditionalExtensions = direct.Slice_ToProto(mapCtx, in.AdditionalExtensions, CertificateAuthority_X509Extension_ToProto)
	out.CaOptions = X509Parameters_CAOptions_ToProto(mapCtx, in.CaOptions)
	out.KeyUsage = CertificateAuthority_KeyUsage_ToProto(mapCtx, in.KeyUsage)
	out.PolicyIds = direct.Slice_ToProto(mapCtx, in.PolicyIds, CertificateAuthority_ObjectID_ToProto)
	return out
}

func CertificateDescription_X509Description_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters) *krm.CertificateDescription_X509Description {
	if in == nil {
		return nil
	}
	out := &krm.CertificateDescription_X509Description{}
	out.AdditionalExtensions = direct.Slice_FromProto(mapCtx, in.GetAdditionalExtensions(), CertificateAuthority_X509ExtensionStatus_FromProto)
	out.AiaOcspServers = in.GetAiaOcspServers()
	out.CaOptions = CertificateAuthority_CaOptionsStatus_FromProto(mapCtx, in.GetCaOptions())
	out.KeyUsage = CertificateAuthority_KeyUsageStatus_FromProto(mapCtx, in.GetKeyUsage())
	out.PolicyIds = direct.Slice_FromProto(mapCtx, in.GetPolicyIds(), CertificateAuthority_ObjectIDStatus_FromProto)
	return out
}

func CertificateDescription_X509Description_ToProto(mapCtx *direct.MapContext, in *krm.CertificateDescription_X509Description) *pb.X509Parameters {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters{}
	out.AdditionalExtensions = direct.Slice_ToProto(mapCtx, in.AdditionalExtensions, CertificateAuthority_X509ExtensionStatus_ToProto)
	out.AiaOcspServers = in.AiaOcspServers
	out.CaOptions = CertificateAuthority_CaOptionsStatus_ToProto(mapCtx, in.CaOptions)
	out.KeyUsage = CertificateAuthority_KeyUsageStatus_ToProto(mapCtx, in.KeyUsage)
	out.PolicyIds = direct.Slice_ToProto(mapCtx, in.PolicyIds, CertificateAuthority_ObjectIDStatus_ToProto)
	return out
}

func CertificateAuthority_CaOptionsStatus_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters_CaOptions) *krm.CertificateAuthority_CaOptionsStatus {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_CaOptionsStatus{}
	out.IsCa = in.IsCa
	if in.MaxIssuerPathLength != nil {
		val := int64(*in.MaxIssuerPathLength)
		out.MaxIssuerPathLength = &val
	}
	return out
}

func CertificateAuthority_CaOptionsStatus_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_CaOptionsStatus) *pb.X509Parameters_CaOptions {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters_CaOptions{}
	out.IsCa = in.IsCa
	if in.MaxIssuerPathLength != nil {
		val := int32(*in.MaxIssuerPathLength)
		out.MaxIssuerPathLength = &val
	}
	return out
}

func PublicKey_FromProto(mapCtx *direct.MapContext, in *pb.PublicKey) *krm.PublicKey {
	if in == nil {
		return nil
	}
	out := &krm.PublicKey{}
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	if key := in.GetKey(); key != nil {
		strKey := string(key)
		out.Key = &strKey
	}
	return out
}

func PublicKey_ToProto(mapCtx *direct.MapContext, in *krm.PublicKey) *pb.PublicKey {
	if in == nil {
		return nil
	}
	out := &pb.PublicKey{}
	out.Format = direct.Enum_ToProto[pb.PublicKey_KeyFormat](mapCtx, in.Format)
	if in.Key != nil {
		out.Key = []byte(*in.Key)
	}
	return out
}

func SubjectAltNamesStatus_FromProto(mapCtx *direct.MapContext, in *pb.SubjectAltNames) *krm.SubjectAltNamesStatus {
	if in == nil {
		return nil
	}
	out := &krm.SubjectAltNamesStatus{}
	out.DNSNames = in.DnsNames
	out.Uris = in.Uris
	out.EmailAddresses = in.EmailAddresses
	out.IPAddresses = in.IpAddresses
	out.CustomSans = direct.Slice_FromProto(mapCtx, in.CustomSans, CertificateAuthority_X509ExtensionStatus_FromProto)
	return out
}

func SubjectAltNamesStatus_ToProto(mapCtx *direct.MapContext, in *krm.SubjectAltNamesStatus) *pb.SubjectAltNames {
	if in == nil {
		return nil
	}
	out := &pb.SubjectAltNames{}
	out.DnsNames = in.DNSNames
	out.Uris = in.Uris
	out.EmailAddresses = in.EmailAddresses
	out.IpAddresses = in.IPAddresses
	out.CustomSans = direct.Slice_ToProto(mapCtx, in.CustomSans, CertificateAuthority_X509ExtensionStatus_ToProto)
	return out
}

func PrivateCACertificateAuthoritySpec_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority) *krm.PrivateCACertificateAuthoritySpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCACertificateAuthoritySpec{}
	if t := direct.Enum_FromProto(mapCtx, in.GetType()); t != nil {
		out.Type = *t
	}
	if config := CertificateConfig_FromProto(mapCtx, in.GetConfig()); config != nil {
		out.Config = *config
	}
	if lifetime := direct.StringDuration_FromProto(mapCtx, in.GetLifetime()); lifetime != nil {
		out.Lifetime = *lifetime
	}
	if keySpec := CertificateAuthority_KeyVersionSpec_FromProto(mapCtx, in.GetKeySpec()); keySpec != nil {
		out.KeySpec = *keySpec
	}
	if in.GetGcsBucket() != "" {
		out.GcsBucketRef = &storagev1beta1.StorageBucketRef{External: in.GetGcsBucket()}
	}
	return out
}

func PrivateCACertificateAuthoritySpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCACertificateAuthoritySpec) *pb.CertificateAuthority {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority{}
	out.Type = direct.Enum_ToProto[pb.CertificateAuthority_Type](mapCtx, &in.Type)
	out.Config = CertificateConfig_ToProto(mapCtx, &in.Config)
	out.Lifetime = direct.StringDuration_ToProto(mapCtx, &in.Lifetime)
	out.KeySpec = CertificateAuthority_KeyVersionSpec_ToProto(mapCtx, &in.KeySpec)
	if in.GcsBucketRef != nil {
		out.GcsBucket = in.GcsBucketRef.External
	}
	return out
}

func PrivateCACertificateAuthorityStatus_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority) *krm.PrivateCACertificateAuthorityStatus {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCACertificateAuthorityStatus{}
	out.AccessUrls = CertificateAuthority_AccessUrls_FromProto(mapCtx, in.GetAccessUrls())
	out.CaCertificateDescriptions = direct.Slice_FromProto(mapCtx, in.GetCaCertificateDescriptions(), CertificateDescription_FromProto)
	if in.GetConfig() != nil {
		out.Config = &krm.CertificateAuthority_ConfigStatus{}
		out.Config.PublicKey = PublicKey_FromProto(mapCtx, in.GetConfig().GetPublicKey())
		if in.GetConfig().GetX509Config() != nil {
			out.Config.X509Config = &krm.CertificateAuthority_X509ConfigStatus{
				AiaOcspServers: in.GetConfig().GetX509Config().GetAiaOcspServers(),
			}
		}
	}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.PemCaCertificates = in.GetPemCaCertificates()
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.SubordinateConfig = SubordinateConfig_FromProto(mapCtx, in.GetSubordinateConfig())
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func PrivateCACertificateAuthorityStatus_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCACertificateAuthorityStatus) *pb.CertificateAuthority {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority{}
	out.AccessUrls = CertificateAuthority_AccessUrls_ToProto(mapCtx, in.AccessUrls)
	out.CaCertificateDescriptions = direct.Slice_ToProto(mapCtx, in.CaCertificateDescriptions, CertificateDescription_ToProto)
	if in.Config != nil {
		out.Config = &pb.CertificateConfig{}
		out.Config.PublicKey = PublicKey_ToProto(mapCtx, in.Config.PublicKey)
		if in.Config.X509Config != nil {
			out.Config.X509Config = &pb.X509Parameters{
				AiaOcspServers: in.Config.X509Config.AiaOcspServers,
			}
		}
	}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.PemCaCertificates = in.PemCaCertificates
	out.State = direct.Enum_ToProto[pb.CertificateAuthority_State](mapCtx, in.State)
	out.SubordinateConfig = SubordinateConfig_ToProto(mapCtx, in.SubordinateConfig)
	out.Tier = direct.Enum_ToProto[pb.CaPool_Tier](mapCtx, in.Tier)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func CertificateAuthority_KeyVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_KeyVersionSpec) *krm.CertificateAuthority_KeyVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_KeyVersionSpec{}
	if in.GetKeyVersion() != nil {
		if alg, ok := in.GetKeyVersion().(*pb.CertificateAuthority_KeyVersionSpec_Algorithm); ok {
			out.Algorithm = direct.Enum_FromProto(mapCtx, alg.Algorithm)
		}
		if kms, ok := in.GetKeyVersion().(*pb.CertificateAuthority_KeyVersionSpec_CloudKmsKeyVersion); ok {
			out.CloudKmsKeyVersionRef = &kmsrefs.KMSCryptoKeyVersionRef{
				External: kms.CloudKmsKeyVersion,
			}
		}
	}
	return out
}

func CertificateAuthority_KeyVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_KeyVersionSpec) *pb.CertificateAuthority_KeyVersionSpec {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_KeyVersionSpec{}
	if in.Algorithm != nil {
		out.KeyVersion = &pb.CertificateAuthority_KeyVersionSpec_Algorithm{
			Algorithm: direct.Enum_ToProto[pb.CertificateAuthority_SignHashAlgorithm](mapCtx, in.Algorithm),
		}
	}
	if in.CloudKmsKeyVersionRef != nil {
		out.KeyVersion = &pb.CertificateAuthority_KeyVersionSpec_CloudKmsKeyVersion{
			CloudKmsKeyVersion: in.CloudKmsKeyVersionRef.External,
		}
	}
	return out
}

func CertificateAuthority_AccessUrls_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_AccessUrls) *krm.CertificateAuthority_AccessUrls {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_AccessUrls{}
	out.CaCertificateAccessUrl = direct.LazyPtr(in.GetCaCertificateAccessUrl())
	out.CrlAccessUrls = in.GetCrlAccessUrls()
	return out
}

func CertificateAuthority_AccessUrls_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_AccessUrls) *pb.CertificateAuthority_AccessUrls {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_AccessUrls{}
	out.CaCertificateAccessUrl = direct.ValueOf(in.CaCertificateAccessUrl)
	out.CrlAccessUrls = in.CrlAccessUrls
	return out
}
