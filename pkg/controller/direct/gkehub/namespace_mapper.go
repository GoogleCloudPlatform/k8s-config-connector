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

package gkehub

import (
	gkehubv1 "google.golang.org/api/gkehub/v1"
	"google.golang.org/protobuf/proto"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GKEHubNamespaceSpec_ToAPI(mapCtx *direct.MapContext, r *krm.GKEHubNamespaceSpec) *gkehubv1.Namespace {
	if r == nil {
		return nil
	}
	out := &gkehubv1.Namespace{}
	out.Labels = r.Labels
	out.NamespaceLabels = r.NamespaceLabels
	return out
}

func GKEHubNamespaceSpec_FromAPI(mapCtx *direct.MapContext, r *gkehubv1.Namespace, id *krm.GKEHubNamespaceIdentity) *krm.GKEHubNamespaceSpec {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubNamespaceSpec{}
	out.ProjectRef = &refs.ProjectRef{External: "projects/" + id.ProjectID}
	out.Location = direct.LazyPtr(id.Location)
	out.ScopeRef = &krm.GKEHubScopeRef{External: id.Parent().String()}
	out.NamespaceID = direct.LazyPtr(id.ID())
	out.ResourceID = direct.LazyPtr(id.ID())
	out.Labels = r.Labels
	out.NamespaceLabels = r.NamespaceLabels
	return out
}

func GKEHubNamespaceStatus_FromAPI(mapCtx *direct.MapContext, r *gkehubv1.Namespace) *krm.GKEHubNamespaceStatus {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubNamespaceStatus{}
	out.ObservedState = &krm.GKEHubNamespaceObservedState{
		CreateTime: direct.LazyPtr(r.CreateTime),
		UpdateTime: direct.LazyPtr(r.UpdateTime),
		DeleteTime: direct.LazyPtr(r.DeleteTime),
		Uid:        direct.LazyPtr(r.Uid),
	}
	if r.State != nil {
		out.ObservedState.State = direct.LazyPtr(r.State.Code)
	}
	return out
}

func GKEHubNamespaceStatus_ToAPI(mapCtx *direct.MapContext, r *krm.GKEHubNamespaceStatus) *gkehubv1.Namespace {
	if r == nil {
		return nil
	}
	out := &gkehubv1.Namespace{}
	if r.ObservedState != nil {
		out.CreateTime = direct.ValueOf(r.ObservedState.CreateTime)
		out.UpdateTime = direct.ValueOf(r.ObservedState.UpdateTime)
		out.DeleteTime = direct.ValueOf(r.ObservedState.DeleteTime)
		out.Uid = direct.ValueOf(r.ObservedState.Uid)
		if r.ObservedState.State != nil {
			out.State = &gkehubv1.NamespaceLifecycleState{
				Code: direct.ValueOf(r.ObservedState.State),
			}
		}
	}
	return out
}

// GKEHubNamespaceSpec_ToProto is a placeholder for the proto mapping.
// Currently GKEHub Namespace proto is not available in the repo.
func GKEHubNamespaceSpec_ToProto(mapCtx *direct.MapContext, r *krm.GKEHubNamespaceSpec) proto.Message {
	return nil
}

// GKEHubNamespaceStatus_FromProto is a placeholder for the proto mapping.
func GKEHubNamespaceStatus_FromProto(mapCtx *direct.MapContext, r proto.Message) *krm.GKEHubNamespaceStatus {
	return nil
}

// GKEHubNamespace_FromAPItoProto is a placeholder for the proto mapping.
func GKEHubNamespace_FromAPItoProto(mapCtx *direct.MapContext, r *gkehubv1.Namespace) proto.Message {
	return nil
}
