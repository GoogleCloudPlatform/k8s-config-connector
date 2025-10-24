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
// krm.group: datastream.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datastream.v1

package datastream

import (
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
func DatastreamPrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.DatastreamPrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamPrivateConnectionObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Error_FromProto(mapCtx, in.GetError())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	// MISSING: PSCInterfaceConfig
	return out
}
func DatastreamPrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamPrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.PrivateConnection_State](mapCtx, in.State)
	out.Error = Error_ToProto(mapCtx, in.Error)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	// MISSING: PSCInterfaceConfig
	return out
}
func DatastreamRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.DatastreamRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamRouteObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func DatastreamRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamRouteObservedState) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func DatastreamRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.DatastreamRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamRouteSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DestinationAddress = direct.LazyPtr(in.GetDestinationAddress())
	out.DestinationPort = direct.LazyPtr(in.GetDestinationPort())
	return out
}
func DatastreamRouteSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamRouteSpec) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DestinationAddress = direct.ValueOf(in.DestinationAddress)
	out.DestinationPort = direct.ValueOf(in.DestinationPort)
	return out
}
func GCSProfile_FromProto(mapCtx *direct.MapContext, in *pb.GcsProfile) *krm.GCSProfile {
	if in == nil {
		return nil
	}
	out := &krm.GCSProfile{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.RootPath = direct.LazyPtr(in.GetRootPath())
	return out
}
func GCSProfile_ToProto(mapCtx *direct.MapContext, in *krm.GCSProfile) *pb.GcsProfile {
	if in == nil {
		return nil
	}
	out := &pb.GcsProfile{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.RootPath = direct.ValueOf(in.RootPath)
	return out
}
func HostAddress_FromProto(mapCtx *direct.MapContext, in *pb.HostAddress) *krm.HostAddress {
	if in == nil {
		return nil
	}
	out := &krm.HostAddress{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	return out
}
func HostAddress_ToProto(mapCtx *direct.MapContext, in *krm.HostAddress) *pb.HostAddress {
	if in == nil {
		return nil
	}
	out := &pb.HostAddress{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	return out
}
func MongodbProfile_FromProto(mapCtx *direct.MapContext, in *pb.MongodbProfile) *krm.MongodbProfile {
	if in == nil {
		return nil
	}
	out := &krm.MongodbProfile{}
	out.HostAddresses = direct.Slice_FromProto(mapCtx, in.HostAddresses, HostAddress_FromProto)
	out.ReplicaSet = direct.LazyPtr(in.GetReplicaSet())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.SecretManagerStoredPassword = direct.LazyPtr(in.GetSecretManagerStoredPassword())
	out.SSLConfig = MongodbSSLConfig_FromProto(mapCtx, in.GetSslConfig())
	out.SrvConnectionFormat = SrvConnectionFormat_FromProto(mapCtx, in.GetSrvConnectionFormat())
	out.StandardConnectionFormat = StandardConnectionFormat_FromProto(mapCtx, in.GetStandardConnectionFormat())
	return out
}
func MongodbProfile_ToProto(mapCtx *direct.MapContext, in *krm.MongodbProfile) *pb.MongodbProfile {
	if in == nil {
		return nil
	}
	out := &pb.MongodbProfile{}
	out.HostAddresses = direct.Slice_ToProto(mapCtx, in.HostAddresses, HostAddress_ToProto)
	out.ReplicaSet = direct.ValueOf(in.ReplicaSet)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.SecretManagerStoredPassword = direct.ValueOf(in.SecretManagerStoredPassword)
	out.SslConfig = MongodbSSLConfig_ToProto(mapCtx, in.SSLConfig)
	if oneof := SrvConnectionFormat_ToProto(mapCtx, in.SrvConnectionFormat); oneof != nil {
		out.MongodbConnectionFormat = &pb.MongodbProfile_SrvConnectionFormat{SrvConnectionFormat: oneof}
	}
	if oneof := StandardConnectionFormat_ToProto(mapCtx, in.StandardConnectionFormat); oneof != nil {
		out.MongodbConnectionFormat = &pb.MongodbProfile_StandardConnectionFormat{StandardConnectionFormat: oneof}
	}
	return out
}
func MongodbProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MongodbProfile) *krm.MongodbProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MongodbProfileObservedState{}
	// MISSING: HostAddresses
	// MISSING: ReplicaSet
	// MISSING: Username
	// MISSING: Password
	// MISSING: SecretManagerStoredPassword
	out.SSLConfig = MongodbSSLConfigObservedState_FromProto(mapCtx, in.GetSslConfig())
	// MISSING: SrvConnectionFormat
	// MISSING: StandardConnectionFormat
	return out
}
func MongodbProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MongodbProfileObservedState) *pb.MongodbProfile {
	if in == nil {
		return nil
	}
	out := &pb.MongodbProfile{}
	// MISSING: HostAddresses
	// MISSING: ReplicaSet
	// MISSING: Username
	// MISSING: Password
	// MISSING: SecretManagerStoredPassword
	out.SslConfig = MongodbSSLConfigObservedState_ToProto(mapCtx, in.SSLConfig)
	// MISSING: SrvConnectionFormat
	// MISSING: StandardConnectionFormat
	return out
}
func MongodbSSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.MongodbSslConfig) *krm.MongodbSSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.MongodbSSLConfig{}
	out.ClientKey = direct.LazyPtr(in.GetClientKey())
	// MISSING: ClientKeySet
	out.ClientCertificate = direct.LazyPtr(in.GetClientCertificate())
	// MISSING: ClientCertificateSet
	out.CACertificate = direct.LazyPtr(in.GetCaCertificate())
	// MISSING: CACertificateSet
	out.SecretManagerStoredClientKey = direct.LazyPtr(in.GetSecretManagerStoredClientKey())
	return out
}
func MongodbSSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.MongodbSSLConfig) *pb.MongodbSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.MongodbSslConfig{}
	out.ClientKey = direct.ValueOf(in.ClientKey)
	// MISSING: ClientKeySet
	out.ClientCertificate = direct.ValueOf(in.ClientCertificate)
	// MISSING: ClientCertificateSet
	out.CaCertificate = direct.ValueOf(in.CACertificate)
	// MISSING: CACertificateSet
	out.SecretManagerStoredClientKey = direct.ValueOf(in.SecretManagerStoredClientKey)
	return out
}
func MongodbSSLConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MongodbSslConfig) *krm.MongodbSSLConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MongodbSSLConfigObservedState{}
	// MISSING: ClientKey
	out.ClientKeySet = direct.LazyPtr(in.GetClientKeySet())
	// MISSING: ClientCertificate
	out.ClientCertificateSet = direct.LazyPtr(in.GetClientCertificateSet())
	// MISSING: CACertificate
	out.CACertificateSet = direct.LazyPtr(in.GetCaCertificateSet())
	// MISSING: SecretManagerStoredClientKey
	return out
}
func MongodbSSLConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MongodbSSLConfigObservedState) *pb.MongodbSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.MongodbSslConfig{}
	// MISSING: ClientKey
	out.ClientKeySet = direct.ValueOf(in.ClientKeySet)
	// MISSING: ClientCertificate
	out.ClientCertificateSet = direct.ValueOf(in.ClientCertificateSet)
	// MISSING: CACertificate
	out.CaCertificateSet = direct.ValueOf(in.CACertificateSet)
	// MISSING: SecretManagerStoredClientKey
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
	out.SSLConfig = MysqlSSLConfigObservedState_FromProto(mapCtx, in.GetSslConfig())
	// MISSING: SecretManagerStoredPassword
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
	out.SslConfig = MysqlSSLConfigObservedState_ToProto(mapCtx, in.SSLConfig)
	// MISSING: SecretManagerStoredPassword
	return out
}
func MysqlSSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSslConfig) *krm.MysqlSSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.MysqlSSLConfig{}
	out.ClientKey = direct.LazyPtr(in.GetClientKey())
	// MISSING: ClientKeySet
	out.ClientCertificate = direct.LazyPtr(in.GetClientCertificate())
	// MISSING: ClientCertificateSet
	out.CACertificate = direct.LazyPtr(in.GetCaCertificate())
	// MISSING: CACertificateSet
	return out
}
func MysqlSSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.MysqlSSLConfig) *pb.MysqlSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSslConfig{}
	out.ClientKey = direct.ValueOf(in.ClientKey)
	// MISSING: ClientKeySet
	out.ClientCertificate = direct.ValueOf(in.ClientCertificate)
	// MISSING: ClientCertificateSet
	out.CaCertificate = direct.ValueOf(in.CACertificate)
	// MISSING: CACertificateSet
	return out
}
func MysqlSSLConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSslConfig) *krm.MysqlSSLConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MysqlSSLConfigObservedState{}
	// MISSING: ClientKey
	out.ClientKeySet = direct.LazyPtr(in.GetClientKeySet())
	// MISSING: ClientCertificate
	out.ClientCertificateSet = direct.LazyPtr(in.GetClientCertificateSet())
	// MISSING: CACertificate
	out.CACertificateSet = direct.LazyPtr(in.GetCaCertificateSet())
	return out
}
func MysqlSSLConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MysqlSSLConfigObservedState) *pb.MysqlSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSslConfig{}
	// MISSING: ClientKey
	out.ClientKeySet = direct.ValueOf(in.ClientKeySet)
	// MISSING: ClientCertificate
	out.ClientCertificateSet = direct.ValueOf(in.ClientCertificateSet)
	// MISSING: CACertificate
	out.CaCertificateSet = direct.ValueOf(in.CACertificateSet)
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
	out.OracleSSLConfig = OracleSSLConfigObservedState_FromProto(mapCtx, in.GetOracleSslConfig())
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
	out.OracleSslConfig = OracleSSLConfigObservedState_ToProto(mapCtx, in.OracleSSLConfig)
	// MISSING: OracleAsmConfig
	// MISSING: SecretManagerStoredPassword
	return out
}
func OracleSSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.OracleSslConfig) *krm.OracleSSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.OracleSSLConfig{}
	out.CACertificate = direct.LazyPtr(in.GetCaCertificate())
	// MISSING: CACertificateSet
	out.ServerCertificateDistinguishedName = direct.LazyPtr(in.GetServerCertificateDistinguishedName())
	return out
}
func OracleSSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.OracleSSLConfig) *pb.OracleSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleSslConfig{}
	out.CaCertificate = direct.ValueOf(in.CACertificate)
	// MISSING: CACertificateSet
	out.ServerCertificateDistinguishedName = direct.ValueOf(in.ServerCertificateDistinguishedName)
	return out
}
func OracleSSLConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OracleSslConfig) *krm.OracleSSLConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracleSSLConfigObservedState{}
	// MISSING: CACertificate
	out.CACertificateSet = direct.LazyPtr(in.GetCaCertificateSet())
	// MISSING: ServerCertificateDistinguishedName
	return out
}
func OracleSSLConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracleSSLConfigObservedState) *pb.OracleSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleSslConfig{}
	// MISSING: CACertificate
	out.CaCertificateSet = direct.ValueOf(in.CACertificateSet)
	// MISSING: ServerCertificateDistinguishedName
	return out
}
func PSCInterfaceConfig_FromProto(mapCtx *direct.MapContext, in *pb.PscInterfaceConfig) *krm.PSCInterfaceConfig {
	if in == nil {
		return nil
	}
	out := &krm.PSCInterfaceConfig{}
	out.NetworkAttachment = direct.LazyPtr(in.GetNetworkAttachment())
	return out
}
func PSCInterfaceConfig_ToProto(mapCtx *direct.MapContext, in *krm.PSCInterfaceConfig) *pb.PscInterfaceConfig {
	if in == nil {
		return nil
	}
	out := &pb.PscInterfaceConfig{}
	out.NetworkAttachment = direct.ValueOf(in.NetworkAttachment)
	return out
}
func PostgresqlSSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlSslConfig) *krm.PostgresqlSSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlSSLConfig{}
	out.ServerVerification = PostgresqlSSLConfig_ServerVerification_FromProto(mapCtx, in.GetServerVerification())
	out.ServerAndClientVerification = PostgresqlSSLConfig_ServerAndClientVerification_FromProto(mapCtx, in.GetServerAndClientVerification())
	return out
}
func PostgresqlSSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.PostgresqlSSLConfig) *pb.PostgresqlSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlSslConfig{}
	if oneof := PostgresqlSSLConfig_ServerVerification_ToProto(mapCtx, in.ServerVerification); oneof != nil {
		out.EncryptionSetting = &pb.PostgresqlSslConfig_ServerVerification_{ServerVerification: oneof}
	}
	if oneof := PostgresqlSSLConfig_ServerAndClientVerification_ToProto(mapCtx, in.ServerAndClientVerification); oneof != nil {
		out.EncryptionSetting = &pb.PostgresqlSslConfig_ServerAndClientVerification_{ServerAndClientVerification: oneof}
	}
	return out
}
func PostgresqlSSLConfig_ServerAndClientVerification_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlSslConfig_ServerAndClientVerification) *krm.PostgresqlSSLConfig_ServerAndClientVerification {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlSSLConfig_ServerAndClientVerification{}
	out.ClientCertificate = direct.LazyPtr(in.GetClientCertificate())
	out.ClientKey = direct.LazyPtr(in.GetClientKey())
	out.CACertificate = direct.LazyPtr(in.GetCaCertificate())
	out.ServerCertificateHostname = direct.LazyPtr(in.GetServerCertificateHostname())
	return out
}
func PostgresqlSSLConfig_ServerAndClientVerification_ToProto(mapCtx *direct.MapContext, in *krm.PostgresqlSSLConfig_ServerAndClientVerification) *pb.PostgresqlSslConfig_ServerAndClientVerification {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlSslConfig_ServerAndClientVerification{}
	out.ClientCertificate = direct.ValueOf(in.ClientCertificate)
	out.ClientKey = direct.ValueOf(in.ClientKey)
	out.CaCertificate = direct.ValueOf(in.CACertificate)
	out.ServerCertificateHostname = direct.ValueOf(in.ServerCertificateHostname)
	return out
}
func PostgresqlSSLConfig_ServerVerification_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlSslConfig_ServerVerification) *krm.PostgresqlSSLConfig_ServerVerification {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlSSLConfig_ServerVerification{}
	out.CACertificate = direct.LazyPtr(in.GetCaCertificate())
	out.ServerCertificateHostname = direct.LazyPtr(in.GetServerCertificateHostname())
	return out
}
func PostgresqlSSLConfig_ServerVerification_ToProto(mapCtx *direct.MapContext, in *krm.PostgresqlSSLConfig_ServerVerification) *pb.PostgresqlSslConfig_ServerVerification {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlSslConfig_ServerVerification{}
	out.CaCertificate = direct.ValueOf(in.CACertificate)
	out.ServerCertificateHostname = direct.ValueOf(in.ServerCertificateHostname)
	return out
}
func PrivateConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnectivity) *pb.PrivateConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnectivity{}
	if in.PrivateConnectionRef != nil {
		out.PrivateConnection = in.PrivateConnectionRef.External
	}
	return out
}
func SalesforceProfile_FromProto(mapCtx *direct.MapContext, in *pb.SalesforceProfile) *krm.SalesforceProfile {
	if in == nil {
		return nil
	}
	out := &krm.SalesforceProfile{}
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.UserCredentials = SalesforceProfile_UserCredentials_FromProto(mapCtx, in.GetUserCredentials())
	out.OAUTH2ClientCredentials = SalesforceProfile_OAUTH2ClientCredentials_FromProto(mapCtx, in.GetOauth2ClientCredentials())
	return out
}
func SalesforceProfile_ToProto(mapCtx *direct.MapContext, in *krm.SalesforceProfile) *pb.SalesforceProfile {
	if in == nil {
		return nil
	}
	out := &pb.SalesforceProfile{}
	out.Domain = direct.ValueOf(in.Domain)
	if oneof := SalesforceProfile_UserCredentials_ToProto(mapCtx, in.UserCredentials); oneof != nil {
		out.Credentials = &pb.SalesforceProfile_UserCredentials_{UserCredentials: oneof}
	}
	if oneof := SalesforceProfile_OAUTH2ClientCredentials_ToProto(mapCtx, in.OAUTH2ClientCredentials); oneof != nil {
		out.Credentials = &pb.SalesforceProfile_Oauth2ClientCredentials_{Oauth2ClientCredentials: oneof}
	}
	return out
}
func SalesforceProfile_OAUTH2ClientCredentials_FromProto(mapCtx *direct.MapContext, in *pb.SalesforceProfile_Oauth2ClientCredentials) *krm.SalesforceProfile_OAUTH2ClientCredentials {
	if in == nil {
		return nil
	}
	out := &krm.SalesforceProfile_OAUTH2ClientCredentials{}
	out.ClientID = direct.LazyPtr(in.GetClientId())
	out.ClientSecret = direct.LazyPtr(in.GetClientSecret())
	out.SecretManagerStoredClientSecret = direct.LazyPtr(in.GetSecretManagerStoredClientSecret())
	return out
}
func SalesforceProfile_OAUTH2ClientCredentials_ToProto(mapCtx *direct.MapContext, in *krm.SalesforceProfile_OAUTH2ClientCredentials) *pb.SalesforceProfile_Oauth2ClientCredentials {
	if in == nil {
		return nil
	}
	out := &pb.SalesforceProfile_Oauth2ClientCredentials{}
	out.ClientId = direct.ValueOf(in.ClientID)
	out.ClientSecret = direct.ValueOf(in.ClientSecret)
	out.SecretManagerStoredClientSecret = direct.ValueOf(in.SecretManagerStoredClientSecret)
	return out
}
func SalesforceProfile_UserCredentials_FromProto(mapCtx *direct.MapContext, in *pb.SalesforceProfile_UserCredentials) *krm.SalesforceProfile_UserCredentials {
	if in == nil {
		return nil
	}
	out := &krm.SalesforceProfile_UserCredentials{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.SecurityToken = direct.LazyPtr(in.GetSecurityToken())
	out.SecretManagerStoredPassword = direct.LazyPtr(in.GetSecretManagerStoredPassword())
	out.SecretManagerStoredSecurityToken = direct.LazyPtr(in.GetSecretManagerStoredSecurityToken())
	return out
}
func SalesforceProfile_UserCredentials_ToProto(mapCtx *direct.MapContext, in *krm.SalesforceProfile_UserCredentials) *pb.SalesforceProfile_UserCredentials {
	if in == nil {
		return nil
	}
	out := &pb.SalesforceProfile_UserCredentials{}
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.SecurityToken = direct.ValueOf(in.SecurityToken)
	out.SecretManagerStoredPassword = direct.ValueOf(in.SecretManagerStoredPassword)
	out.SecretManagerStoredSecurityToken = direct.ValueOf(in.SecretManagerStoredSecurityToken)
	return out
}
func SrvConnectionFormat_FromProto(mapCtx *direct.MapContext, in *pb.SrvConnectionFormat) *krm.SrvConnectionFormat {
	if in == nil {
		return nil
	}
	out := &krm.SrvConnectionFormat{}
	return out
}
func SrvConnectionFormat_ToProto(mapCtx *direct.MapContext, in *krm.SrvConnectionFormat) *pb.SrvConnectionFormat {
	if in == nil {
		return nil
	}
	out := &pb.SrvConnectionFormat{}
	return out
}
func StandardConnectionFormat_FromProto(mapCtx *direct.MapContext, in *pb.StandardConnectionFormat) *krm.StandardConnectionFormat {
	if in == nil {
		return nil
	}
	out := &krm.StandardConnectionFormat{}
	out.DirectConnection = direct.LazyPtr(in.GetDirectConnection())
	return out
}
func StandardConnectionFormat_ToProto(mapCtx *direct.MapContext, in *krm.StandardConnectionFormat) *pb.StandardConnectionFormat {
	if in == nil {
		return nil
	}
	out := &pb.StandardConnectionFormat{}
	out.DirectConnection = direct.ValueOf(in.DirectConnection)
	return out
}
func StaticServiceIPConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.StaticServiceIpConnectivity) *krm.StaticServiceIPConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.StaticServiceIPConnectivity{}
	return out
}
func StaticServiceIPConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.StaticServiceIPConnectivity) *pb.StaticServiceIpConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.StaticServiceIpConnectivity{}
	return out
}
