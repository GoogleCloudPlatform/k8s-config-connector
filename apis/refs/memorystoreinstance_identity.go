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

package refs

import (
	"fmt"
	"strings"
)

// MemorystoreInstanceIdentity defines the resource reference to MemorystoreInstance, which "External" field
// holds the GCP identifier for the KRM object.
type MemorystoreInstanceIdentity struct {
	Parent_ *MemorystoreInstanceParent
	ID_     string
}

func (i *MemorystoreInstanceIdentity) String() string {
	return i.Parent_.String() + "/instances/" + i.ID_
}

func (i *MemorystoreInstanceIdentity) ID() string {
	return i.ID_
}

func (i *MemorystoreInstanceIdentity) Parent() *MemorystoreInstanceParent {
	return i.Parent_
}

type MemorystoreInstanceParent struct {
	ProjectID string
	Location  string
}

func (p *MemorystoreInstanceParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func ParseInstanceExternal(external string) (parent *MemorystoreInstanceParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "instances" {
		return nil, "", fmt.Errorf("format of MemorystoreInstance external=%q was not known (use projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}})", external)
	}
	parent = &MemorystoreInstanceParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
