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
	krmcomputev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeManagedSSLCertificateObservedState_v1alpha1_FromProto converts pb.SslCertificate to krmcomputev1alpha1.ComputeManagedSSLCertificateObservedState.
// Handcoded due to a type mismatch: the proto defines 'id' as *uint64, but KRM represents it as *int64.
func ComputeManagedSSLCertificateObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SslCertificate) *krmcomputev1alpha1.ComputeManagedSSLCertificateObservedState {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeManagedSSLCertificateObservedState{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpireTime = in.ExpireTime
	if in.Id != nil {
		idVal := int64(*in.Id)
		out.ID = &idVal
	}
	out.SelfLink = in.SelfLink
	out.SubjectAlternativeNames = in.SubjectAlternativeNames
	return out
}

// ComputeManagedSSLCertificateObservedState_v1alpha1_ToProto converts krmcomputev1alpha1.ComputeManagedSSLCertificateObservedState to pb.SslCertificate.
// Handcoded due to a type mismatch: the proto defines 'id' as *uint64, but KRM represents it as *int64.
func ComputeManagedSSLCertificateObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeManagedSSLCertificateObservedState) *pb.SslCertificate {
	if in == nil {
		return nil
	}
	out := &pb.SslCertificate{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpireTime = in.ExpireTime
	if in.ID != nil {
		idVal := uint64(*in.ID)
		out.Id = &idVal
	}
	out.SelfLink = in.SelfLink
	out.SubjectAlternativeNames = in.SubjectAlternativeNames
	return out
}

// ComputeManagedSSLCertificateObservedState_v1beta1_FromProto converts pb.SslCertificate to krm.ComputeManagedSSLCertificateObservedState.
// Handcoded due to a type mismatch: the proto defines 'id' as *uint64, but KRM represents it as *int64.
func ComputeManagedSSLCertificateObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SslCertificate) *krm.ComputeManagedSSLCertificateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeManagedSSLCertificateObservedState{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpireTime = in.ExpireTime
	if in.Id != nil {
		idVal := int64(*in.Id)
		out.ID = &idVal
	}
	out.SelfLink = in.SelfLink
	out.SubjectAlternativeNames = in.SubjectAlternativeNames
	return out
}

// ComputeManagedSSLCertificateObservedState_v1beta1_ToProto converts krm.ComputeManagedSSLCertificateObservedState to pb.SslCertificate.
// Handcoded due to a type mismatch: the proto defines 'id' as *uint64, but KRM represents it as *int64.
func ComputeManagedSSLCertificateObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeManagedSSLCertificateObservedState) *pb.SslCertificate {
	if in == nil {
		return nil
	}
	out := &pb.SslCertificate{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpireTime = in.ExpireTime
	if in.ID != nil {
		idVal := uint64(*in.ID)
		out.Id = &idVal
	}
	out.SelfLink = in.SelfLink
	out.SubjectAlternativeNames = in.SubjectAlternativeNames
	return out
}
