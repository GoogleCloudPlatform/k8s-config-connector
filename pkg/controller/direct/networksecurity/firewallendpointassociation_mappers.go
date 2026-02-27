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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	api "google.golang.org/api/networksecurity/v1beta1"
)

func FirewallEndpointAssociationSpec_ToAPI(ctx *direct.MapContext, in *v1beta1.NetworkSecurityFirewallEndpointAssociationSpec) *api.FirewallEndpointAssociation {
	if in == nil {
		return nil
	}

	out := &api.FirewallEndpointAssociation{}
	if in.Disabled != nil {
		out.Disabled = *in.Disabled
	}
	if len(in.Labels) > 0 {
		out.Labels = in.Labels
	}
	return out
}

func FirewallEndpointAssociationSpec_FromAPI(ctx *direct.MapContext, in *api.FirewallEndpointAssociation) *v1beta1.NetworkSecurityFirewallEndpointAssociationSpec {
	if in == nil {
		return nil
	}

	out := &v1beta1.NetworkSecurityFirewallEndpointAssociationSpec{}
	out.Disabled = direct.LazyPtr(in.Disabled)
	if len(in.Labels) > 0 {
		out.Labels = in.Labels
	}
	// Note: Reference fields are populated by export and shouldn't be mapped here.
	return out
}

func FirewallEndpointAssociationObservedState_FromAPI(ctx *direct.MapContext, in *api.FirewallEndpointAssociation) *v1beta1.NetworkSecurityFirewallEndpointAssociationObservedState {
	if in == nil {
		return nil
	}

	out := &v1beta1.NetworkSecurityFirewallEndpointAssociationObservedState{}
	if in.CreateTime != "" {
		out.CreateTime = direct.LazyPtr(in.CreateTime)
	}
	if in.UpdateTime != "" {
		out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	}
	if in.State != "" {
		out.State = direct.LazyPtr(in.State)
	}
	out.Reconciling = direct.LazyPtr(in.Reconciling)
	return out
}
