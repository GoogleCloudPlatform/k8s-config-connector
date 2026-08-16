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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(mapConfigFuzzer())
}

func mapConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MapConfig{},
		MapManagementMapConfigSpec_FromProto, MapManagementMapConfigSpec_ToProto,
		MapManagementMapConfigObservedState_FromProto, MapManagementMapConfigObservedState_ToProto,
	)

	// Identity Field
	f.Unimplemented_Identity(".name")

	// Spec Fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".map_features")
	f.SpecField(".map_features.simple_features")
	f.SpecField(".map_features.poi_boost_level")
	f.SpecField(".map_type")

	// ObservedState Status Fields
	f.StatusField(".map_id")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
