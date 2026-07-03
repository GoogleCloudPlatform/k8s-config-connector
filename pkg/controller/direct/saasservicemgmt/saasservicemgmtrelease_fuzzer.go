// Copyright 2026 Google LLC
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
// proto.message: google.cloud.saasplatform.saasservicemgmt.v1beta1.Release
// api.group: saasservicemgmt.cnrm.cloud.google.com

package saasservicemgmt

import (
	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(saasServiceMgmtReleaseFuzzer())
}

func saasServiceMgmtReleaseFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Release{},
		SaasServiceMgmtReleaseSpec_FromProto, SaasServiceMgmtReleaseSpec_ToProto,
		SaasServiceMgmtReleaseObservedState_FromProto, SaasServiceMgmtReleaseObservedState_ToProto,
	)

	// Identity fields that are not in KRM fields
	f.Unimplemented_Identity(".name")

	// Spec fields to fuzz
	f.SpecField(".unit_kind")
	f.SpecField(".blueprint")
	f.SpecField(".release_requirements")
	f.SpecField(".input_variable_defaults")

	// Status fields (ObservedState) to fuzz
	f.StatusField(".blueprint")
	f.StatusField(".input_variables")
	f.StatusField(".output_variables")
	f.StatusField(".uid")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	// Unimplemented or ignore fields
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_LabelsAnnotations(".annotations")
	f.Unimplemented_Etag()
	f.Unimplemented_NotYetTriaged(".application_template_component")

	return f
}
