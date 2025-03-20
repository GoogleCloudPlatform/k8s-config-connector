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

package netapp

import (
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetAppBackupPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPolicy) *krm.NetAppBackupPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetAppBackupPolicySpec{}
	out.DailyBackupLimit = in.DailyBackupLimit
	out.WeeklyBackupLimit = in.WeeklyBackupLimit
	out.MonthlyBackupLimit = in.MonthlyBackupLimit
	out.Description = in.Description
	out.Enabled = in.Enabled

	// out.Labels = in.Labels // NOT YET

	return out
}
func NetAppBackupPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.NetAppBackupPolicySpec) *pb.BackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackupPolicy{}
	out.DailyBackupLimit = in.DailyBackupLimit
	out.WeeklyBackupLimit = in.WeeklyBackupLimit
	out.MonthlyBackupLimit = in.MonthlyBackupLimit
	out.Description = in.Description
	out.Enabled = in.Enabled
	// out.Labels = in.Labels // NOT YET
	return out
}
func NetAppBackupPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPolicy) *krm.NetAppBackupPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetAppBackupPolicyObservedState{}

	out.AssignedVolumeCount = in.AssignedVolumeCount
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func NetAppBackupPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetAppBackupPolicyObservedState) *pb.BackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackupPolicy{}

	out.AssignedVolumeCount = in.AssignedVolumeCount
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[pb.BackupPolicy_State](mapCtx, in.State)

	return out
}

func ActiveDirectorySpec_FromProto(mapCtx *direct.MapContext, in *pb.ActiveDirectory) *krm.ActiveDirectorySpec {
	if in == nil {
		return nil
	}
	out := &krm.ActiveDirectorySpec{}
	// out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: State
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.Site = direct.LazyPtr(in.GetSite())
	out.DNS = direct.LazyPtr(in.GetDns())
	out.NetBiosPrefix = direct.LazyPtr(in.GetNetBiosPrefix())
	out.OrganizationalUnit = direct.LazyPtr(in.GetOrganizationalUnit())
	out.AesEncryption = direct.LazyPtr(in.GetAesEncryption())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.BackupOperators = in.BackupOperators
	out.Administrators = in.Administrators
	out.SecurityOperators = in.SecurityOperators
	out.KdcHostname = direct.LazyPtr(in.GetKdcHostname())
	out.KdcIP = direct.LazyPtr(in.GetKdcIp())
	out.NfsUsersWithLdap = direct.LazyPtr(in.GetNfsUsersWithLdap())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.LdapSigning = direct.LazyPtr(in.GetLdapSigning())
	out.EncryptDcConnections = direct.LazyPtr(in.GetEncryptDcConnections())
	// out.Labels = in.Labels
	// MISSING: StateDetails
	return out
}
func ActiveDirectorySpec_ToProto(mapCtx *direct.MapContext, in *krm.ActiveDirectorySpec) *pb.ActiveDirectory {
	if in == nil {
		return nil
	}
	out := &pb.ActiveDirectory{}
	// MISSING: CreateTime
	// MISSING: State
	out.Domain = direct.ValueOf(in.Domain)
	out.Site = direct.ValueOf(in.Site)
	out.Dns = direct.ValueOf(in.DNS)
	out.NetBiosPrefix = direct.ValueOf(in.NetBiosPrefix)
	out.OrganizationalUnit = direct.ValueOf(in.OrganizationalUnit)
	out.AesEncryption = direct.ValueOf(in.AesEncryption)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.BackupOperators = in.BackupOperators
	out.Administrators = in.Administrators
	out.SecurityOperators = in.SecurityOperators
	out.KdcHostname = direct.ValueOf(in.KdcHostname)
	out.KdcIp = direct.ValueOf(in.KdcIP)
	out.NfsUsersWithLdap = direct.ValueOf(in.NfsUsersWithLdap)
	out.Description = direct.ValueOf(in.Description)
	out.LdapSigning = direct.ValueOf(in.LdapSigning)
	out.EncryptDcConnections = direct.ValueOf(in.EncryptDcConnections)
	//out.Labels = in.Labels
	// MISSING: StateDetails
	return out
}
func ActiveDirectoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ActiveDirectory) *krm.ActiveDirectoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ActiveDirectoryObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Domain
	// MISSING: Site
	// MISSING: DNS
	// MISSING: NetBiosPrefix
	// MISSING: OrganizationalUnit
	// MISSING: AesEncryption
	// MISSING: Username
	// MISSING: Password
	// MISSING: BackupOperators
	// MISSING: Administrators
	// MISSING: SecurityOperators
	// MISSING: KdcHostname
	// MISSING: KdcIP
	// MISSING: NfsUsersWithLdap
	// MISSING: Description
	// MISSING: LdapSigning
	// MISSING: EncryptDcConnections
	// MISSING: Labels
	out.StateDetails = direct.LazyPtr(in.GetStateDetails())
	return out
}
func ActiveDirectoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ActiveDirectoryObservedState) *pb.ActiveDirectory {
	if in == nil {
		return nil
	}
	out := &pb.ActiveDirectory{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[pb.ActiveDirectory_State](mapCtx, in.State)
	// MISSING: Domain
	// MISSING: Site
	// MISSING: DNS
	// MISSING: NetBiosPrefix
	// MISSING: OrganizationalUnit
	// MISSING: AesEncryption
	// MISSING: Username
	// MISSING: Password
	// MISSING: BackupOperators
	// MISSING: Administrators
	// MISSING: SecurityOperators
	// MISSING: KdcHostname
	// MISSING: KdcIP
	// MISSING: NfsUsersWithLdap
	// MISSING: Description
	// MISSING: LdapSigning
	// MISSING: EncryptDcConnections
	// MISSING: Labels
	out.StateDetails = direct.ValueOf(in.StateDetails)
	return out
}
