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

package session

import (
	pb "cloud.google.com/go/devicestreaming/apiv1/devicestreamingpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzDeviceStreamingSession())
}

func fuzzDeviceStreamingSession() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DeviceSession{},
		DeviceStreamingSessionSpec_FromProto, DeviceStreamingSessionSpec_ToProto,
		DeviceStreamingSessionObservedState_FromProto, DeviceStreamingSessionObservedState_ToProto,
	)
	f.UnimplementedFields.Insert(".name")

	f.SpecFields.Insert(".android_device")
	f.SpecFields.Insert(".android_device.android_model_id")
	f.SpecFields.Insert(".android_device.android_version_id")
	f.SpecFields.Insert(".android_device.locale")
	f.SpecFields.Insert(".android_device.orientation")
	f.SpecFields.Insert(".ttl")
	f.SpecFields.Insert(".expire_time")

	f.StatusFields.Insert(".display_name")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_histories")
	f.StatusFields.Insert(".inactivity_timeout")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".active_start_time")

	return f
}
