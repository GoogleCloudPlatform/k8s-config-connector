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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TagBindingIdentity defines the resource reference to TagsLocationTagBinding, which "External" field
// holds the GCP identifier for the KRM object.
type TagsLocationTagBindingIdentity struct {
	parent   *TagsLocationTagBindingParent
	location string
	id       string
}

func (i *TagsLocationTagBindingIdentity) String() string {
	return i.id
}

func (i *TagsLocationTagBindingIdentity) ID() string {
	return i.id
}

func (i *TagsLocationTagBindingIdentity) Location() string {
	return i.location
}

type TagsLocationTagBindingParent struct {
	parentResource string
}

func (i *TagsLocationTagBindingIdentity) Parent() *TagsLocationTagBindingParent {
	return i.parent
}

func (p *TagsLocationTagBindingParent) String() string {
	return p.parentResource
}

// New builds a TagsLocationTagBindingIdentity from the Config Connector TagsLocationTagBinding object.
func NewTagsLocationTagBindingIdentity(ctx context.Context, reader client.Reader, obj *TagsLocationTagBinding) (*TagsLocationTagBindingIdentity, error) {

	// Get the resource name of the resource that the tag binds to.
	// Uses the "spec.ParentRef"
	var parent string

	if obj.Spec.ParentRef.External != "" {
		parent = obj.Spec.ParentRef.External
	} else {
		return nil, fmt.Errorf("cannot resolve parentRef")
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Get desired Location
	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseTagsLocationTagBindingExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.parentResource != parent {
			return nil, fmt.Errorf("spec.parentRef changed, expect %s, got %s", actualParent.parentResource, parent)
		}

		if resourceID != "" && actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
		if resourceID == "" {
			resourceID = actualResourceID
		}
	}
	return &TagsLocationTagBindingIdentity{
		parent: &TagsLocationTagBindingParent{
			parentResource: parent,
		},
		location: location,
		id:       resourceID,
	}, nil
}

func ParseTagsLocationTagBindingExternal(external string) (parent *TagsLocationTagBindingParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "tagBindings" || tokens[3] != "tagValues" {
		return nil, "", fmt.Errorf("format of TagsLocationTagsLocationTagBinding external=%q was not known (use tagBindings/{{ResourceName}}/tagValues/{{tagValue}}", external)
	}

	parent = &TagsLocationTagBindingParent{
		parentResource: tokens[1],
	}

	resourceID = external
	return parent, resourceID, nil
}
