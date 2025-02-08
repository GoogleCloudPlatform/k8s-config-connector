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
func ReusableConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReusableConfig) *krm.ReusableConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReusableConfig{}
	// MISSING: Name
	out.Values = ReusableConfigValues_FromProto(mapCtx, in.GetValues())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	return out
}
func ReusableConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReusableConfig) *pb.ReusableConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReusableConfig{}
	// MISSING: Name
	out.Values = ReusableConfigValues_ToProto(mapCtx, in.Values)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	return out
}
func ReusableConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReusableConfig) *krm.ReusableConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReusableConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Values
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	return out
}
func ReusableConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReusableConfigObservedState) *pb.ReusableConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReusableConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Values
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
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
func SecurityReusableConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReusableConfig) *krm.SecurityReusableConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecurityReusableConfigObservedState{}
	// MISSING: Name
	// MISSING: Values
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	return out
}
func SecurityReusableConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecurityReusableConfigObservedState) *pb.ReusableConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReusableConfig{}
	// MISSING: Name
	// MISSING: Values
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	return out
}
func SecurityReusableConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReusableConfig) *krm.SecurityReusableConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecurityReusableConfigSpec{}
	// MISSING: Name
	// MISSING: Values
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	return out
}
func SecurityReusableConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecurityReusableConfigSpec) *pb.ReusableConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReusableConfig{}
	// MISSING: Name
	// MISSING: Values
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
