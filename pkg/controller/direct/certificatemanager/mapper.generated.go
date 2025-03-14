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

package certificatemanager

import (
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CertificateManagerDNSAuthorizationSpec_FromProto(mapCtx *direct.MapContext, in *pb.DnsAuthorization) *krm.CertificateManagerDNSAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManagerDNSAuthorizationSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Domain = in.GetDomain()
	return out
}
func CertificateManagerDNSAuthorizationSpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManagerDNSAuthorizationSpec) *pb.DnsAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.DnsAuthorization{}
	out.Description = direct.ValueOf(in.Description)
	out.Domain = in.Domain
	return out
}
func DnsAuthorization_DnsResourceRecord_FromProto(mapCtx *direct.MapContext, in *pb.DnsAuthorization_DnsResourceRecord) *krm.DnsAuthorization_DnsResourceRecord {
	if in == nil {
		return nil
	}
	out := &krm.DnsAuthorization_DnsResourceRecord{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	out.Data = direct.LazyPtr(in.GetData())
	return out
}
func DnsAuthorization_DnsResourceRecord_ToProto(mapCtx *direct.MapContext, in *krm.DnsAuthorization_DnsResourceRecord) *pb.DnsAuthorization_DnsResourceRecord {
	if in == nil {
		return nil
	}
	out := &pb.DnsAuthorization_DnsResourceRecord{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	out.Data = direct.ValueOf(in.Data)
	return out
}

func TrustConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig) *krmv1alpha1.TrustConfigSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TrustConfigSpec{}
	//out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.TrustStores = direct.Slice_FromProto(mapCtx, in.TrustStores, TrustConfig_TrustStore_FromProto)
	return out
}
func TrustConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TrustConfigSpec) *pb.TrustConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	//out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Etag = direct.ValueOf(in.Etag)
	out.TrustStores = direct.Slice_ToProto(mapCtx, in.TrustStores, TrustConfig_TrustStore_ToProto)
	return out
}
func TrustConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig) *krmv1alpha1.TrustConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TrustConfigObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Etag
	// MISSING: TrustStores
	return out
}
func TrustConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TrustConfigObservedState) *pb.TrustConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Etag
	// MISSING: TrustStores
	return out
}
func TrustConfig_IntermediateCA_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig_IntermediateCA) *krmv1alpha1.TrustConfig_IntermediateCA {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TrustConfig_IntermediateCA{}
	out.PemCertificate = direct.LazyPtr(in.GetPemCertificate())
	return out
}
func TrustConfig_IntermediateCA_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TrustConfig_IntermediateCA) *pb.TrustConfig_IntermediateCA {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig_IntermediateCA{}
	if in.PemCertificate != nil {
		out.Kind = &pb.TrustConfig_IntermediateCA_PemCertificate{PemCertificate: *in.PemCertificate}
	}
	return out
}
func TrustConfig_TrustAnchor_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig_TrustAnchor) *krmv1alpha1.TrustConfig_TrustAnchor {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TrustConfig_TrustAnchor{}
	out.PemCertificate = direct.LazyPtr(in.GetPemCertificate())
	return out
}
func TrustConfig_TrustAnchor_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TrustConfig_TrustAnchor) *pb.TrustConfig_TrustAnchor {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig_TrustAnchor{}
	if in.PemCertificate != nil {
		out.Kind = &pb.TrustConfig_TrustAnchor_PemCertificate{PemCertificate: *in.PemCertificate}
	}

	return out
}
func TrustConfig_TrustStore_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig_TrustStore) *krmv1alpha1.TrustConfig_TrustStore {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TrustConfig_TrustStore{}
	out.TrustAnchors = direct.Slice_FromProto(mapCtx, in.TrustAnchors, TrustConfig_TrustAnchor_FromProto)
	out.IntermediateCas = direct.Slice_FromProto(mapCtx, in.IntermediateCas, TrustConfig_IntermediateCA_FromProto)
	return out
}
func TrustConfig_TrustStore_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TrustConfig_TrustStore) *pb.TrustConfig_TrustStore {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig_TrustStore{}
	out.TrustAnchors = direct.Slice_ToProto(mapCtx, in.TrustAnchors, TrustConfig_TrustAnchor_ToProto)
	out.IntermediateCas = direct.Slice_ToProto(mapCtx, in.IntermediateCas, TrustConfig_IntermediateCA_ToProto)
	return out
}
