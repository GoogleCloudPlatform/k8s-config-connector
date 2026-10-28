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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeManagedSSLCertificateObservedState_v1beta1_FromProto maps a protobuf SslCertificate to its KRM ComputeManagedSSLCertificateObservedState representation.
// This is handcoded because CertificateID in KRM is an *int (OpenAPI integer without format)
// while Id in protobuf is a *uint64.
func ComputeManagedSSLCertificateObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SslCertificate) *krm.ComputeManagedSSLCertificateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeManagedSSLCertificateObservedState{}
	if in.Id != nil {
		id := int(*in.Id)
		out.CertificateID = &id
	}
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpireTime = in.ExpireTime
	out.SelfLink = in.SelfLink
	out.SubjectAlternativeNames = in.SubjectAlternativeNames
	return out
}

// ComputeManagedSSLCertificateObservedState_v1beta1_ToProto maps a KRM ComputeManagedSSLCertificateObservedState to its protobuf representation.
// This is handcoded because CertificateID in KRM is an *int (OpenAPI integer without format)
// while Id in protobuf is a *uint64.
func ComputeManagedSSLCertificateObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeManagedSSLCertificateObservedState) *pb.SslCertificate {
	if in == nil {
		return nil
	}
	out := &pb.SslCertificate{}
	if in.CertificateID != nil {
		id := uint64(*in.CertificateID)
		out.Id = &id
	}
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpireTime = in.ExpireTime
	out.SelfLink = in.SelfLink
	out.SubjectAlternativeNames = in.SubjectAlternativeNames
	return out
}

// ComputeManagedSSLCertificateSpec_v1beta1_FromProto maps a protobuf SslCertificate to its KRM ComputeManagedSSLCertificateSpec representation.
// This is handcoded to map the protobuf Name field to KRM ResourceID.
func ComputeManagedSSLCertificateSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SslCertificate) *krm.ComputeManagedSSLCertificateSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeManagedSSLCertificateSpec{}
	out.Description = in.Description
	out.Managed = ComputeManagedSSLCertificateManaged_v1beta1_FromProto(mapCtx, in.GetManaged())
	out.Type = in.Type
	out.ResourceID = in.Name
	return out
}

// ComputeManagedSSLCertificateSpec_v1beta1_ToProto maps a KRM ComputeManagedSSLCertificateSpec to its protobuf representation.
// This is handcoded to map the KRM ResourceID field to protobuf Name.
func ComputeManagedSSLCertificateSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeManagedSSLCertificateSpec) *pb.SslCertificate {
	if in == nil {
		return nil
	}
	out := &pb.SslCertificate{}
	out.Description = in.Description
	out.Managed = ComputeManagedSSLCertificateManaged_v1beta1_ToProto(mapCtx, in.Managed)
	out.Type = in.Type
	out.Name = in.ResourceID
	return out
}
