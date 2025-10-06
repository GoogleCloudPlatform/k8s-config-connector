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
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TagBindingIdentity defines the resource reference to TagBinding, which "External" field
// holds the GCP identifier for the KRM object.
type TagBindingIdentity struct {
	parent   *TagBindingParent
	location string
	id       string
}

func (i *TagBindingIdentity) String() string {
	return i.id
}

func (i *TagBindingIdentity) ID() string {
	return i.id
}

func (i *TagBindingIdentity) Parent() *TagBindingParent {
	return i.parent
}

type TagBindingParent struct {
	parentResource string
}

func (p *TagBindingParent) String() string {
	return p.parentResource
}

// New builds a TagBindingIdentity from the Config Connector TagBinding object.
func NewTagBindingIdentity(ctx context.Context, reader client.Reader, obj *TagBinding) (*TagBindingIdentity, error) {

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
		location = "us"
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseTagBindingExternal(externalRef)
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
	return &TagBindingIdentity{
		parent: &TagBindingParent{
			parentResource: parent,
		},
		location: location,
		id:       resourceID,
	}, nil
}

func ParseTagBindingExternal(external string) (parent *TagBindingParent, resourceID string, err error) {
	if !strings.HasPrefix(external, "tagBindings/") {
		return nil, "", fmt.Errorf("invalid format of TagsLocationTagsLocationTagBinding external=%q, expected prefix 'tagBindings/'", external)
	}
	// This gives us "{parent}/tagValues/{tag_value}"
	suffix := strings.TrimPrefix(external, "tagBindings/")

	tokens := strings.Split(suffix, "/tagValues/")
	if len(tokens) != 2 {
		return nil, "", fmt.Errorf("format of TagsLocationTagsLocationTagBinding external=%q was not known (use tagBindings/{{ResourceName}}/tagValues/{{tagValue}}", external)
	}

	// The parent part of the externalRef has been URL-encoded by the framework, so slashes are represented as `%2F` not `/`.
	// We need re-endcode the parent. Then do the look up
	parentResource, err := url.PathUnescape(tokens[0])
	if err != nil {
		return nil, "", err
	}
	parent = &TagBindingParent{
		parentResource: parentResource,
	}

	resourceID = external
	return parent, resourceID, nil
}
