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

// +generated:mapper
// krm.group: certificatemanager.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.certificatemanager.v1

package certificatemanager

import (
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krmcertificatemanagerv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CertificateIssuanceConfig_CertificateAuthorityConfig_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig_CertificateAuthorityConfig) *krmcertificatemanagerv1alpha1.CertificateIssuanceConfig_CertificateAuthorityConfig {
	if in == nil {
		return nil
	}
	out := &krmcertificatemanagerv1alpha1.CertificateIssuanceConfig_CertificateAuthorityConfig{}
	out.CertificateAuthorityServiceConfig = CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_FromProto(mapCtx, in.GetCertificateAuthorityServiceConfig())
	return out
}
func CertificateIssuanceConfig_CertificateAuthorityConfig_ToProto(mapCtx *direct.MapContext, in *krmcertificatemanagerv1alpha1.CertificateIssuanceConfig_CertificateAuthorityConfig) *pb.CertificateIssuanceConfig_CertificateAuthorityConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig_CertificateAuthorityConfig{}
	if oneof := CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_ToProto(mapCtx, in.CertificateAuthorityServiceConfig); oneof != nil {
		out.Kind = &pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_{CertificateAuthorityServiceConfig: oneof}
	}
	return out
}
func CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig) *krmcertificatemanagerv1alpha1.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig {
	if in == nil {
		return nil
	}
	out := &krmcertificatemanagerv1alpha1.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig{}
	if in.GetCaPool() != "" {
		out.CAPoolRef = &refsv1beta1.PrivateCACAPoolRef{External: in.GetCaPool()}
	}
	return out
}
func CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_ToProto(mapCtx *direct.MapContext, in *krmcertificatemanagerv1alpha1.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig) *pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig{}
	if in.CAPoolRef != nil {
		out.CaPool = in.CAPoolRef.External
	}
	return out
}
func CertificateManagerCertificateIssuanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig) *krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	return out
}
func CertificateManagerCertificateIssuanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigObservedState) *pb.CertificateIssuanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	return out
}
func CertificateManagerCertificateIssuanceConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig) *krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigSpec {
	if in == nil {
		return nil
	}
	out := &krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigSpec{}
	// MISSING: Name
	// MISSING: Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CertificateAuthorityConfig = CertificateIssuanceConfig_CertificateAuthorityConfig_FromProto(mapCtx, in.GetCertificateAuthorityConfig())
	out.Lifetime = direct.StringDuration_FromProto(mapCtx, in.GetLifetime())
	out.RotationWindowPercentage = direct.LazyPtr(in.GetRotationWindowPercentage())
	out.KeyAlgorithm = direct.Enum_FromProto(mapCtx, in.GetKeyAlgorithm())
	return out
}
func CertificateManagerCertificateIssuanceConfigSpec_ToProto(mapCtx *direct.MapContext, in *krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigSpec) *pb.CertificateIssuanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig{}
	// MISSING: Name
	// MISSING: Labels
	out.Description = direct.ValueOf(in.Description)
	out.CertificateAuthorityConfig = CertificateIssuanceConfig_CertificateAuthorityConfig_ToProto(mapCtx, in.CertificateAuthorityConfig)
	out.Lifetime = direct.StringDuration_ToProto(mapCtx, in.Lifetime)
	out.RotationWindowPercentage = direct.ValueOf(in.RotationWindowPercentage)
	out.KeyAlgorithm = direct.Enum_ToProto[pb.CertificateIssuanceConfig_KeyAlgorithm](mapCtx, in.KeyAlgorithm)
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
	// MISSING: DNSResourceRecord
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
	out.Domain = direct.ValueOf(in.Domain)
	// MISSING: DNSResourceRecord
	// MISSING: Type
	return out
}
func CertificateManagerDNSAuthorizationStatus_FromProto(mapCtx *direct.MapContext, in *pb.DnsAuthorization) *krm.CertificateManagerDNSAuthorizationStatus {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManagerDNSAuthorizationStatus{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Domain
	if v := in.GetDnsResourceRecord(); v != nil {
		out.DNSResourceRecord = []*krm.DNSAuthorization_DNSResourceRecordObservedState{DNSAuthorization_DNSResourceRecordObservedState_FromProto(mapCtx, v)}
	}
	// MISSING: Type
	return out
}
func CertificateManagerDNSAuthorizationStatus_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManagerDNSAuthorizationStatus) *pb.DnsAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.DnsAuthorization{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Domain
	if len(in.DNSResourceRecord) > 0 && in.DNSResourceRecord[0] != nil {
		out.DnsResourceRecord = DNSAuthorization_DNSResourceRecordObservedState_ToProto(mapCtx, in.DNSResourceRecord[0])
	}
	// MISSING: Type
	return out
}
func DNSAuthorization_DNSResourceRecord_FromProto(mapCtx *direct.MapContext, in *pb.DnsAuthorization_DnsResourceRecord) *krm.DNSAuthorization_DNSResourceRecord {
	if in == nil {
		return nil
	}
	out := &krm.DNSAuthorization_DNSResourceRecord{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: Data
	return out
}
func DNSAuthorization_DNSResourceRecord_ToProto(mapCtx *direct.MapContext, in *krm.DNSAuthorization_DNSResourceRecord) *pb.DnsAuthorization_DnsResourceRecord {
	if in == nil {
		return nil
	}
	out := &pb.DnsAuthorization_DnsResourceRecord{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: Data
	return out
}
func DNSAuthorization_DNSResourceRecordObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DnsAuthorization_DnsResourceRecord) *krm.DNSAuthorization_DNSResourceRecordObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DNSAuthorization_DNSResourceRecordObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	out.Data = direct.LazyPtr(in.GetData())
	return out
}
func DNSAuthorization_DNSResourceRecordObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DNSAuthorization_DNSResourceRecordObservedState) *pb.DnsAuthorization_DnsResourceRecord {
	if in == nil {
		return nil
	}
	out := &pb.DnsAuthorization_DnsResourceRecord{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	out.Data = direct.ValueOf(in.Data)
	return out
}
