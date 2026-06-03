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
// proto.message: google.cloud.networkservices.v1.WasmPlugin
// api.group: networkservices.cnrm.cloud.google.com

package wasmplugin

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(wasmPluginFuzzer())
}

func wasmPluginFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.WasmPlugin{},
		NetworkServicesWasmPluginSpec_FromProto, NetworkServicesWasmPluginSpec_ToProto,
		NetworkServicesWasmPluginObservedState_FromProto, NetworkServicesWasmPluginObservedState_ToProto,
	)

	f.SpecField(".projectRef")
	f.SpecField(".location")
	f.SpecField(".resourceID")

	f.IdentityField(".name")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".used_by")

	f.SpecField(".description")
	f.SpecField(".main_version_id")
	f.SpecField(".log_config")
	f.SpecField(".versions")

	// Labels are handled by the controller from metadata.labels
	f.Unimplemented_LabelsAnnotations(".labels")

	f.FilterSpec = func(in *pb.WasmPlugin) {
		for _, v := range in.GetVersions() {
			v.CreateTime = nil
			v.UpdateTime = nil
			v.ImageDigest = ""
			v.PluginConfigDigest = ""
		}
	}

	return f
}
