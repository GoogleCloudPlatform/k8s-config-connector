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

package v1alpha1

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

// type: "networkservices.googleapis.com/WasmPlugin"
// pattern: "projects/{project}/locations/{location}/wasmPlugins/{wasm_plugin_id}"
// parent_type: "networkservices.googleapis.com/Location"
// parent_name_extractor: "projects/{project}/locations/{location}"

var wasmPluginURL = gcpurls.Template[NetworkServicesWasmPluginIdentity](
	"networkservices.googleapis.com",
	"projects/{projectID}/locations/{location}/wasmPlugins/{wasmPluginID}",
)

// NetworkServicesWasmPluginIdentity defines the resource reference to NetworkServicesWasmPlugin, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type NetworkServicesWasmPluginIdentity struct {
	ProjectID    string
	Location     string
	WasmPluginID string
}

func (i *NetworkServicesWasmPluginIdentity) String() string {
	return wasmPluginURL.ToString(*i)
}

func (i *NetworkServicesWasmPluginIdentity) ID() string {
	return i.WasmPluginID
}

func (i *NetworkServicesWasmPluginIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.ProjectID, i.Location)
}

func (i *NetworkServicesWasmPluginIdentity) FromExternal(external string) error {
	out, match, err := wasmPluginURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesWasmPlugin external=%q was not known (use %s)", external, wasmPluginURL.CanonicalForm())
	}
	*i = *out
	return nil
}

// Helper to construct Identity from components
func NewNetworkServicesWasmPluginIdentity(project, location, wasmPluginID string) *NetworkServicesWasmPluginIdentity {
	return &NetworkServicesWasmPluginIdentity{
		ProjectID:    project,
		Location:     location,
		WasmPluginID: wasmPluginID,
	}
}

// Common functions using "common" package
func (i *NetworkServicesWasmPluginIdentity) DefaultProjectState(project string) {
	if i.ProjectID == "" {
		i.ProjectID = project
	}
}

func (i *NetworkServicesWasmPluginIdentity) DefaultLocationState(location string) {
	if i.Location == "" {
		i.Location = location
	}
}
