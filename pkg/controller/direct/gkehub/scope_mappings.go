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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	gkehubv1 "google.golang.org/api/gkehub/v1"
)

func GKEHubScopeSpec_ToAPI(mapCtx *direct.MapContext, in *krm.GKEHubScopeSpec) *gkehubv1.Scope {
	if in == nil {
		return nil
	}
	out := &gkehubv1.Scope{}
	out.Labels = in.Labels
	out.NamespaceLabels = in.NamespaceLabels
	return out
}

func GKEHubScopeStatus_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.Scope) *krm.GKEHubScopeStatus {
	if in == nil {
		return nil
	}
	out := &krm.GKEHubScopeStatus{}
	out.Uid = direct.LazyPtr(in.Uid)
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	out.DeleteTime = direct.LazyPtr(in.DeleteTime)
	out.State = ScopeLifecycleStateStatus_FromAPI(mapCtx, in.State)
	return out
}

func ScopeLifecycleStateStatus_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.ScopeLifecycleState) *krm.ScopeLifecycleStateStatus {
	if in == nil {
		return nil
	}
	out := &krm.ScopeLifecycleStateStatus{}
	out.Code = direct.LazyPtr(in.Code)
	return out
}
