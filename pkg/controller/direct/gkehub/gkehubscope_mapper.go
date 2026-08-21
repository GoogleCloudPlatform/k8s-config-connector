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
	gkehubapi "google.golang.org/api/gkehub/v1beta"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GKEHubScopeSpecKRMtoAPI(mapCtx *direct.MapContext, r *krm.GKEHubScopeSpec) *gkehubapi.Scope {
	if r == nil {
		return nil
	}
	out := &gkehubapi.Scope{}
	out.NamespaceLabels = r.NamespaceLabels
	return out
}

func GKEHubScopeSpecAPIToKRM(mapCtx *direct.MapContext, r *gkehubapi.Scope, id *krm.GKEHubScopeIdentity) *krm.GKEHubScopeSpec {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubScopeSpec{}
	out.ProjectRef = refs.ProjectRef{External: "projects/" + id.ProjectID}
	out.Location = direct.LazyPtr(id.Location)
	out.NamespaceLabels = r.NamespaceLabels
	out.ResourceID = direct.LazyPtr(id.ID())
	return out
}

func GKEHubScopeStatusAPIToKRM(mapCtx *direct.MapContext, r *gkehubapi.Scope) *krm.GKEHubScopeStatus {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubScopeStatus{}
	out.CreateTime = direct.LazyPtr(r.CreateTime)
	out.UpdateTime = direct.LazyPtr(r.UpdateTime)
	out.DeleteTime = direct.LazyPtr(r.DeleteTime)
	out.Uid = direct.LazyPtr(r.Uid)
	if r.State != nil {
		out.State = &krm.GKEHubScopeStateStatus{
			Code: direct.LazyPtr(r.State.Code),
		}
	}
	return out
}
