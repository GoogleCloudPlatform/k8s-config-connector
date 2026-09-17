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
// proto.message: google.cloud.aiplatform.v1.PersistentResource
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(aiplatformPersistentResourceFuzzer())
}

func aiplatformPersistentResourceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.PersistentResource{},
		AIPlatformPersistentResourceSpec_FromProto, AIPlatformPersistentResourceSpec_ToProto,
		AIPlatformPersistentResourceObservedState_FromProto, AIPlatformPersistentResourceObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".resource_pools[].machine_spec.gpu_partition_size")
	f.Unimplemented_NotYetTriaged(".resource_pools[].used_replica_count")
	f.Unimplemented_NotYetTriaged(".resource_runtime.access_uris")
	f.Unimplemented_NotYetTriaged(".error.details")

	f.SpecField(".display_name")
	f.SpecField(".resource_pools")
	f.SpecField(".labels")
	f.SpecField(".network")
	f.SpecField(".psc_interface_config")
	f.SpecField(".encryption_spec")
	f.SpecField(".resource_runtime_spec")
	f.SpecField(".reserved_ip_ranges")

	f.StatusField(".state")
	f.StatusField(".error")
	f.StatusField(".create_time")
	f.StatusField(".start_time")
	f.StatusField(".update_time")
	f.StatusField(".resource_runtime")

	return f
}
