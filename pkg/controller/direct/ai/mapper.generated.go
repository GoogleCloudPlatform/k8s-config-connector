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

package ai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta/generativelanguagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AiCorpusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.AiCorpusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiCorpusObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiCorpusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiCorpusObservedState) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiCorpusSpec_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.AiCorpusSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiCorpusSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiCorpusSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiCorpusSpec) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiPermissionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Permission) *krm.AiPermissionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiPermissionObservedState{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiPermissionObservedState) *pb.Permission {
	if in == nil {
		return nil
	}
	out := &pb.Permission{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Permission) *krm.AiPermissionSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiPermissionSpec{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiPermissionSpec) *pb.Permission {
	if in == nil {
		return nil
	}
	out := &pb.Permission{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func Corpus_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.Corpus {
	if in == nil {
		return nil
	}
	out := &krm.Corpus{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Corpus_ToProto(mapCtx *direct.MapContext, in *krm.Corpus) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func CorpusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.CorpusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CorpusObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func CorpusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CorpusObservedState) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
