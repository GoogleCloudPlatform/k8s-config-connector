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

package parent

import (
	"fmt"
	"strings"
)

var _ Parent = &ComputeParent{}

type ComputeParent struct {
	ProjectID string
	Location  string
}

func (p *ComputeParent) String() string {
	if p.Location == "global" {
		return "projects/" + p.ProjectID + "/global"
	} else {
		return "projects/" + p.ProjectID + "/regions/" + p.Location
	}
}

func (p *ComputeParent) MatchActual(actualI Parent) error {
	actual, ok := actualI.(*ComputeParent)
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

func ParseComputeParent(parent string) (*ComputeParent, error) {
	tokens := strings.Split(strings.TrimPrefix(parent, "/"), "/")
	if len(tokens) == 3 && tokens[0] == "projects" && tokens[2] == "global" {
		return &ComputeParent{ProjectID: tokens[1], Location: "global"}, nil
	} else if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "regions" {
		return &ComputeParent{ProjectID: tokens[1], Location: tokens[3]}, nil
	}
	return nil, fmt.Errorf("format of Compute parent %s was not known (use projects/{{projectID}}/global or projects/{{projectID}}/regions/{{region}})", parent)
}
