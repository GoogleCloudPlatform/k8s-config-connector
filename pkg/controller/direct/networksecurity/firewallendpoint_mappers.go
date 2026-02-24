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
	networksecuritypb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirewallEndpointSpec_ToProto(ctx *direct.MapContext, in *v1beta1.NetworkSecurityFirewallEndpointSpec) *networksecuritypb.FirewallEndpoint {
	if in == nil {
		return nil
	}

	out := &networksecuritypb.FirewallEndpoint{}
	if in.Description != nil {
		out.Description = *in.Description
	}
	if in.Labels != nil {
		out.Labels = in.Labels
	}
	if in.BillingProjectID != nil {
		out.BillingProjectId = *in.BillingProjectID
	}
	return out
}

func FirewallEndpointObservedState_FromProto(ctx *direct.MapContext, in *networksecuritypb.FirewallEndpoint) *v1beta1.NetworkSecurityFirewallEndpointStatus {
	if in == nil {
		return nil
	}

	out := &v1beta1.NetworkSecurityFirewallEndpointStatus{}
	out.AssociatedNetworks = in.AssociatedNetworks
	// TODO: map timestamps or other fields if needed
	return out
}
