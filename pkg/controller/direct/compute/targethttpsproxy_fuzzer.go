// Copyright 2025 Google LLC
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

	// KRM-only spec fields
	f.SpecFields.Insert(".location")
	f.SpecFields.Insert(".resourceID")
	f.SpecFields.Insert(".projectRef")

	// KRM-only status fields
	f.StatusFields.Insert(".observedGeneration")
	f.StatusFields.Insert(".externalRef")

	// Volatile status fields (also in status but not in observedState)
	f.StatusFields.Insert(".creationTimestamp")
	f.StatusFields.Insert(".proxyId")
	f.StatusFields.Insert(".selfLink")

	// Unimplemented proto fields
	f.UnimplementedFields.Insert(".id")
	f.UnimplementedFields.Insert(".kind")
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".region")
	f.UnimplementedFields.Insert(".authorization_policy")
	f.UnimplementedFields.Insert(".tls_early_data")

	f.FilterSpec = func(in *pb.TargetHttpsProxy) {
		in.CreationTimestamp = nil
		in.Id = nil
		in.SelfLink = nil
		in.Fingerprint = nil
	}
	f.FilterStatus = func(in *pb.TargetHttpsProxy) {
		in.CertificateMap = nil
		in.HttpKeepAliveTimeoutSec = nil
		in.ProxyBind = nil
		in.QuicOverride = nil
		in.ServerTlsPolicy = nil
		in.SslCertificates = nil
		in.SslPolicy = nil
		in.UrlMap = nil
		in.Description = nil
		in.CreationTimestamp = nil
		in.Id = nil
		in.SelfLink = nil
	}

	return f
}
