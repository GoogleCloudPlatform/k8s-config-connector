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
// proto.message: google.cloud.compute.v1.NodeTemplate
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeNodeTemplateFuzzer())
}

func computeNodeTemplateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NodeTemplate{},
		ComputeNodeTemplateSpec_v1beta1_FromProto, ComputeNodeTemplateSpec_v1beta1_ToProto,
		ComputeNodeTemplateStatus_v1beta1_FromProto, ComputeNodeTemplateStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".cpu_overcommit_type")
	f.SpecField(".description")
	f.SpecField(".node_type")
	f.SpecField(".node_type_flexibility")
	f.SpecField(".region")
	f.SpecField(".server_binding")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")

	// Unimplemented fields
	f.Unimplemented_Identity(".id")
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".kind")

	// Fields not yet triaged or unimplemented
	f.Unimplemented_NotYetTriaged(".accelerators")
	f.Unimplemented_NotYetTriaged(".disks")
	f.Unimplemented_NotYetTriaged(".node_affinity_labels")
	f.Unimplemented_NotYetTriaged(".status")
	f.Unimplemented_NotYetTriaged(".status_message")

	return f
}
