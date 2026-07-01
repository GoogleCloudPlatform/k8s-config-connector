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
// proto.message: google.cloud.eventarc.v1.GoogleApiSource
// api.group: eventarc.cnrm.cloud.google.com

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(eventarcGoogleApiSourceFuzzer())
}

func eventarcGoogleApiSourceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.GoogleApiSource{},
		EventarcGoogleAPISourceSpec_FromProto, EventarcGoogleAPISourceSpec_ToProto,
		EventarcGoogleAPISourceObservedState_FromProto, EventarcGoogleAPISourceObservedState_ToProto,
	)

	f.IdentityField(".name")

	f.SpecField(".display_name")
	f.SpecField(".destination")
	f.SpecField(".crypto_key_name")
	f.SpecField(".logging_config")

	f.StatusField(".uid")
	f.StatusField(".etag")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_LabelsAnnotations(".annotations")

	f.Unimplemented_NotYetTriaged(".project_subscriptions")
	f.Unimplemented_NotYetTriaged(".organization_subscription")

	return f
}
