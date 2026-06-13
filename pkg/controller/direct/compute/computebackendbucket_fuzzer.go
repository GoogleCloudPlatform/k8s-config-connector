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
// proto.message: google.cloud.compute.v1.BackendBucket
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeBackendBucketFuzzer())
}

func computeBackendBucketFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackendBucket{},
		ComputeBackendBucketSpec_v1beta1_FromProto, ComputeBackendBucketSpec_v1beta1_ToProto,
		ComputeBackendBucketStatus_v1beta1_FromProto, ComputeBackendBucketStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".bucket_name")
	f.SpecField(".cdn_policy")
	f.SpecField(".cdn_policy.bypass_cache_on_request_headers")
	f.SpecField(".cdn_policy.bypass_cache_on_request_headers[].header_name")
	f.SpecField(".cdn_policy.cache_key_policy")
	f.SpecField(".cdn_policy.cache_key_policy.include_http_headers")
	f.SpecField(".cdn_policy.cache_key_policy.query_string_whitelist")
	f.SpecField(".cdn_policy.cache_mode")
	f.SpecField(".cdn_policy.client_ttl")
	f.SpecField(".cdn_policy.default_ttl")
	f.SpecField(".cdn_policy.max_ttl")
	f.SpecField(".cdn_policy.negative_caching")
	f.SpecField(".cdn_policy.negative_caching_policy")
	f.SpecField(".cdn_policy.negative_caching_policy[].code")
	f.SpecField(".cdn_policy.negative_caching_policy[].ttl")
	f.SpecField(".cdn_policy.request_coalescing")
	f.SpecField(".cdn_policy.serve_while_stale")
	f.SpecField(".cdn_policy.signed_url_cache_max_age_sec")
	f.SpecField(".compression_mode")
	f.SpecField(".custom_response_headers")
	f.SpecField(".description")
	f.SpecField(".edge_security_policy")
	f.SpecField(".enable_cdn")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_Internal(".id")
	f.Unimplemented_Internal(".kind")
	f.Unimplemented_NotYetTriaged(".params")
	f.Unimplemented_NotYetTriaged(".used_by")
	f.Unimplemented_NotYetTriaged(".cdn_policy.signed_url_key_names")
	f.Unimplemented_NotYetTriaged(".load_balancing_scheme")
	f.Unimplemented_NotYetTriaged(".region")

	return f
}
