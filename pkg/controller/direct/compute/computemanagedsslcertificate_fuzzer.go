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
// proto.message: google.cloud.compute.v1.SslCertificate
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeManagedSSLCertificateFuzzer())
}

func computeManagedSSLCertificateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SslCertificate{},
		ComputeManagedSSLCertificateSpec_v1beta1_FromProto, ComputeManagedSSLCertificateSpec_v1beta1_ToProto,
		ComputeManagedSSLCertificateObservedState_v1beta1_FromProto, ComputeManagedSSLCertificateObservedState_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".managed")
	f.SpecField(".managed.domains")
	f.SpecField(".type")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".expire_time")
	f.StatusField(".id")
	f.StatusField(".self_link")
	f.StatusField(".subject_alternative_names")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".certificate")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".private_key")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".self_managed")

	return f
}
