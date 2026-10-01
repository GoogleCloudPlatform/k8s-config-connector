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

package mapmanagement

import (
	pb "cloud.google.com/go/maps/mapmanagement/apiv2beta/mapmanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/mapmanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(styleConfigFuzzer())
}

func styleConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*pb.StyleConfig, krm.MapManagementStyleConfigSpec, krm.MapManagementStyleConfigObservedState](
		&pb.StyleConfig{},
		MapManagementStyleConfigSpec_FromProto,
		MapManagementStyleConfigSpec_ToProto,
		MapManagementStyleConfigObservedState_FromProto,
		MapManagementStyleConfigObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".json_style_sheet")

	f.StatusField(".style_id")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.IdentityField(".name")
	return f
}
