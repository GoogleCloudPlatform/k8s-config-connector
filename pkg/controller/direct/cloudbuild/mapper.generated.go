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

package cloudbuild

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudbuild/apiv2/cloudbuildpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CloudbuildRepositoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.CloudbuildRepositoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudbuildRepositoryObservedState{}
	// MISSING: Name
	// MISSING: RemoteURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: WebhookID
	return out
}
func CloudbuildRepositoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudbuildRepositoryObservedState) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	// MISSING: Name
	// MISSING: RemoteURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: WebhookID
	return out
}
func CloudbuildRepositorySpec_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.CloudbuildRepositorySpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudbuildRepositorySpec{}
	// MISSING: Name
	// MISSING: RemoteURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: WebhookID
	return out
}
func CloudbuildRepositorySpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudbuildRepositorySpec) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	// MISSING: Name
	// MISSING: RemoteURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: WebhookID
	return out
}
func Repository_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.Repository {
	if in == nil {
		return nil
	}
	out := &krm.Repository{}
	out.Name = direct.LazyPtr(in.GetName())
	out.RemoteURI = direct.LazyPtr(in.GetRemoteUri())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Annotations = in.Annotations
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: WebhookID
	return out
}
func Repository_ToProto(mapCtx *direct.MapContext, in *krm.Repository) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	out.Name = direct.ValueOf(in.Name)
	out.RemoteUri = direct.ValueOf(in.RemoteURI)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Annotations = in.Annotations
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: WebhookID
	return out
}
func RepositoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.RepositoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RepositoryObservedState{}
	// MISSING: Name
	// MISSING: RemoteURI
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Annotations
	// MISSING: Etag
	out.WebhookID = direct.LazyPtr(in.GetWebhookId())
	return out
}
func RepositoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RepositoryObservedState) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	// MISSING: Name
	// MISSING: RemoteURI
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Annotations
	// MISSING: Etag
	out.WebhookId = direct.ValueOf(in.WebhookID)
	return out
}
