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

package developerconnect

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/developerconnect/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DeveloperconnectGitRepositoryLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GitRepositoryLink) *krm.DeveloperconnectGitRepositoryLinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeveloperconnectGitRepositoryLinkObservedState{}
	// MISSING: Name
	// MISSING: CloneURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Uid
	// MISSING: WebhookID
	return out
}
func DeveloperconnectGitRepositoryLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeveloperconnectGitRepositoryLinkObservedState) *pb.GitRepositoryLink {
	if in == nil {
		return nil
	}
	out := &pb.GitRepositoryLink{}
	// MISSING: Name
	// MISSING: CloneURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Uid
	// MISSING: WebhookID
	return out
}
func DeveloperconnectGitRepositoryLinkSpec_FromProto(mapCtx *direct.MapContext, in *pb.GitRepositoryLink) *krm.DeveloperconnectGitRepositoryLinkSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeveloperconnectGitRepositoryLinkSpec{}
	// MISSING: Name
	// MISSING: CloneURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Uid
	// MISSING: WebhookID
	return out
}
func DeveloperconnectGitRepositoryLinkSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeveloperconnectGitRepositoryLinkSpec) *pb.GitRepositoryLink {
	if in == nil {
		return nil
	}
	out := &pb.GitRepositoryLink{}
	// MISSING: Name
	// MISSING: CloneURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Uid
	// MISSING: WebhookID
	return out
}
func GitRepositoryLink_FromProto(mapCtx *direct.MapContext, in *pb.GitRepositoryLink) *krm.GitRepositoryLink {
	if in == nil {
		return nil
	}
	out := &krm.GitRepositoryLink{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CloneURI = direct.LazyPtr(in.GetCloneUri())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: Reconciling
	out.Annotations = in.Annotations
	// MISSING: Uid
	// MISSING: WebhookID
	return out
}
func GitRepositoryLink_ToProto(mapCtx *direct.MapContext, in *krm.GitRepositoryLink) *pb.GitRepositoryLink {
	if in == nil {
		return nil
	}
	out := &pb.GitRepositoryLink{}
	out.Name = direct.ValueOf(in.Name)
	out.CloneUri = direct.ValueOf(in.CloneURI)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: Reconciling
	out.Annotations = in.Annotations
	// MISSING: Uid
	// MISSING: WebhookID
	return out
}
func GitRepositoryLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GitRepositoryLink) *krm.GitRepositoryLinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GitRepositoryLinkObservedState{}
	// MISSING: Name
	// MISSING: CloneURI
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	// MISSING: Labels
	// MISSING: Etag
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: Annotations
	out.Uid = direct.LazyPtr(in.GetUid())
	out.WebhookID = direct.LazyPtr(in.GetWebhookId())
	return out
}
func GitRepositoryLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GitRepositoryLinkObservedState) *pb.GitRepositoryLink {
	if in == nil {
		return nil
	}
	out := &pb.GitRepositoryLink{}
	// MISSING: Name
	// MISSING: CloneURI
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	// MISSING: Labels
	// MISSING: Etag
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: Annotations
	out.Uid = direct.ValueOf(in.Uid)
	out.WebhookId = direct.ValueOf(in.WebhookID)
	return out
}
