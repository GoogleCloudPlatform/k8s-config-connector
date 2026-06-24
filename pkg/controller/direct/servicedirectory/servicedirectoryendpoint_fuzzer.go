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

package servicedirectory

import (
	pb "cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzServiceDirectoryEndpoint())
}

func fuzzServiceDirectoryEndpoint() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Endpoint{},
		ServiceDirectoryEndpointSpec_FromProto, ServiceDirectoryEndpointSpec_ToProto,
		ServiceDirectoryEndpointObservedState_FromProto, ServiceDirectoryEndpointObservedState_ToProto,
	)

	// Field Comparison Table (KRM Spec/Status vs GCP Proto Fields):
	//
	// KRM Spec Field       | GCP Proto Field | Fuzzer/Proto Field Mapping Status
	// ---------------------|-----------------|-----------------------------------
	// serviceRef           | -               | Ignored (maps to parent service path/identity)
	// addressRef           | .address        | f.SpecField(".address")
	// networkRef           | .network        | f.SpecField(".network")
	// port                 | .port           | f.SpecField(".port")
	// resourceID           | -               | Ignored (KCC resource ID identity)
	//
	// KRM Status Field     | GCP Proto Field | Fuzzer/Proto Field Mapping Status
	// ---------------------|-----------------|-----------------------------------
	// name                 | .name           | f.Unimplemented_Identity(".name")
	//
	// Unimplemented Proto  | GCP Proto Field | Fuzzer/Proto Field Mapping Status
	// ---------------------|-----------------|-----------------------------------
	// metadata             | .metadata       | f.Unimplemented_NotYetTriaged(".metadata")
	// createTime           | .create_time    | f.Unimplemented_NotYetTriaged(".create_time")
	// updateTime           | .update_time    | f.Unimplemented_NotYetTriaged(".update_time")
	// uid                  | .uid            | f.Unimplemented_NotYetTriaged(".uid")

	// Spec fields
	f.SpecField(".address")
	f.SpecField(".port")
	f.SpecField(".network")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".uid")
	f.Unimplemented_NotYetTriaged(".metadata")

	return f
}
