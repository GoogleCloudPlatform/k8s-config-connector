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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/support/apiv2/supportpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/support/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
func Comment_FromProto(mapCtx *direct.MapContext, in *pb.Comment) *krm.Comment {
	if in == nil {
		return nil
	}
	out := &krm.Comment{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	out.Body = direct.LazyPtr(in.GetBody())
	// MISSING: PlainTextBody
	return out
}
func Comment_ToProto(mapCtx *direct.MapContext, in *krm.Comment) *pb.Comment {
	if in == nil {
		return nil
	}
	out := &pb.Comment{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	out.Body = direct.ValueOf(in.Body)
	// MISSING: PlainTextBody
	return out
}
func CommentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Comment) *krm.CommentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CommentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Creator = Actor_FromProto(mapCtx, in.GetCreator())
	// MISSING: Body
	out.PlainTextBody = direct.LazyPtr(in.GetPlainTextBody())
	return out
}
func CommentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CommentObservedState) *pb.Comment {
	if in == nil {
		return nil
	}
	out := &pb.Comment{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Creator = Actor_ToProto(mapCtx, in.Creator)
	// MISSING: Body
	out.PlainTextBody = direct.ValueOf(in.PlainTextBody)
	return out
}
func SupportCommentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Comment) *krm.SupportCommentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SupportCommentObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: Body
	// MISSING: PlainTextBody
	return out
}
func SupportCommentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SupportCommentObservedState) *pb.Comment {
	if in == nil {
		return nil
	}
	out := &pb.Comment{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: Body
	// MISSING: PlainTextBody
	return out
}
func SupportCommentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Comment) *krm.SupportCommentSpec {
	if in == nil {
		return nil
	}
	out := &krm.SupportCommentSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: Body
	// MISSING: PlainTextBody
	return out
}
func SupportCommentSpec_ToProto(mapCtx *direct.MapContext, in *krm.SupportCommentSpec) *pb.Comment {
	if in == nil {
		return nil
	}
	out := &pb.Comment{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: Body
	// MISSING: PlainTextBody
	return out
}
