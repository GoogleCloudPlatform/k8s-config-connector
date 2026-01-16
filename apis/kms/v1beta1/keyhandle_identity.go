// Copyright 2024 Google LLC
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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +k8s:deepcopy-gen=false
type KMSKeyHandleIdentity struct {
	id     string
	parent *KMSKeyHandleParent
}

func (i *KMSKeyHandleIdentity) String() string {
	return i.parent.String() + "/keyHandles/" + i.id
}

func (r *KMSKeyHandleIdentity) ID() string {
	return r.id
}

func (r *KMSKeyHandleIdentity) Parent() *KMSKeyHandleParent {
	return r.parent
}

type KMSKeyHandleParent struct {
	ProjectID string
	Location  string
}

func (p *KMSKeyHandleParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func NewKMSKeyHandleIdentity(ctx context.Context, reader client.Reader, obj *KMSKeyHandle) (*KMSKeyHandleIdentity, error) {
	id := &KMSKeyHandleIdentity{}

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := valueOf(obj.Spec.Location)
	id.parent = &KMSKeyHandleParent{ProjectID: projectID, Location: location}

	// Get desired ID
	resourceID := valueOf(obj.Spec.ResourceID)

	// Use approved External
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualParent, actualHandleID, err := ParseKMSKeyHandleExternal(externalRef)
		if err != nil {
			return nil, err
		}
		// Validate desired with actual
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if resourceID != "" && actualHandleID != resourceID {
			return nil, fmt.Errorf("cannot reset `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualHandleID)
		}
		resourceID = actualHandleID
	}
	return &KMSKeyHandleIdentity{
		parent: &KMSKeyHandleParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseKMSKeyHandleExternal(external string) (parent *KMSKeyHandleParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "keyHandles" {
		return nil, "", fmt.Errorf("format of KMSKeyHandle external=%q was not known (use projects/{{projectId}}/locations/{{location}}/keyHandles/{{keyhandleID}})", external)
	}
	parent = &KMSKeyHandleParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}

// var _ identity.Identity = &KMSKeyHandleIdentity{} // Tracking in issue #6073
