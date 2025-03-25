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
// proto.message: google.storage.control.v2.Folder
// api.group: storage.cnrm.cloud.google.com

package storage

import (
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(StorageFolderFuzzer())
}

func StorageFolderFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Folder{},
		StorageFolderSpec_FromProto, StorageFolderSpec_ToProto,
		StorageFolderObservedState_FromProto, StorageFolderObservedState_ToProto,
	)
	f.UnimplementedFields.Insert(".name")

	f.StatusFields.Insert(".metageneration")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".pending_rename_info")

	return f
}
