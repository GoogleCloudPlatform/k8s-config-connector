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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/v1beta1"
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
