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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

// type: "gkehub.googleapis.com/Membership"
// pattern: "projects/{project}/locations/{location}/memberships/{membership_id}"
// parent_type: "gkehub.googleapis.com/Location"
// parent_name_extractor: "projects/{project}/locations/{location}"

var membershipURL = gcpurls.Template[GKEHubMembershipIdentity](
	"gkehub.googleapis.com",
	"projects/{projectID}/locations/{location}/memberships/{membershipID}",
)

// GKEHubMembershipIdentity defines the resource reference to GKEHubMembership, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type GKEHubMembershipIdentity struct {
	ProjectID    string
	Location     string
	MembershipID string
}

func (i *GKEHubMembershipIdentity) String() string {
	return membershipURL.ToString(*i)
}

func (i *GKEHubMembershipIdentity) ID() string {
	return i.MembershipID
}

func (i *GKEHubMembershipIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.ProjectID, i.Location)
}

func (i *GKEHubMembershipIdentity) FromExternal(external string) error {
	out, match, err := membershipURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of GKEHubMembership external=%q was not known (use %s)", external, membershipURL.CanonicalForm())
	}
	*i = *out
	return nil
}

// Helper to construct Identity from components
func NewGKEHubMembershipIdentity(project, location, membershipID string) *GKEHubMembershipIdentity {
	return &GKEHubMembershipIdentity{
		ProjectID:    project,
		Location:     location,
		MembershipID: membershipID,
	}
}

// Common functions using "common" package
func (i *GKEHubMembershipIdentity) DefaultProjectState(project string) {
	if i.ProjectID == "" {
		i.ProjectID = project
	}
}

func (i *GKEHubMembershipIdentity) DefaultLocationState(location string) {
	if i.Location == "" {
		i.Location = location
	}
}
