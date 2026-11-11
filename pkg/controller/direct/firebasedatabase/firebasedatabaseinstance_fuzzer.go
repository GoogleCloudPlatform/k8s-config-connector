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
// proto.message: google.firebasedatabase.v1beta.DatabaseInstance
// api.group: firebasedatabase.cnrm.cloud.google.com

package firebasedatabase

import (
	pb "google.golang.org/api/firebasedatabase/v1beta"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(firebaseDatabaseInstanceFuzzer())
}

func firebaseDatabaseInstanceFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&pb.DatabaseInstance{},
		FirebaseDatabaseInstanceSpec_FromAPI, FirebaseDatabaseInstanceSpec_ToAPI,
		FirebaseDatabaseInstanceStatus_FromAPI, FirebaseDatabaseInstanceStatus_ToAPI,
	)

	// Field Comparison & Mapping Status:
	// KRM Spec Fields:
	// - desiredState: Controls runtime lifecycle state in GCP via update/patch requests, not a direct field of DatabaseInstance
	// - projectRef: References parent GCP Project, not an API field (corresponds to identity .Project)
	// - region: Location for resource creation/identity, not a separate field of DatabaseInstance
	// - resourceID: Resource identity, not an API field
	// - type maps to .Type
	//
	// KRM Status Fields:
	// - conditions: Standard KCC Status condition list, not an API field
	// - databaseUrl maps to .DatabaseUrl
	// - externalRef: Standard KCC Status external reference, not an API field
	// - name maps to .Name
	// - observedGeneration: Standard KCC Status generation tracking, not an API field
	// - state maps to .State

	// Spec fields
	f.SpecField(".Type")

	// Status fields
	f.StatusField(".DatabaseUrl")
	f.StatusField(".Name")
	f.StatusField(".State")

	// Identity fields (API resource parent project identifier)
	f.IdentityField(".Project")

	// JSON bookkeeping fields
	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	return f
}
