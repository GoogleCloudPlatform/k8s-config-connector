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
	gkehubapi "google.golang.org/api/gkehub/v1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GKEHubScopeRBACRoleBindingSpec_ToAPI(mapCtx *direct.MapContext, r *krm.GKEHubScopeRBACRoleBindingSpec) *gkehubapi.RBACRoleBinding {
	if r == nil {
		return nil
	}
	out := &gkehubapi.RBACRoleBinding{}
	out.Role = GKEHubScopeRBACRoleBindingRole_ToAPI(mapCtx, r.Role)
	out.User = direct.ValueOf(r.User)
	out.Group = direct.ValueOf(r.Group)
	return out
}

func GKEHubScopeRBACRoleBindingRole_ToAPI(mapCtx *direct.MapContext, r *krm.GKEHubScopeRBACRoleBindingRole) *gkehubapi.Role {
	if r == nil {
		return nil
	}
	out := &gkehubapi.Role{}
	out.PredefinedRole = direct.ValueOf(r.PredefinedRole)
	out.CustomRole = direct.ValueOf(r.CustomRole)
	return out
}

func GKEHubScopeRBACRoleBindingObservedState_FromProto(mapCtx *direct.MapContext, r *gkehubapi.RBACRoleBinding) *krm.GKEHubScopeRBACRoleBindingObservedState {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubScopeRBACRoleBindingObservedState{}
	out.CreateTime = direct.LazyPtr(r.CreateTime)
	out.UpdateTime = direct.LazyPtr(r.UpdateTime)
	out.DeleteTime = direct.LazyPtr(r.DeleteTime)
	out.Uid = direct.LazyPtr(r.Uid)
	if r.State != nil {
		out.State = &krm.GKEHubScopeRBACRoleBindingStateStatus{
			Code: direct.LazyPtr(r.State.Code),
		}
	}
	return out
}

func GKEHubScopeRBACRoleBindingStatus_FromProto(mapCtx *direct.MapContext, r *gkehubapi.RBACRoleBinding) *krm.GKEHubScopeRBACRoleBindingStatus {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubScopeRBACRoleBindingStatus{}
	out.ObservedState = GKEHubScopeRBACRoleBindingObservedState_FromProto(mapCtx, r)
	return out
}
