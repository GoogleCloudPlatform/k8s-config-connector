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

// +generated:mapper
// krm.group: networksecurity.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.networksecurity.v1

package networksecurity

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/networksecurity/v1"
)

func NetworkSecurityAddressGroupObservedState_FromAPI(mapCtx *direct.MapContext, in *api.AddressGroup) *krm.NetworkSecurityAddressGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityAddressGroupObservedState{}
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	// MISSING: Name
	out.SelfLink = direct.LazyPtr(in.SelfLink)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	return out
}
func NetworkSecurityAddressGroupObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.NetworkSecurityAddressGroupObservedState) *api.AddressGroup {
	if in == nil {
		return nil
	}
	out := &api.AddressGroup{}
	out.CreateTime = direct.ValueOf(in.CreateTime)
	// MISSING: Name
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.UpdateTime = direct.ValueOf(in.UpdateTime)
	return out
}
func NetworkSecurityAddressGroupSpec_FromAPI(mapCtx *direct.MapContext, in *api.AddressGroup) *krm.NetworkSecurityAddressGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityAddressGroupSpec{}
	out.Capacity = direct.LazyPtr(in.Capacity)
	out.Description = direct.LazyPtr(in.Description)
	out.Items = in.Items
	out.Labels = in.Labels
	// MISSING: Name
	out.Purpose = in.Purpose
	out.Type = direct.LazyPtr(in.Type)
	return out
}
func NetworkSecurityAddressGroupSpec_ToAPI(mapCtx *direct.MapContext, in *krm.NetworkSecurityAddressGroupSpec) *api.AddressGroup {
	if in == nil {
		return nil
	}
	out := &api.AddressGroup{}
	out.Capacity = direct.ValueOf(in.Capacity)
	out.Description = direct.ValueOf(in.Description)
	out.Items = in.Items
	out.Labels = in.Labels
	// MISSING: Name
	out.Purpose = in.Purpose
	out.Type = direct.ValueOf(in.Type)
	return out
}
