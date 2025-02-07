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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
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
func CertificateMap_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMap) *krm.CertificateMap {
	if in == nil {
		return nil
	}
	out := &krm.CertificateMap{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: GclbTargets
	return out
}
func CertificateMap_ToProto(mapCtx *direct.MapContext, in *krm.CertificateMap) *pb.CertificateMap {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMap{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: GclbTargets
	return out
}
func CertificateMapObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMap) *krm.CertificateMapObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateMapObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.GclbTargets = direct.Slice_FromProto(mapCtx, in.GclbTargets, CertificateMap_GclbTarget_FromProto)
	return out
}
func CertificateMapObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateMapObservedState) *pb.CertificateMap {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMap{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.GclbTargets = direct.Slice_ToProto(mapCtx, in.GclbTargets, CertificateMap_GclbTarget_ToProto)
	return out
}
func CertificateMap_GclbTarget_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMap_GclbTarget) *krm.CertificateMap_GclbTarget {
	if in == nil {
		return nil
	}
	out := &krm.CertificateMap_GclbTarget{}
	// MISSING: TargetHTTPSProxy
	// MISSING: TargetSslProxy
	// MISSING: IPConfigs
	return out
}
func CertificateMap_GclbTarget_ToProto(mapCtx *direct.MapContext, in *krm.CertificateMap_GclbTarget) *pb.CertificateMap_GclbTarget {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMap_GclbTarget{}
	// MISSING: TargetHTTPSProxy
	// MISSING: TargetSslProxy
	// MISSING: IPConfigs
	return out
}
func CertificateMap_GclbTargetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMap_GclbTarget) *krm.CertificateMap_GclbTargetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateMap_GclbTargetObservedState{}
	out.TargetHTTPSProxy = direct.LazyPtr(in.GetTargetHttpsProxy())
	out.TargetSslProxy = direct.LazyPtr(in.GetTargetSslProxy())
	out.IPConfigs = direct.Slice_FromProto(mapCtx, in.IPConfigs, CertificateMap_GclbTarget_IpConfig_FromProto)
	return out
}
func CertificateMap_GclbTargetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateMap_GclbTargetObservedState) *pb.CertificateMap_GclbTarget {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMap_GclbTarget{}
	if oneof := CertificateMap_GclbTargetObservedState_TargetHttpsProxy_ToProto(mapCtx, in.TargetHTTPSProxy); oneof != nil {
		out.TargetProxy = oneof
	}
	if oneof := CertificateMap_GclbTargetObservedState_TargetSslProxy_ToProto(mapCtx, in.TargetSslProxy); oneof != nil {
		out.TargetProxy = oneof
	}
	out.IpConfigs = direct.Slice_ToProto(mapCtx, in.IPConfigs, CertificateMap_GclbTarget_IpConfig_ToProto)
	return out
}
func CertificateMap_GclbTarget_IpConfig_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMap_GclbTarget_IpConfig) *krm.CertificateMap_GclbTarget_IpConfig {
	if in == nil {
		return nil
	}
	out := &krm.CertificateMap_GclbTarget_IpConfig{}
	// MISSING: IPAddress
	// MISSING: Ports
	return out
}
func CertificateMap_GclbTarget_IpConfig_ToProto(mapCtx *direct.MapContext, in *krm.CertificateMap_GclbTarget_IpConfig) *pb.CertificateMap_GclbTarget_IpConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMap_GclbTarget_IpConfig{}
	// MISSING: IPAddress
	// MISSING: Ports
	return out
}
func CertificateMap_GclbTarget_IpConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMap_GclbTarget_IpConfig) *krm.CertificateMap_GclbTarget_IpConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateMap_GclbTarget_IpConfigObservedState{}
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.Ports = in.Ports
	return out
}
func CertificateMap_GclbTarget_IpConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateMap_GclbTarget_IpConfigObservedState) *pb.CertificateMap_GclbTarget_IpConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMap_GclbTarget_IpConfig{}
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.Ports = in.Ports
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
func CertificatemanagerCertificateMapObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMap) *krm.CertificatemanagerCertificateMapObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificatemanagerCertificateMapObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: GclbTargets
	return out
}
func CertificatemanagerCertificateMapObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificatemanagerCertificateMapObservedState) *pb.CertificateMap {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMap{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: GclbTargets
	return out
}
func CertificatemanagerCertificateMapSpec_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMap) *krm.CertificatemanagerCertificateMapSpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificatemanagerCertificateMapSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: GclbTargets
	return out
}
func CertificatemanagerCertificateMapSpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificatemanagerCertificateMapSpec) *pb.CertificateMap {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMap{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: GclbTargets
	return out
}
func CertificatemanagerCertificateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) *krm.CertificatemanagerCertificateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificatemanagerCertificateObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: SelfManaged
	// MISSING: Managed
	// MISSING: SanDnsnames
	// MISSING: PemCertificate
	// MISSING: ExpireTime
	// MISSING: Scope
	return out
}
func CertificatemanagerCertificateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificatemanagerCertificateObservedState) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: SelfManaged
	// MISSING: Managed
	// MISSING: SanDnsnames
	// MISSING: PemCertificate
	// MISSING: ExpireTime
	// MISSING: Scope
	return out
}
func CertificatemanagerCertificateSpec_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) *krm.CertificatemanagerCertificateSpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificatemanagerCertificateSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: SelfManaged
	// MISSING: Managed
	// MISSING: SanDnsnames
	// MISSING: PemCertificate
	// MISSING: ExpireTime
	// MISSING: Scope
	return out
}
func CertificatemanagerCertificateSpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificatemanagerCertificateSpec) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: SelfManaged
	// MISSING: Managed
	// MISSING: SanDnsnames
	// MISSING: PemCertificate
	// MISSING: ExpireTime
	// MISSING: Scope
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
