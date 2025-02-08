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

package iot

import (
	pb "cloud.google.com/go/iot/apiv1/iotpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iot/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func DeviceRegistry_FromProto(mapCtx *direct.MapContext, in *pb.DeviceRegistry) *krm.DeviceRegistry {
	if in == nil {
		return nil
	}
	out := &krm.DeviceRegistry{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Name = direct.LazyPtr(in.GetName())
	out.EventNotificationConfigs = direct.Slice_FromProto(mapCtx, in.EventNotificationConfigs, EventNotificationConfig_FromProto)
	out.StateNotificationConfig = StateNotificationConfig_FromProto(mapCtx, in.GetStateNotificationConfig())
	out.MqttConfig = MqttConfig_FromProto(mapCtx, in.GetMqttConfig())
	out.HTTPConfig = HttpConfig_FromProto(mapCtx, in.GetHttpConfig())
	out.LogLevel = direct.Enum_FromProto(mapCtx, in.GetLogLevel())
	out.Credentials = direct.Slice_FromProto(mapCtx, in.Credentials, RegistryCredential_FromProto)
	return out
}
func DeviceRegistry_ToProto(mapCtx *direct.MapContext, in *krm.DeviceRegistry) *pb.DeviceRegistry {
	if in == nil {
		return nil
	}
	out := &pb.DeviceRegistry{}
	out.Id = direct.ValueOf(in.ID)
	out.Name = direct.ValueOf(in.Name)
	out.EventNotificationConfigs = direct.Slice_ToProto(mapCtx, in.EventNotificationConfigs, EventNotificationConfig_ToProto)
	out.StateNotificationConfig = StateNotificationConfig_ToProto(mapCtx, in.StateNotificationConfig)
	out.MqttConfig = MqttConfig_ToProto(mapCtx, in.MqttConfig)
	out.HttpConfig = HttpConfig_ToProto(mapCtx, in.HTTPConfig)
	out.LogLevel = direct.Enum_ToProto[pb.LogLevel](mapCtx, in.LogLevel)
	out.Credentials = direct.Slice_ToProto(mapCtx, in.Credentials, RegistryCredential_ToProto)
	return out
}
func EventNotificationConfig_FromProto(mapCtx *direct.MapContext, in *pb.EventNotificationConfig) *krm.EventNotificationConfig {
	if in == nil {
		return nil
	}
	out := &krm.EventNotificationConfig{}
	out.SubfolderMatches = direct.LazyPtr(in.GetSubfolderMatches())
	out.PubsubTopicName = direct.LazyPtr(in.GetPubsubTopicName())
	return out
}
func EventNotificationConfig_ToProto(mapCtx *direct.MapContext, in *krm.EventNotificationConfig) *pb.EventNotificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.EventNotificationConfig{}
	out.SubfolderMatches = direct.ValueOf(in.SubfolderMatches)
	out.PubsubTopicName = direct.ValueOf(in.PubsubTopicName)
	return out
}
func HttpConfig_FromProto(mapCtx *direct.MapContext, in *pb.HttpConfig) *krm.HttpConfig {
	if in == nil {
		return nil
	}
	out := &krm.HttpConfig{}
	out.HTTPEnabledState = direct.Enum_FromProto(mapCtx, in.GetHttpEnabledState())
	return out
}
func HttpConfig_ToProto(mapCtx *direct.MapContext, in *krm.HttpConfig) *pb.HttpConfig {
	if in == nil {
		return nil
	}
	out := &pb.HttpConfig{}
	out.HttpEnabledState = direct.Enum_ToProto[pb.HttpState](mapCtx, in.HTTPEnabledState)
	return out
}
func IotDeviceRegistryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeviceRegistry) *krm.IotDeviceRegistryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IotDeviceRegistryObservedState{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: EventNotificationConfigs
	// MISSING: StateNotificationConfig
	// MISSING: MqttConfig
	// MISSING: HTTPConfig
	// MISSING: LogLevel
	// MISSING: Credentials
	return out
}
func IotDeviceRegistryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IotDeviceRegistryObservedState) *pb.DeviceRegistry {
	if in == nil {
		return nil
	}
	out := &pb.DeviceRegistry{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: EventNotificationConfigs
	// MISSING: StateNotificationConfig
	// MISSING: MqttConfig
	// MISSING: HTTPConfig
	// MISSING: LogLevel
	// MISSING: Credentials
	return out
}
func IotDeviceRegistrySpec_FromProto(mapCtx *direct.MapContext, in *pb.DeviceRegistry) *krm.IotDeviceRegistrySpec {
	if in == nil {
		return nil
	}
	out := &krm.IotDeviceRegistrySpec{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: EventNotificationConfigs
	// MISSING: StateNotificationConfig
	// MISSING: MqttConfig
	// MISSING: HTTPConfig
	// MISSING: LogLevel
	// MISSING: Credentials
	return out
}
func IotDeviceRegistrySpec_ToProto(mapCtx *direct.MapContext, in *krm.IotDeviceRegistrySpec) *pb.DeviceRegistry {
	if in == nil {
		return nil
	}
	out := &pb.DeviceRegistry{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: EventNotificationConfigs
	// MISSING: StateNotificationConfig
	// MISSING: MqttConfig
	// MISSING: HTTPConfig
	// MISSING: LogLevel
	// MISSING: Credentials
	return out
}
func MqttConfig_FromProto(mapCtx *direct.MapContext, in *pb.MqttConfig) *krm.MqttConfig {
	if in == nil {
		return nil
	}
	out := &krm.MqttConfig{}
	out.MqttEnabledState = direct.Enum_FromProto(mapCtx, in.GetMqttEnabledState())
	return out
}
func MqttConfig_ToProto(mapCtx *direct.MapContext, in *krm.MqttConfig) *pb.MqttConfig {
	if in == nil {
		return nil
	}
	out := &pb.MqttConfig{}
	out.MqttEnabledState = direct.Enum_ToProto[pb.MqttState](mapCtx, in.MqttEnabledState)
	return out
}
func PublicKeyCertificate_FromProto(mapCtx *direct.MapContext, in *pb.PublicKeyCertificate) *krm.PublicKeyCertificate {
	if in == nil {
		return nil
	}
	out := &krm.PublicKeyCertificate{}
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	out.Certificate = direct.LazyPtr(in.GetCertificate())
	out.X509Details = X509CertificateDetails_FromProto(mapCtx, in.GetX509Details())
	return out
}
func PublicKeyCertificate_ToProto(mapCtx *direct.MapContext, in *krm.PublicKeyCertificate) *pb.PublicKeyCertificate {
	if in == nil {
		return nil
	}
	out := &pb.PublicKeyCertificate{}
	out.Format = direct.Enum_ToProto[pb.PublicKeyCertificateFormat](mapCtx, in.Format)
	out.Certificate = direct.ValueOf(in.Certificate)
	out.X509Details = X509CertificateDetails_ToProto(mapCtx, in.X509Details)
	return out
}
func RegistryCredential_FromProto(mapCtx *direct.MapContext, in *pb.RegistryCredential) *krm.RegistryCredential {
	if in == nil {
		return nil
	}
	out := &krm.RegistryCredential{}
	out.PublicKeyCertificate = PublicKeyCertificate_FromProto(mapCtx, in.GetPublicKeyCertificate())
	return out
}
func RegistryCredential_ToProto(mapCtx *direct.MapContext, in *krm.RegistryCredential) *pb.RegistryCredential {
	if in == nil {
		return nil
	}
	out := &pb.RegistryCredential{}
	if oneof := PublicKeyCertificate_ToProto(mapCtx, in.PublicKeyCertificate); oneof != nil {
		out.Credential = &pb.RegistryCredential_PublicKeyCertificate{PublicKeyCertificate: oneof}
	}
	return out
}
func StateNotificationConfig_FromProto(mapCtx *direct.MapContext, in *pb.StateNotificationConfig) *krm.StateNotificationConfig {
	if in == nil {
		return nil
	}
	out := &krm.StateNotificationConfig{}
	out.PubsubTopicName = direct.LazyPtr(in.GetPubsubTopicName())
	return out
}
func StateNotificationConfig_ToProto(mapCtx *direct.MapContext, in *krm.StateNotificationConfig) *pb.StateNotificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.StateNotificationConfig{}
	out.PubsubTopicName = direct.ValueOf(in.PubsubTopicName)
	return out
}
func X509CertificateDetails_FromProto(mapCtx *direct.MapContext, in *pb.X509CertificateDetails) *krm.X509CertificateDetails {
	if in == nil {
		return nil
	}
	out := &krm.X509CertificateDetails{}
	out.Issuer = direct.LazyPtr(in.GetIssuer())
	out.Subject = direct.LazyPtr(in.GetSubject())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.ExpiryTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpiryTime())
	out.SignatureAlgorithm = direct.LazyPtr(in.GetSignatureAlgorithm())
	out.PublicKeyType = direct.LazyPtr(in.GetPublicKeyType())
	return out
}
func X509CertificateDetails_ToProto(mapCtx *direct.MapContext, in *krm.X509CertificateDetails) *pb.X509CertificateDetails {
	if in == nil {
		return nil
	}
	out := &pb.X509CertificateDetails{}
	out.Issuer = direct.ValueOf(in.Issuer)
	out.Subject = direct.ValueOf(in.Subject)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.ExpiryTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpiryTime)
	out.SignatureAlgorithm = direct.ValueOf(in.SignatureAlgorithm)
	out.PublicKeyType = direct.ValueOf(in.PublicKeyType)
	return out
}
