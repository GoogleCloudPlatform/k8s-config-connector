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

package fuzztests

import (
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/workstations"
)

func init() {
	RegisterKRMFuzzer(workstationConfigFuzzer())
}

func workstationConfigFuzzer() KRMFuzzer {
	f := NewKRMTypedFuzzer(&pb.WorkstationConfig{},
		workstations.WorkstationsWorkstationConfigSpec_FromProto, workstations.WorkstationsWorkstationConfigSpec_ToProto,
		workstations.WorkstationsWorkstationConfigObservedState_FromProto, workstations.WorkstationsWorkstationConfigObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")

	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".reconciling")
	f.UnimplementedFields.Insert(".degraded")
	f.UnimplementedFields.Insert(".conditions")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".running_timeout")
	f.SpecFields.Insert(".host")
	f.SpecFields.Insert(".persistent_directories")
	f.SpecFields.Insert(".container")
	f.SpecFields.Insert(".idle_timeout")
	f.SpecFields.Insert(".replica_zones")
	f.SpecFields.Insert(".readiness_checks")
	f.SpecFields.Insert(".encryption_key")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".uid")

	return f
}
