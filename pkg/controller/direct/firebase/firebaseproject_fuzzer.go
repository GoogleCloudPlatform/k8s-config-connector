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
// proto.message: google.firebase.v1beta1.FirebaseProject
// api.group: firebase.cnrm.cloud.google.com

package firebase

import (
	pb "google.golang.org/api/firebase/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(firebaseProjectFuzzer())
}

func firebaseProjectFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&pb.FirebaseProject{},
		FirebaseProjectSpec_FromAPI, FirebaseProjectSpec_ToAPI,
		FirebaseProjectStatus_FromAPI, FirebaseProjectStatus_ToAPI,
	)

	// Field Comparison & Mapping Status:
	// KRM Spec Fields:
	// - projectRef: References GCP Project (Standard KCC reference/identity field, not a direct API field)
	// - resourceID: Resource identity (Standard KCC identity field, not a direct API field)
	//
	// KRM Status Fields:
	// - displayName maps to .DisplayName
	// - projectNumber maps to .ProjectNumber
	// - conditions: Standard KCC Status condition list, not an API field
	// - observedGeneration: Standard KCC Status generation tracking, not an API field

	// Status fields
	f.StatusField(".DisplayName")
	f.StatusField(".ProjectNumber")

	// Identity fields (API resource name & unique identifier)
	f.IdentityField(".Name")
	f.IdentityField(".ProjectId")

	// JSON bookkeeping fields
	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	// Untriaged / unimplemented API fields
	f.Unimplemented_NotYetTriaged(".Annotations")
	f.Unimplemented_NotYetTriaged(".Etag")
	f.Unimplemented_NotYetTriaged(".Resources")
	f.Unimplemented_NotYetTriaged(".State")

	return f
}
