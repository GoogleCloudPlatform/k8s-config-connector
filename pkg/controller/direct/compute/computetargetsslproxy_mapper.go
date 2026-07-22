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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	certificatemanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeTargetSSLProxySpec_v1beta1_FromProto converts the protobuf TargetSslProxy message to ComputeTargetSSLProxySpec KRM representation.
// This function is handcoded because the KRM struct field names deviate from default protobuf field naming (e.g., BackendServiceRef maps to Service,
// CertificateMapRef maps to CertificateMap, and SslPolicyRef maps to SslPolicy).
func ComputeTargetSSLProxySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetSslProxy) *krm.ComputeTargetSSLProxySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetSSLProxySpec{}
	out.Description = in.Description
	out.ProxyHeader = in.ProxyHeader
	if in.Service != nil {
		out.BackendServiceRef = krm.ComputeBackendServiceRef{External: direct.ValueOf(in.Service)}
	}
	if in.CertificateMap != nil {
		out.CertificateMapRef = &certificatemanagerv1beta1.CertificateManagerCertificateMapRef{External: *in.CertificateMap}
	}
	if in.SslPolicy != nil {
		out.SslPolicyRef = &krm.ComputeSSLPolicyRef{External: *in.SslPolicy}
	}
	for _, cert := range in.SslCertificates {
		out.SslCertificates = append(out.SslCertificates, krm.ComputeSSLCertificateRef{External: cert})
	}
	return out
}

// ComputeTargetSSLProxySpec_v1beta1_ToProto converts ComputeTargetSSLProxySpec KRM representation to protobuf TargetSslProxy message.
// This function is handcoded because KRM fields use different naming and structures compared to the protobuf counterpart (e.g., BackendServiceRef is mapped
// to the Service proto field, and SslCertificates utilizes custom reference objects).
func ComputeTargetSSLProxySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetSSLProxySpec) *pb.TargetSslProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetSslProxy{}
	out.Description = in.Description
	out.ProxyHeader = in.ProxyHeader
	if in.BackendServiceRef.External != "" {
		out.Service = direct.LazyPtr(in.BackendServiceRef.External)
	}
	if in.CertificateMapRef != nil {
		out.CertificateMap = direct.LazyPtr(in.CertificateMapRef.External)
	}
	if in.SslPolicyRef != nil {
		out.SslPolicy = direct.LazyPtr(in.SslPolicyRef.External)
	}
	for _, certRef := range in.SslCertificates {
		if certRef.External != "" {
			out.SslCertificates = append(out.SslCertificates, certRef.External)
		}
	}
	return out
}

// ComputeTargetSSLProxyStatus_v1beta1_FromProto converts protobuf TargetSslProxy status fields to ComputeTargetSSLProxyStatus representation.
// This function is handcoded because the ID field is mapped to ProxyId (using customized int64 type conversion).
func ComputeTargetSSLProxyStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetSslProxy) *krm.ComputeTargetSSLProxyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetSSLProxyStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ProxyId = direct.LazyPtr(int64(in.GetId()))
	out.SelfLink = in.SelfLink
	return out
}

// ComputeTargetSSLProxyStatus_v1beta1_ToProto converts ComputeTargetSSLProxyStatus representation to protobuf TargetSslProxy message.
// This function is handcoded because ProxyId requires a conversion back to the protobuf uint64 ID field.
func ComputeTargetSSLProxyStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetSSLProxyStatus) *pb.TargetSslProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetSslProxy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = direct.PtrInt64ToPtrUint64(in.ProxyId)
	out.SelfLink = in.SelfLink
	return out
}
