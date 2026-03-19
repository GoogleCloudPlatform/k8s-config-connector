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
// proto.message: google.cloud.memorystore.v1.Instance
// api.group: memorystore.cnrm.cloud.google.com

package memorystore

import (
	pb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(memorystoreInstanceBackupFuzzer())
}

func memorystoreInstanceBackupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Backup{},
		MemorystoreInstanceBackupSpec_FromProto, MemorystoreInstanceBackupSpec_ToProto,
		MemorystoreInstanceBackupObservedState_FromProto, MemorystoreInstanceBackupObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // Special field: resource name
	f.UnimplementedFields.Insert(".encryption_info")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".instance")
	f.StatusFields.Insert(".instance_uid")
	f.StatusFields.Insert(".total_size_bytes")
	f.StatusFields.Insert(".expire_time")
	f.StatusFields.Insert(".engine_version")
	f.StatusFields.Insert(".backup_files")
	f.StatusFields.Insert(".node_type")
	f.StatusFields.Insert(".replica_count")
	f.StatusFields.Insert(".shard_count")
	f.StatusFields.Insert(".backup_type")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".uid")

	return f
}
