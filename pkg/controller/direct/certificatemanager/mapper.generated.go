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
func Certificate_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) *krm.Certificate {
	if in == nil {
		return nil
	}
	out := &krm.Certificate{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.SelfManaged = Certificate_SelfManagedCertificate_FromProto(mapCtx, in.GetSelfManaged())
	out.Managed = Certificate_ManagedCertificate_FromProto(mapCtx, in.GetManaged())
	// MISSING: SanDnsnames
	// MISSING: PemCertificate
	// MISSING: ExpireTime
	out.Scope = direct.Enum_FromProto(mapCtx, in.GetScope())
	return out
}
func Certificate_ToProto(mapCtx *direct.MapContext, in *krm.Certificate) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	if oneof := Certificate_SelfManagedCertificate_ToProto(mapCtx, in.SelfManaged); oneof != nil {
		out.Type = &pb.Certificate_SelfManaged{SelfManaged: oneof}
	}
	if oneof := Certificate_ManagedCertificate_ToProto(mapCtx, in.Managed); oneof != nil {
		out.Type = &pb.Certificate_Managed{Managed: oneof}
	}
	// MISSING: SanDnsnames
	// MISSING: PemCertificate
	// MISSING: ExpireTime
	out.Scope = direct.Enum_ToProto[pb.Certificate_Scope](mapCtx, in.Scope)
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
func CertificateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) *krm.CertificateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: SelfManaged
	out.Managed = Certificate_ManagedCertificateObservedState_FromProto(mapCtx, in.GetManaged())
	out.SanDnsnames = in.SanDnsnames
	out.PemCertificate = direct.LazyPtr(in.GetPemCertificate())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: Scope
	return out
}
func CertificateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateObservedState) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: SelfManaged
	if oneof := Certificate_ManagedCertificateObservedState_ToProto(mapCtx, in.Managed); oneof != nil {
		out.Type = &pb.Certificate_Managed{Managed: oneof}
	}
	out.SanDnsnames = in.SanDnsnames
	out.PemCertificate = direct.ValueOf(in.PemCertificate)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: Scope
	return out
}
func Certificate_ManagedCertificate_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_ManagedCertificate) *krm.Certificate_ManagedCertificate {
	if in == nil {
		return nil
	}
	out := &krm.Certificate_ManagedCertificate{}
	out.Domains = in.Domains
	out.DnsAuthorizations = in.DnsAuthorizations
	out.IssuanceConfig = direct.LazyPtr(in.GetIssuanceConfig())
	// MISSING: State
	// MISSING: ProvisioningIssue
	// MISSING: AuthorizationAttemptInfo
	return out
}
func Certificate_ManagedCertificate_ToProto(mapCtx *direct.MapContext, in *krm.Certificate_ManagedCertificate) *pb.Certificate_ManagedCertificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_ManagedCertificate{}
	out.Domains = in.Domains
	out.DnsAuthorizations = in.DnsAuthorizations
	out.IssuanceConfig = direct.ValueOf(in.IssuanceConfig)
	// MISSING: State
	// MISSING: ProvisioningIssue
	// MISSING: AuthorizationAttemptInfo
	return out
}
func Certificate_ManagedCertificateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_ManagedCertificate) *krm.Certificate_ManagedCertificateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Certificate_ManagedCertificateObservedState{}
	// MISSING: Domains
	// MISSING: DnsAuthorizations
	// MISSING: IssuanceConfig
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ProvisioningIssue = Certificate_ManagedCertificate_ProvisioningIssue_FromProto(mapCtx, in.GetProvisioningIssue())
	out.AuthorizationAttemptInfo = direct.Slice_FromProto(mapCtx, in.AuthorizationAttemptInfo, Certificate_ManagedCertificate_AuthorizationAttemptInfo_FromProto)
	return out
}
func Certificate_ManagedCertificateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Certificate_ManagedCertificateObservedState) *pb.Certificate_ManagedCertificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_ManagedCertificate{}
	// MISSING: Domains
	// MISSING: DnsAuthorizations
	// MISSING: IssuanceConfig
	out.State = direct.Enum_ToProto[pb.Certificate_ManagedCertificate_State](mapCtx, in.State)
	out.ProvisioningIssue = Certificate_ManagedCertificate_ProvisioningIssue_ToProto(mapCtx, in.ProvisioningIssue)
	out.AuthorizationAttemptInfo = direct.Slice_ToProto(mapCtx, in.AuthorizationAttemptInfo, Certificate_ManagedCertificate_AuthorizationAttemptInfo_ToProto)
	return out
}
func Certificate_ManagedCertificate_AuthorizationAttemptInfo_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo) *krm.Certificate_ManagedCertificate_AuthorizationAttemptInfo {
	if in == nil {
		return nil
	}
	out := &krm.Certificate_ManagedCertificate_AuthorizationAttemptInfo{}
	out.Domain = direct.LazyPtr(in.GetDomain())
	// MISSING: State
	// MISSING: FailureReason
	// MISSING: Details
	return out
}
func Certificate_ManagedCertificate_AuthorizationAttemptInfo_ToProto(mapCtx *direct.MapContext, in *krm.Certificate_ManagedCertificate_AuthorizationAttemptInfo) *pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo{}
	out.Domain = direct.ValueOf(in.Domain)
	// MISSING: State
	// MISSING: FailureReason
	// MISSING: Details
	return out
}
func Certificate_ManagedCertificate_AuthorizationAttemptInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo) *krm.Certificate_ManagedCertificate_AuthorizationAttemptInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Certificate_ManagedCertificate_AuthorizationAttemptInfoObservedState{}
	// MISSING: Domain
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.FailureReason = direct.Enum_FromProto(mapCtx, in.GetFailureReason())
	out.Details = direct.LazyPtr(in.GetDetails())
	return out
}
func Certificate_ManagedCertificate_AuthorizationAttemptInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Certificate_ManagedCertificate_AuthorizationAttemptInfoObservedState) *pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo{}
	// MISSING: Domain
	out.State = direct.Enum_ToProto[pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo_State](mapCtx, in.State)
	out.FailureReason = direct.Enum_ToProto[pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo_FailureReason](mapCtx, in.FailureReason)
	out.Details = direct.ValueOf(in.Details)
	return out
}
func Certificate_ManagedCertificate_ProvisioningIssue_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_ManagedCertificate_ProvisioningIssue) *krm.Certificate_ManagedCertificate_ProvisioningIssue {
	if in == nil {
		return nil
	}
	out := &krm.Certificate_ManagedCertificate_ProvisioningIssue{}
	// MISSING: Reason
	// MISSING: Details
	return out
}
func Certificate_ManagedCertificate_ProvisioningIssue_ToProto(mapCtx *direct.MapContext, in *krm.Certificate_ManagedCertificate_ProvisioningIssue) *pb.Certificate_ManagedCertificate_ProvisioningIssue {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_ManagedCertificate_ProvisioningIssue{}
	// MISSING: Reason
	// MISSING: Details
	return out
}
func Certificate_ManagedCertificate_ProvisioningIssueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_ManagedCertificate_ProvisioningIssue) *krm.Certificate_ManagedCertificate_ProvisioningIssueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Certificate_ManagedCertificate_ProvisioningIssueObservedState{}
	out.Reason = direct.Enum_FromProto(mapCtx, in.GetReason())
	out.Details = direct.LazyPtr(in.GetDetails())
	return out
}
func Certificate_ManagedCertificate_ProvisioningIssueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Certificate_ManagedCertificate_ProvisioningIssueObservedState) *pb.Certificate_ManagedCertificate_ProvisioningIssue {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_ManagedCertificate_ProvisioningIssue{}
	out.Reason = direct.Enum_ToProto[pb.Certificate_ManagedCertificate_ProvisioningIssue_Reason](mapCtx, in.Reason)
	out.Details = direct.ValueOf(in.Details)
	return out
}
func Certificate_SelfManagedCertificate_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_SelfManagedCertificate) *krm.Certificate_SelfManagedCertificate {
	if in == nil {
		return nil
	}
	out := &krm.Certificate_SelfManagedCertificate{}
	out.PemCertificate = direct.LazyPtr(in.GetPemCertificate())
	out.PemPrivateKey = direct.LazyPtr(in.GetPemPrivateKey())
	return out
}
func Certificate_SelfManagedCertificate_ToProto(mapCtx *direct.MapContext, in *krm.Certificate_SelfManagedCertificate) *pb.Certificate_SelfManagedCertificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_SelfManagedCertificate{}
	out.PemCertificate = direct.ValueOf(in.PemCertificate)
	out.PemPrivateKey = direct.ValueOf(in.PemPrivateKey)
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
