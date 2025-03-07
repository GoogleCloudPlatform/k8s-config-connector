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

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ Parent = &ProjectAndLocationParent{}

type ProjectAndLocationParent struct {
	ProjectID string
	Location  string
}

func (p *ProjectAndLocationParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func (p *ProjectAndLocationParent) MatchActual(actualI Parent) error {
	actual, ok := actualI.(*ProjectAndLocationParent)
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

func (i *ProjectAndLocationParent) FromExternal(external string) error {
	tokens := strings.Split(external, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		i.ProjectID = tokens[1]
		i.Location = tokens[3]
		return nil
	}
	return fmt.Errorf("format of project external=%q was not known (use projects/{projectID}/locations/{location})", external)
}

var _ ParentBuilder = &ProjectAndLocationRef{}

// ProjectAndLocationParent specifies the resource's GCP hierarchy (Project/Folder/Organization) and its geographical location.
// +kubebuilder:object:generate:=true
type ProjectAndLocationRef struct {
	// +required
	ProjectRef *ProjectRef `json:"projectRef"`

	// +required
	Location string `json:"location"`
}

func (p *ProjectAndLocationRef) Build(ctx context.Context, reader client.Reader, othernamespace string, parent Parent) error {
	projectAndLocation, ok := parent.(*ProjectAndLocationParent)
	if !ok {
		return fmt.Errorf("build invalid parent, except %T", &ProjectAndLocationParent{})
	}
	project := new(ProjectParent)
	if err := p.ProjectRef.Build(ctx, reader, othernamespace, project); err != nil {
		return err
	}
	if project.ProjectID == "" {
		return fmt.Errorf("cannot resolve project")
	}
	projectAndLocation.ProjectID = project.ProjectID
	projectAndLocation.Location = p.Location
	return nil
}

func ParseProjectAndLocationParent(external string) (*ProjectAndLocationParent, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, fmt.Errorf("format of ProjectAndLocation external=%q was not known (use projects/<projectId>/locations/<location>)", external)
	}

	return &ProjectAndLocationParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}, nil
}
