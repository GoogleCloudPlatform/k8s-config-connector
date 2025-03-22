// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.eventarc.v1.Channel
// api.group: eventarc.cnrm.cloud.google.com

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(eventarcChannelFuzzer())
}

func eventarcChannelFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Channel{},
		EventarcChannelSpec_FromProto, EventarcChannelSpec_ToProto,
		EventarcChannelObservedState_FromProto, EventarcChannelObservedState_ToProto,
	)

	f.SpecFields.Insert(".provider")
	f.SpecFields.Insert(".kms_key_name")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".pubsub_topic")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".activation_token")
	f.StatusFields.Insert(".satisfies_pzs")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
