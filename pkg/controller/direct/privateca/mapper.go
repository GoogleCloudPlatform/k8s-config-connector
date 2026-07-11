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

func CertificatePublicKey_FromProto(mapCtx *direct.MapContext, in *pb.PublicKey) *krm.CertificatePublicKey {
	if in == nil {
		return nil
	}
	out := &krm.CertificatePublicKey{}
	out.Key = string(in.GetKey())
	if format := direct.Enum_FromProto(mapCtx, in.GetFormat()); format != nil {
		out.Format = *format
	}
	return out
}

func CertificatePublicKey_ToProto(mapCtx *direct.MapContext, in *krm.CertificatePublicKey) *pb.PublicKey {
	if in == nil {
		return nil
	}
	out := &pb.PublicKey{}
	out.Key = []byte(in.Key)
	out.Format = direct.Enum_ToProto[pb.PublicKey_KeyFormat](mapCtx, direct.LazyPtr(in.Format))
	return out
}

func CertificateSubjectAltName_FromProto(mapCtx *direct.MapContext, in *pb.SubjectAltNames) *krm.CertificateSubjectAltName {
	if in == nil {
		return nil
	}
	out := &krm.CertificateSubjectAltName{}
	out.DnsNames = in.GetDnsNames()
	out.EmailAddresses = in.GetEmailAddresses()
	out.IpAddresses = in.GetIpAddresses()
	out.Uris = in.GetUris()
	return out
}

func CertificateSubjectAltName_ToProto(mapCtx *direct.MapContext, in *krm.CertificateSubjectAltName) *pb.SubjectAltNames {
	if in == nil {
		return nil
	}
	out := &pb.SubjectAltNames{}
	out.DnsNames = in.DnsNames
	out.EmailAddresses = in.EmailAddresses
	out.IpAddresses = in.IpAddresses
	out.Uris = in.Uris
	return out
}

func CertificateAdditionalExtensions_FromProto(mapCtx *direct.MapContext, in *pb.X509Extension) *krm.CertificateAdditionalExtensions {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAdditionalExtensions{}
	if objID := in.GetObjectId(); objID != nil {
		out.ObjectId = krm.CertificateObjectId{
			ObjectIdPath: int32To64Slice(objID.GetObjectIdPath()),
		}
	}
	out.Critical = direct.LazyPtr(in.GetCritical())
	out.Value = string(in.GetValue())
	return out
}

func CertificateAdditionalExtensions_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAdditionalExtensions) *pb.X509Extension {
	if in == nil {
		return nil
	}
	out := &pb.X509Extension{}
	if in.ObjectId.ObjectIdPath != nil {
		out.ObjectId = &pb.ObjectId{
			ObjectIdPath: int64To32Slice(in.ObjectId.ObjectIdPath),
		}
	}
	out.Critical = direct.ValueOf(in.Critical)
	out.Value = []byte(in.Value)
	return out
}

func CertificatePolicyIds_FromProto(mapCtx *direct.MapContext, in *pb.ObjectId) *krm.CertificatePolicyIds {
	if in == nil {
		return nil
	}
	out := &krm.CertificatePolicyIds{}
	out.ObjectIdPath = int32To64Slice(in.GetObjectIdPath())
	return out
}

func CertificatePolicyIds_ToProto(mapCtx *direct.MapContext, in *krm.CertificatePolicyIds) *pb.ObjectId {
	if in == nil {
		return nil
	}
	out := &pb.ObjectId{}
	out.ObjectIdPath = int64To32Slice(in.ObjectIdPath)
	return out
}

func CertificateUnknownExtendedKeyUsages_FromProto(mapCtx *direct.MapContext, in *pb.ObjectId) *krm.CertificateUnknownExtendedKeyUsages {
	if in == nil {
		return nil
	}
	out := &krm.CertificateUnknownExtendedKeyUsages{}
	out.ObjectIdPath = in.GetObjectIdPath()
	return out
}

func CertificateUnknownExtendedKeyUsages_ToProto(mapCtx *direct.MapContext, in *krm.CertificateUnknownExtendedKeyUsages) *pb.ObjectId {
	if in == nil {
		return nil
	}
	out := &pb.ObjectId{}
	out.ObjectIdPath = in.ObjectIdPath
	return out
}

func CertificateCaOptions_FromProto(mapCtx *direct.MapContext, in *pb.X509Parameters_CaOptions) *krm.CertificateCaOptions {
	if in == nil {
		return nil
	}
	out := &krm.CertificateCaOptions{}
	if in.IsCa != nil {
		out.IsCa = direct.PtrTo(*in.IsCa)
	}
	if in.MaxIssuerPathLength != nil {
		out.MaxIssuerPathLength = direct.PtrTo(int64(*in.MaxIssuerPathLength))
	}
	return out
}

func CertificateCaOptions_ToProto(mapCtx *direct.MapContext, in *krm.CertificateCaOptions) *pb.X509Parameters_CaOptions {
	if in == nil {
		return nil
	}
	out := &pb.X509Parameters_CaOptions{}
	if in.IsCa != nil {
		out.IsCa = direct.PtrTo(direct.ValueOf(in.IsCa))
	} else if in.NonCa != nil && direct.ValueOf(in.NonCa) {
		out.IsCa = direct.PtrTo(false)
	}

	if in.MaxIssuerPathLength != nil {
		out.MaxIssuerPathLength = direct.PtrTo(int32(direct.ValueOf(in.MaxIssuerPathLength)))
	} else if in.ZeroMaxIssuerPathLength != nil && direct.ValueOf(in.ZeroMaxIssuerPathLength) {
		out.MaxIssuerPathLength = direct.PtrTo(int32(0))
	}
	return out
}

func PrivateCACertificateSpec_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) *krm.PrivateCACertificateSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCACertificateSpec{}
	if in.GetPemCsr() != "" {
		out.PemCsr = direct.LazyPtr(in.GetPemCsr())
	}
	out.Config = Certificate_Config_FromProto(mapCtx, in.GetConfig())
	if lifetime := direct.StringDuration_FromProto(mapCtx, in.GetLifetime()); lifetime != nil {
		out.Lifetime = *lifetime
	}
	if in.GetCertificateTemplate() != "" {
		out.CertificateTemplateRef = &krm.PrivateCACertificateTemplateRef{External: in.GetCertificateTemplate()}
	}
	out.SubjectMode = direct.Enum_FromProto(mapCtx, in.GetSubjectMode())
	return out
}

func PrivateCACertificateSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCACertificateSpec) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	if in.PemCsr != nil {
		out.CertificateConfig = &pb.Certificate_PemCsr{
			PemCsr: direct.ValueOf(in.PemCsr),
		}
	}
	if in.Config != nil {
		out.CertificateConfig = &pb.Certificate_Config{
			Config: Certificate_Config_ToProto(mapCtx, in.Config),
		}
	}
	out.Lifetime = direct.StringDuration_ToProto(mapCtx, direct.LazyPtr(in.Lifetime))
	if in.CertificateTemplateRef != nil {
		out.CertificateTemplate = in.CertificateTemplateRef.External
	}
	out.SubjectMode = direct.Enum_ToProto[pb.SubjectRequestMode](mapCtx, in.SubjectMode)
	return out
}

func CertificateRevocationDetailsStatus_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_RevocationDetails) *krm.CertificateRevocationDetailsStatus {
	if in == nil {
		return nil
	}
	out := &krm.CertificateRevocationDetailsStatus{}
	out.RevocationState = direct.Enum_FromProto(mapCtx, in.GetRevocationState())
	out.RevocationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevocationTime())
	return out
}

func CertificateRevocationDetailsStatus_ToProto(mapCtx *direct.MapContext, in *krm.CertificateRevocationDetailsStatus) *pb.Certificate_RevocationDetails {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_RevocationDetails{}
	out.RevocationState = direct.Enum_ToProto[pb.RevocationReason](mapCtx, in.RevocationState)
	out.RevocationTime = direct.StringTimestamp_ToProto(mapCtx, in.RevocationTime)
	return out
}

func CertificateCertificateDescriptionStatus_FromProto(mapCtx *direct.MapContext, in *pb.CertificateDescription) *krm.CertificateCertificateDescriptionStatus {
	if in == nil {
		return nil
	}
	out := &krm.CertificateCertificateDescriptionStatus{}
	out.AiaIssuingCertificateUrls = in.GetAiaIssuingCertificateUrls()
	if in.GetAuthorityKeyId() != nil {
		out.AuthorityKeyId = &krm.CertificateAuthorityKeyIdStatus{
			KeyId: direct.LazyPtr(in.GetAuthorityKeyId().GetKeyId()),
		}
	}
	if in.GetCertFingerprint() != nil {
		out.CertFingerprint = &krm.CertificateCertFingerprintStatus{
			Sha256Hash: direct.LazyPtr(in.GetCertFingerprint().GetSha256Hash()),
		}
	}
	out.CrlDistributionPoints = in.GetCrlDistributionPoints()
	if in.GetPublicKey() != nil {
		out.PublicKey = &krm.CertificatePublicKeyStatus{
			Format: direct.Enum_FromProto(mapCtx, in.GetPublicKey().GetFormat()),
			Key:    direct.LazyPtr(string(in.GetPublicKey().GetKey())),
		}
	}
	if in.GetSubjectDescription() != nil {
		sd := in.GetSubjectDescription()
		out.SubjectDescription = &krm.CertificateSubjectDescriptionStatus{
			HexSerialNumber: direct.LazyPtr(sd.GetHexSerialNumber()),
		}
		if sd.Lifetime != nil {
			out.SubjectDescription.Lifetime = direct.StringDuration_FromProto(mapCtx, sd.Lifetime)
		}
		if sd.NotAfterTime != nil {
			out.SubjectDescription.NotAfterTime = direct.StringTimestamp_FromProto(mapCtx, sd.NotAfterTime)
		}
		if sd.NotBeforeTime != nil {
			out.SubjectDescription.NotBeforeTime = direct.StringTimestamp_FromProto(mapCtx, sd.NotBeforeTime)
		}
		if sd.GetSubject() != nil {
			subj := sd.GetSubject()
			out.SubjectDescription.Subject = &krm.CertificateSubjectStatus{
				CommonName:         direct.LazyPtr(subj.GetCommonName()),
				CountryCode:        direct.LazyPtr(subj.GetCountryCode()),
				Locality:           direct.LazyPtr(subj.GetLocality()),
				Organization:       direct.LazyPtr(subj.GetOrganization()),
				OrganizationalUnit: direct.LazyPtr(subj.GetOrganizationalUnit()),
				PostalCode:         direct.LazyPtr(subj.GetPostalCode()),
				Province:           direct.LazyPtr(subj.GetProvince()),
				StreetAddress:      direct.LazyPtr(subj.GetStreetAddress()),
			}
		}
		if sd.GetSubjectAltName() != nil {
			san := sd.GetSubjectAltName()
			out.SubjectDescription.SubjectAltName = &krm.CertificateSubjectAltNameStatus{
				DnsNames:       san.GetDnsNames(),
				EmailAddresses: san.GetEmailAddresses(),
				IpAddresses:    san.GetIpAddresses(),
				Uris:           san.GetUris(),
			}
			for _, cs := range san.GetCustomSans() {
				var objID *krm.CertificateObjectIdStatus
				if cs.ObjectId != nil {
					objID = &krm.CertificateObjectIdStatus{
						ObjectIdPath: int32To64Slice(cs.GetObjectId().GetObjectIdPath()),
					}
				}
				out.SubjectDescription.SubjectAltName.CustomSans = append(out.SubjectDescription.SubjectAltName.CustomSans, krm.CertificateCustomSansStatus{
					Critical: direct.LazyPtr(cs.GetCritical()),
					Value:    direct.LazyPtr(string(cs.GetValue())),
					ObjectId: objID,
				})
			}
		}
	}
	if in.GetSubjectKeyId() != nil {
		out.SubjectKeyId = &krm.CertificateSubjectKeyIdStatus{
			KeyId: direct.LazyPtr(in.GetSubjectKeyId().GetKeyId()),
		}
	}
	if in.GetX509Description() != nil {
		x509 := in.GetX509Description()
		out.X509Description = &krm.CertificateX509DescriptionStatus{
			AiaOcspServers: x509.GetAiaOcspServers(),
		}
		if x509.GetCaOptions() != nil {
			out.X509Description.CaOptions = &krm.CertificateCaOptionsStatus{}
			if x509.GetCaOptions().IsCa != nil {
				out.X509Description.CaOptions.IsCa = direct.PtrTo(*x509.GetCaOptions().IsCa)
			}
			if x509.GetCaOptions().MaxIssuerPathLength != nil {
				out.X509Description.CaOptions.MaxIssuerPathLength = direct.PtrTo(int64(*x509.GetCaOptions().MaxIssuerPathLength))
			}
		}
		if x509.GetKeyUsage() != nil {
			ku := x509.GetKeyUsage()
			out.X509Description.KeyUsage = &krm.CertificateKeyUsageStatus{}
			if ku.GetBaseKeyUsage() != nil {
				bku := ku.GetBaseKeyUsage()
				out.X509Description.KeyUsage.BaseKeyUsage = &krm.CertificateBaseKeyUsageStatus{
					CertSign:          direct.LazyPtr(bku.GetCertSign()),
					ContentCommitment: direct.LazyPtr(bku.GetContentCommitment()),
					CrlSign:           direct.LazyPtr(bku.GetCrlSign()),
					DataEncipherment:  direct.LazyPtr(bku.GetDataEncipherment()),
					DecipherOnly:      direct.LazyPtr(bku.GetDecipherOnly()),
					DigitalSignature:  direct.LazyPtr(bku.GetDigitalSignature()),
					EncipherOnly:      direct.LazyPtr(bku.GetEncipherOnly()),
					KeyAgreement:      direct.LazyPtr(bku.GetKeyAgreement()),
					KeyEncipherment:   direct.LazyPtr(bku.GetKeyEncipherment()),
				}
			}
			if ku.GetExtendedKeyUsage() != nil {
				eku := ku.GetExtendedKeyUsage()
				out.X509Description.KeyUsage.ExtendedKeyUsage = &krm.CertificateExtendedKeyUsageStatus{
					ClientAuth:      direct.LazyPtr(eku.GetClientAuth()),
					CodeSigning:     direct.LazyPtr(eku.GetCodeSigning()),
					EmailProtection: direct.LazyPtr(eku.GetEmailProtection()),
					OcspSigning:     direct.LazyPtr(eku.GetOcspSigning()),
					ServerAuth:      direct.LazyPtr(eku.GetServerAuth()),
					TimeStamping:    direct.LazyPtr(eku.GetTimeStamping()),
				}
			}
			for _, ueku := range ku.GetUnknownExtendedKeyUsages() {
				out.X509Description.KeyUsage.UnknownExtendedKeyUsages = append(out.X509Description.KeyUsage.UnknownExtendedKeyUsages, krm.CertificateUnknownExtendedKeyUsagesStatus{
					ObjectIdPath: int32To64Slice(ueku.GetObjectIdPath()),
				})
			}
		}
		for _, pol := range x509.GetPolicyIds() {
			out.X509Description.PolicyIds = append(out.X509Description.PolicyIds, krm.CertificatePolicyIdsStatus{
				ObjectIdPath: int32To64Slice(pol.GetObjectIdPath()),
			})
		}
		for _, ext := range x509.GetAdditionalExtensions() {
			var objID *krm.CertificateObjectIdStatus
			if ext.ObjectId != nil {
				objID = &krm.CertificateObjectIdStatus{
					ObjectIdPath: int32To64Slice(ext.GetObjectId().GetObjectIdPath()),
				}
			}
			out.X509Description.AdditionalExtensions = append(out.X509Description.AdditionalExtensions, krm.CertificateAdditionalExtensionsStatus{
				Critical: direct.LazyPtr(ext.GetCritical()),
				Value:    direct.LazyPtr(string(ext.GetValue())),
				ObjectId: objID,
			})
		}
	}
	return out
}

func PrivateCACertificateStatus_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) *krm.PrivateCACertificateStatus {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCACertificateStatus{}
	out.CertificateDescription = CertificateCertificateDescriptionStatus_FromProto(mapCtx, in.GetCertificateDescription())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	if in.GetIssuerCertificateAuthority() != "" {
		out.IssuerCertificateAuthority = direct.LazyPtr(in.GetIssuerCertificateAuthority())
	}
	out.PemCertificate = direct.LazyPtr(in.GetPemCertificate())
	out.PemCertificateChain = in.GetPemCertificateChain()
	out.RevocationDetails = CertificateRevocationDetailsStatus_FromProto(mapCtx, in.GetRevocationDetails())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func PrivateCACertificateStatus_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCACertificateStatus) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	out.CertificateDescription = CertificateCertificateDescriptionStatus_ToProto(mapCtx, in.CertificateDescription)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	if in.IssuerCertificateAuthority != nil {
		out.IssuerCertificateAuthority = *in.IssuerCertificateAuthority
	}
	if in.PemCertificate != nil {
		out.PemCertificate = *in.PemCertificate
	}
	out.PemCertificateChain = in.PemCertificateChain
	out.RevocationDetails = CertificateRevocationDetailsStatus_ToProto(mapCtx, in.RevocationDetails)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func CertificateCertificateDescriptionStatus_ToProto(mapCtx *direct.MapContext, in *krm.CertificateCertificateDescriptionStatus) *pb.CertificateDescription {
	if in == nil {
		return nil
	}
	out := &pb.CertificateDescription{}
	out.AiaIssuingCertificateUrls = in.AiaIssuingCertificateUrls
	if in.AuthorityKeyId != nil {
		out.AuthorityKeyId = &pb.CertificateDescription_KeyId{
			KeyId: direct.ValueOf(in.AuthorityKeyId.KeyId),
		}
	}
	if in.CertFingerprint != nil {
		out.CertFingerprint = &pb.CertificateDescription_CertificateFingerprint{
			Sha256Hash: direct.ValueOf(in.CertFingerprint.Sha256Hash),
		}
	}
	out.CrlDistributionPoints = in.CrlDistributionPoints
	if in.PublicKey != nil {
		out.PublicKey = &pb.PublicKey{
			Format: direct.Enum_ToProto[pb.PublicKey_KeyFormat](mapCtx, in.PublicKey.Format),
			Key:    []byte(direct.ValueOf(in.PublicKey.Key)),
		}
	}
	if in.SubjectDescription != nil {
		sd := in.SubjectDescription
		out.SubjectDescription = &pb.CertificateDescription_SubjectDescription{
			HexSerialNumber: direct.ValueOf(sd.HexSerialNumber),
			NotBeforeTime:   direct.StringTimestamp_ToProto(mapCtx, sd.NotBeforeTime),
			NotAfterTime:    direct.StringTimestamp_ToProto(mapCtx, sd.NotAfterTime),
		}
		if sd.Lifetime != nil {
			out.SubjectDescription.Lifetime = direct.StringDuration_ToProto(mapCtx, sd.Lifetime)
		}
		if sd.Subject != nil {
			subj := sd.Subject
			out.SubjectDescription.Subject = &pb.Subject{
				CommonName:         direct.ValueOf(subj.CommonName),
				CountryCode:        direct.ValueOf(subj.CountryCode),
				Locality:           direct.ValueOf(subj.Locality),
				Organization:       direct.ValueOf(subj.Organization),
				OrganizationalUnit: direct.ValueOf(subj.OrganizationalUnit),
				PostalCode:         direct.ValueOf(subj.PostalCode),
				Province:           direct.ValueOf(subj.Province),
				StreetAddress:      direct.ValueOf(subj.StreetAddress),
			}
		}
		if sd.SubjectAltName != nil {
			san := sd.SubjectAltName
			out.SubjectDescription.SubjectAltName = &pb.SubjectAltNames{
				DnsNames:       san.DnsNames,
				EmailAddresses: san.EmailAddresses,
				IpAddresses:    san.IpAddresses,
				Uris:           san.Uris,
			}
			for _, cs := range san.CustomSans {
				var objID *pb.ObjectId
				if cs.ObjectId != nil {
					objID = &pb.ObjectId{
						ObjectIdPath: int64To32Slice(cs.ObjectId.ObjectIdPath),
					}
				}
				out.SubjectDescription.SubjectAltName.CustomSans = append(out.SubjectDescription.SubjectAltName.CustomSans, &pb.X509Extension{
					Critical: direct.ValueOf(cs.Critical),
					Value:    []byte(direct.ValueOf(cs.Value)),
					ObjectId: objID,
				})
			}
		}
	}
	if in.SubjectKeyId != nil {
		out.SubjectKeyId = &pb.CertificateDescription_KeyId{
			KeyId: direct.ValueOf(in.SubjectKeyId.KeyId),
		}
	}
	if in.X509Description != nil {
		x509 := in.X509Description
		out.X509Description = &pb.X509Parameters{
			AiaOcspServers: x509.AiaOcspServers,
		}
		if x509.CaOptions != nil {
			out.X509Description.CaOptions = &pb.X509Parameters_CaOptions{}
			if x509.CaOptions.IsCa != nil {
				isCa := direct.ValueOf(x509.CaOptions.IsCa)
				out.X509Description.CaOptions.IsCa = &isCa
			}
			if x509.CaOptions.MaxIssuerPathLength != nil {
				maxLen := int32(direct.ValueOf(x509.CaOptions.MaxIssuerPathLength))
				out.X509Description.CaOptions.MaxIssuerPathLength = &maxLen
			}
		}
		if x509.KeyUsage != nil {
			ku := x509.KeyUsage
			out.X509Description.KeyUsage = &pb.KeyUsage{}
			if ku.BaseKeyUsage != nil {
				bku := ku.BaseKeyUsage
				out.X509Description.KeyUsage.BaseKeyUsage = &pb.KeyUsage_KeyUsageOptions{
					CertSign:          direct.ValueOf(bku.CertSign),
					ContentCommitment: direct.ValueOf(bku.ContentCommitment),
					CrlSign:           direct.ValueOf(bku.CrlSign),
					DataEncipherment:  direct.ValueOf(bku.DataEncipherment),
					DecipherOnly:      direct.ValueOf(bku.DecipherOnly),
					DigitalSignature:  direct.ValueOf(bku.DigitalSignature),
					EncipherOnly:      direct.ValueOf(bku.EncipherOnly),
					KeyAgreement:      direct.ValueOf(bku.KeyAgreement),
					KeyEncipherment:   direct.ValueOf(bku.KeyEncipherment),
				}
			}
			if ku.ExtendedKeyUsage != nil {
				eku := ku.ExtendedKeyUsage
				out.X509Description.KeyUsage.ExtendedKeyUsage = &pb.KeyUsage_ExtendedKeyUsageOptions{
					ClientAuth:      direct.ValueOf(eku.ClientAuth),
					CodeSigning:     direct.ValueOf(eku.CodeSigning),
					EmailProtection: direct.ValueOf(eku.EmailProtection),
					OcspSigning:     direct.ValueOf(eku.OcspSigning),
					ServerAuth:      direct.ValueOf(eku.ServerAuth),
					TimeStamping:    direct.ValueOf(eku.TimeStamping),
				}
			}
			for _, ueku := range ku.UnknownExtendedKeyUsages {
				out.X509Description.KeyUsage.UnknownExtendedKeyUsages = append(out.X509Description.KeyUsage.UnknownExtendedKeyUsages, &pb.ObjectId{
					ObjectIdPath: int64To32Slice(ueku.ObjectIdPath),
				})
			}
		}
		for _, pol := range x509.PolicyIds {
			out.X509Description.PolicyIds = append(out.X509Description.PolicyIds, &pb.ObjectId{
				ObjectIdPath: int64To32Slice(pol.ObjectIdPath),
			})
		}
		for _, ext := range x509.AdditionalExtensions {
			var objID *pb.ObjectId
			if ext.ObjectId != nil {
				objID = &pb.ObjectId{
					ObjectIdPath: int64To32Slice(ext.ObjectId.ObjectIdPath),
				}
			}
			out.X509Description.AdditionalExtensions = append(out.X509Description.AdditionalExtensions, &pb.X509Extension{
				Critical: direct.ValueOf(ext.Critical),
				Value:    []byte(direct.ValueOf(ext.Value)),
				ObjectId: objID,
			})
		}
	}
	return out
}

func CertificateKeyUsage_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage) *krm.CertificateKeyUsage {
	if in == nil {
		return nil
	}
	out := &krm.CertificateKeyUsage{}
	out.BaseKeyUsage = CertificateBaseKeyUsage_FromProto(mapCtx, in.GetBaseKeyUsage())
	out.ExtendedKeyUsage = CertificateExtendedKeyUsage_FromProto(mapCtx, in.GetExtendedKeyUsage())
	out.UnknownExtendedKeyUsages = direct.Slice_FromProto(mapCtx, in.GetUnknownExtendedKeyUsages(), CertificateUnknownExtendedKeyUsages_FromProto)
	return out
}

func CertificateKeyUsage_ToProto(mapCtx *direct.MapContext, in *krm.CertificateKeyUsage) *pb.KeyUsage {
	if in == nil {
		return nil
	}
	out := &pb.KeyUsage{}
	out.BaseKeyUsage = CertificateBaseKeyUsage_ToProto(mapCtx, in.BaseKeyUsage)
	out.ExtendedKeyUsage = CertificateExtendedKeyUsage_ToProto(mapCtx, in.ExtendedKeyUsage)
	out.UnknownExtendedKeyUsages = direct.Slice_ToProto(mapCtx, in.UnknownExtendedKeyUsages, CertificateUnknownExtendedKeyUsages_ToProto)
	return out
}

func CertificateBaseKeyUsage_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage_KeyUsageOptions) *krm.CertificateBaseKeyUsage {
	if in == nil {
		return nil
	}
	out := &krm.CertificateBaseKeyUsage{}
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

func CertificateBaseKeyUsage_ToProto(mapCtx *direct.MapContext, in *krm.CertificateBaseKeyUsage) *pb.KeyUsage_KeyUsageOptions {
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

func CertificateExtendedKeyUsage_FromProto(mapCtx *direct.MapContext, in *pb.KeyUsage_ExtendedKeyUsageOptions) *krm.CertificateExtendedKeyUsage {
	if in == nil {
		return nil
	}
	out := &krm.CertificateExtendedKeyUsage{}
	out.ClientAuth = direct.LazyPtr(in.GetClientAuth())
	out.CodeSigning = direct.LazyPtr(in.GetCodeSigning())
	out.EmailProtection = direct.LazyPtr(in.GetEmailProtection())
	out.OcspSigning = direct.LazyPtr(in.GetOcspSigning())
	out.ServerAuth = direct.LazyPtr(in.GetServerAuth())
	out.TimeStamping = direct.LazyPtr(in.GetTimeStamping())
	return out
}

func CertificateExtendedKeyUsage_ToProto(mapCtx *direct.MapContext, in *krm.CertificateExtendedKeyUsage) *pb.KeyUsage_ExtendedKeyUsageOptions {
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

func int32To64Slice(in []int32) []int64 {
	if in == nil {
		return nil
	}
	out := make([]int64, len(in))
	for i, v := range in {
		out[i] = int64(v)
	}
	return out
}

func int64To32Slice(in []int64) []int32 {
	if in == nil {
		return nil
	}
	out := make([]int32, len(in))
	for i, v := range in {
		out[i] = int32(v)
	}
	return out
}
