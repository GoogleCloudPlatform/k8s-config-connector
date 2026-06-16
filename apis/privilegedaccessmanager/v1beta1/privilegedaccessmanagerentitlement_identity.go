// Copyright 2026 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &PrivilegedAccessManagerEntitlementIdentity{}
	_ identity.Resource   = &PrivilegedAccessManagerEntitlement{}
)

var (
	ProjectEntitlementIdentityFormat      = gcpurls.Template[PrivilegedAccessManagerEntitlementIdentity]("privilegedaccessmanager.googleapis.com", "projects/{project}/locations/{location}/entitlements/{entitlement}")
	FolderEntitlementIdentityFormat       = gcpurls.Template[PrivilegedAccessManagerEntitlementIdentity]("privilegedaccessmanager.googleapis.com", "folders/{folder}/locations/{location}/entitlements/{entitlement}")
	OrganizationEntitlementIdentityFormat = gcpurls.Template[PrivilegedAccessManagerEntitlementIdentity]("privilegedaccessmanager.googleapis.com", "organizations/{organization}/locations/{location}/entitlements/{entitlement}")
)

// +k8s:deepcopy-gen=false
// PrivilegedAccessManagerEntitlementIdentity is the identity of a Google Cloud PrivilegedAccessManagerEntitlement resource.
type PrivilegedAccessManagerEntitlementIdentity struct {
	Project      string
	Folder       string
	Organization string
	Location     string
	Entitlement  string
}

func (i *PrivilegedAccessManagerEntitlementIdentity) String() string {
	if i.Project != "" {
		return ProjectEntitlementIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return FolderEntitlementIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationEntitlementIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *PrivilegedAccessManagerEntitlementIdentity) Container() string {
	if i.Project != "" {
		return "projects/" + i.Project
	}
	if i.Folder != "" {
		return "folders/" + i.Folder
	}
	if i.Organization != "" {
		return "organizations/" + i.Organization
	}
	return ""
}

func (i *PrivilegedAccessManagerEntitlementIdentity) ParentString() string {
	return fmt.Sprintf("%s/locations/%s", i.Container(), i.Location)
}

func (i *PrivilegedAccessManagerEntitlementIdentity) FullyQualifiedName() string {
	return i.String()
}

func (i *PrivilegedAccessManagerEntitlementIdentity) AsExternalRef() *string {
	e := "//privilegedaccessmanager.googleapis.com/" + i.FullyQualifiedName()
	return &e
}

func (i *PrivilegedAccessManagerEntitlementIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectEntitlementIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := FolderEntitlementIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationEntitlementIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of PrivilegedAccessManagerEntitlement external=%q was not known (use projects/{{projectID}}/locations/{{location}}/entitlements/{{entitlementID}})", ref)
}

func (i *PrivilegedAccessManagerEntitlementIdentity) Host() string {
	return "privilegedaccessmanager.googleapis.com"
}

func getIdentityFromPrivilegedAccessManagerEntitlementSpec(ctx context.Context, reader client.Reader, obj *PrivilegedAccessManagerEntitlement) (*PrivilegedAccessManagerEntitlementIdentity, error) {
	// Get user-configured ID
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location := ""
	if obj.Spec.Location != nil {
		location = *obj.Spec.Location
	}

	identity := &PrivilegedAccessManagerEntitlementIdentity{
		Entitlement: resourceID,
		Location:    location,
	}

	// Verify that at most one parent reference is set
	parentCount := 0
	if obj.Spec.ProjectRef != nil {
		parentCount++
	}
	if obj.Spec.FolderRef != nil {
		parentCount++
	}
	if obj.Spec.OrganizationRef != nil {
		parentCount++
	}
	if parentCount > 1 {
		return nil, fmt.Errorf("at most one of spec.projectRef, spec.folderRef and spec.organizationRef can be set")
	}

	// Resolve parent references
	if obj.Spec.ProjectRef != nil {
		projectID, err := refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = projectID
	} else if obj.Spec.FolderRef != nil {
		folderRef := &refs.FolderRef{
			External:  obj.Spec.FolderRef.External,
			Name:      obj.Spec.FolderRef.Name,
			Namespace: obj.Spec.FolderRef.Namespace,
		}
		folder, err := refs.ResolveFolder(ctx, reader, obj, folderRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.folderRef: %w", err)
		}
		identity.Folder = folder.FolderID
	} else if obj.Spec.OrganizationRef != nil {
		orgRef := &refs.OrganizationRef{
			External: obj.Spec.OrganizationRef.External,
		}
		org, err := refs.ResolveOrganization(ctx, reader, obj, orgRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.organizationRef: %w", err)
		}
		identity.Organization = org.OrganizationID
	} else {
		// Fallback to project ID from namespace
		projectID, err := refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = projectID
	}

	return identity, nil
}

func (obj *PrivilegedAccessManagerEntitlement) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromPrivilegedAccessManagerEntitlementSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &PrivilegedAccessManagerEntitlementIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change PrivilegedAccessManagerEntitlement identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
