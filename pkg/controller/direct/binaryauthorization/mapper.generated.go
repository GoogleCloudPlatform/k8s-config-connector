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

package binaryauthorization

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/binaryauthorization/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Attestor_FromProto(mapCtx *direct.MapContext, in *pb.Attestor) *krm.Attestor {
	if in == nil {
		return nil
	}
	out := &krm.Attestor{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.UserOwnedDrydockNote = UserOwnedDrydockNote_FromProto(mapCtx, in.GetUserOwnedDrydockNote())
	// MISSING: UpdateTime
	return out
}
func Attestor_ToProto(mapCtx *direct.MapContext, in *krm.Attestor) *pb.Attestor {
	if in == nil {
		return nil
	}
	out := &pb.Attestor{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	if oneof := UserOwnedDrydockNote_ToProto(mapCtx, in.UserOwnedDrydockNote); oneof != nil {
		out.AttestorType = &pb.Attestor_UserOwnedDrydockNote{UserOwnedDrydockNote: oneof}
	}
	// MISSING: UpdateTime
	return out
}
func AttestorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Attestor) *krm.AttestorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AttestorObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.UserOwnedDrydockNote = UserOwnedDrydockNoteObservedState_FromProto(mapCtx, in.GetUserOwnedDrydockNote())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func AttestorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AttestorObservedState) *pb.Attestor {
	if in == nil {
		return nil
	}
	out := &pb.Attestor{}
	// MISSING: Name
	// MISSING: Description
	if oneof := UserOwnedDrydockNoteObservedState_ToProto(mapCtx, in.UserOwnedDrydockNote); oneof != nil {
		out.AttestorType = &pb.Attestor_UserOwnedDrydockNote{UserOwnedDrydockNote: oneof}
	}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func AttestorPublicKey_FromProto(mapCtx *direct.MapContext, in *pb.AttestorPublicKey) *krm.AttestorPublicKey {
	if in == nil {
		return nil
	}
	out := &krm.AttestorPublicKey{}
	out.Comment = direct.LazyPtr(in.GetComment())
	out.ID = direct.LazyPtr(in.GetId())
	out.AsciiArmoredPgpPublicKey = direct.LazyPtr(in.GetAsciiArmoredPgpPublicKey())
	out.PkixPublicKey = PkixPublicKey_FromProto(mapCtx, in.GetPkixPublicKey())
	return out
}
func AttestorPublicKey_ToProto(mapCtx *direct.MapContext, in *krm.AttestorPublicKey) *pb.AttestorPublicKey {
	if in == nil {
		return nil
	}
	out := &pb.AttestorPublicKey{}
	out.Comment = direct.ValueOf(in.Comment)
	out.Id = direct.ValueOf(in.ID)
	if oneof := AttestorPublicKey_AsciiArmoredPgpPublicKey_ToProto(mapCtx, in.AsciiArmoredPgpPublicKey); oneof != nil {
		out.PublicKey = oneof
	}
	if oneof := PkixPublicKey_ToProto(mapCtx, in.PkixPublicKey); oneof != nil {
		out.PublicKey = &pb.AttestorPublicKey_PkixPublicKey{PkixPublicKey: oneof}
	}
	return out
}
func PkixPublicKey_FromProto(mapCtx *direct.MapContext, in *pb.PkixPublicKey) *krm.PkixPublicKey {
	if in == nil {
		return nil
	}
	out := &krm.PkixPublicKey{}
	out.PublicKeyPem = direct.LazyPtr(in.GetPublicKeyPem())
	out.SignatureAlgorithm = direct.Enum_FromProto(mapCtx, in.GetSignatureAlgorithm())
	return out
}
func PkixPublicKey_ToProto(mapCtx *direct.MapContext, in *krm.PkixPublicKey) *pb.PkixPublicKey {
	if in == nil {
		return nil
	}
	out := &pb.PkixPublicKey{}
	out.PublicKeyPem = direct.ValueOf(in.PublicKeyPem)
	out.SignatureAlgorithm = direct.Enum_ToProto[pb.PkixPublicKey_SignatureAlgorithm](mapCtx, in.SignatureAlgorithm)
	return out
}
func UserOwnedDrydockNote_FromProto(mapCtx *direct.MapContext, in *pb.UserOwnedDrydockNote) *krm.UserOwnedDrydockNote {
	if in == nil {
		return nil
	}
	out := &krm.UserOwnedDrydockNote{}
	out.NoteReference = direct.LazyPtr(in.GetNoteReference())
	out.PublicKeys = direct.Slice_FromProto(mapCtx, in.PublicKeys, AttestorPublicKey_FromProto)
	// MISSING: DelegationServiceAccountEmail
	return out
}
func UserOwnedDrydockNote_ToProto(mapCtx *direct.MapContext, in *krm.UserOwnedDrydockNote) *pb.UserOwnedDrydockNote {
	if in == nil {
		return nil
	}
	out := &pb.UserOwnedDrydockNote{}
	out.NoteReference = direct.ValueOf(in.NoteReference)
	out.PublicKeys = direct.Slice_ToProto(mapCtx, in.PublicKeys, AttestorPublicKey_ToProto)
	// MISSING: DelegationServiceAccountEmail
	return out
}
func UserOwnedDrydockNoteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UserOwnedDrydockNote) *krm.UserOwnedDrydockNoteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UserOwnedDrydockNoteObservedState{}
	// MISSING: NoteReference
	// MISSING: PublicKeys
	out.DelegationServiceAccountEmail = direct.LazyPtr(in.GetDelegationServiceAccountEmail())
	return out
}
func UserOwnedDrydockNoteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UserOwnedDrydockNoteObservedState) *pb.UserOwnedDrydockNote {
	if in == nil {
		return nil
	}
	out := &pb.UserOwnedDrydockNote{}
	// MISSING: NoteReference
	// MISSING: PublicKeys
	out.DelegationServiceAccountEmail = direct.ValueOf(in.DelegationServiceAccountEmail)
	return out
}
