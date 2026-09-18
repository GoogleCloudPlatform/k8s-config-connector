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

package connectors

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/connectors/v1"
)

func ConnectorsConnectionSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ConnectorsConnectionSpec) *api.Connection {
	if in == nil {
		return nil
	}
	out := &api.Connection{}
	if in.Labels != nil {
		out.Labels = make(map[string]string, len(in.Labels))
		for k, v := range in.Labels {
			out.Labels[k] = v
		}
	}
	out.Description = direct.ValueOf(in.Description)
	if in.ConnectorVersionRef != nil {
		out.ConnectorVersion = in.ConnectorVersionRef.External
	}
	out.ConfigVariables = direct.Slice_ToProto(mapCtx, in.ConfigVariables, ConfigVariable_ToAPI)
	out.AuthConfig = AuthConfig_ToAPI(mapCtx, in.AuthConfig)
	out.LockConfig = LockConfig_ToAPI(mapCtx, in.LockConfig)
	out.DestinationConfigs = direct.Slice_ToProto(mapCtx, in.DestinationConfigs, DestinationConfig_ToAPI)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Suspended = direct.ValueOf(in.Suspended)
	out.NodeConfig = NodeConfig_ToAPI(mapCtx, in.NodeConfig)
	out.SslConfig = SSLConfig_ToAPI(mapCtx, in.SSLConfig)
	return out
}

func ConnectorsConnectionSpec_FromAPI(mapCtx *direct.MapContext, in *api.Connection) *krm.ConnectorsConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsConnectionSpec{}
	if in.Labels != nil {
		out.Labels = make(map[string]string, len(in.Labels))
		for k, v := range in.Labels {
			out.Labels[k] = v
		}
	}
	out.Description = direct.LazyPtr(in.Description)
	if in.ConnectorVersion != "" {
		out.ConnectorVersionRef = &krm.ConnectorsConnectorVersionRef{External: in.ConnectorVersion}
	}
	out.ConfigVariables = direct.Slice_FromProto(mapCtx, in.ConfigVariables, ConfigVariable_FromAPI)
	out.AuthConfig = AuthConfig_FromAPI(mapCtx, in.AuthConfig)
	out.LockConfig = LockConfig_FromAPI(mapCtx, in.LockConfig)
	out.DestinationConfigs = direct.Slice_FromProto(mapCtx, in.DestinationConfigs, DestinationConfig_FromAPI)
	if in.ServiceAccount != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.ServiceAccount}
	}
	out.Suspended = direct.LazyPtr(in.Suspended)
	out.NodeConfig = NodeConfig_FromAPI(mapCtx, in.NodeConfig)
	out.SSLConfig = SSLConfig_FromAPI(mapCtx, in.SslConfig)
	return out
}

func ConnectorsConnectionObservedState_FromAPI(mapCtx *direct.MapContext, in *api.Connection) *krm.ConnectorsConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsConnectionObservedState{}
	out.Name = direct.LazyPtr(in.Name)
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	out.Status = ConnectionStatus_FromAPI(mapCtx, in.Status)
	out.ImageLocation = direct.LazyPtr(in.ImageLocation)
	out.ServiceDirectory = direct.LazyPtr(in.ServiceDirectory)
	out.EnvoyImageLocation = direct.LazyPtr(in.EnvoyImageLocation)
	return out
}

func ConfigVariable_ToAPI(mapCtx *direct.MapContext, in *krm.ConfigVariable) *api.ConfigVariable {
	if in == nil {
		return nil
	}
	out := &api.ConfigVariable{}
	out.Key = direct.ValueOf(in.Key)
	out.IntValue = direct.ValueOf(in.IntValue)
	out.BoolValue = direct.ValueOf(in.BoolValue)
	out.StringValue = direct.ValueOf(in.StringValue)
	if in.SecretValueRef != nil {
		out.SecretValue = &api.Secret{SecretVersion: in.SecretValueRef.External}
	}
	return out
}

func ConfigVariable_FromAPI(mapCtx *direct.MapContext, in *api.ConfigVariable) *krm.ConfigVariable {
	if in == nil {
		return nil
	}
	out := &krm.ConfigVariable{}
	out.Key = direct.LazyPtr(in.Key)
	out.IntValue = direct.LazyPtr(in.IntValue)
	out.BoolValue = direct.LazyPtr(in.BoolValue)
	out.StringValue = direct.LazyPtr(in.StringValue)
	if in.SecretValue != nil && in.SecretValue.SecretVersion != "" {
		out.SecretValueRef = &secretmanagerv1beta1.SecretRef{External: in.SecretValue.SecretVersion}
	}
	return out
}

func AuthConfig_ToAPI(mapCtx *direct.MapContext, in *krm.AuthConfig) *api.AuthConfig {
	if in == nil {
		return nil
	}
	out := &api.AuthConfig{}
	out.AuthType = direct.ValueOf(in.AuthType)
	out.UserPassword = AuthConfig_UserPassword_ToAPI(mapCtx, in.UserPassword)
	out.Oauth2JwtBearer = AuthConfig_OAUTH2JwtBearer_ToAPI(mapCtx, in.OAUTH2JwtBearer)
	out.Oauth2ClientCredentials = AuthConfig_OAUTH2ClientCredentials_ToAPI(mapCtx, in.OAUTH2ClientCredentials)
	out.SshPublicKey = AuthConfig_SSHPublicKey_ToAPI(mapCtx, in.SSHPublicKey)
	out.AdditionalVariables = direct.Slice_ToProto(mapCtx, in.AdditionalVariables, ConfigVariable_ToAPI)
	return out
}

func AuthConfig_FromAPI(mapCtx *direct.MapContext, in *api.AuthConfig) *krm.AuthConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig{}
	out.AuthType = direct.LazyPtr(in.AuthType)
	out.UserPassword = AuthConfig_UserPassword_FromAPI(mapCtx, in.UserPassword)
	out.OAUTH2JwtBearer = AuthConfig_OAUTH2JwtBearer_FromAPI(mapCtx, in.Oauth2JwtBearer)
	out.OAUTH2ClientCredentials = AuthConfig_OAUTH2ClientCredentials_FromAPI(mapCtx, in.Oauth2ClientCredentials)
	out.SSHPublicKey = AuthConfig_SSHPublicKey_FromAPI(mapCtx, in.SshPublicKey)
	out.AdditionalVariables = direct.Slice_FromProto(mapCtx, in.AdditionalVariables, ConfigVariable_FromAPI)
	return out
}

func AuthConfig_UserPassword_ToAPI(mapCtx *direct.MapContext, in *krm.AuthConfig_UserPassword) *api.UserPassword {
	if in == nil {
		return nil
	}
	out := &api.UserPassword{}
	out.Username = direct.ValueOf(in.Username)
	if in.SecretRef != nil {
		out.Password = &api.Secret{SecretVersion: in.SecretRef.External}
	}
	return out
}

func AuthConfig_UserPassword_FromAPI(mapCtx *direct.MapContext, in *api.UserPassword) *krm.AuthConfig_UserPassword {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_UserPassword{}
	out.Username = direct.LazyPtr(in.Username)
	if in.Password != nil && in.Password.SecretVersion != "" {
		out.SecretRef = &secretmanagerv1beta1.SecretRef{External: in.Password.SecretVersion}
	}
	return out
}

func AuthConfig_OAUTH2JwtBearer_ToAPI(mapCtx *direct.MapContext, in *krm.AuthConfig_OAUTH2JwtBearer) *api.Oauth2JwtBearer {
	if in == nil {
		return nil
	}
	out := &api.Oauth2JwtBearer{}
	if in.ClientKeyRef != nil {
		out.ClientKey = &api.Secret{SecretVersion: in.ClientKeyRef.External}
	}
	out.JwtClaims = AuthConfig_OAUTH2JwtBearer_JwtClaims_ToAPI(mapCtx, in.JwtClaims)
	return out
}

func AuthConfig_OAUTH2JwtBearer_FromAPI(mapCtx *direct.MapContext, in *api.Oauth2JwtBearer) *krm.AuthConfig_OAUTH2JwtBearer {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_OAUTH2JwtBearer{}
	if in.ClientKey != nil && in.ClientKey.SecretVersion != "" {
		out.ClientKeyRef = &secretmanagerv1beta1.SecretRef{External: in.ClientKey.SecretVersion}
	}
	out.JwtClaims = AuthConfig_OAUTH2JwtBearer_JwtClaims_FromAPI(mapCtx, in.JwtClaims)
	return out
}

func AuthConfig_OAUTH2JwtBearer_JwtClaims_ToAPI(mapCtx *direct.MapContext, in *krm.AuthConfig_OAUTH2JwtBearer_JwtClaims) *api.JwtClaims {
	if in == nil {
		return nil
	}
	out := &api.JwtClaims{}
	out.Issuer = direct.ValueOf(in.Issuer)
	out.Subject = direct.ValueOf(in.Subject)
	out.Audience = direct.ValueOf(in.Audience)
	return out
}

func AuthConfig_OAUTH2JwtBearer_JwtClaims_FromAPI(mapCtx *direct.MapContext, in *api.JwtClaims) *krm.AuthConfig_OAUTH2JwtBearer_JwtClaims {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_OAUTH2JwtBearer_JwtClaims{}
	out.Issuer = direct.LazyPtr(in.Issuer)
	out.Subject = direct.LazyPtr(in.Subject)
	out.Audience = direct.LazyPtr(in.Audience)
	return out
}

func AuthConfig_OAUTH2ClientCredentials_ToAPI(mapCtx *direct.MapContext, in *krm.AuthConfig_OAUTH2ClientCredentials) *api.Oauth2ClientCredentials {
	if in == nil {
		return nil
	}
	out := &api.Oauth2ClientCredentials{}
	out.ClientId = direct.ValueOf(in.ClientID)
	if in.ClientSecretRef != nil {
		out.ClientSecret = &api.Secret{SecretVersion: in.ClientSecretRef.External}
	}
	return out
}

func AuthConfig_OAUTH2ClientCredentials_FromAPI(mapCtx *direct.MapContext, in *api.Oauth2ClientCredentials) *krm.AuthConfig_OAUTH2ClientCredentials {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_OAUTH2ClientCredentials{}
	out.ClientID = direct.LazyPtr(in.ClientId)
	if in.ClientSecret != nil && in.ClientSecret.SecretVersion != "" {
		out.ClientSecretRef = &secretmanagerv1beta1.SecretRef{External: in.ClientSecret.SecretVersion}
	}
	return out
}

func AuthConfig_SSHPublicKey_ToAPI(mapCtx *direct.MapContext, in *krm.AuthConfig_SSHPublicKey) *api.SshPublicKey {
	if in == nil {
		return nil
	}
	out := &api.SshPublicKey{}
	out.Username = direct.ValueOf(in.Username)
	if in.SSHClientCertRef != nil {
		out.SshClientCert = &api.Secret{SecretVersion: in.SSHClientCertRef.External}
	}
	out.CertType = direct.ValueOf(in.CertType)
	if in.SSHClientCertPassRef != nil {
		out.SshClientCertPass = &api.Secret{SecretVersion: in.SSHClientCertPassRef.External}
	}
	return out
}

func AuthConfig_SSHPublicKey_FromAPI(mapCtx *direct.MapContext, in *api.SshPublicKey) *krm.AuthConfig_SSHPublicKey {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_SSHPublicKey{}
	out.Username = direct.LazyPtr(in.Username)
	if in.SshClientCert != nil && in.SshClientCert.SecretVersion != "" {
		out.SSHClientCertRef = &secretmanagerv1beta1.SecretRef{External: in.SshClientCert.SecretVersion}
	}
	out.CertType = direct.LazyPtr(in.CertType)
	if in.SshClientCertPass != nil && in.SshClientCertPass.SecretVersion != "" {
		out.SSHClientCertPassRef = &secretmanagerv1beta1.SecretRef{External: in.SshClientCertPass.SecretVersion}
	}
	return out
}

func LockConfig_ToAPI(mapCtx *direct.MapContext, in *krm.LockConfig) *api.LockConfig {
	if in == nil {
		return nil
	}
	out := &api.LockConfig{}
	out.Locked = direct.ValueOf(in.Locked)
	out.Reason = direct.ValueOf(in.Reason)
	return out
}

func LockConfig_FromAPI(mapCtx *direct.MapContext, in *api.LockConfig) *krm.LockConfig {
	if in == nil {
		return nil
	}
	out := &krm.LockConfig{}
	out.Locked = direct.LazyPtr(in.Locked)
	out.Reason = direct.LazyPtr(in.Reason)
	return out
}

func DestinationConfig_ToAPI(mapCtx *direct.MapContext, in *krm.DestinationConfig) *api.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &api.DestinationConfig{}
	out.Key = direct.ValueOf(in.Key)
	out.Destinations = direct.Slice_ToProto(mapCtx, in.Destinations, Destination_ToAPI)
	return out
}

func DestinationConfig_FromAPI(mapCtx *direct.MapContext, in *api.DestinationConfig) *krm.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &krm.DestinationConfig{}
	out.Key = direct.LazyPtr(in.Key)
	out.Destinations = direct.Slice_FromProto(mapCtx, in.Destinations, Destination_FromAPI)
	return out
}

func Destination_ToAPI(mapCtx *direct.MapContext, in *krm.Destination) *api.Destination {
	if in == nil {
		return nil
	}
	out := &api.Destination{}
	if in.ServiceAttachmentRef != nil {
		out.ServiceAttachment = in.ServiceAttachmentRef.External
	}
	out.Host = direct.ValueOf(in.Host)
	out.Port = int64(direct.ValueOf(in.Port))
	return out
}

func Destination_FromAPI(mapCtx *direct.MapContext, in *api.Destination) *krm.Destination {
	if in == nil {
		return nil
	}
	out := &krm.Destination{}
	if in.ServiceAttachment != "" {
		out.ServiceAttachmentRef = &refsv1beta1.ComputeServiceAttachmentRef{External: in.ServiceAttachment}
	}
	out.Host = direct.LazyPtr(in.Host)
	if in.Port != 0 {
		out.Port = direct.LazyPtr(int32(in.Port))
	}
	return out
}

func NodeConfig_ToAPI(mapCtx *direct.MapContext, in *krm.NodeConfig) *api.NodeConfig {
	if in == nil {
		return nil
	}
	out := &api.NodeConfig{}
	out.MinNodeCount = int64(direct.ValueOf(in.MinNodeCount))
	out.MaxNodeCount = int64(direct.ValueOf(in.MaxNodeCount))
	return out
}

func NodeConfig_FromAPI(mapCtx *direct.MapContext, in *api.NodeConfig) *krm.NodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig{}
	if in.MinNodeCount != 0 {
		out.MinNodeCount = direct.LazyPtr(int32(in.MinNodeCount))
	}
	if in.MaxNodeCount != 0 {
		out.MaxNodeCount = direct.LazyPtr(int32(in.MaxNodeCount))
	}
	return out
}

func SSLConfig_ToAPI(mapCtx *direct.MapContext, in *krm.SSLConfig) *api.SslConfig {
	if in == nil {
		return nil
	}
	out := &api.SslConfig{}
	out.Type = direct.ValueOf(in.Type)
	out.TrustModel = direct.ValueOf(in.TrustModel)
	if in.PrivateServerCertificateRef != nil {
		out.PrivateServerCertificate = &api.Secret{SecretVersion: in.PrivateServerCertificateRef.External}
	}
	if in.ClientCertificateRef != nil {
		out.ClientCertificate = &api.Secret{SecretVersion: in.ClientCertificateRef.External}
	}
	if in.ClientPrivateKeyRef != nil {
		out.ClientPrivateKey = &api.Secret{SecretVersion: in.ClientPrivateKeyRef.External}
	}
	if in.ClientPrivateKeyPassRef != nil {
		out.ClientPrivateKeyPass = &api.Secret{SecretVersion: in.ClientPrivateKeyPassRef.External}
	}
	out.ServerCertType = direct.ValueOf(in.ServerCertType)
	out.ClientCertType = direct.ValueOf(in.ClientCertType)
	out.UseSsl = direct.ValueOf(in.UseSSL)
	out.AdditionalVariables = direct.Slice_ToProto(mapCtx, in.AdditionalVariables, ConfigVariable_ToAPI)
	return out
}

func SSLConfig_FromAPI(mapCtx *direct.MapContext, in *api.SslConfig) *krm.SSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.SSLConfig{}
	out.Type = direct.LazyPtr(in.Type)
	out.TrustModel = direct.LazyPtr(in.TrustModel)
	if in.PrivateServerCertificate != nil && in.PrivateServerCertificate.SecretVersion != "" {
		out.PrivateServerCertificateRef = &secretmanagerv1beta1.SecretRef{External: in.PrivateServerCertificate.SecretVersion}
	}
	if in.ClientCertificate != nil && in.ClientCertificate.SecretVersion != "" {
		out.ClientCertificateRef = &secretmanagerv1beta1.SecretRef{External: in.ClientCertificate.SecretVersion}
	}
	if in.ClientPrivateKey != nil && in.ClientPrivateKey.SecretVersion != "" {
		out.ClientPrivateKeyRef = &secretmanagerv1beta1.SecretRef{External: in.ClientPrivateKey.SecretVersion}
	}
	if in.ClientPrivateKeyPass != nil && in.ClientPrivateKeyPass.SecretVersion != "" {
		out.ClientPrivateKeyPassRef = &secretmanagerv1beta1.SecretRef{External: in.ClientPrivateKeyPass.SecretVersion}
	}
	out.ServerCertType = direct.LazyPtr(in.ServerCertType)
	out.ClientCertType = direct.LazyPtr(in.ClientCertType)
	out.UseSSL = direct.LazyPtr(in.UseSsl)
	out.AdditionalVariables = direct.Slice_FromProto(mapCtx, in.AdditionalVariables, ConfigVariable_FromAPI)
	return out
}

func ConnectionStatus_FromAPI(mapCtx *direct.MapContext, in *api.ConnectionStatus) *krm.ConnectionStatus {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionStatus{}
	out.State = direct.LazyPtr(in.State)
	out.Description = direct.LazyPtr(in.Description)
	out.Status = direct.LazyPtr(in.Status)
	return out
}

func ConnectorsConnectionObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.ConnectorsConnectionObservedState) *api.Connection {
	if in == nil {
		return nil
	}
	out := &api.Connection{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.ValueOf(in.CreateTime)
	out.UpdateTime = direct.ValueOf(in.UpdateTime)
	out.Status = ConnectionStatus_ToAPI(mapCtx, in.Status)
	out.ImageLocation = direct.ValueOf(in.ImageLocation)
	out.ServiceDirectory = direct.ValueOf(in.ServiceDirectory)
	out.EnvoyImageLocation = direct.ValueOf(in.EnvoyImageLocation)
	return out
}

func ConnectionStatus_ToAPI(mapCtx *direct.MapContext, in *krm.ConnectionStatus) *api.ConnectionStatus {
	if in == nil {
		return nil
	}
	out := &api.ConnectionStatus{}
	out.State = direct.ValueOf(in.State)
	out.Description = direct.ValueOf(in.Description)
	out.Status = direct.ValueOf(in.Status)
	return out
}
