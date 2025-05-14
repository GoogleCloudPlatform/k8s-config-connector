// Copyright 2024 Google LLC
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

package firewallpolicyrule

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeFirewallPolicyRuleSpec_TargetResources_ToProto(mapCtx *direct.MapContext, in []*computev1beta1.ComputeNetworkRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i == nil {
			continue
		}
		if i.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}

func ComputeFirewallPolicyRuleSpec_TargetServiceAccounts_ToProto(mapCtx *direct.MapContext, in []*refs.IAMServiceAccountRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i == nil {
			continue
		}
		if i.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}

func ComputeFirewallPolicyRuleSpec_TargetResources_FromProto(mapCtx *direct.MapContext, in []string) []*computev1beta1.ComputeNetworkRef {
	if in == nil {
		return nil
	}
	var out []*computev1beta1.ComputeNetworkRef
	for _, i := range in {
		out = append(out, &computev1beta1.ComputeNetworkRef{
			External: i,
		})
	}
	return out
}

func ComputeFirewallPolicyRuleSpec_TargetServiceAccounts_FromProto(mapCtx *direct.MapContext, in []string) []*refs.IAMServiceAccountRef {
	if in == nil {
		return nil
	}
	var out []*refs.IAMServiceAccountRef
	for _, i := range in {
		out = append(out, &refs.IAMServiceAccountRef{
			External: i,
		})
	}
	return out
}
