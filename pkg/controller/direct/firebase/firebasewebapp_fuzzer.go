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
// proto.message: google.firebase.v1beta1.WebApp
// api.group: firebase.cnrm.cloud.google.com

package firebase

import (
	pb "google.golang.org/api/firebase/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(firebaseWebAppFuzzer())
}

func firebaseWebAppFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&pb.WebApp{},
		FirebaseWebAppSpec_FromAPI, FirebaseWebAppSpec_ToAPI,
		FirebaseWebAppStatus_FromAPI, FirebaseWebAppStatus_ToAPI,
	)

	// Field Comparison & Mapping Status:
	// KRM Spec Fields:
	// - apiKeyId maps to .ApiKeyId
	// - deletionPolicy: KCC deletion policy, not an API field
	// - displayName maps to .DisplayName
	// - project: References parent GCP Project, not an API field (corresponds to identity .ProjectId)
	// - resourceID: Resource identity, not an API field
	//
	// KRM Status Fields:
	// - appId maps to .AppId
	// - appUrls maps to .AppUrls
	// - conditions: Standard KCC Status condition list, not an API field
	// - externalRef: Standard KCC Status external reference, not an API field
	// - name maps to .Name
	// - observedGeneration: Standard KCC Status generation tracking, not an API field

	// Spec fields
	f.SpecField(".ApiKeyId")
	f.SpecField(".DisplayName")

	// Status fields
	f.StatusField(".AppId")
	f.StatusField(".AppUrls")
	f.StatusField(".Name")

	// Identity fields (API resource parent project identifier)
	f.IdentityField(".ProjectId")

	// JSON bookkeeping fields
	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	// Untriaged / unimplemented API fields
	f.Unimplemented_NotYetTriaged(".Etag")
	f.Unimplemented_NotYetTriaged(".ExpireTime")
	f.Unimplemented_NotYetTriaged(".State")
	f.Unimplemented_NotYetTriaged(".WebId")

	return f
}
