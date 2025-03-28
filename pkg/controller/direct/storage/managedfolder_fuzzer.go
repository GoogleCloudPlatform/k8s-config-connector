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
// proto.message: google.storage.control.v2.ManagedFolder
// api.group: storage.cnrm.cloud.google.com

package storage

import (
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(storageManagedFolderFuzzer())
}

func storageManagedFolderFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ManagedFolder{},
		StorageManagedFolderSpec_FromProto, StorageManagedFolderSpec_ToProto,
		StorageManagedFolderObservedState_FromProto, StorageManagedFolderObservedState_ToProto,
	)

	f.StatusFields.Insert(".metageneration")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
