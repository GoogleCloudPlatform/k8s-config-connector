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
	_ identity.IdentityV2 = &ModelArmorFloorSettingIdentity{}
	_ identity.Resource   = &ModelArmorFloorSetting{}
)

var (
	ProjectModelArmorFloorSettingIdentityFormat      = gcpurls.Template[ModelArmorFloorSettingIdentity]("modelarmor.googleapis.com", "projects/{project}/locations/{location}/floorSetting")
	FolderModelArmorFloorSettingIdentityFormat       = gcpurls.Template[ModelArmorFloorSettingIdentity]("modelarmor.googleapis.com", "folders/{folder}/locations/{location}/floorSetting")
	OrganizationModelArmorFloorSettingIdentityFormat = gcpurls.Template[ModelArmorFloorSettingIdentity]("modelarmor.googleapis.com", "organizations/{organization}/locations/{location}/floorSetting")
)

// ModelArmorFloorSettingIdentity is the identity of a GCP ModelArmorFloorSetting resource.
// +k8s:deepcopy-gen=false
type ModelArmorFloorSettingIdentity struct {
	Project      string
	Folder       string
	Organization string
	Location     string
}

func (i *ModelArmorFloorSettingIdentity) String() string {
	if i.Project != "" {
		return ProjectModelArmorFloorSettingIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return FolderModelArmorFloorSettingIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationModelArmorFloorSettingIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *ModelArmorFloorSettingIdentity) ParentString() string {
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

func (i *ModelArmorFloorSettingIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectModelArmorFloorSettingIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := FolderModelArmorFloorSettingIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationModelArmorFloorSettingIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ModelArmorFloorSetting external=%q was not known (use %s)", ref, ProjectModelArmorFloorSettingIdentityFormat.CanonicalForm())
}

func (i *ModelArmorFloorSettingIdentity) Host() string {
	return "modelarmor.googleapis.com"
}

func getIdentityFromModelArmorFloorSettingSpec(ctx context.Context, reader client.Reader, obj *ModelArmorFloorSetting) (*ModelArmorFloorSettingIdentity, error) {
	location := obj.Spec.Location
	if location == nil || *location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	identity := &ModelArmorFloorSettingIdentity{
		Location: *location,
	}

	if obj.Spec.ProjectRef != nil {
		projectID, err := refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = projectID
	} else if obj.Spec.FolderRef != nil {
		folder, err := refs.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.folderRef: %w", err)
		}
		identity.Folder = folder.FolderID
	} else if obj.Spec.OrganizationRef != nil {
		org, err := refs.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.organizationRef: %w", err)
		}
		identity.Organization = org.OrganizationID
	} else {
		projectID, err := refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = projectID
	}

	return identity, nil
}

func (obj *ModelArmorFloorSetting) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromModelArmorFloorSettingSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ModelArmorFloorSettingIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ModelArmorFloorSetting identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
