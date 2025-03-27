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

package datastream

import (
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DatastreamConnectionProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.DatastreamConnectionProfileSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamConnectionProfileSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.OracleProfile = OracleProfile_FromProto(mapCtx, in.GetOracleProfile())
	out.GCSProfile = GCSProfile_FromProto(mapCtx, in.GetGcsProfile())
	out.MySQLProfile = MysqlProfile_FromProto(mapCtx, in.GetMysqlProfile())
	out.BigQueryProfile = BigQueryProfile_FromProto(mapCtx, in.GetBigqueryProfile())
	out.SQLServerProfile = SQLServerProfile_FromProto(mapCtx, in.GetSqlServerProfile())
	out.StaticServiceIPConnectivity = StaticServiceIPConnectivity_FromProto(mapCtx, in.GetStaticServiceIpConnectivity())
	out.ForwardSSHConnectivity = ForwardSSHTunnelConnectivity_FromProto(mapCtx, in.GetForwardSshConnectivity())
	out.PrivateConnectivity = PrivateConnectivity_FromProto(mapCtx, in.GetPrivateConnectivity())
	return out
}
func DatastreamConnectionProfileSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamConnectionProfileSpec) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := OracleProfile_ToProto(mapCtx, in.OracleProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_OracleProfile{OracleProfile: oneof}
	}
	if oneof := GCSProfile_ToProto(mapCtx, in.GCSProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_GcsProfile{GcsProfile: oneof}
	}
	if oneof := MysqlProfile_ToProto(mapCtx, in.MySQLProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_MysqlProfile{MysqlProfile: oneof}
	}
	if oneof := BigQueryProfile_ToProto(mapCtx, in.BigQueryProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_BigqueryProfile{BigqueryProfile: oneof}
	}
	if oneof := SQLServerProfile_ToProto(mapCtx, in.SQLServerProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_SqlServerProfile{SqlServerProfile: oneof}
	}
	if oneof := StaticServiceIPConnectivity_ToProto(mapCtx, in.StaticServiceIPConnectivity); oneof != nil {
		out.Connectivity = &pb.ConnectionProfile_StaticServiceIpConnectivity{StaticServiceIpConnectivity: oneof}
	}
	if oneof := ForwardSSHTunnelConnectivity_ToProto(mapCtx, in.ForwardSSHConnectivity); oneof != nil {
		out.Connectivity = &pb.ConnectionProfile_ForwardSshConnectivity{ForwardSshConnectivity: oneof}
	}
	if oneof := PrivateConnectivity_ToProto(mapCtx, in.PrivateConnectivity); oneof != nil {
		out.Connectivity = &pb.ConnectionProfile_PrivateConnectivity{PrivateConnectivity: oneof}
	}
	return out
}
func DatastreamConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.DatastreamConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamConnectionProfileObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.OracleProfile = OracleProfileObservedState_FromProto(mapCtx, in.GetOracleProfile())
	out.MySQLProfile = MysqlProfileObservedState_FromProto(mapCtx, in.GetMysqlProfile())
	// MISSING: BigqueryProfile
	// MISSING: PostgresqlProfile
	return out
}
func DatastreamConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamConnectionProfileObservedState) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	if oneof := OracleProfileObservedState_ToProto(mapCtx, in.OracleProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_OracleProfile{OracleProfile: oneof}
	}
	if oneof := MysqlProfileObservedState_ToProto(mapCtx, in.MySQLProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_MysqlProfile{MysqlProfile: oneof}
	}
	// MISSING: BigqueryProfile
	// MISSING: PostgresqlProfile
	return out
}
func OracleAsmConfig_FromProto(mapCtx *direct.MapContext, in *pb.OracleAsmConfig) *krm.OracleAsmConfig {
	if in == nil {
		return nil
	}
	out := &krm.OracleAsmConfig{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	// out.SecretRef is sensitive data, so we don't output it.
	out.ASMService = direct.LazyPtr(in.GetAsmService())
	out.ConnectionAttributes = in.ConnectionAttributes
	out.OracleSSLConfig = OracleSSLConfig_FromProto(mapCtx, in.GetOracleSslConfig())
	return out
}
func OracleAsmConfig_ToProto(mapCtx *direct.MapContext, in *krm.OracleAsmConfig) *pb.OracleAsmConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleAsmConfig{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	if in.SecretRef != nil {
		out.Username = in.SecretRef.Username
		out.Password = in.SecretRef.Password
	}
	out.AsmService = direct.ValueOf(in.ASMService)
	out.ConnectionAttributes = in.ConnectionAttributes
	out.OracleSslConfig = OracleSSLConfig_ToProto(mapCtx, in.OracleSSLConfig)
	return out
}
func OracleProfile_FromProto(mapCtx *direct.MapContext, in *pb.OracleProfile) *krm.OracleProfile {
	if in == nil {
		return nil
	}
	out := &krm.OracleProfile{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	// out.SecretRef is sensitive data, so we don't output it.
	out.DatabaseService = direct.LazyPtr(in.GetDatabaseService())
	out.ConnectionAttributes = in.ConnectionAttributes
	out.OracleSSLConfig = OracleSSLConfig_FromProto(mapCtx, in.GetOracleSslConfig())
	out.OracleASMConfig = OracleAsmConfig_FromProto(mapCtx, in.GetOracleAsmConfig())
	if in.GetSecretManagerStoredPassword() != "" {
		out.SecreteManagerSecretRef = &refsv1beta1.SecretManagerSecretRef{
			External: in.GetSecretManagerStoredPassword(),
		}
	}
	return out
}
func OracleProfile_ToProto(mapCtx *direct.MapContext, in *krm.OracleProfile) *pb.OracleProfile {
	if in == nil {
		return nil
	}
	out := &pb.OracleProfile{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	if in.SecretRef != nil {
		out.Username = in.SecretRef.Username
		out.Password = in.SecretRef.Password
	}
	out.DatabaseService = direct.ValueOf(in.DatabaseService)
	out.ConnectionAttributes = in.ConnectionAttributes
	out.OracleSslConfig = OracleSSLConfig_ToProto(mapCtx, in.OracleSSLConfig)
	out.OracleAsmConfig = OracleAsmConfig_ToProto(mapCtx, in.OracleASMConfig)
	if in.SecreteManagerSecretRef != nil {
		out.SecretManagerStoredPassword = in.SecreteManagerSecretRef.External
	}
	return out
}
func PrivateConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnectivity) *krm.PrivateConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.PrivateConnectivity{}
	if in.GetPrivateConnection() != "" {
		out.PrivateConnectionRef = &krm.PrivateConnectionRef{External: in.GetPrivateConnection()}
	}
	return out
}
func ForwardSSHTunnelConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.ForwardSshTunnelConnectivity) *krm.ForwardSSHTunnelConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.ForwardSSHTunnelConnectivity{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	// out.SecretRef is sensitive data, so we don't output it.
	out.PrivateKey = direct.LazyPtr(in.GetPrivateKey())
	return out
}
func ForwardSSHTunnelConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.ForwardSSHTunnelConnectivity) *pb.ForwardSshTunnelConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.ForwardSshTunnelConnectivity{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	if in.SecretRef != nil {
		out.Username = in.SecretRef.Username
	}
	if oneof := ForwardSSHTunnelConnectivity_Password_ToProto(mapCtx, in.SecretRef); oneof != nil {
		out.AuthenticationMethod = oneof
	}
	if oneof := ForwardSSHTunnelConnectivity_PrivateKey_ToProto(mapCtx, in.PrivateKey); oneof != nil {
		out.AuthenticationMethod = oneof
	}
	return out
}
func ForwardSSHTunnelConnectivity_Password_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.BasicAuthSecretRef) *pb.ForwardSshTunnelConnectivity_Password {
	if in == nil {
		return nil
	}
	if in.Password == "" {
		return nil
	}
	out := &pb.ForwardSshTunnelConnectivity_Password{}
	out.Password = in.Password
	return out
}
func ForwardSSHTunnelConnectivity_PrivateKey_ToProto(mapCtx *direct.MapContext, in *string) *pb.ForwardSshTunnelConnectivity_PrivateKey {
	if in == nil {
		return nil
	}
	out := &pb.ForwardSshTunnelConnectivity_PrivateKey{}
	out.PrivateKey = direct.ValueOf(in)
	return out
}
func MysqlProfile_FromProto(mapCtx *direct.MapContext, in *pb.MysqlProfile) *krm.MysqlProfile {
	if in == nil {
		return nil
	}
	out := &krm.MysqlProfile{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	// out.SecretRef is sensitive data, so we don't output it.
	out.SSLConfig = MysqlSSLConfig_FromProto(mapCtx, in.GetSslConfig())
	return out
}
func MysqlProfile_ToProto(mapCtx *direct.MapContext, in *krm.MysqlProfile) *pb.MysqlProfile {
	if in == nil {
		return nil
	}
	out := &pb.MysqlProfile{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	if in.SecretRef != nil {
		out.Username = in.SecretRef.Username
		out.Password = in.SecretRef.Password
	}
	out.SslConfig = MysqlSSLConfig_ToProto(mapCtx, in.SSLConfig)
	return out
}
func PostgresqlProfile_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlProfile) *krm.PostgresqlProfile {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlProfile{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	// out.SecretRef is sensitive data, so we don't output it.
	out.Database = direct.LazyPtr(in.GetDatabase())
	return out
}
func PostgresqlProfile_ToProto(mapCtx *direct.MapContext, in *krm.PostgresqlProfile) *pb.PostgresqlProfile {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlProfile{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	if in.SecretRef != nil {
		out.Username = in.SecretRef.Username
		out.Password = in.SecretRef.Password
	}
	out.Database = direct.ValueOf(in.Database)
	return out
}
func SQLServerProfile_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerProfile) *krm.SQLServerProfile {
	if in == nil {
		return nil
	}
	out := &krm.SQLServerProfile{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	// out.SecretRef is sensitive data, so we don't output it.
	out.Database = direct.LazyPtr(in.GetDatabase())
	return out
}
func SQLServerProfile_ToProto(mapCtx *direct.MapContext, in *krm.SQLServerProfile) *pb.SqlServerProfile {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerProfile{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	if in.SecretRef != nil {
		out.Username = in.SecretRef.Username
		out.Password = in.SecretRef.Password
	}
	out.Database = direct.ValueOf(in.Database)
	return out
}
