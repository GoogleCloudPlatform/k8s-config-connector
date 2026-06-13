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
// proto.message: google.cloud.compute.v1.ServiceAttachment
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeServiceAttachmentFuzzer())
}

func computeServiceAttachmentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ServiceAttachment{},
		ComputeServiceAttachmentSpec_v1beta1_FromProto, ComputeServiceAttachmentSpec_v1beta1_ToProto,
		ComputeServiceAttachmentStatus_v1beta1_FromProto, ComputeServiceAttachmentStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".connection_preference")
	f.SpecField(".consumer_accept_lists")
	f.SpecField(".consumer_accept_lists[].connection_limit")
	f.SpecField(".consumer_accept_lists[].project_id_or_num")
	f.SpecField(".consumer_reject_lists")
	f.SpecField(".description")
	f.SpecField(".enable_proxy_protocol")
	f.SpecField(".nat_subnets")
	f.SpecField(".target_service")

	// Status fields
	f.StatusField(".connected_endpoints")
	f.StatusField(".connected_endpoints[].endpoint")
	f.StatusField(".connected_endpoints[].psc_connection_id")
	f.StatusField(".connected_endpoints[].status")
	f.StatusField(".fingerprint")
	f.StatusField(".id")
	f.StatusField(".psc_service_attachment_id")
	f.StatusField(".psc_service_attachment_id.high")
	f.StatusField(".psc_service_attachment_id.low")
	f.StatusField(".region")
	f.StatusField(".self_link")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".connected_endpoints[].consumer_network")
	f.Unimplemented_NotYetTriaged(".connected_endpoints[].endpoint_with_id")
	f.Unimplemented_NotYetTriaged(".connected_endpoints[].nat_ips")
	f.Unimplemented_NotYetTriaged(".connected_endpoints[].propagated_connection_count")
	f.Unimplemented_NotYetTriaged(".consumer_accept_lists[].endpoint_url")
	f.Unimplemented_NotYetTriaged(".consumer_accept_lists[].network_url")
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".domain_names")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".metadata")
	f.Unimplemented_NotYetTriaged(".producer_forwarding_rule")
	f.Unimplemented_NotYetTriaged(".propagated_connection_limit")
	f.Unimplemented_NotYetTriaged(".reconcile_connections")

	return f
}
