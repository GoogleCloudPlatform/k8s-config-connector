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
// proto.message: google.cloud.compute.v1.TargetHttpsProxy
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeTargetHTTPSProxyFuzzer())
}

func computeTargetHTTPSProxyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TargetHttpsProxy{},
		ComputeTargetHTTPSProxySpec_v1beta1_FromProto, ComputeTargetHTTPSProxySpec_v1beta1_ToProto,
		ComputeTargetHTTPSProxyObservedState_v1beta1_FromProto, ComputeTargetHTTPSProxyObservedState_v1beta1_ToProto,
	)

	// Field Comparison & Mapping Documentation:
	//
	// KRM Spec Fields:
	// - certificateManagerCertificates -> f.SpecField(".ssl_certificates")
	// - certificateMapRef              -> f.SpecField(".certificate_map")
	// - description                    -> f.SpecField(".description")
	// - httpKeepAliveTimeoutSec        -> f.SpecField(".http_keep_alive_timeout_sec")
	// - location                       -> f.Unimplemented_Identity(".region") (part of Resource Identity/URL)
	// - proxyBind                      -> f.SpecField(".proxy_bind")
	// - quicOverride                   -> f.SpecField(".quic_override")
	// - resourceID                     -> f.Unimplemented_Identity(".name") (part of Resource Identity/URL)
	// - serverTlsPolicyRef             -> f.SpecField(".server_tls_policy")
	// - sslCertificates                -> f.SpecField(".ssl_certificates")
	// - sslPolicyRef                   -> f.SpecField(".ssl_policy")
	// - urlMapRef                      -> f.SpecField(".url_map")
	//
	// KRM Status Fields:
	// - creationTimestamp              -> f.Unimplemented_NotYetTriaged(".creation_timestamp") (not mapped in ObservedState)
	// - proxyId                        -> f.Unimplemented_NotYetTriaged(".id") (not mapped in ObservedState)
	// - selfLink                       -> f.Unimplemented_NotYetTriaged(".self_link") (not mapped in ObservedState)
	// - observedState.fingerprint      -> f.StatusField(".fingerprint")

	// Spec fields
	f.SpecField(".certificate_map")
	f.SpecField(".description")
	f.SpecField(".http_keep_alive_timeout_sec")
	f.SpecField(".proxy_bind")
	f.SpecField(".quic_override")
	f.SpecField(".server_tls_policy")
	f.SpecField(".ssl_certificates")
	f.SpecField(".ssl_policy")
	f.SpecField(".url_map")

	// Status fields
	f.StatusField(".fingerprint")

	// Unimplemented / Identity fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".region")

	f.Unimplemented_Internal(".kind")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".authorization_policy")
	f.Unimplemented_NotYetTriaged(".tls_early_data")
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".self_link")

	f.FilterSpec = func(in *pb.TargetHttpsProxy) {
		if in.CertificateMap != nil && *in.CertificateMap != "" {
			val := *in.CertificateMap
			if !strings.HasPrefix(val, "//certificatemanager.googleapis.com/") {
				val = "//certificatemanager.googleapis.com/" + val
				in.CertificateMap = &val
			}
		}
	}
	f.FilterStatus = func(in *pb.TargetHttpsProxy) {
	}

	return f
}
