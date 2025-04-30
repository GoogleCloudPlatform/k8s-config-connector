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

package cloudidentity

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1beta1"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/cloudidentity/v1beta1"
)

func CloudIdentityGroupObservedState_FromAPI(mapCtx *direct.MapContext, in *api.Group) *krm.CloudIdentityGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityGroupObservedState{}
	out.AdditionalGroupKeys = direct.Slice_FromProto(mapCtx, in.AdditionalGroupKeys, EntityKey_FromAPI)
	return out
}
func CloudIdentityGroupObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.CloudIdentityGroupObservedState) *api.Group {
	if in == nil {
		return nil
	}
	out := &api.Group{}
	out.AdditionalGroupKeys = direct.Slice_ToProto(mapCtx, in.AdditionalGroupKeys, EntityKey_ToAPI)
	return out
}
func CloudIdentityGroupSpec_FromAPI(mapCtx *direct.MapContext, in *api.Group) *krm.CloudIdentityGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityGroupSpec{}
	out.Description = direct.LazyPtr(in.Description)
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.GroupKey = EntityKey_FromAPI(mapCtx, in.GroupKey)
	out.Labels = in.Labels
	out.Parent = direct.LazyPtr(in.Parent)
	return out
}
func CloudIdentityGroupSpec_ToAPI(mapCtx *direct.MapContext, in *krm.CloudIdentityGroupSpec) *api.Group {
	if in == nil {
		return nil
	}
	out := &api.Group{}
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.GroupKey = EntityKey_ToAPI(mapCtx, in.GroupKey)
	out.Labels = in.Labels
	out.Parent = direct.ValueOf(in.Parent)
	return out
}
func CloudIdentityGroupStatus_FromAPI(mapCtx *direct.MapContext, in *api.Group) *krm.CloudIdentityGroupStatus {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityGroupStatus{}
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.Name = direct.LazyPtr(in.Name)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	out.ObservedState = CloudIdentityGroupObservedState_FromAPI(mapCtx, in)
	return out
}
func CloudIdentityGroupStatus_ToAPI(mapCtx *direct.MapContext, in *krm.CloudIdentityGroupStatus) *api.Group {
	if in == nil {
		return nil
	}
	out := &api.Group{}
	out.CreateTime = direct.ValueOf(in.CreateTime)
	out.Name = direct.ValueOf(in.Name)
	out.UpdateTime = direct.ValueOf(in.UpdateTime)
	observedState := CloudIdentityGroupObservedState_ToAPI(mapCtx, in.ObservedState)
	if observedState != nil {
		out.AdditionalGroupKeys = observedState.AdditionalGroupKeys
	}
	return out
}
func EntityKey_FromAPI(mapCtx *direct.MapContext, in *api.EntityKey) *krm.EntityKey {
	if in == nil {
		return nil
	}
	out := &krm.EntityKey{}
	out.ID = in.Id
	out.Namespace = direct.LazyPtr(in.Namespace)
	return out
}
func EntityKey_ToAPI(mapCtx *direct.MapContext, in *krm.EntityKey) *api.EntityKey {
	if in == nil {
		return nil
	}
	out := &api.EntityKey{}
	out.Id = in.ID
	out.Namespace = direct.ValueOf(in.Namespace)
	return out
}
