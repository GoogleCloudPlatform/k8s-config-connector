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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
func ServerTlsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ServerTlsPolicy) *krm.ServerTlsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ServerTlsPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.AllowOpen = direct.LazyPtr(in.GetAllowOpen())
	out.ServerCertificate = CertificateProvider_FromProto(mapCtx, in.GetServerCertificate())
	out.MtlsPolicy = ServerTlsPolicy_MTLSPolicy_FromProto(mapCtx, in.GetMtlsPolicy())
	return out
}
func ServerTlsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ServerTlsPolicy) *pb.ServerTlsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServerTlsPolicy{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.AllowOpen = direct.ValueOf(in.AllowOpen)
	out.ServerCertificate = CertificateProvider_ToProto(mapCtx, in.ServerCertificate)
	out.MtlsPolicy = ServerTlsPolicy_MTLSPolicy_ToProto(mapCtx, in.MtlsPolicy)
	return out
}
func ServerTlsPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServerTlsPolicy) *krm.ServerTlsPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServerTlsPolicyObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: AllowOpen
	// MISSING: ServerCertificate
	// MISSING: MtlsPolicy
	return out
}
func ServerTlsPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServerTlsPolicyObservedState) *pb.ServerTlsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServerTlsPolicy{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: AllowOpen
	// MISSING: ServerCertificate
	// MISSING: MtlsPolicy
	return out
}
func ServerTlsPolicy_MTLSPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ServerTlsPolicy_MTLSPolicy) *krm.ServerTlsPolicy_MTLSPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ServerTlsPolicy_MTLSPolicy{}
	out.ClientValidationCa = direct.Slice_FromProto(mapCtx, in.ClientValidationCa, ValidationCA_FromProto)
	return out
}
func ServerTlsPolicy_MTLSPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ServerTlsPolicy_MTLSPolicy) *pb.ServerTlsPolicy_MTLSPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServerTlsPolicy_MTLSPolicy{}
	out.ClientValidationCa = direct.Slice_ToProto(mapCtx, in.ClientValidationCa, ValidationCA_ToProto)
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
