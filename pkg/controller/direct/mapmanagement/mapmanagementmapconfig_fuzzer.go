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
	fuzztesting.RegisterKRMFuzzer(fuzzMapManagementMapConfig())
}

func fuzzMapManagementMapConfig() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MapConfig{},
		MapManagementMapConfigSpec_FromProto, MapManagementMapConfigSpec_ToProto,
		MapManagementMapConfigObservedState_FromProto, MapManagementMapConfigObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".map_id")
	f.Unimplemented_Identity(".create_time")
	f.Unimplemented_Identity(".update_time")

	return f
}
