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
	RegisterKRMFuzzer(workstationFuzzer())
}

func workstationFuzzer() KRMFuzzer {
	f := NewKRMTypedFuzzer(&pb.Workstation{},
		workstations.WorkstationsWorkstationSpec_FromProto, workstations.WorkstationsWorkstationSpec_ToProto,
		workstations.WorkstationsWorkstationObservedState_FromProto, workstations.WorkstationsWorkstationObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")

	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".reconciling")
	f.UnimplementedFields.Insert(".degraded")
	f.UnimplementedFields.Insert(".conditions")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".annotations")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".start_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".host")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".uid")

	return f
}
