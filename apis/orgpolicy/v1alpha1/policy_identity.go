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

// PolicyIdentity defines the resource reference to OrgPolicyPolicy, which "External" field
// holds the GCP identifier for the KRM object.
type PolicyIdentity struct {
	parent *PolicyParent
	id     string
}

func (i *PolicyIdentity) String() string {
	return i.parent.String() + "/policies/" + i.id
}

func (i *PolicyIdentity) ID() string {
	return i.id
}

func (i *PolicyIdentity) Parent() *PolicyParent {
	return i.parent
}

type PolicyParent struct {
	ProjectID      string
	FolderID       string
	OrganizationID string
}

// Format: "projects/{projectID}" or "folders/{folderID}" or "organizations/{organizationID}"
func (p *PolicyParent) String() string {
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

// New builds a PolicyIdentity from the Config Connector Policy object.
func NewPolicyIdentity(ctx context.Context, reader client.Reader, obj *OrgPolicyPolicy) (*PolicyIdentity, error) {
	var projectID, folderID, organizationID string
	// Get Parent
	if obj.Spec.ProjectRef != nil {
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
		if err != nil {
			return nil, err
		}
		projectID = projectRef.ProjectID
		if projectID == "" {
			return nil, fmt.Errorf("cannot resolve project")
		}
	} else if obj.Spec.FolderRef != nil {
		folderRef, err := refsv1beta1.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
		if err != nil {
			return nil, err
		}
		folderID = folderRef.FolderID
		if folderID == "" {
			return nil, fmt.Errorf("cannot resolve folder")
		}
	} else if obj.Spec.OrganizationRef != nil {
		organizationRef, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
		if err != nil {
			return nil, err
		}
		organizationID = organizationRef.OrganizationID
		if organizationID == "" {
			return nil, fmt.Errorf("cannot resolve organization")
		}
	} else {
		return nil, fmt.Errorf("one of spec.parent.projectRef, spec.parent.folderRef, or spec.parent.organizationRef must be set")
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParsePolicyExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.FolderID != folderID {
			return nil, fmt.Errorf("spec.folderRef changed, expect %s, got %s", actualParent.FolderID, folderID)
		}
		if actualParent.OrganizationID != organizationID {
			return nil, fmt.Errorf("spec.organizationRef changed, expect %s, got %s", actualParent.OrganizationID, organizationID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &PolicyIdentity{
		parent: &PolicyParent{
			ProjectID:      projectID,
			FolderID:       folderID,
			OrganizationID: organizationID,
		},
		id: resourceID,
	}, nil
}

func ParsePolicyExternal(external string) (parent *PolicyParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")

	if len(tokens) != 4 || tokens[2] != "policies" {
		return nil, "", fmt.Errorf("format of OrgPolicyPolicy external=%q was not known (use projects|folders|organizations/{{ID}}/policies/{{policyID}})", external)
	}
	parentType := tokens[0]
	parentVal := tokens[1]
	parent = &PolicyParent{}
	switch parentType {
	case "projects":
		parent.ProjectID = parentVal
	case "folders":
		parent.FolderID = parentVal
	case "organizations":
		parent.OrganizationID = parentVal
	default:
		return nil, "", fmt.Errorf("format of OrgPolicyPolicy external=%q was not known (use projects|folders|organizations/{{ID}}/policies/{{policyID}})", external)
	}

	resourceID = tokens[3]
	return parent, resourceID, nil
}
