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
	krmsecretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/connectors/pb"
)

func AuthConfig_UserPassword_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfig_UserPassword) *krm.AuthConfig_UserPassword {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_UserPassword{}
	out.Username = direct.LazyPtr(in.GetUsername())
	if in.GetPassword() != nil {
		out.SecretRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetPassword().GetSecretVersion()}
	}
	return out
}

func AuthConfig_UserPassword_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_UserPassword) *pb.AuthConfig_UserPassword {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfig_UserPassword{}
	out.Username = direct.ValueOf(in.Username)
	if in.SecretRef != nil {
		out.Password = &pb.Secret{SecretVersion: in.SecretRef.External}
	}
	return out
}

func AuthConfig_OAUTH2ClientCredentials_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfig_Oauth2ClientCredentials) *krm.AuthConfig_OAUTH2ClientCredentials {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_OAUTH2ClientCredentials{}
	out.ClientID = direct.LazyPtr(in.GetClientId())
	if in.GetClientSecret() != nil {
		out.ClientSecretRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetClientSecret().GetSecretVersion()}
	}
	return out
}

func AuthConfig_OAUTH2ClientCredentials_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_OAUTH2ClientCredentials) *pb.AuthConfig_Oauth2ClientCredentials {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfig_Oauth2ClientCredentials{}
	out.ClientId = direct.ValueOf(in.ClientID)
	if in.ClientSecretRef != nil {
		out.ClientSecret = &pb.Secret{SecretVersion: in.ClientSecretRef.External}
	}
	return out
}

func AuthConfig_OAUTH2JwtBearer_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfig_Oauth2JwtBearer) *krm.AuthConfig_OAUTH2JwtBearer {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_OAUTH2JwtBearer{}
	if in.GetClientKey() != nil {
		out.ClientKeyRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetClientKey().GetSecretVersion()}
	}
	out.JwtClaims = AuthConfig_OAUTH2JwtBearer_JwtClaims_FromProto(mapCtx, in.GetJwtClaims())
	return out
}

func AuthConfig_OAUTH2JwtBearer_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_OAUTH2JwtBearer) *pb.AuthConfig_Oauth2JwtBearer {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfig_Oauth2JwtBearer{}
	if in.ClientKeyRef != nil {
		out.ClientKey = &pb.Secret{SecretVersion: in.ClientKeyRef.External}
	}
	out.JwtClaims = AuthConfig_OAUTH2JwtBearer_JwtClaims_ToProto(mapCtx, in.JwtClaims)
	return out
}

func AuthConfig_SSHPublicKey_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfig_SshPublicKey) *krm.AuthConfig_SSHPublicKey {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_SSHPublicKey{}
	out.Username = direct.LazyPtr(in.GetUsername())
	if in.GetSshClientCert() != nil {
		out.SSHClientCertRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetSshClientCert().GetSecretVersion()}
	}
	out.CertType = direct.LazyPtr(in.GetCertType())
	if in.GetSshClientCertPass() != nil {
		out.SSHClientCertPassRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetSshClientCertPass().GetSecretVersion()}
	}
	return out
}

func AuthConfig_SSHPublicKey_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_SSHPublicKey) *pb.AuthConfig_SshPublicKey {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfig_SshPublicKey{}
	out.Username = direct.ValueOf(in.Username)
	if in.SSHClientCertRef != nil {
		out.SshClientCert = &pb.Secret{SecretVersion: in.SSHClientCertRef.External}
	}
	out.CertType = direct.ValueOf(in.CertType)
	if in.SSHClientCertPassRef != nil {
		out.SshClientCertPass = &pb.Secret{SecretVersion: in.SSHClientCertPassRef.External}
	}
	return out
}

func ConfigVariable_FromProto(mapCtx *direct.MapContext, in *pb.ConfigVariable) *krm.ConfigVariable {
	if in == nil {
		return nil
	}
	out := &krm.ConfigVariable{}
	out.Key = direct.LazyPtr(in.GetKey())
	switch v := in.GetValue().(type) {
	case *pb.ConfigVariable_IntValue:
		val := v.IntValue
		out.IntValue = &val
	case *pb.ConfigVariable_BoolValue:
		val := v.BoolValue
		out.BoolValue = &val
	case *pb.ConfigVariable_StringValue:
		val := v.StringValue
		out.StringValue = &val
	case *pb.ConfigVariable_SecretValue:
		if v.SecretValue != nil {
			out.SecretValueRef = &krmsecretmanagerv1beta1.SecretRef{External: v.SecretValue.GetSecretVersion()}
		}
	}
	return out
}

func ConfigVariable_ToProto(mapCtx *direct.MapContext, in *krm.ConfigVariable) *pb.ConfigVariable {
	if in == nil {
		return nil
	}
	out := &pb.ConfigVariable{}
	out.Key = direct.ValueOf(in.Key)
	if in.IntValue != nil {
		out.Value = &pb.ConfigVariable_IntValue{IntValue: *in.IntValue}
	}
	if in.BoolValue != nil {
		out.Value = &pb.ConfigVariable_BoolValue{BoolValue: *in.BoolValue}
	}
	if in.StringValue != nil {
		out.Value = &pb.ConfigVariable_StringValue{StringValue: *in.StringValue}
	}
	if in.SecretValueRef != nil {
		out.Value = &pb.ConfigVariable_SecretValue{SecretValue: &pb.Secret{SecretVersion: in.SecretValueRef.External}}
	}
	return out
}

func Destination_FromProto(mapCtx *direct.MapContext, in *pb.Destination) *krm.Destination {
	if in == nil {
		return nil
	}
	out := &krm.Destination{}
	if in.GetServiceAttachment() != "" {
		out.ServiceAttachmentRef = &refsv1beta1.ComputeServiceAttachmentRef{External: in.GetServiceAttachment()}
	}
	if in.GetHost() != "" {
		out.Host = direct.LazyPtr(in.GetHost())
	}
	out.Port = direct.LazyPtr(in.GetPort())
	return out
}

func Destination_ToProto(mapCtx *direct.MapContext, in *krm.Destination) *pb.Destination {
	if in == nil {
		return nil
	}
	out := &pb.Destination{}
	if in.ServiceAttachmentRef != nil {
		out.Destination = &pb.Destination_ServiceAttachment{ServiceAttachment: in.ServiceAttachmentRef.External}
	}
	if in.Host != nil {
		out.Destination = &pb.Destination_Host{Host: *in.Host}
	}
	out.Port = direct.ValueOf(in.Port)
	return out
}

func SSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krm.SSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.SSLConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.TrustModel = direct.Enum_FromProto(mapCtx, in.GetTrustModel())
	if in.GetPrivateServerCertificate() != nil {
		out.PrivateServerCertificateRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetPrivateServerCertificate().GetSecretVersion()}
	}
	if in.GetClientCertificate() != nil {
		out.ClientCertificateRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetClientCertificate().GetSecretVersion()}
	}
	if in.GetClientPrivateKey() != nil {
		out.ClientPrivateKeyRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetClientPrivateKey().GetSecretVersion()}
	}
	if in.GetClientPrivateKeyPass() != nil {
		out.ClientPrivateKeyPassRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetClientPrivateKeyPass().GetSecretVersion()}
	}
	out.ServerCertType = direct.Enum_FromProto(mapCtx, in.GetServerCertType())
	out.ClientCertType = direct.Enum_FromProto(mapCtx, in.GetClientCertType())
	out.UseSSL = direct.LazyPtr(in.GetUseSsl())
	out.AdditionalVariables = direct.Slice_FromProto(mapCtx, in.AdditionalVariables, ConfigVariable_FromProto)
	return out
}

func SSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.SSLConfig) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	out.Type = direct.Enum_ToProto[pb.SslType](mapCtx, in.Type)
	out.TrustModel = direct.Enum_ToProto[pb.SslConfig_TrustModel](mapCtx, in.TrustModel)
	if in.PrivateServerCertificateRef != nil {
		out.PrivateServerCertificate = &pb.Secret{SecretVersion: in.PrivateServerCertificateRef.External}
	}
	if in.ClientCertificateRef != nil {
		out.ClientCertificate = &pb.Secret{SecretVersion: in.ClientCertificateRef.External}
	}
	if in.ClientPrivateKeyRef != nil {
		out.ClientPrivateKey = &pb.Secret{SecretVersion: in.ClientPrivateKeyRef.External}
	}
	if in.ClientPrivateKeyPassRef != nil {
		out.ClientPrivateKeyPass = &pb.Secret{SecretVersion: in.ClientPrivateKeyPassRef.External}
	}
	out.ServerCertType = direct.Enum_ToProto[pb.CertType](mapCtx, in.ServerCertType)
	out.ClientCertType = direct.Enum_ToProto[pb.CertType](mapCtx, in.ClientCertType)
	out.UseSsl = direct.ValueOf(in.UseSSL)
	out.AdditionalVariables = direct.Slice_ToProto(mapCtx, in.AdditionalVariables, ConfigVariable_ToProto)
	return out
}
