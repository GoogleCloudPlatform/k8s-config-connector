// Copyright 2026 Google LLC
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
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetAppActiveDirectorySpec_Password_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.Legacy) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.Value)
}

func NetAppActiveDirectorySpec_Password_FromProto(mapCtx *direct.MapContext, in string) *refsv1beta1secret.Legacy {
	return nil
}

func NetAppActiveDirectorySpec_FromProto(mapCtx *direct.MapContext, in *pb.ActiveDirectory) *krm.NetAppActiveDirectorySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetAppActiveDirectorySpec{}
	// MISSING: Name
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.Site = direct.LazyPtr(in.GetSite())
	out.DNS = direct.LazyPtr(in.GetDns())
	out.NetBiosPrefix = direct.LazyPtr(in.GetNetBiosPrefix())
	out.OrganizationalUnit = direct.LazyPtr(in.GetOrganizationalUnit())
	out.AesEncryption = direct.LazyPtr(in.GetAesEncryption())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = NetAppActiveDirectorySpec_Password_FromProto(mapCtx, in.GetPassword())
	out.BackupOperators = in.BackupOperators
	out.Administrators = in.Administrators
	out.SecurityOperators = in.SecurityOperators
	out.KdcHostname = direct.LazyPtr(in.GetKdcHostname())
	out.KdcIP = direct.LazyPtr(in.GetKdcIp())
	out.NfsUsersWithLdap = direct.LazyPtr(in.GetNfsUsersWithLdap())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.LdapSigning = direct.LazyPtr(in.GetLdapSigning())
	out.EncryptDcConnections = direct.LazyPtr(in.GetEncryptDcConnections())
	out.Labels = in.Labels
	return out
}

func NetAppActiveDirectorySpec_ToProto(mapCtx *direct.MapContext, in *krm.NetAppActiveDirectorySpec) *pb.ActiveDirectory {
	if in == nil {
		return nil
	}
	out := &pb.ActiveDirectory{}
	// MISSING: Name
	out.Domain = direct.ValueOf(in.Domain)
	out.Site = direct.ValueOf(in.Site)
	out.Dns = direct.ValueOf(in.DNS)
	out.NetBiosPrefix = direct.ValueOf(in.NetBiosPrefix)
	out.OrganizationalUnit = direct.ValueOf(in.OrganizationalUnit)
	out.AesEncryption = direct.ValueOf(in.AesEncryption)
	out.Username = direct.ValueOf(in.Username)
	out.Password = NetAppActiveDirectorySpec_Password_ToProto(mapCtx, in.Password)
	out.BackupOperators = in.BackupOperators
	out.Administrators = in.Administrators
	out.SecurityOperators = in.SecurityOperators
	out.KdcHostname = direct.ValueOf(in.KdcHostname)
	out.KdcIp = direct.ValueOf(in.KdcIP)
	out.NfsUsersWithLdap = direct.ValueOf(in.NfsUsersWithLdap)
	out.Description = direct.ValueOf(in.Description)
	out.LdapSigning = direct.ValueOf(in.LdapSigning)
	out.EncryptDcConnections = direct.ValueOf(in.EncryptDcConnections)
	out.Labels = in.Labels
	return out
}
