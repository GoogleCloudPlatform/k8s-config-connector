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

package security

import (
	pb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/security/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func CertificateRevocationList_FromProto(mapCtx *direct.MapContext, in *pb.CertificateRevocationList) *krm.CertificateRevocationList {
	if in == nil {
		return nil
	}
	out := &krm.CertificateRevocationList{}
	// MISSING: Name
	// MISSING: SequenceNumber
	// MISSING: RevokedCertificates
	// MISSING: PemCrl
	// MISSING: AccessURL
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RevisionID
	out.Labels = in.Labels
	return out
}
func CertificateRevocationList_ToProto(mapCtx *direct.MapContext, in *krm.CertificateRevocationList) *pb.CertificateRevocationList {
	if in == nil {
		return nil
	}
	out := &pb.CertificateRevocationList{}
	// MISSING: Name
	// MISSING: SequenceNumber
	// MISSING: RevokedCertificates
	// MISSING: PemCrl
	// MISSING: AccessURL
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RevisionID
	out.Labels = in.Labels
	return out
}
func CertificateRevocationListObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateRevocationList) *krm.CertificateRevocationListObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CertificateRevocationListObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SequenceNumber = direct.LazyPtr(in.GetSequenceNumber())
	out.RevokedCertificates = direct.Slice_FromProto(mapCtx, in.RevokedCertificates, CertificateRevocationList_RevokedCertificate_FromProto)
	out.PemCrl = direct.LazyPtr(in.GetPemCrl())
	out.AccessURL = direct.LazyPtr(in.GetAccessUrl())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	// MISSING: Labels
	return out
}
func CertificateRevocationListObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CertificateRevocationListObservedState) *pb.CertificateRevocationList {
	if in == nil {
		return nil
	}
	out := &pb.CertificateRevocationList{}
	out.Name = direct.ValueOf(in.Name)
	out.SequenceNumber = direct.ValueOf(in.SequenceNumber)
	out.RevokedCertificates = direct.Slice_ToProto(mapCtx, in.RevokedCertificates, CertificateRevocationList_RevokedCertificate_ToProto)
	out.PemCrl = direct.ValueOf(in.PemCrl)
	out.AccessUrl = direct.ValueOf(in.AccessURL)
	out.State = direct.Enum_ToProto[pb.CertificateRevocationList_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.RevisionId = direct.ValueOf(in.RevisionID)
	// MISSING: Labels
	return out
}
func CertificateRevocationList_RevokedCertificate_FromProto(mapCtx *direct.MapContext, in *pb.CertificateRevocationList_RevokedCertificate) *krm.CertificateRevocationList_RevokedCertificate {
	if in == nil {
		return nil
	}
	out := &krm.CertificateRevocationList_RevokedCertificate{}
	out.Certificate = direct.LazyPtr(in.GetCertificate())
	out.HexSerialNumber = direct.LazyPtr(in.GetHexSerialNumber())
	out.RevocationReason = direct.Enum_FromProto(mapCtx, in.GetRevocationReason())
	return out
}
func CertificateRevocationList_RevokedCertificate_ToProto(mapCtx *direct.MapContext, in *krm.CertificateRevocationList_RevokedCertificate) *pb.CertificateRevocationList_RevokedCertificate {
	if in == nil {
		return nil
	}
	out := &pb.CertificateRevocationList_RevokedCertificate{}
	out.Certificate = direct.ValueOf(in.Certificate)
	out.HexSerialNumber = direct.ValueOf(in.HexSerialNumber)
	out.RevocationReason = direct.Enum_ToProto[pb.RevocationReason](mapCtx, in.RevocationReason)
	return out
}
func SecurityCertificateRevocationListObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CertificateRevocationList) *krm.SecurityCertificateRevocationListObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecurityCertificateRevocationListObservedState{}
	// MISSING: Name
	// MISSING: SequenceNumber
	// MISSING: RevokedCertificates
	// MISSING: PemCrl
	// MISSING: AccessURL
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Labels
	return out
}
func SecurityCertificateRevocationListObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecurityCertificateRevocationListObservedState) *pb.CertificateRevocationList {
	if in == nil {
		return nil
	}
	out := &pb.CertificateRevocationList{}
	// MISSING: Name
	// MISSING: SequenceNumber
	// MISSING: RevokedCertificates
	// MISSING: PemCrl
	// MISSING: AccessURL
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Labels
	return out
}
func SecurityCertificateRevocationListSpec_FromProto(mapCtx *direct.MapContext, in *pb.CertificateRevocationList) *krm.SecurityCertificateRevocationListSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecurityCertificateRevocationListSpec{}
	// MISSING: Name
	// MISSING: SequenceNumber
	// MISSING: RevokedCertificates
	// MISSING: PemCrl
	// MISSING: AccessURL
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Labels
	return out
}
func SecurityCertificateRevocationListSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecurityCertificateRevocationListSpec) *pb.CertificateRevocationList {
	if in == nil {
		return nil
	}
	out := &pb.CertificateRevocationList{}
	// MISSING: Name
	// MISSING: SequenceNumber
	// MISSING: RevokedCertificates
	// MISSING: PemCrl
	// MISSING: AccessURL
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Labels
	return out
}
