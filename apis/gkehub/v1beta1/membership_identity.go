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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
)

// MembershipIdentity defines the resource reference to GKEHubMembership, which "External" field
// holds the GCP identifier for the KRM object.
type MembershipIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *MembershipIdentity) String() string {
	return i.parent.String() + "/memberships/" + i.id
}

func (i *MembershipIdentity) ID() string {
	return i.id
}

func (i *MembershipIdentity) Parent() *parent.ProjectAndLocationParent { return i.parent }

func ParseMembershipExternal(external string) (*MembershipIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "memberships" {
		return nil, fmt.Errorf("format of GKEHubMembership external=%q was not known (use projects/{{projectID}}/locations/{{location}}/memberships/{{membershipID}})", external)
	}
	return &MembershipIdentity{parent: &parent.ProjectAndLocationParent{ProjectID: tokens[1], Location: tokens[3]}, id: tokens[5]}, nil
}
