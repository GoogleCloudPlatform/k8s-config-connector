// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package parent

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ Parent = &ProjectAndLocation{}
var _ identity.Identity = &ProjectAndLocation{}

type ProjectAndLocation struct {
	ProjectID string
	Location  string
}

func (p *ProjectAndLocation) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func (p *ProjectAndLocation) URL() string {
	return "projects/{{projectId}}/locations/{{location}}"
}

func (p *ProjectAndLocation) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return fmt.Errorf("format of ProjectAndLocation external=%q was not known (use %s)", ref, ProjectAndLocationURL)
	}
	p.ProjectID = tokens[1]
	p.Location = tokens[3]
	return nil
}

func (p *ProjectAndLocation) MatchActual(actualI Parent) error {
	actual, ok := actualI.(*ProjectAndLocation)
	if !ok {
		return fmt.Errorf("parent format changed, desired %T", p)
	}
	if p.ProjectID != actual.ProjectID {
		return fmt.Errorf("spec.projectRef changed, desired %s, actual %s", p.ProjectID, actual.ProjectID)
	}
	if p.Location != actual.Location {
		return fmt.Errorf("spec.location changed, desired %s, actual %s", p.Location, actual.Location)
	}
	return nil
}

var _ refsv1beta1.ExternalNormalizer = &ProjectAndLocationRef{}

// ProjectAndLocation specifies the resource's GCP hierarchy (Project/Folder/Organization) and its geographical location.
// +kubebuilder:object:generate:=true
type ProjectAndLocationRef struct {
	// +required
	ProjectRef *ProjectRef `json:"projectRef"`

	// +required
	Location string `json:"location"`
}

func (p *ProjectAndLocationRef) NormalizedExternal(ctx context.Context, reader client.Reader, othernamespace string) (string, error) {
	return nil
}
