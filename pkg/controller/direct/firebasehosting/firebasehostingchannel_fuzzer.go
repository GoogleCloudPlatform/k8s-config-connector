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
// proto.message: google.firebasehosting.v1beta1.Channel
// api.group: firebasehosting.cnrm.cloud.google.com

package firebasehosting

import (
	pb "google.golang.org/api/firebasehosting/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(firebaseHostingChannelFuzzer())
}

func firebaseHostingChannelFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&pb.Channel{},
		FirebaseHostingChannelSpec_FromAPI, FirebaseHostingChannelSpec_ToAPI,
		FirebaseHostingChannelStatus_FromAPI, FirebaseHostingChannelStatus_ToAPI,
	)

	// Field Comparison & Mapping Status:
	// KRM FirebaseHostingChannelSpec Fields:
	// - expireTime maps to .ExpireTime
	// - resourceID: Standard KCC identity field, not an API field
	// - retainedReleaseCount maps to .RetainedReleaseCount
	// - siteId: Parent site ID in GCP REST URL path (sites/{siteId}/channels/{channelId}), not part of Channel struct
	// - ttl maps to .Ttl
	//
	// KRM FirebaseHostingChannelStatus Fields:
	// - conditions: Standard KCC Status condition list, not an API field
	// - externalRef: Standard KCC Status external reference, not an API field
	// - name maps to .Name
	// - observedGeneration: Standard KCC Status generation tracking, not an API field

	// Spec fields
	f.SpecField(".ExpireTime")
	f.SpecField(".RetainedReleaseCount")
	f.SpecField(".Ttl")

	// Status fields
	f.StatusField(".Name")

	// Untriaged / unimplemented API fields
	f.Unimplemented_NotYetTriaged(".CreateTime")
	f.Unimplemented_NotYetTriaged(".Labels")
	f.Unimplemented_NotYetTriaged(".Release")
	f.Unimplemented_NotYetTriaged(".UpdateTime")
	f.Unimplemented_NotYetTriaged(".Url")

	// JSON bookkeeping fields
	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	return f
}
