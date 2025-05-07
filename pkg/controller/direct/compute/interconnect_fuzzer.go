// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.compute.v1.Interconnect
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeInterconnectFuzzer())
}

func computeInterconnectFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Interconnect{},
		ComputeInterconnectSpec_FromProto, ComputeInterconnectSpec_ToProto,
		ComputeInterconnectObservedState_FromProto, ComputeInterconnectObservedState_ToProto,
	)

	f.SpecFields.Insert(".customer_name")
	f.SpecFields.Insert(".label_fingerprint")
	f.SpecFields.Insert(".macsec_enabled")
	f.SpecFields.Insert(".noc_contact_email")
	f.SpecFields.Insert(".remote_location")
	f.SpecFields.Insert(".requested_link_count")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".macsec")
	f.SpecFields.Insert(".requested_features")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".location")
	f.SpecFields.Insert(".admin_enabled")
	f.SpecFields.Insert(".interconnect_type")
	f.SpecFields.Insert(".link_type")

	f.StatusFields.Insert(".available_features")
	f.StatusFields.Insert(".expected_outages")
	f.StatusFields.Insert(".operational_status")
	f.StatusFields.Insert(".provisioned_link_count")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".circuit_infos")
	f.StatusFields.Insert(".kind")
	f.StatusFields.Insert(".interconnect_attachments")
	f.StatusFields.Insert(".satisfies_pzs")
	f.StatusFields.Insert(".google_ip_address")
	f.StatusFields.Insert(".google_reference_id")
	f.StatusFields.Insert(".id")
	f.StatusFields.Insert(".self_link")
	f.StatusFields.Insert(".creation_timestamp")
	f.StatusFields.Insert(".peer_ip_address")

	f.UnimplementedFields.Insert(".name") // special field
	return f
}
