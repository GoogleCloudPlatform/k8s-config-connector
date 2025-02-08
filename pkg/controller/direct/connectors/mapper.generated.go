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

package connectors

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/connectors/apiv1/connectorspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
)
func AuthConfig_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfig) *krm.AuthConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig{}
	out.AuthType = direct.Enum_FromProto(mapCtx, in.GetAuthType())
	out.UserPassword = AuthConfig_UserPassword_FromProto(mapCtx, in.GetUserPassword())
	out.Oauth2JwtBearer = AuthConfig_Oauth2JwtBearer_FromProto(mapCtx, in.GetOauth2JwtBearer())
	out.Oauth2ClientCredentials = AuthConfig_Oauth2ClientCredentials_FromProto(mapCtx, in.GetOauth2ClientCredentials())
	out.SSHPublicKey = AuthConfig_SshPublicKey_FromProto(mapCtx, in.GetSshPublicKey())
	out.AdditionalVariables = direct.Slice_FromProto(mapCtx, in.AdditionalVariables, ConfigVariable_FromProto)
	return out
}
func AuthConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig) *pb.AuthConfig {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfig{}
	out.AuthType = direct.Enum_ToProto[pb.AuthType](mapCtx, in.AuthType)
	if oneof := AuthConfig_UserPassword_ToProto(mapCtx, in.UserPassword); oneof != nil {
		out.Type = &pb.AuthConfig_UserPassword_{UserPassword: oneof}
	}
	if oneof := AuthConfig_Oauth2JwtBearer_ToProto(mapCtx, in.Oauth2JwtBearer); oneof != nil {
		out.Type = &pb.AuthConfig_Oauth2JwtBearer_{Oauth2JwtBearer: oneof}
	}
	if oneof := AuthConfig_Oauth2ClientCredentials_ToProto(mapCtx, in.Oauth2ClientCredentials); oneof != nil {
		out.Type = &pb.AuthConfig_Oauth2ClientCredentials_{Oauth2ClientCredentials: oneof}
	}
	if oneof := AuthConfig_SshPublicKey_ToProto(mapCtx, in.SSHPublicKey); oneof != nil {
		out.Type = &pb.AuthConfig_SshPublicKey_{SshPublicKey: oneof}
	}
	out.AdditionalVariables = direct.Slice_ToProto(mapCtx, in.AdditionalVariables, ConfigVariable_ToProto)
	return out
}
func AuthConfig_Oauth2ClientCredentials_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfig_Oauth2ClientCredentials) *krm.AuthConfig_Oauth2ClientCredentials {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_Oauth2ClientCredentials{}
	out.ClientID = direct.LazyPtr(in.GetClientId())
	out.ClientSecret = Secret_FromProto(mapCtx, in.GetClientSecret())
	return out
}
func AuthConfig_Oauth2ClientCredentials_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_Oauth2ClientCredentials) *pb.AuthConfig_Oauth2ClientCredentials {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfig_Oauth2ClientCredentials{}
	out.ClientId = direct.ValueOf(in.ClientID)
	out.ClientSecret = Secret_ToProto(mapCtx, in.ClientSecret)
	return out
}
func AuthConfig_Oauth2JwtBearer_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfig_Oauth2JwtBearer) *krm.AuthConfig_Oauth2JwtBearer {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_Oauth2JwtBearer{}
	out.ClientKey = Secret_FromProto(mapCtx, in.GetClientKey())
	out.JwtClaims = AuthConfig_Oauth2JwtBearer_JwtClaims_FromProto(mapCtx, in.GetJwtClaims())
	return out
}
func AuthConfig_Oauth2JwtBearer_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_Oauth2JwtBearer) *pb.AuthConfig_Oauth2JwtBearer {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfig_Oauth2JwtBearer{}
	out.ClientKey = Secret_ToProto(mapCtx, in.ClientKey)
	out.JwtClaims = AuthConfig_Oauth2JwtBearer_JwtClaims_ToProto(mapCtx, in.JwtClaims)
	return out
}
func AuthConfig_Oauth2JwtBearer_JwtClaims_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfig_Oauth2JwtBearer_JwtClaims) *krm.AuthConfig_Oauth2JwtBearer_JwtClaims {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_Oauth2JwtBearer_JwtClaims{}
	out.Issuer = direct.LazyPtr(in.GetIssuer())
	out.Subject = direct.LazyPtr(in.GetSubject())
	out.Audience = direct.LazyPtr(in.GetAudience())
	return out
}
func AuthConfig_Oauth2JwtBearer_JwtClaims_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_Oauth2JwtBearer_JwtClaims) *pb.AuthConfig_Oauth2JwtBearer_JwtClaims {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfig_Oauth2JwtBearer_JwtClaims{}
	out.Issuer = direct.ValueOf(in.Issuer)
	out.Subject = direct.ValueOf(in.Subject)
	out.Audience = direct.ValueOf(in.Audience)
	return out
}
func AuthConfig_SshPublicKey_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfig_SshPublicKey) *krm.AuthConfig_SshPublicKey {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_SshPublicKey{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.SSHClientCert = Secret_FromProto(mapCtx, in.GetSshClientCert())
	out.CertType = direct.LazyPtr(in.GetCertType())
	out.SSHClientCertPass = Secret_FromProto(mapCtx, in.GetSshClientCertPass())
	return out
}
func AuthConfig_SshPublicKey_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_SshPublicKey) *pb.AuthConfig_SshPublicKey {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfig_SshPublicKey{}
	out.Username = direct.ValueOf(in.Username)
	out.SshClientCert = Secret_ToProto(mapCtx, in.SSHClientCert)
	out.CertType = direct.ValueOf(in.CertType)
	out.SshClientCertPass = Secret_ToProto(mapCtx, in.SSHClientCertPass)
	return out
}
func AuthConfig_UserPassword_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfig_UserPassword) *krm.AuthConfig_UserPassword {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_UserPassword{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = Secret_FromProto(mapCtx, in.GetPassword())
	return out
}
func AuthConfig_UserPassword_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_UserPassword) *pb.AuthConfig_UserPassword {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfig_UserPassword{}
	out.Username = direct.ValueOf(in.Username)
	out.Password = Secret_ToProto(mapCtx, in.Password)
	return out
}
func ConfigVariable_FromProto(mapCtx *direct.MapContext, in *pb.ConfigVariable) *krm.ConfigVariable {
	if in == nil {
		return nil
	}
	out := &krm.ConfigVariable{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.IntValue = direct.LazyPtr(in.GetIntValue())
	out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	out.SecretValue = Secret_FromProto(mapCtx, in.GetSecretValue())
	return out
}
func ConfigVariable_ToProto(mapCtx *direct.MapContext, in *krm.ConfigVariable) *pb.ConfigVariable {
	if in == nil {
		return nil
	}
	out := &pb.ConfigVariable{}
	out.Key = direct.ValueOf(in.Key)
	if oneof := ConfigVariable_IntValue_ToProto(mapCtx, in.IntValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := ConfigVariable_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := ConfigVariable_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := Secret_ToProto(mapCtx, in.SecretValue); oneof != nil {
		out.Value = &pb.ConfigVariable_SecretValue{SecretValue: oneof}
	}
	return out
}
func Connection_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.Connection {
	if in == nil {
		return nil
	}
	out := &krm.Connection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ConnectorVersion = direct.LazyPtr(in.GetConnectorVersion())
	// MISSING: Status
	out.ConfigVariables = direct.Slice_FromProto(mapCtx, in.ConfigVariables, ConfigVariable_FromProto)
	out.AuthConfig = AuthConfig_FromProto(mapCtx, in.GetAuthConfig())
	out.LockConfig = LockConfig_FromProto(mapCtx, in.GetLockConfig())
	out.DestinationConfigs = direct.Slice_FromProto(mapCtx, in.DestinationConfigs, DestinationConfig_FromProto)
	// MISSING: ImageLocation
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	// MISSING: ServiceDirectory
	// MISSING: EnvoyImageLocation
	out.Suspended = direct.LazyPtr(in.GetSuspended())
	out.NodeConfig = NodeConfig_FromProto(mapCtx, in.GetNodeConfig())
	out.SslConfig = SslConfig_FromProto(mapCtx, in.GetSslConfig())
	return out
}
func Connection_ToProto(mapCtx *direct.MapContext, in *krm.Connection) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.ConnectorVersion = direct.ValueOf(in.ConnectorVersion)
	// MISSING: Status
	out.ConfigVariables = direct.Slice_ToProto(mapCtx, in.ConfigVariables, ConfigVariable_ToProto)
	out.AuthConfig = AuthConfig_ToProto(mapCtx, in.AuthConfig)
	out.LockConfig = LockConfig_ToProto(mapCtx, in.LockConfig)
	out.DestinationConfigs = direct.Slice_ToProto(mapCtx, in.DestinationConfigs, DestinationConfig_ToProto)
	// MISSING: ImageLocation
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	// MISSING: ServiceDirectory
	// MISSING: EnvoyImageLocation
	out.Suspended = direct.ValueOf(in.Suspended)
	out.NodeConfig = NodeConfig_ToProto(mapCtx, in.NodeConfig)
	out.SslConfig = SslConfig_ToProto(mapCtx, in.SslConfig)
	return out
}
func ConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.ConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: ConnectorVersion
	out.Status = ConnectionStatus_FromProto(mapCtx, in.GetStatus())
	// MISSING: ConfigVariables
	// MISSING: AuthConfig
	// MISSING: LockConfig
	// MISSING: DestinationConfigs
	out.ImageLocation = direct.LazyPtr(in.GetImageLocation())
	// MISSING: ServiceAccount
	out.ServiceDirectory = direct.LazyPtr(in.GetServiceDirectory())
	out.EnvoyImageLocation = direct.LazyPtr(in.GetEnvoyImageLocation())
	// MISSING: Suspended
	// MISSING: NodeConfig
	// MISSING: SslConfig
	return out
}
func ConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionObservedState) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: ConnectorVersion
	out.Status = ConnectionStatus_ToProto(mapCtx, in.Status)
	// MISSING: ConfigVariables
	// MISSING: AuthConfig
	// MISSING: LockConfig
	// MISSING: DestinationConfigs
	out.ImageLocation = direct.ValueOf(in.ImageLocation)
	// MISSING: ServiceAccount
	out.ServiceDirectory = direct.ValueOf(in.ServiceDirectory)
	out.EnvoyImageLocation = direct.ValueOf(in.EnvoyImageLocation)
	// MISSING: Suspended
	// MISSING: NodeConfig
	// MISSING: SslConfig
	return out
}
func ConnectionStatus_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionStatus) *krm.ConnectionStatus {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Status = direct.LazyPtr(in.GetStatus())
	return out
}
func ConnectionStatus_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionStatus) *pb.ConnectionStatus {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionStatus{}
	out.State = direct.Enum_ToProto[pb.ConnectionStatus_State](mapCtx, in.State)
	out.Description = direct.ValueOf(in.Description)
	out.Status = direct.ValueOf(in.Status)
	return out
}
func ConnectorsConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.ConnectorsConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsConnectionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: ConnectorVersion
	// MISSING: Status
	// MISSING: ConfigVariables
	// MISSING: AuthConfig
	// MISSING: LockConfig
	// MISSING: DestinationConfigs
	// MISSING: ImageLocation
	// MISSING: ServiceAccount
	// MISSING: ServiceDirectory
	// MISSING: EnvoyImageLocation
	// MISSING: Suspended
	// MISSING: NodeConfig
	// MISSING: SslConfig
	return out
}
func ConnectorsConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsConnectionObservedState) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: ConnectorVersion
	// MISSING: Status
	// MISSING: ConfigVariables
	// MISSING: AuthConfig
	// MISSING: LockConfig
	// MISSING: DestinationConfigs
	// MISSING: ImageLocation
	// MISSING: ServiceAccount
	// MISSING: ServiceDirectory
	// MISSING: EnvoyImageLocation
	// MISSING: Suspended
	// MISSING: NodeConfig
	// MISSING: SslConfig
	return out
}
func ConnectorsConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.ConnectorsConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsConnectionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: ConnectorVersion
	// MISSING: Status
	// MISSING: ConfigVariables
	// MISSING: AuthConfig
	// MISSING: LockConfig
	// MISSING: DestinationConfigs
	// MISSING: ImageLocation
	// MISSING: ServiceAccount
	// MISSING: ServiceDirectory
	// MISSING: EnvoyImageLocation
	// MISSING: Suspended
	// MISSING: NodeConfig
	// MISSING: SslConfig
	return out
}
func ConnectorsConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsConnectionSpec) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: ConnectorVersion
	// MISSING: Status
	// MISSING: ConfigVariables
	// MISSING: AuthConfig
	// MISSING: LockConfig
	// MISSING: DestinationConfigs
	// MISSING: ImageLocation
	// MISSING: ServiceAccount
	// MISSING: ServiceDirectory
	// MISSING: EnvoyImageLocation
	// MISSING: Suspended
	// MISSING: NodeConfig
	// MISSING: SslConfig
	return out
}
func Destination_FromProto(mapCtx *direct.MapContext, in *pb.Destination) *krm.Destination {
	if in == nil {
		return nil
	}
	out := &krm.Destination{}
	out.ServiceAttachment = direct.LazyPtr(in.GetServiceAttachment())
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	return out
}
func Destination_ToProto(mapCtx *direct.MapContext, in *krm.Destination) *pb.Destination {
	if in == nil {
		return nil
	}
	out := &pb.Destination{}
	if oneof := Destination_ServiceAttachment_ToProto(mapCtx, in.ServiceAttachment); oneof != nil {
		out.Destination = oneof
	}
	if oneof := Destination_Host_ToProto(mapCtx, in.Host); oneof != nil {
		out.Destination = oneof
	}
	out.Port = direct.ValueOf(in.Port)
	return out
}
func DestinationConfig_FromProto(mapCtx *direct.MapContext, in *pb.DestinationConfig) *krm.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &krm.DestinationConfig{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Destinations = direct.Slice_FromProto(mapCtx, in.Destinations, Destination_FromProto)
	return out
}
func DestinationConfig_ToProto(mapCtx *direct.MapContext, in *krm.DestinationConfig) *pb.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &pb.DestinationConfig{}
	out.Key = direct.ValueOf(in.Key)
	out.Destinations = direct.Slice_ToProto(mapCtx, in.Destinations, Destination_ToProto)
	return out
}
func LockConfig_FromProto(mapCtx *direct.MapContext, in *pb.LockConfig) *krm.LockConfig {
	if in == nil {
		return nil
	}
	out := &krm.LockConfig{}
	out.Locked = direct.LazyPtr(in.GetLocked())
	out.Reason = direct.LazyPtr(in.GetReason())
	return out
}
func LockConfig_ToProto(mapCtx *direct.MapContext, in *krm.LockConfig) *pb.LockConfig {
	if in == nil {
		return nil
	}
	out := &pb.LockConfig{}
	out.Locked = direct.ValueOf(in.Locked)
	out.Reason = direct.ValueOf(in.Reason)
	return out
}
func NodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeConfig) *krm.NodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig{}
	out.MinNodeCount = direct.LazyPtr(in.GetMinNodeCount())
	out.MaxNodeCount = direct.LazyPtr(in.GetMaxNodeCount())
	return out
}
func NodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeConfig) *pb.NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeConfig{}
	out.MinNodeCount = direct.ValueOf(in.MinNodeCount)
	out.MaxNodeCount = direct.ValueOf(in.MaxNodeCount)
	return out
}
func Secret_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.Secret {
	if in == nil {
		return nil
	}
	out := &krm.Secret{}
	out.SecretVersion = direct.LazyPtr(in.GetSecretVersion())
	return out
}
func Secret_ToProto(mapCtx *direct.MapContext, in *krm.Secret) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	out.SecretVersion = direct.ValueOf(in.SecretVersion)
	return out
}
func SslConfig_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krm.SslConfig {
	if in == nil {
		return nil
	}
	out := &krm.SslConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.TrustModel = direct.Enum_FromProto(mapCtx, in.GetTrustModel())
	out.PrivateServerCertificate = Secret_FromProto(mapCtx, in.GetPrivateServerCertificate())
	out.ClientCertificate = Secret_FromProto(mapCtx, in.GetClientCertificate())
	out.ClientPrivateKey = Secret_FromProto(mapCtx, in.GetClientPrivateKey())
	out.ClientPrivateKeyPass = Secret_FromProto(mapCtx, in.GetClientPrivateKeyPass())
	out.ServerCertType = direct.Enum_FromProto(mapCtx, in.GetServerCertType())
	out.ClientCertType = direct.Enum_FromProto(mapCtx, in.GetClientCertType())
	out.UseSsl = direct.LazyPtr(in.GetUseSsl())
	out.AdditionalVariables = direct.Slice_FromProto(mapCtx, in.AdditionalVariables, ConfigVariable_FromProto)
	return out
}
func SslConfig_ToProto(mapCtx *direct.MapContext, in *krm.SslConfig) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	out.Type = direct.Enum_ToProto[pb.SslType](mapCtx, in.Type)
	out.TrustModel = direct.Enum_ToProto[pb.SslConfig_TrustModel](mapCtx, in.TrustModel)
	out.PrivateServerCertificate = Secret_ToProto(mapCtx, in.PrivateServerCertificate)
	out.ClientCertificate = Secret_ToProto(mapCtx, in.ClientCertificate)
	out.ClientPrivateKey = Secret_ToProto(mapCtx, in.ClientPrivateKey)
	out.ClientPrivateKeyPass = Secret_ToProto(mapCtx, in.ClientPrivateKeyPass)
	out.ServerCertType = direct.Enum_ToProto[pb.CertType](mapCtx, in.ServerCertType)
	out.ClientCertType = direct.Enum_ToProto[pb.CertType](mapCtx, in.ClientCertType)
	out.UseSsl = direct.ValueOf(in.UseSsl)
	out.AdditionalVariables = direct.Slice_ToProto(mapCtx, in.AdditionalVariables, ConfigVariable_ToProto)
	return out
}
