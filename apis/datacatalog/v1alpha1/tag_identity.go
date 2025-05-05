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
	// refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1" // No longer needed directly
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TagIdentity defines the resource reference to DataCatalogTag, which "External" field
// holds the GCP identifier for the KRM object.
type TagIdentity struct {
	parent *TagParent
	id     string
}

// String returns the full GCP resource name for the Tag.
func (i *TagIdentity) String() string {
	// Format: projects/{project}/locations/{location}/entryGroups/{entry_group}/entries/{entry}/tags/{tagID}
	return i.parent.String() + "/tags/" + i.id
}

// ID returns the Tag ID (the last segment of the resource name).
func (i *TagIdentity) ID() string {
	return i.id
}

// Parent returns the TagParent struct containing the parent hierarchy information.
func (i *TagIdentity) Parent() *TagParent {
	return i.parent
}

// TagParent represents the parent of a DataCatalogTag, which is a DataCatalogEntry.
// It holds the components of the DataCatalogEntry's resource name.
type TagParent struct {
	ProjectID    string
	LocationID   string
	EntryGroupID string
	EntryID      string
}

// String returns the GCP resource name for the parent DataCatalogEntry.
func (p *TagParent) String() string {
	// Format: projects/{project}/locations/{location}/entryGroups/{entry_group}/entries/{entry}
	return fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s",
		p.ProjectID, p.LocationID, p.EntryGroupID, p.EntryID)
}

// NewTagIdentity builds a TagIdentity from the Config Connector Tag object.
func NewTagIdentity(ctx context.Context, reader client.Reader, obj *DataCatalogTag) (*TagIdentity, error) {
	// --- Determine Parent ---
	if obj.Spec.EntryRef == nil || (obj.Spec.EntryRef.External == "" && obj.Spec.EntryRef.Name == "") {
		// Based on the API structure (CreateEntry requires entry group parent),
		// EntryGroupRef is implicitly required.
		return nil, fmt.Errorf("spec.entryRef.external or spec.entryRef.name is required to identify the parent Entry")
	}

	if obj.Spec.EntryRef.External == "" {
		entryRef, err := obj.Spec.EntryRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return nil, fmt.Errorf("resolving entry reference: %w", err)
		}
		obj.Spec.EntryRef.External = entryRef
	}

	// Parse parent info from the Entry reference
	entryParent, entryID, err := ParseEntryExternal(obj.Spec.EntryRef.External)
	if err != nil {
		return nil, fmt.Errorf("cannot parse spec.entryRef.external %q: %w", obj.Spec.EntryRef.External, err)
	}

	// Get desired ID
	// This resource has a server-generated ID. This means user should not know
	// the ID before the resource is created, and 'metadata.name' won't be used
	// as the default resource ID. So empty value for 'spec.resourceID' should
	// also be valid:
	// 1. When 'spec.resourceID' is not set or set to an empty value, the
	//    intention is to create the resource.
	// 2. When 'spec.resourceID' is set, the intention is to acquire an existing
	//    resource.
	//    2.1. When 'spec.resourceID' is set but the corresponding GCP resource
	//         is not found, then it is a real error.
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	parent := &TagParent{
		ProjectID:    entryParent.ProjectID,
		LocationID:   entryParent.Location,
		EntryGroupID: entryParent.EntryGroupID,
		EntryID:      entryID,
	}

	// Use approved External (status.externalRef) for validation if available
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired parent and resourceID with actual values derived from status.externalRef
		actualParent, actualResourceID, err := ParseTagExternal(externalRef)
		if err != nil {
			// If parsing fails, the externalRef might be malformed or represent an unexpected state.
			// Proceeding might lead to incorrect state assumptions. It's safer to error out.
			return nil, fmt.Errorf("failed to parse existing status.externalRef %q: %w", externalRef, err)
		}

		// Check if the parent components derived from spec match the parent components from status
		if actualParent.ProjectID != parent.ProjectID ||
			actualParent.LocationID != parent.LocationID ||
			actualParent.EntryGroupID != parent.EntryGroupID ||
			actualParent.EntryID != parent.EntryID {
			// This indicates an attempt to reparent the tag, which is likely unsupported or unintended.
			return nil, fmt.Errorf("parent derived from spec.entryRef (%s) conflicts with parent from status.externalRef (%s)", parent.String(), actualParent.String())
		}

		// Check if the resourceID derived from spec/metadata matches the resourceID from status
		if resourceID != "" && actualResourceID != resourceID {
			// This indicates an attempt to change the tag's ID after it has been established in GCP.
			// The `spec.resourceID` or `metadata.name` determines the *desired* ID.
			// The `status.externalRef` reflects the *actual* ID in GCP. Changing this is not allowed.
			return nil, fmt.Errorf("cannot change tag ID from %q (derived from status.externalRef) to %q (derived from spec.resourceID/metadata.name)",
				actualResourceID, resourceID)
		}
		resourceID = actualResourceID
	}

	// If validation passes or status.externalRef is empty, return the identity based on spec.
	return &TagIdentity{
		parent: parent,
		id:     resourceID,
	}, nil
}

// ParseTagExternal parses the full GCP resource name string for a DataCatalogTag
// into its parent components and the tag ID.
func ParseTagExternal(external string) (parent *TagParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	// Expected format: projects/{project}/locations/{location}/entryGroups/{entry_group}/entries/{entry}/tags/{tagID}
	// indices:         0        1        2          3          4            5            6       7      8     9
	if len(tokens) != 10 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "entryGroups" || tokens[6] != "entries" || tokens[8] != "tags" {
		// Provide a more specific error message including the expected format.
		return nil, "", fmt.Errorf("format of DataCatalogTag externalRef %q is invalid, expected format: projects/{projectID}/locations/{locationID}/entryGroups/{entryGroupID}/entries/{entryID}/tags/{tagID}", external)
	}
	parent = &TagParent{
		ProjectID:    tokens[1],
		LocationID:   tokens[3], // Use LocationID consistently
		EntryGroupID: tokens[5],
		EntryID:      tokens[7],
	}
	resourceID = tokens[9]
	return parent, resourceID, nil
}
