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

package redis

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/redis/cluster/apiv1beta1/clusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CertificateAuthority_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority) *krm.CertificateAuthority {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority{}
	out.ManagedServerCa = CertificateAuthority_ManagedCertificateAuthority_FromProto(mapCtx, in.GetManagedServerCa())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func CertificateAuthority_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority) *pb.CertificateAuthority {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority{}
	if oneof := CertificateAuthority_ManagedCertificateAuthority_ToProto(mapCtx, in.ManagedServerCa); oneof != nil {
		out.ServerCa = &pb.CertificateAuthority_ManagedServerCa{ManagedServerCa: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func CertificateAuthority_ManagedCertificateAuthority_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_ManagedCertificateAuthority) *krm.CertificateAuthority_ManagedCertificateAuthority {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_ManagedCertificateAuthority{}
	out.CaCerts = direct.Slice_FromProto(mapCtx, in.CaCerts, CertificateAuthority_ManagedCertificateAuthority_CertChain_FromProto)
	return out
}
func CertificateAuthority_ManagedCertificateAuthority_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_ManagedCertificateAuthority) *pb.CertificateAuthority_ManagedCertificateAuthority {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_ManagedCertificateAuthority{}
	out.CaCerts = direct.Slice_ToProto(mapCtx, in.CaCerts, CertificateAuthority_ManagedCertificateAuthority_CertChain_ToProto)
	return out
}
func CertificateAuthority_ManagedCertificateAuthority_CertChain_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_ManagedCertificateAuthority_CertChain) *krm.CertificateAuthority_ManagedCertificateAuthority_CertChain {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_ManagedCertificateAuthority_CertChain{}
	out.Certificates = in.Certificates
	return out
}
func CertificateAuthority_ManagedCertificateAuthority_CertChain_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_ManagedCertificateAuthority_CertChain) *pb.CertificateAuthority_ManagedCertificateAuthority_CertChain {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_ManagedCertificateAuthority_CertChain{}
	out.Certificates = in.Certificates
	return out
}
