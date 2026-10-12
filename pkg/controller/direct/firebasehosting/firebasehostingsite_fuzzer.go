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

package firebasehosting

import (
	pb "google.golang.org/api/firebasehosting/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(firebaseHostingSiteFuzzer())
}

func firebaseHostingSiteFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&pb.Site{},
		FirebaseHostingSiteSpec_FromAPI, FirebaseHostingSiteSpec_ToAPI,
		FirebaseHostingSiteStatus_FromAPI, FirebaseHostingSiteStatus_ToAPI,
	)

	// Spec fields
	// KRM FirebaseHostingSiteSpec:
	// - AppID maps to .AppId
	// - ProjectRef and ResourceID are standard KCC resource identity/reference fields
	f.SpecField(".AppId")

	// Status fields
	// KRM FirebaseHostingSiteStatus:
	// - DefaultURL maps to .DefaultUrl
	// - Name maps to .Name
	f.StatusField(".DefaultUrl")
	f.StatusField(".Name")

	// Fields not implemented in KRM Spec/Status yet
	f.Unimplemented_NotYetTriaged(".Labels")
	f.Unimplemented_NotYetTriaged(".Type")

	return f
}
