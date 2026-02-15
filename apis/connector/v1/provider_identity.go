// Copyright 2025 Google LLC
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

package v1

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

// Based on google/cloud/connectors/v1/provider.proto

var ProviderIdentityFormat = gcpurls.Template[ProviderIdentity]("connectors.googleapis.com", "projects/{project}/locations/{location}/providers/{provider}")

// ProviderIdentity defines the resource reference to Provider, which "External" field
// holds the GCP identifier for the KRM object.
type ProviderIdentity struct {
	Project  string
	Location string
	Provider string
}

func (i *ProviderIdentity) String() string {
	return ProviderIdentityFormat.ToString(*i)
}

func (i *ProviderIdentity) ID() string {
	return i.Provider
}

func (i *ProviderIdentity) Parent() *ProviderParent {
	return &ProviderParent{
		ProjectID: i.Project,
		Location:  i.Location,
	}
}

func (i *ProviderIdentity) FromExternal(ref string) error {
	parsed, match, err := ProviderIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of Provider external=%q was not known (use %s): %w", ref, ProviderIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of Provider external=%q was not known (use %s)", ref, ProviderIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

// ProviderParent defines the Provider's parent type.
// No changes were needed in this file.
// The ProviderParent struct is correct as is, and so are the methods that use it.
type ProviderParent struct {
	ProjectID string
	Location  string
}

// String returns the fully qualified Provider in the format of projects/{{project_id}}/locations/{{location}}/providers/{{provider_id}}
func (p *ProviderParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// ParseProviderExternal parses the Provider's parent and ID from a string.
func ParseProviderExternal(external string) (parent *ProviderParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "providers" {
		return nil, "", fmt.Errorf("format of Provider external=%q was not known (use projects/{{projectID}}/locations/{{location}}/providers/{{providerID}})", external)
	}
	parent = &ProviderParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
