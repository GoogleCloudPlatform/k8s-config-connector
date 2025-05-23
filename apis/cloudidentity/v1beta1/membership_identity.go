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
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// type: "cloudidentity.googleapis.com/Membership"
// pattern: "groups/{group_id}/memberships/{member_id}"
// parent_type: "cloudidentity.googleapis.com/Group"
// parent_name_extractor: "groups/{group_id}"

// MembershipIdentity defines the resource reference to CloudIdentityMembership, which "External" field
// holds the GCP identifier for the KRM object.
type MembershipIdentity struct {
	parent *GroupIdentity
	id     string
}

func (i *MembershipIdentity) String() string {
	return i.Parent() + "/memberships/" + i.id
}

func (i *MembershipIdentity) ID() string {
	return i.id
}

func (i *MembershipIdentity) Parent() string {
	return i.parent.String()
}

// New builds a MembershipIdentity from the Config Connector Membership object.
func NewMembershipIdentity(ctx context.Context, reader client.Reader, obj *CloudIdentityMembership) (*MembershipIdentity, error) {
	// Get Parent
	groupExternal, err := obj.Spec.GroupRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	groupID, err := ParseGroupExternal(groupExternal)
	if err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseMembershipExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.id != groupID {
			return nil, fmt.Errorf("spec.groupRef changed, expect %s, got %s", actualParent.id, groupID)
		}
		if resourceID != "" && actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
		resourceID = actualResourceID
	}
	return &MembershipIdentity{
		parent: &GroupIdentity{
			id: groupID,
		},
		id: resourceID,
	}, nil
}

func ParseMembershipExternal(external string) (parent *GroupIdentity, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "groups" || tokens[2] != "memberships" {
		return nil, "", fmt.Errorf("format of CloudIdentityMembership external=%q was not known (use groups/{{groupID}}/memberships/{{membershipID}})", external)
	}
	parent = &GroupIdentity{
		id: tokens[1],
	}
	resourceID = tokens[3]
	return parent, resourceID, nil
}
