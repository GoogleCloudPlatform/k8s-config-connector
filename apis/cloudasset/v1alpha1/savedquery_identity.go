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

package v1alpha1

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
	_ identity.IdentityV2 = &SavedQueryIdentity{}
	_ identity.Resource   = &CloudAssetSavedQuery{}
)

var SavedQueryProjectIdentityFormat = gcpurls.Template[SavedQueryIdentity]("cloudasset.googleapis.com", "projects/{project}/savedQueries/{savedQuery}")
var SavedQueryFolderIdentityFormat = gcpurls.Template[SavedQueryIdentity]("cloudasset.googleapis.com", "folders/{folder}/savedQueries/{savedQuery}")
var SavedQueryOrganizationIdentityFormat = gcpurls.Template[SavedQueryIdentity]("cloudasset.googleapis.com", "organizations/{organization}/savedQueries/{savedQuery}")

// +k8s:deepcopy-gen=false
type SavedQueryIdentity struct {
	Project      string
	Folder       string
	Organization string
	SavedQuery   string
}

func (i *SavedQueryIdentity) String() string {
	if i.Project != "" {
		return SavedQueryProjectIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return SavedQueryFolderIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return SavedQueryOrganizationIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *SavedQueryIdentity) Host() string {
	return "cloudasset.googleapis.com"
}

func (i *SavedQueryIdentity) FromExternal(ref string) error {
	if parsed, match, err := SavedQueryProjectIdentityFormat.Parse(ref); err == nil && match {
		*i = *parsed
		return nil
	}
	if parsed, match, err := SavedQueryFolderIdentityFormat.Parse(ref); err == nil && match {
		*i = *parsed
		return nil
	}
	if parsed, match, err := SavedQueryOrganizationIdentityFormat.Parse(ref); err == nil && match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of CloudAssetSavedQuery external=%q was not known (use %s, %s, or %s)", ref, SavedQueryProjectIdentityFormat.CanonicalForm(), SavedQueryFolderIdentityFormat.CanonicalForm(), SavedQueryOrganizationIdentityFormat.CanonicalForm())
}

// ExternalIdentifier implements the IdentityV2 interface.
func (c *CloudAssetSavedQuery) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	id := &SavedQueryIdentity{}

	if c.Spec.Parent.ProjectRef != nil {
		project, err := refs.ResolveProject(ctx, reader, c.GetNamespace(), c.Spec.Parent.ProjectRef)
		if err != nil {
			return nil, err
		}
		id.Project = project.ProjectID
	} else if c.Spec.Parent.FolderRef != nil {
		folder, err := refs.ResolveFolder(ctx, reader, c, c.Spec.Parent.FolderRef)
		if err != nil {
			return nil, err
		}
		id.Folder = folder.FolderID
	} else if c.Spec.Parent.OrganizationRef != nil {
		organization, err := refs.ResolveOrganization(ctx, reader, c, c.Spec.Parent.OrganizationRef)
		if err != nil {
			return nil, err
		}
		id.Organization = organization.OrganizationID
	} else {
		return nil, fmt.Errorf("one of spec.parent.projectRef, spec.parent.folderRef, or spec.parent.organizationRef must be set")
	}

	resourceID := common.ValueOf(c.Spec.ResourceID)
	if resourceID == "" {
		resourceID = c.GetName()
	}
	id.SavedQuery = resourceID

	if externalRef := common.ValueOf(c.Status.ExternalRef); externalRef != "" {
		actualIdentity := &SavedQueryIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if actualIdentity.Project != id.Project {
			return nil, fmt.Errorf("spec.parent.projectRef changed, expect %s, got %s", actualIdentity.Project, id.Project)
		}
		if actualIdentity.Folder != id.Folder {
			return nil, fmt.Errorf("spec.parent.folderRef changed, expect %s, got %s", actualIdentity.Folder, id.Folder)
		}
		if actualIdentity.Organization != id.Organization {
			return nil, fmt.Errorf("spec.parent.organizationRef changed, expect %s, got %s", actualIdentity.Organization, id.Organization)
		}

		if actualIdentity.SavedQuery != id.SavedQuery {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already been assigned to %s",
				id.SavedQuery, actualIdentity.SavedQuery)
		}
	}

	return id, nil
}

// ExternalIdentifier implements the IdentityV2 interface.
func (c *CloudAssetSavedQuery) ExternalIdentifier() *string {
	return c.Status.ExternalRef
}
