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

package v1beta1

import (
	"fmt"
	"strings"
)

// type: "gkehub.googleapis.com/Membership"
// pattern: "projects/{project}/locations/{location}/memberships/{membership_id}"
// parent_type: "gkehub.googleapis.com/Location"
// parent_name_extractor: "projects/{project}/locations/{location}"

// GKEHubMembershipIdentity defines the resource reference to GKEHubMembership, which "External" field
// holds the GCP identifier for the KRM object.
type GKEHubMembershipIdentity struct {
	parent *ParentIdentity
	id     string
}

func (i *GKEHubMembershipIdentity) String() string {
	return i.Parent() + "/memberships/" + i.id
}

func (i *GKEHubMembershipIdentity) ID() string {
	return i.id
}

func (i *GKEHubMembershipIdentity) Parent() string {
	return i.parent.String()
}

// ParentIdentity defines the parent of a Membership, which is a Location.
// Format: projects/{project}/locations/{location}
type ParentIdentity struct {
	Project  string
	Location string
}

func (p *ParentIdentity) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", p.Project, p.Location)
}

func ParseMembershipExternal(external string) (parent *ParentIdentity, resourceID string, err error) {
	// Custom parser for GKEHubMembership external ID
	// Format: projects/{project}/locations/{location}/memberships/{membership_id}
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "memberships" {
		return nil, "", fmt.Errorf("format of GKEHubMembership external=%q was not known (use projects/{{project}}/locations/{{location}}/memberships/{{membership_id}})", external)
	}
	parent = &ParentIdentity{
		Project:  tokens[1],
		Location: tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}

func (i *GKEHubMembershipIdentity) FromExternal(external string) error {
	parent, id, err := ParseMembershipExternal(external)
	if err != nil {
		return err
	}
	i.parent = parent
	i.id = id
	return nil
}

// Note: We do not implement NewGKEHubMembershipIdentity here yet because we want to avoid
// importing the specific KRM type if possible, or we will implement it when needed
// by looking up the generic unstructured object or the DCL type.
// For references, FromExternal is the most critical part.

// Helper to construct Identity from components
func NewGKEHubMembershipIdentity(project, location, membershipID string) *GKEHubMembershipIdentity {
	return &GKEHubMembershipIdentity{
		parent: &ParentIdentity{
			Project:  project,
			Location: location,
		},
		id: membershipID,
	}
}

// Common functions using "common" package
func (i *GKEHubMembershipIdentity) DefaultProjectState(project string) {
	if i.parent == nil {
		i.parent = &ParentIdentity{}
	}
	if i.parent.Project == "" {
		i.parent.Project = project
	}
}

func (i *GKEHubMembershipIdentity) DefaultLocationState(location string) {
	if i.parent == nil {
		i.parent = &ParentIdentity{}
	}
	if i.parent.Location == "" {
		i.parent.Location = location
	}
}
