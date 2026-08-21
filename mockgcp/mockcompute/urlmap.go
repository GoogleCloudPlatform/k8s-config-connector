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

package mockcompute

import (
	"context"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

func (s *MockService) populateURLMapDefaults(ctx context.Context, obj *pb.UrlMap) {
	if obj.DefaultService != nil {
		obj.DefaultService = PtrTo(ExpandComputeLink(ctx, *obj.DefaultService))
	}

	for _, pm := range obj.PathMatchers {
		if pm.DefaultService != nil {
			pm.DefaultService = PtrTo(ExpandComputeLink(ctx, *pm.DefaultService))
		}

		for _, pr := range pm.PathRules {
			if pr.RouteAction != nil {
				for _, wbs := range pr.RouteAction.WeightedBackendServices {
					if wbs.BackendService != nil {
						wbs.BackendService = PtrTo(ExpandComputeLink(ctx, *wbs.BackendService))
					}
				}
				if pr.RouteAction.RequestMirrorPolicy != nil {
					if pr.RouteAction.RequestMirrorPolicy.BackendService != nil {
						pr.RouteAction.RequestMirrorPolicy.BackendService = PtrTo(ExpandComputeLink(ctx, *pr.RouteAction.RequestMirrorPolicy.BackendService))
					}
				}
				if pr.RouteAction.RetryPolicy != nil {
					if pr.RouteAction.RetryPolicy.PerTryTimeout != nil {
						if pr.RouteAction.RetryPolicy.PerTryTimeout.Nanos == nil {
							pr.RouteAction.RetryPolicy.PerTryTimeout.Nanos = PtrTo[int32](0)
						}
					}
				}
			}
		}
	}

	for _, t := range obj.Tests {
		if t.Description == nil {
			t.Description = PtrTo("")
		}
	}
}
