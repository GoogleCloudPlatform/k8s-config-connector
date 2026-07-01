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
// proto.message: google.cloud.eventarc.v1.ChannelConnection
// api.group: eventarc.cnrm.cloud.google.com

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(eventarcChannelConnectionFuzzer())
}

func eventarcChannelConnectionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ChannelConnection{},
		EventarcChannelConnectionSpec_FromProto, EventarcChannelConnectionSpec_ToProto,
		EventarcChannelConnectionObservedState_FromProto, EventarcChannelConnectionObservedState_ToProto,
	)

	// Spec fields
	f.SpecField(".channel")
	f.SpecField(".activation_token")

	// Status fields
	f.StatusField(".uid")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	// Unimplemented fields
	f.Unimplemented_Identity(".name")

	return f
}
