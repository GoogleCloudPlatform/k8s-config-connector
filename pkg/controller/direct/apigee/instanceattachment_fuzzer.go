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
	fuzztesting.RegisterKRMFuzzer_NoProto(instanceAttachmentFuzzer())
}

func instanceAttachmentFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.GoogleCloudApigeeV1InstanceAttachment{},
		ApigeeInstanceAttachmentSpec_FromAPI, ApigeeInstanceAttachmentSpec_ToAPI,
		ApigeeInstanceAttachmentObservedState_FromAPI, ApigeeInstanceAttachmentObservedState_ToAPI,
	)

	f.SpecField(".Environment")
	f.StatusField(".CreatedAt")

	f.Unimplemented_NotYetTriaged(".Name")
	f.Unimplemented_NotYetTriaged(".ForceSendFields")
	f.Unimplemented_NotYetTriaged(".NullFields")
	f.Unimplemented_NotYetTriaged(".ServerResponse")

	f.FilterStatus = func(in *api.GoogleCloudApigeeV1InstanceAttachment) {
		// time.RFC3339 format drops the milliseconds, so we zero them to pass roundtrip.
		// also keep it within a reasonable range for time parsing.
		in.CreatedAt = in.CreatedAt % 253402300799000
		if in.CreatedAt < 0 {
			in.CreatedAt = -in.CreatedAt
		}
		in.CreatedAt = (in.CreatedAt / 1000) * 1000
	}

	return f
}
