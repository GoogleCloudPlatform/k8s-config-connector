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

func FirewallEndpointSpec_ToAPI(ctx *direct.MapContext, in *v1beta1.NetworkSecurityFirewallEndpointSpec) *api.FirewallEndpoint {
	if in == nil {
		return nil
	}

	out := &api.FirewallEndpoint{}
	if in.Description != nil {
		out.Description = *in.Description
	}
	if len(in.Labels) > 0 {
		out.Labels = in.Labels
	}
	if in.BillingProjectID != nil {
		out.BillingProjectId = *in.BillingProjectID
	}
	return out
}

func FirewallEndpointSpec_FromAPI(ctx *direct.MapContext, in *api.FirewallEndpoint) *v1beta1.NetworkSecurityFirewallEndpointSpec {
	if in == nil {
		return nil
	}

	out := &v1beta1.NetworkSecurityFirewallEndpointSpec{}
	out.Description = direct.LazyPtr(in.Description)
	if len(in.Labels) > 0 {
		out.Labels = in.Labels
	}
	out.BillingProjectID = direct.LazyPtr(in.BillingProjectId)
	// Note: Parent fields (OrganizationRef, Location) are populated by Export.
	return out
}

func FirewallEndpointObservedState_FromAPI(ctx *direct.MapContext, in *api.FirewallEndpoint) *v1beta1.NetworkSecurityFirewallEndpointObservedState {
	if in == nil {
		return nil
	}

	out := &v1beta1.NetworkSecurityFirewallEndpointObservedState{}
	out.AssociatedNetworks = in.AssociatedNetworks
	if in.UpdateTime != "" {
		out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	}
	return out
}
