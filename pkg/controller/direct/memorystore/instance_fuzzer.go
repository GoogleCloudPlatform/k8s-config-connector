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
// proto.message: google.cloud.memorystore.v1beta.Instance
// api.group: memorystore.cnrm.cloud.google.com

package memorystore

import (
	pb "cloud.google.com/go/memorystore/apiv1beta/memorystorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(memorystoreInstanceFuzzer())
}

func memorystoreInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		MemorystoreInstanceSpec_FromProto, MemorystoreInstanceSpec_ToProto,
		MemorystoreInstanceObservedState_FromProto, MemorystoreInstanceObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // Special field: resource name
	f.UnimplementedFields.Insert(".satisfies_pzi")
	f.UnimplementedFields.Insert(".satisfies_pzs")
	f.UnimplementedFields.Insert(".psc_auto_connections")
	f.UnimplementedFields.Insert(".discover_endpoints")
	f.UnimplementedFields.Insert(".node_config") // Handled in status, but not spec.
	// The `port` field in PscAutoConnection is a repeated field in proto
	// but a singular field in KRM, making it not round-trippable.
	f.UnimplementedFields.Insert(".endpoints[].connections[].psc_auto_connection.port")
	// The `state_info` struct in KRM is empty, meaning its subfields are not propagated.
	f.UnimplementedFields.Insert(".state_info.update_info")

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".replica_count")
	f.SpecFields.Insert(".authorization_mode")
	f.SpecFields.Insert(".transit_encryption_mode")
	f.SpecFields.Insert(".shard_count")
	f.SpecFields.Insert(".node_type")
	f.SpecFields.Insert(".persistence_config")
	f.SpecFields.Insert(".engine_version")
	f.SpecFields.Insert(".engine_configs")
	f.SpecFields.Insert(".zone_distribution_config")
	f.SpecFields.Insert(".deletion_protection_enabled")
	f.SpecFields.Insert(".endpoints")
	f.SpecFields.Insert(".mode")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_info")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".node_config") // NodeConfig is an output-only field in the proto, but KRM needs it for status.
	f.StatusFields.Insert(".endpoints")

	return f
}
