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

func GKEHubMembershipBindingSpec_ToAPI(mapCtx *direct.MapContext, r *krm.GKEHubMembershipBindingSpec) *gkehubv1.MembershipBinding {
	if r == nil {
		return nil
	}
	out := &gkehubv1.MembershipBinding{}
	out.Labels = r.Labels
	return out
}

func GKEHubMembershipBindingSpec_FromAPI(mapCtx *direct.MapContext, r *gkehubv1.MembershipBinding, id *krm.GKEHubMembershipBindingIdentity) *krm.GKEHubMembershipBindingSpec {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubMembershipBindingSpec{}
	out.MembershipRef.External = id.Parent().String()
	out.ScopeRef.External = r.Scope
	out.ResourceID = direct.LazyPtr(id.ID())
	out.Labels = r.Labels
	return out
}

func GKEHubMembershipBindingStatus_FromAPI(mapCtx *direct.MapContext, r *gkehubv1.MembershipBinding) *krm.GKEHubMembershipBindingStatus {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubMembershipBindingStatus{}
	out.ObservedState = &krm.GKEHubMembershipBindingObservedState{
		CreateTime: direct.LazyPtr(r.CreateTime),
		UpdateTime: direct.LazyPtr(r.UpdateTime),
		DeleteTime: direct.LazyPtr(r.DeleteTime),
		Uid:        direct.LazyPtr(r.Uid),
	}
	if r.State != nil {
		out.ObservedState.State = &krm.GKEHubMembershipBindingStateStatus{
			Code: direct.LazyPtr(r.State.Code),
		}
	}
	return out
}

func GKEHubMembershipBindingStatus_ToAPI(mapCtx *direct.MapContext, r *krm.GKEHubMembershipBindingStatus) *gkehubv1.MembershipBinding {
	if r == nil {
		return nil
	}
	out := &gkehubv1.MembershipBinding{}
	if r.ObservedState != nil {
		out.CreateTime = direct.ValueOf(r.ObservedState.CreateTime)
		out.UpdateTime = direct.ValueOf(r.ObservedState.UpdateTime)
		out.DeleteTime = direct.ValueOf(r.ObservedState.DeleteTime)
		out.Uid = direct.ValueOf(r.ObservedState.Uid)
		if r.ObservedState.State != nil {
			out.State = &gkehubv1.MembershipBindingLifecycleState{
				Code: direct.ValueOf(r.ObservedState.State.Code),
			}
		}
	}
	return out
}

// GKEHubMembershipBindingSpec_ToProto is a placeholder for the proto mapping.
// Currently GKEHub MembershipBinding proto is not available in the repo.
func GKEHubMembershipBindingSpec_ToProto(mapCtx *direct.MapContext, r *krm.GKEHubMembershipBindingSpec) proto.Message {
	return nil
}

// GKEHubMembershipBindingStatus_FromProto is a placeholder for the proto mapping.
func GKEHubMembershipBindingStatus_FromProto(mapCtx *direct.MapContext, r proto.Message) *krm.GKEHubMembershipBindingStatus {
	return nil
}

// GKEHubMembershipBinding_FromAPItoProto is a placeholder for the proto mapping.
func GKEHubMembershipBinding_FromAPItoProto(mapCtx *direct.MapContext, r *gkehubv1.MembershipBinding) proto.Message {
	return nil
}
