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

type KMSKeyHandleIdentity struct {
	id     string
	parent *KMSKeyHandleParent
}

func (i *KMSKeyHandleIdentity) String() string {
	return i.parent.String() + "/keyHandles/" + i.id
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

func asKMSKeyHandleExternal(parent *KMSKeyHandleParent, resourceID string) (external string) {
	return parent.String() + "/keyHandles/" + resourceID
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
	desiredHandleID := valueOf(obj.Spec.ResourceID)
	// set the desired keyhandle id in the identity
	id.id = desiredHandleID

	// At this point we are expecting desiredHandleID to be either empty or valid uuid
	// 1. if desiredHandleID empty:
	// id.external will be projects/{{pid}}/locations/{{loc}}/keyHandles/. i.e without resourceID.
	// A call will be made to find() with invalid externalID which will return false.
	// 2. if desiredHandleID is a valid UUID: id.external will be valid.

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
		if desiredHandleID != "" && (actualHandleID != desiredHandleID) {
			return nil, fmt.Errorf("cannot reset `spec.resourceID` to %s, since it has already assigned to %s",
				desiredHandleID, actualHandleID)
		}
		return id, nil
	}
	return id, nil
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

func (r *KMSKeyHandleIdentity) KeyHandleID() (string, bool) {
	return r.id, r.id != ""
}

func asKMSKeyHandleExternal_FromSpec(spec *KMSKeyHandleSpec) (parent *KMSKeyHandleParent, resourceID string, err error) {
	external := strings.TrimPrefix(spec.ProjectRef.External, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, "", fmt.Errorf("invalid projectRef found in KMSKeyHandle=%q was not known (use projects/{{projectId}})", external)
	}
	parent = &KMSKeyHandleParent{
		ProjectID: tokens[1],
		Location:  valueOf(spec.Location),
	}
	return parent, valueOf(spec.ResourceID), nil
}
