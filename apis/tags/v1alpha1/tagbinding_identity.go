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

// TagsTagBindingIdentity defines the resource reference to TagsTagBinding, which "External" field
// holds the GCP identifier for the KRM object.
type TagsTagBindingIdentity struct {
	parent   *TagsTagBindingParent
	location string
	id       string
}

func (i *TagsTagBindingIdentity) String() string {
	return i.id
}

func (i *TagsTagBindingIdentity) ID() string {
	return i.id
}

func (i *TagsTagBindingIdentity) Parent() *TagsTagBindingParent {
	return i.parent
}

type TagsTagBindingParent struct {
	parentResource string
}

func (p *TagsTagBindingParent) String() string {
	return p.parentResource
}

// New builds a TagsTagBindingIdentity from the Config Connector TagsTagBinding object.
func NewTagsTagBindingIdentity(ctx context.Context, reader client.Reader, obj *TagsTagBinding) (*TagsTagBindingIdentity, error) {

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
		actualParent, actualResourceID, err := ParseTagsTagBindingExternal(externalRef)
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
	return &TagsTagBindingIdentity{
		parent: &TagsTagBindingParent{
			parentResource: parent,
		},
		location: location,
		id:       resourceID,
	}, nil
}

func ParseTagsTagBindingExternal(external string) (parent *TagsTagBindingParent, resourceID string, err error) {
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
	parent = &TagsTagBindingParent{
		parentResource: parentResource,
	}

	resourceID = external
	return parent, resourceID, nil
}
