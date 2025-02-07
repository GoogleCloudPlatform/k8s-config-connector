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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta/generativelanguagepb"
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
func AiDocumentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.AiDocumentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiDocumentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiDocumentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiDocumentObservedState) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiDocumentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.AiDocumentSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiDocumentSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiDocumentSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiDocumentSpec) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
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
func Document_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.Document {
	if in == nil {
		return nil
	}
	out := &krm.Document{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.CustomMetadata = direct.Slice_FromProto(mapCtx, in.CustomMetadata, CustomMetadata_FromProto)
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func Document_ToProto(mapCtx *direct.MapContext, in *krm.Document) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.CustomMetadata = direct.Slice_ToProto(mapCtx, in.CustomMetadata, CustomMetadata_ToProto)
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func DocumentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.DocumentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func DocumentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentObservedState) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func StringList_FromProto(mapCtx *direct.MapContext, in *pb.StringList) *krm.StringList {
	if in == nil {
		return nil
	}
	out := &krm.StringList{}
	out.Values = in.Values
	return out
}
func StringList_ToProto(mapCtx *direct.MapContext, in *krm.StringList) *pb.StringList {
	if in == nil {
		return nil
	}
	out := &pb.StringList{}
	out.Values = in.Values
	return out
}
