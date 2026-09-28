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
// proto.message: google.cloud.lustre.v1.Instance
// api.group: lustre.cnrm.cloud.google.com

package lustre

import (
	pb "cloud.google.com/go/lustre/apiv1/lustrepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(lustreInstanceFuzzer())
}

func lustreInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		LustreInstanceSpec_FromProto, LustreInstanceSpec_ToProto,
		LustreInstanceObservedState_FromProto, LustreInstanceObservedState_ToProto,
	)

	f.SpecField(".filesystem")
	f.SpecField(".capacity_gib")
	f.SpecField(".network")
	f.SpecField(".description")
	f.SpecField(".per_unit_storage_throughput")

	f.StatusField(".state")
	f.StatusField(".mount_point")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".gke_support_enabled")

	return f
}
