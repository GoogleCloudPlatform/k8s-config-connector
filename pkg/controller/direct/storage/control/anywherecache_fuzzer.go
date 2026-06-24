// Copyright 2025 Google LLC
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
// proto.message: google.storage.control.v2.AnywhereCache
// api.group: storage.cnrm.cloud.google.com

package storagecontrol

import (
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(StorageAnywhereCacheFuzzer())
}

func StorageAnywhereCacheFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AnywhereCache{},
		StorageAnywhereCacheSpec_FromProto, StorageAnywhereCacheSpec_ToProto,
		StorageAnywhereCacheObservedState_FromProto, StorageAnywhereCacheObservedState_ToProto,
	)

	f.SpecField(".zone")
	f.SpecField(".ttl")
	f.SpecField(".admission_policy")
	f.SpecField(".desired_state")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".pending_update")

	f.Unimplemented_Identity(".name") // special field

	return f
}
