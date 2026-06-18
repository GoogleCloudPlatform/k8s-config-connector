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
	_ identity.IdentityV2 = &CloudSecurityComplianceFrameworkIdentity{}
	_ identity.Resource   = &CloudSecurityComplianceFramework{}
)

var (
	ProjectCloudSecurityComplianceFrameworkIdentityFormat      = gcpurls.Template[CloudSecurityComplianceFrameworkIdentity]("cloudsecuritycompliance.googleapis.com", "projects/{project}/locations/{location}/frameworks/{framework}")
	OrganizationCloudSecurityComplianceFrameworkIdentityFormat = gcpurls.Template[CloudSecurityComplianceFrameworkIdentity]("cloudsecuritycompliance.googleapis.com", "organizations/{organization}/locations/{location}/frameworks/{framework}")
	CloudSecurityComplianceFrameworkIdentityFormat             = OrganizationCloudSecurityComplianceFrameworkIdentityFormat
)

// +k8s:deepcopy-gen=false
type CloudSecurityComplianceFrameworkIdentity struct {
	Project      string
	Organization string
	Location     string
	Framework    string
}

func (i *CloudSecurityComplianceFrameworkIdentity) String() string {
	if i.Organization != "" {
		return OrganizationCloudSecurityComplianceFrameworkIdentityFormat.ToString(*i)
	}
	if i.Project != "" {
		return ProjectCloudSecurityComplianceFrameworkIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *CloudSecurityComplianceFrameworkIdentity) ParentString() string {
	if i.Organization != "" {
		return fmt.Sprintf("organizations/%s/locations/%s", i.Organization, i.Location)
	}
	if i.Project != "" {
		return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
	}
	return ""
}

func (i *CloudSecurityComplianceFrameworkIdentity) ID() string {
	return i.Framework
}

func (i *CloudSecurityComplianceFrameworkIdentity) FromExternal(ref string) error {
	if parsed, match, _ := OrganizationCloudSecurityComplianceFrameworkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := ProjectCloudSecurityComplianceFrameworkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of CloudSecurityComplianceFramework external=%q was not known (use organizations/{{organizationID}}/locations/{{location}}/frameworks/{{frameworkID}} or projects/{{projectID}}/locations/{{location}}/frameworks/{{frameworkID}})", ref)
}

func (i *CloudSecurityComplianceFrameworkIdentity) Host() string {
	return "cloudsecuritycompliance.googleapis.com"
}

func getIdentityFromCloudSecurityComplianceFrameworkSpec(ctx context.Context, reader client.Reader, obj client.Object) (*CloudSecurityComplianceFrameworkIdentity, error) {
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
	} else if typed, ok := obj.(*CloudSecurityComplianceFramework); ok {
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

	identity := &CloudSecurityComplianceFrameworkIdentity{
		Project:      projectID,
		Organization: organizationID,
		Location:     location,
		Framework:    resourceID,
	}
	return identity, nil
}

func (obj *CloudSecurityComplianceFramework) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudSecurityComplianceFrameworkSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CloudSecurityComplianceFrameworkIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudSecurityComplianceFramework identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *CloudSecurityComplianceFramework) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
