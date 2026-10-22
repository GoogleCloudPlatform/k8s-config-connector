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

package apigee

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/apigee/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(syncAuthorizationFuzzer())
}

func syncAuthorizationFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.GoogleCloudApigeeV1SyncAuthorization{},
		ApigeeSyncAuthorizationSpec_FromAPI, ApigeeSyncAuthorizationSpec_ToAPI,
		ApigeeSyncAuthorizationStatus_FromAPI, ApigeeSyncAuthorizationStatus_ToAPI,
	)

	// Field comparison:
	// KRM ApigeeSyncAuthorizationSpec:
	// - Identities -> .Identities
	// - ResourceID is the standard KCC resource ID
	// KRM ApigeeSyncAuthorizationStatus:
	// - Etag -> .Etag
	// - Conditions and ObservedGeneration are standard KCC status fields
	f.SpecField(".Identities")
	f.StatusField(".Etag")

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	return f
}
