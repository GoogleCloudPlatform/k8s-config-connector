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
// proto.message: google.cloud.healthcare.v1.ConsentStore
// api.group: healthcare.cnrm.cloud.google.com

package healthcare

import (
	pb "google.golang.org/api/healthcare/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(healthcareConsentStoreFuzzer())
}

func healthcareConsentStoreFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&pb.ConsentStore{},
		HealthcareConsentStoreSpec_FromAPI, HealthcareConsentStoreSpec_ToAPI,
		HealthcareConsentStoreStatus_FromAPI, HealthcareConsentStoreStatus_ToAPI,
	)

	// Field Comparison & Mapping Status:
	// KRM Spec Fields:
	// - dataset: Identifies parent dataset (KCC hierarchical identity field, not in ConsentStore body)
	// - defaultConsentTtl maps to .DefaultConsentTtl
	// - enableConsentCreateOnUpdate maps to .EnableConsentCreateOnUpdate
	// - resourceID: Resource identity (Standard KCC identity field, not a direct API body field)
	//
	// KRM Status Fields:
	// - conditions: Standard KCC Status condition list, not an API field
	// - observedGeneration: Standard KCC Status generation tracking, not an API field

	// Spec fields
	f.SpecField(".DefaultConsentTtl")
	f.SpecField(".EnableConsentCreateOnUpdate")

	// Identity fields (API resource name)
	f.IdentityField(".Name")

	// JSON bookkeeping fields
	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	// Untriaged / unimplemented API fields
	f.Unimplemented_NotYetTriaged(".Labels")

	return f
}
