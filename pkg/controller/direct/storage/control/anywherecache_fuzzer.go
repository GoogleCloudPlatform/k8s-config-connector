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

	f.SpecFields.Insert(".zone")
	f.SpecFields.Insert(".ttl")
	f.SpecFields.Insert(".admission_policy")
	f.SpecFields.Insert(".desired_state")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".pending_update")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
