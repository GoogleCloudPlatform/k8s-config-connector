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

package networksecurity

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
)
func CertificateProvider_FromProto(mapCtx *direct.MapContext, in *pb.CertificateProvider) *krm.CertificateProvider {
	if in == nil {
		return nil
	}
	out := &krm.CertificateProvider{}
	out.GrpcEndpoint = GrpcEndpoint_FromProto(mapCtx, in.GetGrpcEndpoint())
	out.CertificateProviderInstance = CertificateProviderInstance_FromProto(mapCtx, in.GetCertificateProviderInstance())
	return out
}
func CertificateProvider_ToProto(mapCtx *direct.MapContext, in *krm.CertificateProvider) *pb.CertificateProvider {
	if in == nil {
		return nil
	}
	out := &pb.CertificateProvider{}
	if oneof := GrpcEndpoint_ToProto(mapCtx, in.GrpcEndpoint); oneof != nil {
		out.Type = &pb.CertificateProvider_GrpcEndpoint{GrpcEndpoint: oneof}
	}
	if oneof := CertificateProviderInstance_ToProto(mapCtx, in.CertificateProviderInstance); oneof != nil {
		out.Type = &pb.CertificateProvider_CertificateProviderInstance{CertificateProviderInstance: oneof}
	}
	return out
}
func CertificateProviderInstance_FromProto(mapCtx *direct.MapContext, in *pb.CertificateProviderInstance) *krm.CertificateProviderInstance {
	if in == nil {
		return nil
	}
	out := &krm.CertificateProviderInstance{}
	out.PluginInstance = direct.LazyPtr(in.GetPluginInstance())
	return out
}
func CertificateProviderInstance_ToProto(mapCtx *direct.MapContext, in *krm.CertificateProviderInstance) *pb.CertificateProviderInstance {
	if in == nil {
		return nil
	}
	out := &pb.CertificateProviderInstance{}
	out.PluginInstance = direct.ValueOf(in.PluginInstance)
	return out
}
func ClientTlsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ClientTlsPolicy) *krm.ClientTlsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ClientTlsPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Sni = direct.LazyPtr(in.GetSni())
	out.ClientCertificate = CertificateProvider_FromProto(mapCtx, in.GetClientCertificate())
	out.ServerValidationCa = direct.Slice_FromProto(mapCtx, in.ServerValidationCa, ValidationCA_FromProto)
	return out
}
func ClientTlsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ClientTlsPolicy) *pb.ClientTlsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ClientTlsPolicy{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Sni = direct.ValueOf(in.Sni)
	out.ClientCertificate = CertificateProvider_ToProto(mapCtx, in.ClientCertificate)
	out.ServerValidationCa = direct.Slice_ToProto(mapCtx, in.ServerValidationCa, ValidationCA_ToProto)
	return out
}
func ClientTlsPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ClientTlsPolicy) *krm.ClientTlsPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClientTlsPolicyObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Sni
	// MISSING: ClientCertificate
	// MISSING: ServerValidationCa
	return out
}
func ClientTlsPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClientTlsPolicyObservedState) *pb.ClientTlsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ClientTlsPolicy{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Sni
	// MISSING: ClientCertificate
	// MISSING: ServerValidationCa
	return out
}
func GrpcEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.GrpcEndpoint) *krm.GrpcEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.GrpcEndpoint{}
	out.TargetURI = direct.LazyPtr(in.GetTargetUri())
	return out
}
func GrpcEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.GrpcEndpoint) *pb.GrpcEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.GrpcEndpoint{}
	out.TargetUri = direct.ValueOf(in.TargetURI)
	return out
}
func ValidationCA_FromProto(mapCtx *direct.MapContext, in *pb.ValidationCA) *krm.ValidationCA {
	if in == nil {
		return nil
	}
	out := &krm.ValidationCA{}
	out.GrpcEndpoint = GrpcEndpoint_FromProto(mapCtx, in.GetGrpcEndpoint())
	out.CertificateProviderInstance = CertificateProviderInstance_FromProto(mapCtx, in.GetCertificateProviderInstance())
	return out
}
func ValidationCA_ToProto(mapCtx *direct.MapContext, in *krm.ValidationCA) *pb.ValidationCA {
	if in == nil {
		return nil
	}
	out := &pb.ValidationCA{}
	if oneof := GrpcEndpoint_ToProto(mapCtx, in.GrpcEndpoint); oneof != nil {
		out.Type = &pb.ValidationCA_GrpcEndpoint{GrpcEndpoint: oneof}
	}
	if oneof := CertificateProviderInstance_ToProto(mapCtx, in.CertificateProviderInstance); oneof != nil {
		out.Type = &pb.ValidationCA_CertificateProviderInstance{CertificateProviderInstance: oneof}
	}
	return out
}
