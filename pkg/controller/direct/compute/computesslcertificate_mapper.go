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
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeSSLCertificateSpec_v1beta1_ToProto maps a KRM ComputeSSLCertificateSpec to its protobuf representation.
// This is hand-coded because Certificate and PrivateKey fields are modeled as SecretKeyRefs in KRM
// but are represented as string pointers in the Compute API protobuf.
func ComputeSSLCertificateSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSSLCertificateSpec) *pb.SslCertificate {
	if in == nil {
		return nil
	}
	out := &pb.SslCertificate{}
	out.Description = in.Description
	if in.Certificate.Value != nil {
		out.Certificate = in.Certificate.Value
	}
	if in.PrivateKey.Value != nil {
		out.PrivateKey = in.PrivateKey.Value
	}
	if in.ResourceID != nil {
		out.Name = in.ResourceID
	}
	return out
}

// ComputeSSLCertificateSpec_v1beta1_FromProto maps a protobuf SslCertificate to its KRM ComputeSSLCertificateSpec representation.
// This is hand-coded because Certificate and PrivateKey fields are modeled as SecretKeyRefs in KRM
// but are represented as string pointers in the Compute API protobuf.
func ComputeSSLCertificateSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SslCertificate) *krm.ComputeSSLCertificateSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSSLCertificateSpec{}
	out.Description = in.Description
	if in.Certificate != nil {
		out.Certificate.Value = in.Certificate
	}
	if in.PrivateKey != nil {
		out.PrivateKey.Value = in.PrivateKey
	}
	if in.Name != nil {
		out.ResourceID = in.Name
	}
	if in.Region != nil && *in.Region != "" {
		parts := strings.Split(*in.Region, "/")
		out.Location = parts[len(parts)-1]
	} else {
		out.Location = "global"
	}
	return out
}

// ComputeSSLCertificateStatus_v1beta1_FromProto maps a protobuf SslCertificate to its KRM ComputeSSLCertificateStatus representation.
func ComputeSSLCertificateStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SslCertificate) *krm.ComputeSSLCertificateStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSSLCertificateStatus{}
	if in.Id != nil {
		id := int64(*in.Id)
		out.CertificateId = &id
	}
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpireTime = in.ExpireTime
	out.SelfLink = in.SelfLink
	return out
}

// ComputeSSLCertificateStatus_v1beta1_ToProto maps a KRM ComputeSSLCertificateStatus to its protobuf representation.
func ComputeSSLCertificateStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSSLCertificateStatus) *pb.SslCertificate {
	if in == nil {
		return nil
	}
	out := &pb.SslCertificate{}
	if in.CertificateId != nil {
		id := uint64(*in.CertificateId)
		out.Id = &id
	}
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpireTime = in.ExpireTime
	out.SelfLink = in.SelfLink
	return out
}
