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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
)
func BigQueryProfile_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryProfile) *krm.BigQueryProfile {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryProfile{}
	return out
}
func BigQueryProfile_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryProfile) *pb.BigQueryProfile {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryProfile{}
	return out
}
func ConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionProfile{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.OracleProfile = OracleProfile_FromProto(mapCtx, in.GetOracleProfile())
	out.GcsProfile = GcsProfile_FromProto(mapCtx, in.GetGcsProfile())
	out.MysqlProfile = MysqlProfile_FromProto(mapCtx, in.GetMysqlProfile())
	out.BigqueryProfile = BigQueryProfile_FromProto(mapCtx, in.GetBigqueryProfile())
	out.PostgresqlProfile = PostgresqlProfile_FromProto(mapCtx, in.GetPostgresqlProfile())
	out.SqlServerProfile = SqlServerProfile_FromProto(mapCtx, in.GetSqlServerProfile())
	out.StaticServiceIPConnectivity = StaticServiceIpConnectivity_FromProto(mapCtx, in.GetStaticServiceIpConnectivity())
	out.ForwardSSHConnectivity = ForwardSshTunnelConnectivity_FromProto(mapCtx, in.GetForwardSshConnectivity())
	out.PrivateConnectivity = PrivateConnectivity_FromProto(mapCtx, in.GetPrivateConnectivity())
	return out
}
func ConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionProfile) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := OracleProfile_ToProto(mapCtx, in.OracleProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_OracleProfile{OracleProfile: oneof}
	}
	if oneof := GcsProfile_ToProto(mapCtx, in.GcsProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_GcsProfile{GcsProfile: oneof}
	}
	if oneof := MysqlProfile_ToProto(mapCtx, in.MysqlProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_MysqlProfile{MysqlProfile: oneof}
	}
	if oneof := BigQueryProfile_ToProto(mapCtx, in.BigqueryProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_BigqueryProfile{BigqueryProfile: oneof}
	}
	if oneof := PostgresqlProfile_ToProto(mapCtx, in.PostgresqlProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_PostgresqlProfile{PostgresqlProfile: oneof}
	}
	if oneof := SqlServerProfile_ToProto(mapCtx, in.SqlServerProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_SqlServerProfile{SqlServerProfile: oneof}
	}
	if oneof := StaticServiceIpConnectivity_ToProto(mapCtx, in.StaticServiceIPConnectivity); oneof != nil {
		out.Connectivity = &pb.ConnectionProfile_StaticServiceIpConnectivity{StaticServiceIpConnectivity: oneof}
	}
	if oneof := ForwardSshTunnelConnectivity_ToProto(mapCtx, in.ForwardSSHConnectivity); oneof != nil {
		out.Connectivity = &pb.ConnectionProfile_ForwardSshConnectivity{ForwardSshConnectivity: oneof}
	}
	if oneof := PrivateConnectivity_ToProto(mapCtx, in.PrivateConnectivity); oneof != nil {
		out.Connectivity = &pb.ConnectionProfile_PrivateConnectivity{PrivateConnectivity: oneof}
	}
	return out
}
func ConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.ConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionProfileObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	out.OracleProfile = OracleProfileObservedState_FromProto(mapCtx, in.GetOracleProfile())
	// MISSING: GcsProfile
	out.MysqlProfile = MysqlProfileObservedState_FromProto(mapCtx, in.GetMysqlProfile())
	// MISSING: BigqueryProfile
	// MISSING: PostgresqlProfile
	// MISSING: SqlServerProfile
	// MISSING: StaticServiceIPConnectivity
	// MISSING: ForwardSSHConnectivity
	// MISSING: PrivateConnectivity
	return out
}
func ConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionProfileObservedState) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	if oneof := OracleProfileObservedState_ToProto(mapCtx, in.OracleProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_OracleProfile{OracleProfile: oneof}
	}
	// MISSING: GcsProfile
	if oneof := MysqlProfileObservedState_ToProto(mapCtx, in.MysqlProfile); oneof != nil {
		out.Profile = &pb.ConnectionProfile_MysqlProfile{MysqlProfile: oneof}
	}
	// MISSING: BigqueryProfile
	// MISSING: PostgresqlProfile
	// MISSING: SqlServerProfile
	// MISSING: StaticServiceIPConnectivity
	// MISSING: ForwardSSHConnectivity
	// MISSING: PrivateConnectivity
	return out
}
func DatastreamConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.DatastreamConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamConnectionProfileObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: OracleProfile
	// MISSING: GcsProfile
	// MISSING: MysqlProfile
	// MISSING: BigqueryProfile
	// MISSING: PostgresqlProfile
	// MISSING: SqlServerProfile
	// MISSING: StaticServiceIPConnectivity
	// MISSING: ForwardSSHConnectivity
	// MISSING: PrivateConnectivity
	return out
}
func DatastreamConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamConnectionProfileObservedState) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: OracleProfile
	// MISSING: GcsProfile
	// MISSING: MysqlProfile
	// MISSING: BigqueryProfile
	// MISSING: PostgresqlProfile
	// MISSING: SqlServerProfile
	// MISSING: StaticServiceIPConnectivity
	// MISSING: ForwardSSHConnectivity
	// MISSING: PrivateConnectivity
	return out
}
func DatastreamConnectionProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.DatastreamConnectionProfileSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamConnectionProfileSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: OracleProfile
	// MISSING: GcsProfile
	// MISSING: MysqlProfile
	// MISSING: BigqueryProfile
	// MISSING: PostgresqlProfile
	// MISSING: SqlServerProfile
	// MISSING: StaticServiceIPConnectivity
	// MISSING: ForwardSSHConnectivity
	// MISSING: PrivateConnectivity
	return out
}
func DatastreamConnectionProfileSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamConnectionProfileSpec) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: OracleProfile
	// MISSING: GcsProfile
	// MISSING: MysqlProfile
	// MISSING: BigqueryProfile
	// MISSING: PostgresqlProfile
	// MISSING: SqlServerProfile
	// MISSING: StaticServiceIPConnectivity
	// MISSING: ForwardSSHConnectivity
	// MISSING: PrivateConnectivity
	return out
}
func ForwardSshTunnelConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.ForwardSshTunnelConnectivity) *krm.ForwardSshTunnelConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.ForwardSshTunnelConnectivity{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.PrivateKey = direct.LazyPtr(in.GetPrivateKey())
	return out
}
func ForwardSshTunnelConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.ForwardSshTunnelConnectivity) *pb.ForwardSshTunnelConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.ForwardSshTunnelConnectivity{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Username = direct.ValueOf(in.Username)
	out.Port = direct.ValueOf(in.Port)
	if oneof := ForwardSshTunnelConnectivity_Password_ToProto(mapCtx, in.Password); oneof != nil {
		out.AuthenticationMethod = oneof
	}
	if oneof := ForwardSshTunnelConnectivity_PrivateKey_ToProto(mapCtx, in.PrivateKey); oneof != nil {
		out.AuthenticationMethod = oneof
	}
	return out
}
func GcsProfile_FromProto(mapCtx *direct.MapContext, in *pb.GcsProfile) *krm.GcsProfile {
	if in == nil {
		return nil
	}
	out := &krm.GcsProfile{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.RootPath = direct.LazyPtr(in.GetRootPath())
	return out
}
func GcsProfile_ToProto(mapCtx *direct.MapContext, in *krm.GcsProfile) *pb.GcsProfile {
	if in == nil {
		return nil
	}
	out := &pb.GcsProfile{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.RootPath = direct.ValueOf(in.RootPath)
	return out
}
func MysqlProfile_FromProto(mapCtx *direct.MapContext, in *pb.MysqlProfile) *krm.MysqlProfile {
	if in == nil {
		return nil
	}
	out := &krm.MysqlProfile{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.SslConfig = MysqlSslConfig_FromProto(mapCtx, in.GetSslConfig())
	return out
}
func MysqlProfile_ToProto(mapCtx *direct.MapContext, in *krm.MysqlProfile) *pb.MysqlProfile {
	if in == nil {
		return nil
	}
	out := &pb.MysqlProfile{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.SslConfig = MysqlSslConfig_ToProto(mapCtx, in.SslConfig)
	return out
}
func MysqlProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MysqlProfile) *krm.MysqlProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MysqlProfileObservedState{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.SslConfig = MysqlSslConfigObservedState_FromProto(mapCtx, in.GetSslConfig())
	return out
}
func MysqlProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MysqlProfileObservedState) *pb.MysqlProfile {
	if in == nil {
		return nil
	}
	out := &pb.MysqlProfile{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.SslConfig = MysqlSslConfigObservedState_ToProto(mapCtx, in.SslConfig)
	return out
}
func MysqlSslConfig_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSslConfig) *krm.MysqlSslConfig {
	if in == nil {
		return nil
	}
	out := &krm.MysqlSslConfig{}
	out.ClientKey = direct.LazyPtr(in.GetClientKey())
	// MISSING: ClientKeySet
	out.ClientCertificate = direct.LazyPtr(in.GetClientCertificate())
	// MISSING: ClientCertificateSet
	out.CaCertificate = direct.LazyPtr(in.GetCaCertificate())
	// MISSING: CaCertificateSet
	return out
}
func MysqlSslConfig_ToProto(mapCtx *direct.MapContext, in *krm.MysqlSslConfig) *pb.MysqlSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSslConfig{}
	out.ClientKey = direct.ValueOf(in.ClientKey)
	// MISSING: ClientKeySet
	out.ClientCertificate = direct.ValueOf(in.ClientCertificate)
	// MISSING: ClientCertificateSet
	out.CaCertificate = direct.ValueOf(in.CaCertificate)
	// MISSING: CaCertificateSet
	return out
}
func MysqlSslConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSslConfig) *krm.MysqlSslConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MysqlSslConfigObservedState{}
	// MISSING: ClientKey
	out.ClientKeySet = direct.LazyPtr(in.GetClientKeySet())
	// MISSING: ClientCertificate
	out.ClientCertificateSet = direct.LazyPtr(in.GetClientCertificateSet())
	// MISSING: CaCertificate
	out.CaCertificateSet = direct.LazyPtr(in.GetCaCertificateSet())
	return out
}
func MysqlSslConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MysqlSslConfigObservedState) *pb.MysqlSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSslConfig{}
	// MISSING: ClientKey
	out.ClientKeySet = direct.ValueOf(in.ClientKeySet)
	// MISSING: ClientCertificate
	out.ClientCertificateSet = direct.ValueOf(in.ClientCertificateSet)
	// MISSING: CaCertificate
	out.CaCertificateSet = direct.ValueOf(in.CaCertificateSet)
	return out
}
func OracleAsmConfig_FromProto(mapCtx *direct.MapContext, in *pb.OracleAsmConfig) *krm.OracleAsmConfig {
	if in == nil {
		return nil
	}
	out := &krm.OracleAsmConfig{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.AsmService = direct.LazyPtr(in.GetAsmService())
	out.ConnectionAttributes = in.ConnectionAttributes
	out.OracleSslConfig = OracleSslConfig_FromProto(mapCtx, in.GetOracleSslConfig())
	return out
}
func OracleAsmConfig_ToProto(mapCtx *direct.MapContext, in *krm.OracleAsmConfig) *pb.OracleAsmConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleAsmConfig{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.AsmService = direct.ValueOf(in.AsmService)
	out.ConnectionAttributes = in.ConnectionAttributes
	out.OracleSslConfig = OracleSslConfig_ToProto(mapCtx, in.OracleSslConfig)
	return out
}
func OracleProfile_FromProto(mapCtx *direct.MapContext, in *pb.OracleProfile) *krm.OracleProfile {
	if in == nil {
		return nil
	}
	out := &krm.OracleProfile{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.DatabaseService = direct.LazyPtr(in.GetDatabaseService())
	out.ConnectionAttributes = in.ConnectionAttributes
	out.OracleSslConfig = OracleSslConfig_FromProto(mapCtx, in.GetOracleSslConfig())
	out.OracleAsmConfig = OracleAsmConfig_FromProto(mapCtx, in.GetOracleAsmConfig())
	out.SecretManagerStoredPassword = direct.LazyPtr(in.GetSecretManagerStoredPassword())
	return out
}
func OracleProfile_ToProto(mapCtx *direct.MapContext, in *krm.OracleProfile) *pb.OracleProfile {
	if in == nil {
		return nil
	}
	out := &pb.OracleProfile{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.DatabaseService = direct.ValueOf(in.DatabaseService)
	out.ConnectionAttributes = in.ConnectionAttributes
	out.OracleSslConfig = OracleSslConfig_ToProto(mapCtx, in.OracleSslConfig)
	out.OracleAsmConfig = OracleAsmConfig_ToProto(mapCtx, in.OracleAsmConfig)
	out.SecretManagerStoredPassword = direct.ValueOf(in.SecretManagerStoredPassword)
	return out
}
func OracleProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OracleProfile) *krm.OracleProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracleProfileObservedState{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	// MISSING: DatabaseService
	// MISSING: ConnectionAttributes
	out.OracleSslConfig = OracleSslConfigObservedState_FromProto(mapCtx, in.GetOracleSslConfig())
	// MISSING: OracleAsmConfig
	// MISSING: SecretManagerStoredPassword
	return out
}
func OracleProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracleProfileObservedState) *pb.OracleProfile {
	if in == nil {
		return nil
	}
	out := &pb.OracleProfile{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	// MISSING: DatabaseService
	// MISSING: ConnectionAttributes
	out.OracleSslConfig = OracleSslConfigObservedState_ToProto(mapCtx, in.OracleSslConfig)
	// MISSING: OracleAsmConfig
	// MISSING: SecretManagerStoredPassword
	return out
}
func OracleSslConfig_FromProto(mapCtx *direct.MapContext, in *pb.OracleSslConfig) *krm.OracleSslConfig {
	if in == nil {
		return nil
	}
	out := &krm.OracleSslConfig{}
	out.CaCertificate = direct.LazyPtr(in.GetCaCertificate())
	// MISSING: CaCertificateSet
	return out
}
func OracleSslConfig_ToProto(mapCtx *direct.MapContext, in *krm.OracleSslConfig) *pb.OracleSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleSslConfig{}
	out.CaCertificate = direct.ValueOf(in.CaCertificate)
	// MISSING: CaCertificateSet
	return out
}
func OracleSslConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OracleSslConfig) *krm.OracleSslConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracleSslConfigObservedState{}
	// MISSING: CaCertificate
	out.CaCertificateSet = direct.LazyPtr(in.GetCaCertificateSet())
	return out
}
func OracleSslConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracleSslConfigObservedState) *pb.OracleSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleSslConfig{}
	// MISSING: CaCertificate
	out.CaCertificateSet = direct.ValueOf(in.CaCertificateSet)
	return out
}
func PostgresqlProfile_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlProfile) *krm.PostgresqlProfile {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlProfile{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
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
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.Database = direct.ValueOf(in.Database)
	return out
}
func PrivateConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnectivity) *krm.PrivateConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.PrivateConnectivity{}
	out.PrivateConnection = direct.LazyPtr(in.GetPrivateConnection())
	return out
}
func PrivateConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnectivity) *pb.PrivateConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnectivity{}
	out.PrivateConnection = direct.ValueOf(in.PrivateConnection)
	return out
}
func SqlServerProfile_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerProfile) *krm.SqlServerProfile {
	if in == nil {
		return nil
	}
	out := &krm.SqlServerProfile{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.Database = direct.LazyPtr(in.GetDatabase())
	return out
}
func SqlServerProfile_ToProto(mapCtx *direct.MapContext, in *krm.SqlServerProfile) *pb.SqlServerProfile {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerProfile{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.Database = direct.ValueOf(in.Database)
	return out
}
func StaticServiceIpConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.StaticServiceIpConnectivity) *krm.StaticServiceIpConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.StaticServiceIpConnectivity{}
	return out
}
func StaticServiceIpConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.StaticServiceIpConnectivity) *pb.StaticServiceIpConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.StaticServiceIpConnectivity{}
	return out
}
