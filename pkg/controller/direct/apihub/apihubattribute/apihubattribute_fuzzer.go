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

package apihubattribute

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apihubAttributeFuzzer())
}

func apihubAttributeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Attribute{},
		apihub.APIHubAttributeSpec_FromProto, apihub.APIHubAttributeSpec_ToProto,
		apihub.APIHubAttributeObservedState_FromProto, apihub.APIHubAttributeObservedState_ToProto,
	)

	// Identity Field
	f.Unimplemented_Identity(".name")

	// Spec Fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".scope")
	f.SpecField(".data_type")
	f.SpecField(".allowed_values")
	f.SpecField(".cardinality")

	// Status Fields
	f.StatusField(".definition_type")
	f.StatusField(".mandatory")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.FilterSpec = func(p *pb.Attribute) {
		if p.Cardinality == 0 {
			p.Cardinality = 1
		}
	}

	return f
}
