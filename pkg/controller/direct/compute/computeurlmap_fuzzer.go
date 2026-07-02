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
// proto.message: google.cloud.compute.v1.UrlMap
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeURLMapFuzzer())
}

func computeURLMapFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.UrlMap{},
		ComputeURLMapSpec_v1beta1_FromProto, ComputeURLMapSpec_v1beta1_ToProto,
		ComputeURLMapStatus_v1beta1_FromProto, ComputeURLMapStatus_v1beta1_ToProto,
	)

	// Field comparison: ComputeURLMap Spec vs pb.UrlMap Proto
	// - Spec.DefaultCustomErrorResponsePolicy maps to proto field .default_custom_error_response_policy
	// - Spec.DefaultRouteAction              maps to proto field .default_route_action
	// - Spec.DefaultService                  maps to proto field .default_service
	// - Spec.DefaultUrlRedirect              maps to proto field .default_url_redirect
	// - Spec.Description                      maps to proto field .description
	// - Spec.HeaderAction                     maps to proto field .header_action
	// - Spec.HostRule                         maps to proto field .host_rules
	// - Spec.Location                         maps to proto field .region (or global if empty)
	// - Spec.PathMatcher                      maps to proto field .path_matchers
	// - Spec.ResourceID                       maps to proto field .name (handled as Unimplemented_Identity)
	// - Spec.Test                             maps to proto field .tests

	// Spec fields
	f.SpecField(".default_custom_error_response_policy")
	f.SpecField(".default_route_action")
	f.SpecField(".default_service")
	f.SpecField(".default_url_redirect")
	f.SpecField(".description")
	f.SpecField(".header_action")
	f.SpecField(".host_rules")
	f.SpecField(".region")
	f.SpecField(".path_matchers")
	f.SpecField(".tests")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".fingerprint")
	f.StatusField(".id")
	f.StatusField(".self_link")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".default_route_action.max_stream_duration")
	f.Unimplemented_NotYetTriaged(".default_route_action.url_rewrite.path_template_rewrite")
	f.Unimplemented_NotYetTriaged(".default_route_action.request_mirror_policy.mirror_percent")
	f.Unimplemented_NotYetTriaged(".path_matchers[].default_route_action.max_stream_duration")
	f.Unimplemented_NotYetTriaged(".path_matchers[].default_route_action.url_rewrite.path_template_rewrite")
	f.Unimplemented_NotYetTriaged(".path_matchers[].default_route_action.request_mirror_policy.mirror_percent")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].custom_error_response_policy")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].route_action.max_stream_duration")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].route_action.url_rewrite.path_template_rewrite")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].route_action.request_mirror_policy.mirror_percent")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].custom_error_response_policy")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].description")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].route_action.max_stream_duration")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].route_action.url_rewrite.path_template_rewrite")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].route_action.request_mirror_policy.mirror_percent")
	f.Unimplemented_NotYetTriaged(".tests[].expected_output_url")
	f.Unimplemented_NotYetTriaged(".tests[].expected_redirect_response_code")
	f.Unimplemented_NotYetTriaged(".tests[].headers")
	f.Unimplemented_NotYetTriaged(".default_route_action.cache_policy")
	f.Unimplemented_NotYetTriaged(".path_matchers[].default_route_action.cache_policy")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].route_action.cache_policy")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].route_action.cache_policy")

	f.FilterSpec = func(in *pb.UrlMap) {
		normalizeRedirect := func(r *pb.HttpRedirectAction) {
			if r != nil {
				if r.StripQuery != nil && !*r.StripQuery {
					r.StripQuery = nil
				}
			}
		}
		normalizeHeaderAction := func(h *pb.HttpHeaderAction) {
			if h == nil {
				return
			}
			for _, o := range h.RequestHeadersToAdd {
				if o != nil {
					if o.Replace != nil && !*o.Replace {
						o.Replace = nil
					}
				}
			}
			for _, o := range h.ResponseHeadersToAdd {
				if o != nil {
					if o.Replace != nil && !*o.Replace {
						o.Replace = nil
					}
				}
			}
		}
		normalizeCorsPolicy := func(c *pb.CorsPolicy) {
			if c != nil {
				if c.Disabled != nil && !*c.Disabled {
					c.Disabled = nil
				}
			}
		}
		normalizeDuration := func(d *pb.Duration) *pb.Duration {
			if d == nil {
				return nil
			}
			if d.Nanos == nil && d.Seconds == nil {
				return nil
			}
			return d
		}
		normalizeFaultDelay := func(fd *pb.HttpFaultDelay) {
			if fd == nil {
				return
			}
			fd.FixedDelay = normalizeDuration(fd.FixedDelay)
		}
		normalizeFaultInjection := func(fi *pb.HttpFaultInjection) {
			if fi == nil {
				return
			}
			normalizeFaultDelay(fi.Delay)
			if fi.Abort != nil && fi.Abort.HttpStatus == nil && fi.Abort.Percentage == nil {
				fi.Abort = nil
			}
		}
		normalizeRetryPolicy := func(rp *pb.HttpRetryPolicy) {
			if rp == nil {
				return
			}
			rp.PerTryTimeout = normalizeDuration(rp.PerTryTimeout)
		}
		normalizeRouteAction := func(ra *pb.HttpRouteAction) {
			if ra == nil {
				return
			}
			normalizeCorsPolicy(ra.CorsPolicy)
			normalizeFaultInjection(ra.FaultInjectionPolicy)
			normalizeRetryPolicy(ra.RetryPolicy)
			ra.Timeout = normalizeDuration(ra.Timeout)
			for _, wbs := range ra.WeightedBackendServices {
				if wbs != nil {
					normalizeHeaderAction(wbs.HeaderAction)
				}
			}
		}

		normalizeHeaderAction(in.HeaderAction)
		normalizeRedirect(in.DefaultUrlRedirect)
		normalizeRouteAction(in.DefaultRouteAction)

		for _, pm := range in.PathMatchers {
			if pm != nil {
				normalizeHeaderAction(pm.HeaderAction)
				normalizeRedirect(pm.DefaultUrlRedirect)
				normalizeRouteAction(pm.DefaultRouteAction)

				for _, pr := range pm.PathRules {
					if pr != nil {
						normalizeRedirect(pr.UrlRedirect)
						normalizeRouteAction(pr.RouteAction)
					}
				}
				for _, rr := range pm.RouteRules {
					if rr != nil {
						normalizeHeaderAction(rr.HeaderAction)
						normalizeRedirect(rr.UrlRedirect)
						normalizeRouteAction(rr.RouteAction)
					}
				}
			}
		}
	}

	return f
}
