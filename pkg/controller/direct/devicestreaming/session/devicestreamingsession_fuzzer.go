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
// proto.message: google.cloud.devicestreaming.v1.DeviceSession
// api.group: devicestreaming.cnrm.cloud.google.com

package session

import (
	pb "cloud.google.com/go/devicestreaming/apiv1/devicestreamingpb"
	mapper "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/devicestreaming"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(devicestreamingSessionFuzzer())
}

func devicestreamingSessionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DeviceSession{},
		mapper.DeviceStreamingSessionSpec_FromProto, mapper.DeviceStreamingSessionSpec_ToProto,
		mapper.DeviceStreamingSessionObservedState_FromProto, mapper.DeviceStreamingSessionObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".expiration")

	f.SpecField(".ttl")
	f.SpecField(".expire_time")
	f.SpecField(".android_device")

	f.StatusField(".display_name")
	f.StatusField(".state")
	f.StatusField(".state_histories")
	f.StatusField(".inactivity_timeout")
	f.StatusField(".create_time")
	f.StatusField(".active_start_time")

	return f
}
