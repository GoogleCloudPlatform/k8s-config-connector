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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
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
func CertificateMapEntry_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMapEntry) *krm.CertificateMapEntry {
	if in == nil {
		return nil
	}
	out := &krm.CertificateMapEntry{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Matcher = direct.Enum_FromProto(mapCtx, in.GetMatcher())
	out.Certificates = in.Certificates
	// MISSING: State
	return out
}
func CertificateMapEntry_ToProto(mapCtx *direct.MapContext, in *krm.CertificateMapEntry) *pb.CertificateMapEntry {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMapEntry{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	if oneof := CertificateMapEntry_Hostname_ToProto(mapCtx, in.Hostname); oneof != nil {
		out.Match = oneof
	}
	if oneof := CertificateMapEntry_Matcher_ToProto(mapCtx, in.Matcher); oneof != nil {
		out.Match = oneof
	}
	out.Certificates = in.Certificates
	// MISSING: State
	return out
}
func CertificateMapEntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMapEntry) *krm.CertificateMapEntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateMapEntryObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Hostname
	// MISSING: Matcher
	// MISSING: Certificates
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func CertificateMapEntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateMapEntryObservedState) *pb.CertificateMapEntry {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMapEntry{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Hostname
	// MISSING: Matcher
	// MISSING: Certificates
	out.State = direct.Enum_ToProto[pb.ServingState](mapCtx, in.State)
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
func CertificatemanagerCertificateMapEntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMapEntry) *krm.CertificatemanagerCertificateMapEntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificatemanagerCertificateMapEntryObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Hostname
	// MISSING: Matcher
	// MISSING: Certificates
	// MISSING: State
	return out
}
func CertificatemanagerCertificateMapEntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificatemanagerCertificateMapEntryObservedState) *pb.CertificateMapEntry {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMapEntry{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Hostname
	// MISSING: Matcher
	// MISSING: Certificates
	// MISSING: State
	return out
}
func CertificatemanagerCertificateMapEntrySpec_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMapEntry) *krm.CertificatemanagerCertificateMapEntrySpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificatemanagerCertificateMapEntrySpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Hostname
	// MISSING: Matcher
	// MISSING: Certificates
	// MISSING: State
	return out
}
func CertificatemanagerCertificateMapEntrySpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificatemanagerCertificateMapEntrySpec) *pb.CertificateMapEntry {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMapEntry{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Hostname
	// MISSING: Matcher
	// MISSING: Certificates
	// MISSING: State
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
