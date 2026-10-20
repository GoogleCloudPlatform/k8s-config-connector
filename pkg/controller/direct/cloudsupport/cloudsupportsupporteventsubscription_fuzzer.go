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
// proto.message: google.cloud.support.v2.SupportEventSubscription
// api.group: cloudsupport.cnrm.cloud.google.com

package cloudsupport

import (
	pb "cloud.google.com/go/support/apiv2/supportpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(supportEventSubscriptionFuzzer())
}

func supportEventSubscriptionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SupportEventSubscription{},
		CloudSupportSupportEventSubscriptionSpec_FromProto, CloudSupportSupportEventSubscriptionSpec_ToProto,
		CloudSupportSupportEventSubscriptionObservedState_FromProto, CloudSupportSupportEventSubscriptionObservedState_ToProto,
	)

	f.SpecField(".pub_sub_topic")

	f.StatusField(".state")
	f.StatusField(".failure_reason")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".delete_time")
	f.StatusField(".purge_time")

	f.Unimplemented_Identity(".name")

	return f
}
