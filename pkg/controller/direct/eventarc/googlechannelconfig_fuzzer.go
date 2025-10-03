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

// +tool:fuzz-gen
// proto.message: google.cloud.eventarc.v1.GoogleChannelConfig
// api.group: eventarc.cnrm.cloud.google.com

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(eventarcGoogleChannelConfigFuzzer())
}

func eventarcGoogleChannelConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.GoogleChannelConfig{},
		EventarcGoogleChannelConfigSpec_FromProto, EventarcGoogleChannelConfigSpec_ToProto,
		EventarcGoogleChannelConfigObservedState_FromProto, EventarcGoogleChannelConfigObservedState_ToProto,
	)

	f.SpecFields.Insert(".crypto_key_name")

	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
