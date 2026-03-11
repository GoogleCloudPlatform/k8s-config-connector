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

func GKEHubNamespaceSpec_ToAPI(mapCtx *direct.MapContext, in *krm.GKEHubNamespaceSpec) *gkehubv1.Namespace {
	if in == nil {
		return nil
	}
	out := &gkehubv1.Namespace{}
	out.Labels = in.Labels
	out.NamespaceLabels = in.NamespaceLabels
	return out
}

func GKEHubNamespaceStatus_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.Namespace) *krm.GKEHubNamespaceStatus {
	if in == nil {
		return nil
	}
	out := &krm.GKEHubNamespaceStatus{}
	out.ObservedState = &krm.GKEHubNamespaceObservedState{}
	out.ObservedState.Uid = direct.LazyPtr(in.Uid)
	out.ObservedState.CreateTime = direct.LazyPtr(in.CreateTime)
	out.ObservedState.UpdateTime = direct.LazyPtr(in.UpdateTime)
	out.ObservedState.DeleteTime = direct.LazyPtr(in.DeleteTime)
	out.ObservedState.State = NamespaceLifecycleStateStatus_FromAPI(mapCtx, in.State)
	return out
}

func NamespaceLifecycleStateStatus_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.NamespaceLifecycleState) *krm.NamespaceLifecycleStateStatus {
	if in == nil {
		return nil
	}
	out := &krm.NamespaceLifecycleStateStatus{}
	out.Code = direct.LazyPtr(in.Code)
	return out
}
