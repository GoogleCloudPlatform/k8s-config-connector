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

// FeedIdentity defines the resource reference to AssetFeed, which "External" field
// holds the GCP identifier for the KRM object.
type FeedIdentity struct {
	parent *FeedParent
	id     string
}

func (i *FeedIdentity) String() string {
	if i.parent == nil || i.parent.String() == "" {
		return "" // Or handle error appropriately
	}
	return i.parent.String() + "/feeds/" + i.id
}

func (i *FeedIdentity) ID() string {
	return i.id
}

func (i *FeedIdentity) Parent() *FeedParent {
	return i.parent
}

// FeedParent represents the parent of an AssetFeed, which can be a Project, Folder, or Organization.
type FeedParent struct {
	ProjectID      string
	FolderID       string
	OrganizationID string
}

// String returns the string representation of the FeedParent.
// Format: "projects/{projectID}" or "folders/{folderID}" or "organizations/{organizationID}"
func (p *FeedParent) String() string {
	if p.ProjectID != "" {
		return "projects/" + p.ProjectID
	}
	if p.FolderID != "" {
		return "folders/" + p.FolderID
	}
	if p.OrganizationID != "" {
		return "organizations/" + p.OrganizationID
	}
	// Should not happen if constructed correctly
	return ""
}

// NewFeedIdentity builds a FeedIdentity from the Config Connector Feed object.
func NewFeedIdentity(ctx context.Context, reader client.Reader, obj *AssetFeed) (*FeedIdentity, error) {
	desiredParent := &FeedParent{}
	var err error

	// Try parsing parent from Spec.Name first
	// Example Name formats:
	// projects/{project_number}/feeds/{feed_id}
	// folders/{folder_number}/feeds/{feed_id}
	// organizations/{organization_number}/feeds/{feed_id}
	if obj.Spec.Name != nil && *obj.Spec.Name != "" {
		nameTokens := strings.Split(*obj.Spec.Name, "/")
		// Check for formats like "type/id/feeds/feed_id" (4 segments)
		if len(nameTokens) == 4 && nameTokens[2] == "feeds" {
			switch nameTokens[0] {
			case "projects":
				desiredParent.ProjectID = nameTokens[1]
			case "folders":
				desiredParent.FolderID = nameTokens[1]
			case "organizations":
				desiredParent.OrganizationID = nameTokens[1]
				// else: invalid format in Name field, fallback below
			}
		}
		// Allow short formats like "type/id" for parent only, if name doesn't include /feeds/ part
		if len(nameTokens) == 2 && desiredParent.ProjectID == "" && desiredParent.FolderID == "" && desiredParent.OrganizationID == "" {
			switch nameTokens[0] {
			case "projects":
				desiredParent.ProjectID = nameTokens[1]
			case "folders":
				desiredParent.FolderID = nameTokens[1]
			case "organizations":
				desiredParent.OrganizationID = nameTokens[1]
			}
		}
	}

	// If parent not determined from Name, fallback to ProjectRef (implies project parent)
	if desiredParent.ProjectID == "" && desiredParent.FolderID == "" && desiredParent.OrganizationID == "" {
		// If Name wasn't set/valid, we assume it's a project parent resolved via ProjectRef or context.
		// KCC usually resolves project from context if ProjectRef is nil.
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
		if err != nil {
			// If ProjectRef is explicitly nil and context resolution also fails (which ResolveProject handles)
			if obj.Spec.ProjectRef == nil {
				return nil, fmt.Errorf("cannot determine parent: spec.name does not specify parent and spec.projectRef is nil (or project context cannot be resolved)")
			}
			return nil, fmt.Errorf("cannot resolve projectRef: %w", err)
		}
		if projectRef.ProjectID == "" {
			// Should not happen if ResolveProject succeeds without error, but check anyway.
			return nil, fmt.Errorf("resolved projectRef has empty ProjectID")
		}
		desiredParent.ProjectID = projectRef.ProjectID
	}

	// Get desired resource ID from spec or metadata.name
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName() // K8s metadata.name
	}
	// If spec.name specifies a full resource name, extract the ID from there as highest priority for *desired* ID
	if obj.Spec.Name != nil && *obj.Spec.Name != "" {
		_, nameResourceID, nameErr := ParseFeedExternal(*obj.Spec.Name)
		if nameErr == nil && nameResourceID != "" {
			resourceID = nameResourceID
		}
	}

	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID from spec.resourceID, metadata.name, or spec.name")
	}

	// Use approved ExternalRef from Status if available and validate against desired state
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualParent, actualResourceID, err := ParseFeedExternal(externalRef)
		if err != nil {
			// Existing resource has unparsable externalRef, treat as error
			return nil, fmt.Errorf("cannot parse existing status.externalRef %q: %w", externalRef, err)
		}

		// Validate parent hasn't changed
		// Compare normalized string representations
		if actualParent.String() != desiredParent.String() {
			return nil, fmt.Errorf("parent resource changed, expected %q, got %q", actualParent.String(), desiredParent.String())
		}

		// Validate resource ID hasn't changed
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot change resource ID to %q once set to %q (via status.externalRef). Ensure spec.resourceID, metadata.name, or the ID in spec.name (%q) matches the existing resource ID",
				resourceID, actualResourceID, common.ValueOf(obj.Spec.Name))
		}
		// If validation passes, ensure we use the actual ID
		resourceID = actualResourceID
	}

	// Final check: Ensure exactly one parent ID field is set
	parentCount := 0
	if desiredParent.ProjectID != "" {
		parentCount++
	}
	if desiredParent.FolderID != "" {
		parentCount++
	}
	if desiredParent.OrganizationID != "" {
		parentCount++
	}
	if parentCount != 1 {
		return nil, fmt.Errorf("invalid parent configuration: exactly one of ProjectID, FolderID, or OrganizationID must be specified or resolvable (got: %+v)", desiredParent)
	}

	return &FeedIdentity{
		parent: desiredParent,
		id:     resourceID,
	}, nil
}

// ParseFeedExternal parses the external identifier string into its components.
// Format: "projects/{projectID}/feeds/{feedID}"
//
//	or "folders/{folderID}/feeds/{feedID}"
//	or "organizations/{organizationID}/feeds/{feedID}"
func ParseFeedExternal(external string) (parent *FeedParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	// Expect 4 tokens, e.g., projects/{id}/feeds/{feedId}
	if len(tokens) != 4 || tokens[2] != "feeds" {
		return nil, "", fmt.Errorf("format of AssetFeed external=%q was not known (use projects/{projectID}/feeds/{feedID}, folders/{folderID}/feeds/{feedID}, or organizations/{organizationID}/feeds/{feedID})", external)
	}

	parent = &FeedParent{}
	parentType := tokens[0]
	parentVal := tokens[1]
	resourceID = tokens[3]

	switch parentType {
	case "projects":
		parent.ProjectID = parentVal
	case "folders":
		parent.FolderID = parentVal
	case "organizations":
		parent.OrganizationID = parentVal
	default:
		return nil, "", fmt.Errorf("format of AssetFeed external=%q was not known (unknown parent type %q)", external, parentType)
	}

	if parentVal == "" {
		return nil, "", fmt.Errorf("format of AssetFeed external=%q was not known (empty parent ID)", external)
	}
	if resourceID == "" {
		return nil, "", fmt.Errorf("format of AssetFeed external=%q was not known (empty resource ID)", external)
	}

	return parent, resourceID, nil
}
