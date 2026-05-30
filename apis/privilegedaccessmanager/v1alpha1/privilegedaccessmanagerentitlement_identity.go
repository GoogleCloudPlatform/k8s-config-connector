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
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &PrivilegedAccessManagerEntitlementIdentity{}
	_ identity.Resource   = &PrivilegedAccessManagerEntitlement{}
)

var (
	PrivilegedAccessManagerEntitlementProjectIdentityFormat      = gcpurls.Template[PrivilegedAccessManagerEntitlementIdentity]("privilegedaccessmanager.googleapis.com", "projects/{project}/locations/{location}/entitlements/{entitlement}")
	PrivilegedAccessManagerEntitlementFolderIdentityFormat       = gcpurls.Template[PrivilegedAccessManagerEntitlementIdentity]("privilegedaccessmanager.googleapis.com", "folders/{folder}/locations/{location}/entitlements/{entitlement}")
	PrivilegedAccessManagerEntitlementOrganizationIdentityFormat = gcpurls.Template[PrivilegedAccessManagerEntitlementIdentity]("privilegedaccessmanager.googleapis.com", "organizations/{organization}/locations/{location}/entitlements/{entitlement}")
)

// +k8s:deepcopy-gen=false
type PrivilegedAccessManagerEntitlementIdentity struct {
	Project      string
	Folder       string
	Organization string
	Location     string
	Entitlement  string
}

func (i *PrivilegedAccessManagerEntitlementIdentity) String() string {
	if i.Project != "" {
		return PrivilegedAccessManagerEntitlementProjectIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return PrivilegedAccessManagerEntitlementFolderIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return PrivilegedAccessManagerEntitlementOrganizationIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *PrivilegedAccessManagerEntitlementIdentity) FromExternal(ref string) error {
	if parsed, match, err := PrivilegedAccessManagerEntitlementProjectIdentityFormat.Parse(ref); match {
		if err != nil {
			return err
		}
		*i = *parsed
		return nil
	}
	if parsed, match, err := PrivilegedAccessManagerEntitlementFolderIdentityFormat.Parse(ref); match {
		if err != nil {
			return err
		}
		*i = *parsed
		return nil
	}
	if parsed, match, err := PrivilegedAccessManagerEntitlementOrganizationIdentityFormat.Parse(ref); match {
		if err != nil {
			return err
		}
		*i = *parsed
		return nil
	}

	return fmt.Errorf("format of PrivilegedAccessManagerEntitlement external=%q was not known", ref)
}

func (i *PrivilegedAccessManagerEntitlementIdentity) Host() string {
	return PrivilegedAccessManagerEntitlementProjectIdentityFormat.Host()
}

func (i *PrivilegedAccessManagerEntitlementIdentity) Parent() string {
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

func (i *PrivilegedAccessManagerEntitlementIdentity) FullParent() string {
	return i.Parent() + "/locations/" + i.Location
}

func getIdentityFromPrivilegedAccessManagerEntitlementSpec(ctx context.Context, reader client.Reader, obj client.Object) (*PrivilegedAccessManagerEntitlementIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	identity := &PrivilegedAccessManagerEntitlementIdentity{
		Location:    location,
		Entitlement: resourceID,
	}

	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to unstructured: %w", err)
		}
		u = &unstructured.Unstructured{Object: m}
	}

	if _, found, _ := unstructured.NestedMap(u.Object, "spec", "projectRef"); found {
		projectID, err := refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = projectID
	} else if _, found, _ := unstructured.NestedMap(u.Object, "spec", "folderRef"); found {
		folderID, err := refs.ResolveFolderID(ctx, reader, u)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve folder: %w", err)
		}
		identity.Folder = folderID
	} else if _, found, _ := unstructured.NestedMap(u.Object, "spec", "organizationRef"); found {
		organizationID, err := refs.ResolveOrganizationID(ctx, reader, u)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		identity.Organization = organizationID
	} else {
		// Default to project if none set (though one is required by CRD)
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
