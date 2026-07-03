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

package apihubplugin

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apihubPluginFuzzer())
}

func apihubPluginFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Plugin{},
		apihub.APIHubPluginSpec_FromProto, apihub.APIHubPluginSpec_ToProto,
		apihub.APIHubPluginObservedState_FromProto, apihub.APIHubPluginObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")
	f.StatusField(".state")

	f.Unimplemented_NotYetTriaged(".ownership_type")
	f.Unimplemented_NotYetTriaged(".hosting_service")
	f.Unimplemented_NotYetTriaged(".actions_config")
	f.Unimplemented_NotYetTriaged(".documentation")
	f.Unimplemented_NotYetTriaged(".plugin_category")
	f.Unimplemented_NotYetTriaged(".config_template")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".gateway_type")

	f.Unimplemented_NotYetTriaged(".description")
	f.Unimplemented_NotYetTriaged(".display_name")
	f.Unimplemented_NotYetTriaged(".type.attribute")
	f.Unimplemented_NotYetTriaged(".type.enum_values")
	f.Unimplemented_NotYetTriaged(".type.enum_values.values")
	f.Unimplemented_NotYetTriaged(".type.json_values")
	f.Unimplemented_NotYetTriaged(".type.json_values.values")
	f.Unimplemented_NotYetTriaged(".type.string_values")
	f.Unimplemented_NotYetTriaged(".type.uri_values")
	f.Unimplemented_NotYetTriaged(".type.uri_values.values")

	return f
}
