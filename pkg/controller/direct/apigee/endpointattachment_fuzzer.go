// Copyright 2024 Google LLC
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
        api "google.golang.org/api/apigee/v1"
        "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
        fuzztesting.RegisterKRMFuzzer_NoProto(endpointAttachmentFuzzer())
}

func endpointAttachmentFuzzer() fuzztesting.KRMFuzzer_NoProto {
        f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.GoogleCloudApigeeV1EndpointAttachment{},
                ApigeeEndpointAttachmentSpec_FromAPI, ApigeeEndpointAttachmentSpec_ToAPI,
                ApigeeEndpointAttachmentObservedState_FromAPI, ApigeeEndpointAttachmentObservedState_ToAPI,
        )

        f.SpecField(".Location")
        f.SpecField(".ServiceAttachment")
        
        f.StatusField(".ConnectionState")
        f.StatusField(".Host")
        f.StatusField(".State")

        f.Unimplemented_NotYetTriaged(".Name")
        f.Unimplemented_NotYetTriaged(".ForceSendFields")
        f.Unimplemented_NotYetTriaged(".NullFields")
        f.Unimplemented_NotYetTriaged(".ServerResponse")

        return f
}
