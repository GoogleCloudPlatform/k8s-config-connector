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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CloudSecurityFrameworkIdentity{}
	_ identity.Resource   = &CloudSecurityFramework{}
)

var (
	ProjectCloudSecurityFrameworkIdentityFormat      = gcpurls.Template[CloudSecurityFrameworkIdentity]("cloudsecuritycompliance.googleapis.com", "projects/{project}/locations/{location}/frameworks/{framework}")
	OrganizationCloudSecurityFrameworkIdentityFormat = gcpurls.Template[CloudSecurityFrameworkIdentity]("cloudsecuritycompliance.googleapis.com", "organizations/{organization}/locations/{location}/frameworks/{framework}")
)

// +k8s:deepcopy-gen=false
type CloudSecurityFrameworkIdentity struct {
	Project      string
	Organization string
	Location     string
	Framework    string
}

func (i *CloudSecurityFrameworkIdentity) String() string {
	if i.Organization != "" {
		return OrganizationCloudSecurityFrameworkIdentityFormat.ToString(*i)
	}
	if i.Project != "" {
		return ProjectCloudSecurityFrameworkIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *CloudSecurityFrameworkIdentity) ParentString() string {
	if i.Organization != "" {
		return fmt.Sprintf("organizations/%s/locations/%s", i.Organization, i.Location)
	}
	if i.Project != "" {
		return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
	}
	return ""
}

func (i *CloudSecurityFrameworkIdentity) ID() string {
	return i.Framework
}

func (i *CloudSecurityFrameworkIdentity) FromExternal(ref string) error {
	if parsed, match, _ := OrganizationCloudSecurityFrameworkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := ProjectCloudSecurityFrameworkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of CloudSecurityFramework external=%q was not known (use organizations/{{organizationID}}/locations/{{location}}/frameworks/{{frameworkID}} or projects/{{projectID}}/locations/{{location}}/frameworks/{{frameworkID}})", ref)
}

func (i *CloudSecurityFrameworkIdentity) Host() string {
	return "cloudsecuritycompliance.googleapis.com"
}

func GetIdentityFromCloudSecurityFrameworkSpec(ctx context.Context, reader client.Reader, obj client.Object) (*CloudSecurityFrameworkIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	// Resolve Project or Organization
	var projectID string
	var organizationID string
	if u, ok := obj.(*unstructured.Unstructured); ok {
		if _, found, _ := unstructured.NestedMap(u.Object, "spec", "projectRef"); found {
			projID, err := refs.ResolveProjectID(ctx, reader, u)
			if err != nil {
				return nil, fmt.Errorf("cannot resolve project: %w", err)
			}
			projectID = projID
		} else if _, found, _ := unstructured.NestedMap(u.Object, "spec", "organizationRef"); found {
			orgID, err := refs.ResolveOrganizationID(ctx, reader, u)
			if err != nil {
				return nil, fmt.Errorf("cannot resolve organization: %w", err)
			}
			organizationID = orgID
		} else {
			return nil, fmt.Errorf("exactly one parent field (projectRef or organizationRef) must be set")
		}
	} else if typed, ok := obj.(*CloudSecurityFramework); ok {
		if typed.Spec.ProjectRef != nil {
			proj, err := refs.ResolveProject(ctx, reader, typed.GetNamespace(), typed.Spec.ProjectRef)
			if err != nil {
				return nil, fmt.Errorf("cannot resolve project: %w", err)
			}
			if proj != nil {
				projectID = proj.ProjectID
			}
		} else if typed.Spec.OrganizationRef != nil {
			org, err := refs.ResolveOrganization(ctx, reader, typed, typed.Spec.OrganizationRef)
			if err != nil {
				return nil, fmt.Errorf("cannot resolve organization: %w", err)
			}
			if org != nil {
				organizationID = org.OrganizationID
			}
		} else {
			return nil, fmt.Errorf("exactly one parent field (projectRef or organizationRef) must be set")
		}
	}

	if projectID == "" && organizationID == "" {
		return nil, fmt.Errorf("exactly one parent field (projectRef or organizationRef) must be set")
	}

	identity := &CloudSecurityFrameworkIdentity{
		Project:      projectID,
		Organization: organizationID,
		Location:     location,
		Framework:    resourceID,
	}
	return identity, nil
}

func (obj *CloudSecurityFramework) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := GetIdentityFromCloudSecurityFrameworkSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CloudSecurityFrameworkIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudSecurityFramework identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *CloudSecurityFramework) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
