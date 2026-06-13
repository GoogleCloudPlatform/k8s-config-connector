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

func ComputeSSLCertificateSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SslCertificate) *krm.ComputeSSLCertificateSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSSLCertificateSpec{}
	if in.Certificate != nil {
		out.Certificate.Value = in.Certificate
	}
	out.Description = in.Description
	if in.PrivateKey != nil {
		out.PrivateKey.Value = in.PrivateKey
	}
	return out
}

func ComputeSSLCertificateSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSSLCertificateSpec) *pb.SslCertificate {
	if in == nil {
		return nil
	}
	out := &pb.SslCertificate{}
	if in.Certificate.Value != nil {
		out.Certificate = in.Certificate.Value
	}
	out.Description = in.Description
	if in.PrivateKey.Value != nil {
		out.PrivateKey = in.PrivateKey.Value
	}
	return out
}

func ComputeSSLCertificateStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SslCertificate) *krm.ComputeSSLCertificateStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSSLCertificateStatus{}
	if in.Id != nil {
		out.CertificateId = direct.LazyPtr(int64(*in.Id))
	}
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpireTime = in.ExpireTime
	out.SelfLink = in.SelfLink
	return out
}

func ComputeSSLCertificateStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSSLCertificateStatus) *pb.SslCertificate {
	if in == nil {
		return nil
	}
	out := &pb.SslCertificate{}
	if in.CertificateId != nil {
		out.Id = direct.LazyPtr(uint64(*in.CertificateId))
	}
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpireTime = in.ExpireTime
	out.SelfLink = in.SelfLink
	return out
}
