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

func NetworkSecuritySecurityProfileSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecuritySecurityProfileSpec) *api.SecurityProfile {
	if in == nil {
		return nil
	}
	out := &api.SecurityProfile{}
	out.Type = direct.ValueOf(in.Type)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	if in.CustomInterceptProfile != nil {
		out.CustomInterceptProfile = &api.CustomInterceptProfile{}
		if in.CustomInterceptProfile.InterceptEndpointGroupRef != nil {
			out.CustomInterceptProfile.InterceptEndpointGroup = in.CustomInterceptProfile.InterceptEndpointGroupRef.External
		}
	}
	if in.CustomMirroringProfile != nil {
		out.CustomMirroringProfile = &api.CustomMirroringProfile{}
		if in.CustomMirroringProfile.MirroringEndpointGroupRef != nil {
			out.CustomMirroringProfile.MirroringEndpointGroup = in.CustomMirroringProfile.MirroringEndpointGroupRef.External
		}
	}
	if in.ThreatPreventionProfile != nil {
		out.ThreatPreventionProfile = &api.ThreatPreventionProfile{}
		for _, v := range in.ThreatPreventionProfile.SeverityOverrides {
			out.ThreatPreventionProfile.SeverityOverrides = append(out.ThreatPreventionProfile.SeverityOverrides, &api.SeverityOverride{
				Action:   direct.ValueOf(v.Action),
				Severity: direct.ValueOf(v.Severity),
			})
		}
		for _, v := range in.ThreatPreventionProfile.ThreatOverrides {
			out.ThreatPreventionProfile.ThreatOverrides = append(out.ThreatPreventionProfile.ThreatOverrides, &api.ThreatOverride{
				Action:   direct.ValueOf(v.Action),
				ThreatId: direct.ValueOf(v.ThreatID),
			})
		}
		for _, v := range in.ThreatPreventionProfile.AntivirusOverrides {
			out.ThreatPreventionProfile.AntivirusOverrides = append(out.ThreatPreventionProfile.AntivirusOverrides, &api.AntivirusOverride{
				Action:   direct.ValueOf(v.Action),
				Protocol: direct.ValueOf(v.Protocol),
			})
		}
	}
	if in.URLFilteringProfile != nil {
		out.UrlFilteringProfile = &api.UrlFilteringProfile{}
		for _, v := range in.URLFilteringProfile.URLFilters {
			filter := &api.UrlFilter{
				FilteringAction: direct.ValueOf(v.FilteringAction),
			}
			if v.Priority != nil {
				filter.Priority = int64(*v.Priority)
			}
			if len(v.Urls) > 0 {
				filter.Urls = v.Urls
			}
			out.UrlFilteringProfile.UrlFilters = append(out.UrlFilteringProfile.UrlFilters, filter)
		}
	}
	return out
}

func NetworkSecuritySecurityProfileObservedState_FromProto(mapCtx *direct.MapContext, in *api.SecurityProfile) *krm.NetworkSecuritySecurityProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecuritySecurityProfileObservedState{}
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	out.Etag = direct.LazyPtr(in.Etag)
	if in.ThreatPreventionProfile != nil {
		out.ThreatPreventionProfile = &krm.ThreatPreventionProfileObservedState{}
		for _, v := range in.ThreatPreventionProfile.ThreatOverrides {
			out.ThreatPreventionProfile.ThreatOverrides = append(out.ThreatPreventionProfile.ThreatOverrides, krm.ThreatOverrideObservedState{
				Type: direct.LazyPtr(v.Type),
			})
		}
	}
	return out
}

func NetworkSecuritySecurityProfileSpec_FromProto(mapCtx *direct.MapContext, in *api.SecurityProfile) *krm.NetworkSecuritySecurityProfileSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecuritySecurityProfileSpec{}
	out.Type = direct.LazyPtr(in.Type)
	out.Description = direct.LazyPtr(in.Description)
	out.Labels = in.Labels
	if in.CustomInterceptProfile != nil {
		out.CustomInterceptProfile = &krm.CustomInterceptProfile{}
		// References cannot be easily resolved back to k8s references from the API self-link
		// because we don't store Name/Namespace, but we could set External.
		// Usually we don't map references back in FromProto for direct controllers if they are managed by users.
	}
	if in.CustomMirroringProfile != nil {
		out.CustomMirroringProfile = &krm.CustomMirroringProfile{}
	}
	return out
}
