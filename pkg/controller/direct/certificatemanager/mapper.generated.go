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

package certificatemanager

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
)
func CertificateIssuanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig) *krm.CertificateIssuanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.CertificateIssuanceConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CertificateAuthorityConfig = CertificateIssuanceConfig_CertificateAuthorityConfig_FromProto(mapCtx, in.GetCertificateAuthorityConfig())
	out.Lifetime = direct.StringDuration_FromProto(mapCtx, in.GetLifetime())
	out.RotationWindowPercentage = direct.LazyPtr(in.GetRotationWindowPercentage())
	out.KeyAlgorithm = direct.Enum_FromProto(mapCtx, in.GetKeyAlgorithm())
	return out
}
func CertificateIssuanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.CertificateIssuanceConfig) *pb.CertificateIssuanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.CertificateAuthorityConfig = CertificateIssuanceConfig_CertificateAuthorityConfig_ToProto(mapCtx, in.CertificateAuthorityConfig)
	out.Lifetime = direct.StringDuration_ToProto(mapCtx, in.Lifetime)
	out.RotationWindowPercentage = direct.ValueOf(in.RotationWindowPercentage)
	out.KeyAlgorithm = direct.Enum_ToProto[pb.CertificateIssuanceConfig_KeyAlgorithm](mapCtx, in.KeyAlgorithm)
	return out
}
func CertificateIssuanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig) *krm.CertificateIssuanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateIssuanceConfigObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: CertificateAuthorityConfig
	// MISSING: Lifetime
	// MISSING: RotationWindowPercentage
	// MISSING: KeyAlgorithm
	return out
}
func CertificateIssuanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateIssuanceConfigObservedState) *pb.CertificateIssuanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: CertificateAuthorityConfig
	// MISSING: Lifetime
	// MISSING: RotationWindowPercentage
	// MISSING: KeyAlgorithm
	return out
}
func CertificateIssuanceConfig_CertificateAuthorityConfig_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig_CertificateAuthorityConfig) *krm.CertificateIssuanceConfig_CertificateAuthorityConfig {
	if in == nil {
		return nil
	}
	out := &krm.CertificateIssuanceConfig_CertificateAuthorityConfig{}
	out.CertificateAuthorityServiceConfig = CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_FromProto(mapCtx, in.GetCertificateAuthorityServiceConfig())
	return out
}
func CertificateIssuanceConfig_CertificateAuthorityConfig_ToProto(mapCtx *direct.MapContext, in *krm.CertificateIssuanceConfig_CertificateAuthorityConfig) *pb.CertificateIssuanceConfig_CertificateAuthorityConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig_CertificateAuthorityConfig{}
	if oneof := CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_ToProto(mapCtx, in.CertificateAuthorityServiceConfig); oneof != nil {
		out.Kind = &pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_{CertificateAuthorityServiceConfig: oneof}
	}
	return out
}
func CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig) *krm.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig {
	if in == nil {
		return nil
	}
	out := &krm.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig{}
	out.CaPool = direct.LazyPtr(in.GetCaPool())
	return out
}
func CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_ToProto(mapCtx *direct.MapContext, in *krm.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig) *pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig{}
	out.CaPool = direct.ValueOf(in.CaPool)
	return out
}
func CertificateManagerDNSAuthorizationSpec_FromProto(mapCtx *direct.MapContext, in *pb.DnsAuthorization) *krm.CertificateManagerDNSAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManagerDNSAuthorizationSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Domain = direct.LazyPtr(in.GetDomain())
	// MISSING: DnsResourceRecord
	// MISSING: Type
	return out
}
func CertificateManagerDNSAuthorizationSpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManagerDNSAuthorizationSpec) *pb.DnsAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.DnsAuthorization{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.Description = direct.ValueOf(in.Description)
	out.Domain = CertificateManagerDNSAuthorizationSpec_Domain_ToProto(mapCtx, in.Domain)
	// MISSING: DnsResourceRecord
	// MISSING: Type
	return out
}
func CertificatemanagerCertificateIssuanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig) *krm.CertificatemanagerCertificateIssuanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificatemanagerCertificateIssuanceConfigObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: CertificateAuthorityConfig
	// MISSING: Lifetime
	// MISSING: RotationWindowPercentage
	// MISSING: KeyAlgorithm
	return out
}
func CertificatemanagerCertificateIssuanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificatemanagerCertificateIssuanceConfigObservedState) *pb.CertificateIssuanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: CertificateAuthorityConfig
	// MISSING: Lifetime
	// MISSING: RotationWindowPercentage
	// MISSING: KeyAlgorithm
	return out
}
func CertificatemanagerCertificateIssuanceConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig) *krm.CertificatemanagerCertificateIssuanceConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificatemanagerCertificateIssuanceConfigSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: CertificateAuthorityConfig
	// MISSING: Lifetime
	// MISSING: RotationWindowPercentage
	// MISSING: KeyAlgorithm
	return out
}
func CertificatemanagerCertificateIssuanceConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificatemanagerCertificateIssuanceConfigSpec) *pb.CertificateIssuanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: CertificateAuthorityConfig
	// MISSING: Lifetime
	// MISSING: RotationWindowPercentage
	// MISSING: KeyAlgorithm
	return out
}
func DnsAuthorization_DnsResourceRecord_FromProto(mapCtx *direct.MapContext, in *pb.DnsAuthorization_DnsResourceRecord) *krm.DnsAuthorization_DnsResourceRecord {
	if in == nil {
		return nil
	}
	out := &krm.DnsAuthorization_DnsResourceRecord{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	out.Data = direct.LazyPtr(in.GetData())
	return out
}
func DnsAuthorization_DnsResourceRecord_ToProto(mapCtx *direct.MapContext, in *krm.DnsAuthorization_DnsResourceRecord) *pb.DnsAuthorization_DnsResourceRecord {
	if in == nil {
		return nil
	}
	out := &pb.DnsAuthorization_DnsResourceRecord{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	out.Data = direct.ValueOf(in.Data)
	return out
}
