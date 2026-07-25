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
// proto.message: google.cloud.edgecontainer.v1.NodePool
// api.group: edgecontainer.cnrm.cloud.google.com

package edgecontainer

import (
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(edgeContainerNodePoolFuzzer())
}

func edgeContainerNodePoolFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NodePool{},
		EdgeContainerNodePoolSpec_FromProto, EdgeContainerNodePoolSpec_ToProto,
		NodePoolObservedState_FromProto, NodePoolObservedState_ToProto,
	)

	f.SpecFields.Insert(".node_location")
	f.SpecFields.Insert(".node_count")
	f.SpecFields.Insert(".machine_filter")
	f.SpecFields.Insert(".local_disk_encryption.kms_key")
	f.SpecFields.Insert(".node_config")
	f.SpecFields.Insert(".node_config.labels")
	f.SpecFields.Insert(".node_config.node_storage_schema")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".node_version")
	f.StatusFields.Insert(".local_disk_encryption.kms_key_active_version")
	f.StatusFields.Insert(".local_disk_encryption.kms_key_state")
	f.StatusFields.Insert(".local_disk_encryption.kms_status")
	f.StatusFields.Insert(".local_disk_encryption.resource_state")

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".local_disk_encryption.kms_status.details")

	return f
}
