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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &IamAccessPolicyIdentity{}
	_ identity.Resource   = &IamAccessPolicy{}
)

var (
	ProjectAccessPolicyIdentityFormat      = gcpurls.Template[IamAccessPolicyIdentity]("iam.googleapis.com", "projects/{project}/locations/{location}/accessPolicies/{accessPolicy}")
	FolderAccessPolicyIdentityFormat       = gcpurls.Template[IamAccessPolicyIdentity]("iam.googleapis.com", "folders/{folder}/locations/{location}/accessPolicies/{accessPolicy}")
	OrganizationAccessPolicyIdentityFormat = gcpurls.Template[IamAccessPolicyIdentity]("iam.googleapis.com", "organizations/{organization}/locations/{location}/accessPolicies/{accessPolicy}")
)

// +k8s:deepcopy-gen=false
type IamAccessPolicyIdentity struct {
	Project      string
	Folder       string
	Organization string
	Location     string
	AccessPolicy string
}

func (i *IamAccessPolicyIdentity) String() string {
	if i.Project != "" {
		return ProjectAccessPolicyIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return FolderAccessPolicyIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationAccessPolicyIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *IamAccessPolicyIdentity) ParentString() string {
	if i.Project != "" {
		return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
	}
	if i.Folder != "" {
		return fmt.Sprintf("folders/%s/locations/%s", i.Folder, i.Location)
	}
	if i.Organization != "" {
		return fmt.Sprintf("organizations/%s/locations/%s", i.Organization, i.Location)
	}
	return ""
}

func (i *IamAccessPolicyIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectAccessPolicyIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := FolderAccessPolicyIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationAccessPolicyIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of IamAccessPolicy external=%q was not known (use %s, %s, or %s)",
		ref,
		ProjectAccessPolicyIdentityFormat.CanonicalForm(),
		FolderAccessPolicyIdentityFormat.CanonicalForm(),
		OrganizationAccessPolicyIdentityFormat.CanonicalForm(),
	)
}

func (i *IamAccessPolicyIdentity) Host() string {
	return "iam.googleapis.com"
}

func getIdentityFromIamAccessPolicySpec(ctx context.Context, reader client.Reader, obj *IamAccessPolicy) (*IamAccessPolicyIdentity, error) {
	var project, folder, organization string

	if obj.Spec.ProjectRef != nil {
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.projectRef: %w", err)
		}
		project = projectRef.ProjectID
	} else if obj.Spec.FolderRef != nil {
		folderRef, err := refsv1beta1.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.folderRef: %w", err)
		}
		folder = folderRef.FolderID
	} else if obj.Spec.OrganizationRef != nil {
		orgRef, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.organizationRef: %w", err)
		}
		organization = orgRef.OrganizationID
	} else {
		// Fall back to namespace project reference by default if no parent is provided
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), nil)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.projectRef: %w", err)
		}
		project = projectRef.ProjectID
	}

	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	identity := &IamAccessPolicyIdentity{
		Project:      project,
		Folder:       folder,
		Organization: organization,
		Location:     obj.Spec.Location,
		AccessPolicy: resourceID,
	}

	return identity, nil
}

func (obj *IamAccessPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromIamAccessPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
