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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EntryIdentity defines the resource reference to DataCatalogEntry, which "External" field
// holds the GCP identifier for the KRM object.
type EntryIdentity struct {
	parent *EntryParent
	id     string
}

// String returns the GCP resource name for the DataCatalogEntry.
func (i *EntryIdentity) String() string {
	return i.parent.String() + "/entries/" + i.id
}

// ID returns the resource ID part of the identity.
func (i *EntryIdentity) ID() string {
	return i.id
}

// Parent returns the parent identity.
func (i *EntryIdentity) Parent() *EntryParent {
	return i.parent
}

// EntryParent represents the parent (EntryGroup) of a DataCatalogEntry.
type EntryParent struct {
	ProjectID    string
	Location     string
	EntryGroupID string
}

// String returns the GCP resource name for the parent EntryGroup.
func (p *EntryParent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", p.ProjectID, p.Location, p.EntryGroupID)
}

// NewEntryIdentity builds a EntryIdentity from the Config Connector DataCatalogEntry object.
func NewEntryIdentity(ctx context.Context, reader client.Reader, obj *DataCatalogEntry) (*EntryIdentity, error) {
	// --- Determine Parent ---
	if obj.Spec.EntryGroupRef == nil || obj.Spec.EntryGroupRef.External == "" {
		// Based on the API structure (CreateEntry requires entry group parent),
		// EntryGroupRef is implicitly required.
		return nil, fmt.Errorf("spec.entryGroupRef.external is required to identify the parent EntryGroup")
	}

	// Parse parent info from the EntryGroup reference
	entryGroupParent, entryGroupID, err := ParseEntryGroupExternal(obj.Spec.EntryGroupRef.External)
	if err != nil {
		return nil, fmt.Errorf("cannot parse spec.entryGroupRef.external %q: %w", obj.Spec.EntryGroupRef.External, err)
	}

	// Optional: Validate against legacy fields if they exist.
	// This maintains consistency if users specify redundant fields, though EntryGroupRef should be authoritative.
	if obj.Spec.ProjectRef != nil {
		resolvedProjectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
		if err != nil {
			// Don't fail validation if resolution fails, maybe ref is invalid or project doesn't exist yet.
			// Log warning? For now proceed, EntryGroupRef is the primary source.
			_ = err // Placeholder
		} else if resolvedProjectRef.ProjectID != "" && resolvedProjectRef.ProjectID != entryGroupParent.ProjectID {
			return nil, fmt.Errorf("project ID mismatch between spec.entryGroupRef (%s) and spec.projectRef (%s)", entryGroupParent.ProjectID, resolvedProjectRef.ProjectID)
		}
	}
	if obj.Spec.Location != nil {
		if *obj.Spec.Location != entryGroupParent.Location {
			return nil, fmt.Errorf("location mismatch between spec.entryGroupRef (%s) and spec.location (%s)", entryGroupParent.Location, *obj.Spec.Location)
		}
	}

	// --- Determine Resource ID ---
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		// This case should ideally not happen in K8s with a named object
		return nil, fmt.Errorf("cannot resolve resource ID (metadata.name or spec.resourceID is empty)")
	}

	desiredParent := &EntryParent{
		ProjectID:    entryGroupParent.ProjectID,
		Location:     entryGroupParent.Location,
		EntryGroupID: entryGroupID,
	}

	// --- Validate against Status.ExternalRef if present ---
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Parse the full Entry external ref from status
		actualParent, actualResourceID, err := ParseEntryExternal(externalRef)
		if err != nil {
			// If status is set but unparsable, it might indicate a problem. Treat as error.
			return nil, fmt.Errorf("cannot parse status.externalRef %q: %w", externalRef, err)
		}

		// Validate desired state against actual state from status
		if actualParent.ProjectID != desiredParent.ProjectID {
			return nil, fmt.Errorf("parent project ID cannot be changed, expected %q (from status), got %q (from spec)", actualParent.ProjectID, desiredParent.ProjectID)
		}
		if actualParent.Location != desiredParent.Location {
			return nil, fmt.Errorf("parent location cannot be changed, expected %q (from status), got %q (from spec)", actualParent.Location, desiredParent.Location)
		}
		if actualParent.EntryGroupID != desiredParent.EntryGroupID {
			return nil, fmt.Errorf("parent entry group ID cannot be changed, expected %q (from status), got %q (from spec)", actualParent.EntryGroupID, desiredParent.EntryGroupID)
		}
		if actualResourceID != resourceID {
			// This validation prevents changing spec.resourceID or metadata.name after creation/status update.
			return nil, fmt.Errorf("resource ID cannot be changed, expected %q (from status), got %q (from spec.resourceID or metadata.name)", actualResourceID, resourceID)
		}
	}

	// --- Construct Identity ---
	return &EntryIdentity{
		parent: desiredParent,
		id:     resourceID,
	}, nil
}

// ParseEntryExternal parses the external format of a DataCatalogEntry resource name.
// Format: projects/{projectID}/locations/{location}/entryGroups/{entryGroupID}/entries/{entryID}
func ParseEntryExternal(external string) (parent *EntryParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	// Expect 8 tokens: projects/PROJECT/locations/LOCATION/entryGroups/ENTRYGROUP/entries/ENTRY
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "entryGroups" || tokens[6] != "entries" {
		return nil, "", fmt.Errorf("format of DataCatalogEntry external=%q was not known (use projects/{{projectID}}/locations/{{location}}/entryGroups/{{entryGroupID}}/entries/{{entryID}})", external)
	}
	parent = &EntryParent{
		ProjectID:    tokens[1],
		Location:     tokens[3],
		EntryGroupID: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
