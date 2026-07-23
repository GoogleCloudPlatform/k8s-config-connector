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

// +tool:fuzz-gen
// api.group: networksecurity.cnrm.cloud.google.com

package partnerssegateway

import (
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(partnerSSEGatewayFuzzer())
}

func partnerSSEGatewayFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto[*gcpPartnerSSEGateway, krm.NetworkSecurityPartnerSSEGatewaySpec, krm.NetworkSecurityPartnerSSEGatewayStatus](&gcpPartnerSSEGateway{},
		SpecFromAPI, SpecToAPI,
		StatusFromAPI, StatusToAPI,
	)

	f.SpecField(".Labels")
	f.SpecField(".SseGatewayReferenceId")
	f.SpecField(".PartnerVpcSubnetRange")
	f.SpecField(".SseSubnetRange")
	f.SpecField(".PartnerSubnetRange")
	f.SpecField(".Vni")
	f.SpecField(".SymantecOptions")
	f.SpecField(".SymantecOptions.SymantecSiteTargetHost")

	f.StatusField(".CreateTime")
	f.StatusField(".UpdateTime")
	f.StatusField(".SseVpcSubnetRange")
	f.StatusField(".SseVpcTargetIp")
	f.StatusField(".SseBgpIps")
	f.StatusField(".SseBgpAsn")
	f.StatusField(".PartnerSseRealm")
	f.StatusField(".SseTargetIp")
	f.StatusField(".SymantecOptionsStatus")
	f.StatusField(".SymantecOptionsStatus.SymantecLocationUuid")
	f.StatusField(".SymantecOptionsStatus.SymantecSite")
	f.StatusField(".SseProject")
	f.StatusField(".SseNetwork")
	f.StatusField(".PartnerSseEnvironment")
	f.StatusField(".Country")
	f.StatusField(".Timezone")
	f.StatusField(".CapacityBps")
	f.StatusField(".State")
	f.StatusField(".ProberSubnetRanges")

	f.UnimplementedFields.Insert(".Name")

	f.FilterSpec = func(in *gcpPartnerSSEGateway) {
		in.CreateTime = ""
		in.UpdateTime = ""
		in.SseVpcSubnetRange = ""
		in.SseVpcTargetIp = ""
		in.SseBgpIps = nil
		in.SseBgpAsn = 0
		in.PartnerSseRealm = ""
		in.SseTargetIp = ""
		in.SymantecOptionsStatus = nil
		in.SseProject = ""
		in.SseNetwork = ""
		in.PartnerSseEnvironment = ""
		in.Country = ""
		in.Timezone = ""
		in.CapacityBps = ""
		in.State = ""
		in.ProberSubnetRanges = nil
	}

	f.FilterStatus = func(in *gcpPartnerSSEGateway) {
		in.Labels = nil
		in.SseGatewayReferenceId = ""
		in.PartnerVpcSubnetRange = ""
		in.SseSubnetRange = ""
		in.PartnerSubnetRange = ""
		in.Vni = 0
		in.SymantecOptions = nil

		// Sanitize CapacityBps string to be a valid positive integer string
		if in.CapacityBps != "" {
			var hash int64
			for _, c := range in.CapacityBps {
				hash = hash*31 + int64(c)
			}
			if hash < 0 {
				hash = -hash
			}
			in.CapacityBps = fmt.Sprintf("%d", hash)
		}
	}

	return f
}

func SpecFromAPI(mapCtx *direct.MapContext, in *gcpPartnerSSEGateway) *krm.NetworkSecurityPartnerSSEGatewaySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityPartnerSSEGatewaySpec{}
	if in.Labels != nil {
		out.Labels = in.Labels
	}
	if in.SseGatewayReferenceId != "" {
		out.SseGatewayReferenceID = &in.SseGatewayReferenceId
	}
	if in.PartnerVpcSubnetRange != "" {
		out.PartnerVPCSubnetRange = &in.PartnerVpcSubnetRange
	}
	if in.SseSubnetRange != "" {
		out.SseSubnetRange = &in.SseSubnetRange
	}
	if in.PartnerSubnetRange != "" {
		out.PartnerSubnetRange = &in.PartnerSubnetRange
	}
	if in.Vni != 0 {
		out.Vni = &in.Vni
	}
	if in.SymantecOptions != nil {
		out.SymantecOptions = &krm.PartnerSSEGatewaySymantecOptions{}
		if in.SymantecOptions.SymantecSiteTargetHost != "" {
			out.SymantecOptions.SymantecSiteTargetHost = &in.SymantecOptions.SymantecSiteTargetHost
		}
	}
	return out
}

func SpecToAPI(mapCtx *direct.MapContext, in *krm.NetworkSecurityPartnerSSEGatewaySpec) *gcpPartnerSSEGateway {
	if in == nil {
		return nil
	}
	return KRMtoGCP(in)
}

func StatusFromAPI(mapCtx *direct.MapContext, in *gcpPartnerSSEGateway) *krm.NetworkSecurityPartnerSSEGatewayStatus {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityPartnerSSEGatewayStatus{}
	state := &krm.NetworkSecurityPartnerSSEGatewayObservedState{}

	if in.CreateTime != "" {
		state.CreateTime = &in.CreateTime
	}
	if in.UpdateTime != "" {
		state.UpdateTime = &in.UpdateTime
	}
	if in.SseVpcSubnetRange != "" {
		state.SseVPCSubnetRange = &in.SseVpcSubnetRange
	}
	if in.SseVpcTargetIp != "" {
		state.SseVPCTargetIP = &in.SseVpcTargetIp
	}
	if len(in.SseBgpIps) > 0 {
		state.SseBGPIPs = in.SseBgpIps
	}
	if in.SseBgpAsn != 0 {
		state.SseBGPAsn = &in.SseBgpAsn
	}
	if in.PartnerSseRealm != "" {
		state.PartnerSSERealm = &in.PartnerSseRealm
	}
	if in.SseTargetIp != "" {
		state.SseTargetIP = &in.SseTargetIp
	}
	if in.SymantecOptionsStatus != nil {
		state.SymantecOptions = &krm.PartnerSSEGatewaySymantecOptionsObservedState{}
		if in.SymantecOptionsStatus.SymantecLocationUuid != "" {
			state.SymantecOptions.SymantecLocationUuid = &in.SymantecOptionsStatus.SymantecLocationUuid
		}
		if in.SymantecOptionsStatus.SymantecSite != "" {
			state.SymantecOptions.SymantecSite = &in.SymantecOptionsStatus.SymantecSite
		}
	}
	if in.SseProject != "" {
		state.SseProject = &in.SseProject
	}
	if in.SseNetwork != "" {
		state.SseNetwork = &in.SseNetwork
	}
	if in.PartnerSseEnvironment != "" {
		state.PartnerSSEEnvironment = &in.PartnerSseEnvironment
	}
	if in.Country != "" {
		state.Country = &in.Country
	}
	if in.Timezone != "" {
		state.Timezone = &in.Timezone
	}
	if in.CapacityBps != "" {
		var val int64
		if _, err := fmt.Sscanf(in.CapacityBps, "%d", &val); err == nil {
			state.CapacityBps = &val
		}
	}
	if in.State != "" {
		state.State = &in.State
	}
	if len(in.ProberSubnetRanges) > 0 {
		state.ProberSubnetRanges = in.ProberSubnetRanges
	}

	out.ObservedState = state
	return out
}

func StatusToAPI(mapCtx *direct.MapContext, in *krm.NetworkSecurityPartnerSSEGatewayStatus) *gcpPartnerSSEGateway {
	if in == nil {
		return nil
	}
	out := &gcpPartnerSSEGateway{}
	state := in.ObservedState
	if state != nil {
		if state.CreateTime != nil {
			out.CreateTime = *state.CreateTime
		}
		if state.UpdateTime != nil {
			out.UpdateTime = *state.UpdateTime
		}
		if state.SseVPCSubnetRange != nil {
			out.SseVpcSubnetRange = *state.SseVPCSubnetRange
		}
		if state.SseVPCTargetIP != nil {
			out.SseVpcTargetIp = *state.SseVPCTargetIP
		}
		if len(state.SseBGPIPs) > 0 {
			out.SseBgpIps = state.SseBGPIPs
		}
		if state.SseBGPAsn != nil {
			out.SseBgpAsn = *state.SseBGPAsn
		}
		if state.PartnerSSERealm != nil {
			out.PartnerSseRealm = *state.PartnerSSERealm
		}
		if state.SseTargetIP != nil {
			out.SseTargetIp = *state.SseTargetIP
		}
		if state.SymantecOptions != nil {
			out.SymantecOptionsStatus = &gcpSymantecOptionsObservedState{}
			if state.SymantecOptions.SymantecLocationUuid != nil {
				out.SymantecOptionsStatus.SymantecLocationUuid = *state.SymantecOptions.SymantecLocationUuid
			}
			if state.SymantecOptions.SymantecSite != nil {
				out.SymantecOptionsStatus.SymantecSite = *state.SymantecOptions.SymantecSite
			}
		}
		if state.SseProject != nil {
			out.SseProject = *state.SseProject
		}
		if state.SseNetwork != nil {
			out.SseNetwork = *state.SseNetwork
		}
		if state.PartnerSSEEnvironment != nil {
			out.PartnerSseEnvironment = *state.PartnerSSEEnvironment
		}
		if state.Country != nil {
			out.Country = *state.Country
		}
		if state.Timezone != nil {
			out.Timezone = *state.Timezone
		}
		if state.CapacityBps != nil {
			out.CapacityBps = fmt.Sprintf("%d", *state.CapacityBps)
		}
		if state.State != nil {
			out.State = *state.State
		}
		if len(state.ProberSubnetRanges) > 0 {
			out.ProberSubnetRanges = state.ProberSubnetRanges
		}
	}
	return out
}
