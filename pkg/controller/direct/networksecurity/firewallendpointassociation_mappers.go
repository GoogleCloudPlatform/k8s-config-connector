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

package networksecurity

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/networksecurity/v1"
)

func NetworkSecurityFirewallEndpointAssociationSpec_FromAPI(mapCtx *direct.MapContext, in *api.FirewallEndpointAssociation) *krm.NetworkSecurityFirewallEndpointAssociationSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityFirewallEndpointAssociationSpec{}
	out.Labels = in.Labels
	if in.Disabled {
		out.Disabled = direct.LazyPtr(in.Disabled)
	} else {
		// Because it's a bool, and if false, we might want to return pointer to false if it was set,
		// but api doesn't distinguish absent from false easily without NullFields.
		// For now we just return nil if false, but if the API returns false, we should probably output false pointer
		// Let's check ForceSendFields
		hasDisabled := false
		for _, f := range in.ForceSendFields {
			if f == "Disabled" {
				hasDisabled = true
				break
			}
		}
		if hasDisabled || in.Disabled {
			out.Disabled = direct.LazyPtr(in.Disabled)
		}
	}
	// Refs are handled in controller
	return out
}

func NetworkSecurityFirewallEndpointAssociationSpec_ToAPI(mapCtx *direct.MapContext, in *krm.NetworkSecurityFirewallEndpointAssociationSpec) *api.FirewallEndpointAssociation {
	if in == nil {
		return nil
	}
	out := &api.FirewallEndpointAssociation{}
	out.Labels = in.Labels
	if in.Disabled != nil {
		out.Disabled = *in.Disabled
		out.ForceSendFields = append(out.ForceSendFields, "Disabled")
	}
	// Refs are handled in controller
	return out
}

func NetworkSecurityFirewallEndpointAssociationObservedState_FromAPI(mapCtx *direct.MapContext, in *api.FirewallEndpointAssociation) *krm.NetworkSecurityFirewallEndpointAssociationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityFirewallEndpointAssociationObservedState{}
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	if in.State != "" {
		out.State = direct.LazyPtr(in.State)
	}
	out.Reconciling = direct.LazyPtr(in.Reconciling)
	return out
}

func NetworkSecurityFirewallEndpointAssociationObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.NetworkSecurityFirewallEndpointAssociationObservedState) *api.FirewallEndpointAssociation {
	if in == nil {
		return nil
	}
	out := &api.FirewallEndpointAssociation{}
	if in.CreateTime != nil {
		out.CreateTime = *in.CreateTime
	}
	if in.UpdateTime != nil {
		out.UpdateTime = *in.UpdateTime
	}
	if in.State != nil {
		out.State = *in.State
	}
	if in.Reconciling != nil {
		out.Reconciling = *in.Reconciling
	}
	return out
}
