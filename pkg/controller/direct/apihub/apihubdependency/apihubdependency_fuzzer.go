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

package apihubdependency

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apihubDependencyFuzzer())
}

func apihubDependencyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Dependency{},
		apihub.APIHubDependencySpec_FromProto, apihub.APIHubDependencySpec_ToProto,
		apihub.APIHubDependencyObservedState_FromProto, apihub.APIHubDependencyObservedState_ToProto,
	)

	// Identity Field
	f.Unimplemented_Identity(".name")

	// Spec Fields
	f.SpecField(".consumer")
	f.SpecField(".supplier")
	f.SpecField(".description")
	f.SpecField(".attributes")

	f.Unimplemented_NotYetTriaged(".consumer.display_name")
	f.Unimplemented_NotYetTriaged(".supplier.display_name")

	// Unimplemented Attributes Subfields (Attributes is a map)
	f.Unimplemented_NotYetTriaged(".attributes")
	f.Unimplemented_NotYetTriaged(".attributes[].attribute")
	f.Unimplemented_NotYetTriaged(".attributes[].string_values")
	f.Unimplemented_NotYetTriaged(".attributes[].json_values")
	f.Unimplemented_NotYetTriaged(".attributes[].uri_values")
	f.Unimplemented_NotYetTriaged(".attributes[].enum_values.values[].display_name")
	f.Unimplemented_NotYetTriaged(".attributes[].enum_values.values[].description")
	f.Unimplemented_NotYetTriaged(".attributes[].enum_values.values[].immutable")

	// Status Fields
	f.StatusField(".state")
	f.StatusField(".discovery_mode")
	f.StatusField(".error_detail")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
