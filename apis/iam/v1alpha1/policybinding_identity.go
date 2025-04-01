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

// PolicyBindingIdentity defines the resource reference to IAMPolicyBinding, which "External" field
// holds the GCP identifier for the KRM object.
type PolicyBindingIdentity struct {
	parent *PolicyBindingParent
	id     string
}

func (i *PolicyBindingIdentity) String() string {
	return i.parent.String() + "/policyBindings/" + i.id
}

func (i *PolicyBindingIdentity) ID() string {
	return i.id
}

func (i *PolicyBindingIdentity) Parent() *PolicyBindingParent {
	return i.parent
}

// PolicyBindingParent holds the parent information for a PolicyBinding.
// Only one of ProjectID, FolderID, or OrganizationID should be set.
type PolicyBindingParent struct {
	ProjectID      string
	FolderID       string
	OrganizationID string
	Location       string
}

func (p *PolicyBindingParent) String() string {
	var parentPrefix string
	if p.ProjectID != "" {
		parentPrefix = "projects/" + p.ProjectID
	} else if p.FolderID != "" {
		parentPrefix = "folders/" + p.FolderID
	} else if p.OrganizationID != "" {
		parentPrefix = "organizations/" + p.OrganizationID
	}
	// Should ideally not happen if constructed correctly, but return empty if no parent ID is set.
	if parentPrefix == "" {
		return ""
	}
	return parentPrefix + "/locations/" + p.Location
}

// New builds a PolicyBindingIdentity from the Config Connector PolicyBinding object.
func NewPolicyBindingIdentity(ctx context.Context, reader client.Reader, obj *IAMPolicyBinding) (*PolicyBindingIdentity, error) {
	parent := &PolicyBindingParent{
		Location: obj.Spec.Location,
	}

	// Determine and resolve the parent reference
	parentCount := 0
	if obj.Spec.ProjectRef != nil {
		parentCount++
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
		if err != nil {
			return nil, fmt.Errorf("resolving projectRef: %w", err)
		}
		parent.ProjectID = projectRef.ProjectID
		if parent.ProjectID == "" {
			return nil, fmt.Errorf("projectRef resolved to an empty project ID")
		}
	}
	if obj.Spec.FolderRef != nil {
		parentCount++
		folderRef, err := refsv1beta1.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
		if err != nil {
			return nil, fmt.Errorf("resolving folderRef: %w", err)
		}
		parent.FolderID = folderRef.FolderID // Note: FolderID field name
		if parent.FolderID == "" {
			return nil, fmt.Errorf("folderRef resolved to an empty folder ID")
		}
	}
	if obj.Spec.OrganizationRef != nil {
		parentCount++
		orgRef, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
		if err != nil {
			return nil, fmt.Errorf("resolving organizationRef: %w", err)
		}
		parent.OrganizationID = orgRef.OrganizationID // Note: OrganizationID field name
		if parent.OrganizationID == "" {
			return nil, fmt.Errorf("organizationRef resolved to an empty organization ID")
		}
	}

	if parentCount == 0 {
		return nil, fmt.Errorf("no parent reference (projectRef, folderRef, or organizationRef) specified in spec")
	}
	if parentCount > 1 {
		return nil, fmt.Errorf("multiple parent references specified in spec; only one of projectRef, folderRef, or organizationRef can be set")
	}
	if parent.Location == "" {
		return nil, fmt.Errorf("spec.location must be specified")
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID (metadata.name or spec.resourceID)")
	}

	// Use approved External if available and validate against desired state
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualParent, actualResourceID, err := ParsePolicyBindingExternal(externalRef)
		if err != nil {
			return nil, fmt.Errorf("parsing existing externalRef %q: %w", externalRef, err)
		}

		// Validate parsed parent against desired parent
		if actualParent.ProjectID != parent.ProjectID {
			return nil, fmt.Errorf("parent reference changed: external ref has project ID %q, but spec resolves to %q", actualParent.ProjectID, parent.ProjectID)
		}
		if actualParent.FolderID != parent.FolderID {
			return nil, fmt.Errorf("parent reference changed: external ref has folder ID %q, but spec resolves to %q", actualParent.FolderID, parent.FolderID)
		}
		if actualParent.OrganizationID != parent.OrganizationID {
			return nil, fmt.Errorf("parent reference changed: external ref has organization ID %q, but spec resolves to %q", actualParent.OrganizationID, parent.OrganizationID)
		}
		if actualParent.Location != parent.Location {
			return nil, fmt.Errorf("spec.location changed: external ref has location %q, got %q", actualParent.Location, parent.Location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot change resource ID: external ref has ID %q, but desired ID is %q (from metadata.name or spec.resourceID)",
				actualResourceID, resourceID)
		}
	}

	return &PolicyBindingIdentity{
		parent: parent,
		id:     resourceID,
	}, nil
}

func ParsePolicyBindingExternal(external string) (parent *PolicyBindingParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 {
		return nil, "", fmt.Errorf("invalid format for IAMPolicyBinding external reference %q: expected 6 segments separated by '/'", external)
	}

	if tokens[2] != "locations" {
		return nil, "", fmt.Errorf("invalid format for IAMPolicyBinding external reference %q: segment 3 should be 'locations', got %q", external, tokens[2])
	}
	if tokens[4] != "policyBindings" {
		return nil, "", fmt.Errorf("invalid format for IAMPolicyBinding external reference %q: segment 5 should be 'policyBindings', got %q", external, tokens[4])
	}

	parent = &PolicyBindingParent{
		Location: tokens[3],
	}
	resourceID = tokens[5]

	switch tokens[0] {
	case "projects":
		parent.ProjectID = tokens[1]
	case "folders":
		parent.FolderID = tokens[1]
	case "organizations":
		parent.OrganizationID = tokens[1]
	default:
		return nil, "", fmt.Errorf("invalid format for IAMPolicyBinding external reference %q: segment 1 should be 'projects', 'folders', or 'organizations', got %q", external, tokens[0])
	}

	// Validate IDs are not empty
	if (parent.ProjectID == "" && parent.FolderID == "" && parent.OrganizationID == "") || parent.Location == "" || resourceID == "" {
		return nil, "", fmt.Errorf("invalid format for IAMPolicyBinding external reference %q: contains empty segments", external)
	}

	return parent, resourceID, nil
}
