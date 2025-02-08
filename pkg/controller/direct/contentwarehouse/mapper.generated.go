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

package contentwarehouse

import (
	pb "cloud.google.com/go/contentwarehouse/apiv1/contentwarehousepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ContentwarehouseDocumentLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DocumentLink) *krm.ContentwarehouseDocumentLinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContentwarehouseDocumentLinkObservedState{}
	// MISSING: Name
	// MISSING: SourceDocumentReference
	// MISSING: TargetDocumentReference
	// MISSING: Description
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func ContentwarehouseDocumentLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContentwarehouseDocumentLinkObservedState) *pb.DocumentLink {
	if in == nil {
		return nil
	}
	out := &pb.DocumentLink{}
	// MISSING: Name
	// MISSING: SourceDocumentReference
	// MISSING: TargetDocumentReference
	// MISSING: Description
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func ContentwarehouseDocumentLinkSpec_FromProto(mapCtx *direct.MapContext, in *pb.DocumentLink) *krm.ContentwarehouseDocumentLinkSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContentwarehouseDocumentLinkSpec{}
	// MISSING: Name
	// MISSING: SourceDocumentReference
	// MISSING: TargetDocumentReference
	// MISSING: Description
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func ContentwarehouseDocumentLinkSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContentwarehouseDocumentLinkSpec) *pb.DocumentLink {
	if in == nil {
		return nil
	}
	out := &pb.DocumentLink{}
	// MISSING: Name
	// MISSING: SourceDocumentReference
	// MISSING: TargetDocumentReference
	// MISSING: Description
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func DocumentLink_FromProto(mapCtx *direct.MapContext, in *pb.DocumentLink) *krm.DocumentLink {
	if in == nil {
		return nil
	}
	out := &krm.DocumentLink{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SourceDocumentReference = DocumentReference_FromProto(mapCtx, in.GetSourceDocumentReference())
	out.TargetDocumentReference = DocumentReference_FromProto(mapCtx, in.GetTargetDocumentReference())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: UpdateTime
	// MISSING: CreateTime
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func DocumentLink_ToProto(mapCtx *direct.MapContext, in *krm.DocumentLink) *pb.DocumentLink {
	if in == nil {
		return nil
	}
	out := &pb.DocumentLink{}
	out.Name = direct.ValueOf(in.Name)
	out.SourceDocumentReference = DocumentReference_ToProto(mapCtx, in.SourceDocumentReference)
	out.TargetDocumentReference = DocumentReference_ToProto(mapCtx, in.TargetDocumentReference)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: UpdateTime
	// MISSING: CreateTime
	out.State = direct.Enum_ToProto[pb.DocumentLink_State](mapCtx, in.State)
	return out
}
func DocumentLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DocumentLink) *krm.DocumentLinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentLinkObservedState{}
	// MISSING: Name
	out.SourceDocumentReference = DocumentReferenceObservedState_FromProto(mapCtx, in.GetSourceDocumentReference())
	// MISSING: TargetDocumentReference
	// MISSING: Description
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: State
	return out
}
func DocumentLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentLinkObservedState) *pb.DocumentLink {
	if in == nil {
		return nil
	}
	out := &pb.DocumentLink{}
	// MISSING: Name
	out.SourceDocumentReference = DocumentReferenceObservedState_ToProto(mapCtx, in.SourceDocumentReference)
	// MISSING: TargetDocumentReference
	// MISSING: Description
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: State
	return out
}
func DocumentReference_FromProto(mapCtx *direct.MapContext, in *pb.DocumentReference) *krm.DocumentReference {
	if in == nil {
		return nil
	}
	out := &krm.DocumentReference{}
	out.DocumentName = direct.LazyPtr(in.GetDocumentName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Snippet = direct.LazyPtr(in.GetSnippet())
	out.DocumentIsFolder = direct.LazyPtr(in.GetDocumentIsFolder())
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: DeleteTime
	out.DocumentIsRetentionFolder = direct.LazyPtr(in.GetDocumentIsRetentionFolder())
	out.DocumentIsLegalHoldFolder = direct.LazyPtr(in.GetDocumentIsLegalHoldFolder())
	return out
}
func DocumentReference_ToProto(mapCtx *direct.MapContext, in *krm.DocumentReference) *pb.DocumentReference {
	if in == nil {
		return nil
	}
	out := &pb.DocumentReference{}
	out.DocumentName = direct.ValueOf(in.DocumentName)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Snippet = direct.ValueOf(in.Snippet)
	out.DocumentIsFolder = direct.ValueOf(in.DocumentIsFolder)
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: DeleteTime
	out.DocumentIsRetentionFolder = direct.ValueOf(in.DocumentIsRetentionFolder)
	out.DocumentIsLegalHoldFolder = direct.ValueOf(in.DocumentIsLegalHoldFolder)
	return out
}
func DocumentReferenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DocumentReference) *krm.DocumentReferenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentReferenceObservedState{}
	// MISSING: DocumentName
	// MISSING: DisplayName
	// MISSING: Snippet
	// MISSING: DocumentIsFolder
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	// MISSING: DocumentIsRetentionFolder
	// MISSING: DocumentIsLegalHoldFolder
	return out
}
func DocumentReferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentReferenceObservedState) *pb.DocumentReference {
	if in == nil {
		return nil
	}
	out := &pb.DocumentReference{}
	// MISSING: DocumentName
	// MISSING: DisplayName
	// MISSING: Snippet
	// MISSING: DocumentIsFolder
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	// MISSING: DocumentIsRetentionFolder
	// MISSING: DocumentIsLegalHoldFolder
	return out
}
