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

package support

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/support/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/support/apiv2/supportpb"
)
func Actor_FromProto(mapCtx *direct.MapContext, in *pb.Actor) *krm.Actor {
	if in == nil {
		return nil
	}
	out := &krm.Actor{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Email = direct.LazyPtr(in.GetEmail())
	// MISSING: GoogleSupport
	return out
}
func Actor_ToProto(mapCtx *direct.MapContext, in *krm.Actor) *pb.Actor {
	if in == nil {
		return nil
	}
	out := &pb.Actor{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Email = direct.ValueOf(in.Email)
	// MISSING: GoogleSupport
	return out
}
func ActorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Actor) *krm.ActorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ActorObservedState{}
	// MISSING: DisplayName
	// MISSING: Email
	out.GoogleSupport = direct.LazyPtr(in.GetGoogleSupport())
	return out
}
func ActorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ActorObservedState) *pb.Actor {
	if in == nil {
		return nil
	}
	out := &pb.Actor{}
	// MISSING: DisplayName
	// MISSING: Email
	out.GoogleSupport = direct.ValueOf(in.GoogleSupport)
	return out
}
func Attachment_FromProto(mapCtx *direct.MapContext, in *pb.Attachment) *krm.Attachment {
	if in == nil {
		return nil
	}
	out := &krm.Attachment{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	out.Filename = direct.LazyPtr(in.GetFilename())
	// MISSING: MimeType
	// MISSING: SizeBytes
	return out
}
func Attachment_ToProto(mapCtx *direct.MapContext, in *krm.Attachment) *pb.Attachment {
	if in == nil {
		return nil
	}
	out := &pb.Attachment{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	out.Filename = direct.ValueOf(in.Filename)
	// MISSING: MimeType
	// MISSING: SizeBytes
	return out
}
func AttachmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Attachment) *krm.AttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AttachmentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Creator = Actor_FromProto(mapCtx, in.GetCreator())
	// MISSING: Filename
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	return out
}
func AttachmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AttachmentObservedState) *pb.Attachment {
	if in == nil {
		return nil
	}
	out := &pb.Attachment{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Creator = Actor_ToProto(mapCtx, in.Creator)
	// MISSING: Filename
	out.MimeType = direct.ValueOf(in.MimeType)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	return out
}
func SupportAttachmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Attachment) *krm.SupportAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SupportAttachmentObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: Filename
	// MISSING: MimeType
	// MISSING: SizeBytes
	return out
}
func SupportAttachmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SupportAttachmentObservedState) *pb.Attachment {
	if in == nil {
		return nil
	}
	out := &pb.Attachment{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: Filename
	// MISSING: MimeType
	// MISSING: SizeBytes
	return out
}
func SupportAttachmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Attachment) *krm.SupportAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.SupportAttachmentSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: Filename
	// MISSING: MimeType
	// MISSING: SizeBytes
	return out
}
func SupportAttachmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.SupportAttachmentSpec) *pb.Attachment {
	if in == nil {
		return nil
	}
	out := &pb.Attachment{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: Filename
	// MISSING: MimeType
	// MISSING: SizeBytes
	return out
}
