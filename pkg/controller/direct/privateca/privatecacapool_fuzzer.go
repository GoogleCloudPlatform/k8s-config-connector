// Copyright 2025 Google LLC
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
// proto.message: google.cloud.security.privateca.v1.CaPool
// api.group: privateca.cnrm.cloud.google.com

package privateca

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(privatecaCAPoolFuzzer())
}

func privatecaCAPoolFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CaPool{},
		PrivateCACAPoolSpec_FromProto, PrivateCACAPoolSpec_ToProto,
		PrivateCACAPoolObservedState_FromProto, PrivateCACAPoolObservedState_ToProto,
	)

	f.IdentityField(".name")

	f.SpecField(".issuance_policy")
	f.SpecField(".tier")
	f.SpecField(".publishing_options")

	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".issuance_policy.backdate_duration")
	f.Unimplemented_NotYetTriaged(".issuance_policy.baseline_values.name_constraints")

	f.Unimplemented_NotYetTriaged(".publishing_options.encoding_format")

	return f
}
