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

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	krmkmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AlloyDbSettings_UserPassword_Password_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.Legacy) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.Value)
}

func AlloyDbSettings_UserPassword_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings_UserPassword) *krm.AlloyDbSettings_UserPassword {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings_UserPassword{}
	out.User = direct.LazyPtr(in.GetUser())
	return out
}

func CloudSQLSettings_RootPassword_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.Legacy) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.Value)
}

func CloudSQLSettings_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlSettings) *krm.CloudSQLSettings {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLSettings{}
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	out.UserLabels = in.UserLabels
	out.Tier = direct.LazyPtr(in.GetTier())
	out.StorageAutoResizeLimit = direct.Int64Value_FromProto(mapCtx, in.GetStorageAutoResizeLimit())
	out.ActivationPolicy = direct.Enum_FromProto(mapCtx, in.GetActivationPolicy())
	out.IPConfig = SQLIPConfig_FromProto(mapCtx, in.GetIpConfig())
	out.AutoStorageIncrease = direct.BoolValue_FromProto(mapCtx, in.GetAutoStorageIncrease())
	out.DatabaseFlags = in.DatabaseFlags
	out.DataDiskType = direct.Enum_FromProto(mapCtx, in.GetDataDiskType())
	out.DataDiskSizeGB = direct.Int64Value_FromProto(mapCtx, in.GetDataDiskSizeGb())
	out.Zone = direct.LazyPtr(in.GetZone())
	out.SecondaryZone = direct.LazyPtr(in.GetSecondaryZone())
	if in.GetSourceId() != "" {
		out.SourceIDRef = &krm.CloudDMSConnectionProfileRef{External: in.GetSourceId()}
	}
	out.Collation = direct.LazyPtr(in.GetCollation())
	if in.GetCmekKeyName() != "" {
		out.CmekKeyNameRef = &krmkmsv1beta1.KMSCryptoKeyRef{External: in.GetCmekKeyName()}
	}
	out.AvailabilityType = direct.Enum_FromProto(mapCtx, in.GetAvailabilityType())
	out.Edition = direct.Enum_FromProto(mapCtx, in.GetEdition())
	return out
}

func MySQLConnectionProfile_Password_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.Legacy) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.Value)
}

func MySQLConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.MySqlConnectionProfile) *krm.MySQLConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.MySQLConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.SSL = SSLConfig_FromProto(mapCtx, in.GetSsl())
	out.CloudSQLID = direct.LazyPtr(in.GetCloudSqlId())
	return out
}

func ForwardSSHTunnelConnectivity_Password_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.Legacy) *pb.ForwardSshTunnelConnectivity_Password {
	if in == nil {
		return nil
	}
	return &pb.ForwardSshTunnelConnectivity_Password{
		Password: direct.ValueOf(in.Value),
	}
}

func ForwardSSHTunnelConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.ForwardSshTunnelConnectivity) *krm.ForwardSSHTunnelConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.ForwardSSHTunnelConnectivity{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Port = direct.LazyPtr(in.GetPort())
	out.PrivateKey = direct.LazyPtr(in.GetPrivateKey())
	return out
}

func OracleConnectionProfile_Password_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.Legacy) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.Value)
}

func OracleConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.OracleConnectionProfile) *krm.OracleConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.OracleConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.DatabaseService = direct.LazyPtr(in.GetDatabaseService())
	out.SSL = SSLConfig_FromProto(mapCtx, in.GetSsl())
	out.StaticServiceIPConnectivity = StaticServiceIPConnectivity_FromProto(mapCtx, in.GetStaticServiceIpConnectivity())
	out.ForwardSSHConnectivity = ForwardSSHTunnelConnectivity_FromProto(mapCtx, in.GetForwardSshConnectivity())
	out.PrivateConnectivity = PrivateConnectivity_FromProto(mapCtx, in.GetPrivateConnectivity())
	return out
}

func PostgreSQLConnectionProfile_Password_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.Legacy) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.Value)
}

func PostgreSQLConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.PostgreSqlConnectionProfile) *krm.PostgreSQLConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.PostgreSQLConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.SSL = SSLConfig_FromProto(mapCtx, in.GetSsl())
	out.CloudSQLID = direct.LazyPtr(in.GetCloudSqlId())
	out.StaticIPConnectivity = StaticIPConnectivity_FromProto(mapCtx, in.GetStaticIpConnectivity())
	out.PrivateServiceConnectConnectivity = PrivateServiceConnectConnectivity_FromProto(mapCtx, in.GetPrivateServiceConnectConnectivity())
	return out
}
