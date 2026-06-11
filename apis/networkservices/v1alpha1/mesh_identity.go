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

// type: "networkservices.googleapis.com/Mesh"
// pattern: "projects/{project}/locations/{location}/meshes/{mesh_id}"
// parent_type: "networkservices.googleapis.com/Location"
// parent_name_extractor: "projects/{project}/locations/{location}"

var meshURL = gcpurls.Template[NetworkServicesMeshIdentity](
	"networkservices.googleapis.com",
	"projects/{projectID}/locations/{location}/meshes/{meshID}",
)

// NetworkServicesMeshIdentity is the identity of a NetworkServicesMesh.
// +k8s:deepcopy-gen=false
type NetworkServicesMeshIdentity struct {
	ProjectID string
	Location  string
	MeshID    string
}

func (i *NetworkServicesMeshIdentity) String() string {
	return meshURL.ToString(*i)
}

func (i *NetworkServicesMeshIdentity) ID() string {
	return i.MeshID
}

func (i *NetworkServicesMeshIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.ProjectID, i.Location)
}

func (i *NetworkServicesMeshIdentity) FromExternal(external string) error {
	out, match, err := meshURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesMesh external=%q was not known (use %s)", external, meshURL.CanonicalForm())
	}
	*i = *out
	return nil
}

// Helper to construct Identity from components
func NewNetworkServicesMeshIdentity(project, location, meshID string) *NetworkServicesMeshIdentity {
	return &NetworkServicesMeshIdentity{
		ProjectID: project,
		Location:  location,
		MeshID:    meshID,
	}
}

// Common functions using "common" package
func (i *NetworkServicesMeshIdentity) DefaultProjectState(project string) {
	if i.ProjectID == "" {
		i.ProjectID = project
	}
}

func (i *NetworkServicesMeshIdentity) DefaultLocationState(location string) {
	if i.Location == "" {
		i.Location = location
	}
}
