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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CertificateManagerDNSAuthorizationSpec_FromProto(mapCtx *direct.MapContext, in *pb.DnsAuthorization) *krm.CertificateManagerDNSAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManagerDNSAuthorizationSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Domain = direct.LazyPtr(in.GetDomain())
	// MISSING: DNSResourceRecord
	// MISSING: Type
	return out
}
func CertificateManagerDNSAuthorizationSpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManagerDNSAuthorizationSpec) *pb.DnsAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.DnsAuthorization{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.Description = direct.ValueOf(in.Description)
	out.Domain = CertificateManagerDNSAuthorizationSpec_Domain_ToProto(mapCtx, in.Domain)
	// MISSING: DNSResourceRecord
	// MISSING: Type
	return out
}
func CertificateManagerTrustConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig) *krm.CertificateManagerTrustConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManagerTrustConfigObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Etag
	// MISSING: TrustStores
	return out
}
func CertificateManagerTrustConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManagerTrustConfigObservedState) *pb.TrustConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Etag
	// MISSING: TrustStores
	return out
}
func CertificateManagerTrustConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig) *krm.CertificateManagerTrustConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManagerTrustConfigSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Etag
	// MISSING: TrustStores
	return out
}
func CertificateManagerTrustConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManagerTrustConfigSpec) *pb.TrustConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Etag
	// MISSING: TrustStores
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
func TrustConfig_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig) *krm.TrustConfig {
	if in == nil {
		return nil
	}
	out := &krm.TrustConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.TrustStores = direct.Slice_FromProto(mapCtx, in.TrustStores, TrustConfig_TrustStore_FromProto)
	return out
}
func TrustConfig_ToProto(mapCtx *direct.MapContext, in *krm.TrustConfig) *pb.TrustConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Etag = direct.ValueOf(in.Etag)
	out.TrustStores = direct.Slice_ToProto(mapCtx, in.TrustStores, TrustConfig_TrustStore_ToProto)
	return out
}
func TrustConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig) *krm.TrustConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TrustConfigObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Etag
	// MISSING: TrustStores
	return out
}
func TrustConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TrustConfigObservedState) *pb.TrustConfig {
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
func TrustConfig_IntermediateCA_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig_IntermediateCA) *krm.TrustConfig_IntermediateCA {
	if in == nil {
		return nil
	}
	out := &krm.TrustConfig_IntermediateCA{}
	out.PemCertificate = direct.LazyPtr(in.GetPemCertificate())
	return out
}
func TrustConfig_IntermediateCA_ToProto(mapCtx *direct.MapContext, in *krm.TrustConfig_IntermediateCA) *pb.TrustConfig_IntermediateCA {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig_IntermediateCA{}
	if oneof := TrustConfig_IntermediateCA_PemCertificate_ToProto(mapCtx, in.PemCertificate); oneof != nil {
		out.Kind = oneof
	}
	return out
}
func TrustConfig_TrustAnchor_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig_TrustAnchor) *krm.TrustConfig_TrustAnchor {
	if in == nil {
		return nil
	}
	out := &krm.TrustConfig_TrustAnchor{}
	out.PemCertificate = direct.LazyPtr(in.GetPemCertificate())
	return out
}
func TrustConfig_TrustAnchor_ToProto(mapCtx *direct.MapContext, in *krm.TrustConfig_TrustAnchor) *pb.TrustConfig_TrustAnchor {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig_TrustAnchor{}
	if oneof := TrustConfig_TrustAnchor_PemCertificate_ToProto(mapCtx, in.PemCertificate); oneof != nil {
		out.Kind = oneof
	}
	return out
}
func TrustConfig_TrustStore_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig_TrustStore) *krm.TrustConfig_TrustStore {
	if in == nil {
		return nil
	}
	out := &krm.TrustConfig_TrustStore{}
	out.TrustAnchors = direct.Slice_FromProto(mapCtx, in.TrustAnchors, TrustConfig_TrustAnchor_FromProto)
	out.IntermediateCas = direct.Slice_FromProto(mapCtx, in.IntermediateCas, TrustConfig_IntermediateCA_FromProto)
	return out
}
func TrustConfig_TrustStore_ToProto(mapCtx *direct.MapContext, in *krm.TrustConfig_TrustStore) *pb.TrustConfig_TrustStore {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig_TrustStore{}
	out.TrustAnchors = direct.Slice_ToProto(mapCtx, in.TrustAnchors, TrustConfig_TrustAnchor_ToProto)
	out.IntermediateCas = direct.Slice_ToProto(mapCtx, in.IntermediateCas, TrustConfig_IntermediateCA_ToProto)
	return out
}
